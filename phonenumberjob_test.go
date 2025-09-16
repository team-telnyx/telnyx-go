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

func TestPhoneNumberJobGet(t *testing.T) {
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
	_, err := client.PhoneNumbers.Jobs.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberJobListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.Jobs.List(context.TODO(), telnyx.PhoneNumberJobListParams{
		Filter: telnyx.PhoneNumberJobListParamsFilter{
			Type: "update_emergency_settings",
		},
		Page: telnyx.PhoneNumberJobListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.PhoneNumberJobListParamsSortCreatedAt,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberJobDeleteBatch(t *testing.T) {
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
	_, err := client.PhoneNumbers.Jobs.DeleteBatch(context.TODO(), telnyx.PhoneNumberJobDeleteBatchParams{
		PhoneNumbers: []string{"+19705555098", "+19715555098", "32873127836"},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberJobUpdateBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.Jobs.UpdateBatch(context.TODO(), telnyx.PhoneNumberJobUpdateBatchParams{
		PhoneNumbers: []string{"1583466971586889004", "+13127367254"},
		Filter: telnyx.PhoneNumberJobUpdateBatchParamsFilter{
			BillingGroupID:     telnyx.String("62e4bf2e-c278-4282-b524-488d9c9c43b2"),
			ConnectionID:       telnyx.String("1521916448077776306"),
			CustomerReference:  telnyx.String("customer_reference"),
			EmergencyAddressID: telnyx.String("9102160989215728032"),
			HasBundle:          telnyx.String("has_bundle"),
			PhoneNumber:        telnyx.String("phone_number"),
			Status:             "active",
			Tag:                telnyx.String("tag"),
			VoiceConnectionName: telnyx.PhoneNumberJobUpdateBatchParamsFilterVoiceConnectionName{
				Contains:   telnyx.String("test"),
				EndsWith:   telnyx.String("test"),
				Eq:         telnyx.String("test"),
				StartsWith: telnyx.String("test"),
			},
			VoiceUsagePaymentMethod: "channel",
		},
		BillingGroupID:    telnyx.String("dc8e4d67-33a0-4cbb-af74-7b58f05bd494"),
		ConnectionID:      telnyx.String("dc8e4d67-33a0-4cbb-af74-7b58f05bd494"),
		CustomerReference: telnyx.String("customer-reference"),
		ExternalPin:       telnyx.String("123456"),
		HDVoiceEnabled:    telnyx.Bool(true),
		Tags:              []string{"tag"},
		Voice: telnyx.UpdateVoiceSettingsParam{
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
			TranslatedNumber:   telnyx.String("translated_number"),
			UsagePaymentMethod: telnyx.UpdateVoiceSettingsUsagePaymentMethodPayPerMinute,
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

func TestPhoneNumberJobUpdateEmergencySettingsBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.Jobs.UpdateEmergencySettingsBatch(context.TODO(), telnyx.PhoneNumberJobUpdateEmergencySettingsBatchParams{
		EmergencyEnabled:   true,
		PhoneNumbers:       []string{"+19705555098", "+19715555098", "32873127836"},
		EmergencyAddressID: telnyx.String("53829456729313"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
