// Package lib provides custom utilities for the Telnyx SDK.
//
// This file contains WebSocket infrastructure for streaming APIs.
package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketState represents the state of a WebSocket connection.
type WebSocketState int

const (
	// StateConnecting indicates the WebSocket is connecting.
	StateConnecting WebSocketState = iota
	// StateOpen indicates the WebSocket is open and ready.
	StateOpen
	// StateClosing indicates the WebSocket is closing.
	StateClosing
	// StateClosed indicates the WebSocket is closed.
	StateClosed
)

// WebSocketError represents an error that occurred during WebSocket operations.
type WebSocketError struct {
	Message string
	Cause   error
}

func (e *WebSocketError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *WebSocketError) Unwrap() error {
	return e.Cause
}

// BaseWebSocket provides common WebSocket functionality.
type BaseWebSocket struct {
	conn      *websocket.Conn
	state     WebSocketState
	stateMu   sync.RWMutex
	closeOnce sync.Once
	closeCh   chan struct{}
	openCh    chan struct{}
	openErr   error
	openOnce  sync.Once
}

// newBaseWebSocket creates a new BaseWebSocket.
func newBaseWebSocket() *BaseWebSocket {
	return &BaseWebSocket{
		state:   StateConnecting,
		closeCh: make(chan struct{}),
		openCh:  make(chan struct{}),
	}
}

// State returns the current WebSocket state.
func (ws *BaseWebSocket) State() WebSocketState {
	ws.stateMu.RLock()
	defer ws.stateMu.RUnlock()
	return ws.state
}

// setState sets the WebSocket state.
func (ws *BaseWebSocket) setState(state WebSocketState) {
	ws.stateMu.Lock()
	defer ws.stateMu.Unlock()
	ws.state = state
}

// IsOpen returns true if the WebSocket is open.
func (ws *BaseWebSocket) IsOpen() bool {
	return ws.State() == StateOpen
}

// markOpen marks the WebSocket as open.
func (ws *BaseWebSocket) markOpen() {
	ws.openOnce.Do(func() {
		ws.setState(StateOpen)
		close(ws.openCh)
	})
}

// markOpenError marks the WebSocket as failed to open.
func (ws *BaseWebSocket) markOpenError(err error) {
	ws.openOnce.Do(func() {
		ws.openErr = err
		ws.setState(StateClosed)
		close(ws.openCh)
	})
}

// WaitForOpen blocks until the WebSocket is open or returns an error if connection fails.
func (ws *BaseWebSocket) WaitForOpen() error {
	<-ws.openCh
	return ws.openErr
}

// WaitForOpenWithTimeout blocks until the WebSocket is open or timeout occurs.
func (ws *BaseWebSocket) WaitForOpenWithTimeout(timeout time.Duration) error {
	select {
	case <-ws.openCh:
		return ws.openErr
	case <-time.After(timeout):
		return &WebSocketError{Message: "connection timeout"}
	}
}

// buildWebSocketURL constructs a WebSocket URL from a base URL and path with query parameters.
func buildWebSocketURL(baseURL, path string, params url.Values) (*url.URL, error) {
	// Parse the base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	// Convert HTTP to WebSocket protocol
	switch u.Scheme {
	case "https":
		u.Scheme = "wss"
	case "http":
		u.Scheme = "ws"
	case "wss", "ws":
		// Already WebSocket protocol
	default:
		return nil, fmt.Errorf("unsupported URL scheme: %s", u.Scheme)
	}

	// Append path
	u.Path = strings.TrimSuffix(u.Path, "/") + path

	// Add query parameters
	if params != nil && len(params) > 0 {
		u.RawQuery = params.Encode()
	}

	return u, nil
}

// dial establishes a WebSocket connection.
func dial(wsURL *url.URL, apiKey string) (*websocket.Conn, *http.Response, error) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 30 * time.Second,
	}

	headers := http.Header{}
	if apiKey != "" {
		headers.Set("Authorization", "Bearer "+apiKey)
	}

	return dialer.Dial(wsURL.String(), headers)
}

// sendJSON sends a JSON message over the WebSocket.
func (ws *BaseWebSocket) sendJSON(v interface{}) error {
	if ws.conn == nil {
		return &WebSocketError{Message: "connection not established"}
	}

	data, err := json.Marshal(v)
	if err != nil {
		return &WebSocketError{Message: "failed to marshal JSON", Cause: err}
	}

	if err := ws.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return &WebSocketError{Message: "failed to send message", Cause: err}
	}

	return nil
}

// sendBinary sends binary data over the WebSocket.
func (ws *BaseWebSocket) sendBinary(data []byte) error {
	if ws.conn == nil {
		return &WebSocketError{Message: "connection not established"}
	}

	if err := ws.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		return &WebSocketError{Message: "failed to send binary data", Cause: err}
	}

	return nil
}

// close closes the WebSocket connection.
func (ws *BaseWebSocket) close(code int, reason string) error {
	var closeErr error
	ws.closeOnce.Do(func() {
		ws.setState(StateClosing)

		if ws.conn != nil {
			// Send close message
			message := websocket.FormatCloseMessage(code, reason)
			ws.conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))

			// Close the connection
			closeErr = ws.conn.Close()
		}

		ws.setState(StateClosed)
		close(ws.closeCh)
	})
	return closeErr
}

// ErrWebSocketClosed is returned when operations are attempted on a closed WebSocket.
var ErrWebSocketClosed = errors.New("websocket is closed")
