package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var _ = Describe("Retirement", func() {
	Context("given input from user should calcutate if / when retirement goals are met", func() {
		It("should return whether goals are met as a bool", func() {
			ok, _ := CalculateRetirement(30, 10000, 10000)
			Expect(ok).To(BeAssignableToTypeOf(true))
		})

		It("should return an int of age when goals are met ", func() {
			_, age := CalculateRetirement(30, 10000, 10000)
			Expect(age).To(BeAssignableToTypeOf(1))
		})

		It("should determine yearly savings with employer matching", func() {
			Expect(YearlySavings(10, 100000)).To(Equal(20000.00))
		})

		It("should return true if retirement goals are met", func() {
			ok, _ := CalculateRetirement(30, 10000, 10000)
			Expect(ok).To(Equal(true))
		})

		It("should return the correct age when goals reached", func() {
			_, age := CalculateRetirement(30, 10000, 100000)
			Expect(age).To(Equal(40))
			_, age = CalculateRetirement(30, 10000, 300000)
			Expect(age).To(Equal(60))
		})

		It("should return false when goal is not reached", func() {
			met, _ := CalculateRetirement(30, 10000, 3000000)
			Expect(met).To(Equal(false))
			met, _ = CalculateRetirement(99, 10000, 100000)
			Expect(met).To(Equal(false))
		})

		It("should return 0 for age when goal is not reached", func() {
			_, age := CalculateRetirement(30, 10000, 3000000)
			Expect(age).To(Equal(0))
		})

	})
})
