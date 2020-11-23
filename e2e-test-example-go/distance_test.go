package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Distance Function", func() {
	Context("given 2 points, return the distance", func() {

		// Not required for Go - only interpreted / dynamic languages
		It("should return a float", func() {
			Expect(CalculateDistance(1, 2, 6, 7)).To(BeAssignableToTypeOf(0.0))
		})

		It("should square two numbers", func() {
			Expect(Squared(2)).To(Equal(4.0))
			Expect(Squared(5)).To(Equal(25.0))
		})

		It("should calculate the correct distance given 4 inputs", func() {
			Expect(CalculateDistance(1, 1, 1, 1)).To(Equal(0.0))
			Expect(fmt.Sprintf("%.2f", CalculateDistance(-3, 0, 0.2, -3))).To(Equal("4.39"))
			Expect(fmt.Sprintf("%.2f", CalculateDistance(-12, 18, 0.99, -19))).To(Equal("36.05"))
		})
	})
})
