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

func TestAgentGatewayAPINewWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.APIs.New(
		context.TODO(),
		"32ce035d-3165-4031-b1a7-a7c0c259590a",
		telnyx.AgentGatewayAPINewParams{
			Name:    "petstore",
			SpecURL: "https://petstore3.swagger.io/api/v3/openapi.json",
			Auth: telnyx.AgentGatewayAPINewParamsAuthUnion{
				OfNoAuthentication: &telnyx.AgentGatewayNoAuthParam{
					Type: telnyx.AgentGatewayNoAuthTypeNone,
				},
			},
			BaseURL:          telnyx.String("https://petstore3.swagger.io/api/v3"),
			ForwardedHeaders: []string{"string"},
			OpenToEveryKey:   telnyx.Bool(true),
			StaticHeaders: map[string]string{
				"foo": "string",
			},
			Tools: []string{"string"},
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

func TestAgentGatewayAPIUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.APIs.Update(
		context.TODO(),
		"dbe53cc0-bf7e-4f7a-8d84-752d19890eb3",
		telnyx.AgentGatewayAPIUpdateParams{
			InstanceID: "32ce035d-3165-4031-b1a7-a7c0c259590a",
			Auth: telnyx.AgentGatewayAPIUpdateParamsAuthUnion{
				OfNoAuthentication: &telnyx.AgentGatewayNoAuthParam{
					Type: telnyx.AgentGatewayNoAuthTypeNone,
				},
			},
			BaseURL:          telnyx.String("base_url"),
			ForwardedHeaders: []string{"string"},
			Name:             telnyx.String("name"),
			OpenToEveryKey:   telnyx.Bool(true),
			SpecURL:          telnyx.String("spec_url"),
			StaticHeaders: map[string]string{
				"foo": "string",
			},
			Tools: []string{"findpetsbystatus", "getpetbyid"},
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

func TestAgentGatewayAPIList(t *testing.T) {
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
	_, err := client.AgentGateways.APIs.List(context.TODO(), "32ce035d-3165-4031-b1a7-a7c0c259590a")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGatewayAPIDelete(t *testing.T) {
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
	_, err := client.AgentGateways.APIs.Delete(
		context.TODO(),
		"dbe53cc0-bf7e-4f7a-8d84-752d19890eb3",
		telnyx.AgentGatewayAPIDeleteParams{
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
