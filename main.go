package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var coll *mgo.Collection
var sleep = time.Sleep
var logFatal = log.Fatal
var logPrintf = log.Printf
var httpListenAndServe = http.ListenAndServe
var serviceName = "go-demo"

type Person struct {
	Name string
}

var (
	histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: "http_server",
		Name: "resp_time",
		Help: "Request response time",
	}, []string{
		"service",
		"code",
		"method",
		"path",
	})
)

func main() {
	if len(os.Getenv("SERVICE_NAME")) > 0 {
		serviceName = os.Getenv("SERVICE_NAME")
	}
	setupDb()
	RunServer()
}

func init() {
	prometheus.MustRegister(histogram)
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
	mux := http.NewServeMux()
	mux.HandleFunc("/demo/hello", HelloServer)
	mux.HandleFunc("/demo/person", PersonServer)
	mux.HandleFunc("/demo/random-error", RandomErrorServer)
	mux.Handle("/metrics", prometheusHandler())
	logFatal("ListenAndServe: ", httpListenAndServe(":8080", mux))
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	defer func() { recordMetrics(start, req, http.StatusOK) }()

	logPrintf("%s request to %s\n", req.Method, req.RequestURI)
	delay := req.URL.Query().Get("delay")
	if len(delay) > 0 {
		delayNum, _ := strconv.Atoi(delay)
		sleep(time.Duration(delayNum) * time.Millisecond)
	}
	io.WriteString(w, "hello, world!\n")
}

func RandomErrorServer(w http.ResponseWriter, req *http.Request) {
	code := http.StatusOK
	start := time.Now()
	defer func() { recordMetrics(start, req, code) }()

	logPrintf("%s request to %s\n", req.Method, req.RequestURI)
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	msg := "Everything is still OK\n"
	if n == 0 {
		code = http.StatusInternalServerError
		msg = "ERROR: Something, somewhere, went wrong!\n"
		logPrintf(msg)
	}
	w.WriteHeader(code)
	io.WriteString(w, msg)
}

func PersonServer(w http.ResponseWriter, req *http.Request) {
	code := http.StatusOK
	start := time.Now()
	defer func() { recordMetrics(start, req, code) }()

	logPrintf("%s request to %s\n", req.Method, req.RequestURI)
	msg := "Everything is OK"
	if req.Method == "PUT" {
		name := req.URL.Query().Get("name")
		if _, err := upsertId(name, &Person{
			Name: name,
		}); err != nil {
			code = http.StatusInternalServerError
			msg = err.Error()
		}
	} else {
		var res []Person
		if err := findPeople(&res); err != nil {
			panic(err)
		}
		var names []string
		for _, p := range res {
			names = append(names, p.Name)
		}
		msg = strings.Join(names, "\n")
	}
	w.WriteHeader(code)
	io.WriteString(w, msg)
}

var prometheusHandler = func() http.Handler {
	return prometheus.Handler()
}

var findPeople = func(res *[]Person) error {
	return coll.Find(bson.M{}).All(res)
}

var upsertId = func(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return coll.UpsertId(id, update)
}

func recordMetrics(start time.Time, req *http.Request, code int) {
	duration := time.Since(start)
	histogram.With(
		prometheus.Labels{
			"service": serviceName,
			"code": fmt.Sprintf("%d", code),
			"method": req.Method,
			"path": req.URL.Path,
		},
	).Observe(duration.Seconds())
}
