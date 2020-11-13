package db

import (
	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
)

type RepoClient interface {
	Close() error
	Collection(path string) *firestore.CollectionRef
}

type RepoCollectionRef interface {
	Doc(id string) RepoDocumentRef
}

type RepoDocumentRef interface {
	Get(ctx context.Context) (RepoDocumentSnapshot, error)
}

type RepoDocumentSnapshot interface {
	DataTo(p interface{}) error
}
