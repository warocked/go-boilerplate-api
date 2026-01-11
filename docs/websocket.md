# WebSocket Support

This application includes WebSocket support for real-time bidirectional communication.

## Endpoint

- **WebSocket URL**: `ws://localhost:8080/ws` (development)
- **WebSocket URL**: `wss://yourdomain.com/ws` (production with SSL)

## Usage

### Client Connection (JavaScript)

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function(event) {
    console.log('WebSocket connected');
    ws.send('Hello Server!');
};

ws.onmessage = function(event) {
    console.log('Message from server:', event.data);
};

ws.onerror = function(error) {
    console.error('WebSocket error:', error);
};

ws.onclose = function(event) {
    console.log('WebSocket closed');
};
```

### Client Connection (Go)

```go
import (
    "github.com/gorilla/websocket"
    "net/url"
)

u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
if err != nil {
    log.Fatal("dial:", err)
}
defer conn.Close()

// Send message
err = conn.WriteMessage(websocket.TextMessage, []byte("Hello Server!"))
if err != nil {
    log.Println("write:", err)
}

// Read messages
for {
    _, message, err := conn.ReadMessage()
    if err != nil {
        log.Println("read:", err)
        break
    }
    log.Printf("recv: %s", message)
}
```

## Server-Side Broadcasting

Broadcast messages to all connected clients:

```go
import "go-boilerplate-api/internal/api/handlers"

// Broadcast to all connected clients
handlers.BroadcastMessage([]byte("Hello all clients!"))
```

## Get Connected Clients Count

```go
import "go-boilerplate-api/internal/api/handlers"

count := handlers.GetConnectedClientsCount()
fmt.Printf("Connected clients: %d\n", count)
```

## Features

- **Broadcast messaging**: Messages are broadcast to all connected clients
- **Connection management**: Automatic cleanup on disconnect
- **Thread-safe**: Uses mutex for concurrent access
- **Error handling**: Graceful error handling and connection cleanup

## Production Considerations

1. **SSL/TLS**: Use `wss://` in production (WebSocket over SSL)
2. **Rate Limiting**: Consider adding rate limiting for WebSocket connections
3. **Authentication**: Add authentication if needed (check tokens in upgrade handler)
4. **Message Size Limits**: Consider setting message size limits
5. **Connection Limits**: Monitor and limit concurrent connections
6. **Heartbeat/Ping-Pong**: Consider implementing ping-pong for connection health

## Security

- Implement authentication in the WebSocket upgrade handler if needed
- Validate and sanitize all incoming messages
- Consider rate limiting to prevent abuse
- Use SSL/TLS (wss://) in production

## Example: Adding Authentication

```go
func WebSocketUpgrade(c *fiber.Ctx) error {
    // Check authentication token
    token := c.Query("token")
    if !isValidToken(token) {
        return c.Status(401).SendString("Unauthorized")
    }
    
    if websocket.IsWebSocketUpgrade(c) {
        c.Locals("allowed", true)
        return c.Next()
    }
    return fiber.ErrUpgradeRequired
}
```

## Architecture

- **Connection Pool**: All connections are stored in a thread-safe map
- **Broadcasting**: Messages from one client are broadcast to all connected clients
- **Cleanup**: Automatic cleanup on disconnect or error
- **Concurrent Safe**: Uses RWMutex for thread-safe access
