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
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Background jobs performed over a batch of phone numbers
//
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

// Returns the status and details of the phone-number background job identified by
// `id`.
func (r *PhoneNumberJobService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberJobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("phone_numbers/jobs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns background jobs that operate on phone numbers. Filter by job type,
// target phone numbers, job status, or creation-time customer reference, and sort
// by creation time. Multiple phone-number or status values use OR semantics within
// that filter; different filter categories use AND semantics. Results include
// pagination metadata.
func (r *PhoneNumberJobService) List(ctx context.Context, query PhoneNumberJobListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[NumbersPhoneNumbersJob], err error) {
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

// Returns background jobs that operate on phone numbers. Filter by job type,
// target phone numbers, job status, or creation-time customer reference, and sort
// by creation time. Multiple phone-number or status values use OR semantics within
// that filter; different filter categories use AND semantics. Results include
// pagination metadata.
func (r *PhoneNumberJobService) ListAutoPaging(ctx context.Context, query PhoneNumberJobListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[NumbersPhoneNumbersJob] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Creates a new background job to delete a batch of numbers. At most one thousand
// numbers can be updated per API call.
func (r *PhoneNumberJobService) DeleteBatch(ctx context.Context, body PhoneNumberJobDeleteBatchParams, opts ...option.RequestOption) (res *PhoneNumberJobDeleteBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/jobs/delete_phone_numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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
	return res, err
}

// Creates a background job to update the emergency settings of a collection of
// phone numbers. At most one thousand numbers can be updated per API call.
func (r *PhoneNumberJobService) UpdateEmergencySettingsBatch(ctx context.Context, body PhoneNumberJobUpdateEmergencySettingsBatchParams, opts ...option.RequestOption) (res *PhoneNumberJobUpdateEmergencySettingsBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/jobs/update_emergency_settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type NumbersPhoneNumbersJob struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date indicating when the estimated time of completion of the
	// background job.
	Etc              time.Time                               `json:"etc" format:"date-time"`
	FailedOperations []NumbersPhoneNumbersJobFailedOperation `json:"failed_operations"`
	// Effective parameters saved when the job was created, not a snapshot of the
	// original HTTP request or current phone-number settings. Update jobs include the
	// saved update fields; emergency-settings jobs include the saved emergency fields;
	// delete jobs include only targets. Absent fields are omitted, and saved false,
	// null and empty values are preserved. Tags may be normalized to arrays;
	// filter-created jobs contain resolved target identifiers rather than the original
	// filters. Historical jobs with no saved arguments return an empty object.
	Parameters        NumbersPhoneNumbersJobParameters         `json:"parameters"`
	PendingOperations []NumbersPhoneNumbersJobPendingOperation `json:"pending_operations"`
	PhoneNumbers      []shared.PhoneNumbersJobPhoneNumber      `json:"phone_numbers"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates the completion status of the background update.
	//
	// Any of "pending", "in_progress", "completed", "failed", "expired".
	Status               NumbersPhoneNumbersJobStatus                `json:"status"`
	SuccessfulOperations []NumbersPhoneNumbersJobSuccessfulOperation `json:"successful_operations"`
	// Identifies the type of the background job.
	//
	// Any of "update_emergency_settings", "delete_phone_numbers",
	// "update_phone_numbers".
	Type NumbersPhoneNumbersJobType `json:"type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		Etc                  respjson.Field
		FailedOperations     respjson.Field
		Parameters           respjson.Field
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
func (r NumbersPhoneNumbersJob) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersPhoneNumbersJobFailedOperation struct {
	// The phone number's ID
	ID     string                   `json:"id"`
	Errors []NumbersError73c9788e53 `json:"errors"`
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
func (r NumbersPhoneNumbersJobFailedOperation) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobFailedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effective parameters saved when the job was created, not a snapshot of the
// original HTTP request or current phone-number settings. Update jobs include the
// saved update fields; emergency-settings jobs include the saved emergency fields;
// delete jobs include only targets. Absent fields are omitted, and saved false,
// null and empty values are preserved. Tags may be normalized to arrays;
// filter-created jobs contain resolved target identifiers rather than the original
// filters. Historical jobs with no saved arguments return an empty object.
type NumbersPhoneNumbersJobParameters struct {
	// Saved add_tags for an update_phone_numbers job. Values reflect creation-time
	// normalization.
	AddTags []string `json:"add_tags" api:"nullable"`
	// Saved billing_group_id supplied for an update_phone_numbers job.
	BillingGroupID string `json:"billing_group_id" api:"nullable"`
	// Saved connection_id supplied for an update_phone_numbers job.
	ConnectionID string `json:"connection_id" api:"nullable"`
	// Saved customer_reference supplied for an update_phone_numbers job.
	CustomerReference string `json:"customer_reference" api:"nullable"`
	// Saved deletion_lock_enabled supplied for an update_phone_numbers job.
	DeletionLockEnabled bool `json:"deletion_lock_enabled" api:"nullable"`
	// Saved emergency address identifier for an update_emergency_settings job. An
	// explicitly supplied null is preserved.
	EmergencyAddressID string `json:"emergency_address_id" api:"nullable"`
	// Saved emergency-services setting for an update_emergency_settings job.
	EmergencyEnabled bool `json:"emergency_enabled" api:"nullable"`
	// Saved external_pin supplied for an update_phone_numbers job.
	ExternalPin string `json:"external_pin" api:"nullable"`
	// Saved hd_voice_enabled supplied for an update_phone_numbers job.
	HDVoiceEnabled bool `json:"hd_voice_enabled" api:"nullable"`
	// Saved target phone numbers in E.164 format and/or identifiers. Jobs created
	// using filters contain the resolved identifiers, not the original filter
	// expression. Stored identifier types are preserved.
	PhoneNumbers []NumbersPhoneNumbersJobParametersPhoneNumbersUnion `json:"phone_numbers"`
	// Saved remove_tags for an update_phone_numbers job. Values reflect creation-time
	// normalization.
	RemoveTags []string `json:"remove_tags" api:"nullable"`
	// Saved tags for an update_phone_numbers job. Values reflect creation-time
	// normalization.
	Tags []string `json:"tags" api:"nullable"`
	// Saved voice settings for an update_phone_numbers job, using the public field
	// names and nesting. Only settings present at creation are returned.
	Voice NumbersPhoneNumbersJobParametersVoice `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddTags             respjson.Field
		BillingGroupID      respjson.Field
		ConnectionID        respjson.Field
		CustomerReference   respjson.Field
		DeletionLockEnabled respjson.Field
		EmergencyAddressID  respjson.Field
		EmergencyEnabled    respjson.Field
		ExternalPin         respjson.Field
		HDVoiceEnabled      respjson.Field
		PhoneNumbers        respjson.Field
		RemoveTags          respjson.Field
		Tags                respjson.Field
		Voice               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersPhoneNumbersJobParameters) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// NumbersPhoneNumbersJobParametersPhoneNumbersUnion contains all possible
// properties and values from [string], [int64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type NumbersPhoneNumbersJobParametersPhoneNumbersUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (u NumbersPhoneNumbersJobParametersPhoneNumbersUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u NumbersPhoneNumbersJobParametersPhoneNumbersUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u NumbersPhoneNumbersJobParametersPhoneNumbersUnion) RawJSON() string { return u.JSON.raw }

func (r *NumbersPhoneNumbersJobParametersPhoneNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Saved voice settings for an update_phone_numbers job, using the public field
// names and nesting. Only settings present at creation are returned.
type NumbersPhoneNumbersJobParametersVoice struct {
	CallForwarding NumbersPhoneNumbersJobParametersVoiceCallForwarding `json:"call_forwarding"`
	CallRecording  NumbersPhoneNumbersJobParametersVoiceCallRecording  `json:"call_recording"`
	// Saved caller_id_name_enabled setting.
	CallerIDNameEnabled bool                                             `json:"caller_id_name_enabled" api:"nullable"`
	CnamListing         NumbersPhoneNumbersJobParametersVoiceCnamListing `json:"cnam_listing"`
	// Saved inbound_call_screening setting.
	InboundCallScreening string                                             `json:"inbound_call_screening" api:"nullable"`
	MediaFeatures        NumbersPhoneNumbersJobParametersVoiceMediaFeatures `json:"media_features"`
	// Saved tech_prefix_enabled setting.
	TechPrefixEnabled bool `json:"tech_prefix_enabled" api:"nullable"`
	// Saved usage_payment_method setting.
	UsagePaymentMethod string `json:"usage_payment_method" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallForwarding       respjson.Field
		CallRecording        respjson.Field
		CallerIDNameEnabled  respjson.Field
		CnamListing          respjson.Field
		InboundCallScreening respjson.Field
		MediaFeatures        respjson.Field
		TechPrefixEnabled    respjson.Field
		UsagePaymentMethod   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersPhoneNumbersJobParametersVoice) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobParametersVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersPhoneNumbersJobParametersVoiceCallForwarding struct {
	// Saved call_forwarding_enabled setting.
	CallForwardingEnabled bool `json:"call_forwarding_enabled" api:"nullable"`
	// Saved forwarding_type setting.
	ForwardingType string `json:"forwarding_type" api:"nullable"`
	// Saved forwards_to setting.
	ForwardsTo string `json:"forwards_to" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallForwardingEnabled respjson.Field
		ForwardingType        respjson.Field
		ForwardsTo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersPhoneNumbersJobParametersVoiceCallForwarding) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobParametersVoiceCallForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersPhoneNumbersJobParametersVoiceCallRecording struct {
	// Saved inbound_call_recording_channels setting.
	InboundCallRecordingChannels string `json:"inbound_call_recording_channels" api:"nullable"`
	// Saved inbound_call_recording_enabled setting.
	InboundCallRecordingEnabled bool `json:"inbound_call_recording_enabled" api:"nullable"`
	// Saved inbound_call_recording_format setting.
	InboundCallRecordingFormat string `json:"inbound_call_recording_format" api:"nullable"`
	// Saved record_continue_on_hold setting.
	RecordContinueOnHold bool `json:"record_continue_on_hold" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundCallRecordingChannels respjson.Field
		InboundCallRecordingEnabled  respjson.Field
		InboundCallRecordingFormat   respjson.Field
		RecordContinueOnHold         respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersPhoneNumbersJobParametersVoiceCallRecording) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobParametersVoiceCallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersPhoneNumbersJobParametersVoiceCnamListing struct {
	// Saved cnam_listing_details setting.
	CnamListingDetails string `json:"cnam_listing_details" api:"nullable"`
	// Saved cnam_listing_enabled setting.
	CnamListingEnabled bool `json:"cnam_listing_enabled" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnamListingDetails respjson.Field
		CnamListingEnabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersPhoneNumbersJobParametersVoiceCnamListing) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobParametersVoiceCnamListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumbersPhoneNumbersJobParametersVoiceMediaFeatures struct {
	// Saved accept_any_rtp_packets_enabled setting.
	AcceptAnyRtpPacketsEnabled bool `json:"accept_any_rtp_packets_enabled" api:"nullable"`
	// Saved rtp_auto_adjust_enabled setting.
	RtpAutoAdjustEnabled bool `json:"rtp_auto_adjust_enabled" api:"nullable"`
	// Saved t38_fax_gateway_enabled setting.
	T38FaxGatewayEnabled bool `json:"t38_fax_gateway_enabled" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAnyRtpPacketsEnabled respjson.Field
		RtpAutoAdjustEnabled       respjson.Field
		T38FaxGatewayEnabled       respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumbersPhoneNumbersJobParametersVoiceMediaFeatures) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobParametersVoiceMediaFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The phone numbers pending confirmation on update results. Entries in this list
// are transient, and will be moved to either successful_operations or
// failed_operations once the processing is done.
type NumbersPhoneNumbersJobPendingOperation struct {
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
func (r NumbersPhoneNumbersJobPendingOperation) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobPendingOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the completion status of the background update.
type NumbersPhoneNumbersJobStatus string

const (
	NumbersPhoneNumbersJobStatusPending    NumbersPhoneNumbersJobStatus = "pending"
	NumbersPhoneNumbersJobStatusInProgress NumbersPhoneNumbersJobStatus = "in_progress"
	NumbersPhoneNumbersJobStatusCompleted  NumbersPhoneNumbersJobStatus = "completed"
	NumbersPhoneNumbersJobStatusFailed     NumbersPhoneNumbersJobStatus = "failed"
	NumbersPhoneNumbersJobStatusExpired    NumbersPhoneNumbersJobStatus = "expired"
)

// The phone numbers successfully updated.
type NumbersPhoneNumbersJobSuccessfulOperation struct {
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
func (r NumbersPhoneNumbersJobSuccessfulOperation) RawJSON() string { return r.JSON.raw }
func (r *NumbersPhoneNumbersJobSuccessfulOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the background job.
type NumbersPhoneNumbersJobType string

const (
	NumbersPhoneNumbersJobTypeUpdateEmergencySettings NumbersPhoneNumbersJobType = "update_emergency_settings"
	NumbersPhoneNumbersJobTypeDeletePhoneNumbers      NumbersPhoneNumbersJobType = "delete_phone_numbers"
	NumbersPhoneNumbersJobTypeUpdatePhoneNumbers      NumbersPhoneNumbersJobType = "update_phone_numbers"
)

type PhoneNumberJobGetResponse struct {
	Data NumbersPhoneNumbersJob `json:"data"`
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
	Data NumbersPhoneNumbersJob `json:"data"`
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
	Data NumbersPhoneNumbersJob `json:"data"`
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
	Data NumbersPhoneNumbersJob `json:"data"`
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
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[type],
	// filter[phone_number], filter[phone_number][], filter[status][],
	// filter[customer_reference]
	Filter PhoneNumberJobListParamsFilter `query:"filter,omitzero" json:"-"`
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

// Consolidated filter parameter (deepObject style). Originally: filter[type],
// filter[phone_number], filter[phone_number][], filter[status][],
// filter[customer_reference]
type PhoneNumberJobListParamsFilter struct {
	// Exact, case-sensitive match for the customer reference supplied when the job was
	// created. Only jobs created within the last 90 days are returned, including for
	// accounts otherwise exempt from job retention limits. Searches the stored job
	// reference, not the current reference on a phone number. Multiple jobs may share
	// a reference. Combines with the other filters and account ownership. Supply a
	// single non-blank string of up to 250 characters, without null characters, using
	// `filter[customer_reference]=yourcustomerrefhere`. Whitespace is preserved; blank
	// strings, arrays, and objects return HTTP 422. An invalid filter value returns
	// HTTP 422. The literal string `null` is matched as text.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Returns jobs that targeted any of the supplied account-owned phone numbers.
	// Values beginning with `+` must contain 1 to 20 digits after the plus sign. The
	// 10-value limit is enforced before duplicate values are removed. Unmatched or
	// non-account-owned identifiers return an empty result. Phone-number filtering
	// must be enabled for the account.
	PhoneNumber PhoneNumberJobListParamsFilterPhoneNumberUnion `query:"phone_number,omitzero" json:"-"`
	// Returns jobs with any of the supplied statuses. Use repeated `filter[status][]`
	// parameters; scalar and comma-separated status values are not accepted.
	//
	// Any of "pending", "in_progress", "completed", "failed", "expired".
	Status []string `query:"status,omitzero" json:"-"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PhoneNumberJobListParamsFilterPhoneNumberUnion struct {
	OfString      param.Opt[string] `query:",omitzero,inline"`
	OfStringArray []string          `query:",omitzero,inline"`
	paramUnion
}

func (u *PhoneNumberJobListParamsFilterPhoneNumberUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Specifies the sort order for results. If not given, results are sorted by
// created_at in descending order.
type PhoneNumberJobListParamsSort string

const (
	PhoneNumberJobListParamsSortCreatedAt PhoneNumberJobListParamsSort = "created_at"
)

type PhoneNumberJobDeleteBatchParams struct {
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
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
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	// Identifies the billing group associated with the phone number.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero"`
	// Identifies the connection associated with the phone number.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// Sets the customer reference on the selected phone numbers and retains the
	// supplied value on the job. Recover the job without its ID using GET
	// /phone_numbers/jobs with filter[customer_reference] within 90 days of job
	// creation. This body field updates the numbers; the separate
	// filter[customer_reference] query parameter selects numbers by their current
	// reference.
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
	EmergencyEnabled bool     `json:"emergency_enabled" api:"required"`
	PhoneNumbers     []string `json:"phone_numbers,omitzero" api:"required"`
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
