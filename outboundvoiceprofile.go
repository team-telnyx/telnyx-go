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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// OutboundVoiceProfileService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOutboundVoiceProfileService] method instead.
type OutboundVoiceProfileService struct {
	Options []option.RequestOption
}

// NewOutboundVoiceProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOutboundVoiceProfileService(opts ...option.RequestOption) (r OutboundVoiceProfileService) {
	r = OutboundVoiceProfileService{}
	r.Options = opts
	return
}

// Create an outbound voice profile.
func (r *OutboundVoiceProfileService) New(ctx context.Context, body OutboundVoiceProfileNewParams, opts ...option.RequestOption) (res *OutboundVoiceProfileNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "outbound_voice_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of an existing outbound voice profile.
func (r *OutboundVoiceProfileService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OutboundVoiceProfileGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("outbound_voice_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing outbound voice profile.
func (r *OutboundVoiceProfileService) Update(ctx context.Context, id string, body OutboundVoiceProfileUpdateParams, opts ...option.RequestOption) (res *OutboundVoiceProfileUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("outbound_voice_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get all outbound voice profiles belonging to the user that match the given
// filters.
func (r *OutboundVoiceProfileService) List(ctx context.Context, query OutboundVoiceProfileListParams, opts ...option.RequestOption) (res *OutboundVoiceProfileListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "outbound_voice_profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing outbound voice profile.
func (r *OutboundVoiceProfileService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OutboundVoiceProfileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("outbound_voice_profiles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type OutboundCallRecording struct {
	// When call_recording_type is 'by_caller_phone_number', only outbound calls using
	// one of these numbers will be recorded. Numbers must be specified in E164 format.
	CallRecordingCallerPhoneNumbers []string `json:"call_recording_caller_phone_numbers"`
	// When using 'dual' channels, the final audio file will be a stereo recording with
	// the first leg on channel A, and the rest on channel B.
	//
	// Any of "single", "dual".
	CallRecordingChannels OutboundCallRecordingCallRecordingChannels `json:"call_recording_channels"`
	// The audio file format for calls being recorded.
	//
	// Any of "wav", "mp3".
	CallRecordingFormat OutboundCallRecordingCallRecordingFormat `json:"call_recording_format"`
	// Specifies which calls are recorded.
	//
	// Any of "all", "none", "by_caller_phone_number".
	CallRecordingType OutboundCallRecordingCallRecordingType `json:"call_recording_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallRecordingCallerPhoneNumbers respjson.Field
		CallRecordingChannels           respjson.Field
		CallRecordingFormat             respjson.Field
		CallRecordingType               respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundCallRecording) RawJSON() string { return r.JSON.raw }
func (r *OutboundCallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OutboundCallRecording to a OutboundCallRecordingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OutboundCallRecordingParam.Overrides()
func (r OutboundCallRecording) ToParam() OutboundCallRecordingParam {
	return param.Override[OutboundCallRecordingParam](json.RawMessage(r.RawJSON()))
}

// When using 'dual' channels, the final audio file will be a stereo recording with
// the first leg on channel A, and the rest on channel B.
type OutboundCallRecordingCallRecordingChannels string

const (
	OutboundCallRecordingCallRecordingChannelsSingle OutboundCallRecordingCallRecordingChannels = "single"
	OutboundCallRecordingCallRecordingChannelsDual   OutboundCallRecordingCallRecordingChannels = "dual"
)

// The audio file format for calls being recorded.
type OutboundCallRecordingCallRecordingFormat string

const (
	OutboundCallRecordingCallRecordingFormatWav OutboundCallRecordingCallRecordingFormat = "wav"
	OutboundCallRecordingCallRecordingFormatMP3 OutboundCallRecordingCallRecordingFormat = "mp3"
)

// Specifies which calls are recorded.
type OutboundCallRecordingCallRecordingType string

const (
	OutboundCallRecordingCallRecordingTypeAll                 OutboundCallRecordingCallRecordingType = "all"
	OutboundCallRecordingCallRecordingTypeNone                OutboundCallRecordingCallRecordingType = "none"
	OutboundCallRecordingCallRecordingTypeByCallerPhoneNumber OutboundCallRecordingCallRecordingType = "by_caller_phone_number"
)

type OutboundCallRecordingParam struct {
	// When call_recording_type is 'by_caller_phone_number', only outbound calls using
	// one of these numbers will be recorded. Numbers must be specified in E164 format.
	CallRecordingCallerPhoneNumbers []string `json:"call_recording_caller_phone_numbers,omitzero"`
	// When using 'dual' channels, the final audio file will be a stereo recording with
	// the first leg on channel A, and the rest on channel B.
	//
	// Any of "single", "dual".
	CallRecordingChannels OutboundCallRecordingCallRecordingChannels `json:"call_recording_channels,omitzero"`
	// The audio file format for calls being recorded.
	//
	// Any of "wav", "mp3".
	CallRecordingFormat OutboundCallRecordingCallRecordingFormat `json:"call_recording_format,omitzero"`
	// Specifies which calls are recorded.
	//
	// Any of "all", "none", "by_caller_phone_number".
	CallRecordingType OutboundCallRecordingCallRecordingType `json:"call_recording_type,omitzero"`
	paramObj
}

func (r OutboundCallRecordingParam) MarshalJSON() (data []byte, err error) {
	type shadow OutboundCallRecordingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutboundCallRecordingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfile struct {
	// A user-supplied name to help with organization.
	Name string `json:"name,required"`
	// Identifies the resource.
	ID string `json:"id"`
	// The ID of the billing group associated with the outbound proflile. Defaults to
	// null (for no group assigned).
	BillingGroupID string                `json:"billing_group_id,nullable" format:"uuid"`
	CallRecording  OutboundCallRecording `json:"call_recording"`
	// Must be no more than your global concurrent call limit. Null means no limit.
	ConcurrentCallLimit int64 `json:"concurrent_call_limit,nullable"`
	// Amount of connections associated with this outbound voice profile.
	ConnectionsCount int64 `json:"connections_count"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// The maximum amount of usage charges, in USD, you want Telnyx to allow on this
	// outbound voice profile in a day before disallowing new calls.
	DailySpendLimit string `json:"daily_spend_limit"`
	// Specifies whether to enforce the daily_spend_limit on this outbound voice
	// profile.
	DailySpendLimitEnabled bool `json:"daily_spend_limit_enabled"`
	// Specifies whether the outbound voice profile can be used. Disabled profiles will
	// result in outbound calls being blocked for the associated Connections.
	Enabled bool `json:"enabled"`
	// Maximum rate (price per minute) for a Destination to be allowed when making
	// outbound calls.
	MaxDestinationRate float64 `json:"max_destination_rate"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates the coverage of the termination regions.
	//
	// Any of "global".
	ServicePlan ServicePlan `json:"service_plan"`
	Tags        []string    `json:"tags"`
	// Specifies the type of traffic allowed in this profile.
	//
	// Any of "conversational".
	TrafficType TrafficType `json:"traffic_type"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// Setting for how costs for outbound profile are calculated.
	//
	// Any of "rate-deck".
	UsagePaymentMethod UsagePaymentMethod `json:"usage_payment_method"`
	// The list of destinations you want to be able to call using this outbound voice
	// profile formatted in alpha2.
	WhitelistedDestinations []string `json:"whitelisted_destinations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                    respjson.Field
		ID                      respjson.Field
		BillingGroupID          respjson.Field
		CallRecording           respjson.Field
		ConcurrentCallLimit     respjson.Field
		ConnectionsCount        respjson.Field
		CreatedAt               respjson.Field
		DailySpendLimit         respjson.Field
		DailySpendLimitEnabled  respjson.Field
		Enabled                 respjson.Field
		MaxDestinationRate      respjson.Field
		RecordType              respjson.Field
		ServicePlan             respjson.Field
		Tags                    respjson.Field
		TrafficType             respjson.Field
		UpdatedAt               respjson.Field
		UsagePaymentMethod      respjson.Field
		WhitelistedDestinations respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundVoiceProfile) RawJSON() string { return r.JSON.raw }
func (r *OutboundVoiceProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the coverage of the termination regions.
type ServicePlan string

const (
	ServicePlanGlobal ServicePlan = "global"
)

// Specifies the type of traffic allowed in this profile.
type TrafficType string

const (
	TrafficTypeConversational TrafficType = "conversational"
)

// Setting for how costs for outbound profile are calculated.
type UsagePaymentMethod string

const (
	UsagePaymentMethodRateDeck UsagePaymentMethod = "rate-deck"
)

type OutboundVoiceProfileNewResponse struct {
	Data OutboundVoiceProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundVoiceProfileNewResponse) RawJSON() string { return r.JSON.raw }
func (r *OutboundVoiceProfileNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileGetResponse struct {
	Data OutboundVoiceProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundVoiceProfileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OutboundVoiceProfileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileUpdateResponse struct {
	Data OutboundVoiceProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundVoiceProfileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *OutboundVoiceProfileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileListResponse struct {
	Data []OutboundVoiceProfile `json:"data"`
	Meta PaginationMeta         `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundVoiceProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *OutboundVoiceProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileDeleteResponse struct {
	Data OutboundVoiceProfile `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundVoiceProfileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OutboundVoiceProfileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileNewParams struct {
	// A user-supplied name to help with organization.
	Name string `json:"name,required"`
	// The ID of the billing group associated with the outbound proflile. Defaults to
	// null (for no group assigned).
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero" format:"uuid"`
	// Must be no more than your global concurrent call limit. Null means no limit.
	ConcurrentCallLimit param.Opt[int64] `json:"concurrent_call_limit,omitzero"`
	// The maximum amount of usage charges, in USD, you want Telnyx to allow on this
	// outbound voice profile in a day before disallowing new calls.
	DailySpendLimit param.Opt[string] `json:"daily_spend_limit,omitzero"`
	// Specifies whether to enforce the daily_spend_limit on this outbound voice
	// profile.
	DailySpendLimitEnabled param.Opt[bool] `json:"daily_spend_limit_enabled,omitzero"`
	// Specifies whether the outbound voice profile can be used. Disabled profiles will
	// result in outbound calls being blocked for the associated Connections.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Maximum rate (price per minute) for a Destination to be allowed when making
	// outbound calls.
	MaxDestinationRate param.Opt[float64]         `json:"max_destination_rate,omitzero"`
	CallRecording      OutboundCallRecordingParam `json:"call_recording,omitzero"`
	// Indicates the coverage of the termination regions.
	//
	// Any of "global".
	ServicePlan ServicePlan `json:"service_plan,omitzero"`
	Tags        []string    `json:"tags,omitzero"`
	// Specifies the type of traffic allowed in this profile.
	//
	// Any of "conversational".
	TrafficType TrafficType `json:"traffic_type,omitzero"`
	// Setting for how costs for outbound profile are calculated.
	//
	// Any of "rate-deck".
	UsagePaymentMethod UsagePaymentMethod `json:"usage_payment_method,omitzero"`
	// The list of destinations you want to be able to call using this outbound voice
	// profile formatted in alpha2.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitzero"`
	paramObj
}

func (r OutboundVoiceProfileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OutboundVoiceProfileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutboundVoiceProfileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileUpdateParams struct {
	// A user-supplied name to help with organization.
	Name string `json:"name,required"`
	// The ID of the billing group associated with the outbound proflile. Defaults to
	// null (for no group assigned).
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero" format:"uuid"`
	// Must be no more than your global concurrent call limit. Null means no limit.
	ConcurrentCallLimit param.Opt[int64] `json:"concurrent_call_limit,omitzero"`
	// The maximum amount of usage charges, in USD, you want Telnyx to allow on this
	// outbound voice profile in a day before disallowing new calls.
	DailySpendLimit param.Opt[string] `json:"daily_spend_limit,omitzero"`
	// Specifies whether to enforce the daily_spend_limit on this outbound voice
	// profile.
	DailySpendLimitEnabled param.Opt[bool] `json:"daily_spend_limit_enabled,omitzero"`
	// Specifies whether the outbound voice profile can be used. Disabled profiles will
	// result in outbound calls being blocked for the associated Connections.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Maximum rate (price per minute) for a Destination to be allowed when making
	// outbound calls.
	MaxDestinationRate param.Opt[float64]         `json:"max_destination_rate,omitzero"`
	CallRecording      OutboundCallRecordingParam `json:"call_recording,omitzero"`
	// Indicates the coverage of the termination regions.
	//
	// Any of "global".
	ServicePlan ServicePlan `json:"service_plan,omitzero"`
	Tags        []string    `json:"tags,omitzero"`
	// Specifies the type of traffic allowed in this profile.
	//
	// Any of "conversational".
	TrafficType TrafficType `json:"traffic_type,omitzero"`
	// Setting for how costs for outbound profile are calculated.
	//
	// Any of "rate-deck".
	UsagePaymentMethod UsagePaymentMethod `json:"usage_payment_method,omitzero"`
	// The list of destinations you want to be able to call using this outbound voice
	// profile formatted in alpha2.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitzero"`
	paramObj
}

func (r OutboundVoiceProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OutboundVoiceProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutboundVoiceProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundVoiceProfileListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[name][contains]
	Filter OutboundVoiceProfileListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page OutboundVoiceProfileListParamsPage `query:"page,omitzero" json:"-"`
	// Specifies the sort order for results. By default sorting direction is ascending.
	// To have the results sorted in descending order add the <code>-</code>
	// prefix.<br/><br/> That is: <ul>
	//
	//	<li>
	//	  <code>name</code>: sorts the result by the
	//	  <code>name</code> field in ascending order.
	//	</li>
	//
	//	<li>
	//	  <code>-name</code>: sorts the result by the
	//	  <code>name</code> field in descending order.
	//	</li>
	//
	// </ul> <br/>
	//
	// Any of "enabled", "-enabled", "created_at", "-created_at", "name", "-name",
	// "service_plan", "-service_plan", "traffic_type", "-traffic_type",
	// "usage_payment_method", "-usage_payment_method".
	Sort OutboundVoiceProfileListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OutboundVoiceProfileListParams]'s query parameters as
// `url.Values`.
func (r OutboundVoiceProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[name][contains]
type OutboundVoiceProfileListParamsFilter struct {
	// Name filtering operations
	Name OutboundVoiceProfileListParamsFilterName `query:"name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OutboundVoiceProfileListParamsFilter]'s query parameters as
// `url.Values`.
func (r OutboundVoiceProfileListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Name filtering operations
type OutboundVoiceProfileListParamsFilterName struct {
	// Optional filter on outbound voice profile name.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OutboundVoiceProfileListParamsFilterName]'s query
// parameters as `url.Values`.
func (r OutboundVoiceProfileListParamsFilterName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type OutboundVoiceProfileListParamsPage struct {
	// The page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OutboundVoiceProfileListParamsPage]'s query parameters as
// `url.Values`.
func (r OutboundVoiceProfileListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. By default sorting direction is ascending.
// To have the results sorted in descending order add the <code>-</code>
// prefix.<br/><br/> That is: <ul>
//
//	<li>
//	  <code>name</code>: sorts the result by the
//	  <code>name</code> field in ascending order.
//	</li>
//
//	<li>
//	  <code>-name</code>: sorts the result by the
//	  <code>name</code> field in descending order.
//	</li>
//
// </ul> <br/>
type OutboundVoiceProfileListParamsSort string

const (
	OutboundVoiceProfileListParamsSortEnabled                OutboundVoiceProfileListParamsSort = "enabled"
	OutboundVoiceProfileListParamsSortEnabledDesc            OutboundVoiceProfileListParamsSort = "-enabled"
	OutboundVoiceProfileListParamsSortCreatedAt              OutboundVoiceProfileListParamsSort = "created_at"
	OutboundVoiceProfileListParamsSortCreatedAtDesc          OutboundVoiceProfileListParamsSort = "-created_at"
	OutboundVoiceProfileListParamsSortName                   OutboundVoiceProfileListParamsSort = "name"
	OutboundVoiceProfileListParamsSortNameDesc               OutboundVoiceProfileListParamsSort = "-name"
	OutboundVoiceProfileListParamsSortServicePlan            OutboundVoiceProfileListParamsSort = "service_plan"
	OutboundVoiceProfileListParamsSortServicePlanDesc        OutboundVoiceProfileListParamsSort = "-service_plan"
	OutboundVoiceProfileListParamsSortTrafficType            OutboundVoiceProfileListParamsSort = "traffic_type"
	OutboundVoiceProfileListParamsSortTrafficTypeDesc        OutboundVoiceProfileListParamsSort = "-traffic_type"
	OutboundVoiceProfileListParamsSortUsagePaymentMethod     OutboundVoiceProfileListParamsSort = "usage_payment_method"
	OutboundVoiceProfileListParamsSortUsagePaymentMethodDesc OutboundVoiceProfileListParamsSort = "-usage_payment_method"
)
