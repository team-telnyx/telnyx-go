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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Number portout operations
//
// PortoutReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortoutReportService] method instead.
type PortoutReportService struct {
	Options []option.RequestOption
}

// NewPortoutReportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortoutReportService(opts ...option.RequestOption) (r PortoutReportService) {
	r = PortoutReportService{}
	r.Options = opts
	return
}

// Generate reports about port-out operations.
func (r *PortoutReportService) New(ctx context.Context, body PortoutReportNewParams, opts ...option.RequestOption) (res *PortoutReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "portouts/reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific report generated.
func (r *PortoutReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortoutReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the reports generated about port-out operations.
func (r *PortoutReportService) List(ctx context.Context, query PortoutReportListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortoutReport], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "portouts/reports"
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

// List the reports generated about port-out operations.
func (r *PortoutReportService) ListAutoPaging(ctx context.Context, query PortoutReportListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortoutReport] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// The parameters for generating a port-outs CSV report.
type ExportPortoutsCsvReport struct {
	// The filters to apply to the export port-out CSV report.
	Filters ExportPortoutsCsvReportFilters `json:"filters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filters     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPortoutsCsvReport) RawJSON() string { return r.JSON.raw }
func (r *ExportPortoutsCsvReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExportPortoutsCsvReport to a ExportPortoutsCsvReportParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExportPortoutsCsvReportParam.Overrides()
func (r ExportPortoutsCsvReport) ToParam() ExportPortoutsCsvReportParam {
	return param.Override[ExportPortoutsCsvReportParam](json.RawMessage(r.RawJSON()))
}

// The filters to apply to the export port-out CSV report.
type ExportPortoutsCsvReportFilters struct {
	// The date and time the port-out was created after.
	CreatedAtGt time.Time `json:"created_at__gt" format:"date-time"`
	// The date and time the port-out was created before.
	CreatedAtLt time.Time `json:"created_at__lt" format:"date-time"`
	// The customer reference of the port-outs to include in the report.
	CustomerReferenceIn []string `json:"customer_reference__in"`
	// The end user name of the port-outs to include in the report.
	EndUserName string `json:"end_user_name"`
	// A list of phone numbers that the port-outs phone numbers must overlap with.
	PhoneNumbersOverlaps []string `json:"phone_numbers__overlaps"`
	// The status of the port-outs to include in the report.
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	StatusIn []string `json:"status__in"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAtGt          respjson.Field
		CreatedAtLt          respjson.Field
		CustomerReferenceIn  respjson.Field
		EndUserName          respjson.Field
		PhoneNumbersOverlaps respjson.Field
		StatusIn             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExportPortoutsCsvReportFilters) RawJSON() string { return r.JSON.raw }
func (r *ExportPortoutsCsvReportFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters for generating a port-outs CSV report.
//
// The property Filters is required.
type ExportPortoutsCsvReportParam struct {
	// The filters to apply to the export port-out CSV report.
	Filters ExportPortoutsCsvReportFiltersParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r ExportPortoutsCsvReportParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPortoutsCsvReportParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPortoutsCsvReportParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The filters to apply to the export port-out CSV report.
type ExportPortoutsCsvReportFiltersParam struct {
	// The date and time the port-out was created after.
	CreatedAtGt param.Opt[time.Time] `json:"created_at__gt,omitzero" format:"date-time"`
	// The date and time the port-out was created before.
	CreatedAtLt param.Opt[time.Time] `json:"created_at__lt,omitzero" format:"date-time"`
	// The end user name of the port-outs to include in the report.
	EndUserName param.Opt[string] `json:"end_user_name,omitzero"`
	// The customer reference of the port-outs to include in the report.
	CustomerReferenceIn []string `json:"customer_reference__in,omitzero"`
	// A list of phone numbers that the port-outs phone numbers must overlap with.
	PhoneNumbersOverlaps []string `json:"phone_numbers__overlaps,omitzero"`
	// The status of the port-outs to include in the report.
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	StatusIn []string `json:"status__in,omitzero"`
	paramObj
}

func (r ExportPortoutsCsvReportFiltersParam) MarshalJSON() (data []byte, err error) {
	type shadow ExportPortoutsCsvReportFiltersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExportPortoutsCsvReportFiltersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutReport struct {
	// Uniquely identifies the report.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the document that was uploaded when report was generated. This field
	// is only populated when the report is under completed status.
	DocumentID string `json:"document_id" format:"uuid"`
	// The parameters for generating a port-outs CSV report.
	Params ExportPortoutsCsvReport `json:"params"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Identifies the type of report
	//
	// Any of "export_portouts_csv".
	ReportType PortoutReportReportType `json:"report_type"`
	// The current status of the report generation.
	//
	// Any of "pending", "completed".
	Status PortoutReportStatus `json:"status"`
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
func (r PortoutReport) RawJSON() string { return r.JSON.raw }
func (r *PortoutReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of report
type PortoutReportReportType string

const (
	PortoutReportReportTypeExportPortoutsCsv PortoutReportReportType = "export_portouts_csv"
)

// The current status of the report generation.
type PortoutReportStatus string

const (
	PortoutReportStatusPending   PortoutReportStatus = "pending"
	PortoutReportStatusCompleted PortoutReportStatus = "completed"
)

type PortoutReportNewResponse struct {
	Data PortoutReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutReportGetResponse struct {
	Data PortoutReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutReportNewParams struct {
	// The parameters for generating a port-outs CSV report.
	Params ExportPortoutsCsvReportParam `json:"params,omitzero" api:"required"`
	// Identifies the type of report
	//
	// Any of "export_portouts_csv".
	ReportType PortoutReportNewParamsReportType `json:"report_type,omitzero" api:"required"`
	paramObj
}

func (r PortoutReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PortoutReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortoutReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of report
type PortoutReportNewParamsReportType string

const (
	PortoutReportNewParamsReportTypeExportPortoutsCsv PortoutReportNewParamsReportType = "export_portouts_csv"
)

type PortoutReportListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[report_type], filter[status]
	Filter PortoutReportListParamsFilter `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutReportListParams]'s query parameters as
// `url.Values`.
func (r PortoutReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[report_type], filter[status]
type PortoutReportListParamsFilter struct {
	// Filter reports of a specific type
	//
	// Any of "export_portouts_csv".
	ReportType string `query:"report_type,omitzero" json:"-"`
	// Filter reports of a specific status
	//
	// Any of "pending", "completed".
	Status string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutReportListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortoutReportListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
