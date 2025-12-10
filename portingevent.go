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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
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
func (r *PortingEventService) List(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PortingEventListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "porting/events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of all porting events.
func (r *PortingEventService) ListAutoPaging(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PortingEventListResponse] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Republish a specific porting event.
func (r *PortingEventService) Republish(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	EventType PortingEventListResponseEventType `json:"event_type"`
	// The webhook payload for the porting_order.deleted event
	Payload PortingEventListResponsePayloadUnion `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventListResponsePayloadStatus `json:"payload_status"`
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
func (r PortingEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventListResponseEventType string

const (
	PortingEventListResponseEventTypePortingOrderDeleted             PortingEventListResponseEventType = "porting_order.deleted"
	PortingEventListResponseEventTypePortingOrderLoaUpdated          PortingEventListResponseEventType = "porting_order.loa_updated"
	PortingEventListResponseEventTypePortingOrderMessagingChanged    PortingEventListResponseEventType = "porting_order.messaging_changed"
	PortingEventListResponseEventTypePortingOrderStatusChanged       PortingEventListResponseEventType = "porting_order.status_changed"
	PortingEventListResponseEventTypePortingOrderSharingTokenExpired PortingEventListResponseEventType = "porting_order.sharing_token_expired"
	PortingEventListResponseEventTypePortingOrderNewComment          PortingEventListResponseEventType = "porting_order.new_comment"
	PortingEventListResponseEventTypePortingOrderSplit               PortingEventListResponseEventType = "porting_order.split"
)

// PortingEventListResponsePayloadUnion contains all possible properties and values
// from [PortingEventListResponsePayloadWebhookPortingOrderDeletedPayload],
// [PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload],
// [PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload],
// [PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload],
// [PortingEventListResponsePayloadWebhookPortingOrderSplitPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortingEventListResponsePayloadUnion struct {
	ID                string `json:"id"`
	CustomerReference string `json:"customer_reference"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderDeletedPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload].
	Messaging  PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayloadMessaging `json:"messaging"`
	SupportKey string                                                                             `json:"support_key"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload].
	Comment PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayloadComment `json:"comment"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderSplitPayload].
	From PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadFrom `json:"from"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderSplitPayload].
	PortingPhoneNumbers []PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant
	// [PortingEventListResponsePayloadWebhookPortingOrderSplitPayload].
	To   PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadTo `json:"to"`
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

func (u PortingEventListResponsePayloadUnion) AsPortingEventListResponsePayloadWebhookPortingOrderDeletedPayload() (v PortingEventListResponsePayloadWebhookPortingOrderDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponsePayloadUnion) AsPortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload() (v PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponsePayloadUnion) AsPortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload() (v PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponsePayloadUnion) AsPortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload() (v PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponsePayloadUnion) AsPortingEventListResponsePayloadWebhookPortingOrderSplitPayload() (v PortingEventListResponsePayloadWebhookPortingOrderSplitPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortingEventListResponsePayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *PortingEventListResponsePayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.deleted event
type PortingEventListResponsePayloadWebhookPortingOrderDeletedPayload struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderDeletedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderDeletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.messaging_changed event
type PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// The messaging portability status of the porting order.
	Messaging PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayloadMessaging `json:"messaging"`
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
func (r PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The messaging portability status of the porting order.
type PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayloadMessaging struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayloadMessaging) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderMessagingChangedPayloadMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.status_changed event
type PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.new_comment event
type PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload struct {
	// The comment that was added to the porting order.
	Comment PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayloadComment `json:"comment"`
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
func (r PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comment that was added to the porting order.
type PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayloadComment struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayloadComment) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderNewCommentPayloadComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.split event
type PortingEventListResponsePayloadWebhookPortingOrderSplitPayload struct {
	// The porting order that was split.
	From PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadFrom `json:"from"`
	// The list of porting phone numbers that were moved to the new porting order.
	PortingPhoneNumbers []PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// The new porting order that the phone numbers was moved to.
	To PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadTo `json:"to"`
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
func (r PortingEventListResponsePayloadWebhookPortingOrderSplitPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderSplitPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The porting order that was split.
type PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadFrom struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadPortingPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new porting order that the phone numbers was moved to.
type PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadTo struct {
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
func (r PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadTo) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePayloadWebhookPortingOrderSplitPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the payload generation.
type PortingEventListResponsePayloadStatus string

const (
	PortingEventListResponsePayloadStatusCreated   PortingEventListResponsePayloadStatus = "created"
	PortingEventListResponsePayloadStatusCompleted PortingEventListResponsePayloadStatus = "completed"
)

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
