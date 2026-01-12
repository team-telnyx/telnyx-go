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

func TestInexplicitNumberOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.InexplicitNumberOrders.New(context.TODO(), telnyx.InexplicitNumberOrderNewParams{
		OrderingGroups: []telnyx.InexplicitNumberOrderNewParamsOrderingGroup{{
			CountRequested:          "count_requested",
			CountryISO:              "US",
			PhoneNumberType:         "phone_number_type",
			AdministrativeArea:      telnyx.String("administrative_area"),
			ExcludeHeldNumbers:      telnyx.Bool(true),
			Features:                []string{"string"},
			Locality:                telnyx.String("locality"),
			NationalDestinationCode: telnyx.String("national_destination_code"),
			PhoneNumber: telnyx.InexplicitNumberOrderNewParamsOrderingGroupPhoneNumber{
				Contains:   telnyx.String("contains"),
				EndsWith:   telnyx.String("ends_with"),
				StartsWith: telnyx.String("starts_with"),
			},
			Quickship: telnyx.Bool(true),
			Strategy:  "always",
		}},
		BillingGroupID:     telnyx.String("billing_group_id"),
		ConnectionID:       telnyx.String("connection_id"),
		CustomerReference:  telnyx.String("customer_reference"),
		MessagingProfileID: telnyx.String("messaging_profile_id"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInexplicitNumberOrderGet(t *testing.T) {
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
	_, err := client.InexplicitNumberOrders.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInexplicitNumberOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.InexplicitNumberOrders.List(context.TODO(), telnyx.InexplicitNumberOrderListParams{
		PageNumber: telnyx.Int(1),
		PageSize:   telnyx.Int(1),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
