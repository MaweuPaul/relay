package relay

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MaweuPaul/relay/broadcast"
	"github.com/MaweuPaul/relay/config"
)

type Server struct {
	cfg *config.Config
	hub *broadcast.Hub
}

func NewServer() *Server {
	return &Server{
		cfg: config.Load(),
		hub: broadcast.NewHub(),
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
	s.hub.Register(conn)
	defer s.hub.Unregister(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received message from client:", message)
		s.hub.Broadcast(message, conn)
	}
	fmt.Println("Client disconnected")
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

	// goroutine to receive incoming messages from server
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// send outgoing messages to server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.Write([]byte(scanner.Text() + "\n"))
	}
	return nil
}
