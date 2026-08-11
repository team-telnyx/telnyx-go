// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Manage the legal business entities that operate RCS agents.
//
// RcBrandService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRcBrandService] method instead.
type RcBrandService struct {
	Options []option.RequestOption
}

// NewRcBrandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRcBrandService(opts ...option.RequestOption) (r RcBrandService) {
	r = RcBrandService{}
	r.Options = opts
	return
}

// Creates an editable RCS brand draft. Creating the draft does not begin external
// review.
func (r *RcBrandService) New(ctx context.Context, body RcBrandNewParams, opts ...option.RequestOption) (res *BrandResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rcs/brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an RCS brand and its current lifecycle status.
func (r *RcBrandService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BrandResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates one or more fields on a brand while its status is `CREATED`. Submitted
// brands cannot be changed.
func (r *RcBrandService) Update(ctx context.Context, id string, body RcBrandUpdateParams, opts ...option.RequestOption) (res *BrandResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/brands/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists RCS brands owned by the authenticated organization.
func (r *RcBrandService) List(ctx context.Context, opts ...option.RequestOption) (res *[]BrandResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rcs/brands"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Starts asynchronous provider provisioning and external review for a brand.
// Repeating this request for an in-progress brand returns its current state
// without creating new work.
func (r *RcBrandService) Submit(ctx context.Context, id string, opts ...option.RequestOption) (res *BrandResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/brands/%s/submit", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type BrandContact struct {
	// Any of "BRAND", "PRIMARY", "OFFICER", "AGENT", "RESPONSIBLE_PARTY", "BILLING",
	// "UNKNOWN".
	ContactType BrandContactContactType `json:"contact_type" api:"required"`
	Email       string                  `json:"email" api:"required" format:"email"`
	FirstName   string                  `json:"first_name" api:"required"`
	LastName    string                  `json:"last_name" api:"required"`
	PhoneNumber string                  `json:"phone_number" api:"required"`
	Title       string                  `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContactType respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		PhoneNumber respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandContact) RawJSON() string { return r.JSON.raw }
func (r *BrandContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BrandContact to a BrandContactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BrandContactParam.Overrides()
func (r BrandContact) ToParam() BrandContactParam {
	return param.Override[BrandContactParam](json.RawMessage(r.RawJSON()))
}

type BrandContactContactType string

const (
	BrandContactContactTypeBrand            BrandContactContactType = "BRAND"
	BrandContactContactTypePrimary          BrandContactContactType = "PRIMARY"
	BrandContactContactTypeOfficer          BrandContactContactType = "OFFICER"
	BrandContactContactTypeAgent            BrandContactContactType = "AGENT"
	BrandContactContactTypeResponsibleParty BrandContactContactType = "RESPONSIBLE_PARTY"
	BrandContactContactTypeBilling          BrandContactContactType = "BILLING"
	BrandContactContactTypeUnknown          BrandContactContactType = "UNKNOWN"
)

// The properties ContactType, Email, FirstName, LastName, PhoneNumber are
// required.
type BrandContactParam struct {
	// Any of "BRAND", "PRIMARY", "OFFICER", "AGENT", "RESPONSIBLE_PARTY", "BILLING",
	// "UNKNOWN".
	ContactType BrandContactContactType `json:"contact_type,omitzero" api:"required"`
	Email       string                  `json:"email" api:"required" format:"email"`
	FirstName   string                  `json:"first_name" api:"required"`
	LastName    string                  `json:"last_name" api:"required"`
	PhoneNumber string                  `json:"phone_number" api:"required"`
	Title       param.Opt[string]       `json:"title,omitzero"`
	paramObj
}

func (r BrandContactParam) MarshalJSON() (data []byte, err error) {
	type shadow BrandContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BrandContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandLegalEntityType string

const (
	BrandLegalEntityTypeLimitedLiabilityCompany BrandLegalEntityType = "LIMITED_LIABILITY_COMPANY"
	BrandLegalEntityTypeSoleProprietorship      BrandLegalEntityType = "SOLE_PROPRIETORSHIP"
	BrandLegalEntityTypePartnership             BrandLegalEntityType = "PARTNERSHIP"
	BrandLegalEntityTypeCorporation             BrandLegalEntityType = "CORPORATION"
	BrandLegalEntityTypeSCorporation            BrandLegalEntityType = "S_CORPORATION"
)

type BrandOrganizationType string

const (
	BrandOrganizationTypePrivateProfit BrandOrganizationType = "PRIVATE_PROFIT"
	BrandOrganizationTypePublicProfit  BrandOrganizationType = "PUBLIC_PROFIT"
	BrandOrganizationTypeNonProfit     BrandOrganizationType = "NON_PROFIT"
	BrandOrganizationTypeGovernment    BrandOrganizationType = "GOVERNMENT"
	BrandOrganizationTypeUnknown       BrandOrganizationType = "UNKNOWN"
)

type BrandResponse struct {
	Addresses        map[string]BrandResponseAddress          `json:"addresses" api:"required"`
	BrandID          string                                   `json:"brand_id" api:"required" format:"uuid"`
	Capabilities     CapabilitiesResponse                     `json:"capabilities" api:"required"`
	Contacts         map[string]BrandContact                  `json:"contacts" api:"required"`
	DisplayName      string                                   `json:"display_name" api:"required"`
	Identifiers      map[string]BrandResponseIdentifiersUnion `json:"identifiers" api:"required"`
	LegalEntityType  string                                   `json:"legal_entity_type" api:"required"`
	LegalName        string                                   `json:"legal_name" api:"required"`
	OrganizationType string                                   `json:"organization_type" api:"required"`
	ProfileID        string                                   `json:"profile_id" api:"required"`
	// Any of "CREATED", "CONFIGURED", "SUBMITTED", "REVIEWING", "VETTING", "VERIFIED",
	// "REJECTED", "FAILED".
	Status     BrandResponseStatus `json:"status" api:"required"`
	WebsiteURL string              `json:"website_url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addresses        respjson.Field
		BrandID          respjson.Field
		Capabilities     respjson.Field
		Contacts         respjson.Field
		DisplayName      respjson.Field
		Identifiers      respjson.Field
		LegalEntityType  respjson.Field
		LegalName        respjson.Field
		OrganizationType respjson.Field
		ProfileID        respjson.Field
		Status           respjson.Field
		WebsiteURL       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandResponse) RawJSON() string { return r.JSON.raw }
func (r *BrandResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandResponseAddress struct {
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// The two-letter ISO 3166-1 country code.
	CountryCode string `json:"country_code" api:"required"`
	Line1       string `json:"line_1" api:"required"`
	PostalCode  string `json:"postal_code" api:"required"`
	Line2       string `json:"line_2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea respjson.Field
		City               respjson.Field
		CountryCode        respjson.Field
		Line1              respjson.Field
		PostalCode         respjson.Field
		Line2              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BrandResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *BrandResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BrandResponseIdentifiersUnion contains all possible properties and values from
// [EinBrandIdentifier], [StockSymbolBrandIdentifier].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BrandResponseIdentifiersUnion struct {
	IdentifierType string `json:"identifier_type"`
	Value          string `json:"value"`
	JSON           struct {
		IdentifierType respjson.Field
		Value          respjson.Field
		raw            string
	} `json:"-"`
}

func (u BrandResponseIdentifiersUnion) AsEinBrandIdentifier() (v EinBrandIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BrandResponseIdentifiersUnion) AsStockSymbolBrandIdentifier() (v StockSymbolBrandIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BrandResponseIdentifiersUnion) RawJSON() string { return u.JSON.raw }

func (r *BrandResponseIdentifiersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BrandResponseStatus string

const (
	BrandResponseStatusCreated    BrandResponseStatus = "CREATED"
	BrandResponseStatusConfigured BrandResponseStatus = "CONFIGURED"
	BrandResponseStatusSubmitted  BrandResponseStatus = "SUBMITTED"
	BrandResponseStatusReviewing  BrandResponseStatus = "REVIEWING"
	BrandResponseStatusVetting    BrandResponseStatus = "VETTING"
	BrandResponseStatusVerified   BrandResponseStatus = "VERIFIED"
	BrandResponseStatusRejected   BrandResponseStatus = "REJECTED"
	BrandResponseStatusFailed     BrandResponseStatus = "FAILED"
)

type EinBrandIdentifier struct {
	IdentifierType constant.Ein `json:"identifier_type" default:"EIN"`
	// Nine digits, optionally formatted as NN-NNNNNNN.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IdentifierType respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EinBrandIdentifier) RawJSON() string { return r.JSON.raw }
func (r *EinBrandIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EinBrandIdentifier to a EinBrandIdentifierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EinBrandIdentifierParam.Overrides()
func (r EinBrandIdentifier) ToParam() EinBrandIdentifierParam {
	return param.Override[EinBrandIdentifierParam](json.RawMessage(r.RawJSON()))
}

// The properties IdentifierType, Value are required.
type EinBrandIdentifierParam struct {
	// Nine digits, optionally formatted as NN-NNNNNNN.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "EIN".
	IdentifierType constant.Ein `json:"identifier_type" default:"EIN"`
	paramObj
}

func (r EinBrandIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow EinBrandIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EinBrandIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StockSymbolBrandIdentifier struct {
	IdentifierType constant.StockSymbol `json:"identifier_type" default:"STOCK_SYMBOL"`
	// A stock symbol using EXCHANGE:SYMBOL.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IdentifierType respjson.Field
		Value          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StockSymbolBrandIdentifier) RawJSON() string { return r.JSON.raw }
func (r *StockSymbolBrandIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StockSymbolBrandIdentifier to a
// StockSymbolBrandIdentifierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StockSymbolBrandIdentifierParam.Overrides()
func (r StockSymbolBrandIdentifier) ToParam() StockSymbolBrandIdentifierParam {
	return param.Override[StockSymbolBrandIdentifierParam](json.RawMessage(r.RawJSON()))
}

// The properties IdentifierType, Value are required.
type StockSymbolBrandIdentifierParam struct {
	// A stock symbol using EXCHANGE:SYMBOL.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "STOCK_SYMBOL".
	IdentifierType constant.StockSymbol `json:"identifier_type" default:"STOCK_SYMBOL"`
	paramObj
}

func (r StockSymbolBrandIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow StockSymbolBrandIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StockSymbolBrandIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcBrandNewParams struct {
	Addresses map[string]RcBrandNewParamsAddresses `json:"addresses,omitzero" api:"required"`
	// Named business contacts. Use the `brand` key for the required BRAND contact.
	Contacts    RcBrandNewParamsContacts `json:"contacts,omitzero" api:"required"`
	DisplayName string                   `json:"display_name" api:"required"`
	// Named business identifiers. Use the `ein` key for the required EIN and
	// `stock_symbol` for a public-profit brand's stock symbol.
	Identifiers RcBrandNewParamsIdentifiers `json:"identifiers,omitzero" api:"required"`
	// Any of "LIMITED_LIABILITY_COMPANY", "SOLE_PROPRIETORSHIP", "PARTNERSHIP",
	// "CORPORATION", "S_CORPORATION".
	LegalEntityType BrandLegalEntityType `json:"legal_entity_type,omitzero" api:"required"`
	LegalName       string               `json:"legal_name" api:"required"`
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT", "UNKNOWN".
	OrganizationType BrandOrganizationType `json:"organization_type,omitzero" api:"required"`
	WebsiteURL       string                `json:"website_url" api:"required" format:"uri"`
	// A Messaging Profile owned by the authenticated organization. Agents inherit this
	// value when they do not provide their own profile.
	ProfileID param.Opt[string] `json:"profile_id,omitzero"`
	paramObj
}

func (r RcBrandNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcBrandNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AdministrativeArea, City, CountryCode, Line1, PostalCode are
// required.
type RcBrandNewParamsAddresses struct {
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// The two-letter ISO 3166-1 country code.
	CountryCode string            `json:"country_code" api:"required"`
	Line1       string            `json:"line_1" api:"required"`
	PostalCode  string            `json:"postal_code" api:"required"`
	Line2       param.Opt[string] `json:"line_2,omitzero"`
	paramObj
}

func (r RcBrandNewParamsAddresses) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandNewParamsAddresses
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcBrandNewParamsAddresses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Named business contacts. Use the `brand` key for the required BRAND contact.
//
// The property Brand is required.
type RcBrandNewParamsContacts struct {
	Brand       RcBrandNewParamsContactsBrand `json:"brand,omitzero" api:"required"`
	ExtraFields map[string]BrandContactParam  `json:"-"`
	paramObj
}

func (r RcBrandNewParamsContacts) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandNewParamsContacts
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *RcBrandNewParamsContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcBrandNewParamsContactsBrand struct {
	ContactType string `json:"contact_type,omitzero"`
	BrandContactParam
}

func (r RcBrandNewParamsContactsBrand) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*RcBrandNewParamsContactsBrand
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Named business identifiers. Use the `ein` key for the required EIN and
// `stock_symbol` for a public-profit brand's stock symbol.
//
// The property Ein is required.
type RcBrandNewParamsIdentifiers struct {
	Ein         EinBrandIdentifierParam                     `json:"ein,omitzero" api:"required"`
	StockSymbol StockSymbolBrandIdentifierParam             `json:"stock_symbol,omitzero"`
	ExtraFields map[string]RcBrandNewParamsIdentifiersUnion `json:"-"`
	paramObj
}

func (r RcBrandNewParamsIdentifiers) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandNewParamsIdentifiers
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *RcBrandNewParamsIdentifiers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RcBrandNewParamsIdentifiersUnion struct {
	OfEinBrandIdentifier         *EinBrandIdentifierParam         `json:",omitzero,inline"`
	OfStockSymbolBrandIdentifier *StockSymbolBrandIdentifierParam `json:",omitzero,inline"`
	paramUnion
}

func (u RcBrandNewParamsIdentifiersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEinBrandIdentifier, u.OfStockSymbolBrandIdentifier)
}
func (u *RcBrandNewParamsIdentifiersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RcBrandNewParamsIdentifiersUnion) asAny() any {
	if !param.IsOmitted(u.OfEinBrandIdentifier) {
		return u.OfEinBrandIdentifier
	} else if !param.IsOmitted(u.OfStockSymbolBrandIdentifier) {
		return u.OfStockSymbolBrandIdentifier
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RcBrandNewParamsIdentifiersUnion) GetIdentifierType() *string {
	if vt := u.OfEinBrandIdentifier; vt != nil {
		return (*string)(&vt.IdentifierType)
	} else if vt := u.OfStockSymbolBrandIdentifier; vt != nil {
		return (*string)(&vt.IdentifierType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RcBrandNewParamsIdentifiersUnion) GetValue() *string {
	if vt := u.OfEinBrandIdentifier; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfStockSymbolBrandIdentifier; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

type RcBrandUpdateParams struct {
	DisplayName param.Opt[string]                       `json:"display_name,omitzero"`
	LegalName   param.Opt[string]                       `json:"legal_name,omitzero"`
	ProfileID   param.Opt[string]                       `json:"profile_id,omitzero"`
	WebsiteURL  param.Opt[string]                       `json:"website_url,omitzero" format:"uri"`
	Addresses   map[string]RcBrandUpdateParamsAddresses `json:"addresses,omitzero"`
	// Named business contacts. Use the `brand` key for the required BRAND contact.
	Contacts RcBrandUpdateParamsContacts `json:"contacts,omitzero"`
	// Named business identifiers. Use the `ein` key for the required EIN and
	// `stock_symbol` for a public-profit brand's stock symbol.
	Identifiers RcBrandUpdateParamsIdentifiers `json:"identifiers,omitzero"`
	// Any of "LIMITED_LIABILITY_COMPANY", "SOLE_PROPRIETORSHIP", "PARTNERSHIP",
	// "CORPORATION", "S_CORPORATION".
	LegalEntityType BrandLegalEntityType `json:"legal_entity_type,omitzero"`
	// Any of "PRIVATE_PROFIT", "PUBLIC_PROFIT", "NON_PROFIT", "GOVERNMENT", "UNKNOWN".
	OrganizationType BrandOrganizationType `json:"organization_type,omitzero"`
	paramObj
}

func (r RcBrandUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcBrandUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AdministrativeArea, City, CountryCode, Line1, PostalCode are
// required.
type RcBrandUpdateParamsAddresses struct {
	AdministrativeArea string `json:"administrative_area" api:"required"`
	City               string `json:"city" api:"required"`
	// The two-letter ISO 3166-1 country code.
	CountryCode string            `json:"country_code" api:"required"`
	Line1       string            `json:"line_1" api:"required"`
	PostalCode  string            `json:"postal_code" api:"required"`
	Line2       param.Opt[string] `json:"line_2,omitzero"`
	paramObj
}

func (r RcBrandUpdateParamsAddresses) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandUpdateParamsAddresses
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcBrandUpdateParamsAddresses) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Named business contacts. Use the `brand` key for the required BRAND contact.
//
// The property Brand is required.
type RcBrandUpdateParamsContacts struct {
	Brand       RcBrandUpdateParamsContactsBrand `json:"brand,omitzero" api:"required"`
	ExtraFields map[string]BrandContactParam     `json:"-"`
	paramObj
}

func (r RcBrandUpdateParamsContacts) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandUpdateParamsContacts
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *RcBrandUpdateParamsContacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcBrandUpdateParamsContactsBrand struct {
	ContactType string `json:"contact_type,omitzero"`
	BrandContactParam
}

func (r RcBrandUpdateParamsContactsBrand) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*RcBrandUpdateParamsContactsBrand
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// Named business identifiers. Use the `ein` key for the required EIN and
// `stock_symbol` for a public-profit brand's stock symbol.
//
// The property Ein is required.
type RcBrandUpdateParamsIdentifiers struct {
	Ein         EinBrandIdentifierParam                        `json:"ein,omitzero" api:"required"`
	StockSymbol StockSymbolBrandIdentifierParam                `json:"stock_symbol,omitzero"`
	ExtraFields map[string]RcBrandUpdateParamsIdentifiersUnion `json:"-"`
	paramObj
}

func (r RcBrandUpdateParamsIdentifiers) MarshalJSON() (data []byte, err error) {
	type shadow RcBrandUpdateParamsIdentifiers
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *RcBrandUpdateParamsIdentifiers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type RcBrandUpdateParamsIdentifiersUnion struct {
	OfEinBrandIdentifier         *EinBrandIdentifierParam         `json:",omitzero,inline"`
	OfStockSymbolBrandIdentifier *StockSymbolBrandIdentifierParam `json:",omitzero,inline"`
	paramUnion
}

func (u RcBrandUpdateParamsIdentifiersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEinBrandIdentifier, u.OfStockSymbolBrandIdentifier)
}
func (u *RcBrandUpdateParamsIdentifiersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *RcBrandUpdateParamsIdentifiersUnion) asAny() any {
	if !param.IsOmitted(u.OfEinBrandIdentifier) {
		return u.OfEinBrandIdentifier
	} else if !param.IsOmitted(u.OfStockSymbolBrandIdentifier) {
		return u.OfStockSymbolBrandIdentifier
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RcBrandUpdateParamsIdentifiersUnion) GetIdentifierType() *string {
	if vt := u.OfEinBrandIdentifier; vt != nil {
		return (*string)(&vt.IdentifierType)
	} else if vt := u.OfStockSymbolBrandIdentifier; vt != nil {
		return (*string)(&vt.IdentifierType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u RcBrandUpdateParamsIdentifiersUnion) GetValue() *string {
	if vt := u.OfEinBrandIdentifier; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfStockSymbolBrandIdentifier; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}
