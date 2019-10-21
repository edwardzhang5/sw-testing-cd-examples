package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUnitTestDepsHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnitTestDepsHttp Suite")
}
