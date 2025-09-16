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

func TestAvailablePhoneNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.AvailablePhoneNumbers.List(context.TODO(), telnyx.AvailablePhoneNumberListParams{
		Filter: telnyx.AvailablePhoneNumberListParamsFilter{
			AdministrativeArea:      telnyx.String("administrative_area"),
			BestEffort:              telnyx.Bool(true),
			CountryCode:             telnyx.String("country_code"),
			ExcludeHeldNumbers:      telnyx.Bool(true),
			Features:                []string{"sms"},
			Limit:                   telnyx.Int(0),
			Locality:                telnyx.String("locality"),
			NationalDestinationCode: telnyx.String("national_destination_code"),
			PhoneNumber: telnyx.AvailablePhoneNumberListParamsFilterPhoneNumber{
				Contains:   telnyx.String("contains"),
				EndsWith:   telnyx.String("ends_with"),
				StartsWith: telnyx.String("starts_with"),
			},
			PhoneNumberType: "local",
			Quickship:       telnyx.Bool(true),
			RateCenter:      telnyx.String("rate_center"),
			Reservable:      telnyx.Bool(true),
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
