package main

import (
	"log"

	"github.com/MaweuPaul/relay/relay"
)

func main() {
	server := relay.NewServer()
	if err := server.Listen(); err != nil {
		log.Fatal(err)
	}
}
