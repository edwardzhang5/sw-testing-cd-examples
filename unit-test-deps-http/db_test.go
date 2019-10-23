package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var session Session
var _ = Describe("Db", func() {
	BeforeEach(func() {
		session = NewMockSession()
	})
	It("should properly setup a database session", func() {
		documents, _ := session.DB("test").C("other_test").GetDBEntries()
		Expect(len(documents)).Should(BeNumerically(">", 0))
	})

	It("should retreive mocked database entries", func() {
		var correctData bool
    documents, _ := session.DB("test").C("other_test").GetDBEntries()
		for _, m := range documents {
      mapData := m.(map[string]interface{})
			if val, ok := mapData["testdata1"]; ok {
				if val == "software testing" {
					correctData = true
				}
			}
		}
		Expect(correctData).Should(BeTrue())
	})
})
