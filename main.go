package main

import (
	"fmt"
	"log"
	"net"

	"main.go/config"
)

func main() {

	cfg := config.Load()
	port := cfg.Port

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is listening on port " + port + "...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue

		}
		fmt.Println("Client connected")
		conn.Close()

	}

}
