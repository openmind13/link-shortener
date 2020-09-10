package store

import (
	"context"
	"time"

	"github.com/openmind13/link-shortener/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Store struct
type Store struct {
	config *Config
	client *mongo.Client
}

// New - return new store
func New(config *Config) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongodbConnection))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	store := &Store{
		config: config,
		client: client,
	}

	return store, nil
}

// AddURL - save short url into db
func (s *Store) AddURL(longurl, shorturl string) error {
	type data struct {
		longurl  string
		shorturl string
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := s.client.Database("linkshortener").Collection("links")
	if _, err := collection.InsertOne(ctx, bson.M{"longurl": longurl, "shorturl": shorturl}); err != nil {
		return err
	}

	return nil
}

// GetLongURL - return long url
func (s *Store) GetLongURL(shorturl string) (string, error) {
	collection := s.client.Database("linkshortener").Collection("links")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var d model.RequestAddCustom
	if err := collection.FindOne(ctx, bson.M{"shorturl": shorturl}).Decode(&d); err != nil {
		return "", err
	}

	return d.LongURL, nil
}
