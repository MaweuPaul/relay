package relay

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"

	"github.com/MaweuPaul/relay/config"
)

type Server struct {
	cfg     *config.Config
	clients map[net.Conn]bool
	mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		cfg:     config.Load(),
		clients: make(map[net.Conn]bool),
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

func (s *Server) register(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[conn] = true
}

func (s *Server) unregister(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, conn)
	conn.Close()
}

func (s *Server) broadcast(message string, sender net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.clients {
		if client != sender {
			client.Write([]byte(message + "\n"))
		}
	}
}

func (s *Server) handleClient(conn net.Conn) {
	s.register(conn)
	defer s.unregister(conn)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received message from client:", message)
		s.broadcast(message, conn)
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
