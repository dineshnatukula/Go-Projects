// client.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	// Send a message
	message := "Hello from client!"
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		log.Fatal("Write error:", err)
	}

	// Read echoed message
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Printf("Received from server: %s\n", msg)

	time.Sleep(1 * time.Second)
}
