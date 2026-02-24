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

func TestPhoneNumberVoiceGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.PhoneNumbers.Voice.Get(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberVoiceUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.PhoneNumbers.Voice.Update(
		context.TODO(),
		"1293384261075731499",
		telnyx.PhoneNumberVoiceUpdateParams{
			UpdateVoiceSettings: telnyx.UpdateVoiceSettingsParam{
				CallForwarding: telnyx.CallForwardingParam{
					CallForwardingEnabled: telnyx.Bool(true),
					ForwardingType:        telnyx.CallForwardingForwardingTypeAlways,
					ForwardsTo:            telnyx.String("+13035559123"),
				},
				CallRecording: telnyx.CallRecordingParam{
					InboundCallRecordingChannels: telnyx.CallRecordingInboundCallRecordingChannelsSingle,
					InboundCallRecordingEnabled:  telnyx.Bool(true),
					InboundCallRecordingFormat:   telnyx.CallRecordingInboundCallRecordingFormatWav,
				},
				CallerIDNameEnabled: telnyx.Bool(true),
				CnamListing: telnyx.CnamListingParam{
					CnamListingDetails: telnyx.String("example"),
					CnamListingEnabled: telnyx.Bool(true),
				},
				InboundCallScreening: telnyx.UpdateVoiceSettingsInboundCallScreeningDisabled,
				MediaFeatures: telnyx.MediaFeaturesParam{
					AcceptAnyRtpPacketsEnabled: telnyx.Bool(true),
					RtpAutoAdjustEnabled:       telnyx.Bool(true),
					T38FaxGatewayEnabled:       telnyx.Bool(true),
				},
				TechPrefixEnabled:  telnyx.Bool(true),
				TranslatedNumber:   telnyx.String("+13035559999"),
				UsagePaymentMethod: telnyx.UpdateVoiceSettingsUsagePaymentMethodPayPerMinute,
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

func TestPhoneNumberVoiceListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.PhoneNumbers.Voice.List(context.TODO(), telnyx.PhoneNumberVoiceListParams{
		Filter: telnyx.PhoneNumberVoiceListParamsFilter{
			ConnectionName: telnyx.PhoneNumberVoiceListParamsFilterConnectionName{
				Contains: telnyx.String("test"),
			},
			CustomerReference:       telnyx.String("customer_reference"),
			PhoneNumber:             telnyx.String("phone_number"),
			VoiceUsagePaymentMethod: "channel",
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort:       telnyx.PhoneNumberVoiceListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
