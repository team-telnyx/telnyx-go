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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// LegacyReportingUsageReportMessagingService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingUsageReportMessagingService] method instead.
type LegacyReportingUsageReportMessagingService struct {
	Options []option.RequestOption
}

// NewLegacyReportingUsageReportMessagingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLegacyReportingUsageReportMessagingService(opts ...option.RequestOption) (r LegacyReportingUsageReportMessagingService) {
	r = LegacyReportingUsageReportMessagingService{}
	r.Options = opts
	return
}

// Creates a new legacy usage V2 MDR report request with the specified filters
func (r *LegacyReportingUsageReportMessagingService) New(ctx context.Context, body LegacyReportingUsageReportMessagingNewParams, opts ...option.RequestOption) (res *LegacyReportingUsageReportMessagingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/usage_reports/messaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch single MDR usage report by id.
func (r *LegacyReportingUsageReportMessagingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingUsageReportMessagingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/messaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch all previous requests for MDR usage reports.
func (r *LegacyReportingUsageReportMessagingService) List(ctx context.Context, query LegacyReportingUsageReportMessagingListParams, opts ...option.RequestOption) (res *pagination.PerPagePagination[MdrUsageReportResponseLegacy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "legacy/reporting/usage_reports/messaging"
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

// Fetch all previous requests for MDR usage reports.
func (r *LegacyReportingUsageReportMessagingService) ListAutoPaging(ctx context.Context, query LegacyReportingUsageReportMessagingListParams, opts ...option.RequestOption) *pagination.PerPagePaginationAutoPager[MdrUsageReportResponseLegacy] {
	return pagination.NewPerPagePaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific V2 legacy usage MDR report request by ID
func (r *LegacyReportingUsageReportMessagingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingUsageReportMessagingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/messaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Legacy V2 MDR usage report response
type MdrUsageReportResponseLegacy struct {
	// Identifies the resource
	ID string `json:"id" format:"uuid"`
	// Aggregation type: No aggregation = 0, By Messaging Profile = 1, By Tags = 2
	AggregationType int64     `json:"aggregation_type"`
	Connections     []string  `json:"connections"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	EndTime         time.Time `json:"end_time" format:"date-time"`
	// List of messaging profile IDs
	Profiles   []string       `json:"profiles" format:"uuid"`
	RecordType string         `json:"record_type"`
	ReportURL  string         `json:"report_url"`
	Result     map[string]any `json:"result"`
	StartTime  time.Time      `json:"start_time" format:"date-time"`
	// Status of the report (Pending = 1, Complete = 2, Failed = 3, Expired = 4)
	Status    int64     `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AggregationType respjson.Field
		Connections     respjson.Field
		CreatedAt       respjson.Field
		EndTime         respjson.Field
		Profiles        respjson.Field
		RecordType      respjson.Field
		ReportURL       respjson.Field
		Result          respjson.Field
		StartTime       respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MdrUsageReportResponseLegacy) RawJSON() string { return r.JSON.raw }
func (r *MdrUsageReportResponseLegacy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StandardPaginationMeta struct {
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
func (r StandardPaginationMeta) RawJSON() string { return r.JSON.raw }
func (r *StandardPaginationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportMessagingNewResponse struct {
	// Legacy V2 MDR usage report response
	Data MdrUsageReportResponseLegacy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportMessagingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportMessagingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportMessagingGetResponse struct {
	// Legacy V2 MDR usage report response
	Data MdrUsageReportResponseLegacy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportMessagingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportMessagingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportMessagingDeleteResponse struct {
	// Legacy V2 MDR usage report response
	Data MdrUsageReportResponseLegacy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportMessagingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportMessagingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportMessagingNewParams struct {
	// Aggregation type: No aggregation = 0, By Messaging Profile = 1, By Tags = 2
	AggregationType          int64                `json:"aggregation_type,required"`
	EndTime                  param.Opt[time.Time] `json:"end_time,omitzero" format:"date-time"`
	SelectAllManagedAccounts param.Opt[bool]      `json:"select_all_managed_accounts,omitzero"`
	StartTime                param.Opt[time.Time] `json:"start_time,omitzero" format:"date-time"`
	// List of managed accounts to include
	ManagedAccounts []string `json:"managed_accounts,omitzero" format:"uuid"`
	// List of messaging profile IDs to filter by
	Profiles []string `json:"profiles,omitzero" format:"uuid"`
	paramObj
}

func (r LegacyReportingUsageReportMessagingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingUsageReportMessagingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingUsageReportMessagingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportMessagingListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Size of the page
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LegacyReportingUsageReportMessagingListParams]'s query
// parameters as `url.Values`.
func (r LegacyReportingUsageReportMessagingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
