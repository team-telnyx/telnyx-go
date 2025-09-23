// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go"
	"github.com/team-telnyx/telnyx-go/internal/testutil"
	"github.com/team-telnyx/telnyx-go/option"
)

func TestPortingOrderPhoneNumberBlockNew(t *testing.T) {
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
	_, err := client.PortingOrders.PhoneNumberBlocks.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderPhoneNumberBlockNewParams{
			ActivationRanges: []telnyx.PortingOrderPhoneNumberBlockNewParamsActivationRange{{
				EndAt:   "+4930244999910",
				StartAt: "+4930244999901",
			}},
			PhoneNumberRange: telnyx.PortingOrderPhoneNumberBlockNewParamsPhoneNumberRange{
				EndAt:   "+4930244999910",
				StartAt: "+4930244999901",
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

func TestPortingOrderPhoneNumberBlockListWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.PhoneNumberBlocks.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderPhoneNumberBlockListParams{
			Filter: telnyx.PortingOrderPhoneNumberBlockListParamsFilter{
				ActivationStatus:  "Active",
				PhoneNumber:       []string{"+12003151212"},
				PortabilityStatus: "confirmed",
				PortingOrderID:    []string{"f3575e15-32ce-400e-a4c0-dd78800c20b0"},
				Status: telnyx.PortingOrderPhoneNumberBlockListParamsFilterStatusUnion{
					OfPortingOrderPhoneNumberBlockListsFilterStatusString: telnyx.String("in-process"),
				},
				SupportKey: telnyx.PortingOrderPhoneNumberBlockListParamsFilterSupportKeyUnion{
					OfString: telnyx.String("sr_a12345"),
				},
			},
			Page: telnyx.PortingOrderPhoneNumberBlockListParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(1),
			},
			Sort: telnyx.PortingOrderPhoneNumberBlockListParamsSort{
				Value: "-created_at",
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

func TestPortingOrderPhoneNumberBlockDelete(t *testing.T) {
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
	_, err := client.PortingOrders.PhoneNumberBlocks.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.PortingOrderPhoneNumberBlockDeleteParams{
			PortingOrderID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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
