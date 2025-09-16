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

func TestWebhookDeliveryGet(t *testing.T) {
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
	_, err := client.WebhookDeliveries.Get(context.TODO(), "C9C0797E-901D-4349-A33C-C2C8F31A92C2")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebhookDeliveryListWithOptionalParams(t *testing.T) {
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
	_, err := client.WebhookDeliveries.List(context.TODO(), telnyx.WebhookDeliveryListParams{
		Filter: telnyx.WebhookDeliveryListParamsFilter{
			Attempts: telnyx.WebhookDeliveryListParamsFilterAttempts{
				Contains: telnyx.String("https://fallback.example.com/webhooks"),
			},
			EventType: telnyx.String("call_initiated,call.initiated"),
			FinishedAt: telnyx.WebhookDeliveryListParamsFilterFinishedAt{
				Gte: telnyx.String("2019-03-29T11:10:00Z"),
				Lte: telnyx.String("2019-03-29T11:10:00Z"),
			},
			StartedAt: telnyx.WebhookDeliveryListParamsFilterStartedAt{
				Gte: telnyx.String("2019-03-29T11:10:00Z"),
				Lte: telnyx.String("2019-03-29T11:10:00Z"),
			},
			Status: telnyx.WebhookDeliveryListParamsFilterStatus{
				Eq: "delivered",
			},
			Webhook: telnyx.WebhookDeliveryListParamsFilterWebhook{
				Contains: telnyx.String("call.initiated"),
			},
		},
		Page: telnyx.WebhookDeliveryListParamsPage{
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
