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

func TestTtsIntegrationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TtsIntegrations.New(context.TODO(), telnyx.TtsIntegrationNewParams{
		Capabilities: telnyx.CommunityTtsEngineCapabilitiesParam{
			StreamedAudio:         true,
			FlushSupport:          telnyx.Bool(true),
			ForceInterruptSupport: telnyx.Bool(true),
			SsmlSupport:           telnyx.Bool(true),
			VoiceSettingsSchema: map[string]any{
				"voice_speed": map[string]any{},
				"emotion":     map[string]any{},
			},
		},
		Name:                   "acme",
		SupportedAudioFormats:  []telnyx.CommunityTtsAudioFormat{telnyx.CommunityTtsAudioFormatLinear16},
		VoicesURL:              "https://tts.acme.example/voices",
		WebsocketURL:           "wss://tts.acme.example/v1",
		ConnectTimeoutMs:       telnyx.Int(500),
		ContactEmail:           telnyx.String("contact_email"),
		DataRetentionPolicyURL: telnyx.String("data_retention_policy_url"),
		DocumentationURL:       telnyx.String("https://tts.acme.example/docs"),
		Visibility:             telnyx.CommunityTtsEngineVisibilityPrivate,
		VoicesAPIKeyRef:        telnyx.String("voices_api_key_ref"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTtsIntegrationGet(t *testing.T) {
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
	_, err := client.TtsIntegrations.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTtsIntegrationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.TtsIntegrations.Update(
		context.TODO(),
		"id",
		telnyx.TtsIntegrationUpdateParams{
			UpdateCommunityTtsEngineRequest: telnyx.UpdateCommunityTtsEngineRequestParam{
				Capabilities: telnyx.CommunityTtsEngineCapabilitiesParam{
					StreamedAudio:         true,
					FlushSupport:          telnyx.Bool(true),
					ForceInterruptSupport: telnyx.Bool(true),
					SsmlSupport:           telnyx.Bool(true),
					VoiceSettingsSchema: map[string]any{
						"voice_speed": map[string]any{},
						"emotion":     map[string]any{},
					},
				},
				ConnectTimeoutMs:       telnyx.Int(500),
				ContactEmail:           telnyx.String("contact_email"),
				DataRetentionPolicyURL: telnyx.String("data_retention_policy_url"),
				DocumentationURL:       telnyx.String("https://"),
				SupportedAudioFormats:  []telnyx.CommunityTtsAudioFormat{telnyx.CommunityTtsAudioFormatMP3},
				Visibility:             telnyx.CommunityTtsEngineVisibilityPublic,
				VoicesAPIKeyRef:        telnyx.String("voices_api_key_ref"),
				VoicesURL:              telnyx.String("https://"),
				WebsocketURL:           telnyx.String("wss://"),
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

func TestTtsIntegrationListWithOptionalParams(t *testing.T) {
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
	_, err := client.TtsIntegrations.List(context.TODO(), telnyx.TtsIntegrationListParams{
		Page: telnyx.TtsIntegrationListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(20),
		},
		Status:     telnyx.CommunityTtsEngineStatusDraft,
		Visibility: telnyx.CommunityTtsEngineVisibilityPrivate,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTtsIntegrationDelete(t *testing.T) {
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
	err := client.TtsIntegrations.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTtsIntegrationReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.TtsIntegrations.Replace(
		context.TODO(),
		"id",
		telnyx.TtsIntegrationReplaceParams{
			UpdateCommunityTtsEngineRequest: telnyx.UpdateCommunityTtsEngineRequestParam{
				Capabilities: telnyx.CommunityTtsEngineCapabilitiesParam{
					StreamedAudio:         true,
					FlushSupport:          telnyx.Bool(true),
					ForceInterruptSupport: telnyx.Bool(true),
					SsmlSupport:           telnyx.Bool(true),
					VoiceSettingsSchema: map[string]any{
						"voice_speed": map[string]any{},
						"emotion":     map[string]any{},
					},
				},
				ConnectTimeoutMs:       telnyx.Int(500),
				ContactEmail:           telnyx.String("voices@acme.example"),
				DataRetentionPolicyURL: telnyx.String("data_retention_policy_url"),
				DocumentationURL:       telnyx.String("https://tts.acme.example/docs"),
				SupportedAudioFormats:  []telnyx.CommunityTtsAudioFormat{telnyx.CommunityTtsAudioFormatMP3},
				Visibility:             telnyx.CommunityTtsEngineVisibilityPrivate,
				VoicesAPIKeyRef:        telnyx.String("voices_api_key_ref"),
				VoicesURL:              telnyx.String("https://"),
				WebsocketURL:           telnyx.String("wss://"),
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

func TestTtsIntegrationGetConformanceRunsWithOptionalParams(t *testing.T) {
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
	_, err := client.TtsIntegrations.GetConformanceRuns(
		context.TODO(),
		"id",
		telnyx.TtsIntegrationGetConformanceRunsParams{
			Page: telnyx.TtsIntegrationGetConformanceRunsParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(20),
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

func TestTtsIntegrationGetDiscoverableWithOptionalParams(t *testing.T) {
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
	_, err := client.TtsIntegrations.GetDiscoverable(context.TODO(), telnyx.TtsIntegrationGetDiscoverableParams{
		Page: telnyx.TtsIntegrationGetDiscoverableParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(20),
		},
		Status:     telnyx.CommunityTtsEngineStatusDraft,
		Visibility: telnyx.CommunityTtsEngineVisibilityPrivate,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTtsIntegrationGetVoices(t *testing.T) {
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
	_, err := client.TtsIntegrations.GetVoices(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
