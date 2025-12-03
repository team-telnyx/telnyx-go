// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestAIChatNewCompletionWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AI.Chat.NewCompletion(context.TODO(), telnyx.AIChatNewCompletionParams{
		Messages: []telnyx.AIChatNewCompletionParamsMessage{{
			Content: telnyx.AIChatNewCompletionParamsMessageContentUnion{
				OfString: telnyx.String("You are a friendly chatbot."),
			},
			Role: "system",
		}, {
			Content: telnyx.AIChatNewCompletionParamsMessageContentUnion{
				OfString: telnyx.String("Hello, world!"),
			},
			Role: "user",
		}},
		APIKeyRef:        telnyx.String("api_key_ref"),
		BestOf:           telnyx.Int(0),
		EarlyStopping:    telnyx.Bool(true),
		FrequencyPenalty: telnyx.Float(0),
		GuidedChoice:     []string{"string"},
		GuidedJson: map[string]any{
			"foo": "bar",
		},
		GuidedRegex:     telnyx.String("guided_regex"),
		LengthPenalty:   telnyx.Float(0),
		Logprobs:        telnyx.Bool(true),
		MaxTokens:       telnyx.Int(0),
		MinP:            telnyx.Float(0),
		Model:           telnyx.String("model"),
		N:               telnyx.Float(0),
		PresencePenalty: telnyx.Float(0),
		ResponseFormat: telnyx.AIChatNewCompletionParamsResponseFormat{
			Type: "text",
		},
		Stream:      telnyx.Bool(true),
		Temperature: telnyx.Float(0),
		ToolChoice:  telnyx.AIChatNewCompletionParamsToolChoiceNone,
		Tools: []telnyx.AIChatNewCompletionParamsToolUnion{{
			OfFunction: &telnyx.AIChatNewCompletionParamsToolFunction{
				Function: telnyx.AIChatNewCompletionParamsToolFunctionFunction{
					Name:        "name",
					Description: telnyx.String("description"),
					Parameters: map[string]any{
						"foo": "bar",
					},
				},
			},
		}},
		TopLogprobs:   telnyx.Int(0),
		TopP:          telnyx.Float(0),
		UseBeamSearch: telnyx.Bool(true),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
