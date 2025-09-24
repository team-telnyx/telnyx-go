// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// TexmlCallService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlCallService] method instead.
type TexmlCallService struct {
	Options []option.RequestOption
}

// NewTexmlCallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTexmlCallService(opts ...option.RequestOption) (r TexmlCallService) {
	r = TexmlCallService{}
	r.Options = opts
	return
}

// Update TeXML call. Please note that the keys present in the payload MUST BE
// formatted in CamelCase as specified in the example.
func (r *TexmlCallService) Update(ctx context.Context, callSid string, body TexmlCallUpdateParams, opts ...option.RequestOption) (res *TexmlCallUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/calls/%s/update", callSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate an outbound TeXML call. Telnyx will request TeXML from the XML Request
// URL configured for the connection in the Mission Control Portal.
func (r *TexmlCallService) Initiate(ctx context.Context, applicationID string, body TexmlCallInitiateParams, opts ...option.RequestOption) (res *TexmlCallInitiateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return
	}
	path := fmt.Sprintf("texml/calls/%s", applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TexmlCallUpdateResponse struct {
	Data TexmlCallUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlCallUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlCallUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlCallUpdateResponseData struct {
	Sid    string `json:"sid"`
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sid         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlCallUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *TexmlCallUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlCallInitiateResponse struct {
	Data TexmlCallInitiateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlCallInitiateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlCallInitiateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlCallInitiateResponseData struct {
	From   string `json:"from"`
	Status string `json:"status"`
	To     string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlCallInitiateResponseData) RawJSON() string { return r.JSON.raw }
func (r *TexmlCallInitiateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlCallUpdateParams struct {
	UpdateCall UpdateCallParam
	paramObj
}

func (r TexmlCallUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCall)
}
func (r *TexmlCallUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateCall)
}

type TexmlCallInitiateParams struct {
	// The phone number of the party that initiated the call. Phone numbers are
	// formatted with a `+` and country code.
	From string `json:"From,required"`
	// The phone number of the called party. Phone numbers are formatted with a `+` and
	// country code.
	To string `json:"To,required"`
	// Select whether to perform answering machine detection in the background. By
	// default execution is blocked until Answering Machine Detection is completed.
	AsyncAmd param.Opt[bool] `json:"AsyncAmd,omitzero"`
	// URL destination for Telnyx to send AMD callback events to for the call.
	AsyncAmdStatusCallback param.Opt[string] `json:"AsyncAmdStatusCallback,omitzero"`
	// To be used as the caller id name (SIP From Display Name) presented to the
	// destination (`To` number). The string should have a maximum of 128 characters,
	// containing only letters, numbers, spaces, and `-_~!.+` special characters. If
	// ommited, the display name will be the same as the number in the `From` field.
	CallerID param.Opt[string] `json:"CallerId,omitzero"`
	// Whether to cancel ongoing playback on `greeting ended` detection. Defaults to
	// `true`.
	CancelPlaybackOnDetectMessageEnd param.Opt[bool] `json:"CancelPlaybackOnDetectMessageEnd,omitzero"`
	// Whether to cancel ongoing playback on `machine` detection. Defaults to `true`.
	CancelPlaybackOnMachineDetection param.Opt[bool] `json:"CancelPlaybackOnMachineDetection,omitzero"`
	// A failover URL for which Telnyx will retrieve the TeXML call instructions if the
	// `Url` is not responding.
	FallbackURL param.Opt[string] `json:"FallbackUrl,omitzero"`
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
	// The list of comma-separated codecs to be offered on a call.
	PreferredCodecs param.Opt[string] `json:"PreferredCodecs,omitzero"`
	// Whether to record the entire participant's call leg. Defaults to `false`.
	Record param.Opt[bool] `json:"Record,omitzero"`
	// The URL the recording callbacks will be sent to.
	RecordingStatusCallback param.Opt[string] `json:"RecordingStatusCallback,omitzero"`
	// The changes to the recording's state that should generate a call to
	// `RecoridngStatusCallback`. Can be: `in-progress`, `completed` and `absent`.
	// Separate multiple values with a space. Defaults to `completed`.
	RecordingStatusCallbackEvent param.Opt[string] `json:"RecordingStatusCallbackEvent,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected. The timer only starts when the speech is detected. Please
	// note that the transcription is used to detect silence and the related charge
	// will be applied. The minimum value is 0. The default value is 0 (infinite)
	RecordingTimeout param.Opt[int64] `json:"RecordingTimeout,omitzero"`
	// The password to use for SIP authentication.
	SipAuthPassword param.Opt[string] `json:"SipAuthPassword,omitzero"`
	// The username to use for SIP authentication.
	SipAuthUsername param.Opt[string] `json:"SipAuthUsername,omitzero"`
	// URL destination for Telnyx to send status callback events to for the call.
	StatusCallback param.Opt[string] `json:"StatusCallback,omitzero"`
	// The URL from which Telnyx will retrieve the TeXML call instructions.
	URL param.Opt[string] `json:"Url,omitzero"`
	// HTTP request type used for `AsyncAmdStatusCallback`. The default value is
	// inherited from TeXML Application setting.
	//
	// Any of "GET", "POST".
	AsyncAmdStatusCallbackMethod TexmlCallInitiateParamsAsyncAmdStatusCallbackMethod `json:"AsyncAmdStatusCallbackMethod,omitzero"`
	// Allows you to chose between Premium and Standard detections.
	//
	// Any of "Premium", "Regular".
	DetectionMode TexmlCallInitiateParamsDetectionMode `json:"DetectionMode,omitzero"`
	// Enables Answering Machine Detection.
	//
	// Any of "Enable", "Disable", "DetectMessageEnd".
	MachineDetection TexmlCallInitiateParamsMachineDetection `json:"MachineDetection,omitzero"`
	// The number of channels in the final recording. Defaults to `mono`.
	//
	// Any of "mono", "dual".
	RecordingChannels TexmlCallInitiateParamsRecordingChannels `json:"RecordingChannels,omitzero"`
	// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
	//
	// Any of "GET", "POST".
	RecordingStatusCallbackMethod TexmlCallInitiateParamsRecordingStatusCallbackMethod `json:"RecordingStatusCallbackMethod,omitzero"`
	// The audio track to record for the call. The default is `both`.
	//
	// Any of "inbound", "outbound", "both".
	RecordingTrack TexmlCallInitiateParamsRecordingTrack `json:"RecordingTrack,omitzero"`
	// The call events for which Telnyx should send a webhook. Multiple events can be
	// defined when separated by a space.
	//
	// Any of "initiated", "ringing", "answered", "completed".
	StatusCallbackEvent TexmlCallInitiateParamsStatusCallbackEvent `json:"StatusCallbackEvent,omitzero"`
	// HTTP request type used for `StatusCallback`.
	//
	// Any of "GET", "POST".
	StatusCallbackMethod TexmlCallInitiateParamsStatusCallbackMethod `json:"StatusCallbackMethod,omitzero"`
	// Whether to trim any leading and trailing silence from the recording. Defaults to
	// `trim-silence`.
	//
	// Any of "trim-silence", "do-not-trim".
	Trim TexmlCallInitiateParamsTrim `json:"Trim,omitzero"`
	// HTTP request type used for `Url`. The default value is inherited from TeXML
	// Application setting.
	//
	// Any of "GET", "POST".
	URLMethod TexmlCallInitiateParamsURLMethod `json:"UrlMethod,omitzero"`
	paramObj
}

func (r TexmlCallInitiateParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlCallInitiateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlCallInitiateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request type used for `AsyncAmdStatusCallback`. The default value is
// inherited from TeXML Application setting.
type TexmlCallInitiateParamsAsyncAmdStatusCallbackMethod string

const (
	TexmlCallInitiateParamsAsyncAmdStatusCallbackMethodGet  TexmlCallInitiateParamsAsyncAmdStatusCallbackMethod = "GET"
	TexmlCallInitiateParamsAsyncAmdStatusCallbackMethodPost TexmlCallInitiateParamsAsyncAmdStatusCallbackMethod = "POST"
)

// Allows you to chose between Premium and Standard detections.
type TexmlCallInitiateParamsDetectionMode string

const (
	TexmlCallInitiateParamsDetectionModePremium TexmlCallInitiateParamsDetectionMode = "Premium"
	TexmlCallInitiateParamsDetectionModeRegular TexmlCallInitiateParamsDetectionMode = "Regular"
)

// Enables Answering Machine Detection.
type TexmlCallInitiateParamsMachineDetection string

const (
	TexmlCallInitiateParamsMachineDetectionEnable           TexmlCallInitiateParamsMachineDetection = "Enable"
	TexmlCallInitiateParamsMachineDetectionDisable          TexmlCallInitiateParamsMachineDetection = "Disable"
	TexmlCallInitiateParamsMachineDetectionDetectMessageEnd TexmlCallInitiateParamsMachineDetection = "DetectMessageEnd"
)

// The number of channels in the final recording. Defaults to `mono`.
type TexmlCallInitiateParamsRecordingChannels string

const (
	TexmlCallInitiateParamsRecordingChannelsMono TexmlCallInitiateParamsRecordingChannels = "mono"
	TexmlCallInitiateParamsRecordingChannelsDual TexmlCallInitiateParamsRecordingChannels = "dual"
)

// HTTP request type used for `RecordingStatusCallback`. Defaults to `POST`.
type TexmlCallInitiateParamsRecordingStatusCallbackMethod string

const (
	TexmlCallInitiateParamsRecordingStatusCallbackMethodGet  TexmlCallInitiateParamsRecordingStatusCallbackMethod = "GET"
	TexmlCallInitiateParamsRecordingStatusCallbackMethodPost TexmlCallInitiateParamsRecordingStatusCallbackMethod = "POST"
)

// The audio track to record for the call. The default is `both`.
type TexmlCallInitiateParamsRecordingTrack string

const (
	TexmlCallInitiateParamsRecordingTrackInbound  TexmlCallInitiateParamsRecordingTrack = "inbound"
	TexmlCallInitiateParamsRecordingTrackOutbound TexmlCallInitiateParamsRecordingTrack = "outbound"
	TexmlCallInitiateParamsRecordingTrackBoth     TexmlCallInitiateParamsRecordingTrack = "both"
)

// The call events for which Telnyx should send a webhook. Multiple events can be
// defined when separated by a space.
type TexmlCallInitiateParamsStatusCallbackEvent string

const (
	TexmlCallInitiateParamsStatusCallbackEventInitiated TexmlCallInitiateParamsStatusCallbackEvent = "initiated"
	TexmlCallInitiateParamsStatusCallbackEventRinging   TexmlCallInitiateParamsStatusCallbackEvent = "ringing"
	TexmlCallInitiateParamsStatusCallbackEventAnswered  TexmlCallInitiateParamsStatusCallbackEvent = "answered"
	TexmlCallInitiateParamsStatusCallbackEventCompleted TexmlCallInitiateParamsStatusCallbackEvent = "completed"
)

// HTTP request type used for `StatusCallback`.
type TexmlCallInitiateParamsStatusCallbackMethod string

const (
	TexmlCallInitiateParamsStatusCallbackMethodGet  TexmlCallInitiateParamsStatusCallbackMethod = "GET"
	TexmlCallInitiateParamsStatusCallbackMethodPost TexmlCallInitiateParamsStatusCallbackMethod = "POST"
)

// Whether to trim any leading and trailing silence from the recording. Defaults to
// `trim-silence`.
type TexmlCallInitiateParamsTrim string

const (
	TexmlCallInitiateParamsTrimTrimSilence TexmlCallInitiateParamsTrim = "trim-silence"
	TexmlCallInitiateParamsTrimDoNotTrim   TexmlCallInitiateParamsTrim = "do-not-trim"
)

// HTTP request type used for `Url`. The default value is inherited from TeXML
// Application setting.
type TexmlCallInitiateParamsURLMethod string

const (
	TexmlCallInitiateParamsURLMethodGet  TexmlCallInitiateParamsURLMethod = "GET"
	TexmlCallInitiateParamsURLMethodPost TexmlCallInitiateParamsURLMethod = "POST"
)
