// Package lib provides custom utilities for the Telnyx SDK.
//
// This file contains ED25519 webhook signature verification logic.
// It is kept separate from generated code to avoid merge conflicts
// when Stainless updates webhook event types.
package lib

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	// SignatureHeader is the HTTP header containing the ED25519 signature.
	SignatureHeader = "telnyx-signature-ed25519"

	// TimestampHeader is the HTTP header containing the Unix timestamp.
	TimestampHeader = "telnyx-timestamp"

	// TimestampTolerance is the maximum allowed time difference for webhook timestamps.
	// Webhooks older than 5 minutes or more than 5 minutes in the future are rejected.
	TimestampTolerance = 5 * time.Minute
)

// WebhookVerificationError represents an error during webhook verification.
type WebhookVerificationError struct {
	Message string
}

func (e *WebhookVerificationError) Error() string {
	return e.Message
}

// VerifyWebhookSignature verifies the ED25519 signature of a Telnyx webhook.
//
// Telnyx webhooks are signed using ED25519 with the following format:
//   - Header "Telnyx-Signature-Ed25519": Base64-encoded ED25519 signature (64 bytes)
//   - Header "Telnyx-Timestamp": Unix timestamp in seconds
//   - Signed payload: "{timestamp}|{payload}"
//
// Parameters:
//   - payload: The raw webhook body as bytes
//   - headers: HTTP headers from the webhook request
//   - publicKeyB64: Base64-encoded ED25519 public key from Telnyx Mission Control
//
// Returns nil if verification succeeds, or an error describing why it failed.
//
// Example:
//
//	err := lib.VerifyWebhookSignature(body, req.Header, os.Getenv("TELNYX_PUBLIC_KEY"))
//	if err != nil {
//	    // Reject the webhook
//	}
func VerifyWebhookSignature(payload []byte, headers http.Header, publicKeyB64 string) error {
	// Get required headers (http.Header.Get is case-insensitive)
	signature := headers.Get(SignatureHeader)
	timestamp := headers.Get(TimestampHeader)

	if signature == "" {
		return &WebhookVerificationError{Message: "missing required header: " + SignatureHeader}
	}

	if timestamp == "" {
		return &WebhookVerificationError{Message: "missing required header: " + TimestampHeader}
	}

	// Validate timestamp to prevent replay attacks
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return &WebhookVerificationError{Message: fmt.Sprintf("invalid timestamp format: %s", timestamp)}
	}

	webhookTime := time.Unix(ts, 0)
	now := time.Now()
	timeDiff := now.Sub(webhookTime)

	if timeDiff > TimestampTolerance {
		return &WebhookVerificationError{Message: fmt.Sprintf("webhook timestamp too old: %s difference", timeDiff)}
	}

	if timeDiff < -TimestampTolerance {
		return &WebhookVerificationError{Message: fmt.Sprintf("webhook timestamp too far in the future: %s difference", -timeDiff)}
	}

	// Decode the base64-encoded public key
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyB64)
	if err != nil {
		return &WebhookVerificationError{Message: fmt.Sprintf("invalid public key encoding: %v", err)}
	}

	if len(publicKeyBytes) != ed25519.PublicKeySize {
		return &WebhookVerificationError{
			Message: fmt.Sprintf("invalid public key size: expected %d bytes, got %d", ed25519.PublicKeySize, len(publicKeyBytes)),
		}
	}

	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return &WebhookVerificationError{Message: fmt.Sprintf("invalid signature encoding: %v", err)}
	}

	if len(signatureBytes) != ed25519.SignatureSize {
		return &WebhookVerificationError{
			Message: fmt.Sprintf("invalid signature size: expected %d bytes, got %d", ed25519.SignatureSize, len(signatureBytes)),
		}
	}

	// Build the signed payload: "{timestamp}|{payload}"
	signedPayload := []byte(timestamp + "|" + string(payload))

	// Verify the ED25519 signature
	if !ed25519.Verify(publicKeyBytes, signedPayload, signatureBytes) {
		return &WebhookVerificationError{Message: "signature verification failed: signature does not match payload"}
	}

	return nil
}

// MustVerifyWebhookSignature is like VerifyWebhookSignature but panics on error.
// Use this only in contexts where panicking is acceptable (e.g., tests).
func MustVerifyWebhookSignature(payload []byte, headers http.Header, publicKeyB64 string) {
	if err := VerifyWebhookSignature(payload, headers, publicKeyB64); err != nil {
		panic(err)
	}
}

// IsWebhookVerificationError checks if an error is a WebhookVerificationError.
func IsWebhookVerificationError(err error) bool {
	var verifyErr *WebhookVerificationError
	return errors.As(err, &verifyErr)
}
