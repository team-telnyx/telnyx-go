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

func TestAIKnowledgeCollectionGetDocumentsWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Knowledge.Collections.GetDocuments(
		context.TODO(),
		"support-transcripts",
		telnyx.AIKnowledgeCollectionGetDocumentsParams{
			Filter: map[string]any{
				"foo": "bar",
			},
			PageNumber:    telnyx.Int(1),
			PageSize:      telnyx.Int(20),
			Query:         telnyx.String("customer called about billing issue"),
			RetrievalType: telnyx.AIKnowledgeCollectionGetDocumentsParamsRetrievalTypeVector,
			Sources:       telnyx.String("voice,message"),
			TopK:          telnyx.Int(10),
		},
	)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
