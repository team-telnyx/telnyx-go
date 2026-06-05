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

func TestEnterpriseNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.New(context.TODO(), telnyx.EnterpriseNewParams{
		BillingAddress: telnyx.BillingAddressParam{
			AdministrativeArea: "IL",
			City:               "Chicago",
			Country:            "US",
			PostalCode:         "60601",
			StreetAddress:      "100 Main St",
			ExtendedAddress:    telnyx.String("Suite 504"),
		},
		BillingContact: telnyx.BillingContactParam{
			Email:       "billing@run065.example.com",
			FirstName:   "Alex",
			LastName:    "Bill",
			PhoneNumber: "+13125550001",
		},
		CountryCode:                 "US",
		DoingBusinessAs:             "Run 065 Debug",
		Fein:                        "12-3456789",
		Industry:                    telnyx.EnterpriseNewParamsIndustryTechnology,
		JurisdictionOfIncorporation: "Delaware",
		LegalName:                   "Run 065 Debug Co",
		NumberOfEmployees:           telnyx.EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees51_200,
		OrganizationContact: telnyx.OrganizationContactParam{
			Email:       "org@run065.example.com",
			FirstName:   "Sam",
			JobTitle:    "Compliance Lead",
			LastName:    "Org",
			PhoneNumber: "+13125550000",
		},
		OrganizationLegalType: telnyx.EnterpriseNewParamsOrganizationLegalTypeLlc,
		OrganizationPhysicalAddress: telnyx.PhysicalAddressParam{
			AdministrativeArea: "IL",
			City:               "Chicago",
			Country:            "US",
			PostalCode:         "60601",
			StreetAddress:      "100 Main St",
			ExtendedAddress:    telnyx.String("Suite 504"),
		},
		OrganizationType:             telnyx.EnterpriseNewParamsOrganizationTypeCommercial,
		Website:                      "https://run065.example.com",
		CorporateRegistrationNumber:  telnyx.String("corporate_registration_number"),
		CustomerReference:            telnyx.String("internal-id-12345"),
		DunBradstreetNumber:          telnyx.String("dun_bradstreet_number"),
		PrimaryBusinessDomainSicCode: telnyx.String("primary_business_domain_sic_code"),
		ProfessionalLicenseNumber:    telnyx.String("professional_license_number"),
		RoleType:                     telnyx.EnterpriseNewParamsRoleTypeEnterprise,
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnterpriseGet(t *testing.T) {
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
	_, err := client.Enterprises.Get(context.TODO(), "4a6192a4-573d-446d-b3ce-aff9117272a6")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnterpriseUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.Update(
		context.TODO(),
		"4a6192a4-573d-446d-b3ce-aff9117272a6",
		telnyx.EnterpriseUpdateParams{
			BillingAddress: telnyx.BillingAddressParam{
				AdministrativeArea: "IL",
				City:               "Chicago",
				Country:            "US",
				PostalCode:         "60601",
				StreetAddress:      "100 Main St",
				ExtendedAddress:    telnyx.String("Suite 504"),
			},
			BillingContact: telnyx.BillingContactParam{
				Email:       "billing@acmeplumbing.example.com",
				FirstName:   "Alex",
				LastName:    "Bill",
				PhoneNumber: "+13125550001",
			},
			CorporateRegistrationNumber: telnyx.String("corporate_registration_number"),
			CustomerReference:           telnyx.String("internal-ref-2026Q2"),
			DoingBusinessAs:             telnyx.String("Acme Plumbing"),
			DunBradstreetNumber:         telnyx.String("dun_bradstreet_number"),
			Fein:                        telnyx.String("12-3456789"),
			Industry:                    telnyx.EnterpriseUpdateParamsIndustryBusiness,
			JurisdictionOfIncorporation: telnyx.String("Delaware"),
			LegalName:                   telnyx.String("Acme Plumbing LLC"),
			NumberOfEmployees:           telnyx.String("51-200"),
			OrganizationContact: telnyx.OrganizationContactParam{
				Email:       "sam@acmeplumbing.example.com",
				FirstName:   "Sam",
				JobTitle:    "Compliance Lead",
				LastName:    "Owner",
				PhoneNumber: "+13125550000",
			},
			OrganizationLegalType: telnyx.String("llc"),
			OrganizationPhysicalAddress: telnyx.PhysicalAddressParam{
				AdministrativeArea: "IL",
				City:               "Chicago",
				Country:            "US",
				PostalCode:         "60601",
				StreetAddress:      "100 Main St",
				ExtendedAddress:    telnyx.String("Suite 504"),
			},
			PrimaryBusinessDomainSicCode: telnyx.String("primary_business_domain_sic_code"),
			ProfessionalLicenseNumber:    telnyx.String("professional_license_number"),
			Website:                      telnyx.String("https://acmeplumbing.example.com"),
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

func TestEnterpriseListWithOptionalParams(t *testing.T) {
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
	_, err := client.Enterprises.List(context.TODO(), telnyx.EnterpriseListParams{
		LegalName:  telnyx.String("Acme"),
		PageNumber: telnyx.Int(1),
		PageSize:   telnyx.Int(10),
	})
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnterpriseDelete(t *testing.T) {
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
	err := client.Enterprises.Delete(context.TODO(), "4a6192a4-573d-446d-b3ce-aff9117272a6")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEnterpriseActivateBrandedCalling(t *testing.T) {
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
	_, err := client.Enterprises.ActivateBrandedCalling(context.TODO(), "4a6192a4-573d-446d-b3ce-aff9117272a6")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
