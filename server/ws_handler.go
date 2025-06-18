package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// handleWebSocket manages a single WebSocket connection
func handleWebSocket(conn *websocket.Conn) {
	// Create channels for audio streaming and responses
	audioChan := make(chan []byte, 128)  // PCM frames ▶ Python
	resultChan := make(chan []byte, 128) // Python ▶ client

	// Start goroutine to read audio data from client
	go func() {
		defer close(audioChan)
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error reading message: %v", err)
				return
			}

			// Only process binary messages (audio data)
			if messageType == websocket.BinaryMessage {
				audioChan <- data
			}
		}
	}()

	// Start goroutine to write responses back to client
	go func() {
		for response := range resultChan {
			if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
				log.Printf("Error writing message: %v", err)
				return
			}
		}
	}()

	// TODO: Add real-time Whisper transcription
	// TODO: Add emotion detection logic
	// TODO: Add GPT-based character prompt generation
	// TODO: Add voice synthesis and audio streaming back to client

	// Start processing audio through Python bridge
	processAudioStream(audioChan, resultChan)
}
