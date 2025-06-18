// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all CORS origins for development
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// Initialize Python bridge
	if err := initPythonBridge(); err != nil {
		log.Fatalf("Failed to initialize Python bridge: %v", err)
	}

	// WebSocket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}
		defer conn.Close()

		// Handle the WebSocket connection
		handleWebSocket(conn)
	})

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
