package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func main() {
	// Connect to the WebSocket server
	serverAddr := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error dialing WebSocket:", err)
	}
	defer conn.Close()

	// Send 1.8 million messages in half an hour
	for i := 0; i < 1800000; i++ {
		// Send a message to the WebSocket server
		err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Message %d", i)))
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}

		// Sleep for a fraction of a second to match 1000 requests per second
		time.Sleep(time.Millisecond) // 1000 requests per second

		// Optionally, read the response from the server
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Printf("Received response: %s\n", msg)
	}
}
