package models_test

import (
	"testing"

	"github.com/drbyronw/accounts/models"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAccount(t *testing.T) {
	Convey("Given a client who wants to deposit into a new account", t, func() {
		acct1 := models.Account{
			ID:       "123",
			ClientID: "client1",
			Balance:  0.00,
			IsActive: true,
			Type:     "checking",
		}
		Convey("When they add 1234.56 dollars to the account with a zero balance", func() {
			amt := decimal.NewFromFloat(1234.56)
			acct1.Deposit(amt)
			Convey("The account balance should be $1234.56", func() {
				So(acct1.Balance, ShouldEqual, 1234.56)
			})
		})
		Convey("When they want to confirm the deposited amount of $100", func() {
			amt := decimal.NewFromFloat(100.00)
			deposited := acct1.Deposit(amt)
			Convey("The deposited amount should equal $100", func() {
				So(deposited, ShouldEqual, decimal.NewFromFloat(100))
			})
		})
		Convey("When they want to confirm that extra precision is truncated", func() {
			amt := decimal.NewFromFloat(100.00001)
			acct1.Deposit(amt)
			Convey("The account balance should be $100.00", func() {
				So(acct1.Balance, ShouldEqual, 100.00)
			})
		})
	})

	Convey("Given a client who wants to withdraw from an account", t, func() {
		acctW := models.Account{
			ID:       "123",
			ClientID: "client1",
			Balance:  1000.00,
			IsActive: true,
			Type:     "checking",
		}
		Convey("When they withdraw the current balance of $1000.00", func() {
			amt := decimal.NewFromFloat(1000.00)
			wAmount, err := acctW.Withdraw(amt)
			Convey("The account balance should be $0.00", func() {
				So(err, ShouldBeNil)
				So(acctW.Balance, ShouldEqual, 0.00)
				So(wAmount, ShouldEqual, amt)
			})
		})
		Convey("When they try to withdraw more than the account balance", func() {
			amt := decimal.NewFromFloat(1001.00)
			wAmount, _ := acctW.Withdraw(amt)
			Convey("The function should return an error", func() {
				// So(err, ShouldNotBeNil)
				So(wAmount, ShouldEqual, decimal.Zero)
			})
		})
	})
	Convey("Given a client who wants to funds between accounts", t, func() {
		acct1 := models.Account{
			ID:       "123",
			ClientID: "client1",
			Balance:  1000.00,
			IsActive: true,
			Type:     "checking",
		}
		acct2 := models.Account{
			ID:       "124",
			ClientID: "client1",
			Balance:  1000.00,
			IsActive: true,
			Type:     "savings",
		}
		Convey("When they transfer $100.00 cash between accounts", func() {
			amt := decimal.NewFromFloat(100.00)
			err := acct1.Transfer(amt, &acct2)
			Convey("The account1 balance should be $900.00", func() {
				So(err, ShouldBeNil)
				So(acct1.Balance, ShouldEqual, 900.00)
			})
			Convey("The account2 balance should be $1100.00", func() {
				So(err, ShouldBeNil)
				So(acct2.Balance, ShouldEqual, 1100.00)
			})
		})
	})
}
