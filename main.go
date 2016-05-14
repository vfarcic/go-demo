package main

import (
	"net/http"
	"log"
	"io"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
	"time"
	"strconv"
)

var coll *mgo.Collection

type Person struct {
	Name string
}

func init() {
	db := os.Getenv("DB")
	if len(db) == 0 {
		db = "localhost"
	}
	session, err := mgo.Dial(db)
	if err != nil {
		panic(err)
	}
	coll = session.DB("test").C("people")
}

func main() {
	http.HandleFunc("/demo/hello", HelloServer)
	http.HandleFunc("/demo/person", PersonServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s request to %s\n", req.Method, req.RequestURI)
	delay := req.URL.Query().Get("delay")
	if len(delay) > 0 {
		delayNum, _ := strconv.Atoi(delay)
		time.Sleep(time.Duration(delayNum) * time.Millisecond)
	}
	io.WriteString(w, "hello, world!\n")
}

func PersonServer(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s request to %s\n", req.Method, req.RequestURI)
	if req.Method == "PUT" {
		name := req.URL.Query().Get("name")
		if _, err := coll.UpsertId(name, &Person{
			Name: name,
		}); err != nil {
			panic(err)
		}
	} else {
		var res []Person
		if err := coll.Find(bson.M{}).All(&res); err != nil {
			panic(err)
		}
		for _, p := range res {
			io.WriteString(w, fmt.Sprintln(p.Name))
		}
	}
}