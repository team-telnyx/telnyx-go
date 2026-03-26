// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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
func (r *EnterpriseService) List(ctx context.Context, query EnterpriseListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EnterprisePublic], err error) {
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
func (r *EnterpriseService) ListAutoPaging(ctx context.Context, query EnterpriseListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EnterprisePublic] {
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

type BillingAddress struct {
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
func (r BillingAddress) RawJSON() string { return r.JSON.raw }
func (r *BillingAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BillingAddress to a BillingAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BillingAddressParam.Overrides()
func (r BillingAddress) ToParam() BillingAddressParam {
	return param.Override[BillingAddressParam](json.RawMessage(r.RawJSON()))
}

// The properties AdministrativeArea, City, Country, PostalCode, StreetAddress are
// required.
type BillingAddressParam struct {
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

func (r BillingAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow BillingAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingContact struct {
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
func (r BillingContact) RawJSON() string { return r.JSON.raw }
func (r *BillingContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BillingContact to a BillingContactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BillingContactParam.Overrides()
func (r BillingContact) ToParam() BillingContactParam {
	return param.Override[BillingContactParam](json.RawMessage(r.RawJSON()))
}

// The properties Email, FirstName, LastName, PhoneNumber are required.
type BillingContactParam struct {
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

func (r BillingContactParam) MarshalJSON() (data []byte, err error) {
	type shadow BillingContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterprisePublic struct {
	// Unique identifier of the enterprise
	ID             string         `json:"id" format:"uuid"`
	BillingAddress BillingAddress `json:"billing_address"`
	BillingContact BillingContact `json:"billing_contact"`
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
	NumberOfEmployees EnterprisePublicNumberOfEmployees `json:"number_of_employees" api:"nullable"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact OrganizationContact `json:"organization_contact"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterprisePublicOrganizationLegalType `json:"organization_legal_type" api:"nullable"`
	OrganizationPhysicalAddress PhysicalAddress                       `json:"organization_physical_address"`
	// Type of organization
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType EnterprisePublicOrganizationType `json:"organization_type"`
	// SIC Code
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// Professional license number
	ProfessionalLicenseNumber string `json:"professional_license_number" api:"nullable"`
	// Role type in Branded Calling / Number Reputation services
	//
	// Any of "enterprise", "bpo".
	RoleType EnterprisePublicRoleType `json:"role_type"`
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
func (r EnterprisePublic) RawJSON() string { return r.JSON.raw }
func (r *EnterprisePublic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employee count range
type EnterprisePublicNumberOfEmployees string

const (
	EnterprisePublicNumberOfEmployees1_10       EnterprisePublicNumberOfEmployees = "1-10"
	EnterprisePublicNumberOfEmployees11_50      EnterprisePublicNumberOfEmployees = "11-50"
	EnterprisePublicNumberOfEmployees51_200     EnterprisePublicNumberOfEmployees = "51-200"
	EnterprisePublicNumberOfEmployees201_500    EnterprisePublicNumberOfEmployees = "201-500"
	EnterprisePublicNumberOfEmployees501_2000   EnterprisePublicNumberOfEmployees = "501-2000"
	EnterprisePublicNumberOfEmployees2001_10000 EnterprisePublicNumberOfEmployees = "2001-10000"
	EnterprisePublicNumberOfEmployees10001      EnterprisePublicNumberOfEmployees = "10001+"
)

// Legal structure type
type EnterprisePublicOrganizationLegalType string

const (
	EnterprisePublicOrganizationLegalTypeCorporation EnterprisePublicOrganizationLegalType = "corporation"
	EnterprisePublicOrganizationLegalTypeLlc         EnterprisePublicOrganizationLegalType = "llc"
	EnterprisePublicOrganizationLegalTypePartnership EnterprisePublicOrganizationLegalType = "partnership"
	EnterprisePublicOrganizationLegalTypeNonprofit   EnterprisePublicOrganizationLegalType = "nonprofit"
	EnterprisePublicOrganizationLegalTypeOther       EnterprisePublicOrganizationLegalType = "other"
)

// Type of organization
type EnterprisePublicOrganizationType string

const (
	EnterprisePublicOrganizationTypeCommercial EnterprisePublicOrganizationType = "commercial"
	EnterprisePublicOrganizationTypeGovernment EnterprisePublicOrganizationType = "government"
	EnterprisePublicOrganizationTypeNonProfit  EnterprisePublicOrganizationType = "non_profit"
)

// Role type in Branded Calling / Number Reputation services
type EnterprisePublicRoleType string

const (
	EnterprisePublicRoleTypeEnterprise EnterprisePublicRoleType = "enterprise"
	EnterprisePublicRoleTypeBpo        EnterprisePublicRoleType = "bpo"
)

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
type OrganizationContact struct {
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
func (r OrganizationContact) RawJSON() string { return r.JSON.raw }
func (r *OrganizationContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OrganizationContact to a OrganizationContactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OrganizationContactParam.Overrides()
func (r OrganizationContact) ToParam() OrganizationContactParam {
	return param.Override[OrganizationContactParam](json.RawMessage(r.RawJSON()))
}

// Organization contact information. Note: the response returns this object with
// the phone field as 'phone' (not 'phone_number').
//
// The properties Email, FirstName, JobTitle, LastName, Phone are required.
type OrganizationContactParam struct {
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

func (r OrganizationContactParam) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhysicalAddress struct {
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
func (r PhysicalAddress) RawJSON() string { return r.JSON.raw }
func (r *PhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PhysicalAddress to a PhysicalAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PhysicalAddressParam.Overrides()
func (r PhysicalAddress) ToParam() PhysicalAddressParam {
	return param.Override[PhysicalAddressParam](json.RawMessage(r.RawJSON()))
}

// The properties AdministrativeArea, City, Country, PostalCode, StreetAddress are
// required.
type PhysicalAddressParam struct {
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

func (r PhysicalAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow PhysicalAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhysicalAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseNewResponse struct {
	Data EnterprisePublic `json:"data"`
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

type EnterpriseGetResponse struct {
	Data EnterprisePublic `json:"data"`
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

type EnterpriseUpdateResponse struct {
	Data EnterprisePublic `json:"data"`
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

type EnterpriseNewParams struct {
	BillingAddress BillingAddressParam `json:"billing_address,omitzero" api:"required"`
	BillingContact BillingContactParam `json:"billing_contact,omitzero" api:"required"`
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
	OrganizationContact OrganizationContactParam `json:"organization_contact,omitzero" api:"required"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterpriseNewParamsOrganizationLegalType `json:"organization_legal_type,omitzero" api:"required"`
	OrganizationPhysicalAddress PhysicalAddressParam                     `json:"organization_physical_address,omitzero" api:"required"`
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

// Legal structure type
type EnterpriseNewParamsOrganizationLegalType string

const (
	EnterpriseNewParamsOrganizationLegalTypeCorporation EnterpriseNewParamsOrganizationLegalType = "corporation"
	EnterpriseNewParamsOrganizationLegalTypeLlc         EnterpriseNewParamsOrganizationLegalType = "llc"
	EnterpriseNewParamsOrganizationLegalTypePartnership EnterpriseNewParamsOrganizationLegalType = "partnership"
	EnterpriseNewParamsOrganizationLegalTypeNonprofit   EnterpriseNewParamsOrganizationLegalType = "nonprofit"
	EnterpriseNewParamsOrganizationLegalTypeOther       EnterpriseNewParamsOrganizationLegalType = "other"
)

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
	Website        param.Opt[string]   `json:"website,omitzero"`
	BillingAddress BillingAddressParam `json:"billing_address,omitzero"`
	BillingContact BillingContactParam `json:"billing_contact,omitzero"`
	// Employee count range
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees EnterpriseUpdateParamsNumberOfEmployees `json:"number_of_employees,omitzero"`
	// Organization contact information. Note: the response returns this object with
	// the phone field as 'phone' (not 'phone_number').
	OrganizationContact OrganizationContactParam `json:"organization_contact,omitzero"`
	// Legal structure type
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterpriseUpdateParamsOrganizationLegalType `json:"organization_legal_type,omitzero"`
	OrganizationPhysicalAddress PhysicalAddressParam                        `json:"organization_physical_address,omitzero"`
	paramObj
}

func (r EnterpriseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParams) UnmarshalJSON(data []byte) error {
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

// Legal structure type
type EnterpriseUpdateParamsOrganizationLegalType string

const (
	EnterpriseUpdateParamsOrganizationLegalTypeCorporation EnterpriseUpdateParamsOrganizationLegalType = "corporation"
	EnterpriseUpdateParamsOrganizationLegalTypeLlc         EnterpriseUpdateParamsOrganizationLegalType = "llc"
	EnterpriseUpdateParamsOrganizationLegalTypePartnership EnterpriseUpdateParamsOrganizationLegalType = "partnership"
	EnterpriseUpdateParamsOrganizationLegalTypeNonprofit   EnterpriseUpdateParamsOrganizationLegalType = "nonprofit"
	EnterpriseUpdateParamsOrganizationLegalTypeOther       EnterpriseUpdateParamsOrganizationLegalType = "other"
)

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
