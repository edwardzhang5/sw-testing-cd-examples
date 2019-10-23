package main_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var _ = Describe("BMI Function", func() {
	Context("Given height, feet, and pounds, it returns the BMI and category", func() {
		It("should convert feet and inches to total number of inches", func() {
			Expect(InchesHeight(6, 0)).To(Equal(72.0))
		})

		It("should convert height in inches to meters", func() {
			Expect(MetersHeightFromInches(72)).To(Equal(1.8))
		})

		It("should return the correct BMI category for any value", func() {
			Expect(BMICategory(18)).To(Equal(Underweight))
			Expect(BMICategory(24)).To(Equal(Normal))
			Expect(BMICategory(25.0001)).To(Equal(Overweight))
			Expect(BMICategory(30.001)).To(Equal(Obese))
			Expect(BMICategory(30)).To(Equal(Obese))
			Expect(BMICategory(29.9999)).To(Equal(Overweight))
		})
	})

	Context("given valid unacceptable input", func() {
		It("should fail given negative values for weight and height", func() {
			_, err := CalculateBMI(-1, 5, 5)
			Expect(err).Should(HaveOccurred())
			_, err = CalculateBMI(5, -5, 100)
			Expect(err).Should(HaveOccurred())
			_, err = CalculateBMI(5, 5, -500)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("BMIHandler should support access via http GET requests", func() {
		It("should calculate the BMI given the correct url string", func() {
			router := SetupRouter(&MDBHandler)
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/bmi/6/0/170", nil)
			if err != nil {
				Fail("could not create request: " + err.Error())
			}

			var results map[string]interface{}
			router.ServeHTTP(w, req)
			json.NewDecoder(w.Body).Decode(&results)

			fmt.Printf("RESULTS, %+v BODY: %+v\n", results, w.Body)
			bmi, _ := strconv.ParseFloat(fmt.Sprintf("%v", results["bmi"]), 64)
			fmt.Println("WCODE------", w.Code)

			Expect(bmi).Should(BeNumerically("<", 24.99))
			Expect(bmi).Should(BeNumerically(">", 18.5))

		})
		It("should return StatusOK (200) for request to correct endpoint", func() {
			router := SetupRouter(&MDBHandler)
			w := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/bmi/6/0/170", nil)
			if err != nil {
				Fail("could not create request: " + err.Error())
			}
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})

// var _ = Describe("BMIInterface Check", func() {
// 	It("should error if no user input", func() {
// 		err := BMIInterface()
// 		Expect(err).Should(HaveOccurred())
// 	})
// })
