// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func TestRoomSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.Get(
		context.TODO(),
		"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		telnyx.RoomSessionGetParams{
			IncludeParticipants: telnyx.Bool(true),
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

func TestRoomSessionList0WithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.List0(context.TODO(), telnyx.RoomSessionList0Params{
		Filter: telnyx.RoomSessionList0ParamsFilter{
			Active: telnyx.Bool(true),
			DateCreatedAt: telnyx.RoomSessionList0ParamsFilterDateCreatedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DateEndedAt: telnyx.RoomSessionList0ParamsFilterDateEndedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DateUpdatedAt: telnyx.RoomSessionList0ParamsFilterDateUpdatedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			RoomID: telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
		},
		IncludeParticipants: telnyx.Bool(true),
		PageNumber:          telnyx.Int(0),
		PageSize:            telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomSessionList1WithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.List1(
		context.TODO(),
		"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		telnyx.RoomSessionList1Params{
			Filter: telnyx.RoomSessionList1ParamsFilter{
				Active: telnyx.Bool(true),
				DateCreatedAt: telnyx.RoomSessionList1ParamsFilterDateCreatedAt{
					Eq:  telnyx.Time(time.Now()),
					Gte: telnyx.Time(time.Now()),
					Lte: telnyx.Time(time.Now()),
				},
				DateEndedAt: telnyx.RoomSessionList1ParamsFilterDateEndedAt{
					Eq:  telnyx.Time(time.Now()),
					Gte: telnyx.Time(time.Now()),
					Lte: telnyx.Time(time.Now()),
				},
				DateUpdatedAt: telnyx.RoomSessionList1ParamsFilterDateUpdatedAt{
					Eq:  telnyx.Time(time.Now()),
					Gte: telnyx.Time(time.Now()),
					Lte: telnyx.Time(time.Now()),
				},
			},
			IncludeParticipants: telnyx.Bool(true),
			PageNumber:          telnyx.Int(0),
			PageSize:            telnyx.Int(0),
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

func TestRoomSessionGetParticipantsWithOptionalParams(t *testing.T) {
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
	_, err := client.Rooms.Sessions.GetParticipants(
		context.TODO(),
		"0ccc7b54-4df3-4bca-a65a-3da1ecc777f0",
		telnyx.RoomSessionGetParticipantsParams{
			Filter: telnyx.RoomSessionGetParticipantsParamsFilter{
				Context: telnyx.String("Alice"),
				DateJoinedAt: telnyx.RoomSessionGetParticipantsParamsFilterDateJoinedAt{
					Eq:  telnyx.Time(time.Now()),
					Gte: telnyx.Time(time.Now()),
					Lte: telnyx.Time(time.Now()),
				},
				DateLeftAt: telnyx.RoomSessionGetParticipantsParamsFilterDateLeftAt{
					Eq:  telnyx.Time(time.Now()),
					Gte: telnyx.Time(time.Now()),
					Lte: telnyx.Time(time.Now()),
				},
				DateUpdatedAt: telnyx.RoomSessionGetParticipantsParamsFilterDateUpdatedAt{
					Eq:  telnyx.Time(time.Now()),
					Gte: telnyx.Time(time.Now()),
					Lte: telnyx.Time(time.Now()),
				},
			},
			PageNumber: telnyx.Int(0),
			PageSize:   telnyx.Int(0),
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
