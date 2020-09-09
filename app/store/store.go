package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store struct
type Store struct {
	client *mongo.Client
	config *Config
}

// New - return new store
func New(config *Config) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongodbConnection))
	if err != nil {
		return nil, err
	}

	store := &Store{
		client: client,
		config: config,
	}

	return store, nil
}

// AddURL ...
func (s *Store) AddURL(longurl, shorturl string) error {
	type data struct {
		longurl  string
		shorturl string
	}

	collection := s.client.Database("linkshortener").Collection("links")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, data{
		longurl:  longurl,
		shorturl: shorturl,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)

	return nil
}

// GetLongURL ...
func (s *Store) GetLongURL(shorturl string) (string, error) {
	collection := s.client.Database("linkshortener").Collection("links")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return "", err
	}
	defer cursor.Close(ctx)

	var longurl string
	cursor.All(ctx, longurl)

	// for cursor.Next(ctx) {
	// 	var data string
	// 	cursor.Decode(data)
	// 	datas = append(datas, data)
	// }
	if err := cursor.Err(); err != nil {
		return "", err
	}

	return longurl, nil
}
