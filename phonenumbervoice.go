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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Configure your phone numbers
//
// PhoneNumberVoiceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberVoiceService] method instead.
type PhoneNumberVoiceService struct {
	Options []option.RequestOption
}

// NewPhoneNumberVoiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberVoiceService(opts ...option.RequestOption) (r PhoneNumberVoiceService) {
	r = PhoneNumberVoiceService{}
	r.Options = opts
	return
}

// Retrieve a phone number with voice settings
func (r *PhoneNumberVoiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberVoiceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/voice", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a phone number with voice settings
func (r *PhoneNumberVoiceService) Update(ctx context.Context, id string, body PhoneNumberVoiceUpdateParams, opts ...option.RequestOption) (res *PhoneNumberVoiceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/voice", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List phone numbers with voice settings
func (r *PhoneNumberVoiceService) List(ctx context.Context, query PhoneNumberVoiceListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PhoneNumberWithVoiceSettings], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "phone_numbers/voice"
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

// List phone numbers with voice settings
func (r *PhoneNumberVoiceService) ListAutoPaging(ctx context.Context, query PhoneNumberVoiceListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PhoneNumberWithVoiceSettings] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// The call forwarding settings for a phone number.
type CallForwarding struct {
	// Indicates if call forwarding will be enabled for this number if forwards_to and
	// forwarding_type are filled in. Defaults to true for backwards compatibility with
	// APIV1 use of numbers endpoints.
	CallForwardingEnabled bool `json:"call_forwarding_enabled"`
	// Call forwarding type. 'forwards_to' must be set for this to have an effect.
	//
	// Any of "always", "on-failure".
	ForwardingType CallForwardingForwardingType `json:"forwarding_type"`
	// The phone number to which inbound calls to this number are forwarded. Inbound
	// calls will not be forwarded if this field is left blank. If set, must be a
	// +E.164-formatted phone number.
	ForwardsTo string `json:"forwards_to"`
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
func (r CallForwarding) RawJSON() string { return r.JSON.raw }
func (r *CallForwarding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CallForwarding to a CallForwardingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CallForwardingParam.Overrides()
func (r CallForwarding) ToParam() CallForwardingParam {
	return param.Override[CallForwardingParam](json.RawMessage(r.RawJSON()))
}

// Call forwarding type. 'forwards_to' must be set for this to have an effect.
type CallForwardingForwardingType string

const (
	CallForwardingForwardingTypeAlways    CallForwardingForwardingType = "always"
	CallForwardingForwardingTypeOnFailure CallForwardingForwardingType = "on-failure"
)

// The call forwarding settings for a phone number.
type CallForwardingParam struct {
	// Indicates if call forwarding will be enabled for this number if forwards_to and
	// forwarding_type are filled in. Defaults to true for backwards compatibility with
	// APIV1 use of numbers endpoints.
	CallForwardingEnabled param.Opt[bool] `json:"call_forwarding_enabled,omitzero"`
	// The phone number to which inbound calls to this number are forwarded. Inbound
	// calls will not be forwarded if this field is left blank. If set, must be a
	// +E.164-formatted phone number.
	ForwardsTo param.Opt[string] `json:"forwards_to,omitzero"`
	// Call forwarding type. 'forwards_to' must be set for this to have an effect.
	//
	// Any of "always", "on-failure".
	ForwardingType CallForwardingForwardingType `json:"forwarding_type,omitzero"`
	paramObj
}

func (r CallForwardingParam) MarshalJSON() (data []byte, err error) {
	type shadow CallForwardingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallForwardingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The call recording settings for a phone number.
type CallRecording struct {
	// When using 'dual' channels, final audio file will be stereo recorded with the
	// first leg on channel A, and the rest on channel B.
	//
	// Any of "single", "dual".
	InboundCallRecordingChannels CallRecordingInboundCallRecordingChannels `json:"inbound_call_recording_channels"`
	// When enabled, any inbound call to this number will be recorded.
	InboundCallRecordingEnabled bool `json:"inbound_call_recording_enabled"`
	// The audio file format for calls being recorded.
	//
	// Any of "wav", "mp3".
	InboundCallRecordingFormat CallRecordingInboundCallRecordingFormat `json:"inbound_call_recording_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InboundCallRecordingChannels respjson.Field
		InboundCallRecordingEnabled  respjson.Field
		InboundCallRecordingFormat   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecording) RawJSON() string { return r.JSON.raw }
func (r *CallRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CallRecording to a CallRecordingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CallRecordingParam.Overrides()
func (r CallRecording) ToParam() CallRecordingParam {
	return param.Override[CallRecordingParam](json.RawMessage(r.RawJSON()))
}

// When using 'dual' channels, final audio file will be stereo recorded with the
// first leg on channel A, and the rest on channel B.
type CallRecordingInboundCallRecordingChannels string

const (
	CallRecordingInboundCallRecordingChannelsSingle CallRecordingInboundCallRecordingChannels = "single"
	CallRecordingInboundCallRecordingChannelsDual   CallRecordingInboundCallRecordingChannels = "dual"
)

// The audio file format for calls being recorded.
type CallRecordingInboundCallRecordingFormat string

const (
	CallRecordingInboundCallRecordingFormatWav CallRecordingInboundCallRecordingFormat = "wav"
	CallRecordingInboundCallRecordingFormatMP3 CallRecordingInboundCallRecordingFormat = "mp3"
)

// The call recording settings for a phone number.
type CallRecordingParam struct {
	// When enabled, any inbound call to this number will be recorded.
	InboundCallRecordingEnabled param.Opt[bool] `json:"inbound_call_recording_enabled,omitzero"`
	// When using 'dual' channels, final audio file will be stereo recorded with the
	// first leg on channel A, and the rest on channel B.
	//
	// Any of "single", "dual".
	InboundCallRecordingChannels CallRecordingInboundCallRecordingChannels `json:"inbound_call_recording_channels,omitzero"`
	// The audio file format for calls being recorded.
	//
	// Any of "wav", "mp3".
	InboundCallRecordingFormat CallRecordingInboundCallRecordingFormat `json:"inbound_call_recording_format,omitzero"`
	paramObj
}

func (r CallRecordingParam) MarshalJSON() (data []byte, err error) {
	type shadow CallRecordingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallRecordingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The CNAM listing settings for a phone number.
type CnamListing struct {
	// The CNAM listing details for this number. Must be alphanumeric characters or
	// spaces with a maximum length of 15. Requires cnam_listing_enabled to also be set
	// to true.
	CnamListingDetails string `json:"cnam_listing_details"`
	// Enables CNAM listings for this number. Requires cnam_listing_details to also be
	// set.
	CnamListingEnabled bool `json:"cnam_listing_enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnamListingDetails respjson.Field
		CnamListingEnabled respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CnamListing) RawJSON() string { return r.JSON.raw }
func (r *CnamListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CnamListing to a CnamListingParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CnamListingParam.Overrides()
func (r CnamListing) ToParam() CnamListingParam {
	return param.Override[CnamListingParam](json.RawMessage(r.RawJSON()))
}

// The CNAM listing settings for a phone number.
type CnamListingParam struct {
	// The CNAM listing details for this number. Must be alphanumeric characters or
	// spaces with a maximum length of 15. Requires cnam_listing_enabled to also be set
	// to true.
	CnamListingDetails param.Opt[string] `json:"cnam_listing_details,omitzero"`
	// Enables CNAM listings for this number. Requires cnam_listing_details to also be
	// set.
	CnamListingEnabled param.Opt[bool] `json:"cnam_listing_enabled,omitzero"`
	paramObj
}

func (r CnamListingParam) MarshalJSON() (data []byte, err error) {
	type shadow CnamListingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CnamListingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The media features settings for a phone number.
type MediaFeatures struct {
	// When enabled, Telnyx will accept RTP packets from any customer-side IP address
	// and port, not just those to which Telnyx is sending RTP.
	AcceptAnyRtpPacketsEnabled bool `json:"accept_any_rtp_packets_enabled"`
	// When RTP Auto-Adjust is enabled, the destination RTP address port will be
	// automatically changed to match the source of the incoming RTP packets.
	RtpAutoAdjustEnabled bool `json:"rtp_auto_adjust_enabled"`
	// Controls whether Telnyx will accept a T.38 re-INVITE for this phone number. Note
	// that Telnyx will not send a T.38 re-INVITE; this option only controls whether
	// one will be accepted.
	T38FaxGatewayEnabled bool `json:"t38_fax_gateway_enabled"`
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
func (r MediaFeatures) RawJSON() string { return r.JSON.raw }
func (r *MediaFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MediaFeatures to a MediaFeaturesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MediaFeaturesParam.Overrides()
func (r MediaFeatures) ToParam() MediaFeaturesParam {
	return param.Override[MediaFeaturesParam](json.RawMessage(r.RawJSON()))
}

// The media features settings for a phone number.
type MediaFeaturesParam struct {
	// When enabled, Telnyx will accept RTP packets from any customer-side IP address
	// and port, not just those to which Telnyx is sending RTP.
	AcceptAnyRtpPacketsEnabled param.Opt[bool] `json:"accept_any_rtp_packets_enabled,omitzero"`
	// When RTP Auto-Adjust is enabled, the destination RTP address port will be
	// automatically changed to match the source of the incoming RTP packets.
	RtpAutoAdjustEnabled param.Opt[bool] `json:"rtp_auto_adjust_enabled,omitzero"`
	// Controls whether Telnyx will accept a T.38 re-INVITE for this phone number. Note
	// that Telnyx will not send a T.38 re-INVITE; this option only controls whether
	// one will be accepted.
	T38FaxGatewayEnabled param.Opt[bool] `json:"t38_fax_gateway_enabled,omitzero"`
	paramObj
}

func (r MediaFeaturesParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaFeaturesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaFeaturesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateVoiceSettingsParam struct {
	// Controls whether the caller ID name is enabled for this phone number.
	CallerIDNameEnabled param.Opt[bool] `json:"caller_id_name_enabled,omitzero"`
	// Controls whether a tech prefix is enabled for this phone number.
	TechPrefixEnabled param.Opt[bool] `json:"tech_prefix_enabled,omitzero"`
	// This field allows you to rewrite the destination number of an inbound call
	// before the call is routed to you. The value of this field may be any
	// alphanumeric value, and the value will replace the number originally dialed.
	TranslatedNumber param.Opt[string] `json:"translated_number,omitzero"`
	// The call forwarding settings for a phone number.
	CallForwarding CallForwardingParam `json:"call_forwarding,omitzero"`
	// The call recording settings for a phone number.
	CallRecording CallRecordingParam `json:"call_recording,omitzero"`
	// The CNAM listing settings for a phone number.
	CnamListing CnamListingParam `json:"cnam_listing,omitzero"`
	// The inbound_call_screening setting is a phone number configuration option
	// variable that allows users to configure their settings to block or flag
	// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
	// feature has an additional per-number monthly cost associated with it.
	//
	// Any of "disabled", "reject_calls", "flag_calls".
	InboundCallScreening UpdateVoiceSettingsInboundCallScreening `json:"inbound_call_screening,omitzero"`
	// The media features settings for a phone number.
	MediaFeatures MediaFeaturesParam `json:"media_features,omitzero"`
	// Controls whether a number is billed per minute or uses your concurrent channels.
	//
	// Any of "pay-per-minute", "channel".
	UsagePaymentMethod UpdateVoiceSettingsUsagePaymentMethod `json:"usage_payment_method,omitzero"`
	paramObj
}

func (r UpdateVoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateVoiceSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateVoiceSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The inbound_call_screening setting is a phone number configuration option
// variable that allows users to configure their settings to block or flag
// fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This
// feature has an additional per-number monthly cost associated with it.
type UpdateVoiceSettingsInboundCallScreening string

const (
	UpdateVoiceSettingsInboundCallScreeningDisabled    UpdateVoiceSettingsInboundCallScreening = "disabled"
	UpdateVoiceSettingsInboundCallScreeningRejectCalls UpdateVoiceSettingsInboundCallScreening = "reject_calls"
	UpdateVoiceSettingsInboundCallScreeningFlagCalls   UpdateVoiceSettingsInboundCallScreening = "flag_calls"
)

// Controls whether a number is billed per minute or uses your concurrent channels.
type UpdateVoiceSettingsUsagePaymentMethod string

const (
	UpdateVoiceSettingsUsagePaymentMethodPayPerMinute UpdateVoiceSettingsUsagePaymentMethod = "pay-per-minute"
	UpdateVoiceSettingsUsagePaymentMethodChannel      UpdateVoiceSettingsUsagePaymentMethod = "channel"
)

type PhoneNumberVoiceGetResponse struct {
	Data PhoneNumberWithVoiceSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberVoiceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberVoiceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoiceUpdateResponse struct {
	Data PhoneNumberWithVoiceSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberVoiceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberVoiceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberVoiceUpdateParams struct {
	UpdateVoiceSettings UpdateVoiceSettingsParam
	paramObj
}

func (r PhoneNumberVoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateVoiceSettings)
}
func (r *PhoneNumberVoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateVoiceSettings)
}

type PhoneNumberVoiceListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[phone_number], filter[connection_name], filter[customer_reference],
	// filter[voice.usage_payment_method]
	Filter PhoneNumberVoiceListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "purchased_at", "phone_number", "connection_name",
	// "usage_payment_method".
	Sort PhoneNumberVoiceListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberVoiceListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberVoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[phone_number], filter[connection_name], filter[customer_reference],
// filter[voice.usage_payment_method]
type PhoneNumberVoiceListParamsFilter struct {
	// Filter numbers via the customer_reference set.
	CustomerReference param.Opt[string] `query:"customer_reference,omitzero" json:"-"`
	// Filter by phone number. Requires at least three digits. Non-numerical characters
	// will result in no values being returned.
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	// Filter by connection name pattern matching.
	ConnectionName PhoneNumberVoiceListParamsFilterConnectionName `query:"connection_name,omitzero" json:"-"`
	// Filter by usage_payment_method.
	//
	// Any of "pay-per-minute", "channel".
	VoiceUsagePaymentMethod string `query:"voice.usage_payment_method,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberVoiceListParamsFilter]'s query parameters as
// `url.Values`.
func (r PhoneNumberVoiceListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by connection name pattern matching.
type PhoneNumberVoiceListParamsFilterConnectionName struct {
	// Filter contains connection name. Requires at least three characters.
	Contains param.Opt[string] `query:"contains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberVoiceListParamsFilterConnectionName]'s query
// parameters as `url.Values`.
func (r PhoneNumberVoiceListParamsFilterConnectionName) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// created_at in descending order.
type PhoneNumberVoiceListParamsSort string

const (
	PhoneNumberVoiceListParamsSortPurchasedAt        PhoneNumberVoiceListParamsSort = "purchased_at"
	PhoneNumberVoiceListParamsSortPhoneNumber        PhoneNumberVoiceListParamsSort = "phone_number"
	PhoneNumberVoiceListParamsSortConnectionName     PhoneNumberVoiceListParamsSort = "connection_name"
	PhoneNumberVoiceListParamsSortUsagePaymentMethod PhoneNumberVoiceListParamsSort = "usage_payment_method"
)
