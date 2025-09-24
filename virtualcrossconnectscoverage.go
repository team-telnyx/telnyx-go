// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// VirtualCrossConnectsCoverageService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVirtualCrossConnectsCoverageService] method instead.
type VirtualCrossConnectsCoverageService struct {
	Options []option.RequestOption
}

// NewVirtualCrossConnectsCoverageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewVirtualCrossConnectsCoverageService(opts ...option.RequestOption) (r VirtualCrossConnectsCoverageService) {
	r = VirtualCrossConnectsCoverageService{}
	r.Options = opts
	return
}

// List Virtual Cross Connects Cloud Coverage.<br /><br />This endpoint shows which
// cloud regions are available for the `location_code` your Virtual Cross Connect
// will be provisioned in.
func (r *VirtualCrossConnectsCoverageService) List(ctx context.Context, query VirtualCrossConnectsCoverageListParams, opts ...option.RequestOption) (res *VirtualCrossConnectsCoverageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "virtual_cross_connects_coverage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type VirtualCrossConnectsCoverageListResponse struct {
	Data []VirtualCrossConnectsCoverageListResponseData `json:"data"`
	Meta PaginationMeta                                 `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectsCoverageListResponse) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectsCoverageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectsCoverageListResponseData struct {
	// The available throughput in Megabits per Second (Mbps) for your Virtual Cross
	// Connect.
	AvailableBandwidth []float64 `json:"available_bandwidth"`
	// The Virtual Private Cloud with which you would like to establish a cross
	// connect.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `json:"cloud_provider"`
	// The region where your Virtual Private Cloud hosts are located. Should be
	// identical to how the cloud provider names region, i.e. us-east-1 for AWS but
	// Frankfurt for Azure
	CloudProviderRegion string                                               `json:"cloud_provider_region"`
	Location            VirtualCrossConnectsCoverageListResponseDataLocation `json:"location"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableBandwidth  respjson.Field
		CloudProvider       respjson.Field
		CloudProviderRegion respjson.Field
		Location            respjson.Field
		RecordType          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VirtualCrossConnectsCoverageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectsCoverageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectsCoverageListResponseDataLocation struct {
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
func (r VirtualCrossConnectsCoverageListResponseDataLocation) RawJSON() string { return r.JSON.raw }
func (r *VirtualCrossConnectsCoverageListResponseDataLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VirtualCrossConnectsCoverageListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[cloud_provider], filter[cloud_provider_region], filter[location.region],
	// filter[location.site], filter[location.pop], filter[location.code]
	Filter VirtualCrossConnectsCoverageListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated filters parameter (deepObject style). Originally:
	// filters[available_bandwidth][contains]
	Filters VirtualCrossConnectsCoverageListParamsFilters `query:"filters,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page VirtualCrossConnectsCoverageListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectsCoverageListParams]'s query parameters
// as `url.Values`.
func (r VirtualCrossConnectsCoverageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[cloud_provider], filter[cloud_provider_region], filter[location.region],
// filter[location.site], filter[location.pop], filter[location.code]
type VirtualCrossConnectsCoverageListParamsFilter struct {
	// The region of specific cloud provider.
	CloudProviderRegion param.Opt[string] `query:"cloud_provider_region,omitzero" json:"-"`
	// The code of associated location to filter on.
	LocationCode param.Opt[string] `query:"location.code,omitzero" json:"-"`
	// The POP of associated location to filter on.
	LocationPop param.Opt[string] `query:"location.pop,omitzero" json:"-"`
	// The region of associated location to filter on.
	LocationRegion param.Opt[string] `query:"location.region,omitzero" json:"-"`
	// The site of associated location to filter on.
	LocationSite param.Opt[string] `query:"location.site,omitzero" json:"-"`
	// The Virtual Private Cloud provider.
	//
	// Any of "aws", "azure", "gce".
	CloudProvider string `query:"cloud_provider,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectsCoverageListParamsFilter]'s query
// parameters as `url.Values`.
func (r VirtualCrossConnectsCoverageListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filters parameter (deepObject style). Originally:
// filters[available_bandwidth][contains]
type VirtualCrossConnectsCoverageListParamsFilters struct {
	// Filter by exact available bandwidth match
	AvailableBandwidth VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthUnion `query:"available_bandwidth,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectsCoverageListParamsFilters]'s query
// parameters as `url.Values`.
func (r VirtualCrossConnectsCoverageListParamsFilters) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthUnion struct {
	OfInt                                                                param.Opt[int64]                                                         `query:",omitzero,inline"`
	OfVirtualCrossConnectsCoverageListsFiltersAvailableBandwidthContains *VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthContains `query:",omitzero,inline"`
	paramUnion
}

func (u *VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthUnion) asAny() any {
	if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfVirtualCrossConnectsCoverageListsFiltersAvailableBandwidthContains) {
		return u.OfVirtualCrossConnectsCoverageListsFiltersAvailableBandwidthContains
	}
	return nil
}

// Available bandwidth filtering operations
type VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthContains struct {
	// Filter by available bandwidth containing the specified value
	Contains param.Opt[int64] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthContains]'s
// query parameters as `url.Values`.
func (r VirtualCrossConnectsCoverageListParamsFiltersAvailableBandwidthContains) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type VirtualCrossConnectsCoverageListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VirtualCrossConnectsCoverageListParamsPage]'s query
// parameters as `url.Values`.
func (r VirtualCrossConnectsCoverageListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
