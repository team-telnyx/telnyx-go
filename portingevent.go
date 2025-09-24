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
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
	"github.com/team-telnyx/telnyx-go/shared"
)

// PortingEventService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingEventService] method instead.
type PortingEventService struct {
	Options []option.RequestOption
}

// NewPortingEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortingEventService(opts ...option.RequestOption) (r PortingEventService) {
	r = PortingEventService{}
	r.Options = opts
	return
}

// Show a specific porting event.
func (r *PortingEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/events/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all porting events.
func (r *PortingEventService) List(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) (res *PortingEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "porting/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Republish a specific porting event.
func (r *PortingEventService) Republish(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting/events/%s/republish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type PortingEventGetResponse struct {
	Data PortingEventGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseData struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook", "webhook_v1".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "porting_order.deleted", "porting_order.loa_updated",
	// "porting_order.messaging_changed", "porting_order.status_changed",
	// "porting_order.sharing_token_expired", "porting_order.new_comment",
	// "porting_order.split".
	EventType string `json:"event_type"`
	// The webhook payload for the porting_order.deleted event
	Payload PortingEventGetResponseDataPayloadUnion `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the porting order associated with the event.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		CreatedAt                    respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortingOrderID               respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortingEventGetResponseDataPayloadUnion contains all possible properties and
// values from
// [PortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload],
// [PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload],
// [PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload],
// [PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload],
// [PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortingEventGetResponseDataPayloadUnion struct {
	ID                string `json:"id"`
	CustomerReference string `json:"customer_reference"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload].
	Messaging  PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging `json:"messaging"`
	SupportKey string                                                                                `json:"support_key"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload].
	Comment PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment `json:"comment"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload].
	From PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadFrom `json:"from"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload].
	PortingPhoneNumbers []PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant
	// [PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload].
	To   PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadTo `json:"to"`
	JSON struct {
		ID                  respjson.Field
		CustomerReference   respjson.Field
		DeletedAt           respjson.Field
		Messaging           respjson.Field
		SupportKey          respjson.Field
		Status              respjson.Field
		UpdatedAt           respjson.Field
		WebhookURL          respjson.Field
		Comment             respjson.Field
		PortingOrderID      respjson.Field
		From                respjson.Field
		PortingPhoneNumbers respjson.Field
		To                  respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PortingEventGetResponseDataPayloadUnion) AsPortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload() (v PortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataPayloadUnion) AsPortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload() (v PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataPayloadUnion) AsPortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload() (v PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataPayloadUnion) AsPortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload() (v PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataPayloadUnion) AsPortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload() (v PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortingEventGetResponseDataPayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *PortingEventGetResponseDataPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.deleted event
type PortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload struct {
	// Identifies the porting order that was deleted.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// ISO 8601 formatted date indicating when the porting order was deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CustomerReference respjson.Field
		DeletedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderDeletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.messaging_changed event
type PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// The messaging portability status of the porting order.
	Messaging PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging `json:"messaging"`
	// Identifies the support key associated with the porting order.
	SupportKey string `json:"support_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CustomerReference respjson.Field
		Messaging         respjson.Field
		SupportKey        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The messaging portability status of the porting order.
type PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging struct {
	// Indicates whether Telnyx will port messaging capabilities from the losing
	// carrier. If false, any messaging capabilities will stay with their current
	// provider.
	EnableMessaging bool `json:"enable_messaging"`
	// Indicates whether the porting order is messaging capable.
	MessagingCapable bool `json:"messaging_capable"`
	// Indicates whether the messaging port is completed.
	MessagingPortCompleted bool `json:"messaging_port_completed"`
	// Indicates the messaging port status of the porting order.
	//
	// Any of "not_applicable", "pending", "activating", "exception", "canceled",
	// "partial_port_complete", "ported".
	MessagingPortStatus string `json:"messaging_port_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableMessaging        respjson.Field
		MessagingCapable       respjson.Field
		MessagingPortCompleted respjson.Field
		MessagingPortStatus    respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.status_changed event
type PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// Porting order status
	Status shared.PortingOrderStatus `json:"status"`
	// Identifies the support key associated with the porting order.
	SupportKey string `json:"support_key"`
	// ISO 8601 formatted date indicating when the porting order was moved.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The URL to send the webhook to.
	WebhookURL string `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CustomerReference respjson.Field
		Status            respjson.Field
		SupportKey        respjson.Field
		UpdatedAt         respjson.Field
		WebhookURL        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.new_comment event
type PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload struct {
	// The comment that was added to the porting order.
	Comment PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment `json:"comment"`
	// Identifies the porting order that the comment was added to.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the support key associated with the porting order.
	SupportKey string `json:"support_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comment        respjson.Field
		PortingOrderID respjson.Field
		SupportKey     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comment that was added to the porting order.
type PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment struct {
	// Identifies the comment.
	ID string `json:"id" format:"uuid"`
	// The body of the comment.
	Body string `json:"body"`
	// ISO 8601 formatted date indicating when the comment was created.
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Identifies the user that create the comment.
	UserID string `json:"user_id" format:"uuid"`
	// Identifies the type of the user that created the comment.
	//
	// Any of "user", "admin", "system".
	UserType string `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		InsertedAt  respjson.Field
		UserID      respjson.Field
		UserType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.split event
type PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload struct {
	// The porting order that was split.
	From PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadFrom `json:"from"`
	// The list of porting phone numbers that were moved to the new porting order.
	PortingPhoneNumbers []PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// The new porting order that the phone numbers was moved to.
	To PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadTo `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From                respjson.Field
		PortingPhoneNumbers respjson.Field
		To                  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The porting order that was split.
type PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadFrom struct {
	// Identifies the porting order that was split.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber struct {
	// Identifies the porting phone number that was moved.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new porting order that the phone numbers was moved to.
type PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadTo struct {
	// Identifies the porting order that was split.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadTo) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPayloadWebhookPortingOrderSplitPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponse struct {
	Data []PortingEventListResponseData `json:"data"`
	Meta PaginationMeta                 `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponseData struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook", "webhook_v1".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "porting_order.deleted", "porting_order.loa_updated",
	// "porting_order.messaging_changed", "porting_order.status_changed",
	// "porting_order.sharing_token_expired", "porting_order.new_comment",
	// "porting_order.split".
	EventType string `json:"event_type"`
	// The webhook payload for the porting_order.deleted event
	Payload PortingEventListResponseDataPayloadUnion `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the porting order associated with the event.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		CreatedAt                    respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortingOrderID               respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortingEventListResponseDataPayloadUnion contains all possible properties and
// values from
// [PortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload],
// [PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload],
// [PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload],
// [PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload],
// [PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortingEventListResponseDataPayloadUnion struct {
	ID                string `json:"id"`
	CustomerReference string `json:"customer_reference"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload].
	Messaging  PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging `json:"messaging"`
	SupportKey string                                                                                 `json:"support_key"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload].
	Comment PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment `json:"comment"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload].
	From PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadFrom `json:"from"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload].
	PortingPhoneNumbers []PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant
	// [PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload].
	To   PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadTo `json:"to"`
	JSON struct {
		ID                  respjson.Field
		CustomerReference   respjson.Field
		DeletedAt           respjson.Field
		Messaging           respjson.Field
		SupportKey          respjson.Field
		Status              respjson.Field
		UpdatedAt           respjson.Field
		WebhookURL          respjson.Field
		Comment             respjson.Field
		PortingOrderID      respjson.Field
		From                respjson.Field
		PortingPhoneNumbers respjson.Field
		To                  respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PortingEventListResponseDataPayloadUnion) AsPortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload() (v PortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseDataPayloadUnion) AsPortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload() (v PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseDataPayloadUnion) AsPortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload() (v PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseDataPayloadUnion) AsPortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload() (v PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseDataPayloadUnion) AsPortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload() (v PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortingEventListResponseDataPayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *PortingEventListResponseDataPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.deleted event
type PortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload struct {
	// Identifies the porting order that was deleted.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// ISO 8601 formatted date indicating when the porting order was deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CustomerReference respjson.Field
		DeletedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderDeletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.messaging_changed event
type PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// The messaging portability status of the porting order.
	Messaging PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging `json:"messaging"`
	// Identifies the support key associated with the porting order.
	SupportKey string `json:"support_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CustomerReference respjson.Field
		Messaging         respjson.Field
		SupportKey        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The messaging portability status of the porting order.
type PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging struct {
	// Indicates whether Telnyx will port messaging capabilities from the losing
	// carrier. If false, any messaging capabilities will stay with their current
	// provider.
	EnableMessaging bool `json:"enable_messaging"`
	// Indicates whether the porting order is messaging capable.
	MessagingCapable bool `json:"messaging_capable"`
	// Indicates whether the messaging port is completed.
	MessagingPortCompleted bool `json:"messaging_port_completed"`
	// Indicates the messaging port status of the porting order.
	//
	// Any of "not_applicable", "pending", "activating", "exception", "canceled",
	// "partial_port_complete", "ported".
	MessagingPortStatus string `json:"messaging_port_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableMessaging        respjson.Field
		MessagingCapable       respjson.Field
		MessagingPortCompleted respjson.Field
		MessagingPortStatus    respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderMessagingChangedPayloadMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.status_changed event
type PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// Porting order status
	Status shared.PortingOrderStatus `json:"status"`
	// Identifies the support key associated with the porting order.
	SupportKey string `json:"support_key"`
	// ISO 8601 formatted date indicating when the porting order was moved.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The URL to send the webhook to.
	WebhookURL string `json:"webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CustomerReference respjson.Field
		Status            respjson.Field
		SupportKey        respjson.Field
		UpdatedAt         respjson.Field
		WebhookURL        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.new_comment event
type PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload struct {
	// The comment that was added to the porting order.
	Comment PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment `json:"comment"`
	// Identifies the porting order that the comment was added to.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the support key associated with the porting order.
	SupportKey string `json:"support_key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comment        respjson.Field
		PortingOrderID respjson.Field
		SupportKey     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comment that was added to the porting order.
type PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment struct {
	// Identifies the comment.
	ID string `json:"id" format:"uuid"`
	// The body of the comment.
	Body string `json:"body"`
	// ISO 8601 formatted date indicating when the comment was created.
	InsertedAt time.Time `json:"inserted_at" format:"date-time"`
	// Identifies the user that create the comment.
	UserID string `json:"user_id" format:"uuid"`
	// Identifies the type of the user that created the comment.
	//
	// Any of "user", "admin", "system".
	UserType string `json:"user_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		InsertedAt  respjson.Field
		UserID      respjson.Field
		UserType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderNewCommentPayloadComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.split event
type PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload struct {
	// The porting order that was split.
	From PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadFrom `json:"from"`
	// The list of porting phone numbers that were moved to the new porting order.
	PortingPhoneNumbers []PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// The new porting order that the phone numbers was moved to.
	To PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadTo `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From                respjson.Field
		PortingPhoneNumbers respjson.Field
		To                  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The porting order that was split.
type PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadFrom struct {
	// Identifies the porting order that was split.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber struct {
	// Identifies the porting phone number that was moved.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new porting order that the phone numbers was moved to.
type PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadTo struct {
	// Identifies the porting order that was split.
	ID string `json:"id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadTo) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponseDataPayloadWebhookPortingOrderSplitPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListParams struct {
	// Consolidated filter parameter (deepObject style). Originally: filter[type],
	// filter[porting_order_id], filter[created_at][gte], filter[created_at][lte]
	Filter PortingEventListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page PortingEventListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingEventListParams]'s query parameters as `url.Values`.
func (r PortingEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[type],
// filter[porting_order_id], filter[created_at][gte], filter[created_at][lte]
type PortingEventListParamsFilter struct {
	// Filter by porting order ID.
	PortingOrderID param.Opt[string] `query:"porting_order_id,omitzero" format:"uuid" json:"-"`
	// Created at date range filtering operations
	CreatedAt PortingEventListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	// Filter by event type.
	//
	// Any of "porting_order.deleted", "porting_order.loa_updated",
	// "porting_order.messaging_changed", "porting_order.status_changed",
	// "porting_order.sharing_token_expired", "porting_order.new_comment",
	// "porting_order.split".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingEventListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortingEventListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Created at date range filtering operations
type PortingEventListParamsFilterCreatedAt struct {
	// Filter by created at greater than or equal to.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Filter by created at less than or equal to.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PortingEventListParamsFilterCreatedAt]'s query parameters
// as `url.Values`.
func (r PortingEventListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type PortingEventListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortingEventListParamsPage]'s query parameters as
// `url.Values`.
func (r PortingEventListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
