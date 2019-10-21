package main_test

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sw-testing-cd/unit-test-deps-http"
)

var _ = Describe("Split the Tip", func() {
	Context("given a total and number of guests it evenly distributes the tip", func() {
		It("should return a single item in list for 1 guest", func() {
			singleSplit := []float64{1.15}
			split, _ := SplitTip(1.00, 1)
			Expect(split).To(Equal(singleSplit))
		})

		It("should split tab with no remainder evenly", func() {
			evenSplit := []float64{1.15, 1.15, 1.15}
			split, _ := SplitTip(3, 3)
			Expect(split).To(Equal(evenSplit))
		})

		It("should distribute the splits as evenly as possible", func() {
			evenSplit := []float64{1.15, 1.16, 1.15}
			split, err := SplitTip(3.01, 3)
			if err != nil {
				log.Println(err)
			}
			Expect(split).To(Equal(evenSplit))
		})

	})
})

func TestDollarDecimals(t *testing.T) {
	splits, err := SplitTip(3.3, 3)
	if err == nil {
		t.Fatalf("Should error for more than 2 decimals in dollar amount %#v", splits)
	}
}
