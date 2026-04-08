// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Enterprise management for Branded Calling and Number Reputation services
//
// EnterpriseService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseService] method instead.
type EnterpriseService struct {
	Options []option.RequestOption
	// Manage Number Reputation enrollment and check frequency settings for an
	// enterprise
	Reputation EnterpriseReputationService
}

// NewEnterpriseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnterpriseService(opts ...option.RequestOption) (r EnterpriseService) {
	r = EnterpriseService{}
	r.Options = opts
	r.Reputation = NewEnterpriseReputationService(opts...)
	return
}

// Create a new enterprise for Branded Calling / Number Reputation services.
//
// Registers the enterprise in the Branded Calling / Number Reputation services,
// enabling it to create Display Identity Records (DIRs) or enroll in Number
// Reputation monitoring.
//
// **Required Fields:** `legal_name`, `doing_business_as`, `organization_type`,
// `country_code`, `website`, `fein`, `industry`, `number_of_employees`,
// `organization_legal_type`, `organization_contact`, `billing_contact`,
// `organization_physical_address`, `billing_address`
func (r *EnterpriseService) New(ctx context.Context, body EnterpriseNewParams, opts ...option.RequestOption) (res *EnterpriseNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "enterprises"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details of a specific enterprise by ID.
func (r *EnterpriseService) Get(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *EnterpriseGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update enterprise information. All fields are optional — only the provided
// fields will be updated.
func (r *EnterpriseService) Update(ctx context.Context, enterpriseID string, body EnterpriseUpdateParams, opts ...option.RequestOption) (res *EnterpriseUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of enterprises associated with your account.
func (r *EnterpriseService) List(ctx context.Context, query EnterpriseListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EnterpriseListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "enterprises"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a paginated list of enterprises associated with your account.
func (r *EnterpriseService) ListAutoPaging(ctx context.Context, query EnterpriseListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EnterpriseListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete an enterprise and all associated resources. This action is irreversible.
func (r *EnterpriseService) Delete(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return err
	}
	path := fmt.Sprintf("enterprises/%s", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type EnterpriseNewResponse struct {
	Data EnterpriseNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseNewResponseData struct {
	// Unique identifier of the enterprise
	ID             string                                  `json:"id" format:"uuid"`
	BillingAddress EnterpriseNewResponseDataBillingAddress `json:"billing_address"`
	BillingContact EnterpriseNewResponseDataBillingContact `json:"billing_contact"`
	// Corporate registration number
	CorporateRegistrationNumber string `json:"corporate_registration_number" api:"nullable"`
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// When the enterprise was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Customer reference identifier
	CustomerReference string `json:"customer_reference" api:"nullable"`
	// DBA name
	DoingBusinessAs string `json:"doing_business_as"`
	// D-U-N-S Number
	DunBradstreetNumber string `json:"dun_bradstreet_number" api:"nullable"`
	// Federal Employer Identification Number
	Fein string `json:"fein" api:"nullable"`
	// Industry classification
	Industry string `json:"industry" api:"nullable"`
	// Legal name of the enterprise
	LegalName string `json:"legal_name"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees string `json:"number_of_employees" api:"nullable"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact EnterpriseNewResponseDataOrganizationContact `json:"organization_contact"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       string                                               `json:"organization_legal_type" api:"nullable"`
	OrganizationPhysicalAddress EnterpriseNewResponseDataOrganizationPhysicalAddress `json:"organization_physical_address"`
	// Type of organization
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType string `json:"organization_type"`
	// SIC Code
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// Professional license number
	ProfessionalLicenseNumber string `json:"professional_license_number" api:"nullable"`
	// Role type in Branded Calling / Number Reputation services
	//
	// Any of "enterprise", "bpo".
	RoleType string `json:"role_type"`
	// When the enterprise was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Company website URL
	Website string `json:"website" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		BillingAddress               respjson.Field
		BillingContact               respjson.Field
		CorporateRegistrationNumber  respjson.Field
		CountryCode                  respjson.Field
		CreatedAt                    respjson.Field
		CustomerReference            respjson.Field
		DoingBusinessAs              respjson.Field
		DunBradstreetNumber          respjson.Field
		Fein                         respjson.Field
		Industry                     respjson.Field
		LegalName                    respjson.Field
		NumberOfEmployees            respjson.Field
		OrganizationContact          respjson.Field
		OrganizationLegalType        respjson.Field
		OrganizationPhysicalAddress  respjson.Field
		OrganizationType             respjson.Field
		PrimaryBusinessDomainSicCode respjson.Field
		ProfessionalLicenseNumber    respjson.Field
		RoleType                     respjson.Field
		UpdatedAt                    respjson.Field
		Website                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseNewResponseDataBillingAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseNewResponseDataBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseNewResponseDataBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseNewResponseDataBillingContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number (10-15 digits)
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseNewResponseDataBillingContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseNewResponseDataBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
type EnterpriseNewResponseDataOrganizationContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's job title (required)
	JobTitle string `json:"job_title" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		JobTitle    respjson.Field
		LastName    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseNewResponseDataOrganizationContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseNewResponseDataOrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseNewResponseDataOrganizationPhysicalAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseNewResponseDataOrganizationPhysicalAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseNewResponseDataOrganizationPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseGetResponse struct {
	Data EnterpriseGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseGetResponseData struct {
	// Unique identifier of the enterprise
	ID             string                                  `json:"id" format:"uuid"`
	BillingAddress EnterpriseGetResponseDataBillingAddress `json:"billing_address"`
	BillingContact EnterpriseGetResponseDataBillingContact `json:"billing_contact"`
	// Corporate registration number
	CorporateRegistrationNumber string `json:"corporate_registration_number" api:"nullable"`
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// When the enterprise was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Customer reference identifier
	CustomerReference string `json:"customer_reference" api:"nullable"`
	// DBA name
	DoingBusinessAs string `json:"doing_business_as"`
	// D-U-N-S Number
	DunBradstreetNumber string `json:"dun_bradstreet_number" api:"nullable"`
	// Federal Employer Identification Number
	Fein string `json:"fein" api:"nullable"`
	// Industry classification
	Industry string `json:"industry" api:"nullable"`
	// Legal name of the enterprise
	LegalName string `json:"legal_name"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees string `json:"number_of_employees" api:"nullable"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact EnterpriseGetResponseDataOrganizationContact `json:"organization_contact"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       string                                               `json:"organization_legal_type" api:"nullable"`
	OrganizationPhysicalAddress EnterpriseGetResponseDataOrganizationPhysicalAddress `json:"organization_physical_address"`
	// Type of organization
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType string `json:"organization_type"`
	// SIC Code
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// Professional license number
	ProfessionalLicenseNumber string `json:"professional_license_number" api:"nullable"`
	// Role type in Branded Calling / Number Reputation services
	//
	// Any of "enterprise", "bpo".
	RoleType string `json:"role_type"`
	// When the enterprise was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Company website URL
	Website string `json:"website" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		BillingAddress               respjson.Field
		BillingContact               respjson.Field
		CorporateRegistrationNumber  respjson.Field
		CountryCode                  respjson.Field
		CreatedAt                    respjson.Field
		CustomerReference            respjson.Field
		DoingBusinessAs              respjson.Field
		DunBradstreetNumber          respjson.Field
		Fein                         respjson.Field
		Industry                     respjson.Field
		LegalName                    respjson.Field
		NumberOfEmployees            respjson.Field
		OrganizationContact          respjson.Field
		OrganizationLegalType        respjson.Field
		OrganizationPhysicalAddress  respjson.Field
		OrganizationType             respjson.Field
		PrimaryBusinessDomainSicCode respjson.Field
		ProfessionalLicenseNumber    respjson.Field
		RoleType                     respjson.Field
		UpdatedAt                    respjson.Field
		Website                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseGetResponseDataBillingAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseGetResponseDataBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseGetResponseDataBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseGetResponseDataBillingContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number (10-15 digits)
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseGetResponseDataBillingContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseGetResponseDataBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
type EnterpriseGetResponseDataOrganizationContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's job title (required)
	JobTitle string `json:"job_title" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		JobTitle    respjson.Field
		LastName    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseGetResponseDataOrganizationContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseGetResponseDataOrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseGetResponseDataOrganizationPhysicalAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseGetResponseDataOrganizationPhysicalAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseGetResponseDataOrganizationPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseUpdateResponse struct {
	Data EnterpriseUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseUpdateResponseData struct {
	// Unique identifier of the enterprise
	ID             string                                     `json:"id" format:"uuid"`
	BillingAddress EnterpriseUpdateResponseDataBillingAddress `json:"billing_address"`
	BillingContact EnterpriseUpdateResponseDataBillingContact `json:"billing_contact"`
	// Corporate registration number
	CorporateRegistrationNumber string `json:"corporate_registration_number" api:"nullable"`
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// When the enterprise was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Customer reference identifier
	CustomerReference string `json:"customer_reference" api:"nullable"`
	// DBA name
	DoingBusinessAs string `json:"doing_business_as"`
	// D-U-N-S Number
	DunBradstreetNumber string `json:"dun_bradstreet_number" api:"nullable"`
	// Federal Employer Identification Number
	Fein string `json:"fein" api:"nullable"`
	// Industry classification
	Industry string `json:"industry" api:"nullable"`
	// Legal name of the enterprise
	LegalName string `json:"legal_name"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees string `json:"number_of_employees" api:"nullable"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact EnterpriseUpdateResponseDataOrganizationContact `json:"organization_contact"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       string                                                  `json:"organization_legal_type" api:"nullable"`
	OrganizationPhysicalAddress EnterpriseUpdateResponseDataOrganizationPhysicalAddress `json:"organization_physical_address"`
	// Type of organization
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType string `json:"organization_type"`
	// SIC Code
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// Professional license number
	ProfessionalLicenseNumber string `json:"professional_license_number" api:"nullable"`
	// Role type in Branded Calling / Number Reputation services
	//
	// Any of "enterprise", "bpo".
	RoleType string `json:"role_type"`
	// When the enterprise was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Company website URL
	Website string `json:"website" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		BillingAddress               respjson.Field
		BillingContact               respjson.Field
		CorporateRegistrationNumber  respjson.Field
		CountryCode                  respjson.Field
		CreatedAt                    respjson.Field
		CustomerReference            respjson.Field
		DoingBusinessAs              respjson.Field
		DunBradstreetNumber          respjson.Field
		Fein                         respjson.Field
		Industry                     respjson.Field
		LegalName                    respjson.Field
		NumberOfEmployees            respjson.Field
		OrganizationContact          respjson.Field
		OrganizationLegalType        respjson.Field
		OrganizationPhysicalAddress  respjson.Field
		OrganizationType             respjson.Field
		PrimaryBusinessDomainSicCode respjson.Field
		ProfessionalLicenseNumber    respjson.Field
		RoleType                     respjson.Field
		UpdatedAt                    respjson.Field
		Website                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseUpdateResponseDataBillingAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseUpdateResponseDataBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseUpdateResponseDataBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseUpdateResponseDataBillingContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number (10-15 digits)
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseUpdateResponseDataBillingContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseUpdateResponseDataBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
type EnterpriseUpdateResponseDataOrganizationContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's job title (required)
	JobTitle string `json:"job_title" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		JobTitle    respjson.Field
		LastName    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseUpdateResponseDataOrganizationContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseUpdateResponseDataOrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseUpdateResponseDataOrganizationPhysicalAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseUpdateResponseDataOrganizationPhysicalAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseUpdateResponseDataOrganizationPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseListResponse struct {
	// Unique identifier of the enterprise
	ID             string                               `json:"id" format:"uuid"`
	BillingAddress EnterpriseListResponseBillingAddress `json:"billing_address"`
	BillingContact EnterpriseListResponseBillingContact `json:"billing_contact"`
	// Corporate registration number
	CorporateRegistrationNumber string `json:"corporate_registration_number" api:"nullable"`
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// When the enterprise was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Customer reference identifier
	CustomerReference string `json:"customer_reference" api:"nullable"`
	// DBA name
	DoingBusinessAs string `json:"doing_business_as"`
	// D-U-N-S Number
	DunBradstreetNumber string `json:"dun_bradstreet_number" api:"nullable"`
	// Federal Employer Identification Number
	Fein string `json:"fein" api:"nullable"`
	// Industry classification
	Industry string `json:"industry" api:"nullable"`
	// Legal name of the enterprise
	LegalName string `json:"legal_name"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees EnterpriseListResponseNumberOfEmployees `json:"number_of_employees" api:"nullable"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact EnterpriseListResponseOrganizationContact `json:"organization_contact"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterpriseListResponseOrganizationLegalType       `json:"organization_legal_type" api:"nullable"`
	OrganizationPhysicalAddress EnterpriseListResponseOrganizationPhysicalAddress `json:"organization_physical_address"`
	// Type of organization
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType EnterpriseListResponseOrganizationType `json:"organization_type"`
	// SIC Code
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// Professional license number
	ProfessionalLicenseNumber string `json:"professional_license_number" api:"nullable"`
	// Role type in Branded Calling / Number Reputation services
	//
	// Any of "enterprise", "bpo".
	RoleType EnterpriseListResponseRoleType `json:"role_type"`
	// When the enterprise was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Company website URL
	Website string `json:"website" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		BillingAddress               respjson.Field
		BillingContact               respjson.Field
		CorporateRegistrationNumber  respjson.Field
		CountryCode                  respjson.Field
		CreatedAt                    respjson.Field
		CustomerReference            respjson.Field
		DoingBusinessAs              respjson.Field
		DunBradstreetNumber          respjson.Field
		Fein                         respjson.Field
		Industry                     respjson.Field
		LegalName                    respjson.Field
		NumberOfEmployees            respjson.Field
		OrganizationContact          respjson.Field
		OrganizationLegalType        respjson.Field
		OrganizationPhysicalAddress  respjson.Field
		OrganizationType             respjson.Field
		PrimaryBusinessDomainSicCode respjson.Field
		ProfessionalLicenseNumber    respjson.Field
		RoleType                     respjson.Field
		UpdatedAt                    respjson.Field
		Website                      respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseListResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseListResponseBillingAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseListResponseBillingAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseListResponseBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseListResponseBillingContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number (10-15 digits)
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseListResponseBillingContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseListResponseBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee count range
type EnterpriseListResponseNumberOfEmployees string

const (
	EnterpriseListResponseNumberOfEmployees1_10       EnterpriseListResponseNumberOfEmployees = "1-10"
	EnterpriseListResponseNumberOfEmployees11_50      EnterpriseListResponseNumberOfEmployees = "11-50"
	EnterpriseListResponseNumberOfEmployees51_200     EnterpriseListResponseNumberOfEmployees = "51-200"
	EnterpriseListResponseNumberOfEmployees201_500    EnterpriseListResponseNumberOfEmployees = "201-500"
	EnterpriseListResponseNumberOfEmployees501_2000   EnterpriseListResponseNumberOfEmployees = "501-2000"
	EnterpriseListResponseNumberOfEmployees2001_10000 EnterpriseListResponseNumberOfEmployees = "2001-10000"
	EnterpriseListResponseNumberOfEmployees10001      EnterpriseListResponseNumberOfEmployees = "10001+"
)

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
type EnterpriseListResponseOrganizationContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's job title (required)
	JobTitle string `json:"job_title" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		JobTitle    respjson.Field
		LastName    respjson.Field
		Phone       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseListResponseOrganizationContact) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseListResponseOrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legal structure type
type EnterpriseListResponseOrganizationLegalType string

const (
	EnterpriseListResponseOrganizationLegalTypeCorporation EnterpriseListResponseOrganizationLegalType = "corporation"
	EnterpriseListResponseOrganizationLegalTypeLlc         EnterpriseListResponseOrganizationLegalType = "llc"
	EnterpriseListResponseOrganizationLegalTypePartnership EnterpriseListResponseOrganizationLegalType = "partnership"
	EnterpriseListResponseOrganizationLegalTypeNonprofit   EnterpriseListResponseOrganizationLegalType = "nonprofit"
	EnterpriseListResponseOrganizationLegalTypeOther       EnterpriseListResponseOrganizationLegalType = "other"
)

type EnterpriseListResponseOrganizationPhysicalAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress string `json:"extended_address" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		Country            respjson.Field
		PostalCode         respjson.Field
		StreetAddress      respjson.Field
		ExtendedAddress    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseListResponseOrganizationPhysicalAddress) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseListResponseOrganizationPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of organization
type EnterpriseListResponseOrganizationType string

const (
	EnterpriseListResponseOrganizationTypeCommercial EnterpriseListResponseOrganizationType = "commercial"
	EnterpriseListResponseOrganizationTypeGovernment EnterpriseListResponseOrganizationType = "government"
	EnterpriseListResponseOrganizationTypeNonProfit  EnterpriseListResponseOrganizationType = "non_profit"
)

// Role type in Branded Calling / Number Reputation services
type EnterpriseListResponseRoleType string

const (
	EnterpriseListResponseRoleTypeEnterprise EnterpriseListResponseRoleType = "enterprise"
	EnterpriseListResponseRoleTypeBpo        EnterpriseListResponseRoleType = "bpo"
)

type EnterpriseNewParams struct {
	BillingAddress EnterpriseNewParamsBillingAddress `json:"billing_address,omitzero" api:"required"`
	BillingContact EnterpriseNewParamsBillingContact `json:"billing_contact,omitzero" api:"required"`
	// Country code. Currently only 'US' is accepted.
	CountryCode string `json:"country_code" api:"required"`
	// Primary business name / DBA name
	DoingBusinessAs string `json:"doing_business_as" api:"required"`
	// Federal Employer Identification Number. Format: XX-XXXXXXX or 9-digit number
	// (minimum 9 digits).
	Fein string `json:"fein" api:"required"`
	// Industry classification. Case-insensitive. Accepted values: accounting, finance,
	// billing, collections, business, charity, nonprofit, communications, telecom,
	// customer service, support, delivery, shipping, logistics, education, financial,
	// banking, government, public, healthcare, health, pharmacy, medical, insurance,
	// legal, law, notifications, scheduling, real estate, property, retail, ecommerce,
	// sales, marketing, software, technology, tech, media, surveys, market research,
	// travel, hospitality, hotel
	Industry string `json:"industry" api:"required"`
	// Legal name of the enterprise
	LegalName string `json:"legal_name" api:"required"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees EnterpriseNewParamsNumberOfEmployees `json:"number_of_employees,omitzero" api:"required"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact EnterpriseNewParamsOrganizationContact `json:"organization_contact,omitzero" api:"required"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterpriseNewParamsOrganizationLegalType       `json:"organization_legal_type,omitzero" api:"required"`
	OrganizationPhysicalAddress EnterpriseNewParamsOrganizationPhysicalAddress `json:"organization_physical_address,omitzero" api:"required"`
	// Type of organization
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType EnterpriseNewParamsOrganizationType `json:"organization_type,omitzero" api:"required"`
	// Enterprise website URL. Accepts any string — no URL format validation enforced.
	Website string `json:"website" api:"required"`
	// Corporate registration number (optional)
	CorporateRegistrationNumber param.Opt[string] `json:"corporate_registration_number,omitzero"`
	// Optional customer reference identifier for your own tracking
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// D-U-N-S Number (optional)
	DunBradstreetNumber param.Opt[string] `json:"dun_bradstreet_number,omitzero"`
	// SIC Code (optional)
	PrimaryBusinessDomainSicCode param.Opt[string] `json:"primary_business_domain_sic_code,omitzero"`
	// Professional license number (optional)
	ProfessionalLicenseNumber param.Opt[string] `json:"professional_license_number,omitzero"`
	// Role type in Branded Calling / Number Reputation services
	//
	// Any of "enterprise", "bpo".
	RoleType EnterpriseNewParamsRoleType `json:"role_type,omitzero"`
	paramObj
}

func (r EnterpriseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AdministrativeArea, City, Country, PostalCode, StreetAddress are
// required.
type EnterpriseNewParamsBillingAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r EnterpriseNewParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseNewParamsBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseNewParamsBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Email, FirstName, LastName, PhoneNumber are required.
type EnterpriseNewParamsBillingContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number (10-15 digits)
	PhoneNumber string `json:"phone_number" api:"required"`
	paramObj
}

func (r EnterpriseNewParamsBillingContact) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseNewParamsBillingContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseNewParamsBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee count range
type EnterpriseNewParamsNumberOfEmployees string

const (
	EnterpriseNewParamsNumberOfEmployees1_10       EnterpriseNewParamsNumberOfEmployees = "1-10"
	EnterpriseNewParamsNumberOfEmployees11_50      EnterpriseNewParamsNumberOfEmployees = "11-50"
	EnterpriseNewParamsNumberOfEmployees51_200     EnterpriseNewParamsNumberOfEmployees = "51-200"
	EnterpriseNewParamsNumberOfEmployees201_500    EnterpriseNewParamsNumberOfEmployees = "201-500"
	EnterpriseNewParamsNumberOfEmployees501_2000   EnterpriseNewParamsNumberOfEmployees = "501-2000"
	EnterpriseNewParamsNumberOfEmployees2001_10000 EnterpriseNewParamsNumberOfEmployees = "2001-10000"
	EnterpriseNewParamsNumberOfEmployees10001      EnterpriseNewParamsNumberOfEmployees = "10001+"
)

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
//
// The properties Email, FirstName, JobTitle, LastName, Phone are required.
type EnterpriseNewParamsOrganizationContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's job title (required)
	JobTitle string `json:"job_title" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	paramObj
}

func (r EnterpriseNewParamsOrganizationContact) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseNewParamsOrganizationContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseNewParamsOrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legal structure type
type EnterpriseNewParamsOrganizationLegalType string

const (
	EnterpriseNewParamsOrganizationLegalTypeCorporation EnterpriseNewParamsOrganizationLegalType = "corporation"
	EnterpriseNewParamsOrganizationLegalTypeLlc         EnterpriseNewParamsOrganizationLegalType = "llc"
	EnterpriseNewParamsOrganizationLegalTypePartnership EnterpriseNewParamsOrganizationLegalType = "partnership"
	EnterpriseNewParamsOrganizationLegalTypeNonprofit   EnterpriseNewParamsOrganizationLegalType = "nonprofit"
	EnterpriseNewParamsOrganizationLegalTypeOther       EnterpriseNewParamsOrganizationLegalType = "other"
)

// The properties AdministrativeArea, City, Country, PostalCode, StreetAddress are
// required.
type EnterpriseNewParamsOrganizationPhysicalAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r EnterpriseNewParamsOrganizationPhysicalAddress) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseNewParamsOrganizationPhysicalAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseNewParamsOrganizationPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of organization
type EnterpriseNewParamsOrganizationType string

const (
	EnterpriseNewParamsOrganizationTypeCommercial EnterpriseNewParamsOrganizationType = "commercial"
	EnterpriseNewParamsOrganizationTypeGovernment EnterpriseNewParamsOrganizationType = "government"
	EnterpriseNewParamsOrganizationTypeNonProfit  EnterpriseNewParamsOrganizationType = "non_profit"
)

// Role type in Branded Calling / Number Reputation services
type EnterpriseNewParamsRoleType string

const (
	EnterpriseNewParamsRoleTypeEnterprise EnterpriseNewParamsRoleType = "enterprise"
	EnterpriseNewParamsRoleTypeBpo        EnterpriseNewParamsRoleType = "bpo"
)

type EnterpriseUpdateParams struct {
	// Corporate registration number
	CorporateRegistrationNumber param.Opt[string] `json:"corporate_registration_number,omitzero"`
	// Customer reference identifier
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// DBA name
	DoingBusinessAs param.Opt[string] `json:"doing_business_as,omitzero"`
	// D-U-N-S Number
	DunBradstreetNumber param.Opt[string] `json:"dun_bradstreet_number,omitzero"`
	// Federal Employer Identification Number. Format: XX-XXXXXXX or XXXXXXXXX
	Fein param.Opt[string] `json:"fein,omitzero"`
	// Industry classification
	Industry param.Opt[string] `json:"industry,omitzero"`
	// Legal name of the enterprise
	LegalName param.Opt[string] `json:"legal_name,omitzero"`
	// SIC Code
	PrimaryBusinessDomainSicCode param.Opt[string] `json:"primary_business_domain_sic_code,omitzero"`
	// Professional license number
	ProfessionalLicenseNumber param.Opt[string] `json:"professional_license_number,omitzero"`
	// Company website URL
	Website        param.Opt[string]                    `json:"website,omitzero"`
	BillingAddress EnterpriseUpdateParamsBillingAddress `json:"billing_address,omitzero"`
	BillingContact EnterpriseUpdateParamsBillingContact `json:"billing_contact,omitzero"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees EnterpriseUpdateParamsNumberOfEmployees `json:"number_of_employees,omitzero"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact EnterpriseUpdateParamsOrganizationContact `json:"organization_contact,omitzero"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterpriseUpdateParamsOrganizationLegalType       `json:"organization_legal_type,omitzero"`
	OrganizationPhysicalAddress EnterpriseUpdateParamsOrganizationPhysicalAddress `json:"organization_physical_address,omitzero"`
	paramObj
}

func (r EnterpriseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AdministrativeArea, City, Country, PostalCode, StreetAddress are
// required.
type EnterpriseUpdateParamsBillingAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r EnterpriseUpdateParamsBillingAddress) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParamsBillingAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParamsBillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Email, FirstName, LastName, PhoneNumber are required.
type EnterpriseUpdateParamsBillingContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number (10-15 digits)
	PhoneNumber string `json:"phone_number" api:"required"`
	paramObj
}

func (r EnterpriseUpdateParamsBillingContact) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParamsBillingContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParamsBillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee count range
type EnterpriseUpdateParamsNumberOfEmployees string

const (
	EnterpriseUpdateParamsNumberOfEmployees1_10       EnterpriseUpdateParamsNumberOfEmployees = "1-10"
	EnterpriseUpdateParamsNumberOfEmployees11_50      EnterpriseUpdateParamsNumberOfEmployees = "11-50"
	EnterpriseUpdateParamsNumberOfEmployees51_200     EnterpriseUpdateParamsNumberOfEmployees = "51-200"
	EnterpriseUpdateParamsNumberOfEmployees201_500    EnterpriseUpdateParamsNumberOfEmployees = "201-500"
	EnterpriseUpdateParamsNumberOfEmployees501_2000   EnterpriseUpdateParamsNumberOfEmployees = "501-2000"
	EnterpriseUpdateParamsNumberOfEmployees2001_10000 EnterpriseUpdateParamsNumberOfEmployees = "2001-10000"
	EnterpriseUpdateParamsNumberOfEmployees10001      EnterpriseUpdateParamsNumberOfEmployees = "10001+"
)

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
//
// The properties Email, FirstName, JobTitle, LastName, Phone are required.
type EnterpriseUpdateParamsOrganizationContact struct {
	// Contact's email address
	Email string `json:"email" api:"required" format:"email"`
	// Contact's first name
	FirstName string `json:"first_name" api:"required"`
	// Contact's job title (required)
	JobTitle string `json:"job_title" api:"required"`
	// Contact's last name
	LastName string `json:"last_name" api:"required"`
	// Contact's phone number in E.164 format
	Phone string `json:"phone" api:"required"`
	paramObj
}

func (r EnterpriseUpdateParamsOrganizationContact) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParamsOrganizationContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParamsOrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legal structure type
type EnterpriseUpdateParamsOrganizationLegalType string

const (
	EnterpriseUpdateParamsOrganizationLegalTypeCorporation EnterpriseUpdateParamsOrganizationLegalType = "corporation"
	EnterpriseUpdateParamsOrganizationLegalTypeLlc         EnterpriseUpdateParamsOrganizationLegalType = "llc"
	EnterpriseUpdateParamsOrganizationLegalTypePartnership EnterpriseUpdateParamsOrganizationLegalType = "partnership"
	EnterpriseUpdateParamsOrganizationLegalTypeNonprofit   EnterpriseUpdateParamsOrganizationLegalType = "nonprofit"
	EnterpriseUpdateParamsOrganizationLegalTypeOther       EnterpriseUpdateParamsOrganizationLegalType = "other"
)

// The properties AdministrativeArea, City, Country, PostalCode, StreetAddress are
// required.
type EnterpriseUpdateParamsOrganizationPhysicalAddress struct {
	// State or province
	AdministrativeArea string `json:"administrative_area" api:"required"`
	// City name
	City string `json:"city" api:"required"`
	// Country name (e.g., United States)
	Country string `json:"country" api:"required"`
	// ZIP or postal code
	PostalCode string `json:"postal_code" api:"required"`
	// Street address
	StreetAddress string `json:"street_address" api:"required"`
	// Additional address line (suite, apt, etc.)
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	paramObj
}

func (r EnterpriseUpdateParamsOrganizationPhysicalAddress) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParamsOrganizationPhysicalAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParamsOrganizationPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseListParams struct {
	// Filter by legal name (partial match)
	LegalName param.Opt[string] `query:"legal_name,omitzero" json:"-"`
	// Page number (1-indexed)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseListParams]'s query parameters as `url.Values`.
func (r EnterpriseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
