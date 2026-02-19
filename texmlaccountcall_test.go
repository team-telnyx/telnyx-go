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

func TestTexmlAccountCallGet(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.Get(
		context.TODO(),
		"call_sid",
		telnyx.TexmlAccountCallGetParams{
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

func TestTexmlAccountCallUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.Update(
		context.TODO(),
		"call_sid",
		telnyx.TexmlAccountCallUpdateParams{
			AccountSid: "account_sid",
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

func TestTexmlAccountCallCallsWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.Calls(
		context.TODO(),
		"account_sid",
		telnyx.TexmlAccountCallCallsParams{
			ApplicationSid:                   "example-app-sid",
			From:                             "+13120001234",
			To:                               "+13121230000",
			AsyncAmd:                         telnyx.Bool(true),
			AsyncAmdStatusCallback:           telnyx.String("https://www.example.com/callback"),
			AsyncAmdStatusCallbackMethod:     telnyx.TexmlAccountCallCallsParamsAsyncAmdStatusCallbackMethodGet,
			CallerID:                         telnyx.String("Info"),
			CancelPlaybackOnDetectMessageEnd: telnyx.Bool(false),
			CancelPlaybackOnMachineDetection: telnyx.Bool(false),
			CustomHeaders: []telnyx.TexmlAccountCallCallsParamsCustomHeader{{
				Name:  "X-Custom-Header",
				Value: "custom-value",
			}},
			DetectionMode:                      telnyx.TexmlAccountCallCallsParamsDetectionModePremium,
			FallbackURL:                        telnyx.String("https://www.example.com/instructions-fallback.xml"),
			MachineDetection:                   telnyx.TexmlAccountCallCallsParamsMachineDetectionEnable,
			MachineDetectionSilenceTimeout:     telnyx.Int(2000),
			MachineDetectionSpeechEndThreshold: telnyx.Int(2000),
			MachineDetectionSpeechThreshold:    telnyx.Int(2000),
			MachineDetectionTimeout:            telnyx.Int(5000),
			PreferredCodecs:                    telnyx.String("PCMA,PCMU"),
			Record:                             telnyx.Bool(false),
			RecordingChannels:                  telnyx.TexmlAccountCallCallsParamsRecordingChannelsDual,
			RecordingStatusCallback:            telnyx.String("https://example.com/recording_status_callback"),
			RecordingStatusCallbackEvent:       telnyx.String("in-progress completed absent"),
			RecordingStatusCallbackMethod:      telnyx.TexmlAccountCallCallsParamsRecordingStatusCallbackMethodGet,
			RecordingTimeout:                   telnyx.Int(5),
			RecordingTrack:                     telnyx.TexmlAccountCallCallsParamsRecordingTrackInbound,
			SendRecordingURL:                   telnyx.Bool(false),
			SipAuthPassword:                    telnyx.String("1234"),
			SipAuthUsername:                    telnyx.String("user"),
			SipRegion:                          telnyx.TexmlAccountCallCallsParamsSipRegionCanada,
			StatusCallback:                     telnyx.String("https://www.example.com/statuscallback-listener"),
			StatusCallbackEvent:                telnyx.TexmlAccountCallCallsParamsStatusCallbackEventInitiated,
			StatusCallbackMethod:               telnyx.TexmlAccountCallCallsParamsStatusCallbackMethodGet,
			SuperviseCallSid:                   telnyx.String("v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"),
			SupervisingRole:                    telnyx.TexmlAccountCallCallsParamsSupervisingRoleMonitor,
			Texml:                              telnyx.String(`<?xml version="1.0" encoding="UTF-8"?><Response><Say>Hello</Say></Response>`),
			TimeLimit:                          telnyx.Int(3600),
			Timeout:                            telnyx.Int(60),
			Trim:                               telnyx.TexmlAccountCallCallsParamsTrimTrimSilence,
			URL:                                telnyx.String("https://www.example.com/texml.xml"),
			URLMethod:                          telnyx.TexmlAccountCallCallsParamsURLMethodGet,
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

func TestTexmlAccountCallGetCallsWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.GetCalls(
		context.TODO(),
		"account_sid",
		telnyx.TexmlAccountCallGetCallsParams{
			EndTime:     telnyx.String("EndTime"),
			EndTimeGt:   telnyx.String("EndTime_gt"),
			EndTimeLt:   telnyx.String("EndTime_lt"),
			From:        telnyx.String("From"),
			Page:        telnyx.Int(0),
			PageSize:    telnyx.Int(0),
			PageToken:   telnyx.String("PageToken"),
			StartTime:   telnyx.String("StartTime"),
			StartTimeGt: telnyx.String("StartTime_gt"),
			StartTimeLt: telnyx.String("StartTime_lt"),
			Status:      telnyx.TexmlAccountCallGetCallsParamsStatusCanceled,
			To:          telnyx.String("To"),
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

func TestTexmlAccountCallSiprecJsonWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.SiprecJson(
		context.TODO(),
		"call_sid",
		telnyx.TexmlAccountCallSiprecJsonParams{
			AccountSid:                   "account_sid",
			ConnectorName:                telnyx.String("my_connector"),
			IncludeMetadataCustomHeaders: true,
			Name:                         telnyx.String("my_siprec_session"),
			Secure:                       true,
			SessionTimeoutSecs:           telnyx.Int(900),
			SipTransport:                 telnyx.TexmlAccountCallSiprecJsonParamsSipTransportTcp,
			StatusCallback:               telnyx.String("https://www.example.com/callback"),
			StatusCallbackMethod:         telnyx.TexmlAccountCallSiprecJsonParamsStatusCallbackMethodGet,
			Track:                        telnyx.TexmlAccountCallSiprecJsonParamsTrackBothTracks,
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

func TestTexmlAccountCallStreamsJsonWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Calls.StreamsJson(
		context.TODO(),
		"call_sid",
		telnyx.TexmlAccountCallStreamsJsonParams{
			AccountSid:           "account_sid",
			BidirectionalCodec:   telnyx.TexmlAccountCallStreamsJsonParamsBidirectionalCodecG722,
			BidirectionalMode:    telnyx.TexmlAccountCallStreamsJsonParamsBidirectionalModeRtp,
			Name:                 telnyx.String("My stream"),
			StatusCallback:       telnyx.String("http://webhook.com/callback"),
			StatusCallbackMethod: telnyx.TexmlAccountCallStreamsJsonParamsStatusCallbackMethodGet,
			Track:                telnyx.TexmlAccountCallStreamsJsonParamsTrackBothTracks,
			URL:                  telnyx.String("wss://www.example.com/websocket"),
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
