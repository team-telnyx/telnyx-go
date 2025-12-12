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
)

// PortoutEventService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortoutEventService] method instead.
type PortoutEventService struct {
	Options []option.RequestOption
}

// NewPortoutEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPortoutEventService(opts ...option.RequestOption) (r PortoutEventService) {
	r = PortoutEventService{}
	r.Options = opts
	return
}

// Show a specific port-out event.
func (r *PortoutEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PortoutEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/events/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all port-out events.
func (r *PortoutEventService) List(ctx context.Context, query PortoutEventListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[PortoutEventListResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "portouts/events"
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

// Returns a list of all port-out events.
func (r *PortoutEventService) ListAutoPaging(ctx context.Context, query PortoutEventListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[PortoutEventListResponseUnion] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Republish a specific port-out event.
func (r *PortoutEventService) Republish(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/events/%s/republish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type PortoutEventGetResponse struct {
	Data PortoutEventGetResponseDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortoutEventGetResponseDataUnion contains all possible properties and values
// from [PortoutEventGetResponseDataWebhookPortoutStatusChanged],
// [PortoutEventGetResponseDataWebhookPortoutNewComment],
// [PortoutEventGetResponseDataWebhookPortoutFocDateChanged].
//
// Use the [PortoutEventGetResponseDataUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortoutEventGetResponseDataUnion struct {
	ID                           string    `json:"id"`
	AvailableNotificationMethods []string  `json:"available_notification_methods"`
	CreatedAt                    time.Time `json:"created_at"`
	// Any of nil, nil, nil.
	EventType string `json:"event_type"`
	// This field is a union of
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload],
	// [PortoutEventGetResponseDataWebhookPortoutNewCommentPayload],
	// [PortoutEventGetResponseDataWebhookPortoutFocDateChangedPayload]
	Payload       PortoutEventGetResponseDataUnionPayload `json:"payload"`
	PayloadStatus string                                  `json:"payload_status"`
	PortoutID     string                                  `json:"portout_id"`
	RecordType    string                                  `json:"record_type"`
	UpdatedAt     time.Time                               `json:"updated_at"`
	JSON          struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		CreatedAt                    respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PortoutEventGetResponseDataUnion) AsPortoutEventGetResponseDataWebhookPortoutStatusChanged() (v PortoutEventGetResponseDataWebhookPortoutStatusChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventGetResponseDataUnion) AsPortoutEventGetResponseDataWebhookPortoutNewComment() (v PortoutEventGetResponseDataWebhookPortoutNewComment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventGetResponseDataUnion) AsPortoutEventGetResponseDataWebhookPortoutFocDateChanged() (v PortoutEventGetResponseDataWebhookPortoutFocDateChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortoutEventGetResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *PortoutEventGetResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortoutEventGetResponseDataUnionPayload is an implicit subunion of
// [PortoutEventGetResponseDataUnion]. PortoutEventGetResponseDataUnionPayload
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PortoutEventGetResponseDataUnion].
type PortoutEventGetResponseDataUnionPayload struct {
	ID string `json:"id"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	AttemptedPin string `json:"attempted_pin"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	CarrierName string `json:"carrier_name"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	PhoneNumbers []string `json:"phone_numbers"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	RejectionReason string `json:"rejection_reason"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	Spid string `json:"spid"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	Status string `json:"status"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload].
	SubscriberName string `json:"subscriber_name"`
	UserID         string `json:"user_id"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutNewCommentPayload].
	Comment string `json:"comment"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutNewCommentPayload].
	PortoutID string `json:"portout_id"`
	// This field is from variant
	// [PortoutEventGetResponseDataWebhookPortoutFocDateChangedPayload].
	FocDate time.Time `json:"foc_date"`
	JSON    struct {
		ID              respjson.Field
		AttemptedPin    respjson.Field
		CarrierName     respjson.Field
		PhoneNumbers    respjson.Field
		RejectionReason respjson.Field
		Spid            respjson.Field
		Status          respjson.Field
		SubscriberName  respjson.Field
		UserID          respjson.Field
		Comment         respjson.Field
		PortoutID       respjson.Field
		FocDate         respjson.Field
		raw             string
	} `json:"-"`
}

func (r *PortoutEventGetResponseDataUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventGetResponseDataWebhookPortoutStatusChanged struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "portout.status_changed", "portout.foc_date_changed",
	// "portout.new_comment".
	EventType string `json:"event_type"`
	// The webhook payload for the portout.status_changed event
	Payload PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the port-out order associated with the event.
	PortoutID string `json:"portout_id" format:"uuid"`
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
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponseDataWebhookPortoutStatusChanged) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventGetResponseDataWebhookPortoutStatusChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.status_changed event
type PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload struct {
	// Identifies the port out that was moved.
	ID string `json:"id" format:"uuid"`
	// The PIN that was attempted to be used to authorize the port out.
	AttemptedPin string `json:"attempted_pin"`
	// Carrier the number will be ported out to
	CarrierName string `json:"carrier_name"`
	// Phone numbers associated with this port-out order
	PhoneNumbers []string `json:"phone_numbers"`
	// The reason why the order is being rejected by the user. If the order is
	// authorized, this field can be left null
	RejectionReason string `json:"rejection_reason,nullable"`
	// The new carrier SPID.
	Spid string `json:"spid"`
	// The new status of the port out.
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	Status string `json:"status"`
	// The name of the port-out's end user.
	SubscriberName string `json:"subscriber_name"`
	// Identifies the user that the port-out order belongs to.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttemptedPin    respjson.Field
		CarrierName     respjson.Field
		PhoneNumbers    respjson.Field
		RejectionReason respjson.Field
		Spid            respjson.Field
		Status          respjson.Field
		SubscriberName  respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventGetResponseDataWebhookPortoutStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventGetResponseDataWebhookPortoutNewComment struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "portout.status_changed", "portout.foc_date_changed",
	// "portout.new_comment".
	EventType string `json:"event_type"`
	// The webhook payload for the portout.new_comment event
	Payload PortoutEventGetResponseDataWebhookPortoutNewCommentPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the port-out order associated with the event.
	PortoutID string `json:"portout_id" format:"uuid"`
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
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponseDataWebhookPortoutNewComment) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventGetResponseDataWebhookPortoutNewComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.new_comment event
type PortoutEventGetResponseDataWebhookPortoutNewCommentPayload struct {
	// Identifies the comment that was added to the port-out order.
	ID string `json:"id" format:"uuid"`
	// The body of the comment.
	Comment string `json:"comment"`
	// Identifies the port-out order that the comment was added to.
	PortoutID string `json:"portout_id" format:"uuid"`
	// Identifies the user that added the comment.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Comment     respjson.Field
		PortoutID   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponseDataWebhookPortoutNewCommentPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventGetResponseDataWebhookPortoutNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventGetResponseDataWebhookPortoutFocDateChanged struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "portout.status_changed", "portout.foc_date_changed",
	// "portout.new_comment".
	EventType string `json:"event_type"`
	// The webhook payload for the portout.foc_date_changed event
	Payload PortoutEventGetResponseDataWebhookPortoutFocDateChangedPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the port-out order associated with the event.
	PortoutID string `json:"portout_id" format:"uuid"`
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
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponseDataWebhookPortoutFocDateChanged) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventGetResponseDataWebhookPortoutFocDateChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.foc_date_changed event
type PortoutEventGetResponseDataWebhookPortoutFocDateChangedPayload struct {
	// Identifies the port-out order that have the FOC date changed.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating the new FOC date.
	FocDate time.Time `json:"foc_date" format:"date-time"`
	// Identifies the organization that port-out order belongs to.
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FocDate     respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventGetResponseDataWebhookPortoutFocDateChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventGetResponseDataWebhookPortoutFocDateChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortoutEventListResponseUnion contains all possible properties and values from
// [PortoutEventListResponseWebhookPortoutStatusChanged],
// [PortoutEventListResponseWebhookPortoutNewComment],
// [PortoutEventListResponseWebhookPortoutFocDateChanged].
//
// Use the [PortoutEventListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortoutEventListResponseUnion struct {
	ID                           string    `json:"id"`
	AvailableNotificationMethods []string  `json:"available_notification_methods"`
	CreatedAt                    time.Time `json:"created_at"`
	// Any of nil, nil, nil.
	EventType string `json:"event_type"`
	// This field is a union of
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload],
	// [PortoutEventListResponseWebhookPortoutNewCommentPayload],
	// [PortoutEventListResponseWebhookPortoutFocDateChangedPayload]
	Payload       PortoutEventListResponseUnionPayload `json:"payload"`
	PayloadStatus string                               `json:"payload_status"`
	PortoutID     string                               `json:"portout_id"`
	RecordType    string                               `json:"record_type"`
	UpdatedAt     time.Time                            `json:"updated_at"`
	JSON          struct {
		ID                           respjson.Field
		AvailableNotificationMethods respjson.Field
		CreatedAt                    respjson.Field
		EventType                    respjson.Field
		Payload                      respjson.Field
		PayloadStatus                respjson.Field
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		raw                          string
	} `json:"-"`
}

func (u PortoutEventListResponseUnion) AsPortoutEventListResponseWebhookPortoutStatusChanged() (v PortoutEventListResponseWebhookPortoutStatusChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventListResponseUnion) AsPortoutEventListResponseWebhookPortoutNewComment() (v PortoutEventListResponseWebhookPortoutNewComment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventListResponseUnion) AsPortoutEventListResponseWebhookPortoutFocDateChanged() (v PortoutEventListResponseWebhookPortoutFocDateChanged) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortoutEventListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *PortoutEventListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortoutEventListResponseUnionPayload is an implicit subunion of
// [PortoutEventListResponseUnion]. PortoutEventListResponseUnionPayload provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [PortoutEventListResponseUnion].
type PortoutEventListResponseUnionPayload struct {
	ID string `json:"id"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	AttemptedPin string `json:"attempted_pin"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	CarrierName string `json:"carrier_name"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	PhoneNumbers []string `json:"phone_numbers"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	RejectionReason string `json:"rejection_reason"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	Spid string `json:"spid"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	Status string `json:"status"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutStatusChangedPayload].
	SubscriberName string `json:"subscriber_name"`
	UserID         string `json:"user_id"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutNewCommentPayload].
	Comment string `json:"comment"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutNewCommentPayload].
	PortoutID string `json:"portout_id"`
	// This field is from variant
	// [PortoutEventListResponseWebhookPortoutFocDateChangedPayload].
	FocDate time.Time `json:"foc_date"`
	JSON    struct {
		ID              respjson.Field
		AttemptedPin    respjson.Field
		CarrierName     respjson.Field
		PhoneNumbers    respjson.Field
		RejectionReason respjson.Field
		Spid            respjson.Field
		Status          respjson.Field
		SubscriberName  respjson.Field
		UserID          respjson.Field
		Comment         respjson.Field
		PortoutID       respjson.Field
		FocDate         respjson.Field
		raw             string
	} `json:"-"`
}

func (r *PortoutEventListResponseUnionPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventListResponseWebhookPortoutStatusChanged struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "portout.status_changed", "portout.foc_date_changed",
	// "portout.new_comment".
	EventType string `json:"event_type"`
	// The webhook payload for the portout.status_changed event
	Payload PortoutEventListResponseWebhookPortoutStatusChangedPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the port-out order associated with the event.
	PortoutID string `json:"portout_id" format:"uuid"`
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
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventListResponseWebhookPortoutStatusChanged) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventListResponseWebhookPortoutStatusChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.status_changed event
type PortoutEventListResponseWebhookPortoutStatusChangedPayload struct {
	// Identifies the port out that was moved.
	ID string `json:"id" format:"uuid"`
	// The PIN that was attempted to be used to authorize the port out.
	AttemptedPin string `json:"attempted_pin"`
	// Carrier the number will be ported out to
	CarrierName string `json:"carrier_name"`
	// Phone numbers associated with this port-out order
	PhoneNumbers []string `json:"phone_numbers"`
	// The reason why the order is being rejected by the user. If the order is
	// authorized, this field can be left null
	RejectionReason string `json:"rejection_reason,nullable"`
	// The new carrier SPID.
	Spid string `json:"spid"`
	// The new status of the port out.
	//
	// Any of "pending", "authorized", "ported", "rejected", "rejected-pending",
	// "canceled".
	Status string `json:"status"`
	// The name of the port-out's end user.
	SubscriberName string `json:"subscriber_name"`
	// Identifies the user that the port-out order belongs to.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttemptedPin    respjson.Field
		CarrierName     respjson.Field
		PhoneNumbers    respjson.Field
		RejectionReason respjson.Field
		Spid            respjson.Field
		Status          respjson.Field
		SubscriberName  respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventListResponseWebhookPortoutStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventListResponseWebhookPortoutStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventListResponseWebhookPortoutNewComment struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "portout.status_changed", "portout.foc_date_changed",
	// "portout.new_comment".
	EventType string `json:"event_type"`
	// The webhook payload for the portout.new_comment event
	Payload PortoutEventListResponseWebhookPortoutNewCommentPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the port-out order associated with the event.
	PortoutID string `json:"portout_id" format:"uuid"`
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
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventListResponseWebhookPortoutNewComment) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventListResponseWebhookPortoutNewComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.new_comment event
type PortoutEventListResponseWebhookPortoutNewCommentPayload struct {
	// Identifies the comment that was added to the port-out order.
	ID string `json:"id" format:"uuid"`
	// The body of the comment.
	Comment string `json:"comment"`
	// Identifies the port-out order that the comment was added to.
	PortoutID string `json:"portout_id" format:"uuid"`
	// Identifies the user that added the comment.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Comment     respjson.Field
		PortoutID   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventListResponseWebhookPortoutNewCommentPayload) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventListResponseWebhookPortoutNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventListResponseWebhookPortoutFocDateChanged struct {
	// Uniquely identifies the event.
	ID string `json:"id" format:"uuid"`
	// Indicates the notification methods used.
	//
	// Any of "email", "webhook".
	AvailableNotificationMethods []string `json:"available_notification_methods"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifies the event type
	//
	// Any of "portout.status_changed", "portout.foc_date_changed",
	// "portout.new_comment".
	EventType string `json:"event_type"`
	// The webhook payload for the portout.foc_date_changed event
	Payload PortoutEventListResponseWebhookPortoutFocDateChangedPayload `json:"payload"`
	// The status of the payload generation.
	//
	// Any of "created", "completed".
	PayloadStatus string `json:"payload_status"`
	// Identifies the port-out order associated with the event.
	PortoutID string `json:"portout_id" format:"uuid"`
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
		PortoutID                    respjson.Field
		RecordType                   respjson.Field
		UpdatedAt                    respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventListResponseWebhookPortoutFocDateChanged) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventListResponseWebhookPortoutFocDateChanged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.foc_date_changed event
type PortoutEventListResponseWebhookPortoutFocDateChangedPayload struct {
	// Identifies the port-out order that have the FOC date changed.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating the new FOC date.
	FocDate time.Time `json:"foc_date" format:"date-time"`
	// Identifies the organization that port-out order belongs to.
	UserID string `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FocDate     respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortoutEventListResponseWebhookPortoutFocDateChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventListResponseWebhookPortoutFocDateChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[event_type], filter[portout_id], filter[created_at]
	Filter PortoutEventListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page PortoutEventListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutEventListParams]'s query parameters as `url.Values`.
func (r PortoutEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[event_type], filter[portout_id], filter[created_at]
type PortoutEventListParamsFilter struct {
	// Filter by port-out order ID.
	PortoutID param.Opt[string] `query:"portout_id,omitzero" format:"uuid" json:"-"`
	// Filter by created_at date range using nested operations
	CreatedAt PortoutEventListParamsFilterCreatedAt `query:"created_at,omitzero" json:"-"`
	// Filter by event type.
	//
	// Any of "portout.status_changed", "portout.new_comment",
	// "portout.foc_date_changed".
	EventType string `query:"event_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutEventListParamsFilter]'s query parameters as
// `url.Values`.
func (r PortoutEventListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by created_at date range using nested operations
type PortoutEventListParamsFilterCreatedAt struct {
	// Filter by created at greater than or equal to.
	Gte param.Opt[time.Time] `query:"gte,omitzero" format:"date-time" json:"-"`
	// Filter by created at less than or equal to.
	Lte param.Opt[time.Time] `query:"lte,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutEventListParamsFilterCreatedAt]'s query parameters
// as `url.Values`.
func (r PortoutEventListParamsFilterCreatedAt) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type PortoutEventListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PortoutEventListParamsPage]'s query parameters as
// `url.Values`.
func (r PortoutEventListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
