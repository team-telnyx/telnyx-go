// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/internal/testutil"
	"github.com/team-telnyx/telnyx-go/option"
)

func TestRoomParticipantGet(t *testing.T) {
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
	_, err := client.RoomParticipants.Get(context.TODO(), "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomParticipantListWithOptionalParams(t *testing.T) {
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
	_, err := client.RoomParticipants.List(context.TODO(), telnyx.RoomParticipantListParams{
		Filter: telnyx.RoomParticipantListParamsFilter{
			Context: telnyx.String("Alice"),
			DateJoinedAt: telnyx.RoomParticipantListParamsFilterDateJoinedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DateLeftAt: telnyx.RoomParticipantListParamsFilterDateLeftAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DateUpdatedAt: telnyx.RoomParticipantListParamsFilterDateUpdatedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			SessionID: telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
		},
		Page: telnyx.RoomParticipantListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
