package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"sync"
)

// WebSocket Handler
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow any origin for testing purposes
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")

	var count int
	var mu sync.Mutex // Mutex for synchronized access to file

	// Continuously receive data from WebSocket
	for {
		// Read message from WebSocket connection
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Process binary message
		if messageType == websocket.BinaryMessage {
			// Store the binary data into a file
			mu.Lock()
			fileName := fmt.Sprintf("data_%d.bin", count)
			err := os.WriteFile(fileName, p, 0644)
			if err != nil {
				log.Println("Error writing data to file:", err)
				mu.Unlock()
				break
			}
			fmt.Printf("Received and stored data of size: %d bytes in %s\n", len(p), fileName)
			mu.Unlock()
			count++
		}
	}
}

func main() {
	// Handle WebSocket requests
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

