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

func TestConferenceActionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Update(
		context.TODO(),
		"id",
		telnyx.ConferenceActionUpdateParams{
			UpdateConference: telnyx.UpdateConferenceParam{
				CallControlID:         "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
				SupervisorRole:        telnyx.UpdateConferenceSupervisorRoleWhisper,
				CommandID:             telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
				Region:                telnyx.UpdateConferenceRegionUs,
				WhisperCallControlIDs: []string{"v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ", "v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw"},
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

func TestConferenceActionEndConferenceWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.EndConference(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.ConferenceActionEndConferenceParams{
			CommandID: telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestConferenceActionGatherDtmfAudioWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.GatherDtmfAudio(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.ConferenceActionGatherDtmfAudioParams{
			CallControlID:           "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			AudioURL:                telnyx.String("http://example.com/gather_prompt.wav"),
			ClientState:             telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			GatherID:                telnyx.String("gather_id"),
			InitialTimeoutMillis:    telnyx.Int(10000),
			InterDigitTimeoutMillis: telnyx.Int(3000),
			InvalidAudioURL:         telnyx.String("invalid_audio_url"),
			InvalidMediaName:        telnyx.String("invalid_media_name"),
			MaximumDigits:           telnyx.Int(4),
			MaximumTries:            telnyx.Int(3),
			MediaName:               telnyx.String("media_name"),
			MinimumDigits:           telnyx.Int(1),
			StopPlaybackOnDtmf:      telnyx.Bool(true),
			TerminatingDigit:        telnyx.String("#"),
			TimeoutMillis:           telnyx.Int(30000),
			ValidDigits:             telnyx.String("0123456789"),
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

func TestConferenceActionHoldWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Hold(
		context.TODO(),
		"id",
		telnyx.ConferenceActionHoldParams{
			AudioURL:       telnyx.String("http://example.com/message.wav"),
			CallControlIDs: []string{"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"},
			MediaName:      telnyx.String("my_media_uploaded_to_media_storage_api"),
			Region:         telnyx.ConferenceActionHoldParamsRegionUs,
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

func TestConferenceActionJoinWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Join(
		context.TODO(),
		"id",
		telnyx.ConferenceActionJoinParams{
			CallControlID:           "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			BeepEnabled:             telnyx.ConferenceActionJoinParamsBeepEnabledAlways,
			ClientState:             telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:               telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			EndConferenceOnExit:     telnyx.Bool(true),
			Hold:                    telnyx.Bool(true),
			HoldAudioURL:            telnyx.String("http://www.example.com/audio.wav"),
			HoldMediaName:           telnyx.String("my_media_uploaded_to_media_storage_api"),
			Mute:                    telnyx.Bool(true),
			Region:                  telnyx.ConferenceActionJoinParamsRegionUs,
			SoftEndConferenceOnExit: telnyx.Bool(true),
			StartConferenceOnEnter:  telnyx.Bool(true),
			SupervisorRole:          telnyx.ConferenceActionJoinParamsSupervisorRoleWhisper,
			WhisperCallControlIDs:   []string{"v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ", "v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw"},
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

func TestConferenceActionLeaveWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Leave(
		context.TODO(),
		"id",
		telnyx.ConferenceActionLeaveParams{
			CallControlID: "c46e06d7-b78f-4b13-96b6-c576af9640ff",
			BeepEnabled:   telnyx.ConferenceActionLeaveParamsBeepEnabledNever,
			CommandID:     telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Region:        telnyx.ConferenceActionLeaveParamsRegionUs,
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

func TestConferenceActionMuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Mute(
		context.TODO(),
		"id",
		telnyx.ConferenceActionMuteParams{
			CallControlIDs: []string{"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"},
			Region:         telnyx.ConferenceActionMuteParamsRegionUs,
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

func TestConferenceActionPlayWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Play(
		context.TODO(),
		"id",
		telnyx.ConferenceActionPlayParams{
			AudioURL:       telnyx.String("http://www.example.com/sounds/greeting.wav"),
			CallControlIDs: []string{"string"},
			Loop: telnyx.LoopcountUnionParam{
				OfString: telnyx.String("infinity"),
			},
			MediaName: telnyx.String("my_media_uploaded_to_media_storage_api"),
			Region:    telnyx.ConferenceActionPlayParamsRegionUs,
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

func TestConferenceActionRecordPauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.RecordPause(
		context.TODO(),
		"id",
		telnyx.ConferenceActionRecordPauseParams{
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			RecordingID: telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Region:      telnyx.ConferenceActionRecordPauseParamsRegionUs,
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

func TestConferenceActionRecordResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.RecordResume(
		context.TODO(),
		"id",
		telnyx.ConferenceActionRecordResumeParams{
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			RecordingID: telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Region:      telnyx.ConferenceActionRecordResumeParamsRegionUs,
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

func TestConferenceActionRecordStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.RecordStart(
		context.TODO(),
		"id",
		telnyx.ConferenceActionRecordStartParams{
			Format:         telnyx.ConferenceActionRecordStartParamsFormatWav,
			CommandID:      telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			CustomFileName: telnyx.String("my_recording_file_name"),
			PlayBeep:       telnyx.Bool(true),
			Region:         telnyx.ConferenceActionRecordStartParamsRegionUs,
			Trim:           telnyx.ConferenceActionRecordStartParamsTrimTrimSilence,
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

func TestConferenceActionRecordStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.RecordStop(
		context.TODO(),
		"id",
		telnyx.ConferenceActionRecordStopParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			RecordingID: telnyx.String("6e00ab49-9487-4364-8ad6-23965965afb2"),
			Region:      telnyx.ConferenceActionRecordStopParamsRegionUs,
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

func TestConferenceActionSendDtmfWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.SendDtmf(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.ConferenceActionSendDtmfParams{
			Digits:         "1234#",
			CallControlIDs: []string{"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"},
			ClientState:    telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			DurationMillis: telnyx.Int(250),
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

func TestConferenceActionSpeakWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Speak(
		context.TODO(),
		"id",
		telnyx.ConferenceActionSpeakParams{
			Payload:        "Say this to participants",
			Voice:          "female",
			CallControlIDs: []string{"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"},
			CommandID:      telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Language:       telnyx.ConferenceActionSpeakParamsLanguageEnUs,
			PayloadType:    telnyx.ConferenceActionSpeakParamsPayloadTypeText,
			Region:         telnyx.ConferenceActionSpeakParamsRegionUs,
			VoiceSettings: telnyx.ConferenceActionSpeakParamsVoiceSettingsUnion{
				OfElevenlabs: &telnyx.ElevenLabsVoiceSettingsParam{
					Type:      telnyx.ElevenLabsVoiceSettingsTypeElevenlabs,
					APIKeyRef: telnyx.String("my_elevenlabs_api_key"),
				},
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

func TestConferenceActionStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Stop(
		context.TODO(),
		"id",
		telnyx.ConferenceActionStopParams{
			CallControlIDs: []string{"string"},
			Region:         telnyx.ConferenceActionStopParamsRegionUs,
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

func TestConferenceActionUnholdWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Unhold(
		context.TODO(),
		"id",
		telnyx.ConferenceActionUnholdParams{
			CallControlIDs: []string{"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"},
			Region:         telnyx.ConferenceActionUnholdParamsRegionUs,
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

func TestConferenceActionUnmuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Actions.Unmute(
		context.TODO(),
		"id",
		telnyx.ConferenceActionUnmuteParams{
			CallControlIDs: []string{"v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"},
			Region:         telnyx.ConferenceActionUnmuteParamsRegionUs,
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
