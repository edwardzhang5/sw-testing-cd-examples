package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "github.com/sw-testing-cd/unit-test-deps-http"
)
var MSession Session
var MDBHandler DBHandler
func TestUnitTestDepsHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnitTestDepsHttp Suite")
}

var _ = BeforeSuite(func() {
  MSession = NewMockSession()
  MDBHandler.Session = MSession
})

var _ = AfterSuite(func() {
  MSession.Close()
})

