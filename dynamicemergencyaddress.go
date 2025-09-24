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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// DynamicEmergencyAddressService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDynamicEmergencyAddressService] method instead.
type DynamicEmergencyAddressService struct {
	Options []option.RequestOption
}

// NewDynamicEmergencyAddressService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDynamicEmergencyAddressService(opts ...option.RequestOption) (r DynamicEmergencyAddressService) {
	r = DynamicEmergencyAddressService{}
	r.Options = opts
	return
}

// Creates a dynamic emergency address.
func (r *DynamicEmergencyAddressService) New(ctx context.Context, body DynamicEmergencyAddressNewParams, opts ...option.RequestOption) (res *DynamicEmergencyAddressNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dynamic_emergency_addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the dynamic emergency address based on the ID provided
func (r *DynamicEmergencyAddressService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DynamicEmergencyAddressGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("dynamic_emergency_addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the dynamic emergency addresses according to filters
func (r *DynamicEmergencyAddressService) List(ctx context.Context, query DynamicEmergencyAddressListParams, opts ...option.RequestOption) (res *DynamicEmergencyAddressListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "dynamic_emergency_addresses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the dynamic emergency address based on the ID provided
func (r *DynamicEmergencyAddressService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DynamicEmergencyAddressDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("dynamic_emergency_addresses/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DynamicEmergencyAddress struct {
	AdministrativeArea string `json:"administrative_area,required"`
	// Any of "US", "CA", "PR".
	CountryCode DynamicEmergencyAddressCountryCode `json:"country_code,required"`
	HouseNumber string                             `json:"house_number,required"`
	Locality    string                             `json:"locality,required"`
	PostalCode  string                             `json:"postal_code,required"`
	StreetName  string                             `json:"street_name,required"`
	ID          string                             `json:"id"`
	// ISO 8601 formatted date of when the resource was created
	CreatedAt       string `json:"created_at"`
	ExtendedAddress string `json:"extended_address"`
	HouseSuffix     string `json:"house_suffix"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Unique location reference string to be used in SIP INVITE from / p-asserted
	// headers.
	SipGeolocationID string `json:"sip_geolocation_id"`
	// Status of dynamic emergency address
	//
	// Any of "pending", "activated", "rejected".
	Status                DynamicEmergencyAddressStatus `json:"status"`
	StreetPostDirectional string                        `json:"street_post_directional"`
	StreetPreDirectional  string                        `json:"street_pre_directional"`
	StreetSuffix          string                        `json:"street_suffix"`
	// ISO 8601 formatted date of when the resource was last updated
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdministrativeArea    respjson.Field
		CountryCode           respjson.Field
		HouseNumber           respjson.Field
		Locality              respjson.Field
		PostalCode            respjson.Field
		StreetName            respjson.Field
		ID                    respjson.Field
		CreatedAt             respjson.Field
		ExtendedAddress       respjson.Field
		HouseSuffix           respjson.Field
		RecordType            respjson.Field
		SipGeolocationID      respjson.Field
		Status                respjson.Field
		StreetPostDirectional respjson.Field
		StreetPreDirectional  respjson.Field
		StreetSuffix          respjson.Field
		UpdatedAt             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyAddress) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DynamicEmergencyAddress to a DynamicEmergencyAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DynamicEmergencyAddressParam.Overrides()
func (r DynamicEmergencyAddress) ToParam() DynamicEmergencyAddressParam {
	return param.Override[DynamicEmergencyAddressParam](json.RawMessage(r.RawJSON()))
}

type DynamicEmergencyAddressCountryCode string

const (
	DynamicEmergencyAddressCountryCodeUs DynamicEmergencyAddressCountryCode = "US"
	DynamicEmergencyAddressCountryCodeCa DynamicEmergencyAddressCountryCode = "CA"
	DynamicEmergencyAddressCountryCodePr DynamicEmergencyAddressCountryCode = "PR"
)

// Status of dynamic emergency address
type DynamicEmergencyAddressStatus string

const (
	DynamicEmergencyAddressStatusPending   DynamicEmergencyAddressStatus = "pending"
	DynamicEmergencyAddressStatusActivated DynamicEmergencyAddressStatus = "activated"
	DynamicEmergencyAddressStatusRejected  DynamicEmergencyAddressStatus = "rejected"
)

// The properties AdministrativeArea, CountryCode, HouseNumber, Locality,
// PostalCode, StreetName are required.
type DynamicEmergencyAddressParam struct {
	AdministrativeArea string `json:"administrative_area,required"`
	// Any of "US", "CA", "PR".
	CountryCode DynamicEmergencyAddressCountryCode `json:"country_code,omitzero,required"`
	HouseNumber string                             `json:"house_number,required"`
	Locality    string                             `json:"locality,required"`
	PostalCode  string                             `json:"postal_code,required"`
	StreetName  string                             `json:"street_name,required"`
	ID          param.Opt[string]                  `json:"id,omitzero"`
	// ISO 8601 formatted date of when the resource was created
	CreatedAt       param.Opt[string] `json:"created_at,omitzero"`
	ExtendedAddress param.Opt[string] `json:"extended_address,omitzero"`
	HouseSuffix     param.Opt[string] `json:"house_suffix,omitzero"`
	// Identifies the type of the resource.
	RecordType param.Opt[string] `json:"record_type,omitzero"`
	// Unique location reference string to be used in SIP INVITE from / p-asserted
	// headers.
	SipGeolocationID      param.Opt[string] `json:"sip_geolocation_id,omitzero"`
	StreetPostDirectional param.Opt[string] `json:"street_post_directional,omitzero"`
	StreetPreDirectional  param.Opt[string] `json:"street_pre_directional,omitzero"`
	StreetSuffix          param.Opt[string] `json:"street_suffix,omitzero"`
	// ISO 8601 formatted date of when the resource was last updated
	UpdatedAt param.Opt[string] `json:"updated_at,omitzero"`
	// Status of dynamic emergency address
	//
	// Any of "pending", "activated", "rejected".
	Status DynamicEmergencyAddressStatus `json:"status,omitzero"`
	paramObj
}

func (r DynamicEmergencyAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow DynamicEmergencyAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DynamicEmergencyAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyAddressNewResponse struct {
	Data DynamicEmergencyAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyAddressNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyAddressNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyAddressGetResponse struct {
	Data DynamicEmergencyAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyAddressGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyAddressGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyAddressListResponse struct {
	Data []DynamicEmergencyAddress `json:"data"`
	Meta shared.Metadata           `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyAddressListResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyAddressListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyAddressDeleteResponse struct {
	Data DynamicEmergencyAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicEmergencyAddressDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DynamicEmergencyAddressDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicEmergencyAddressNewParams struct {
	DynamicEmergencyAddress DynamicEmergencyAddressParam
	paramObj
}

func (r DynamicEmergencyAddressNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DynamicEmergencyAddress)
}
func (r *DynamicEmergencyAddressNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DynamicEmergencyAddress)
}

type DynamicEmergencyAddressListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[status],
	// filter[country_code]
	Filter DynamicEmergencyAddressListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page DynamicEmergencyAddressListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DynamicEmergencyAddressListParams]'s query parameters as
// `url.Values`.
func (r DynamicEmergencyAddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[status],
// filter[country_code]
type DynamicEmergencyAddressListParamsFilter struct {
	// Filter by country code.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter by status.
	//
	// Any of "pending", "activated", "rejected".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DynamicEmergencyAddressListParamsFilter]'s query parameters
// as `url.Values`.
func (r DynamicEmergencyAddressListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type DynamicEmergencyAddressListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DynamicEmergencyAddressListParamsPage]'s query parameters
// as `url.Values`.
func (r DynamicEmergencyAddressListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
