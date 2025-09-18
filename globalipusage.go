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

// GlobalIPUsageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPUsageService] method instead.
type GlobalIPUsageService struct {
	Options []option.RequestOption
}

// NewGlobalIPUsageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGlobalIPUsageService(opts ...option.RequestOption) (r GlobalIPUsageService) {
	r = GlobalIPUsageService{}
	r.Options = opts
	return
}

// Global IP Usage Metrics
func (r *GlobalIPUsageService) Get(ctx context.Context, query GlobalIPUsageGetParams, opts ...option.RequestOption) (res *GlobalIPUsageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global_ip_usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GlobalIPUsageGetResponse struct {
	Data []GlobalIPUsageGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPUsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPUsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPUsageGetResponseData struct {
	GlobalIP GlobalIPUsageGetResponseDataGlobalIP `json:"global_ip"`
	Received GlobalIPUsageGetResponseDataReceived `json:"received"`
	// The timestamp of the metric.
	Timestamp   time.Time                               `json:"timestamp" format:"date-time"`
	Transmitted GlobalIPUsageGetResponseDataTransmitted `json:"transmitted"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIP    respjson.Field
		Received    respjson.Field
		Timestamp   respjson.Field
		Transmitted respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPUsageGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPUsageGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPUsageGetResponseDataGlobalIP struct {
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
func (r GlobalIPUsageGetResponseDataGlobalIP) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPUsageGetResponseDataGlobalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPUsageGetResponseDataReceived struct {
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
func (r GlobalIPUsageGetResponseDataReceived) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPUsageGetResponseDataReceived) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPUsageGetResponseDataTransmitted struct {
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
func (r GlobalIPUsageGetResponseDataTransmitted) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPUsageGetResponseDataTransmitted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPUsageGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[global_ip_id][in]
	Filter GlobalIPUsageGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPUsageGetParams]'s query parameters as `url.Values`.
func (r GlobalIPUsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[global_ip_id][in]
type GlobalIPUsageGetParamsFilter struct {
	// Filter by exact Global IP ID
	GlobalIPID GlobalIPUsageGetParamsFilterGlobalIPIDUnion `query:"global_ip_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPUsageGetParamsFilter]'s query parameters as
// `url.Values`.
func (r GlobalIPUsageGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GlobalIPUsageGetParamsFilterGlobalIPIDUnion struct {
	OfString                              param.Opt[string]                         `query:",omitzero,inline"`
	OfGlobalIPUsageGetsFilterGlobalIPIDIn *GlobalIPUsageGetParamsFilterGlobalIPIDIn `query:",omitzero,inline"`
	paramUnion
}

func (u *GlobalIPUsageGetParamsFilterGlobalIPIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGlobalIPUsageGetsFilterGlobalIPIDIn) {
		return u.OfGlobalIPUsageGetsFilterGlobalIPIDIn
	}
	return nil
}

// Filtering operations
type GlobalIPUsageGetParamsFilterGlobalIPIDIn struct {
	// Filter by Global IP ID(s) separated by commas
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPUsageGetParamsFilterGlobalIPIDIn]'s query
// parameters as `url.Values`.
func (r GlobalIPUsageGetParamsFilterGlobalIPIDIn) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
