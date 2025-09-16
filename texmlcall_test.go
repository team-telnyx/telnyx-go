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

func TestTexmlCallUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Calls.Update(
		context.TODO(),
		"call_sid",
		telnyx.TexmlCallUpdateParams{
			UpdateCall: telnyx.UpdateCallParam{
				FallbackMethod:       telnyx.UpdateCallFallbackMethodGet,
				FallbackURL:          telnyx.String("https://www.example.com/intruction-c.xml"),
				Method:               telnyx.UpdateCallMethodGet,
				Status:               telnyx.String("completed"),
				StatusCallback:       telnyx.String("https://www.example.com/callback"),
				StatusCallbackMethod: telnyx.UpdateCallStatusCallbackMethodGet,
				Texml:                telnyx.String(`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello</Say></Response>`),
				URL:                  telnyx.String("https://www.example.com/intruction-b.xml"),
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

func TestTexmlCallInitiateWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Calls.Initiate(
		context.TODO(),
		"application_id",
		telnyx.TexmlCallInitiateParams{
			From:                               "+13120001234",
			To:                                 "+13121230000",
			AsyncAmd:                           telnyx.Bool(true),
			AsyncAmdStatusCallback:             telnyx.String("https://www.example.com/callback"),
			AsyncAmdStatusCallbackMethod:       telnyx.TexmlCallInitiateParamsAsyncAmdStatusCallbackMethodGet,
			CallerID:                           telnyx.String("Info"),
			CancelPlaybackOnDetectMessageEnd:   telnyx.Bool(false),
			CancelPlaybackOnMachineDetection:   telnyx.Bool(false),
			DetectionMode:                      telnyx.TexmlCallInitiateParamsDetectionModePremium,
			FallbackURL:                        telnyx.String("https://www.example.com/instructions-fallback.xml"),
			MachineDetection:                   telnyx.TexmlCallInitiateParamsMachineDetectionEnable,
			MachineDetectionSilenceTimeout:     telnyx.Int(2000),
			MachineDetectionSpeechEndThreshold: telnyx.Int(2000),
			MachineDetectionSpeechThreshold:    telnyx.Int(2000),
			MachineDetectionTimeout:            telnyx.Int(5000),
			PreferredCodecs:                    telnyx.String("PCMA,PCMU"),
			Record:                             telnyx.Bool(false),
			RecordingChannels:                  telnyx.TexmlCallInitiateParamsRecordingChannelsDual,
			RecordingStatusCallback:            telnyx.String("https://example.com/recording_status_callback"),
			RecordingStatusCallbackEvent:       telnyx.String("in-progress completed absent"),
			RecordingStatusCallbackMethod:      telnyx.TexmlCallInitiateParamsRecordingStatusCallbackMethodGet,
			RecordingTimeout:                   telnyx.Int(5),
			RecordingTrack:                     telnyx.TexmlCallInitiateParamsRecordingTrackInbound,
			SipAuthPassword:                    telnyx.String("1234"),
			SipAuthUsername:                    telnyx.String("user"),
			StatusCallback:                     telnyx.String("https://www.example.com/statuscallback-listener"),
			StatusCallbackEvent:                telnyx.TexmlCallInitiateParamsStatusCallbackEventInitiated,
			StatusCallbackMethod:               telnyx.TexmlCallInitiateParamsStatusCallbackMethodGet,
			Trim:                               telnyx.TexmlCallInitiateParamsTrimTrimSilence,
			URL:                                telnyx.String("https://www.example.com/texml.xml"),
			URLMethod:                          telnyx.TexmlCallInitiateParamsURLMethodGet,
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
