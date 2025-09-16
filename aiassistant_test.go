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

func TestAIAssistantNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.New(context.TODO(), telnyx.AIAssistantNewParams{
		Instructions: "instructions",
		Model:        "model",
		Name:         "name",
		Description:  telnyx.String("description"),
		DynamicVariables: map[string]any{
			"foo": "bar",
		},
		DynamicVariablesWebhookURL: telnyx.String("dynamic_variables_webhook_url"),
		EnabledFeatures:            []telnyx.EnabledFeatures{telnyx.EnabledFeaturesTelephony},
		Greeting:                   telnyx.String("greeting"),
		InsightSettings: telnyx.InsightSettingsParam{
			InsightGroupID: telnyx.String("insight_group_id"),
		},
		LlmAPIKeyRef: telnyx.String("llm_api_key_ref"),
		MessagingSettings: telnyx.MessagingSettingsParam{
			DefaultMessagingProfileID: telnyx.String("default_messaging_profile_id"),
			DeliveryStatusWebhookURL:  telnyx.String("delivery_status_webhook_url"),
		},
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
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantGetWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Get(
		context.TODO(),
		"assistant_id",
		telnyx.AIAssistantGetParams{
			CallControlID:                    telnyx.String("call_control_id"),
			FetchDynamicVariablesFromWebhook: telnyx.Bool(true),
			From:                             telnyx.String("from"),
			To:                               telnyx.String("to"),
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

func TestAIAssistantUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Update(
		context.TODO(),
		"assistant_id",
		telnyx.AIAssistantUpdateParams{
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
			PromoteToMain: telnyx.Bool(true),
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
	)
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantList(t *testing.T) {
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
	_, err := client.AI.Assistants.List(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantDelete(t *testing.T) {
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
	_, err := client.AI.Assistants.Delete(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantChatWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Chat(
		context.TODO(),
		"assistant_id",
		telnyx.AIAssistantChatParams{
			Content:        "Tell me a joke about cats",
			ConversationID: "42b20469-1215-4a9a-8964-c36f66b406f4",
			Name:           telnyx.String("Charlie"),
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

func TestAIAssistantClone(t *testing.T) {
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
	_, err := client.AI.Assistants.Clone(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantGetTexml(t *testing.T) {
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
	_, err := client.AI.Assistants.GetTexml(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantImport(t *testing.T) {
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
	_, err := client.AI.Assistants.Import(context.TODO(), telnyx.AIAssistantImportParams{
		APIKeyRef: "api_key_ref",
		Provider:  telnyx.AIAssistantImportParamsProviderElevenlabs,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
