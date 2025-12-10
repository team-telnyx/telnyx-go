// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// MessageService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
	Rcs     MessageRcService
}

// NewMessageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMessageService(opts ...option.RequestOption) (r MessageService) {
	r = MessageService{}
	r.Options = opts
	r.Rcs = NewMessageRcService(opts...)
	return
}

// Note: This API endpoint can only retrieve messages that are no older than 10
// days since their creation. If you require messages older than this, please
// generate an
// [MDR report.](https://developers.telnyx.com/api/v1/mission-control/add-mdr-request)
func (r *MessageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a scheduled message that has not yet been sent. Only messages with
// `status=scheduled` and `send_at` more than a minute from now can be cancelled.
func (r *MessageService) CancelScheduled(ctx context.Context, id string, opts ...option.RequestOption) (res *MessageCancelScheduledResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Schedule a message with a Phone Number, Alphanumeric Sender ID, Short Code or
// Number Pool.
//
// This endpoint allows you to schedule a message with any messaging resource.
// Current messaging resources include: long-code, short-code, number-pool, and
// alphanumeric-sender-id.
func (r *MessageService) Schedule(ctx context.Context, body MessageScheduleParams, opts ...option.RequestOption) (res *MessageScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/schedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a message with a Phone Number, Alphanumeric Sender ID, Short Code or Number
// Pool.
//
// This endpoint allows you to send a message with any messaging resource. Current
// messaging resources include: long-code, short-code, number-pool, and
// alphanumeric-sender-id.
func (r *MessageService) Send(ctx context.Context, body MessageSendParams, opts ...option.RequestOption) (res *MessageSendResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a group MMS message
func (r *MessageService) SendGroupMms(ctx context.Context, body MessageSendGroupMmsParams, opts ...option.RequestOption) (res *MessageSendGroupMmsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/group_mms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a long code message
func (r *MessageService) SendLongCode(ctx context.Context, body MessageSendLongCodeParams, opts ...option.RequestOption) (res *MessageSendLongCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/long_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a message using number pool
func (r *MessageService) SendNumberPool(ctx context.Context, body MessageSendNumberPoolParams, opts ...option.RequestOption) (res *MessageSendNumberPoolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/number_pool"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Send a short code message
func (r *MessageService) SendShortCode(ctx context.Context, body MessageSendShortCodeParams, opts ...option.RequestOption) (res *MessageSendShortCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/short_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessagingError struct {
	Code   string               `json:"code,required" format:"integer"`
	Title  string               `json:"title,required"`
	Detail string               `json:"detail"`
	Meta   map[string]any       `json:"meta"`
	Source MessagingErrorSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Title       respjson.Field
		Detail      respjson.Field
		Meta        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingError) RawJSON() string { return r.JSON.raw }
func (r *MessagingError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingErrorSource struct {
	// Indicates which query parameter caused the error.
	Parameter string `json:"parameter"`
	// JSON pointer (RFC6901) to the offending entity.
	Pointer string `json:"pointer" format:"json-pointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameter   respjson.Field
		Pointer     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingErrorSource) RawJSON() string { return r.JSON.raw }
func (r *MessagingErrorSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessagePayload struct {
	// Identifies the type of resource.
	ID string                     `json:"id" format:"uuid"`
	Cc []OutboundMessagePayloadCc `json:"cc"`
	// ISO 8601 formatted date indicating when the message was finalized.
	CompletedAt time.Time                  `json:"completed_at,nullable" format:"date-time"`
	Cost        OutboundMessagePayloadCost `json:"cost,nullable"`
	// Detailed breakdown of the message cost components.
	CostBreakdown OutboundMessagePayloadCostBreakdown `json:"cost_breakdown,nullable"`
	// The direction of the message. Inbound messages are sent to you whereas outbound
	// messages are sent from you.
	//
	// Any of "outbound".
	Direction OutboundMessagePayloadDirection `json:"direction"`
	// Encoding scheme used for the message body.
	Encoding string `json:"encoding"`
	// These errors may point at addressees when referring to unsuccessful/unconfirmed
	// delivery statuses.
	Errors []MessagingError              `json:"errors"`
	From   OutboundMessagePayloadFrom    `json:"from"`
	Media  []OutboundMessagePayloadMedia `json:"media"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The id of the organization the messaging profile belongs to.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Number of parts into which the message's body must be split.
	Parts int64 `json:"parts"`
	// ISO 8601 formatted date indicating when the message request was received.
	ReceivedAt time.Time `json:"received_at" format:"date-time"`
	// Identifies the type of the resource.
	//
	// Any of "message".
	RecordType OutboundMessagePayloadRecordType `json:"record_type"`
	// ISO 8601 formatted date indicating when the message was sent.
	SentAt time.Time `json:"sent_at,nullable" format:"date-time"`
	// Subject of multimedia message
	Subject string `json:"subject,nullable"`
	// Tags associated with the resource.
	Tags []string `json:"tags"`
	// Indicates whether the TCR campaign is billable.
	TcrCampaignBillable bool `json:"tcr_campaign_billable"`
	// The Campaign Registry (TCR) campaign ID associated with the message.
	TcrCampaignID string `json:"tcr_campaign_id,nullable"`
	// The registration status of the TCR campaign.
	TcrCampaignRegistered string `json:"tcr_campaign_registered,nullable"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text string                     `json:"text"`
	To   []OutboundMessagePayloadTo `json:"to"`
	// The type of message.
	//
	// Any of "SMS", "MMS".
	Type OutboundMessagePayloadType `json:"type"`
	// Message must be out of the queue by this time or else it will be discarded and
	// marked as 'sending_failed'. Once the message moves out of the queue, this field
	// will be nulled
	ValidUntil time.Time `json:"valid_until,nullable" format:"date-time"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL string `json:"webhook_failover_url,nullable" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL string `json:"webhook_url,nullable" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Cc                    respjson.Field
		CompletedAt           respjson.Field
		Cost                  respjson.Field
		CostBreakdown         respjson.Field
		Direction             respjson.Field
		Encoding              respjson.Field
		Errors                respjson.Field
		From                  respjson.Field
		Media                 respjson.Field
		MessagingProfileID    respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		Subject               respjson.Field
		Tags                  respjson.Field
		TcrCampaignBillable   respjson.Field
		TcrCampaignID         respjson.Field
		TcrCampaignRegistered respjson.Field
		Text                  respjson.Field
		To                    respjson.Field
		Type                  respjson.Field
		ValidUntil            respjson.Field
		WebhookFailoverURL    respjson.Field
		WebhookURL            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayload) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessagePayloadCc struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// Any of "queued", "sending", "sent", "delivered", "sending_failed",
	// "delivery_failed", "delivery_unconfirmed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadCc) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadCc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessagePayloadCost struct {
	// The amount deducted from your account.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadCost) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of the message cost components.
type OutboundMessagePayloadCostBreakdown struct {
	CarrierFee OutboundMessagePayloadCostBreakdownCarrierFee `json:"carrier_fee"`
	Rate       OutboundMessagePayloadCostBreakdownRate       `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierFee  respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadCostBreakdown) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadCostBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessagePayloadCostBreakdownCarrierFee struct {
	// The carrier fee amount.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadCostBreakdownCarrierFee) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadCostBreakdownCarrierFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessagePayloadCostBreakdownRate struct {
	// The rate amount applied.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadCostBreakdownRate) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadCostBreakdownRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the message. Inbound messages are sent to you whereas outbound
// messages are sent from you.
type OutboundMessagePayloadDirection string

const (
	OutboundMessagePayloadDirectionOutbound OutboundMessagePayloadDirection = "outbound"
)

type OutboundMessagePayloadFrom struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadFrom) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutboundMessagePayloadMedia struct {
	// The MIME type of the requested media.
	ContentType string `json:"content_type,nullable" format:"mime-type"`
	// The SHA256 hash of the requested media.
	Sha256 string `json:"sha256,nullable"`
	// The size of the requested media.
	Size int64 `json:"size,nullable"`
	// The url of the media requested to be sent.
	URL string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Sha256      respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadMedia) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type OutboundMessagePayloadRecordType string

const (
	OutboundMessagePayloadRecordTypeMessage OutboundMessagePayloadRecordType = "message"
)

type OutboundMessagePayloadTo struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// The delivery status of the message.
	//
	// Any of "queued", "sending", "sent", "expired", "sending_failed",
	// "delivery_unconfirmed", "delivered", "delivery_failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutboundMessagePayloadTo) RawJSON() string { return r.JSON.raw }
func (r *OutboundMessagePayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of message.
type OutboundMessagePayloadType string

const (
	OutboundMessagePayloadTypeSMS OutboundMessagePayloadType = "SMS"
	OutboundMessagePayloadTypeMms OutboundMessagePayloadType = "MMS"
)

type MessageGetResponse struct {
	Data MessageGetResponseDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnion contains all possible properties and values from
// [OutboundMessagePayload], [shared.InboundMessagePayload].
//
// Use the [MessageGetResponseDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageGetResponseDataUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]OutboundMessagePayloadCc],
	// [[]shared.InboundMessagePayloadCc]
	Cc          MessageGetResponseDataUnionCc `json:"cc"`
	CompletedAt time.Time                     `json:"completed_at"`
	// This field is a union of [OutboundMessagePayloadCost],
	// [shared.InboundMessagePayloadCost]
	Cost MessageGetResponseDataUnionCost `json:"cost"`
	// This field is a union of [OutboundMessagePayloadCostBreakdown],
	// [shared.InboundMessagePayloadCostBreakdown]
	CostBreakdown MessageGetResponseDataUnionCostBreakdown `json:"cost_breakdown"`
	// Any of "outbound", "inbound".
	Direction string           `json:"direction"`
	Encoding  string           `json:"encoding"`
	Errors    []MessagingError `json:"errors"`
	// This field is a union of [OutboundMessagePayloadFrom],
	// [shared.InboundMessagePayloadFrom]
	From MessageGetResponseDataUnionFrom `json:"from"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media                 MessageGetResponseDataUnionMedia `json:"media"`
	MessagingProfileID    string                           `json:"messaging_profile_id"`
	OrganizationID        string                           `json:"organization_id"`
	Parts                 int64                            `json:"parts"`
	ReceivedAt            time.Time                        `json:"received_at"`
	RecordType            string                           `json:"record_type"`
	SentAt                time.Time                        `json:"sent_at"`
	Subject               string                           `json:"subject"`
	Tags                  []string                         `json:"tags"`
	TcrCampaignBillable   bool                             `json:"tcr_campaign_billable"`
	TcrCampaignID         string                           `json:"tcr_campaign_id"`
	TcrCampaignRegistered string                           `json:"tcr_campaign_registered"`
	Text                  string                           `json:"text"`
	// This field is a union of [[]OutboundMessagePayloadTo],
	// [[]shared.InboundMessagePayloadTo]
	To                 MessageGetResponseDataUnionTo `json:"to"`
	Type               string                        `json:"type"`
	ValidUntil         time.Time                     `json:"valid_until"`
	WebhookFailoverURL string                        `json:"webhook_failover_url"`
	WebhookURL         string                        `json:"webhook_url"`
	JSON               struct {
		ID                    respjson.Field
		Cc                    respjson.Field
		CompletedAt           respjson.Field
		Cost                  respjson.Field
		CostBreakdown         respjson.Field
		Direction             respjson.Field
		Encoding              respjson.Field
		Errors                respjson.Field
		From                  respjson.Field
		Media                 respjson.Field
		MessagingProfileID    respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		Subject               respjson.Field
		Tags                  respjson.Field
		TcrCampaignBillable   respjson.Field
		TcrCampaignID         respjson.Field
		TcrCampaignRegistered respjson.Field
		Text                  respjson.Field
		To                    respjson.Field
		Type                  respjson.Field
		ValidUntil            respjson.Field
		WebhookFailoverURL    respjson.Field
		WebhookURL            respjson.Field
		raw                   string
	} `json:"-"`
}

// anyMessageGetResponseData is implemented by each variant of
// [MessageGetResponseDataUnion] to add type safety for the return type of
// [MessageGetResponseDataUnion.AsAny]
type anyMessageGetResponseData interface {
	ImplMessageGetResponseDataUnion()
}

func (OutboundMessagePayload) ImplMessageGetResponseDataUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageGetResponseDataUnion.AsAny().(type) {
//	case telnyx.OutboundMessagePayload:
//	case shared.InboundMessagePayload:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MessageGetResponseDataUnion) AsAny() anyMessageGetResponseData {
	switch u.Direction {
	case "outbound":
		return u.AsOutbound()
	case "inbound":
		return u.AsInbound()
	}
	return nil
}

func (u MessageGetResponseDataUnion) AsOutbound() (v OutboundMessagePayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageGetResponseDataUnion) AsInbound() (v shared.InboundMessagePayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageGetResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageGetResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionCc is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionCc provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadCcArray OfInboundMessagePayloadCcArray]
type MessageGetResponseDataUnionCc struct {
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

func (r *MessageGetResponseDataUnionCc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionCost is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionCost provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
type MessageGetResponseDataUnionCost struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *MessageGetResponseDataUnionCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionCostBreakdown is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionCostBreakdown provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
type MessageGetResponseDataUnionCostBreakdown struct {
	// This field is a union of [OutboundMessagePayloadCostBreakdownCarrierFee],
	// [shared.InboundMessagePayloadCostBreakdownCarrierFee]
	CarrierFee MessageGetResponseDataUnionCostBreakdownCarrierFee `json:"carrier_fee"`
	// This field is a union of [OutboundMessagePayloadCostBreakdownRate],
	// [shared.InboundMessagePayloadCostBreakdownRate]
	Rate MessageGetResponseDataUnionCostBreakdownRate `json:"rate"`
	JSON struct {
		CarrierFee respjson.Field
		Rate       respjson.Field
		raw        string
	} `json:"-"`
}

func (r *MessageGetResponseDataUnionCostBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionCostBreakdownCarrierFee is an implicit subunion of
// [MessageGetResponseDataUnion].
// MessageGetResponseDataUnionCostBreakdownCarrierFee provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
type MessageGetResponseDataUnionCostBreakdownCarrierFee struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *MessageGetResponseDataUnionCostBreakdownCarrierFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionCostBreakdownRate is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionCostBreakdownRate
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
type MessageGetResponseDataUnionCostBreakdownRate struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	JSON     struct {
		Amount   respjson.Field
		Currency respjson.Field
		raw      string
	} `json:"-"`
}

func (r *MessageGetResponseDataUnionCostBreakdownRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionFrom is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionFrom provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
type MessageGetResponseDataUnionFrom struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	// This field is from variant [shared.InboundMessagePayloadFrom].
	Status string `json:"status"`
	JSON   struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		raw         string
	} `json:"-"`
}

func (r *MessageGetResponseDataUnionFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionMedia is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionMedia provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadMedia OfInboundMessagePayloadMedia]
type MessageGetResponseDataUnionMedia struct {
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

func (r *MessageGetResponseDataUnionMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageGetResponseDataUnionTo is an implicit subunion of
// [MessageGetResponseDataUnion]. MessageGetResponseDataUnionTo provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageGetResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutboundMessagePayloadToArray OfInboundMessagePayloadToArray]
type MessageGetResponseDataUnionTo struct {
	// This field will be present if the value is a [[]OutboundMessagePayloadTo]
	// instead of an object.
	OfOutboundMessagePayloadToArray []OutboundMessagePayloadTo `json:",inline"`
	// This field will be present if the value is a [[]shared.InboundMessagePayloadTo]
	// instead of an object.
	OfInboundMessagePayloadToArray []shared.InboundMessagePayloadTo `json:",inline"`
	JSON                           struct {
		OfOutboundMessagePayloadToArray respjson.Field
		OfInboundMessagePayloadToArray  respjson.Field
		raw                             string
	} `json:"-"`
}

func (r *MessageGetResponseDataUnionTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageCancelScheduledResponse struct {
	// Identifies the type of resource.
	ID string                             `json:"id" format:"uuid"`
	Cc []MessageCancelScheduledResponseCc `json:"cc"`
	// ISO 8601 formatted date indicating when the message was finalized.
	CompletedAt time.Time                          `json:"completed_at,nullable" format:"date-time"`
	Cost        MessageCancelScheduledResponseCost `json:"cost,nullable"`
	// Detailed breakdown of the message cost components.
	CostBreakdown MessageCancelScheduledResponseCostBreakdown `json:"cost_breakdown,nullable"`
	// The direction of the message. Inbound messages are sent to you whereas outbound
	// messages are sent from you.
	//
	// Any of "outbound".
	Direction MessageCancelScheduledResponseDirection `json:"direction"`
	// Encoding scheme used for the message body.
	Encoding string `json:"encoding"`
	// These errors may point at addressees when referring to unsuccessful/unconfirmed
	// delivery statuses.
	Errors []MessagingError                      `json:"errors"`
	From   MessageCancelScheduledResponseFrom    `json:"from"`
	Media  []MessageCancelScheduledResponseMedia `json:"media"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The id of the organization the messaging profile belongs to.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Number of parts into which the message's body must be split.
	Parts int64 `json:"parts"`
	// ISO 8601 formatted date indicating when the message request was received.
	ReceivedAt time.Time `json:"received_at" format:"date-time"`
	// Identifies the type of the resource.
	//
	// Any of "message".
	RecordType MessageCancelScheduledResponseRecordType `json:"record_type"`
	// ISO 8601 formatted date indicating when the message was sent.
	SentAt time.Time `json:"sent_at,nullable" format:"date-time"`
	// Subject of multimedia message
	Subject string `json:"subject,nullable"`
	// Tags associated with the resource.
	Tags []string `json:"tags"`
	// Indicates whether the TCR campaign is billable.
	TcrCampaignBillable bool `json:"tcr_campaign_billable"`
	// The Campaign Registry (TCR) campaign ID associated with the message.
	TcrCampaignID string `json:"tcr_campaign_id,nullable"`
	// The registration status of the TCR campaign.
	TcrCampaignRegistered string `json:"tcr_campaign_registered,nullable"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text string                             `json:"text"`
	To   []MessageCancelScheduledResponseTo `json:"to"`
	// The type of message.
	//
	// Any of "SMS", "MMS".
	Type MessageCancelScheduledResponseType `json:"type"`
	// Message must be out of the queue by this time or else it will be discarded and
	// marked as 'sending_failed'. Once the message moves out of the queue, this field
	// will be nulled
	ValidUntil time.Time `json:"valid_until,nullable" format:"date-time"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL string `json:"webhook_failover_url,nullable" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL string `json:"webhook_url,nullable" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Cc                    respjson.Field
		CompletedAt           respjson.Field
		Cost                  respjson.Field
		CostBreakdown         respjson.Field
		Direction             respjson.Field
		Encoding              respjson.Field
		Errors                respjson.Field
		From                  respjson.Field
		Media                 respjson.Field
		MessagingProfileID    respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		Subject               respjson.Field
		Tags                  respjson.Field
		TcrCampaignBillable   respjson.Field
		TcrCampaignID         respjson.Field
		TcrCampaignRegistered respjson.Field
		Text                  respjson.Field
		To                    respjson.Field
		Type                  respjson.Field
		ValidUntil            respjson.Field
		WebhookFailoverURL    respjson.Field
		WebhookURL            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageCancelScheduledResponseCc struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// The delivery status of the message.
	//
	// Any of "scheduled", "queued", "sending", "sent", "cancelled", "expired",
	// "sending_failed", "delivery_unconfirmed", "delivered", "delivery_failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseCc) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseCc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageCancelScheduledResponseCost struct {
	// The amount deducted from your account.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseCost) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseCost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed breakdown of the message cost components.
type MessageCancelScheduledResponseCostBreakdown struct {
	CarrierFee MessageCancelScheduledResponseCostBreakdownCarrierFee `json:"carrier_fee"`
	Rate       MessageCancelScheduledResponseCostBreakdownRate       `json:"rate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarrierFee  respjson.Field
		Rate        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseCostBreakdown) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseCostBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageCancelScheduledResponseCostBreakdownCarrierFee struct {
	// The carrier fee amount.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseCostBreakdownCarrierFee) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseCostBreakdownCarrierFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageCancelScheduledResponseCostBreakdownRate struct {
	// The rate amount applied.
	Amount string `json:"amount" format:"decimal"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency" format:"iso4217"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseCostBreakdownRate) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseCostBreakdownRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The direction of the message. Inbound messages are sent to you whereas outbound
// messages are sent from you.
type MessageCancelScheduledResponseDirection string

const (
	MessageCancelScheduledResponseDirectionOutbound MessageCancelScheduledResponseDirection = "outbound"
)

type MessageCancelScheduledResponseFrom struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseFrom) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageCancelScheduledResponseMedia struct {
	// The MIME type of the requested media.
	ContentType string `json:"content_type,nullable" format:"mime-type"`
	// The SHA256 hash of the requested media.
	Sha256 string `json:"sha256,nullable"`
	// The size of the requested media.
	Size int64 `json:"size,nullable"`
	// The url of the media requested to be sent.
	URL string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Sha256      respjson.Field
		Size        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseMedia) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type MessageCancelScheduledResponseRecordType string

const (
	MessageCancelScheduledResponseRecordTypeMessage MessageCancelScheduledResponseRecordType = "message"
)

type MessageCancelScheduledResponseTo struct {
	// The carrier of the receiver.
	Carrier string `json:"carrier"`
	// The line-type of the receiver.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Receiving address (+E.164 formatted phone number or short code).
	PhoneNumber string `json:"phone_number" format:"address"`
	// The delivery status of the message.
	//
	// Any of "scheduled", "queued", "sending", "sent", "cancelled", "expired",
	// "sending_failed", "delivery_unconfirmed", "delivered", "delivery_failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier     respjson.Field
		LineType    respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageCancelScheduledResponseTo) RawJSON() string { return r.JSON.raw }
func (r *MessageCancelScheduledResponseTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of message.
type MessageCancelScheduledResponseType string

const (
	MessageCancelScheduledResponseTypeSMS MessageCancelScheduledResponseType = "SMS"
	MessageCancelScheduledResponseTypeMms MessageCancelScheduledResponseType = "MMS"
)

type MessageScheduleResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageScheduleResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageScheduleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendGroupMmsResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendGroupMmsResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendGroupMmsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendLongCodeResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendLongCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendLongCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendNumberPoolResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendNumberPoolResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendNumberPoolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendShortCodeResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendShortCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendShortCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageScheduleParams struct {
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required" format:"address"`
	// Automatically detect if an SMS message is unusually long and exceeds a
	// recommended limit of message parts.
	AutoDetect param.Opt[bool] `json:"auto_detect,omitzero"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	//
	// **Required if sending with a phone number, short code, or alphanumeric sender
	// ID.**
	From param.Opt[string] `json:"from,omitzero" format:"address"`
	// Unique identifier for a messaging profile.
	//
	// **Required if sending via number pool or with an alphanumeric sender ID.**
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	// ISO 8601 formatted date indicating when to send the message - accurate up till a
	// minute.
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	// Subject of multimedia message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text param.Opt[string] `json:"text,omitzero"`
	// If the profile this number is associated with has webhooks, use them for
	// delivery notifications. If webhooks are also specified on the message itself,
	// they will be attempted first, then those on the profile.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// A list of media URLs. The total media size must be less than 1 MB.
	//
	// **Required for MMS**
	MediaURLs []string `json:"media_urls,omitzero" format:"url"`
	// The protocol for sending the message, either SMS or MMS.
	//
	// Any of "SMS", "MMS".
	Type MessageScheduleParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageScheduleParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageScheduleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageScheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol for sending the message, either SMS or MMS.
type MessageScheduleParamsType string

const (
	MessageScheduleParamsTypeSMS MessageScheduleParamsType = "SMS"
	MessageScheduleParamsTypeMms MessageScheduleParamsType = "MMS"
)

type MessageSendParams struct {
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required" format:"address"`
	// ISO 8601 formatted date indicating when to send the message - accurate up till a
	// minute.
	SendAt param.Opt[time.Time] `json:"send_at,omitzero" format:"date-time"`
	// Automatically detect if an SMS message is unusually long and exceeds a
	// recommended limit of message parts.
	AutoDetect param.Opt[bool] `json:"auto_detect,omitzero"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	//
	// **Required if sending with a phone number, short code, or alphanumeric sender
	// ID.**
	From param.Opt[string] `json:"from,omitzero" format:"address"`
	// Unique identifier for a messaging profile.
	//
	// **Required if sending via number pool or with an alphanumeric sender ID.**
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	// Subject of multimedia message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text param.Opt[string] `json:"text,omitzero"`
	// If the profile this number is associated with has webhooks, use them for
	// delivery notifications. If webhooks are also specified on the message itself,
	// they will be attempted first, then those on the profile.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// A list of media URLs. The total media size must be less than 1 MB.
	//
	// **Required for MMS**
	MediaURLs []string `json:"media_urls,omitzero" format:"url"`
	// The protocol for sending the message, either SMS or MMS.
	//
	// Any of "SMS", "MMS".
	Type MessageSendParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageSendParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol for sending the message, either SMS or MMS.
type MessageSendParamsType string

const (
	MessageSendParamsTypeSMS MessageSendParamsType = "SMS"
	MessageSendParamsTypeMms MessageSendParamsType = "MMS"
)

type MessageSendGroupMmsParams struct {
	// Phone number, in +E.164 format, used to send the message.
	From string `json:"from,required" format:"address"`
	// A list of destinations. No more than 8 destinations are allowed.
	To []string `json:"to,omitzero,required" format:"address"`
	// Subject of multimedia message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Message body (i.e., content) as a non-empty string.
	Text param.Opt[string] `json:"text,omitzero"`
	// If the profile this number is associated with has webhooks, use them for
	// delivery notifications. If webhooks are also specified on the message itself,
	// they will be attempted first, then those on the profile.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// A list of media URLs. The total media size must be less than 1 MB.
	MediaURLs []string `json:"media_urls,omitzero" format:"url"`
	paramObj
}

func (r MessageSendGroupMmsParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendGroupMmsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendGroupMmsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendLongCodeParams struct {
	// Phone number, in +E.164 format, used to send the message.
	From string `json:"from,required" format:"address"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required" format:"address"`
	// Automatically detect if an SMS message is unusually long and exceeds a
	// recommended limit of message parts.
	AutoDetect param.Opt[bool] `json:"auto_detect,omitzero"`
	// Subject of multimedia message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text param.Opt[string] `json:"text,omitzero"`
	// If the profile this number is associated with has webhooks, use them for
	// delivery notifications. If webhooks are also specified on the message itself,
	// they will be attempted first, then those on the profile.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// A list of media URLs. The total media size must be less than 1 MB.
	//
	// **Required for MMS**
	MediaURLs []string `json:"media_urls,omitzero" format:"url"`
	// The protocol for sending the message, either SMS or MMS.
	//
	// Any of "SMS", "MMS".
	Type MessageSendLongCodeParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageSendLongCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendLongCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendLongCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol for sending the message, either SMS or MMS.
type MessageSendLongCodeParamsType string

const (
	MessageSendLongCodeParamsTypeSMS MessageSendLongCodeParamsType = "SMS"
	MessageSendLongCodeParamsTypeMms MessageSendLongCodeParamsType = "MMS"
)

type MessageSendNumberPoolParams struct {
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id,required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required" format:"address"`
	// Automatically detect if an SMS message is unusually long and exceeds a
	// recommended limit of message parts.
	AutoDetect param.Opt[bool] `json:"auto_detect,omitzero"`
	// Subject of multimedia message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text param.Opt[string] `json:"text,omitzero"`
	// If the profile this number is associated with has webhooks, use them for
	// delivery notifications. If webhooks are also specified on the message itself,
	// they will be attempted first, then those on the profile.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// A list of media URLs. The total media size must be less than 1 MB.
	//
	// **Required for MMS**
	MediaURLs []string `json:"media_urls,omitzero" format:"url"`
	// The protocol for sending the message, either SMS or MMS.
	//
	// Any of "SMS", "MMS".
	Type MessageSendNumberPoolParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageSendNumberPoolParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendNumberPoolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendNumberPoolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol for sending the message, either SMS or MMS.
type MessageSendNumberPoolParamsType string

const (
	MessageSendNumberPoolParamsTypeSMS MessageSendNumberPoolParamsType = "SMS"
	MessageSendNumberPoolParamsTypeMms MessageSendNumberPoolParamsType = "MMS"
)

type MessageSendShortCodeParams struct {
	// Phone number, in +E.164 format, used to send the message.
	From string `json:"from,required" format:"address"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required" format:"address"`
	// Automatically detect if an SMS message is unusually long and exceeds a
	// recommended limit of message parts.
	AutoDetect param.Opt[bool] `json:"auto_detect,omitzero"`
	// Subject of multimedia message
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Message body (i.e., content) as a non-empty string.
	//
	// **Required for SMS**
	Text param.Opt[string] `json:"text,omitzero"`
	// If the profile this number is associated with has webhooks, use them for
	// delivery notifications. If webhooks are also specified on the message itself,
	// they will be attempted first, then those on the profile.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// A list of media URLs. The total media size must be less than 1 MB.
	//
	// **Required for MMS**
	MediaURLs []string `json:"media_urls,omitzero" format:"url"`
	// The protocol for sending the message, either SMS or MMS.
	//
	// Any of "SMS", "MMS".
	Type MessageSendShortCodeParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageSendShortCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendShortCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendShortCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol for sending the message, either SMS or MMS.
type MessageSendShortCodeParamsType string

const (
	MessageSendShortCodeParamsTypeSMS MessageSendShortCodeParamsType = "SMS"
	MessageSendShortCodeParamsTypeMms MessageSendShortCodeParamsType = "MMS"
)
