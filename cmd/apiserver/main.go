package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/openmind13/link-shortener/app/server"
)

var (
	configPath = "config/server.toml"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("start application")

	urlLength, err := strconv.Atoi(os.Getenv("URL_LENGTH"))
	if err != nil {
		return err
	}

	config := &server.Config{
		BindAddr:          os.Getenv("BIND_ADDR"),
		ShortURLLength:    urlLength,
		MongodbConnection: os.Getenv("MONGODB_CONNECTION"),
		DBName:            os.Getenv("DBNAME"),
		CollectionName:    os.Getenv("COLLECTION"),
	}

	server, err := server.New(config)
	if err != nil {
		return err
	}

	fmt.Printf("Server running on: %v\n", config.BindAddr)

	if err := server.Start(); err != nil {
		return err
	}

	return nil
}
