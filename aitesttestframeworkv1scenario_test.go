// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestAITestTestFrameworkV1ScenarioNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.New(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioNewParams{
		Name:                  "Refund escalation flow",
		Personality:           4643,
		Agent:                 telnyx.Int(1259),
		ExpectedOutcomePrompt: telnyx.String("The assistant explains the refund policy and offers escalation."),
		FirstMessage:          telnyx.String("Hi, I want a refund."),
		FolderPath:            telnyx.String("folder_path"),
		Instructions:          telnyx.String("Call the support line and demand a refund for a damaged product."),
		Metrics:               []int64{266881},
		Tags:                  []string{"support"},
		TestProfile:           telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioGet(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.Get(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.Update(
		context.TODO(),
		0,
		telnyx.AITestTestFrameworkV1ScenarioUpdateParams{
			Agent:                 telnyx.Int(0),
			ExpectedOutcomePrompt: telnyx.String("expected_outcome_prompt"),
			FirstMessage:          telnyx.String("first_message"),
			FolderPath:            telnyx.String("folder_path"),
			InboundPhoneNumber:    telnyx.Int(0),
			Instructions:          telnyx.String("Call the support line, demand a refund, and mention the damaged packaging."),
			MaxDuration:           telnyx.Int(0),
			Metrics:               []int64{266881, 266882},
			Name:                  telnyx.String("Refund escalation flow - updated"),
			Personality:           telnyx.Int(0),
			PhoneNumber:           telnyx.Int(0),
			ScenarioLanguage:      telnyx.String("scenario_language"),
			Tags:                  []string{"string"},
			TestProfile:           telnyx.Int(0),
			ToolIDs:               []string{"string"},
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

func TestAITestTestFrameworkV1ScenarioListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.List(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioListParams{
		AgentID:     telnyx.Int(0),
		AssistantID: telnyx.String("assistant_id"),
		Page:        telnyx.Int(1),
		PageSize:    telnyx.Int(1),
		Tags:        telnyx.String("tags"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioDelete(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioGenerateBgWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.GenerateBg(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioGenerateBgParams{
		NumScenarios:             5,
		AgentID:                  telnyx.Int(1259),
		ExtraInstructions:        telnyx.String("Focus on angry customers asking about refunds."),
		FirstMessage:             telnyx.String("first_message"),
		FolderPath:               telnyx.String("folder_path"),
		GenerateExpectedOutcomes: telnyx.Bool(true),
		InboundPhoneNumber:       telnyx.Int(0),
		Personalities:            []int64{0},
		ScenarioType:             telnyx.AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioTypeWorkflow,
		Tags:                     []string{"string"},
		TestProfile:              telnyx.Int(0),
		ToolIDs:                  []string{"string"},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioGetExportScenariosWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := telnyx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.AI.Tests.TestFramework.V1.Scenarios.GetExportScenarios(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioGetExportScenariosParams{
		ExportType:  telnyx.AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportTypeCsv,
		ScenarioIDs: "scenario_ids",
		AgentID:     telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAITestTestFrameworkV1ScenarioGetGenerateProgress(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.GetGenerateProgress(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioGetGenerateProgressParams{
		ProgressID: "progress_id",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioGetScenarioAgentProgress(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.GetScenarioAgentProgress(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioGetScenarioAgentProgressParams{
		ProgressID: "progress_id",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioRunScenariosWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.RunScenarios(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioRunScenariosParams{
		AgentID:             telnyx.Int(1259),
		AgentNumber:         telnyx.String("agent_number"),
		AssistantID:         telnyx.String("assistant_id"),
		ConcurrencyLimit:    telnyx.Int(3),
		Frequency:           telnyx.Int(0),
		Mode:                telnyx.AITestTestFrameworkV1ScenarioRunScenariosParamsModeSameNumber,
		Name:                telnyx.String("name"),
		OutboundPhoneNumber: telnyx.String("outbound_phone_number"),
		PersonalityIDs:      []int64{0},
		Scenarios: telnyx.AITestTestFrameworkV1ScenarioRunScenariosParamsScenariosUnion{
			OfAnyArray: []any{9512},
		},
		Tags:           []string{"regression"},
		TestProfileIDs: []int64{0},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1ScenarioScenarioAgentWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Scenarios.ScenarioAgent(context.TODO(), telnyx.AITestTestFrameworkV1ScenarioScenarioAgentParams{
		Scenarios: telnyx.AITestTestFrameworkV1ScenarioScenarioAgentParamsScenariosUnion{
			OfString: telnyx.String("all"),
		},
		UserMessage:             "Make the callers ask about damaged packaging too.",
		AgentID:                 telnyx.Int(1259),
		AgentResolutionResponse: telnyx.String("agent_resolution_response"),
		ApplyChanges:            telnyx.Bool(true),
		Context: []telnyx.AITestTestFrameworkV1ScenarioScenarioAgentParamsContext{{
			Content:   telnyx.String("content"),
			Role:      telnyx.String("role"),
			Timestamp: telnyx.String("timestamp"),
		}},
		KnowledgeBase: []int64{0},
		NumScenarios:  telnyx.Int(0),
		PendingToolCalls: []map[string]any{{
			"foo": "bar",
		}},
		ScenariosNeedingAgent: []int64{0},
		SessionID:             telnyx.Int(0),
		UpdatedScenarios: []map[string]any{{
			"foo": "bar",
		}},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
