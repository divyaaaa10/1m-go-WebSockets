package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var messageCounter = 0 // Counter to track the number of received messages

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// Increment the message counter
		messageCounter++

		// Print received message with the number of the message
		fmt.Printf("Message #%d received: %s\n", messageCounter, string(p))

		// Send a numbered response back to the client
		err = conn.WriteMessage(messageType, []byte(fmt.Sprintf("Message #%d received by the server", messageCounter)))
		if err != nil {
			fmt.Println("Error writing message:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handleConnection)
	fmt.Println("WebSocket server started at ws://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
