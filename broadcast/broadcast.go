package broadcast

import (
	"net"
	"sync"
)

type Hub struct {
	clients map[net.Conn]bool
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[net.Conn]bool),
	}
}

func (s *Hub) Register(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[conn] = true
}

func (s *Hub) Unregister(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, conn)
	conn.Close()
}

func (s *Hub) Broadcast(message string, sender net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for client := range s.clients {
		if client != sender {
			client.Write([]byte(message + "\n"))
		}
	}
}
