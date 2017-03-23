// +build integration

package main

// $ export HOST_IP=<HOST_IP>
// $ docker-compose -f docker-compose-test.yml down

// Unit tests
// $ docker-compose -f docker-compose-test.yml run --rm unit

// Build
// $ docker-compose build app

// Staging tests
// $ docker-compose -f docker-compose-test.yml up -d staging-dep
// $ docker-compose -f docker-compose-test.yml run --rm staging
// $ docker-compose -f docker-compose-test.yml down

// Push
// $ docker push vfarcic/go-demo

// Production tests
// $ docker-compose -f docker-compose-test.yml up -d staging-dep
// $ docker-compose -f docker-compose-test.yml run --rm production

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

type IntegrationTestSuite struct {
	suite.Suite
	hostIp string
}

func (s *IntegrationTestSuite) SetupTest() {
}

// Integration

func (s IntegrationTestSuite) Test_Hello_ReturnsStatus200() {
	address := fmt.Sprintf("http://%s/demo/hello", s.hostIp)
	resp, err := http.Get(address)

	s.NoError(err)
	s.Equal(200, resp.StatusCode)
}

func (s IntegrationTestSuite) Test_Person_ReturnsStatus200() {
	address := fmt.Sprintf("http://%s/demo/person", s.hostIp)
	resp, err := http.Get(address)

	s.NoError(err)
	s.Equal(200, resp.StatusCode)
}

// Suite

func TestIntegrationTestSuite(t *testing.T) {
	s := new(IntegrationTestSuite)
	s.hostIp = os.Getenv("HOST_IP")
	suite.Run(t, s)
}
