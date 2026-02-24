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

func TestCallDialWithOptionalParams(t *testing.T) {
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
	_, err := client.Calls.Dial(context.TODO(), telnyx.CallDialParams{
		ConnectionID: "7267xxxxxxxxxxxxxx",
		From:         "+18005550101",
		To: telnyx.CallDialParamsToUnion{
			OfString: telnyx.String("+18005550100 or sip:username@sip.telnyx.com"),
		},
		AnsweringMachineDetection: telnyx.CallDialParamsAnsweringMachineDetectionDetect,
		AnsweringMachineDetectionConfig: telnyx.CallDialParamsAnsweringMachineDetectionConfig{
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
		AudioURL:       telnyx.String("http://www.example.com/sounds/greeting.wav"),
		BillingGroupID: telnyx.String("f5586561-8ff0-4291-a0ac-84fe544797bd"),
		BridgeIntent:   telnyx.Bool(true),
		BridgeOnAnswer: telnyx.Bool(true),
		ClientState:    telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
		CommandID:      telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
		ConferenceConfig: telnyx.CallDialParamsConferenceConfig{
			ID:                      telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			BeepEnabled:             "on_exit",
			ConferenceName:          telnyx.String("telnyx-conference"),
			EarlyMedia:              telnyx.Bool(false),
			EndConferenceOnExit:     telnyx.Bool(true),
			Hold:                    telnyx.Bool(true),
			HoldAudioURL:            telnyx.String("http://example.com/message.wav"),
			HoldMediaName:           telnyx.String("my_media_uploaded_to_media_storage_api"),
			Mute:                    telnyx.Bool(true),
			SoftEndConferenceOnExit: telnyx.Bool(true),
			StartConferenceOnCreate: telnyx.Bool(false),
			StartConferenceOnEnter:  telnyx.Bool(true),
			SupervisorRole:          "whisper",
			WhisperCallControlIDs:   []string{"v2:Sg1xxxQ_U3ixxxyXT_VDNI3xxxazZdg6Vxxxs4-GNYxxxVaJPOhFMRQ", "v2:qqpb0mmvd-ovhhBr0BUQQn0fld5jIboaaX3-De0DkqXHzbf8d75xkw"},
		},
		CustomHeaders: []telnyx.CustomSipHeaderParam{{
			Name:  "head_1",
			Value: "val_1",
		}, {
			Name:  "head_2",
			Value: "val_2",
		}},
		DialogflowConfig: telnyx.DialogflowConfigParam{
			AnalyzeSentiment:           telnyx.Bool(false),
			PartialAutomatedAgentReply: telnyx.Bool(false),
		},
		EnableDialogflow:     telnyx.Bool(false),
		FromDisplayName:      telnyx.String("Company Name"),
		LinkTo:               telnyx.String("ilditnZK_eVysupV21KzmzN_sM29ygfauQojpm4BgFtfX5hXAcjotg=="),
		MediaEncryption:      telnyx.CallDialParamsMediaEncryptionSrtp,
		MediaName:            telnyx.String("my_media_uploaded_to_media_storage_api"),
		ParkAfterUnbridge:    telnyx.String("self"),
		PreferredCodecs:      telnyx.String("G722,PCMU,PCMA,G729,OPUS,VP8,H264"),
		Record:               telnyx.CallDialParamsRecordRecordFromAnswer,
		RecordChannels:       telnyx.CallDialParamsRecordChannelsSingle,
		RecordCustomFileName: telnyx.String("my_recording_file_name"),
		RecordFormat:         telnyx.CallDialParamsRecordFormatWav,
		RecordMaxLength:      telnyx.Int(1000),
		RecordTimeoutSecs:    telnyx.Int(100),
		RecordTrack:          telnyx.CallDialParamsRecordTrackOutbound,
		RecordTrim:           telnyx.CallDialParamsRecordTrimTrimSilence,
		SendSilenceWhenIdle:  telnyx.Bool(true),
		SipAuthPassword:      telnyx.String("password"),
		SipAuthUsername:      telnyx.String("username"),
		SipHeaders: []telnyx.SipHeaderParam{{
			Name:  telnyx.SipHeaderNameUserToUser,
			Value: "12345",
		}},
		SipRegion:            telnyx.CallDialParamsSipRegionCanada,
		SipTransportProtocol: telnyx.CallDialParamsSipTransportProtocolTls,
		SoundModifications: telnyx.SoundModificationsParam{
			Octaves:  telnyx.Float(0.1),
			Pitch:    telnyx.Float(0.8),
			Semitone: telnyx.Float(-2),
			Track:    telnyx.String("both"),
		},
		StreamAuthToken:                    telnyx.String("your-auth-token"),
		StreamBidirectionalCodec:           telnyx.StreamBidirectionalCodecG722,
		StreamBidirectionalMode:            telnyx.StreamBidirectionalModeRtp,
		StreamBidirectionalSamplingRate:    16000,
		StreamBidirectionalTargetLegs:      telnyx.StreamBidirectionalTargetLegsBoth,
		StreamCodec:                        telnyx.StreamCodecPcma,
		StreamEstablishBeforeCallOriginate: telnyx.Bool(true),
		StreamTrack:                        telnyx.CallDialParamsStreamTrackBothTracks,
		StreamURL:                          telnyx.String("wss://www.example.com/websocket"),
		SuperviseCallControlID:             telnyx.String("v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg"),
		SupervisorRole:                     telnyx.CallDialParamsSupervisorRoleBarge,
		TimeLimitSecs:                      telnyx.Int(60),
		TimeoutSecs:                        telnyx.Int(60),
		Transcription:                      telnyx.Bool(true),
		TranscriptionConfig: telnyx.TranscriptionStartRequestParam{
			ClientState:         telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
			CommandID:           telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
			TranscriptionEngine: telnyx.TranscriptionStartRequestTranscriptionEngineGoogle,
			TranscriptionEngineConfig: telnyx.TranscriptionStartRequestTranscriptionEngineConfigUnionParam{
				OfGoogle: &telnyx.TranscriptionEngineGoogleConfigParam{
					EnableSpeakerDiarization: telnyx.Bool(true),
					Hints:                    []string{"string"},
					InterimResults:           telnyx.Bool(true),
					Language:                 telnyx.GoogleTranscriptionLanguageEn,
					MaxSpeakerCount:          telnyx.Int(4),
					MinSpeakerCount:          telnyx.Int(4),
					Model:                    telnyx.TranscriptionEngineGoogleConfigModelLatestLong,
					ProfanityFilter:          telnyx.Bool(true),
					SpeechContext: []telnyx.TranscriptionEngineGoogleConfigSpeechContextParam{{
						Boost:   telnyx.Float(1),
						Phrases: []string{"string"},
					}},
					TranscriptionEngine: telnyx.TranscriptionEngineGoogleConfigTranscriptionEngineGoogle,
					UseEnhanced:         telnyx.Bool(true),
				},
			},
			TranscriptionTracks: telnyx.String("both"),
		},
		WebhookURL:       telnyx.String("https://www.example.com/server-b/"),
		WebhookURLMethod: telnyx.CallDialParamsWebhookURLMethodPost,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCallGetStatus(t *testing.T) {
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
	_, err := client.Calls.GetStatus(context.TODO(), "call_control_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
