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
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// CallService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
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
//   - `call.machine.detection.ended` if `answering_machine_detection` was requested
//   - `call.machine.greeting.ended` if `answering_machine_detection` was requested
//     to detect the end of machine greeting
//   - `call.machine.premium.detection.ended` if
//     `answering_machine_detection=premium` was requested
//   - `call.machine.premium.greeting.ended` if `answering_machine_detection=premium`
//     was requested and a beep was detected
//   - `streaming.started`, `streaming.stopped` or `streaming.failed` if `stream_url`
//     was set
//
// When the `record` parameter is set to `record-from-answer`, the response will
// include a `recording_id` field.
func (r *CallService) Dial(ctx context.Context, body CallDialParams, opts ...option.RequestOption) (res *CallDialResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "calls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the status of a call (data is available 10 minutes after call ended).
func (r *CallService) GetStatus(ctx context.Context, callControlID string, opts ...option.RequestOption) (res *CallGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if callControlID == "" {
		err = errors.New("missing required call_control_id parameter")
		return
	}
	path := fmt.Sprintf("calls/%s", callControlID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomSipHeader struct {
	// The name of the header to add.
	Name string `json:"name,required"`
	// The value of the header.
	Value string `json:"value,required"`
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
	Name string `json:"name,required"`
	// The value of the header.
	Value string `json:"value,required"`
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
	Name SipHeaderName `json:"name,required"`
	// The value of the header.
	Value string `json:"value,required"`
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
	Name SipHeaderName `json:"name,omitzero,required"`
	// The value of the header.
	Value string `json:"value,required"`
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
	CallControlID string `json:"call_control_id,required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id,required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id,required"`
	// Indicates whether the call is alive or not. For Dial command it will always be
	// `false` (dialing is asynchronous).
	IsAlive bool `json:"is_alive,required"`
	// Any of "call".
	RecordType string `json:"record_type,required"`
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
	CallControlID string `json:"call_control_id,required"`
	// ID that is unique to the call and can be used to correlate webhook events
	CallLegID string `json:"call_leg_id,required"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call
	CallSessionID string `json:"call_session_id,required"`
	// Indicates whether the call is alive or not. For Dial command it will always be
	// `false` (dialing is asynchronous).
	IsAlive bool `json:"is_alive,required"`
	// Any of "call".
	RecordType string `json:"record_type,required"`
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
	ConnectionID string `json:"connection_id,required"`
	// The `from` number to be used as the caller id presented to the destination (`to`
	// number). The number should be in +E164 format.
	From string `json:"from,required"`
	// The DID or SIP URI to dial out to. Multiple DID or SIP URIs can be provided
	// using an array of strings
	To CallDialParamsToUnion `json:"to,omitzero,required"`
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
	// The list of comma-separated codecs in a preferred order for the forked media to
	// be received.
	PreferredCodecs param.Opt[string] `json:"preferred_codecs,omitzero"`
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
	// Generate silence RTP packets when no transmission available.
	SendSilenceWhenIdle param.Opt[bool] `json:"send_silence_when_idle,omitzero"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword param.Opt[string] `json:"sip_auth_password,omitzero"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername param.Opt[string] `json:"sip_auth_username,omitzero"`
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
	// performance.
	AnsweringMachineDetectionConfig CallDialParamsAnsweringMachineDetectionConfig `json:"answering_machine_detection_config,omitzero"`
	// Optional configuration parameters to dial new participant into a conference.
	ConferenceConfig CallDialParamsConferenceConfig `json:"conference_config,omitzero"`
	// Custom headers to be added to the SIP INVITE.
	CustomHeaders    []CustomSipHeaderParam `json:"custom_headers,omitzero"`
	DialogflowConfig DialogflowConfigParam  `json:"dialogflow_config,omitzero"`
	// Defines whether media should be encrypted on the call.
	//
	// Any of "disabled", "SRTP", "DTLS".
	MediaEncryption CallDialParamsMediaEncryption `json:"media_encryption,omitzero"`
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
	// HTTP request type used for `webhook_url`.
	//
	// Any of "POST", "GET".
	WebhookURLMethod CallDialParamsWebhookURLMethod `json:"webhook_url_method,omitzero"`
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
// performance.
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

// Defines whether media should be encrypted on the call.
type CallDialParamsMediaEncryption string

const (
	CallDialParamsMediaEncryptionDisabled CallDialParamsMediaEncryption = "disabled"
	CallDialParamsMediaEncryptionSrtp     CallDialParamsMediaEncryption = "SRTP"
	CallDialParamsMediaEncryptionDtls     CallDialParamsMediaEncryption = "DTLS"
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

// HTTP request type used for `webhook_url`.
type CallDialParamsWebhookURLMethod string

const (
	CallDialParamsWebhookURLMethodPost CallDialParamsWebhookURLMethod = "POST"
	CallDialParamsWebhookURLMethodGet  CallDialParamsWebhookURLMethod = "GET"
)
