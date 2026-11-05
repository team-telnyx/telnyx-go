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

func TestAgentGatewayPolicyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Policies.New(
		context.TODO(),
		"32ce035d-3165-4031-b1a7-a7c0c259590a",
		telnyx.AgentGatewayPolicyNewParams{
			Name:        "baseline",
			Base:        telnyx.String("base"),
			Description: telnyx.String("Mask personal data"),
			Guardrails: telnyx.AgentGatewayPolicyGuardrailsParam{
				Add:    []string{"mask-pii"},
				Remove: []string{"string"},
			},
			Model: telnyx.String("model"),
			Pipeline: telnyx.AgentGatewayPolicyPipelineParam{
				Hook: telnyx.AgentGatewayPolicyPipelineHookBeforeRequest,
				Steps: []telnyx.AgentGatewayPolicyPipelineStepParam{{
					Guardrail: "guardrail",
					Message:   telnyx.String("message"),
					OnError:   "next",
					OnFail:    "next",
					OnPass:    "next",
					PassData:  telnyx.Bool(true),
				}},
			},
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

func TestAgentGatewayPolicyGet(t *testing.T) {
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
	_, err := client.AgentGateways.Policies.Get(
		context.TODO(),
		"strict",
		telnyx.AgentGatewayPolicyGetParams{
			InstanceID: "32ce035d-3165-4031-b1a7-a7c0c259590a",
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

func TestAgentGatewayPolicyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.Policies.Update(
		context.TODO(),
		"strict",
		telnyx.AgentGatewayPolicyUpdateParams{
			InstanceID:  "32ce035d-3165-4031-b1a7-a7c0c259590a",
			Base:        telnyx.String("base"),
			Description: telnyx.String("Baseline plus a word check"),
			Guardrails: telnyx.AgentGatewayPolicyGuardrailsParam{
				Add:    []string{"string"},
				Remove: []string{"string"},
			},
			Model: telnyx.String("model"),
			Pipeline: telnyx.AgentGatewayPolicyPipelineParam{
				Hook: telnyx.AgentGatewayPolicyPipelineHookBeforeRequest,
				Steps: []telnyx.AgentGatewayPolicyPipelineStepParam{{
					Guardrail: "guardrail",
					Message:   telnyx.String("message"),
					OnError:   "next",
					OnFail:    "next",
					OnPass:    "next",
					PassData:  telnyx.Bool(true),
				}},
			},
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

func TestAgentGatewayPolicyList(t *testing.T) {
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
	_, err := client.AgentGateways.Policies.List(context.TODO(), "32ce035d-3165-4031-b1a7-a7c0c259590a")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGatewayPolicyDelete(t *testing.T) {
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
	_, err := client.AgentGateways.Policies.Delete(
		context.TODO(),
		"strict",
		telnyx.AgentGatewayPolicyDeleteParams{
			InstanceID: "32ce035d-3165-4031-b1a7-a7c0c259590a",
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
