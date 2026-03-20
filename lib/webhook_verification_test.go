package lib_test

import (
	"crypto/ed25519"
	"encoding/base64"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/lib"
)

func TestVerifyWebhookSignature_ValidSignature(t *testing.T) {
	// Generate ED25519 keypair for testing
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal("Failed to generate keypair:", err)
	}
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	payload := []byte(`{"data":{"id":"test-id","event_type":"call.initiated"}}`)

	// Create timestamp and sign the payload
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
	signedPayload := []byte(timestamp + "|" + string(payload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, signatureB64)
	headers.Set(lib.TimestampHeader, timestamp)

	err = lib.VerifyWebhookSignature(payload, headers, publicKeyB64)
	if err != nil {
		t.Errorf("Expected successful verification, got error: %v", err)
	}
}

func TestVerifyWebhookSignature_InvalidSignature(t *testing.T) {
	publicKey, _, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatal("Failed to generate keypair:", err)
	}
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	payload := []byte(`{"data":{}}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Use a fake signature (64 zero bytes)
	fakeSignature := base64.StdEncoding.EncodeToString(make([]byte, 64))

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, fakeSignature)
	headers.Set(lib.TimestampHeader, timestamp)

	err = lib.VerifyWebhookSignature(payload, headers, publicKeyB64)
	if err == nil {
		t.Error("Expected error for invalid signature, got nil")
	}

	if !lib.IsWebhookVerificationError(err) {
		t.Errorf("Expected WebhookVerificationError, got %T", err)
	}
}

func TestVerifyWebhookSignature_MissingSignatureHeader(t *testing.T) {
	publicKey, _, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	payload := []byte(`{}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	headers := make(http.Header)
	headers.Set(lib.TimestampHeader, timestamp)
	// Missing signature header

	err := lib.VerifyWebhookSignature(payload, headers, publicKeyB64)
	if err == nil {
		t.Error("Expected error for missing signature header")
	}
}

func TestVerifyWebhookSignature_MissingTimestampHeader(t *testing.T) {
	publicKey, _, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	payload := []byte(`{}`)

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, base64.StdEncoding.EncodeToString(make([]byte, 64)))
	// Missing timestamp header

	err := lib.VerifyWebhookSignature(payload, headers, publicKeyB64)
	if err == nil {
		t.Error("Expected error for missing timestamp header")
	}
}

func TestVerifyWebhookSignature_ExpiredTimestamp(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	payload := []byte(`{}`)

	// Timestamp from 10 minutes ago (beyond 5 min tolerance)
	oldTime := time.Now().Add(-10 * time.Minute)
	timestamp := strconv.FormatInt(oldTime.Unix(), 10)

	signedPayload := []byte(timestamp + "|" + string(payload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, signatureB64)
	headers.Set(lib.TimestampHeader, timestamp)

	err := lib.VerifyWebhookSignature(payload, headers, publicKeyB64)
	if err == nil {
		t.Error("Expected error for expired timestamp")
	}
}

func TestVerifyWebhookSignature_FutureTimestamp(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	payload := []byte(`{}`)

	// Timestamp 10 minutes in the future
	futureTime := time.Now().Add(10 * time.Minute)
	timestamp := strconv.FormatInt(futureTime.Unix(), 10)

	signedPayload := []byte(timestamp + "|" + string(payload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, signatureB64)
	headers.Set(lib.TimestampHeader, timestamp)

	err := lib.VerifyWebhookSignature(payload, headers, publicKeyB64)
	if err == nil {
		t.Error("Expected error for future timestamp")
	}
}

func TestVerifyWebhookSignature_TamperedPayload(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	// Sign the original payload
	originalPayload := []byte(`{"data":{"id":"original"}}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signedPayload := []byte(timestamp + "|" + string(originalPayload))
	signature := ed25519.Sign(privateKey, signedPayload)
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	// But verify with a tampered payload
	tamperedPayload := []byte(`{"data":{"id":"TAMPERED"}}`)

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, signatureB64)
	headers.Set(lib.TimestampHeader, timestamp)

	err := lib.VerifyWebhookSignature(tamperedPayload, headers, publicKeyB64)
	if err == nil {
		t.Error("Expected error for tampered payload")
	}
}

func TestVerifyWebhookSignature_InvalidPublicKey(t *testing.T) {
	payload := []byte(`{}`)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	headers := make(http.Header)
	headers.Set(lib.SignatureHeader, base64.StdEncoding.EncodeToString(make([]byte, 64)))
	headers.Set(lib.TimestampHeader, timestamp)

	// Invalid public key (not base64)
	err := lib.VerifyWebhookSignature(payload, headers, "not-valid-base64!!!")
	if err == nil {
		t.Error("Expected error for invalid public key encoding")
	}

	// Invalid public key size (too short)
	shortKey := base64.StdEncoding.EncodeToString([]byte("too-short"))
	err = lib.VerifyWebhookSignature(payload, headers, shortKey)
	if err == nil {
		t.Error("Expected error for invalid public key size")
	}
}
