// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// WebhookService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEventUnion, error) {
	res := &UnsafeUnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.PublicKey
	if key == "" {
		return nil, errors.New("The PublicKey option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type CallStreamingFailed struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "streaming.failed".
	EventType CallStreamingFailedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                  `json:"occurred_at" format:"date-time"`
	Payload    CallStreamingFailedPayload `json:"payload"`
	// Identifies the resource.
	//
	// Any of "event".
	RecordType CallStreamingFailedRecordType `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingFailed) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallStreamingFailedEventType string

const (
	CallStreamingFailedEventTypeStreamingFailed CallStreamingFailedEventType = "streaming.failed"
)

type CallStreamingFailedPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// A short description explaning why the media streaming failed.
	FailureReason string `json:"failure_reason"`
	// Identifies the streaming.
	StreamID string `json:"stream_id" format:"uuid"`
	// Streaming parameters as they were originally given to the Call Control API.
	StreamParams CallStreamingFailedPayloadStreamParams `json:"stream_params"`
	// The type of stream connection the stream is performing.
	//
	// Any of "websocket", "dialogflow".
	StreamType string `json:"stream_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		FailureReason respjson.Field
		StreamID      respjson.Field
		StreamParams  respjson.Field
		StreamType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingFailedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingFailedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Streaming parameters as they were originally given to the Call Control API.
type CallStreamingFailedPayloadStreamParams struct {
	// The destination WebSocket address where the stream is going to be delivered.
	StreamURL string `json:"stream_url"`
	// Specifies which track should be streamed.
	//
	// Any of "inbound_track", "outbound_track", "both_tracks".
	Track string `json:"track"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StreamURL   respjson.Field
		Track       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingFailedPayloadStreamParams) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingFailedPayloadStreamParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the resource.
type CallStreamingFailedRecordType string

const (
	CallStreamingFailedRecordTypeEvent CallStreamingFailedRecordType = "event"
)

type CallStreamingStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "streaming.started".
	EventType CallStreamingStartedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                   `json:"occurred_at" format:"date-time"`
	Payload    CallStreamingStartedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallStreamingStartedRecordType `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingStarted) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallStreamingStartedEventType string

const (
	CallStreamingStartedEventTypeStreamingStarted CallStreamingStartedEventType = "streaming.started"
)

type CallStreamingStartedPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Destination WebSocket address where the stream is going to be delivered.
	StreamURL string `json:"stream_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		StreamURL     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallStreamingStartedRecordType string

const (
	CallStreamingStartedRecordTypeEvent CallStreamingStartedRecordType = "event"
)

type CallStreamingStopped struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "streaming.stopped".
	EventType CallStreamingStoppedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                   `json:"occurred_at" format:"date-time"`
	Payload    CallStreamingStoppedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallStreamingStoppedRecordType `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingStopped) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingStopped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallStreamingStoppedEventType string

const (
	CallStreamingStoppedEventTypeStreamingStopped CallStreamingStoppedEventType = "streaming.stopped"
)

type CallStreamingStoppedPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Destination WebSocket address where the stream is going to be delivered.
	StreamURL string `json:"stream_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		StreamURL     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingStoppedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingStoppedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallStreamingStoppedRecordType string

const (
	CallStreamingStoppedRecordTypeEvent CallStreamingStoppedRecordType = "event"
)

type CallAIGatherEndedWebhookEvent struct {
	Data CallAIGatherEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.ai_gather.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    CallAIGatherEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Telnyx connection ID used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// The history of the messages exchanged during the AI gather
	MessageHistory []CallAIGatherEndedWebhookEventDataPayloadMessageHistory `json:"message_history"`
	// The result of the AI gather, its type depends of the `parameters` provided in
	// the command
	Result map[string]any `json:"result"`
	// Reflects how command ended.
	//
	// Any of "valid", "invalid".
	Status string `json:"status"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID  respjson.Field
		CallLegID      respjson.Field
		CallSessionID  respjson.Field
		ClientState    respjson.Field
		ConnectionID   respjson.Field
		From           respjson.Field
		MessageHistory respjson.Field
		Result         respjson.Field
		Status         respjson.Field
		To             respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherEndedWebhookEventDataPayloadMessageHistory struct {
	// The content of the message
	Content string `json:"content"`
	// The role of the message sender
	//
	// Any of "assistant", "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherEndedWebhookEventDataPayloadMessageHistory) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEndedWebhookEventDataPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherMessageHistoryUpdatedWebhookEvent struct {
	Data CallAIGatherMessageHistoryUpdatedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherMessageHistoryUpdatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherMessageHistoryUpdatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherMessageHistoryUpdatedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.ai_gather.message_history_updated".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                                `json:"occurred_at" format:"date-time"`
	Payload    CallAIGatherMessageHistoryUpdatedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherMessageHistoryUpdatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherMessageHistoryUpdatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherMessageHistoryUpdatedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Telnyx connection ID used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// The history of the messages exchanged during the AI gather
	MessageHistory []CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory `json:"message_history"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID  respjson.Field
		CallLegID      respjson.Field
		CallSessionID  respjson.Field
		ClientState    respjson.Field
		ConnectionID   respjson.Field
		From           respjson.Field
		MessageHistory respjson.Field
		To             respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherMessageHistoryUpdatedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherMessageHistoryUpdatedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory struct {
	// The content of the message
	Content string `json:"content"`
	// The role of the message sender
	//
	// Any of "assistant", "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory) RawJSON() string {
	return r.JSON.raw
}
func (r *CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherPartialResultsWebhookEvent struct {
	Data CallAIGatherPartialResultsWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherPartialResultsWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherPartialResultsWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherPartialResultsWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.ai_gather.partial_results".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                         `json:"occurred_at" format:"date-time"`
	Payload    CallAIGatherPartialResultsWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherPartialResultsWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherPartialResultsWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherPartialResultsWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Telnyx connection ID used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// The history of the messages exchanged during the AI gather
	MessageHistory []CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory `json:"message_history"`
	// The partial result of the AI gather, its type depends of the `parameters`
	// provided in the command
	PartialResults map[string]any `json:"partial_results"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID  respjson.Field
		CallLegID      respjson.Field
		CallSessionID  respjson.Field
		ClientState    respjson.Field
		ConnectionID   respjson.Field
		From           respjson.Field
		MessageHistory respjson.Field
		PartialResults respjson.Field
		To             respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherPartialResultsWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherPartialResultsWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory struct {
	// The content of the message
	Content string `json:"content"`
	// The role of the message sender
	//
	// Any of "assistant", "user".
	Role string `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory) RawJSON() string {
	return r.JSON.raw
}
func (r *CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAnsweredWebhookEvent struct {
	Data CallAnsweredWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAnsweredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallAnsweredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAnsweredWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.answered".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                           `json:"occurred_at" format:"date-time"`
	Payload    CallAnsweredWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAnsweredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallAnsweredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAnsweredWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Custom headers set on answer command
	CustomHeaders []CustomSipHeader `json:"custom_headers"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// User-to-User and Diversion headers from sip invite.
	SipHeaders []SipHeader `json:"sip_headers"`
	// ISO 8601 datetime of when the call started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// State received from a command.
	//
	// Any of "answered".
	State string `json:"state"`
	// Array of tags associated to number.
	Tags []string `json:"tags"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		CustomHeaders respjson.Field
		From          respjson.Field
		SipHeaders    respjson.Field
		StartTime     respjson.Field
		State         respjson.Field
		Tags          respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallAnsweredWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAnsweredWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallBridgedWebhookEvent struct {
	Data CallBridgedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallBridgedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallBridgedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallBridgedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.bridged".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                          `json:"occurred_at" format:"date-time"`
	Payload    CallBridgedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallBridgedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallBridgedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallBridgedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		From          respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallBridgedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallBridgedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationEndedWebhookEvent struct {
	Data CallConversationEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallConversationEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationEndedWebhookEventData struct {
	// Unique identifier for the event.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the event was created in the system.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The type of event being delivered.
	//
	// Any of "call.conversation.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                    `json:"occurred_at" format:"date-time"`
	Payload    CallConversationEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallConversationEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationEndedWebhookEventDataPayload struct {
	// Unique identifier of the assistant involved in the call.
	AssistantID string `json:"assistant_id"`
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call leg.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session (group of related call legs).
	CallSessionID string `json:"call_session_id"`
	// The type of calling party connection.
	//
	// Any of "pstn", "sip".
	CallingPartyType string `json:"calling_party_type"`
	// Base64-encoded state received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID unique to the conversation or insight group generated for the call.
	ConversationID string `json:"conversation_id"`
	// Duration of the conversation in seconds.
	DurationSec int64 `json:"duration_sec"`
	// The caller's number or identifier.
	From string `json:"from"`
	// The large language model used during the conversation.
	LlmModel string `json:"llm_model"`
	// The speech-to-text model used in the conversation.
	SttModel string `json:"stt_model"`
	// The callee's number or SIP address.
	To string `json:"to"`
	// The model ID used for text-to-speech synthesis.
	TtsModelID string `json:"tts_model_id"`
	// The text-to-speech provider used in the call.
	TtsProvider string `json:"tts_provider"`
	// Voice ID used for TTS.
	TtsVoiceID string `json:"tts_voice_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantID      respjson.Field
		CallControlID    respjson.Field
		CallLegID        respjson.Field
		CallSessionID    respjson.Field
		CallingPartyType respjson.Field
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		ConversationID   respjson.Field
		DurationSec      respjson.Field
		From             respjson.Field
		LlmModel         respjson.Field
		SttModel         respjson.Field
		To               respjson.Field
		TtsModelID       respjson.Field
		TtsProvider      respjson.Field
		TtsVoiceID       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallConversationEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationInsightsGeneratedWebhookEvent struct {
	Data CallConversationInsightsGeneratedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationInsightsGeneratedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallConversationInsightsGeneratedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationInsightsGeneratedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.conversation_insights.generated".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                                `json:"occurred_at" format:"date-time"`
	Payload    CallConversationInsightsGeneratedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationInsightsGeneratedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallConversationInsightsGeneratedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationInsightsGeneratedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// The type of calling party connection.
	//
	// Any of "pstn", "sip".
	CallingPartyType string `json:"calling_party_type"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the insight group being generated for the call.
	InsightGroupID string `json:"insight_group_id"`
	// Array of insight results being generated for the call.
	Results []CallConversationInsightsGeneratedWebhookEventDataPayloadResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID    respjson.Field
		CallLegID        respjson.Field
		CallSessionID    respjson.Field
		CallingPartyType respjson.Field
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		InsightGroupID   respjson.Field
		Results          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationInsightsGeneratedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallConversationInsightsGeneratedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationInsightsGeneratedWebhookEventDataPayloadResult struct {
	// ID that is unique to the insight result being generated for the call.
	InsightID string `json:"insight_id"`
	// The result of the insight.
	Result CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsightID   respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationInsightsGeneratedWebhookEventDataPayloadResult) RawJSON() string {
	return r.JSON.raw
}
func (r *CallConversationInsightsGeneratedWebhookEventDataPayloadResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion
// contains all possible properties and values from [map[string]any], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCallConversationInsightsGeneratedWebhookEventDataPayloadResultResultInsightObjectResultItem
// OfString]
type CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfCallConversationInsightsGeneratedWebhookEventDataPayloadResultResultInsightObjectResultItem any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfCallConversationInsightsGeneratedWebhookEventDataPayloadResultResultInsightObjectResultItem respjson.Field
		OfString                                                                                      respjson.Field
		raw                                                                                           string
	} `json:"-"`
}

func (u CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion) AsInsightObjectResult() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CallConversationInsightsGeneratedWebhookEventDataPayloadResultResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallDtmfReceivedWebhookEvent struct {
	Data CallDtmfReceivedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallDtmfReceivedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallDtmfReceivedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallDtmfReceivedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.dtmf.received".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                               `json:"occurred_at" format:"date-time"`
	Payload    CallDtmfReceivedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallDtmfReceivedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallDtmfReceivedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallDtmfReceivedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Identifies the type of resource.
	ConnectionID string `json:"connection_id"`
	// The received DTMF digit or symbol.
	Digit string `json:"digit"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		Digit         respjson.Field
		From          respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallDtmfReceivedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallDtmfReceivedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallEnqueuedWebhookEvent struct {
	Data CallEnqueuedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallEnqueuedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallEnqueuedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallEnqueuedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.enqueued".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                           `json:"occurred_at" format:"date-time"`
	Payload    CallEnqueuedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallEnqueuedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallEnqueuedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallEnqueuedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Current position of the call in the queue.
	CurrentPosition int64 `json:"current_position"`
	// The name of the queue
	Queue string `json:"queue"`
	// Average time call spends in the queue in seconds.
	QueueAvgWaitTimeSecs int64 `json:"queue_avg_wait_time_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID        respjson.Field
		CallLegID            respjson.Field
		CallSessionID        respjson.Field
		ClientState          respjson.Field
		ConnectionID         respjson.Field
		CurrentPosition      respjson.Field
		Queue                respjson.Field
		QueueAvgWaitTimeSecs respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallEnqueuedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallEnqueuedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallForkStartedWebhookEvent struct {
	Data CallForkStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallForkStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallForkStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallForkStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.fork.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                              `json:"occurred_at" format:"date-time"`
	Payload    CallForkStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallForkStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallForkStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallForkStartedWebhookEventDataPayload struct {
	// Unique ID for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Type of media streamed. It can be either 'raw' or 'decrypted'.
	//
	// Any of "decrypted".
	StreamType string `json:"stream_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		StreamType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallForkStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallForkStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallForkStoppedWebhookEvent struct {
	Data CallForkStoppedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallForkStoppedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallForkStoppedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallForkStoppedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.fork.stopped".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                              `json:"occurred_at" format:"date-time"`
	Payload    CallForkStoppedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallForkStoppedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallForkStoppedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallForkStoppedWebhookEventDataPayload struct {
	// Unique ID for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Type of media streamed. It can be either 'raw' or 'decrypted'.
	//
	// Any of "decrypted".
	StreamType string `json:"stream_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		StreamType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallForkStoppedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallForkStoppedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGatherEndedWebhookEvent struct {
	Data CallGatherEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGatherEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallGatherEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGatherEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.gather.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                              `json:"occurred_at" format:"date-time"`
	Payload    CallGatherEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGatherEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallGatherEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGatherEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// The received DTMF digit or symbol.
	Digits string `json:"digits"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Reflects how command ended.
	//
	// Any of "valid", "invalid", "call_hangup", "cancelled", "cancelled_amd",
	// "timeout".
	Status string `json:"status"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		Digits        respjson.Field
		From          respjson.Field
		Status        respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGatherEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallGatherEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallHangupWebhookEvent struct {
	Data CallHangupWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallHangupWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallHangupWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.hangup".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                         `json:"occurred_at" format:"date-time"`
	Payload    CallHangupWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallHangupWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallHangupWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// Call quality statistics aggregated from the CHANNEL_HANGUP_COMPLETE event. Only
	// includes metrics that are available (filters out nil values). Returns nil if no
	// metrics are available.
	CallQualityStats CallHangupWebhookEventDataPayloadCallQualityStats `json:"call_quality_stats,nullable"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Custom headers set on answer command
	CustomHeaders []CustomSipHeader `json:"custom_headers"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// The reason the call was ended (`call_rejected`, `normal_clearing`,
	// `originator_cancel`, `timeout`, `time_limit`, `user_busy`, `not_found`,
	// `no_answer` or `unspecified`).
	//
	// Any of "call_rejected", "normal_clearing", "originator_cancel", "timeout",
	// "time_limit", "user_busy", "not_found", "no_answer", "unspecified".
	HangupCause string `json:"hangup_cause"`
	// The party who ended the call (`callee`, `caller`, `unknown`).
	//
	// Any of "caller", "callee", "unknown".
	HangupSource string `json:"hangup_source"`
	// The reason the call was ended (SIP response code). If the SIP response is
	// unavailable (in inbound calls for example) this is set to `unspecified`.
	SipHangupCause string `json:"sip_hangup_cause"`
	// User-to-User and Diversion headers from sip invite.
	SipHeaders []SipHeader `json:"sip_headers"`
	// ISO 8601 datetime of when the call started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// State received from a command.
	//
	// Any of "hangup".
	State string `json:"state"`
	// Array of tags associated to number.
	Tags []string `json:"tags"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID    respjson.Field
		CallLegID        respjson.Field
		CallQualityStats respjson.Field
		CallSessionID    respjson.Field
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		CustomHeaders    respjson.Field
		From             respjson.Field
		HangupCause      respjson.Field
		HangupSource     respjson.Field
		SipHangupCause   respjson.Field
		SipHeaders       respjson.Field
		StartTime        respjson.Field
		State            respjson.Field
		Tags             respjson.Field
		To               respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallHangupWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Call quality statistics aggregated from the CHANNEL_HANGUP_COMPLETE event. Only
// includes metrics that are available (filters out nil values). Returns nil if no
// metrics are available.
type CallHangupWebhookEventDataPayloadCallQualityStats struct {
	// Inbound call quality statistics.
	Inbound CallHangupWebhookEventDataPayloadCallQualityStatsInbound `json:"inbound"`
	// Outbound call quality statistics.
	Outbound CallHangupWebhookEventDataPayloadCallQualityStatsOutbound `json:"outbound"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inbound     respjson.Field
		Outbound    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupWebhookEventDataPayloadCallQualityStats) RawJSON() string { return r.JSON.raw }
func (r *CallHangupWebhookEventDataPayloadCallQualityStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inbound call quality statistics.
type CallHangupWebhookEventDataPayloadCallQualityStatsInbound struct {
	// Maximum jitter variance for inbound audio.
	JitterMaxVariance string `json:"jitter_max_variance"`
	// Number of packets used for jitter calculation on inbound audio.
	JitterPacketCount string `json:"jitter_packet_count"`
	// Mean Opinion Score (MOS) for inbound audio quality.
	Mos string `json:"mos"`
	// Total number of inbound audio packets.
	PacketCount string `json:"packet_count"`
	// Number of skipped inbound packets (packet loss).
	SkipPacketCount string `json:"skip_packet_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JitterMaxVariance respjson.Field
		JitterPacketCount respjson.Field
		Mos               respjson.Field
		PacketCount       respjson.Field
		SkipPacketCount   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupWebhookEventDataPayloadCallQualityStatsInbound) RawJSON() string { return r.JSON.raw }
func (r *CallHangupWebhookEventDataPayloadCallQualityStatsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound call quality statistics.
type CallHangupWebhookEventDataPayloadCallQualityStatsOutbound struct {
	// Total number of outbound audio packets.
	PacketCount string `json:"packet_count"`
	// Number of skipped outbound packets (packet loss).
	SkipPacketCount string `json:"skip_packet_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PacketCount     respjson.Field
		SkipPacketCount respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupWebhookEventDataPayloadCallQualityStatsOutbound) RawJSON() string {
	return r.JSON.raw
}
func (r *CallHangupWebhookEventDataPayloadCallQualityStatsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallInitiatedWebhookEvent struct {
	Data CallInitiatedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInitiatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallInitiatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallInitiatedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.initiated".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                            `json:"occurred_at" format:"date-time"`
	Payload    CallInitiatedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInitiatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallInitiatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallInitiatedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// Call screening result.
	CallScreeningResult string `json:"call_screening_result"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// Caller id.
	CallerIDName string `json:"caller_id_name"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// The list of comma-separated codecs enabled for the connection.
	ConnectionCodecs string `json:"connection_codecs"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Custom headers from sip invite
	CustomHeaders []CustomSipHeader `json:"custom_headers"`
	// Whether the call is `incoming` or `outgoing`.
	//
	// Any of "incoming", "outgoing".
	Direction string `json:"direction"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// The list of comma-separated codecs offered by caller.
	OfferedCodecs string `json:"offered_codecs"`
	// SHAKEN/STIR attestation level.
	ShakenStirAttestation string `json:"shaken_stir_attestation"`
	// Whether attestation was successfully validated or not.
	ShakenStirValidated bool `json:"shaken_stir_validated"`
	// User-to-User and Diversion headers from sip invite.
	SipHeaders []SipHeader `json:"sip_headers"`
	// ISO 8601 datetime of when the call started.
	StartTime time.Time `json:"start_time" format:"date-time"`
	// State received from a command.
	//
	// Any of "parked", "bridging".
	State string `json:"state"`
	// Array of tags associated to number.
	Tags []string `json:"tags"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID         respjson.Field
		CallLegID             respjson.Field
		CallScreeningResult   respjson.Field
		CallSessionID         respjson.Field
		CallerIDName          respjson.Field
		ClientState           respjson.Field
		ConnectionCodecs      respjson.Field
		ConnectionID          respjson.Field
		CustomHeaders         respjson.Field
		Direction             respjson.Field
		From                  respjson.Field
		OfferedCodecs         respjson.Field
		ShakenStirAttestation respjson.Field
		ShakenStirValidated   respjson.Field
		SipHeaders            respjson.Field
		StartTime             respjson.Field
		State                 respjson.Field
		Tags                  respjson.Field
		To                    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInitiatedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallInitiatedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallLeftQueueWebhookEvent struct {
	Data CallLeftQueueWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallLeftQueueWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallLeftQueueWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallLeftQueueWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.dequeued".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                            `json:"occurred_at" format:"date-time"`
	Payload    CallLeftQueueWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallLeftQueueWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallLeftQueueWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallLeftQueueWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// The name of the queue
	Queue string `json:"queue"`
	// Last position of the call in the queue.
	QueuePosition int64 `json:"queue_position"`
	// The reason for leaving the queue
	//
	// Any of "bridged", "bridging-in-process", "hangup", "leave", "timeout".
	Reason string `json:"reason"`
	// Time call spent in the queue in seconds.
	WaitTimeSecs int64 `json:"wait_time_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		Queue         respjson.Field
		QueuePosition respjson.Field
		Reason        respjson.Field
		WaitTimeSecs  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallLeftQueueWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallLeftQueueWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachineDetectionEndedWebhookEvent struct {
	Data CallMachineDetectionEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachineDetectionEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallMachineDetectionEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachineDetectionEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.detection.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                        `json:"occurred_at" format:"date-time"`
	Payload    CallMachineDetectionEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachineDetectionEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallMachineDetectionEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachineDetectionEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Answering machine detection result.
	//
	// Any of "human", "machine", "not_sure".
	Result string `json:"result"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		From          respjson.Field
		Result        respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachineDetectionEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachineDetectionEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachineGreetingEndedWebhookEvent struct {
	Data CallMachineGreetingEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachineGreetingEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallMachineGreetingEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachineGreetingEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.greeting.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                       `json:"occurred_at" format:"date-time"`
	Payload    CallMachineGreetingEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachineGreetingEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallMachineGreetingEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachineGreetingEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Answering machine greeting ended result.
	//
	// Any of "beep_detected", "ended", "not_sure".
	Result string `json:"result"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		From          respjson.Field
		Result        respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachineGreetingEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachineGreetingEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachinePremiumDetectionEndedWebhookEvent struct {
	Data CallMachinePremiumDetectionEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachinePremiumDetectionEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumDetectionEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachinePremiumDetectionEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.premium.detection.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                               `json:"occurred_at" format:"date-time"`
	Payload    CallMachinePremiumDetectionEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachinePremiumDetectionEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumDetectionEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachinePremiumDetectionEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Premium Answering Machine Detection result.
	//
	// Any of "human_residence", "human_business", "machine", "silence",
	// "fax_detected", "not_sure".
	Result string `json:"result"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		From          respjson.Field
		Result        respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachinePremiumDetectionEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumDetectionEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachinePremiumGreetingEndedWebhookEvent struct {
	Data CallMachinePremiumGreetingEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachinePremiumGreetingEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumGreetingEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachinePremiumGreetingEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.premium.greeting.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                              `json:"occurred_at" format:"date-time"`
	Payload    CallMachinePremiumGreetingEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachinePremiumGreetingEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumGreetingEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallMachinePremiumGreetingEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// Premium Answering Machine Greeting Ended result.
	//
	// Any of "beep_detected", "no_beep_detected".
	Result string `json:"result"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		From          respjson.Field
		Result        respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallMachinePremiumGreetingEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumGreetingEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallPlaybackEndedWebhookEvent struct {
	Data CallPlaybackEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallPlaybackEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallPlaybackEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.playback.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    CallPlaybackEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallPlaybackEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallPlaybackEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// The name of the audio media file being played back, if media_name has been used
	// to start.
	MediaName string `json:"media_name"`
	// The audio URL being played back, if audio_url has been used to start.
	MediaURL string `json:"media_url"`
	// Whether the stopped audio was in overlay mode or not.
	Overlay bool `json:"overlay"`
	// Reflects how command ended.
	//
	// Any of "file_not_found", "call_hangup", "unknown", "cancelled", "cancelled_amd",
	// "completed", "failed".
	Status string `json:"status"`
	// Provides details in case of failure.
	StatusDetail string `json:"status_detail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		MediaName     respjson.Field
		MediaURL      respjson.Field
		Overlay       respjson.Field
		Status        respjson.Field
		StatusDetail  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallPlaybackEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallPlaybackStartedWebhookEvent struct {
	Data CallPlaybackStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallPlaybackStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallPlaybackStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.playback.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                  `json:"occurred_at" format:"date-time"`
	Payload    CallPlaybackStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallPlaybackStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallPlaybackStartedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// The name of the audio media file being played back, if media_name has been used
	// to start.
	MediaName string `json:"media_name"`
	// The audio URL being played back, if audio_url has been used to start.
	MediaURL string `json:"media_url"`
	// Whether the audio is going to be played in overlay mode or not.
	Overlay bool `json:"overlay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		MediaName     respjson.Field
		MediaURL      respjson.Field
		Overlay       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallPlaybackStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingErrorWebhookEvent struct {
	Data CallRecordingErrorWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingErrorWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingErrorWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingErrorWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.recording.error".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                 `json:"occurred_at" format:"date-time"`
	Payload    CallRecordingErrorWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingErrorWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingErrorWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingErrorWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Indication that there was a problem recording the call.
	//
	// Any of "Failed to authorize with storage using custom credentials", "Invalid
	// credentials json", "Unsupported backend", "Internal server error".
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingErrorWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingErrorWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingSavedWebhookEvent struct {
	Data CallRecordingSavedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingSavedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.recording.saved".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                 `json:"occurred_at" format:"date-time"`
	Payload    CallRecordingSavedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingSavedWebhookEventDataPayload struct {
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// Whether recording was recorded in `single` or `dual` channel.
	//
	// Any of "single", "dual".
	Channels string `json:"channels"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Recording URLs in requested format. The URL is valid for as long as the file
	// exists. For security purposes, this feature is activated on a per request basis.
	// Please contact customer support with your Account ID to request activation.
	PublicRecordingURLs CallRecordingSavedWebhookEventDataPayloadPublicRecordingURLs `json:"public_recording_urls"`
	// ISO 8601 datetime of when recording ended.
	RecordingEndedAt time.Time `json:"recording_ended_at" format:"date-time"`
	// ISO 8601 datetime of when recording started.
	RecordingStartedAt time.Time `json:"recording_started_at" format:"date-time"`
	// Recording URLs in requested format. These URLs are valid for 10 minutes. After
	// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
	// documentation, or via Mission Control under Reporting -> Recordings.
	RecordingURLs CallRecordingSavedWebhookEventDataPayloadRecordingURLs `json:"recording_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallLegID           respjson.Field
		CallSessionID       respjson.Field
		Channels            respjson.Field
		ClientState         respjson.Field
		ConnectionID        respjson.Field
		PublicRecordingURLs respjson.Field
		RecordingEndedAt    respjson.Field
		RecordingStartedAt  respjson.Field
		RecordingURLs       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. The URL is valid for as long as the file
// exists. For security purposes, this feature is activated on a per request basis.
// Please contact customer support with your Account ID to request activation.
type CallRecordingSavedWebhookEventDataPayloadPublicRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3,nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedWebhookEventDataPayloadPublicRecordingURLs) RawJSON() string {
	return r.JSON.raw
}
func (r *CallRecordingSavedWebhookEventDataPayloadPublicRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. These URLs are valid for 10 minutes. After
// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
// documentation, or via Mission Control under Reporting -> Recordings.
type CallRecordingSavedWebhookEventDataPayloadRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3,nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedWebhookEventDataPayloadRecordingURLs) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedWebhookEventDataPayloadRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingTranscriptionSavedWebhookEvent struct {
	Data CallRecordingTranscriptionSavedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingTranscriptionSavedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingTranscriptionSavedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingTranscriptionSavedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.recording.transcription.saved".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                              `json:"occurred_at" format:"date-time"`
	Payload    CallRecordingTranscriptionSavedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingTranscriptionSavedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingTranscriptionSavedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallRecordingTranscriptionSavedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// The type of calling party connection.
	//
	// Any of "pstn", "sip".
	CallingPartyType string `json:"calling_party_type"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the recording session and can be used to correlate webhook
	// events.
	RecordingID string `json:"recording_id"`
	// ID that is unique to the transcription process and can be used to correlate
	// webhook events.
	RecordingTranscriptionID string `json:"recording_transcription_id"`
	// The transcription status.
	//
	// Any of "completed".
	Status string `json:"status"`
	// The transcribed text
	TranscriptionText string `json:"transcription_text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID            respjson.Field
		CallLegID                respjson.Field
		CallSessionID            respjson.Field
		CallingPartyType         respjson.Field
		ClientState              respjson.Field
		ConnectionID             respjson.Field
		RecordingID              respjson.Field
		RecordingTranscriptionID respjson.Field
		Status                   respjson.Field
		TranscriptionText        respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingTranscriptionSavedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingTranscriptionSavedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferCompletedWebhookEvent struct {
	Data CallReferCompletedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferCompletedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallReferCompletedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferCompletedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.refer.completed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                 `json:"occurred_at" format:"date-time"`
	Payload    CallReferCompletedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferCompletedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallReferCompletedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferCompletedWebhookEventDataPayload struct {
	// Unique ID for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// SIP NOTIFY event status for tracking the REFER attempt.
	SipNotifyResponse int64 `json:"sip_notify_response"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID     respjson.Field
		CallLegID         respjson.Field
		CallSessionID     respjson.Field
		ClientState       respjson.Field
		ConnectionID      respjson.Field
		From              respjson.Field
		SipNotifyResponse respjson.Field
		To                respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferCompletedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallReferCompletedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferFailedWebhookEvent struct {
	Data CallReferFailedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallReferFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferFailedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.refer.failed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                              `json:"occurred_at" format:"date-time"`
	Payload    CallReferFailedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallReferFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferFailedWebhookEventDataPayload struct {
	// Unique ID for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// SIP NOTIFY event status for tracking the REFER attempt.
	SipNotifyResponse int64 `json:"sip_notify_response"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID     respjson.Field
		CallLegID         respjson.Field
		CallSessionID     respjson.Field
		ClientState       respjson.Field
		ConnectionID      respjson.Field
		From              respjson.Field
		SipNotifyResponse respjson.Field
		To                respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferFailedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallReferFailedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferStartedWebhookEvent struct {
	Data CallReferStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallReferStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.refer.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                               `json:"occurred_at" format:"date-time"`
	Payload    CallReferStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallReferStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallReferStartedWebhookEventDataPayload struct {
	// Unique ID for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Number or SIP URI placing the call.
	From string `json:"from"`
	// SIP NOTIFY event status for tracking the REFER attempt.
	SipNotifyResponse int64 `json:"sip_notify_response"`
	// Destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID     respjson.Field
		CallLegID         respjson.Field
		CallSessionID     respjson.Field
		ClientState       respjson.Field
		ConnectionID      respjson.Field
		From              respjson.Field
		SipNotifyResponse respjson.Field
		To                respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallReferStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallReferStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecFailedWebhookEvent struct {
	Data CallSiprecFailedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecFailedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "siprec.failed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                               `json:"occurred_at" format:"date-time"`
	Payload    CallSiprecFailedWebhookEventDataPayload `json:"payload"`
	// Identifies the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecFailedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Q850 reason why siprec session failed.
	FailureCause string `json:"failure_cause"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		FailureCause  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecFailedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecFailedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecStartedWebhookEvent struct {
	Data CallSiprecStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "siprec.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    CallSiprecStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecStartedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecStoppedWebhookEvent struct {
	Data CallSiprecStoppedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecStoppedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStoppedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecStoppedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "siprec.stopped".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    CallSiprecStoppedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecStoppedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStoppedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSiprecStoppedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Q850 reason why the SIPREC session was stopped.
	HangupCause string `json:"hangup_cause"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		HangupCause   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSiprecStoppedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStoppedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSpeakEndedWebhookEvent struct {
	Data CallSpeakEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSpeakEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSpeakEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.speak.ended".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                             `json:"occurred_at" format:"date-time"`
	Payload    CallSpeakEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSpeakEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSpeakEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// Reflects how the command ended.
	//
	// Any of "completed", "call_hangup", "cancelled_amd".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSpeakEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSpeakStartedWebhookEvent struct {
	Data CallSpeakStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSpeakStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSpeakStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.speak.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                               `json:"occurred_at" format:"date-time"`
	Payload    CallSpeakStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSpeakStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallSpeakStartedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConnectionID  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallSpeakStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallStreamingFailedWebhookEvent struct {
	Data CallStreamingFailed `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallStreamingStartedWebhookEvent struct {
	Data CallStreamingStarted `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallStreamingStoppedWebhookEvent struct {
	Data CallStreamingStopped `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallStreamingStoppedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CallStreamingStoppedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignStatusUpdateWebhookEvent struct {
	// Brand ID associated with the campaign.
	BrandID string `json:"brandId"`
	// The ID of the campaign.
	CampaignID string `json:"campaignId"`
	// Unix timestamp when campaign was created.
	CreateDate string `json:"createDate"`
	// Alphanumeric identifier of the CSP associated with this campaign.
	CspID string `json:"cspId"`
	// Description of the event.
	Description string `json:"description"`
	// Indicates whether the campaign is registered with T-Mobile.
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// The status of the campaign.
	//
	// Any of "ACCEPTED", "REJECTED", "DORMANT", "success", "failed".
	Status CampaignStatusUpdateWebhookEventStatus `json:"status"`
	// Any of "TELNYX_EVENT", "REGISTRATION", "MNO_REVIEW", "TELNYX_REVIEW",
	// "NUMBER_POOL_PROVISIONED", "NUMBER_POOL_DEPROVISIONED", "TCR_EVENT", "VERIFIED".
	Type CampaignStatusUpdateWebhookEventType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrandID             respjson.Field
		CampaignID          respjson.Field
		CreateDate          respjson.Field
		CspID               respjson.Field
		Description         respjson.Field
		IsTMobileRegistered respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CampaignStatusUpdateWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *CampaignStatusUpdateWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the campaign.
type CampaignStatusUpdateWebhookEventStatus string

const (
	CampaignStatusUpdateWebhookEventStatusAccepted CampaignStatusUpdateWebhookEventStatus = "ACCEPTED"
	CampaignStatusUpdateWebhookEventStatusRejected CampaignStatusUpdateWebhookEventStatus = "REJECTED"
	CampaignStatusUpdateWebhookEventStatusDormant  CampaignStatusUpdateWebhookEventStatus = "DORMANT"
	CampaignStatusUpdateWebhookEventStatusSuccess  CampaignStatusUpdateWebhookEventStatus = "success"
	CampaignStatusUpdateWebhookEventStatusFailed   CampaignStatusUpdateWebhookEventStatus = "failed"
)

type CampaignStatusUpdateWebhookEventType string

const (
	CampaignStatusUpdateWebhookEventTypeTelnyxEvent             CampaignStatusUpdateWebhookEventType = "TELNYX_EVENT"
	CampaignStatusUpdateWebhookEventTypeRegistration            CampaignStatusUpdateWebhookEventType = "REGISTRATION"
	CampaignStatusUpdateWebhookEventTypeMnoReview               CampaignStatusUpdateWebhookEventType = "MNO_REVIEW"
	CampaignStatusUpdateWebhookEventTypeTelnyxReview            CampaignStatusUpdateWebhookEventType = "TELNYX_REVIEW"
	CampaignStatusUpdateWebhookEventTypeNumberPoolProvisioned   CampaignStatusUpdateWebhookEventType = "NUMBER_POOL_PROVISIONED"
	CampaignStatusUpdateWebhookEventTypeNumberPoolDeprovisioned CampaignStatusUpdateWebhookEventType = "NUMBER_POOL_DEPROVISIONED"
	CampaignStatusUpdateWebhookEventTypeTcrEvent                CampaignStatusUpdateWebhookEventType = "TCR_EVENT"
	CampaignStatusUpdateWebhookEventTypeVerified                CampaignStatusUpdateWebhookEventType = "VERIFIED"
)

type ConferenceCreatedWebhookEvent struct {
	Data ConferenceCreatedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceCreatedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceCreatedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceCreatedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.created".
	EventType string                                   `json:"event_type"`
	Payload   ConferenceCreatedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceCreatedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceCreatedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceCreatedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Conference ID that the participant joined.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConferenceID  respjson.Field
		ConnectionID  respjson.Field
		OccurredAt    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceCreatedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceCreatedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceEndedWebhookEvent struct {
	Data ConferenceEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.ended".
	EventType string                                 `json:"event_type"`
	Payload   ConferenceEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceEndedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Conference ID that the participant joined.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// Reason the conference ended.
	//
	// Any of "all_left", "host_left", "time_exceeded".
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConferenceID  respjson.Field
		ConnectionID  respjson.Field
		OccurredAt    respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceFloorChangedWebhookEvent struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.floor.changed".
	EventType ConferenceFloorChangedWebhookEventEventType `json:"event_type"`
	Payload   ConferenceFloorChangedWebhookEventPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceFloorChangedWebhookEventRecordType `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceFloorChangedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceFloorChangedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceFloorChangedWebhookEventEventType string

const (
	ConferenceFloorChangedWebhookEventEventTypeConferenceFloorChanged ConferenceFloorChangedWebhookEventEventType = "conference.floor.changed"
)

type ConferenceFloorChangedWebhookEventPayload struct {
	// Call Control ID of the new speaker.
	CallControlID string `json:"call_control_id"`
	// Call Leg ID of the new speaker.
	CallLegID string `json:"call_leg_id"`
	// Call Session ID of the new speaker.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Conference ID that had a speaker change event.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConferenceID  respjson.Field
		ConnectionID  respjson.Field
		OccurredAt    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceFloorChangedWebhookEventPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceFloorChangedWebhookEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceFloorChangedWebhookEventRecordType string

const (
	ConferenceFloorChangedWebhookEventRecordTypeEvent ConferenceFloorChangedWebhookEventRecordType = "event"
)

type ConferenceParticipantJoinedWebhookEvent struct {
	Data ConferenceParticipantJoinedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantJoinedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantJoinedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantJoinedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.joined".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                          `json:"occurred_at" format:"date-time"`
	Payload    ConferenceParticipantJoinedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantJoinedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantJoinedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantJoinedWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Conference ID that the participant joined.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConferenceID  respjson.Field
		ConnectionID  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantJoinedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantJoinedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantLeftWebhookEvent struct {
	Data ConferenceParticipantLeftWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantLeftWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantLeftWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantLeftWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.left".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                        `json:"occurred_at" format:"date-time"`
	Payload    ConferenceParticipantLeftWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantLeftWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantLeftWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantLeftWebhookEventDataPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// Conference ID that the participant joined.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID respjson.Field
		CallLegID     respjson.Field
		CallSessionID respjson.Field
		ClientState   respjson.Field
		ConferenceID  respjson.Field
		ConnectionID  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantLeftWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantLeftWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantPlaybackEndedWebhookEvent struct {
	Data ConferenceParticipantPlaybackEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantPlaybackEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantPlaybackEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.playback.ended".
	EventType string                                                    `json:"event_type"`
	Payload   ConferenceParticipantPlaybackEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantPlaybackEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantPlaybackEndedWebhookEventDataPayload struct {
	// Participant's call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// The name of the audio media file being played back, if media_name has been used
	// to start.
	MediaName string `json:"media_name"`
	// The audio URL being played back, if audio_url has been used to start.
	MediaURL string `json:"media_url"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID        respjson.Field
		CallLegID            respjson.Field
		CallSessionID        respjson.Field
		ClientState          respjson.Field
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		MediaName            respjson.Field
		MediaURL             respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantPlaybackEndedWebhookEventDataPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *ConferenceParticipantPlaybackEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantPlaybackStartedWebhookEvent struct {
	Data ConferenceParticipantPlaybackStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantPlaybackStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantPlaybackStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.playback.started".
	EventType string                                                      `json:"event_type"`
	Payload   ConferenceParticipantPlaybackStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantPlaybackStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantPlaybackStartedWebhookEventDataPayload struct {
	// Participant's call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// The name of the audio media file being played back, if media_name has been used
	// to start.
	MediaName string `json:"media_name"`
	// The audio URL being played back, if audio_url has been used to start.
	MediaURL string `json:"media_url"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID        respjson.Field
		CallLegID            respjson.Field
		CallSessionID        respjson.Field
		ClientState          respjson.Field
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		MediaName            respjson.Field
		MediaURL             respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantPlaybackStartedWebhookEventDataPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *ConferenceParticipantPlaybackStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantSpeakEndedWebhookEvent struct {
	Data ConferenceParticipantSpeakEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantSpeakEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantSpeakEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.speak.ended".
	EventType string                                                 `json:"event_type"`
	Payload   ConferenceParticipantSpeakEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantSpeakEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantSpeakEndedWebhookEventDataPayload struct {
	// Participant's call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID        respjson.Field
		CallLegID            respjson.Field
		CallSessionID        respjson.Field
		ClientState          respjson.Field
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantSpeakEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantSpeakStartedWebhookEvent struct {
	Data ConferenceParticipantSpeakStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantSpeakStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantSpeakStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.speak.started".
	EventType string                                                   `json:"event_type"`
	Payload   ConferenceParticipantSpeakStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantSpeakStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceParticipantSpeakStartedWebhookEventDataPayload struct {
	// Participant's call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID        respjson.Field
		CallLegID            respjson.Field
		CallSessionID        respjson.Field
		ClientState          respjson.Field
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceParticipantSpeakStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferencePlaybackEndedWebhookEvent struct {
	Data ConferencePlaybackEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferencePlaybackEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferencePlaybackEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.playback.ended".
	EventType string                                         `json:"event_type"`
	Payload   ConferencePlaybackEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferencePlaybackEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferencePlaybackEndedWebhookEventDataPayload struct {
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// The name of the audio media file being played back, if media_name has been used
	// to start.
	MediaName string `json:"media_name"`
	// The audio URL being played back, if audio_url has been used to start.
	MediaURL string `json:"media_url"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		MediaName            respjson.Field
		MediaURL             respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferencePlaybackEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferencePlaybackStartedWebhookEvent struct {
	Data ConferencePlaybackStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferencePlaybackStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferencePlaybackStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.playback.started".
	EventType string                                           `json:"event_type"`
	Payload   ConferencePlaybackStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferencePlaybackStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferencePlaybackStartedWebhookEventDataPayload struct {
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// The name of the audio media file being played back, if media_name has been used
	// to start.
	MediaName string `json:"media_name"`
	// The audio URL being played back, if audio_url has been used to start.
	MediaURL string `json:"media_url"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		MediaName            respjson.Field
		MediaURL             respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferencePlaybackStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceRecordingSavedWebhookEvent struct {
	Data ConferenceRecordingSavedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSavedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceRecordingSavedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.recording.saved".
	EventType string                                          `json:"event_type"`
	Payload   ConferenceRecordingSavedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSavedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceRecordingSavedWebhookEventDataPayload struct {
	// Participant's call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// Whether recording was recorded in `single` or `dual` channel.
	//
	// Any of "single", "dual".
	Channels string `json:"channels"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// ID of the conference that is being recorded.
	ConferenceID string `json:"conference_id" format:"uuid"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// The audio file format used when storing the call recording. Can be either `mp3`
	// or `wav`.
	//
	// Any of "wav", "mp3".
	Format string `json:"format"`
	// Recording URLs in requested format. The URL is valid for as long as the file
	// exists. For security purposes, this feature is activated on a per request basis.
	// Please contact customer support with your Account ID to request activation.
	PublicRecordingURLs ConferenceRecordingSavedWebhookEventDataPayloadPublicRecordingURLs `json:"public_recording_urls"`
	// ISO 8601 datetime of when recording ended.
	RecordingEndedAt time.Time `json:"recording_ended_at" format:"date-time"`
	// ID of the conference recording.
	RecordingID string `json:"recording_id" format:"uuid"`
	// ISO 8601 datetime of when recording started.
	RecordingStartedAt time.Time `json:"recording_started_at" format:"date-time"`
	// Recording URLs in requested format. These URLs are valid for 10 minutes. After
	// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
	// documentation, or via Mission Control under Reporting -> Recordings.
	RecordingURLs ConferenceRecordingSavedWebhookEventDataPayloadRecordingURLs `json:"recording_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID       respjson.Field
		CallSessionID       respjson.Field
		Channels            respjson.Field
		ClientState         respjson.Field
		ConferenceID        respjson.Field
		ConnectionID        respjson.Field
		Format              respjson.Field
		PublicRecordingURLs respjson.Field
		RecordingEndedAt    respjson.Field
		RecordingID         respjson.Field
		RecordingStartedAt  respjson.Field
		RecordingURLs       respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSavedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. The URL is valid for as long as the file
// exists. For security purposes, this feature is activated on a per request basis.
// Please contact customer support with your Account ID to request activation.
type ConferenceRecordingSavedWebhookEventDataPayloadPublicRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3,nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedWebhookEventDataPayloadPublicRecordingURLs) RawJSON() string {
	return r.JSON.raw
}
func (r *ConferenceRecordingSavedWebhookEventDataPayloadPublicRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. These URLs are valid for 10 minutes. After
// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
// documentation, or via Mission Control under Reporting -> Recordings.
type ConferenceRecordingSavedWebhookEventDataPayloadRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3,nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedWebhookEventDataPayloadRecordingURLs) RawJSON() string {
	return r.JSON.raw
}
func (r *ConferenceRecordingSavedWebhookEventDataPayloadRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceSpeakEndedWebhookEvent struct {
	Data ConferenceSpeakEndedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceSpeakEndedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakEndedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceSpeakEndedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.speak.ended".
	EventType string                                      `json:"event_type"`
	Payload   ConferenceSpeakEndedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceSpeakEndedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakEndedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceSpeakEndedWebhookEventDataPayload struct {
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceSpeakEndedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakEndedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceSpeakStartedWebhookEvent struct {
	Data ConferenceSpeakStartedWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceSpeakStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceSpeakStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.speak.started".
	EventType string                                        `json:"event_type"`
	Payload   ConferenceSpeakStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceSpeakStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConferenceSpeakStartedWebhookEventDataPayload struct {
	// ID of the conference the text was spoken in.
	ConferenceID string `json:"conference_id"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID string `json:"connection_id"`
	// ID that is unique to the call session that started the conference.
	CreatorCallSessionID string `json:"creator_call_session_id"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time `json:"occurred_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConferenceID         respjson.Field
		ConnectionID         respjson.Field
		CreatorCallSessionID respjson.Field
		OccurredAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceSpeakStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliveryUpdateWebhookEvent struct {
	Data DeliveryUpdateWebhookEventData `json:"data"`
	Meta DeliveryUpdateWebhookEventMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryUpdateWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *DeliveryUpdateWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliveryUpdateWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "message.sent", "message.finalized".
	EventType string `json:"event_type"`
	// ISO 8601 formatted date indicating when the resource was created.
	OccurredAt time.Time              `json:"occurred_at" format:"date-time"`
	Payload    OutboundMessagePayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryUpdateWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *DeliveryUpdateWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeliveryUpdateWebhookEventMeta struct {
	// Number of attempts to deliver the webhook event.
	Attempt int64 `json:"attempt"`
	// The webhook URL the event was delivered to.
	DeliveredTo string `json:"delivered_to" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryUpdateWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *DeliveryUpdateWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxDeliveredWebhookEvent struct {
	Data FaxDeliveredWebhookEventData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxDeliveredWebhookEventMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxDeliveredWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxDeliveredWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.delivered".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                           `json:"occurred_at" format:"date-time"`
	Payload    FaxDeliveredWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxDeliveredWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxDeliveredWebhookEventDataPayload struct {
	// The duration of the call in seconds.
	CallDurationSecs int64 `json:"call_duration_secs"`
	// State received from a command.
	ClientState string `json:"client_state"`
	// The ID of the connection used to send the fax.
	ConnectionID string `json:"connection_id"`
	// The direction of the fax.
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// Identifies the fax.
	FaxID string `json:"fax_id" format:"uuid"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName string `json:"media_name"`
	// The original URL to the PDF used for the fax's media. If media_name was
	// supplied, this is omitted
	OriginalMediaURL string `json:"original_media_url"`
	// Number of transferred pages
	PageCount int64 `json:"page_count"`
	// The status of the fax.
	//
	// Any of "delivered".
	Status string `json:"status"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// Identifier of the user to whom the fax belongs
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallDurationSecs respjson.Field
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		Direction        respjson.Field
		FaxID            respjson.Field
		From             respjson.Field
		MediaName        respjson.Field
		OriginalMediaURL respjson.Field
		PageCount        respjson.Field
		Status           respjson.Field
		To               respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxDeliveredWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxDeliveredWebhookEventMeta struct {
	// The delivery attempt number.
	Attempt int64 `json:"attempt"`
	// The URL the webhook was delivered to.
	DeliveredTo string `json:"delivered_to" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxDeliveredWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxFailedWebhookEvent struct {
	Data FaxFailedWebhookEventData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxFailedWebhookEventMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxFailedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxFailedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.failed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                        `json:"occurred_at" format:"date-time"`
	Payload    FaxFailedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxFailedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxFailedWebhookEventDataPayload struct {
	// State received from a command.
	ClientState string `json:"client_state"`
	// The ID of the connection used to send the fax.
	ConnectionID string `json:"connection_id"`
	// The direction of the fax.
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// Cause of the sending failure
	//
	// Any of "rejected".
	FailureReason string `json:"failure_reason"`
	// Identifies the fax.
	FaxID string `json:"fax_id" format:"uuid"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName string `json:"media_name"`
	// The original URL to the PDF used for the fax's media. If media_name was
	// supplied, this is omitted
	OriginalMediaURL string `json:"original_media_url"`
	// The status of the fax.
	//
	// Any of "failed".
	Status string `json:"status"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// Identifier of the user to whom the fax belongs
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		Direction        respjson.Field
		FailureReason    respjson.Field
		FaxID            respjson.Field
		From             respjson.Field
		MediaName        respjson.Field
		OriginalMediaURL respjson.Field
		Status           respjson.Field
		To               respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxFailedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxFailedWebhookEventMeta struct {
	// The delivery attempt number.
	Attempt int64 `json:"attempt"`
	// The URL the webhook was delivered to.
	DeliveredTo string `json:"delivered_to" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxFailedWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxMediaProcessedWebhookEvent struct {
	Data FaxMediaProcessedWebhookEventData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxMediaProcessedWebhookEventMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxMediaProcessedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxMediaProcessedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.media.processed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    FaxMediaProcessedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxMediaProcessedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxMediaProcessedWebhookEventDataPayload struct {
	// State received from a command.
	ClientState string `json:"client_state"`
	// The ID of the connection used to send the fax.
	ConnectionID string `json:"connection_id"`
	// The direction of the fax.
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// Identifies the fax.
	FaxID string `json:"fax_id" format:"uuid"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName string `json:"media_name"`
	// The original URL to the PDF used for the fax's media. If media_name was
	// supplied, this is omitted
	OriginalMediaURL string `json:"original_media_url"`
	// The status of the fax.
	//
	// Any of "media.processed".
	Status string `json:"status"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// Identifier of the user to whom the fax belongs
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		Direction        respjson.Field
		FaxID            respjson.Field
		From             respjson.Field
		MediaName        respjson.Field
		OriginalMediaURL respjson.Field
		Status           respjson.Field
		To               respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxMediaProcessedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxMediaProcessedWebhookEventMeta struct {
	// The delivery attempt number.
	Attempt int64 `json:"attempt"`
	// The URL the webhook was delivered to.
	DeliveredTo string `json:"delivered_to" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxMediaProcessedWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxQueuedWebhookEvent struct {
	Data FaxQueuedWebhookEventData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxQueuedWebhookEventMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxQueuedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxQueuedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.queued".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                        `json:"occurred_at" format:"date-time"`
	Payload    FaxQueuedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxQueuedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxQueuedWebhookEventDataPayload struct {
	// State received from a command.
	ClientState string `json:"client_state"`
	// The ID of the connection used to send the fax.
	ConnectionID string `json:"connection_id"`
	// The direction of the fax.
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// Identifies the fax.
	FaxID string `json:"fax_id" format:"uuid"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName string `json:"media_name"`
	// The original URL to the PDF used for the fax's media. If media_name was
	// supplied, this is omitted
	OriginalMediaURL string `json:"original_media_url"`
	// The status of the fax.
	//
	// Any of "queued".
	Status string `json:"status"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// Identifier of the user to whom the fax belongs
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		Direction        respjson.Field
		FaxID            respjson.Field
		From             respjson.Field
		MediaName        respjson.Field
		OriginalMediaURL respjson.Field
		Status           respjson.Field
		To               respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxQueuedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxQueuedWebhookEventMeta struct {
	// The delivery attempt number.
	Attempt int64 `json:"attempt"`
	// The URL the webhook was delivered to.
	DeliveredTo string `json:"delivered_to" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxQueuedWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxSendingStartedWebhookEvent struct {
	Data FaxSendingStartedWebhookEventData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxSendingStartedWebhookEventMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxSendingStartedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxSendingStartedWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.sending.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    FaxSendingStartedWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxSendingStartedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxSendingStartedWebhookEventDataPayload struct {
	// State received from a command.
	ClientState string `json:"client_state"`
	// The ID of the connection used to send the fax.
	ConnectionID string `json:"connection_id"`
	// The direction of the fax.
	//
	// Any of "inbound", "outbound".
	Direction string `json:"direction"`
	// Identifies the fax.
	FaxID string `json:"fax_id" format:"uuid"`
	// The phone number, in E.164 format, the fax will be sent from.
	From string `json:"from"`
	// The media_name used for the fax's media. Must point to a file previously
	// uploaded to api.telnyx.com/v2/media by the same user/organization. media_name
	// and media_url/contents can't be submitted together.
	MediaName string `json:"media_name"`
	// The original URL to the PDF used for the fax's media. If media_name was
	// supplied, this is omitted
	OriginalMediaURL string `json:"original_media_url"`
	// The status of the fax.
	//
	// Any of "sending".
	Status string `json:"status"`
	// The phone number, in E.164 format, the fax will be sent to or SIP URI
	To string `json:"to"`
	// Identifier of the user to whom the fax belongs
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientState      respjson.Field
		ConnectionID     respjson.Field
		Direction        respjson.Field
		FaxID            respjson.Field
		From             respjson.Field
		MediaName        respjson.Field
		OriginalMediaURL respjson.Field
		Status           respjson.Field
		To               respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxSendingStartedWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxSendingStartedWebhookEventMeta struct {
	// The delivery attempt number.
	Attempt int64 `json:"attempt"`
	// The URL the webhook was delivered to.
	DeliveredTo string `json:"delivered_to" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxSendingStartedWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessageWebhookEvent struct {
	Data InboundMessageWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessageWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *InboundMessageWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessageWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "message.received".
	EventType string `json:"event_type"`
	// ISO 8601 formatted date indicating when the resource was created.
	OccurredAt time.Time                    `json:"occurred_at" format:"date-time"`
	Payload    shared.InboundMessagePayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboundMessageWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *InboundMessageWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderStatusUpdateWebhookEvent struct {
	Data NumberOrderStatusUpdateWebhookEventData `json:"data,required"`
	Meta NumberOrderStatusUpdateWebhookEventMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderStatusUpdateWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderStatusUpdateWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderStatusUpdateWebhookEventData struct {
	// Unique identifier for the event
	ID string `json:"id,required" format:"uuid"`
	// The type of event being sent
	EventType string `json:"event_type,required"`
	// ISO 8601 timestamp of when the event occurred
	OccurredAt time.Time                   `json:"occurred_at,required" format:"date-time"`
	Payload    NumberOrderWithPhoneNumbers `json:"payload,required"`
	// Type of record
	RecordType string `json:"record_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderStatusUpdateWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderStatusUpdateWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderStatusUpdateWebhookEventMeta struct {
	// Webhook delivery attempt number
	Attempt int64 `json:"attempt,required"`
	// URL where the webhook was delivered
	DeliveredTo string `json:"delivered_to,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderStatusUpdateWebhookEventMeta) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderStatusUpdateWebhookEventMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplacedLinkClickWebhookEvent struct {
	Data ReplacedLinkClickWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplacedLinkClickWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *ReplacedLinkClickWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplacedLinkClickWebhookEventData struct {
	// The message ID associated with the clicked link.
	MessageID string `json:"message_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the message request was received.
	TimeClicked time.Time `json:"time_clicked" format:"date-time"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	To string `json:"to" format:"address"`
	// The original link that was sent in the message.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		RecordType  respjson.Field
		TimeClicked respjson.Field
		To          respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplacedLinkClickWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *ReplacedLinkClickWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionWebhookEvent struct {
	Data TranscriptionWebhookEventData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionWebhookEventData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.transcription".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                            `json:"occurred_at" format:"date-time"`
	Payload    TranscriptionWebhookEventDataPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionWebhookEventDataPayload struct {
	// Unique identifier and token for controlling the call.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// ID that is unique to the call session and can be used to correlate webhook
	// events. Call session is a group of related call legs that logically belong to
	// the same phone call, e.g. an inbound and outbound leg of a transferred call.
	CallSessionID string `json:"call_session_id"`
	// Use this field to add state to every subsequent webhook. It must be a valid
	// Base-64 encoded string.
	ClientState string `json:"client_state"`
	// Call Control App ID (formerly Telnyx connection ID) used in the call.
	ConnectionID      string                                                `json:"connection_id"`
	TranscriptionData TranscriptionWebhookEventDataPayloadTranscriptionData `json:"transcription_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallControlID     respjson.Field
		CallLegID         respjson.Field
		CallSessionID     respjson.Field
		ClientState       respjson.Field
		ConnectionID      respjson.Field
		TranscriptionData respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionWebhookEventDataPayload) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionWebhookEventDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionWebhookEventDataPayloadTranscriptionData struct {
	// Speech recognition confidence level.
	Confidence float64 `json:"confidence"`
	// When false, it means that this is an interim result.
	IsFinal bool `json:"is_final"`
	// Recognized text.
	Transcript string `json:"transcript"`
	// Indicates which leg of the call has been transcribed. This is only available
	// when `transcription_engine` is set to `B`.
	//
	// Any of "inbound", "outbound".
	TranscriptionTrack string `json:"transcription_track"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence         respjson.Field
		IsFinal            respjson.Field
		Transcript         respjson.Field
		TranscriptionTrack respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionWebhookEventDataPayloadTranscriptionData) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionWebhookEventDataPayloadTranscriptionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnion contains all possible properties and values from
// [CallAIGatherEndedWebhookEvent],
// [CallAIGatherMessageHistoryUpdatedWebhookEvent],
// [CallAIGatherPartialResultsWebhookEvent], [CallAnsweredWebhookEvent],
// [CallBridgedWebhookEvent], [CallConversationEndedWebhookEvent],
// [CallConversationInsightsGeneratedWebhookEvent], [CallDtmfReceivedWebhookEvent],
// [CallEnqueuedWebhookEvent], [CallForkStartedWebhookEvent],
// [CallForkStoppedWebhookEvent], [CallGatherEndedWebhookEvent],
// [CallHangupWebhookEvent], [CallInitiatedWebhookEvent],
// [CallLeftQueueWebhookEvent], [CallMachineDetectionEndedWebhookEvent],
// [CallMachineGreetingEndedWebhookEvent],
// [CallMachinePremiumDetectionEndedWebhookEvent],
// [CallMachinePremiumGreetingEndedWebhookEvent], [CallPlaybackEndedWebhookEvent],
// [CallPlaybackStartedWebhookEvent], [CallRecordingErrorWebhookEvent],
// [CallRecordingSavedWebhookEvent], [CallRecordingTranscriptionSavedWebhookEvent],
// [CallReferCompletedWebhookEvent], [CallReferFailedWebhookEvent],
// [CallReferStartedWebhookEvent], [CallSiprecFailedWebhookEvent],
// [CallSiprecStartedWebhookEvent], [CallSiprecStoppedWebhookEvent],
// [CallSpeakEndedWebhookEvent], [CallSpeakStartedWebhookEvent],
// [CallStreamingFailedWebhookEvent], [CallStreamingStartedWebhookEvent],
// [CallStreamingStoppedWebhookEvent], [CampaignStatusUpdateWebhookEvent],
// [ConferenceCreatedWebhookEvent], [ConferenceEndedWebhookEvent],
// [ConferenceFloorChangedWebhookEvent], [ConferenceParticipantJoinedWebhookEvent],
// [ConferenceParticipantLeftWebhookEvent],
// [ConferenceParticipantPlaybackEndedWebhookEvent],
// [ConferenceParticipantPlaybackStartedWebhookEvent],
// [ConferenceParticipantSpeakEndedWebhookEvent],
// [ConferenceParticipantSpeakStartedWebhookEvent],
// [ConferencePlaybackEndedWebhookEvent], [ConferencePlaybackStartedWebhookEvent],
// [ConferenceRecordingSavedWebhookEvent], [ConferenceSpeakEndedWebhookEvent],
// [ConferenceSpeakStartedWebhookEvent], [DeliveryUpdateWebhookEvent],
// [FaxDeliveredWebhookEvent], [FaxFailedWebhookEvent],
// [FaxMediaProcessedWebhookEvent], [FaxQueuedWebhookEvent],
// [FaxSendingStartedWebhookEvent], [InboundMessageWebhookEvent],
// [NumberOrderStatusUpdateWebhookEvent], [ReplacedLinkClickWebhookEvent],
// [TranscriptionWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	// This field is a union of [CallAIGatherEndedWebhookEventData],
	// [CallAIGatherMessageHistoryUpdatedWebhookEventData],
	// [CallAIGatherPartialResultsWebhookEventData], [CallAnsweredWebhookEventData],
	// [CallBridgedWebhookEventData], [CallConversationEndedWebhookEventData],
	// [CallConversationInsightsGeneratedWebhookEventData],
	// [CallDtmfReceivedWebhookEventData], [CallEnqueuedWebhookEventData],
	// [CallForkStartedWebhookEventData], [CallForkStoppedWebhookEventData],
	// [CallGatherEndedWebhookEventData], [CallHangupWebhookEventData],
	// [CallInitiatedWebhookEventData], [CallLeftQueueWebhookEventData],
	// [CallMachineDetectionEndedWebhookEventData],
	// [CallMachineGreetingEndedWebhookEventData],
	// [CallMachinePremiumDetectionEndedWebhookEventData],
	// [CallMachinePremiumGreetingEndedWebhookEventData],
	// [CallPlaybackEndedWebhookEventData], [CallPlaybackStartedWebhookEventData],
	// [CallRecordingErrorWebhookEventData], [CallRecordingSavedWebhookEventData],
	// [CallRecordingTranscriptionSavedWebhookEventData],
	// [CallReferCompletedWebhookEventData], [CallReferFailedWebhookEventData],
	// [CallReferStartedWebhookEventData], [CallSiprecFailedWebhookEventData],
	// [CallSiprecStartedWebhookEventData], [CallSiprecStoppedWebhookEventData],
	// [CallSpeakEndedWebhookEventData], [CallSpeakStartedWebhookEventData],
	// [CallStreamingFailed], [CallStreamingStarted], [CallStreamingStopped],
	// [ConferenceCreatedWebhookEventData], [ConferenceEndedWebhookEventData],
	// [ConferenceParticipantJoinedWebhookEventData],
	// [ConferenceParticipantLeftWebhookEventData],
	// [ConferenceParticipantPlaybackEndedWebhookEventData],
	// [ConferenceParticipantPlaybackStartedWebhookEventData],
	// [ConferenceParticipantSpeakEndedWebhookEventData],
	// [ConferenceParticipantSpeakStartedWebhookEventData],
	// [ConferencePlaybackEndedWebhookEventData],
	// [ConferencePlaybackStartedWebhookEventData],
	// [ConferenceRecordingSavedWebhookEventData],
	// [ConferenceSpeakEndedWebhookEventData],
	// [ConferenceSpeakStartedWebhookEventData], [DeliveryUpdateWebhookEventData],
	// [FaxDeliveredWebhookEventData], [FaxFailedWebhookEventData],
	// [FaxMediaProcessedWebhookEventData], [FaxQueuedWebhookEventData],
	// [FaxSendingStartedWebhookEventData], [InboundMessageWebhookEventData],
	// [NumberOrderStatusUpdateWebhookEventData], [ReplacedLinkClickWebhookEventData],
	// [TranscriptionWebhookEventData]
	Data UnsafeUnwrapWebhookEventUnionData `json:"data"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	BrandID string `json:"brandId"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	CampaignID string `json:"campaignId"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	CreateDate string `json:"createDate"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	CspID string `json:"cspId"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	Description string `json:"description"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	Status CampaignStatusUpdateWebhookEventStatus `json:"status"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	Type CampaignStatusUpdateWebhookEventType `json:"type"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	ID string `json:"id"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	EventType ConferenceFloorChangedWebhookEventEventType `json:"event_type"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	Payload ConferenceFloorChangedWebhookEventPayload `json:"payload"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	RecordType ConferenceFloorChangedWebhookEventRecordType `json:"record_type"`
	// This field is a union of [DeliveryUpdateWebhookEventMeta],
	// [FaxDeliveredWebhookEventMeta], [FaxFailedWebhookEventMeta],
	// [FaxMediaProcessedWebhookEventMeta], [FaxQueuedWebhookEventMeta],
	// [FaxSendingStartedWebhookEventMeta], [NumberOrderStatusUpdateWebhookEventMeta]
	Meta UnsafeUnwrapWebhookEventUnionMeta `json:"meta"`
	JSON struct {
		Data                respjson.Field
		BrandID             respjson.Field
		CampaignID          respjson.Field
		CreateDate          respjson.Field
		CspID               respjson.Field
		Description         respjson.Field
		IsTMobileRegistered respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		ID                  respjson.Field
		EventType           respjson.Field
		Payload             respjson.Field
		RecordType          respjson.Field
		Meta                respjson.Field
		raw                 string
	} `json:"-"`
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallAIGatherEndedEvent() (v CallAIGatherEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallAIGatherMessageHistoryUpdatedEvent() (v CallAIGatherMessageHistoryUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallAIGatherPartialResultsEvent() (v CallAIGatherPartialResultsWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallAnsweredEvent() (v CallAnsweredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallBridgedEvent() (v CallBridgedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallConversationEndedEvent() (v CallConversationEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallConversationInsightsGeneratedEvent() (v CallConversationInsightsGeneratedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallDtmfReceivedEvent() (v CallDtmfReceivedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallEnqueuedEvent() (v CallEnqueuedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallForkStartedEvent() (v CallForkStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallForkStoppedEvent() (v CallForkStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallGatherEndedEvent() (v CallGatherEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallHangupEvent() (v CallHangupWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallInitiatedEvent() (v CallInitiatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallLeftQueueEvent() (v CallLeftQueueWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallMachineDetectionEndedEvent() (v CallMachineDetectionEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallMachineGreetingEndedEvent() (v CallMachineGreetingEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallMachinePremiumDetectionEndedEvent() (v CallMachinePremiumDetectionEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallMachinePremiumGreetingEndedEvent() (v CallMachinePremiumGreetingEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallPlaybackEndedEvent() (v CallPlaybackEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallPlaybackStartedEvent() (v CallPlaybackStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallRecordingErrorEvent() (v CallRecordingErrorWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallRecordingSavedEvent() (v CallRecordingSavedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallRecordingTranscriptionSavedEvent() (v CallRecordingTranscriptionSavedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallReferCompletedEvent() (v CallReferCompletedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallReferFailedEvent() (v CallReferFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallReferStartedEvent() (v CallReferStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSiprecFailedEvent() (v CallSiprecFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSiprecStartedEvent() (v CallSiprecStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsSiprecStoppedEvent() (v CallSiprecStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallSpeakEndedEvent() (v CallSpeakEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCallSpeakStartedEvent() (v CallSpeakStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsStreamingFailedEvent() (v CallStreamingFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsStreamingStartedEvent() (v CallStreamingStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsStreamingStoppedEvent() (v CallStreamingStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsCampaignStatusUpdateEvent() (v CampaignStatusUpdateWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceCreatedEvent() (v ConferenceCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceEndedEvent() (v ConferenceEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceFloorChanged() (v ConferenceFloorChangedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceParticipantJoinedEvent() (v ConferenceParticipantJoinedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceParticipantLeftEvent() (v ConferenceParticipantLeftWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceParticipantPlaybackEndedEvent() (v ConferenceParticipantPlaybackEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceParticipantPlaybackStartedEvent() (v ConferenceParticipantPlaybackStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceParticipantSpeakEndedEvent() (v ConferenceParticipantSpeakEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceParticipantSpeakStartedEvent() (v ConferenceParticipantSpeakStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferencePlaybackEndedEvent() (v ConferencePlaybackEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferencePlaybackStartedEvent() (v ConferencePlaybackStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceRecordingSavedEvent() (v ConferenceRecordingSavedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceSpeakEndedEvent() (v ConferenceSpeakEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceSpeakStartedEvent() (v ConferenceSpeakStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsDeliveryUpdateWebhookEvent() (v DeliveryUpdateWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxDeliveredWebhookEvent() (v FaxDeliveredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxFailedWebhookEvent() (v FaxFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxMediaProcessedWebhookEvent() (v FaxMediaProcessedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxQueuedWebhookEvent() (v FaxQueuedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxSendingStartedWebhookEvent() (v FaxSendingStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsInboundMessageWebhookEvent() (v InboundMessageWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsNumberOrderEvent() (v NumberOrderStatusUpdateWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsReplacedLinkClickWebhookEvent() (v ReplacedLinkClickWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsTranscriptionEvent() (v TranscriptionWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnsafeUnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnsafeUnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionData is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionData struct {
	ID         string    `json:"id"`
	EventType  string    `json:"event_type"`
	OccurredAt time.Time `json:"occurred_at"`
	// This field is a union of [CallAIGatherEndedWebhookEventDataPayload],
	// [CallAIGatherMessageHistoryUpdatedWebhookEventDataPayload],
	// [CallAIGatherPartialResultsWebhookEventDataPayload],
	// [CallAnsweredWebhookEventDataPayload], [CallBridgedWebhookEventDataPayload],
	// [CallConversationEndedWebhookEventDataPayload],
	// [CallConversationInsightsGeneratedWebhookEventDataPayload],
	// [CallDtmfReceivedWebhookEventDataPayload],
	// [CallEnqueuedWebhookEventDataPayload], [CallForkStartedWebhookEventDataPayload],
	// [CallForkStoppedWebhookEventDataPayload],
	// [CallGatherEndedWebhookEventDataPayload], [CallHangupWebhookEventDataPayload],
	// [CallInitiatedWebhookEventDataPayload], [CallLeftQueueWebhookEventDataPayload],
	// [CallMachineDetectionEndedWebhookEventDataPayload],
	// [CallMachineGreetingEndedWebhookEventDataPayload],
	// [CallMachinePremiumDetectionEndedWebhookEventDataPayload],
	// [CallMachinePremiumGreetingEndedWebhookEventDataPayload],
	// [CallPlaybackEndedWebhookEventDataPayload],
	// [CallPlaybackStartedWebhookEventDataPayload],
	// [CallRecordingErrorWebhookEventDataPayload],
	// [CallRecordingSavedWebhookEventDataPayload],
	// [CallRecordingTranscriptionSavedWebhookEventDataPayload],
	// [CallReferCompletedWebhookEventDataPayload],
	// [CallReferFailedWebhookEventDataPayload],
	// [CallReferStartedWebhookEventDataPayload],
	// [CallSiprecFailedWebhookEventDataPayload],
	// [CallSiprecStartedWebhookEventDataPayload],
	// [CallSiprecStoppedWebhookEventDataPayload],
	// [CallSpeakEndedWebhookEventDataPayload],
	// [CallSpeakStartedWebhookEventDataPayload], [CallStreamingFailedPayload],
	// [CallStreamingStartedPayload], [CallStreamingStoppedPayload],
	// [ConferenceCreatedWebhookEventDataPayload],
	// [ConferenceEndedWebhookEventDataPayload],
	// [ConferenceParticipantJoinedWebhookEventDataPayload],
	// [ConferenceParticipantLeftWebhookEventDataPayload],
	// [ConferenceParticipantPlaybackEndedWebhookEventDataPayload],
	// [ConferenceParticipantPlaybackStartedWebhookEventDataPayload],
	// [ConferenceParticipantSpeakEndedWebhookEventDataPayload],
	// [ConferenceParticipantSpeakStartedWebhookEventDataPayload],
	// [ConferencePlaybackEndedWebhookEventDataPayload],
	// [ConferencePlaybackStartedWebhookEventDataPayload],
	// [ConferenceRecordingSavedWebhookEventDataPayload],
	// [ConferenceSpeakEndedWebhookEventDataPayload],
	// [ConferenceSpeakStartedWebhookEventDataPayload], [OutboundMessagePayload],
	// [FaxDeliveredWebhookEventDataPayload], [FaxFailedWebhookEventDataPayload],
	// [FaxMediaProcessedWebhookEventDataPayload], [FaxQueuedWebhookEventDataPayload],
	// [FaxSendingStartedWebhookEventDataPayload], [shared.InboundMessagePayload],
	// [NumberOrderWithPhoneNumbers], [TranscriptionWebhookEventDataPayload]
	Payload    UnsafeUnwrapWebhookEventUnionDataPayload `json:"payload"`
	RecordType string                                   `json:"record_type"`
	// This field is from variant [CallConversationEndedWebhookEventData].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	MessageID string `json:"message_id"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	TimeClicked time.Time `json:"time_clicked"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	To string `json:"to"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	URL  string `json:"url"`
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		CreatedAt   respjson.Field
		MessageID   respjson.Field
		TimeClicked respjson.Field
		To          respjson.Field
		URL         respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayload is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayload
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayload struct {
	CallControlID string `json:"call_control_id"`
	CallLegID     string `json:"call_leg_id"`
	CallSessionID string `json:"call_session_id"`
	ClientState   string `json:"client_state"`
	ConnectionID  string `json:"connection_id"`
	// This field is a union of [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [OutboundMessagePayloadFrom], [string],
	// [string], [string], [string], [string], [shared.InboundMessagePayloadFrom]
	From UnsafeUnwrapWebhookEventUnionDataPayloadFrom `json:"from"`
	// This field is a union of
	// [[]CallAIGatherEndedWebhookEventDataPayloadMessageHistory],
	// [[]CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory],
	// [[]CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory]
	MessageHistory UnsafeUnwrapWebhookEventUnionDataPayloadMessageHistory `json:"message_history"`
	// This field is a union of [map[string]any], [string], [string], [string],
	// [string]
	Result UnsafeUnwrapWebhookEventUnionDataPayloadResult `json:"result"`
	Status string                                         `json:"status"`
	// This field is a union of [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [[]OutboundMessagePayloadTo], [string],
	// [string], [string], [string], [string], [[]shared.InboundMessagePayloadTo]
	To UnsafeUnwrapWebhookEventUnionDataPayloadTo `json:"to"`
	// This field is from variant [CallAIGatherPartialResultsWebhookEventDataPayload].
	PartialResults map[string]any    `json:"partial_results"`
	CustomHeaders  []CustomSipHeader `json:"custom_headers"`
	SipHeaders     []SipHeader       `json:"sip_headers"`
	StartTime      time.Time         `json:"start_time"`
	State          string            `json:"state"`
	Tags           []string          `json:"tags"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	AssistantID      string `json:"assistant_id"`
	CallingPartyType string `json:"calling_party_type"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	ConversationID string `json:"conversation_id"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	DurationSec int64 `json:"duration_sec"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	LlmModel string `json:"llm_model"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	SttModel string `json:"stt_model"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	TtsModelID string `json:"tts_model_id"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	TtsProvider string `json:"tts_provider"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	TtsVoiceID string `json:"tts_voice_id"`
	// This field is from variant
	// [CallConversationInsightsGeneratedWebhookEventDataPayload].
	InsightGroupID string `json:"insight_group_id"`
	// This field is from variant
	// [CallConversationInsightsGeneratedWebhookEventDataPayload].
	Results []CallConversationInsightsGeneratedWebhookEventDataPayloadResult `json:"results"`
	// This field is from variant [CallDtmfReceivedWebhookEventDataPayload].
	Digit string `json:"digit"`
	// This field is from variant [CallEnqueuedWebhookEventDataPayload].
	CurrentPosition int64  `json:"current_position"`
	Queue           string `json:"queue"`
	// This field is from variant [CallEnqueuedWebhookEventDataPayload].
	QueueAvgWaitTimeSecs int64  `json:"queue_avg_wait_time_secs"`
	StreamType           string `json:"stream_type"`
	// This field is from variant [CallGatherEndedWebhookEventDataPayload].
	Digits string `json:"digits"`
	// This field is from variant [CallHangupWebhookEventDataPayload].
	CallQualityStats CallHangupWebhookEventDataPayloadCallQualityStats `json:"call_quality_stats"`
	HangupCause      string                                            `json:"hangup_cause"`
	// This field is from variant [CallHangupWebhookEventDataPayload].
	HangupSource string `json:"hangup_source"`
	// This field is from variant [CallHangupWebhookEventDataPayload].
	SipHangupCause string `json:"sip_hangup_cause"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	CallScreeningResult string `json:"call_screening_result"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	CallerIDName string `json:"caller_id_name"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	ConnectionCodecs string `json:"connection_codecs"`
	Direction        string `json:"direction"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	OfferedCodecs string `json:"offered_codecs"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	ShakenStirAttestation string `json:"shaken_stir_attestation"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	ShakenStirValidated bool `json:"shaken_stir_validated"`
	// This field is from variant [CallLeftQueueWebhookEventDataPayload].
	QueuePosition int64  `json:"queue_position"`
	Reason        string `json:"reason"`
	// This field is from variant [CallLeftQueueWebhookEventDataPayload].
	WaitTimeSecs int64  `json:"wait_time_secs"`
	MediaName    string `json:"media_name"`
	MediaURL     string `json:"media_url"`
	Overlay      bool   `json:"overlay"`
	// This field is from variant [CallPlaybackEndedWebhookEventDataPayload].
	StatusDetail string `json:"status_detail"`
	Channels     string `json:"channels"`
	// This field is a union of
	// [CallRecordingSavedWebhookEventDataPayloadPublicRecordingURLs],
	// [ConferenceRecordingSavedWebhookEventDataPayloadPublicRecordingURLs]
	PublicRecordingURLs UnsafeUnwrapWebhookEventUnionDataPayloadPublicRecordingURLs `json:"public_recording_urls"`
	RecordingEndedAt    time.Time                                                   `json:"recording_ended_at"`
	RecordingStartedAt  time.Time                                                   `json:"recording_started_at"`
	// This field is a union of
	// [CallRecordingSavedWebhookEventDataPayloadRecordingURLs],
	// [ConferenceRecordingSavedWebhookEventDataPayloadRecordingURLs]
	RecordingURLs UnsafeUnwrapWebhookEventUnionDataPayloadRecordingURLs `json:"recording_urls"`
	RecordingID   string                                                `json:"recording_id"`
	// This field is from variant
	// [CallRecordingTranscriptionSavedWebhookEventDataPayload].
	RecordingTranscriptionID string `json:"recording_transcription_id"`
	// This field is from variant
	// [CallRecordingTranscriptionSavedWebhookEventDataPayload].
	TranscriptionText string `json:"transcription_text"`
	SipNotifyResponse int64  `json:"sip_notify_response"`
	// This field is from variant [CallSiprecFailedWebhookEventDataPayload].
	FailureCause  string `json:"failure_cause"`
	FailureReason string `json:"failure_reason"`
	// This field is from variant [CallStreamingFailedPayload].
	StreamID string `json:"stream_id"`
	// This field is from variant [CallStreamingFailedPayload].
	StreamParams         CallStreamingFailedPayloadStreamParams `json:"stream_params"`
	StreamURL            string                                 `json:"stream_url"`
	ConferenceID         string                                 `json:"conference_id"`
	OccurredAt           time.Time                              `json:"occurred_at"`
	CreatorCallSessionID string                                 `json:"creator_call_session_id"`
	// This field is from variant [ConferenceRecordingSavedWebhookEventDataPayload].
	Format string `json:"format"`
	ID     string `json:"id"`
	// This field is a union of [[]OutboundMessagePayloadCc],
	// [[]shared.InboundMessagePayloadCc]
	Cc          UnsafeUnwrapWebhookEventUnionDataPayloadCc `json:"cc"`
	CompletedAt time.Time                                  `json:"completed_at"`
	// This field is a union of [OutboundMessagePayloadCost],
	// [shared.InboundMessagePayloadCost]
	Cost UnsafeUnwrapWebhookEventUnionDataPayloadCost `json:"cost"`
	// This field is a union of [OutboundMessagePayloadCostBreakdown],
	// [shared.InboundMessagePayloadCostBreakdown]
	CostBreakdown UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdown `json:"cost_breakdown"`
	Encoding      string                                                `json:"encoding"`
	Errors        []MessagingError                                      `json:"errors"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media                 UnsafeUnwrapWebhookEventUnionDataPayloadMedia `json:"media"`
	MessagingProfileID    string                                        `json:"messaging_profile_id"`
	OrganizationID        string                                        `json:"organization_id"`
	Parts                 int64                                         `json:"parts"`
	ReceivedAt            time.Time                                     `json:"received_at"`
	RecordType            string                                        `json:"record_type"`
	SentAt                time.Time                                     `json:"sent_at"`
	Subject               string                                        `json:"subject"`
	TcrCampaignBillable   bool                                          `json:"tcr_campaign_billable"`
	TcrCampaignID         string                                        `json:"tcr_campaign_id"`
	TcrCampaignRegistered string                                        `json:"tcr_campaign_registered"`
	Text                  string                                        `json:"text"`
	Type                  string                                        `json:"type"`
	ValidUntil            time.Time                                     `json:"valid_until"`
	WebhookFailoverURL    string                                        `json:"webhook_failover_url"`
	WebhookURL            string                                        `json:"webhook_url"`
	// This field is from variant [FaxDeliveredWebhookEventDataPayload].
	CallDurationSecs int64  `json:"call_duration_secs"`
	FaxID            string `json:"fax_id"`
	OriginalMediaURL string `json:"original_media_url"`
	// This field is from variant [FaxDeliveredWebhookEventDataPayload].
	PageCount int64  `json:"page_count"`
	UserID    string `json:"user_id"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	BillingGroupID string `json:"billing_group_id"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	CustomerReference string `json:"customer_reference"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	PhoneNumbersCount int64 `json:"phone_numbers_count"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	RequirementsMet bool `json:"requirements_met"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	SubNumberOrdersIDs []string `json:"sub_number_orders_ids"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [TranscriptionWebhookEventDataPayload].
	TranscriptionData TranscriptionWebhookEventDataPayloadTranscriptionData `json:"transcription_data"`
	JSON              struct {
		CallControlID            respjson.Field
		CallLegID                respjson.Field
		CallSessionID            respjson.Field
		ClientState              respjson.Field
		ConnectionID             respjson.Field
		From                     respjson.Field
		MessageHistory           respjson.Field
		Result                   respjson.Field
		Status                   respjson.Field
		To                       respjson.Field
		PartialResults           respjson.Field
		CustomHeaders            respjson.Field
		SipHeaders               respjson.Field
		StartTime                respjson.Field
		State                    respjson.Field
		Tags                     respjson.Field
		AssistantID              respjson.Field
		CallingPartyType         respjson.Field
		ConversationID           respjson.Field
		DurationSec              respjson.Field
		LlmModel                 respjson.Field
		SttModel                 respjson.Field
		TtsModelID               respjson.Field
		TtsProvider              respjson.Field
		TtsVoiceID               respjson.Field
		InsightGroupID           respjson.Field
		Results                  respjson.Field
		Digit                    respjson.Field
		CurrentPosition          respjson.Field
		Queue                    respjson.Field
		QueueAvgWaitTimeSecs     respjson.Field
		StreamType               respjson.Field
		Digits                   respjson.Field
		CallQualityStats         respjson.Field
		HangupCause              respjson.Field
		HangupSource             respjson.Field
		SipHangupCause           respjson.Field
		CallScreeningResult      respjson.Field
		CallerIDName             respjson.Field
		ConnectionCodecs         respjson.Field
		Direction                respjson.Field
		OfferedCodecs            respjson.Field
		ShakenStirAttestation    respjson.Field
		ShakenStirValidated      respjson.Field
		QueuePosition            respjson.Field
		Reason                   respjson.Field
		WaitTimeSecs             respjson.Field
		MediaName                respjson.Field
		MediaURL                 respjson.Field
		Overlay                  respjson.Field
		StatusDetail             respjson.Field
		Channels                 respjson.Field
		PublicRecordingURLs      respjson.Field
		RecordingEndedAt         respjson.Field
		RecordingStartedAt       respjson.Field
		RecordingURLs            respjson.Field
		RecordingID              respjson.Field
		RecordingTranscriptionID respjson.Field
		TranscriptionText        respjson.Field
		SipNotifyResponse        respjson.Field
		FailureCause             respjson.Field
		FailureReason            respjson.Field
		StreamID                 respjson.Field
		StreamParams             respjson.Field
		StreamURL                respjson.Field
		ConferenceID             respjson.Field
		OccurredAt               respjson.Field
		CreatorCallSessionID     respjson.Field
		Format                   respjson.Field
		ID                       respjson.Field
		Cc                       respjson.Field
		CompletedAt              respjson.Field
		Cost                     respjson.Field
		CostBreakdown            respjson.Field
		Encoding                 respjson.Field
		Errors                   respjson.Field
		Media                    respjson.Field
		MessagingProfileID       respjson.Field
		OrganizationID           respjson.Field
		Parts                    respjson.Field
		ReceivedAt               respjson.Field
		RecordType               respjson.Field
		SentAt                   respjson.Field
		Subject                  respjson.Field
		TcrCampaignBillable      respjson.Field
		TcrCampaignID            respjson.Field
		TcrCampaignRegistered    respjson.Field
		Text                     respjson.Field
		Type                     respjson.Field
		ValidUntil               respjson.Field
		WebhookFailoverURL       respjson.Field
		WebhookURL               respjson.Field
		CallDurationSecs         respjson.Field
		FaxID                    respjson.Field
		OriginalMediaURL         respjson.Field
		PageCount                respjson.Field
		UserID                   respjson.Field
		BillingGroupID           respjson.Field
		CreatedAt                respjson.Field
		CustomerReference        respjson.Field
		PhoneNumbers             respjson.Field
		PhoneNumbersCount        respjson.Field
		RequirementsMet          respjson.Field
		SubNumberOrdersIDs       respjson.Field
		UpdatedAt                respjson.Field
		TranscriptionData        respjson.Field
		raw                      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadFrom is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayloadFrom
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type UnsafeUnwrapWebhookEventUnionDataPayloadFrom struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString    string `json:",inline"`
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	// This field is from variant [shared.InboundMessagePayloadFrom].
	Status string `json:"status"`
	JSON   struct {
		OfString    respjson.Field
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadMessageHistory is an implicit subunion
// of [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataPayloadMessageHistory provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCallAIGatherEndedWebhookEventDataPayloadMessageHistoryArray
// OfCallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistoryArray
// OfCallAIGatherPartialResultsWebhookEventDataPayloadMessageHistoryArray]
type UnsafeUnwrapWebhookEventUnionDataPayloadMessageHistory struct {
	// This field will be present if the value is a
	// [[]CallAIGatherEndedWebhookEventDataPayloadMessageHistory] instead of an object.
	OfCallAIGatherEndedWebhookEventDataPayloadMessageHistoryArray []CallAIGatherEndedWebhookEventDataPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory]
	// instead of an object.
	OfCallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistoryArray []CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory] instead of
	// an object.
	OfCallAIGatherPartialResultsWebhookEventDataPayloadMessageHistoryArray []CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory `json:",inline"`
	JSON                                                                   struct {
		OfCallAIGatherEndedWebhookEventDataPayloadMessageHistoryArray                 respjson.Field
		OfCallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistoryArray respjson.Field
		OfCallAIGatherPartialResultsWebhookEventDataPayloadMessageHistoryArray        respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadResult is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayloadResult
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCallAIGatherEndedWebhookEventDataPayloadResult
// OfCallMachinePremiumGreetingEndedWebhookEventDataPayloadResult]
type UnsafeUnwrapWebhookEventUnionDataPayloadResult struct {
	// This field will be present if the value is a [any] instead of an object.
	OfCallAIGatherEndedWebhookEventDataPayloadResult any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfCallMachinePremiumGreetingEndedWebhookEventDataPayloadResult string `json:",inline"`
	JSON                                                           struct {
		OfCallAIGatherEndedWebhookEventDataPayloadResult               respjson.Field
		OfCallMachinePremiumGreetingEndedWebhookEventDataPayloadResult respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadTo is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayloadTo
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOutboundMessagePayloadToArray
// OfInboundMessagePayloadToArray]
type UnsafeUnwrapWebhookEventUnionDataPayloadTo struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]OutboundMessagePayloadTo]
	// instead of an object.
	OfOutboundMessagePayloadToArray []OutboundMessagePayloadTo `json:",inline"`
	// This field will be present if the value is a [[]shared.InboundMessagePayloadTo]
	// instead of an object.
	OfInboundMessagePayloadToArray []shared.InboundMessagePayloadTo `json:",inline"`
	JSON                           struct {
		OfString                        respjson.Field
		OfOutboundMessagePayloadToArray respjson.Field
		OfInboundMessagePayloadToArray  respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadPublicRecordingURLs is an implicit
// subunion of [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataPayloadPublicRecordingURLs provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayloadPublicRecordingURLs struct {
	MP3  string `json:"mp3"`
	Wav  string `json:"wav"`
	JSON struct {
		MP3 respjson.Field
		Wav respjson.Field
		raw string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadPublicRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadRecordingURLs is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataPayloadRecordingURLs provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayloadRecordingURLs struct {
	MP3  string `json:"mp3"`
	Wav  string `json:"wav"`
	JSON struct {
		MP3 respjson.Field
		Wav respjson.Field
		raw string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadCc is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayloadCc
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadCcArray OfInboundMessagePayloadCcArray]
type UnsafeUnwrapWebhookEventUnionDataPayloadCc struct {
	// This field will be present if the value is a [[]OutboundMessagePayloadCc]
	// instead of an object.
	OfOutboundMessagePayloadCcArray []OutboundMessagePayloadCc `json:",inline"`
	// This field will be present if the value is a [[]shared.InboundMessagePayloadCc]
	// instead of an object.
	OfInboundMessagePayloadCcArray []shared.InboundMessagePayloadCc `json:",inline"`
	JSON                           struct {
		OfOutboundMessagePayloadCcArray respjson.Field
		OfInboundMessagePayloadCcArray  respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadCc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadCost is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayloadCost
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayloadCost struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdown is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdown provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdown struct {
	// This field is a union of [OutboundMessagePayloadCostBreakdownCarrierFee],
	// [shared.InboundMessagePayloadCostBreakdownCarrierFee]
	CarrierFee UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee `json:"carrier_fee"`
	// This field is a union of [OutboundMessagePayloadCostBreakdownRate],
	// [shared.InboundMessagePayloadCostBreakdownRate]
	Rate UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownRate `json:"rate"`
	JSON struct {
		CarrierFee respjson.Field
		Rate       respjson.Field
		raw        string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee is an implicit
// subunion of [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownRate is an implicit
// subunion of [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownRate provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownRate struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadCostBreakdownRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataPayloadMedia is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataPayloadMedia
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadMedia OfInboundMessagePayloadMedia]
type UnsafeUnwrapWebhookEventUnionDataPayloadMedia struct {
	// This field will be present if the value is a [[]OutboundMessagePayloadMedia]
	// instead of an object.
	OfOutboundMessagePayloadMedia []OutboundMessagePayloadMedia `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.InboundMessagePayloadMedia] instead of an object.
	OfInboundMessagePayloadMedia []shared.InboundMessagePayloadMedia `json:",inline"`
	JSON                         struct {
		OfOutboundMessagePayloadMedia respjson.Field
		OfInboundMessagePayloadMedia  respjson.Field
		raw                           string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataPayloadMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionMeta is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionMeta provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionMeta struct {
	Attempt     int64  `json:"attempt"`
	DeliveredTo string `json:"delivered_to"`
	JSON        struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnion contains all possible properties and values from
// [CallAIGatherEndedWebhookEvent],
// [CallAIGatherMessageHistoryUpdatedWebhookEvent],
// [CallAIGatherPartialResultsWebhookEvent], [CallAnsweredWebhookEvent],
// [CallBridgedWebhookEvent], [CallConversationEndedWebhookEvent],
// [CallConversationInsightsGeneratedWebhookEvent], [CallDtmfReceivedWebhookEvent],
// [CallEnqueuedWebhookEvent], [CallForkStartedWebhookEvent],
// [CallForkStoppedWebhookEvent], [CallGatherEndedWebhookEvent],
// [CallHangupWebhookEvent], [CallInitiatedWebhookEvent],
// [CallLeftQueueWebhookEvent], [CallMachineDetectionEndedWebhookEvent],
// [CallMachineGreetingEndedWebhookEvent],
// [CallMachinePremiumDetectionEndedWebhookEvent],
// [CallMachinePremiumGreetingEndedWebhookEvent], [CallPlaybackEndedWebhookEvent],
// [CallPlaybackStartedWebhookEvent], [CallRecordingErrorWebhookEvent],
// [CallRecordingSavedWebhookEvent], [CallRecordingTranscriptionSavedWebhookEvent],
// [CallReferCompletedWebhookEvent], [CallReferFailedWebhookEvent],
// [CallReferStartedWebhookEvent], [CallSiprecFailedWebhookEvent],
// [CallSiprecStartedWebhookEvent], [CallSiprecStoppedWebhookEvent],
// [CallSpeakEndedWebhookEvent], [CallSpeakStartedWebhookEvent],
// [CallStreamingFailedWebhookEvent], [CallStreamingStartedWebhookEvent],
// [CallStreamingStoppedWebhookEvent], [CampaignStatusUpdateWebhookEvent],
// [ConferenceCreatedWebhookEvent], [ConferenceEndedWebhookEvent],
// [ConferenceFloorChangedWebhookEvent], [ConferenceParticipantJoinedWebhookEvent],
// [ConferenceParticipantLeftWebhookEvent],
// [ConferenceParticipantPlaybackEndedWebhookEvent],
// [ConferenceParticipantPlaybackStartedWebhookEvent],
// [ConferenceParticipantSpeakEndedWebhookEvent],
// [ConferenceParticipantSpeakStartedWebhookEvent],
// [ConferencePlaybackEndedWebhookEvent], [ConferencePlaybackStartedWebhookEvent],
// [ConferenceRecordingSavedWebhookEvent], [ConferenceSpeakEndedWebhookEvent],
// [ConferenceSpeakStartedWebhookEvent], [DeliveryUpdateWebhookEvent],
// [FaxDeliveredWebhookEvent], [FaxFailedWebhookEvent],
// [FaxMediaProcessedWebhookEvent], [FaxQueuedWebhookEvent],
// [FaxSendingStartedWebhookEvent], [InboundMessageWebhookEvent],
// [NumberOrderStatusUpdateWebhookEvent], [ReplacedLinkClickWebhookEvent],
// [TranscriptionWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	// This field is a union of [CallAIGatherEndedWebhookEventData],
	// [CallAIGatherMessageHistoryUpdatedWebhookEventData],
	// [CallAIGatherPartialResultsWebhookEventData], [CallAnsweredWebhookEventData],
	// [CallBridgedWebhookEventData], [CallConversationEndedWebhookEventData],
	// [CallConversationInsightsGeneratedWebhookEventData],
	// [CallDtmfReceivedWebhookEventData], [CallEnqueuedWebhookEventData],
	// [CallForkStartedWebhookEventData], [CallForkStoppedWebhookEventData],
	// [CallGatherEndedWebhookEventData], [CallHangupWebhookEventData],
	// [CallInitiatedWebhookEventData], [CallLeftQueueWebhookEventData],
	// [CallMachineDetectionEndedWebhookEventData],
	// [CallMachineGreetingEndedWebhookEventData],
	// [CallMachinePremiumDetectionEndedWebhookEventData],
	// [CallMachinePremiumGreetingEndedWebhookEventData],
	// [CallPlaybackEndedWebhookEventData], [CallPlaybackStartedWebhookEventData],
	// [CallRecordingErrorWebhookEventData], [CallRecordingSavedWebhookEventData],
	// [CallRecordingTranscriptionSavedWebhookEventData],
	// [CallReferCompletedWebhookEventData], [CallReferFailedWebhookEventData],
	// [CallReferStartedWebhookEventData], [CallSiprecFailedWebhookEventData],
	// [CallSiprecStartedWebhookEventData], [CallSiprecStoppedWebhookEventData],
	// [CallSpeakEndedWebhookEventData], [CallSpeakStartedWebhookEventData],
	// [CallStreamingFailed], [CallStreamingStarted], [CallStreamingStopped],
	// [ConferenceCreatedWebhookEventData], [ConferenceEndedWebhookEventData],
	// [ConferenceParticipantJoinedWebhookEventData],
	// [ConferenceParticipantLeftWebhookEventData],
	// [ConferenceParticipantPlaybackEndedWebhookEventData],
	// [ConferenceParticipantPlaybackStartedWebhookEventData],
	// [ConferenceParticipantSpeakEndedWebhookEventData],
	// [ConferenceParticipantSpeakStartedWebhookEventData],
	// [ConferencePlaybackEndedWebhookEventData],
	// [ConferencePlaybackStartedWebhookEventData],
	// [ConferenceRecordingSavedWebhookEventData],
	// [ConferenceSpeakEndedWebhookEventData],
	// [ConferenceSpeakStartedWebhookEventData], [DeliveryUpdateWebhookEventData],
	// [FaxDeliveredWebhookEventData], [FaxFailedWebhookEventData],
	// [FaxMediaProcessedWebhookEventData], [FaxQueuedWebhookEventData],
	// [FaxSendingStartedWebhookEventData], [InboundMessageWebhookEventData],
	// [NumberOrderStatusUpdateWebhookEventData], [ReplacedLinkClickWebhookEventData],
	// [TranscriptionWebhookEventData]
	Data UnwrapWebhookEventUnionData `json:"data"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	BrandID string `json:"brandId"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	CampaignID string `json:"campaignId"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	CreateDate string `json:"createDate"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	CspID string `json:"cspId"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	Description string `json:"description"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	Status CampaignStatusUpdateWebhookEventStatus `json:"status"`
	// This field is from variant [CampaignStatusUpdateWebhookEvent].
	Type CampaignStatusUpdateWebhookEventType `json:"type"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	ID string `json:"id"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	EventType ConferenceFloorChangedWebhookEventEventType `json:"event_type"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	Payload ConferenceFloorChangedWebhookEventPayload `json:"payload"`
	// This field is from variant [ConferenceFloorChangedWebhookEvent].
	RecordType ConferenceFloorChangedWebhookEventRecordType `json:"record_type"`
	// This field is a union of [DeliveryUpdateWebhookEventMeta],
	// [FaxDeliveredWebhookEventMeta], [FaxFailedWebhookEventMeta],
	// [FaxMediaProcessedWebhookEventMeta], [FaxQueuedWebhookEventMeta],
	// [FaxSendingStartedWebhookEventMeta], [NumberOrderStatusUpdateWebhookEventMeta]
	Meta UnwrapWebhookEventUnionMeta `json:"meta"`
	JSON struct {
		Data                respjson.Field
		BrandID             respjson.Field
		CampaignID          respjson.Field
		CreateDate          respjson.Field
		CspID               respjson.Field
		Description         respjson.Field
		IsTMobileRegistered respjson.Field
		Status              respjson.Field
		Type                respjson.Field
		ID                  respjson.Field
		EventType           respjson.Field
		Payload             respjson.Field
		RecordType          respjson.Field
		Meta                respjson.Field
		raw                 string
	} `json:"-"`
}

func (u UnwrapWebhookEventUnion) AsCallAIGatherEndedEvent() (v CallAIGatherEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallAIGatherMessageHistoryUpdatedEvent() (v CallAIGatherMessageHistoryUpdatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallAIGatherPartialResultsEvent() (v CallAIGatherPartialResultsWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallAnsweredEvent() (v CallAnsweredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallBridgedEvent() (v CallBridgedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallConversationEndedEvent() (v CallConversationEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallConversationInsightsGeneratedEvent() (v CallConversationInsightsGeneratedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallDtmfReceivedEvent() (v CallDtmfReceivedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallEnqueuedEvent() (v CallEnqueuedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallForkStartedEvent() (v CallForkStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallForkStoppedEvent() (v CallForkStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallGatherEndedEvent() (v CallGatherEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallHangupEvent() (v CallHangupWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallInitiatedEvent() (v CallInitiatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallLeftQueueEvent() (v CallLeftQueueWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallMachineDetectionEndedEvent() (v CallMachineDetectionEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallMachineGreetingEndedEvent() (v CallMachineGreetingEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallMachinePremiumDetectionEndedEvent() (v CallMachinePremiumDetectionEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallMachinePremiumGreetingEndedEvent() (v CallMachinePremiumGreetingEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallPlaybackEndedEvent() (v CallPlaybackEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallPlaybackStartedEvent() (v CallPlaybackStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallRecordingErrorEvent() (v CallRecordingErrorWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallRecordingSavedEvent() (v CallRecordingSavedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallRecordingTranscriptionSavedEvent() (v CallRecordingTranscriptionSavedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallReferCompletedEvent() (v CallReferCompletedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallReferFailedEvent() (v CallReferFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallReferStartedEvent() (v CallReferStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSiprecFailedEvent() (v CallSiprecFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSiprecStartedEvent() (v CallSiprecStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsSiprecStoppedEvent() (v CallSiprecStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallSpeakEndedEvent() (v CallSpeakEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCallSpeakStartedEvent() (v CallSpeakStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsStreamingFailedEvent() (v CallStreamingFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsStreamingStartedEvent() (v CallStreamingStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsStreamingStoppedEvent() (v CallStreamingStoppedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsCampaignStatusUpdateEvent() (v CampaignStatusUpdateWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceCreatedEvent() (v ConferenceCreatedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceEndedEvent() (v ConferenceEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceFloorChanged() (v ConferenceFloorChangedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceParticipantJoinedEvent() (v ConferenceParticipantJoinedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceParticipantLeftEvent() (v ConferenceParticipantLeftWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceParticipantPlaybackEndedEvent() (v ConferenceParticipantPlaybackEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceParticipantPlaybackStartedEvent() (v ConferenceParticipantPlaybackStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceParticipantSpeakEndedEvent() (v ConferenceParticipantSpeakEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceParticipantSpeakStartedEvent() (v ConferenceParticipantSpeakStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferencePlaybackEndedEvent() (v ConferencePlaybackEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferencePlaybackStartedEvent() (v ConferencePlaybackStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceRecordingSavedEvent() (v ConferenceRecordingSavedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceSpeakEndedEvent() (v ConferenceSpeakEndedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsConferenceSpeakStartedEvent() (v ConferenceSpeakStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsDeliveryUpdateWebhookEvent() (v DeliveryUpdateWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxDeliveredWebhookEvent() (v FaxDeliveredWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxFailedWebhookEvent() (v FaxFailedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxMediaProcessedWebhookEvent() (v FaxMediaProcessedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxQueuedWebhookEvent() (v FaxQueuedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxSendingStartedWebhookEvent() (v FaxSendingStartedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsInboundMessageWebhookEvent() (v InboundMessageWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNumberOrderEvent() (v NumberOrderStatusUpdateWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsReplacedLinkClickWebhookEvent() (v ReplacedLinkClickWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsTranscriptionEvent() (v TranscriptionWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	ID         string    `json:"id"`
	EventType  string    `json:"event_type"`
	OccurredAt time.Time `json:"occurred_at"`
	// This field is a union of [CallAIGatherEndedWebhookEventDataPayload],
	// [CallAIGatherMessageHistoryUpdatedWebhookEventDataPayload],
	// [CallAIGatherPartialResultsWebhookEventDataPayload],
	// [CallAnsweredWebhookEventDataPayload], [CallBridgedWebhookEventDataPayload],
	// [CallConversationEndedWebhookEventDataPayload],
	// [CallConversationInsightsGeneratedWebhookEventDataPayload],
	// [CallDtmfReceivedWebhookEventDataPayload],
	// [CallEnqueuedWebhookEventDataPayload], [CallForkStartedWebhookEventDataPayload],
	// [CallForkStoppedWebhookEventDataPayload],
	// [CallGatherEndedWebhookEventDataPayload], [CallHangupWebhookEventDataPayload],
	// [CallInitiatedWebhookEventDataPayload], [CallLeftQueueWebhookEventDataPayload],
	// [CallMachineDetectionEndedWebhookEventDataPayload],
	// [CallMachineGreetingEndedWebhookEventDataPayload],
	// [CallMachinePremiumDetectionEndedWebhookEventDataPayload],
	// [CallMachinePremiumGreetingEndedWebhookEventDataPayload],
	// [CallPlaybackEndedWebhookEventDataPayload],
	// [CallPlaybackStartedWebhookEventDataPayload],
	// [CallRecordingErrorWebhookEventDataPayload],
	// [CallRecordingSavedWebhookEventDataPayload],
	// [CallRecordingTranscriptionSavedWebhookEventDataPayload],
	// [CallReferCompletedWebhookEventDataPayload],
	// [CallReferFailedWebhookEventDataPayload],
	// [CallReferStartedWebhookEventDataPayload],
	// [CallSiprecFailedWebhookEventDataPayload],
	// [CallSiprecStartedWebhookEventDataPayload],
	// [CallSiprecStoppedWebhookEventDataPayload],
	// [CallSpeakEndedWebhookEventDataPayload],
	// [CallSpeakStartedWebhookEventDataPayload], [CallStreamingFailedPayload],
	// [CallStreamingStartedPayload], [CallStreamingStoppedPayload],
	// [ConferenceCreatedWebhookEventDataPayload],
	// [ConferenceEndedWebhookEventDataPayload],
	// [ConferenceParticipantJoinedWebhookEventDataPayload],
	// [ConferenceParticipantLeftWebhookEventDataPayload],
	// [ConferenceParticipantPlaybackEndedWebhookEventDataPayload],
	// [ConferenceParticipantPlaybackStartedWebhookEventDataPayload],
	// [ConferenceParticipantSpeakEndedWebhookEventDataPayload],
	// [ConferenceParticipantSpeakStartedWebhookEventDataPayload],
	// [ConferencePlaybackEndedWebhookEventDataPayload],
	// [ConferencePlaybackStartedWebhookEventDataPayload],
	// [ConferenceRecordingSavedWebhookEventDataPayload],
	// [ConferenceSpeakEndedWebhookEventDataPayload],
	// [ConferenceSpeakStartedWebhookEventDataPayload], [OutboundMessagePayload],
	// [FaxDeliveredWebhookEventDataPayload], [FaxFailedWebhookEventDataPayload],
	// [FaxMediaProcessedWebhookEventDataPayload], [FaxQueuedWebhookEventDataPayload],
	// [FaxSendingStartedWebhookEventDataPayload], [shared.InboundMessagePayload],
	// [NumberOrderWithPhoneNumbers], [TranscriptionWebhookEventDataPayload]
	Payload    UnwrapWebhookEventUnionDataPayload `json:"payload"`
	RecordType string                             `json:"record_type"`
	// This field is from variant [CallConversationEndedWebhookEventData].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	MessageID string `json:"message_id"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	TimeClicked time.Time `json:"time_clicked"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	To string `json:"to"`
	// This field is from variant [ReplacedLinkClickWebhookEventData].
	URL  string `json:"url"`
	JSON struct {
		ID          respjson.Field
		EventType   respjson.Field
		OccurredAt  respjson.Field
		Payload     respjson.Field
		RecordType  respjson.Field
		CreatedAt   respjson.Field
		MessageID   respjson.Field
		TimeClicked respjson.Field
		To          respjson.Field
		URL         respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayload is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayload provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayload struct {
	CallControlID string `json:"call_control_id"`
	CallLegID     string `json:"call_leg_id"`
	CallSessionID string `json:"call_session_id"`
	ClientState   string `json:"client_state"`
	ConnectionID  string `json:"connection_id"`
	// This field is a union of [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [OutboundMessagePayloadFrom], [string],
	// [string], [string], [string], [string], [shared.InboundMessagePayloadFrom]
	From UnwrapWebhookEventUnionDataPayloadFrom `json:"from"`
	// This field is a union of
	// [[]CallAIGatherEndedWebhookEventDataPayloadMessageHistory],
	// [[]CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory],
	// [[]CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory]
	MessageHistory UnwrapWebhookEventUnionDataPayloadMessageHistory `json:"message_history"`
	// This field is a union of [map[string]any], [string], [string], [string],
	// [string]
	Result UnwrapWebhookEventUnionDataPayloadResult `json:"result"`
	Status string                                   `json:"status"`
	// This field is a union of [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [string], [string], [string], [string],
	// [string], [string], [string], [string], [[]OutboundMessagePayloadTo], [string],
	// [string], [string], [string], [string], [[]shared.InboundMessagePayloadTo]
	To UnwrapWebhookEventUnionDataPayloadTo `json:"to"`
	// This field is from variant [CallAIGatherPartialResultsWebhookEventDataPayload].
	PartialResults map[string]any    `json:"partial_results"`
	CustomHeaders  []CustomSipHeader `json:"custom_headers"`
	SipHeaders     []SipHeader       `json:"sip_headers"`
	StartTime      time.Time         `json:"start_time"`
	State          string            `json:"state"`
	Tags           []string          `json:"tags"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	AssistantID      string `json:"assistant_id"`
	CallingPartyType string `json:"calling_party_type"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	ConversationID string `json:"conversation_id"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	DurationSec int64 `json:"duration_sec"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	LlmModel string `json:"llm_model"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	SttModel string `json:"stt_model"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	TtsModelID string `json:"tts_model_id"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	TtsProvider string `json:"tts_provider"`
	// This field is from variant [CallConversationEndedWebhookEventDataPayload].
	TtsVoiceID string `json:"tts_voice_id"`
	// This field is from variant
	// [CallConversationInsightsGeneratedWebhookEventDataPayload].
	InsightGroupID string `json:"insight_group_id"`
	// This field is from variant
	// [CallConversationInsightsGeneratedWebhookEventDataPayload].
	Results []CallConversationInsightsGeneratedWebhookEventDataPayloadResult `json:"results"`
	// This field is from variant [CallDtmfReceivedWebhookEventDataPayload].
	Digit string `json:"digit"`
	// This field is from variant [CallEnqueuedWebhookEventDataPayload].
	CurrentPosition int64  `json:"current_position"`
	Queue           string `json:"queue"`
	// This field is from variant [CallEnqueuedWebhookEventDataPayload].
	QueueAvgWaitTimeSecs int64  `json:"queue_avg_wait_time_secs"`
	StreamType           string `json:"stream_type"`
	// This field is from variant [CallGatherEndedWebhookEventDataPayload].
	Digits string `json:"digits"`
	// This field is from variant [CallHangupWebhookEventDataPayload].
	CallQualityStats CallHangupWebhookEventDataPayloadCallQualityStats `json:"call_quality_stats"`
	HangupCause      string                                            `json:"hangup_cause"`
	// This field is from variant [CallHangupWebhookEventDataPayload].
	HangupSource string `json:"hangup_source"`
	// This field is from variant [CallHangupWebhookEventDataPayload].
	SipHangupCause string `json:"sip_hangup_cause"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	CallScreeningResult string `json:"call_screening_result"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	CallerIDName string `json:"caller_id_name"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	ConnectionCodecs string `json:"connection_codecs"`
	Direction        string `json:"direction"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	OfferedCodecs string `json:"offered_codecs"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	ShakenStirAttestation string `json:"shaken_stir_attestation"`
	// This field is from variant [CallInitiatedWebhookEventDataPayload].
	ShakenStirValidated bool `json:"shaken_stir_validated"`
	// This field is from variant [CallLeftQueueWebhookEventDataPayload].
	QueuePosition int64  `json:"queue_position"`
	Reason        string `json:"reason"`
	// This field is from variant [CallLeftQueueWebhookEventDataPayload].
	WaitTimeSecs int64  `json:"wait_time_secs"`
	MediaName    string `json:"media_name"`
	MediaURL     string `json:"media_url"`
	Overlay      bool   `json:"overlay"`
	// This field is from variant [CallPlaybackEndedWebhookEventDataPayload].
	StatusDetail string `json:"status_detail"`
	Channels     string `json:"channels"`
	// This field is a union of
	// [CallRecordingSavedWebhookEventDataPayloadPublicRecordingURLs],
	// [ConferenceRecordingSavedWebhookEventDataPayloadPublicRecordingURLs]
	PublicRecordingURLs UnwrapWebhookEventUnionDataPayloadPublicRecordingURLs `json:"public_recording_urls"`
	RecordingEndedAt    time.Time                                             `json:"recording_ended_at"`
	RecordingStartedAt  time.Time                                             `json:"recording_started_at"`
	// This field is a union of
	// [CallRecordingSavedWebhookEventDataPayloadRecordingURLs],
	// [ConferenceRecordingSavedWebhookEventDataPayloadRecordingURLs]
	RecordingURLs UnwrapWebhookEventUnionDataPayloadRecordingURLs `json:"recording_urls"`
	RecordingID   string                                          `json:"recording_id"`
	// This field is from variant
	// [CallRecordingTranscriptionSavedWebhookEventDataPayload].
	RecordingTranscriptionID string `json:"recording_transcription_id"`
	// This field is from variant
	// [CallRecordingTranscriptionSavedWebhookEventDataPayload].
	TranscriptionText string `json:"transcription_text"`
	SipNotifyResponse int64  `json:"sip_notify_response"`
	// This field is from variant [CallSiprecFailedWebhookEventDataPayload].
	FailureCause  string `json:"failure_cause"`
	FailureReason string `json:"failure_reason"`
	// This field is from variant [CallStreamingFailedPayload].
	StreamID string `json:"stream_id"`
	// This field is from variant [CallStreamingFailedPayload].
	StreamParams         CallStreamingFailedPayloadStreamParams `json:"stream_params"`
	StreamURL            string                                 `json:"stream_url"`
	ConferenceID         string                                 `json:"conference_id"`
	OccurredAt           time.Time                              `json:"occurred_at"`
	CreatorCallSessionID string                                 `json:"creator_call_session_id"`
	// This field is from variant [ConferenceRecordingSavedWebhookEventDataPayload].
	Format string `json:"format"`
	ID     string `json:"id"`
	// This field is a union of [[]OutboundMessagePayloadCc],
	// [[]shared.InboundMessagePayloadCc]
	Cc          UnwrapWebhookEventUnionDataPayloadCc `json:"cc"`
	CompletedAt time.Time                            `json:"completed_at"`
	// This field is a union of [OutboundMessagePayloadCost],
	// [shared.InboundMessagePayloadCost]
	Cost UnwrapWebhookEventUnionDataPayloadCost `json:"cost"`
	// This field is a union of [OutboundMessagePayloadCostBreakdown],
	// [shared.InboundMessagePayloadCostBreakdown]
	CostBreakdown UnwrapWebhookEventUnionDataPayloadCostBreakdown `json:"cost_breakdown"`
	Encoding      string                                          `json:"encoding"`
	Errors        []MessagingError                                `json:"errors"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media                 UnwrapWebhookEventUnionDataPayloadMedia `json:"media"`
	MessagingProfileID    string                                  `json:"messaging_profile_id"`
	OrganizationID        string                                  `json:"organization_id"`
	Parts                 int64                                   `json:"parts"`
	ReceivedAt            time.Time                               `json:"received_at"`
	RecordType            string                                  `json:"record_type"`
	SentAt                time.Time                               `json:"sent_at"`
	Subject               string                                  `json:"subject"`
	TcrCampaignBillable   bool                                    `json:"tcr_campaign_billable"`
	TcrCampaignID         string                                  `json:"tcr_campaign_id"`
	TcrCampaignRegistered string                                  `json:"tcr_campaign_registered"`
	Text                  string                                  `json:"text"`
	Type                  string                                  `json:"type"`
	ValidUntil            time.Time                               `json:"valid_until"`
	WebhookFailoverURL    string                                  `json:"webhook_failover_url"`
	WebhookURL            string                                  `json:"webhook_url"`
	// This field is from variant [FaxDeliveredWebhookEventDataPayload].
	CallDurationSecs int64  `json:"call_duration_secs"`
	FaxID            string `json:"fax_id"`
	OriginalMediaURL string `json:"original_media_url"`
	// This field is from variant [FaxDeliveredWebhookEventDataPayload].
	PageCount int64  `json:"page_count"`
	UserID    string `json:"user_id"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	BillingGroupID string `json:"billing_group_id"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	CustomerReference string `json:"customer_reference"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	PhoneNumbersCount int64 `json:"phone_numbers_count"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	RequirementsMet bool `json:"requirements_met"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	SubNumberOrdersIDs []string `json:"sub_number_orders_ids"`
	// This field is from variant [NumberOrderWithPhoneNumbers].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [TranscriptionWebhookEventDataPayload].
	TranscriptionData TranscriptionWebhookEventDataPayloadTranscriptionData `json:"transcription_data"`
	JSON              struct {
		CallControlID            respjson.Field
		CallLegID                respjson.Field
		CallSessionID            respjson.Field
		ClientState              respjson.Field
		ConnectionID             respjson.Field
		From                     respjson.Field
		MessageHistory           respjson.Field
		Result                   respjson.Field
		Status                   respjson.Field
		To                       respjson.Field
		PartialResults           respjson.Field
		CustomHeaders            respjson.Field
		SipHeaders               respjson.Field
		StartTime                respjson.Field
		State                    respjson.Field
		Tags                     respjson.Field
		AssistantID              respjson.Field
		CallingPartyType         respjson.Field
		ConversationID           respjson.Field
		DurationSec              respjson.Field
		LlmModel                 respjson.Field
		SttModel                 respjson.Field
		TtsModelID               respjson.Field
		TtsProvider              respjson.Field
		TtsVoiceID               respjson.Field
		InsightGroupID           respjson.Field
		Results                  respjson.Field
		Digit                    respjson.Field
		CurrentPosition          respjson.Field
		Queue                    respjson.Field
		QueueAvgWaitTimeSecs     respjson.Field
		StreamType               respjson.Field
		Digits                   respjson.Field
		CallQualityStats         respjson.Field
		HangupCause              respjson.Field
		HangupSource             respjson.Field
		SipHangupCause           respjson.Field
		CallScreeningResult      respjson.Field
		CallerIDName             respjson.Field
		ConnectionCodecs         respjson.Field
		Direction                respjson.Field
		OfferedCodecs            respjson.Field
		ShakenStirAttestation    respjson.Field
		ShakenStirValidated      respjson.Field
		QueuePosition            respjson.Field
		Reason                   respjson.Field
		WaitTimeSecs             respjson.Field
		MediaName                respjson.Field
		MediaURL                 respjson.Field
		Overlay                  respjson.Field
		StatusDetail             respjson.Field
		Channels                 respjson.Field
		PublicRecordingURLs      respjson.Field
		RecordingEndedAt         respjson.Field
		RecordingStartedAt       respjson.Field
		RecordingURLs            respjson.Field
		RecordingID              respjson.Field
		RecordingTranscriptionID respjson.Field
		TranscriptionText        respjson.Field
		SipNotifyResponse        respjson.Field
		FailureCause             respjson.Field
		FailureReason            respjson.Field
		StreamID                 respjson.Field
		StreamParams             respjson.Field
		StreamURL                respjson.Field
		ConferenceID             respjson.Field
		OccurredAt               respjson.Field
		CreatorCallSessionID     respjson.Field
		Format                   respjson.Field
		ID                       respjson.Field
		Cc                       respjson.Field
		CompletedAt              respjson.Field
		Cost                     respjson.Field
		CostBreakdown            respjson.Field
		Encoding                 respjson.Field
		Errors                   respjson.Field
		Media                    respjson.Field
		MessagingProfileID       respjson.Field
		OrganizationID           respjson.Field
		Parts                    respjson.Field
		ReceivedAt               respjson.Field
		RecordType               respjson.Field
		SentAt                   respjson.Field
		Subject                  respjson.Field
		TcrCampaignBillable      respjson.Field
		TcrCampaignID            respjson.Field
		TcrCampaignRegistered    respjson.Field
		Text                     respjson.Field
		Type                     respjson.Field
		ValidUntil               respjson.Field
		WebhookFailoverURL       respjson.Field
		WebhookURL               respjson.Field
		CallDurationSecs         respjson.Field
		FaxID                    respjson.Field
		OriginalMediaURL         respjson.Field
		PageCount                respjson.Field
		UserID                   respjson.Field
		BillingGroupID           respjson.Field
		CreatedAt                respjson.Field
		CustomerReference        respjson.Field
		PhoneNumbers             respjson.Field
		PhoneNumbersCount        respjson.Field
		RequirementsMet          respjson.Field
		SubNumberOrdersIDs       respjson.Field
		UpdatedAt                respjson.Field
		TranscriptionData        respjson.Field
		raw                      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadFrom is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadFrom provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type UnwrapWebhookEventUnionDataPayloadFrom struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString    string `json:",inline"`
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	// This field is from variant [shared.InboundMessagePayloadFrom].
	Status string `json:"status"`
	JSON   struct {
		OfString    respjson.Field
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadMessageHistory is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadMessageHistory
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCallAIGatherEndedWebhookEventDataPayloadMessageHistoryArray
// OfCallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistoryArray
// OfCallAIGatherPartialResultsWebhookEventDataPayloadMessageHistoryArray]
type UnwrapWebhookEventUnionDataPayloadMessageHistory struct {
	// This field will be present if the value is a
	// [[]CallAIGatherEndedWebhookEventDataPayloadMessageHistory] instead of an object.
	OfCallAIGatherEndedWebhookEventDataPayloadMessageHistoryArray []CallAIGatherEndedWebhookEventDataPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory]
	// instead of an object.
	OfCallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistoryArray []CallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory] instead of
	// an object.
	OfCallAIGatherPartialResultsWebhookEventDataPayloadMessageHistoryArray []CallAIGatherPartialResultsWebhookEventDataPayloadMessageHistory `json:",inline"`
	JSON                                                                   struct {
		OfCallAIGatherEndedWebhookEventDataPayloadMessageHistoryArray                 respjson.Field
		OfCallAIGatherMessageHistoryUpdatedWebhookEventDataPayloadMessageHistoryArray respjson.Field
		OfCallAIGatherPartialResultsWebhookEventDataPayloadMessageHistoryArray        respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadResult is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadResult provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCallAIGatherEndedWebhookEventDataPayloadResult
// OfCallMachinePremiumGreetingEndedWebhookEventDataPayloadResult]
type UnwrapWebhookEventUnionDataPayloadResult struct {
	// This field will be present if the value is a [any] instead of an object.
	OfCallAIGatherEndedWebhookEventDataPayloadResult any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfCallMachinePremiumGreetingEndedWebhookEventDataPayloadResult string `json:",inline"`
	JSON                                                           struct {
		OfCallAIGatherEndedWebhookEventDataPayloadResult               respjson.Field
		OfCallMachinePremiumGreetingEndedWebhookEventDataPayloadResult respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadTo is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadTo provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfOutboundMessagePayloadToArray
// OfInboundMessagePayloadToArray]
type UnwrapWebhookEventUnionDataPayloadTo struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]OutboundMessagePayloadTo]
	// instead of an object.
	OfOutboundMessagePayloadToArray []OutboundMessagePayloadTo `json:",inline"`
	// This field will be present if the value is a [[]shared.InboundMessagePayloadTo]
	// instead of an object.
	OfInboundMessagePayloadToArray []shared.InboundMessagePayloadTo `json:",inline"`
	JSON                           struct {
		OfString                        respjson.Field
		OfOutboundMessagePayloadToArray respjson.Field
		OfInboundMessagePayloadToArray  respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadPublicRecordingURLs is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadPublicRecordingURLs
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayloadPublicRecordingURLs struct {
	MP3  string `json:"mp3"`
	Wav  string `json:"wav"`
	JSON struct {
		MP3 respjson.Field
		Wav respjson.Field
		raw string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadPublicRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadRecordingURLs is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadRecordingURLs
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayloadRecordingURLs struct {
	MP3  string `json:"mp3"`
	Wav  string `json:"wav"`
	JSON struct {
		MP3 respjson.Field
		Wav respjson.Field
		raw string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadCc is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadCc provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadCcArray OfInboundMessagePayloadCcArray]
type UnwrapWebhookEventUnionDataPayloadCc struct {
	// This field will be present if the value is a [[]OutboundMessagePayloadCc]
	// instead of an object.
	OfOutboundMessagePayloadCcArray []OutboundMessagePayloadCc `json:",inline"`
	// This field will be present if the value is a [[]shared.InboundMessagePayloadCc]
	// instead of an object.
	OfInboundMessagePayloadCcArray []shared.InboundMessagePayloadCc `json:",inline"`
	JSON                           struct {
		OfOutboundMessagePayloadCcArray respjson.Field
		OfInboundMessagePayloadCcArray  respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadCc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadCost is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadCost provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayloadCost struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadCostBreakdown is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadCostBreakdown
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayloadCostBreakdown struct {
	// This field is a union of [OutboundMessagePayloadCostBreakdownCarrierFee],
	// [shared.InboundMessagePayloadCostBreakdownCarrierFee]
	CarrierFee UnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee `json:"carrier_fee"`
	// This field is a union of [OutboundMessagePayloadCostBreakdownRate],
	// [shared.InboundMessagePayloadCostBreakdownRate]
	Rate UnwrapWebhookEventUnionDataPayloadCostBreakdownRate `json:"rate"`
	JSON struct {
		CarrierFee respjson.Field
		Rate       respjson.Field
		raw        string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadCostBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee is an implicit
// subunion of [UnwrapWebhookEventUnion].
// UnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadCostBreakdownCarrierFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadCostBreakdownRate is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadCostBreakdownRate
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataPayloadCostBreakdownRate struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadCostBreakdownRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataPayloadMedia is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataPayloadMedia provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadMedia OfInboundMessagePayloadMedia]
type UnwrapWebhookEventUnionDataPayloadMedia struct {
	// This field will be present if the value is a [[]OutboundMessagePayloadMedia]
	// instead of an object.
	OfOutboundMessagePayloadMedia []OutboundMessagePayloadMedia `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.InboundMessagePayloadMedia] instead of an object.
	OfInboundMessagePayloadMedia []shared.InboundMessagePayloadMedia `json:",inline"`
	JSON                         struct {
		OfOutboundMessagePayloadMedia respjson.Field
		OfInboundMessagePayloadMedia  respjson.Field
		raw                           string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataPayloadMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionMeta is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionMeta provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionMeta struct {
	Attempt     int64  `json:"attempt"`
	DeliveredTo string `json:"delivered_to"`
	JSON        struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		raw         string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
