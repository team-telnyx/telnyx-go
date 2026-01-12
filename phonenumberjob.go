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

// PhoneNumberJobService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberJobService] method instead.
type PhoneNumberJobService struct {
	Options []option.RequestOption
}

// NewPhoneNumberJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhoneNumberJobService(opts ...option.RequestOption) (r PhoneNumberJobService) {
	r = PhoneNumberJobService{}
	r.Options = opts
	return
}

// Retrieve a phone numbers job
func (r *PhoneNumberJobService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberJobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/jobs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the phone numbers jobs
func (r *PhoneNumberJobService) List(ctx context.Context, query PhoneNumberJobListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PhoneNumbersJob], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "phone_numbers/jobs"
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

// Lists the phone numbers jobs
func (r *PhoneNumberJobService) ListAutoPaging(ctx context.Context, query PhoneNumberJobListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PhoneNumbersJob] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Creates a new background job to delete a batch of numbers. At most one thousand
// numbers can be updated per API call.
func (r *PhoneNumberJobService) DeleteBatch(ctx context.Context, body PhoneNumberJobDeleteBatchParams, opts ...option.RequestOption) (res *PhoneNumberJobDeleteBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/jobs/delete_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a new background job to update a batch of numbers. At most one thousand
// numbers can be updated per API call. At least one of the updateable fields must
// be submitted. IMPORTANT: You must either specify filters (using the filter
// parameters) or specific phone numbers (using the phone_numbers parameter in the
// request body). If you specify filters, ALL phone numbers that match the given
// filters (up to 1000 at a time) will be updated. If you want to update only
// specific numbers, you must use the phone_numbers parameter in the request body.
// When using the phone_numbers parameter, ensure you follow the correct format as
// shown in the example (either phone number IDs or phone numbers in E164 format).
func (r *PhoneNumberJobService) UpdateBatch(ctx context.Context, params PhoneNumberJobUpdateBatchParams, opts ...option.RequestOption) (res *PhoneNumberJobUpdateBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/jobs/update_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Creates a background job to update the emergency settings of a collection of
// phone numbers. At most one thousand numbers can be updated per API call.
func (r *PhoneNumberJobService) UpdateEmergencySettingsBatch(ctx context.Context, body PhoneNumberJobUpdateEmergencySettingsBatchParams, opts ...option.RequestOption) (res *PhoneNumberJobUpdateEmergencySettingsBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/jobs/update_emergency_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PhoneNumbersJob struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date indicating when the estimated time of completion of the
	// background job.
	Etc               time.Time                         `json:"etc" format:"date-time"`
	FailedOperations  []PhoneNumbersJobFailedOperation  `json:"failed_operations"`
	PendingOperations []PhoneNumbersJobPendingOperation `json:"pending_operations"`
	PhoneNumbers      []PhoneNumbersJobPhoneNumber      `json:"phone_numbers"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates the completion status of the background update.
	//
	// Any of "pending", "in_progress", "completed", "failed", "expired".
	Status               PhoneNumbersJobStatus                `json:"status"`
	SuccessfulOperations []PhoneNumbersJobSuccessfulOperation `json:"successful_operations"`
	// Identifies the type of the background job.
	//
	// Any of "update_emergency_settings", "delete_phone_numbers",
	// "update_phone_numbers".
	Type PhoneNumbersJobType `json:"type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		Etc                  respjson.Field
		FailedOperations     respjson.Field
		PendingOperations    respjson.Field
		PhoneNumbers         respjson.Field
		RecordType           respjson.Field
		Status               respjson.Field
		SuccessfulOperations respjson.Field
		Type                 respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersJob) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumbersJobFailedOperation struct {
	// The phone number's ID
	ID     string     `json:"id"`
	Errors []JobError `json:"errors"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Errors      respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersJobFailedOperation) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersJobFailedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The phone numbers pending confirmation on update results. Entries in this list
// are transient, and will be moved to either successful_operations or
// failed_operations once the processing is done.
type PhoneNumbersJobPendingOperation struct {
	// The phone number's ID
	ID string `json:"id"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersJobPendingOperation) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersJobPendingOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The unique phone numbers given as arguments in the job creation.
type PhoneNumbersJobPhoneNumber struct {
	// The phone number's ID
	ID string `json:"id"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersJobPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersJobPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the completion status of the background update.
type PhoneNumbersJobStatus string

const (
	PhoneNumbersJobStatusPending    PhoneNumbersJobStatus = "pending"
	PhoneNumbersJobStatusInProgress PhoneNumbersJobStatus = "in_progress"
	PhoneNumbersJobStatusCompleted  PhoneNumbersJobStatus = "completed"
	PhoneNumbersJobStatusFailed     PhoneNumbersJobStatus = "failed"
	PhoneNumbersJobStatusExpired    PhoneNumbersJobStatus = "expired"
)

// The phone numbers successfully updated.
type PhoneNumbersJobSuccessfulOperation struct {
	// The phone number's ID
	ID string `json:"id"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumbersJobSuccessfulOperation) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumbersJobSuccessfulOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the background job.
type PhoneNumbersJobType string

const (
	PhoneNumbersJobTypeUpdateEmergencySettings PhoneNumbersJobType = "update_emergency_settings"
	PhoneNumbersJobTypeDeletePhoneNumbers      PhoneNumbersJobType = "delete_phone_numbers"
	PhoneNumbersJobTypeUpdatePhoneNumbers      PhoneNumbersJobType = "update_phone_numbers"
)

type PhoneNumberJobGetResponse struct {
	Data PhoneNumbersJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberJobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberJobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberJobDeleteBatchResponse struct {
	Data PhoneNumbersJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberJobDeleteBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberJobDeleteBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberJobUpdateBatchResponse struct {
	Data PhoneNumbersJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberJobUpdateBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberJobUpdateBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberJobUpdateEmergencySettingsBatchResponse struct {
	Data PhoneNumbersJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberJobUpdateEmergencySettingsBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberJobUpdateEmergencySettingsBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberJobListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[type]
	Filter PhoneNumberJobListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PhoneNumberJobListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at".
	Sort PhoneNumberJobListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberJobListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberJobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[type]
type PhoneNumberJobListParamsFilter struct {
	// Identifies the type of the background job.
	//
	// Any of "update_emergency_settings", "delete_phone_numbers",
	// "update_phone_numbers".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberJobListParamsFilter]'s query parameters as
// `url.Values`.
func (r PhoneNumberJobListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PhoneNumberJobListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberJobListParamsPage]'s query parameters as
// `url.Values`.
func (r PhoneNumberJobListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// created_at in descending order.
type PhoneNumberJobListParamsSort string

const (
	PhoneNumberJobListParamsSortCreatedAt PhoneNumberJobListParamsSort = "created_at"
)

type PhoneNumberJobDeleteBatchParams struct {
	PhoneNumbers []string `json:"phone_numbers,omitzero,required"`
	paramObj
}

func (r PhoneNumberJobDeleteBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberJobDeleteBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberJobDeleteBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberJobUpdateBatchParams struct {
	// Array of phone number ids and/or phone numbers in E164 format to update. This
	// parameter is required if no filter parameters are provided. If you want to
	// update specific numbers rather than all numbers matching a filter, you must use
	// this parameter. Each item must be either a valid phone number ID or a phone
	// number in E164 format (e.g., '+13127367254').
	PhoneNumbers []string `json:"phone_numbers,omitzero,required"`
	// Identifies the billing group associated with the phone number.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero"`
	// Identifies the connection associated with the phone number.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// A customer reference string for customer look ups.
	CustomerReference param.Opt[string] `json:"customer_reference,omitzero"`
	// Indicates whether to enable or disable the deletion lock on each phone number.
	// When enabled, this prevents the phone number from being deleted via the API or
	// Telnyx portal.
	DeletionLockEnabled param.Opt[bool] `json:"deletion_lock_enabled,omitzero"`
	// If someone attempts to port your phone number away from Telnyx and your phone
	// number has an external PIN set, we will attempt to verify that you provided the
	// correct external PIN to the winning carrier. Note that not all carriers
	// cooperate with this security mechanism.
	ExternalPin param.Opt[string] `json:"external_pin,omitzero"`
	// Indicates whether to enable or disable HD Voice on each phone number. HD Voice
	// is a paid feature and may not be available for all phone numbers, more details
	// about it can be found in the Telnyx support documentation.
	HDVoiceEnabled param.Opt[bool] `json:"hd_voice_enabled,omitzero"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[has_bundle], filter[tag], filter[connection_id], filter[phone_number],
	// filter[status], filter[voice.connection_name],
	// filter[voice.usage_payment_method], filter[billing_group_id],
	// filter[emergency_address_id], filter[customer_reference]
	Filter PhoneNumberJobUpdateBatchParamsFilter `query:"filter,omitzero" json:"-"`
	// A list of user-assigned tags to help organize phone numbers.
	Tags  []string                 `json:"tags,omitzero"`
	Voice UpdateVoiceSettingsParam `json:"voice,omitzero"`
	paramObj
}

func (r PhoneNumberJobUpdateBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberJobUpdateBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberJobUpdateBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [PhoneNumberJobUpdateBatchParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberJobUpdateBatchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[has_bundle], filter[tag], filter[connection_id], filter[phone_number],
// filter[status], filter[voice.connection_name],
// filter[voice.usage_payment_method], filter[billing_group_id],
// filter[emergency_address_id], filter[customer_reference]
type PhoneNumberJobUpdateBatchParamsFilter struct {
	// Filter by the billing_group_id associated with phone numbers. To filter to only
	// phone numbers that have no billing group associated them, set the value of this
	// filter to the string 'null'.
	BillingGroupID param.Opt[string] `query:"billing_group_id,omitzero" json:"-"`
	// Filter by connection_id.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Filter numbers via the customer_reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter by the emergency_address_id associated with phone numbers. To filter only
	// phone numbers that have no emergency address associated with them, set the value
	// of this filter to the string 'null'.
	EmergencyAddressID param.Opt[string] `query:"emergency_address_id,omitzero" json:"-"`
	// Filter by phone number that have bundles.
	HasBundle param.Opt[string] `query:"has_bundle,omitzero" json:"-"`
	// Filter by phone number. Requires at least three digits. Non-numerical characters
	// will result in no values being returned.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter by phone number tags.
	Tag param.Opt[string] `query:"tag,omitzero" json:"-"`
	// Filter by phone number status.
	//
	// Any of "purchase-pending", "purchase-failed", "port-pending", "active",
	// "deleted", "port-failed", "emergency-only", "ported-out", "port-out-pending".
	Status string `query:"status,omitzero" json:"-"`
	// Filter by voice connection name pattern matching.
	VoiceConnectionName PhoneNumberJobUpdateBatchParamsFilterVoiceConnectionName `query:"voice.connection_name,omitzero" json:"-"`
	// Filter by usage_payment_method.
	//
	// Any of "pay-per-minute", "channel".
	VoiceUsagePaymentMethod string `query:"voice.usage_payment_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberJobUpdateBatchParamsFilter]'s query parameters
// as `url.Values`.
func (r PhoneNumberJobUpdateBatchParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by voice connection name pattern matching.
type PhoneNumberJobUpdateBatchParamsFilterVoiceConnectionName struct {
	// Filter contains connection name. Requires at least three characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	// Filter ends with connection name. Requires at least three characters.
	EndsWith param.Opt[string] `query:"ends_with,omitzero" json:"-"`
	// Filter by connection name.
	Eq param.Opt[string] `query:"eq,omitzero" json:"-"`
	// Filter starts with connection name. Requires at least three characters.
	StartsWith param.Opt[string] `query:"starts_with,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberJobUpdateBatchParamsFilterVoiceConnectionName]'s
// query parameters as `url.Values`.
func (r PhoneNumberJobUpdateBatchParamsFilterVoiceConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PhoneNumberJobUpdateEmergencySettingsBatchParams struct {
	// Indicates whether to enable or disable emergency services on the numbers.
	EmergencyEnabled bool     `json:"emergency_enabled,required"`
	PhoneNumbers     []string `json:"phone_numbers,omitzero,required"`
	// Identifies the address to be used with emergency services. Required if
	// emergency_enabled is true, must be null or omitted if emergency_enabled is
	// false.
	EmergencyAddressID param.Opt[string] `json:"emergency_address_id,omitzero"`
	paramObj
}

func (r PhoneNumberJobUpdateEmergencySettingsBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberJobUpdateEmergencySettingsBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberJobUpdateEmergencySettingsBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
