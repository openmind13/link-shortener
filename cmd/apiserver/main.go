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

func init() {
	os.Setenv("CONFIG_PATH", "config/server.toml")
}

func loadConfigFromTomlFile(configPath string) (*server.Config, error) {
	config := &server.Config{}

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	fmt.Println("start application")

	configPath := os.Getenv("CONFIG_PATH")

	config, err := loadConfigFromTomlFile(configPath)
	if err != nil {
		fmt.Println("Error in decoding toml file")
		log.Fatal(err)
	}

	server, err := server.New(config)
	if err != nil {
		fmt.Printf("Server not created!\n")
		log.Fatal(err)
	}

	fmt.Printf("Start api server on port: %v\n", config.BindAddr)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
