package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

var (
	upgrader   = websocket.Upgrader{}
	reqCounter uint64
	dataSent   uint64
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		atomic.AddUint64(&reqCounter, 1)
		atomic.AddUint64(&dataSent, uint64(len(message)))
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	go func() {
		for {
			fmt.Printf("Requests Received: %d | Data Sent: %.2f MB\n",
				atomic.LoadUint64(&reqCounter), float64(atomic.LoadUint64(&dataSent))/(1024*1024))
		}
	}()

	log.Println("WebSocket server listening on ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
