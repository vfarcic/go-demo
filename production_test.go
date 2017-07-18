// +build production

package main

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

type ProductionTestSuite struct {
	suite.Suite
	hostIp string
}

func (s *ProductionTestSuite) SetupTest() {
}

// Production

func (s ProductionTestSuite) Test_Hello_ReturnsStatus200() {
	start := time.Now()
	if len(os.Getenv("DURATION")) > 0 {
		max, _ := strconv.ParseFloat(os.Getenv("DURATION"), 64)
		minutes := float64(0)
		failures := 0
		errorMsg := ""
		for time.Since(start).Minutes() < max {
			address := fmt.Sprintf("http://%s/demo/hello", s.hostIp)
			resp, err := http.Get(address)

			if err != nil {
				failures++
				errorMsg = err.Error()
			} else if resp == nil {
				failures++
				errorMsg = "Got no response"
			} else {
				s.Equal(200, resp.StatusCode)
			}
			if time.Since(start).Minutes() > minutes {
				fmt.Printf("%2.0f out of %2.0f minutes passed\n", minutes, max)
				minutes++
			}
			if failures > 1 {
				s.Fail(errorMsg)
			}
		}
	} else {
		address := fmt.Sprintf("http://%s/demo/hello", s.hostIp)
		resp, err := http.Get(address)

		if err != nil {
			s.Fail(err.Error())
		} else if resp == nil {
			s.Fail("Got no response")
		} else {
			s.Equal(200, resp.StatusCode)
		}
	}
}

// Suite

func TestProductionTestSuite(t *testing.T) {
	s := new(ProductionTestSuite)
	s.hostIp = os.Getenv("HOST_IP")
	suite.Run(t, s)
}
