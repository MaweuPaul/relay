# Relay

> Takes a message from one client. Sends it to everyone else. That's it.

Built in Go. Concurrent by default. Drop it into any application that needs real-time messaging.

---

## Quick Start

You can use Relay as a Go library. The module path is `github.com/MaweuPaul/relay`.

### Server

```go
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
```

### Client

```go
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
```

---

## Development

### Prerequisites

- Go 1.26.2+
- Git

### Running locally

```bash
git clone https://github.com/MaweuPaul/relay.git
cd relay
cp .env.example .env

# Terminal 1 — run the server
go run cmd/server/main.go

# Terminal 2 — run a client
go run cmd/client/main.go

# Terminal 3 — run another client
go run cmd/client/main.go
```

Type a message in one client terminal. It broadcasts to all other connected clients.

### Environment variables

| Variable | Default   | Description                                       |
| -------- | --------- | ------------------------------------------------- |
| PORT     | 9090      | Port the server listens on and client connects to |
| HOST     | localhost | Host the client connects to                       |

### Testing the connection

**Mac/Linux:**

```bash
telnet localhost 9090
```

**Windows — run the built in test client:**

```bash
go run cmd/client/main.go
```

Or enable telnet on Windows:

```bash
dism /online /Enable-Feature /FeatureName:TelnetClient
```

---

## Architecture

```
Client → TCP Listener → Broadcast Hub → All Connected Clients
```

- The **Server** accepts TCP connections. Each client gets its own goroutine.
- The **Broadcast Hub** maintains the list of connected clients and distributes messages. Protected by a mutex for concurrent safe access.
- The **Client** connects via TCP, sends messages from stdin, and receives broadcasts from the server in a separate goroutine.

TCP is a stream protocol. Relay uses newline delimited message framing to ensure messages arrive cleanly and never merge unexpectedly.

---

## Roadmap

### Foundation

- [x] TCP server
- [x] Client test utility
- [x] Broadcast to all connected clients
- [x] Concurrent safe connection management
- [ ] Message framing
- [ ] Graceful disconnect and reconnect handling
- [ ] Tests
- [ ] Benchmarks

### Expansion

- [ ] WebSocket transport
- [ ] Room based messaging
- [ ] Authentication and rate limiting
- [ ] JavaScript and Python client SDKs

---

## Built with

- Go 1.26.2
- [godotenv](https://github.com/joho/godotenv)

---

## License

MIT
