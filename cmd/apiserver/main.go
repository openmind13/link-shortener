package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/openmind13/link-shortener/app/server"
)

var (
	configPath = "config/server.toml"
)

func main() {
	fmt.Println("start application")

	configPath := os.Getenv("CONFIG_PATH")

	config := &server.Config{}
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	server, err := server.New(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server running on: %v\n", config.BindAddr)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
