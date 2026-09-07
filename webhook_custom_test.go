package telnyx_test

import (
	"crypto/ed25519"
	"encoding/base64"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

const webhookPayload = "{\n  \"data\": {\"id\": \"event-123\", \"event_type\": \"call.initiated\", \"record_type\": \"event\", \"payload\": {}}\n}"

type signedWebhook struct {
	payload   []byte
	headers   http.Header
	publicKey string
}

func newSignedWebhook(t *testing.T, payload []byte, timestamp string) signedWebhook {
	t.Helper()

	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("generate Ed25519 key: %v", err)
	}

	signedPayload := append([]byte(timestamp+"|"), payload...)
	signature := ed25519.Sign(privateKey, signedPayload)
	headers := make(http.Header)
	headers.Set("TeLnYx-SiGnAtUrE-Ed25519", base64.StdEncoding.EncodeToString(signature))
	headers.Set("tElNyX-TiMeStAmP", timestamp)

	return signedWebhook{
		payload:   payload,
		headers:   headers,
		publicKey: base64.StdEncoding.EncodeToString(publicKey),
	}
}

func currentSignedWebhook(t *testing.T) signedWebhook {
	t.Helper()
	return newSignedWebhook(t, []byte(webhookPayload), strconv.FormatInt(time.Now().Unix(), 10))
}

func requireWebhookError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected webhook verification to fail")
	}
}

func TestWebhookUnwrapUsesPublicClientAndReturnsTypedEvent(t *testing.T) {
	signed := currentSignedWebhook(t)
	client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))

	event, err := client.Webhooks.Unwrap(signed.payload, signed.headers)
	if err != nil {
		t.Fatalf("unwrap valid webhook: %v", err)
	}
	if got := event.AsCallInitiatedEvent().Data.EventType; got != telnyx.CallInitiatedEventTypeCallInitiated {
		t.Fatalf("event type = %q, want %q", got, telnyx.CallInitiatedEventTypeCallInitiated)
	}
}

func TestWebhookUnwrapSupportsPerCallKeyOverride(t *testing.T) {
	signed := currentSignedWebhook(t)
	_, wrongPrivateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("generate wrong key: %v", err)
	}
	wrongKey := base64.StdEncoding.EncodeToString(wrongPrivateKey.Public().(ed25519.PublicKey))
	client := telnyx.NewClient(option.WithPublicKey(wrongKey))

	if _, err := client.Webhooks.Unwrap(signed.payload, signed.headers); err == nil {
		t.Fatal("expected client-level wrong key to fail")
	}
	if _, err := client.Webhooks.Unwrap(signed.payload, signed.headers, option.WithPublicKey(signed.publicKey)); err != nil {
		t.Fatalf("per-call key override failed: %v", err)
	}
}

func TestWebhookUnwrapFailsClosedWithoutVerificationMaterial(t *testing.T) {
	signed := currentSignedWebhook(t)

	t.Run("public key", func(t *testing.T) {
		t.Setenv("TELNYX_PUBLIC_KEY", "")
		client := telnyx.NewClient()
		_, err := client.Webhooks.Unwrap(signed.payload, signed.headers)
		requireWebhookError(t, err)
	})

	t.Run("signature header", func(t *testing.T) {
		headers := signed.headers.Clone()
		headers.Del("telnyx-signature-ed25519")
		client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))
		_, err := client.Webhooks.Unwrap(signed.payload, headers)
		requireWebhookError(t, err)
	})

	t.Run("timestamp header", func(t *testing.T) {
		headers := signed.headers.Clone()
		headers.Del("telnyx-timestamp")
		client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))
		_, err := client.Webhooks.Unwrap(signed.payload, headers)
		requireWebhookError(t, err)
	})

	t.Run("malformed timestamp", func(t *testing.T) {
		headers := signed.headers.Clone()
		headers.Set("telnyx-timestamp", "not-a-timestamp")
		client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))
		_, err := client.Webhooks.Unwrap(signed.payload, headers)
		requireWebhookError(t, err)
	})
}

func TestWebhookUnwrapRejectsInvalidCryptographicInput(t *testing.T) {
	signed := currentSignedWebhook(t)
	client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))

	tests := map[string]struct {
		payload []byte
		headers http.Header
		key     string
	}{
		"tampered payload": {
			payload: append(append([]byte(nil), signed.payload...), ' '),
			headers: signed.headers,
			key:     signed.publicKey,
		},
		"malformed signature": {
			payload: signed.payload,
			headers: func() http.Header {
				headers := signed.headers.Clone()
				headers.Set("telnyx-signature-ed25519", "not-base64")
				return headers
			}(),
			key: signed.publicKey,
		},
		"malformed public key": {
			payload: signed.payload,
			headers: signed.headers,
			key:     "not-base64",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := client.Webhooks.Unwrap(test.payload, test.headers, option.WithPublicKey(test.key))
			requireWebhookError(t, err)
		})
	}
}

func TestWebhookUnwrapRejectsStaleAndFutureTimestamps(t *testing.T) {
	for name, timestamp := range map[string]string{
		"stale":  strconv.FormatInt(time.Now().Add(-10*time.Minute).Unix(), 10),
		"future": "9999999999",
	} {
		t.Run(name, func(t *testing.T) {
			signed := newSignedWebhook(t, []byte(webhookPayload), timestamp)
			client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))
			_, err := client.Webhooks.Unwrap(signed.payload, signed.headers)
			requireWebhookError(t, err)
		})
	}
}

func TestWebhookUnwrapVerifiesBeforeParsing(t *testing.T) {
	signed := currentSignedWebhook(t)
	malformedPayload := []byte("{")

	client := telnyx.NewClient(option.WithPublicKey(signed.publicKey))
	_, err := client.Webhooks.Unwrap(malformedPayload, signed.headers)
	requireWebhookError(t, err)
	if !strings.Contains(strings.ToLower(err.Error()), "signature") {
		t.Fatalf("error = %q, want signature verification failure before JSON parsing", err)
	}
}

func TestWebhookUnsafeUnwrapIsExplicitlyUnverified(t *testing.T) {
	client := telnyx.NewClient()
	event, err := client.Webhooks.UnsafeUnwrap([]byte(webhookPayload))
	if err != nil {
		t.Fatalf("unsafe unwrap: %v", err)
	}
	if got := event.AsCallInitiatedEvent().Data.EventType; got != telnyx.CallInitiatedEventTypeCallInitiated {
		t.Fatalf("event type = %q, want %q", got, telnyx.CallInitiatedEventTypeCallInitiated)
	}
}
