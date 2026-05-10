package main

import (
	"log"

	"github.com/MaweuPaul/relay/relay"
)

func main() {
	client := relay.NewClient()
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
}
