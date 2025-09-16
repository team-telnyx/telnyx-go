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

func TestNumberOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.NumberOrders.New(context.TODO(), telnyx.NumberOrderNewParams{
		BillingGroupID:     telnyx.String("abc85f64-5717-4562-b3fc-2c9600"),
		ConnectionID:       telnyx.String("346789098765567"),
		CustomerReference:  telnyx.String("MY REF 001"),
		MessagingProfileID: telnyx.String("abc85f64-5717-4562-b3fc-2c9600"),
		PhoneNumbers: []telnyx.NumberOrderNewParamsPhoneNumber{{
			PhoneNumber:        "+19705555098",
			BundleID:           telnyx.String("bc8e4d67-33a0-4cbb-af74-7b58f05bd494"),
			RequirementGroupID: telnyx.String("dbbd94ee-9079-488f-80ba-f566b247fd7"),
		}, {
			PhoneNumber:        "+492111609539",
			BundleID:           telnyx.String("bc8e4d67-33a0-4cbb-af74-7b58f05bd494"),
			RequirementGroupID: telnyx.String("dbbd94ee-9079-488f-80ba-f566b247fd79"),
		}},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNumberOrderGet(t *testing.T) {
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
	_, err := client.NumberOrders.Get(context.TODO(), "number_order_id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNumberOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.NumberOrders.Update(
		context.TODO(),
		"number_order_id",
		telnyx.NumberOrderUpdateParams{
			CustomerReference: telnyx.String("MY REF 001"),
			RegulatoryRequirements: []telnyx.UpdateRegulatoryRequirementParam{{
				FieldValue:    telnyx.String("45f45a04-b4be-4592-95b1-9306b9db2b21"),
				RequirementID: telnyx.String("8ffb3622-7c6b-4ccc-b65f-7a3dc0099576"),
			}},
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

func TestNumberOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.NumberOrders.List(context.TODO(), telnyx.NumberOrderListParams{
		Filter: telnyx.NumberOrderListParamsFilter{
			CreatedAt: telnyx.NumberOrderListParamsFilterCreatedAt{
				Gt: telnyx.String("gt"),
				Lt: telnyx.String("lt"),
			},
			CustomerReference: telnyx.String("customer_reference"),
			PhoneNumbersCount: telnyx.String("phone_numbers_count"),
			RequirementsMet:   telnyx.Bool(true),
			Status:            telnyx.String("status"),
		},
		Page: telnyx.NumberOrderListParamsPage{
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
