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
// [MDR report.](https://developers.telnyx.com/api-reference/mdr-usage-reports/create-mdr-usage-report)
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

// Send a Whatsapp message
func (r *MessageService) SendWhatsapp(ctx context.Context, body MessageSendWhatsappParams, opts ...option.RequestOption) (res *MessageSendWhatsappResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "messages/whatsapp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessagingError struct {
	Code   string               `json:"code,required"`
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
	// Indicates whether smart encoding was applied to this message. When `true`, one
	// or more Unicode characters were automatically replaced with GSM-7 equivalents to
	// reduce segment count and cost. The original message text is preserved in
	// webhooks.
	SmartEncodingApplied bool `json:"smart_encoding_applied"`
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
	ContentType string `json:"content_type,nullable"`
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
	CardContents []RcsCardContent `json:"card_contents,required"`
	// The width of the cards in the carousel.
	//
	// Any of "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM".
	CardWidth string `json:"card_width,required"`
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
	CardContent RcsCardContent `json:"card_content,required"`
	// Orientation of the card.
	//
	// Any of "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL".
	CardOrientation string `json:"card_orientation,required"`
	// Image preview alignment for standalone cards with horizontal layout.
	//
	// Any of "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT".
	ThumbnailImageAlignment string `json:"thumbnail_image_alignment,required"`
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
	CardContents []RcsCardContentParam `json:"card_contents,omitzero,required"`
	// The width of the cards in the carousel.
	//
	// Any of "CARD_WIDTH_UNSPECIFIED", "SMALL", "MEDIUM".
	CardWidth string `json:"card_width,omitzero,required"`
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
	CardContent RcsCardContentParam `json:"card_content,omitzero,required"`
	// Orientation of the card.
	//
	// Any of "CARD_ORIENTATION_UNSPECIFIED", "HORIZONTAL", "VERTICAL".
	CardOrientation string `json:"card_orientation,omitzero,required"`
	// Image preview alignment for standalone cards with horizontal layout.
	//
	// Any of "THUMBNAIL_IMAGE_ALIGNMENT_UNSPECIFIED", "LEFT", "RIGHT".
	ThumbnailImageAlignment string `json:"thumbnail_image_alignment,omitzero,required"`
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
	FileURL string `json:"file_url,required" format:"url"`
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
	FileURL string `json:"file_url,required" format:"url"`
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
	PhoneNumber string `json:"phone_number,required"`
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
	Application string `json:"application,required"`
	URL         string `json:"url,required" format:"url"`
	// Any of "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL".
	WebviewViewMode string `json:"webview_view_mode,required"`
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
	Latitude float64 `json:"latitude,required"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude,required"`
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
	PhoneNumber string `json:"phone_number,required"`
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
	Application string `json:"application,omitzero,required"`
	URL         string `json:"url,required" format:"url"`
	// Any of "WEBVIEW_VIEW_MODE_UNSPECIFIED", "FULL", "HALF", "TALL".
	WebviewViewMode string `json:"webview_view_mode,omitzero,required"`
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
	Latitude float64 `json:"latitude,required"`
	// The longitude in degrees. It must be in the range [-180.0, +180.0].
	Longitude float64 `json:"longitude,required"`
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

type WhatsappMedia struct {
	// media caption
	Caption string `json:"caption"`
	// file name with extension
	Filename string `json:"filename"`
	// media URL
	Link string `json:"link" format:"url"`
	// true if voice message
	Voice bool `json:"voice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Filename    respjson.Field
		Link        respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappMedia) RawJSON() string { return r.JSON.raw }
func (r *WhatsappMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WhatsappMedia to a WhatsappMediaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WhatsappMediaParam.Overrides()
func (r WhatsappMedia) ToParam() WhatsappMediaParam {
	return param.Override[WhatsappMediaParam](json.RawMessage(r.RawJSON()))
}

type WhatsappMediaParam struct {
	// media caption
	Caption param.Opt[string] `json:"caption,omitzero"`
	// file name with extension
	Filename param.Opt[string] `json:"filename,omitzero"`
	// media URL
	Link param.Opt[string] `json:"link,omitzero" format:"url"`
	// true if voice message
	Voice param.Opt[bool] `json:"voice,omitzero"`
	paramObj
}

func (r WhatsappMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappMediaParam) UnmarshalJSON(data []byte) error {
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
	Direction string           `json:"direction"`
	Encoding  string           `json:"encoding"`
	Errors    []MessagingError `json:"errors"`
	// This field is a union of [OutboundMessagePayloadFrom],
	// [shared.InboundMessagePayloadFrom]
	From MessageGetResponseDataUnionFrom `json:"from"`
	// This field is a union of [[]OutboundMessagePayloadMedia],
	// [[]shared.InboundMessagePayloadMedia]
	Media              MessageGetResponseDataUnionMedia `json:"media"`
	MessagingProfileID string                           `json:"messaging_profile_id"`
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
	// Indicates whether smart encoding was applied to this message. When `true`, one
	// or more Unicode characters were automatically replaced with GSM-7 equivalents to
	// reduce segment count and cost. The original message text is preserved in
	// webhooks.
	SmartEncodingApplied bool `json:"smart_encoding_applied"`
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
	ContentType string `json:"content_type,nullable"`
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

type MessageSendWhatsappResponse struct {
	Data MessageSendWhatsappResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponse) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseData struct {
	// message ID
	ID                 string                              `json:"id"`
	Body               MessageSendWhatsappResponseDataBody `json:"body"`
	Direction          string                              `json:"direction"`
	Encoding           string                              `json:"encoding"`
	From               MessageSendWhatsappResponseDataFrom `json:"from"`
	MessagingProfileID string                              `json:"messaging_profile_id"`
	OrganizationID     string                              `json:"organization_id"`
	ReceivedAt         time.Time                           `json:"received_at" format:"date-time"`
	RecordType         string                              `json:"record_type"`
	To                 []MessageSendWhatsappResponseDataTo `json:"to"`
	Type               string                              `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Body               respjson.Field
		Direction          respjson.Field
		Encoding           respjson.Field
		From               respjson.Field
		MessagingProfileID respjson.Field
		OrganizationID     respjson.Field
		ReceivedAt         respjson.Field
		RecordType         respjson.Field
		To                 respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBody struct {
	Audio WhatsappMedia `json:"audio"`
	// custom data to return with status update
	BizOpaqueCallbackData string                                         `json:"biz_opaque_callback_data"`
	Contacts              []MessageSendWhatsappResponseDataBodyContact   `json:"contacts"`
	Document              WhatsappMedia                                  `json:"document"`
	Image                 WhatsappMedia                                  `json:"image"`
	Interactive           MessageSendWhatsappResponseDataBodyInteractive `json:"interactive"`
	Location              MessageSendWhatsappResponseDataBodyLocation    `json:"location"`
	Reaction              MessageSendWhatsappResponseDataBodyReaction    `json:"reaction"`
	Sticker               WhatsappMedia                                  `json:"sticker"`
	// Any of "audio", "document", "image", "sticker", "video", "interactive",
	// "location", "template", "reaction", "contacts".
	Type  string        `json:"type"`
	Video WhatsappMedia `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio                 respjson.Field
		BizOpaqueCallbackData respjson.Field
		Contacts              respjson.Field
		Document              respjson.Field
		Image                 respjson.Field
		Interactive           respjson.Field
		Location              respjson.Field
		Reaction              respjson.Field
		Sticker               respjson.Field
		Type                  respjson.Field
		Video                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBody) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyContact struct {
	Addresses []MessageSendWhatsappResponseDataBodyContactAddress `json:"addresses"`
	Birthday  string                                              `json:"birthday"`
	Emails    []MessageSendWhatsappResponseDataBodyContactEmail   `json:"emails"`
	Name      string                                              `json:"name"`
	Org       MessageSendWhatsappResponseDataBodyContactOrg       `json:"org"`
	Phones    []MessageSendWhatsappResponseDataBodyContactPhone   `json:"phones"`
	URLs      []MessageSendWhatsappResponseDataBodyContactURL     `json:"urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Addresses   respjson.Field
		Birthday    respjson.Field
		Emails      respjson.Field
		Name        respjson.Field
		Org         respjson.Field
		Phones      respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyContact) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyContactAddress struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	Street      string `json:"street"`
	Type        string `json:"type"`
	Zip         string `json:"zip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		State       respjson.Field
		Street      respjson.Field
		Type        respjson.Field
		Zip         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyContactAddress) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyContactAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyContactEmail struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyContactEmail) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyContactEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyContactOrg struct {
	Company    string `json:"company"`
	Department string `json:"department"`
	Title      string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Department  respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyContactOrg) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyContactOrg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyContactPhone struct {
	Phone string `json:"phone"`
	Type  string `json:"type"`
	WaID  string `json:"wa_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Phone       respjson.Field
		Type        respjson.Field
		WaID        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyContactPhone) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyContactPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyContactURL struct {
	Type string `json:"type"`
	URL  string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyContactURL) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyContactURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractive struct {
	Action MessageSendWhatsappResponseDataBodyInteractiveAction `json:"action"`
	Body   MessageSendWhatsappResponseDataBodyInteractiveBody   `json:"body"`
	Footer MessageSendWhatsappResponseDataBodyInteractiveFooter `json:"footer"`
	Header MessageSendWhatsappResponseDataBodyInteractiveHeader `json:"header"`
	// Any of "cta_url", "list", "carousel", "button", "location_request_message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Body        respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractive) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyInteractive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveAction struct {
	Button            string                                                         `json:"button"`
	Buttons           []MessageSendWhatsappResponseDataBodyInteractiveActionButton   `json:"buttons"`
	Cards             []MessageSendWhatsappResponseDataBodyInteractiveActionCard     `json:"cards"`
	CatalogID         string                                                         `json:"catalog_id"`
	Mode              string                                                         `json:"mode"`
	Name              string                                                         `json:"name"`
	Parameters        MessageSendWhatsappResponseDataBodyInteractiveActionParameters `json:"parameters"`
	ProductRetailerID string                                                         `json:"product_retailer_id"`
	Sections          []MessageSendWhatsappResponseDataBodyInteractiveActionSection  `json:"sections"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Button            respjson.Field
		Buttons           respjson.Field
		Cards             respjson.Field
		CatalogID         respjson.Field
		Mode              respjson.Field
		Name              respjson.Field
		Parameters        respjson.Field
		ProductRetailerID respjson.Field
		Sections          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveAction) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyInteractiveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionButton struct {
	Reply MessageSendWhatsappResponseDataBodyInteractiveActionButtonReply `json:"reply"`
	// Any of "reply".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reply       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionButton) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionButtonReply struct {
	// unique identifier for each button, 256 character maximum
	ID string `json:"id"`
	// button label, 20 character maximum
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionButtonReply) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionButtonReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionCard struct {
	Action MessageSendWhatsappResponseDataBodyInteractiveActionCardAction `json:"action"`
	Body   MessageSendWhatsappResponseDataBodyInteractiveActionCardBody   `json:"body"`
	// unique index for each card (0-9)
	CardIndex int64                                                          `json:"card_index"`
	Header    MessageSendWhatsappResponseDataBodyInteractiveActionCardHeader `json:"header"`
	// Any of "cta_url".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Body        respjson.Field
		CardIndex   respjson.Field
		Header      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionCard) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionCardAction struct {
	// the unique ID of the catalog
	CatalogID string `json:"catalog_id"`
	// the unique retailer ID of the product
	ProductRetailerID string `json:"product_retailer_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CatalogID         respjson.Field
		ProductRetailerID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionCardAction) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionCardAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionCardBody struct {
	// 160 character maximum, up to 2 line breaks
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionCardBody) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionCardBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionCardHeader struct {
	Image WhatsappMedia `json:"image"`
	// Any of "image", "video".
	Type  string        `json:"type"`
	Video WhatsappMedia `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionCardHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionCardHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionParameters struct {
	// button label text, 20 character maximum
	DisplayText string `json:"display_text"`
	// button URL to load when tapped by the user
	URL string `json:"url" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayText respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionParameters) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionSection struct {
	ProductItems []MessageSendWhatsappResponseDataBodyInteractiveActionSectionProductItem `json:"product_items"`
	Rows         []MessageSendWhatsappResponseDataBodyInteractiveActionSectionRow         `json:"rows"`
	// section title, 24 character maximum
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductItems respjson.Field
		Rows         respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionSection) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionSectionProductItem struct {
	ProductRetailerID string `json:"product_retailer_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductRetailerID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionSectionProductItem) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionSectionProductItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveActionSectionRow struct {
	// arbitrary string identifying the row, 200 character maximum
	ID string `json:"id"`
	// row description, 72 character maximum
	Description string `json:"description"`
	// row title, 24 character maximum
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveActionSectionRow) RawJSON() string {
	return r.JSON.raw
}
func (r *MessageSendWhatsappResponseDataBodyInteractiveActionSectionRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveBody struct {
	// body text, 1024 character maximum
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveBody) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyInteractiveBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveFooter struct {
	// footer text, 60 character maximum
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveFooter) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyInteractiveFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyInteractiveHeader struct {
	Document WhatsappMedia `json:"document"`
	Image    WhatsappMedia `json:"image"`
	SubText  string        `json:"sub_text"`
	// header text, 60 character maximum
	Text  string        `json:"text"`
	Video WhatsappMedia `json:"video"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		Image       respjson.Field
		SubText     respjson.Field
		Text        respjson.Field
		Video       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyInteractiveHeader) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyInteractiveHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyLocation struct {
	Address   string `json:"address"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Latitude    respjson.Field
		Longitude   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyLocation) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataBodyReaction struct {
	Emoji     string `json:"emoji"`
	MessageID string `json:"message_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emoji       respjson.Field
		MessageID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageSendWhatsappResponseDataBodyReaction) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataBodyReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataFrom struct {
	// The carrier of the sender.
	Carrier string `json:"carrier"`
	// The line-type of the sender.
	//
	// Any of "Wireline", "Wireless", "VoWiFi", "VoIP", "Pre-Paid Wireless", "".
	LineType string `json:"line_type"`
	// Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short
	// code).
	PhoneNumber string `json:"phone_number"`
	// Any of "received", "delivered".
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
func (r MessageSendWhatsappResponseDataFrom) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappResponseDataTo struct {
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
func (r MessageSendWhatsappResponseDataTo) RawJSON() string { return r.JSON.raw }
func (r *MessageSendWhatsappResponseDataTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageScheduleParams struct {
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required"`
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
	To string `json:"to,required"`
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
	From string `json:"from,required"`
	// A list of destinations. No more than 8 destinations are allowed.
	To []string `json:"to,omitzero,required"`
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
	From string `json:"from,required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required"`
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
	MessagingProfileID string `json:"messaging_profile_id,required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required"`
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
	From string `json:"from,required"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to,required"`
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

type MessageSendWhatsappParams struct {
	// Phone number in +E.164 format associated with Whatsapp account
	From string `json:"from,required"`
	// Phone number in +E.164 format
	To              string                                   `json:"to,required"`
	WhatsappMessage MessageSendWhatsappParamsWhatsappMessage `json:"whatsapp_message,omitzero,required"`
	// The URL where webhooks related to this message will be sent.
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero" format:"url"`
	// Message type - must be set to "WHATSAPP"
	//
	// Any of "WHATSAPP".
	Type MessageSendWhatsappParamsType `json:"type,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParams) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessage struct {
	// custom data to return with status update
	BizOpaqueCallbackData param.Opt[string]                                   `json:"biz_opaque_callback_data,omitzero"`
	Audio                 WhatsappMediaParam                                  `json:"audio,omitzero"`
	Contacts              []MessageSendWhatsappParamsWhatsappMessageContact   `json:"contacts,omitzero"`
	Document              WhatsappMediaParam                                  `json:"document,omitzero"`
	Image                 WhatsappMediaParam                                  `json:"image,omitzero"`
	Interactive           MessageSendWhatsappParamsWhatsappMessageInteractive `json:"interactive,omitzero"`
	Location              MessageSendWhatsappParamsWhatsappMessageLocation    `json:"location,omitzero"`
	Reaction              MessageSendWhatsappParamsWhatsappMessageReaction    `json:"reaction,omitzero"`
	Sticker               WhatsappMediaParam                                  `json:"sticker,omitzero"`
	// Any of "audio", "document", "image", "sticker", "video", "interactive",
	// "location", "template", "reaction", "contacts".
	Type  string             `json:"type,omitzero"`
	Video WhatsappMediaParam `json:"video,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessage) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageSendWhatsappParamsWhatsappMessage](
		"type", "audio", "document", "image", "sticker", "video", "interactive", "location", "template", "reaction", "contacts",
	)
}

type MessageSendWhatsappParamsWhatsappMessageContact struct {
	Birthday  param.Opt[string]                                        `json:"birthday,omitzero"`
	Name      param.Opt[string]                                        `json:"name,omitzero"`
	Addresses []MessageSendWhatsappParamsWhatsappMessageContactAddress `json:"addresses,omitzero"`
	Emails    []MessageSendWhatsappParamsWhatsappMessageContactEmail   `json:"emails,omitzero"`
	Org       MessageSendWhatsappParamsWhatsappMessageContactOrg       `json:"org,omitzero"`
	Phones    []MessageSendWhatsappParamsWhatsappMessageContactPhone   `json:"phones,omitzero"`
	URLs      []MessageSendWhatsappParamsWhatsappMessageContactURL     `json:"urls,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageContact) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageContact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageContactAddress struct {
	City        param.Opt[string] `json:"city,omitzero"`
	Country     param.Opt[string] `json:"country,omitzero"`
	CountryCode param.Opt[string] `json:"country_code,omitzero"`
	State       param.Opt[string] `json:"state,omitzero"`
	Street      param.Opt[string] `json:"street,omitzero"`
	Type        param.Opt[string] `json:"type,omitzero"`
	Zip         param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageContactAddress) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageContactAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageContactAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageContactEmail struct {
	Email param.Opt[string] `json:"email,omitzero"`
	Type  param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageContactEmail) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageContactEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageContactEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageContactOrg struct {
	Company    param.Opt[string] `json:"company,omitzero"`
	Department param.Opt[string] `json:"department,omitzero"`
	Title      param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageContactOrg) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageContactOrg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageContactOrg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageContactPhone struct {
	Phone param.Opt[string] `json:"phone,omitzero"`
	Type  param.Opt[string] `json:"type,omitzero"`
	WaID  param.Opt[string] `json:"wa_id,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageContactPhone) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageContactPhone
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageContactPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageContactURL struct {
	Type param.Opt[string] `json:"type,omitzero"`
	URL  param.Opt[string] `json:"url,omitzero" format:"url"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageContactURL) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageContactURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageContactURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractive struct {
	Action MessageSendWhatsappParamsWhatsappMessageInteractiveAction `json:"action,omitzero"`
	Body   MessageSendWhatsappParamsWhatsappMessageInteractiveBody   `json:"body,omitzero"`
	Footer MessageSendWhatsappParamsWhatsappMessageInteractiveFooter `json:"footer,omitzero"`
	Header MessageSendWhatsappParamsWhatsappMessageInteractiveHeader `json:"header,omitzero"`
	// Any of "cta_url", "list", "carousel", "button", "location_request_message".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractive) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractive
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageSendWhatsappParamsWhatsappMessageInteractive](
		"type", "cta_url", "list", "carousel", "button", "location_request_message",
	)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveAction struct {
	Button            param.Opt[string]                                                   `json:"button,omitzero"`
	CatalogID         param.Opt[string]                                                   `json:"catalog_id,omitzero"`
	Mode              param.Opt[string]                                                   `json:"mode,omitzero"`
	Name              param.Opt[string]                                                   `json:"name,omitzero"`
	ProductRetailerID param.Opt[string]                                                   `json:"product_retailer_id,omitzero"`
	Buttons           []MessageSendWhatsappParamsWhatsappMessageInteractiveActionButton   `json:"buttons,omitzero"`
	Cards             []MessageSendWhatsappParamsWhatsappMessageInteractiveActionCard     `json:"cards,omitzero"`
	Parameters        MessageSendWhatsappParamsWhatsappMessageInteractiveActionParameters `json:"parameters,omitzero"`
	Sections          []MessageSendWhatsappParamsWhatsappMessageInteractiveActionSection  `json:"sections,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveAction) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionButton struct {
	Reply MessageSendWhatsappParamsWhatsappMessageInteractiveActionButtonReply `json:"reply,omitzero"`
	// Any of "reply".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionButton) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionButton
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionButton) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageSendWhatsappParamsWhatsappMessageInteractiveActionButton](
		"type", "reply",
	)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionButtonReply struct {
	// unique identifier for each button, 256 character maximum
	ID param.Opt[string] `json:"id,omitzero"`
	// button label, 20 character maximum
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionButtonReply) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionButtonReply
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionButtonReply) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionCard struct {
	// unique index for each card (0-9)
	CardIndex param.Opt[int64]                                                    `json:"card_index,omitzero"`
	Action    MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardAction `json:"action,omitzero"`
	Body      MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardBody   `json:"body,omitzero"`
	Header    MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardHeader `json:"header,omitzero"`
	// Any of "cta_url".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionCard) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionCard
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageSendWhatsappParamsWhatsappMessageInteractiveActionCard](
		"type", "cta_url",
	)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardAction struct {
	// the unique ID of the catalog
	CatalogID param.Opt[string] `json:"catalog_id,omitzero"`
	// the unique retailer ID of the product
	ProductRetailerID param.Opt[string] `json:"product_retailer_id,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardAction) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardBody struct {
	// 160 character maximum, up to 2 line breaks
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardBody) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardHeader struct {
	Image WhatsappMediaParam `json:"image,omitzero"`
	// Any of "image", "video".
	Type  string             `json:"type,omitzero"`
	Video WhatsappMediaParam `json:"video,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardHeader) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MessageSendWhatsappParamsWhatsappMessageInteractiveActionCardHeader](
		"type", "image", "video",
	)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionParameters struct {
	// button label text, 20 character maximum
	DisplayText param.Opt[string] `json:"display_text,omitzero"`
	// button URL to load when tapped by the user
	URL param.Opt[string] `json:"url,omitzero" format:"url"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionParameters) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionSection struct {
	// section title, 24 character maximum
	Title        param.Opt[string]                                                             `json:"title,omitzero"`
	ProductItems []MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem `json:"product_items,omitzero"`
	Rows         []MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionRow         `json:"rows,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionSection) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem struct {
	ProductRetailerID param.Opt[string] `json:"product_retailer_id,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionProductItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionRow struct {
	// arbitrary string identifying the row, 200 character maximum
	ID param.Opt[string] `json:"id,omitzero"`
	// row description, 72 character maximum
	Description param.Opt[string] `json:"description,omitzero"`
	// row title, 24 character maximum
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionRow) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionRow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveActionSectionRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveBody struct {
	// body text, 1024 character maximum
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveBody) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveFooter struct {
	// footer text, 60 character maximum
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveFooter) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveFooter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveFooter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageInteractiveHeader struct {
	SubText param.Opt[string] `json:"sub_text,omitzero"`
	// header text, 60 character maximum
	Text     param.Opt[string]  `json:"text,omitzero"`
	Document WhatsappMediaParam `json:"document,omitzero"`
	Image    WhatsappMediaParam `json:"image,omitzero"`
	Video    WhatsappMediaParam `json:"video,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageInteractiveHeader) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageInteractiveHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageInteractiveHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageLocation struct {
	Address   param.Opt[string] `json:"address,omitzero"`
	Latitude  param.Opt[string] `json:"latitude,omitzero"`
	Longitude param.Opt[string] `json:"longitude,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageLocation) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageSendWhatsappParamsWhatsappMessageReaction struct {
	Emoji     param.Opt[string] `json:"emoji,omitzero"`
	MessageID param.Opt[string] `json:"message_id,omitzero"`
	paramObj
}

func (r MessageSendWhatsappParamsWhatsappMessageReaction) MarshalJSON() (data []byte, err error) {
	type shadow MessageSendWhatsappParamsWhatsappMessageReaction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageSendWhatsappParamsWhatsappMessageReaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Message type - must be set to "WHATSAPP"
type MessageSendWhatsappParamsType string

const (
	MessageSendWhatsappParamsTypeWhatsapp MessageSendWhatsappParamsType = "WHATSAPP"
)
