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

## Roadmap

- [x] TCP server
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
