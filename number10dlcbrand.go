// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Number10dlcBrandService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcBrandService] method instead.
type Number10dlcBrandService struct {
	Options         []option.RequestOption
	ExternalVetting Number10dlcBrandExternalVettingService
}

// NewNumber10dlcBrandService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNumber10dlcBrandService(opts ...option.RequestOption) (r Number10dlcBrandService) {
	r = Number10dlcBrandService{}
	r.Options = opts
	r.ExternalVetting = NewNumber10dlcBrandExternalVettingService(opts...)
	return
}

// This endpoint is used to create a new brand. A brand is an entity created by The
// Campaign Registry (TCR) that represents an organization or a company. It is this
// entity that TCR created campaigns will be associated with. Each brand creation
// will entail an upfront, non-refundable $4 expense.
func (r *Number10dlcBrandService) New(ctx context.Context, body Number10dlcBrandNewParams, opts ...option.RequestOption) (res *TelnyxBrand, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "10dlc/brand"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a brand by `brandId`.
func (r *Number10dlcBrandService) Get(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Number10dlcBrandGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a brand's attributes by `brandId`.
func (r *Number10dlcBrandService) Update(ctx context.Context, brandID string, body Number10dlcBrandUpdateParams, opts ...option.RequestOption) (res *TelnyxBrand, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint is used to list all brands associated with your organization.
func (r *Number10dlcBrandService) List(ctx context.Context, query Number10dlcBrandListParams, opts ...option.RequestOption) (res *pagination.PerPagePaginationV2[Number10dlcBrandListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "10dlc/brand"
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

// This endpoint is used to list all brands associated with your organization.
func (r *Number10dlcBrandService) ListAutoPaging(ctx context.Context, query Number10dlcBrandListParams, opts ...option.RequestOption) *pagination.PerPagePaginationV2AutoPager[Number10dlcBrandListResponse] {
	return pagination.NewPerPagePaginationV2AutoPager(r.List(ctx, query, opts...))
}

// Delete Brand. This endpoint is used to delete a brand. Note the brand cannot be
// deleted if it contains one or more active campaigns, the campaigns need to be
// inactive and at least 3 months old due to billing purposes.
func (r *Number10dlcBrandService) Delete(ctx context.Context, brandID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get feedback about a brand by ID. This endpoint can be used after creating or
// revetting a brand.
//
// Possible values for `.category[].id`:
//
//   - `TAX_ID` - Data mismatch related to tax id and its associated properties.
//   - `STOCK_SYMBOL` - Non public entity registered as a public for profit entity or
//     the stock information mismatch.
//   - `GOVERNMENT_ENTITY` - Non government entity registered as a government entity.
//     Must be a U.S. government entity.
//   - `NONPROFIT` - Not a recognized non-profit entity. No IRS tax-exempt status
//     found.
//   - `OTHERS` - Details of the data misrepresentation if any.
func (r *Number10dlcBrandService) GetFeedback(ctx context.Context, brandID string, opts ...option.RequestOption) (res *Number10dlcBrandGetFeedbackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/feedback/%s", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resend brand 2FA email
func (r *Number10dlcBrandService) Resend2faEmail(ctx context.Context, brandID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/2faEmail", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// This operation allows you to revet the brand. However, revetting is allowed once
// after the successful brand registration and thereafter limited to once every 3
// months.
func (r *Number10dlcBrandService) Revet(ctx context.Context, brandID string, opts ...option.RequestOption) (res *TelnyxBrand, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/revet", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// An enumeration.
type AltBusinessIDType string

const (
	AltBusinessIDTypeNone AltBusinessIDType = "NONE"
	AltBusinessIDTypeDuns AltBusinessIDType = "DUNS"
	AltBusinessIDTypeGiin AltBusinessIDType = "GIIN"
	AltBusinessIDTypeLei  AltBusinessIDType = "LEI"
)

// The verification status of an active brand
type BrandIdentityStatus string

const (
	BrandIdentityStatusVerified       BrandIdentityStatus = "VERIFIED"
	BrandIdentityStatusUnverified     BrandIdentityStatus = "UNVERIFIED"
	BrandIdentityStatusSelfDeclared   BrandIdentityStatus = "SELF_DECLARED"
	BrandIdentityStatusVettedVerified BrandIdentityStatus = "VETTED_VERIFIED"
)

// Entity type behind the brand. This is the form of business establishment.
type EntityType string

const (
	EntityTypePrivateProfit EntityType = "PRIVATE_PROFIT"
	EntityTypePublicProfit  EntityType = "PUBLIC_PROFIT"
	EntityTypeNonProfit     EntityType = "NON_PROFIT"
	EntityTypeGovernment    EntityType = "GOVERNMENT"
)

// (Required for public company) stock exchange.
type StockExchange string

const (
	StockExchangeNone   StockExchange = "NONE"
	StockExchangeNasdaq StockExchange = "NASDAQ"
	StockExchangeNyse   StockExchange = "NYSE"
	StockExchangeAmex   StockExchange = "AMEX"
	StockExchangeAmx    StockExchange = "AMX"
	StockExchangeAsx    StockExchange = "ASX"
	StockExchangeB3     StockExchange = "B3"
	StockExchangeBme    StockExchange = "BME"
	StockExchangeBse    StockExchange = "BSE"
	StockExchangeFra    StockExchange = "FRA"
	StockExchangeIcex   StockExchange = "ICEX"
	StockExchangeJpx    StockExchange = "JPX"
	StockExchangeJse    StockExchange = "JSE"
	StockExchangeKrx    StockExchange = "KRX"
	StockExchangeLon    StockExchange = "LON"
	StockExchangeNse    StockExchange = "NSE"
	StockExchangeOmx    StockExchange = "OMX"
	StockExchangeSehk   StockExchange = "SEHK"
	StockExchangeSse    StockExchange = "SSE"
	StockExchangeSto    StockExchange = "STO"
	StockExchangeSwx    StockExchange = "SWX"
	StockExchangeSzse   StockExchange = "SZSE"
	StockExchangeTsx    StockExchange = "TSX"
	StockExchangeTwse   StockExchange = "TWSE"
	StockExchangeVse    StockExchange = "VSE"
)

// Telnyx-specific extensions to The Campaign Registry's `Brand` type
type TelnyxBrand struct {
	// Brand relationship to the CSP.
	//
	// Any of "BASIC_ACCOUNT", "SMALL_ACCOUNT", "MEDIUM_ACCOUNT", "LARGE_ACCOUNT",
	// "KEY_ACCOUNT".
	BrandRelationship TelnyxBrandBrandRelationship `json:"brandRelationship,required"`
	// ISO2 2 characters country code. Example: US - United States
	Country string `json:"country,required"`
	// Display or marketing name of the brand.
	DisplayName string `json:"displayName,required"`
	// Valid email address of brand support contact.
	Email string `json:"email,required"`
	// Entity type behind the brand. This is the form of business establishment.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT".
	EntityType EntityType `json:"entityType,required"`
	// Vertical or industry segment of the brand.
	Vertical string `json:"vertical,required"`
	// Alternate business identifier such as DUNS, LEI, or GIIN
	AltBusinessID string `json:"altBusinessId"`
	// An enumeration.
	//
	// Any of "NONE", "DUNS", "GIIN", "LEI".
	AltBusinessIDType AltBusinessIDType `json:"altBusinessIdType"`
	// Unique identifier assigned to the brand.
	BrandID string `json:"brandId"`
	// Business contact email.
	//
	// Required if `entityType` is `PUBLIC_PROFIT`.
	BusinessContactEmail string `json:"businessContactEmail"`
	// City name
	City string `json:"city"`
	// (Required for Non-profit/private/public) Legal company name.
	CompanyName string `json:"companyName"`
	// Date and time that the brand was created at.
	CreatedAt string `json:"createdAt"`
	// Unique identifier assigned to the csp by the registry.
	CspID string `json:"cspId"`
	// (Required for Non-profit) Government assigned corporate tax ID. EIN is 9-digits
	// in U.S.
	Ein string `json:"ein"`
	// Failure reasons for brand
	FailureReasons string `json:"failureReasons"`
	// First name of business contact.
	FirstName string `json:"firstName"`
	// The verification status of an active brand
	//
	// Any of "VERIFIED", "UNVERIFIED", "SELF_DECLARED", "VETTED_VERIFIED".
	IdentityStatus BrandIdentityStatus `json:"identityStatus"`
	// IP address of the browser requesting to create brand identity.
	IPAddress string `json:"ipAddress"`
	// Indicates whether this brand is known to be a reseller
	IsReseller bool `json:"isReseller"`
	// Last name of business contact.
	LastName string `json:"lastName"`
	// Valid mobile phone number in e.164 international format.
	MobilePhone string `json:"mobilePhone"`
	// Mock brand for testing purposes
	Mock               bool                          `json:"mock"`
	OptionalAttributes TelnyxBrandOptionalAttributes `json:"optionalAttributes"`
	// Valid phone number in e.164 international format.
	Phone string `json:"phone"`
	// Postal codes. Use 5 digit zipcode for United States
	PostalCode string `json:"postalCode"`
	// Unique identifier Telnyx assigned to the brand - the brandId
	ReferenceID string `json:"referenceId"`
	// State. Must be 2 letters code for United States.
	State string `json:"state"`
	// Status of the brand
	//
	// Any of "OK", "REGISTRATION_PENDING", "REGISTRATION_FAILED".
	Status TelnyxBrandStatus `json:"status"`
	// (Required for public company) stock exchange.
	//
	// Any of "NONE", "NASDAQ", "NYSE", "AMEX", "AMX", "ASX", "B3", "BME", "BSE",
	// "FRA", "ICEX", "JPX", "JSE", "KRX", "LON", "NSE", "OMX", "SEHK", "SSE", "STO",
	// "SWX", "SZSE", "TSX", "TWSE", "VSE".
	StockExchange StockExchange `json:"stockExchange"`
	// (Required for public company) stock symbol.
	StockSymbol string `json:"stockSymbol"`
	// Street number and name.
	Street string `json:"street"`
	// Unique identifier assigned to the brand by the registry.
	TcrBrandID string `json:"tcrBrandId"`
	// Universal EIN of Brand, Read Only.
	UniversalEin string `json:"universalEin"`
	// Date and time that the brand was last updated at.
	UpdatedAt string `json:"updatedAt"`
	// Failover webhook to which brand status updates are sent.
	WebhookFailoverURL string `json:"webhookFailoverURL"`
	// Webhook to which brand status updates are sent.
	WebhookURL string `json:"webhookURL"`
	// Brand website URL.
	Website string `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandRelationship    respjson.Field
		Country              respjson.Field
		DisplayName          respjson.Field
		Email                respjson.Field
		EntityType           respjson.Field
		Vertical             respjson.Field
		AltBusinessID        respjson.Field
		AltBusinessIDType    respjson.Field
		BrandID              respjson.Field
		BusinessContactEmail respjson.Field
		City                 respjson.Field
		CompanyName          respjson.Field
		CreatedAt            respjson.Field
		CspID                respjson.Field
		Ein                  respjson.Field
		FailureReasons       respjson.Field
		FirstName            respjson.Field
		IdentityStatus       respjson.Field
		IPAddress            respjson.Field
		IsReseller           respjson.Field
		LastName             respjson.Field
		MobilePhone          respjson.Field
		Mock                 respjson.Field
		OptionalAttributes   respjson.Field
		Phone                respjson.Field
		PostalCode           respjson.Field
		ReferenceID          respjson.Field
		State                respjson.Field
		Status               respjson.Field
		StockExchange        respjson.Field
		StockSymbol          respjson.Field
		Street               respjson.Field
		TcrBrandID           respjson.Field
		UniversalEin         respjson.Field
		UpdatedAt            respjson.Field
		WebhookFailoverURL   respjson.Field
		WebhookURL           respjson.Field
		Website              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelnyxBrand) RawJSON() string { return r.JSON.raw }
func (r *TelnyxBrand) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brand relationship to the CSP.
type TelnyxBrandBrandRelationship string

const (
	TelnyxBrandBrandRelationshipBasicAccount  TelnyxBrandBrandRelationship = "BASIC_ACCOUNT"
	TelnyxBrandBrandRelationshipSmallAccount  TelnyxBrandBrandRelationship = "SMALL_ACCOUNT"
	TelnyxBrandBrandRelationshipMediumAccount TelnyxBrandBrandRelationship = "MEDIUM_ACCOUNT"
	TelnyxBrandBrandRelationshipLargeAccount  TelnyxBrandBrandRelationship = "LARGE_ACCOUNT"
	TelnyxBrandBrandRelationshipKeyAccount    TelnyxBrandBrandRelationship = "KEY_ACCOUNT"
)

type TelnyxBrandOptionalAttributes struct {
	// The tax exempt status of the brand
	TaxExemptStatus string `json:"taxExemptStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaxExemptStatus respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelnyxBrandOptionalAttributes) RawJSON() string { return r.JSON.raw }
func (r *TelnyxBrandOptionalAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the brand
type TelnyxBrandStatus string

const (
	TelnyxBrandStatusOk                  TelnyxBrandStatus = "OK"
	TelnyxBrandStatusRegistrationPending TelnyxBrandStatus = "REGISTRATION_PENDING"
	TelnyxBrandStatusRegistrationFailed  TelnyxBrandStatus = "REGISTRATION_FAILED"
)

// Vertical or industry segment of the brand or campaign.
type Vertical string

const (
	VerticalRealEstate    Vertical = "REAL_ESTATE"
	VerticalHealthcare    Vertical = "HEALTHCARE"
	VerticalEnergy        Vertical = "ENERGY"
	VerticalEntertainment Vertical = "ENTERTAINMENT"
	VerticalRetail        Vertical = "RETAIL"
	VerticalAgriculture   Vertical = "AGRICULTURE"
	VerticalInsurance     Vertical = "INSURANCE"
	VerticalEducation     Vertical = "EDUCATION"
	VerticalHospitality   Vertical = "HOSPITALITY"
	VerticalFinancial     Vertical = "FINANCIAL"
	VerticalGambling      Vertical = "GAMBLING"
	VerticalConstruction  Vertical = "CONSTRUCTION"
	VerticalNgo           Vertical = "NGO"
	VerticalManufacturing Vertical = "MANUFACTURING"
	VerticalGovernment    Vertical = "GOVERNMENT"
	VerticalTechnology    Vertical = "TECHNOLOGY"
	VerticalCommunication Vertical = "COMMUNICATION"
)

// Telnyx-specific extensions to The Campaign Registry's `Brand` type
type Number10dlcBrandGetResponse struct {
	// Number of campaigns associated with the brand
	AssignedCampaignsCount float64 `json:"assignedCampaignsCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssignedCampaignsCount respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
	TelnyxBrand
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandGetResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandListResponse struct {
	// Number of campaigns associated with the brand
	AssignedCampaingsCount int64 `json:"assignedCampaingsCount"`
	// Unique identifier assigned to the brand.
	BrandID string `json:"brandId"`
	// (Required for Non-profit/private/public) Legal company name.
	CompanyName string `json:"companyName"`
	// Date and time that the brand was created at.
	CreatedAt string `json:"createdAt"`
	// Display or marketing name of the brand.
	DisplayName string `json:"displayName"`
	// Valid email address of brand support contact.
	Email string `json:"email"`
	// Entity type behind the brand. This is the form of business establishment.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT".
	EntityType EntityType `json:"entityType"`
	// Failure reasons for brand
	FailureReasons string `json:"failureReasons"`
	// The verification status of an active brand
	//
	// Any of "VERIFIED", "UNVERIFIED", "SELF_DECLARED", "VETTED_VERIFIED".
	IdentityStatus BrandIdentityStatus `json:"identityStatus"`
	// Status of the brand
	//
	// Any of "OK", "REGISTRATION_PENDING", "REGISTRATION_FAILED".
	Status Number10dlcBrandListResponseStatus `json:"status"`
	// Unique identifier assigned to the brand by the registry.
	TcrBrandID string `json:"tcrBrandId"`
	// Date and time that the brand was last updated at.
	UpdatedAt string `json:"updatedAt"`
	// Brand website URL.
	Website string `json:"website"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssignedCampaingsCount respjson.Field
		BrandID                respjson.Field
		CompanyName            respjson.Field
		CreatedAt              respjson.Field
		DisplayName            respjson.Field
		Email                  respjson.Field
		EntityType             respjson.Field
		FailureReasons         respjson.Field
		IdentityStatus         respjson.Field
		Status                 respjson.Field
		TcrBrandID             respjson.Field
		UpdatedAt              respjson.Field
		Website                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandListResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the brand
type Number10dlcBrandListResponseStatus string

const (
	Number10dlcBrandListResponseStatusOk                  Number10dlcBrandListResponseStatus = "OK"
	Number10dlcBrandListResponseStatusRegistrationPending Number10dlcBrandListResponseStatus = "REGISTRATION_PENDING"
	Number10dlcBrandListResponseStatusRegistrationFailed  Number10dlcBrandListResponseStatus = "REGISTRATION_FAILED"
)

type Number10dlcBrandGetFeedbackResponse struct {
	// ID of the brand being queried about
	BrandID string `json:"brandId,required"`
	// A list of reasons why brand creation/revetting didn't go as planned
	Category []Number10dlcBrandGetFeedbackResponseCategory `json:"category,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID     respjson.Field
		Category    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandGetFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandGetFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandGetFeedbackResponseCategory struct {
	// One of `TAX_ID`, `STOCK_SYMBOL`, `GOVERNMENT_ENTITY`, `NONPROFIT`, and `OTHERS`
	ID string `json:"id,required"`
	// Long-form description of the feedback with additional information
	Description string `json:"description,required"`
	// Human-readable version of the `id` field
	DisplayName string `json:"displayName,required"`
	// List of relevant fields in the originally-submitted brand json
	Fields []string `json:"fields,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Fields      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandGetFeedbackResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandGetFeedbackResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandNewParams struct {
	// ISO2 2 characters country code. Example: US - United States
	Country string `json:"country,required"`
	// Display name, marketing name, or DBA name of the brand.
	DisplayName string `json:"displayName,required"`
	// Valid email address of brand support contact.
	Email string `json:"email,required"`
	// Entity type behind the brand. This is the form of business establishment.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT".
	EntityType EntityType `json:"entityType,omitzero,required"`
	// Vertical or industry segment of the brand or campaign.
	//
	// Any of "REAL_ESTATE", "HEALTHCARE", "ENERGY", "ENTERTAINMENT", "RETAIL",
	// "AGRICULTURE", "INSURANCE", "EDUCATION", "HOSPITALITY", "FINANCIAL", "GAMBLING",
	// "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical Vertical `json:"vertical,omitzero,required"`
	// Business contact email.
	//
	// Required if `entityType` is `PUBLIC_PROFIT`. Otherwise, it is recommended to
	// either omit this field or set it to `null`.
	BusinessContactEmail param.Opt[string] `json:"businessContactEmail,omitzero"`
	// City name
	City param.Opt[string] `json:"city,omitzero"`
	// (Required for Non-profit/private/public) Legal company name.
	CompanyName param.Opt[string] `json:"companyName,omitzero"`
	// (Required for Non-profit) Government assigned corporate tax ID. EIN is 9-digits
	// in U.S.
	Ein param.Opt[string] `json:"ein,omitzero"`
	// First name of business contact.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// IP address of the browser requesting to create brand identity.
	IPAddress  param.Opt[string] `json:"ipAddress,omitzero"`
	IsReseller param.Opt[bool]   `json:"isReseller,omitzero"`
	// Last name of business contact.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Valid mobile phone number in e.164 international format.
	MobilePhone param.Opt[string] `json:"mobilePhone,omitzero"`
	// Mock brand for testing purposes. Defaults to false.
	Mock param.Opt[bool] `json:"mock,omitzero"`
	// Valid phone number in e.164 international format.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Postal codes. Use 5 digit zipcode for United States
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State. Must be 2 letters code for United States.
	State param.Opt[string] `json:"state,omitzero"`
	// (Required for public company) stock symbol.
	StockSymbol param.Opt[string] `json:"stockSymbol,omitzero"`
	// Street number and name.
	Street param.Opt[string] `json:"street,omitzero"`
	// Webhook failover URL for brand status updates.
	WebhookFailoverURL param.Opt[string] `json:"webhookFailoverURL,omitzero"`
	// Webhook URL for brand status updates.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	// Brand website URL.
	Website param.Opt[string] `json:"website,omitzero"`
	// (Required for public company) stock exchange.
	//
	// Any of "NONE", "NASDAQ", "NYSE", "AMEX", "AMX", "ASX", "B3", "BME", "BSE",
	// "FRA", "ICEX", "JPX", "JSE", "KRX", "LON", "NSE", "OMX", "SEHK", "SSE", "STO",
	// "SWX", "SZSE", "TSX", "TWSE", "VSE".
	StockExchange StockExchange `json:"stockExchange,omitzero"`
	paramObj
}

func (r Number10dlcBrandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandUpdateParams struct {
	// ISO2 2 characters country code. Example: US - United States
	Country string `json:"country,required"`
	// Display or marketing name of the brand.
	DisplayName string `json:"displayName,required"`
	// Valid email address of brand support contact.
	Email string `json:"email,required"`
	// Entity type behind the brand. This is the form of business establishment.
	//
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT".
	EntityType EntityType `json:"entityType,omitzero,required"`
	// Vertical or industry segment of the brand or campaign.
	//
	// Any of "REAL_ESTATE", "HEALTHCARE", "ENERGY", "ENTERTAINMENT", "RETAIL",
	// "AGRICULTURE", "INSURANCE", "EDUCATION", "HOSPITALITY", "FINANCIAL", "GAMBLING",
	// "CONSTRUCTION", "NGO", "MANUFACTURING", "GOVERNMENT", "TECHNOLOGY",
	// "COMMUNICATION".
	Vertical Vertical `json:"vertical,omitzero,required"`
	// Alternate business identifier such as DUNS, LEI, or GIIN
	AltBusinessID param.Opt[string] `json:"altBusiness_id,omitzero"`
	// Business contact email.
	//
	// Required if `entityType` will be changed to `PUBLIC_PROFIT`. Otherwise, it is
	// recommended to either omit this field or set it to `null`.
	BusinessContactEmail param.Opt[string] `json:"businessContactEmail,omitzero"`
	// City name
	City param.Opt[string] `json:"city,omitzero"`
	// (Required for Non-profit/private/public) Legal company name.
	CompanyName param.Opt[string] `json:"companyName,omitzero"`
	// (Required for Non-profit) Government assigned corporate tax ID. EIN is 9-digits
	// in U.S.
	Ein param.Opt[string] `json:"ein,omitzero"`
	// First name of business contact.
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// IP address of the browser requesting to create brand identity.
	IPAddress  param.Opt[string] `json:"ipAddress,omitzero"`
	IsReseller param.Opt[bool]   `json:"isReseller,omitzero"`
	// Last name of business contact.
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Valid phone number in e.164 international format.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// Postal codes. Use 5 digit zipcode for United States
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// State. Must be 2 letters code for United States.
	State param.Opt[string] `json:"state,omitzero"`
	// (Required for public company) stock symbol.
	StockSymbol param.Opt[string] `json:"stockSymbol,omitzero"`
	// Street number and name.
	Street param.Opt[string] `json:"street,omitzero"`
	// Webhook failover URL for brand status updates.
	WebhookFailoverURL param.Opt[string] `json:"webhookFailoverURL,omitzero"`
	// Webhook URL for brand status updates.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	// Brand website URL.
	Website param.Opt[string] `json:"website,omitzero"`
	// An enumeration.
	//
	// Any of "NONE", "DUNS", "GIIN", "LEI".
	AltBusinessIDType AltBusinessIDType `json:"altBusinessIdType,omitzero"`
	// The verification status of an active brand
	//
	// Any of "VERIFIED", "UNVERIFIED", "SELF_DECLARED", "VETTED_VERIFIED".
	IdentityStatus BrandIdentityStatus `json:"identityStatus,omitzero"`
	// (Required for public company) stock exchange.
	//
	// Any of "NONE", "NASDAQ", "NYSE", "AMEX", "AMX", "ASX", "B3", "BME", "BSE",
	// "FRA", "ICEX", "JPX", "JSE", "KRX", "LON", "NSE", "OMX", "SEHK", "SSE", "STO",
	// "SWX", "SZSE", "TSX", "TWSE", "VSE".
	StockExchange StockExchange `json:"stockExchange,omitzero"`
	paramObj
}

func (r Number10dlcBrandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandListParams struct {
	// Filter results by the Telnyx Brand id
	BrandID     param.Opt[string] `query:"brandId,omitzero" json:"-"`
	Country     param.Opt[string] `query:"country,omitzero" json:"-"`
	DisplayName param.Opt[string] `query:"displayName,omitzero" json:"-"`
	EntityType  param.Opt[string] `query:"entityType,omitzero" json:"-"`
	Page        param.Opt[int64]  `query:"page,omitzero" json:"-"`
	// number of records per page. maximum of 500
	RecordsPerPage param.Opt[int64]  `query:"recordsPerPage,omitzero" json:"-"`
	State          param.Opt[string] `query:"state,omitzero" json:"-"`
	// Filter results by the TCR Brand id
	TcrBrandID param.Opt[string] `query:"tcrBrandId,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// createdAt in descending order.
	//
	// Any of "assignedCampaignsCount", "-assignedCampaignsCount", "brandId",
	// "-brandId", "createdAt", "-createdAt", "displayName", "-displayName",
	// "identityStatus", "-identityStatus", "status", "-status", "tcrBrandId",
	// "-tcrBrandId".
	Sort Number10dlcBrandListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [Number10dlcBrandListParams]'s query parameters as
// `url.Values`.
func (r Number10dlcBrandListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// createdAt in descending order.
type Number10dlcBrandListParamsSort string

const (
	Number10dlcBrandListParamsSortAssignedCampaignsCount     Number10dlcBrandListParamsSort = "assignedCampaignsCount"
	Number10dlcBrandListParamsSortAssignedCampaignsCountDesc Number10dlcBrandListParamsSort = "-assignedCampaignsCount"
	Number10dlcBrandListParamsSortBrandID                    Number10dlcBrandListParamsSort = "brandId"
	Number10dlcBrandListParamsSortBrandIDDesc                Number10dlcBrandListParamsSort = "-brandId"
	Number10dlcBrandListParamsSortCreatedAt                  Number10dlcBrandListParamsSort = "createdAt"
	Number10dlcBrandListParamsSortCreatedAtDesc              Number10dlcBrandListParamsSort = "-createdAt"
	Number10dlcBrandListParamsSortDisplayName                Number10dlcBrandListParamsSort = "displayName"
	Number10dlcBrandListParamsSortDisplayNameDesc            Number10dlcBrandListParamsSort = "-displayName"
	Number10dlcBrandListParamsSortIdentityStatus             Number10dlcBrandListParamsSort = "identityStatus"
	Number10dlcBrandListParamsSortIdentityStatusDesc         Number10dlcBrandListParamsSort = "-identityStatus"
	Number10dlcBrandListParamsSortStatus                     Number10dlcBrandListParamsSort = "status"
	Number10dlcBrandListParamsSortStatusDesc                 Number10dlcBrandListParamsSort = "-status"
	Number10dlcBrandListParamsSortTcrBrandID                 Number10dlcBrandListParamsSort = "tcrBrandId"
	Number10dlcBrandListParamsSortTcrBrandIDDesc             Number10dlcBrandListParamsSort = "-tcrBrandId"
)
