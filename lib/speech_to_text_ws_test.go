package lib

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestSpeechToTextStreamParams_toURLValues(t *testing.T) {
	params := SpeechToTextStreamParams{
		TranscriptionEngine: "Deepgram",
		InputFormat:         "wav",
		Language:            "en-US",
		SampleRate:          16000,
		Channels:            1,
		InterimResults:      true,
		Model:               "nova-2",
		Encoding:            "linear16",
	}

	values := params.toURLValues()

	if values.Get("transcription_engine") != "Deepgram" {
		t.Errorf("expected transcription_engine=Deepgram, got %s", values.Get("transcription_engine"))
	}
	if values.Get("input_format") != "wav" {
		t.Errorf("expected input_format=wav, got %s", values.Get("input_format"))
	}
	if values.Get("language") != "en-US" {
		t.Errorf("expected language=en-US, got %s", values.Get("language"))
	}
	if values.Get("sample_rate") != "16000" {
		t.Errorf("expected sample_rate=16000, got %s", values.Get("sample_rate"))
	}
	if values.Get("channels") != "1" {
		t.Errorf("expected channels=1, got %s", values.Get("channels"))
	}
	if values.Get("interim_results") != "true" {
		t.Errorf("expected interim_results=true, got %s", values.Get("interim_results"))
	}
	if values.Get("model") != "nova-2" {
		t.Errorf("expected model=nova-2, got %s", values.Get("model"))
	}
	if values.Get("encoding") != "linear16" {
		t.Errorf("expected encoding=linear16, got %s", values.Get("encoding"))
	}
}

func TestSpeechToTextStreamParams_toURLValues_empty(t *testing.T) {
	params := SpeechToTextStreamParams{}
	values := params.toURLValues()

	if len(values) != 0 {
		t.Errorf("expected empty values, got %v", values)
	}
}

func TestSttEvent_Unmarshal(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected SttEvent
	}{
		{
			name: "transcript event",
			json: `{"type":"transcript","transcript":"Hello world","is_final":true,"confidence":0.95}`,
			expected: SttEvent{
				Type:       "transcript",
				Transcript: "Hello world",
				IsFinal:    true,
				Confidence: 0.95,
			},
		},
		{
			name: "error event",
			json: `{"type":"error","error":"connection failed"}`,
			expected: SttEvent{
				Type:  "error",
				Error: "connection failed",
			},
		},
		{
			name: "interim transcript",
			json: `{"type":"transcript","transcript":"Hello","is_final":false,"confidence":0.8}`,
			expected: SttEvent{
				Type:       "transcript",
				Transcript: "Hello",
				IsFinal:    false,
				Confidence: 0.8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event SttEvent
			if err := json.Unmarshal([]byte(tt.json), &event); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if event.Type != tt.expected.Type {
				t.Errorf("expected Type=%s, got %s", tt.expected.Type, event.Type)
			}
			if event.Transcript != tt.expected.Transcript {
				t.Errorf("expected Transcript=%s, got %s", tt.expected.Transcript, event.Transcript)
			}
			if event.IsFinal != tt.expected.IsFinal {
				t.Errorf("expected IsFinal=%v, got %v", tt.expected.IsFinal, event.IsFinal)
			}
			if event.Error != tt.expected.Error {
				t.Errorf("expected Error=%s, got %s", tt.expected.Error, event.Error)
			}
		})
	}
}

func TestSpeechToTextWS_Integration(t *testing.T) {
	// Create a test WebSocket server
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	var serverConn *websocket.Conn
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify Authorization header
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Verify query parameters
		if r.URL.Query().Get("transcription_engine") != "Deepgram" {
			http.Error(w, "missing transcription_engine", http.StatusBadRequest)
			return
		}

		var err error
		serverConn, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("upgrade error: %v", err)
			return
		}
		defer serverConn.Close()

		// Echo back transcripts for any received audio
		for {
			messageType, data, err := serverConn.ReadMessage()
			if err != nil {
				return
			}

			if messageType == websocket.BinaryMessage {
				// Send a transcript response
				response := SttEvent{
					Type:       "transcript",
					Transcript: "test transcript",
					IsFinal:    true,
					Confidence: 0.99,
				}
				respBytes, _ := json.Marshal(response)
				serverConn.WriteMessage(websocket.TextMessage, respBytes)

				// Acknowledge received audio size
				t.Logf("server received %d bytes of audio", len(data))
			}
		}
	}))
	defer server.Close()

	// Convert to WebSocket URL
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Create client
	ws, err := NewSpeechToTextWS(SpeechToTextWSConfig{
		APIKey:  "test-api-key",
		BaseURL: wsURL,
	}, SpeechToTextStreamParams{
		TranscriptionEngine: "Deepgram",
		InputFormat:         "wav",
		Language:            "en-US",
	})
	if err != nil {
		t.Fatalf("failed to create WebSocket: %v", err)
	}
	defer ws.Close()

	// Wait for connection
	if err := ws.WaitForOpenWithTimeout(5 * time.Second); err != nil {
		t.Fatalf("failed to wait for open: %v", err)
	}

	// Verify connection is open
	if !ws.IsOpen() {
		t.Error("expected WebSocket to be open")
	}

	// Send audio data
	audioData := []byte("fake audio data")
	if err := ws.Send(audioData); err != nil {
		t.Fatalf("failed to send audio: %v", err)
	}

	// Receive transcript
	select {
	case event := <-ws.Events():
		if event.Type != "transcript" {
			t.Errorf("expected type=transcript, got %s", event.Type)
		}
		if event.Transcript != "test transcript" {
			t.Errorf("expected transcript='test transcript', got %s", event.Transcript)
		}
		if !event.IsFinal {
			t.Error("expected is_final=true")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for transcript")
	}
}

func TestSpeechToTextWS_WaitForOpen_Timeout(t *testing.T) {
	// Create a non-responsive server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Don't upgrade, just hang
		time.Sleep(10 * time.Second)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// This will fail to upgrade
	_, err := NewSpeechToTextWS(SpeechToTextWSConfig{
		APIKey:  "test-api-key",
		BaseURL: wsURL,
	}, SpeechToTextStreamParams{})

	if err == nil {
		t.Error("expected error, got nil")
	}
}
