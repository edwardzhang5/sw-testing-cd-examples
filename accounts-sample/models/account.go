package models

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

// AccountType type for enum to indicate the type of account.
type AccountType string

// AccountType type.
const (
	Checking    AccountType = "checking"
	Savings     AccountType = "savings"
	MoneyMarket AccountType = "money-market"
	Loan        AccountType = "loan"
	Investment  AccountType = "investment"
	CreditCard  AccountType = "credit-card"
)

// Account struct for account info.
type Account struct {
	ID        string      `json:"id,omitempty" firestore:"id"`
	ClientID  string      `json:"client_id,omitempty" firestore:"client_id"`
	Type      AccountType `json:"type,omitempty" firestore:"type"`
	Balance   float64     `json:"balance,omitempty" firestore:"balance"`
	IsActive  bool        `json:"is_active,omitempty" firestore:"is_active"`
	CeatedAt  time.Time   `json:"ceated_at,omitempty" firestore:"created_at"`
	UpdatedAt time.Time   `json:"updated_at,omitempty" firestore:"updated_at"`
}

// Transfer transfers funds between 2 accounts
func (a *Account) Transfer(amount decimal.Decimal, toAcct *Account) error {
	amt := amount.Truncate(2)
	balanceD := decimal.NewFromFloat(a.Balance).Truncate(2)
	if balanceD.GreaterThan(amt) {
		wAmt, err := a.Withdraw(amt)
		if err != nil {
			return err
		}
		_ = toAcct.Deposit(wAmt)
		return nil
	}

	err := errors.New("Insufficient funds for this transfer - transaction cancelled")
	return err
}

// Deposit adds the specified amount ot the account.
func (a *Account) Deposit(amount decimal.Decimal) decimal.Decimal {
	balanceD := decimal.NewFromFloat(a.Balance).Truncate(2)
	balanceD = balanceD.Add(amount.Truncate(2))
	a.Balance, _ = balanceD.Float64()
	return balanceD
}

// Withdraw pulls cash from the account and returns the amount pulled
// for further processing
func (a *Account) Withdraw(amount decimal.Decimal) (decimal.Decimal, error) {
	amt := amount.Truncate(2)
	balanceD := decimal.NewFromFloat(a.Balance).Truncate(2)
	if balanceD.GreaterThanOrEqual(amt) {
		balanceD = balanceD.Sub(amt)
		a.Balance, _ = balanceD.Float64()
		return amt, nil
	}

	err := errors.New("Insufficient funds for this withdrawal - transaction cancelled")
	return decimal.Zero, err
}
