// Telnyx webhook verification using ED25519 signatures.
//
// This file provides ED25519 signature verification for Telnyx webhooks,
// matching the implementation pattern used in the Python and Node SDKs.
//
// Example usage:
//
//	client := telnyx.NewClient(
//		option.WithAPIKey(os.Getenv("TELNYX_API_KEY")),
//		option.WithPublicKey(os.Getenv("TELNYX_PUBLIC_KEY")), // Base64 from Mission Control
//	)
//
//	// In your webhook handler:
//	event, err := client.Webhooks.Unwrap(payload, req.Header)
//	if err != nil {
//		// Signature verification failed
//	}

package telnyx

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

const (
	// Telnyx webhook signature headers (case-insensitive per HTTP spec)
	webhookSignatureHeader = "telnyx-signature-ed25519"
	webhookTimestampHeader = "telnyx-timestamp"
	// Tolerance for timestamp validation (5 minutes)
	webhookTimestampTolerance = 5 * time.Minute
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

	// Verify the webhook signature using ED25519
	if err := verifyWebhookSignature(payload, headers, key); err != nil {
		return nil, err
	}

	res := &UnwrapWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

// verifyWebhookSignature verifies the ED25519 signature of a Telnyx webhook.
// Telnyx webhooks are signed using ED25519 with the following format:
//   - Header "Telnyx-Signature-Ed25519": Base64-encoded ED25519 signature
//   - Header "Telnyx-Timestamp": Unix timestamp (seconds)
//   - Signed payload: "{timestamp}|{payload}"
func verifyWebhookSignature(payload []byte, headers http.Header, publicKeyB64 string) error {
	// Get required headers
	signature := headers.Get(webhookSignatureHeader)
	timestamp := headers.Get(webhookTimestampHeader)

	if signature == "" || timestamp == "" {
		return errors.New("missing required webhook headers: telnyx-signature-ed25519 and telnyx-timestamp")
	}

	// Decode the base64-encoded public key
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyB64)
	if err != nil {
		return fmt.Errorf("invalid public key encoding: %w", err)
	}

	if len(publicKeyBytes) != ed25519.PublicKeySize {
		return fmt.Errorf("invalid public key size: expected %d bytes, got %d", ed25519.PublicKeySize, len(publicKeyBytes))
	}

	// Decode the base64-encoded signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("invalid signature encoding: %w", err)
	}

	// Build the signed payload: "{timestamp}|{payload}"
	signedPayload := []byte(timestamp + "|" + string(payload))

	// Verify the ED25519 signature
	if !ed25519.Verify(publicKeyBytes, signedPayload, signatureBytes) {
		return errors.New("webhook signature verification failed")
	}

	// Validate timestamp to prevent replay attacks
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %w", err)
	}

	webhookTime := time.Unix(ts, 0)
	now := time.Now()

	if now.Sub(webhookTime) > webhookTimestampTolerance {
		return errors.New("webhook timestamp is too old")
	}

	if webhookTime.Sub(now) > webhookTimestampTolerance {
		return errors.New("webhook timestamp is too far in the future")
	}

	return nil
}

type CallAIGatherEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.ai_gather.ended".
	EventType CallAIGatherEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                `json:"occurred_at" format:"date-time"`
	Payload    CallAIGatherEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallAIGatherEndedRecordType `json:"record_type"`
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
func (r CallAIGatherEnded) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallAIGatherEndedEventType string

const (
	CallAIGatherEndedEventTypeCallAIGatherEnded CallAIGatherEndedEventType = "call.ai_gather.ended"
)

type CallAIGatherEndedPayload struct {
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
	MessageHistory []CallAIGatherEndedPayloadMessageHistory `json:"message_history"`
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
func (r CallAIGatherEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherEndedPayloadMessageHistory struct {
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
func (r CallAIGatherEndedPayloadMessageHistory) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherEndedPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallAIGatherEndedRecordType string

const (
	CallAIGatherEndedRecordTypeEvent CallAIGatherEndedRecordType = "event"
)

type CallAIGatherMessageHistoryUpdated struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.ai_gather.message_history_updated".
	EventType CallAIGatherMessageHistoryUpdatedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    CallAIGatherMessageHistoryUpdatedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallAIGatherMessageHistoryUpdatedRecordType `json:"record_type"`
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
func (r CallAIGatherMessageHistoryUpdated) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherMessageHistoryUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallAIGatherMessageHistoryUpdatedEventType string

const (
	CallAIGatherMessageHistoryUpdatedEventTypeCallAIGatherMessageHistoryUpdated CallAIGatherMessageHistoryUpdatedEventType = "call.ai_gather.message_history_updated"
)

type CallAIGatherMessageHistoryUpdatedPayload struct {
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
	MessageHistory []CallAIGatherMessageHistoryUpdatedPayloadMessageHistory `json:"message_history"`
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
func (r CallAIGatherMessageHistoryUpdatedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherMessageHistoryUpdatedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherMessageHistoryUpdatedPayloadMessageHistory struct {
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
func (r CallAIGatherMessageHistoryUpdatedPayloadMessageHistory) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherMessageHistoryUpdatedPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallAIGatherMessageHistoryUpdatedRecordType string

const (
	CallAIGatherMessageHistoryUpdatedRecordTypeEvent CallAIGatherMessageHistoryUpdatedRecordType = "event"
)

type CallAIGatherPartialResults struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.ai_gather.partial_results".
	EventType CallAIGatherPartialResultsEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                         `json:"occurred_at" format:"date-time"`
	Payload    CallAIGatherPartialResultsPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallAIGatherPartialResultsRecordType `json:"record_type"`
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
func (r CallAIGatherPartialResults) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherPartialResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallAIGatherPartialResultsEventType string

const (
	CallAIGatherPartialResultsEventTypeCallAIGatherPartialResults CallAIGatherPartialResultsEventType = "call.ai_gather.partial_results"
)

type CallAIGatherPartialResultsPayload struct {
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
	MessageHistory []CallAIGatherPartialResultsPayloadMessageHistory `json:"message_history"`
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
func (r CallAIGatherPartialResultsPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherPartialResultsPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallAIGatherPartialResultsPayloadMessageHistory struct {
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
func (r CallAIGatherPartialResultsPayloadMessageHistory) RawJSON() string { return r.JSON.raw }
func (r *CallAIGatherPartialResultsPayloadMessageHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallAIGatherPartialResultsRecordType string

const (
	CallAIGatherPartialResultsRecordTypeEvent CallAIGatherPartialResultsRecordType = "event"
)

type CallAnswered struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.answered".
	EventType CallAnsweredEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time           `json:"occurred_at" format:"date-time"`
	Payload    CallAnsweredPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallAnsweredRecordType `json:"record_type"`
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
func (r CallAnswered) RawJSON() string { return r.JSON.raw }
func (r *CallAnswered) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallAnsweredEventType string

const (
	CallAnsweredEventTypeCallAnswered CallAnsweredEventType = "call.answered"
)

type CallAnsweredPayload struct {
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
func (r CallAnsweredPayload) RawJSON() string { return r.JSON.raw }
func (r *CallAnsweredPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallAnsweredRecordType string

const (
	CallAnsweredRecordTypeEvent CallAnsweredRecordType = "event"
)

type CallBridged struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.bridged".
	EventType CallBridgedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time          `json:"occurred_at" format:"date-time"`
	Payload    CallBridgedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallBridgedRecordType `json:"record_type"`
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
func (r CallBridged) RawJSON() string { return r.JSON.raw }
func (r *CallBridged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallBridgedEventType string

const (
	CallBridgedEventTypeCallBridged CallBridgedEventType = "call.bridged"
)

type CallBridgedPayload struct {
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
func (r CallBridgedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallBridgedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallBridgedRecordType string

const (
	CallBridgedRecordTypeEvent CallBridgedRecordType = "event"
)

type CallConversationEnded struct {
	// Unique identifier for the event.
	ID string `json:"id" format:"uuid"`
	// Timestamp when the event was created in the system.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The type of event being delivered.
	//
	// Any of "call.conversation.ended".
	EventType CallConversationEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                    `json:"occurred_at" format:"date-time"`
	Payload    CallConversationEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallConversationEndedRecordType `json:"record_type"`
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
func (r CallConversationEnded) RawJSON() string { return r.JSON.raw }
func (r *CallConversationEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallConversationEndedEventType string

const (
	CallConversationEndedEventTypeCallConversationEnded CallConversationEndedEventType = "call.conversation.ended"
)

type CallConversationEndedPayload struct {
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
func (r CallConversationEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallConversationEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallConversationEndedRecordType string

const (
	CallConversationEndedRecordTypeEvent CallConversationEndedRecordType = "event"
)

type CallConversationInsightsGenerated struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.conversation_insights.generated".
	EventType CallConversationInsightsGeneratedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                                `json:"occurred_at" format:"date-time"`
	Payload    CallConversationInsightsGeneratedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallConversationInsightsGeneratedRecordType `json:"record_type"`
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
func (r CallConversationInsightsGenerated) RawJSON() string { return r.JSON.raw }
func (r *CallConversationInsightsGenerated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallConversationInsightsGeneratedEventType string

const (
	CallConversationInsightsGeneratedEventTypeCallConversationInsightsGenerated CallConversationInsightsGeneratedEventType = "call.conversation_insights.generated"
)

type CallConversationInsightsGeneratedPayload struct {
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
	Results []CallConversationInsightsGeneratedPayloadResult `json:"results"`
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
func (r CallConversationInsightsGeneratedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallConversationInsightsGeneratedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallConversationInsightsGeneratedPayloadResult struct {
	// ID that is unique to the insight result being generated for the call.
	InsightID string `json:"insight_id"`
	// The result of the insight.
	Result CallConversationInsightsGeneratedPayloadResultResultUnion `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsightID   respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallConversationInsightsGeneratedPayloadResult) RawJSON() string { return r.JSON.raw }
func (r *CallConversationInsightsGeneratedPayloadResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallConversationInsightsGeneratedPayloadResultResultUnion contains all possible
// properties and values from [map[string]any], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfCallConversationInsightsGeneratedPayloadResultResultInsightObjectResultItem
// OfString]
type CallConversationInsightsGeneratedPayloadResultResultUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfCallConversationInsightsGeneratedPayloadResultResultInsightObjectResultItem any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfCallConversationInsightsGeneratedPayloadResultResultInsightObjectResultItem respjson.Field
		OfString                                                                      respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u CallConversationInsightsGeneratedPayloadResultResultUnion) AsInsightObjectResult() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallConversationInsightsGeneratedPayloadResultResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CallConversationInsightsGeneratedPayloadResultResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CallConversationInsightsGeneratedPayloadResultResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallConversationInsightsGeneratedRecordType string

const (
	CallConversationInsightsGeneratedRecordTypeEvent CallConversationInsightsGeneratedRecordType = "event"
)

type CallDtmfReceived struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.dtmf.received".
	EventType CallDtmfReceivedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time               `json:"occurred_at" format:"date-time"`
	Payload    CallDtmfReceivedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallDtmfReceivedRecordType `json:"record_type"`
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
func (r CallDtmfReceived) RawJSON() string { return r.JSON.raw }
func (r *CallDtmfReceived) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallDtmfReceivedEventType string

const (
	CallDtmfReceivedEventTypeCallDtmfReceived CallDtmfReceivedEventType = "call.dtmf.received"
)

type CallDtmfReceivedPayload struct {
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
func (r CallDtmfReceivedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallDtmfReceivedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallDtmfReceivedRecordType string

const (
	CallDtmfReceivedRecordTypeEvent CallDtmfReceivedRecordType = "event"
)

type CallEnqueued struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.enqueued".
	EventType CallEnqueuedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time           `json:"occurred_at" format:"date-time"`
	Payload    CallEnqueuedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallEnqueuedRecordType `json:"record_type"`
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
func (r CallEnqueued) RawJSON() string { return r.JSON.raw }
func (r *CallEnqueued) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallEnqueuedEventType string

const (
	CallEnqueuedEventTypeCallEnqueued CallEnqueuedEventType = "call.enqueued"
)

type CallEnqueuedPayload struct {
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
func (r CallEnqueuedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallEnqueuedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallEnqueuedRecordType string

const (
	CallEnqueuedRecordTypeEvent CallEnqueuedRecordType = "event"
)

type CallForkStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.fork.started".
	EventType CallForkStartedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time              `json:"occurred_at" format:"date-time"`
	Payload    CallForkStartedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallForkStartedRecordType `json:"record_type"`
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
func (r CallForkStarted) RawJSON() string { return r.JSON.raw }
func (r *CallForkStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallForkStartedEventType string

const (
	CallForkStartedEventTypeCallForkStarted CallForkStartedEventType = "call.fork.started"
)

type CallForkStartedPayload struct {
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
func (r CallForkStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallForkStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallForkStartedRecordType string

const (
	CallForkStartedRecordTypeEvent CallForkStartedRecordType = "event"
)

type CallForkStopped struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.fork.stopped".
	EventType CallForkStoppedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time              `json:"occurred_at" format:"date-time"`
	Payload    CallForkStoppedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallForkStoppedRecordType `json:"record_type"`
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
func (r CallForkStopped) RawJSON() string { return r.JSON.raw }
func (r *CallForkStopped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallForkStoppedEventType string

const (
	CallForkStoppedEventTypeCallForkStopped CallForkStoppedEventType = "call.fork.stopped"
)

type CallForkStoppedPayload struct {
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
func (r CallForkStoppedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallForkStoppedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallForkStoppedRecordType string

const (
	CallForkStoppedRecordTypeEvent CallForkStoppedRecordType = "event"
)

type CallGatherEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.gather.ended".
	EventType CallGatherEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time              `json:"occurred_at" format:"date-time"`
	Payload    CallGatherEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallGatherEndedRecordType `json:"record_type"`
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
func (r CallGatherEnded) RawJSON() string { return r.JSON.raw }
func (r *CallGatherEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallGatherEndedEventType string

const (
	CallGatherEndedEventTypeCallGatherEnded CallGatherEndedEventType = "call.gather.ended"
)

type CallGatherEndedPayload struct {
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
func (r CallGatherEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallGatherEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallGatherEndedRecordType string

const (
	CallGatherEndedRecordTypeEvent CallGatherEndedRecordType = "event"
)

type CallHangup struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.hangup".
	EventType CallHangupEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time         `json:"occurred_at" format:"date-time"`
	Payload    CallHangupPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallHangupRecordType `json:"record_type"`
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
func (r CallHangup) RawJSON() string { return r.JSON.raw }
func (r *CallHangup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallHangupEventType string

const (
	CallHangupEventTypeCallHangup CallHangupEventType = "call.hangup"
)

type CallHangupPayload struct {
	// Call ID used to issue commands via Call Control API.
	CallControlID string `json:"call_control_id"`
	// ID that is unique to the call and can be used to correlate webhook events.
	CallLegID string `json:"call_leg_id"`
	// Call quality statistics aggregated from the CHANNEL_HANGUP_COMPLETE event. Only
	// includes metrics that are available (filters out nil values). Returns nil if no
	// metrics are available.
	CallQualityStats CallHangupPayloadCallQualityStats `json:"call_quality_stats" api:"nullable"`
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
func (r CallHangupPayload) RawJSON() string { return r.JSON.raw }
func (r *CallHangupPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Call quality statistics aggregated from the CHANNEL_HANGUP_COMPLETE event. Only
// includes metrics that are available (filters out nil values). Returns nil if no
// metrics are available.
type CallHangupPayloadCallQualityStats struct {
	// Inbound call quality statistics.
	Inbound CallHangupPayloadCallQualityStatsInbound `json:"inbound"`
	// Outbound call quality statistics.
	Outbound CallHangupPayloadCallQualityStatsOutbound `json:"outbound"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inbound     respjson.Field
		Outbound    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallHangupPayloadCallQualityStats) RawJSON() string { return r.JSON.raw }
func (r *CallHangupPayloadCallQualityStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inbound call quality statistics.
type CallHangupPayloadCallQualityStatsInbound struct {
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
func (r CallHangupPayloadCallQualityStatsInbound) RawJSON() string { return r.JSON.raw }
func (r *CallHangupPayloadCallQualityStatsInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound call quality statistics.
type CallHangupPayloadCallQualityStatsOutbound struct {
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
func (r CallHangupPayloadCallQualityStatsOutbound) RawJSON() string { return r.JSON.raw }
func (r *CallHangupPayloadCallQualityStatsOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallHangupRecordType string

const (
	CallHangupRecordTypeEvent CallHangupRecordType = "event"
)

type CallInitiated struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.initiated".
	EventType CallInitiatedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time            `json:"occurred_at" format:"date-time"`
	Payload    CallInitiatedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallInitiatedRecordType `json:"record_type"`
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
func (r CallInitiated) RawJSON() string { return r.JSON.raw }
func (r *CallInitiated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallInitiatedEventType string

const (
	CallInitiatedEventTypeCallInitiated CallInitiatedEventType = "call.initiated"
)

type CallInitiatedPayload struct {
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
func (r CallInitiatedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallInitiatedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallInitiatedRecordType string

const (
	CallInitiatedRecordTypeEvent CallInitiatedRecordType = "event"
)

type CallLeftQueue struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.dequeued".
	EventType CallLeftQueueEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time            `json:"occurred_at" format:"date-time"`
	Payload    CallLeftQueuePayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallLeftQueueRecordType `json:"record_type"`
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
func (r CallLeftQueue) RawJSON() string { return r.JSON.raw }
func (r *CallLeftQueue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallLeftQueueEventType string

const (
	CallLeftQueueEventTypeCallDequeued CallLeftQueueEventType = "call.dequeued"
)

type CallLeftQueuePayload struct {
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
func (r CallLeftQueuePayload) RawJSON() string { return r.JSON.raw }
func (r *CallLeftQueuePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallLeftQueueRecordType string

const (
	CallLeftQueueRecordTypeEvent CallLeftQueueRecordType = "event"
)

type CallMachineDetectionEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.detection.ended".
	EventType CallMachineDetectionEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                        `json:"occurred_at" format:"date-time"`
	Payload    CallMachineDetectionEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallMachineDetectionEndedRecordType `json:"record_type"`
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
func (r CallMachineDetectionEnded) RawJSON() string { return r.JSON.raw }
func (r *CallMachineDetectionEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallMachineDetectionEndedEventType string

const (
	CallMachineDetectionEndedEventTypeCallMachineDetectionEnded CallMachineDetectionEndedEventType = "call.machine.detection.ended"
)

type CallMachineDetectionEndedPayload struct {
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
func (r CallMachineDetectionEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachineDetectionEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallMachineDetectionEndedRecordType string

const (
	CallMachineDetectionEndedRecordTypeEvent CallMachineDetectionEndedRecordType = "event"
)

type CallMachineGreetingEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.greeting.ended".
	EventType CallMachineGreetingEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                       `json:"occurred_at" format:"date-time"`
	Payload    CallMachineGreetingEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallMachineGreetingEndedRecordType `json:"record_type"`
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
func (r CallMachineGreetingEnded) RawJSON() string { return r.JSON.raw }
func (r *CallMachineGreetingEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallMachineGreetingEndedEventType string

const (
	CallMachineGreetingEndedEventTypeCallMachineGreetingEnded CallMachineGreetingEndedEventType = "call.machine.greeting.ended"
)

type CallMachineGreetingEndedPayload struct {
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
func (r CallMachineGreetingEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachineGreetingEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallMachineGreetingEndedRecordType string

const (
	CallMachineGreetingEndedRecordTypeEvent CallMachineGreetingEndedRecordType = "event"
)

type CallMachinePremiumDetectionEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.premium.detection.ended".
	EventType CallMachinePremiumDetectionEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                               `json:"occurred_at" format:"date-time"`
	Payload    CallMachinePremiumDetectionEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallMachinePremiumDetectionEndedRecordType `json:"record_type"`
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
func (r CallMachinePremiumDetectionEnded) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumDetectionEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallMachinePremiumDetectionEndedEventType string

const (
	CallMachinePremiumDetectionEndedEventTypeCallMachinePremiumDetectionEnded CallMachinePremiumDetectionEndedEventType = "call.machine.premium.detection.ended"
)

type CallMachinePremiumDetectionEndedPayload struct {
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
func (r CallMachinePremiumDetectionEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumDetectionEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallMachinePremiumDetectionEndedRecordType string

const (
	CallMachinePremiumDetectionEndedRecordTypeEvent CallMachinePremiumDetectionEndedRecordType = "event"
)

type CallMachinePremiumGreetingEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.machine.premium.greeting.ended".
	EventType CallMachinePremiumGreetingEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                              `json:"occurred_at" format:"date-time"`
	Payload    CallMachinePremiumGreetingEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallMachinePremiumGreetingEndedRecordType `json:"record_type"`
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
func (r CallMachinePremiumGreetingEnded) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumGreetingEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallMachinePremiumGreetingEndedEventType string

const (
	CallMachinePremiumGreetingEndedEventTypeCallMachinePremiumGreetingEnded CallMachinePremiumGreetingEndedEventType = "call.machine.premium.greeting.ended"
)

type CallMachinePremiumGreetingEndedPayload struct {
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
func (r CallMachinePremiumGreetingEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallMachinePremiumGreetingEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallMachinePremiumGreetingEndedRecordType string

const (
	CallMachinePremiumGreetingEndedRecordTypeEvent CallMachinePremiumGreetingEndedRecordType = "event"
)

type CallPlaybackEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.playback.ended".
	EventType CallPlaybackEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                `json:"occurred_at" format:"date-time"`
	Payload    CallPlaybackEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallPlaybackEndedRecordType `json:"record_type"`
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
func (r CallPlaybackEnded) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallPlaybackEndedEventType string

const (
	CallPlaybackEndedEventTypeCallPlaybackEnded CallPlaybackEndedEventType = "call.playback.ended"
)

type CallPlaybackEndedPayload struct {
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
func (r CallPlaybackEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallPlaybackEndedRecordType string

const (
	CallPlaybackEndedRecordTypeEvent CallPlaybackEndedRecordType = "event"
)

type CallPlaybackStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.playback.started".
	EventType CallPlaybackStartedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                  `json:"occurred_at" format:"date-time"`
	Payload    CallPlaybackStartedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallPlaybackStartedRecordType `json:"record_type"`
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
func (r CallPlaybackStarted) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallPlaybackStartedEventType string

const (
	CallPlaybackStartedEventTypeCallPlaybackStarted CallPlaybackStartedEventType = "call.playback.started"
)

type CallPlaybackStartedPayload struct {
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
func (r CallPlaybackStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallPlaybackStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallPlaybackStartedRecordType string

const (
	CallPlaybackStartedRecordTypeEvent CallPlaybackStartedRecordType = "event"
)

type CallRecordingError struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.recording.error".
	EventType CallRecordingErrorEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                 `json:"occurred_at" format:"date-time"`
	Payload    CallRecordingErrorPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallRecordingErrorRecordType `json:"record_type"`
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
func (r CallRecordingError) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallRecordingErrorEventType string

const (
	CallRecordingErrorEventTypeCallRecordingError CallRecordingErrorEventType = "call.recording.error"
)

type CallRecordingErrorPayload struct {
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
func (r CallRecordingErrorPayload) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingErrorPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallRecordingErrorRecordType string

const (
	CallRecordingErrorRecordTypeEvent CallRecordingErrorRecordType = "event"
)

type CallRecordingSaved struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.recording.saved".
	EventType CallRecordingSavedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                 `json:"occurred_at" format:"date-time"`
	Payload    CallRecordingSavedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallRecordingSavedRecordType `json:"record_type"`
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
func (r CallRecordingSaved) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSaved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallRecordingSavedEventType string

const (
	CallRecordingSavedEventTypeCallRecordingSaved CallRecordingSavedEventType = "call.recording.saved"
)

type CallRecordingSavedPayload struct {
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
	PublicRecordingURLs CallRecordingSavedPayloadPublicRecordingURLs `json:"public_recording_urls"`
	// ISO 8601 datetime of when recording ended.
	RecordingEndedAt time.Time `json:"recording_ended_at" format:"date-time"`
	// ISO 8601 datetime of when recording started.
	RecordingStartedAt time.Time `json:"recording_started_at" format:"date-time"`
	// Recording URLs in requested format. These URLs are valid for 10 minutes. After
	// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
	// documentation, or via Mission Control under Reporting -> Recordings.
	RecordingURLs CallRecordingSavedPayloadRecordingURLs `json:"recording_urls"`
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
func (r CallRecordingSavedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. The URL is valid for as long as the file
// exists. For security purposes, this feature is activated on a per request basis.
// Please contact customer support with your Account ID to request activation.
type CallRecordingSavedPayloadPublicRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3" api:"nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedPayloadPublicRecordingURLs) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedPayloadPublicRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. These URLs are valid for 10 minutes. After
// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
// documentation, or via Mission Control under Reporting -> Recordings.
type CallRecordingSavedPayloadRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3" api:"nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallRecordingSavedPayloadRecordingURLs) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingSavedPayloadRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallRecordingSavedRecordType string

const (
	CallRecordingSavedRecordTypeEvent CallRecordingSavedRecordType = "event"
)

type CallRecordingTranscriptionSaved struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.recording.transcription.saved".
	EventType CallRecordingTranscriptionSavedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                              `json:"occurred_at" format:"date-time"`
	Payload    CallRecordingTranscriptionSavedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallRecordingTranscriptionSavedRecordType `json:"record_type"`
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
func (r CallRecordingTranscriptionSaved) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingTranscriptionSaved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallRecordingTranscriptionSavedEventType string

const (
	CallRecordingTranscriptionSavedEventTypeCallRecordingTranscriptionSaved CallRecordingTranscriptionSavedEventType = "call.recording.transcription.saved"
)

type CallRecordingTranscriptionSavedPayload struct {
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
func (r CallRecordingTranscriptionSavedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallRecordingTranscriptionSavedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallRecordingTranscriptionSavedRecordType string

const (
	CallRecordingTranscriptionSavedRecordTypeEvent CallRecordingTranscriptionSavedRecordType = "event"
)

type CallReferCompleted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.refer.completed".
	EventType CallReferCompletedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                 `json:"occurred_at" format:"date-time"`
	Payload    CallReferCompletedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallReferCompletedRecordType `json:"record_type"`
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
func (r CallReferCompleted) RawJSON() string { return r.JSON.raw }
func (r *CallReferCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallReferCompletedEventType string

const (
	CallReferCompletedEventTypeCallReferCompleted CallReferCompletedEventType = "call.refer.completed"
)

type CallReferCompletedPayload struct {
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
func (r CallReferCompletedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallReferCompletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallReferCompletedRecordType string

const (
	CallReferCompletedRecordTypeEvent CallReferCompletedRecordType = "event"
)

type CallReferFailed struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.refer.failed".
	EventType CallReferFailedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time              `json:"occurred_at" format:"date-time"`
	Payload    CallReferFailedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallReferFailedRecordType `json:"record_type"`
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
func (r CallReferFailed) RawJSON() string { return r.JSON.raw }
func (r *CallReferFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallReferFailedEventType string

const (
	CallReferFailedEventTypeCallReferFailed CallReferFailedEventType = "call.refer.failed"
)

type CallReferFailedPayload struct {
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
func (r CallReferFailedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallReferFailedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallReferFailedRecordType string

const (
	CallReferFailedRecordTypeEvent CallReferFailedRecordType = "event"
)

type CallReferStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.refer.started".
	EventType CallReferStartedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time               `json:"occurred_at" format:"date-time"`
	Payload    CallReferStartedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallReferStartedRecordType `json:"record_type"`
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
func (r CallReferStarted) RawJSON() string { return r.JSON.raw }
func (r *CallReferStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallReferStartedEventType string

const (
	CallReferStartedEventTypeCallReferStarted CallReferStartedEventType = "call.refer.started"
)

type CallReferStartedPayload struct {
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
func (r CallReferStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallReferStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallReferStartedRecordType string

const (
	CallReferStartedRecordTypeEvent CallReferStartedRecordType = "event"
)

type CallSiprecFailed struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "siprec.failed".
	EventType CallSiprecFailedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time               `json:"occurred_at" format:"date-time"`
	Payload    CallSiprecFailedPayload `json:"payload"`
	// Identifies the resource.
	//
	// Any of "event".
	RecordType CallSiprecFailedRecordType `json:"record_type"`
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
func (r CallSiprecFailed) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallSiprecFailedEventType string

const (
	CallSiprecFailedEventTypeSiprecFailed CallSiprecFailedEventType = "siprec.failed"
)

type CallSiprecFailedPayload struct {
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
func (r CallSiprecFailedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecFailedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the resource.
type CallSiprecFailedRecordType string

const (
	CallSiprecFailedRecordTypeEvent CallSiprecFailedRecordType = "event"
)

type CallSiprecStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "siprec.started".
	EventType CallSiprecStartedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                `json:"occurred_at" format:"date-time"`
	Payload    CallSiprecStartedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallSiprecStartedRecordType `json:"record_type"`
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
func (r CallSiprecStarted) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallSiprecStartedEventType string

const (
	CallSiprecStartedEventTypeSiprecStarted CallSiprecStartedEventType = "siprec.started"
)

type CallSiprecStartedPayload struct {
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
func (r CallSiprecStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallSiprecStartedRecordType string

const (
	CallSiprecStartedRecordTypeEvent CallSiprecStartedRecordType = "event"
)

type CallSiprecStopped struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "siprec.stopped".
	EventType CallSiprecStoppedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                `json:"occurred_at" format:"date-time"`
	Payload    CallSiprecStoppedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallSiprecStoppedRecordType `json:"record_type"`
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
func (r CallSiprecStopped) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStopped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallSiprecStoppedEventType string

const (
	CallSiprecStoppedEventTypeSiprecStopped CallSiprecStoppedEventType = "siprec.stopped"
)

type CallSiprecStoppedPayload struct {
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
func (r CallSiprecStoppedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSiprecStoppedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallSiprecStoppedRecordType string

const (
	CallSiprecStoppedRecordTypeEvent CallSiprecStoppedRecordType = "event"
)

type CallSpeakEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.speak.ended".
	EventType CallSpeakEndedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time             `json:"occurred_at" format:"date-time"`
	Payload    CallSpeakEndedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallSpeakEndedRecordType `json:"record_type"`
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
func (r CallSpeakEnded) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallSpeakEndedEventType string

const (
	CallSpeakEndedEventTypeCallSpeakEnded CallSpeakEndedEventType = "call.speak.ended"
)

type CallSpeakEndedPayload struct {
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
func (r CallSpeakEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallSpeakEndedRecordType string

const (
	CallSpeakEndedRecordTypeEvent CallSpeakEndedRecordType = "event"
)

type CallSpeakStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.speak.started".
	EventType CallSpeakStartedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time               `json:"occurred_at" format:"date-time"`
	Payload    CallSpeakStartedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType CallSpeakStartedRecordType `json:"record_type"`
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
func (r CallSpeakStarted) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type CallSpeakStartedEventType string

const (
	CallSpeakStartedEventTypeCallSpeakStarted CallSpeakStartedEventType = "call.speak.started"
)

type CallSpeakStartedPayload struct {
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
func (r CallSpeakStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *CallSpeakStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type CallSpeakStartedRecordType string

const (
	CallSpeakStartedRecordTypeEvent CallSpeakStartedRecordType = "event"
)

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

type CampaignStatusUpdate struct {
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
	Status CampaignStatusUpdateStatus `json:"status"`
	// Any of "TELNYX_EVENT", "REGISTRATION", "MNO_REVIEW", "TELNYX_REVIEW",
	// "NUMBER_POOL_PROVISIONED", "NUMBER_POOL_DEPROVISIONED", "TCR_EVENT", "VERIFIED".
	Type CampaignStatusUpdateType `json:"type"`
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
func (r CampaignStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *CampaignStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the campaign.
type CampaignStatusUpdateStatus string

const (
	CampaignStatusUpdateStatusAccepted CampaignStatusUpdateStatus = "ACCEPTED"
	CampaignStatusUpdateStatusRejected CampaignStatusUpdateStatus = "REJECTED"
	CampaignStatusUpdateStatusDormant  CampaignStatusUpdateStatus = "DORMANT"
	CampaignStatusUpdateStatusSuccess  CampaignStatusUpdateStatus = "success"
	CampaignStatusUpdateStatusFailed   CampaignStatusUpdateStatus = "failed"
)

type CampaignStatusUpdateType string

const (
	CampaignStatusUpdateTypeTelnyxEvent             CampaignStatusUpdateType = "TELNYX_EVENT"
	CampaignStatusUpdateTypeRegistration            CampaignStatusUpdateType = "REGISTRATION"
	CampaignStatusUpdateTypeMnoReview               CampaignStatusUpdateType = "MNO_REVIEW"
	CampaignStatusUpdateTypeTelnyxReview            CampaignStatusUpdateType = "TELNYX_REVIEW"
	CampaignStatusUpdateTypeNumberPoolProvisioned   CampaignStatusUpdateType = "NUMBER_POOL_PROVISIONED"
	CampaignStatusUpdateTypeNumberPoolDeprovisioned CampaignStatusUpdateType = "NUMBER_POOL_DEPROVISIONED"
	CampaignStatusUpdateTypeTcrEvent                CampaignStatusUpdateType = "TCR_EVENT"
	CampaignStatusUpdateTypeVerified                CampaignStatusUpdateType = "VERIFIED"
)

type ConferenceCreated struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.created".
	EventType ConferenceCreatedEventType `json:"event_type"`
	Payload   ConferenceCreatedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceCreatedRecordType `json:"record_type"`
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
func (r ConferenceCreated) RawJSON() string { return r.JSON.raw }
func (r *ConferenceCreated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceCreatedEventType string

const (
	ConferenceCreatedEventTypeConferenceCreated ConferenceCreatedEventType = "conference.created"
)

type ConferenceCreatedPayload struct {
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
func (r ConferenceCreatedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceCreatedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceCreatedRecordType string

const (
	ConferenceCreatedRecordTypeEvent ConferenceCreatedRecordType = "event"
)

type ConferenceEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.ended".
	EventType ConferenceEndedEventType `json:"event_type"`
	Payload   ConferenceEndedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceEndedRecordType `json:"record_type"`
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
func (r ConferenceEnded) RawJSON() string { return r.JSON.raw }
func (r *ConferenceEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceEndedEventType string

const (
	ConferenceEndedEventTypeConferenceEnded ConferenceEndedEventType = "conference.ended"
)

type ConferenceEndedPayload struct {
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
func (r ConferenceEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceEndedRecordType string

const (
	ConferenceEndedRecordTypeEvent ConferenceEndedRecordType = "event"
)

type ConferenceFloorChanged struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.floor.changed".
	EventType ConferenceFloorChangedEventType `json:"event_type"`
	Payload   ConferenceFloorChangedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceFloorChangedRecordType `json:"record_type"`
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
func (r ConferenceFloorChanged) RawJSON() string { return r.JSON.raw }
func (r *ConferenceFloorChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceFloorChangedEventType string

const (
	ConferenceFloorChangedEventTypeConferenceFloorChanged ConferenceFloorChangedEventType = "conference.floor.changed"
)

type ConferenceFloorChangedPayload struct {
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
func (r ConferenceFloorChangedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceFloorChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceFloorChangedRecordType string

const (
	ConferenceFloorChangedRecordTypeEvent ConferenceFloorChangedRecordType = "event"
)

type ConferenceParticipantJoined struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.joined".
	EventType ConferenceParticipantJoinedEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                          `json:"occurred_at" format:"date-time"`
	Payload    ConferenceParticipantJoinedPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceParticipantJoinedRecordType `json:"record_type"`
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
func (r ConferenceParticipantJoined) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantJoined) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceParticipantJoinedEventType string

const (
	ConferenceParticipantJoinedEventTypeConferenceParticipantJoined ConferenceParticipantJoinedEventType = "conference.participant.joined"
)

type ConferenceParticipantJoinedPayload struct {
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
func (r ConferenceParticipantJoinedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantJoinedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceParticipantJoinedRecordType string

const (
	ConferenceParticipantJoinedRecordTypeEvent ConferenceParticipantJoinedRecordType = "event"
)

type ConferenceParticipantLeft struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.left".
	EventType ConferenceParticipantLeftEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                        `json:"occurred_at" format:"date-time"`
	Payload    ConferenceParticipantLeftPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceParticipantLeftRecordType `json:"record_type"`
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
func (r ConferenceParticipantLeft) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantLeft) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceParticipantLeftEventType string

const (
	ConferenceParticipantLeftEventTypeConferenceParticipantLeft ConferenceParticipantLeftEventType = "conference.participant.left"
)

type ConferenceParticipantLeftPayload struct {
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
func (r ConferenceParticipantLeftPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantLeftPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceParticipantLeftRecordType string

const (
	ConferenceParticipantLeftRecordTypeEvent ConferenceParticipantLeftRecordType = "event"
)

type ConferenceParticipantPlaybackEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.playback.ended".
	EventType ConferenceParticipantPlaybackEndedEventType `json:"event_type"`
	Payload   ConferenceParticipantPlaybackEndedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceParticipantPlaybackEndedRecordType `json:"record_type"`
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
func (r ConferenceParticipantPlaybackEnded) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceParticipantPlaybackEndedEventType string

const (
	ConferenceParticipantPlaybackEndedEventTypeConferenceParticipantPlaybackEnded ConferenceParticipantPlaybackEndedEventType = "conference.participant.playback.ended"
)

type ConferenceParticipantPlaybackEndedPayload struct {
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
func (r ConferenceParticipantPlaybackEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceParticipantPlaybackEndedRecordType string

const (
	ConferenceParticipantPlaybackEndedRecordTypeEvent ConferenceParticipantPlaybackEndedRecordType = "event"
)

type ConferenceParticipantPlaybackStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.playback.started".
	EventType ConferenceParticipantPlaybackStartedEventType `json:"event_type"`
	Payload   ConferenceParticipantPlaybackStartedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceParticipantPlaybackStartedRecordType `json:"record_type"`
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
func (r ConferenceParticipantPlaybackStarted) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceParticipantPlaybackStartedEventType string

const (
	ConferenceParticipantPlaybackStartedEventTypeConferenceParticipantPlaybackStarted ConferenceParticipantPlaybackStartedEventType = "conference.participant.playback.started"
)

type ConferenceParticipantPlaybackStartedPayload struct {
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
func (r ConferenceParticipantPlaybackStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantPlaybackStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceParticipantPlaybackStartedRecordType string

const (
	ConferenceParticipantPlaybackStartedRecordTypeEvent ConferenceParticipantPlaybackStartedRecordType = "event"
)

type ConferenceParticipantSpeakEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.speak.ended".
	EventType ConferenceParticipantSpeakEndedEventType `json:"event_type"`
	Payload   ConferenceParticipantSpeakEndedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceParticipantSpeakEndedRecordType `json:"record_type"`
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
func (r ConferenceParticipantSpeakEnded) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceParticipantSpeakEndedEventType string

const (
	ConferenceParticipantSpeakEndedEventTypeConferenceParticipantSpeakEnded ConferenceParticipantSpeakEndedEventType = "conference.participant.speak.ended"
)

type ConferenceParticipantSpeakEndedPayload struct {
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
func (r ConferenceParticipantSpeakEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceParticipantSpeakEndedRecordType string

const (
	ConferenceParticipantSpeakEndedRecordTypeEvent ConferenceParticipantSpeakEndedRecordType = "event"
)

type ConferenceParticipantSpeakStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.participant.speak.started".
	EventType ConferenceParticipantSpeakStartedEventType `json:"event_type"`
	Payload   ConferenceParticipantSpeakStartedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceParticipantSpeakStartedRecordType `json:"record_type"`
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
func (r ConferenceParticipantSpeakStarted) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceParticipantSpeakStartedEventType string

const (
	ConferenceParticipantSpeakStartedEventTypeConferenceParticipantSpeakStarted ConferenceParticipantSpeakStartedEventType = "conference.participant.speak.started"
)

type ConferenceParticipantSpeakStartedPayload struct {
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
func (r ConferenceParticipantSpeakStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceParticipantSpeakStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceParticipantSpeakStartedRecordType string

const (
	ConferenceParticipantSpeakStartedRecordTypeEvent ConferenceParticipantSpeakStartedRecordType = "event"
)

type ConferencePlaybackEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.playback.ended".
	EventType ConferencePlaybackEndedEventType `json:"event_type"`
	Payload   ConferencePlaybackEndedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferencePlaybackEndedRecordType `json:"record_type"`
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
func (r ConferencePlaybackEnded) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferencePlaybackEndedEventType string

const (
	ConferencePlaybackEndedEventTypeConferencePlaybackEnded ConferencePlaybackEndedEventType = "conference.playback.ended"
)

type ConferencePlaybackEndedPayload struct {
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
func (r ConferencePlaybackEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferencePlaybackEndedRecordType string

const (
	ConferencePlaybackEndedRecordTypeEvent ConferencePlaybackEndedRecordType = "event"
)

type ConferencePlaybackStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.playback.started".
	EventType ConferencePlaybackStartedEventType `json:"event_type"`
	Payload   ConferencePlaybackStartedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferencePlaybackStartedRecordType `json:"record_type"`
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
func (r ConferencePlaybackStarted) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferencePlaybackStartedEventType string

const (
	ConferencePlaybackStartedEventTypeConferencePlaybackStarted ConferencePlaybackStartedEventType = "conference.playback.started"
)

type ConferencePlaybackStartedPayload struct {
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
func (r ConferencePlaybackStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferencePlaybackStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferencePlaybackStartedRecordType string

const (
	ConferencePlaybackStartedRecordTypeEvent ConferencePlaybackStartedRecordType = "event"
)

type ConferenceRecordingSaved struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.recording.saved".
	EventType ConferenceRecordingSavedEventType `json:"event_type"`
	Payload   ConferenceRecordingSavedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceRecordingSavedRecordType `json:"record_type"`
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
func (r ConferenceRecordingSaved) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSaved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceRecordingSavedEventType string

const (
	ConferenceRecordingSavedEventTypeConferenceRecordingSaved ConferenceRecordingSavedEventType = "conference.recording.saved"
)

type ConferenceRecordingSavedPayload struct {
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
	PublicRecordingURLs ConferenceRecordingSavedPayloadPublicRecordingURLs `json:"public_recording_urls"`
	// ISO 8601 datetime of when recording ended.
	RecordingEndedAt time.Time `json:"recording_ended_at" format:"date-time"`
	// ID of the conference recording.
	RecordingID string `json:"recording_id" format:"uuid"`
	// ISO 8601 datetime of when recording started.
	RecordingStartedAt time.Time `json:"recording_started_at" format:"date-time"`
	// Recording URLs in requested format. These URLs are valid for 10 minutes. After
	// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
	// documentation, or via Mission Control under Reporting -> Recordings.
	RecordingURLs ConferenceRecordingSavedPayloadRecordingURLs `json:"recording_urls"`
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
func (r ConferenceRecordingSavedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSavedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. The URL is valid for as long as the file
// exists. For security purposes, this feature is activated on a per request basis.
// Please contact customer support with your Account ID to request activation.
type ConferenceRecordingSavedPayloadPublicRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3" api:"nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedPayloadPublicRecordingURLs) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSavedPayloadPublicRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recording URLs in requested format. These URLs are valid for 10 minutes. After
// 10 minutes, you may retrieve recordings via API using Reports -> Call Recordings
// documentation, or via Mission Control under Reporting -> Recordings.
type ConferenceRecordingSavedPayloadRecordingURLs struct {
	// Recording URL in requested `mp3` format.
	MP3 string `json:"mp3" api:"nullable"`
	// Recording URL in requested `wav` format.
	Wav string `json:"wav" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MP3         respjson.Field
		Wav         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConferenceRecordingSavedPayloadRecordingURLs) RawJSON() string { return r.JSON.raw }
func (r *ConferenceRecordingSavedPayloadRecordingURLs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceRecordingSavedRecordType string

const (
	ConferenceRecordingSavedRecordTypeEvent ConferenceRecordingSavedRecordType = "event"
)

type ConferenceSpeakEnded struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.speak.ended".
	EventType ConferenceSpeakEndedEventType `json:"event_type"`
	Payload   ConferenceSpeakEndedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceSpeakEndedRecordType `json:"record_type"`
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
func (r ConferenceSpeakEnded) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakEnded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceSpeakEndedEventType string

const (
	ConferenceSpeakEndedEventTypeConferenceSpeakEnded ConferenceSpeakEndedEventType = "conference.speak.ended"
)

type ConferenceSpeakEndedPayload struct {
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
func (r ConferenceSpeakEndedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakEndedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceSpeakEndedRecordType string

const (
	ConferenceSpeakEndedRecordTypeEvent ConferenceSpeakEndedRecordType = "event"
)

type ConferenceSpeakStarted struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "conference.speak.started".
	EventType ConferenceSpeakStartedEventType `json:"event_type"`
	Payload   ConferenceSpeakStartedPayload   `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType ConferenceSpeakStartedRecordType `json:"record_type"`
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
func (r ConferenceSpeakStarted) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type ConferenceSpeakStartedEventType string

const (
	ConferenceSpeakStartedEventTypeConferenceSpeakStarted ConferenceSpeakStartedEventType = "conference.speak.started"
)

type ConferenceSpeakStartedPayload struct {
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
func (r ConferenceSpeakStartedPayload) RawJSON() string { return r.JSON.raw }
func (r *ConferenceSpeakStartedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type ConferenceSpeakStartedRecordType string

const (
	ConferenceSpeakStartedRecordTypeEvent ConferenceSpeakStartedRecordType = "event"
)

type FaxDelivered struct {
	Data FaxDeliveredData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxDeliveredMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxDelivered) RawJSON() string { return r.JSON.raw }
func (r *FaxDelivered) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxDeliveredData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.delivered".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time               `json:"occurred_at" format:"date-time"`
	Payload    FaxDeliveredDataPayload `json:"payload"`
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
func (r FaxDeliveredData) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxDeliveredDataPayload struct {
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
func (r FaxDeliveredDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxDeliveredMeta struct {
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
func (r FaxDeliveredMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxDeliveredMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxFailed struct {
	Data FaxFailedData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxFailedMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxFailed) RawJSON() string { return r.JSON.raw }
func (r *FaxFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxFailedData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.failed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time            `json:"occurred_at" format:"date-time"`
	Payload    FaxFailedDataPayload `json:"payload"`
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
func (r FaxFailedData) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxFailedDataPayload struct {
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
func (r FaxFailedDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxFailedMeta struct {
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
func (r FaxFailedMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxFailedMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxMediaProcessed struct {
	Data FaxMediaProcessedData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxMediaProcessedMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxMediaProcessed) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxMediaProcessedData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.media.processed".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                    `json:"occurred_at" format:"date-time"`
	Payload    FaxMediaProcessedDataPayload `json:"payload"`
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
func (r FaxMediaProcessedData) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxMediaProcessedDataPayload struct {
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
func (r FaxMediaProcessedDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxMediaProcessedMeta struct {
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
func (r FaxMediaProcessedMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxMediaProcessedMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxQueued struct {
	Data FaxQueuedData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxQueuedMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxQueued) RawJSON() string { return r.JSON.raw }
func (r *FaxQueued) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxQueuedData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.queued".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time            `json:"occurred_at" format:"date-time"`
	Payload    FaxQueuedDataPayload `json:"payload"`
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
func (r FaxQueuedData) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxQueuedDataPayload struct {
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
func (r FaxQueuedDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxQueuedMeta struct {
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
func (r FaxQueuedMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxQueuedMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxSendingStarted struct {
	Data FaxSendingStartedData `json:"data"`
	// Metadata about the webhook delivery.
	Meta FaxSendingStartedMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FaxSendingStarted) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxSendingStartedData struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "fax.sending.started".
	EventType string `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time                    `json:"occurred_at" format:"date-time"`
	Payload    FaxSendingStartedDataPayload `json:"payload"`
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
func (r FaxSendingStartedData) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FaxSendingStartedDataPayload struct {
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
func (r FaxSendingStartedDataPayload) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedDataPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the webhook delivery.
type FaxSendingStartedMeta struct {
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
func (r FaxSendingStartedMeta) RawJSON() string { return r.JSON.raw }
func (r *FaxSendingStartedMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboundMessage struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "message.received".
	EventType InboundMessageEventType `json:"event_type"`
	// ISO 8601 formatted date indicating when the resource was created.
	OccurredAt time.Time                    `json:"occurred_at" format:"date-time"`
	Payload    shared.InboundMessagePayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType InboundMessageRecordType `json:"record_type"`
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
func (r InboundMessage) RawJSON() string { return r.JSON.raw }
func (r *InboundMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type InboundMessageEventType string

const (
	InboundMessageEventTypeMessageReceived InboundMessageEventType = "message.received"
)

// Identifies the type of the resource.
type InboundMessageRecordType string

const (
	InboundMessageRecordTypeEvent InboundMessageRecordType = "event"
)

type NumberOrderStatusUpdate struct {
	Data NumberOrderStatusUpdateData `json:"data" api:"required"`
	Meta NumberOrderStatusUpdateMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderStatusUpdate) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderStatusUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderStatusUpdateData struct {
	// Unique identifier for the event
	ID string `json:"id" api:"required" format:"uuid"`
	// The type of event being sent
	EventType string `json:"event_type" api:"required"`
	// ISO 8601 timestamp of when the event occurred
	OccurredAt time.Time                   `json:"occurred_at" api:"required" format:"date-time"`
	Payload    NumberOrderWithPhoneNumbers `json:"payload" api:"required"`
	// Type of record
	RecordType string `json:"record_type" api:"required"`
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
func (r NumberOrderStatusUpdateData) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderStatusUpdateData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberOrderStatusUpdateMeta struct {
	// Webhook delivery attempt number
	Attempt int64 `json:"attempt" api:"required"`
	// URL where the webhook was delivered
	DeliveredTo string `json:"delivered_to" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attempt     respjson.Field
		DeliveredTo respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberOrderStatusUpdateMeta) RawJSON() string { return r.JSON.raw }
func (r *NumberOrderStatusUpdateMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessage struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "message.sent", "message.finalized".
	EventType OutboundMessageEventType `json:"event_type"`
	// ISO 8601 formatted date indicating when the resource was created.
	OccurredAt time.Time              `json:"occurred_at" format:"date-time"`
	Payload    OutboundMessagePayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType OutboundMessageRecordType `json:"record_type"`
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
func (r OutboundMessage) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type OutboundMessageEventType string

const (
	OutboundMessageEventTypeMessageSent      OutboundMessageEventType = "message.sent"
	OutboundMessageEventTypeMessageFinalized OutboundMessageEventType = "message.finalized"
)

// Identifies the type of the resource.
type OutboundMessageRecordType string

const (
	OutboundMessageRecordTypeEvent OutboundMessageRecordType = "event"
)

type ReplacedLinkClick struct {
	// The message ID associated with the clicked link.
	MessageID string `json:"message_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the message request was received.
	TimeClicked time.Time `json:"time_clicked" format:"date-time"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	To string `json:"to"`
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
func (r ReplacedLinkClick) RawJSON() string { return r.JSON.raw }
func (r *ReplacedLinkClick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Transcription struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// The type of event being delivered.
	//
	// Any of "call.transcription".
	EventType TranscriptionEventType `json:"event_type"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt time.Time            `json:"occurred_at" format:"date-time"`
	Payload    TranscriptionPayload `json:"payload"`
	// Identifies the type of the resource.
	//
	// Any of "event".
	RecordType TranscriptionRecordType `json:"record_type"`
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
func (r Transcription) RawJSON() string { return r.JSON.raw }
func (r *Transcription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of event being delivered.
type TranscriptionEventType string

const (
	TranscriptionEventTypeCallTranscription TranscriptionEventType = "call.transcription"
)

type TranscriptionPayload struct {
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
	ConnectionID      string                                `json:"connection_id"`
	TranscriptionData TranscriptionPayloadTranscriptionData `json:"transcription_data"`
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
func (r TranscriptionPayload) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionPayloadTranscriptionData struct {
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
func (r TranscriptionPayloadTranscriptionData) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionPayloadTranscriptionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type TranscriptionRecordType string

const (
	TranscriptionRecordTypeEvent TranscriptionRecordType = "event"
)

type CallAIGatherEndedWebhookEvent struct {
	Data CallAIGatherEnded `json:"data"`
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

type CallAIGatherMessageHistoryUpdatedWebhookEvent struct {
	Data CallAIGatherMessageHistoryUpdated `json:"data"`
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

type CallAIGatherPartialResultsWebhookEvent struct {
	Data CallAIGatherPartialResults `json:"data"`
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

type CallAnsweredWebhookEvent struct {
	Data CallAnswered `json:"data"`
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

type CallBridgedWebhookEvent struct {
	Data CallBridged `json:"data"`
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

type CallConversationEndedWebhookEvent struct {
	Data CallConversationEnded `json:"data"`
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

type CallConversationInsightsGeneratedWebhookEvent struct {
	Data CallConversationInsightsGenerated `json:"data"`
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

type CallDtmfReceivedWebhookEvent struct {
	Data CallDtmfReceived `json:"data"`
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

type CallEnqueuedWebhookEvent struct {
	Data CallEnqueued `json:"data"`
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

type CallForkStartedWebhookEvent struct {
	Data CallForkStarted `json:"data"`
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

type CallForkStoppedWebhookEvent struct {
	Data CallForkStopped `json:"data"`
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

type CallGatherEndedWebhookEvent struct {
	Data CallGatherEnded `json:"data"`
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

type CallHangupWebhookEvent struct {
	Data CallHangup `json:"data"`
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

type CallInitiatedWebhookEvent struct {
	Data CallInitiated `json:"data"`
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

type CallLeftQueueWebhookEvent struct {
	Data CallLeftQueue `json:"data"`
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

type CallMachineDetectionEndedWebhookEvent struct {
	Data CallMachineDetectionEnded `json:"data"`
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

type CallMachineGreetingEndedWebhookEvent struct {
	Data CallMachineGreetingEnded `json:"data"`
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

type CallMachinePremiumDetectionEndedWebhookEvent struct {
	Data CallMachinePremiumDetectionEnded `json:"data"`
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

type CallMachinePremiumGreetingEndedWebhookEvent struct {
	Data CallMachinePremiumGreetingEnded `json:"data"`
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

type CallPlaybackEndedWebhookEvent struct {
	Data CallPlaybackEnded `json:"data"`
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

type CallPlaybackStartedWebhookEvent struct {
	Data CallPlaybackStarted `json:"data"`
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

type CallRecordingErrorWebhookEvent struct {
	Data CallRecordingError `json:"data"`
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

type CallRecordingSavedWebhookEvent struct {
	Data CallRecordingSaved `json:"data"`
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

type CallRecordingTranscriptionSavedWebhookEvent struct {
	Data CallRecordingTranscriptionSaved `json:"data"`
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

type CallReferCompletedWebhookEvent struct {
	Data CallReferCompleted `json:"data"`
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

type CallReferFailedWebhookEvent struct {
	Data CallReferFailed `json:"data"`
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

type CallReferStartedWebhookEvent struct {
	Data CallReferStarted `json:"data"`
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

type CallSiprecFailedWebhookEvent struct {
	Data CallSiprecFailed `json:"data"`
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

type CallSiprecStartedWebhookEvent struct {
	Data CallSiprecStarted `json:"data"`
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

type CallSiprecStoppedWebhookEvent struct {
	Data CallSiprecStopped `json:"data"`
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

type CallSpeakEndedWebhookEvent struct {
	Data CallSpeakEnded `json:"data"`
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

type CallSpeakStartedWebhookEvent struct {
	Data CallSpeakStarted `json:"data"`
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

type ConferenceCreatedWebhookEvent struct {
	Data ConferenceCreated `json:"data"`
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

type ConferenceEndedWebhookEvent struct {
	Data ConferenceEnded `json:"data"`
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

type ConferenceParticipantJoinedWebhookEvent struct {
	Data ConferenceParticipantJoined `json:"data"`
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

type ConferenceParticipantLeftWebhookEvent struct {
	Data ConferenceParticipantLeft `json:"data"`
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

type ConferenceParticipantPlaybackEndedWebhookEvent struct {
	Data ConferenceParticipantPlaybackEnded `json:"data"`
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

type ConferenceParticipantPlaybackStartedWebhookEvent struct {
	Data ConferenceParticipantPlaybackStarted `json:"data"`
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

type ConferenceParticipantSpeakEndedWebhookEvent struct {
	Data ConferenceParticipantSpeakEnded `json:"data"`
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

type ConferenceParticipantSpeakStartedWebhookEvent struct {
	Data ConferenceParticipantSpeakStarted `json:"data"`
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

type ConferencePlaybackEndedWebhookEvent struct {
	Data ConferencePlaybackEnded `json:"data"`
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

type ConferencePlaybackStartedWebhookEvent struct {
	Data ConferencePlaybackStarted `json:"data"`
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

type ConferenceRecordingSavedWebhookEvent struct {
	Data ConferenceRecordingSaved `json:"data"`
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

type ConferenceSpeakEndedWebhookEvent struct {
	Data ConferenceSpeakEnded `json:"data"`
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

type ConferenceSpeakStartedWebhookEvent struct {
	Data ConferenceSpeakStarted `json:"data"`
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

type DeliveryUpdateWebhookEvent struct {
	Data OutboundMessage                `json:"data"`
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

type InboundMessageWebhookEvent struct {
	Data InboundMessage `json:"data"`
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

type ReplacedLinkClickWebhookEvent struct {
	Data ReplacedLinkClick `json:"data"`
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

type TranscriptionWebhookEvent struct {
	Data Transcription `json:"data"`
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
// [CallStreamingStoppedWebhookEvent], [CampaignStatusUpdate],
// [ConferenceCreatedWebhookEvent], [ConferenceEndedWebhookEvent],
// [ConferenceFloorChanged], [ConferenceParticipantJoinedWebhookEvent],
// [ConferenceParticipantLeftWebhookEvent],
// [ConferenceParticipantPlaybackEndedWebhookEvent],
// [ConferenceParticipantPlaybackStartedWebhookEvent],
// [ConferenceParticipantSpeakEndedWebhookEvent],
// [ConferenceParticipantSpeakStartedWebhookEvent],
// [ConferencePlaybackEndedWebhookEvent], [ConferencePlaybackStartedWebhookEvent],
// [ConferenceRecordingSavedWebhookEvent], [ConferenceSpeakEndedWebhookEvent],
// [ConferenceSpeakStartedWebhookEvent], [DeliveryUpdateWebhookEvent],
// [FaxDelivered], [FaxFailed], [FaxMediaProcessed], [FaxQueued],
// [FaxSendingStarted], [InboundMessageWebhookEvent], [NumberOrderStatusUpdate],
// [ReplacedLinkClickWebhookEvent], [TranscriptionWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	// This field is a union of [CallAIGatherEnded],
	// [CallAIGatherMessageHistoryUpdated], [CallAIGatherPartialResults],
	// [CallAnswered], [CallBridged], [CallConversationEnded],
	// [CallConversationInsightsGenerated], [CallDtmfReceived], [CallEnqueued],
	// [CallForkStarted], [CallForkStopped], [CallGatherEnded], [CallHangup],
	// [CallInitiated], [CallLeftQueue], [CallMachineDetectionEnded],
	// [CallMachineGreetingEnded], [CallMachinePremiumDetectionEnded],
	// [CallMachinePremiumGreetingEnded], [CallPlaybackEnded], [CallPlaybackStarted],
	// [CallRecordingError], [CallRecordingSaved], [CallRecordingTranscriptionSaved],
	// [CallReferCompleted], [CallReferFailed], [CallReferStarted], [CallSiprecFailed],
	// [CallSiprecStarted], [CallSiprecStopped], [CallSpeakEnded], [CallSpeakStarted],
	// [CallStreamingFailed], [CallStreamingStarted], [CallStreamingStopped],
	// [ConferenceCreated], [ConferenceEnded], [ConferenceParticipantJoined],
	// [ConferenceParticipantLeft], [ConferenceParticipantPlaybackEnded],
	// [ConferenceParticipantPlaybackStarted], [ConferenceParticipantSpeakEnded],
	// [ConferenceParticipantSpeakStarted], [ConferencePlaybackEnded],
	// [ConferencePlaybackStarted], [ConferenceRecordingSaved], [ConferenceSpeakEnded],
	// [ConferenceSpeakStarted], [OutboundMessage], [FaxDeliveredData],
	// [FaxFailedData], [FaxMediaProcessedData], [FaxQueuedData],
	// [FaxSendingStartedData], [InboundMessage], [NumberOrderStatusUpdateData],
	// [ReplacedLinkClick], [Transcription]
	Data UnsafeUnwrapWebhookEventUnionData `json:"data"`
	// This field is from variant [CampaignStatusUpdate].
	BrandID string `json:"brandId"`
	// This field is from variant [CampaignStatusUpdate].
	CampaignID string `json:"campaignId"`
	// This field is from variant [CampaignStatusUpdate].
	CreateDate string `json:"createDate"`
	// This field is from variant [CampaignStatusUpdate].
	CspID string `json:"cspId"`
	// This field is from variant [CampaignStatusUpdate].
	Description string `json:"description"`
	// This field is from variant [CampaignStatusUpdate].
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// This field is from variant [CampaignStatusUpdate].
	Status CampaignStatusUpdateStatus `json:"status"`
	// This field is from variant [CampaignStatusUpdate].
	Type CampaignStatusUpdateType `json:"type"`
	// This field is from variant [ConferenceFloorChanged].
	ID string `json:"id"`
	// This field is from variant [ConferenceFloorChanged].
	EventType ConferenceFloorChangedEventType `json:"event_type"`
	// This field is from variant [ConferenceFloorChanged].
	Payload ConferenceFloorChangedPayload `json:"payload"`
	// This field is from variant [ConferenceFloorChanged].
	RecordType ConferenceFloorChangedRecordType `json:"record_type"`
	// This field is a union of [DeliveryUpdateWebhookEventMeta], [FaxDeliveredMeta],
	// [FaxFailedMeta], [FaxMediaProcessedMeta], [FaxQueuedMeta],
	// [FaxSendingStartedMeta], [NumberOrderStatusUpdateMeta]
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

func (u UnsafeUnwrapWebhookEventUnion) AsCampaignStatusUpdateEvent() (v CampaignStatusUpdate) {
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

func (u UnsafeUnwrapWebhookEventUnion) AsConferenceFloorChanged() (v ConferenceFloorChanged) {
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

func (u UnsafeUnwrapWebhookEventUnion) AsFaxDelivered() (v FaxDelivered) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxFailed() (v FaxFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxMediaProcessed() (v FaxMediaProcessed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxQueued() (v FaxQueued) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsFaxSendingStarted() (v FaxSendingStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsInboundMessageWebhookEvent() (v InboundMessageWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsNumberOrderEvent() (v NumberOrderStatusUpdate) {
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
	// This field is a union of [CallAIGatherEndedPayload],
	// [CallAIGatherMessageHistoryUpdatedPayload], [CallAIGatherPartialResultsPayload],
	// [CallAnsweredPayload], [CallBridgedPayload], [CallConversationEndedPayload],
	// [CallConversationInsightsGeneratedPayload], [CallDtmfReceivedPayload],
	// [CallEnqueuedPayload], [CallForkStartedPayload], [CallForkStoppedPayload],
	// [CallGatherEndedPayload], [CallHangupPayload], [CallInitiatedPayload],
	// [CallLeftQueuePayload], [CallMachineDetectionEndedPayload],
	// [CallMachineGreetingEndedPayload], [CallMachinePremiumDetectionEndedPayload],
	// [CallMachinePremiumGreetingEndedPayload], [CallPlaybackEndedPayload],
	// [CallPlaybackStartedPayload], [CallRecordingErrorPayload],
	// [CallRecordingSavedPayload], [CallRecordingTranscriptionSavedPayload],
	// [CallReferCompletedPayload], [CallReferFailedPayload],
	// [CallReferStartedPayload], [CallSiprecFailedPayload],
	// [CallSiprecStartedPayload], [CallSiprecStoppedPayload], [CallSpeakEndedPayload],
	// [CallSpeakStartedPayload], [CallStreamingFailedPayload],
	// [CallStreamingStartedPayload], [CallStreamingStoppedPayload],
	// [ConferenceCreatedPayload], [ConferenceEndedPayload],
	// [ConferenceParticipantJoinedPayload], [ConferenceParticipantLeftPayload],
	// [ConferenceParticipantPlaybackEndedPayload],
	// [ConferenceParticipantPlaybackStartedPayload],
	// [ConferenceParticipantSpeakEndedPayload],
	// [ConferenceParticipantSpeakStartedPayload], [ConferencePlaybackEndedPayload],
	// [ConferencePlaybackStartedPayload], [ConferenceRecordingSavedPayload],
	// [ConferenceSpeakEndedPayload], [ConferenceSpeakStartedPayload],
	// [OutboundMessagePayload], [FaxDeliveredDataPayload], [FaxFailedDataPayload],
	// [FaxMediaProcessedDataPayload], [FaxQueuedDataPayload],
	// [FaxSendingStartedDataPayload], [shared.InboundMessagePayload],
	// [NumberOrderWithPhoneNumbers], [TranscriptionPayload]
	Payload    UnsafeUnwrapWebhookEventUnionDataPayload `json:"payload"`
	RecordType string                                   `json:"record_type"`
	// This field is from variant [CallConversationEnded].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [ReplacedLinkClick].
	MessageID string `json:"message_id"`
	// This field is from variant [ReplacedLinkClick].
	TimeClicked time.Time `json:"time_clicked"`
	// This field is from variant [ReplacedLinkClick].
	To string `json:"to"`
	// This field is from variant [ReplacedLinkClick].
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
	// This field is a union of [[]CallAIGatherEndedPayloadMessageHistory],
	// [[]CallAIGatherMessageHistoryUpdatedPayloadMessageHistory],
	// [[]CallAIGatherPartialResultsPayloadMessageHistory]
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
	// This field is from variant [CallAIGatherPartialResultsPayload].
	PartialResults map[string]any    `json:"partial_results"`
	CustomHeaders  []CustomSipHeader `json:"custom_headers"`
	SipHeaders     []SipHeader       `json:"sip_headers"`
	StartTime      time.Time         `json:"start_time"`
	State          string            `json:"state"`
	Tags           []string          `json:"tags"`
	// This field is from variant [CallConversationEndedPayload].
	AssistantID      string `json:"assistant_id"`
	CallingPartyType string `json:"calling_party_type"`
	// This field is from variant [CallConversationEndedPayload].
	ConversationID string `json:"conversation_id"`
	// This field is from variant [CallConversationEndedPayload].
	DurationSec int64 `json:"duration_sec"`
	// This field is from variant [CallConversationEndedPayload].
	LlmModel string `json:"llm_model"`
	// This field is from variant [CallConversationEndedPayload].
	SttModel string `json:"stt_model"`
	// This field is from variant [CallConversationEndedPayload].
	TtsModelID string `json:"tts_model_id"`
	// This field is from variant [CallConversationEndedPayload].
	TtsProvider string `json:"tts_provider"`
	// This field is from variant [CallConversationEndedPayload].
	TtsVoiceID string `json:"tts_voice_id"`
	// This field is from variant [CallConversationInsightsGeneratedPayload].
	InsightGroupID string `json:"insight_group_id"`
	// This field is from variant [CallConversationInsightsGeneratedPayload].
	Results []CallConversationInsightsGeneratedPayloadResult `json:"results"`
	// This field is from variant [CallDtmfReceivedPayload].
	Digit string `json:"digit"`
	// This field is from variant [CallEnqueuedPayload].
	CurrentPosition int64  `json:"current_position"`
	Queue           string `json:"queue"`
	// This field is from variant [CallEnqueuedPayload].
	QueueAvgWaitTimeSecs int64  `json:"queue_avg_wait_time_secs"`
	StreamType           string `json:"stream_type"`
	// This field is from variant [CallGatherEndedPayload].
	Digits string `json:"digits"`
	// This field is from variant [CallHangupPayload].
	CallQualityStats CallHangupPayloadCallQualityStats `json:"call_quality_stats"`
	HangupCause      string                            `json:"hangup_cause"`
	// This field is from variant [CallHangupPayload].
	HangupSource string `json:"hangup_source"`
	// This field is from variant [CallHangupPayload].
	SipHangupCause string `json:"sip_hangup_cause"`
	// This field is from variant [CallInitiatedPayload].
	CallScreeningResult string `json:"call_screening_result"`
	// This field is from variant [CallInitiatedPayload].
	CallerIDName string `json:"caller_id_name"`
	// This field is from variant [CallInitiatedPayload].
	ConnectionCodecs string `json:"connection_codecs"`
	Direction        string `json:"direction"`
	// This field is from variant [CallInitiatedPayload].
	OfferedCodecs string `json:"offered_codecs"`
	// This field is from variant [CallInitiatedPayload].
	ShakenStirAttestation string `json:"shaken_stir_attestation"`
	// This field is from variant [CallInitiatedPayload].
	ShakenStirValidated bool `json:"shaken_stir_validated"`
	// This field is from variant [CallLeftQueuePayload].
	QueuePosition int64  `json:"queue_position"`
	Reason        string `json:"reason"`
	// This field is from variant [CallLeftQueuePayload].
	WaitTimeSecs int64  `json:"wait_time_secs"`
	MediaName    string `json:"media_name"`
	MediaURL     string `json:"media_url"`
	Overlay      bool   `json:"overlay"`
	// This field is from variant [CallPlaybackEndedPayload].
	StatusDetail string `json:"status_detail"`
	Channels     string `json:"channels"`
	// This field is a union of [CallRecordingSavedPayloadPublicRecordingURLs],
	// [ConferenceRecordingSavedPayloadPublicRecordingURLs]
	PublicRecordingURLs UnsafeUnwrapWebhookEventUnionDataPayloadPublicRecordingURLs `json:"public_recording_urls"`
	RecordingEndedAt    time.Time                                                   `json:"recording_ended_at"`
	RecordingStartedAt  time.Time                                                   `json:"recording_started_at"`
	// This field is a union of [CallRecordingSavedPayloadRecordingURLs],
	// [ConferenceRecordingSavedPayloadRecordingURLs]
	RecordingURLs UnsafeUnwrapWebhookEventUnionDataPayloadRecordingURLs `json:"recording_urls"`
	RecordingID   string                                                `json:"recording_id"`
	// This field is from variant [CallRecordingTranscriptionSavedPayload].
	RecordingTranscriptionID string `json:"recording_transcription_id"`
	// This field is from variant [CallRecordingTranscriptionSavedPayload].
	TranscriptionText string `json:"transcription_text"`
	SipNotifyResponse int64  `json:"sip_notify_response"`
	// This field is from variant [CallSiprecFailedPayload].
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
	// This field is from variant [ConferenceRecordingSavedPayload].
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
	Errors        []shared.MessagingError                               `json:"errors"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media              UnsafeUnwrapWebhookEventUnionDataPayloadMedia `json:"media"`
	MessagingProfileID string                                        `json:"messaging_profile_id"`
	OrganizationID     string                                        `json:"organization_id"`
	Parts              int64                                         `json:"parts"`
	ReceivedAt         time.Time                                     `json:"received_at"`
	RecordType         string                                        `json:"record_type"`
	SentAt             time.Time                                     `json:"sent_at"`
	// This field is from variant [OutboundMessagePayload].
	SmartEncodingApplied  bool      `json:"smart_encoding_applied"`
	Subject               string    `json:"subject"`
	TcrCampaignBillable   bool      `json:"tcr_campaign_billable"`
	TcrCampaignID         string    `json:"tcr_campaign_id"`
	TcrCampaignRegistered string    `json:"tcr_campaign_registered"`
	Text                  string    `json:"text"`
	Type                  string    `json:"type"`
	ValidUntil            time.Time `json:"valid_until"`
	WebhookFailoverURL    string    `json:"webhook_failover_url"`
	WebhookURL            string    `json:"webhook_url"`
	// This field is from variant [FaxDeliveredDataPayload].
	CallDurationSecs int64  `json:"call_duration_secs"`
	FaxID            string `json:"fax_id"`
	OriginalMediaURL string `json:"original_media_url"`
	// This field is from variant [FaxDeliveredDataPayload].
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
	// This field is from variant [TranscriptionPayload].
	TranscriptionData TranscriptionPayloadTranscriptionData `json:"transcription_data"`
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
		SmartEncodingApplied     respjson.Field
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
// will be valid: OfCallAIGatherEndedPayloadMessageHistoryArray
// OfCallAIGatherMessageHistoryUpdatedPayloadMessageHistoryArray
// OfCallAIGatherPartialResultsPayloadMessageHistoryArray]
type UnsafeUnwrapWebhookEventUnionDataPayloadMessageHistory struct {
	// This field will be present if the value is a
	// [[]CallAIGatherEndedPayloadMessageHistory] instead of an object.
	OfCallAIGatherEndedPayloadMessageHistoryArray []CallAIGatherEndedPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherMessageHistoryUpdatedPayloadMessageHistory] instead of an object.
	OfCallAIGatherMessageHistoryUpdatedPayloadMessageHistoryArray []CallAIGatherMessageHistoryUpdatedPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherPartialResultsPayloadMessageHistory] instead of an object.
	OfCallAIGatherPartialResultsPayloadMessageHistoryArray []CallAIGatherPartialResultsPayloadMessageHistory `json:",inline"`
	JSON                                                   struct {
		OfCallAIGatherEndedPayloadMessageHistoryArray                 respjson.Field
		OfCallAIGatherMessageHistoryUpdatedPayloadMessageHistoryArray respjson.Field
		OfCallAIGatherPartialResultsPayloadMessageHistoryArray        respjson.Field
		raw                                                           string
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
// will be valid: OfCallAIGatherEndedPayloadResult
// OfCallMachinePremiumGreetingEndedPayloadResult]
type UnsafeUnwrapWebhookEventUnionDataPayloadResult struct {
	// This field will be present if the value is a [any] instead of an object.
	OfCallAIGatherEndedPayloadResult any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfCallMachinePremiumGreetingEndedPayloadResult string `json:",inline"`
	JSON                                           struct {
		OfCallAIGatherEndedPayloadResult               respjson.Field
		OfCallMachinePremiumGreetingEndedPayloadResult respjson.Field
		raw                                            string
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
// [CallStreamingStoppedWebhookEvent], [CampaignStatusUpdate],
// [ConferenceCreatedWebhookEvent], [ConferenceEndedWebhookEvent],
// [ConferenceFloorChanged], [ConferenceParticipantJoinedWebhookEvent],
// [ConferenceParticipantLeftWebhookEvent],
// [ConferenceParticipantPlaybackEndedWebhookEvent],
// [ConferenceParticipantPlaybackStartedWebhookEvent],
// [ConferenceParticipantSpeakEndedWebhookEvent],
// [ConferenceParticipantSpeakStartedWebhookEvent],
// [ConferencePlaybackEndedWebhookEvent], [ConferencePlaybackStartedWebhookEvent],
// [ConferenceRecordingSavedWebhookEvent], [ConferenceSpeakEndedWebhookEvent],
// [ConferenceSpeakStartedWebhookEvent], [DeliveryUpdateWebhookEvent],
// [FaxDelivered], [FaxFailed], [FaxMediaProcessed], [FaxQueued],
// [FaxSendingStarted], [InboundMessageWebhookEvent], [NumberOrderStatusUpdate],
// [ReplacedLinkClickWebhookEvent], [TranscriptionWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	// This field is a union of [CallAIGatherEnded],
	// [CallAIGatherMessageHistoryUpdated], [CallAIGatherPartialResults],
	// [CallAnswered], [CallBridged], [CallConversationEnded],
	// [CallConversationInsightsGenerated], [CallDtmfReceived], [CallEnqueued],
	// [CallForkStarted], [CallForkStopped], [CallGatherEnded], [CallHangup],
	// [CallInitiated], [CallLeftQueue], [CallMachineDetectionEnded],
	// [CallMachineGreetingEnded], [CallMachinePremiumDetectionEnded],
	// [CallMachinePremiumGreetingEnded], [CallPlaybackEnded], [CallPlaybackStarted],
	// [CallRecordingError], [CallRecordingSaved], [CallRecordingTranscriptionSaved],
	// [CallReferCompleted], [CallReferFailed], [CallReferStarted], [CallSiprecFailed],
	// [CallSiprecStarted], [CallSiprecStopped], [CallSpeakEnded], [CallSpeakStarted],
	// [CallStreamingFailed], [CallStreamingStarted], [CallStreamingStopped],
	// [ConferenceCreated], [ConferenceEnded], [ConferenceParticipantJoined],
	// [ConferenceParticipantLeft], [ConferenceParticipantPlaybackEnded],
	// [ConferenceParticipantPlaybackStarted], [ConferenceParticipantSpeakEnded],
	// [ConferenceParticipantSpeakStarted], [ConferencePlaybackEnded],
	// [ConferencePlaybackStarted], [ConferenceRecordingSaved], [ConferenceSpeakEnded],
	// [ConferenceSpeakStarted], [OutboundMessage], [FaxDeliveredData],
	// [FaxFailedData], [FaxMediaProcessedData], [FaxQueuedData],
	// [FaxSendingStartedData], [InboundMessage], [NumberOrderStatusUpdateData],
	// [ReplacedLinkClick], [Transcription]
	Data UnwrapWebhookEventUnionData `json:"data"`
	// This field is from variant [CampaignStatusUpdate].
	BrandID string `json:"brandId"`
	// This field is from variant [CampaignStatusUpdate].
	CampaignID string `json:"campaignId"`
	// This field is from variant [CampaignStatusUpdate].
	CreateDate string `json:"createDate"`
	// This field is from variant [CampaignStatusUpdate].
	CspID string `json:"cspId"`
	// This field is from variant [CampaignStatusUpdate].
	Description string `json:"description"`
	// This field is from variant [CampaignStatusUpdate].
	IsTMobileRegistered bool `json:"isTMobileRegistered"`
	// This field is from variant [CampaignStatusUpdate].
	Status CampaignStatusUpdateStatus `json:"status"`
	// This field is from variant [CampaignStatusUpdate].
	Type CampaignStatusUpdateType `json:"type"`
	// This field is from variant [ConferenceFloorChanged].
	ID string `json:"id"`
	// This field is from variant [ConferenceFloorChanged].
	EventType ConferenceFloorChangedEventType `json:"event_type"`
	// This field is from variant [ConferenceFloorChanged].
	Payload ConferenceFloorChangedPayload `json:"payload"`
	// This field is from variant [ConferenceFloorChanged].
	RecordType ConferenceFloorChangedRecordType `json:"record_type"`
	// This field is a union of [DeliveryUpdateWebhookEventMeta], [FaxDeliveredMeta],
	// [FaxFailedMeta], [FaxMediaProcessedMeta], [FaxQueuedMeta],
	// [FaxSendingStartedMeta], [NumberOrderStatusUpdateMeta]
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

func (u UnwrapWebhookEventUnion) AsCampaignStatusUpdateEvent() (v CampaignStatusUpdate) {
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

func (u UnwrapWebhookEventUnion) AsConferenceFloorChanged() (v ConferenceFloorChanged) {
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

func (u UnwrapWebhookEventUnion) AsFaxDelivered() (v FaxDelivered) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxFailed() (v FaxFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxMediaProcessed() (v FaxMediaProcessed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxQueued() (v FaxQueued) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsFaxSendingStarted() (v FaxSendingStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsInboundMessageWebhookEvent() (v InboundMessageWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsNumberOrderEvent() (v NumberOrderStatusUpdate) {
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
	// This field is a union of [CallAIGatherEndedPayload],
	// [CallAIGatherMessageHistoryUpdatedPayload], [CallAIGatherPartialResultsPayload],
	// [CallAnsweredPayload], [CallBridgedPayload], [CallConversationEndedPayload],
	// [CallConversationInsightsGeneratedPayload], [CallDtmfReceivedPayload],
	// [CallEnqueuedPayload], [CallForkStartedPayload], [CallForkStoppedPayload],
	// [CallGatherEndedPayload], [CallHangupPayload], [CallInitiatedPayload],
	// [CallLeftQueuePayload], [CallMachineDetectionEndedPayload],
	// [CallMachineGreetingEndedPayload], [CallMachinePremiumDetectionEndedPayload],
	// [CallMachinePremiumGreetingEndedPayload], [CallPlaybackEndedPayload],
	// [CallPlaybackStartedPayload], [CallRecordingErrorPayload],
	// [CallRecordingSavedPayload], [CallRecordingTranscriptionSavedPayload],
	// [CallReferCompletedPayload], [CallReferFailedPayload],
	// [CallReferStartedPayload], [CallSiprecFailedPayload],
	// [CallSiprecStartedPayload], [CallSiprecStoppedPayload], [CallSpeakEndedPayload],
	// [CallSpeakStartedPayload], [CallStreamingFailedPayload],
	// [CallStreamingStartedPayload], [CallStreamingStoppedPayload],
	// [ConferenceCreatedPayload], [ConferenceEndedPayload],
	// [ConferenceParticipantJoinedPayload], [ConferenceParticipantLeftPayload],
	// [ConferenceParticipantPlaybackEndedPayload],
	// [ConferenceParticipantPlaybackStartedPayload],
	// [ConferenceParticipantSpeakEndedPayload],
	// [ConferenceParticipantSpeakStartedPayload], [ConferencePlaybackEndedPayload],
	// [ConferencePlaybackStartedPayload], [ConferenceRecordingSavedPayload],
	// [ConferenceSpeakEndedPayload], [ConferenceSpeakStartedPayload],
	// [OutboundMessagePayload], [FaxDeliveredDataPayload], [FaxFailedDataPayload],
	// [FaxMediaProcessedDataPayload], [FaxQueuedDataPayload],
	// [FaxSendingStartedDataPayload], [shared.InboundMessagePayload],
	// [NumberOrderWithPhoneNumbers], [TranscriptionPayload]
	Payload    UnwrapWebhookEventUnionDataPayload `json:"payload"`
	RecordType string                             `json:"record_type"`
	// This field is from variant [CallConversationEnded].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [ReplacedLinkClick].
	MessageID string `json:"message_id"`
	// This field is from variant [ReplacedLinkClick].
	TimeClicked time.Time `json:"time_clicked"`
	// This field is from variant [ReplacedLinkClick].
	To string `json:"to"`
	// This field is from variant [ReplacedLinkClick].
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
	// This field is a union of [[]CallAIGatherEndedPayloadMessageHistory],
	// [[]CallAIGatherMessageHistoryUpdatedPayloadMessageHistory],
	// [[]CallAIGatherPartialResultsPayloadMessageHistory]
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
	// This field is from variant [CallAIGatherPartialResultsPayload].
	PartialResults map[string]any    `json:"partial_results"`
	CustomHeaders  []CustomSipHeader `json:"custom_headers"`
	SipHeaders     []SipHeader       `json:"sip_headers"`
	StartTime      time.Time         `json:"start_time"`
	State          string            `json:"state"`
	Tags           []string          `json:"tags"`
	// This field is from variant [CallConversationEndedPayload].
	AssistantID      string `json:"assistant_id"`
	CallingPartyType string `json:"calling_party_type"`
	// This field is from variant [CallConversationEndedPayload].
	ConversationID string `json:"conversation_id"`
	// This field is from variant [CallConversationEndedPayload].
	DurationSec int64 `json:"duration_sec"`
	// This field is from variant [CallConversationEndedPayload].
	LlmModel string `json:"llm_model"`
	// This field is from variant [CallConversationEndedPayload].
	SttModel string `json:"stt_model"`
	// This field is from variant [CallConversationEndedPayload].
	TtsModelID string `json:"tts_model_id"`
	// This field is from variant [CallConversationEndedPayload].
	TtsProvider string `json:"tts_provider"`
	// This field is from variant [CallConversationEndedPayload].
	TtsVoiceID string `json:"tts_voice_id"`
	// This field is from variant [CallConversationInsightsGeneratedPayload].
	InsightGroupID string `json:"insight_group_id"`
	// This field is from variant [CallConversationInsightsGeneratedPayload].
	Results []CallConversationInsightsGeneratedPayloadResult `json:"results"`
	// This field is from variant [CallDtmfReceivedPayload].
	Digit string `json:"digit"`
	// This field is from variant [CallEnqueuedPayload].
	CurrentPosition int64  `json:"current_position"`
	Queue           string `json:"queue"`
	// This field is from variant [CallEnqueuedPayload].
	QueueAvgWaitTimeSecs int64  `json:"queue_avg_wait_time_secs"`
	StreamType           string `json:"stream_type"`
	// This field is from variant [CallGatherEndedPayload].
	Digits string `json:"digits"`
	// This field is from variant [CallHangupPayload].
	CallQualityStats CallHangupPayloadCallQualityStats `json:"call_quality_stats"`
	HangupCause      string                            `json:"hangup_cause"`
	// This field is from variant [CallHangupPayload].
	HangupSource string `json:"hangup_source"`
	// This field is from variant [CallHangupPayload].
	SipHangupCause string `json:"sip_hangup_cause"`
	// This field is from variant [CallInitiatedPayload].
	CallScreeningResult string `json:"call_screening_result"`
	// This field is from variant [CallInitiatedPayload].
	CallerIDName string `json:"caller_id_name"`
	// This field is from variant [CallInitiatedPayload].
	ConnectionCodecs string `json:"connection_codecs"`
	Direction        string `json:"direction"`
	// This field is from variant [CallInitiatedPayload].
	OfferedCodecs string `json:"offered_codecs"`
	// This field is from variant [CallInitiatedPayload].
	ShakenStirAttestation string `json:"shaken_stir_attestation"`
	// This field is from variant [CallInitiatedPayload].
	ShakenStirValidated bool `json:"shaken_stir_validated"`
	// This field is from variant [CallLeftQueuePayload].
	QueuePosition int64  `json:"queue_position"`
	Reason        string `json:"reason"`
	// This field is from variant [CallLeftQueuePayload].
	WaitTimeSecs int64  `json:"wait_time_secs"`
	MediaName    string `json:"media_name"`
	MediaURL     string `json:"media_url"`
	Overlay      bool   `json:"overlay"`
	// This field is from variant [CallPlaybackEndedPayload].
	StatusDetail string `json:"status_detail"`
	Channels     string `json:"channels"`
	// This field is a union of [CallRecordingSavedPayloadPublicRecordingURLs],
	// [ConferenceRecordingSavedPayloadPublicRecordingURLs]
	PublicRecordingURLs UnwrapWebhookEventUnionDataPayloadPublicRecordingURLs `json:"public_recording_urls"`
	RecordingEndedAt    time.Time                                             `json:"recording_ended_at"`
	RecordingStartedAt  time.Time                                             `json:"recording_started_at"`
	// This field is a union of [CallRecordingSavedPayloadRecordingURLs],
	// [ConferenceRecordingSavedPayloadRecordingURLs]
	RecordingURLs UnwrapWebhookEventUnionDataPayloadRecordingURLs `json:"recording_urls"`
	RecordingID   string                                          `json:"recording_id"`
	// This field is from variant [CallRecordingTranscriptionSavedPayload].
	RecordingTranscriptionID string `json:"recording_transcription_id"`
	// This field is from variant [CallRecordingTranscriptionSavedPayload].
	TranscriptionText string `json:"transcription_text"`
	SipNotifyResponse int64  `json:"sip_notify_response"`
	// This field is from variant [CallSiprecFailedPayload].
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
	// This field is from variant [ConferenceRecordingSavedPayload].
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
	Errors        []shared.MessagingError                         `json:"errors"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media              UnwrapWebhookEventUnionDataPayloadMedia `json:"media"`
	MessagingProfileID string                                  `json:"messaging_profile_id"`
	OrganizationID     string                                  `json:"organization_id"`
	Parts              int64                                   `json:"parts"`
	ReceivedAt         time.Time                               `json:"received_at"`
	RecordType         string                                  `json:"record_type"`
	SentAt             time.Time                               `json:"sent_at"`
	// This field is from variant [OutboundMessagePayload].
	SmartEncodingApplied  bool      `json:"smart_encoding_applied"`
	Subject               string    `json:"subject"`
	TcrCampaignBillable   bool      `json:"tcr_campaign_billable"`
	TcrCampaignID         string    `json:"tcr_campaign_id"`
	TcrCampaignRegistered string    `json:"tcr_campaign_registered"`
	Text                  string    `json:"text"`
	Type                  string    `json:"type"`
	ValidUntil            time.Time `json:"valid_until"`
	WebhookFailoverURL    string    `json:"webhook_failover_url"`
	WebhookURL            string    `json:"webhook_url"`
	// This field is from variant [FaxDeliveredDataPayload].
	CallDurationSecs int64  `json:"call_duration_secs"`
	FaxID            string `json:"fax_id"`
	OriginalMediaURL string `json:"original_media_url"`
	// This field is from variant [FaxDeliveredDataPayload].
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
	// This field is from variant [TranscriptionPayload].
	TranscriptionData TranscriptionPayloadTranscriptionData `json:"transcription_data"`
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
		SmartEncodingApplied     respjson.Field
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
// will be valid: OfCallAIGatherEndedPayloadMessageHistoryArray
// OfCallAIGatherMessageHistoryUpdatedPayloadMessageHistoryArray
// OfCallAIGatherPartialResultsPayloadMessageHistoryArray]
type UnwrapWebhookEventUnionDataPayloadMessageHistory struct {
	// This field will be present if the value is a
	// [[]CallAIGatherEndedPayloadMessageHistory] instead of an object.
	OfCallAIGatherEndedPayloadMessageHistoryArray []CallAIGatherEndedPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherMessageHistoryUpdatedPayloadMessageHistory] instead of an object.
	OfCallAIGatherMessageHistoryUpdatedPayloadMessageHistoryArray []CallAIGatherMessageHistoryUpdatedPayloadMessageHistory `json:",inline"`
	// This field will be present if the value is a
	// [[]CallAIGatherPartialResultsPayloadMessageHistory] instead of an object.
	OfCallAIGatherPartialResultsPayloadMessageHistoryArray []CallAIGatherPartialResultsPayloadMessageHistory `json:",inline"`
	JSON                                                   struct {
		OfCallAIGatherEndedPayloadMessageHistoryArray                 respjson.Field
		OfCallAIGatherMessageHistoryUpdatedPayloadMessageHistoryArray respjson.Field
		OfCallAIGatherPartialResultsPayloadMessageHistoryArray        respjson.Field
		raw                                                           string
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
// will be valid: OfCallAIGatherEndedPayloadResult
// OfCallMachinePremiumGreetingEndedPayloadResult]
type UnwrapWebhookEventUnionDataPayloadResult struct {
	// This field will be present if the value is a [any] instead of an object.
	OfCallAIGatherEndedPayloadResult any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfCallMachinePremiumGreetingEndedPayloadResult string `json:",inline"`
	JSON                                           struct {
		OfCallAIGatherEndedPayloadResult               respjson.Field
		OfCallMachinePremiumGreetingEndedPayloadResult respjson.Field
		raw                                            string
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
