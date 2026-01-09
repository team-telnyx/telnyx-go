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

func TestOutboundVoiceProfileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.OutboundVoiceProfiles.New(context.TODO(), telnyx.OutboundVoiceProfileNewParams{
		Name:           "office",
		BillingGroupID: telnyx.String("6a09cdc3-8948-47f0-aa62-74ac943d6c58"),
		CallRecording: telnyx.OutboundCallRecordingParam{
			CallRecordingCallerPhoneNumbers: []string{"+19705555098"},
			CallRecordingChannels:           telnyx.OutboundCallRecordingCallRecordingChannelsDual,
			CallRecordingFormat:             telnyx.OutboundCallRecordingCallRecordingFormatMP3,
			CallRecordingType:               telnyx.OutboundCallRecordingCallRecordingTypeByCallerPhoneNumber,
		},
		CallingWindow: telnyx.OutboundVoiceProfileNewParamsCallingWindow{
			CallsPerCld: telnyx.Int(5),
			EndTime:     telnyx.String("20:00:00.00Z"),
			StartTime:   telnyx.String("08:00:00.00Z"),
		},
		ConcurrentCallLimit:     telnyx.Int(10),
		DailySpendLimit:         telnyx.String("100.00"),
		DailySpendLimitEnabled:  telnyx.Bool(true),
		Enabled:                 telnyx.Bool(true),
		MaxDestinationRate:      telnyx.Float(10),
		ServicePlan:             telnyx.ServicePlanGlobal,
		Tags:                    []string{"office-profile"},
		TrafficType:             telnyx.TrafficTypeConversational,
		UsagePaymentMethod:      telnyx.UsagePaymentMethodRateDeck,
		WhitelistedDestinations: []string{"US", "BR", "AU"},
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOutboundVoiceProfileGet(t *testing.T) {
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
	_, err := client.OutboundVoiceProfiles.Get(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOutboundVoiceProfileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.OutboundVoiceProfiles.Update(
		context.TODO(),
		"1293384261075731499",
		telnyx.OutboundVoiceProfileUpdateParams{
			Name:           "office",
			BillingGroupID: telnyx.String("6a09cdc3-8948-47f0-aa62-74ac943d6c58"),
			CallRecording: telnyx.OutboundCallRecordingParam{
				CallRecordingCallerPhoneNumbers: []string{"+19705555098"},
				CallRecordingChannels:           telnyx.OutboundCallRecordingCallRecordingChannelsDual,
				CallRecordingFormat:             telnyx.OutboundCallRecordingCallRecordingFormatMP3,
				CallRecordingType:               telnyx.OutboundCallRecordingCallRecordingTypeByCallerPhoneNumber,
			},
			CallingWindow: telnyx.OutboundVoiceProfileUpdateParamsCallingWindow{
				CallsPerCld: telnyx.Int(5),
				EndTime:     telnyx.String("20:00:00.00Z"),
				StartTime:   telnyx.String("08:00:00.00Z"),
			},
			ConcurrentCallLimit:     telnyx.Int(10),
			DailySpendLimit:         telnyx.String("100.00"),
			DailySpendLimitEnabled:  telnyx.Bool(true),
			Enabled:                 telnyx.Bool(true),
			MaxDestinationRate:      telnyx.Float(10),
			ServicePlan:             telnyx.ServicePlanGlobal,
			Tags:                    []string{"office-profile"},
			TrafficType:             telnyx.TrafficTypeConversational,
			UsagePaymentMethod:      telnyx.UsagePaymentMethodRateDeck,
			WhitelistedDestinations: []string{"US", "BR", "AU"},
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

func TestOutboundVoiceProfileListWithOptionalParams(t *testing.T) {
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
	_, err := client.OutboundVoiceProfiles.List(context.TODO(), telnyx.OutboundVoiceProfileListParams{
		Filter: telnyx.OutboundVoiceProfileListParamsFilter{
			Name: telnyx.OutboundVoiceProfileListParamsFilterName{
				Contains: telnyx.String("office-profile"),
			},
		},
		Page: telnyx.OutboundVoiceProfileListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.OutboundVoiceProfileListParamsSortName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOutboundVoiceProfileDelete(t *testing.T) {
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
	_, err := client.OutboundVoiceProfiles.Delete(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
