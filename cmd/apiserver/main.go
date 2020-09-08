package main

import (
	"log"

	"github.com/openmind13/link-shortener/app/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
