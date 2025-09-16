// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/telnyx-go"
	"github.com/stainless-sdks/telnyx-go/internal/testutil"
	"github.com/stainless-sdks/telnyx-go/option"
)

func TestAIAssistantVersionGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AI.Assistants.Versions.Get(
		context.TODO(),
		"version_id",
		telnyx.AIAssistantVersionGetParams{
			AssistantID:       "assistant_id",
			IncludeMcpServers: telnyx.Bool(true),
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

func TestAIAssistantVersionUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AI.Assistants.Versions.Update(
		context.TODO(),
		"version_id",
		telnyx.AIAssistantVersionUpdateParams{
			AssistantID: "assistant_id",
			UpdateAssistant: telnyx.UpdateAssistantParam{
				Description: telnyx.String("description"),
				DynamicVariables: map[string]any{
					"foo": "bar",
				},
				DynamicVariablesWebhookURL: telnyx.String("dynamic_variables_webhook_url"),
				EnabledFeatures:            []telnyx.EnabledFeatures{telnyx.EnabledFeaturesTelephony},
				Greeting:                   telnyx.String("greeting"),
				InsightSettings: telnyx.InsightSettingsParam{
					InsightGroupID: telnyx.String("insight_group_id"),
				},
				Instructions: telnyx.String("instructions"),
				LlmAPIKeyRef: telnyx.String("llm_api_key_ref"),
				MessagingSettings: telnyx.MessagingSettingsParam{
					DefaultMessagingProfileID: telnyx.String("default_messaging_profile_id"),
					DeliveryStatusWebhookURL:  telnyx.String("delivery_status_webhook_url"),
				},
				Model: telnyx.String("model"),
				Name:  telnyx.String("name"),
				PrivacySettings: telnyx.PrivacySettingsParam{
					DataRetention: telnyx.Bool(true),
				},
				TelephonySettings: telnyx.TelephonySettingsParam{
					DefaultTexmlAppID:               telnyx.String("default_texml_app_id"),
					SupportsUnauthenticatedWebCalls: telnyx.Bool(true),
				},
				Tools: []telnyx.AssistantToolUnionParam{{
					OfWebhookTool: &telnyx.WebhookToolParam{
						Type: telnyx.WebhookToolTypeWebhook,
						Webhook: telnyx.InferenceEmbeddingWebhookToolParams{
							Description: "description",
							Name:        "name",
							URL:         "https://example.com/api/v1/function",
							BodyParameters: telnyx.InferenceEmbeddingWebhookToolParamsBodyParameters{
								Properties: map[string]any{
									"age":      "bar",
									"location": "bar",
								},
								Required: []string{"age", "location"},
								Type:     "object",
							},
							Headers: []telnyx.InferenceEmbeddingWebhookToolParamsHeader{{
								Name:  telnyx.String("name"),
								Value: telnyx.String("value"),
							}},
							Method: telnyx.InferenceEmbeddingWebhookToolParamsMethodGet,
							PathParameters: telnyx.InferenceEmbeddingWebhookToolParamsPathParameters{
								Properties: map[string]any{
									"id": "bar",
								},
								Required: []string{"id"},
								Type:     "object",
							},
							QueryParameters: telnyx.InferenceEmbeddingWebhookToolParamsQueryParameters{
								Properties: map[string]any{
									"page": "bar",
								},
								Required: []string{"page"},
								Type:     "object",
							},
						},
					},
				}},
				Transcription: telnyx.TranscriptionSettingsParam{
					Language: telnyx.String("language"),
					Model:    telnyx.String("model"),
				},
				VoiceSettings: telnyx.VoiceSettingsParam{
					Voice:      "voice",
					APIKeyRef:  telnyx.String("api_key_ref"),
					VoiceSpeed: telnyx.Float(0),
				},
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

func TestAIAssistantVersionList(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AI.Assistants.Versions.List(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantVersionDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	err := client.AI.Assistants.Versions.Delete(
		context.TODO(),
		"version_id",
		telnyx.AIAssistantVersionDeleteParams{
			AssistantID: "assistant_id",
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

func TestAIAssistantVersionPromote(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.AI.Assistants.Versions.Promote(
		context.TODO(),
		"version_id",
		telnyx.AIAssistantVersionPromoteParams{
			AssistantID: "assistant_id",
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
