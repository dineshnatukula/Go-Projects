// server.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	client := &Client{conn: ws, send: make(chan []byte)}
	clients[client] = true

	go client.read()
	go client.write()
}

func (c *Client) read() {
	defer func() {
		c.conn.Close()
		delete(clients, c)
	}()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		broadcast <- msg
	}
}

func (c *Client) write() {
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		for client := range clients {
			client.send <- msg
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	go handleBroadcast()

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
