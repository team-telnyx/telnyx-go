// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// TeXML REST Commands
//
// TexmlService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlService] method instead.
type TexmlService struct {
	Options []option.RequestOption
	// TeXML REST Commands
	Accounts TexmlAccountService
}

// NewTexmlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTexmlService(opts ...option.RequestOption) (r TexmlService) {
	r = TexmlService{}
	r.Options = opts
	r.Accounts = NewTexmlAccountService(opts...)
	return
}

// Initiate an outbound AI call with warm-up support. Validates parameters, builds
// an internal TeXML with an AI Assistant configuration, encodes instructions into
// client state, and calls the dial API. The Twiml, Texml, and Url parameters are
// not allowed and will result in a 422 error.
func (r *TexmlService) InitiateAICall(ctx context.Context, connectionID string, body TexmlInitiateAICallParams, opts ...option.RequestOption) (res *TexmlInitiateAICallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("texml/ai_calls/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a TeXML secret which can be later used as a Dynamic Parameter for TeXML
// when using Mustache Templates in your TeXML. In your TeXML you will be able to
// use your secret name, and this name will be replaced by the actual secret value
// when processing the TeXML on Telnyx side. The secrets are not visible in any
// logs.
func (r *TexmlService) Secrets(ctx context.Context, body TexmlSecretsParams, opts ...option.RequestOption) (res *TexmlSecretsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "texml/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TexmlInitiateAICallResponse struct {
	CallSid string `json:"call_sid"`
	From    string `json:"from"`
	Status  string `json:"status"`
	To      string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSid     respjson.Field
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlInitiateAICallResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlInitiateAICallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlSecretsResponse struct {
	Data TexmlSecretsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlSecretsResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlSecretsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlSecretsResponseData struct {
	Name string `json:"name"`
	// Any of "REDACTED".
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlSecretsResponseData) RawJSON() string { return r.JSON.raw }
func (r *TexmlSecretsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlInitiateAICallParams struct {
	// The ID of the AI assistant to use for the call.
	AIAssistantID string `json:"AIAssistantId" api:"required"`
	// The phone number of the party initiating the call. Phone numbers are formatted
	// with a `+` and country code.
	From string `json:"From" api:"required"`
	// The phone number of the called party. Phone numbers are formatted with a `+` and
	// country code.
	To string `json:"To" api:"required"`
	// The version of the AI assistant to use.
	AIAssistantVersion param.Opt[string] `json:"AIAssistantVersion,omitzero"`
	// Select whether to perform answering machine detection in the background. By
	// default execution is blocked until Answering Machine Detection is completed.
	AsyncAmd param.Opt[bool] `json:"AsyncAmd,omitzero"`
	// URL destination for Telnyx to send AMD callback events to for the call.
	AsyncAmdStatusCallback param.Opt[string] `json:"AsyncAmdStatusCallback,omitzero"`
	// To be used as the caller id name (SIP From Display Name) presented to the
	// destination (`To` number). The string should have a maximum of 128 characters,
	// containing only letters, numbers, spaces, and `-_~!.+` special characters. If
	// omitted, the display name will be the same as the number in the `From` field.
	CallerID param.Opt[string] `json:"CallerId,omitzero"`
	// URL destination for Telnyx to send conversation callback events to.
	ConversationCallback param.Opt[string] `json:"ConversationCallback,omitzero"`
	// If initial silence duration is greater than this value, consider it a machine.
	// Ignored when `premium` detection is used.
	MachineDetectionSilenceTimeout param.Opt[int64] `json:"MachineDetectionSilenceTimeout,omitzero"`
	// Silence duration threshold after a greeting message or voice for it be
	// considered human. Ignored when `premium` detection is used.
	MachineDetectionSpeechEndThreshold param.Opt[int64] `json:"MachineDetectionSpeechEndThreshold,omitzero"`
	// Maximum threshold of a human greeting. If greeting longer than this value,
	// considered machine. Ignored when `premium` detection is used.
	MachineDetectionSpeechThreshold param.Opt[int64] `json:"MachineDetectionSpeechThreshold,omitzero"`
	// Maximum timeout threshold in milliseconds for overall detection.
	MachineDetectionTimeout param.Opt[int64] `json:"MachineDetectionTimeout,omitzero"`
	// A string of passport identifiers to associate with the call.
	Passports param.Opt[string] `json:"Passports,omitzero"`
	// The list of comma-separated codecs to be offered on a call.
	PreferredCodecs param.Opt[string] `json:"PreferredCodecs,omitzero"`
	// Whether to record the entire participant's call leg. Defaults to `false`.
	Record param.Opt[bool] `json:"Record,omitzero"`
	// The URL the recording callbacks will be sent to.
	RecordingStatusCallback param.Opt[string] `json:"RecordingStatusCallback,omitzero"`
	// The changes to the recording's state that should generate a call to
	// `RecordingStatusCallback`. Can be: `in-progress`, `completed` and `absent`.
	// Separate multiple values with a space. Defaults to `completed`.
	RecordingStatusCallbackEvent param.Opt[string] `json:"RecordingStatusCallbackEvent,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected. The timer only starts when the speech is detected. The
	// minimum value is 0. The default value is 0 (infinite).
	RecordingTimeout param.Opt[int64] `json:"RecordingTimeout,omitzero"`
	// Whether to send RecordingUrl in webhooks.
	SendRecordingURL param.Opt[bool] `json:"SendRecordingUrl,omitzero"`
	// The password to use for SIP authentication.
	SipAuthPassword param.Opt[string] `json:"SipAuthPassword,omitzero"`
	// The username to use for SIP authentication.
	SipAuthUsername param.Opt[string] `json:"SipAuthUsername,omitzero"`
	// URL destination for Telnyx to send status callback events to for the call.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero"`
	// The call events for which Telnyx should send a webhook. Multiple events can be
	// defined when separated by a space. Valid values: initiated, ringing, answered,
	// completed.
	StatusCallbackEvent param.Opt[string] `json:"StatusCallbackEvent,omitzero"`
	// The maximum duration of the call in seconds. The minimum value is 30 and the
	// maximum value is 14400 (4 hours). Default is 14400 seconds.
	TimeLimit param.Opt[int64] `json:"TimeLimit,omitzero"`
	// The number of seconds to wait for the called party to answer the call before the
	// call is canceled. The minimum value is 5 and the maximum value is 120. Default
	// is 30 seconds.
	TimeoutSeconds param.Opt[int64] `json:"Timeout,omitzero"`
	// Key-value map of dynamic variables to pass to the AI assistant.
	AIAssistantDynamicVariables map[string]string `json:"AIAssistantDynamicVariables,omitzero"`
	// HTTP request type used for `AsyncAmdStatusCallback`.
	//
	// Any of "GET", "POST".
	AsyncAmdStatusCallbackMethod TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethod `json:"AsyncAmdStatusCallbackMethod,omitzero"`
	// HTTP request type used for `ConversationCallback`.
	//
	// Any of "GET", "POST".
	ConversationCallbackMethod TexmlInitiateAICallParamsConversationCallbackMethod `json:"ConversationCallbackMethod,omitzero"`
	// An array of URL destinations for conversation callback events.
	ConversationCallbacks []string `json:"ConversationCallbacks,omitzero"`
	// Custom HTTP headers to be sent with the call. Each header should be an object
	// with 'name' and 'value' properties.
	CustomHeaders []TexmlInitiateAICallParamsCustomHeader `json:"CustomHeaders,omitzero"`
	// Allows you to choose between Premium and Standard detections.
	//
	// Any of "Premium", "Regular".
	DetectionMode TexmlInitiateAICallParamsDetectionMode `json:"DetectionMode,omitzero"`
	// Enables Answering Machine Detection.
	//
	// Any of "Enable", "Disable", "DetectMessageEnd".
	MachineDetection TexmlInitiateAICallParamsMachineDetection `json:"MachineDetection,omitzero"`
	// The number of channels in the final recording. Defaults to `mono`.
	//
	// Any of "mono", "dual".
	RecordingChannels TexmlInitiateAICallParamsRecordingChannels `json:"RecordingChannels,omitzero"`
	// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	RecordingStatusCallbackMethod TexmlInitiateAICallParamsRecordingStatusCallbackMethod `json:"RecordingStatusCallbackMethod,omitzero"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	RecordingTrack TexmlInitiateAICallParamsRecordingTrack `json:"RecordingTrack,omitzero"`
	// Defines the SIP region to be used for the call.
	//
	// Any of "US", "Europe", "Canada", "Australia", "Middle East".
	SipRegion TexmlInitiateAICallParamsSipRegion `json:"SipRegion,omitzero"`
	// HTTP request type used for `StatusCallback`.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod TexmlInitiateAICallParamsStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	// An array of URL destinations for Telnyx to send status callback events to for
	// the call.
	StatusCallbacks []string `json:"StatusCallbacks,omitzero"`
	// Whether to trim any leading and trailing silence from the recording. Defaults to
	// `trim-silence`.
	//
	// Any of "trim-silence", "do-not-trim".
	Trim TexmlInitiateAICallParamsTrim `json:"Trim,omitzero"`
	paramObj
}

func (r TexmlInitiateAICallParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlInitiateAICallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlInitiateAICallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request type used for `AsyncAmdStatusCallback`.
type TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethod string

const (
	TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethodGet  TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethod = "GET"
	TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethodPost TexmlInitiateAICallParamsAsyncAmdStatusCallbackMethod = "POST"
)

// HTTP request type used for `ConversationCallback`.
type TexmlInitiateAICallParamsConversationCallbackMethod string

const (
	TexmlInitiateAICallParamsConversationCallbackMethodGet  TexmlInitiateAICallParamsConversationCallbackMethod = "GET"
	TexmlInitiateAICallParamsConversationCallbackMethodPost TexmlInitiateAICallParamsConversationCallbackMethod = "POST"
)

// The properties Name, Value are required.
type TexmlInitiateAICallParamsCustomHeader struct {
	// The name of the custom header
	Name string `json:"name" api:"required"`
	// The value of the custom header
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TexmlInitiateAICallParamsCustomHeader) MarshalJSON() (data []byte, err error) {
	type shadow TexmlInitiateAICallParamsCustomHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlInitiateAICallParamsCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows you to choose between Premium and Standard detections.
type TexmlInitiateAICallParamsDetectionMode string

const (
	TexmlInitiateAICallParamsDetectionModePremium TexmlInitiateAICallParamsDetectionMode = "Premium"
	TexmlInitiateAICallParamsDetectionModeRegular TexmlInitiateAICallParamsDetectionMode = "Regular"
)

// Enables Answering Machine Detection.
type TexmlInitiateAICallParamsMachineDetection string

const (
	TexmlInitiateAICallParamsMachineDetectionEnable           TexmlInitiateAICallParamsMachineDetection = "Enable"
	TexmlInitiateAICallParamsMachineDetectionDisable          TexmlInitiateAICallParamsMachineDetection = "Disable"
	TexmlInitiateAICallParamsMachineDetectionDetectMessageEnd TexmlInitiateAICallParamsMachineDetection = "DetectMessageEnd"
)

// The number of channels in the final recording. Defaults to `mono`.
type TexmlInitiateAICallParamsRecordingChannels string

const (
	TexmlInitiateAICallParamsRecordingChannelsMono TexmlInitiateAICallParamsRecordingChannels = "mono"
	TexmlInitiateAICallParamsRecordingChannelsDual TexmlInitiateAICallParamsRecordingChannels = "dual"
)

// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
type TexmlInitiateAICallParamsRecordingStatusCallbackMethod string

const (
	TexmlInitiateAICallParamsRecordingStatusCallbackMethodGet  TexmlInitiateAICallParamsRecordingStatusCallbackMethod = "GET"
	TexmlInitiateAICallParamsRecordingStatusCallbackMethodPost TexmlInitiateAICallParamsRecordingStatusCallbackMethod = "POST"
)

// The audio track to record for the call. The default is `both`.
type TexmlInitiateAICallParamsRecordingTrack string

const (
	TexmlInitiateAICallParamsRecordingTrackInbound  TexmlInitiateAICallParamsRecordingTrack = "inbound"
	TexmlInitiateAICallParamsRecordingTrackOutbound TexmlInitiateAICallParamsRecordingTrack = "outbound"
	TexmlInitiateAICallParamsRecordingTrackBoth     TexmlInitiateAICallParamsRecordingTrack = "both"
)

// Defines the SIP region to be used for the call.
type TexmlInitiateAICallParamsSipRegion string

const (
	TexmlInitiateAICallParamsSipRegionUs         TexmlInitiateAICallParamsSipRegion = "US"
	TexmlInitiateAICallParamsSipRegionEurope     TexmlInitiateAICallParamsSipRegion = "Europe"
	TexmlInitiateAICallParamsSipRegionCanada     TexmlInitiateAICallParamsSipRegion = "Canada"
	TexmlInitiateAICallParamsSipRegionAustralia  TexmlInitiateAICallParamsSipRegion = "Australia"
	TexmlInitiateAICallParamsSipRegionMiddleEast TexmlInitiateAICallParamsSipRegion = "Middle East"
)

// HTTP request type used for `StatusCallback`.
type TexmlInitiateAICallParamsStatusCallbackMethod string

const (
	TexmlInitiateAICallParamsStatusCallbackMethodGet  TexmlInitiateAICallParamsStatusCallbackMethod = "GET"
	TexmlInitiateAICallParamsStatusCallbackMethodPost TexmlInitiateAICallParamsStatusCallbackMethod = "POST"
)

// Whether to trim any leading and trailing silence from the recording. Defaults to
// `trim-silence`.
type TexmlInitiateAICallParamsTrim string

const (
	TexmlInitiateAICallParamsTrimTrimSilence TexmlInitiateAICallParamsTrim = "trim-silence"
	TexmlInitiateAICallParamsTrimDoNotTrim   TexmlInitiateAICallParamsTrim = "do-not-trim"
)

type TexmlSecretsParams struct {
	// Name used as a reference for the secret, if the name already exists within the
	// account its value will be replaced
	Name string `json:"name" api:"required"`
	// Secret value which will be used when rendering the TeXML template
	Value string `json:"value" api:"required"`
	paramObj
}

func (r TexmlSecretsParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlSecretsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlSecretsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
