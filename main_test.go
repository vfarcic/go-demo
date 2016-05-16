package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"reflect"
)

// Setup

type MainTestSuite struct {
	suite.Suite
}

func (s *MainTestSuite) SetupTest() {
}

// RunServer

func (s *MainTestSuite) Test_RunServer_InvokesHandleFuncWithHelloServer() {
	httpHandleFuncOrig := httpHandleFunc
	defer func() { httpHandleFunc = httpHandleFuncOrig }()
	wasCalled := false
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		if pattern == "/demo/hello" && reflect.ValueOf(handler) == reflect.ValueOf(HelloServer) {
			wasCalled = true
		}
	}

	RunServer()

	s.True(wasCalled)
}

func (s *MainTestSuite) Test_RunServer_InvokesHandleFuncWithPersonServer() {
	httpHandleFuncOrig := httpHandleFunc
	defer func() { httpHandleFunc = httpHandleFuncOrig }()
	wasCalled := false
	httpHandleFunc = func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		if pattern == "/demo/person" && reflect.ValueOf(handler) == reflect.ValueOf(PersonServer) {
			wasCalled = true
		}
	}

	RunServer()

	s.True(wasCalled)
}

func (s *MainTestSuite) Test_RunServer_InvokesListenAndServe() {
	actual := ""
	httpListenAndServe = func(addr string, handler http.Handler) error {
		actual = addr
		return nil
	}

	RunServer()

	s.Equal(":8080", actual)
}

// HelloServer

func (s *MainTestSuite) Test_HelloServer_WritesHelloWorld() {
	req, _ := http.NewRequest("GET", "/demo/hello", nil)
	w := getResponseWriterMock()

	HelloServer(w, req)

	w.AssertCalled(s.T(), "Write", []byte("hello, world!\n"))
}

func (s *MainTestSuite) Test_HelloServer_Waits_WhenDelayIsPresent() {
	sleepOrig := sleep
	defer func() { sleep = sleepOrig }()
	var actual time.Duration
	sleep = func(d time.Duration) {
		actual = d
	}
	req, _ := http.NewRequest("GET", "/demo/hello?delay=10", nil)
	w := getResponseWriterMock()

	HelloServer(w, req)

	s.Equal(10*time.Millisecond, actual)
}

// PersonServer

func (s *MainTestSuite) Test_PersonServer_InvokesUpsertId_WhenPutPerson() {
	name := "Viktor"
	upsertIdOrig := upsertId
	defer func() { upsertId = upsertIdOrig }()
	var actualId interface{}
	var actualUpdate interface{}
	upsertId = func(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
		actualId = id
		actualUpdate = update
		return nil, nil
	}
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/demo/person?name=%s", name), nil)
	w := getResponseWriterMock()

	PersonServer(w, req)

	s.Equal(name, actualId)
	s.Equal(&Person{Name: name}, actualUpdate)
}

func (s *MainTestSuite) Test_PersonServer_Panics_WhenUpsertIdReturnsError() {
	upsertIdOrig := upsertId
	defer func() { upsertId = upsertIdOrig }()
	upsertId = func(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
		return nil, fmt.Errorf("This is an error")
	}
	req, _ := http.NewRequest("PUT", "/demo/person?name=Viktor", nil)
	w := getResponseWriterMock()

	s.Panics(func() { PersonServer(w, req) })
}

func (s *MainTestSuite) Test_PersonServer_WritesPeople() {
	findPeopleOrig := findPeople
	people := []Person{
		Person{Name: "Viktor"},
		Person{Name: "Sara"},
	}
	defer func() { findPeople = findPeopleOrig }()
	findPeople = func(res *[]Person) error {
		*res = people
		return nil
	}
	req, _ := http.NewRequest("GET", "/demo/person", nil)
	w := getResponseWriterMock()

	PersonServer(w, req)

	w.AssertCalled(s.T(), "Write", []byte("Viktor\nSara"))
}

func (s *MainTestSuite) Test_PersonServer_Panics_WhenFindReturnsError() {
	findPeopleOrig := findPeople
	defer func() { findPeople = findPeopleOrig }()
	findPeople = func(res *[]Person) error {
		return fmt.Errorf("This is an error")
	}
	req, _ := http.NewRequest("GET", "/demo/person", nil)
	w := getResponseWriterMock()

	s.Panics(func() { PersonServer(w, req) })
}

// Suite

func TestMainSuite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HelloServer))
	//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//		actualPath := r.URL.Path
	//		if r.Method == "PUT" {
	//			defer r.Body.Close()
	//			body, _ := ioutil.ReadAll(r.Body)
	//			switch actualPath {
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/color", s.ServiceName):
	//				s.ConsulRequestBody.ServiceColor = string(body)
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/path", s.ServiceName):
	//				s.ConsulRequestBody.ServicePath = strings.Split(string(body), ",")
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/domain", s.ServiceName):
	//				s.ConsulRequestBody.ServiceDomain = string(body)
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/pathtype", s.ServiceName):
	//				s.ConsulRequestBody.PathType = string(body)
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/skipcheck", s.ServiceName):
	//				v, _ := strconv.ParseBool(string(body))
	//				s.ConsulRequestBody.SkipCheck = v
	//			}
	//		} else if r.Method == "GET" {
	//			switch actualPath {
	//			case "/v1/catalog/services":
	//				w.WriteHeader(http.StatusOK)
	//				w.Header().Set("Content-Type", "application/json")
	//				data := map[string][]string{"service1": []string{}, "service2": []string{}, s.ServiceName: []string{}}
	//				js, _ := json.Marshal(data)
	//				w.Write(js)
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/%s", s.ServiceName, PATH_KEY):
	//				if r.URL.RawQuery == "raw" {
	//					w.WriteHeader(http.StatusOK)
	//					w.Write([]byte(strings.Join(s.ServicePath, ",")))
	//				}
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/%s", s.ServiceName, COLOR_KEY):
	//				if r.URL.RawQuery == "raw" {
	//					w.WriteHeader(http.StatusOK)
	//					w.Write([]byte("orange"))
	//				}
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/%s", s.ServiceName, DOMAIN_KEY):
	//				if r.URL.RawQuery == "raw" {
	//					w.WriteHeader(http.StatusOK)
	//					w.Write([]byte(s.ServiceDomain))
	//				}
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/%s", s.ServiceName, PATH_TYPE_KEY):
	//				if r.URL.RawQuery == "raw" {
	//					w.WriteHeader(http.StatusOK)
	//					w.Write([]byte(s.PathType))
	//				}
	//			case fmt.Sprintf("/v1/kv/docker-flow/%s/%s", s.ServiceName, SKIP_CHECK_KEY):
	//				if r.URL.RawQuery == "raw" {
	//					w.WriteHeader(http.StatusOK)
	//					w.Write([]byte(fmt.Sprintf("%t", s.SkipCheck)))
	//				}
	//			default:
	//				w.WriteHeader(http.StatusNotFound)
	//			}
	//		}
	//	}))
	defer server.Close()
	logFatalOrig := logFatal
	defer func() { logFatal = logFatalOrig }()
	logFatal = func(v ...interface{}) {}
	logPrintfOrig := logPrintf
	defer func() { logPrintf = logPrintfOrig }()
	logPrintf = func(format string, v ...interface{}) {}
	httpListenAndServeOrig := httpListenAndServe
	defer func() { httpListenAndServe = httpListenAndServeOrig }()
	httpListenAndServe = func(addr string, handler http.Handler) error { return nil }
	suite.Run(t, new(MainTestSuite))
}

type ResponseWriterMock struct {
	mock.Mock
}

func (m *ResponseWriterMock) Header() http.Header {
	m.Called()
	return make(map[string][]string)
}

func (m *ResponseWriterMock) Write(data []byte) (int, error) {
	params := m.Called(data)
	return params.Int(0), params.Error(1)
}

func (m *ResponseWriterMock) WriteHeader(header int) {
	m.Called(header)
}

func getResponseWriterMock() *ResponseWriterMock {
	mockObj := new(ResponseWriterMock)
	mockObj.On("Header").Return(nil)
	mockObj.On("Write", mock.Anything).Return(0, nil)
	mockObj.On("WriteHeader", mock.Anything)
	return mockObj
}
