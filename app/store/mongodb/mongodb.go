package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/openmind13/link-shortener/app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

// Config - db configuration
type Config struct {
	MongodbConnection string
	DBName            string
	CollectionName    string
}

// Store struct
type Store struct {
	config *Config
	client *mongo.Client
}

// NewMongodbStore - return new store
func NewMongodbStore(config *Config) (*Store, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongodbConnection))
	if err != nil {
		fmt.Println("in connect new client")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		fmt.Println("error in connect")
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("error in ping")
		return nil, err
	}
	store := &Store{
		config: config,
		client: client,
	}
	return store, nil
}

// Add - save short url into db
func (s *Store) Add(data model.Data) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := s.client.Database(s.config.DBName).Collection(s.config.CollectionName)
	if _, err := collection.InsertOne(ctx, bson.M{"longurl": data.LongURL, "shorturl": data.ShortURL}); err != nil {
		return err
	}
	return nil
}

// Get - return long url
func (s *Store) Get(data model.Data) (model.Data, error) {
	collection := s.client.Database(s.config.DBName).Collection(s.config.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection.FindOne(ctx, bson.M{"shorturl": data.ShortURL}).Decode(data.LongURL); err != nil {
		return model.Data{}, err
	}
	return data, nil
}
