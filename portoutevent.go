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
func (r *PortoutEventService) List(ctx context.Context, query PortoutEventListParams, opts ...option.RequestOption) (res *PortoutEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "portouts/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Republish a specific port-out event.
func (r *PortoutEventService) Republish(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("portouts/events/%s/republish", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

type PortoutEventGetResponse struct {
	Data PortoutEventGetResponseData `json:"data"`
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

type PortoutEventGetResponseData struct {
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
	Payload PortoutEventGetResponseDataPayloadUnion `json:"payload"`
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
func (r PortoutEventGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortoutEventGetResponseDataPayloadUnion contains all possible properties and
// values from
// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload],
// [PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload],
// [PortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortoutEventGetResponseDataPayloadUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	AttemptedPin string `json:"attempted_pin"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	CarrierName string `json:"carrier_name"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	PhoneNumbers []string `json:"phone_numbers"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	RejectionReason string `json:"rejection_reason"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	Spid string `json:"spid"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	Status string `json:"status"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload].
	SubscriberName string `json:"subscriber_name"`
	UserID         string `json:"user_id"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload].
	Comment string `json:"comment"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload].
	PortoutID string `json:"portout_id"`
	// This field is from variant
	// [PortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload].
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

func (u PortoutEventGetResponseDataPayloadUnion) AsPortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload() (v PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventGetResponseDataPayloadUnion) AsPortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload() (v PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventGetResponseDataPayloadUnion) AsPortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload() (v PortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortoutEventGetResponseDataPayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *PortoutEventGetResponseDataPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.status_changed event
type PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload struct {
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
	RejectionReason string `json:"rejection_reason"`
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
func (r PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventGetResponseDataPayloadWebhookPortoutStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.new_comment event
type PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload struct {
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
func (r PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventGetResponseDataPayloadWebhookPortoutNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.foc_date_changed event
type PortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload struct {
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
func (r PortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventGetResponseDataPayloadWebhookPortoutFocDateChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventListResponse struct {
	Data []PortoutEventListResponseData `json:"data"`
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
func (r PortoutEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortoutEventListResponseData struct {
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
	Payload PortoutEventListResponseDataPayloadUnion `json:"payload"`
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
func (r PortoutEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortoutEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PortoutEventListResponseDataPayloadUnion contains all possible properties and
// values from
// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload],
// [PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload],
// [PortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PortoutEventListResponseDataPayloadUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	AttemptedPin string `json:"attempted_pin"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	CarrierName string `json:"carrier_name"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	PhoneNumbers []string `json:"phone_numbers"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	RejectionReason string `json:"rejection_reason"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	Spid string `json:"spid"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	Status string `json:"status"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload].
	SubscriberName string `json:"subscriber_name"`
	UserID         string `json:"user_id"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload].
	Comment string `json:"comment"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload].
	PortoutID string `json:"portout_id"`
	// This field is from variant
	// [PortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload].
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

func (u PortoutEventListResponseDataPayloadUnion) AsPortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload() (v PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventListResponseDataPayloadUnion) AsPortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload() (v PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PortoutEventListResponseDataPayloadUnion) AsPortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload() (v PortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PortoutEventListResponseDataPayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *PortoutEventListResponseDataPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.status_changed event
type PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload struct {
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
	RejectionReason string `json:"rejection_reason"`
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
func (r PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventListResponseDataPayloadWebhookPortoutStatusChangedPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.new_comment event
type PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload struct {
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
func (r PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventListResponseDataPayloadWebhookPortoutNewCommentPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The webhook payload for the portout.foc_date_changed event
type PortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload struct {
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
func (r PortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *PortoutEventListResponseDataPayloadWebhookPortoutFocDateChangedPayload) UnmarshalJSON(data []byte) error {
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

func init() {
	apijson.RegisterFieldValidator[PortoutEventListParamsFilter](
		"event_type", "portout.status_changed", "portout.new_comment", "portout.foc_date_changed",
	)
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
