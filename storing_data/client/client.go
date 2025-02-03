package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"

	"time"
)

// Function to generate random data of specified size
func generateRandomData(size int) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		log.Fatal("Error generating random data:", err)
	}
	return data
}

func main() {
	// Connect to the WebSocket server
	serverAddr := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer conn.Close()

	log.Println("Connected to server at", serverAddr)

	// Send data every second for 30 minutes (1800 seconds)
	for i := 0; i < 10; i++ {
		data := generateRandomData(1024) // 1KB of data

		// Send binary data to server
		err := conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			log.Println("Error sending data:", err)
			break
		}

		// Log sent data
		fmt.Printf("Sent data of size: %d bytes\n", len(data))

		// Wait 1 second before sending the next data
		time.Sleep(1 * time.Second)
	}

	log.Println("Finished sending data.")
}
