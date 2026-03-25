// Example: Speech-to-Text WebSocket streaming
//
// This example demonstrates how to use the Telnyx Go SDK to stream
// audio to the Speech-to-Text API and receive real-time transcriptions.
//
// Usage:
//
//	export TELNYX_API_KEY=your-api-key
//	go run main.go path/to/audio.wav
package main

import (
	"fmt"
	"io"
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

	// Get audio file path from arguments
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <audio-file>")
	}
	audioPath := os.Args[1]

	// Read audio file
	audioData, err := os.ReadFile(audioPath)
	if err != nil {
		log.Fatalf("Failed to read audio file: %v", err)
	}

	fmt.Printf("Read %d bytes from %s\n", len(audioData), audioPath)

	// Create WebSocket connection
	ws, err := lib.NewSpeechToTextWS(lib.SpeechToTextWSConfig{
		APIKey: apiKey,
	}, lib.SpeechToTextStreamParams{
		TranscriptionEngine: "Deepgram",
		InputFormat:         "wav",
		Language:            "en-US",
		InterimResults:      true,
	})
	if err != nil {
		log.Fatalf("Failed to create WebSocket: %v", err)
	}
	defer ws.Close()

	// Wait for connection to be ready
	fmt.Println("Connecting to Speech-to-Text WebSocket...")
	if err := ws.WaitForOpenWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	fmt.Println("Connected!")

	// Start receiving transcripts in a goroutine
	done := make(chan struct{})
	go func() {
		defer close(done)
		for event := range ws.Events() {
			switch event.Type {
			case "transcript":
				if event.IsFinal {
					fmt.Printf("[FINAL] %s (confidence: %.2f)\n", event.Transcript, event.Confidence)
				} else {
					fmt.Printf("[INTERIM] %s\n", event.Transcript)
				}
			case "error":
				fmt.Printf("[ERROR] %s\n", event.Error)
			default:
				fmt.Printf("[%s] %+v\n", event.Type, event)
			}
		}
		fmt.Println("Events channel closed")
	}()

	// Send audio data in chunks (simulating real-time streaming)
	chunkSize := 4096 // 4KB chunks
	for i := 0; i < len(audioData); i += chunkSize {
		end := i + chunkSize
		if end > len(audioData) {
			end = len(audioData)
		}

		chunk := audioData[i:end]
		if err := ws.Send(chunk); err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Warning: failed to send chunk: %v", err)
		}

		// Small delay to simulate real-time streaming
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("Finished sending audio, waiting for final transcripts...")

	// Wait a bit for final transcripts
	time.Sleep(2 * time.Second)

	// Close the connection
	fmt.Println("Closing connection...")
	ws.Close()

	// Wait for event processing to complete
	<-done
	fmt.Println("Done!")
}
