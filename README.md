# Relay

> A basic TCP server that receives messages from clients and logs them to the console.

Built in Go.

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
	// Listen() reads the PORT from your environment or defaults to 8080
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
	// Connect() dials HOST:PORT from your environment
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

# Run the server
go run cmd/server/main.go
```

In a separate terminal, run the test client:

```bash
# Run the client
go run cmd/client/main.go
```

Type a message in the client terminal and hit enter. You should see it appear in the server terminal as `Received message from client: <message>`.

### Environment Variables

Configure your server and client using a `.env` file at the root of the project.

| Variable | Default   | Description                                       |
| -------- | --------- | ------------------------------------------------- |
| PORT     | 8080      | Port the server listens on and client connects to |
| HOST     | localhost | Host the client connects to                       |

*(Note: The example `.env.example` sets `PORT=9090` and `HOST=localhost`)*

---

## Architecture

Currently, the architecture is a straightforward TCP client-server interaction:
```
Client (os.Stdin) → TCP Dial → TCP Listener → Server Console Output
```

- The **Server** listens on the configured TCP port. For each connected client, it spawns a goroutine that reads from the connection into a 1024-byte buffer in a loop and prints the received bytes to the console.
- The **Client** dials the configured host and port, reads input from `os.Stdin` using a scanner, and writes the raw text directly to the TCP connection.

*Note: Broadcasting, channels, and message framing are not yet implemented.*

---

## Roadmap

### Done
- [x] TCP server (accepts connections and logs incoming data)
- [x] Client test utility (reads from stdin and sends to server)
- [x] Configuration loading via `.env` files

### Not Yet Implemented
- [ ] Broadcast messages to all connected clients
- [ ] Message framing
- [ ] Concurrent safe connection management
- [ ] Graceful disconnect and reconnect handling
- [ ] Tests and Benchmarks
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
