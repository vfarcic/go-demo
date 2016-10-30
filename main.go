package main

import (
	"gopkg.in/mgo.v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"strings"
	"math/rand"
)

var coll *mgo.Collection
var sleep = time.Sleep
var logFatal = log.Fatal
var logPrintf = log.Printf
var httpHandleFunc = http.HandleFunc
var httpListenAndServe = http.ListenAndServe

type Person struct {
	Name string
}

// TODO: Test

func main() {
	setupDb()
	RunServer()
}

// TODO: Test

func setupDb() {
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

func RunServer() {
	httpHandleFunc("/demo/hello", HelloServer)
	httpHandleFunc("/demo/person", PersonServer)
	httpHandleFunc("/demo/random-error", RandomErrorServer)
	logFatal("ListenAndServe: ", httpListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	logPrintf("%s request to %s\n", req.Method, req.RequestURI)
	delay := req.URL.Query().Get("delay")
	if len(delay) > 0 {
		delayNum, _ := strconv.Atoi(delay)
		sleep(time.Duration(delayNum) * time.Millisecond)
	}
	io.WriteString(w, "hello, world!\n")
}

func RandomErrorServer(w http.ResponseWriter, req *http.Request) {
	logPrintf("%s request to %s\n", req.Method, req.RequestURI)
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	if n == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "ERROR: Something, somewhere, went wrong!\n"
		logPrintf(msg)
		io.WriteString(w, msg)
	} else {
		io.WriteString(w, "Everything is still OK\n")
	}
}

func PersonServer(w http.ResponseWriter, req *http.Request) {
	logPrintf("%s request to %s\n", req.Method, req.RequestURI)

	if req.Method == "PUT" {
		name := req.URL.Query().Get("name")
		if _, err := upsertId(name, &Person{
			Name: name,
		}); err != nil {
			panic(err)
		}
	} else {
		var res []Person
		if err := findPeople(&res); err != nil {
			panic(err)
		}
		var names []string
		for _, p := range res {
			names = append(names, p.Name)
			io.WriteString(w, fmt.Sprintln(p.Name))
		}
		io.WriteString(w, strings.Join(names, "\n"))
	}
}

var findPeople = func(res *[]Person) error {
	return coll.Find(bson.M{}).All(res)
}

var upsertId = func(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return coll.UpsertId(id, update)
}

