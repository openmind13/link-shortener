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

// func init() {
// 	os.Setenv("CONFIG_PATH", "config/server.toml")
// }

// func getConfigFromEnv() (*server.Config, error) {
// 	shortURLLength, err := strconv.Atoi(os.Getenv("SHORT_URL_LENGTH"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	config := &server.Config{
// 		BindAddr:          os.Getenv("BIND_ADDR"),
// 		ShortURLLength:    shortURLLength,
// 		MongodbConnection: os.Getenv("MONGODB_CONNECTION"),
// 		DBName:            os.Getenv("DATABASE_NAME"),
// 		CollectionName:    os.Getenv("MONGODB_CONNECTION"),
// 	}

// 	return config, nil
// }

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

	// config, err := getConfigFromEnv()
	// if err != nil {
	// 	log.Fatal(err)
	// }

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
