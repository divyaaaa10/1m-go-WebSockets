package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	serverAddr  = "ws://localhost:8080/ws"
	messageSize = 500 * 1024 // 500 KB
	rate        = 1000       // 1000 requests per second
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	message := make([]byte, messageSize) // 500 KB binary payload
	ticker := time.NewTicker(time.Second / time.Duration(rate))
	defer ticker.Stop()

	sentCount := 0
	totalDataSent := 0

	for range ticker.C {
		if err := conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
			log.Println("Write error:", err)
			break
		}
		sentCount++
		totalDataSent += messageSize
		log.Printf("Sent: %d requests | Data Sent: %.2f MB", sentCount, float64(totalDataSent)/(1024*1024))
	}
}
