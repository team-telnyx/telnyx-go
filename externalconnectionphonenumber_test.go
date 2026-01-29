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

func TestExternalConnectionPhoneNumberGet(t *testing.T) {
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
	_, err := client.ExternalConnections.PhoneNumbers.Get(
		context.TODO(),
		"1234567889",
		telnyx.ExternalConnectionPhoneNumberGetParams{
			ID: "id",
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

func TestExternalConnectionPhoneNumberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalConnections.PhoneNumbers.Update(
		context.TODO(),
		"1234567889",
		telnyx.ExternalConnectionPhoneNumberUpdateParams{
			ID:         "id",
			LocationID: telnyx.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestExternalConnectionPhoneNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExternalConnections.PhoneNumbers.List(
		context.TODO(),
		"id",
		telnyx.ExternalConnectionPhoneNumberListParams{
			Filter: telnyx.ExternalConnectionPhoneNumberListParamsFilter{
				CivicAddressID: telnyx.ExternalConnectionPhoneNumberListParamsFilterCivicAddressID{
					Eq: telnyx.String("19990261512338516954"),
				},
				LocationID: telnyx.ExternalConnectionPhoneNumberListParamsFilterLocationID{
					Eq: telnyx.String("19995665508264022121"),
				},
				PhoneNumber: telnyx.ExternalConnectionPhoneNumberListParamsFilterPhoneNumber{
					Contains: telnyx.String("+1970"),
					Eq:       telnyx.String("+19705555098"),
				},
			},
			Page: telnyx.ExternalConnectionPhoneNumberListParamsPage{
				Number: telnyx.Int(1),
				Size:   telnyx.Int(1),
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
