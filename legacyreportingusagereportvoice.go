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

// Voice usage reports
//
// LegacyReportingUsageReportVoiceService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingUsageReportVoiceService] method instead.
type LegacyReportingUsageReportVoiceService struct {
	Options []option.RequestOption
}

// NewLegacyReportingUsageReportVoiceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLegacyReportingUsageReportVoiceService(opts ...option.RequestOption) (r LegacyReportingUsageReportVoiceService) {
	r = LegacyReportingUsageReportVoiceService{}
	r.Options = opts
	return
}

// Creates a new legacy usage V2 CDR report request with the specified filters
func (r *LegacyReportingUsageReportVoiceService) New(ctx context.Context, body LegacyReportingUsageReportVoiceNewParams, opts ...option.RequestOption) (res *LegacyReportingUsageReportVoiceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/usage_reports/voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch single cdr usage report by id.
func (r *LegacyReportingUsageReportVoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingUsageReportVoiceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/voice/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch all previous requests for cdr usage reports.
func (r *LegacyReportingUsageReportVoiceService) List(ctx context.Context, query LegacyReportingUsageReportVoiceListParams, opts ...option.RequestOption) (res *pagination.PerPagePagination[CdrUsageReportResponseLegacy], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "legacy/reporting/usage_reports/voice"
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

// Fetch all previous requests for cdr usage reports.
func (r *LegacyReportingUsageReportVoiceService) ListAutoPaging(ctx context.Context, query LegacyReportingUsageReportVoiceListParams, opts ...option.RequestOption) *pagination.PerPagePaginationAutoPager[CdrUsageReportResponseLegacy] {
	return pagination.NewPerPagePaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a specific V2 legacy usage CDR report request by ID
func (r *LegacyReportingUsageReportVoiceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingUsageReportVoiceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/usage_reports/voice/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Legacy V2 CDR usage report response
type CdrUsageReportResponseLegacy struct {
	// Identifies the resource
	ID string `json:"id" format:"uuid"`
	// Aggregation type: All = 0, By Connections = 1, By Tags = 2, By Billing Group = 3
	AggregationType int64     `json:"aggregation_type"`
	Connections     []string  `json:"connections"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	EndTime         time.Time `json:"end_time" format:"date-time"`
	// Product breakdown type: No breakdown = 0, DID vs Toll-free = 1, Country = 2, DID
	// vs Toll-free per Country = 3
	ProductBreakdown int64          `json:"product_breakdown"`
	RecordType       string         `json:"record_type"`
	ReportURL        string         `json:"report_url"`
	Result           map[string]any `json:"result"`
	StartTime        time.Time      `json:"start_time" format:"date-time"`
	// Status of the report: Pending = 1, Complete = 2, Failed = 3, Expired = 4
	Status    int64     `json:"status"`
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
func (r CdrUsageReportResponseLegacy) RawJSON() string { return r.JSON.raw }
func (r *CdrUsageReportResponseLegacy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportVoiceNewResponse struct {
	// Legacy V2 CDR usage report response
	Data CdrUsageReportResponseLegacy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportVoiceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportVoiceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportVoiceGetResponse struct {
	// Legacy V2 CDR usage report response
	Data CdrUsageReportResponseLegacy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportVoiceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportVoiceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportVoiceDeleteResponse struct {
	// Legacy V2 CDR usage report response
	Data CdrUsageReportResponseLegacy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingUsageReportVoiceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingUsageReportVoiceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportVoiceNewParams struct {
	// End time in ISO format
	EndTime time.Time `json:"end_time" api:"required" format:"date-time"`
	// Start time in ISO format
	StartTime time.Time `json:"start_time" api:"required" format:"date-time"`
	// Aggregation type: All = 0, By Connections = 1, By Tags = 2, By Billing Group = 3
	AggregationType param.Opt[int64] `json:"aggregation_type,omitzero"`
	// Product breakdown type: No breakdown = 0, DID vs Toll-free = 1, Country = 2, DID
	// vs Toll-free per Country = 3
	ProductBreakdown param.Opt[int64] `json:"product_breakdown,omitzero"`
	// Whether to select all managed accounts
	SelectAllManagedAccounts param.Opt[bool] `json:"select_all_managed_accounts,omitzero"`
	// List of connections to filter by
	Connections []int64 `json:"connections,omitzero"`
	// List of managed accounts to include
	ManagedAccounts []string `json:"managed_accounts,omitzero" format:"uuid"`
	paramObj
}

func (r LegacyReportingUsageReportVoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingUsageReportVoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingUsageReportVoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingUsageReportVoiceListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Size of the page
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LegacyReportingUsageReportVoiceListParams]'s query
// parameters as `url.Values`.
func (r LegacyReportingUsageReportVoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
