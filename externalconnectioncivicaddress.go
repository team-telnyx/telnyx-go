// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// External Connections operations
//
// ExternalConnectionCivicAddressService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalConnectionCivicAddressService] method instead.
type ExternalConnectionCivicAddressService struct {
	Options []option.RequestOption
}

// NewExternalConnectionCivicAddressService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalConnectionCivicAddressService(opts ...option.RequestOption) (r ExternalConnectionCivicAddressService) {
	r = ExternalConnectionCivicAddressService{}
	r.Options = opts
	return
}

// Return the details of an existing Civic Address with its Locations inside the
// 'data' attribute of the response.
func (r *ExternalConnectionCivicAddressService) Get(ctx context.Context, addressID string, query ExternalConnectionCivicAddressGetParams, opts ...option.RequestOption) (res *ExternalConnectionCivicAddressGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if addressID == "" {
		err = errors.New("missing required address_id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/civic_addresses/%s", query.ID, addressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the civic addresses and locations from Microsoft Teams.
func (r *ExternalConnectionCivicAddressService) List(ctx context.Context, id string, query ExternalConnectionCivicAddressListParams, opts ...option.RequestOption) (res *ExternalConnectionCivicAddressListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("external_connections/%s/civic_addresses", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CivicAddress struct {
	// Uniquely identifies the resource.
	ID                string `json:"id" format:"uuid"`
	CityOrTown        string `json:"city_or_town"`
	CityOrTownAlias   string `json:"city_or_town_alias"`
	CompanyName       string `json:"company_name"`
	Country           string `json:"country"`
	CountryOrDistrict string `json:"country_or_district"`
	// Identifies what is the default location in the list of locations.
	DefaultLocationID string     `json:"default_location_id" format:"uuid"`
	Description       string     `json:"description"`
	HouseNumber       string     `json:"house_number"`
	HouseNumberSuffix string     `json:"house_number_suffix"`
	Locations         []Location `json:"locations"`
	PostalOrZipCode   string     `json:"postal_or_zip_code"`
	// Identifies the type of the resource.
	RecordType      string `json:"record_type"`
	StateOrProvince string `json:"state_or_province"`
	StreetName      string `json:"street_name"`
	StreetSuffix    string `json:"street_suffix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CityOrTown        respjson.Field
		CityOrTownAlias   respjson.Field
		CompanyName       respjson.Field
		Country           respjson.Field
		CountryOrDistrict respjson.Field
		DefaultLocationID respjson.Field
		Description       respjson.Field
		HouseNumber       respjson.Field
		HouseNumberSuffix respjson.Field
		Locations         respjson.Field
		PostalOrZipCode   respjson.Field
		RecordType        respjson.Field
		StateOrProvince   respjson.Field
		StreetName        respjson.Field
		StreetSuffix      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CivicAddress) RawJSON() string { return r.JSON.raw }
func (r *CivicAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Location struct {
	// Uniquely identifies the resource.
	ID             string `json:"id" format:"uuid"`
	AdditionalInfo string `json:"additional_info"`
	Description    string `json:"description"`
	// Represents whether the location is the default or not.
	IsDefault bool `json:"is_default"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AdditionalInfo respjson.Field
		Description    respjson.Field
		IsDefault      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionCivicAddressGetResponse struct {
	Data CivicAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionCivicAddressGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionCivicAddressGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionCivicAddressListResponse struct {
	Data []CivicAddress `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalConnectionCivicAddressListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalConnectionCivicAddressListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalConnectionCivicAddressGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type ExternalConnectionCivicAddressListParams struct {
	// Filter parameter for civic addresses (deepObject style). Supports filtering by
	// country.
	Filter ExternalConnectionCivicAddressListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionCivicAddressListParams]'s query
// parameters as `url.Values`.
func (r ExternalConnectionCivicAddressListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter parameter for civic addresses (deepObject style). Supports filtering by
// country.
type ExternalConnectionCivicAddressListParamsFilter struct {
	// The country (or countries) to filter addresses by.
	Country []string `query:"country,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExternalConnectionCivicAddressListParamsFilter]'s query
// parameters as `url.Values`.
func (r ExternalConnectionCivicAddressListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
