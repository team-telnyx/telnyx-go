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

func TestPhoneNumberGet(t *testing.T) {
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
	_, err := client.PhoneNumbers.Get(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.Update(
		context.TODO(),
		"1293384261075731499",
		telnyx.PhoneNumberUpdateParams{
			BillingGroupID:    telnyx.String("dc8e4d67-33a0-4cbb-af74-7b58f05bd494"),
			ConnectionID:      telnyx.String("dc8e4d67-33a0-4cbb-af74-7b58f05bd494"),
			CustomerReference: telnyx.String("customer-reference"),
			ExternalPin:       telnyx.String("1234"),
			HDVoiceEnabled:    telnyx.Bool(true),
			Tags:              []string{"tag"},
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

func TestPhoneNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.List(context.TODO(), telnyx.PhoneNumberListParams{
		Filter: telnyx.PhoneNumberListParamsFilter{
			BillingGroupID: telnyx.String("62e4bf2e-c278-4282-b524-488d9c9c43b2"),
			ConnectionID:   telnyx.String("1521916448077776306"),
			CountryISOAlpha2: telnyx.PhoneNumberListParamsFilterCountryISOAlpha2Union{
				OfString: telnyx.String("US"),
			},
			CustomerReference:  telnyx.String("customer_reference"),
			EmergencyAddressID: telnyx.String("9102160989215728032"),
			NumberType: telnyx.PhoneNumberListParamsFilterNumberType{
				Eq: "local",
			},
			PhoneNumber: telnyx.String("phone_number"),
			Source:      "ported",
			Status:      "active",
			Tag:         telnyx.String("tag"),
			VoiceConnectionName: telnyx.PhoneNumberListParamsFilterVoiceConnectionName{
				Contains:   telnyx.String("test"),
				EndsWith:   telnyx.String("test"),
				Eq:         telnyx.String("test"),
				StartsWith: telnyx.String("test"),
			},
			VoiceUsagePaymentMethod: "channel",
			WithoutTags:             "true",
		},
		HandleMessagingProfileError: telnyx.PhoneNumberListParamsHandleMessagingProfileErrorFalse,
		Page: telnyx.PhoneNumberListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.PhoneNumberListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberDelete(t *testing.T) {
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
	_, err := client.PhoneNumbers.Delete(context.TODO(), "1293384261075731499")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberSlimListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.SlimList(context.TODO(), telnyx.PhoneNumberSlimListParams{
		Filter: telnyx.PhoneNumberSlimListParamsFilter{
			BillingGroupID: telnyx.String("62e4bf2e-c278-4282-b524-488d9c9c43b2"),
			ConnectionID:   telnyx.String("1521916448077776306"),
			CountryISOAlpha2: telnyx.PhoneNumberSlimListParamsFilterCountryISOAlpha2Union{
				OfString: telnyx.String("US"),
			},
			CustomerReference:  telnyx.String("customer_reference"),
			EmergencyAddressID: telnyx.String("9102160989215728032"),
			NumberType: telnyx.PhoneNumberSlimListParamsFilterNumberType{
				Eq: "local",
			},
			PhoneNumber: telnyx.String("phone_number"),
			Source:      "ported",
			Status:      "active",
			Tag:         telnyx.String("tag"),
			VoiceConnectionName: telnyx.PhoneNumberSlimListParamsFilterVoiceConnectionName{
				Contains:   telnyx.String("test"),
				EndsWith:   telnyx.String("test"),
				Eq:         telnyx.String("test"),
				StartsWith: telnyx.String("test"),
			},
			VoiceUsagePaymentMethod: "channel",
		},
		IncludeConnection: telnyx.Bool(true),
		IncludeTags:       telnyx.Bool(true),
		Page: telnyx.PhoneNumberSlimListParamsPage{
			Number: telnyx.Int(1),
			Size:   telnyx.Int(1),
		},
		Sort: telnyx.PhoneNumberSlimListParamsSortConnectionName,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
