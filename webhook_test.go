// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestWebhookUnwrap(t *testing.T) {
	client := telnyx.NewClient(
		option.WithPublicKey("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
	)
	payload := []byte(`{"id":"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0","event_type":"conference.floor.changed","payload":{"call_control_id":"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg","call_leg_id":"428c31b6-7af4-4bcb-b7f5-5013ef9657c1","call_session_id":"428c31b6-7af4-4bcb-b7f5-5013ef9657c1","client_state":"aGF2ZSBhIG5pY2UgZGF5ID1d","conference_id":"428c31b6-abf3-3bc1-b7f4-5013ef9657c1","connection_id":"7267xxxxxxxxxxxxxx","occurred_at":"2018-02-02T22:25:27.521Z"},"record_type":"event"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Fatal("Failed to sign test webhook message", err)
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatal("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}
}
