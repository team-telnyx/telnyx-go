// Package lib provides custom utilities for the Telnyx SDK.
//
// This file contains WebSocket support for Speech-to-Text streaming.
package lib

import (
	"encoding/json"
	"net/url"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

// SpeechToTextStreamParams contains parameters for establishing a Speech-to-Text WebSocket connection.
type SpeechToTextStreamParams struct {
	// TranscriptionEngine specifies the STT engine to use (e.g., "Deepgram", "Telnyx").
	TranscriptionEngine string `json:"transcription_engine,omitempty"`

	// InputFormat specifies the audio format (e.g., "wav", "mp3", "raw").
	InputFormat string `json:"input_format,omitempty"`

	// Language specifies the language code (e.g., "en-US", "es-ES").
	Language string `json:"language,omitempty"`

	// SampleRate specifies the audio sample rate in Hz (e.g., 8000, 16000, 44100).
	SampleRate int `json:"sample_rate,omitempty"`

	// Channels specifies the number of audio channels (1 for mono, 2 for stereo).
	Channels int `json:"channels,omitempty"`

	// InterimResults enables interim (non-final) transcription results.
	InterimResults bool `json:"interim_results,omitempty"`

	// Model specifies the transcription model to use.
	Model string `json:"model,omitempty"`

	// Encoding specifies the audio encoding (e.g., "linear16", "mulaw").
	Encoding string `json:"encoding,omitempty"`
}

// toURLValues converts the params to URL query values.
func (p SpeechToTextStreamParams) toURLValues() url.Values {
	v := url.Values{}

	if p.TranscriptionEngine != "" {
		v.Set("transcription_engine", p.TranscriptionEngine)
	}
	if p.InputFormat != "" {
		v.Set("input_format", p.InputFormat)
	}
	if p.Language != "" {
		v.Set("language", p.Language)
	}
	if p.SampleRate > 0 {
		v.Set("sample_rate", strconv.Itoa(p.SampleRate))
	}
	if p.Channels > 0 {
		v.Set("channels", strconv.Itoa(p.Channels))
	}
	if p.InterimResults {
		v.Set("interim_results", "true")
	}
	if p.Model != "" {
		v.Set("model", p.Model)
	}
	if p.Encoding != "" {
		v.Set("encoding", p.Encoding)
	}

	return v
}

// SttEvent represents an event received from the STT WebSocket.
type SttEvent struct {
	// Type is the event type: "transcript" or "error".
	Type string `json:"type"`

	// Transcript is the transcribed text (for "transcript" events).
	Transcript string `json:"transcript,omitempty"`

	// IsFinal indicates if this is a final transcription result.
	IsFinal bool `json:"is_final,omitempty"`

	// Confidence is the confidence score (0-1) for the transcription.
	Confidence float64 `json:"confidence,omitempty"`

	// Error is the error message (for "error" events).
	Error string `json:"error,omitempty"`

	// Words contains word-level timing information if available.
	Words []SttWord `json:"words,omitempty"`

	// Channel indicates which audio channel the transcript is from.
	Channel int `json:"channel,omitempty"`

	// StartTime is the start time of the transcript in seconds.
	StartTime float64 `json:"start_time,omitempty"`

	// Duration is the duration of the transcript in seconds.
	Duration float64 `json:"duration,omitempty"`
}

// SttWord represents word-level timing information.
type SttWord struct {
	Word       string  `json:"word"`
	Start      float64 `json:"start"`
	End        float64 `json:"end"`
	Confidence float64 `json:"confidence,omitempty"`
}

// SpeechToTextWS provides WebSocket-based Speech-to-Text streaming.
type SpeechToTextWS struct {
	*BaseWebSocket

	params   SpeechToTextStreamParams
	apiKey   string
	baseURL  string
	eventsCh chan SttEvent
	errorsCh chan error
	wg       sync.WaitGroup
}

// SpeechToTextWSConfig contains configuration for the STT WebSocket.
type SpeechToTextWSConfig struct {
	// APIKey is the Telnyx API key for authentication.
	APIKey string

	// BaseURL is the base API URL (defaults to "https://api.telnyx.com/v2").
	BaseURL string

	// EventBufferSize is the size of the events channel buffer (defaults to 100).
	EventBufferSize int
}

// NewSpeechToTextWS creates a new Speech-to-Text WebSocket client and connects.
//
// Example:
//
//	ws, err := lib.NewSpeechToTextWS(lib.SpeechToTextWSConfig{
//	    APIKey: os.Getenv("TELNYX_API_KEY"),
//	}, lib.SpeechToTextStreamParams{
//	    TranscriptionEngine: "Deepgram",
//	    InputFormat:         "wav",
//	    Language:            "en-US",
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
//	// Send audio data
//	ws.Send(audioBytes)
//
//	// Receive transcripts
//	for event := range ws.Events() {
//	    fmt.Printf("Transcript: %s (final: %v)\n", event.Transcript, event.IsFinal)
//	}
func NewSpeechToTextWS(config SpeechToTextWSConfig, params SpeechToTextStreamParams) (*SpeechToTextWS, error) {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.telnyx.com/v2"
	}

	bufferSize := config.EventBufferSize
	if bufferSize <= 0 {
		bufferSize = 100
	}

	ws := &SpeechToTextWS{
		BaseWebSocket: newBaseWebSocket(),
		params:        params,
		apiKey:        config.APIKey,
		baseURL:       baseURL,
		eventsCh:      make(chan SttEvent, bufferSize),
		errorsCh:      make(chan error, 10),
	}

	// Build WebSocket URL
	wsURL, err := buildWebSocketURL(baseURL, "/speech-to-text/transcription", params.toURLValues())
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

// Send sends binary audio data to the server for transcription.
func (ws *SpeechToTextWS) Send(data []byte) error {
	if !ws.IsOpen() {
		return ErrWebSocketClosed
	}
	return ws.sendBinary(data)
}

// Events returns a channel that receives STT events.
// The channel is closed when the WebSocket connection closes.
func (ws *SpeechToTextWS) Events() <-chan SttEvent {
	return ws.eventsCh
}

// Errors returns a channel that receives errors.
func (ws *SpeechToTextWS) Errors() <-chan error {
	return ws.errorsCh
}

// Close closes the WebSocket connection gracefully.
func (ws *SpeechToTextWS) Close() error {
	err := ws.close(websocket.CloseNormalClosure, "OK")
	ws.wg.Wait()
	return err
}

// receiveMessages reads messages from the WebSocket and sends them to the events channel.
func (ws *SpeechToTextWS) receiveMessages() {
	defer ws.wg.Done()
	defer close(ws.eventsCh)

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
				// Send error to errors channel (non-blocking)
				select {
				case ws.errorsCh <- &WebSocketError{Message: "failed to read message", Cause: err}:
				default:
				}
				return
			}

			if messageType != websocket.TextMessage {
				continue
			}

			var event SttEvent
			if err := json.Unmarshal(message, &event); err != nil {
				select {
				case ws.errorsCh <- &WebSocketError{Message: "failed to parse event", Cause: err}:
				default:
				}
				continue
			}

			select {
			case ws.eventsCh <- event:
			case <-ws.closeCh:
				return
			}
		}
	}
}
