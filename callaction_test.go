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

func TestCallActionAnswerWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Answer(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionAnswerParams{
			BillingGroupID: telnyx.String("f5586561-8ff0-4291-a0ac-84fe544797bd"),
			ClientState:    telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:      telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			CustomHeaders: []telnyx.CustomSipHeaderParam{{
				Name:  "head_1",
				Value: "val_1",
			}, {
				Name:  "head_2",
				Value: "val_2",
			}},
			PreferredCodecs:      telnyx.CallActionAnswerParamsPreferredCodecsG722PcmuPcmaG729OpusVp8H264,
			Record:               telnyx.CallActionAnswerParamsRecordRecordFromAnswer,
			RecordChannels:       telnyx.CallActionAnswerParamsRecordChannelsSingle,
			RecordCustomFileName: telnyx.String("my_recording_file_name"),
			RecordFormat:         telnyx.CallActionAnswerParamsRecordFormatWav,
			RecordMaxLength:      telnyx.Int(1000),
			RecordTimeoutSecs:    telnyx.Int(100),
			RecordTrack:          telnyx.CallActionAnswerParamsRecordTrackOutbound,
			RecordTrim:           telnyx.CallActionAnswerParamsRecordTrimTrimSilence,
			SendSilenceWhenIdle:  telnyx.Bool(true),
			SipHeaders: []telnyx.SipHeaderParam{{
				Name:  telnyx.SipHeaderNameUserToUser,
				Value: "value",
			}},
			SoundModifications: telnyx.SoundModificationsParam{
				Octaves:  telnyx.Float(0.1),
				Pitch:    telnyx.Float(0),
				Semitone: telnyx.Float(-2),
				Track:    telnyx.String("both"),
			},
			StreamBidirectionalCodec:      telnyx.StreamBidirectionalCodecG722,
			StreamBidirectionalMode:       telnyx.StreamBidirectionalModeRtp,
			StreamBidirectionalTargetLegs: telnyx.StreamBidirectionalTargetLegsBoth,
			StreamCodec:                   telnyx.StreamCodecPcma,
			StreamTrack:                   telnyx.CallActionAnswerParamsStreamTrackBothTracks,
			StreamURL:                     telnyx.String("wss://www.example.com/websocket"),
			Transcription:                 telnyx.Bool(true),
			TranscriptionConfig: telnyx.TranscriptionStartRequestParam{
				ClientState:         telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
				CommandID:           telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
				TranscriptionEngine: telnyx.TranscriptionStartRequestTranscriptionEngineA,
				TranscriptionEngineConfig: telnyx.TranscriptionStartRequestTranscriptionEngineConfigUnionParam{
					OfA: &telnyx.TranscriptionEngineAConfigParam{
						EnableSpeakerDiarization: telnyx.Bool(true),
						Hints:                    []string{"Telnyx"},
						InterimResults:           telnyx.Bool(true),
						Language:                 telnyx.GoogleTranscriptionLanguageEn,
						MaxSpeakerCount:          telnyx.Int(4),
						MinSpeakerCount:          telnyx.Int(4),
						Model:                    telnyx.TranscriptionEngineAConfigModelLatestLong,
						ProfanityFilter:          telnyx.Bool(true),
						SpeechContext: []telnyx.TranscriptionEngineAConfigSpeechContextParam{{
							Boost:   telnyx.Float(1),
							Phrases: []string{"Telnyx"},
						}},
						TranscriptionEngine: telnyx.TranscriptionEngineAConfigTranscriptionEngineA,
						UseEnhanced:         telnyx.Bool(true),
					},
				},
				TranscriptionTracks: telnyx.String("both"),
			},
			WebhookURL:       telnyx.String("https://www.example.com/server-b/"),
			WebhookURLMethod: telnyx.CallActionAnswerParamsWebhookURLMethodPost,
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

func TestCallActionBridgeWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Bridge(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionBridgeParams{
			CallControlID:        "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
			ClientState:          telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:            telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			MuteDtmf:             telnyx.CallActionBridgeParamsMuteDtmfOpposite,
			ParkAfterUnbridge:    telnyx.String("self"),
			PlayRingtone:         telnyx.Bool(true),
			Queue:                telnyx.String("support"),
			Record:               telnyx.CallActionBridgeParamsRecordRecordFromAnswer,
			RecordChannels:       telnyx.CallActionBridgeParamsRecordChannelsSingle,
			RecordCustomFileName: telnyx.String("my_recording_file_name"),
			RecordFormat:         telnyx.CallActionBridgeParamsRecordFormatWav,
			RecordMaxLength:      telnyx.Int(1000),
			RecordTimeoutSecs:    telnyx.Int(100),
			RecordTrack:          telnyx.CallActionBridgeParamsRecordTrackOutbound,
			RecordTrim:           telnyx.CallActionBridgeParamsRecordTrimTrimSilence,
			Ringtone:             telnyx.CallActionBridgeParamsRingtonePl,
			VideoRoomContext:     telnyx.String("Alice"),
			VideoRoomID:          telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
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

func TestCallActionEnqueueWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Enqueue(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionEnqueueParams{
			QueueName:       "support",
			ClientState:     telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:       telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			MaxSize:         telnyx.Int(20),
			MaxWaitTimeSecs: telnyx.Int(600),
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

func TestCallActionGatherWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Gather(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionGatherParams{
			ClientState:             telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:               telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			GatherID:                telnyx.String("my_gather_id"),
			InitialTimeoutMillis:    telnyx.Int(10000),
			InterDigitTimeoutMillis: telnyx.Int(10000),
			MaximumDigits:           telnyx.Int(10),
			MinimumDigits:           telnyx.Int(1),
			TerminatingDigit:        telnyx.String("#"),
			TimeoutMillis:           telnyx.Int(60000),
			ValidDigits:             telnyx.String("123"),
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

func TestCallActionGatherUsingAIWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.GatherUsingAI(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionGatherUsingAIParams{
			Parameters: map[string]interface{}{
				"properties": map[string]interface{}{
					"age": map[string]interface{}{
						"description": "The age of the customer.",
						"type":        "integer",
					},
					"location": map[string]interface{}{
						"description": "The location of the customer.",
						"type":        "string",
					},
				},
				"required": map[string]interface{}{
					"0": "age",
					"1": "location",
				},
				"type": "object",
			},
			Assistant: telnyx.AssistantParam{
				Instructions:    telnyx.String("You are a friendly voice assistant."),
				Model:           telnyx.String("meta-llama/Meta-Llama-3.1-70B-Instruct"),
				OpenAIAPIKeyRef: telnyx.String("my_openai_api_key"),
				Tools: []telnyx.AssistantToolUnionParam{{
					OfBookAppointmentTool: &telnyx.AssistantToolBookAppointmentToolParam{
						BookAppointment: telnyx.AssistantToolBookAppointmentToolBookAppointmentParam{
							APIKeyRef:        "my_calcom_api_key",
							EventTypeID:      0,
							AttendeeName:     telnyx.String("attendee_name"),
							AttendeeTimezone: telnyx.String("attendee_timezone"),
						},
						Type: "book_appointment",
					},
				}},
			},
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Greeting:    telnyx.String("Hello, can you tell me your age and where you live?"),
			InterruptionSettings: telnyx.InterruptionSettingsParam{
				Enable: telnyx.Bool(true),
			},
			Language: telnyx.GoogleTranscriptionLanguageEn,
			MessageHistory: []telnyx.CallActionGatherUsingAIParamsMessageHistory{{
				Content: telnyx.String("Hello, what's your name?"),
				Role:    "assistant",
			}, {
				Content: telnyx.String("Hello, I'm John."),
				Role:    "user",
			}},
			SendMessageHistoryUpdates: telnyx.Bool(true),
			SendPartialResults:        telnyx.Bool(true),
			Transcription: telnyx.TranscriptionConfigParam{
				Model: telnyx.String("distil-whisper/distil-large-v2"),
			},
			UserResponseTimeoutMs: telnyx.Int(5000),
			Voice:                 telnyx.String("Telnyx.KokoroTTS.af"),
			VoiceSettings: telnyx.CallActionGatherUsingAIParamsVoiceSettingsUnion{
				OfElevenLabsVoiceSettings: &telnyx.ElevenLabsVoiceSettingsParam{
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

func TestCallActionGatherUsingAudioWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.GatherUsingAudio(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionGatherUsingAudioParams{
			AudioURL:                telnyx.String("http://example.com/message.wav"),
			ClientState:             telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:               telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			InterDigitTimeoutMillis: telnyx.Int(10000),
			InvalidAudioURL:         telnyx.String("http://example.com/message.wav"),
			InvalidMediaName:        telnyx.String("my_media_uploaded_to_media_storage_api"),
			MaximumDigits:           telnyx.Int(10),
			MaximumTries:            telnyx.Int(3),
			MediaName:               telnyx.String("my_media_uploaded_to_media_storage_api"),
			MinimumDigits:           telnyx.Int(1),
			TerminatingDigit:        telnyx.String("#"),
			TimeoutMillis:           telnyx.Int(10000),
			ValidDigits:             telnyx.String("123"),
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

func TestCallActionGatherUsingSpeakWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.GatherUsingSpeak(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionGatherUsingSpeakParams{
			Payload:                 "say this on call",
			Voice:                   "male",
			ClientState:             telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:               telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			InterDigitTimeoutMillis: telnyx.Int(10000),
			InvalidPayload:          telnyx.String("say this on call"),
			Language:                telnyx.CallActionGatherUsingSpeakParamsLanguageArb,
			MaximumDigits:           telnyx.Int(10),
			MaximumTries:            telnyx.Int(3),
			MinimumDigits:           telnyx.Int(1),
			PayloadType:             telnyx.CallActionGatherUsingSpeakParamsPayloadTypeText,
			ServiceLevel:            telnyx.CallActionGatherUsingSpeakParamsServiceLevelPremium,
			TerminatingDigit:        telnyx.String("#"),
			TimeoutMillis:           telnyx.Int(60000),
			ValidDigits:             telnyx.String("123"),
			VoiceSettings: telnyx.CallActionGatherUsingSpeakParamsVoiceSettingsUnion{
				OfElevenLabsVoiceSettings: &telnyx.ElevenLabsVoiceSettingsParam{
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

func TestCallActionHangupWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Hangup(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionHangupParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionLeaveQueueWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.LeaveQueue(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionLeaveQueueParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionPauseRecordingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.PauseRecording(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionPauseRecordingParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			RecordingID: telnyx.String("6e00ab49-9487-4364-8ad6-23965965afb2"),
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

func TestCallActionReferWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Refer(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionReferParams{
			SipAddress:  "sip:username@sip.non-telnyx-address.com",
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			CustomHeaders: []telnyx.CustomSipHeaderParam{{
				Name:  "head_1",
				Value: "val_1",
			}, {
				Name:  "head_2",
				Value: "val_2",
			}},
			SipAuthPassword: telnyx.String("sip_auth_password"),
			SipAuthUsername: telnyx.String("sip_auth_username"),
			SipHeaders: []telnyx.SipHeaderParam{{
				Name:  telnyx.SipHeaderNameUserToUser,
				Value: "value",
			}},
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

func TestCallActionRejectWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Reject(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionRejectParams{
			Cause:       telnyx.CallActionRejectParamsCauseUserBusy,
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionResumeRecordingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.ResumeRecording(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionResumeRecordingParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			RecordingID: telnyx.String("6e00ab49-9487-4364-8ad6-23965965afb2"),
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

func TestCallActionSendDtmfWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.SendDtmf(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionSendDtmfParams{
			Digits:         "1www2WABCDw9",
			ClientState:    telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:      telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			DurationMillis: telnyx.Int(500),
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

func TestCallActionSendSipInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.SendSipInfo(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionSendSipInfoParams{
			Body:        `{"key": "value", "numValue": 100}`,
			ContentType: "application/json",
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionSpeakWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Speak(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionSpeakParams{
			Payload:      "Say this on the call",
			Voice:        "female",
			ClientState:  telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:    telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Language:     telnyx.CallActionSpeakParamsLanguageArb,
			PayloadType:  telnyx.CallActionSpeakParamsPayloadTypeText,
			ServiceLevel: telnyx.CallActionSpeakParamsServiceLevelBasic,
			Stop:         telnyx.String("current"),
			VoiceSettings: telnyx.CallActionSpeakParamsVoiceSettingsUnion{
				OfElevenLabsVoiceSettings: &telnyx.ElevenLabsVoiceSettingsParam{
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

func TestCallActionStartAIAssistantWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartAIAssistant(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartAIAssistantParams{
			Assistant: telnyx.CallActionStartAIAssistantParamsAssistant{
				ID:              telnyx.String("id"),
				Instructions:    telnyx.String("You are a friendly voice assistant."),
				OpenAIAPIKeyRef: telnyx.String("openai_api_key_ref"),
			},
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Greeting:    telnyx.String("Hello, can you tell me your age and where you live?"),
			InterruptionSettings: telnyx.InterruptionSettingsParam{
				Enable: telnyx.Bool(true),
			},
			Transcription: telnyx.TranscriptionConfigParam{
				Model: telnyx.String("distil-whisper/distil-large-v2"),
			},
			Voice: telnyx.String("Telnyx.KokoroTTS.af"),
			VoiceSettings: telnyx.CallActionStartAIAssistantParamsVoiceSettingsUnion{
				OfElevenLabsVoiceSettings: &telnyx.ElevenLabsVoiceSettingsParam{
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

func TestCallActionStartForkingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartForking(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartForkingParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Rx:          telnyx.String("udp:192.0.2.1:9000"),
			StreamType:  telnyx.CallActionStartForkingParamsStreamTypeDecrypted,
			Tx:          telnyx.String("udp:192.0.2.1:9001"),
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

func TestCallActionStartNoiseSuppressionWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartNoiseSuppression(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartNoiseSuppressionParams{
			ClientState:            telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:              telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Direction:              telnyx.CallActionStartNoiseSuppressionParamsDirectionBoth,
			NoiseSuppressionEngine: telnyx.CallActionStartNoiseSuppressionParamsNoiseSuppressionEngineA,
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

func TestCallActionStartPlaybackWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartPlayback(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartPlaybackParams{
			AudioType:   telnyx.CallActionStartPlaybackParamsAudioTypeWav,
			AudioURL:    telnyx.String("http://www.example.com/sounds/greeting.wav"),
			CacheAudio:  telnyx.Bool(true),
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Loop: telnyx.LoopcountUnionParam{
				OfString: telnyx.String("infinity"),
			},
			MediaName:       telnyx.String("my_media_uploaded_to_media_storage_api"),
			Overlay:         telnyx.Bool(true),
			PlaybackContent: telnyx.String("SUQzAwAAAAADf1..."),
			Stop:            telnyx.String("current"),
			TargetLegs:      telnyx.String("self"),
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

func TestCallActionStartRecordingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartRecording(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartRecordingParams{
			Channels:                        telnyx.CallActionStartRecordingParamsChannelsSingle,
			Format:                          telnyx.CallActionStartRecordingParamsFormatWav,
			ClientState:                     telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:                       telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			CustomFileName:                  telnyx.String("my_recording_file_name"),
			MaxLength:                       telnyx.Int(0),
			PlayBeep:                        telnyx.Bool(true),
			RecordingTrack:                  telnyx.CallActionStartRecordingParamsRecordingTrackOutbound,
			TimeoutSecs:                     telnyx.Int(0),
			Transcription:                   telnyx.Bool(true),
			TranscriptionEngine:             telnyx.String("B"),
			TranscriptionLanguage:           telnyx.CallActionStartRecordingParamsTranscriptionLanguageEnUs,
			TranscriptionMaxSpeakerCount:    telnyx.Int(4),
			TranscriptionMinSpeakerCount:    telnyx.Int(4),
			TranscriptionProfanityFilter:    telnyx.Bool(true),
			TranscriptionSpeakerDiarization: telnyx.Bool(true),
			Trim:                            telnyx.CallActionStartRecordingParamsTrimTrimSilence,
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

func TestCallActionStartSiprecWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartSiprec(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartSiprecParams{
			ClientState:                  telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			ConnectorName:                telnyx.String("my-siprec-connector"),
			IncludeMetadataCustomHeaders: true,
			Secure:                       true,
			SessionTimeoutSecs:           telnyx.Int(900),
			SipTransport:                 telnyx.CallActionStartSiprecParamsSipTransportTcp,
			SiprecTrack:                  telnyx.CallActionStartSiprecParamsSiprecTrackBothTracks,
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

func TestCallActionStartStreamingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartStreaming(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartStreamingParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			DialogflowConfig: telnyx.DialogflowConfigParam{
				AnalyzeSentiment:           telnyx.Bool(false),
				PartialAutomatedAgentReply: telnyx.Bool(false),
			},
			EnableDialogflow:              telnyx.Bool(false),
			StreamBidirectionalCodec:      telnyx.StreamBidirectionalCodecG722,
			StreamBidirectionalMode:       telnyx.StreamBidirectionalModeRtp,
			StreamBidirectionalTargetLegs: telnyx.StreamBidirectionalTargetLegsBoth,
			StreamCodec:                   telnyx.StreamCodecPcma,
			StreamTrack:                   telnyx.CallActionStartStreamingParamsStreamTrackBothTracks,
			StreamURL:                     telnyx.String("wss://www.example.com/websocket"),
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

func TestCallActionStartTranscriptionWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StartTranscription(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStartTranscriptionParams{
			TranscriptionStartRequest: telnyx.TranscriptionStartRequestParam{
				ClientState:         telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
				CommandID:           telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
				TranscriptionEngine: telnyx.TranscriptionStartRequestTranscriptionEngineA,
				TranscriptionEngineConfig: telnyx.TranscriptionStartRequestTranscriptionEngineConfigUnionParam{
					OfA: &telnyx.TranscriptionEngineAConfigParam{
						EnableSpeakerDiarization: telnyx.Bool(true),
						Hints:                    []string{"Telnyx"},
						InterimResults:           telnyx.Bool(true),
						Language:                 telnyx.GoogleTranscriptionLanguageEn,
						MaxSpeakerCount:          telnyx.Int(4),
						MinSpeakerCount:          telnyx.Int(4),
						Model:                    telnyx.TranscriptionEngineAConfigModelLatestLong,
						ProfanityFilter:          telnyx.Bool(true),
						SpeechContext: []telnyx.TranscriptionEngineAConfigSpeechContextParam{{
							Boost:   telnyx.Float(1),
							Phrases: []string{"Telnyx"},
						}},
						TranscriptionEngine: telnyx.TranscriptionEngineAConfigTranscriptionEngineA,
						UseEnhanced:         telnyx.Bool(true),
					},
				},
				TranscriptionTracks: telnyx.String("both"),
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

func TestCallActionStopAIAssistantWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopAIAssistant(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopAIAssistantParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionStopForkingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopForking(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopForkingParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			StreamType:  telnyx.CallActionStopForkingParamsStreamTypeDecrypted,
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

func TestCallActionStopGatherWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopGather(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopGatherParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionStopNoiseSuppressionWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopNoiseSuppression(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopNoiseSuppressionParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionStopPlaybackWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopPlayback(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopPlaybackParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			Overlay:     telnyx.Bool(false),
			Stop:        telnyx.String("all"),
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

func TestCallActionStopRecordingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopRecording(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopRecordingParams{
			StopRecordingRequest: telnyx.StopRecordingRequestParam{
				ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
				CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
				RecordingID: telnyx.String("6e00ab49-9487-4364-8ad6-23965965afb2"),
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

func TestCallActionStopSiprecWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopSiprec(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopSiprecParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionStopStreamingWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopStreaming(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopStreamingParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			StreamID:    telnyx.String("1edb94f9-7ef0-4150-b502-e0ebadfd9491"),
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

func TestCallActionStopTranscriptionWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.StopTranscription(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionStopTranscriptionParams{
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
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

func TestCallActionSwitchSupervisorRole(t *testing.T) {
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
	_, err := client.Calls.Actions.SwitchSupervisorRole(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionSwitchSupervisorRoleParams{
			Role: telnyx.CallActionSwitchSupervisorRoleParamsRoleBarge,
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

func TestCallActionTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Actions.Transfer(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionTransferParams{
			To:                        "+18005550100 or sip:username@sip.telnyx.com",
			AnsweringMachineDetection: telnyx.CallActionTransferParamsAnsweringMachineDetectionDetect,
			AnsweringMachineDetectionConfig: telnyx.CallActionTransferParamsAnsweringMachineDetectionConfig{
				AfterGreetingSilenceMillis:      telnyx.Int(1000),
				BetweenWordsSilenceMillis:       telnyx.Int(1000),
				GreetingDurationMillis:          telnyx.Int(1000),
				GreetingSilenceDurationMillis:   telnyx.Int(2000),
				GreetingTotalAnalysisTimeMillis: telnyx.Int(50000),
				InitialSilenceMillis:            telnyx.Int(1000),
				MaximumNumberOfWords:            telnyx.Int(1000),
				MaximumWordLengthMillis:         telnyx.Int(2000),
				SilenceThreshold:                telnyx.Int(512),
				TotalAnalysisTimeMillis:         telnyx.Int(5000),
			},
			AudioURL:    telnyx.String("http://www.example.com/sounds/greeting.wav"),
			ClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:   telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			CustomHeaders: []telnyx.CustomSipHeaderParam{{
				Name:  "head_1",
				Value: "val_1",
			}, {
				Name:  "head_2",
				Value: "val_2",
			}},
			EarlyMedia:        telnyx.Bool(true),
			From:              telnyx.String("+18005550101"),
			FromDisplayName:   telnyx.String("Company Name"),
			MediaEncryption:   telnyx.CallActionTransferParamsMediaEncryptionSrtp,
			MediaName:         telnyx.String("my_media_uploaded_to_media_storage_api"),
			MuteDtmf:          telnyx.CallActionTransferParamsMuteDtmfOpposite,
			ParkAfterUnbridge: telnyx.String("self"),
			SipAuthPassword:   telnyx.String("password"),
			SipAuthUsername:   telnyx.String("username"),
			SipHeaders: []telnyx.SipHeaderParam{{
				Name:  telnyx.SipHeaderNameUserToUser,
				Value: "value",
			}},
			SipTransportProtocol: telnyx.CallActionTransferParamsSipTransportProtocolTls,
			SoundModifications: telnyx.SoundModificationsParam{
				Octaves:  telnyx.Float(0.1),
				Pitch:    telnyx.Float(0),
				Semitone: telnyx.Float(-2),
				Track:    telnyx.String("both"),
			},
			TargetLegClientState: telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			TimeLimitSecs:        telnyx.Int(600),
			TimeoutSecs:          telnyx.Int(60),
			WebhookURL:           telnyx.String("https://www.example.com/server-b/"),
			WebhookURLMethod:     telnyx.CallActionTransferParamsWebhookURLMethodPost,
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

func TestCallActionUpdateClientState(t *testing.T) {
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
	_, err := client.Calls.Actions.UpdateClientState(
		context.TODO(),
		"call_control_id",
		telnyx.CallActionUpdateClientStateParams{
			ClientState: "aGF2ZSBhIG5pY2UgZGF5ID1d",
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
