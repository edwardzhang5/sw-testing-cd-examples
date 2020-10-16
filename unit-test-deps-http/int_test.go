// +build integration

package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var _ = Describe("Db", func() {
	var session Session
	var dbHandler DBHandler
	var router *gin.Engine
	BeforeEach(func() {
		session = NewSession()
		dbHandler.Session = session
		router = SetupRouter(&dbHandler)
	})
	AfterEach(func() {
		session.Close()
	})
	It("should properly setup a database session", func() {
		session.DB("test").C("other_test").Insert(bson.M{"test1": "data"})
		documents, _ := session.DB("test").C("other_test").GetDBEntries()
		Expect(len(documents)).Should(BeNumerically(">", 0))
	})

	It("", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/bmi/6/0/170", nil)
		if err != nil {
			Fail("could not create request: " + err.Error())
		}

		var results map[string]interface{}
		router.ServeHTTP(w, req)
		json.NewDecoder(w.Body).Decode(&results)
		documents, _ := session.DB("test").C("other_test").GetDBEntries()

		Expect(len(documents)).Should(BeNumerically(">", 0))

		// var bmiData BMI

	})

})
