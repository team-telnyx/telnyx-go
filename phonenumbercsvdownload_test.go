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

func TestPhoneNumberCsvDownloadNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.CsvDownloads.New(context.TODO(), telnyx.PhoneNumberCsvDownloadNewParams{
		CsvFormat: telnyx.PhoneNumberCsvDownloadNewParamsCsvFormatV2,
		Filter: telnyx.PhoneNumberCsvDownloadNewParamsFilter{
			BillingGroupID:     telnyx.String("62e4bf2e-c278-4282-b524-488d9c9c43b2"),
			ConnectionID:       telnyx.String("1521916448077776306"),
			CustomerReference:  telnyx.String("customer_reference"),
			EmergencyAddressID: telnyx.String("9102160989215728032"),
			HasBundle:          telnyx.String("has_bundle"),
			PhoneNumber:        telnyx.String("phone_number"),
			Status:             "active",
			Tag:                telnyx.String("tag"),
			VoiceConnectionName: telnyx.PhoneNumberCsvDownloadNewParamsFilterVoiceConnectionName{
				Contains:   telnyx.String("test"),
				EndsWith:   telnyx.String("test"),
				Eq:         telnyx.String("test"),
				StartsWith: telnyx.String("test"),
			},
			VoiceUsagePaymentMethod: "channel",
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

func TestPhoneNumberCsvDownloadGet(t *testing.T) {
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
	_, err := client.PhoneNumbers.CsvDownloads.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPhoneNumberCsvDownloadListWithOptionalParams(t *testing.T) {
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
	_, err := client.PhoneNumbers.CsvDownloads.List(context.TODO(), telnyx.PhoneNumberCsvDownloadListParams{
		Page: telnyx.PhoneNumberCsvDownloadListParamsPage{
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
