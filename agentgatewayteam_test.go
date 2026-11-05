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

func TestAgentGatewayTeamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Teams.New(
		context.TODO(),
		"ecffdf4b-93ea-4a7b-8295-51dfd2077996",
		telnyx.AgentGatewayTeamNewParams{
			Name:              "support",
			BudgetResetsEvery: telnyx.String("30d"),
			MaxBudget:         telnyx.Float(50),
			Models:            []string{"GLM-5.3-Flash", "Qwen3.8-27B-nothink"},
			RequestsPerMinute: telnyx.Int(120),
			TokensPerMinute:   telnyx.Int(1),
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

func TestAgentGatewayTeamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Teams.Update(
		context.TODO(),
		"0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e",
		telnyx.AgentGatewayTeamUpdateParams{
			InstanceID:        "ecffdf4b-93ea-4a7b-8295-51dfd2077996",
			Blocked:           telnyx.Bool(true),
			BudgetResetsEvery: telnyx.String("30d"),
			MaxBudget:         telnyx.Float(100),
			Models:            []string{"x"},
			Name:              telnyx.String("x"),
			RequestsPerMinute: telnyx.Int(1),
			TokensPerMinute:   telnyx.Int(1),
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

func TestAgentGatewayTeamList(t *testing.T) {
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
	_, err := client.AgentGateways.Teams.List(context.TODO(), "ecffdf4b-93ea-4a7b-8295-51dfd2077996")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGatewayTeamDelete(t *testing.T) {
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
	_, err := client.AgentGateways.Teams.Delete(
		context.TODO(),
		"0b6c5d4e-3f2a-4b1c-9d8e-7f6a5b4c3d2e",
		telnyx.AgentGatewayTeamDeleteParams{
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
