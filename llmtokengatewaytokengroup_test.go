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

func TestLlmTokenGatewayTokenGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenGroups.New(context.TODO(), telnyx.LlmTokenGatewayTokenGroupNewParams{
		AllowedModels:  []string{"Kimi-K3", "GLM-5.3-Flash"},
		Name:           "support-backend",
		IdempotencyKey: "Idempotency-Key",
		Blocked:        telnyx.Bool(true),
		BudgetDuration: telnyx.BudgetDurationBudgetDuration1D,
		CacheTtl:       telnyx.Int(1),
		Fallbacks: map[string][]string{
			"Kimi-K3": {"GLM-5.3-Flash"},
		},
		Guardrails: telnyx.LlmTokenGatewayTokenGroupNewParamsGuardrails{
			Dlp: telnyx.LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp{
				Profiles: []string{"financial"},
				Prompt:   "ignore",
				Response: "ignore",
			},
			Secrets: telnyx.LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets{
				Prompt:   "ignore",
				Response: "ignore",
			},
			Streaming: "buffered",
		},
		MaxBudget:           telnyx.Float(0),
		MaxParallelRequests: telnyx.Int(20),
		ProviderKeyIDs:      []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Retries: telnyx.LlmTokenGatewayTokenGroupNewParamsRetries{
			MaxAttempts:      2,
			Backoff:          "exponential",
			RequestTimeoutMs: telnyx.Int(10000),
			RetryDelayMs:     telnyx.Int(250),
		},
		RpmLimit: telnyx.Int(0),
		TpmLimit: telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmTokenGatewayTokenGroupGet(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenGroups.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmTokenGatewayTokenGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenGroups.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.LlmTokenGatewayTokenGroupUpdateParams{
			IdempotencyKey: "Idempotency-Key",
			IfMatch:        "If-Match",
			AllowedModels:  []string{"Kimi-K3", "GLM-5.3-Flash"},
			Blocked:        telnyx.Bool(true),
			BudgetDuration: telnyx.BudgetDurationBudgetDuration1D,
			CacheTtl:       telnyx.Int(1),
			Fallbacks: map[string][]string{
				"Kimi-K3": {"GLM-5.3-Flash"},
			},
			Guardrails: telnyx.LlmTokenGatewayTokenGroupUpdateParamsGuardrails{
				Dlp: telnyx.LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp{
					Profiles: []string{"financial"},
					Prompt:   "ignore",
					Response: "ignore",
				},
				Secrets: telnyx.LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets{
					Prompt:   "ignore",
					Response: "ignore",
				},
				Streaming: "buffered",
			},
			MaxBudget:           telnyx.Float(0),
			MaxParallelRequests: telnyx.Int(20),
			Name:                telnyx.String("support-backend"),
			ProviderKeyIDs:      []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
			Retries: telnyx.LlmTokenGatewayTokenGroupUpdateParamsRetries{
				MaxAttempts:      2,
				Backoff:          "exponential",
				RequestTimeoutMs: telnyx.Int(10000),
				RetryDelayMs:     telnyx.Int(250),
			},
			RpmLimit: telnyx.Int(0),
			TpmLimit: telnyx.Int(0),
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

func TestLlmTokenGatewayTokenGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.LlmTokenGateway.TokenGroups.List(context.TODO(), telnyx.LlmTokenGatewayTokenGroupListParams{
		PageNumber:   telnyx.Int(1),
		PageSize:     telnyx.Int(1),
		PageSnapshot: telnyx.String("page[snapshot]"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLlmTokenGatewayTokenGroupDelete(t *testing.T) {
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
	err := client.LlmTokenGateway.TokenGroups.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.LlmTokenGatewayTokenGroupDeleteParams{
			IdempotencyKey: "Idempotency-Key",
			IfMatch:        "If-Match",
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
