// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Mobile network operators operations
//
// MobileNetworkOperatorService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMobileNetworkOperatorService] method instead.
type MobileNetworkOperatorService struct {
	Options []option.RequestOption
}

// NewMobileNetworkOperatorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMobileNetworkOperatorService(opts ...option.RequestOption) (r MobileNetworkOperatorService) {
	r = MobileNetworkOperatorService{}
	r.Options = opts
	return
}

// Telnyx has a set of GSM mobile operators partners that are available through our
// mobile network roaming. This resource is entirely managed by Telnyx and may
// change over time. That means that this resource won't allow any write operations
// for it. Still, it's available so it can be used as a support resource that can
// be related to other resources or become a configuration option.
func (r *MobileNetworkOperatorService) List(ctx context.Context, query MobileNetworkOperatorListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[MobileNetworkOperatorListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "mobile_network_operators"
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

// Telnyx has a set of GSM mobile operators partners that are available through our
// mobile network roaming. This resource is entirely managed by Telnyx and may
// change over time. That means that this resource won't allow any write operations
// for it. Still, it's available so it can be used as a support resource that can
// be related to other resources or become a configuration option.
func (r *MobileNetworkOperatorService) ListAutoPaging(ctx context.Context, query MobileNetworkOperatorListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[MobileNetworkOperatorListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type MobileNetworkOperatorListResponse struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// The mobile operator two-character (ISO 3166-1 alpha-2) origin country code.
	CountryCode string `json:"country_code"`
	// MCC stands for Mobile Country Code. It's a three decimal digit that identifies a
	// country.<br/><br/> This code is commonly seen joined with a Mobile Network Code
	// (MNC) in a tuple that allows identifying a carrier known as PLMN (Public Land
	// Mobile Network) code.
	Mcc string `json:"mcc"`
	// MNC stands for Mobile Network Code. It's a two to three decimal digits that
	// identify a network.<br/><br/> This code is commonly seen joined with a Mobile
	// Country Code (MCC) in a tuple that allows identifying a carrier known as PLMN
	// (Public Land Mobile Network) code.
	Mnc string `json:"mnc"`
	// The network operator name.
	Name string `json:"name"`
	// Indicate whether the mobile network operator can be set as preferred in the
	// Network Preferences API.
	NetworkPreferencesEnabled bool `json:"network_preferences_enabled"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// TADIG stands for Transferred Account Data Interchange Group. The TADIG code is a
	// unique identifier for network operators in GSM mobile networks.
	Tadig string `json:"tadig"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CountryCode               respjson.Field
		Mcc                       respjson.Field
		Mnc                       respjson.Field
		Name                      respjson.Field
		NetworkPreferencesEnabled respjson.Field
		RecordType                respjson.Field
		Tadig                     respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobileNetworkOperatorListResponse) RawJSON() string { return r.JSON.raw }
func (r *MobileNetworkOperatorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobileNetworkOperatorListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter for mobile network operators (deepObject style).
	// Originally: filter[name][starts_with], filter[name][contains],
	// filter[name][ends_with], filter[country_code], filter[mcc], filter[mnc],
	// filter[tadig], filter[network_preferences_enabled]
	Filter MobileNetworkOperatorListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobileNetworkOperatorListParams]'s query parameters as
// `url.Values`.
func (r MobileNetworkOperatorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter for mobile network operators (deepObject style).
// Originally: filter[name][starts_with], filter[name][contains],
// filter[name][ends_with], filter[country_code], filter[mcc], filter[mnc],
// filter[tadig], filter[network_preferences_enabled]
type MobileNetworkOperatorListParamsFilter struct {
	// Filter by exact country_code.
	CountryCode param.Opt[string] `query:"country_code,omitzero" json:"-"`
	// Filter by exact MCC.
	Mcc param.Opt[string] `query:"mcc,omitzero" json:"-"`
	// Filter by exact MNC.
	Mnc param.Opt[string] `query:"mnc,omitzero" json:"-"`
	// Filter by network_preferences_enabled.
	NetworkPreferencesEnabled param.Opt[bool] `query:"network_preferences_enabled,omitzero" json:"-"`
	// Filter by exact TADIG.
	Tadig param.Opt[string] `query:"tadig,omitzero" json:"-"`
	// Advanced name filtering operations
	Name MobileNetworkOperatorListParamsFilterName `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobileNetworkOperatorListParamsFilter]'s query parameters
// as `url.Values`.
func (r MobileNetworkOperatorListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Advanced name filtering operations
type MobileNetworkOperatorListParamsFilterName struct {
	// Filter by name containing match.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter by name ending with.
	EndsWith param.Opt[string] `query:"ends_with,omitzero" json:"-"`
	// Filter by name starting with.
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobileNetworkOperatorListParamsFilterName]'s query
// parameters as `url.Values`.
func (r MobileNetworkOperatorListParamsFilterName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
