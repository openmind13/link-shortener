package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config - server config
type Config struct {
	BindAddr          string `envconfig:"BIND_ADDR"`
	ShortURLLength    int    `envconfig:"SHORTURL_LENGTH"`
	MongodbConnection string `envconfig:"MONGODB_CONN"`
	DBName            string `envconfig:"DBNAME"`
	CollectionName    string `envconfig:"COLLECTION_NAME"`
}

var (
	once   sync.Once
	config Config

	errLoadingConfig     = errors.New("Error in loading config")
	errEmptyConfigFields = errors.New("Error: some config fields is empty")
)

// Get reads evn config once
func Get() *Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			log.Fatal(errLoadingConfig)
		}
		configBytes, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Config: ", string(configBytes))
		if config.BindAddr == "" ||
			config.CollectionName == "" ||
			config.DBName == "" ||
			config.MongodbConnection == "" {
			log.Fatal(errEmptyConfigFields)
		}
	})
	return &config
}
