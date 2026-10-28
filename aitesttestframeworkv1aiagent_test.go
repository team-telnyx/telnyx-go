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

func TestAITestTestFrameworkV1AiagentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Aiagents.New(context.TODO(), telnyx.AITestTestFrameworkV1AiagentNewParams{
		AgentName:              "Support Bot Under Test",
		AssistantID:            "e1a2b3c4-d5e6-f7a8-b9c0-d1e2f3a4b5c6",
		AutoFetchCallsEnabled:  telnyx.Bool(true),
		AutoSyncPromptEnabled:  telnyx.Bool(true),
		AutoUpdateDropoffNodes: telnyx.Bool(false),
		AutoUpdateTopicNodes:   telnyx.Bool(false),
		ContactNumber:          telnyx.String("contact_number"),
		Description:            telnyx.String("Evaluates the support assistant against generated callers."),
		Inbound:                telnyx.Bool(false),
		Language:               telnyx.String("en"),
		LlmMaxTokens:           telnyx.Int(512),
		LlmModel:               telnyx.String("openai/gpt-oss-120b"),
		LlmSystemPrompt:        telnyx.String("llm_system_prompt"),
		LlmTemperature:         telnyx.Float(0.7),
		OutboundAutoCall:       telnyx.Bool(false),
		SipEndpoint:            telnyx.String(""),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1AiagentGet(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Aiagents.Get(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1AiagentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Aiagents.Update(
		context.TODO(),
		0,
		telnyx.AITestTestFrameworkV1AiagentUpdateParams{
			AgentName:              telnyx.String("Support Bot Under Test - renamed"),
			AssistantID:            telnyx.String("assistant_id"),
			AutoFetchCallsEnabled:  telnyx.Bool(true),
			AutoSyncPromptEnabled:  telnyx.Bool(true),
			AutoUpdateDropoffNodes: telnyx.Bool(true),
			AutoUpdateTopicNodes:   telnyx.Bool(true),
			ContactNumber:          telnyx.String("contact_number"),
			Description:            telnyx.String("description"),
			Inbound:                telnyx.Bool(true),
			Language:               telnyx.String("language"),
			LlmMaxTokens:           telnyx.Int(256),
			LlmModel:               telnyx.String("llm_model"),
			LlmSystemPrompt:        telnyx.String("llm_system_prompt"),
			LlmTemperature:         telnyx.Float(0.5),
			OutboundAutoCall:       telnyx.Bool(true),
			SipEndpoint:            telnyx.String("sip_endpoint"),
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

func TestAITestTestFrameworkV1AiagentListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Aiagents.List(context.TODO(), telnyx.AITestTestFrameworkV1AiagentListParams{
		Page:     telnyx.Int(1),
		PageSize: telnyx.Int(1),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1AiagentDelete(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Aiagents.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
