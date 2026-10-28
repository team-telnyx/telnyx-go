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

func TestAITestTestFrameworkV1MetricNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.New(context.TODO(), telnyx.AITestTestFrameworkV1MetricNewParams{
		CreateAITestMetricRequest: telnyx.CreateAITestMetricRequestParam{
			Name:                 "Escalates on refund denial",
			AddToNewAgents:       telnyx.Bool(false),
			Agents:               []int64{1259},
			AssistantID:          telnyx.String("assistant_id"),
			AudioEnabled:         telnyx.Bool(true),
			Description:          telnyx.String("Returns a score for whether the assistant escalates when the caller's refund is denied."),
			DisplayOrder:         telnyx.Int(0),
			EnumValues:           []string{"string"},
			EvalType:             telnyx.String("continuous_qualitative"),
			ObservabilityEnabled: telnyx.Bool(true),
			Prompt:               telnyx.String("Score the assistant's escalation behaviour between 0.0 and 1.0."),
			SamplingEnabled:      telnyx.Bool(true),
			SimulationEnabled:    telnyx.Bool(true),
			Type:                 telnyx.String("llm_judge"),
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

func TestAITestTestFrameworkV1MetricGet(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.Get(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1MetricUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.Update(
		context.TODO(),
		0,
		telnyx.AITestTestFrameworkV1MetricUpdateParams{
			AddToNewAgents:       telnyx.Bool(true),
			Agents:               []int64{0},
			AudioEnabled:         telnyx.Bool(true),
			Description:          telnyx.String("description"),
			DisplayOrder:         telnyx.Int(0),
			EnumValues:           []string{"string"},
			EvalType:             telnyx.String("eval_type"),
			Name:                 telnyx.String("Escalates on refund denial"),
			ObservabilityEnabled: telnyx.Bool(true),
			Prompt:               telnyx.String("Score the assistant's escalation behaviour between 0.0 and 1.0, considering tone."),
			SamplingEnabled:      telnyx.Bool(true),
			SimulationEnabled:    telnyx.Bool(true),
			Type:                 telnyx.String("type"),
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

func TestAITestTestFrameworkV1MetricListWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.List(context.TODO(), telnyx.AITestTestFrameworkV1MetricListParams{
		AgentID:     telnyx.Int(0),
		AssistantID: telnyx.String("assistant_id"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1MetricDelete(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1MetricBulkNew(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.BulkNew(context.TODO(), telnyx.AITestTestFrameworkV1MetricBulkNewParams{
		Body: []telnyx.CreateAITestMetricRequestParam{{
			Name:                 "Greets within 5 seconds",
			AddToNewAgents:       telnyx.Bool(false),
			Agents:               []int64{1259},
			AssistantID:          telnyx.String("assistant_id"),
			AudioEnabled:         telnyx.Bool(true),
			Description:          telnyx.String("Returns whether the assistant greeted the caller within five seconds."),
			DisplayOrder:         telnyx.Int(0),
			EnumValues:           []string{"string"},
			EvalType:             telnyx.String("continuous_qualitative"),
			ObservabilityEnabled: telnyx.Bool(true),
			Prompt:               telnyx.String("Score whether the greeting was timely and natural."),
			SamplingEnabled:      telnyx.Bool(true),
			SimulationEnabled:    telnyx.Bool(true),
			Type:                 telnyx.String("llm_judge"),
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

func TestAITestTestFrameworkV1MetricGenerateMetricsWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.GenerateMetrics(context.TODO(), telnyx.AITestTestFrameworkV1MetricGenerateMetricsParams{
		AgentID:    1259,
		NumMetrics: 5,
		Scenarios:  []int64{9512},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAITestTestFrameworkV1MetricGetGenerateMetricsProgress(t *testing.T) {
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
	_, err := client.AI.Tests.TestFramework.V1.Metrics.GetGenerateMetricsProgress(context.TODO(), telnyx.AITestTestFrameworkV1MetricGetGenerateMetricsProgressParams{
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
