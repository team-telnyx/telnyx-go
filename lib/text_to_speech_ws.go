// Package lib provides custom utilities for the Telnyx SDK.
//
// This file contains WebSocket support for Text-to-Speech streaming.
package lib

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

// TextToSpeechStreamParams contains parameters for establishing a TTS WebSocket connection.
type TextToSpeechStreamParams struct {
	// Voice specifies the voice to use (e.g., "telnyx.NaturalHD.Alloy", "Telnyx.Ultra.<voice_id>").
	Voice string `json:"voice,omitempty"`

	// Provider specifies the TTS provider (e.g., "telnyx", "aws", "azure", "elevenlabs").
	Provider string `json:"provider,omitempty"`

	// OutputFormat specifies the audio output format (e.g., "mp3", "wav", "pcm").
	OutputFormat string `json:"output_format,omitempty"`

	// SampleRate specifies the audio sample rate in Hz.
	SampleRate int `json:"sample_rate,omitempty"`

	// Language specifies the language code (e.g., "en-US").
	Language string `json:"language,omitempty"`
}

// toURLValues converts the params to URL query values.
func (p TextToSpeechStreamParams) toURLValues() url.Values {
	v := url.Values{}

	if p.Voice != "" {
		v.Set("voice", p.Voice)
	}
	if p.Provider != "" {
		v.Set("provider", p.Provider)
	}
	if p.OutputFormat != "" {
		v.Set("output_format", p.OutputFormat)
	}
	if p.SampleRate > 0 {
		v.Set("sample_rate", strconv.Itoa(p.SampleRate))
	}
	if p.Language != "" {
		v.Set("language", p.Language)
	}

	return v
}

// TtsClientEvent represents a client-to-server frame for TTS.
type TtsClientEvent struct {
	// Text is the text to convert to speech. Send " " (single space) as an initial
	// handshake with optional VoiceSettings.
	Text string `json:"text"`

	// Force, when true, stops the current synthesis worker and starts a new one.
	// Used to interrupt speech mid-stream and begin synthesizing new text.
	Force bool `json:"force,omitempty"`

	// VoiceSettings contains provider-specific voice settings sent with the initial handshake.
	// Contents vary by provider — e.g. {"speed": 1.2} for Minimax, {"voice_speed": 1.5} for Telnyx.
	VoiceSettings map[string]interface{} `json:"voice_settings,omitempty"`
}

// TtsEvent represents an event received from the TTS WebSocket.
type TtsEvent struct {
	// Type is the event type: "audio_chunk", "final", or "error".
	Type string `json:"type"`

	// Audio is base64-encoded audio data (for "audio_chunk" events).
	Audio string `json:"audio,omitempty"`

	// Text is the text that was synthesized.
	Text string `json:"text,omitempty"`

	// IsFinal indicates if this is the final audio chunk.
	IsFinal bool `json:"isFinal,omitempty"`

	// Cached indicates if the audio was served from cache.
	Cached bool `json:"cached,omitempty"`

	// TimeToFirstAudioFrameMs is the time to first audio frame in milliseconds.
	TimeToFirstAudioFrameMs int `json:"timeToFirstAudioFrameMs,omitempty"`

	// Error is the error message (for "error" events).
	Error string `json:"error,omitempty"`
}

// AudioBytes returns the decoded audio bytes from the Audio field.
// Returns nil if the event has no audio or decoding fails.
func (e TtsEvent) AudioBytes() []byte {
	if e.Audio == "" {
		return nil
	}
	data, err := base64.StdEncoding.DecodeString(e.Audio)
	if err != nil {
		return nil
	}
	return data
}

// TtsStreamMessage represents a message from the TTS stream iterator.
type TtsStreamMessage struct {
	// Type is the message type: "connecting", "open", "closing", "close", "message", or "error".
	Type string

	// Message is the TTS event (for "message" type).
	Message *TtsEvent

	// Error is the error (for "error" type).
	Error error
}

// TextToSpeechWS provides WebSocket-based Text-to-Speech streaming.
type TextToSpeechWS struct {
	*BaseWebSocket

	params    TextToSpeechStreamParams
	apiKey    string
	baseURL   string
	eventsCh  chan TtsEvent
	streamCh  chan TtsStreamMessage
	errorsCh  chan error
	wg        sync.WaitGroup
	streamMu  sync.Mutex
	streaming bool
}

// TextToSpeechWSConfig contains configuration for the TTS WebSocket.
type TextToSpeechWSConfig struct {
	// APIKey is the Telnyx API key for authentication.
	APIKey string

	// BaseURL is the base API URL (defaults to "https://api.telnyx.com/v2").
	BaseURL string

	// EventBufferSize is the size of the events channel buffer (defaults to 100).
	EventBufferSize int
}

// NewTextToSpeechWS creates a new Text-to-Speech WebSocket client and connects.
//
// Example:
//
//	ws, err := lib.NewTextToSpeechWS(lib.TextToSpeechWSConfig{
//	    APIKey: os.Getenv("TELNYX_API_KEY"),
//	}, lib.TextToSpeechStreamParams{
//	    Voice:        "telnyx.NaturalHD.Alloy",
//	    OutputFormat: "mp3",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer ws.Close()
//
//	// Wait for connection
//	if err := ws.WaitForOpen(); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Send initial handshake
//	ws.Send(lib.TtsClientEvent{Text: " "})
//
//	// Send text to synthesize
//	ws.Send(lib.TtsClientEvent{Text: "Hello, world!"})
//
//	// Receive audio chunks
//	for event := range ws.Events() {
//	    if event.Type == "audio_chunk" {
//	        audioData := event.AudioBytes()
//	        // Process audio data...
//	    }
//	}
func NewTextToSpeechWS(config TextToSpeechWSConfig, params TextToSpeechStreamParams) (*TextToSpeechWS, error) {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.telnyx.com/v2"
	}

	bufferSize := config.EventBufferSize
	if bufferSize <= 0 {
		bufferSize = 100
	}

	ws := &TextToSpeechWS{
		BaseWebSocket: newBaseWebSocket(),
		params:        params,
		apiKey:        config.APIKey,
		baseURL:       baseURL,
		eventsCh:      make(chan TtsEvent, bufferSize),
		streamCh:      make(chan TtsStreamMessage, bufferSize),
		errorsCh:      make(chan error, 10),
	}

	// Build WebSocket URL
	wsURL, err := buildWebSocketURL(baseURL, "/text-to-speech/speech", params.toURLValues())
	if err != nil {
		return nil, &WebSocketError{Message: "failed to build WebSocket URL", Cause: err}
	}

	// Connect
	conn, _, err := dial(wsURL, config.APIKey)
	if err != nil {
		ws.markOpenError(err)
		return nil, &WebSocketError{Message: "failed to connect to WebSocket", Cause: err}
	}

	ws.conn = conn
	ws.markOpen()

	// Start message receiver
	ws.wg.Add(1)
	go ws.receiveMessages()

	return ws, nil
}

// Send sends a text event to the server for synthesis.
func (ws *TextToSpeechWS) Send(event TtsClientEvent) error {
	if !ws.IsOpen() {
		return ErrWebSocketClosed
	}
	return ws.sendJSON(event)
}

// Events returns a channel that receives TTS events.
// The channel is closed when the WebSocket connection closes.
func (ws *TextToSpeechWS) Events() <-chan TtsEvent {
	return ws.eventsCh
}

// Errors returns a channel that receives errors.
func (ws *TextToSpeechWS) Errors() <-chan error {
	return ws.errorsCh
}

// Stream returns a channel that provides an iterator-like interface over WebSocket
// lifecycle and message events. This provides an alternative to the event-based API.
// The iterator will exit if the socket closes, but breaking out of the iterator
// does not close the socket.
//
// Example:
//
//	for msg := range ws.Stream() {
//	    switch msg.Type {
//	    case "message":
//	        fmt.Println("received:", msg.Message)
//	    case "error":
//	        fmt.Println("error:", msg.Error)
//	    case "close":
//	        fmt.Println("connection closed")
//	    }
//	}
func (ws *TextToSpeechWS) Stream() <-chan TtsStreamMessage {
	ws.streamMu.Lock()
	defer ws.streamMu.Unlock()

	// Return existing stream if already streaming
	if ws.streaming {
		return ws.streamCh
	}

	ws.streaming = true

	// Send initial state based on connection state
	switch ws.State() {
	case StateConnecting:
		ws.streamCh <- TtsStreamMessage{Type: "connecting"}
	case StateOpen:
		ws.streamCh <- TtsStreamMessage{Type: "open"}
	case StateClosing:
		ws.streamCh <- TtsStreamMessage{Type: "closing"}
	case StateClosed:
		ws.streamCh <- TtsStreamMessage{Type: "close"}
	}

	return ws.streamCh
}

// Close closes the WebSocket connection gracefully.
func (ws *TextToSpeechWS) Close() error {
	err := ws.close(websocket.CloseNormalClosure, "OK")
	ws.wg.Wait()
	return err
}

// receiveMessages reads messages from the WebSocket and sends them to the channels.
func (ws *TextToSpeechWS) receiveMessages() {
	defer ws.wg.Done()
	defer func() {
		close(ws.eventsCh)
		ws.streamMu.Lock()
		if ws.streaming {
			ws.streamCh <- TtsStreamMessage{Type: "close"}
			close(ws.streamCh)
		}
		ws.streamMu.Unlock()
	}()

	for {
		select {
		case <-ws.closeCh:
			return
		default:
			messageType, message, err := ws.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					return
				}
				// Send error to channels (non-blocking)
				wsErr := &WebSocketError{Message: "failed to read message", Cause: err}
				select {
				case ws.errorsCh <- wsErr:
				default:
				}
				ws.streamMu.Lock()
				if ws.streaming {
					select {
					case ws.streamCh <- TtsStreamMessage{Type: "error", Error: wsErr}:
					default:
					}
				}
				ws.streamMu.Unlock()
				return
			}

			if messageType != websocket.TextMessage {
				continue
			}

			var event TtsEvent
			if err := json.Unmarshal(message, &event); err != nil {
				wsErr := &WebSocketError{Message: "failed to parse event", Cause: err}
				select {
				case ws.errorsCh <- wsErr:
				default:
				}
				ws.streamMu.Lock()
				if ws.streaming {
					select {
					case ws.streamCh <- TtsStreamMessage{Type: "error", Error: wsErr}:
					default:
					}
				}
				ws.streamMu.Unlock()
				continue
			}

			// Send to events channel
			select {
			case ws.eventsCh <- event:
			case <-ws.closeCh:
				return
			}

			// Send to stream channel if streaming
			ws.streamMu.Lock()
			if ws.streaming {
				if event.Type == "error" {
					select {
					case ws.streamCh <- TtsStreamMessage{Type: "error", Error: &WebSocketError{Message: event.Error}}:
					default:
					}
				} else {
					select {
					case ws.streamCh <- TtsStreamMessage{Type: "message", Message: &event}:
					default:
					}
				}
			}
			ws.streamMu.Unlock()
		}
	}
}
