// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// BrandService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandService] method instead.
type BrandService struct {
	Options         []option.RequestOption
	ExternalVetting BrandExternalVettingService
}

// NewBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBrandService(opts ...option.RequestOption) (r BrandService) {
	r = BrandService{}
	r.Options = opts
	r.ExternalVetting = NewBrandExternalVettingService(opts...)
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
