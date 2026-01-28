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

func TestPortingOrderPhoneNumberConfigurationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.PhoneNumberConfigurations.New(context.TODO(), telnyx.PortingOrderPhoneNumberConfigurationNewParams{
		PhoneNumberConfigurations: []telnyx.PortingOrderPhoneNumberConfigurationNewParamsPhoneNumberConfiguration{{
			PortingPhoneNumberID: "927f4687-318c-44bc-9f2f-22a5898143a4",
			UserBundleID:         "ff901545-3e27-462a-ba9d-2b34654cab82",
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

func TestPortingOrderPhoneNumberConfigurationListWithOptionalParams(t *testing.T) {
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
	_, err := client.PortingOrders.PhoneNumberConfigurations.List(context.TODO(), telnyx.PortingOrderPhoneNumberConfigurationListParams{
		Filter: telnyx.PortingOrderPhoneNumberConfigurationListParamsFilter{
			PortingOrder: telnyx.PortingOrderPhoneNumberConfigurationListParamsFilterPortingOrder{
				Status: []string{"in-process"},
			},
			PortingPhoneNumber: []string{"5d6f7ede-1961-4717-bfb5-db392c5efc2d"},
			UserBundleID:       []string{"5d6f7ede-1961-4717-bfb5-db392c5efc2d"},
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort: telnyx.PortingOrderPhoneNumberConfigurationListParamsSort{
			Value: "created_at",
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
