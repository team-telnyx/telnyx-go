// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// CallService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
	// Call Control command operations
	Actions CallActionService
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r CallService) {
	r = CallService{}
	r.Options = opts
	r.Actions = NewCallActionService(opts...)
	return
}

// Dial a number or SIP URI from a given connection. A successful response will
// include a `call_leg_id` which can be used to correlate the command with
// subsequent webhooks.
//
// **Expected Webhooks:**
//
//   - `call.initiated`
//   - `call.answered` or `call.hangup`
//   - `call.hold` and `call.unhold` if the call is held/unheld
//   - `call.machine.detection.ended` if `answering_machine_detection` was requested
//   - `call.machine.greeting.ended` if `answering_machine_detection` was requested
//     to detect the end of machine greeting
//   - `call.machine.premium.detection.ended` if
//     `answering_machine_detection=premium` was requested
//   - `call.machine.premium.greeting.ended` if `answering_machine_detection=premium`
//     was requested and a beep was detected
//   - `call.deepfake_detection.result` if `deepfake_detection` was enabled
//   - `call.deepfake_detection.error` if `deepfake_detection` was enabled and an
//     error occurred
//   - `streaming.started`, `streaming.stopped` or `streaming.failed` if `stream_url`
//     was set
//
// When the `record` parameter is set to `record-from-answer`, the response will
// include a `recording_id` field.
func (r *CallService) Dial(ctx context.Context, body CallDialParams, opts ...option.RequestOption) (res *CallDialResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the status of a call (data is available 10 minutes after call ended).
func (r *CallService) GetStatus(ctx context.Context, callControlID string, opts ...option.RequestOption) (res *CallGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("calls/%s", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// AI Assistant configuration. All fields except `id` are optional — the
// assistant's stored configuration will be used as fallback for any omitted
// fields.
//
// The property ID is required.
type CallAssistantRequestParam struct {
	// The identifier of the AI assistant to use.
	ID string `json:"id" api:"required"`
	// Initial greeting text spoken when the assistant starts. Can be plain text for
	// any voice or SSML for `AWS.Polly.<voice_id>` voices. There is a 3,000 character
	// limit.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// System instructions for the voice assistant. Can be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables).
	// This will overwrite the instructions set in the assistant configuration.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Integration secret identifier for the LLM provider API key. Use this field to
	// reference an
	// [integration secret](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// containing your LLM provider API key. Supports any LLM provider (OpenAI,
	// Anthropic, etc.).
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// LLM model override for this call. If omitted, the assistant's configured model
	// is used.
	Model param.Opt[string] `json:"model,omitzero"`
	// Assistant name override for this call.
	Name param.Opt[string] `json:"name,omitzero"`
	// Deprecated — use `llm_api_key_ref` instead. Integration secret identifier for
	// the OpenAI API key. This field is maintained for backward compatibility;
	// `llm_api_key_ref` is the canonical field name and supports all LLM providers.
	//
	// Deprecated: This field is deprecated and will be removed soon
	OpenAIAPIKeyRef param.Opt[string] `json:"openai_api_key_ref,omitzero"`
	// Map of dynamic variables and their default values. Dynamic variables can be
	// referenced in instructions, greeting, and tool definitions using the
	// `{{variable_name}}` syntax. Call-control-agent automatically merges in
	// `telnyx_call_*` variables (telnyx_call_to, telnyx_call_from,
	// telnyx_conversation_channel, telnyx_agent_target, telnyx_end_user_target,
	// telnyx_call_caller_id_name) and custom header variables.
	DynamicVariables map[string]CallAssistantRequestDynamicVariableUnionParam `json:"dynamic_variables,omitzero"`
	// External LLM configuration for bringing your own LLM endpoint.
	ExternalLlm CallAssistantRequestExternalLlmParam `json:"external_llm,omitzero"`
	// Fallback LLM configuration used when the primary LLM provider is unavailable.
	FallbackConfig CallAssistantRequestFallbackConfigParam `json:"fallback_config,omitzero"`
	// MCP (Model Context Protocol) server configurations for extending the assistant's
	// capabilities with external tools and data sources.
	McpServers []map[string]any `json:"mcp_servers,omitzero"`
	// Observability configuration for the assistant session, including Langfuse
	// integration for tracing and monitoring.
	ObservabilitySettings map[string]any `json:"observability_settings,omitzero"`
	// Inline tool definitions available to the assistant (webhook, retrieval,
	// transfer, hangup, etc.). Overrides the assistant's stored tools if provided.
	Tools []CallAssistantRequestToolUnionParam `json:"tools,omitzero"`
	paramObj
}

func (r CallAssistantRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CallAssistantRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallAssistantRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallAssistantRequestDynamicVariableUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CallAssistantRequestDynamicVariableUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CallAssistantRequestDynamicVariableUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallAssistantRequestDynamicVariableUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// External LLM configuration for bringing your own LLM endpoint.
type CallAssistantRequestExternalLlmParam struct {
	// Base URL for the external LLM endpoint.
	BaseURL param.Opt[string] `json:"base_url,omitzero"`
	// Integration secret identifier for the client certificate used with certificate
	// authentication.
	CertificateRef param.Opt[string] `json:"certificate_ref,omitzero"`
	// When enabled, Telnyx forwards the assistant's dynamic variables to the external
	// LLM endpoint. Defaults to false. The chat completion request includes a
	// top-level `extra_metadata` object when dynamic variables are available. For
	// example:
	// `{"extra_metadata":{"customer_name":"Jane","account_id":"acct_789","telnyx_agent_target":"+13125550100","telnyx_end_user_target":"+13125550123"}}`.
	ForwardMetadata param.Opt[bool] `json:"forward_metadata,omitzero"`
	// Integration secret identifier for the external LLM API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Model identifier to use with the external LLM endpoint.
	Model param.Opt[string] `json:"model,omitzero"`
	// URL used to retrieve an access token when certificate authentication is enabled.
	TokenRetrievalURL param.Opt[string] `json:"token_retrieval_url,omitzero"`
	// Authentication method used when connecting to the external LLM endpoint.
	//
	// Any of "token", "certificate".
	AuthenticationMethod string         `json:"authentication_method,omitzero"`
	ExtraFields          map[string]any `json:"-"`
	paramObj
}

func (r CallAssistantRequestExternalLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow CallAssistantRequestExternalLlmParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CallAssistantRequestExternalLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallAssistantRequestExternalLlmParam](
		"authentication_method", "token", "certificate",
	)
}

// Fallback LLM configuration used when the primary LLM provider is unavailable.
type CallAssistantRequestFallbackConfigParam struct {
	// Integration secret identifier for the fallback model API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Fallback Telnyx-hosted model to use when the primary LLM provider is
	// unavailable.
	Model param.Opt[string] `json:"model,omitzero"`
	// External LLM fallback configuration.
	ExternalLlm CallAssistantRequestFallbackConfigExternalLlmParam `json:"external_llm,omitzero"`
	ExtraFields map[string]any                                     `json:"-"`
	paramObj
}

func (r CallAssistantRequestFallbackConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CallAssistantRequestFallbackConfigParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CallAssistantRequestFallbackConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// External LLM fallback configuration.
type CallAssistantRequestFallbackConfigExternalLlmParam struct {
	// Base URL for the external LLM endpoint.
	BaseURL param.Opt[string] `json:"base_url,omitzero"`
	// Integration secret identifier for the client certificate used with certificate
	// authentication.
	CertificateRef param.Opt[string] `json:"certificate_ref,omitzero"`
	// When enabled, Telnyx forwards the assistant's dynamic variables to the external
	// LLM endpoint. Defaults to false. The chat completion request includes a
	// top-level `extra_metadata` object when dynamic variables are available. For
	// example:
	// `{"extra_metadata":{"customer_name":"Jane","account_id":"acct_789","telnyx_agent_target":"+13125550100","telnyx_end_user_target":"+13125550123"}}`.
	ForwardMetadata param.Opt[bool] `json:"forward_metadata,omitzero"`
	// Integration secret identifier for the external LLM API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Model identifier to use with the external LLM endpoint.
	Model param.Opt[string] `json:"model,omitzero"`
	// URL used to retrieve an access token when certificate authentication is enabled.
	TokenRetrievalURL param.Opt[string] `json:"token_retrieval_url,omitzero"`
	// Authentication method used when connecting to the external LLM endpoint.
	//
	// Any of "token", "certificate".
	AuthenticationMethod string `json:"authentication_method,omitzero"`
	paramObj
}

func (r CallAssistantRequestFallbackConfigExternalLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow CallAssistantRequestFallbackConfigExternalLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallAssistantRequestFallbackConfigExternalLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallAssistantRequestFallbackConfigExternalLlmParam](
		"authentication_method", "token", "certificate",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallAssistantRequestToolUnionParam struct {
	OfBookAppointment   *shared.BookAppointmentToolParam      `json:",omitzero,inline"`
	OfCheckAvailability *shared.CheckAvailabilityToolParam    `json:",omitzero,inline"`
	OfWebhook           *WebhookToolParam                     `json:",omitzero,inline"`
	OfHangup            *HangupToolParam                      `json:",omitzero,inline"`
	OfTransfer          *TransferToolParam                    `json:",omitzero,inline"`
	OfRetrieval         *shared.CallControlRetrievalToolParam `json:",omitzero,inline"`
	paramUnion
}

func (u CallAssistantRequestToolUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBookAppointment,
		u.OfCheckAvailability,
		u.OfWebhook,
		u.OfHangup,
		u.OfTransfer,
		u.OfRetrieval)
}
func (u *CallAssistantRequestToolUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallAssistantRequestToolUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBookAppointment) {
		return u.OfBookAppointment
	} else if !param.IsOmitted(u.OfCheckAvailability) {
		return u.OfCheckAvailability
	} else if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfHangup) {
		return u.OfHangup
	} else if !param.IsOmitted(u.OfTransfer) {
		return u.OfTransfer
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetBookAppointment() *shared.BookAppointmentToolParams {
	if vt := u.OfBookAppointment; vt != nil {
		return &vt.BookAppointment
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetCheckAvailability() *shared.CheckAvailabilityToolParams {
	if vt := u.OfCheckAvailability; vt != nil {
		return &vt.CheckAvailability
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetWebhook() *WebhookToolWebhookParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetHangup() *HangupToolParams {
	if vt := u.OfHangup; vt != nil {
		return &vt.Hangup
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetTransfer() *TransferToolTransferParam {
	if vt := u.OfTransfer; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetRetrieval() *shared.CallControlBucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallAssistantRequestToolUnionParam) GetType() *string {
	if vt := u.OfBookAppointment; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCheckAvailability; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHangup; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTransfer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[CallAssistantRequestToolUnionParam](
		"type",
		apijson.Discriminator[shared.BookAppointmentToolParam]("book_appointment"),
		apijson.Discriminator[shared.CheckAvailabilityToolParam]("check_availability"),
		apijson.Discriminator[WebhookToolParam]("webhook"),
		apijson.Discriminator[HangupToolParam]("hangup"),
		apijson.Discriminator[TransferToolParam]("transfer"),
		apijson.Discriminator[shared.CallControlRetrievalToolParam]("retrieval"),
	)
}

type CustomSipHeader struct {
	// The name of the header to add.
	Name string `json:"name" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomSipHeader) RawJSON() string { return r.JSON.raw }
func (r *CustomSipHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CustomSipHeader to a CustomSipHeaderParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CustomSipHeaderParam.Overrides()
func (r CustomSipHeader) ToParam() CustomSipHeaderParam {
	return param.Override[CustomSipHeaderParam](json.RawMessage(r.RawJSON()))
}

// The properties Name, Value are required.
type CustomSipHeaderParam struct {
	// The name of the header to add.
	Name string `json:"name" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r CustomSipHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomSipHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomSipHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConfigParam struct {
	// Enable sentiment analysis from Dialogflow.
	AnalyzeSentiment param.Opt[bool] `json:"analyze_sentiment,omitzero"`
	// Enable partial automated agent reply from Dialogflow.
	PartialAutomatedAgentReply param.Opt[bool] `json:"partial_automated_agent_reply,omitzero"`
	paramObj
}

func (r DialogflowConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow DialogflowConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DialogflowConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipHeader struct {
	// The name of the header to add.
	//
	// Any of "User-to-User".
	Name SipHeaderName `json:"name" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SipHeader) RawJSON() string { return r.JSON.raw }
func (r *SipHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SipHeader to a SipHeaderParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SipHeaderParam.Overrides()
func (r SipHeader) ToParam() SipHeaderParam {
	return param.Override[SipHeaderParam](json.RawMessage(r.RawJSON()))
}

// The name of the header to add.
type SipHeaderName string

const (
	SipHeaderNameUserToUser SipHeaderName = "User-to-User"
)

// The properties Name, Value are required.
type SipHeaderParam struct {
	// The name of the header to add.
	//
	// Any of "User-to-User".
	Name SipHeaderName `json:"name,omitzero" api:"required"`
	// The value of the header.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r SipHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow SipHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SipHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Use this field to modify sound effects, for example adjust the pitch.
type SoundModificationsParam struct {
	// Adjust the pitch in octaves, values should be between -1 and 1, default 0
	Octaves param.Opt[float64] `json:"octaves,omitzero"`
	// Set the pitch directly, value should be > 0, default 1 (lower = lower tone)
	Pitch param.Opt[float64] `json:"pitch,omitzero"`
	// Adjust the pitch in semitones, values should be between -14 and 14, default 0
	Semitone param.Opt[float64] `json:"semitone,omitzero"`
	// The track to which the sound modifications will be applied. Accepted values are
	// `inbound` or `outbound`
	Track param.Opt[string] `json:"track,omitzero"`
	paramObj
}

func (r SoundModificationsParam) MarshalJSON() (data []byte, err error) {
	type shadow SoundModificationsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoundModificationsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates codec for bidirectional streaming RTP payloads. Used only with
// stream_bidirectional_mode=rtp. Case sensitive.
type StreamBidirectionalCodec string

const (
	StreamBidirectionalCodecPcmu  StreamBidirectionalCodec = "PCMU"
	StreamBidirectionalCodecPcma  StreamBidirectionalCodec = "PCMA"
	StreamBidirectionalCodecG722  StreamBidirectionalCodec = "G722"
	StreamBidirectionalCodecOpus  StreamBidirectionalCodec = "OPUS"
	StreamBidirectionalCodecAmrWb StreamBidirectionalCodec = "AMR-WB"
	StreamBidirectionalCodecL16   StreamBidirectionalCodec = "L16"
)

// Configures method of bidirectional streaming (mp3, rtp).
type StreamBidirectionalMode string

const (
	StreamBidirectionalModeMP3 StreamBidirectionalMode = "mp3"
	StreamBidirectionalModeRtp StreamBidirectionalMode = "rtp"
)

// Audio sampling rate.
type StreamBidirectionalSamplingRate int64

const (
	StreamBidirectionalSamplingRateRate8000  StreamBidirectionalSamplingRate = 8000
	StreamBidirectionalSamplingRateRate16000 StreamBidirectionalSamplingRate = 16000
	StreamBidirectionalSamplingRateRate22050 StreamBidirectionalSamplingRate = 22050
	StreamBidirectionalSamplingRateRate24000 StreamBidirectionalSamplingRate = 24000
	StreamBidirectionalSamplingRateRate48000 StreamBidirectionalSamplingRate = 48000
)

// Specifies which call legs should receive the bidirectional stream audio.
type StreamBidirectionalTargetLegs string

const (
	StreamBidirectionalTargetLegsBoth     StreamBidirectionalTargetLegs = "both"
	StreamBidirectionalTargetLegsSelf     StreamBidirectionalTargetLegs = "self"
	StreamBidirectionalTargetLegsOpposite StreamBidirectionalTargetLegs = "opposite"
)

// Specifies the codec to be used for the streamed audio. When set to 'default' or
// when transcoding is not possible, the codec from the call will be used.
type StreamCodec string

const (
	StreamCodecPcmu    StreamCodec = "PCMU"
	StreamCodecPcma    StreamCodec = "PCMA"
	StreamCodecG722    StreamCodec = "G722"
	StreamCodecOpus    StreamCodec = "OPUS"
	StreamCodecAmrWb   StreamCodec = "AMR-WB"
	StreamCodecL16     StreamCodec = "L16"
	StreamCodecDefault StreamCodec = "default"
)

type CallDialResponse struct {
	Data CallDialResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallDialResponse) RawJSON() string { return r.JSON.raw }
func (r *CallDialResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallDialResponseData struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id" api:"required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id" api:"required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id" api:"required"`
	// Indicates whether the call is alive or not. For Dial command it will always be
	// `false` (dialing is asynchronous).
	IsAlive bool `json:"is_alive" api:"required"`
	// Any of "call".
	RecordType string `json:"record_type" api:"required"`
	// Indicates the duration of the call in seconds
	CallDuration int64 `json:"call_duration"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ISO 8601 formatted date indicating when the call ended. Only present when the
	// call is not alive
	EndTime string `json:"end_time"`
	// The ID of the recording. Only present when the record parameter is set to
	// record-from-answer.
	RecordingID string `json:"recording_id" format:"uuid"`
	// ISO 8601 formatted date indicating when the call started
	StartTime string `json:"start_time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		IsAlive       respjson.Field
		RecordType    respjson.Field
		CallDuration  respjson.Field
		ClientState   respjson.Field
		EndTime       respjson.Field
		RecordingID   respjson.Field
		StartTime     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallDialResponseData) RawJSON() string { return r.JSON.raw }
func (r *CallDialResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetStatusResponse struct {
	Data CallGetStatusResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *CallGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetStatusResponseData struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id" api:"required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id" api:"required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id" api:"required"`
	// Indicates whether the call is alive or not. For Dial command it will always be
	// `false` (dialing is asynchronous).
	IsAlive bool `json:"is_alive" api:"required"`
	// Any of "call".
	RecordType string `json:"record_type" api:"required"`
	// Indicates the duration of the call in seconds
	CallDuration int64 `json:"call_duration"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ISO 8601 formatted date indicating when the call ended. Only present when the
	// call is not alive
	EndTime string `json:"end_time"`
	// ISO 8601 formatted date indicating when the call started
	StartTime string `json:"start_time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		IsAlive       respjson.Field
		RecordType    respjson.Field
		CallDuration  respjson.Field
		ClientState   respjson.Field
		EndTime       respjson.Field
		StartTime     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetStatusResponseData) RawJSON() string { return r.JSON.raw }
func (r *CallGetStatusResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallDialParams struct {
	// The ID of the Call Control App (formerly ID of the connection) to be used when
	// dialing the destination.
	ConnectionID string `json:"connection_id" api:"required"`
	// The `from` number to be used as the caller id presented to the destination (`to`
	// number). The number should be in +E164 format.
	From string `json:"from" api:"required"`
	// The DID or SIP URI to dial out to. Multiple DID or SIP URIs can be provided
	// using an array of strings. For SIP URI destinations, append `;secure=true` or
	// `;secure=srtp` to enable SRTP media encryption for that endpoint, or
	// `;secure=dtls` to enable DTLS media encryption for that endpoint. If
	// `media_encryption` is set to `SRTP` or `DTLS`, it takes precedence over any
	// per-endpoint `secure` URI parameter. For a single string destination, you may
	// append a comma followed by DTMF digits (e.g. `+18004247767,200`) to play those
	// digits as DTMF once the called party answers — equivalent to setting
	// `send_digits_on_answer` separately. If both are present, the explicit
	// `send_digits_on_answer` parameter takes precedence. This shorthand is not
	// supported when `to` is an array.
	To CallDialParamsToUnion `json:"to,omitzero" api:"required"`
	// The URL of a file to be played back to the callee when the call is answered. The
	// URL can point to either a WAV or MP3 file. media_name and audio_url cannot be
	// used together in one request.
	AudioURL param.Opt[string] `json:"audio_url,omitzero"`
	// Use this field to set the Billing Group ID for the call. Must be a valid and
	// existing Billing Group ID.
	BillingGroupID param.Opt[string] `json:"billing_group_id,omitzero" format:"uuid"`
	// Indicates the intent to bridge this call with the call specified in link_to.
	// When bridge_intent is true, link_to becomes required and the from number will be
	// overwritten by the from number from the linked call.
	BridgeIntent param.Opt[bool] `json:"bridge_intent,omitzero"`
	// Whether to automatically bridge answered call to the call specified in link_to.
	// When bridge_on_answer is true, link_to becomes required.
	BridgeOnAnswer param.Opt[bool] `json:"bridge_on_answer,omitzero"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState param.Opt[string] `json:"client_state,omitzero"`
	// Use this field to avoid duplicate commands. Telnyx will ignore others Dial
	// commands with the same `command_id`.
	CommandID param.Opt[string] `json:"command_id,omitzero"`
	// Enables Dialogflow for the current call. The default value is false.
	EnableDialogflow param.Opt[bool] `json:"enable_dialogflow,omitzero"`
	// The `from_display_name` string to be used as the caller id name (SIP From
	// Display Name) presented to the destination (`to` number). The string should have
	// a maximum of 128 characters, containing only letters, numbers, spaces, and
	// -\_~!.+ special characters. If ommited, the display name will be the same as the
	// number in the `from` field.
	FromDisplayName param.Opt[string] `json:"from_display_name,omitzero"`
	// Use another call's control id for sharing the same call session id
	LinkTo param.Opt[string] `json:"link_to,omitzero"`
	// The media_name of a file to be played back to the callee when the call is
	// answered. The media_name must point to a file previously uploaded to
	// api.telnyx.com/v2/media by the same user/organization. The file must either be a
	// WAV or MP3 file.
	MediaName param.Opt[string] `json:"media_name,omitzero"`
	// If supplied with the value `self`, the current leg will be parked after
	// unbridge. If not set, the default behavior is to hang up the leg. When
	// park_after_unbridge is set, link_to becomes required.
	ParkAfterUnbridge param.Opt[string] `json:"park_after_unbridge,omitzero"`
	// The list of comma-separated codecs in a preferred order for the forked media to
	// be received.
	PreferredCodecs param.Opt[string] `json:"preferred_codecs,omitzero"`
	// Prevents bridging and hangs up the call if the target is already bridged.
	// Disabled by default.
	PreventDoubleBridge param.Opt[bool] `json:"prevent_double_bridge,omitzero"`
	// The custom recording file name to be used instead of the default `call_leg_id`.
	// Telnyx will still add a Unix timestamp suffix.
	RecordCustomFileName param.Opt[string] `json:"record_custom_file_name,omitzero"`
	// Defines the maximum length for the recording in seconds when `record` is
	// specified. The minimum value is 0. The maximum value is 43200. The default value
	// is 0 (infinite).
	RecordMaxLength param.Opt[int64] `json:"record_max_length,omitzero"`
	// The number of seconds that Telnyx will wait for the recording to be stopped if
	// silence is detected when `record` is specified. The timer only starts when the
	// speech is detected. Please note that call transcription is used to detect
	// silence and the related charge will be applied. The minimum value is 0. The
	// default value is 0 (infinite).
	RecordTimeoutSecs param.Opt[int64] `json:"record_timeout_secs,omitzero"`
	// DTMF digits to send automatically after the called party answers. Useful for
	// reaching an extension behind an IVR (e.g. `"200"` to dial extension 200 once the
	// called party picks up). Allowed characters: `0-9`, `A-D`, `w` (0.5s pause), `W`
	// (1s pause), `*`, `#`. Maximum 64 characters. When omitted, no automatic DTMF is
	// sent. May also be supplied inline by appending `,<digits>` to `to` (e.g.
	// `to=+18004247767,200`); if both forms are present, this explicit field takes
	// precedence.
	SendDigitsOnAnswer param.Opt[string] `json:"send_digits_on_answer,omitzero"`
	// Generate silence RTP packets when no transmission available.
	SendSilenceWhenIdle param.Opt[bool] `json:"send_silence_when_idle,omitzero"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword param.Opt[string] `json:"sip_auth_password,omitzero"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername param.Opt[string] `json:"sip_auth_username,omitzero"`
	// An authentication token to be sent as part of the WebSocket connection when
	// using streaming. Maximum length is 4000 characters.
	StreamAuthToken param.Opt[string] `json:"stream_auth_token,omitzero"`
	// Establish websocket connection before dialing the destination. This is useful
	// for cases where the websocket connection takes a long time to establish.
	StreamEstablishBeforeCallOriginate param.Opt[bool] `json:"stream_establish_before_call_originate,omitzero"`
	// The destination WebSocket address where the stream is going to be delivered.
	StreamURL param.Opt[string] `json:"stream_url,omitzero"`
	// The call leg which will be supervised by the new call.
	SuperviseCallControlID param.Opt[string] `json:"supervise_call_control_id,omitzero"`
	// Sets the maximum duration of a Call Control Leg in seconds. If the time limit is
	// reached, the call will hangup and a `call.hangup` webhook with a `hangup_cause`
	// of `time_limit` will be sent. For example, by setting a time limit of 120
	// seconds, a Call Leg will be automatically terminated two minutes after being
	// answered. The default time limit is 14400 seconds or 4 hours and this is also
	// the maximum allowed call length.
	TimeLimitSecs param.Opt[int64] `json:"time_limit_secs,omitzero"`
	// The number of seconds that Telnyx will wait for the call to be answered by the
	// destination to which it is being called. If the timeout is reached before an
	// answer is received, the call will hangup and a `call.hangup` webhook with a
	// `hangup_cause` of `timeout` will be sent. Minimum value is 5 seconds. Maximum
	// value is 600 seconds.
	TimeoutSecs param.Opt[int64] `json:"timeout_secs,omitzero"`
	// Enable transcription upon call answer. The default value is false.
	Transcription param.Opt[bool] `json:"transcription,omitzero"`
	// Use this field to override the URL for which Telnyx will send subsequent
	// webhooks to for this call.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Enables Answering Machine Detection. Telnyx offers Premium and Standard
	// detections. With Premium detection, when a call is answered, Telnyx runs
	// real-time detection and sends a `call.machine.premium.detection.ended` webhook
	// with one of the following results: `human_residence`, `human_business`,
	// `machine`, `silence` or `fax_detected`. If we detect a beep, we also send a
	// `call.machine.premium.greeting.ended` webhook with the result of
	// `beep_detected`. If we detect a beep before
	// `call.machine.premium.detection.ended` we only send
	// `call.machine.premium.greeting.ended`, and if we detect a beep after
	// `call.machine.premium.detection.ended`, we send both webhooks. With Standard
	// detection, when a call is answered, Telnyx runs real-time detection to determine
	// if it was picked up by a human or a machine and sends an
	// `call.machine.detection.ended` webhook with the analysis result. If
	// `greeting_end` or `detect_words` is used and a `machine` is detected, you will
	// receive another `call.machine.greeting.ended` webhook when the answering machine
	// greeting ends with a beep or silence. If `detect_beep` is used, you will only
	// receive `call.machine.greeting.ended` if a beep is detected.
	//
	// Any of "premium", "detect", "detect_beep", "detect_words", "greeting_end",
	// "disabled".
	AnsweringMachineDetection CallDialParamsAnsweringMachineDetection `json:"answering_machine_detection,omitzero"`
	// Optional configuration parameters to modify 'answering_machine_detection'
	// performance. Only `total_analysis_time_millis` and `greeting_duration_millis`
	// parameters are applicable when `premium` is selected as
	// answering_machine_detection.
	AnsweringMachineDetectionConfig CallDialParamsAnsweringMachineDetectionConfig `json:"answering_machine_detection_config,omitzero"`
	// AI Assistant configuration. All fields except `id` are optional — the
	// assistant's stored configuration will be used as fallback for any omitted
	// fields.
	Assistant CallAssistantRequestParam `json:"assistant,omitzero"`
	// Optional configuration parameters to dial new participant into a conference.
	ConferenceConfig CallDialParamsConferenceConfig `json:"conference_config,omitzero"`
	// Starts a Conversation Relay session automatically when the answered/dialed call
	// is answered. This embedded shape is supported on `answer` and `dial`. It uses
	// public field names (`url`, `dtmf_detection`, `greeting`, `voice`, `language`,
	// etc.) and maps them to the underlying Conversation Relay action. `client_state`,
	// `tts_language`, and `transcription_language` inside this object are ignored; use
	// the parent command's `client_state` and `command_id` fields instead.
	ConversationRelayConfig CallDialParamsConversationRelayConfig `json:"conversation_relay_config,omitzero"`
	// Custom headers to be added to the SIP INVITE.
	CustomHeaders []CustomSipHeaderParam `json:"custom_headers,omitzero"`
	// Enables deepfake detection on the call. When enabled, audio from the remote
	// party is streamed to a detection service that analyzes whether the voice is
	// AI-generated. Results are delivered via the `call.deepfake_detection.result`
	// webhook.
	DeepfakeDetection CallDialParamsDeepfakeDetection `json:"deepfake_detection,omitzero"`
	DialogflowConfig  DialogflowConfigParam           `json:"dialogflow_config,omitzero"`
	// Defines whether media should be encrypted on the call. For SIP URI destinations,
	// media encryption can also be requested per endpoint with the `secure` URI
	// parameter: `;secure=true` or `;secure=srtp` enables SRTP, and `;secure=dtls`
	// enables DTLS. This parameter, when set to `SRTP` or `DTLS`, takes precedence
	// over the per-endpoint `secure` value.
	//
	// Any of "disabled", "SRTP", "DTLS".
	MediaEncryption CallDialParamsMediaEncryption `json:"media_encryption,omitzero"`
	// Indicates the privacy level to be used for the call. When set to `id`, caller ID
	// information (name and number) will be hidden from the called party. When set to
	// `none` or omitted, caller ID will be shown normally.
	//
	// Any of "id", "none".
	Privacy CallDialParamsPrivacy `json:"privacy,omitzero"`
	// Start recording automatically after an event. Disabled by default.
	//
	// Any of "record-from-answer".
	Record CallDialParamsRecord `json:"record,omitzero"`
	// Defines which channel should be recorded ('single' or 'dual') when `record` is
	// specified.
	//
	// Any of "single", "dual".
	RecordChannels CallDialParamsRecordChannels `json:"record_channels,omitzero"`
	// Defines the format of the recording ('wav' or 'mp3') when `record` is specified.
	//
	// Any of "wav", "mp3".
	RecordFormat CallDialParamsRecordFormat `json:"record_format,omitzero"`
	// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
	// If only single track is specified (`inbound`, `outbound`), `channels`
	// configuration is ignored and it will be recorded as mono (single channel).
	//
	// Any of "both", "inbound", "outbound".
	RecordTrack CallDialParamsRecordTrack `json:"record_track,omitzero"`
	// When set to `trim-silence`, silence will be removed from the beginning and end
	// of the recording.
	//
	// Any of "trim-silence".
	RecordTrim CallDialParamsRecordTrim `json:"record_trim,omitzero"`
	// SIP headers to be added to the SIP INVITE request. Currently only User-to-User
	// header is supported.
	SipHeaders []SipHeaderParam `json:"sip_headers,omitzero"`
	// Defines the SIP region to be used for the call.
	//
	// Any of "US", "Europe", "Canada", "Australia", "Middle East".
	SipRegion CallDialParamsSipRegion `json:"sip_region,omitzero"`
	// Defines SIP transport protocol to be used on the call.
	//
	// Any of "UDP", "TCP", "TLS".
	SipTransportProtocol CallDialParamsSipTransportProtocol `json:"sip_transport_protocol,omitzero"`
	// Use this field to modify sound effects, for example adjust the pitch.
	SoundModifications SoundModificationsParam `json:"sound_modifications,omitzero"`
	// Indicates codec for bidirectional streaming RTP payloads. Used only with
	// stream_bidirectional_mode=rtp. Case sensitive.
	//
	// Any of "PCMU", "PCMA", "G722", "OPUS", "AMR-WB", "L16".
	StreamBidirectionalCodec StreamBidirectionalCodec `json:"stream_bidirectional_codec,omitzero"`
	// Configures method of bidirectional streaming (mp3, rtp).
	//
	// Any of "mp3", "rtp".
	StreamBidirectionalMode StreamBidirectionalMode `json:"stream_bidirectional_mode,omitzero"`
	// Audio sampling rate.
	//
	// Any of 8000, 16000, 22050, 24000, 48000.
	StreamBidirectionalSamplingRate int64 `json:"stream_bidirectional_sampling_rate,omitzero"`
	// Specifies which call legs should receive the bidirectional stream audio.
	//
	// Any of "both", "self", "opposite".
	StreamBidirectionalTargetLegs StreamBidirectionalTargetLegs `json:"stream_bidirectional_target_legs,omitzero"`
	// Specifies the codec to be used for the streamed audio. When set to 'default' or
	// when transcoding is not possible, the codec from the call will be used.
	//
	// Any of "PCMU", "PCMA", "G722", "OPUS", "AMR-WB", "L16", "default".
	StreamCodec StreamCodec `json:"stream_codec,omitzero"`
	// Specifies which track should be streamed.
	//
	// Any of "inbound_track", "outbound_track", "both_tracks".
	StreamTrack CallDialParamsStreamTrack `json:"stream_track,omitzero"`
	// The role of the supervisor call. 'barge' means that supervisor call hears and is
	// being heard by both ends of the call (caller & callee). 'whisper' means that
	// only supervised_call_control_id hears supervisor but supervisor can hear
	// everything. 'monitor' means that nobody can hear supervisor call, but supervisor
	// can hear everything on the call.
	//
	// Any of "barge", "whisper", "monitor".
	SupervisorRole      CallDialParamsSupervisorRole   `json:"supervisor_role,omitzero"`
	TranscriptionConfig TranscriptionStartRequestParam `json:"transcription_config,omitzero"`
	// A map of event types to retry policies. Each retry policy contains an array of
	// `retries_ms` specifying the delays between retry attempts in milliseconds.
	// Maximum 5 retries, total delay cannot exceed 60 seconds.
	WebhookRetriesPolicies map[string]CallDialParamsWebhookRetriesPolicy `json:"webhook_retries_policies,omitzero"`
	// HTTP request type used for `webhook_url`.
	//
	// Any of "POST", "GET".
	WebhookURLMethod CallDialParamsWebhookURLMethod `json:"webhook_url_method,omitzero"`
	// A map of event types to webhook URLs. When an event of the specified type
	// occurs, the webhook URL associated with that event type will be called instead
	// of the default webhook URL. Events not mapped here will use the default webhook
	// URL.
	WebhookURLs map[string]string `json:"webhook_urls,omitzero" format:"uri"`
	// HTTP request method to invoke `webhook_urls`.
	//
	// Any of "POST", "GET".
	WebhookURLsMethod CallDialParamsWebhookURLsMethod `json:"webhook_urls_method,omitzero"`
	paramObj
}

func (r CallDialParams) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallDialParamsToUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CallDialParamsToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CallDialParamsToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallDialParamsToUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Enables Answering Machine Detection. Telnyx offers Premium and Standard
// detections. With Premium detection, when a call is answered, Telnyx runs
// real-time detection and sends a `call.machine.premium.detection.ended` webhook
// with one of the following results: `human_residence`, `human_business`,
// `machine`, `silence` or `fax_detected`. If we detect a beep, we also send a
// `call.machine.premium.greeting.ended` webhook with the result of
// `beep_detected`. If we detect a beep before
// `call.machine.premium.detection.ended` we only send
// `call.machine.premium.greeting.ended`, and if we detect a beep after
// `call.machine.premium.detection.ended`, we send both webhooks. With Standard
// detection, when a call is answered, Telnyx runs real-time detection to determine
// if it was picked up by a human or a machine and sends an
// `call.machine.detection.ended` webhook with the analysis result. If
// `greeting_end` or `detect_words` is used and a `machine` is detected, you will
// receive another `call.machine.greeting.ended` webhook when the answering machine
// greeting ends with a beep or silence. If `detect_beep` is used, you will only
// receive `call.machine.greeting.ended` if a beep is detected.
type CallDialParamsAnsweringMachineDetection string

const (
	CallDialParamsAnsweringMachineDetectionPremium     CallDialParamsAnsweringMachineDetection = "premium"
	CallDialParamsAnsweringMachineDetectionDetect      CallDialParamsAnsweringMachineDetection = "detect"
	CallDialParamsAnsweringMachineDetectionDetectBeep  CallDialParamsAnsweringMachineDetection = "detect_beep"
	CallDialParamsAnsweringMachineDetectionDetectWords CallDialParamsAnsweringMachineDetection = "detect_words"
	CallDialParamsAnsweringMachineDetectionGreetingEnd CallDialParamsAnsweringMachineDetection = "greeting_end"
	CallDialParamsAnsweringMachineDetectionDisabled    CallDialParamsAnsweringMachineDetection = "disabled"
)

// Optional configuration parameters to modify 'answering_machine_detection'
// performance. Only `total_analysis_time_millis` and `greeting_duration_millis`
// parameters are applicable when `premium` is selected as
// answering_machine_detection.
type CallDialParamsAnsweringMachineDetectionConfig struct {
	// Silence duration threshold after a greeting message or voice for it be
	// considered human.
	AfterGreetingSilenceMillis param.Opt[int64] `json:"after_greeting_silence_millis,omitzero"`
	// Maximum threshold for silence between words.
	BetweenWordsSilenceMillis param.Opt[int64] `json:"between_words_silence_millis,omitzero"`
	// Maximum threshold of a human greeting. If greeting longer than this value,
	// considered machine.
	GreetingDurationMillis param.Opt[int64] `json:"greeting_duration_millis,omitzero"`
	// If machine already detected, maximum threshold for silence between words. If
	// exceeded, the greeting is considered ended.
	GreetingSilenceDurationMillis param.Opt[int64] `json:"greeting_silence_duration_millis,omitzero"`
	// If machine already detected, maximum timeout threshold to determine the end of
	// the machine greeting.
	GreetingTotalAnalysisTimeMillis param.Opt[int64] `json:"greeting_total_analysis_time_millis,omitzero"`
	// If initial silence duration is greater than this value, consider it a machine.
	InitialSilenceMillis param.Opt[int64] `json:"initial_silence_millis,omitzero"`
	// If number of detected words is greater than this value, consder it a machine.
	MaximumNumberOfWords param.Opt[int64] `json:"maximum_number_of_words,omitzero"`
	// If a single word lasts longer than this threshold, consider it a machine.
	MaximumWordLengthMillis param.Opt[int64] `json:"maximum_word_length_millis,omitzero"`
	// Minimum noise threshold for any analysis.
	SilenceThreshold param.Opt[int64] `json:"silence_threshold,omitzero"`
	// Maximum timeout threshold for overall detection.
	TotalAnalysisTimeMillis param.Opt[int64] `json:"total_analysis_time_millis,omitzero"`
	paramObj
}

func (r CallDialParamsAnsweringMachineDetectionConfig) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsAnsweringMachineDetectionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsAnsweringMachineDetectionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional configuration parameters to dial new participant into a conference.
type CallDialParamsConferenceConfig struct {
	// Conference ID to be joined
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// Conference name to be joined
	ConferenceName param.Opt[string] `json:"conference_name,omitzero"`
	// Controls the moment when dialled call is joined into conference. If set to
	// `true` user will be joined as soon as media is available (ringback). If `false`
	// user will be joined when call is answered. Defaults to `true`
	EarlyMedia param.Opt[bool] `json:"early_media,omitzero"`
	// Whether the conference should end and all remaining participants be hung up
	// after the participant leaves the conference. Defaults to "false".
	EndConferenceOnExit param.Opt[bool] `json:"end_conference_on_exit,omitzero"`
	// Whether the participant should be put on hold immediately after joining the
	// conference. Defaults to "false".
	Hold param.Opt[bool] `json:"hold,omitzero"`
	// The URL of a file to be played to the participant when they are put on hold
	// after joining the conference. hold_media_name and hold_audio_url cannot be used
	// together in one request. Takes effect only when "start_conference_on_create" is
	// set to "false". This property takes effect only if "hold" is set to "true".
	HoldAudioURL param.Opt[string] `json:"hold_audio_url,omitzero"`
	// The media_name of a file to be played to the participant when they are put on
	// hold after joining the conference. The media_name must point to a file
	// previously uploaded to api.telnyx.com/v2/media by the same user/organization.
	// The file must either be a WAV or MP3 file. Takes effect only when
	// "start_conference_on_create" is set to "false". This property takes effect only
	// if "hold" is set to "true".
	HoldMediaName param.Opt[string] `json:"hold_media_name,omitzero"`
	// Whether the participant should be muted immediately after joining the
	// conference. Defaults to "false".
	Mute param.Opt[bool] `json:"mute,omitzero"`
	// Whether the conference should end after the participant leaves the conference.
	// NOTE this doesn't hang up the other participants. Defaults to "false".
	SoftEndConferenceOnExit param.Opt[bool] `json:"soft_end_conference_on_exit,omitzero"`
	// Whether the conference should be started on creation. If the conference isn't
	// started all participants that join are automatically put on hold. Defaults to
	// "true".
	StartConferenceOnCreate param.Opt[bool] `json:"start_conference_on_create,omitzero"`
	// Whether the conference should be started after the participant joins the
	// conference. Defaults to "false".
	StartConferenceOnEnter param.Opt[bool] `json:"start_conference_on_enter,omitzero"`
	// Whether a beep sound should be played when the participant joins and/or leaves
	// the conference. Can be used to override the conference-level setting.
	//
	// Any of "always", "never", "on_enter", "on_exit".
	BeepEnabled string `json:"beep_enabled,omitzero"`
	// Sets the joining participant as a supervisor for the conference. A conference
	// can have multiple supervisors. "barge" means the supervisor enters the
	// conference as a normal participant. This is the same as "none". "monitor" means
	// the supervisor is muted but can hear all participants. "whisper" means that only
	// the specified "whisper_call_control_ids" can hear the supervisor. Defaults to
	// "none".
	//
	// Any of "barge", "monitor", "none", "whisper".
	SupervisorRole string `json:"supervisor_role,omitzero"`
	// Array of unique call_control_ids the joining supervisor can whisper to. If none
	// provided, the supervisor will join the conference as a monitoring participant
	// only.
	WhisperCallControlIDs []string `json:"whisper_call_control_ids,omitzero"`
	paramObj
}

func (r CallDialParamsConferenceConfig) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsConferenceConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsConferenceConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallDialParamsConferenceConfig](
		"beep_enabled", "always", "never", "on_enter", "on_exit",
	)
	apijson.RegisterFieldValidator[CallDialParamsConferenceConfig](
		"supervisor_role", "barge", "monitor", "none", "whisper",
	)
}

// Starts a Conversation Relay session automatically when the answered/dialed call
// is answered. This embedded shape is supported on `answer` and `dial`. It uses
// public field names (`url`, `dtmf_detection`, `greeting`, `voice`, `language`,
// etc.) and maps them to the underlying Conversation Relay action. `client_state`,
// `tts_language`, and `transcription_language` inside this object are ignored; use
// the parent command's `client_state` and `command_id` fields instead.
//
// The property URL is required.
type CallDialParamsConversationRelayConfig struct {
	// WebSocket URL for your Conversation Relay server. Must start with `ws://` or
	// `wss://`.
	URL string `json:"url" api:"required"`
	// Enable DTMF detection for the relay session.
	DtmfDetection param.Opt[bool] `json:"dtmf_detection,omitzero"`
	// Text played when the relay session starts.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// Default language for both text-to-speech and speech recognition.
	Language param.Opt[string] `json:"language,omitzero"`
	// Structured voice provider. Must be supplied together with `structured_provider`.
	Provider param.Opt[string] `json:"provider,omitzero"`
	// Text-to-speech provider. If omitted, Telnyx derives it from `voice` or
	// `provider`.
	TtsProvider param.Opt[string] `json:"tts_provider,omitzero"`
	// The voice to be used by the voice assistant. Currently we support ElevenLabs,
	// Telnyx and AWS voices.
	//
	// **Supported Providers:**
	//
	//   - **AWS:** Use `AWS.Polly.<VoiceId>` (e.g., `AWS.Polly.Joanna`). For neural
	//     voices, which provide more realistic, human-like speech, append `-Neural` to
	//     the `VoiceId` (e.g., `AWS.Polly.Joanna-Neural`). Check the
	//     [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html)
	//     for compatibility.
	//   - **Azure:** Use `Azure.<VoiceId>. (e.g. Azure.en-CA-ClaraNeural,
	//     Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural,
	//     Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to
	//     [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).)
	//   - **ElevenLabs:** Use `ElevenLabs.<ModelId>.<VoiceId>` (e.g.,
	//     `ElevenLabs.BaseModel.John`). The `ModelId` part is optional. To use
	//     ElevenLabs, you must provide your ElevenLabs API key as an integration secret
	//     under `"voice_settings": {"api_key_ref": "<secret_id>"}`. See
	//     [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	//     for details. Check
	//     [available voices](https://elevenlabs.io/docs/api-reference/get-voices).
	//   - **Telnyx:** Use `Telnyx.<model_id>.<voice_id>`
	//   - **Inworld:** Use `Inworld.<ModelId>.<VoiceId>` (e.g., `Inworld.Mini.Loretta`,
	//     `Inworld.Max.Oliver`). Supported models: `Mini`, `Max`.
	//   - **xAI:** Use `xAI.<VoiceId>` (e.g., `xAI.eve`). Available voices: `eve`,
	//     `ara`, `rex`, `sal`, `leo`.
	Voice param.Opt[string] `json:"voice,omitzero"`
	// Custom key-value parameters forwarded to the relay session as assistant dynamic
	// variables.
	CustomParameters map[string]any `json:"custom_parameters,omitzero"`
	// Controls when caller input can interrupt assistant speech. `any` allows speech
	// or DTMF interruptions; `none` disables interruptions; `speech` allows speech
	// only; `dtmf` allows DTMF only.
	//
	// Any of "none", "any", "speech", "dtmf".
	Interruptible string `json:"interruptible,omitzero"`
	// Controls when caller input can interrupt assistant speech. `any` allows speech
	// or DTMF interruptions; `none` disables interruptions; `speech` allows speech
	// only; `dtmf` allows DTMF only.
	//
	// Any of "none", "any", "speech", "dtmf".
	InterruptibleGreeting string `json:"interruptible_greeting,omitzero"`
	// Settings for handling caller interruptions during Conversation Relay speech.
	InterruptionSettings CallDialParamsConversationRelayConfigInterruptionSettings `json:"interruption_settings,omitzero"`
	// Per-language TTS and transcription settings.
	Languages []CallDialParamsConversationRelayConfigLanguage `json:"languages,omitzero"`
	// Provider-specific structured voice settings. Must be supplied together with
	// `provider`; Telnyx sends the value as the nested provider configuration for
	// Conversation Relay.
	StructuredProvider map[string]any `json:"structured_provider,omitzero"`
	// Engine to use for speech recognition. Legacy values `A` - `Google`, `B` -
	// `Telnyx` are supported for backward compatibility. For Conversation Relay, use
	// this field with `transcription_engine_config`; the `transcription` object is not
	// supported.
	//
	// Any of "Google", "Telnyx", "Deepgram", "Azure", "xAI", "AssemblyAI",
	// "Speechmatics", "Soniox", "A", "B".
	TranscriptionEngine string `json:"transcription_engine,omitzero"`
	// Engine-specific transcription settings for Conversation Relay. This accepts the
	// same provider-specific options used by the Call Transcription Start command,
	// such as `transcription_model`, without requiring the engine discriminator to be
	// repeated inside this object.
	TranscriptionEngineConfig map[string]any `json:"transcription_engine_config,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings CallDialParamsConversationRelayConfigVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r CallDialParamsConversationRelayConfig) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsConversationRelayConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsConversationRelayConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfig](
		"interruptible", "none", "any", "speech", "dtmf",
	)
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfig](
		"interruptible_greeting", "none", "any", "speech", "dtmf",
	)
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfig](
		"transcription_engine", "Google", "Telnyx", "Deepgram", "Azure", "xAI", "AssemblyAI", "Speechmatics", "Soniox", "A", "B",
	)
}

// Settings for handling caller interruptions during Conversation Relay speech.
type CallDialParamsConversationRelayConfigInterruptionSettings struct {
	// Legacy boolean form. `true` is equivalent to `interruptible=any`; `false` is
	// equivalent to `interruptible=none`.
	Enable param.Opt[bool] `json:"enable,omitzero"`
	// Controls when caller input can interrupt assistant speech. `any` allows speech
	// or DTMF interruptions; `none` disables interruptions; `speech` allows speech
	// only; `dtmf` allows DTMF only.
	//
	// Any of "none", "any", "speech", "dtmf".
	Interruptible string `json:"interruptible,omitzero"`
	// Controls when caller input can interrupt assistant speech. `any` allows speech
	// or DTMF interruptions; `none` disables interruptions; `speech` allows speech
	// only; `dtmf` allows DTMF only.
	//
	// Any of "none", "any", "speech", "dtmf".
	InterruptibleGreeting string `json:"interruptible_greeting,omitzero"`
	// Controls when caller input can interrupt assistant speech. `any` allows speech
	// or DTMF interruptions; `none` disables interruptions; `speech` allows speech
	// only; `dtmf` allows DTMF only.
	//
	// Any of "none", "any", "speech", "dtmf".
	WelcomeGreetingInterruptible string `json:"welcome_greeting_interruptible,omitzero"`
	paramObj
}

func (r CallDialParamsConversationRelayConfigInterruptionSettings) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsConversationRelayConfigInterruptionSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsConversationRelayConfigInterruptionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfigInterruptionSettings](
		"interruptible", "none", "any", "speech", "dtmf",
	)
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfigInterruptionSettings](
		"interruptible_greeting", "none", "any", "speech", "dtmf",
	)
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfigInterruptionSettings](
		"welcome_greeting_interruptible", "none", "any", "speech", "dtmf",
	)
}

// Language-specific TTS and transcription settings for Conversation Relay.
//
// The property Language is required.
type CallDialParamsConversationRelayConfigLanguage struct {
	// BCP 47 language tag for this language configuration.
	Language string `json:"language" api:"required"`
	// Conversation Relay speech model. Prefer
	// `transcription_engine_config.transcription_model` when configuring
	// speech-to-text.
	SpeechModel param.Opt[string] `json:"speech_model,omitzero"`
	// Conversation Relay transcription provider name. Prefer `transcription_engine`
	// when configuring speech-to-text.
	TranscriptionProvider param.Opt[string] `json:"transcription_provider,omitzero"`
	// Text-to-speech provider for this language. If omitted and `voice` is provided,
	// Telnyx derives the provider from the voice identifier.
	TtsProvider param.Opt[string] `json:"tts_provider,omitzero"`
	// Voice identifier for this language.
	Voice param.Opt[string] `json:"voice,omitzero"`
	// Engine to use for speech recognition. Legacy values `A` - `Google`, `B` -
	// `Telnyx` are supported for backward compatibility. When provided in a
	// Conversation Relay language entry, Telnyx derives `transcription_provider` and
	// `speech_model` for that language.
	//
	// Any of "Google", "Telnyx", "Deepgram", "Azure", "xAI", "AssemblyAI",
	// "Speechmatics", "Soniox", "A", "B".
	TranscriptionEngine string `json:"transcription_engine,omitzero"`
	// Engine-specific transcription settings for Conversation Relay. This accepts the
	// same provider-specific options used by the Call Transcription Start command,
	// such as `transcription_model`, without requiring the engine discriminator to be
	// repeated inside this object.
	TranscriptionEngineConfig map[string]any `json:"transcription_engine_config,omitzero"`
	// The settings associated with the voice selected
	VoiceSettings CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion `json:"voice_settings,omitzero"`
	paramObj
}

func (r CallDialParamsConversationRelayConfigLanguage) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsConversationRelayConfigLanguage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsConversationRelayConfigLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CallDialParamsConversationRelayConfigLanguage](
		"transcription_engine", "Google", "Telnyx", "Deepgram", "Azure", "xAI", "AssemblyAI", "Speechmatics", "Soniox", "A", "B",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion struct {
	OfElevenlabs *ElevenLabsVoiceSettingsParam                                      `json:",omitzero,inline"`
	OfTelnyx     *TelnyxVoiceSettingsParam                                          `json:",omitzero,inline"`
	OfAws        *AwsVoiceSettingsParam                                             `json:",omitzero,inline"`
	OfMinimax    *shared.MinimaxVoiceSettingsParam                                  `json:",omitzero,inline"`
	OfAzure      *shared.AzureVoiceSettingsParam                                    `json:",omitzero,inline"`
	OfRime       *shared.RimeVoiceSettingsParam                                     `json:",omitzero,inline"`
	OfResemble   *shared.ResembleVoiceSettingsParam                                 `json:",omitzero,inline"`
	OfInworld    *CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld `json:",omitzero,inline"`
	OfXai        *shared.XaiVoiceSettingsParam                                      `json:",omitzero,inline"`
	paramUnion
}

func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenlabs,
		u.OfTelnyx,
		u.OfAws,
		u.OfMinimax,
		u.OfAzure,
		u.OfRime,
		u.OfResemble,
		u.OfInworld,
		u.OfXai)
}
func (u *CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenlabs) {
		return u.OfElevenlabs
	} else if !param.IsOmitted(u.OfTelnyx) {
		return u.OfTelnyx
	} else if !param.IsOmitted(u.OfAws) {
		return u.OfAws
	} else if !param.IsOmitted(u.OfMinimax) {
		return u.OfMinimax
	} else if !param.IsOmitted(u.OfAzure) {
		return u.OfAzure
	} else if !param.IsOmitted(u.OfRime) {
		return u.OfRime
	} else if !param.IsOmitted(u.OfResemble) {
		return u.OfResemble
	} else if !param.IsOmitted(u.OfInworld) {
		return u.OfInworld
	} else if !param.IsOmitted(u.OfXai) {
		return u.OfXai
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetLanguageBoost() *string {
	if vt := u.OfMinimax; vt != nil {
		return (*string)(&vt.LanguageBoost)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetPitch() *int64 {
	if vt := u.OfMinimax; vt != nil && vt.Pitch.Valid() {
		return &vt.Pitch.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetSpeed() *float64 {
	if vt := u.OfMinimax; vt != nil && vt.Speed.Valid() {
		return &vt.Speed.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetVol() *float64 {
	if vt := u.OfMinimax; vt != nil && vt.Vol.Valid() {
		return &vt.Vol.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetDeploymentID() *string {
	if vt := u.OfAzure; vt != nil && vt.DeploymentID.Valid() {
		return &vt.DeploymentID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetEffect() *string {
	if vt := u.OfAzure; vt != nil {
		return (*string)(&vt.Effect)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetGender() *string {
	if vt := u.OfAzure; vt != nil {
		return (*string)(&vt.Gender)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetRegion() *string {
	if vt := u.OfAzure; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetFormat() *string {
	if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetPrecision() *string {
	if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.Precision)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetSampleRate() *string {
	if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.SampleRate)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetLanguage() *string {
	if vt := u.OfXai; vt != nil && vt.Language.Valid() {
		return &vt.Language.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetType() *string {
	if vt := u.OfElevenlabs; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTelnyx; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAws; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMinimax; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAzure; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInworld; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfXai; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetAPIKeyRef() *string {
	if vt := u.OfElevenlabs; vt != nil && vt.APIKeyRef.Valid() {
		return &vt.APIKeyRef.Value
	} else if vt := u.OfAzure; vt != nil && vt.APIKeyRef.Valid() {
		return &vt.APIKeyRef.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion) GetVoiceSpeed() *float64 {
	if vt := u.OfTelnyx; vt != nil && vt.VoiceSpeed.Valid() {
		return &vt.VoiceSpeed.Value
	} else if vt := u.OfRime; vt != nil && vt.VoiceSpeed.Valid() {
		return &vt.VoiceSpeed.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[CallDialParamsConversationRelayConfigLanguageVoiceSettingsUnion](
		"type",
		apijson.Discriminator[ElevenLabsVoiceSettingsParam]("elevenlabs"),
		apijson.Discriminator[TelnyxVoiceSettingsParam]("telnyx"),
		apijson.Discriminator[AwsVoiceSettingsParam]("aws"),
		apijson.Discriminator[shared.MinimaxVoiceSettingsParam]("minimax"),
		apijson.Discriminator[shared.AzureVoiceSettingsParam]("azure"),
		apijson.Discriminator[shared.RimeVoiceSettingsParam]("rime"),
		apijson.Discriminator[shared.ResembleVoiceSettingsParam]("resemble"),
		apijson.Discriminator[CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld]("inworld"),
		apijson.Discriminator[shared.XaiVoiceSettingsParam]("xai"),
	)
}

func NewCallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld() CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld {
	return CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld{
		Type: "inworld",
	}
}

// This struct has a constant value, construct it with
// [NewCallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld].
type CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld struct {
	// Voice settings provider type
	Type constant.Inworld `json:"type" default:"inworld"`
	paramObj
}

func (r CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsConversationRelayConfigLanguageVoiceSettingsInworld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CallDialParamsConversationRelayConfigVoiceSettingsUnion struct {
	OfElevenlabs *ElevenLabsVoiceSettingsParam                              `json:",omitzero,inline"`
	OfTelnyx     *TelnyxVoiceSettingsParam                                  `json:",omitzero,inline"`
	OfAws        *AwsVoiceSettingsParam                                     `json:",omitzero,inline"`
	OfMinimax    *shared.MinimaxVoiceSettingsParam                          `json:",omitzero,inline"`
	OfAzure      *shared.AzureVoiceSettingsParam                            `json:",omitzero,inline"`
	OfRime       *shared.RimeVoiceSettingsParam                             `json:",omitzero,inline"`
	OfResemble   *shared.ResembleVoiceSettingsParam                         `json:",omitzero,inline"`
	OfInworld    *CallDialParamsConversationRelayConfigVoiceSettingsInworld `json:",omitzero,inline"`
	OfXai        *shared.XaiVoiceSettingsParam                              `json:",omitzero,inline"`
	paramUnion
}

func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfElevenlabs,
		u.OfTelnyx,
		u.OfAws,
		u.OfMinimax,
		u.OfAzure,
		u.OfRime,
		u.OfResemble,
		u.OfInworld,
		u.OfXai)
}
func (u *CallDialParamsConversationRelayConfigVoiceSettingsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CallDialParamsConversationRelayConfigVoiceSettingsUnion) asAny() any {
	if !param.IsOmitted(u.OfElevenlabs) {
		return u.OfElevenlabs
	} else if !param.IsOmitted(u.OfTelnyx) {
		return u.OfTelnyx
	} else if !param.IsOmitted(u.OfAws) {
		return u.OfAws
	} else if !param.IsOmitted(u.OfMinimax) {
		return u.OfMinimax
	} else if !param.IsOmitted(u.OfAzure) {
		return u.OfAzure
	} else if !param.IsOmitted(u.OfRime) {
		return u.OfRime
	} else if !param.IsOmitted(u.OfResemble) {
		return u.OfResemble
	} else if !param.IsOmitted(u.OfInworld) {
		return u.OfInworld
	} else if !param.IsOmitted(u.OfXai) {
		return u.OfXai
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetLanguageBoost() *string {
	if vt := u.OfMinimax; vt != nil {
		return (*string)(&vt.LanguageBoost)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetPitch() *int64 {
	if vt := u.OfMinimax; vt != nil && vt.Pitch.Valid() {
		return &vt.Pitch.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetSpeed() *float64 {
	if vt := u.OfMinimax; vt != nil && vt.Speed.Valid() {
		return &vt.Speed.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetVol() *float64 {
	if vt := u.OfMinimax; vt != nil && vt.Vol.Valid() {
		return &vt.Vol.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetDeploymentID() *string {
	if vt := u.OfAzure; vt != nil && vt.DeploymentID.Valid() {
		return &vt.DeploymentID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetEffect() *string {
	if vt := u.OfAzure; vt != nil {
		return (*string)(&vt.Effect)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetGender() *string {
	if vt := u.OfAzure; vt != nil {
		return (*string)(&vt.Gender)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetRegion() *string {
	if vt := u.OfAzure; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetFormat() *string {
	if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetPrecision() *string {
	if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.Precision)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetSampleRate() *string {
	if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.SampleRate)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetLanguage() *string {
	if vt := u.OfXai; vt != nil && vt.Language.Valid() {
		return &vt.Language.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetType() *string {
	if vt := u.OfElevenlabs; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTelnyx; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAws; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMinimax; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAzure; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRime; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfResemble; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInworld; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfXai; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetAPIKeyRef() *string {
	if vt := u.OfElevenlabs; vt != nil && vt.APIKeyRef.Valid() {
		return &vt.APIKeyRef.Value
	} else if vt := u.OfAzure; vt != nil && vt.APIKeyRef.Valid() {
		return &vt.APIKeyRef.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CallDialParamsConversationRelayConfigVoiceSettingsUnion) GetVoiceSpeed() *float64 {
	if vt := u.OfTelnyx; vt != nil && vt.VoiceSpeed.Valid() {
		return &vt.VoiceSpeed.Value
	} else if vt := u.OfRime; vt != nil && vt.VoiceSpeed.Valid() {
		return &vt.VoiceSpeed.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[CallDialParamsConversationRelayConfigVoiceSettingsUnion](
		"type",
		apijson.Discriminator[ElevenLabsVoiceSettingsParam]("elevenlabs"),
		apijson.Discriminator[TelnyxVoiceSettingsParam]("telnyx"),
		apijson.Discriminator[AwsVoiceSettingsParam]("aws"),
		apijson.Discriminator[shared.MinimaxVoiceSettingsParam]("minimax"),
		apijson.Discriminator[shared.AzureVoiceSettingsParam]("azure"),
		apijson.Discriminator[shared.RimeVoiceSettingsParam]("rime"),
		apijson.Discriminator[shared.ResembleVoiceSettingsParam]("resemble"),
		apijson.Discriminator[CallDialParamsConversationRelayConfigVoiceSettingsInworld]("inworld"),
		apijson.Discriminator[shared.XaiVoiceSettingsParam]("xai"),
	)
}

func NewCallDialParamsConversationRelayConfigVoiceSettingsInworld() CallDialParamsConversationRelayConfigVoiceSettingsInworld {
	return CallDialParamsConversationRelayConfigVoiceSettingsInworld{
		Type: "inworld",
	}
}

// This struct has a constant value, construct it with
// [NewCallDialParamsConversationRelayConfigVoiceSettingsInworld].
type CallDialParamsConversationRelayConfigVoiceSettingsInworld struct {
	// Voice settings provider type
	Type constant.Inworld `json:"type" default:"inworld"`
	paramObj
}

func (r CallDialParamsConversationRelayConfigVoiceSettingsInworld) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsConversationRelayConfigVoiceSettingsInworld
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsConversationRelayConfigVoiceSettingsInworld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enables deepfake detection on the call. When enabled, audio from the remote
// party is streamed to a detection service that analyzes whether the voice is
// AI-generated. Results are delivered via the `call.deepfake_detection.result`
// webhook.
//
// The property Enabled is required.
type CallDialParamsDeepfakeDetection struct {
	// Whether deepfake detection is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Maximum time in seconds to wait for RTP audio before timing out. If no audio is
	// received within this window, detection stops with an error.
	RtpTimeout param.Opt[int64] `json:"rtp_timeout,omitzero"`
	// Maximum time in seconds to wait for a detection result before timing out.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	paramObj
}

func (r CallDialParamsDeepfakeDetection) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsDeepfakeDetection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsDeepfakeDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines whether media should be encrypted on the call. For SIP URI destinations,
// media encryption can also be requested per endpoint with the `secure` URI
// parameter: `;secure=true` or `;secure=srtp` enables SRTP, and `;secure=dtls`
// enables DTLS. This parameter, when set to `SRTP` or `DTLS`, takes precedence
// over the per-endpoint `secure` value.
type CallDialParamsMediaEncryption string

const (
	CallDialParamsMediaEncryptionDisabled CallDialParamsMediaEncryption = "disabled"
	CallDialParamsMediaEncryptionSrtp     CallDialParamsMediaEncryption = "SRTP"
	CallDialParamsMediaEncryptionDtls     CallDialParamsMediaEncryption = "DTLS"
)

// Indicates the privacy level to be used for the call. When set to `id`, caller ID
// information (name and number) will be hidden from the called party. When set to
// `none` or omitted, caller ID will be shown normally.
type CallDialParamsPrivacy string

const (
	CallDialParamsPrivacyID   CallDialParamsPrivacy = "id"
	CallDialParamsPrivacyNone CallDialParamsPrivacy = "none"
)

// Start recording automatically after an event. Disabled by default.
type CallDialParamsRecord string

const (
	CallDialParamsRecordRecordFromAnswer CallDialParamsRecord = "record-from-answer"
)

// Defines which channel should be recorded ('single' or 'dual') when `record` is
// specified.
type CallDialParamsRecordChannels string

const (
	CallDialParamsRecordChannelsSingle CallDialParamsRecordChannels = "single"
	CallDialParamsRecordChannelsDual   CallDialParamsRecordChannels = "dual"
)

// Defines the format of the recording ('wav' or 'mp3') when `record` is specified.
type CallDialParamsRecordFormat string

const (
	CallDialParamsRecordFormatWav CallDialParamsRecordFormat = "wav"
	CallDialParamsRecordFormatMP3 CallDialParamsRecordFormat = "mp3"
)

// The audio track to be recorded. Can be either `both`, `inbound` or `outbound`.
// If only single track is specified (`inbound`, `outbound`), `channels`
// configuration is ignored and it will be recorded as mono (single channel).
type CallDialParamsRecordTrack string

const (
	CallDialParamsRecordTrackBoth     CallDialParamsRecordTrack = "both"
	CallDialParamsRecordTrackInbound  CallDialParamsRecordTrack = "inbound"
	CallDialParamsRecordTrackOutbound CallDialParamsRecordTrack = "outbound"
)

// When set to `trim-silence`, silence will be removed from the beginning and end
// of the recording.
type CallDialParamsRecordTrim string

const (
	CallDialParamsRecordTrimTrimSilence CallDialParamsRecordTrim = "trim-silence"
)

// Defines the SIP region to be used for the call.
type CallDialParamsSipRegion string

const (
	CallDialParamsSipRegionUs         CallDialParamsSipRegion = "US"
	CallDialParamsSipRegionEurope     CallDialParamsSipRegion = "Europe"
	CallDialParamsSipRegionCanada     CallDialParamsSipRegion = "Canada"
	CallDialParamsSipRegionAustralia  CallDialParamsSipRegion = "Australia"
	CallDialParamsSipRegionMiddleEast CallDialParamsSipRegion = "Middle East"
)

// Defines SIP transport protocol to be used on the call.
type CallDialParamsSipTransportProtocol string

const (
	CallDialParamsSipTransportProtocolUdp CallDialParamsSipTransportProtocol = "UDP"
	CallDialParamsSipTransportProtocolTcp CallDialParamsSipTransportProtocol = "TCP"
	CallDialParamsSipTransportProtocolTls CallDialParamsSipTransportProtocol = "TLS"
)

// Specifies which track should be streamed.
type CallDialParamsStreamTrack string

const (
	CallDialParamsStreamTrackInboundTrack  CallDialParamsStreamTrack = "inbound_track"
	CallDialParamsStreamTrackOutboundTrack CallDialParamsStreamTrack = "outbound_track"
	CallDialParamsStreamTrackBothTracks    CallDialParamsStreamTrack = "both_tracks"
)

// The role of the supervisor call. 'barge' means that supervisor call hears and is
// being heard by both ends of the call (caller & callee). 'whisper' means that
// only supervised_call_control_id hears supervisor but supervisor can hear
// everything. 'monitor' means that nobody can hear supervisor call, but supervisor
// can hear everything on the call.
type CallDialParamsSupervisorRole string

const (
	CallDialParamsSupervisorRoleBarge   CallDialParamsSupervisorRole = "barge"
	CallDialParamsSupervisorRoleWhisper CallDialParamsSupervisorRole = "whisper"
	CallDialParamsSupervisorRoleMonitor CallDialParamsSupervisorRole = "monitor"
)

type CallDialParamsWebhookRetriesPolicy struct {
	// Array of delays in milliseconds between retry attempts. Total sum cannot exceed
	// 60000ms.
	RetriesMs []int64 `json:"retries_ms,omitzero"`
	paramObj
}

func (r CallDialParamsWebhookRetriesPolicy) MarshalJSON() (data []byte, err error) {
	type shadow CallDialParamsWebhookRetriesPolicy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CallDialParamsWebhookRetriesPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP request type used for `webhook_url`.
type CallDialParamsWebhookURLMethod string

const (
	CallDialParamsWebhookURLMethodPost CallDialParamsWebhookURLMethod = "POST"
	CallDialParamsWebhookURLMethodGet  CallDialParamsWebhookURLMethod = "GET"
)

// HTTP request method to invoke `webhook_urls`.
type CallDialParamsWebhookURLsMethod string

const (
	CallDialParamsWebhookURLsMethodPost CallDialParamsWebhookURLsMethod = "POST"
	CallDialParamsWebhookURLsMethodGet  CallDialParamsWebhookURLsMethod = "GET"
)
