// +build integration

package service_test

import (
	"github.com/drbyronw/accounts/api"
	"github.com/drbyronw/accounts/db"
	"github.com/drbyronw/accounts/models"
	"github.com/drbyronw/accounts/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

func NewTestWebApp() (*api.WebApp, error) {
	var wa api.WebApp
	var err error

	wa.DB, err = db.NewFSRepoClient()
	if err != nil {
		return nil, err
	}
	wa.Accounts = service.NewAccountsService("int_test_accounts", wa.DB)
	return &wa, err
}

var _ = Describe("Service", func() {

	Context("Account service", func() {
		wa, err := NewTestWebApp()
		as := wa.Accounts

		Context("When we need an accounts service", func() {
			It("Should setup an accounts service", func() {
				Expect(err).To(BeNil())
			})
		})
		Context("When we want to create an account", func() {
			It("Should allow us to create a new account in the DB", func() {
				acct := models.Account{
					ID:       "account1",
					ClientID: "client1",
					Type:     "checking",
					Balance:  100.00,
					IsActive: true,
				}

				accErr := as.Create(acct)
				Expect(accErr).To(BeNil())

				acct2 := models.Account{
					ID:       "account2",
					ClientID: "client1",
					Type:     "savings",
					Balance:  1000.00,
					IsActive: true,
				}

				accErr2 := as.Create(acct2)
				Expect(accErr2).To(BeNil())
			})
		})

		Context("When need to find an account", func() {
			It("Should allow us to get an account by ID", func() {
				_, accErr := as.Find("account1")
				Expect(accErr).To(BeNil())
			})
			It("Should error if the account is not found", func() {
				_, accErr := as.Find("xxxx")
				Expect(accErr).Should(HaveOccurred())
			})
		})
		Context("When need to deposit funds", func() {
			It("Should allow do accurately deposit a dollar amount", func() {
				r, accErr := as.Deposit("account1", 100.00)
				Expect(accErr).To(BeNil())
				Expect(r.Equal(decimal.NewFromFloat(200.00))).To(BeTrue())
			})
		})
		Context("When need to transfer funds", func() {
			It("should error if there are insufficient funds", func() {
				errTrue := as.Transfer("account1", "account2", 10000)
				Expect(errTrue).Should(HaveOccurred())

			})
			It("Should allow do accurately Transfer between client acconts", func() {
				errNo := as.Transfer("account2", "account1", 900.00001)
				Expect(errNo).To(BeNil())

				a, accErr := as.Find("account1")
				Expect(accErr).To(BeNil())
				Expect(a.Balance).To(Equal(1100.00))
			})
		})
	})
})
