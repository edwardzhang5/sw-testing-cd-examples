package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var _ = Describe("Retirement", func() {
	Context("given input from user should calcutate if / when retirement goals are met", func() {
		It("should return a specific message string if goals are not met", func() {
			_, message := CalculateRetirement(60, 1000, 100000)
			Expect(message).To(Equal("Sorry, according to our calculations (dead > 100), you'll be dead before you reach your goal"))
		})

		It("should return an int of age when goals are met ", func() {
			age, _ := CalculateRetirement(30, 10000, 10000)
			Expect(age).To(BeAssignableToTypeOf(1))
		})

		It("should determine yearly savings with employer matching", func() {
			Expect(YearlySavings(10, 100000)).To(Equal(20000.00))
		})

		It("should return the correct age when goals reached", func() {
			age, _ := CalculateRetirement(30, 10000, 100000)
			Expect(age).To(Equal(40))
			age, _ = CalculateRetirement(30, 10000, 300000)
			Expect(age).To(Equal(60))
		})

		It("should return 0 for age when goal is not reached", func() {
			age, _ := CalculateRetirement(30, 10000, 3000000)
			Expect(age).To(Equal(0))
		})

	})
})
