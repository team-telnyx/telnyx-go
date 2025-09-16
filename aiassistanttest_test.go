// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go"
	"github.com/stainless-sdks/telnyx-go/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/option"
)

func TestAIAssistantTestNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Tests.New(context.TODO(), telnyx.AIAssistantTestNewParams{
		Destination:  "x",
		Instructions: "x",
		Name:         "x",
		Rubric: []telnyx.AIAssistantTestNewParamsRubric{{
			Criteria: "criteria",
			Name:     "name",
		}},
		Description:               telnyx.String("description"),
		MaxDurationSeconds:        telnyx.Int(1),
		TelnyxConversationChannel: telnyx.TelnyxConversationChannelPhoneCall,
		TestSuite:                 telnyx.String("test_suite"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantTestGet(t *testing.T) {
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
	_, err := client.AI.Assistants.Tests.Get(context.TODO(), "test_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantTestUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Tests.Update(
		context.TODO(),
		"test_id",
		telnyx.AIAssistantTestUpdateParams{
			Description:        telnyx.String("description"),
			Destination:        telnyx.String("x"),
			Instructions:       telnyx.String("x"),
			MaxDurationSeconds: telnyx.Int(1),
			Name:               telnyx.String("x"),
			Rubric: []telnyx.AIAssistantTestUpdateParamsRubric{{
				Criteria: "criteria",
				Name:     "name",
			}},
			TelnyxConversationChannel: telnyx.TelnyxConversationChannelPhoneCall,
			TestSuite:                 telnyx.String("test_suite"),
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

func TestAIAssistantTestListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Tests.List(context.TODO(), telnyx.AIAssistantTestListParams{
		Destination: telnyx.String("destination"),
		Page: telnyx.AIAssistantTestListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		TelnyxConversationChannel: telnyx.String("telnyx_conversation_channel"),
		TestSuite:                 telnyx.String("test_suite"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantTestDelete(t *testing.T) {
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
	_, err := client.AI.Assistants.Tests.Delete(context.TODO(), "test_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
