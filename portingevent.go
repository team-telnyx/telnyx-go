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
func (r *PortingEventService) List(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PortingEventListResponseUnion], err error) {
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
func (r *PortingEventService) ListAutoPaging(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PortingEventListResponseUnion] {
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
	Data PortingEventGetResponseDataUnion `json:"data"`
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

// PortingEventGetResponseDataUnion contains all possible properties and values
// from [PortingEventGetResponseDataPortingEventDeletedPayload],
// [PortingEventGetResponseDataPortingEventMessagingChangedPayload],
// [PortingEventGetResponseDataPortingEventStatusChangedEvent],
// [PortingEventGetResponseDataPortingEventNewCommentEvent],
// [PortingEventGetResponseDataPortingEventSplitEvent],
// [PortingEventGetResponseDataPortingEventWithoutWebhook].
//
// Use the [PortingEventGetResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortingEventGetResponseDataUnion struct {
	ID                           string   `json:"id"`
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// Any of nil, nil, nil, nil, nil, nil.
	EventType string `json:"event_type"`
	// This field is a union of
	// [PortingEventGetResponseDataPortingEventDeletedPayloadPayload],
	// [PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayload],
	// [PortingEventGetResponseDataPortingEventStatusChangedEventPayload],
	// [PortingEventGetResponseDataPortingEventNewCommentEventPayload],
	// [PortingEventGetResponseDataPortingEventSplitEventPayload], [any]
	Payload        PortingEventGetResponseDataUnionPayload `json:"payload"`
	PayloadStatus  string                                  `json:"payload_status"`
	PortingOrderID string                                  `json:"porting_order_id"`
	CreatedAt      time.Time                               `json:"created_at"`
	RecordType     string                                  `json:"record_type"`
	UpdatedAt      time.Time                               `json:"updated_at"`
	JSON           struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortingOrderID               respjson.Field
		CreatedAt                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PortingEventGetResponseDataUnion) AsPortingEventGetResponseDataPortingEventDeletedPayload() (v PortingEventGetResponseDataPortingEventDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventGetResponseDataPortingEventMessagingChangedPayload() (v PortingEventGetResponseDataPortingEventMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventGetResponseDataPortingEventStatusChangedEvent() (v PortingEventGetResponseDataPortingEventStatusChangedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventGetResponseDataPortingEventNewCommentEvent() (v PortingEventGetResponseDataPortingEventNewCommentEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventGetResponseDataPortingEventSplitEvent() (v PortingEventGetResponseDataPortingEventSplitEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventGetResponseDataPortingEventWithoutWebhook() (v PortingEventGetResponseDataPortingEventWithoutWebhook) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortingEventGetResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *PortingEventGetResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortingEventGetResponseDataUnionPayload is an implicit subunion of
// [PortingEventGetResponseDataUnion]. PortingEventGetResponseDataUnionPayload
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PortingEventGetResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPortingEventGetResponseDataPortingEventWithoutWebhookPayload]
type PortingEventGetResponseDataUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPortingEventGetResponseDataPortingEventWithoutWebhookPayload any    `json:",inline"`
	ID                                                             string `json:"id"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventDeletedPayloadPayload].
	CreatedAt         time.Time `json:"created_at"`
	CustomerReference string    `json:"customer_reference"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventDeletedPayloadPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventDeletedPayloadPayload].
	RecordType string    `json:"record_type"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayload].
	Messaging  PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
	SupportKey string                                                                         `json:"support_key"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventStatusChangedEventPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventStatusChangedEventPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventNewCommentEventPayload].
	Comment PortingEventGetResponseDataPortingEventNewCommentEventPayloadComment `json:"comment"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventNewCommentEventPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventSplitEventPayload].
	From PortingEventGetResponseDataPortingEventSplitEventPayloadFrom `json:"from"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventSplitEventPayload].
	PortingPhoneNumbers []PortingEventGetResponseDataPortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant
	// [PortingEventGetResponseDataPortingEventSplitEventPayload].
	To   PortingEventGetResponseDataPortingEventSplitEventPayloadTo `json:"to"`
	JSON struct {
		OfPortingEventGetResponseDataPortingEventWithoutWebhookPayload respjson.Field
		ID                                                             respjson.Field
		CreatedAt                                                      respjson.Field
		CustomerReference                                              respjson.Field
		DeletedAt                                                      respjson.Field
		RecordType                                                     respjson.Field
		UpdatedAt                                                      respjson.Field
		Messaging                                                      respjson.Field
		SupportKey                                                     respjson.Field
		Status                                                         respjson.Field
		WebhookURL                                                     respjson.Field
		Comment                                                        respjson.Field
		PortingOrderID                                                 respjson.Field
		From                                                           respjson.Field
		PortingPhoneNumbers                                            respjson.Field
		To                                                             respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (r *PortingEventGetResponseDataUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventDeletedPayload struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook", "webhook_v1".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// Identifies the event type
	//
	// Any of "porting_order.deleted", "porting_order.loa_updated",
	// "porting_order.messaging_changed", "porting_order.status_changed",
	// "porting_order.sharing_token_expired", "porting_order.new_comment",
	// "porting_order.split".
	EventType string                                                       `json:"event_type"`
	Payload   PortingEventGetResponseDataPortingEventDeletedPayloadPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the porting order associated with the event.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortingOrderID               respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPortingEventDeletedPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponseDataPortingEventDeletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventDeletedPayloadPayload struct {
	// Identifies the porting order that was deleted.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// ISO 8601 formatted date indicating when the porting order was deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		CustomerReference respjson.Field
		DeletedAt         respjson.Field
		RecordType        respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventGetResponseDataPortingEventDeletedPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventDeletedPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventMessagingChangedPayload struct {
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
	// The webhook payload for the porting_order.messaging_changed event
	Payload PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayload `json:"payload"`
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
func (r PortingEventGetResponseDataPortingEventMessagingChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventMessagingChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.messaging_changed event
type PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// The messaging portability status of the porting order.
	Messaging PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
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
func (r PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The messaging portability status of the porting order.
type PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayloadMessaging struct {
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
func (r PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayloadMessaging) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventMessagingChangedPayloadPayloadMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventStatusChangedEvent struct {
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
	// The webhook payload for the porting_order.status_changed event
	Payload PortingEventGetResponseDataPortingEventStatusChangedEventPayload `json:"payload"`
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
func (r PortingEventGetResponseDataPortingEventStatusChangedEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventStatusChangedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.status_changed event
type PortingEventGetResponseDataPortingEventStatusChangedEventPayload struct {
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
func (r PortingEventGetResponseDataPortingEventStatusChangedEventPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventStatusChangedEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventNewCommentEvent struct {
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
	// The webhook payload for the porting_order.new_comment event
	Payload PortingEventGetResponseDataPortingEventNewCommentEventPayload `json:"payload"`
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
func (r PortingEventGetResponseDataPortingEventNewCommentEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponseDataPortingEventNewCommentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.new_comment event
type PortingEventGetResponseDataPortingEventNewCommentEventPayload struct {
	// The comment that was added to the porting order.
	Comment PortingEventGetResponseDataPortingEventNewCommentEventPayloadComment `json:"comment"`
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
func (r PortingEventGetResponseDataPortingEventNewCommentEventPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventNewCommentEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comment that was added to the porting order.
type PortingEventGetResponseDataPortingEventNewCommentEventPayloadComment struct {
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
func (r PortingEventGetResponseDataPortingEventNewCommentEventPayloadComment) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventNewCommentEventPayloadComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventSplitEvent struct {
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
	// The webhook payload for the porting_order.split event
	Payload PortingEventGetResponseDataPortingEventSplitEventPayload `json:"payload"`
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
func (r PortingEventGetResponseDataPortingEventSplitEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponseDataPortingEventSplitEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.split event
type PortingEventGetResponseDataPortingEventSplitEventPayload struct {
	// The porting order that was split.
	From PortingEventGetResponseDataPortingEventSplitEventPayloadFrom `json:"from"`
	// The list of porting phone numbers that were moved to the new porting order.
	PortingPhoneNumbers []PortingEventGetResponseDataPortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// The new porting order that the phone numbers was moved to.
	To PortingEventGetResponseDataPortingEventSplitEventPayloadTo `json:"to"`
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
func (r PortingEventGetResponseDataPortingEventSplitEventPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponseDataPortingEventSplitEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The porting order that was split.
type PortingEventGetResponseDataPortingEventSplitEventPayloadFrom struct {
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
func (r PortingEventGetResponseDataPortingEventSplitEventPayloadFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventSplitEventPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventSplitEventPayloadPortingPhoneNumber struct {
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
func (r PortingEventGetResponseDataPortingEventSplitEventPayloadPortingPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventSplitEventPayloadPortingPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new porting order that the phone numbers was moved to.
type PortingEventGetResponseDataPortingEventSplitEventPayloadTo struct {
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
func (r PortingEventGetResponseDataPortingEventSplitEventPayloadTo) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventGetResponseDataPortingEventSplitEventPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventGetResponseDataPortingEventWithoutWebhook struct {
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
	Payload   any    `json:"payload,nullable"`
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
func (r PortingEventGetResponseDataPortingEventWithoutWebhook) RawJSON() string { return r.JSON.raw }
func (r *PortingEventGetResponseDataPortingEventWithoutWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortingEventListResponseUnion contains all possible properties and values from
// [PortingEventListResponsePortingEventDeletedPayload],
// [PortingEventListResponsePortingEventMessagingChangedPayload],
// [PortingEventListResponsePortingEventStatusChangedEvent],
// [PortingEventListResponsePortingEventNewCommentEvent],
// [PortingEventListResponsePortingEventSplitEvent],
// [PortingEventListResponsePortingEventWithoutWebhook].
//
// Use the [PortingEventListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortingEventListResponseUnion struct {
	ID                           string   `json:"id"`
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// Any of nil, nil, nil, nil, nil, nil.
	EventType string `json:"event_type"`
	// This field is a union of
	// [PortingEventListResponsePortingEventDeletedPayloadPayload],
	// [PortingEventListResponsePortingEventMessagingChangedPayloadPayload],
	// [PortingEventListResponsePortingEventStatusChangedEventPayload],
	// [PortingEventListResponsePortingEventNewCommentEventPayload],
	// [PortingEventListResponsePortingEventSplitEventPayload], [any]
	Payload        PortingEventListResponseUnionPayload `json:"payload"`
	PayloadStatus  string                               `json:"payload_status"`
	PortingOrderID string                               `json:"porting_order_id"`
	CreatedAt      time.Time                            `json:"created_at"`
	RecordType     string                               `json:"record_type"`
	UpdatedAt      time.Time                            `json:"updated_at"`
	JSON           struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortingOrderID               respjson.Field
		CreatedAt                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PortingEventListResponseUnion) AsPortingEventListResponsePortingEventDeletedPayload() (v PortingEventListResponsePortingEventDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventListResponsePortingEventMessagingChangedPayload() (v PortingEventListResponsePortingEventMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventListResponsePortingEventStatusChangedEvent() (v PortingEventListResponsePortingEventStatusChangedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventListResponsePortingEventNewCommentEvent() (v PortingEventListResponsePortingEventNewCommentEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventListResponsePortingEventSplitEvent() (v PortingEventListResponsePortingEventSplitEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventListResponsePortingEventWithoutWebhook() (v PortingEventListResponsePortingEventWithoutWebhook) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortingEventListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *PortingEventListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortingEventListResponseUnionPayload is an implicit subunion of
// [PortingEventListResponseUnion]. PortingEventListResponseUnionPayload provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PortingEventListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPortingEventListResponsePortingEventWithoutWebhookPayload]
type PortingEventListResponseUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPortingEventListResponsePortingEventWithoutWebhookPayload any    `json:",inline"`
	ID                                                          string `json:"id"`
	// This field is from variant
	// [PortingEventListResponsePortingEventDeletedPayloadPayload].
	CreatedAt         time.Time `json:"created_at"`
	CustomerReference string    `json:"customer_reference"`
	// This field is from variant
	// [PortingEventListResponsePortingEventDeletedPayloadPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant
	// [PortingEventListResponsePortingEventDeletedPayloadPayload].
	RecordType string    `json:"record_type"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant
	// [PortingEventListResponsePortingEventMessagingChangedPayloadPayload].
	Messaging  PortingEventListResponsePortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
	SupportKey string                                                                      `json:"support_key"`
	// This field is from variant
	// [PortingEventListResponsePortingEventStatusChangedEventPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant
	// [PortingEventListResponsePortingEventStatusChangedEventPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant
	// [PortingEventListResponsePortingEventNewCommentEventPayload].
	Comment PortingEventListResponsePortingEventNewCommentEventPayloadComment `json:"comment"`
	// This field is from variant
	// [PortingEventListResponsePortingEventNewCommentEventPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant
	// [PortingEventListResponsePortingEventSplitEventPayload].
	From PortingEventListResponsePortingEventSplitEventPayloadFrom `json:"from"`
	// This field is from variant
	// [PortingEventListResponsePortingEventSplitEventPayload].
	PortingPhoneNumbers []PortingEventListResponsePortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant
	// [PortingEventListResponsePortingEventSplitEventPayload].
	To   PortingEventListResponsePortingEventSplitEventPayloadTo `json:"to"`
	JSON struct {
		OfPortingEventListResponsePortingEventWithoutWebhookPayload respjson.Field
		ID                                                          respjson.Field
		CreatedAt                                                   respjson.Field
		CustomerReference                                           respjson.Field
		DeletedAt                                                   respjson.Field
		RecordType                                                  respjson.Field
		UpdatedAt                                                   respjson.Field
		Messaging                                                   respjson.Field
		SupportKey                                                  respjson.Field
		Status                                                      respjson.Field
		WebhookURL                                                  respjson.Field
		Comment                                                     respjson.Field
		PortingOrderID                                              respjson.Field
		From                                                        respjson.Field
		PortingPhoneNumbers                                         respjson.Field
		To                                                          respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (r *PortingEventListResponseUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventDeletedPayload struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook", "webhook_v1".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// Identifies the event type
	//
	// Any of "porting_order.deleted", "porting_order.loa_updated",
	// "porting_order.messaging_changed", "porting_order.status_changed",
	// "porting_order.sharing_token_expired", "porting_order.new_comment",
	// "porting_order.split".
	EventType string                                                    `json:"event_type"`
	Payload   PortingEventListResponsePortingEventDeletedPayloadPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the porting order associated with the event.
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortingOrderID               respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponsePortingEventDeletedPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventDeletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventDeletedPayloadPayload struct {
	// Identifies the porting order that was deleted.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// ISO 8601 formatted date indicating when the porting order was deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		CustomerReference respjson.Field
		DeletedAt         respjson.Field
		RecordType        respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingEventListResponsePortingEventDeletedPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventDeletedPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventMessagingChangedPayload struct {
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
	// The webhook payload for the porting_order.messaging_changed event
	Payload PortingEventListResponsePortingEventMessagingChangedPayloadPayload `json:"payload"`
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
func (r PortingEventListResponsePortingEventMessagingChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventMessagingChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.messaging_changed event
type PortingEventListResponsePortingEventMessagingChangedPayloadPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// The messaging portability status of the porting order.
	Messaging PortingEventListResponsePortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
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
func (r PortingEventListResponsePortingEventMessagingChangedPayloadPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventMessagingChangedPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The messaging portability status of the porting order.
type PortingEventListResponsePortingEventMessagingChangedPayloadPayloadMessaging struct {
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
func (r PortingEventListResponsePortingEventMessagingChangedPayloadPayloadMessaging) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventMessagingChangedPayloadPayloadMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventStatusChangedEvent struct {
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
	// The webhook payload for the porting_order.status_changed event
	Payload PortingEventListResponsePortingEventStatusChangedEventPayload `json:"payload"`
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
func (r PortingEventListResponsePortingEventStatusChangedEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventStatusChangedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.status_changed event
type PortingEventListResponsePortingEventStatusChangedEventPayload struct {
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
func (r PortingEventListResponsePortingEventStatusChangedEventPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventStatusChangedEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventNewCommentEvent struct {
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
	// The webhook payload for the porting_order.new_comment event
	Payload PortingEventListResponsePortingEventNewCommentEventPayload `json:"payload"`
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
func (r PortingEventListResponsePortingEventNewCommentEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventNewCommentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.new_comment event
type PortingEventListResponsePortingEventNewCommentEventPayload struct {
	// The comment that was added to the porting order.
	Comment PortingEventListResponsePortingEventNewCommentEventPayloadComment `json:"comment"`
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
func (r PortingEventListResponsePortingEventNewCommentEventPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventNewCommentEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comment that was added to the porting order.
type PortingEventListResponsePortingEventNewCommentEventPayloadComment struct {
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
func (r PortingEventListResponsePortingEventNewCommentEventPayloadComment) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventNewCommentEventPayloadComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventSplitEvent struct {
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
	// The webhook payload for the porting_order.split event
	Payload PortingEventListResponsePortingEventSplitEventPayload `json:"payload"`
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
func (r PortingEventListResponsePortingEventSplitEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventSplitEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the porting_order.split event
type PortingEventListResponsePortingEventSplitEventPayload struct {
	// The porting order that was split.
	From PortingEventListResponsePortingEventSplitEventPayloadFrom `json:"from"`
	// The list of porting phone numbers that were moved to the new porting order.
	PortingPhoneNumbers []PortingEventListResponsePortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// The new porting order that the phone numbers was moved to.
	To PortingEventListResponsePortingEventSplitEventPayloadTo `json:"to"`
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
func (r PortingEventListResponsePortingEventSplitEventPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventSplitEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The porting order that was split.
type PortingEventListResponsePortingEventSplitEventPayloadFrom struct {
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
func (r PortingEventListResponsePortingEventSplitEventPayloadFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventSplitEventPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventSplitEventPayloadPortingPhoneNumber struct {
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
func (r PortingEventListResponsePortingEventSplitEventPayloadPortingPhoneNumber) RawJSON() string {
	return r.JSON.raw
}
func (r *PortingEventListResponsePortingEventSplitEventPayloadPortingPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new porting order that the phone numbers was moved to.
type PortingEventListResponsePortingEventSplitEventPayloadTo struct {
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
func (r PortingEventListResponsePortingEventSplitEventPayloadTo) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventSplitEventPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListResponsePortingEventWithoutWebhook struct {
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
	Payload   any    `json:"payload,nullable"`
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
func (r PortingEventListResponsePortingEventWithoutWebhook) RawJSON() string { return r.JSON.raw }
func (r *PortingEventListResponsePortingEventWithoutWebhook) UnmarshalJSON(data []byte) error {
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
