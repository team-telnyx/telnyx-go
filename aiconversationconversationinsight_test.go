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

func TestAIConversationConversationInsightGetAggregatesWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Conversations.ConversationInsights.GetAggregates(context.TODO(), telnyx.AIConversationConversationInsightGetAggregatesParams{
		CreatedAt: telnyx.String("created_at"),
		GroupBy:   []string{"string"},
		InsightID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Metadata: telnyx.AIConversationConversationInsightGetAggregatesParamsMetadata{
			AssistantID: telnyx.String("assistant_id"),
		},
		Show: []string{"string"},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
