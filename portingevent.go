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

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Endpoints related to porting orders management.
//
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
func (r *PortingEventService) List(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PortingEventListResponseUnion], err error) {
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
func (r *PortingEventService) ListAutoPaging(ctx context.Context, query PortingEventListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PortingEventListResponseUnion] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
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

type PortingEventDeletedPayload struct {
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
	EventType PortingEventDeletedPayloadEventType `json:"event_type"`
	Payload   PortingEventDeletedPayloadPayload   `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventDeletedPayloadPayloadStatus `json:"payload_status"`
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
func (r PortingEventDeletedPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventDeletedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventDeletedPayloadEventType string

const (
	PortingEventDeletedPayloadEventTypePortingOrderDeleted             PortingEventDeletedPayloadEventType = "porting_order.deleted"
	PortingEventDeletedPayloadEventTypePortingOrderLoaUpdated          PortingEventDeletedPayloadEventType = "porting_order.loa_updated"
	PortingEventDeletedPayloadEventTypePortingOrderMessagingChanged    PortingEventDeletedPayloadEventType = "porting_order.messaging_changed"
	PortingEventDeletedPayloadEventTypePortingOrderStatusChanged       PortingEventDeletedPayloadEventType = "porting_order.status_changed"
	PortingEventDeletedPayloadEventTypePortingOrderSharingTokenExpired PortingEventDeletedPayloadEventType = "porting_order.sharing_token_expired"
	PortingEventDeletedPayloadEventTypePortingOrderNewComment          PortingEventDeletedPayloadEventType = "porting_order.new_comment"
	PortingEventDeletedPayloadEventTypePortingOrderSplit               PortingEventDeletedPayloadEventType = "porting_order.split"
)

type PortingEventDeletedPayloadPayload struct {
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
func (r PortingEventDeletedPayloadPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventDeletedPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the payload generation.
type PortingEventDeletedPayloadPayloadStatus string

const (
	PortingEventDeletedPayloadPayloadStatusCreated   PortingEventDeletedPayloadPayloadStatus = "created"
	PortingEventDeletedPayloadPayloadStatusCompleted PortingEventDeletedPayloadPayloadStatus = "completed"
)

type PortingEventMessagingChangedPayload struct {
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
	EventType PortingEventMessagingChangedPayloadEventType `json:"event_type"`
	// The webhook payload for the porting_order.messaging_changed event
	Payload PortingEventMessagingChangedPayloadPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventMessagingChangedPayloadPayloadStatus `json:"payload_status"`
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
func (r PortingEventMessagingChangedPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventMessagingChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventMessagingChangedPayloadEventType string

const (
	PortingEventMessagingChangedPayloadEventTypePortingOrderDeleted             PortingEventMessagingChangedPayloadEventType = "porting_order.deleted"
	PortingEventMessagingChangedPayloadEventTypePortingOrderLoaUpdated          PortingEventMessagingChangedPayloadEventType = "porting_order.loa_updated"
	PortingEventMessagingChangedPayloadEventTypePortingOrderMessagingChanged    PortingEventMessagingChangedPayloadEventType = "porting_order.messaging_changed"
	PortingEventMessagingChangedPayloadEventTypePortingOrderStatusChanged       PortingEventMessagingChangedPayloadEventType = "porting_order.status_changed"
	PortingEventMessagingChangedPayloadEventTypePortingOrderSharingTokenExpired PortingEventMessagingChangedPayloadEventType = "porting_order.sharing_token_expired"
	PortingEventMessagingChangedPayloadEventTypePortingOrderNewComment          PortingEventMessagingChangedPayloadEventType = "porting_order.new_comment"
	PortingEventMessagingChangedPayloadEventTypePortingOrderSplit               PortingEventMessagingChangedPayloadEventType = "porting_order.split"
)

// The webhook payload for the porting_order.messaging_changed event
type PortingEventMessagingChangedPayloadPayload struct {
	// Identifies the porting order that was moved.
	ID string `json:"id" format:"uuid"`
	// Identifies the customer reference associated with the porting order.
	CustomerReference string `json:"customer_reference"`
	// The messaging portability status of the porting order.
	Messaging PortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
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
func (r PortingEventMessagingChangedPayloadPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventMessagingChangedPayloadPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The messaging portability status of the porting order.
type PortingEventMessagingChangedPayloadPayloadMessaging struct {
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
func (r PortingEventMessagingChangedPayloadPayloadMessaging) RawJSON() string { return r.JSON.raw }
func (r *PortingEventMessagingChangedPayloadPayloadMessaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the payload generation.
type PortingEventMessagingChangedPayloadPayloadStatus string

const (
	PortingEventMessagingChangedPayloadPayloadStatusCreated   PortingEventMessagingChangedPayloadPayloadStatus = "created"
	PortingEventMessagingChangedPayloadPayloadStatusCompleted PortingEventMessagingChangedPayloadPayloadStatus = "completed"
)

type PortingEventNewCommentEvent struct {
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
	EventType PortingEventNewCommentEventEventType `json:"event_type"`
	// The webhook payload for the porting_order.new_comment event
	Payload PortingEventNewCommentEventPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventNewCommentEventPayloadStatus `json:"payload_status"`
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
func (r PortingEventNewCommentEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventNewCommentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventNewCommentEventEventType string

const (
	PortingEventNewCommentEventEventTypePortingOrderDeleted             PortingEventNewCommentEventEventType = "porting_order.deleted"
	PortingEventNewCommentEventEventTypePortingOrderLoaUpdated          PortingEventNewCommentEventEventType = "porting_order.loa_updated"
	PortingEventNewCommentEventEventTypePortingOrderMessagingChanged    PortingEventNewCommentEventEventType = "porting_order.messaging_changed"
	PortingEventNewCommentEventEventTypePortingOrderStatusChanged       PortingEventNewCommentEventEventType = "porting_order.status_changed"
	PortingEventNewCommentEventEventTypePortingOrderSharingTokenExpired PortingEventNewCommentEventEventType = "porting_order.sharing_token_expired"
	PortingEventNewCommentEventEventTypePortingOrderNewComment          PortingEventNewCommentEventEventType = "porting_order.new_comment"
	PortingEventNewCommentEventEventTypePortingOrderSplit               PortingEventNewCommentEventEventType = "porting_order.split"
)

// The webhook payload for the porting_order.new_comment event
type PortingEventNewCommentEventPayload struct {
	// The comment that was added to the porting order.
	Comment PortingEventNewCommentEventPayloadComment `json:"comment"`
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
func (r PortingEventNewCommentEventPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventNewCommentEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comment that was added to the porting order.
type PortingEventNewCommentEventPayloadComment struct {
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
func (r PortingEventNewCommentEventPayloadComment) RawJSON() string { return r.JSON.raw }
func (r *PortingEventNewCommentEventPayloadComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the payload generation.
type PortingEventNewCommentEventPayloadStatus string

const (
	PortingEventNewCommentEventPayloadStatusCreated   PortingEventNewCommentEventPayloadStatus = "created"
	PortingEventNewCommentEventPayloadStatusCompleted PortingEventNewCommentEventPayloadStatus = "completed"
)

type PortingEventSplitEvent struct {
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
	EventType PortingEventSplitEventEventType `json:"event_type"`
	// The webhook payload for the porting_order.split event
	Payload PortingEventSplitEventPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventSplitEventPayloadStatus `json:"payload_status"`
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
func (r PortingEventSplitEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventSplitEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventSplitEventEventType string

const (
	PortingEventSplitEventEventTypePortingOrderDeleted             PortingEventSplitEventEventType = "porting_order.deleted"
	PortingEventSplitEventEventTypePortingOrderLoaUpdated          PortingEventSplitEventEventType = "porting_order.loa_updated"
	PortingEventSplitEventEventTypePortingOrderMessagingChanged    PortingEventSplitEventEventType = "porting_order.messaging_changed"
	PortingEventSplitEventEventTypePortingOrderStatusChanged       PortingEventSplitEventEventType = "porting_order.status_changed"
	PortingEventSplitEventEventTypePortingOrderSharingTokenExpired PortingEventSplitEventEventType = "porting_order.sharing_token_expired"
	PortingEventSplitEventEventTypePortingOrderNewComment          PortingEventSplitEventEventType = "porting_order.new_comment"
	PortingEventSplitEventEventTypePortingOrderSplit               PortingEventSplitEventEventType = "porting_order.split"
)

// The webhook payload for the porting_order.split event
type PortingEventSplitEventPayload struct {
	// The porting order that was split.
	From PortingEventSplitEventPayloadFrom `json:"from"`
	// The list of porting phone numbers that were moved to the new porting order.
	PortingPhoneNumbers []PortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// The new porting order that the phone numbers was moved to.
	To PortingEventSplitEventPayloadTo `json:"to"`
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
func (r PortingEventSplitEventPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventSplitEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The porting order that was split.
type PortingEventSplitEventPayloadFrom struct {
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
func (r PortingEventSplitEventPayloadFrom) RawJSON() string { return r.JSON.raw }
func (r *PortingEventSplitEventPayloadFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventSplitEventPayloadPortingPhoneNumber struct {
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
func (r PortingEventSplitEventPayloadPortingPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *PortingEventSplitEventPayloadPortingPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new porting order that the phone numbers was moved to.
type PortingEventSplitEventPayloadTo struct {
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
func (r PortingEventSplitEventPayloadTo) RawJSON() string { return r.JSON.raw }
func (r *PortingEventSplitEventPayloadTo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the payload generation.
type PortingEventSplitEventPayloadStatus string

const (
	PortingEventSplitEventPayloadStatusCreated   PortingEventSplitEventPayloadStatus = "created"
	PortingEventSplitEventPayloadStatusCompleted PortingEventSplitEventPayloadStatus = "completed"
)

type PortingEventStatusChangedEvent struct {
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
	EventType PortingEventStatusChangedEventEventType `json:"event_type"`
	// The webhook payload for the porting_order.status_changed event
	Payload PortingEventStatusChangedEventPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventStatusChangedEventPayloadStatus `json:"payload_status"`
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
func (r PortingEventStatusChangedEvent) RawJSON() string { return r.JSON.raw }
func (r *PortingEventStatusChangedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventStatusChangedEventEventType string

const (
	PortingEventStatusChangedEventEventTypePortingOrderDeleted             PortingEventStatusChangedEventEventType = "porting_order.deleted"
	PortingEventStatusChangedEventEventTypePortingOrderLoaUpdated          PortingEventStatusChangedEventEventType = "porting_order.loa_updated"
	PortingEventStatusChangedEventEventTypePortingOrderMessagingChanged    PortingEventStatusChangedEventEventType = "porting_order.messaging_changed"
	PortingEventStatusChangedEventEventTypePortingOrderStatusChanged       PortingEventStatusChangedEventEventType = "porting_order.status_changed"
	PortingEventStatusChangedEventEventTypePortingOrderSharingTokenExpired PortingEventStatusChangedEventEventType = "porting_order.sharing_token_expired"
	PortingEventStatusChangedEventEventTypePortingOrderNewComment          PortingEventStatusChangedEventEventType = "porting_order.new_comment"
	PortingEventStatusChangedEventEventTypePortingOrderSplit               PortingEventStatusChangedEventEventType = "porting_order.split"
)

// The webhook payload for the porting_order.status_changed event
type PortingEventStatusChangedEventPayload struct {
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
func (r PortingEventStatusChangedEventPayload) RawJSON() string { return r.JSON.raw }
func (r *PortingEventStatusChangedEventPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the payload generation.
type PortingEventStatusChangedEventPayloadStatus string

const (
	PortingEventStatusChangedEventPayloadStatusCreated   PortingEventStatusChangedEventPayloadStatus = "created"
	PortingEventStatusChangedEventPayloadStatusCompleted PortingEventStatusChangedEventPayloadStatus = "completed"
)

type PortingEventWithoutWebhook struct {
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
	EventType PortingEventWithoutWebhookEventType `json:"event_type"`
	Payload   any                                 `json:"payload" api:"nullable"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus PortingEventWithoutWebhookPayloadStatus `json:"payload_status"`
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
func (r PortingEventWithoutWebhook) RawJSON() string { return r.JSON.raw }
func (r *PortingEventWithoutWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the event type
type PortingEventWithoutWebhookEventType string

const (
	PortingEventWithoutWebhookEventTypePortingOrderDeleted             PortingEventWithoutWebhookEventType = "porting_order.deleted"
	PortingEventWithoutWebhookEventTypePortingOrderLoaUpdated          PortingEventWithoutWebhookEventType = "porting_order.loa_updated"
	PortingEventWithoutWebhookEventTypePortingOrderMessagingChanged    PortingEventWithoutWebhookEventType = "porting_order.messaging_changed"
	PortingEventWithoutWebhookEventTypePortingOrderStatusChanged       PortingEventWithoutWebhookEventType = "porting_order.status_changed"
	PortingEventWithoutWebhookEventTypePortingOrderSharingTokenExpired PortingEventWithoutWebhookEventType = "porting_order.sharing_token_expired"
	PortingEventWithoutWebhookEventTypePortingOrderNewComment          PortingEventWithoutWebhookEventType = "porting_order.new_comment"
	PortingEventWithoutWebhookEventTypePortingOrderSplit               PortingEventWithoutWebhookEventType = "porting_order.split"
)

// The status of the payload generation.
type PortingEventWithoutWebhookPayloadStatus string

const (
	PortingEventWithoutWebhookPayloadStatusCreated   PortingEventWithoutWebhookPayloadStatus = "created"
	PortingEventWithoutWebhookPayloadStatusCompleted PortingEventWithoutWebhookPayloadStatus = "completed"
)

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
// from [PortingEventDeletedPayload], [PortingEventMessagingChangedPayload],
// [PortingEventStatusChangedEvent], [PortingEventNewCommentEvent],
// [PortingEventSplitEvent], [PortingEventWithoutWebhook].
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
	// This field is a union of [PortingEventDeletedPayloadPayload],
	// [PortingEventMessagingChangedPayloadPayload],
	// [PortingEventStatusChangedEventPayload], [PortingEventNewCommentEventPayload],
	// [PortingEventSplitEventPayload], [any]
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

func (u PortingEventGetResponseDataUnion) AsPortingEventDeletedPayload() (v PortingEventDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventMessagingChangedPayload() (v PortingEventMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventStatusChangedEvent() (v PortingEventStatusChangedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventNewCommentEvent() (v PortingEventNewCommentEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventSplitEvent() (v PortingEventSplitEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventGetResponseDataUnion) AsPortingEventWithoutWebhook() (v PortingEventWithoutWebhook) {
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
// will be valid: OfPortingEventWithoutWebhookPayload]
type PortingEventGetResponseDataUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPortingEventWithoutWebhookPayload any    `json:",inline"`
	ID                                  string `json:"id"`
	// This field is from variant [PortingEventDeletedPayloadPayload].
	CreatedAt         time.Time `json:"created_at"`
	CustomerReference string    `json:"customer_reference"`
	// This field is from variant [PortingEventDeletedPayloadPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant [PortingEventDeletedPayloadPayload].
	RecordType string    `json:"record_type"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant [PortingEventMessagingChangedPayloadPayload].
	Messaging  PortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
	SupportKey string                                              `json:"support_key"`
	// This field is from variant [PortingEventStatusChangedEventPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant [PortingEventStatusChangedEventPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant [PortingEventNewCommentEventPayload].
	Comment PortingEventNewCommentEventPayloadComment `json:"comment"`
	// This field is from variant [PortingEventNewCommentEventPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant [PortingEventSplitEventPayload].
	From PortingEventSplitEventPayloadFrom `json:"from"`
	// This field is from variant [PortingEventSplitEventPayload].
	PortingPhoneNumbers []PortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant [PortingEventSplitEventPayload].
	To   PortingEventSplitEventPayloadTo `json:"to"`
	JSON struct {
		OfPortingEventWithoutWebhookPayload respjson.Field
		ID                                  respjson.Field
		CreatedAt                           respjson.Field
		CustomerReference                   respjson.Field
		DeletedAt                           respjson.Field
		RecordType                          respjson.Field
		UpdatedAt                           respjson.Field
		Messaging                           respjson.Field
		SupportKey                          respjson.Field
		Status                              respjson.Field
		WebhookURL                          respjson.Field
		Comment                             respjson.Field
		PortingOrderID                      respjson.Field
		From                                respjson.Field
		PortingPhoneNumbers                 respjson.Field
		To                                  respjson.Field
		raw                                 string
	} `json:"-"`
}

func (r *PortingEventGetResponseDataUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortingEventListResponseUnion contains all possible properties and values from
// [PortingEventDeletedPayload], [PortingEventMessagingChangedPayload],
// [PortingEventStatusChangedEvent], [PortingEventNewCommentEvent],
// [PortingEventSplitEvent], [PortingEventWithoutWebhook].
//
// Use the [PortingEventListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortingEventListResponseUnion struct {
	ID                           string   `json:"id"`
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// Any of nil, nil, nil, nil, nil, nil.
	EventType string `json:"event_type"`
	// This field is a union of [PortingEventDeletedPayloadPayload],
	// [PortingEventMessagingChangedPayloadPayload],
	// [PortingEventStatusChangedEventPayload], [PortingEventNewCommentEventPayload],
	// [PortingEventSplitEventPayload], [any]
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

func (u PortingEventListResponseUnion) AsPortingEventDeletedPayload() (v PortingEventDeletedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventMessagingChangedPayload() (v PortingEventMessagingChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventStatusChangedEvent() (v PortingEventStatusChangedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventNewCommentEvent() (v PortingEventNewCommentEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventSplitEvent() (v PortingEventSplitEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortingEventListResponseUnion) AsPortingEventWithoutWebhook() (v PortingEventWithoutWebhook) {
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
// will be valid: OfPortingEventWithoutWebhookPayload]
type PortingEventListResponseUnionPayload struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPortingEventWithoutWebhookPayload any    `json:",inline"`
	ID                                  string `json:"id"`
	// This field is from variant [PortingEventDeletedPayloadPayload].
	CreatedAt         time.Time `json:"created_at"`
	CustomerReference string    `json:"customer_reference"`
	// This field is from variant [PortingEventDeletedPayloadPayload].
	DeletedAt time.Time `json:"deleted_at"`
	// This field is from variant [PortingEventDeletedPayloadPayload].
	RecordType string    `json:"record_type"`
	UpdatedAt  time.Time `json:"updated_at"`
	// This field is from variant [PortingEventMessagingChangedPayloadPayload].
	Messaging  PortingEventMessagingChangedPayloadPayloadMessaging `json:"messaging"`
	SupportKey string                                              `json:"support_key"`
	// This field is from variant [PortingEventStatusChangedEventPayload].
	Status shared.PortingOrderStatus `json:"status"`
	// This field is from variant [PortingEventStatusChangedEventPayload].
	WebhookURL string `json:"webhook_url"`
	// This field is from variant [PortingEventNewCommentEventPayload].
	Comment PortingEventNewCommentEventPayloadComment `json:"comment"`
	// This field is from variant [PortingEventNewCommentEventPayload].
	PortingOrderID string `json:"porting_order_id"`
	// This field is from variant [PortingEventSplitEventPayload].
	From PortingEventSplitEventPayloadFrom `json:"from"`
	// This field is from variant [PortingEventSplitEventPayload].
	PortingPhoneNumbers []PortingEventSplitEventPayloadPortingPhoneNumber `json:"porting_phone_numbers"`
	// This field is from variant [PortingEventSplitEventPayload].
	To   PortingEventSplitEventPayloadTo `json:"to"`
	JSON struct {
		OfPortingEventWithoutWebhookPayload respjson.Field
		ID                                  respjson.Field
		CreatedAt                           respjson.Field
		CustomerReference                   respjson.Field
		DeletedAt                           respjson.Field
		RecordType                          respjson.Field
		UpdatedAt                           respjson.Field
		Messaging                           respjson.Field
		SupportKey                          respjson.Field
		Status                              respjson.Field
		WebhookURL                          respjson.Field
		Comment                             respjson.Field
		PortingOrderID                      respjson.Field
		From                                respjson.Field
		PortingPhoneNumbers                 respjson.Field
		To                                  respjson.Field
		raw                                 string
	} `json:"-"`
}

func (r *PortingEventListResponseUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingEventListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[type],
	// filter[porting_order_id], filter[created_at][gte], filter[created_at][lte]
	Filter PortingEventListParamsFilter `query:"filter,omitzero" json:"-"`
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
