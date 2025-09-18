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

// GlobalIPLatencyService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalIPLatencyService] method instead.
type GlobalIPLatencyService struct {
	Options []option.RequestOption
}

// NewGlobalIPLatencyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGlobalIPLatencyService(opts ...option.RequestOption) (r GlobalIPLatencyService) {
	r = GlobalIPLatencyService{}
	r.Options = opts
	return
}

// Global IP Latency Metrics
func (r *GlobalIPLatencyService) Get(ctx context.Context, query GlobalIPLatencyGetParams, opts ...option.RequestOption) (res *GlobalIPLatencyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global_ip_latency"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GlobalIPLatencyGetResponse struct {
	Data []GlobalIPLatencyGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPLatencyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseData struct {
	GlobalIP          GlobalIPLatencyGetResponseDataGlobalIP          `json:"global_ip"`
	MeanLatency       GlobalIPLatencyGetResponseDataMeanLatency       `json:"mean_latency"`
	PercentileLatency GlobalIPLatencyGetResponseDataPercentileLatency `json:"percentile_latency"`
	ProberLocation    GlobalIPLatencyGetResponseDataProberLocation    `json:"prober_location"`
	// The timestamp of the metric.
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GlobalIP          respjson.Field
		MeanLatency       respjson.Field
		PercentileLatency respjson.Field
		ProberLocation    respjson.Field
		Timestamp         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPLatencyGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataGlobalIP struct {
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
func (r GlobalIPLatencyGetResponseDataGlobalIP) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataGlobalIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataMeanLatency struct {
	// The average latency.
	Amount float64 `json:"amount"`
	// The unit of the average latency.
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
func (r GlobalIPLatencyGetResponseDataMeanLatency) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataMeanLatency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency struct {
	Number0   GlobalIPLatencyGetResponseDataPercentileLatency0   `json:"0"`
	Number100 GlobalIPLatencyGetResponseDataPercentileLatency100 `json:"100"`
	Number25  GlobalIPLatencyGetResponseDataPercentileLatency25  `json:"25"`
	Number50  GlobalIPLatencyGetResponseDataPercentileLatency50  `json:"50"`
	Number75  GlobalIPLatencyGetResponseDataPercentileLatency75  `json:"75"`
	Number90  GlobalIPLatencyGetResponseDataPercentileLatency90  `json:"90"`
	Number99  GlobalIPLatencyGetResponseDataPercentileLatency99  `json:"99"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number0     respjson.Field
		Number100   respjson.Field
		Number25    respjson.Field
		Number50    respjson.Field
		Number75    respjson.Field
		Number90    respjson.Field
		Number99    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPLatencyGetResponseDataPercentileLatency) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency0 struct {
	// The minimum latency.
	Amount float64 `json:"amount"`
	// The unit of the minimum latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency0) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency0) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency100 struct {
	// The maximum latency.
	Amount float64 `json:"amount"`
	// The unit of the maximum latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency100) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency100) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency25 struct {
	// The 25th percentile latency.
	Amount float64 `json:"amount"`
	// The unit of the 25th percentile latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency25) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency25) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency50 struct {
	// The 50th percentile latency.
	Amount float64 `json:"amount"`
	// The unit of the 50th percentile latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency50) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency50) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency75 struct {
	// The 75th percentile latency.
	Amount float64 `json:"amount"`
	// The unit of the 75th percentile latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency75) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency75) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency90 struct {
	// The 90th percentile latency.
	Amount float64 `json:"amount"`
	// The unit of the 90th percentile latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency90) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency90) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataPercentileLatency99 struct {
	// The 99th percentile latency.
	Amount float64 `json:"amount"`
	// The unit of the 99th percentile latency.
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
func (r GlobalIPLatencyGetResponseDataPercentileLatency99) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataPercentileLatency99) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetResponseDataProberLocation struct {
	// Location ID.
	ID string `json:"id" format:"uuid"`
	// Latitude.
	Lat float64 `json:"lat"`
	// Longitude.
	Lon float64 `json:"lon"`
	// Location name.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Lat         respjson.Field
		Lon         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalIPLatencyGetResponseDataProberLocation) RawJSON() string { return r.JSON.raw }
func (r *GlobalIPLatencyGetResponseDataProberLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalIPLatencyGetParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[global_ip_id][in]
	Filter GlobalIPLatencyGetParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPLatencyGetParams]'s query parameters as
// `url.Values`.
func (r GlobalIPLatencyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[global_ip_id][in]
type GlobalIPLatencyGetParamsFilter struct {
	// Filter by exact Global IP ID
	GlobalIPID GlobalIPLatencyGetParamsFilterGlobalIPIDUnion `query:"global_ip_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPLatencyGetParamsFilter]'s query parameters as
// `url.Values`.
func (r GlobalIPLatencyGetParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GlobalIPLatencyGetParamsFilterGlobalIPIDUnion struct {
	OfString                                param.Opt[string]                           `query:",omitzero,inline"`
	OfGlobalIPLatencyGetsFilterGlobalIPIDIn *GlobalIPLatencyGetParamsFilterGlobalIPIDIn `query:",omitzero,inline"`
	paramUnion
}

func (u *GlobalIPLatencyGetParamsFilterGlobalIPIDUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfGlobalIPLatencyGetsFilterGlobalIPIDIn) {
		return u.OfGlobalIPLatencyGetsFilterGlobalIPIDIn
	}
	return nil
}

// Filtering operations
type GlobalIPLatencyGetParamsFilterGlobalIPIDIn struct {
	// Filter by Global IP ID(s) separated by commas
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GlobalIPLatencyGetParamsFilterGlobalIPIDIn]'s query
// parameters as `url.Values`.
func (r GlobalIPLatencyGetParamsFilterGlobalIPIDIn) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
