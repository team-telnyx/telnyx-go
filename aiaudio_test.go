// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/internal/testutil"
	"github.com/team-telnyx/telnyx-go/option"
)

func TestAIAudioTranscribeWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Audio.Transcribe(context.TODO(), telnyx.AIAudioTranscribeParams{
		Model:                  telnyx.AIAudioTranscribeParamsModelDistilWhisperDistilLargeV2,
		File:                   io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		FileURL:                telnyx.String("https://example.com/file.mp3"),
		ResponseFormat:         telnyx.AIAudioTranscribeParamsResponseFormatJson,
		TimestampGranularities: telnyx.AIAudioTranscribeParamsTimestampGranularitiesSegment,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
