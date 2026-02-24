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

func TestTexmlAccountCallRecordingsJsonRecordingsJsonWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.RecordingsJson.RecordingsJson(
		context.TODO(),
		"call_sid",
		telnyx.TexmlAccountCallRecordingsJsonRecordingsJsonParams{
			AccountSid:                    "account_sid",
			PlayBeep:                      telnyx.Bool(false),
			RecordingChannels:             telnyx.TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingChannelsSingle,
			RecordingStatusCallback:       telnyx.String("http://webhook.com/callback"),
			RecordingStatusCallbackEvent:  telnyx.String("in-progress completed absent"),
			RecordingStatusCallbackMethod: telnyx.TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingStatusCallbackMethodGet,
			RecordingTrack:                telnyx.TexmlAccountCallRecordingsJsonRecordingsJsonParamsRecordingTrackInbound,
			SendRecordingURL:              telnyx.Bool(false),
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

func TestTexmlAccountCallRecordingsJsonGetRecordingsJson(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.RecordingsJson.GetRecordingsJson(
		context.TODO(),
		"call_sid",
		telnyx.TexmlAccountCallRecordingsJsonGetRecordingsJsonParams{
			AccountSid: "account_sid",
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
