// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestAIAnthropicV1MessagesWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.AI.Anthropic.V1.Messages(context.TODO(), telnyx.AIAnthropicV1MessagesParams{
		MaxTokens: 1024,
		Messages: []map[string]any{{
			"role":    "bar",
			"content": "bar",
		}},
		Model:          "zai-org/GLM-5.2",
		APIKeyRef:      telnyx.String("api_key_ref"),
		BillingGroupID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FallbackConfig: map[string]any{
			"foo": "bar",
		},
		MaxRetries: telnyx.Int(0),
		McpServers: []map[string]any{{
			"foo": "bar",
		}},
		Metadata: map[string]any{
			"foo": "bar",
		},
		ServiceTier:   telnyx.String("service_tier"),
		StopSequences: []string{"string"},
		Stream:        telnyx.Bool(true),
		System: telnyx.AIAnthropicV1MessagesParamsSystemUnion{
			OfString: telnyx.String("You are a friendly chatbot."),
		},
		Temperature: telnyx.Float(0),
		Thinking: map[string]any{
			"foo": "bar",
		},
		Timeout: telnyx.Float(0),
		ToolChoice: map[string]any{
			"foo": "bar",
		},
		Tools: []map[string]any{{
			"foo": "bar",
		}},
		TopK: telnyx.Int(0),
		TopP: telnyx.Float(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
