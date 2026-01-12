// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// GlobalIPAssignmentsUsageService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPAssignmentsUsageService] method instead.
type GlobalIPAssignmentsUsageService struct {
	Options []option.RequestOption
}

// NewGlobalIPAssignmentsUsageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewGlobalIPAssignmentsUsageService(opts ...option.RequestOption) (r GlobalIPAssignmentsUsageService) {
	r = GlobalIPAssignmentsUsageService{}
	r.Options = opts
	return
}

// Global IP Assignment Usage Metrics
func (r *GlobalIPAssignmentsUsageService) Get(ctx context.Context, query GlobalIPAssignmentsUsageGetParams, opts ...option.RequestOption) (res *GlobalIPAssignmentsUsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "global_ip_assignments_usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GlobalIPAssignmentsUsageGetResponse struct {
	Data []GlobalIPAssignmentsUsageGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentsUsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentsUsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetResponseData struct {
	GlobalIP           GlobalIPAssignmentsUsageGetResponseDataGlobalIP           `json:"global_ip"`
	GlobalIPAssignment GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignment `json:"global_ip_assignment"`
	Received           GlobalIPAssignmentsUsageGetResponseDataReceived           `json:"received"`
	// The timestamp of the metric.
	Timestamp   time.Time                                          `json:"timestamp" format:"date-time"`
	Transmitted GlobalIPAssignmentsUsageGetResponseDataTransmitted `json:"transmitted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIP           respjson.Field
		GlobalIPAssignment respjson.Field
		Received           respjson.Field
		Timestamp          respjson.Field
		Transmitted        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentsUsageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentsUsageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetResponseDataGlobalIP struct {
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
func (r GlobalIPAssignmentsUsageGetResponseDataGlobalIP) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentsUsageGetResponseDataGlobalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignment struct {
	// Global IP assignment ID.
	ID            string                                                                 `json:"id" format:"uuid"`
	WireguardPeer GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignmentWireguardPeer `json:"wireguard_peer"`
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
func (r GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignment) RawJSON() string {
	return r.JSON.raw
}
func (r *GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignmentWireguardPeer struct {
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
func (r GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignmentWireguardPeer) RawJSON() string {
	return r.JSON.raw
}
func (r *GlobalIPAssignmentsUsageGetResponseDataGlobalIPAssignmentWireguardPeer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetResponseDataReceived struct {
	// The amount of data received.
	Amount float64 `json:"amount"`
	// The unit of the amount of data received.
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentsUsageGetResponseDataReceived) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentsUsageGetResponseDataReceived) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetResponseDataTransmitted struct {
	// The amount of data transmitted.
	Amount float64 `json:"amount"`
	// The unit of the amount of data transmitted.
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPAssignmentsUsageGetResponseDataTransmitted) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPAssignmentsUsageGetResponseDataTransmitted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPAssignmentsUsageGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[global_ip_assignment_id][in], filter[global_ip_id][in]
	Filter GlobalIPAssignmentsUsageGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentsUsageGetParams]'s query parameters as
// `url.Values`.
func (r GlobalIPAssignmentsUsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[global_ip_assignment_id][in], filter[global_ip_id][in]
type GlobalIPAssignmentsUsageGetParamsFilter struct {
	// Filter by exact Global IP Assignment ID
	GlobalIPAssignmentID GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDUnion `query:"global_ip_assignment_id,omitzero" json:"-"`
	// Filter by exact Global IP ID
	GlobalIPID GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDUnion `query:"global_ip_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentsUsageGetParamsFilter]'s query parameters
// as `url.Values`.
func (r GlobalIPAssignmentsUsageGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDUnion struct {
	OfString                                                   param.Opt[string]                                              `query:",omitzero,inline"`
	OfGlobalIPAssignmentsUsageGetsFilterGlobalIPAssignmentIDIn *GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDIn `query:",omitzero,inline"`
	paramUnion
}

func (u *GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGlobalIPAssignmentsUsageGetsFilterGlobalIPAssignmentIDIn) {
		return u.OfGlobalIPAssignmentsUsageGetsFilterGlobalIPAssignmentIDIn
	}
	return nil
}

// Filtering operations
type GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDIn struct {
	// Filter by Global IP Assignment ID(s) separated by commas
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDIn]'s query
// parameters as `url.Values`.
func (r GlobalIPAssignmentsUsageGetParamsFilterGlobalIPAssignmentIDIn) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDUnion struct {
	OfString                                         param.Opt[string]                                    `query:",omitzero,inline"`
	OfGlobalIPAssignmentsUsageGetsFilterGlobalIPIDIn *GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDIn `query:",omitzero,inline"`
	paramUnion
}

func (u *GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGlobalIPAssignmentsUsageGetsFilterGlobalIPIDIn) {
		return u.OfGlobalIPAssignmentsUsageGetsFilterGlobalIPIDIn
	}
	return nil
}

// Filtering operations
type GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDIn struct {
	// Filter by Global IP ID(s) separated by commas
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDIn]'s
// query parameters as `url.Values`.
func (r GlobalIPAssignmentsUsageGetParamsFilterGlobalIPIDIn) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
