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

// Manage the legal-entity record that owns your DIRs and phone numbers.
//
// EnterpriseService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseService] method instead.
type EnterpriseService struct {
	Options []option.RequestOption
	// Phone-number reputation monitoring (spam-score lookup and tracking).
	Reputation EnterpriseReputationService
	// A Display Identity Record (DIR) is the verified calling identity (display name,
	// logo, call reasons) shown to recipients on outbound calls.
	Dir EnterpriseDirService
}

// NewEnterpriseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnterpriseService(opts ...option.RequestOption) (r EnterpriseService) {
	r = EnterpriseService{}
	r.Options = opts
	r.Reputation = NewEnterpriseReputationService(opts...)
	r.Dir = NewEnterpriseDirService(opts...)
	return
}

// Create the legal entity (enterprise) that represents your business on the Telnyx
// platform.
//
// The response carries a server-assigned `id` you use for every subsequent call.
// An enterprise is created once and reused; the API collects all required fields
// up front.
//
// Common failure modes:
//
//   - `422` - a required field is missing or malformed (the response
//     `errors[].source.pointer` names the field).
//   - `409` - an enterprise with the same identifying details already exists under
//     your account.
func (r *EnterpriseService) New(ctx context.Context, body EnterpriseNewParams, opts ...option.RequestOption) (res *EnterprisePublicWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "enterprises"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a single enterprise by id. Returns `404` if the id does not exist or
// does not belong to your account.
func (r *EnterpriseService) Get(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *EnterprisePublicWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replace the enterprise's mutable fields. Only mutable fields may be sent.
// Server-assigned and immutable fields (`id`, `record_type`, `created_at`,
// `updated_at`, status fields, `organization_type`, `country_code`, `role_type`)
// cannot be changed: including any of them in the body is rejected with
// `400 Bad Request` (`Field 'X' is not allowed in this request`).
func (r *EnterpriseService) Update(ctx context.Context, enterpriseID string, body EnterpriseUpdateParams, opts ...option.RequestOption) (res *EnterprisePublicWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Return the enterprises you own, paginated. The default page size is 20; the
// maximum is 250.
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

// Return the enterprises you own, paginated. The default page size is 20; the
// maximum is 250.
func (r *EnterpriseService) ListAutoPaging(ctx context.Context, query EnterpriseListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EnterprisePublic] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Soft-delete an enterprise.
//
// Failure modes:
//
//   - `400` - the enterprise still has dependent resources in a non-deletable state.
//     Remove those first; the response `detail` identifies what is blocking the
//     delete.
//   - `409` - the enterprise has a dependent resource with an unresolved claim.
//     Resolve it before deleting.
//   - `404` - the enterprise does not exist or does not belong to your account.
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

// Branded Calling is a paid product that must be activated on each enterprise.
// Activation is idempotent:
//
//   - First call: marks the enterprise as activated and begins onboarding it with
//     the Branded Calling platform asynchronously. Returns `200` with
//     `branded_calling_enabled: true`.
//   - Re-call after success: no-op, returns the same enterprise body.
//   - Re-call after a prior failure: re-queues onboarding, returns `200`.
//
// Prerequisite: the calling user must have agreed to the Branded Calling Terms of
// Service (`POST /terms_of_service/branded_calling/agree`). Without that, this
// endpoint returns `403 terms_of_service_not_accepted`.
//
// Failure modes:
//
// - `403` - Branded Calling Terms of Service not accepted.
// - `404` - enterprise does not exist or does not belong to your account.
//
// **Pricing:** This is a billable action. See https://telnyx.com/pricing/numbers
// for current pricing.
func (r *EnterpriseService) BrandedCalling(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *EnterprisePublicWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/branded_calling", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type BillingAddress struct {
	// State or province code (e.g. `IL`, `ON`).
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// ISO 3166-1 alpha-2 code (currently `US` or `CA`).
	Country         string `json:"country" api:"required"`
	PostalCode      string `json:"postal_code" api:"required"`
	StreetAddress   string `json:"street_address" api:"required"`
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
	// State or province code (e.g. `IL`, `ON`).
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// ISO 3166-1 alpha-2 code (currently `US` or `CA`).
	Country         string            `json:"country" api:"required"`
	PostalCode      string            `json:"postal_code" api:"required"`
	StreetAddress   string            `json:"street_address" api:"required"`
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
	Email     string `json:"email" api:"required" format:"email"`
	FirstName string `json:"first_name" api:"required"`
	LastName  string `json:"last_name" api:"required"`
	// E.164 format with leading `+`.
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
	Email     string `json:"email" api:"required" format:"email"`
	FirstName string `json:"first_name" api:"required"`
	LastName  string `json:"last_name" api:"required"`
	// E.164 format with leading `+`.
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
	ID             string         `json:"id" format:"uuid"`
	BillingAddress BillingAddress `json:"billing_address"`
	BillingContact BillingContact `json:"billing_contact"`
	// True once Branded Calling has been activated on this enterprise (see
	// `POST /enterprises/{id}/branded_calling`).
	BrandedCallingEnabled bool `json:"branded_calling_enabled"`
	// Optional corporate-registration / company-number identifier.
	CorporateRegistrationNumber string    `json:"corporate_registration_number" api:"nullable"`
	CountryCode                 string    `json:"country_code"`
	CreatedAt                   time.Time `json:"created_at" format:"date-time"`
	CustomerReference           string    `json:"customer_reference"`
	DoingBusinessAs             string    `json:"doing_business_as"`
	// Optional D-U-N-S Number issued by Dun & Bradstreet.
	DunBradstreetNumber         string `json:"dun_bradstreet_number" api:"nullable"`
	Fein                        string `json:"fein"`
	Industry                    string `json:"industry"`
	JurisdictionOfIncorporation string `json:"jurisdiction_of_incorporation"`
	LegalName                   string `json:"legal_name"`
	NumberOfEmployees           string `json:"number_of_employees"`
	// True once Phone Number Reputation has been enabled on this enterprise (see
	// `POST /enterprises/{id}/reputation`).
	NumberReputationEnabled     bool                `json:"number_reputation_enabled"`
	OrganizationContact         OrganizationContact `json:"organization_contact"`
	OrganizationLegalType       string              `json:"organization_legal_type"`
	OrganizationPhysicalAddress PhysicalAddress     `json:"organization_physical_address"`
	OrganizationType            string              `json:"organization_type"`
	// Optional SIC code for the primary line of business.
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// Optional professional-license number for regulated industries.
	ProfessionalLicenseNumber string    `json:"professional_license_number" api:"nullable"`
	RoleType                  string    `json:"role_type"`
	UpdatedAt                 time.Time `json:"updated_at" format:"date-time"`
	Website                   string    `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		BillingAddress               respjson.Field
		BillingContact               respjson.Field
		BrandedCallingEnabled        respjson.Field
		CorporateRegistrationNumber  respjson.Field
		CountryCode                  respjson.Field
		CreatedAt                    respjson.Field
		CustomerReference            respjson.Field
		DoingBusinessAs              respjson.Field
		DunBradstreetNumber          respjson.Field
		Fein                         respjson.Field
		Industry                     respjson.Field
		JurisdictionOfIncorporation  respjson.Field
		LegalName                    respjson.Field
		NumberOfEmployees            respjson.Field
		NumberReputationEnabled      respjson.Field
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

type EnterprisePublicWrapped struct {
	Data EnterprisePublic `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterprisePublicWrapped) RawJSON() string { return r.JSON.raw }
func (r *EnterprisePublicWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON:API pagination metadata returned with every paginated list response. Page
// numbering is 1-based. `page_size` reports the number of items actually returned
// in `data` for this page; the requested size is taken from the `page[size]` query
// parameter.
type NumberReputationPaginationMeta struct {
	// 1-based index of this page. Echoes the `page[number]` query parameter (default
	// `1`).
	PageNumber int64 `json:"page_number" api:"required"`
	// Number of items returned in this page's `data` array. Capped at 250.
	PageSize int64 `json:"page_size" api:"required"`
	// Total number of pages available given the current `page_size`.
	TotalPages int64 `json:"total_pages" api:"required"`
	// Total number of items across all pages (excludes soft-deleted rows).
	TotalResults int64 `json:"total_results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberReputationPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *NumberReputationPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationContact struct {
	Email     string `json:"email" api:"required" format:"email"`
	FirstName string `json:"first_name" api:"required"`
	JobTitle  string `json:"job_title" api:"required"`
	LastName  string `json:"last_name" api:"required"`
	// E.164 format with leading `+`.
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FirstName   respjson.Field
		JobTitle    respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
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

// The properties Email, FirstName, JobTitle, LastName, PhoneNumber are required.
type OrganizationContactParam struct {
	Email     string `json:"email" api:"required" format:"email"`
	FirstName string `json:"first_name" api:"required"`
	JobTitle  string `json:"job_title" api:"required"`
	LastName  string `json:"last_name" api:"required"`
	// E.164 format with leading `+`.
	PhoneNumber string `json:"phone_number" api:"required"`
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
	// State or province code (e.g. `IL`, `ON`).
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// ISO 3166-1 alpha-2 code (currently `US` or `CA`).
	Country         string `json:"country" api:"required"`
	PostalCode      string `json:"postal_code" api:"required"`
	StreetAddress   string `json:"street_address" api:"required"`
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
	// State or province code (e.g. `IL`, `ON`).
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// ISO 3166-1 alpha-2 code (currently `US` or `CA`).
	Country         string            `json:"country" api:"required"`
	PostalCode      string            `json:"postal_code" api:"required"`
	StreetAddress   string            `json:"street_address" api:"required"`
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

type EnterpriseNewParams struct {
	BillingAddress BillingAddressParam `json:"billing_address,omitzero" api:"required"`
	BillingContact BillingContactParam `json:"billing_contact,omitzero" api:"required"`
	// ISO 3166-1 alpha-2 country code. Currently `US` and `CA` are supported.
	CountryCode     string `json:"country_code" api:"required"`
	DoingBusinessAs string `json:"doing_business_as" api:"required"`
	// US Federal Employer Identification Number (`NN-NNNNNNN`) or Canadian equivalent.
	Fein string `json:"fein" api:"required"`
	// Industry classification.
	//
	// Any of "accounting", "finance", "billing", "collections", "business", "charity",
	// "nonprofit", "communications", "telecom", "customer service", "support",
	// "delivery", "shipping", "logistics", "education", "financial", "banking",
	// "government", "public", "healthcare", "health", "pharmacy", "medical",
	// "insurance", "legal", "law", "notifications", "scheduling", "real estate",
	// "property", "retail", "ecommerce", "sales", "marketing", "software",
	// "technology", "tech", "media", "surveys", "market research", "travel",
	// "hospitality", "hotel".
	Industry                    EnterpriseNewParamsIndustry `json:"industry,omitzero" api:"required"`
	JurisdictionOfIncorporation string                      `json:"jurisdiction_of_incorporation" api:"required"`
	// Legal name of the enterprise.
	LegalName string `json:"legal_name" api:"required"`
	// Approximate headcount range. Used for vetting heuristics; pick the bucket that
	// contains your current employee count.
	//
	// Any of "1-10", "11-50", "51-200", "201-500", "501-2000", "2001-10000", "10001+".
	NumberOfEmployees   EnterpriseNewParamsNumberOfEmployees `json:"number_of_employees,omitzero" api:"required"`
	OrganizationContact OrganizationContactParam             `json:"organization_contact,omitzero" api:"required"`
	// Legal-entity form. Pick the form that matches your incorporation documents:
	//
	//   - `corporation` - C-corp or S-corp.
	//   - `llc` - limited liability company.
	//   - `partnership` - general/limited partnership.
	//   - `nonprofit` - non-profit corporation, charitable trust, or
	//     501(c)(3)/equivalent.
	//   - `other` - anything else (sole proprietorships, government bodies, DBAs, etc.).
	//     You may be asked for additional documents during vetting.
	//
	// Any of "corporation", "llc", "partnership", "nonprofit", "other".
	OrganizationLegalType       EnterpriseNewParamsOrganizationLegalType `json:"organization_legal_type,omitzero" api:"required"`
	OrganizationPhysicalAddress PhysicalAddressParam                     `json:"organization_physical_address,omitzero" api:"required"`
	// Organization category for vetting purposes:
	//
	//   - `commercial` - for-profit business entities (LLC, corp, partnership, sole
	//     proprietorship). Most callers fall here.
	//   - `government` - federal/state/local government bodies.
	//   - `non_profit` - registered 501(c)(3)/equivalent (incl. educational
	//     institutions, charities, religious organisations).
	//
	// Any of "commercial", "government", "non_profit".
	OrganizationType EnterpriseNewParamsOrganizationType `json:"organization_type,omitzero" api:"required"`
	Website          string                              `json:"website" api:"required" format:"uri"`
	// Optional corporate-registration / company-number identifier.
	CorporateRegistrationNumber param.Opt[string] `json:"corporate_registration_number,omitzero"`
	// Optional D-U-N-S Number.
	DunBradstreetNumber param.Opt[string] `json:"dun_bradstreet_number,omitzero"`
	// Optional SIC code for the primary line of business.
	PrimaryBusinessDomainSicCode param.Opt[string] `json:"primary_business_domain_sic_code,omitzero"`
	// Optional professional-license number for regulated industries.
	ProfessionalLicenseNumber param.Opt[string] `json:"professional_license_number,omitzero"`
	// Optional free-form string the caller can attach for their own bookkeeping.
	// Telnyx does not interpret it.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// `enterprise` for an organization registering its own DIRs; `bpo` for a Business
	// Process Outsourcer placing calls on behalf of one or more enterprises.
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

// Industry classification.
type EnterpriseNewParamsIndustry string

const (
	EnterpriseNewParamsIndustryAccounting      EnterpriseNewParamsIndustry = "accounting"
	EnterpriseNewParamsIndustryFinance         EnterpriseNewParamsIndustry = "finance"
	EnterpriseNewParamsIndustryBilling         EnterpriseNewParamsIndustry = "billing"
	EnterpriseNewParamsIndustryCollections     EnterpriseNewParamsIndustry = "collections"
	EnterpriseNewParamsIndustryBusiness        EnterpriseNewParamsIndustry = "business"
	EnterpriseNewParamsIndustryCharity         EnterpriseNewParamsIndustry = "charity"
	EnterpriseNewParamsIndustryNonprofit       EnterpriseNewParamsIndustry = "nonprofit"
	EnterpriseNewParamsIndustryCommunications  EnterpriseNewParamsIndustry = "communications"
	EnterpriseNewParamsIndustryTelecom         EnterpriseNewParamsIndustry = "telecom"
	EnterpriseNewParamsIndustryCustomerService EnterpriseNewParamsIndustry = "customer service"
	EnterpriseNewParamsIndustrySupport         EnterpriseNewParamsIndustry = "support"
	EnterpriseNewParamsIndustryDelivery        EnterpriseNewParamsIndustry = "delivery"
	EnterpriseNewParamsIndustryShipping        EnterpriseNewParamsIndustry = "shipping"
	EnterpriseNewParamsIndustryLogistics       EnterpriseNewParamsIndustry = "logistics"
	EnterpriseNewParamsIndustryEducation       EnterpriseNewParamsIndustry = "education"
	EnterpriseNewParamsIndustryFinancial       EnterpriseNewParamsIndustry = "financial"
	EnterpriseNewParamsIndustryBanking         EnterpriseNewParamsIndustry = "banking"
	EnterpriseNewParamsIndustryGovernment      EnterpriseNewParamsIndustry = "government"
	EnterpriseNewParamsIndustryPublic          EnterpriseNewParamsIndustry = "public"
	EnterpriseNewParamsIndustryHealthcare      EnterpriseNewParamsIndustry = "healthcare"
	EnterpriseNewParamsIndustryHealth          EnterpriseNewParamsIndustry = "health"
	EnterpriseNewParamsIndustryPharmacy        EnterpriseNewParamsIndustry = "pharmacy"
	EnterpriseNewParamsIndustryMedical         EnterpriseNewParamsIndustry = "medical"
	EnterpriseNewParamsIndustryInsurance       EnterpriseNewParamsIndustry = "insurance"
	EnterpriseNewParamsIndustryLegal           EnterpriseNewParamsIndustry = "legal"
	EnterpriseNewParamsIndustryLaw             EnterpriseNewParamsIndustry = "law"
	EnterpriseNewParamsIndustryNotifications   EnterpriseNewParamsIndustry = "notifications"
	EnterpriseNewParamsIndustryScheduling      EnterpriseNewParamsIndustry = "scheduling"
	EnterpriseNewParamsIndustryRealEstate      EnterpriseNewParamsIndustry = "real estate"
	EnterpriseNewParamsIndustryProperty        EnterpriseNewParamsIndustry = "property"
	EnterpriseNewParamsIndustryRetail          EnterpriseNewParamsIndustry = "retail"
	EnterpriseNewParamsIndustryEcommerce       EnterpriseNewParamsIndustry = "ecommerce"
	EnterpriseNewParamsIndustrySales           EnterpriseNewParamsIndustry = "sales"
	EnterpriseNewParamsIndustryMarketing       EnterpriseNewParamsIndustry = "marketing"
	EnterpriseNewParamsIndustrySoftware        EnterpriseNewParamsIndustry = "software"
	EnterpriseNewParamsIndustryTechnology      EnterpriseNewParamsIndustry = "technology"
	EnterpriseNewParamsIndustryTech            EnterpriseNewParamsIndustry = "tech"
	EnterpriseNewParamsIndustryMedia           EnterpriseNewParamsIndustry = "media"
	EnterpriseNewParamsIndustrySurveys         EnterpriseNewParamsIndustry = "surveys"
	EnterpriseNewParamsIndustryMarketResearch  EnterpriseNewParamsIndustry = "market research"
	EnterpriseNewParamsIndustryTravel          EnterpriseNewParamsIndustry = "travel"
	EnterpriseNewParamsIndustryHospitality     EnterpriseNewParamsIndustry = "hospitality"
	EnterpriseNewParamsIndustryHotel           EnterpriseNewParamsIndustry = "hotel"
)

// Approximate headcount range. Used for vetting heuristics; pick the bucket that
// contains your current employee count.
type EnterpriseNewParamsNumberOfEmployees string

const (
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees1_10       EnterpriseNewParamsNumberOfEmployees = "1-10"
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees11_50      EnterpriseNewParamsNumberOfEmployees = "11-50"
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees51_200     EnterpriseNewParamsNumberOfEmployees = "51-200"
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees201_500    EnterpriseNewParamsNumberOfEmployees = "201-500"
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees501_2000   EnterpriseNewParamsNumberOfEmployees = "501-2000"
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees2001_10000 EnterpriseNewParamsNumberOfEmployees = "2001-10000"
	EnterpriseNewParamsNumberOfEmployeesNumberOfEmployees10001Plus  EnterpriseNewParamsNumberOfEmployees = "10001+"
)

// Legal-entity form. Pick the form that matches your incorporation documents:
//
//   - `corporation` - C-corp or S-corp.
//   - `llc` - limited liability company.
//   - `partnership` - general/limited partnership.
//   - `nonprofit` - non-profit corporation, charitable trust, or
//     501(c)(3)/equivalent.
//   - `other` - anything else (sole proprietorships, government bodies, DBAs, etc.).
//     You may be asked for additional documents during vetting.
type EnterpriseNewParamsOrganizationLegalType string

const (
	EnterpriseNewParamsOrganizationLegalTypeCorporation EnterpriseNewParamsOrganizationLegalType = "corporation"
	EnterpriseNewParamsOrganizationLegalTypeLlc         EnterpriseNewParamsOrganizationLegalType = "llc"
	EnterpriseNewParamsOrganizationLegalTypePartnership EnterpriseNewParamsOrganizationLegalType = "partnership"
	EnterpriseNewParamsOrganizationLegalTypeNonprofit   EnterpriseNewParamsOrganizationLegalType = "nonprofit"
	EnterpriseNewParamsOrganizationLegalTypeOther       EnterpriseNewParamsOrganizationLegalType = "other"
)

// Organization category for vetting purposes:
//
//   - `commercial` - for-profit business entities (LLC, corp, partnership, sole
//     proprietorship). Most callers fall here.
//   - `government` - federal/state/local government bodies.
//   - `non_profit` - registered 501(c)(3)/equivalent (incl. educational
//     institutions, charities, religious organisations).
type EnterpriseNewParamsOrganizationType string

const (
	EnterpriseNewParamsOrganizationTypeCommercial EnterpriseNewParamsOrganizationType = "commercial"
	EnterpriseNewParamsOrganizationTypeGovernment EnterpriseNewParamsOrganizationType = "government"
	EnterpriseNewParamsOrganizationTypeNonProfit  EnterpriseNewParamsOrganizationType = "non_profit"
)

// `enterprise` for an organization registering its own DIRs; `bpo` for a Business
// Process Outsourcer placing calls on behalf of one or more enterprises.
type EnterpriseNewParamsRoleType string

const (
	EnterpriseNewParamsRoleTypeEnterprise EnterpriseNewParamsRoleType = "enterprise"
	EnterpriseNewParamsRoleTypeBpo        EnterpriseNewParamsRoleType = "bpo"
)

type EnterpriseUpdateParams struct {
	CorporateRegistrationNumber  param.Opt[string] `json:"corporate_registration_number,omitzero"`
	DunBradstreetNumber          param.Opt[string] `json:"dun_bradstreet_number,omitzero"`
	PrimaryBusinessDomainSicCode param.Opt[string] `json:"primary_business_domain_sic_code,omitzero"`
	ProfessionalLicenseNumber    param.Opt[string] `json:"professional_license_number,omitzero"`
	CustomerReference            param.Opt[string] `json:"customer_reference,omitzero"`
	DoingBusinessAs              param.Opt[string] `json:"doing_business_as,omitzero"`
	Fein                         param.Opt[string] `json:"fein,omitzero"`
	// Updated state/province/country of incorporation. Optional on update.
	JurisdictionOfIncorporation param.Opt[string] `json:"jurisdiction_of_incorporation,omitzero"`
	// Legal name of the enterprise.
	LegalName             param.Opt[string]   `json:"legal_name,omitzero"`
	NumberOfEmployees     param.Opt[string]   `json:"number_of_employees,omitzero"`
	OrganizationLegalType param.Opt[string]   `json:"organization_legal_type,omitzero"`
	Website               param.Opt[string]   `json:"website,omitzero" format:"uri"`
	BillingAddress        BillingAddressParam `json:"billing_address,omitzero"`
	BillingContact        BillingContactParam `json:"billing_contact,omitzero"`
	// Any of "accounting", "finance", "billing", "collections", "business", "charity",
	// "nonprofit", "communications", "telecom", "customer service", "support",
	// "delivery", "shipping", "logistics", "education", "financial", "banking",
	// "government", "public", "healthcare", "health", "pharmacy", "medical",
	// "insurance", "legal", "law", "notifications", "scheduling", "real estate",
	// "property", "retail", "ecommerce", "sales", "marketing", "software",
	// "technology", "tech", "media", "surveys", "market research", "travel",
	// "hospitality", "hotel".
	Industry                    EnterpriseUpdateParamsIndustry `json:"industry,omitzero"`
	OrganizationContact         OrganizationContactParam       `json:"organization_contact,omitzero"`
	OrganizationPhysicalAddress PhysicalAddressParam           `json:"organization_physical_address,omitzero"`
	paramObj
}

func (r EnterpriseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseUpdateParamsIndustry string

const (
	EnterpriseUpdateParamsIndustryAccounting      EnterpriseUpdateParamsIndustry = "accounting"
	EnterpriseUpdateParamsIndustryFinance         EnterpriseUpdateParamsIndustry = "finance"
	EnterpriseUpdateParamsIndustryBilling         EnterpriseUpdateParamsIndustry = "billing"
	EnterpriseUpdateParamsIndustryCollections     EnterpriseUpdateParamsIndustry = "collections"
	EnterpriseUpdateParamsIndustryBusiness        EnterpriseUpdateParamsIndustry = "business"
	EnterpriseUpdateParamsIndustryCharity         EnterpriseUpdateParamsIndustry = "charity"
	EnterpriseUpdateParamsIndustryNonprofit       EnterpriseUpdateParamsIndustry = "nonprofit"
	EnterpriseUpdateParamsIndustryCommunications  EnterpriseUpdateParamsIndustry = "communications"
	EnterpriseUpdateParamsIndustryTelecom         EnterpriseUpdateParamsIndustry = "telecom"
	EnterpriseUpdateParamsIndustryCustomerService EnterpriseUpdateParamsIndustry = "customer service"
	EnterpriseUpdateParamsIndustrySupport         EnterpriseUpdateParamsIndustry = "support"
	EnterpriseUpdateParamsIndustryDelivery        EnterpriseUpdateParamsIndustry = "delivery"
	EnterpriseUpdateParamsIndustryShipping        EnterpriseUpdateParamsIndustry = "shipping"
	EnterpriseUpdateParamsIndustryLogistics       EnterpriseUpdateParamsIndustry = "logistics"
	EnterpriseUpdateParamsIndustryEducation       EnterpriseUpdateParamsIndustry = "education"
	EnterpriseUpdateParamsIndustryFinancial       EnterpriseUpdateParamsIndustry = "financial"
	EnterpriseUpdateParamsIndustryBanking         EnterpriseUpdateParamsIndustry = "banking"
	EnterpriseUpdateParamsIndustryGovernment      EnterpriseUpdateParamsIndustry = "government"
	EnterpriseUpdateParamsIndustryPublic          EnterpriseUpdateParamsIndustry = "public"
	EnterpriseUpdateParamsIndustryHealthcare      EnterpriseUpdateParamsIndustry = "healthcare"
	EnterpriseUpdateParamsIndustryHealth          EnterpriseUpdateParamsIndustry = "health"
	EnterpriseUpdateParamsIndustryPharmacy        EnterpriseUpdateParamsIndustry = "pharmacy"
	EnterpriseUpdateParamsIndustryMedical         EnterpriseUpdateParamsIndustry = "medical"
	EnterpriseUpdateParamsIndustryInsurance       EnterpriseUpdateParamsIndustry = "insurance"
	EnterpriseUpdateParamsIndustryLegal           EnterpriseUpdateParamsIndustry = "legal"
	EnterpriseUpdateParamsIndustryLaw             EnterpriseUpdateParamsIndustry = "law"
	EnterpriseUpdateParamsIndustryNotifications   EnterpriseUpdateParamsIndustry = "notifications"
	EnterpriseUpdateParamsIndustryScheduling      EnterpriseUpdateParamsIndustry = "scheduling"
	EnterpriseUpdateParamsIndustryRealEstate      EnterpriseUpdateParamsIndustry = "real estate"
	EnterpriseUpdateParamsIndustryProperty        EnterpriseUpdateParamsIndustry = "property"
	EnterpriseUpdateParamsIndustryRetail          EnterpriseUpdateParamsIndustry = "retail"
	EnterpriseUpdateParamsIndustryEcommerce       EnterpriseUpdateParamsIndustry = "ecommerce"
	EnterpriseUpdateParamsIndustrySales           EnterpriseUpdateParamsIndustry = "sales"
	EnterpriseUpdateParamsIndustryMarketing       EnterpriseUpdateParamsIndustry = "marketing"
	EnterpriseUpdateParamsIndustrySoftware        EnterpriseUpdateParamsIndustry = "software"
	EnterpriseUpdateParamsIndustryTechnology      EnterpriseUpdateParamsIndustry = "technology"
	EnterpriseUpdateParamsIndustryTech            EnterpriseUpdateParamsIndustry = "tech"
	EnterpriseUpdateParamsIndustryMedia           EnterpriseUpdateParamsIndustry = "media"
	EnterpriseUpdateParamsIndustrySurveys         EnterpriseUpdateParamsIndustry = "surveys"
	EnterpriseUpdateParamsIndustryMarketResearch  EnterpriseUpdateParamsIndustry = "market research"
	EnterpriseUpdateParamsIndustryTravel          EnterpriseUpdateParamsIndustry = "travel"
	EnterpriseUpdateParamsIndustryHospitality     EnterpriseUpdateParamsIndustry = "hospitality"
	EnterpriseUpdateParamsIndustryHotel           EnterpriseUpdateParamsIndustry = "hotel"
)

type EnterpriseListParams struct {
	// Case-insensitive partial match on legal name.
	FilterLegalNameContains param.Opt[string] `query:"filter[legal_name][contains],omitzero" json:"-"`
	// Filter by legal name (partial match).
	LegalName param.Opt[string] `query:"legal_name,omitzero" json:"-"`
	// 1-based page number. Out-of-range values return an empty page with correct meta.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Items per page. Default 10. Maximum 250; values above are clamped to 250.
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
