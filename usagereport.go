// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// UsageReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageReportService] method instead.
type UsageReportService struct {
	Options []option.RequestOption
}

// NewUsageReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsageReportService(opts ...option.RequestOption) (r UsageReportService) {
	r = UsageReportService{}
	r.Options = opts
	return
}

// Get Telnyx usage data by product, broken out by the specified dimensions
func (r *UsageReportService) List(ctx context.Context, params UsageReportListParams, opts ...option.RequestOption) (res *UsageReportListResponse, err error) {
	if !param.IsOmitted(params.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%s", params.AuthorizationBearer.Value)))
	}
	opts = append(r.Options[:], opts...)
	path := "usage_reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get the Usage Reports options for querying usage, including the products
// available and their respective metrics and dimensions
func (r *UsageReportService) GetOptions(ctx context.Context, params UsageReportGetOptionsParams, opts ...option.RequestOption) (res *UsageReportGetOptionsResponse, err error) {
	if !param.IsOmitted(params.AuthorizationBearer) {
		opts = append(opts, option.WithHeader("authorization_bearer", fmt.Sprintf("%s", params.AuthorizationBearer.Value)))
	}
	opts = append(r.Options[:], opts...)
	path := "usage_reports/options"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type UsageReportListResponse struct {
	Data []map[string]any            `json:"data"`
	Meta UsageReportListResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportListResponseMeta struct {
	// Selected page number.
	PageNumber int64 `json:"page_number"`
	// Records per single page
	PageSize int64 `json:"page_size"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages"`
	// Total number of results.
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *UsageReportListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object following one of the schemas published in
// https://developers.telnyx.com/docs/api/v2/detail-records
type UsageReportGetOptionsResponse struct {
	// Collection of product description
	Data []UsageReportGetOptionsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportGetOptionsResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageReportGetOptionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object following one of the schemas published in
// https://developers.telnyx.com/docs/api/v2/detail-records
type UsageReportGetOptionsResponseData struct {
	// Telnyx Product
	Product string `json:"product"`
	// Telnyx Product Dimensions
	ProductDimensions []string `json:"product_dimensions"`
	// Telnyx Product Metrics
	ProductMetrics []string `json:"product_metrics"`
	// Subproducts if applicable
	RecordTypes []UsageReportGetOptionsResponseDataRecordType `json:"record_types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Product           respjson.Field
		ProductDimensions respjson.Field
		ProductMetrics    respjson.Field
		RecordTypes       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportGetOptionsResponseData) RawJSON() string { return r.JSON.raw }
func (r *UsageReportGetOptionsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object following one of the schemas published in
// https://developers.telnyx.com/docs/api/v2/detail-records
type UsageReportGetOptionsResponseDataRecordType struct {
	// Telnyx Product Dimensions
	ProductDimensions []string `json:"product_dimensions"`
	// Telnyx Product Metrics
	ProductMetrics []string `json:"product_metrics"`
	// Telnyx Product type
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductDimensions respjson.Field
		ProductMetrics    respjson.Field
		RecordType        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageReportGetOptionsResponseDataRecordType) RawJSON() string { return r.JSON.raw }
func (r *UsageReportGetOptionsResponseDataRecordType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageReportListParams struct {
	// Breakout by specified product dimensions
	Dimensions []string `query:"dimensions,omitzero,required" json:"-"`
	// Specified product usage values
	Metrics []string `query:"metrics,omitzero,required" json:"-"`
	// Telnyx product
	Product string `query:"product,required" json:"-"`
	// A more user-friendly way to specify the timespan you want to filter by. More
	// options can be found in the Telnyx API Reference docs.
	DateRange param.Opt[string] `query:"date_range,omitzero" json:"-"`
	// The end date for the time range you are interested in. The maximum time range is
	// 31 days. Format: YYYY-MM-DDTHH:mm:ssZ
	EndDate param.Opt[string] `query:"end_date,omitzero" json:"-"`
	// Filter records on dimensions
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Return the aggregations for all Managed Accounts under the user making the
	// request.
	ManagedAccounts param.Opt[bool] `query:"managed_accounts,omitzero" json:"-"`
	// The start date for the time range you are interested in. The maximum time range
	// is 31 days. Format: YYYY-MM-DDTHH:mm:ssZ
	StartDate param.Opt[string] `query:"start_date,omitzero" json:"-"`
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	// Specify the response format (csv or json). JSON is returned by default, even if
	// not specified.
	//
	// Any of "csv", "json".
	Format UsageReportListParamsFormat `query:"format,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page UsageReportListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageReportListParams]'s query parameters as `url.Values`.
func (r UsageReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the response format (csv or json). JSON is returned by default, even if
// not specified.
type UsageReportListParamsFormat string

const (
	UsageReportListParamsFormatCsv  UsageReportListParamsFormat = "csv"
	UsageReportListParamsFormatJson UsageReportListParamsFormat = "json"
)

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type UsageReportListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	Size   param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageReportListParamsPage]'s query parameters as
// `url.Values`.
func (r UsageReportListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UsageReportGetOptionsParams struct {
	// Options (dimensions and metrics) for a given product. If none specified, all
	// products will be returned.
	Product param.Opt[string] `query:"product,omitzero" json:"-"`
	// Authenticates the request with your Telnyx API V2 KEY
	AuthorizationBearer param.Opt[string] `header:"authorization_bearer,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageReportGetOptionsParams]'s query parameters as
// `url.Values`.
func (r UsageReportGetOptionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
