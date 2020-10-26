package main

import (
	"fmt"
	"log"

	"github.com/openmind13/link-shortener/app/config"
	"github.com/openmind13/link-shortener/app/server"
)

// var (
// 	configPath = "config/server.toml"
// )

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config := config.Get()
	server, err := server.New(config)
	if err != nil {
		return err
	}
	fmt.Printf("Starting server on: %v\n", config.BindAddr)
	if err := server.Start(); err != nil {
		return err
	}
	return nil
}
