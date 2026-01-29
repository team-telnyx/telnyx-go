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

func TestPortoutGet(t *testing.T) {
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
	_, err := client.Portouts.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortoutListWithOptionalParams(t *testing.T) {
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
	_, err := client.Portouts.List(context.TODO(), telnyx.PortoutListParams{
		Filter: telnyx.PortoutListParamsFilter{
			CarrierName:   telnyx.String("carrier_name"),
			CountryCode:   telnyx.String("US"),
			CountryCodeIn: []string{"CA", "US"},
			FocDate:       telnyx.Time(time.Now()),
			InsertedAt: telnyx.PortoutListParamsFilterInsertedAt{
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			PhoneNumber: telnyx.String("+13035551212"),
			Pon:         telnyx.String("pon"),
			PortedOutAt: telnyx.PortoutListParamsFilterPortedOutAt{
				Gte: telnyx.Time(time.Now()),
				Lte: telnyx.Time(time.Now()),
			},
			Spid:       telnyx.String("spid"),
			Status:     "pending",
			StatusIn:   []string{"pending"},
			SupportKey: telnyx.String("PO_abc123"),
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPortoutListRejectionCodesWithOptionalParams(t *testing.T) {
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
	_, err := client.Portouts.ListRejectionCodes(
		context.TODO(),
		"329d6658-8f93-405d-862f-648776e8afd7",
		telnyx.PortoutListRejectionCodesParams{
			Filter: telnyx.PortoutListRejectionCodesParamsFilter{
				Code: telnyx.PortoutListRejectionCodesParamsFilterCodeUnion{
					OfInt: telnyx.Int(1002),
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

func TestPortoutUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Portouts.UpdateStatus(
		context.TODO(),
		telnyx.PortoutUpdateStatusParamsStatusAuthorized,
		telnyx.PortoutUpdateStatusParams{
			ID:            "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Reason:        "I do not recognize this transaction",
			HostMessaging: telnyx.Bool(false),
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
