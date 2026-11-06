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

func TestAgentGatewayKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Keys.New(
		context.TODO(),
		"ecffdf4b-93ea-4a7b-8295-51dfd2077996",
		telnyx.AgentGatewayKeyNewParams{
			BudgetResetsEvery:   telnyx.String("30d"),
			ExpiresIn:           telnyx.String("30d"),
			MaxBudget:           telnyx.Float(20),
			MaxParallelRequests: telnyx.Int(4),
			Models:              []string{"GLM-5.3-Flash"},
			Name:                telnyx.String("support-agent"),
			RequestsPerMinute:   telnyx.Int(60),
			TeamID:              telnyx.String("0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e"),
			TokensPerMinute:     telnyx.Int(1),
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

func TestAgentGatewayKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Keys.Update(
		context.TODO(),
		"5f1c3a0e9b7d4c2a8e6f0b1d3c5a7e9f2b4d6c8a0e1f3a5c7e9b1d3f5a7c9e0b",
		telnyx.AgentGatewayKeyUpdateParams{
			InstanceID:          "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			Blocked:             telnyx.Bool(true),
			BudgetResetsEvery:   telnyx.String("30d"),
			ExpiresIn:           telnyx.String("30d"),
			MaxBudget:           telnyx.Float(0),
			MaxParallelRequests: telnyx.Int(1),
			Models:              []string{"x"},
			Name:                telnyx.String("x"),
			RequestsPerMinute:   telnyx.Int(1),
			TeamID:              telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			TokensPerMinute:     telnyx.Int(1),
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

func TestAgentGatewayKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Keys.List(
		context.TODO(),
		"ecffdf4b-93ea-4a7b-8295-51dfd2077996",
		telnyx.AgentGatewayKeyListParams{
			TeamID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestAgentGatewayKeyDelete(t *testing.T) {
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
	_, err := client.AgentGateways.Keys.Delete(
		context.TODO(),
		"5f1c3a0e9b7d4c2a8e6f0b1d3c5a7e9f2b4d6c8a0e1f3a5c7e9b1d3f5a7c9e0b",
		telnyx.AgentGatewayKeyDeleteParams{
			InstanceID: "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
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
