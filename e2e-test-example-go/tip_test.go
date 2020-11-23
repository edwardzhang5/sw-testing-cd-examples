package main_test

import (
	"log"
	"testing"

	. "sw-test-4-cd/unit-test-functions"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		It("should return an error if not a US dollar amount with 2 decimals", func() {
			_, err := SplitTip(3.323, 3)
			Expect(err).Should(HaveOccurred())
			log.Println(err)
		})

		It("should distribute the splits as evenly as possible", func() {
			evenSplit := []float64{1.15, 1.15, 1.16}
			split, err := SplitTip(3.01, 3)
			if err != nil {
				log.Println(err)
			}
			Expect(split).To(Equal(evenSplit))
		})

	})
})

func TestDollarDecimals(t *testing.T) {
	splits, err := SplitTip(3.33, 3)
	if err == nil {
		t.Fatalf("Should error for more than 2 decimals in dollar amount %#v", splits)
	}
}
