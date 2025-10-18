// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// LegacyReportingBatchDetailRecordVoiceService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingBatchDetailRecordVoiceService] method instead.
type LegacyReportingBatchDetailRecordVoiceService struct {
	Options []option.RequestOption
}

// NewLegacyReportingBatchDetailRecordVoiceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLegacyReportingBatchDetailRecordVoiceService(opts ...option.RequestOption) (r LegacyReportingBatchDetailRecordVoiceService) {
	r = LegacyReportingBatchDetailRecordVoiceService{}
	r.Options = opts
	return
}

// Creates a new CDR report request with the specified filters
func (r *LegacyReportingBatchDetailRecordVoiceService) New(ctx context.Context, body LegacyReportingBatchDetailRecordVoiceNewParams, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordVoiceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific CDR report request by ID
func (r *LegacyReportingBatchDetailRecordVoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordVoiceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/batch_detail_records/voice/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all CDR report requests for the authenticated user
func (r *LegacyReportingBatchDetailRecordVoiceService) List(ctx context.Context, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordVoiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific CDR report request by ID
func (r *LegacyReportingBatchDetailRecordVoiceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordVoiceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/batch_detail_records/voice/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieves all available fields that can be used in CDR reports
func (r *LegacyReportingBatchDetailRecordVoiceService) GetFields(ctx context.Context, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordVoiceGetFieldsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/voice/fields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response object for CDR detailed report
type CdrDetailedReqResponse struct {
	// Unique identifier for the report
	ID string `json:"id"`
	// List of call types (Inbound = 1, Outbound = 2)
	CallTypes []int64 `json:"call_types"`
	// List of connections
	Connections []int64 `json:"connections"`
	// Creation date of the report
	CreatedAt string `json:"created_at"`
	// End time in ISO format
	EndTime string `json:"end_time"`
	// List of filters
	Filters []Filter `json:"filters"`
	// List of managed accounts
	ManagedAccounts []string `json:"managed_accounts" format:"uuid"`
	RecordType      string   `json:"record_type"`
	// List of record types (Complete = 1, Incomplete = 2, Errors = 3)
	RecordTypes []int64 `json:"record_types"`
	// Name of the report
	ReportName string `json:"report_name"`
	// URL to download the report
	ReportURL string `json:"report_url"`
	// Number of retries
	Retry int64 `json:"retry"`
	// Source of the report. Valid values: calls (default), call-control, fax-api,
	// webrtc
	Source string `json:"source"`
	// Start time in ISO format
	StartTime string `json:"start_time"`
	// Status of the report (Pending = 1, Complete = 2, Failed = 3, Expired = 4)
	Status int64 `json:"status"`
	// Timezone for the report
	Timezone string `json:"timezone"`
	// Last update date of the report
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CallTypes       respjson.Field
		Connections     respjson.Field
		CreatedAt       respjson.Field
		EndTime         respjson.Field
		Filters         respjson.Field
		ManagedAccounts respjson.Field
		RecordType      respjson.Field
		RecordTypes     respjson.Field
		ReportName      respjson.Field
		ReportURL       respjson.Field
		Retry           respjson.Field
		Source          respjson.Field
		StartTime       respjson.Field
		Status          respjson.Field
		Timezone        respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CdrDetailedReqResponse) RawJSON() string { return r.JSON.raw }
func (r *CdrDetailedReqResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordVoiceNewResponse struct {
	// Response object for CDR detailed report
	Data CdrDetailedReqResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordVoiceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordVoiceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordVoiceGetResponse struct {
	// Response object for CDR detailed report
	Data CdrDetailedReqResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordVoiceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordVoiceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordVoiceListResponse struct {
	Data []CdrDetailedReqResponse `json:"data"`
	Meta BatchCsvPaginationMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordVoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordVoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordVoiceDeleteResponse struct {
	// Response object for CDR detailed report
	Data CdrDetailedReqResponse `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordVoiceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordVoiceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Available CDR report fields grouped by category
type LegacyReportingBatchDetailRecordVoiceGetFieldsResponse struct {
	// Cost and billing related information
	Billing []string `json:"Billing"`
	// Fields related to call interaction and basic call information
	InteractionData []string `json:"Interaction Data"`
	// Geographic and routing information for phone numbers
	NumberInformation []string `json:"Number Information"`
	// Technical telephony and call control information
	TelephonyData []string `json:"Telephony Data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Billing           respjson.Field
		InteractionData   respjson.Field
		NumberInformation respjson.Field
		TelephonyData     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordVoiceGetFieldsResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordVoiceGetFieldsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordVoiceNewParams struct {
	// End time in ISO format
	EndTime time.Time `json:"end_time,required" format:"date-time"`
	// Start time in ISO format
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Whether to include all metadata
	IncludeAllMetadata param.Opt[bool] `json:"include_all_metadata,omitzero"`
	// Name of the report
	ReportName param.Opt[string] `json:"report_name,omitzero"`
	// Whether to select all managed accounts
	SelectAllManagedAccounts param.Opt[bool] `json:"select_all_managed_accounts,omitzero"`
	// Source of the report. Valid values: calls (default), call-control, fax-api,
	// webrtc
	Source param.Opt[string] `json:"source,omitzero"`
	// Timezone for the report
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// List of call types to filter by (Inbound = 1, Outbound = 2)
	CallTypes []int64 `json:"call_types,omitzero"`
	// List of connections to filter by
	Connections []int64 `json:"connections,omitzero"`
	// Set of fields to include in the report
	Fields []string `json:"fields,omitzero"`
	// List of filters to apply
	Filters []FilterParam `json:"filters,omitzero"`
	// List of managed accounts to include
	ManagedAccounts []string `json:"managed_accounts,omitzero" format:"uuid"`
	// List of record types to filter by (Complete = 1, Incomplete = 2, Errors = 3)
	RecordTypes []int64 `json:"record_types,omitzero"`
	paramObj
}

func (r LegacyReportingBatchDetailRecordVoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingBatchDetailRecordVoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingBatchDetailRecordVoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
