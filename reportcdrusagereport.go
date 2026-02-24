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

// ReportCdrUsageReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportCdrUsageReportService] method instead.
type ReportCdrUsageReportService struct {
	Options []option.RequestOption
}

// NewReportCdrUsageReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportCdrUsageReportService(opts ...option.RequestOption) (r ReportCdrUsageReportService) {
	r = ReportCdrUsageReportService{}
	r.Options = opts
	return
}

// Generate and fetch voice usage report synchronously. This endpoint will both
// generate and fetch the voice report over a specified time period. No polling is
// necessary but the response may take up to a couple of minutes.
func (r *ReportCdrUsageReportService) FetchSync(ctx context.Context, query ReportCdrUsageReportFetchSyncParams, opts ...option.RequestOption) (res *ReportCdrUsageReportFetchSyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "reports/cdr_usage_reports/sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ReportCdrUsageReportFetchSyncResponse struct {
	Data ReportCdrUsageReportFetchSyncResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportCdrUsageReportFetchSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportCdrUsageReportFetchSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportCdrUsageReportFetchSyncResponseData struct {
	// Identifies the resource
	ID string `json:"id" format:"uuid"`
	// Any of "NO_AGGREGATION", "CONNECTION", "TAG", "BILLING_GROUP".
	AggregationType string    `json:"aggregation_type"`
	Connections     []int64   `json:"connections"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	EndTime         time.Time `json:"end_time" format:"date-time"`
	// Any of "NO_BREAKDOWN", "DID_VS_TOLL_FREE", "COUNTRY",
	// "DID_VS_TOLL_FREE_PER_COUNTRY".
	ProductBreakdown string         `json:"product_breakdown"`
	RecordType       string         `json:"record_type"`
	ReportURL        string         `json:"report_url"`
	Result           map[string]any `json:"result"`
	StartTime        time.Time      `json:"start_time" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AggregationType  respjson.Field
		Connections      respjson.Field
		CreatedAt        respjson.Field
		EndTime          respjson.Field
		ProductBreakdown respjson.Field
		RecordType       respjson.Field
		ReportURL        respjson.Field
		Result           respjson.Field
		StartTime        respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportCdrUsageReportFetchSyncResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReportCdrUsageReportFetchSyncResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportCdrUsageReportFetchSyncParams struct {
	// Any of "NO_AGGREGATION", "CONNECTION", "TAG", "BILLING_GROUP".
	AggregationType ReportCdrUsageReportFetchSyncParamsAggregationType `query:"aggregation_type,omitzero" api:"required" json:"-"`
	// Any of "NO_BREAKDOWN", "DID_VS_TOLL_FREE", "COUNTRY",
	// "DID_VS_TOLL_FREE_PER_COUNTRY".
	ProductBreakdown ReportCdrUsageReportFetchSyncParamsProductBreakdown `query:"product_breakdown,omitzero" api:"required" json:"-"`
	EndDate          param.Opt[time.Time]                                `query:"end_date,omitzero" format:"date-time" json:"-"`
	StartDate        param.Opt[time.Time]                                `query:"start_date,omitzero" format:"date-time" json:"-"`
	Connections      []float64                                           `query:"connections,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportCdrUsageReportFetchSyncParams]'s query parameters as
// `url.Values`.
func (r ReportCdrUsageReportFetchSyncParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportCdrUsageReportFetchSyncParamsAggregationType string

const (
	ReportCdrUsageReportFetchSyncParamsAggregationTypeNoAggregation ReportCdrUsageReportFetchSyncParamsAggregationType = "NO_AGGREGATION"
	ReportCdrUsageReportFetchSyncParamsAggregationTypeConnection    ReportCdrUsageReportFetchSyncParamsAggregationType = "CONNECTION"
	ReportCdrUsageReportFetchSyncParamsAggregationTypeTag           ReportCdrUsageReportFetchSyncParamsAggregationType = "TAG"
	ReportCdrUsageReportFetchSyncParamsAggregationTypeBillingGroup  ReportCdrUsageReportFetchSyncParamsAggregationType = "BILLING_GROUP"
)

type ReportCdrUsageReportFetchSyncParamsProductBreakdown string

const (
	ReportCdrUsageReportFetchSyncParamsProductBreakdownNoBreakdown             ReportCdrUsageReportFetchSyncParamsProductBreakdown = "NO_BREAKDOWN"
	ReportCdrUsageReportFetchSyncParamsProductBreakdownDidVsTollFree           ReportCdrUsageReportFetchSyncParamsProductBreakdown = "DID_VS_TOLL_FREE"
	ReportCdrUsageReportFetchSyncParamsProductBreakdownCountry                 ReportCdrUsageReportFetchSyncParamsProductBreakdown = "COUNTRY"
	ReportCdrUsageReportFetchSyncParamsProductBreakdownDidVsTollFreePerCountry ReportCdrUsageReportFetchSyncParamsProductBreakdown = "DID_VS_TOLL_FREE_PER_COUNTRY"
)
