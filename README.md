# Relay

> Takes a message from one client. Sends it to everyone else. That's it.

Built in Go. Concurrent by default. Drop it into any application that needs real-time messaging.

---

## Quick Start

```go
import "github.com/yourname/relay"

server := relay.NewServer()
server.Listen(":8080")
```

That's it. Your app has messaging.

---

## What's happening underneath

- Goroutine per connected client
- Channel based message broadcasting
- Concurrent safe connection management
- Automatic cleanup on disconnect
- Handles thousands of simultaneous connections

---

## Development

### Prerequisites

- Go 1.21+
- Git

### Running locally

```bash
git clone https://github.com/MaweuPaul/relay.git
cd relay
cp .env.example .env
go run main.go
```

### Environment variables

| Variable | Default   | Description                |
| -------- | --------- | -------------------------- |
| PORT     | 9090      | Port the server listens on |
| HOST     | localhost | Host the server runs on    |

### Testing the connection

**Mac/Linux:**

```bash
telnet localhost 9090
```

**Windows — run the built in test client:**

```bash
go run client/main.go
```

Type a message and hit enter. You should see it appear in the server terminal.

Or enable telnet on Windows:

```bash
dism /online /Enable-Feature /FeatureName:TelnetClient
```

---

## Roadmap

- [x] TCP server
- [x] Client test utility
- [ ] Broadcast to all connected clients
- [ ] WebSocket support
- [ ] Room based messaging
- [ ] Rate limiting
- [ ] JavaScript and Python client SDKs

---

## Built with

Go

---

## License

MIT
