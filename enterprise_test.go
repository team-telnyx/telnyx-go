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
			AdministrativeArea: "Illinois",
			City:               "Chicago",
			Country:            "United States",
			PostalCode:         "60601",
			StreetAddress:      "123 Main St",
			ExtendedAddress:    telnyx.String("Suite 400"),
		},
		BillingContact: telnyx.BillingContactParam{
			Email:       "billing@acme.com",
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "15551234568",
		},
		CountryCode:       "US",
		DoingBusinessAs:   "Acme",
		Fein:              "12-3456789",
		Industry:          "technology",
		LegalName:         "Acme Corp Inc.",
		NumberOfEmployees: telnyx.EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees51_200,
		OrganizationContact: telnyx.OrganizationContactParam{
			Email:     "jane.smith@acme.com",
			FirstName: "Jane",
			JobTitle:  "VP of Engineering",
			LastName:  "Smith",
			Phone:     "+16035551234",
		},
		OrganizationLegalType: telnyx.EnterpriseNewParamsOrganizationLegalTypeCorporation,
		OrganizationPhysicalAddress: telnyx.PhysicalAddressParam{
			AdministrativeArea: "Illinois",
			City:               "Chicago",
			Country:            "United States",
			PostalCode:         "60601",
			StreetAddress:      "123 Main St",
			ExtendedAddress:    telnyx.String("Suite 400"),
		},
		OrganizationType:             telnyx.EnterpriseNewParamsOrganizationTypeCommercial,
		Website:                      "https://acme.com",
		CorporateRegistrationNumber:  telnyx.String("corporate_registration_number"),
		CustomerReference:            telnyx.String("customer_reference"),
		DunBradstreetNumber:          telnyx.String("dun_bradstreet_number"),
		PrimaryBusinessDomainSicCode: telnyx.String("7372"),
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
	_, err := client.Enterprises.Get(context.TODO(), "6a09cdc3-8948-47f0-aa62-74ac943d6c58")
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
		"6a09cdc3-8948-47f0-aa62-74ac943d6c58",
		telnyx.EnterpriseUpdateParams{
			BillingAddress: telnyx.BillingAddressParam{
				AdministrativeArea: "Illinois",
				City:               "Chicago",
				Country:            "United States",
				PostalCode:         "60601",
				StreetAddress:      "123 Main St",
				ExtendedAddress:    telnyx.String("Suite 400"),
			},
			BillingContact: telnyx.BillingContactParam{
				Email:       "billing@acme.com",
				FirstName:   "John",
				LastName:    "Doe",
				PhoneNumber: "15551234568",
			},
			CorporateRegistrationNumber: telnyx.String("corporate_registration_number"),
			CustomerReference:           telnyx.String("customer_reference"),
			DoingBusinessAs:             telnyx.String("doing_business_as"),
			DunBradstreetNumber:         telnyx.String("dun_bradstreet_number"),
			Fein:                        telnyx.String("fein"),
			Industry:                    telnyx.String("industry"),
			LegalName:                   telnyx.String("xxx"),
			NumberOfEmployees:           telnyx.EnterpriseUpdateParamsNumberOfEmployeesNumberOfEmployees1_10,
			OrganizationContact: telnyx.OrganizationContactParam{
				Email:     "jane.smith@acme.com",
				FirstName: "Jane",
				JobTitle:  "VP of Engineering",
				LastName:  "Smith",
				Phone:     "+16035551234",
			},
			OrganizationLegalType: telnyx.EnterpriseUpdateParamsOrganizationLegalTypeCorporation,
			OrganizationPhysicalAddress: telnyx.PhysicalAddressParam{
				AdministrativeArea: "Illinois",
				City:               "Chicago",
				Country:            "United States",
				PostalCode:         "60601",
				StreetAddress:      "123 Main St",
				ExtendedAddress:    telnyx.String("Suite 400"),
			},
			PrimaryBusinessDomainSicCode: telnyx.String("primary_business_domain_sic_code"),
			ProfessionalLicenseNumber:    telnyx.String("professional_license_number"),
			Website:                      telnyx.String("website"),
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
		PageSize:   telnyx.Int(1),
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
	err := client.Enterprises.Delete(context.TODO(), "6a09cdc3-8948-47f0-aa62-74ac943d6c58")
	if err != nil {
		var apierr *telnyx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
