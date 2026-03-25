package lib

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestTextToSpeechStreamParams_toURLValues(t *testing.T) {
	params := TextToSpeechStreamParams{
		Voice:        "telnyx.NaturalHD.Alloy",
		Provider:     "telnyx",
		OutputFormat: "mp3",
		SampleRate:   44100,
		Language:     "en-US",
	}

	values := params.toURLValues()

	if values.Get("voice") != "telnyx.NaturalHD.Alloy" {
		t.Errorf("expected voice=telnyx.NaturalHD.Alloy, got %s", values.Get("voice"))
	}
	if values.Get("provider") != "telnyx" {
		t.Errorf("expected provider=telnyx, got %s", values.Get("provider"))
	}
	if values.Get("output_format") != "mp3" {
		t.Errorf("expected output_format=mp3, got %s", values.Get("output_format"))
	}
	if values.Get("sample_rate") != "44100" {
		t.Errorf("expected sample_rate=44100, got %s", values.Get("sample_rate"))
	}
	if values.Get("language") != "en-US" {
		t.Errorf("expected language=en-US, got %s", values.Get("language"))
	}
}

func TestTextToSpeechStreamParams_toURLValues_empty(t *testing.T) {
	params := TextToSpeechStreamParams{}
	values := params.toURLValues()

	if len(values) != 0 {
		t.Errorf("expected empty values, got %v", values)
	}
}

func TestTtsEvent_Unmarshal(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected TtsEvent
	}{
		{
			name: "audio chunk event",
			json: `{"type":"audio_chunk","audio":"SGVsbG8=","text":"Hello","cached":false,"timeToFirstAudioFrameMs":150}`,
			expected: TtsEvent{
				Type:                    "audio_chunk",
				Audio:                   "SGVsbG8=",
				Text:                    "Hello",
				Cached:                  false,
				TimeToFirstAudioFrameMs: 150,
			},
		},
		{
			name: "final event",
			json: `{"type":"final","isFinal":true}`,
			expected: TtsEvent{
				Type:    "final",
				IsFinal: true,
			},
		},
		{
			name: "error event",
			json: `{"type":"error","error":"synthesis failed"}`,
			expected: TtsEvent{
				Type:  "error",
				Error: "synthesis failed",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event TtsEvent
			if err := json.Unmarshal([]byte(tt.json), &event); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			if event.Type != tt.expected.Type {
				t.Errorf("expected Type=%s, got %s", tt.expected.Type, event.Type)
			}
			if event.Audio != tt.expected.Audio {
				t.Errorf("expected Audio=%s, got %s", tt.expected.Audio, event.Audio)
			}
			if event.Text != tt.expected.Text {
				t.Errorf("expected Text=%s, got %s", tt.expected.Text, event.Text)
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

func TestTtsEvent_AudioBytes(t *testing.T) {
	tests := []struct {
		name     string
		audio    string
		expected []byte
	}{
		{
			name:     "valid base64",
			audio:    "SGVsbG8gV29ybGQ=", // "Hello World"
			expected: []byte("Hello World"),
		},
		{
			name:     "empty audio",
			audio:    "",
			expected: nil,
		},
		{
			name:     "invalid base64",
			audio:    "not-valid-base64!!!",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := TtsEvent{Audio: tt.audio}
			result := event.AudioBytes()

			if tt.expected == nil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
				return
			}

			if string(result) != string(tt.expected) {
				t.Errorf("expected %s, got %s", string(tt.expected), string(result))
			}
		})
	}
}

func TestTtsClientEvent_Marshal(t *testing.T) {
	event := TtsClientEvent{
		Text:  "Hello, world!",
		Force: true,
		VoiceSettings: map[string]interface{}{
			"speed": 1.2,
		},
	}

	data, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	// Unmarshal back to verify
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("failed to unmarshal result: %v", err)
	}

	if result["text"] != "Hello, world!" {
		t.Errorf("expected text='Hello, world!', got %v", result["text"])
	}
	if result["force"] != true {
		t.Errorf("expected force=true, got %v", result["force"])
	}
	if settings, ok := result["voice_settings"].(map[string]interface{}); !ok || settings["speed"] != 1.2 {
		t.Errorf("expected voice_settings.speed=1.2, got %v", result["voice_settings"])
	}
}

func TestTextToSpeechWS_Integration(t *testing.T) {
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

		var err error
		serverConn, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("upgrade error: %v", err)
			return
		}
		defer serverConn.Close()

		// Handle client messages
		for {
			_, data, err := serverConn.ReadMessage()
			if err != nil {
				return
			}

			var clientEvent TtsClientEvent
			if err := json.Unmarshal(data, &clientEvent); err != nil {
				continue
			}

			// Skip handshake
			if strings.TrimSpace(clientEvent.Text) == "" {
				continue
			}

			// Send audio chunk response
			audioData := []byte("fake audio data")
			response := TtsEvent{
				Type:                    "audio_chunk",
				Audio:                   base64.StdEncoding.EncodeToString(audioData),
				Text:                    clientEvent.Text,
				Cached:                  false,
				TimeToFirstAudioFrameMs: 100,
			}
			respBytes, _ := json.Marshal(response)
			serverConn.WriteMessage(websocket.TextMessage, respBytes)

			// Send final event
			finalResponse := TtsEvent{
				Type:    "final",
				IsFinal: true,
			}
			finalBytes, _ := json.Marshal(finalResponse)
			serverConn.WriteMessage(websocket.TextMessage, finalBytes)
		}
	}))
	defer server.Close()

	// Convert to WebSocket URL
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Create client
	ws, err := NewTextToSpeechWS(TextToSpeechWSConfig{
		APIKey:  "test-api-key",
		BaseURL: wsURL,
	}, TextToSpeechStreamParams{
		Voice:        "telnyx.NaturalHD.Alloy",
		OutputFormat: "mp3",
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

	// Send initial handshake
	if err := ws.Send(TtsClientEvent{Text: " "}); err != nil {
		t.Fatalf("failed to send handshake: %v", err)
	}

	// Send text to synthesize
	if err := ws.Send(TtsClientEvent{Text: "Hello, world!"}); err != nil {
		t.Fatalf("failed to send text: %v", err)
	}

	// Receive audio chunk
	select {
	case event := <-ws.Events():
		if event.Type != "audio_chunk" {
			t.Errorf("expected type=audio_chunk, got %s", event.Type)
		}
		if event.Text != "Hello, world!" {
			t.Errorf("expected text='Hello, world!', got %s", event.Text)
		}
		audioBytes := event.AudioBytes()
		if string(audioBytes) != "fake audio data" {
			t.Errorf("expected audio='fake audio data', got %s", string(audioBytes))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for audio chunk")
	}

	// Receive final event
	select {
	case event := <-ws.Events():
		if event.Type != "final" {
			t.Errorf("expected type=final, got %s", event.Type)
		}
		if !event.IsFinal {
			t.Error("expected isFinal=true")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for final event")
	}
}

func TestTextToSpeechWS_Stream(t *testing.T) {
	// Create a test WebSocket server
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("upgrade error: %v", err)
			return
		}
		defer conn.Close()

		// Send an audio chunk
		response := TtsEvent{
			Type:  "audio_chunk",
			Audio: base64.StdEncoding.EncodeToString([]byte("audio")),
			Text:  "test",
		}
		respBytes, _ := json.Marshal(response)
		conn.WriteMessage(websocket.TextMessage, respBytes)

		// Wait for close
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				return
			}
		}
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	ws, err := NewTextToSpeechWS(TextToSpeechWSConfig{
		APIKey:  "test-api-key",
		BaseURL: wsURL,
	}, TextToSpeechStreamParams{})
	if err != nil {
		t.Fatalf("failed to create WebSocket: %v", err)
	}
	defer ws.Close()

	if err := ws.WaitForOpenWithTimeout(5 * time.Second); err != nil {
		t.Fatalf("failed to wait for open: %v", err)
	}

	// Use Stream interface
	streamCh := ws.Stream()

	// First message should be "open" (since we're already open)
	select {
	case msg := <-streamCh:
		if msg.Type != "open" {
			t.Errorf("expected type=open, got %s", msg.Type)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for open message")
	}

	// Next should be the audio chunk
	select {
	case msg := <-streamCh:
		if msg.Type != "message" {
			t.Errorf("expected type=message, got %s", msg.Type)
		}
		if msg.Message.Type != "audio_chunk" {
			t.Errorf("expected message type=audio_chunk, got %s", msg.Message.Type)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for audio message")
	}
}
