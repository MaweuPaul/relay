package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/MaweuPaul/relay/config"
)

func main() {
	cfg := config.Load()
	conn, err := net.Dial("tcp", cfg.Host+":"+cfg.Port)
	if err != nil {
		fmt.Println("Could not connect:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to Relay!")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		conn.Write([]byte(text))
	}
}
