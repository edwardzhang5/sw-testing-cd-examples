package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var _ = Describe("Main", func() {
	Context("Setup and test router and application connections", func() {

		It("should properly setup router", func() {
			router := SetupRouter(&MDBHandler)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(200))
		})

		It("should pass the ping test", func() {
			router := SetupRouter(&MDBHandler)
			var ping map[string]string

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/ping", nil)
			router.ServeHTTP(w, req)
			json.NewDecoder(w.Body).Decode(&ping)
			Expect(ping["message"]).To(Equal("pong"))
		})
	})

})
