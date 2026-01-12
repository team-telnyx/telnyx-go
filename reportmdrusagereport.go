// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// ReportMdrUsageReportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportMdrUsageReportService] method instead.
type ReportMdrUsageReportService struct {
	Options []option.RequestOption
}

// NewReportMdrUsageReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReportMdrUsageReportService(opts ...option.RequestOption) (r ReportMdrUsageReportService) {
	r = ReportMdrUsageReportService{}
	r.Options = opts
	return
}

// Submit request for new new messaging usage report. This endpoint will pull and
// aggregate messaging data in specified time period.
func (r *ReportMdrUsageReportService) New(ctx context.Context, body ReportMdrUsageReportNewParams, opts ...option.RequestOption) (res *ReportMdrUsageReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "reports/mdr_usage_reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single messaging usage report by id
func (r *ReportMdrUsageReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ReportMdrUsageReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("reports/mdr_usage_reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch all messaging usage reports. Usage reports are aggregated messaging data
// for specified time period and breakdown
func (r *ReportMdrUsageReportService) List(ctx context.Context, query ReportMdrUsageReportListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[MdrUsageReport], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "reports/mdr_usage_reports"
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

// Fetch all messaging usage reports. Usage reports are aggregated messaging data
// for specified time period and breakdown
func (r *ReportMdrUsageReportService) ListAutoPaging(ctx context.Context, query ReportMdrUsageReportListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[MdrUsageReport] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete messaging usage report by id
func (r *ReportMdrUsageReportService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ReportMdrUsageReportDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("reports/mdr_usage_reports/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generate and fetch messaging usage report synchronously. This endpoint will both
// generate and fetch the messaging report over a specified time period. No polling
// is necessary but the response may take up to a couple of minutes.
func (r *ReportMdrUsageReportService) FetchSync(ctx context.Context, query ReportMdrUsageReportFetchSyncParams, opts ...option.RequestOption) (res *ReportMdrUsageReportFetchSyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "reports/mdr_usage_reports/sync"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MdrUsageReport struct {
	// Identifies the resource
	ID string `json:"id" format:"uuid"`
	// Any of "NO_AGGREGATION", "PROFILE", "TAGS".
	AggregationType MdrUsageReportAggregationType `json:"aggregation_type"`
	Connections     []int64                       `json:"connections"`
	CreatedAt       time.Time                     `json:"created_at" format:"date-time"`
	EndDate         time.Time                     `json:"end_date" format:"date-time"`
	Profiles        string                        `json:"profiles"`
	RecordType      string                        `json:"record_type"`
	ReportURL       string                        `json:"report_url"`
	Result          []MdrUsageReportResult        `json:"result"`
	StartDate       time.Time                     `json:"start_date" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status    MdrUsageReportStatus `json:"status"`
	UpdatedAt time.Time            `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AggregationType respjson.Field
		Connections     respjson.Field
		CreatedAt       respjson.Field
		EndDate         respjson.Field
		Profiles        respjson.Field
		RecordType      respjson.Field
		ReportURL       respjson.Field
		Result          respjson.Field
		StartDate       respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MdrUsageReport) RawJSON() string { return r.JSON.raw }
func (r *MdrUsageReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MdrUsageReportAggregationType string

const (
	MdrUsageReportAggregationTypeNoAggregation MdrUsageReportAggregationType = "NO_AGGREGATION"
	MdrUsageReportAggregationTypeProfile       MdrUsageReportAggregationType = "PROFILE"
	MdrUsageReportAggregationTypeTags          MdrUsageReportAggregationType = "TAGS"
)

type MdrUsageReportResult struct {
	CarrierPassthroughFee string `json:"carrier_passthrough_fee"`
	Connection            string `json:"connection"`
	Cost                  string `json:"cost"`
	Currency              string `json:"currency"`
	Delivered             string `json:"delivered"`
	Direction             string `json:"direction"`
	MessageType           string `json:"message_type"`
	Parts                 string `json:"parts"`
	Product               string `json:"product"`
	ProfileID             string `json:"profile_id"`
	Received              string `json:"received"`
	Sent                  string `json:"sent"`
	Tags                  string `json:"tags"`
	TnType                string `json:"tn_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierPassthroughFee respjson.Field
		Connection            respjson.Field
		Cost                  respjson.Field
		Currency              respjson.Field
		Delivered             respjson.Field
		Direction             respjson.Field
		MessageType           respjson.Field
		Parts                 respjson.Field
		Product               respjson.Field
		ProfileID             respjson.Field
		Received              respjson.Field
		Sent                  respjson.Field
		Tags                  respjson.Field
		TnType                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MdrUsageReportResult) RawJSON() string { return r.JSON.raw }
func (r *MdrUsageReportResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MdrUsageReportStatus string

const (
	MdrUsageReportStatusPending  MdrUsageReportStatus = "PENDING"
	MdrUsageReportStatusComplete MdrUsageReportStatus = "COMPLETE"
	MdrUsageReportStatusFailed   MdrUsageReportStatus = "FAILED"
	MdrUsageReportStatusExpired  MdrUsageReportStatus = "EXPIRED"
)

type PaginationMetaReporting struct {
	PageNumber   int64 `json:"page_number,required"`
	TotalPages   int64 `json:"total_pages,required"`
	PageSize     int64 `json:"page_size"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		TotalPages   respjson.Field
		PageSize     respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginationMetaReporting) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetaReporting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportMdrUsageReportNewResponse struct {
	Data MdrUsageReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportMdrUsageReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportMdrUsageReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportMdrUsageReportGetResponse struct {
	Data MdrUsageReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportMdrUsageReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportMdrUsageReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportMdrUsageReportDeleteResponse struct {
	Data MdrUsageReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportMdrUsageReportDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportMdrUsageReportDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportMdrUsageReportFetchSyncResponse struct {
	Data MdrUsageReport `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportMdrUsageReportFetchSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportMdrUsageReportFetchSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportMdrUsageReportNewParams struct {
	// Any of "NO_AGGREGATION", "PROFILE", "TAGS".
	AggregationType ReportMdrUsageReportNewParamsAggregationType `json:"aggregation_type,omitzero,required"`
	EndDate         time.Time                                    `json:"end_date,required" format:"date-time"`
	StartDate       time.Time                                    `json:"start_date,required" format:"date-time"`
	Profiles        param.Opt[string]                            `json:"profiles,omitzero"`
	paramObj
}

func (r ReportMdrUsageReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ReportMdrUsageReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReportMdrUsageReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportMdrUsageReportNewParamsAggregationType string

const (
	ReportMdrUsageReportNewParamsAggregationTypeNoAggregation ReportMdrUsageReportNewParamsAggregationType = "NO_AGGREGATION"
	ReportMdrUsageReportNewParamsAggregationTypeProfile       ReportMdrUsageReportNewParamsAggregationType = "PROFILE"
	ReportMdrUsageReportNewParamsAggregationTypeTags          ReportMdrUsageReportNewParamsAggregationType = "TAGS"
)

type ReportMdrUsageReportListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportMdrUsageReportListParams]'s query parameters as
// `url.Values`.
func (r ReportMdrUsageReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportMdrUsageReportFetchSyncParams struct {
	// Any of "NO_AGGREGATION", "PROFILE", "TAGS".
	AggregationType ReportMdrUsageReportFetchSyncParamsAggregationType `query:"aggregation_type,omitzero,required" json:"-"`
	EndDate         param.Opt[time.Time]                               `query:"end_date,omitzero" format:"date-time" json:"-"`
	StartDate       param.Opt[time.Time]                               `query:"start_date,omitzero" format:"date-time" json:"-"`
	Profiles        []string                                           `query:"profiles,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportMdrUsageReportFetchSyncParams]'s query parameters as
// `url.Values`.
func (r ReportMdrUsageReportFetchSyncParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportMdrUsageReportFetchSyncParamsAggregationType string

const (
	ReportMdrUsageReportFetchSyncParamsAggregationTypeNoAggregation ReportMdrUsageReportFetchSyncParamsAggregationType = "NO_AGGREGATION"
	ReportMdrUsageReportFetchSyncParamsAggregationTypeProfile       ReportMdrUsageReportFetchSyncParamsAggregationType = "PROFILE"
	ReportMdrUsageReportFetchSyncParamsAggregationTypeTags          ReportMdrUsageReportFetchSyncParamsAggregationType = "TAGS"
)
