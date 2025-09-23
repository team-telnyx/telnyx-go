// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// LegacyReportingBatchDetailRecordMessagingService contains methods and other
// services that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLegacyReportingBatchDetailRecordMessagingService] method instead.
type LegacyReportingBatchDetailRecordMessagingService struct {
	Options []option.RequestOption
}

// NewLegacyReportingBatchDetailRecordMessagingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLegacyReportingBatchDetailRecordMessagingService(opts ...option.RequestOption) (r LegacyReportingBatchDetailRecordMessagingService) {
	r = LegacyReportingBatchDetailRecordMessagingService{}
	r.Options = opts
	return
}

// Creates a new MDR detailed report request with the specified filters
func (r *LegacyReportingBatchDetailRecordMessagingService) New(ctx context.Context, body LegacyReportingBatchDetailRecordMessagingNewParams, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordMessagingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/messaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific MDR detailed report request by ID
func (r *LegacyReportingBatchDetailRecordMessagingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordMessagingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/batch_detail_records/messaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves all MDR detailed report requests for the authenticated user
func (r *LegacyReportingBatchDetailRecordMessagingService) List(ctx context.Context, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordMessagingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "legacy/reporting/batch_detail_records/messaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific MDR detailed report request by ID
func (r *LegacyReportingBatchDetailRecordMessagingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *LegacyReportingBatchDetailRecordMessagingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("legacy/reporting/batch_detail_records/messaging/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LegacyReportingBatchDetailRecordMessagingNewResponse struct {
	Data LegacyReportingBatchDetailRecordMessagingNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordMessagingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingNewResponseData struct {
	// Identifies the resource
	ID          string    `json:"id" format:"uuid"`
	Connections []int64   `json:"connections"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	// Any of "INBOUND", "OUTBOUND".
	Directions []string                                                         `json:"directions"`
	EndDate    time.Time                                                        `json:"end_date" format:"date-time"`
	Filters    []LegacyReportingBatchDetailRecordMessagingNewResponseDataFilter `json:"filters"`
	// List of messaging profile IDs
	Profiles   []string `json:"profiles" format:"uuid"`
	RecordType string   `json:"record_type"`
	// Any of "INCOMPLETE", "COMPLETED", "ERRORS".
	RecordTypes []string  `json:"record_types"`
	ReportName  string    `json:"report_name"`
	ReportURL   string    `json:"report_url"`
	StartDate   time.Time `json:"start_date" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Connections respjson.Field
		CreatedAt   respjson.Field
		Directions  respjson.Field
		EndDate     respjson.Field
		Filters     respjson.Field
		Profiles    respjson.Field
		RecordType  respjson.Field
		RecordTypes respjson.Field
		ReportName  respjson.Field
		ReportURL   respjson.Field
		StartDate   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordMessagingNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type LegacyReportingBatchDetailRecordMessagingNewResponseDataFilter struct {
	// Billing group UUID to filter by
	BillingGroup string `json:"billing_group"`
	// Called line identification (destination number)
	Cld string `json:"cld"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter string `json:"cld_filter"`
	// Calling line identification (caller ID)
	Cli string `json:"cli"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter string `json:"cli_filter"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType string `json:"filter_type"`
	// Tag name to filter by
	TagsList string `json:"tags_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingGroup respjson.Field
		Cld          respjson.Field
		CldFilter    respjson.Field
		Cli          respjson.Field
		CliFilter    respjson.Field
		FilterType   respjson.Field
		TagsList     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingNewResponseDataFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingNewResponseDataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingGetResponse struct {
	Data LegacyReportingBatchDetailRecordMessagingGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordMessagingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingGetResponseData struct {
	// Identifies the resource
	ID          string    `json:"id" format:"uuid"`
	Connections []int64   `json:"connections"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	// Any of "INBOUND", "OUTBOUND".
	Directions []string                                                         `json:"directions"`
	EndDate    time.Time                                                        `json:"end_date" format:"date-time"`
	Filters    []LegacyReportingBatchDetailRecordMessagingGetResponseDataFilter `json:"filters"`
	// List of messaging profile IDs
	Profiles   []string `json:"profiles" format:"uuid"`
	RecordType string   `json:"record_type"`
	// Any of "INCOMPLETE", "COMPLETED", "ERRORS".
	RecordTypes []string  `json:"record_types"`
	ReportName  string    `json:"report_name"`
	ReportURL   string    `json:"report_url"`
	StartDate   time.Time `json:"start_date" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Connections respjson.Field
		CreatedAt   respjson.Field
		Directions  respjson.Field
		EndDate     respjson.Field
		Filters     respjson.Field
		Profiles    respjson.Field
		RecordType  respjson.Field
		RecordTypes respjson.Field
		ReportName  respjson.Field
		ReportURL   respjson.Field
		StartDate   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordMessagingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type LegacyReportingBatchDetailRecordMessagingGetResponseDataFilter struct {
	// Billing group UUID to filter by
	BillingGroup string `json:"billing_group"`
	// Called line identification (destination number)
	Cld string `json:"cld"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter string `json:"cld_filter"`
	// Calling line identification (caller ID)
	Cli string `json:"cli"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter string `json:"cli_filter"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType string `json:"filter_type"`
	// Tag name to filter by
	TagsList string `json:"tags_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingGroup respjson.Field
		Cld          respjson.Field
		CldFilter    respjson.Field
		Cli          respjson.Field
		CliFilter    respjson.Field
		FilterType   respjson.Field
		TagsList     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingGetResponseDataFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingGetResponseDataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingListResponse struct {
	Data []LegacyReportingBatchDetailRecordMessagingListResponseData `json:"data"`
	Meta LegacyReportingBatchDetailRecordMessagingListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingListResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordMessagingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingListResponseData struct {
	// Identifies the resource
	ID          string    `json:"id" format:"uuid"`
	Connections []int64   `json:"connections"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	// Any of "INBOUND", "OUTBOUND".
	Directions []string                                                          `json:"directions"`
	EndDate    time.Time                                                         `json:"end_date" format:"date-time"`
	Filters    []LegacyReportingBatchDetailRecordMessagingListResponseDataFilter `json:"filters"`
	// List of messaging profile IDs
	Profiles   []string `json:"profiles" format:"uuid"`
	RecordType string   `json:"record_type"`
	// Any of "INCOMPLETE", "COMPLETED", "ERRORS".
	RecordTypes []string  `json:"record_types"`
	ReportName  string    `json:"report_name"`
	ReportURL   string    `json:"report_url"`
	StartDate   time.Time `json:"start_date" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Connections respjson.Field
		CreatedAt   respjson.Field
		Directions  respjson.Field
		EndDate     respjson.Field
		Filters     respjson.Field
		Profiles    respjson.Field
		RecordType  respjson.Field
		RecordTypes respjson.Field
		ReportName  respjson.Field
		ReportURL   respjson.Field
		StartDate   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingListResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type LegacyReportingBatchDetailRecordMessagingListResponseDataFilter struct {
	// Billing group UUID to filter by
	BillingGroup string `json:"billing_group"`
	// Called line identification (destination number)
	Cld string `json:"cld"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter string `json:"cld_filter"`
	// Calling line identification (caller ID)
	Cli string `json:"cli"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter string `json:"cli_filter"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType string `json:"filter_type"`
	// Tag name to filter by
	TagsList string `json:"tags_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingGroup respjson.Field
		Cld          respjson.Field
		CldFilter    respjson.Field
		Cli          respjson.Field
		CliFilter    respjson.Field
		FilterType   respjson.Field
		TagsList     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingListResponseDataFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingListResponseDataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
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
func (r LegacyReportingBatchDetailRecordMessagingListResponseMeta) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingDeleteResponse struct {
	Data LegacyReportingBatchDetailRecordMessagingDeleteResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *LegacyReportingBatchDetailRecordMessagingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingDeleteResponseData struct {
	// Identifies the resource
	ID          string    `json:"id" format:"uuid"`
	Connections []int64   `json:"connections"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	// Any of "INBOUND", "OUTBOUND".
	Directions []string                                                            `json:"directions"`
	EndDate    time.Time                                                           `json:"end_date" format:"date-time"`
	Filters    []LegacyReportingBatchDetailRecordMessagingDeleteResponseDataFilter `json:"filters"`
	// List of messaging profile IDs
	Profiles   []string `json:"profiles" format:"uuid"`
	RecordType string   `json:"record_type"`
	// Any of "INCOMPLETE", "COMPLETED", "ERRORS".
	RecordTypes []string  `json:"record_types"`
	ReportName  string    `json:"report_name"`
	ReportURL   string    `json:"report_url"`
	StartDate   time.Time `json:"start_date" format:"date-time"`
	// Any of "PENDING", "COMPLETE", "FAILED", "EXPIRED".
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Connections respjson.Field
		CreatedAt   respjson.Field
		Directions  respjson.Field
		EndDate     respjson.Field
		Filters     respjson.Field
		Profiles    respjson.Field
		RecordType  respjson.Field
		RecordTypes respjson.Field
		ReportName  respjson.Field
		ReportURL   respjson.Field
		StartDate   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingDeleteResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type LegacyReportingBatchDetailRecordMessagingDeleteResponseDataFilter struct {
	// Billing group UUID to filter by
	BillingGroup string `json:"billing_group"`
	// Called line identification (destination number)
	Cld string `json:"cld"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter string `json:"cld_filter"`
	// Calling line identification (caller ID)
	Cli string `json:"cli"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter string `json:"cli_filter"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType string `json:"filter_type"`
	// Tag name to filter by
	TagsList string `json:"tags_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingGroup respjson.Field
		Cld          respjson.Field
		CldFilter    respjson.Field
		Cli          respjson.Field
		CliFilter    respjson.Field
		FilterType   respjson.Field
		TagsList     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LegacyReportingBatchDetailRecordMessagingDeleteResponseDataFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *LegacyReportingBatchDetailRecordMessagingDeleteResponseDataFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LegacyReportingBatchDetailRecordMessagingNewParams struct {
	// End time in ISO format. Note: If end time includes the last 4 hours, some MDRs
	// might not appear in this report, due to wait time for downstream message
	// delivery confirmation
	EndTime time.Time `json:"end_time,required" format:"date-time"`
	// Start time in ISO format
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Whether to include message body in the report
	IncludeMessageBody param.Opt[bool] `json:"include_message_body,omitzero"`
	// Name of the report
	ReportName param.Opt[string] `json:"report_name,omitzero"`
	// Whether to select all managed accounts
	SelectAllManagedAccounts param.Opt[bool] `json:"select_all_managed_accounts,omitzero"`
	// Timezone for the report
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// List of connections to filter by
	Connections []int64 `json:"connections,omitzero"`
	// List of directions to filter by (Inbound = 1, Outbound = 2)
	Directions []int64 `json:"directions,omitzero"`
	// List of filters to apply
	Filters []LegacyReportingBatchDetailRecordMessagingNewParamsFilter `json:"filters,omitzero"`
	// List of managed accounts to include
	ManagedAccounts []string `json:"managed_accounts,omitzero" format:"uuid"`
	// List of messaging profile IDs to filter by
	Profiles []string `json:"profiles,omitzero" format:"uuid"`
	// List of record types to filter by (Complete = 1, Incomplete = 2, Errors = 3)
	RecordTypes []int64 `json:"record_types,omitzero"`
	paramObj
}

func (r LegacyReportingBatchDetailRecordMessagingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingBatchDetailRecordMessagingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingBatchDetailRecordMessagingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Query filter criteria. Note: The first filter object must specify filter_type as
// 'and'. You cannot follow an 'or' with another 'and'.
type LegacyReportingBatchDetailRecordMessagingNewParamsFilter struct {
	// Billing group UUID to filter by
	BillingGroup param.Opt[string] `json:"billing_group,omitzero"`
	// Called line identification (destination number)
	Cld param.Opt[string] `json:"cld,omitzero"`
	// Calling line identification (caller ID)
	Cli param.Opt[string] `json:"cli,omitzero"`
	// Tag name to filter by
	TagsList param.Opt[string] `json:"tags_list,omitzero"`
	// Filter type for CLD matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CldFilter string `json:"cld_filter,omitzero"`
	// Filter type for CLI matching
	//
	// Any of "contains", "starts_with", "ends_with".
	CliFilter string `json:"cli_filter,omitzero"`
	// Logical operator for combining filters
	//
	// Any of "and", "or".
	FilterType string `json:"filter_type,omitzero"`
	paramObj
}

func (r LegacyReportingBatchDetailRecordMessagingNewParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow LegacyReportingBatchDetailRecordMessagingNewParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LegacyReportingBatchDetailRecordMessagingNewParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LegacyReportingBatchDetailRecordMessagingNewParamsFilter](
		"cld_filter", "contains", "starts_with", "ends_with",
	)
	apijson.RegisterFieldValidator[LegacyReportingBatchDetailRecordMessagingNewParamsFilter](
		"cli_filter", "contains", "starts_with", "ends_with",
	)
	apijson.RegisterFieldValidator[LegacyReportingBatchDetailRecordMessagingNewParamsFilter](
		"filter_type", "and", "or",
	)
}
