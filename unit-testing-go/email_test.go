package main_test

import (
	. "drbyronw/sw-test-assign1_S17"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Email function", func() {
	Context("given an input string, it returns whether an string is a valid", func() {

		// Not required for Go - only interpreted / dynamic languages
		It("should return a bool", func() {
			Expect(ValidateEmail("string")).To(BeAssignableToTypeOf(true))
		})

		It("should return true if valid email", func() {
			Expect(ValidateEmail("williams@cse.msstate.edu")).To(Equal(true))
		})

		It("should return false if invalid email", func() {
			Expect(ValidateEmail("wil/iams@cse.msstate.edu")).To(Equal(false))
		})

		It("should return false tld is longer than 3 letters", func() {
			Expect(ValidateEmail("williams@cse.msstate.eduu")).To(Equal(false))
			Expect(ValidateEmail("williams@cse.msstate.e3u")).To(Equal(false))
			Expect(ValidateEmail("williams@cse.msstate.eu")).To(Equal(true))
		})
	})
})
