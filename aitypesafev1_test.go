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

func TestAITypesafeV1Systemone(t *testing.T) {
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
	_, err := client.AI.Typesafe.V1.Systemone(context.TODO(), telnyx.AITypesafeV1SystemoneParams{
		Questions: map[string]telnyx.AITypesafeV1SystemoneParamsQuestionsUnion{
			"team": {
				OfChoice: &telnyx.AITypesafeV1SystemoneParamsQuestionsChoice{
					Criteria: map[string]string{
						"billing":           "Payments and refunds",
						"technical_support": "Service faults and technical problems",
						"sales":             "New purchases",
					},
					Instructions: telnyx.AITypesafeV1SystemoneParamsQuestionsChoiceInstructionsUnion{
						OfString: telnyx.String("Choose the team that should handle this incident."),
					},
				},
			},
			"production_incident": {
				OfNoul: &telnyx.AITypesafeV1SystemoneParamsQuestionsNoul{
					Instructions: telnyx.AITypesafeV1SystemoneParamsQuestionsNoulInstructionsUnion{
						OfString: telnyx.String("Does the message describe an active production incident?"),
					},
					Criteria: telnyx.AITypesafeV1SystemoneParamsQuestionsNoulCriteria{
						False: telnyx.String("false"),
						True:  telnyx.String("true"),
					},
				},
			},
			"urgency": {
				OfScore: &telnyx.AITypesafeV1SystemoneParamsQuestionsScore{
					Criteria: []string{"Low", "Normal", "High", "Critical"},
					Instructions: telnyx.AITypesafeV1SystemoneParamsQuestionsScoreInstructionsUnion{
						OfString: telnyx.String("Rate operational urgency."),
					},
				},
			},
		},
		State: telnyx.AITypesafeV1SystemoneParamsStateUnion{
			OfString: telnyx.String("Our production calls are failing. Every customer is affected."),
		},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
