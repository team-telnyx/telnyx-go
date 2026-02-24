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

func TestUserAddressNewWithOptionalParams(t *testing.T) {
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
	_, err := client.UserAddresses.New(context.TODO(), telnyx.UserAddressNewParams{
		BusinessName:            "Toy-O'Kon",
		CountryCode:             "US",
		FirstName:               "Alfred",
		LastName:                "Foster",
		Locality:                "Austin",
		StreetAddress:           "600 Congress Avenue",
		AdministrativeArea:      telnyx.String("TX"),
		Borough:                 telnyx.String("Guadalajara"),
		CustomerReference:       telnyx.String("MY REF 001"),
		ExtendedAddress:         telnyx.String("14th Floor"),
		Neighborhood:            telnyx.String("Ciudad de los deportes"),
		PhoneNumber:             telnyx.String("+12125559000"),
		PostalCode:              telnyx.String("78701"),
		SkipAddressVerification: telnyx.Bool(true),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserAddressGet(t *testing.T) {
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
	_, err := client.UserAddresses.Get(context.TODO(), "id")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserAddressListWithOptionalParams(t *testing.T) {
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
	_, err := client.UserAddresses.List(context.TODO(), telnyx.UserAddressListParams{
		Filter: telnyx.UserAddressListParamsFilter{
			CustomerReference: telnyx.UserAddressListParamsFilterCustomerReference{
				Contains: telnyx.String("contains"),
				Eq:       telnyx.String("eq"),
			},
			StreetAddress: telnyx.UserAddressListParamsFilterStreetAddress{
				Contains: telnyx.String("contains"),
			},
		},
		PageNumber: telnyx.Int(0),
		PageSize:   telnyx.Int(0),
		Sort:       telnyx.UserAddressListParamsSortStreetAddress,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
