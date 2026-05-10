package main

import (
	"fmt"
	"log"
	"net"

	"github.com/MaweuPaul/relay/config"
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
		go handleClient(conn)

	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading from client:", err)
			return
		}
		message := string(buf[:n])
		fmt.Println("Received message from client:", message)
	}
}
