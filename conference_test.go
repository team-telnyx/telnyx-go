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

func TestConferenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.New(context.TODO(), telnyx.ConferenceNewParams{
		CallControlID:           "v3:MdI91X4lWFEs7IgbBEOT9M4AigoY08M0WWZFISt1Yw2axZ_IiE4pqg",
		Name:                    "Business",
		BeepEnabled:             telnyx.ConferenceNewParamsBeepEnabledAlways,
		ClientState:             telnyx.String("aGF2ZSBhIG5pY2UgZGF5ID1d"),
		ComfortNoise:            telnyx.Bool(false),
		CommandID:               telnyx.String("891510ac-f3e4-11e8-af5b-de00688a4901"),
		DurationMinutes:         telnyx.Int(5),
		HoldAudioURL:            telnyx.String("http://www.example.com/audio.wav"),
		HoldMediaName:           telnyx.String("my_media_uploaded_to_media_storage_api"),
		MaxParticipants:         telnyx.Int(250),
		Region:                  telnyx.ConferenceNewParamsRegionUs,
		StartConferenceOnCreate: telnyx.Bool(false),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConferenceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.Get(
		context.TODO(),
		"id",
		telnyx.ConferenceGetParams{
			Region: telnyx.ConferenceGetParamsRegionAustralia,
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

func TestConferenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.List(context.TODO(), telnyx.ConferenceListParams{
		Filter: telnyx.ConferenceListParamsFilter{
			ApplicationName: telnyx.ConferenceListParamsFilterApplicationName{
				Contains: telnyx.String("contains"),
			},
			ApplicationSessionID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ConnectionID:         telnyx.String("connection_id"),
			Failed:               telnyx.Bool(false),
			From:                 telnyx.String("+12025550142"),
			LegID:                telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:                 telnyx.String("name"),
			OccurredAt: telnyx.ConferenceListParamsFilterOccurredAt{
				Eq:  telnyx.String("2019-03-29T11:10:00Z"),
				Gt:  telnyx.String("2019-03-29T11:10:00Z"),
				Gte: telnyx.String("2019-03-29T11:10:00Z"),
				Lt:  telnyx.String("2019-03-29T11:10:00Z"),
				Lte: telnyx.String("2019-03-29T11:10:00Z"),
			},
			OutboundOutboundVoiceProfileID: telnyx.String("1293384261075731499"),
			Product:                        "texml",
			Status:                         "init",
			To:                             telnyx.String("+12025550142"),
			Type:                           "webhook",
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Region:     telnyx.ConferenceListParamsRegionAustralia,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConferenceListParticipantsWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.ListParticipants(
		context.TODO(),
		"conference_id",
		telnyx.ConferenceListParticipantsParams{
			Filter: telnyx.ConferenceListParticipantsParamsFilter{
				Muted:      telnyx.Bool(true),
				OnHold:     telnyx.Bool(true),
				Whispering: telnyx.Bool(true),
			},
			PageNumber: telnyx.Int(0),
			PageSize:   telnyx.Int(0),
			Region:     telnyx.ConferenceListParticipantsParamsRegionAustralia,
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

func TestConferenceGetParticipant(t *testing.T) {
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
	_, err := client.Conferences.GetParticipant(
		context.TODO(),
		"participant_id",
		telnyx.ConferenceGetParticipantParams{
			ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestConferenceUpdateParticipantWithOptionalParams(t *testing.T) {
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
	_, err := client.Conferences.UpdateParticipant(
		context.TODO(),
		"participant_id",
		telnyx.ConferenceUpdateParticipantParams{
			ID:                      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			BeepEnabled:             telnyx.ConferenceUpdateParticipantParamsBeepEnabledNever,
			EndConferenceOnExit:     telnyx.Bool(true),
			SoftEndConferenceOnExit: telnyx.Bool(false),
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
