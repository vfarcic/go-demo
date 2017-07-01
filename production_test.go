// +build production

package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

type ProductionTestSuite struct {
	suite.Suite
	hostIp      string
}

func (s *ProductionTestSuite) SetupTest() {
}

// Production

func (s ProductionTestSuite) Test_Hello_ReturnsStatus200() {
	address := fmt.Sprintf("http://%s/demo/hello", s.hostIp)
	resp, err := http.Get(address)

	s.NoError(err)
	s.Equal(200, resp.StatusCode)
}

func (s ProductionTestSuite) Test_Person_ReturnsStatus200() {
	address := fmt.Sprintf("http://%s/demo/person", s.hostIp)
	resp, err := http.Get(address)

	s.NoError(err)
	s.Equal(200, resp.StatusCode)
}

// Suite

func TestProductionTestSuite(t *testing.T) {
	s := new(ProductionTestSuite)
	s.hostIp = os.Getenv("HOST_IP")
	suite.Run(t, s)
}
