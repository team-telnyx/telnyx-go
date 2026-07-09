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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Messages
//
// MessageService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessageService] method instead.
type MessageService struct {
	Options []option.RequestOption
	// Send RCS messages
	Rcs MessageRcService
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
// [MDR report.](https://developers.telnyx.com/api-reference/mdr-usage-reports/create-mdr-usage-report)
func (r *MessageService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancel a scheduled message that has not yet been sent. Only messages with
// `status=scheduled` and `send_at` more than a minute from now can be cancelled.
func (r *MessageService) CancelScheduled(ctx context.Context, id string, opts ...option.RequestOption) (res *MessageCancelScheduledResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("messages/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieve all messages in a group MMS conversation by the group message ID.
func (r *MessageService) GetGroupMessages(ctx context.Context, messageID string, opts ...option.RequestOption) (res *MessageGetGroupMessagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("messages/group/%s", messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	return res, err
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
	return res, err
}

// Send a group MMS message
func (r *MessageService) SendGroupMms(ctx context.Context, body MessageSendGroupMmsParams, opts ...option.RequestOption) (res *MessageSendGroupMmsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/group_mms"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a long code message
func (r *MessageService) SendLongCode(ctx context.Context, body MessageSendLongCodeParams, opts ...option.RequestOption) (res *MessageSendLongCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/long_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a message using number pool
func (r *MessageService) SendNumberPool(ctx context.Context, body MessageSendNumberPoolParams, opts ...option.RequestOption) (res *MessageSendNumberPoolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/number_pool"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a short code message
func (r *MessageService) SendShortCode(ctx context.Context, body MessageSendShortCodeParams, opts ...option.RequestOption) (res *MessageSendShortCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/short_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send an SMS message using an alphanumeric sender ID. This is SMS only.
func (r *MessageService) SendWithAlphanumericSender(ctx context.Context, body MessageSendWithAlphanumericSenderParams, opts ...option.RequestOption) (res *MessageSendWithAlphanumericSenderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/alphanumeric_sender_id"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type OutboundMessagePayload struct {
	// Identifies the type of resource.
	ID string                     `json:"id" format:"uuid"`
	Cc []OutboundMessagePayloadCc `json:"cc"`
	// ISO 8601 formatted date indicating when the message was finalized.
	CompletedAt time.Time                  `json:"completed_at" api:"nullable" format:"date-time"`
	Cost        OutboundMessagePayloadCost `json:"cost" api:"nullable"`
	// Detailed breakdown of the message cost components.
	CostBreakdown OutboundMessagePayloadCostBreakdown `json:"cost_breakdown" api:"nullable"`
	// The direction of the message. Inbound messages are sent to you whereas outbound
	// messages are sent from you.
	//
	// Any of "outbound".
	Direction OutboundMessagePayloadDirection `json:"direction"`
	// Encoding scheme used for the message body.
	Encoding string `json:"encoding"`
	// These errors may point at addressees when referring to unsuccessful/unconfirmed
	// delivery statuses.
	Errors []shared.MessagingError       `json:"errors"`
	From   OutboundMessagePayloadFrom    `json:"from"`
	Media  []OutboundMessagePayloadMedia `json:"media"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The number of characters in the message text
	NumChars int64 `json:"num_chars"`
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
	SentAt time.Time `json:"sent_at" api:"nullable" format:"date-time"`
	// Indicates whether smart encoding was applied to this message. When `true`, one
	// or more Unicode characters were automatically replaced with GSM-7 equivalents to
	// reduce segment count and cost. The original message text is preserved in
	// webhooks.
	SmartEncodingApplied bool `json:"smart_encoding_applied"`
	// Subject of multimedia message
	Subject string `json:"subject" api:"nullable"`
	// Tags associated with the resource.
	Tags []string `json:"tags"`
	// Indicates whether the TCR campaign is billable.
	TcrCampaignBillable bool `json:"tcr_campaign_billable"`
	// The Campaign Registry (TCR) campaign ID associated with the message.
	TcrCampaignID string `json:"tcr_campaign_id" api:"nullable"`
	// The registration status of the TCR campaign.
	TcrCampaignRegistered string `json:"tcr_campaign_registered" api:"nullable"`
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
	ValidUntil time.Time `json:"valid_until" api:"nullable" format:"date-time"`
	// Seconds the message is queued due to rate limiting before being sent to the
	// carrier. Represents the maximum wait across all applicable rate limits (account,
	// carrier, campaign). 0.0 = no queuing delay.
	WaitSeconds float64 `json:"wait_seconds" api:"nullable"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL string `json:"webhook_failover_url" api:"nullable" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL string `json:"webhook_url" api:"nullable" format:"url"`
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
		NumChars              respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		SmartEncodingApplied  respjson.Field
		Subject               respjson.Field
		Tags                  respjson.Field
		TcrCampaignBillable   respjson.Field
		TcrCampaignID         respjson.Field
		TcrCampaignRegistered respjson.Field
		Text                  respjson.Field
		To                    respjson.Field
		Type                  respjson.Field
		ValidUntil            respjson.Field
		WaitSeconds           respjson.Field
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
	PhoneNumber string `json:"phone_number"`
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
	Amount string `json:"amount"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
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
	Amount string `json:"amount"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
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
	Amount string `json:"amount"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
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
	PhoneNumber string `json:"phone_number"`
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
	ContentType string `json:"content_type" api:"nullable"`
	// The SHA256 hash of the requested media.
	Sha256 string `json:"sha256" api:"nullable"`
	// The size of the requested media.
	Size int64 `json:"size" api:"nullable"`
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
	PhoneNumber string `json:"phone_number"`
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

type RcsAgentMessage struct {
	ContentMessage RcsAgentMessageContentMessage `json:"content_message"`
	// RCS Event to send to the recipient
	Event RcsAgentMessageEvent `json:"event"`
	// Timestamp in UTC of when this message is considered expired
	ExpireTime time.Time `json:"expire_time" format:"date-time"`
	// Duration in seconds ending with 's'
	Ttl string `json:"ttl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentMessage respjson.Field
		Event          respjson.Field
		ExpireTime     respjson.Field
		Ttl            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessage) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsAgentMessage to a RcsAgentMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsAgentMessageParam.Overrides()
func (r RcsAgentMessage) ToParam() RcsAgentMessageParam {
	return param.Override[RcsAgentMessageParam](json.RawMessage(r.RawJSON()))
}

type RcsAgentMessageContentMessage struct {
	ContentInfo RcsContentInfo                        `json:"content_info"`
	RichCard    RcsAgentMessageContentMessageRichCard `json:"rich_card"`
	// List of suggested actions and replies
	Suggestions []RcsSuggestion `json:"suggestions"`
	// Text (maximum 3072 characters)
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentInfo respjson.Field
		RichCard    respjson.Field
		Suggestions respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessage) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageContentMessageRichCard struct {
	// Carousel of cards.
	CarouselCard RcsAgentMessageContentMessageRichCardCarouselCard `json:"carousel_card"`
	// Standalone card
	StandaloneCard RcsAgentMessageContentMessageRichCardStandaloneCard `json:"standalone_card"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CarouselCard   respjson.Field
		StandaloneCard respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessageRichCard) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessageRichCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Carousel of cards.
type RcsAgentMessageContentMessageRichCardCarouselCard struct {
	// The list of contents for each card in the carousel. A carousel can have a
	// minimum of 2 cards and a maximum 10 cards.
	CardContents []RcsCardContent `json:"card_contents" api:"required"`
	// The width of the cards in the carousel.
	//
	// Any of "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM".
	CardWidth string `json:"card_width" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardContents respjson.Field
		CardWidth    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessageRichCardCarouselCard) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessageRichCardCarouselCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Standalone card
type RcsAgentMessageContentMessageRichCardStandaloneCard struct {
	CardContent RcsCardContent `json:"card_content" api:"required"`
	// Orientation of the card.
	//
	// Any of "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL".
	CardOrientation string `json:"card_orientation" api:"required"`
	// Image preview alignment for standalone cards with horizontal layout.
	//
	// Any of "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT".
	ThumbnailImageAlignment string `json:"thumbnail_image_alignment" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardContent             respjson.Field
		CardOrientation         respjson.Field
		ThumbnailImageAlignment respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageContentMessageRichCardStandaloneCard) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageContentMessageRichCardStandaloneCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// RCS Event to send to the recipient
type RcsAgentMessageEvent struct {
	// Any of "TYPE_UNSPECIFIED", "IS_TYPING", "READ".
	EventType string `json:"event_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsAgentMessageEvent) RawJSON() string { return r.JSON.raw }
func (r *RcsAgentMessageEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageParam struct {
	// Timestamp in UTC of when this message is considered expired
	ExpireTime param.Opt[time.Time] `json:"expire_time,omitzero" format:"date-time"`
	// Duration in seconds ending with 's'
	Ttl            param.Opt[string]                  `json:"ttl,omitzero"`
	ContentMessage RcsAgentMessageContentMessageParam `json:"content_message,omitzero"`
	// RCS Event to send to the recipient
	Event RcsAgentMessageEventParam `json:"event,omitzero"`
	paramObj
}

func (r RcsAgentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageContentMessageParam struct {
	// Text (maximum 3072 characters)
	Text        param.Opt[string]                          `json:"text,omitzero"`
	ContentInfo RcsContentInfoParam                        `json:"content_info,omitzero"`
	RichCard    RcsAgentMessageContentMessageRichCardParam `json:"rich_card,omitzero"`
	// List of suggested actions and replies
	Suggestions []RcsSuggestionParam `json:"suggestions,omitzero"`
	paramObj
}

func (r RcsAgentMessageContentMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsAgentMessageContentMessageRichCardParam struct {
	// Carousel of cards.
	CarouselCard RcsAgentMessageContentMessageRichCardCarouselCardParam `json:"carousel_card,omitzero"`
	// Standalone card
	StandaloneCard RcsAgentMessageContentMessageRichCardStandaloneCardParam `json:"standalone_card,omitzero"`
	paramObj
}

func (r RcsAgentMessageContentMessageRichCardParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageRichCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageRichCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Carousel of cards.
//
// The properties CardContents, CardWidth are required.
type RcsAgentMessageContentMessageRichCardCarouselCardParam struct {
	// The list of contents for each card in the carousel. A carousel can have a
	// minimum of 2 cards and a maximum 10 cards.
	CardContents []RcsCardContentParam `json:"card_contents,omitzero" api:"required"`
	// The width of the cards in the carousel.
	//
	// Any of "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM".
	CardWidth string `json:"card_width,omitzero" api:"required"`
	paramObj
}

func (r RcsAgentMessageContentMessageRichCardCarouselCardParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageRichCardCarouselCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageRichCardCarouselCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsAgentMessageContentMessageRichCardCarouselCardParam](
		"card_width", "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM",
	)
}

// Standalone card
//
// The properties CardContent, CardOrientation, ThumbnailImageAlignment are
// required.
type RcsAgentMessageContentMessageRichCardStandaloneCardParam struct {
	CardContent RcsCardContentParam `json:"card_content,omitzero" api:"required"`
	// Orientation of the card.
	//
	// Any of "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL".
	CardOrientation string `json:"card_orientation,omitzero" api:"required"`
	// Image preview alignment for standalone cards with horizontal layout.
	//
	// Any of "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT".
	ThumbnailImageAlignment string `json:"thumbnail_image_alignment,omitzero" api:"required"`
	paramObj
}

func (r RcsAgentMessageContentMessageRichCardStandaloneCardParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageContentMessageRichCardStandaloneCardParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageContentMessageRichCardStandaloneCardParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsAgentMessageContentMessageRichCardStandaloneCardParam](
		"card_orientation", "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL",
	)
	apijson.RegisterFieldValidator[RcsAgentMessageContentMessageRichCardStandaloneCardParam](
		"thumbnail_image_alignment", "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT",
	)
}

// RCS Event to send to the recipient
type RcsAgentMessageEventParam struct {
	// Any of "TYPE_UNSPECIFIED", "IS_TYPING", "READ".
	EventType string `json:"event_type,omitzero"`
	paramObj
}

func (r RcsAgentMessageEventParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsAgentMessageEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsAgentMessageEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsAgentMessageEventParam](
		"event_type", "TYPE_UNSPECIFIED", "IS_TYPING", "READ",
	)
}

type RcsCardContent struct {
	// Description of the card (at most 2000 characters)
	Description string `json:"description"`
	// A media file within a rich card.
	Media RcsCardContentMedia `json:"media"`
	// List of suggestions to include in the card. Maximum 10 suggestions.
	Suggestions []RcsSuggestion `json:"suggestions"`
	// Title of the card (at most 200 characters)
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Media       respjson.Field
		Suggestions respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsCardContent) RawJSON() string { return r.JSON.raw }
func (r *RcsCardContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsCardContent to a RcsCardContentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsCardContentParam.Overrides()
func (r RcsCardContent) ToParam() RcsCardContentParam {
	return param.Override[RcsCardContentParam](json.RawMessage(r.RawJSON()))
}

// A media file within a rich card.
type RcsCardContentMedia struct {
	ContentInfo RcsContentInfo `json:"content_info"`
	// The height of the media within a rich card with a vertical layout. For a
	// standalone card with horizontal layout, height is not customizable, and this
	// field is ignored.
	//
	// Any of "HEIGHT_UNSPECIFIED", "SHORT", "MEDIUM", "TALL".
	Height string `json:"height"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentInfo respjson.Field
		Height      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsCardContentMedia) RawJSON() string { return r.JSON.raw }
func (r *RcsCardContentMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsCardContentParam struct {
	// Description of the card (at most 2000 characters)
	Description param.Opt[string] `json:"description,omitzero"`
	// Title of the card (at most 200 characters)
	Title param.Opt[string] `json:"title,omitzero"`
	// A media file within a rich card.
	Media RcsCardContentMediaParam `json:"media,omitzero"`
	// List of suggestions to include in the card. Maximum 10 suggestions.
	Suggestions []RcsSuggestionParam `json:"suggestions,omitzero"`
	paramObj
}

func (r RcsCardContentParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsCardContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsCardContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A media file within a rich card.
type RcsCardContentMediaParam struct {
	ContentInfo RcsContentInfoParam `json:"content_info,omitzero"`
	// The height of the media within a rich card with a vertical layout. For a
	// standalone card with horizontal layout, height is not customizable, and this
	// field is ignored.
	//
	// Any of "HEIGHT_UNSPECIFIED", "SHORT", "MEDIUM", "TALL".
	Height string `json:"height,omitzero"`
	paramObj
}

func (r RcsCardContentMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsCardContentMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsCardContentMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsCardContentMediaParam](
		"height", "HEIGHT_UNSPECIFIED", "SHORT", "MEDIUM", "TALL",
	)
}

type RcsContentInfo struct {
	// Publicly reachable URL of the file.
	FileURL string `json:"file_url" api:"required" format:"url"`
	// If set the URL content will not be cached.
	ForceRefresh bool `json:"force_refresh"`
	// Publicly reachable URL of the thumbnail. Maximum size of 100 kB.
	ThumbnailURL string `json:"thumbnail_url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileURL      respjson.Field
		ForceRefresh respjson.Field
		ThumbnailURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsContentInfo) RawJSON() string { return r.JSON.raw }
func (r *RcsContentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsContentInfo to a RcsContentInfoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsContentInfoParam.Overrides()
func (r RcsContentInfo) ToParam() RcsContentInfoParam {
	return param.Override[RcsContentInfoParam](json.RawMessage(r.RawJSON()))
}

// The property FileURL is required.
type RcsContentInfoParam struct {
	// Publicly reachable URL of the file.
	FileURL string `json:"file_url" api:"required" format:"url"`
	// If set the URL content will not be cached.
	ForceRefresh param.Opt[bool] `json:"force_refresh,omitzero"`
	// Publicly reachable URL of the thumbnail. Maximum size of 100 kB.
	ThumbnailURL param.Opt[string] `json:"thumbnail_url,omitzero" format:"url"`
	paramObj
}

func (r RcsContentInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsContentInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsContentInfoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestion struct {
	// When tapped, initiates the corresponding native action on the device.
	Action RcsSuggestionAction `json:"action"`
	Reply  RcsSuggestionReply  `json:"reply"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Reply       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestion) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RcsSuggestion to a RcsSuggestionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RcsSuggestionParam.Overrides()
func (r RcsSuggestion) ToParam() RcsSuggestionParam {
	return param.Override[RcsSuggestionParam](json.RawMessage(r.RawJSON()))
}

// When tapped, initiates the corresponding native action on the device.
type RcsSuggestionAction struct {
	// Opens the user's default calendar app and starts the new calendar event flow
	// with the agent-specified event data pre-filled.
	CreateCalendarEventAction RcsSuggestionActionCreateCalendarEventAction `json:"create_calendar_event_action"`
	// Opens the user's default dialer app with the agent-specified phone number filled
	// in.
	DialAction RcsSuggestionActionDialAction `json:"dial_action"`
	// Fallback URL to use if a client doesn't support a suggested action. Fallback
	// URLs open in new browser windows. Maximum 2048 characters.
	FallbackURL string `json:"fallback_url" format:"url"`
	// Opens the user's default web browser app to the specified URL.
	OpenURLAction RcsSuggestionActionOpenURLAction `json:"open_url_action"`
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData string `json:"postback_data"`
	// Opens the RCS app's location chooser so the user can pick a location to send
	// back to the agent.
	ShareLocationAction map[string]any `json:"share_location_action"`
	// Text that is shown in the suggested action. Maximum 25 characters.
	Text string `json:"text"`
	// Opens the user's default map app and selects the agent-specified location.
	ViewLocationAction RcsSuggestionActionViewLocationAction `json:"view_location_action"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateCalendarEventAction respjson.Field
		DialAction                respjson.Field
		FallbackURL               respjson.Field
		OpenURLAction             respjson.Field
		PostbackData              respjson.Field
		ShareLocationAction       respjson.Field
		Text                      respjson.Field
		ViewLocationAction        respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default calendar app and starts the new calendar event flow
// with the agent-specified event data pre-filled.
type RcsSuggestionActionCreateCalendarEventAction struct {
	// Event description. Maximum 500 characters.
	Description string    `json:"description"`
	EndTime     time.Time `json:"end_time" format:"date-time"`
	StartTime   time.Time `json:"start_time" format:"date-time"`
	// Event title. Maximum 100 characters.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		EndTime     respjson.Field
		StartTime   respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionCreateCalendarEventAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionCreateCalendarEventAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default dialer app with the agent-specified phone number filled
// in.
type RcsSuggestionActionDialAction struct {
	// Phone number in +E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionDialAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionDialAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default web browser app to the specified URL.
type RcsSuggestionActionOpenURLAction struct {
	// URL open application, browser or webview.
	//
	// Any of "OPEN_URL_APPLICATION_UNSPECIFIED", "BROWSER", "WEBVIEW".
	Application string `json:"application" api:"required"`
	URL         string `json:"url" api:"required" format:"url"`
	// Any of "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL".
	WebviewViewMode string `json:"webview_view_mode" api:"required"`
	// Accessbility description for webview.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Application     respjson.Field
		URL             respjson.Field
		WebviewViewMode respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionOpenURLAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionOpenURLAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default map app and selects the agent-specified location.
type RcsSuggestionActionViewLocationAction struct {
	// The label of the pin dropped
	Label   string                                       `json:"label"`
	LatLong RcsSuggestionActionViewLocationActionLatLong `json:"lat_long"`
	// query string (Android only)
	Query string `json:"query"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		LatLong     respjson.Field
		Query       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionViewLocationAction) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionViewLocationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionActionViewLocationActionLatLong struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `json:"latitude" api:"required"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionActionViewLocationActionLatLong) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionActionViewLocationActionLatLong) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionReply struct {
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData string `json:"postback_data"`
	// Text that is shown in the suggested reply (maximum 25 characters)
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PostbackData respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RcsSuggestionReply) RawJSON() string { return r.JSON.raw }
func (r *RcsSuggestionReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionParam struct {
	// When tapped, initiates the corresponding native action on the device.
	Action RcsSuggestionActionParam `json:"action,omitzero"`
	Reply  RcsSuggestionReplyParam  `json:"reply,omitzero"`
	paramObj
}

func (r RcsSuggestionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When tapped, initiates the corresponding native action on the device.
type RcsSuggestionActionParam struct {
	// Fallback URL to use if a client doesn't support a suggested action. Fallback
	// URLs open in new browser windows. Maximum 2048 characters.
	FallbackURL param.Opt[string] `json:"fallback_url,omitzero" format:"url"`
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData param.Opt[string] `json:"postback_data,omitzero"`
	// Text that is shown in the suggested action. Maximum 25 characters.
	Text param.Opt[string] `json:"text,omitzero"`
	// Opens the user's default calendar app and starts the new calendar event flow
	// with the agent-specified event data pre-filled.
	CreateCalendarEventAction RcsSuggestionActionCreateCalendarEventActionParam `json:"create_calendar_event_action,omitzero"`
	// Opens the user's default dialer app with the agent-specified phone number filled
	// in.
	DialAction RcsSuggestionActionDialActionParam `json:"dial_action,omitzero"`
	// Opens the user's default web browser app to the specified URL.
	OpenURLAction RcsSuggestionActionOpenURLActionParam `json:"open_url_action,omitzero"`
	// Opens the RCS app's location chooser so the user can pick a location to send
	// back to the agent.
	ShareLocationAction map[string]any `json:"share_location_action,omitzero"`
	// Opens the user's default map app and selects the agent-specified location.
	ViewLocationAction RcsSuggestionActionViewLocationActionParam `json:"view_location_action,omitzero"`
	paramObj
}

func (r RcsSuggestionActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default calendar app and starts the new calendar event flow
// with the agent-specified event data pre-filled.
type RcsSuggestionActionCreateCalendarEventActionParam struct {
	// Event description. Maximum 500 characters.
	Description param.Opt[string]    `json:"description,omitzero"`
	EndTime     param.Opt[time.Time] `json:"end_time,omitzero" format:"date-time"`
	StartTime   param.Opt[time.Time] `json:"start_time,omitzero" format:"date-time"`
	// Event title. Maximum 100 characters.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r RcsSuggestionActionCreateCalendarEventActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionCreateCalendarEventActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionCreateCalendarEventActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default dialer app with the agent-specified phone number filled
// in.
//
// The property PhoneNumber is required.
type RcsSuggestionActionDialActionParam struct {
	// Phone number in +E.164 format
	PhoneNumber string `json:"phone_number" api:"required"`
	paramObj
}

func (r RcsSuggestionActionDialActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionDialActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionDialActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opens the user's default web browser app to the specified URL.
//
// The properties Application, URL, WebviewViewMode are required.
type RcsSuggestionActionOpenURLActionParam struct {
	// URL open application, browser or webview.
	//
	// Any of "OPEN_URL_APPLICATION_UNSPECIFIED", "BROWSER", "WEBVIEW".
	Application string `json:"application,omitzero" api:"required"`
	URL         string `json:"url" api:"required" format:"url"`
	// Any of "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL".
	WebviewViewMode string `json:"webview_view_mode,omitzero" api:"required"`
	// Accessbility description for webview.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r RcsSuggestionActionOpenURLActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionOpenURLActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionOpenURLActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RcsSuggestionActionOpenURLActionParam](
		"application", "OPEN_URL_APPLICATION_UNSPECIFIED", "BROWSER", "WEBVIEW",
	)
	apijson.RegisterFieldValidator[RcsSuggestionActionOpenURLActionParam](
		"webview_view_mode", "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL",
	)
}

// Opens the user's default map app and selects the agent-specified location.
type RcsSuggestionActionViewLocationActionParam struct {
	// The label of the pin dropped
	Label param.Opt[string] `json:"label,omitzero"`
	// query string (Android only)
	Query   param.Opt[string]                                 `json:"query,omitzero"`
	LatLong RcsSuggestionActionViewLocationActionLatLongParam `json:"lat_long,omitzero"`
	paramObj
}

func (r RcsSuggestionActionViewLocationActionParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionViewLocationActionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionViewLocationActionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Latitude, Longitude are required.
type RcsSuggestionActionViewLocationActionLatLongParam struct {
	// The latitude in degrees. It must be in the range [-90.0, +90.0].
	Latitude float64 `json:"latitude" api:"required"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude" api:"required"`
	paramObj
}

func (r RcsSuggestionActionViewLocationActionLatLongParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionActionViewLocationActionLatLongParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionActionViewLocationActionLatLongParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsSuggestionReplyParam struct {
	// Payload (base64 encoded) that will be sent to the agent in the user event that
	// results when the user taps the suggested action. Maximum 2048 characters.
	PostbackData param.Opt[string] `json:"postback_data,omitzero"`
	// Text that is shown in the suggested reply (maximum 25 characters)
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r RcsSuggestionReplyParam) MarshalJSON() (data []byte, err error) {
	type shadow RcsSuggestionReplyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcsSuggestionReplyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcsToItem struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
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
func (r RcsToItem) RawJSON() string { return r.JSON.raw }
func (r *RcsToItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	Direction string                  `json:"direction"`
	Encoding  string                  `json:"encoding"`
	Errors    []shared.MessagingError `json:"errors"`
	// This field is a union of [OutboundMessagePayloadFrom],
	// [shared.InboundMessagePayloadFrom]
	From MessageGetResponseDataUnionFrom `json:"from"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media              MessageGetResponseDataUnionMedia `json:"media"`
	MessagingProfileID string                           `json:"messaging_profile_id"`
	NumChars           int64                            `json:"num_chars"`
	OrganizationID     string                           `json:"organization_id"`
	Parts              int64                            `json:"parts"`
	ReceivedAt         time.Time                        `json:"received_at"`
	RecordType         string                           `json:"record_type"`
	SentAt             time.Time                        `json:"sent_at"`
	// This field is from variant [OutboundMessagePayload].
	SmartEncodingApplied  bool     `json:"smart_encoding_applied"`
	Subject               string   `json:"subject"`
	Tags                  []string `json:"tags"`
	TcrCampaignBillable   bool     `json:"tcr_campaign_billable"`
	TcrCampaignID         string   `json:"tcr_campaign_id"`
	TcrCampaignRegistered string   `json:"tcr_campaign_registered"`
	Text                  string   `json:"text"`
	// This field is a union of [[]OutboundMessagePayloadTo],
	// [[]shared.InboundMessagePayloadTo]
	To         MessageGetResponseDataUnionTo `json:"to"`
	Type       string                        `json:"type"`
	ValidUntil time.Time                     `json:"valid_until"`
	// This field is from variant [OutboundMessagePayload].
	WaitSeconds        float64 `json:"wait_seconds"`
	WebhookFailoverURL string  `json:"webhook_failover_url"`
	WebhookURL         string  `json:"webhook_url"`
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
		NumChars              respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		SmartEncodingApplied  respjson.Field
		Subject               respjson.Field
		Tags                  respjson.Field
		TcrCampaignBillable   respjson.Field
		TcrCampaignID         respjson.Field
		TcrCampaignRegistered respjson.Field
		Text                  respjson.Field
		To                    respjson.Field
		Type                  respjson.Field
		ValidUntil            respjson.Field
		WaitSeconds           respjson.Field
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
	CompletedAt time.Time                          `json:"completed_at" api:"nullable" format:"date-time"`
	Cost        MessageCancelScheduledResponseCost `json:"cost" api:"nullable"`
	// Detailed breakdown of the message cost components.
	CostBreakdown MessageCancelScheduledResponseCostBreakdown `json:"cost_breakdown" api:"nullable"`
	// The direction of the message. Inbound messages are sent to you whereas outbound
	// messages are sent from you.
	//
	// Any of "outbound".
	Direction MessageCancelScheduledResponseDirection `json:"direction"`
	// Encoding scheme used for the message body.
	Encoding string `json:"encoding"`
	// These errors may point at addressees when referring to unsuccessful/unconfirmed
	// delivery statuses.
	Errors []shared.MessagingError               `json:"errors"`
	From   MessageCancelScheduledResponseFrom    `json:"from"`
	Media  []MessageCancelScheduledResponseMedia `json:"media"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id"`
	// The number of characters in the message text
	NumChars int64 `json:"num_chars"`
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
	SentAt time.Time `json:"sent_at" api:"nullable" format:"date-time"`
	// Indicates whether smart encoding was applied to this message. When `true`, one
	// or more Unicode characters were automatically replaced with GSM-7 equivalents to
	// reduce segment count and cost. The original message text is preserved in
	// webhooks.
	SmartEncodingApplied bool `json:"smart_encoding_applied"`
	// Subject of multimedia message
	Subject string `json:"subject" api:"nullable"`
	// Tags associated with the resource.
	Tags []string `json:"tags"`
	// Indicates whether the TCR campaign is billable.
	TcrCampaignBillable bool `json:"tcr_campaign_billable"`
	// The Campaign Registry (TCR) campaign ID associated with the message.
	TcrCampaignID string `json:"tcr_campaign_id" api:"nullable"`
	// The registration status of the TCR campaign.
	TcrCampaignRegistered string `json:"tcr_campaign_registered" api:"nullable"`
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
	ValidUntil time.Time `json:"valid_until" api:"nullable" format:"date-time"`
	// The failover URL where webhooks related to this message will be sent if sending
	// to the primary URL fails.
	WebhookFailoverURL string `json:"webhook_failover_url" api:"nullable" format:"url"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL string `json:"webhook_url" api:"nullable" format:"url"`
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
		NumChars              respjson.Field
		OrganizationID        respjson.Field
		Parts                 respjson.Field
		ReceivedAt            respjson.Field
		RecordType            respjson.Field
		SentAt                respjson.Field
		SmartEncodingApplied  respjson.Field
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
	PhoneNumber string `json:"phone_number"`
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
	Amount string `json:"amount"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
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
	Amount string `json:"amount"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
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
	Amount string `json:"amount"`
	// The ISO 4217 currency identifier.
	Currency string `json:"currency"`
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
	PhoneNumber string `json:"phone_number"`
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
	ContentType string `json:"content_type" api:"nullable"`
	// The SHA256 hash of the requested media.
	Sha256 string `json:"sha256" api:"nullable"`
	// The size of the requested media.
	Size int64 `json:"size" api:"nullable"`
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
	PhoneNumber string `json:"phone_number"`
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

type MessageGetGroupMessagesResponse struct {
	Data []OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageGetGroupMessagesResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageGetGroupMessagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type MessageSendWithAlphanumericSenderResponse struct {
	Data OutboundMessagePayload `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWithAlphanumericSenderResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWithAlphanumericSenderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageScheduleParams struct {
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to" api:"required"`
	// Automatically detect if an SMS message is unusually long and exceeds a
	// recommended limit of message parts.
	AutoDetect param.Opt[bool] `json:"auto_detect,omitzero"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	//
	// **Required if sending with a phone number, short code, or alphanumeric sender
	// ID.**
	From param.Opt[string] `json:"from,omitzero"`
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
	To string `json:"to" api:"required"`
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
	From param.Opt[string] `json:"from,omitzero"`
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
	// Encoding to use for the message. `auto` (default) uses smart encoding to
	// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
	// (returns 400 if message contains characters that cannot be encoded). `ucs2`
	// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
	// messaging profile's `smart_encoding` setting.
	//
	// Any of "auto", "gsm7", "ucs2".
	Encoding MessageSendParamsEncoding `json:"encoding,omitzero"`
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

// Encoding to use for the message. `auto` (default) uses smart encoding to
// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
// (returns 400 if message contains characters that cannot be encoded). `ucs2`
// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
// messaging profile's `smart_encoding` setting.
type MessageSendParamsEncoding string

const (
	MessageSendParamsEncodingAuto MessageSendParamsEncoding = "auto"
	MessageSendParamsEncodingGsm7 MessageSendParamsEncoding = "gsm7"
	MessageSendParamsEncodingUcs2 MessageSendParamsEncoding = "ucs2"
)

// The protocol for sending the message, either SMS or MMS.
type MessageSendParamsType string

const (
	MessageSendParamsTypeSMS MessageSendParamsType = "SMS"
	MessageSendParamsTypeMms MessageSendParamsType = "MMS"
)

type MessageSendGroupMmsParams struct {
	// Phone number, in +E.164 format, used to send the message.
	From string `json:"from" api:"required"`
	// A list of destinations. No more than 8 destinations are allowed.
	To []string `json:"to,omitzero" api:"required"`
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
	From string `json:"from" api:"required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to" api:"required"`
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
	// Encoding to use for the message. `auto` (default) uses smart encoding to
	// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
	// (returns 400 if message contains characters that cannot be encoded). `ucs2`
	// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
	// messaging profile's `smart_encoding` setting.
	//
	// Any of "auto", "gsm7", "ucs2".
	Encoding MessageSendLongCodeParamsEncoding `json:"encoding,omitzero"`
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

// Encoding to use for the message. `auto` (default) uses smart encoding to
// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
// (returns 400 if message contains characters that cannot be encoded). `ucs2`
// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
// messaging profile's `smart_encoding` setting.
type MessageSendLongCodeParamsEncoding string

const (
	MessageSendLongCodeParamsEncodingAuto MessageSendLongCodeParamsEncoding = "auto"
	MessageSendLongCodeParamsEncodingGsm7 MessageSendLongCodeParamsEncoding = "gsm7"
	MessageSendLongCodeParamsEncodingUcs2 MessageSendLongCodeParamsEncoding = "ucs2"
)

// The protocol for sending the message, either SMS or MMS.
type MessageSendLongCodeParamsType string

const (
	MessageSendLongCodeParamsTypeSMS MessageSendLongCodeParamsType = "SMS"
	MessageSendLongCodeParamsTypeMms MessageSendLongCodeParamsType = "MMS"
)

type MessageSendNumberPoolParams struct {
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id" api:"required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to" api:"required"`
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
	// Encoding to use for the message. `auto` (default) uses smart encoding to
	// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
	// (returns 400 if message contains characters that cannot be encoded). `ucs2`
	// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
	// messaging profile's `smart_encoding` setting.
	//
	// Any of "auto", "gsm7", "ucs2".
	Encoding MessageSendNumberPoolParamsEncoding `json:"encoding,omitzero"`
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

// Encoding to use for the message. `auto` (default) uses smart encoding to
// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
// (returns 400 if message contains characters that cannot be encoded). `ucs2`
// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
// messaging profile's `smart_encoding` setting.
type MessageSendNumberPoolParamsEncoding string

const (
	MessageSendNumberPoolParamsEncodingAuto MessageSendNumberPoolParamsEncoding = "auto"
	MessageSendNumberPoolParamsEncodingGsm7 MessageSendNumberPoolParamsEncoding = "gsm7"
	MessageSendNumberPoolParamsEncodingUcs2 MessageSendNumberPoolParamsEncoding = "ucs2"
)

// The protocol for sending the message, either SMS or MMS.
type MessageSendNumberPoolParamsType string

const (
	MessageSendNumberPoolParamsTypeSMS MessageSendNumberPoolParamsType = "SMS"
	MessageSendNumberPoolParamsTypeMms MessageSendNumberPoolParamsType = "MMS"
)

type MessageSendShortCodeParams struct {
	// Phone number, in +E.164 format, used to send the message.
	From string `json:"from" api:"required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to" api:"required"`
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
	// Encoding to use for the message. `auto` (default) uses smart encoding to
	// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
	// (returns 400 if message contains characters that cannot be encoded). `ucs2`
	// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
	// messaging profile's `smart_encoding` setting.
	//
	// Any of "auto", "gsm7", "ucs2".
	Encoding MessageSendShortCodeParamsEncoding `json:"encoding,omitzero"`
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

// Encoding to use for the message. `auto` (default) uses smart encoding to
// automatically select the most efficient encoding. `gsm7` forces GSM-7 encoding
// (returns 400 if message contains characters that cannot be encoded). `ucs2`
// forces UCS-2 encoding and disables smart encoding. When set, this overrides the
// messaging profile's `smart_encoding` setting.
type MessageSendShortCodeParamsEncoding string

const (
	MessageSendShortCodeParamsEncodingAuto MessageSendShortCodeParamsEncoding = "auto"
	MessageSendShortCodeParamsEncodingGsm7 MessageSendShortCodeParamsEncoding = "gsm7"
	MessageSendShortCodeParamsEncodingUcs2 MessageSendShortCodeParamsEncoding = "ucs2"
)

// The protocol for sending the message, either SMS or MMS.
type MessageSendShortCodeParamsType string

const (
	MessageSendShortCodeParamsTypeSMS MessageSendShortCodeParamsType = "SMS"
	MessageSendShortCodeParamsTypeMms MessageSendShortCodeParamsType = "MMS"
)

type MessageSendWithAlphanumericSenderParams struct {
	// A valid alphanumeric sender ID on the user's account.
	From string `json:"from" api:"required"`
	// The messaging profile ID to use.
	MessagingProfileID string `json:"messaging_profile_id" api:"required" format:"uuid"`
	// The message body.
	Text string `json:"text" api:"required"`
	// Receiving address (+E.164 formatted phone number).
	To string `json:"to" api:"required"`
	// Failover callback URL for delivery status updates.
	WebhookFailoverURL param.Opt[string] `json:"webhook_failover_url,omitzero" format:"url"`
	// Callback URL for delivery status updates.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// If true, use the messaging profile's webhook settings.
	UseProfileWebhooks param.Opt[bool] `json:"use_profile_webhooks,omitzero"`
	paramObj
}

func (r MessageSendWithAlphanumericSenderParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWithAlphanumericSenderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWithAlphanumericSenderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
