// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// NetworkCoverageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkCoverageService] method instead.
type NetworkCoverageService struct {
	Options []option.RequestOption
}

// NewNetworkCoverageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkCoverageService(opts ...option.RequestOption) (r NetworkCoverageService) {
	r = NetworkCoverageService{}
	r.Options = opts
	return
}

// List all locations and the interfaces that region supports
func (r *NetworkCoverageService) List(ctx context.Context, query NetworkCoverageListParams, opts ...option.RequestOption) (res *NetworkCoverageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "network_coverage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AvailableService string

const (
	AvailableServiceCloudVpn               AvailableService = "cloud_vpn"
	AvailableServicePrivateWirelessGateway AvailableService = "private_wireless_gateway"
	AvailableServiceVirtualCrossConnect    AvailableService = "virtual_cross_connect"
)

type NetworkCoverageListResponse struct {
	Data []NetworkCoverageListResponseData `json:"data"`
	Meta PaginationMeta                    `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkCoverageListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkCoverageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkCoverageListResponseData struct {
	// List of interface types supported in this region.
	//
	// Any of "cloud_vpn", "private_wireless_gateway", "virtual_cross_connect".
	AvailableServices []string                                `json:"available_services"`
	Location          NetworkCoverageListResponseDataLocation `json:"location"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableServices respjson.Field
		Location          respjson.Field
		RecordType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkCoverageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *NetworkCoverageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkCoverageListResponseDataLocation struct {
	// Location code.
	Code string `json:"code"`
	// Human readable name of location.
	Name string `json:"name"`
	// Point of presence of location.
	Pop string `json:"pop"`
	// Identifies the geographical region of location.
	Region string `json:"region"`
	// Site of location.
	Site string `json:"site"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		Pop         respjson.Field
		Region      respjson.Field
		Site        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkCoverageListResponseDataLocation) RawJSON() string { return r.JSON.raw }
func (r *NetworkCoverageListResponseDataLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkCoverageListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[location.region], filter[location.site], filter[location.pop],
	// filter[location.code]
	Filter NetworkCoverageListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated filters parameter (deepObject style). Originally:
	// filters[available_services][contains]
	Filters NetworkCoverageListParamsFilters `query:"filters,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page NetworkCoverageListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkCoverageListParams]'s query parameters as
// `url.Values`.
func (r NetworkCoverageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[location.region], filter[location.site], filter[location.pop],
// filter[location.code]
type NetworkCoverageListParamsFilter struct {
	// The code of associated location to filter on.
	LocationCode param.Opt[string] `query:"location.code,omitzero" json:"-"`
	// The POP of associated location to filter on.
	LocationPop param.Opt[string] `query:"location.pop,omitzero" json:"-"`
	// The region of associated location to filter on.
	LocationRegion param.Opt[string] `query:"location.region,omitzero" json:"-"`
	// The site of associated location to filter on.
	LocationSite param.Opt[string] `query:"location.site,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkCoverageListParamsFilter]'s query parameters as
// `url.Values`.
func (r NetworkCoverageListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filters parameter (deepObject style). Originally:
// filters[available_services][contains]
type NetworkCoverageListParamsFilters struct {
	// Filter by exact available service match
	AvailableServices NetworkCoverageListParamsFiltersAvailableServicesUnion `query:"available_services,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkCoverageListParamsFilters]'s query parameters as
// `url.Values`.
func (r NetworkCoverageListParamsFilters) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type NetworkCoverageListParamsFiltersAvailableServicesUnion struct {
	// Check if union is this variant with !param.IsOmitted(union.OfAvailableService)
	OfAvailableService                                     param.Opt[AvailableService]                                `query:",omitzero,inline"`
	OfNetworkCoverageListsFiltersAvailableServicesContains *NetworkCoverageListParamsFiltersAvailableServicesContains `query:",omitzero,inline"`
	paramUnion
}

func (u *NetworkCoverageListParamsFiltersAvailableServicesUnion) asAny() any {
	if !param.IsOmitted(u.OfAvailableService) {
		return &u.OfAvailableService
	} else if !param.IsOmitted(u.OfNetworkCoverageListsFiltersAvailableServicesContains) {
		return u.OfNetworkCoverageListsFiltersAvailableServicesContains
	}
	return nil
}

// Available service filtering operations
type NetworkCoverageListParamsFiltersAvailableServicesContains struct {
	// Filter by available services containing the specified service
	//
	// Any of "cloud_vpn", "private_wireless_gateway", "virtual_cross_connect".
	Contains AvailableService `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [NetworkCoverageListParamsFiltersAvailableServicesContains]'s query parameters
// as `url.Values`.
func (r NetworkCoverageListParamsFiltersAvailableServicesContains) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type NetworkCoverageListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkCoverageListParamsPage]'s query parameters as
// `url.Values`.
func (r NetworkCoverageListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
