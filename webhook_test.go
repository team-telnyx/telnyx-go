// Tests for Telnyx webhook ED25519 signature verification.

package telnyx_test

import (
	"crypto/ed25519"
	"encoding/base64"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestWebhookUnwrap(t *testing.T) {
	// Generate ED25519 keypair for testing
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal("Failed to generate keypair:", err)
	}
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	client := telnyx.NewClient(
		option.WithPublicKey(publicKeyB64),
		option.WithAPIKey("My API Key"),
	)

	payload := []byte(`{"id":"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0","event_type":"conference.floor.changed","payload":{"call_control_id":"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg","call_leg_id":"428c31b6-7af4-4bcb-b7f5-5013ef9657c1","call_session_id":"428c31b6-7af4-4bcb-b7f5-5013ef9657c1","client_state":"aGF2ZSBhIG5pY2UgZGF5ID1d","conference_id":"428c31b6-abf3-3bc1-b7f4-5013ef9657c1","connection_id":"7267xxxxxxxxxxxxxx","occurred_at":"2018-02-02T22:25:27.521Z"},"record_type":"event"}`)

	// Create timestamp and sign the payload using ED25519 (Telnyx format)
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)

	// Signed payload format: "{timestamp}|{payload}"
	signedPayload := []byte(timestamp + "|" + string(payload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	// Set headers as Telnyx would send them
	headers := make(http.Header)
	headers.Set("telnyx-signature-ed25519", signatureB64)
	headers.Set("telnyx-timestamp", timestamp)

	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}
}

func TestWebhookUnwrap_InvalidSignature(t *testing.T) {
	// Generate keypair
	publicKey, _, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal("Failed to generate keypair:", err)
	}
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	client := telnyx.NewClient(
		option.WithPublicKey(publicKeyB64),
		option.WithAPIKey("My API Key"),
	)

	payload := []byte(`{"data":{"id":"test"}}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Use a fake signature (64 zero bytes)
	fakeSignature := base64.StdEncoding.EncodeToString(make([]byte, 64))

	headers := make(http.Header)
	headers.Set("telnyx-signature-ed25519", fakeSignature)
	headers.Set("telnyx-timestamp", timestamp)

	_, err = client.Webhooks.Unwrap(payload, headers)
	if err == nil {
		t.Error("Expected error for invalid signature, got nil")
	}
}

func TestWebhookUnwrap_TamperedPayload(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	client := telnyx.NewClient(
		option.WithPublicKey(publicKeyB64),
		option.WithAPIKey("My API Key"),
	)

	// Sign the original payload
	originalPayload := []byte(`{"data":{"id":"original"}}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signedPayload := []byte(timestamp + "|" + string(originalPayload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	// But send a tampered payload
	tamperedPayload := []byte(`{"data":{"id":"TAMPERED"}}`)

	headers := make(http.Header)
	headers.Set("telnyx-signature-ed25519", signatureB64)
	headers.Set("telnyx-timestamp", timestamp)

	_, err := client.Webhooks.Unwrap(tamperedPayload, headers)
	if err == nil {
		t.Error("Expected error for tampered payload, got nil")
	}
}

func TestWebhookUnwrap_MissingHeaders(t *testing.T) {
	publicKey, _, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	client := telnyx.NewClient(
		option.WithPublicKey(publicKeyB64),
		option.WithAPIKey("My API Key"),
	)

	payload := []byte(`{"data":{}}`)

	// Missing timestamp header
	headers := make(http.Header)
	headers.Set("telnyx-signature-ed25519", "somesignature")

	_, err := client.Webhooks.Unwrap(payload, headers)
	if err == nil {
		t.Error("Expected error for missing headers, got nil")
	}
}

func TestWebhookUnwrap_ExpiredTimestamp(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	client := telnyx.NewClient(
		option.WithPublicKey(publicKeyB64),
		option.WithAPIKey("My API Key"),
	)

	payload := []byte(`{"data":{}}`)

	// Timestamp from 10 minutes ago (beyond 5 min tolerance)
	oldTime := time.Now().Add(-10 * time.Minute)
	timestamp := strconv.FormatInt(oldTime.Unix(), 10)

	signedPayload := []byte(timestamp + "|" + string(payload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	headers := make(http.Header)
	headers.Set("telnyx-signature-ed25519", signatureB64)
	headers.Set("telnyx-timestamp", timestamp)

	_, err := client.Webhooks.Unwrap(payload, headers)
	if err == nil {
		t.Error("Expected error for expired timestamp, got nil")
	}
}
