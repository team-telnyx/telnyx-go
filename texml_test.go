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

func TestTexmlInitiateAICallWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.InitiateAICall(
		context.TODO(),
		"connection_id",
		telnyx.TexmlInitiateAICallParams{
			AIAssistantID: "ai-assistant-id-123",
			From:          "+13120001234",
			To:            "+13121230000",
			AIAssistantDynamicVariables: map[string]string{
				"customer_name": "John Doe",
				"account_id":    "12345",
			},
			AIAssistantVersion:           telnyx.String("AIAssistantVersion"),
			AsyncAmd:                     telnyx.Bool(true),
			AsyncAmdStatusCallback:       telnyx.String("https://www.example.com/callback"),
			AsyncAmdStatusCallbackMethod: telnyx.TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethodGet,
			CallerID:                     telnyx.String("Info"),
			ConversationCallback:         telnyx.String("https://www.example.com/conversation-callback"),
			ConversationCallbackMethod:   telnyx.TexmlInitiateAICallParamsConversationCallbackMethodGet,
			ConversationCallbacks:        []string{"https://www.example.com/conversation-callback1", "https://www.example.com/conversation-callback2"},
			CustomHeaders: []telnyx.TexmlInitiateAICallParamsCustomHeader{{
				Name:  "X-Custom-Header",
				Value: "custom-value",
			}},
			DetectionMode:                      telnyx.TexmlInitiateAICallParamsDetectionModePremium,
			MachineDetection:                   telnyx.TexmlInitiateAICallParamsMachineDetectionEnable,
			MachineDetectionSilenceTimeout:     telnyx.Int(2000),
			MachineDetectionSpeechEndThreshold: telnyx.Int(2000),
			MachineDetectionSpeechThreshold:    telnyx.Int(2000),
			MachineDetectionTimeout:            telnyx.Int(5000),
			Passports:                          telnyx.String("Passports"),
			PreferredCodecs:                    telnyx.String("PCMA,PCMU"),
			Record:                             telnyx.Bool(false),
			RecordingChannels:                  telnyx.TexmlInitiateAICallParamsRecordingChannelsDual,
			RecordingStatusCallback:            telnyx.String("https://example.com/recording_status_callback"),
			RecordingStatusCallbackEvent:       telnyx.String("in-progress completed absent"),
			RecordingStatusCallbackMethod:      telnyx.TexmlInitiateAICallParamsRecordingStatusCallbackMethodGet,
			RecordingTimeout:                   telnyx.Int(5),
			RecordingTrack:                     telnyx.TexmlInitiateAICallParamsRecordingTrackInbound,
			SendRecordingURL:                   telnyx.Bool(false),
			SipAuthPassword:                    telnyx.String("1234"),
			SipAuthUsername:                    telnyx.String("user"),
			SipRegion:                          telnyx.TexmlInitiateAICallParamsSipRegionCanada,
			StatusCallback:                     telnyx.String("https://www.example.com/callback"),
			StatusCallbackEvent:                telnyx.String("initiated answered"),
			StatusCallbackMethod:               telnyx.TexmlInitiateAICallParamsStatusCallbackMethodGet,
			StatusCallbacks:                    []string{"https://www.example.com/callback1", "https://www.example.com/callback2"},
			TimeLimit:                          telnyx.Int(3600),
			TimeoutSeconds:                     telnyx.Int(60),
			Trim:                               telnyx.TexmlInitiateAICallParamsTrimTrimSilence,
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

func TestTexmlSecrets(t *testing.T) {
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
	_, err := client.Texml.Secrets(context.TODO(), telnyx.TexmlSecretsParams{
		Name:  "My Secret Name",
		Value: "My Secret Value",
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
