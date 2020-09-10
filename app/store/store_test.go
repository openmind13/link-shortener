package store

import (
	"testing"
)

var (
	MongodbConnection = "mongodb://localhost:27017"
	DBName            = "linkshortener_test"
	CollectionName    = "links_test"
)

func Test_New(t *testing.T) {
	testConfig := &Config{
		MongodbConnection: MongodbConnection,
		DBName:            DBName,
		CollectionName:    CollectionName,
	}

	_, err := New(testConfig)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_AddURL(t *testing.T) {

}
func Test_GetLongURL(t *testing.T) {

}
