package store

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store struct
type Store struct {
	client *mongo.Client
}

// New - return new store
func New(connectionString string) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	store := &Store{
		client: client,
	}

	return store, nil
}

// AddURL ...
func (s *Store) AddURL(longurl, shorturl string) error {
	return nil
}

// GetURL ...
func (s *Store) GetURL() (string, error) {
	return "", nil
}
