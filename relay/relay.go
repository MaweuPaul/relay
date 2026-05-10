package relay

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MaweuPaul/relay/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer() *Server {
	return &Server{
		cfg: config.Load(),
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":"+s.cfg.Port)
	if err != nil {
		return err
	}
	fmt.Println("Server is listening on port " + s.cfg.Port + "...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Client connected")
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading from client:", err)
			return
		}
		fmt.Println("Received message from client:", string(buf[:n]))
	}
}

type Client struct {
	cfg *config.Config
}

func NewClient() *Client {
	return &Client{
		cfg: config.Load(),
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.cfg.Host+":"+c.cfg.Port)
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("Connected to Relay!")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		conn.Write([]byte(text))
	}
	return nil
}
