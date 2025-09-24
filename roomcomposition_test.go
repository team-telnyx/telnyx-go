// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestRoomCompositionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.RoomCompositions.New(context.TODO(), telnyx.RoomCompositionNewParams{
		Format:     telnyx.String("mp4"),
		Resolution: telnyx.String("800x600"),
		SessionID:  telnyx.String("0ccc7b54-4df3-4bca-a65a-3da1ecc777b0"),
		VideoLayout: map[string]telnyx.VideoRegionParam{
			"foo": {
				Height:       telnyx.Int(360),
				MaxColumns:   telnyx.Int(3),
				MaxRows:      telnyx.Int(3),
				VideoSources: []string{"7b61621f-62e0-4aad-ab11-9fd19e272e73"},
				Width:        telnyx.Int(480),
				XPos:         telnyx.Int(100),
				YPos:         telnyx.Int(100),
				ZPos:         telnyx.Int(1),
			},
		},
		WebhookEventFailoverURL: telnyx.String("https://failover.example.com"),
		WebhookEventURL:         telnyx.String("https://example.com"),
		WebhookTimeoutSecs:      telnyx.Int(25),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomCompositionGet(t *testing.T) {
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
	_, err := client.RoomCompositions.Get(context.TODO(), "5219b3af-87c6-4c08-9b58-5a533d893e21")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoomCompositionListWithOptionalParams(t *testing.T) {
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
	_, err := client.RoomCompositions.List(context.TODO(), telnyx.RoomCompositionListParams{
		Filter: telnyx.RoomCompositionListParamsFilter{
			DateCreatedAt: telnyx.RoomCompositionListParamsFilterDateCreatedAt{
				Eq:  telnyx.Time(time.Now()),
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			SessionID: telnyx.String("92e7d459-bcc5-4386-9f5f-6dd14a82588d"),
			Status:    "completed",
		},
		Page: telnyx.RoomCompositionListParamsPage{
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

func TestRoomCompositionDelete(t *testing.T) {
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
	err := client.RoomCompositions.Delete(context.TODO(), "5219b3af-87c6-4c08-9b58-5a533d893e21")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
