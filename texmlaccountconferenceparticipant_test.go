// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestTexmlAccountConferenceParticipantGet(t *testing.T) {
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
	_, err := client.Texml.Accounts.Conferences.Participants.Get(
		context.TODO(),
		"call_sid_or_participant_label",
		telnyx.TexmlAccountConferenceParticipantGetParams{
			AccountSid:    "account_sid",
			ConferenceSid: "conference_sid",
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

func TestTexmlAccountConferenceParticipantUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Conferences.Participants.Update(
		context.TODO(),
		"call_sid_or_participant_label",
		telnyx.TexmlAccountConferenceParticipantUpdateParams{
			AccountSid:          "account_sid",
			ConferenceSid:       "conference_sid",
			AnnounceMethod:      telnyx.TexmlAccountConferenceParticipantUpdateParamsAnnounceMethodGet,
			AnnounceURL:         telnyx.String("https://www.example.com/announce.xml"),
			BeepOnExit:          telnyx.Bool(false),
			CallSidToCoach:      telnyx.String("v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ"),
			Coaching:            telnyx.Bool(false),
			EndConferenceOnExit: telnyx.Bool(false),
			Hold:                telnyx.Bool(true),
			HoldMethod:          telnyx.TexmlAccountConferenceParticipantUpdateParamsHoldMethodPost,
			HoldURL:             telnyx.String("HoldUrl"),
			Muted:               telnyx.Bool(true),
			WaitURL:             telnyx.String("https://www.example.com/wait_music.mp3"),
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

func TestTexmlAccountConferenceParticipantDelete(t *testing.T) {
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
	err := client.Texml.Accounts.Conferences.Participants.Delete(
		context.TODO(),
		"call_sid_or_participant_label",
		telnyx.TexmlAccountConferenceParticipantDeleteParams{
			AccountSid:    "account_sid",
			ConferenceSid: "conference_sid",
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

func TestTexmlAccountConferenceParticipantParticipantsWithOptionalParams(t *testing.T) {
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
	_, err := client.Texml.Accounts.Conferences.Participants.Participants(
		context.TODO(),
		"conference_sid",
		telnyx.TexmlAccountConferenceParticipantParticipantsParams{
			AccountSid:                              "account_sid",
			AmdStatusCallback:                       telnyx.String("https://www.example.com/amd_result"),
			AmdStatusCallbackMethod:                 telnyx.TexmlAccountConferenceParticipantParticipantsParamsAmdStatusCallbackMethodGet,
			Beep:                                    telnyx.TexmlAccountConferenceParticipantParticipantsParamsBeepOnExit,
			CallerID:                                telnyx.String("Info"),
			CallSidToCoach:                          telnyx.String("v3:9X2vxPDFY2RHSJ1EdMS0RHRksMTg7ldNxdjWbVr9zBjbGjGsSe-aiQ"),
			CancelPlaybackOnDetectMessageEnd:        telnyx.Bool(false),
			CancelPlaybackOnMachineDetection:        telnyx.Bool(false),
			Coaching:                                telnyx.Bool(false),
			ConferenceRecord:                        telnyx.TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordRecordFromStart,
			ConferenceRecordingStatusCallback:       telnyx.String("https://example.com/conference_recording_status_callback"),
			ConferenceRecordingStatusCallbackEvent:  telnyx.String("in-progress completed failed absent"),
			ConferenceRecordingStatusCallbackMethod: telnyx.TexmlAccountConferenceParticipantParticipantsParamsConferenceRecordingStatusCallbackMethodGet,
			ConferenceRecordingTimeout:              telnyx.Int(5),
			ConferenceStatusCallback:                telnyx.String("https://example.com/conference_status_callback"),
			ConferenceStatusCallbackEvent:           telnyx.String("start end join leave"),
			ConferenceStatusCallbackMethod:          telnyx.TexmlAccountConferenceParticipantParticipantsParamsConferenceStatusCallbackMethodGet,
			ConferenceTrim:                          telnyx.TexmlAccountConferenceParticipantParticipantsParamsConferenceTrimTrimSilence,
			EarlyMedia:                              telnyx.Bool(true),
			EndConferenceOnExit:                     telnyx.Bool(true),
			From:                                    telnyx.String("+12065550200"),
			MachineDetection:                        telnyx.TexmlAccountConferenceParticipantParticipantsParamsMachineDetectionEnable,
			MachineDetectionSilenceTimeout:          telnyx.Int(2000),
			MachineDetectionSpeechEndThreshold:      telnyx.Int(2000),
			MachineDetectionSpeechThreshold:         telnyx.Int(2000),
			MachineDetectionTimeout:                 telnyx.Int(1000),
			MaxParticipants:                         telnyx.Int(30),
			Muted:                                   telnyx.Bool(true),
			PreferredCodecs:                         telnyx.String("PCMA,PCMU"),
			Record:                                  telnyx.Bool(false),
			RecordingChannels:                       telnyx.TexmlAccountConferenceParticipantParticipantsParamsRecordingChannelsDual,
			RecordingStatusCallback:                 telnyx.String("https://example.com/recording_status_callback"),
			RecordingStatusCallbackEvent:            telnyx.String("in-progress completed absent"),
			RecordingStatusCallbackMethod:           telnyx.TexmlAccountConferenceParticipantParticipantsParamsRecordingStatusCallbackMethodGet,
			RecordingTrack:                          telnyx.TexmlAccountConferenceParticipantParticipantsParamsRecordingTrackInbound,
			SipAuthPassword:                         telnyx.String("1234"),
			SipAuthUsername:                         telnyx.String("user"),
			StartConferenceOnEnter:                  telnyx.Bool(false),
			StatusCallback:                          telnyx.String("https://www.example.com/callback"),
			StatusCallbackEvent:                     telnyx.String("answered completed"),
			StatusCallbackMethod:                    telnyx.TexmlAccountConferenceParticipantParticipantsParamsStatusCallbackMethodGet,
			TimeLimit:                               telnyx.Int(30),
			TimeoutSeconds:                          telnyx.Int(30),
			To:                                      telnyx.String("+12065550100"),
			Trim:                                    telnyx.TexmlAccountConferenceParticipantParticipantsParamsTrimTrimSilence,
			WaitURL:                                 telnyx.String("https://www.example.com/wait_music.mp3"),
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

func TestTexmlAccountConferenceParticipantGetParticipants(t *testing.T) {
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
	_, err := client.Texml.Accounts.Conferences.Participants.GetParticipants(
		context.TODO(),
		"conference_sid",
		telnyx.TexmlAccountConferenceParticipantGetParticipantsParams{
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
