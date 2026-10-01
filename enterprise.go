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
	// Verify ownership of a DIR's authorizer email. A short code is emailed and
	// confirmed; the email must be verified before references can be submitted.
	VerifyEmail EnterpriseVerifyEmailService
}

// NewEnterpriseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnterpriseService(opts ...option.RequestOption) (r EnterpriseService) {
	r = EnterpriseService{}
	r.Options = opts
	r.Reputation = NewEnterpriseReputationService(opts...)
	r.Dir = NewEnterpriseDirService(opts...)
	r.VerifyEmail = NewEnterpriseVerifyEmailService(opts...)
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
//
// For an approved BPO enterprise (`role_type` `bpo`), changing any identity field
// (legal name, DBA, website, FEIN, industry, number of employees, physical
// address, organization contact, D-U-N-S number, legal type, SIC code, corporate
// registration number, professional license number, or jurisdiction of
// incorporation) resets `bpo_verification_status` to `pending` for re-approval and
// sets every DIR authorization for that BPO to `rejected`. After re-approval, link
// it again with a newly signed LOA (a new `loa_document_id`); resending the old
// one keeps the authorization `rejected`. Re-sending an unchanged value does not
// reset anything.
//
// If Number Reputation is enabled on the enterprise, `legal_name`,
// `doing_business_as`, `website`, `fein`, `industry`, `number_of_employees`,
// `organization_physical_address`, `organization_contact`, and
// `dun_bradstreet_number` cannot be changed: the request is rejected with `400`.
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

// Branded Calling must be activated on each enterprise. Activation is idempotent:
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
//   - `400` - the account has no available credit. Add funds and retry.
//   - `400` - the enterprise is not in the United States. Branded Calling is
//     currently available only to US enterprises.
//   - `403` - Branded Calling Terms of Service not accepted.
//   - `404` - enterprise does not exist or does not belong to your account.
//
// **Pricing:** Activation itself is free, but the account must have available
// credit. Branded Calling fees are charged per DIR and per branded call. See
// https://telnyx.com/pricing/branded-calling for current pricing.
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

type BillingContact struct {
	// The email address of the person Telnyx should contact about billing for this
	// account.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the person Telnyx should contact about billing for this
	// account.
	FirstName string `json:"first_name" api:"required"`
	// The last name of the person Telnyx should contact about billing for this
	// account.
	LastName string `json:"last_name" api:"required"`
	// The phone number of the billing contact, in E.164 format, for example
	// +12125551234.
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
	// The email address of the person Telnyx should contact about billing for this
	// account.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the person Telnyx should contact about billing for this
	// account.
	FirstName string `json:"first_name" api:"required"`
	// The last name of the person Telnyx should contact about billing for this
	// account.
	LastName string `json:"last_name" api:"required"`
	// The phone number of the billing contact, in E.164 format, for example
	// +12125551234.
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
	ID             string          `json:"id" format:"uuid"`
	BillingAddress PhysicalAddress `json:"billing_address"`
	BillingContact BillingContact  `json:"billing_contact"`
	// Reason Telnyx rejected the BPO (Business Process Outsourcer) verification, when
	// `bpo_verification_status` is `rejected`; `null` otherwise.
	BpoVerificationRejectionReason string `json:"bpo_verification_rejection_reason" api:"nullable"`
	// Whether Telnyx has approved this BPO (Business Process Outsourcer) account. Only
	// set for accounts created with `role_type` `bpo`; `null` for normal enterprises.
	// A BPO enterprise must be `approved` before a DIR can be linked to it through
	// `bpo_authorizations`.
	//
	// Any of "pending", "approved", "rejected".
	BpoVerificationStatus EnterprisePublicBpoVerificationStatus `json:"bpo_verification_status" api:"nullable"`
	// True once Branded Calling has been activated on this enterprise (see
	// `POST /enterprises/{id}/branded_calling`).
	BrandedCallingEnabled bool `json:"branded_calling_enabled"`
	// The official number your company received when it was legally registered or
	// incorporated (for example from your state or national business registry). It is
	// on your certificate of incorporation.
	CorporateRegistrationNumber string    `json:"corporate_registration_number" api:"nullable"`
	CountryCode                 string    `json:"country_code"`
	CreatedAt                   time.Time `json:"created_at" format:"date-time"`
	// Your own label for this account. Enter any reference that helps you find it in
	// your records. Telnyx does not use it during vetting.
	CustomerReference string `json:"customer_reference"`
	// The trade name your business operates under if it is different from your legal
	// name, also called a Doing Business As (DBA) name. Leave blank if you only use
	// your legal name.
	DoingBusinessAs string `json:"doing_business_as"`
	// Your optional 9-digit D-U-N-S Number issued by Dun & Bradstreet, a unique
	// identifier for your business. Leave blank if you do not have one.
	DunBradstreetNumber string `json:"dun_bradstreet_number" api:"nullable"`
	// US Federal Employer Identification Number (`NN-NNNNNNN`) or Canadian equivalent.
	Fein string `json:"fein"`
	// The industry your business operates in. Choose the closest match from the list;
	// if your value is not accepted, pick the nearest category.
	Industry string `json:"industry"`
	// The state, province, or country where your business was legally incorporated,
	// for example Delaware.
	JurisdictionOfIncorporation string `json:"jurisdiction_of_incorporation"`
	// Your business's full registered legal name, exactly as it appears on your
	// incorporation or tax documents, 3 to 64 characters.
	LegalName string `json:"legal_name"`
	// Approximate headcount range. Used for vetting heuristics; pick the bucket that
	// contains your current employee count.
	NumberOfEmployees string `json:"number_of_employees"`
	// True once Phone Number Reputation has been enabled on this enterprise (see
	// `POST /enterprises/{id}/reputation`).
	NumberReputationEnabled bool                `json:"number_reputation_enabled"`
	OrganizationContact     OrganizationContact `json:"organization_contact"`
	// Legal-entity form. Pick the form that matches your incorporation documents:
	//
	//   - `corporation` - C-corp or S-corp.
	//   - `llc` - limited liability company.
	//   - `partnership` - general/limited partnership.
	//   - `nonprofit` - non-profit corporation, charitable trust, or
	//     501(c)(3)/equivalent.
	//   - `other` - anything else (sole proprietorships, government bodies, DBAs, etc.).
	//     You may be asked for additional documents during vetting.
	OrganizationLegalType       string          `json:"organization_legal_type"`
	OrganizationPhysicalAddress PhysicalAddress `json:"organization_physical_address"`
	OrganizationType            string          `json:"organization_type"`
	// The 4-digit Standard Industrial Classification code for your main line of
	// business, which tells us what industry you operate in. Look it up in the SIC
	// code directory if you are unsure.
	PrimaryBusinessDomainSicCode string `json:"primary_business_domain_sic_code" api:"nullable"`
	// If your business operates under a professional license (for example legal,
	// medical, or financial services), enter the license number issued by the
	// licensing authority. Leave blank if it does not apply.
	ProfessionalLicenseNumber string `json:"professional_license_number" api:"nullable"`
	// Any of "enterprise", "bpo".
	RoleType  EnterprisePublicRoleType `json:"role_type"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	// Your business's public website address, including https://. Leave blank if your
	// business has no website.
	Website string `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                             respjson.Field
		BillingAddress                 respjson.Field
		BillingContact                 respjson.Field
		BpoVerificationRejectionReason respjson.Field
		BpoVerificationStatus          respjson.Field
		BrandedCallingEnabled          respjson.Field
		CorporateRegistrationNumber    respjson.Field
		CountryCode                    respjson.Field
		CreatedAt                      respjson.Field
		CustomerReference              respjson.Field
		DoingBusinessAs                respjson.Field
		DunBradstreetNumber            respjson.Field
		Fein                           respjson.Field
		Industry                       respjson.Field
		JurisdictionOfIncorporation    respjson.Field
		LegalName                      respjson.Field
		NumberOfEmployees              respjson.Field
		NumberReputationEnabled        respjson.Field
		OrganizationContact            respjson.Field
		OrganizationLegalType          respjson.Field
		OrganizationPhysicalAddress    respjson.Field
		OrganizationType               respjson.Field
		PrimaryBusinessDomainSicCode   respjson.Field
		ProfessionalLicenseNumber      respjson.Field
		RoleType                       respjson.Field
		UpdatedAt                      respjson.Field
		Website                        respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterprisePublic) RawJSON() string { return r.JSON.raw }
func (r *EnterprisePublic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether Telnyx has approved this BPO (Business Process Outsourcer) account. Only
// set for accounts created with `role_type` `bpo`; `null` for normal enterprises.
// A BPO enterprise must be `approved` before a DIR can be linked to it through
// `bpo_authorizations`.
type EnterprisePublicBpoVerificationStatus string

const (
	EnterprisePublicBpoVerificationStatusPending  EnterprisePublicBpoVerificationStatus = "pending"
	EnterprisePublicBpoVerificationStatusApproved EnterprisePublicBpoVerificationStatus = "approved"
	EnterprisePublicBpoVerificationStatusRejected EnterprisePublicBpoVerificationStatus = "rejected"
)

type EnterprisePublicRoleType string

const (
	EnterprisePublicRoleTypeEnterprise EnterprisePublicRoleType = "enterprise"
	EnterprisePublicRoleTypeBpo        EnterprisePublicRoleType = "bpo"
)

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
	// The email address of the main person Telnyx should contact about this account.
	// For a call center (BPO) account this is the email you will verify later, so use
	// a mailbox you can access.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the main person Telnyx should contact about this account.
	FirstName string `json:"first_name" api:"required"`
	// The job title of the main person Telnyx should contact about this account.
	JobTitle string `json:"job_title" api:"required"`
	// The last name of the main person Telnyx should contact about this account.
	LastName string `json:"last_name" api:"required"`
	// The phone number of the main contact, in E.164 format, for example +12125551234.
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
	// The email address of the main person Telnyx should contact about this account.
	// For a call center (BPO) account this is the email you will verify later, so use
	// a mailbox you can access.
	Email string `json:"email" api:"required" format:"email"`
	// The first name of the main person Telnyx should contact about this account.
	FirstName string `json:"first_name" api:"required"`
	// The job title of the main person Telnyx should contact about this account.
	JobTitle string `json:"job_title" api:"required"`
	// The last name of the main person Telnyx should contact about this account.
	LastName string `json:"last_name" api:"required"`
	// The phone number of the main contact, in E.164 format, for example +12125551234.
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
	// The city of your registered business address.
	City string `json:"city" api:"required"`
	// ISO 3166-1 alpha-2 code (currently `US` or `CA`).
	Country string `json:"country" api:"required"`
	// The postal or ZIP code of your registered business address.
	PostalCode string `json:"postal_code" api:"required"`
	// The street address of your registered business, including the building number
	// and street name.
	StreetAddress string `json:"street_address" api:"required"`
	// An optional second address line, such as a suite, unit, or floor. Leave blank if
	// it does not apply.
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
	// The city of your registered business address.
	City string `json:"city" api:"required"`
	// ISO 3166-1 alpha-2 code (currently `US` or `CA`).
	Country string `json:"country" api:"required"`
	// The postal or ZIP code of your registered business address.
	PostalCode string `json:"postal_code" api:"required"`
	// The street address of your registered business, including the building number
	// and street name.
	StreetAddress string `json:"street_address" api:"required"`
	// An optional second address line, such as a suite, unit, or floor. Leave blank if
	// it does not apply.
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
	BillingAddress PhysicalAddressParam `json:"billing_address,omitzero" api:"required"`
	BillingContact BillingContactParam  `json:"billing_contact,omitzero" api:"required"`
	// ISO 3166-1 alpha-2 country code. Currently `US` and `CA` are supported.
	CountryCode string `json:"country_code" api:"required"`
	// The trade name your business operates under if it is different from your legal
	// name, also called a Doing Business As (DBA) name. Leave blank if you only use
	// your legal name.
	DoingBusinessAs string `json:"doing_business_as" api:"required"`
	// US Federal Employer Identification Number (`NN-NNNNNNN`) or Canadian equivalent.
	Fein string `json:"fein" api:"required"`
	// The industry your business operates in. Choose the closest match from the list;
	// if your value is not accepted, pick the nearest category.
	//
	// Any of "accounting", "finance", "billing", "collections", "business", "charity",
	// "nonprofit", "communications", "telecom", "customer service", "support",
	// "delivery", "shipping", "logistics", "education", "financial", "banking",
	// "government", "public", "healthcare", "health", "pharmacy", "medical",
	// "insurance", "legal", "law", "notifications", "scheduling", "real estate",
	// "property", "retail", "ecommerce", "sales", "marketing", "software",
	// "technology", "tech", "media", "surveys", "market research", "travel",
	// "hospitality", "hotel".
	Industry EnterpriseNewParamsIndustry `json:"industry,omitzero" api:"required"`
	// The state, province, or country where your business was legally incorporated,
	// for example Delaware.
	JurisdictionOfIncorporation string `json:"jurisdiction_of_incorporation" api:"required"`
	// Your business's full registered legal name, exactly as it appears on your
	// incorporation or tax documents, 3 to 64 characters.
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
	// Your business's public website address, including https://. Leave blank if your
	// business has no website.
	Website string `json:"website" api:"required" format:"uri"`
	// The official number your company received when it was legally registered or
	// incorporated (for example from your state or national business registry). It is
	// on your certificate of incorporation.
	CorporateRegistrationNumber param.Opt[string] `json:"corporate_registration_number,omitzero"`
	// Your optional 9-digit D-U-N-S Number issued by Dun & Bradstreet, a unique
	// identifier for your business. Leave blank if you do not have one.
	DunBradstreetNumber param.Opt[string] `json:"dun_bradstreet_number,omitzero"`
	// The 4-digit Standard Industrial Classification code for your main line of
	// business, which tells us what industry you operate in. Look it up in the SIC
	// code directory if you are unsure.
	PrimaryBusinessDomainSicCode param.Opt[string] `json:"primary_business_domain_sic_code,omitzero"`
	// If your business operates under a professional license (for example legal,
	// medical, or financial services), enter the license number issued by the
	// licensing authority. Leave blank if it does not apply.
	ProfessionalLicenseNumber param.Opt[string] `json:"professional_license_number,omitzero"`
	// Your own label for this account. Enter any reference that helps you find it in
	// your records. Telnyx does not use it during vetting.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// `enterprise` for an organization registering its own DIRs (the default, and the
	// right choice when the calls display your own brand). `bpo` for a Business
	// Process Outsourcer: a call center that places calls on behalf of other
	// enterprises and displays their brand. A `bpo` enterprise describes the call
	// center itself and cannot own a DIR. Each client the call center calls for gets
	// its own `enterprise` in the same account, with the client's DIR under it; that
	// DIR is then linked to the `bpo` enterprise through `bpo_authorizations`. Fixed
	// at creation.
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

// The industry your business operates in. Choose the closest match from the list;
// if your value is not accepted, pick the nearest category.
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

// `enterprise` for an organization registering its own DIRs (the default, and the
// right choice when the calls display your own brand). `bpo` for a Business
// Process Outsourcer: a call center that places calls on behalf of other
// enterprises and displays their brand. A `bpo` enterprise describes the call
// center itself and cannot own a DIR. Each client the call center calls for gets
// its own `enterprise` in the same account, with the client's DIR under it; that
// DIR is then linked to the `bpo` enterprise through `bpo_authorizations`. Fixed
// at creation.
type EnterpriseNewParamsRoleType string

const (
	EnterpriseNewParamsRoleTypeEnterprise EnterpriseNewParamsRoleType = "enterprise"
	EnterpriseNewParamsRoleTypeBpo        EnterpriseNewParamsRoleType = "bpo"
)

type EnterpriseUpdateParams struct {
	// The official number your company received when it was legally registered or
	// incorporated (for example from your state or national business registry). It is
	// on your certificate of incorporation.
	CorporateRegistrationNumber param.Opt[string] `json:"corporate_registration_number,omitzero"`
	// Your optional 9-digit D-U-N-S Number issued by Dun & Bradstreet, a unique
	// identifier for your business. Leave blank if you do not have one.
	DunBradstreetNumber param.Opt[string] `json:"dun_bradstreet_number,omitzero"`
	// The 4-digit Standard Industrial Classification code for your main line of
	// business, which tells us what industry you operate in. Look it up in the SIC
	// code directory if you are unsure.
	PrimaryBusinessDomainSicCode param.Opt[string] `json:"primary_business_domain_sic_code,omitzero"`
	// If your business operates under a professional license (for example legal,
	// medical, or financial services), enter the license number issued by the
	// licensing authority. Leave blank if it does not apply.
	ProfessionalLicenseNumber param.Opt[string] `json:"professional_license_number,omitzero"`
	// Your own label for this account. Enter any reference that helps you find it in
	// your records. Telnyx does not use it during vetting.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// The trade name your business operates under if it is different from your legal
	// name, also called a Doing Business As (DBA) name. Leave blank if you only use
	// your legal name.
	DoingBusinessAs param.Opt[string] `json:"doing_business_as,omitzero"`
	// US Federal Employer Identification Number (`NN-NNNNNNN`) or Canadian equivalent.
	Fein param.Opt[string] `json:"fein,omitzero"`
	// The state, province, or country where your business was legally incorporated,
	// for example Delaware.
	JurisdictionOfIncorporation param.Opt[string] `json:"jurisdiction_of_incorporation,omitzero"`
	// Your business's full registered legal name, exactly as it appears on your
	// incorporation or tax documents, 3 to 64 characters.
	LegalName param.Opt[string] `json:"legal_name,omitzero"`
	// Approximate headcount range. Used for vetting heuristics; pick the bucket that
	// contains your current employee count.
	NumberOfEmployees param.Opt[string] `json:"number_of_employees,omitzero"`
	// Legal-entity form. Pick the form that matches your incorporation documents:
	//
	//   - `corporation` - C-corp or S-corp.
	//   - `llc` - limited liability company.
	//   - `partnership` - general/limited partnership.
	//   - `nonprofit` - non-profit corporation, charitable trust, or
	//     501(c)(3)/equivalent.
	//   - `other` - anything else (sole proprietorships, government bodies, DBAs, etc.).
	//     You may be asked for additional documents during vetting.
	OrganizationLegalType param.Opt[string] `json:"organization_legal_type,omitzero"`
	// Your business's public website address, including https://. Leave blank if your
	// business has no website.
	Website        param.Opt[string]    `json:"website,omitzero" format:"uri"`
	BillingAddress PhysicalAddressParam `json:"billing_address,omitzero"`
	BillingContact BillingContactParam  `json:"billing_contact,omitzero"`
	// The industry your business operates in. Choose the closest match from the list;
	// if your value is not accepted, pick the nearest category.
	//
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

// The industry your business operates in. Choose the closest match from the list;
// if your value is not accepted, pick the nearest category.
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
	// Only return enterprises of this type: `bpo` for call-center (BPO) enterprises,
	// `enterprise` for normal enterprises. Omit to return both.
	//
	// Any of "enterprise", "bpo".
	FilterRoleType EnterpriseListParamsFilterRoleType `query:"filter[role_type],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseListParams]'s query parameters as `url.Values`.
func (r EnterpriseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only return enterprises of this type: `bpo` for call-center (BPO) enterprises,
// `enterprise` for normal enterprises. Omit to return both.
type EnterpriseListParamsFilterRoleType string

const (
	EnterpriseListParamsFilterRoleTypeEnterprise EnterpriseListParamsFilterRoleType = "enterprise"
	EnterpriseListParamsFilterRoleTypeBpo        EnterpriseListParamsFilterRoleType = "bpo"
)
