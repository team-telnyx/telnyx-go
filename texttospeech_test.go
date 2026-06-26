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
			LanguageCode: telnyx.String("language_code"),
			LexiconNames: []string{"string"},
			OutputFormat: telnyx.String("output_format"),
			SampleRate:   telnyx.String("sample_rate"),
			TextType:     "text",
		},
		Azure: telnyx.TextToSpeechGenerateSpeechParamsAzure{
			APIKey:       telnyx.String("api_key"),
			DeploymentID: telnyx.String("deployment_id"),
			Effect:       telnyx.String("effect"),
			Gender:       telnyx.String("gender"),
			LanguageCode: telnyx.String("language_code"),
			OutputFormat: telnyx.String("output_format"),
			Region:       telnyx.String("region"),
			TextType:     "text",
		},
		DisableCache: telnyx.Bool(true),
		Elevenlabs: telnyx.TextToSpeechGenerateSpeechParamsElevenlabs{
			APIKey:       telnyx.String("api_key"),
			LanguageCode: telnyx.String("language_code"),
			VoiceSettings: map[string]any{
				"foo": "bar",
			},
		},
		Language: telnyx.String("language"),
		Minimax: telnyx.TextToSpeechGenerateSpeechParamsMinimax{
			LanguageBoost:  telnyx.String("language_boost"),
			Pitch:          telnyx.Int(0),
			ResponseFormat: telnyx.String("response_format"),
			Speed:          telnyx.Float(0),
			Vol:            telnyx.Float(0),
		},
		OutputType: telnyx.TextToSpeechGenerateSpeechParamsOutputTypeBinaryOutput,
		Provider:   telnyx.TextToSpeechGenerateSpeechParamsProviderAws,
		Resemble: telnyx.TextToSpeechGenerateSpeechParamsResemble{
			APIKey:     telnyx.String("api_key"),
			Format:     telnyx.String("format"),
			Precision:  telnyx.String("precision"),
			SampleRate: telnyx.String("sample_rate"),
		},
		Rime: telnyx.TextToSpeechGenerateSpeechParamsRime{
			ResponseFormat: telnyx.String("response_format"),
			SamplingRate:   telnyx.Int(0),
			VoiceSpeed:     telnyx.Float(0),
		},
		Telnyx: telnyx.TextToSpeechGenerateSpeechParamsTelnyx{
			Emotion:        "neutral",
			ResponseFormat: telnyx.String("response_format"),
			SamplingRate:   telnyx.Int(0),
			Temperature:    telnyx.Float(0),
			VoiceSpeed:     telnyx.Float(0.5),
			Volume:         telnyx.Float(0),
		},
		Text:     telnyx.String("text"),
		TextType: telnyx.TextToSpeechGenerateSpeechParamsTextTypeText,
		Voice:    telnyx.String("voice"),
		VoiceSettings: map[string]any{
			"foo": "bar",
		},
		Xai: telnyx.TextToSpeechGenerateSpeechParamsXai{
			VoiceID:      "eve",
			Language:     telnyx.String("language"),
			OutputFormat: "mp3",
			SampleRate:   8000,
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
