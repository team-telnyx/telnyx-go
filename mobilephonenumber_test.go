// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/team-telnyx/telnyx-go/v3"
	"github.com/team-telnyx/telnyx-go/v3/internal/testutil"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

func TestMobilePhoneNumberGet(t *testing.T) {
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
	_, err := client.MobilePhoneNumbers.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMobilePhoneNumberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MobilePhoneNumbers.Update(
		context.TODO(),
		"id",
		telnyx.MobilePhoneNumberUpdateParams{
			CallForwarding: telnyx.MobilePhoneNumberUpdateParamsCallForwarding{
				CallForwardingEnabled: telnyx.Bool(true),
				ForwardingType:        "always",
				ForwardsTo:            telnyx.String("forwards_to"),
			},
			CallRecording: telnyx.MobilePhoneNumberUpdateParamsCallRecording{
				InboundCallRecordingChannels: "single",
				InboundCallRecordingEnabled:  telnyx.Bool(true),
				InboundCallRecordingFormat:   "wav",
			},
			CallerIDNameEnabled: telnyx.Bool(true),
			CnamListing: telnyx.MobilePhoneNumberUpdateParamsCnamListing{
				CnamListingDetails: telnyx.String("cnam_listing_details"),
				CnamListingEnabled: telnyx.Bool(true),
			},
			ConnectionID:      telnyx.String("connection_id"),
			CustomerReference: telnyx.String("customer_reference"),
			Inbound: telnyx.MobilePhoneNumberUpdateParamsInbound{
				InterceptionAppID: telnyx.String("interception_app_id"),
			},
			InboundCallScreening: telnyx.MobilePhoneNumberUpdateParamsInboundCallScreeningDisabled,
			NoiseSuppression:     telnyx.Bool(true),
			Outbound: telnyx.MobilePhoneNumberUpdateParamsOutbound{
				InterceptionAppID: telnyx.String("interception_app_id"),
			},
			Tags: []string{"string"},
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

func TestMobilePhoneNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.MobilePhoneNumbers.List(context.TODO(), telnyx.MobilePhoneNumberListParams{
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
