// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// GlobalIPAssignmentHealthService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPAssignmentHealthService] method instead.
type GlobalIPAssignmentHealthService struct {
	Options []option.RequestOption
}

// NewGlobalIPAssignmentHealthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGlobalIPAssignmentHealthService(opts ...option.RequestOption) (r GlobalIPAssignmentHealthService) {
	r = GlobalIPAssignmentHealthService{}
	r.Options = opts
	return
}

// Global IP Assignment Health Check Metrics
func (r *GlobalIPAssignmentHealthService) Get(ctx context.Context, query GlobalIPAssignmentHealthGetParams, opts ...option.RequestOption) (res *GlobalIPAssignmentHealthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global_ip_assignment_health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GlobalIPAssignmentHealthGetResponse struct {
	Data []GlobalIPAssignmentHealthGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentHealthGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentHealthGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentHealthGetResponseData struct {
	GlobalIP           GlobalIPAssignmentHealthGetResponseDataGlobalIP           `json:"global_ip"`
	GlobalIPAssignment GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignment `json:"global_ip_assignment"`
	Health             GlobalIPAssignmentHealthGetResponseDataHealth             `json:"health"`
	// The timestamp of the metric.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIP           respjson.Field
		GlobalIPAssignment respjson.Field
		Health             respjson.Field
		Timestamp          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentHealthGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentHealthGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentHealthGetResponseDataGlobalIP struct {
	// Global IP ID.
	ID string `json:"id" format:"uuid"`
	// The Global IP address.
	IPAddress string `json:"ip_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IPAddress   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentHealthGetResponseDataGlobalIP) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentHealthGetResponseDataGlobalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignment struct {
	// Global IP assignment ID.
	ID            string                                                                 `json:"id" format:"uuid"`
	WireguardPeer GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignmentWireguardPeer `json:"wireguard_peer"`
	// Wireguard peer ID.
	WireguardPeerID string `json:"wireguard_peer_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		WireguardPeer   respjson.Field
		WireguardPeerID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignment) RawJSON() string {
	return r.JSON.raw
}
func (r *GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignmentWireguardPeer struct {
	// The IP address of the interface.
	IPAddress string `json:"ip_address"`
	// A user specified name for the interface.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignmentWireguardPeer) RawJSON() string {
	return r.JSON.raw
}
func (r *GlobalIPAssignmentHealthGetResponseDataGlobalIPAssignmentWireguardPeer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentHealthGetResponseDataHealth struct {
	// The number of failed health checks.
	Fail float64 `json:"fail"`
	// The number of successful health checks.
	Pass float64 `json:"pass"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fail        respjson.Field
		Pass        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentHealthGetResponseDataHealth) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentHealthGetResponseDataHealth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentHealthGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[global_ip_id][in], filter[global_ip_assignment_id][in]
	Filter GlobalIPAssignmentHealthGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentHealthGetParams]'s query parameters as
// `url.Values`.
func (r GlobalIPAssignmentHealthGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[global_ip_id][in], filter[global_ip_assignment_id][in]
type GlobalIPAssignmentHealthGetParamsFilter struct {
	// Filter by exact Global IP Assignment ID
	GlobalIPAssignmentID GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDUnion `query:"global_ip_assignment_id,omitzero" json:"-"`
	// Filter by exact Global IP ID
	GlobalIPID GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDUnion `query:"global_ip_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentHealthGetParamsFilter]'s query parameters
// as `url.Values`.
func (r GlobalIPAssignmentHealthGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDUnion struct {
	OfString                                                   param.Opt[string]                                              `query:",omitzero,inline"`
	OfGlobalIPAssignmentHealthGetsFilterGlobalIPAssignmentIDIn *GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDIn `query:",omitzero,inline"`
	paramUnion
}

func (u *GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGlobalIPAssignmentHealthGetsFilterGlobalIPAssignmentIDIn) {
		return u.OfGlobalIPAssignmentHealthGetsFilterGlobalIPAssignmentIDIn
	}
	return nil
}

// Filtering operations
type GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDIn struct {
	// Filter by Global IP Assignment ID(s) separated by commas
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDIn]'s query
// parameters as `url.Values`.
func (r GlobalIPAssignmentHealthGetParamsFilterGlobalIPAssignmentIDIn) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDUnion struct {
	OfString                                         param.Opt[string]                                    `query:",omitzero,inline"`
	OfGlobalIPAssignmentHealthGetsFilterGlobalIPIDIn *GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDIn `query:",omitzero,inline"`
	paramUnion
}

func (u *GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGlobalIPAssignmentHealthGetsFilterGlobalIPIDIn) {
		return u.OfGlobalIPAssignmentHealthGetsFilterGlobalIPIDIn
	}
	return nil
}

// Filtering operations
type GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDIn struct {
	// Filter by Global IP ID(s) separated by commas
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDIn]'s
// query parameters as `url.Values`.
func (r GlobalIPAssignmentHealthGetParamsFilterGlobalIPIDIn) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
