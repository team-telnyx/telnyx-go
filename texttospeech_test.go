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

func TestTextToSpeechGenerateSpeechWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToSpeech.GenerateSpeech(context.TODO(), telnyx.TextToSpeechGenerateSpeechParams{
		Aws: telnyx.TextToSpeechGenerateSpeechParamsAws{
			LanguageCode: telnyx.String("string"),
			LexiconNames: []string{"string"},
			OutputFormat: telnyx.String("string"),
			SampleRate:   telnyx.String("string"),
			TextType:     "text",
		},
		Azure: telnyx.TextToSpeechGenerateSpeechParamsAzure{
			APIKey:       telnyx.String("string"),
			DeploymentID: telnyx.String("string"),
			Effect:       telnyx.String("string"),
			Gender:       telnyx.String("string"),
			LanguageCode: telnyx.String("en-US"),
			OutputFormat: telnyx.String("audio-24khz-160kbitrate-mono-mp3"),
			Region:       telnyx.String("string"),
			TextType:     "text",
		},
		DisableCache: telnyx.Bool(false),
		Elevenlabs: telnyx.TextToSpeechGenerateSpeechParamsElevenlabs{
			APIKey:       telnyx.String("string"),
			LanguageCode: telnyx.String("string"),
			VoiceSettings: map[string]any{
				"foo": "bar",
			},
		},
		Humain: telnyx.TextToSpeechGenerateSpeechParamsHumain{
			VoiceID:       "sara-en",
			TtfbEagerness: telnyx.Float(0),
		},
		Language: telnyx.String("string"),
		Minimax: telnyx.TextToSpeechGenerateSpeechParamsMinimax{
			LanguageBoost:  telnyx.String("string"),
			Pitch:          telnyx.Int(0),
			ResponseFormat: telnyx.String("string"),
			Speed:          telnyx.Float(0),
			Vol:            telnyx.Float(0),
		},
		OutputType: telnyx.TextToSpeechGenerateSpeechParamsOutputTypeBinaryOutput,
		Provider:   telnyx.TextToSpeechGenerateSpeechParamsProviderAws,
		Resemble: telnyx.TextToSpeechGenerateSpeechParamsResemble{
			APIKey:     telnyx.String("string"),
			Format:     telnyx.String("string"),
			Precision:  telnyx.String("string"),
			SampleRate: telnyx.String("string"),
		},
		Telnyx: telnyx.TextToSpeechGenerateSpeechParamsTelnyx{
			Emotion:        "neutral",
			ResponseFormat: telnyx.String("mp3"),
			SamplingRate:   telnyx.Int(24000),
			VoiceSpeed:     telnyx.Float(1),
			Volume:         telnyx.Float(1),
		},
		Text:     telnyx.String("string"),
		TextType: telnyx.TextToSpeechGenerateSpeechParamsTextTypeText,
		Voice:    telnyx.String("string"),
		VoiceSettings: map[string]any{
			"foo": "bar",
		},
		Xai: telnyx.TextToSpeechGenerateSpeechParamsXai{
			VoiceID:      "eve",
			Language:     telnyx.String("auto"),
			OutputFormat: "mp3",
			SampleRate:   24000,
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

func TestTextToSpeechListVoicesWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToSpeech.ListVoices(context.TODO(), telnyx.TextToSpeechListVoicesParams{
		APIKey:   telnyx.String("api_key"),
		Provider: telnyx.TextToSpeechListVoicesParamsProviderAws,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTextToSpeechGetSpeechWithOptionalParams(t *testing.T) {
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
	err := client.TextToSpeech.GetSpeech(context.TODO(), telnyx.TextToSpeechGetSpeechParams{
		AudioFormat:  telnyx.TextToSpeechGetSpeechParamsAudioFormatPcm,
		DisableCache: telnyx.Bool(true),
		ModelID:      telnyx.String("model_id"),
		Provider:     telnyx.TextToSpeechGetSpeechParamsProviderAws,
		SocketID:     telnyx.String("socket_id"),
		Voice:        telnyx.String("voice"),
		VoiceID:      telnyx.String("voice_id"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
