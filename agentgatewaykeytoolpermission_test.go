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

func TestAgentGatewayKeyToolPermissionNew(t *testing.T) {
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
	_, err := client.AgentGateways.Keys.ToolPermissions.New(
		context.TODO(),
		"f739ad190d29e2a5e658c3a486111953e6c18d746fd04374962ab150aae17757",
		telnyx.AgentGatewayKeyToolPermissionNewParams{
			InstanceID: "32ce035d-3165-4031-b1a7-a7c0c259590a",
			AgentGatewayToolPermissions: telnyx.AgentGatewayToolPermissionsParam{
				Servers: []telnyx.AgentGatewayToolPermissionsServerParam{{
					ID:    "88d5a97b-4080-4855-8bac-a39811b1dc24",
					Tools: []string{"get_order"},
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

func TestAgentGatewayKeyToolPermissionList(t *testing.T) {
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
	_, err := client.AgentGateways.Keys.ToolPermissions.List(
		context.TODO(),
		"f739ad190d29e2a5e658c3a486111953e6c18d746fd04374962ab150aae17757",
		telnyx.AgentGatewayKeyToolPermissionListParams{
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
