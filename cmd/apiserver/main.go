package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/openmind13/link-shortener/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/server.toml", "path to configuration file")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	config := server.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	server, err := server.New(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Start api server on port: %v\n", config.BindAddr)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
