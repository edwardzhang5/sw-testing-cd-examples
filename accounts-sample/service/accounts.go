package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/drbyronw/accounts/db"
	"github.com/drbyronw/accounts/models"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
)

const truncateAmt = 2

// AccountsService interface to handle account operations.
type AccountsService interface {
	Create(account models.Account) error
	Find(id string) (*models.Account, error)
	FindAll(clientID string) ([]models.Account, error)
	Init(path string, client *db.FSRepo)
	Transfer(from, to string, amount float64) error
	Deposit(acct string, amount float64) (decimal.Decimal, error)
}

type accountsService struct {
	client *db.FSRepo
	path   string
}

// NewAccountsService initializes the accounts service with a db name/path
// and db client
func NewAccountsService(path string, client *db.FSRepo) AccountsService {
	var as accountsService

	as.Init(path, client)

	return &as
}

func (as *accountsService) Deposit(acct string, amount float64) (decimal.Decimal, error) {
	var err error

	amt := decimal.NewFromFloat(amount).Truncate(truncateAmt)

	account, err := as.Find(acct)
	if err != nil {
		return decimal.Zero, errors.New("account does not exist")
	}

	newBalance := account.Deposit(amt)
	account.UpdatedAt = time.Now()

	ctx := context.Background()

	err = as.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		_, err := as.client.Collection(as.path).Doc(account.ID).Set(ctx, map[string]interface{}{
			"balance":    account.Balance,
			"updated_at": account.UpdatedAt,
		}, firestore.MergeAll)
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return decimal.Zero, err
	}
	return newBalance, nil
}

func (as *accountsService) Transfer(from, to string, amount float64) error {
	var err error

	amt := decimal.NewFromFloat(amount).Truncate(truncateAmt)

	fromAcct, err := as.Find(from)
	if err != nil {
		return errors.New("account does not exist")
	}

	toAcct, err := as.Find(to)
	if err != nil {
		return errors.New("account does not exist")
	}

	if fromAcct.ClientID != toAcct.ClientID {
		return errors.New("Transfers can only take for accounts owned by the same client")
	}

	err = fromAcct.Transfer(amt, toAcct)
	if err != nil {
		return err
	}
	fromAcct.UpdatedAt = time.Now()
	toAcct.UpdatedAt = time.Now()

	ctx := context.Background()
	err = as.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		_, err := as.client.Collection(as.path).Doc(fromAcct.ID).Set(ctx, map[string]interface{}{
			"balance":    fromAcct.Balance,
			"updated_at": fromAcct.UpdatedAt,
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
		_, err = as.client.Collection(as.path).Doc(toAcct.ID).Set(ctx, map[string]interface{}{
			"balance":    toAcct.Balance,
			"updated_at": toAcct.UpdatedAt,
		}, firestore.MergeAll)
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		return err
	}

	return err
}

func (as *accountsService) Init(path string, client *db.FSRepo) {
	as.path = path
	as.client = client
}

func (as *accountsService) FindAll(clientID string) ([]models.Account, error) {
	var accts []models.Account

	ctx := context.Background()

	iter := as.client.Collection(as.path).Where("client_id", "==", clientID).Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var acct models.Account
		if err = doc.DataTo(&acct); err != nil {
			logrus.Errorf("unable to add account: %+v - error: %v", acct, err)
		} else {
			accts = append(accts, acct)
		}
	}

	if len(accts) > 0 {
		return accts, nil
	}

	return nil, fmt.Errorf("no accounts available for client_id: %s", clientID)
}

func (as *accountsService) Find(id string) (*models.Account, error) {
	var acct models.Account

	ctx := context.Background()

	snap, err := as.client.Collection(as.path).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(&acct)
	if err != nil {
		return nil, err
	}

	return &acct, err
}

func (as *accountsService) Create(account models.Account) error {
	var err error
	ctx := context.Background()

	if account.ID == "" {
		acctRef := as.client.Collection(as.path).NewDoc()
		account.ID = acctRef.ID
		_, err = acctRef.Set(ctx, map[string]interface{}{
			"id":         account.ID,
			"client_id":  account.ClientID,
			"type":       account.Type,
			"balance":    account.Balance,
			"is_active":  account.IsActive,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		})
		if err != nil {
			return err
		}
	} else {
		_, err = as.client.Collection(as.path).Doc(account.ID).Set(ctx, map[string]interface{}{
			"id":         account.ID,
			"client_id":  account.ClientID,
			"type":       account.Type,
			"balance":    account.Balance,
			"is_active":  account.IsActive,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		})
		if err != nil {
			return err
		}
	}
	return err
}
