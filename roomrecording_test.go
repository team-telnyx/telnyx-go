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

func TestRoomRecordingGet(t *testing.T) {
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
	_, err := client.RoomRecordings.Get(context.TODO(), "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomRecordingListWithOptionalParams(t *testing.T) {
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
	_, err := client.RoomRecordings.List(context.TODO(), telnyx.RoomRecordingListParams{
		Filter: telnyx.RoomRecordingListParamsFilter{
			DateEndedAt: telnyx.RoomRecordingListParamsFilterDateEndedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DateStartedAt: telnyx.RoomRecordingListParamsFilterDateStartedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DurationSecs:  telnyx.Int(20),
			ParticipantID: telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			RoomID:        telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			SessionID:     telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			Status:        telnyx.String("completed"),
			Type:          telnyx.String("audio"),
		},
		Page: telnyx.RoomRecordingListParamsPage{
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

func TestRoomRecordingDelete(t *testing.T) {
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
	err := client.RoomRecordings.Delete(context.TODO(), "0ccc7b54-4df3-4bca-a65a-3da1ecc777f0")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomRecordingDeleteBulkWithOptionalParams(t *testing.T) {
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
	_, err := client.RoomRecordings.DeleteBulk(context.TODO(), telnyx.RoomRecordingDeleteBulkParams{
		Filter: telnyx.RoomRecordingDeleteBulkParamsFilter{
			DateEndedAt: telnyx.RoomRecordingDeleteBulkParamsFilterDateEndedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DateStartedAt: telnyx.RoomRecordingDeleteBulkParamsFilterDateStartedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			DurationSecs:  telnyx.Int(20),
			ParticipantID: telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			RoomID:        telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			SessionID:     telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777f0"),
			Status:        telnyx.String("completed"),
			Type:          telnyx.String("audio"),
		},
		Page: telnyx.RoomRecordingDeleteBulkParamsPage{
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
