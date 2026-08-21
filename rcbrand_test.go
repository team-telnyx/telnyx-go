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

func TestRcBrandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Rcs.Brands.New(context.TODO(), telnyx.RcBrandNewParams{
		Addresses: map[string]telnyx.BrandAddressParam{
			"primary": {
				AdministrativeArea: "IL",
				City:               "Chicago",
				CountryCode:        "US",
				Line1:              "1 Main Street",
				PostalCode:         "60601",
				Line2:              telnyx.String("x"),
			},
		},
		Contacts: telnyx.RcBrandNewParamsContacts{
			Brand: telnyx.RcBrandNewParamsContactsBrand{
				BrandContactParam: telnyx.BrandContactParam{
					ContactType: telnyx.BrandContactContactTypeBrand,
					Email:       "jane@example.com",
					FirstName:   "Jane",
					LastName:    "Doe",
					PhoneNumber: "+13125550100",
					Title:       telnyx.String("Messaging Operations Manager"),
				},
				ContactType: string(telnyx.BrandContactContactTypeBrand),
			},
		},
		DisplayName: "Acme",
		Identifiers: telnyx.RcBrandNewParamsIdentifiers{
			Ein: telnyx.EinBrandIdentifierParam{
				Value: "12-3456789",
			},
			StockSymbol: telnyx.StockSymbolBrandIdentifierParam{
				Value: "J!Q0Ok0bzJb7:pro",
			},
		},
		LegalEntityType:  telnyx.BrandLegalEntityTypeLimitedLiabilityCompany,
		LegalName:        "Acme LLC",
		OrganizationType: telnyx.BrandOrganizationTypePrivateProfit,
		WebsiteURL:       "https://www.example.com",
		ProfileID:        telnyx.String("40000000-0000-0000-0000-000000000001"),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcBrandGet(t *testing.T) {
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
	_, err := client.Rcs.Brands.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcBrandUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Rcs.Brands.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		telnyx.RcBrandUpdateParams{
			Addresses: map[string]telnyx.BrandAddressParam{
				"foo": {
					AdministrativeArea: "x",
					City:               "x",
					CountryCode:        "SE",
					Line1:              "x",
					PostalCode:         "x",
					Line2:              telnyx.String("x"),
				},
			},
			Contacts: telnyx.RcBrandUpdateParamsContacts{
				Brand: telnyx.RcBrandUpdateParamsContactsBrand{
					BrandContactParam: telnyx.BrandContactParam{
						ContactType: telnyx.BrandContactContactTypeBrand,
						Email:       "dev@stainless.com",
						FirstName:   "x",
						LastName:    "x",
						PhoneNumber: "+49605132",
						Title:       telnyx.String("x"),
					},
					ContactType: string(telnyx.BrandContactContactTypeBrand),
				},
			},
			DisplayName: telnyx.String("Acme Communications"),
			Identifiers: telnyx.RcBrandUpdateParamsIdentifiers{
				Ein: telnyx.EinBrandIdentifierParam{
					Value: "29-1051329",
				},
				StockSymbol: telnyx.StockSymbolBrandIdentifierParam{
					Value: "J!Q0Ok0bzJb7:pro",
				},
			},
			LegalEntityType:  telnyx.BrandLegalEntityTypeLimitedLiabilityCompany,
			LegalName:        telnyx.String("x"),
			OrganizationType: telnyx.BrandOrganizationTypePrivateProfit,
			ProfileID:        telnyx.String("profile_id"),
			WebsiteURL:       telnyx.String("https://example.com"),
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

func TestRcBrandList(t *testing.T) {
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
	_, err := client.Rcs.Brands.List(context.TODO())
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRcBrandSubmit(t *testing.T) {
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
	_, err := client.Rcs.Brands.Submit(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
