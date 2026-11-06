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

func TestAgentGatewayMcpServerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.McpServers.New(
		context.TODO(),
		"32ce035d-3165-4031-b1a7-a7c0c259590a",
		telnyx.AgentGatewayMcpServerNewParams{
			Name: "orders",
			URL:  "http://10.40.3.7:9000/mcp",
			Auth: telnyx.AgentGatewayMcpServerNewParamsAuthUnion{
				OfAPIKey: &telnyx.AgentGatewayStaticAuthParam{
					Type:   telnyx.AgentGatewayStaticAuthTypeAPIKey,
					Value:  "example-orders-key",
					Header: telnyx.String("header"),
				},
			},
			ForwardedHeaders: []string{"X-Request-Id"},
			OpenToEveryKey:   telnyx.Bool(true),
			StaticHeaders: map[string]string{
				"X-Env": "prod",
			},
			Tools:     []string{"string"},
			Transport: telnyx.AgentGatewayMcpServerNewParamsTransportHTTP,
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

func TestAgentGatewayMcpServerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AgentGateways.McpServers.Update(
		context.TODO(),
		"88d5a97b-4080-4855-8bac-a39811b1dc24",
		telnyx.AgentGatewayMcpServerUpdateParams{
			InstanceID: "32ce035d-3165-4031-b1a7-a7c0c259590a",
			Auth: telnyx.AgentGatewayMcpServerUpdateParamsAuthUnion{
				OfNoAuthentication: &telnyx.AgentGatewayNoAuthParam{
					Type: telnyx.AgentGatewayNoAuthTypeNone,
				},
			},
			ForwardedHeaders: []string{"string"},
			Name:             telnyx.String("name"),
			OpenToEveryKey:   telnyx.Bool(true),
			StaticHeaders: map[string]string{
				"foo": "string",
			},
			Tools:     []string{"get_order", "create_ticket"},
			Transport: telnyx.AgentGatewayMcpServerUpdateParamsTransportHTTP,
			URL:       telnyx.String("url"),
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

func TestAgentGatewayMcpServerList(t *testing.T) {
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
	_, err := client.AgentGateways.McpServers.List(context.TODO(), "32ce035d-3165-4031-b1a7-a7c0c259590a")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGatewayMcpServerDelete(t *testing.T) {
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
	_, err := client.AgentGateways.McpServers.Delete(
		context.TODO(),
		"88d5a97b-4080-4855-8bac-a39811b1dc24",
		telnyx.AgentGatewayMcpServerDeleteParams{
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
