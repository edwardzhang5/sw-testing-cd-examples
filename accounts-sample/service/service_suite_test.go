// +build integration

package service_test

import (
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	_ = godotenv.Load("../.env")
	var err error
	if err != nil {
		t.Fatalf("[TestService] : unable to generate test web app: %v", err)
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}
