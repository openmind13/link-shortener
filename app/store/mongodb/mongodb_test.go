package mongodb

import "testing"

var (
	testConfig = &Config{
		MongodbConnection: "mongodb://localhost:27017",
		DBName:            "linkshortener_test",
		CollectionName:    "links_test",
	}
)

func TestNew(t *testing.T) {
	_, err := NewMongodbStore(testConfig)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAdd(t *testing.T) {}

func TestGet(t *testing.T) {}
