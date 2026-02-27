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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Observability into Telnyx platform stability and performance.
//
// SetiService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSetiService] method instead.
type SetiService struct {
	Options []option.RequestOption
}

// NewSetiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSetiService(opts ...option.RequestOption) (r SetiService) {
	r = SetiService{}
	r.Options = opts
	return
}

// Returns the results of the various black box tests
func (r *SetiService) GetBlackBoxTestResults(ctx context.Context, query SetiGetBlackBoxTestResultsParams, opts ...option.RequestOption) (res *SetiGetBlackBoxTestResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "seti/black_box_test_results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SetiGetBlackBoxTestResultsResponse struct {
	Data []SetiGetBlackBoxTestResultsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetiGetBlackBoxTestResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *SetiGetBlackBoxTestResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SetiGetBlackBoxTestResultsResponseData struct {
	BlackBoxTests []SetiGetBlackBoxTestResultsResponseDataBlackBoxTest `json:"black_box_tests"`
	// The product associated with the black box test group.
	Product    string `json:"product"`
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlackBoxTests respjson.Field
		Product       respjson.Field
		RecordType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetiGetBlackBoxTestResultsResponseData) RawJSON() string { return r.JSON.raw }
func (r *SetiGetBlackBoxTestResultsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SetiGetBlackBoxTestResultsResponseDataBlackBoxTest struct {
	// The name of the black box test.
	ID         string `json:"id"`
	RecordType string `json:"record_type"`
	// The average result of the black box test over the last hour.
	Result float64 `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		RecordType  respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SetiGetBlackBoxTestResultsResponseDataBlackBoxTest) RawJSON() string { return r.JSON.raw }
func (r *SetiGetBlackBoxTestResultsResponseDataBlackBoxTest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SetiGetBlackBoxTestResultsParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[product]
	Filter SetiGetBlackBoxTestResultsParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SetiGetBlackBoxTestResultsParams]'s query parameters as
// `url.Values`.
func (r SetiGetBlackBoxTestResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[product]
type SetiGetBlackBoxTestResultsParamsFilter struct {
	// Filter results for a specific product.
	Product param.Opt[string] `query:"product,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SetiGetBlackBoxTestResultsParamsFilter]'s query parameters
// as `url.Values`.
func (r SetiGetBlackBoxTestResultsParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
