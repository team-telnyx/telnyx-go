// Example: Text-to-Speech WebSocket streaming
//
// This example demonstrates how to use the Telnyx Go SDK to stream
// text to the Text-to-Speech API and receive audio chunks in real-time.
//
// Usage:
//
//	export TELNYX_API_KEY=your-api-key
//	go run main.go "Hello, world!"
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/lib"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("TELNYX_API_KEY")
	if apiKey == "" {
		log.Fatal("TELNYX_API_KEY environment variable is required")
	}

	// Get text from arguments
	text := "Hello, this is a test of the Telnyx Text-to-Speech API using WebSocket streaming."
	if len(os.Args) >= 2 {
		text = os.Args[1]
	}

	fmt.Printf("Text to synthesize: %s\n", text)

	// Create WebSocket connection
	ws, err := lib.NewTextToSpeechWS(lib.TextToSpeechWSConfig{
		APIKey: apiKey,
	}, lib.TextToSpeechStreamParams{
		Voice:        "telnyx.NaturalHD.Alloy",
		OutputFormat: "mp3",
	})
	if err != nil {
		log.Fatalf("Failed to create WebSocket: %v", err)
	}
	defer ws.Close()

	// Wait for connection to be ready
	fmt.Println("Connecting to Text-to-Speech WebSocket...")
	if err := ws.WaitForOpenWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Println("Connected!")

	// Collect audio data
	var audioBuffer bytes.Buffer
	done := make(chan struct{})

	// Start receiving audio in a goroutine
	go func() {
		defer close(done)
		for event := range ws.Events() {
			switch event.Type {
			case "audio_chunk":
				audioData := event.AudioBytes()
				if audioData != nil {
					audioBuffer.Write(audioData)
					fmt.Printf("Received audio chunk: %d bytes (total: %d bytes)\n", len(audioData), audioBuffer.Len())
				}
				if event.TimeToFirstAudioFrameMs > 0 {
					fmt.Printf("Time to first audio frame: %d ms\n", event.TimeToFirstAudioFrameMs)
				}
			case "final":
				fmt.Println("Received final event - synthesis complete")
			case "error":
				fmt.Printf("[ERROR] %s\n", event.Error)
			default:
				fmt.Printf("[%s] %+v\n", event.Type, event)
			}
		}
		fmt.Println("Events channel closed")
	}()

	// Send initial handshake with optional voice settings
	fmt.Println("Sending handshake...")
	if err := ws.Send(lib.TtsClientEvent{
		Text: " ", // Space as handshake
		VoiceSettings: map[string]interface{}{
			"voice_speed": 1.0,
		},
	}); err != nil {
		log.Fatalf("Failed to send handshake: %v", err)
	}

	// Small delay after handshake
	time.Sleep(100 * time.Millisecond)

	// Send text to synthesize
	fmt.Println("Sending text for synthesis...")
	if err := ws.Send(lib.TtsClientEvent{
		Text: text,
	}); err != nil {
		log.Fatalf("Failed to send text: %v", err)
	}

	// Wait for synthesis to complete
	fmt.Println("Waiting for audio...")
	time.Sleep(5 * time.Second)

	// Close the connection
	fmt.Println("Closing connection...")
	ws.Close()

	// Wait for event processing to complete
	<-done

	// Save audio to file
	if audioBuffer.Len() > 0 {
		outputPath := "output.mp3"
		if err := os.WriteFile(outputPath, audioBuffer.Bytes(), 0644); err != nil {
			log.Fatalf("Failed to save audio: %v", err)
		}
		fmt.Printf("Saved %d bytes of audio to %s\n", audioBuffer.Len(), outputPath)
	} else {
		fmt.Println("No audio received")
	}

	fmt.Println("Done!")
}
