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

func TestAIAssistantNewWithOptionalParams(t *testing.T) {
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
			ConversationInactivityMinutes: telnyx.Int(1),
			DefaultMessagingProfileID:     telnyx.String("default_messaging_profile_id"),
			DeliveryStatusWebhookURL:      telnyx.String("delivery_status_webhook_url"),
		},
		PrivacySettings: telnyx.PrivacySettingsParam{
			DataRetention: telnyx.Bool(true),
		},
		TelephonySettings: telnyx.TelephonySettingsParam{
			DefaultTexmlAppID: telnyx.String("default_texml_app_id"),
			NoiseSuppression:  telnyx.TelephonySettingsNoiseSuppressionKrisp,
			NoiseSuppressionConfig: telnyx.TelephonySettingsNoiseSuppressionConfigParam{
				AttenuationLimit: telnyx.Int(0),
				Mode:             "advanced",
			},
			RecordingSettings: telnyx.TelephonySettingsRecordingSettingsParam{
				Channels: "single",
				Format:   "wav",
			},
			SupportsUnauthenticatedWebCalls: telnyx.Bool(true),
			TimeLimitSecs:                   telnyx.Int(30),
			UserIdleTimeoutSecs:             telnyx.Int(30),
			VoicemailDetection: telnyx.TelephonySettingsVoicemailDetectionParam{
				OnVoicemailDetected: telnyx.TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam{
					Action: "stop_assistant",
					VoicemailMessage: telnyx.TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam{
						Message: telnyx.String("message"),
						Prompt:  telnyx.String("prompt"),
						Type:    "prompt",
					},
				},
			},
		},
		Tools: []telnyx.AssistantToolsItemsUnionParam{{
			OfWebhook: &telnyx.InferenceEmbeddingWebhookToolParams{
				Type: telnyx.InferenceEmbeddingWebhookToolParamsTypeWebhook,
				Webhook: telnyx.InferenceEmbeddingWebhookToolParamsWebhook{
					Description: "description",
					Name:        "name",
					URL:         "https://example.com/api/v1/function",
					Async:       telnyx.Bool(true),
					BodyParameters: telnyx.InferenceEmbeddingWebhookToolParamsWebhookBodyParameters{
						Properties: map[string]any{
							"age":      "bar",
							"location": "bar",
						},
						Required: []string{"age", "location"},
						Type:     "object",
					},
					Headers: []telnyx.InferenceEmbeddingWebhookToolParamsWebhookHeader{{
						Name:  telnyx.String("name"),
						Value: telnyx.String("value"),
					}},
					Method: "GET",
					PathParameters: telnyx.InferenceEmbeddingWebhookToolParamsWebhookPathParameters{
						Properties: map[string]any{
							"id": "bar",
						},
						Required: []string{"id"},
						Type:     "object",
					},
					QueryParameters: telnyx.InferenceEmbeddingWebhookToolParamsWebhookQueryParameters{
						Properties: map[string]any{
							"page": "bar",
						},
						Required: []string{"page"},
						Type:     "object",
					},
					TimeoutMs: telnyx.Int(500),
				},
			},
		}},
		Transcription: telnyx.TranscriptionSettingsParam{
			Language: telnyx.String("language"),
			Model:    telnyx.TranscriptionSettingsModelDeepgramFlux,
			Region:   telnyx.String("region"),
			Settings: telnyx.TranscriptionSettingsConfigParam{
				EagerEotThreshold: telnyx.Float(0.3),
				EotThreshold:      telnyx.Float(0),
				EotTimeoutMs:      telnyx.Int(0),
				Numerals:          telnyx.Bool(true),
				SmartFormat:       telnyx.Bool(true),
			},
		},
		VoiceSettings: telnyx.VoiceSettingsParam{
			Voice:     "voice",
			APIKeyRef: telnyx.String("api_key_ref"),
			BackgroundAudio: telnyx.VoiceSettingsBackgroundAudioUnionParam{
				OfPredefinedMedia: &telnyx.VoiceSettingsBackgroundAudioPredefinedMediaParam{
					Value: "silence",
				},
			},
			LanguageBoost:   telnyx.VoiceSettingsLanguageBoostAuto,
			SimilarityBoost: telnyx.Float(0),
			Speed:           telnyx.Float(0),
			Style:           telnyx.Float(0),
			Temperature:     telnyx.Float(0),
			UseSpeakerBoost: telnyx.Bool(true),
			VoiceSpeed:      telnyx.Float(0),
		},
		WidgetSettings: telnyx.WidgetSettingsParam{
			AgentThinkingText: telnyx.String("agent_thinking_text"),
			AudioVisualizerConfig: telnyx.AudioVisualizerConfigParam{
				Color:  telnyx.AudioVisualizerConfigColorVerdant,
				Preset: telnyx.String("preset"),
			},
			DefaultState:         telnyx.WidgetSettingsDefaultStateExpanded,
			GiveFeedbackURL:      telnyx.String("give_feedback_url"),
			LogoIconURL:          telnyx.String("logo_icon_url"),
			Position:             telnyx.WidgetSettingsPositionFixed,
			ReportIssueURL:       telnyx.String("report_issue_url"),
			SpeakToInterruptText: telnyx.String("speak_to_interrupt_text"),
			StartCallText:        telnyx.String("start_call_text"),
			Theme:                telnyx.WidgetSettingsThemeLight,
			ViewHistoryURL:       telnyx.String("view_history_url"),
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
				ConversationInactivityMinutes: telnyx.Int(1),
				DefaultMessagingProfileID:     telnyx.String("default_messaging_profile_id"),
				DeliveryStatusWebhookURL:      telnyx.String("delivery_status_webhook_url"),
			},
			Model: telnyx.String("model"),
			Name:  telnyx.String("name"),
			PrivacySettings: telnyx.PrivacySettingsParam{
				DataRetention: telnyx.Bool(true),
			},
			PromoteToMain: telnyx.Bool(true),
			TelephonySettings: telnyx.TelephonySettingsParam{
				DefaultTexmlAppID: telnyx.String("default_texml_app_id"),
				NoiseSuppression:  telnyx.TelephonySettingsNoiseSuppressionKrisp,
				NoiseSuppressionConfig: telnyx.TelephonySettingsNoiseSuppressionConfigParam{
					AttenuationLimit: telnyx.Int(0),
					Mode:             "advanced",
				},
				RecordingSettings: telnyx.TelephonySettingsRecordingSettingsParam{
					Channels: "single",
					Format:   "wav",
				},
				SupportsUnauthenticatedWebCalls: telnyx.Bool(true),
				TimeLimitSecs:                   telnyx.Int(30),
				UserIdleTimeoutSecs:             telnyx.Int(30),
				VoicemailDetection: telnyx.TelephonySettingsVoicemailDetectionParam{
					OnVoicemailDetected: telnyx.TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam{
						Action: "stop_assistant",
						VoicemailMessage: telnyx.TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam{
							Message: telnyx.String("message"),
							Prompt:  telnyx.String("prompt"),
							Type:    "prompt",
						},
					},
				},
			},
			Tools: []telnyx.AssistantToolsItemsUnionParam{{
				OfWebhook: &telnyx.InferenceEmbeddingWebhookToolParams{
					Type: telnyx.InferenceEmbeddingWebhookToolParamsTypeWebhook,
					Webhook: telnyx.InferenceEmbeddingWebhookToolParamsWebhook{
						Description: "description",
						Name:        "name",
						URL:         "https://example.com/api/v1/function",
						Async:       telnyx.Bool(true),
						BodyParameters: telnyx.InferenceEmbeddingWebhookToolParamsWebhookBodyParameters{
							Properties: map[string]any{
								"age":      "bar",
								"location": "bar",
							},
							Required: []string{"age", "location"},
							Type:     "object",
						},
						Headers: []telnyx.InferenceEmbeddingWebhookToolParamsWebhookHeader{{
							Name:  telnyx.String("name"),
							Value: telnyx.String("value"),
						}},
						Method: "GET",
						PathParameters: telnyx.InferenceEmbeddingWebhookToolParamsWebhookPathParameters{
							Properties: map[string]any{
								"id": "bar",
							},
							Required: []string{"id"},
							Type:     "object",
						},
						QueryParameters: telnyx.InferenceEmbeddingWebhookToolParamsWebhookQueryParameters{
							Properties: map[string]any{
								"page": "bar",
							},
							Required: []string{"page"},
							Type:     "object",
						},
						TimeoutMs: telnyx.Int(500),
					},
				},
			}},
			Transcription: telnyx.TranscriptionSettingsParam{
				Language: telnyx.String("language"),
				Model:    telnyx.TranscriptionSettingsModelDeepgramFlux,
				Region:   telnyx.String("region"),
				Settings: telnyx.TranscriptionSettingsConfigParam{
					EagerEotThreshold: telnyx.Float(0.3),
					EotThreshold:      telnyx.Float(0),
					EotTimeoutMs:      telnyx.Int(0),
					Numerals:          telnyx.Bool(true),
					SmartFormat:       telnyx.Bool(true),
				},
			},
			VoiceSettings: telnyx.VoiceSettingsParam{
				Voice:     "voice",
				APIKeyRef: telnyx.String("api_key_ref"),
				BackgroundAudio: telnyx.VoiceSettingsBackgroundAudioUnionParam{
					OfPredefinedMedia: &telnyx.VoiceSettingsBackgroundAudioPredefinedMediaParam{
						Value: "silence",
					},
				},
				LanguageBoost:   telnyx.VoiceSettingsLanguageBoostAuto,
				SimilarityBoost: telnyx.Float(0),
				Speed:           telnyx.Float(0),
				Style:           telnyx.Float(0),
				Temperature:     telnyx.Float(0),
				UseSpeakerBoost: telnyx.Bool(true),
				VoiceSpeed:      telnyx.Float(0),
			},
			WidgetSettings: telnyx.WidgetSettingsParam{
				AgentThinkingText: telnyx.String("agent_thinking_text"),
				AudioVisualizerConfig: telnyx.AudioVisualizerConfigParam{
					Color:  telnyx.AudioVisualizerConfigColorVerdant,
					Preset: telnyx.String("preset"),
				},
				DefaultState:         telnyx.WidgetSettingsDefaultStateExpanded,
				GiveFeedbackURL:      telnyx.String("give_feedback_url"),
				LogoIconURL:          telnyx.String("logo_icon_url"),
				Position:             telnyx.WidgetSettingsPositionFixed,
				ReportIssueURL:       telnyx.String("report_issue_url"),
				SpeakToInterruptText: telnyx.String("speak_to_interrupt_text"),
				StartCallText:        telnyx.String("start_call_text"),
				Theme:                telnyx.WidgetSettingsThemeLight,
				ViewHistoryURL:       telnyx.String("view_history_url"),
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
	_, err := client.AI.Assistants.GetTexml(context.TODO(), "assistant_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantImportsWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.Imports(context.TODO(), telnyx.AIAssistantImportsParams{
		APIKeyRef: "api_key_ref",
		Provider:  telnyx.AIAssistantImportsParamsProviderElevenlabs,
		ImportIDs: []string{"string"},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAssistantSendSMSWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Assistants.SendSMS(
		context.TODO(),
		"assistant_id",
		telnyx.AIAssistantSendSMSParams{
			From: "from",
			To:   "to",
			ConversationMetadata: map[string]telnyx.AIAssistantSendSMSParamsConversationMetadataUnion{
				"foo": {
					OfString: telnyx.String("string"),
				},
			},
			ShouldCreateConversation: telnyx.Bool(true),
			Text:                     telnyx.String("text"),
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
