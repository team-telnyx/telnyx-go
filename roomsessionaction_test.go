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

func TestRoomSessionActionEnd(t *testing.T) {
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
	_, err := client.Rooms.Sessions.Actions.End(context.TODO(), "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomSessionActionKickWithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.Actions.Kick(
		context.TODO(),
		"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		telnyx.RoomSessionActionKickParams{
			ActionsParticipantsRequest: telnyx.ActionsParticipantsRequestParam{
				Exclude: []string{"7b61621f-62e0-4aad-ab11-9fd19e272e73"},
				Participants: telnyx.ActionsParticipantsRequestParticipantsUnionParam{
					OfActionsParticipantsRequestParticipantsString: telnyx.String("all"),
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

func TestRoomSessionActionMuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.Actions.Mute(
		context.TODO(),
		"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		telnyx.RoomSessionActionMuteParams{
			ActionsParticipantsRequest: telnyx.ActionsParticipantsRequestParam{
				Exclude: []string{"7b61621f-62e0-4aad-ab11-9fd19e272e73"},
				Participants: telnyx.ActionsParticipantsRequestParticipantsUnionParam{
					OfActionsParticipantsRequestParticipantsString: telnyx.String("all"),
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

func TestRoomSessionActionUnmuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.Actions.Unmute(
		context.TODO(),
		"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		telnyx.RoomSessionActionUnmuteParams{
			ActionsParticipantsRequest: telnyx.ActionsParticipantsRequestParam{
				Exclude: []string{"7b61621f-62e0-4aad-ab11-9fd19e272e73"},
				Participants: telnyx.ActionsParticipantsRequestParticipantsUnionParam{
					OfActionsParticipantsRequestParticipantsString: telnyx.String("all"),
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
