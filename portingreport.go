// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// PortingReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingReportService] method instead.
type PortingReportService struct {
	Options []option.RequestOption
}

// NewPortingReportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortingReportService(opts ...option.RequestOption) (r PortingReportService) {
	r = PortingReportService{}
	r.Options = opts
	return
}

// Generate reports about porting operations.
func (r *PortingReportService) New(ctx context.Context, body PortingReportNewParams, opts ...option.RequestOption) (res *PortingReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting/reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific report generated.
func (r *PortingReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the reports generated about porting operations.
func (r *PortingReportService) List(ctx context.Context, query PortingReportListParams, opts ...option.RequestOption) (res *PortingReportListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting/reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The parameters for generating a porting orders CSV report.
type ExportPortingOrdersCsvReport struct {
	// The filters to apply to the export porting order CSV report.
	Filters ExportPortingOrdersCsvReportFilters `json:"filters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPortingOrdersCsvReport) RawJSON() string { return r.JSON.raw }
func (r *ExportPortingOrdersCsvReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExportPortingOrdersCsvReport to a
// ExportPortingOrdersCsvReportParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExportPortingOrdersCsvReportParam.Overrides()
func (r ExportPortingOrdersCsvReport) ToParam() ExportPortingOrdersCsvReportParam {
	return param.Override[ExportPortingOrdersCsvReportParam](json.RawMessage(r.RawJSON()))
}

// The filters to apply to the export porting order CSV report.
type ExportPortingOrdersCsvReportFilters struct {
	// The date and time the porting order was created after.
	CreatedAtGt time.Time `json:"created_at__gt" format:"date-time"`
	// The date and time the porting order was created before.
	CreatedAtLt time.Time `json:"created_at__lt" format:"date-time"`
	// The customer reference of the porting orders to include in the report.
	CustomerReferenceIn []string `json:"customer_reference__in"`
	// The status of the porting orders to include in the report.
	//
	// Any of "draft", "in-process", "submitted", "exception", "foc-date-confirmed",
	// "cancel-pending", "ported", "cancelled".
	StatusIn []string `json:"status__in"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAtGt         respjson.Field
		CreatedAtLt         respjson.Field
		CustomerReferenceIn respjson.Field
		StatusIn            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPortingOrdersCsvReportFilters) RawJSON() string { return r.JSON.raw }
func (r *ExportPortingOrdersCsvReportFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters for generating a porting orders CSV report.
//
// The property Filters is required.
type ExportPortingOrdersCsvReportParam struct {
	// The filters to apply to the export porting order CSV report.
	Filters ExportPortingOrdersCsvReportFiltersParam `json:"filters,omitzero,required"`
	paramObj
}

func (r ExportPortingOrdersCsvReportParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPortingOrdersCsvReportParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPortingOrdersCsvReportParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The filters to apply to the export porting order CSV report.
type ExportPortingOrdersCsvReportFiltersParam struct {
	// The date and time the porting order was created after.
	CreatedAtGt param.Opt[time.Time] `json:"created_at__gt,omitzero" format:"date-time"`
	// The date and time the porting order was created before.
	CreatedAtLt param.Opt[time.Time] `json:"created_at__lt,omitzero" format:"date-time"`
	// The customer reference of the porting orders to include in the report.
	CustomerReferenceIn []string `json:"customer_reference__in,omitzero"`
	// The status of the porting orders to include in the report.
	//
	// Any of "draft", "in-process", "submitted", "exception", "foc-date-confirmed",
	// "cancel-pending", "ported", "cancelled".
	StatusIn []string `json:"status__in,omitzero"`
	paramObj
}

func (r ExportPortingOrdersCsvReportFiltersParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPortingOrdersCsvReportFiltersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPortingOrdersCsvReportFiltersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingReport struct {
	// Uniquely identifies the report.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the document that was uploaded when report was generated. This field
	// is only populated when the report is under completed status.
	DocumentID string `json:"document_id" format:"uuid"`
	// The parameters for generating a porting orders CSV report.
	Params ExportPortingOrdersCsvReport `json:"params"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Identifies the type of report
	//
	// Any of "export_porting_orders_csv".
	ReportType PortingReportReportType `json:"report_type"`
	// The current status of the report generation.
	//
	// Any of "pending", "completed".
	Status PortingReportStatus `json:"status"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DocumentID  respjson.Field
		Params      respjson.Field
		RecordType  respjson.Field
		ReportType  respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingReport) RawJSON() string { return r.JSON.raw }
func (r *PortingReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of report
type PortingReportReportType string

const (
	PortingReportReportTypeExportPortingOrdersCsv PortingReportReportType = "export_porting_orders_csv"
)

// The current status of the report generation.
type PortingReportStatus string

const (
	PortingReportStatusPending   PortingReportStatus = "pending"
	PortingReportStatusCompleted PortingReportStatus = "completed"
)

type PortingReportNewResponse struct {
	Data PortingReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingReportGetResponse struct {
	Data PortingReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingReportListResponse struct {
	Data []PortingReport `json:"data"`
	Meta PaginationMeta  `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingReportNewParams struct {
	// The parameters for generating a porting orders CSV report.
	Params ExportPortingOrdersCsvReportParam `json:"params,omitzero,required"`
	// Identifies the type of report
	//
	// Any of "export_porting_orders_csv".
	ReportType PortingReportNewParamsReportType `json:"report_type,omitzero,required"`
	paramObj
}

func (r PortingReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of report
type PortingReportNewParamsReportType string

const (
	PortingReportNewParamsReportTypeExportPortingOrdersCsv PortingReportNewParamsReportType = "export_porting_orders_csv"
)

type PortingReportListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[report_type], filter[status]
	Filter PortingReportListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingReportListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingReportListParams]'s query parameters as
// `url.Values`.
func (r PortingReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[report_type], filter[status]
type PortingReportListParamsFilter struct {
	// Filter reports of a specific type
	//
	// Any of "export_porting_orders_csv".
	ReportType string `query:"report_type,omitzero" json:"-"`
	// Filter reports of a specific status
	//
	// Any of "pending", "completed".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingReportListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortingReportListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingReportListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingReportListParamsPage]'s query parameters as
// `url.Values`.
func (r PortingReportListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
