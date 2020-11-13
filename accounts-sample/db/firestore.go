package db

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FSRepo stuct to hold the firestore client
type FSRepo struct {
	*firestore.Client
}

// NewFSRepoClient returns a RepoClient to access the database
func NewFSRepoClient() (*FSRepo, error) {
	c, err := newFSClient()
	if err != nil {
		log.Fatalf("Unable to connect to Firestore: %v", err)
	}
	var fs = &FSRepo{c.(*firestore.Client)}
	return fs, err
}

// newFSClient returns
func newFSClient() (RepoClient, error) {
	ctx := context.Background()
	env := os.Getenv("ENVIRONMENT")
	if env == "development" {
		sa := option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT"))
		app, err := firebase.NewApp(ctx, nil, sa)
		if err != nil {
			return nil, err
		}
		client, err := app.Firestore(ctx)
		if err != nil {
			return nil, err
		}

		return client, err
	}

	conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, err

}
