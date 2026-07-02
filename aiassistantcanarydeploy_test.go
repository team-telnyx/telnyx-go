// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go/v4"
	"github.com/stainless-sdks/telnyx-go/v4/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/v4/option"
)

func TestAIAssistantCanaryDeployNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.CanaryDeploys.New(
		context.TODO(),
		"assistant_id",
		telnyx.AIAssistantCanaryDeployNewParams{
			CanaryDeploy: telnyx.CanaryDeployParam{
				Rules: []telnyx.CanaryDeployRuleParam{{
					Serve: telnyx.CanaryDeployRuleServeParam{
						Rollout: []telnyx.CanaryDeployRuleServeRolloutParam{{
							VersionID: "version_id",
							Weight:    0,
						}},
						VersionID: telnyx.String("version_id"),
					},
					Match: []telnyx.CanaryDeployRuleMatchParam{{
						Attribute: "attribute",
						Operator:  "in",
						Values:    []string{"string"},
					}},
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

func TestAIAssistantCanaryDeployGet(t *testing.T) {
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
	_, err := client.AI.Assistants.CanaryDeploys.Get(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantCanaryDeployUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.CanaryDeploys.Update(
		context.TODO(),
		"assistant_id",
		telnyx.AIAssistantCanaryDeployUpdateParams{
			CanaryDeploy: telnyx.CanaryDeployParam{
				Rules: []telnyx.CanaryDeployRuleParam{{
					Serve: telnyx.CanaryDeployRuleServeParam{
						Rollout: []telnyx.CanaryDeployRuleServeRolloutParam{{
							VersionID: "version_id",
							Weight:    0,
						}},
						VersionID: telnyx.String("version_id"),
					},
					Match: []telnyx.CanaryDeployRuleMatchParam{{
						Attribute: "attribute",
						Operator:  "in",
						Values:    []string{"string"},
					}},
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

func TestAIAssistantCanaryDeployDelete(t *testing.T) {
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
	err := client.AI.Assistants.CanaryDeploys.Delete(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
