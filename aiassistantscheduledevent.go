// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// AIAssistantScheduledEventService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantScheduledEventService] method instead.
type AIAssistantScheduledEventService struct {
	Options []option.RequestOption
}

// NewAIAssistantScheduledEventService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIAssistantScheduledEventService(opts ...option.RequestOption) (r AIAssistantScheduledEventService) {
	r = AIAssistantScheduledEventService{}
	r.Options = opts
	return
}

// Create a scheduled event for an assistant
func (r *AIAssistantScheduledEventService) New(ctx context.Context, assistantID string, body AIAssistantScheduledEventNewParams, opts ...option.RequestOption) (res *ScheduledEventResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/scheduled_events", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a scheduled event by event ID
func (r *AIAssistantScheduledEventService) Get(ctx context.Context, eventID string, query AIAssistantScheduledEventGetParams, opts ...option.RequestOption) (res *ScheduledEventResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if query.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/scheduled_events/%s", query.AssistantID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get scheduled events for an assistant with pagination and filtering
func (r *AIAssistantScheduledEventService) List(ctx context.Context, assistantID string, query AIAssistantScheduledEventListParams, opts ...option.RequestOption) (res *AIAssistantScheduledEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/scheduled_events", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// If the event is pending, this will cancel the event. Otherwise, this will simply
// remove the record of the event.
func (r *AIAssistantScheduledEventService) Delete(ctx context.Context, eventID string, body AIAssistantScheduledEventDeleteParams, opts ...option.RequestOption) (res *AIAssistantScheduledEventDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/scheduled_events/%s", body.AssistantID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ConversationChannelType string

const (
	ConversationChannelTypePhoneCall ConversationChannelType = "phone_call"
	ConversationChannelTypeSMSChat   ConversationChannelType = "sms_chat"
)

type EventStatus string

const (
	EventStatusPending    EventStatus = "pending"
	EventStatusInProgress EventStatus = "in_progress"
	EventStatusCompleted  EventStatus = "completed"
	EventStatusFailed     EventStatus = "failed"
)

// ScheduledEventResponseUnion contains all possible properties and values from
// [ScheduledPhoneCallEventResponse], [ScheduledSMSEventResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ScheduledEventResponseUnion struct {
	AssistantID              string    `json:"assistant_id"`
	ScheduledAtFixedDatetime time.Time `json:"scheduled_at_fixed_datetime"`
	TelnyxAgentTarget        string    `json:"telnyx_agent_target"`
	// This field is from variant [ScheduledPhoneCallEventResponse].
	TelnyxConversationChannel ConversationChannelType `json:"telnyx_conversation_channel"`
	TelnyxEndUserTarget       string                  `json:"telnyx_end_user_target"`
	ConversationID            string                  `json:"conversation_id"`
	// This field is a union of
	// [map[string]ScheduledPhoneCallEventResponseConversationMetadataUnion],
	// [map[string]ScheduledSMSEventResponseConversationMetadataUnion]
	ConversationMetadata ScheduledEventResponseUnionConversationMetadata `json:"conversation_metadata"`
	CreatedAt            time.Time                                       `json:"created_at"`
	Errors               []string                                        `json:"errors"`
	// This field is from variant [ScheduledPhoneCallEventResponse].
	RetryAttempts    int64  `json:"retry_attempts"`
	RetryCount       int64  `json:"retry_count"`
	ScheduledEventID string `json:"scheduled_event_id"`
	// This field is from variant [ScheduledPhoneCallEventResponse].
	Status EventStatus `json:"status"`
	// This field is from variant [ScheduledSMSEventResponse].
	Text string `json:"text"`
	JSON struct {
		AssistantID               respjson.Field
		ScheduledAtFixedDatetime  respjson.Field
		TelnyxAgentTarget         respjson.Field
		TelnyxConversationChannel respjson.Field
		TelnyxEndUserTarget       respjson.Field
		ConversationID            respjson.Field
		ConversationMetadata      respjson.Field
		CreatedAt                 respjson.Field
		Errors                    respjson.Field
		RetryAttempts             respjson.Field
		RetryCount                respjson.Field
		ScheduledEventID          respjson.Field
		Status                    respjson.Field
		Text                      respjson.Field
		raw                       string
	} `json:"-"`
}

func (u ScheduledEventResponseUnion) AsScheduledPhoneCallEventResponse() (v ScheduledPhoneCallEventResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScheduledEventResponseUnion) AsScheduledSMSEventResponse() (v ScheduledSMSEventResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScheduledEventResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ScheduledEventResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScheduledEventResponseUnionConversationMetadata is an implicit subunion of
// [ScheduledEventResponseUnion]. ScheduledEventResponseUnionConversationMetadata
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ScheduledEventResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type ScheduledEventResponseUnionConversationMetadata struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (r *ScheduledEventResponseUnionConversationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduledPhoneCallEventResponse struct {
	AssistantID              string    `json:"assistant_id,required"`
	ScheduledAtFixedDatetime time.Time `json:"scheduled_at_fixed_datetime,required" format:"date-time"`
	TelnyxAgentTarget        string    `json:"telnyx_agent_target,required"`
	// Any of "phone_call", "sms_chat".
	TelnyxConversationChannel ConversationChannelType                                             `json:"telnyx_conversation_channel,required"`
	TelnyxEndUserTarget       string                                                              `json:"telnyx_end_user_target,required"`
	ConversationID            string                                                              `json:"conversation_id"`
	ConversationMetadata      map[string]ScheduledPhoneCallEventResponseConversationMetadataUnion `json:"conversation_metadata"`
	CreatedAt                 time.Time                                                           `json:"created_at" format:"date-time"`
	Errors                    []string                                                            `json:"errors"`
	RetryAttempts             int64                                                               `json:"retry_attempts"`
	RetryCount                int64                                                               `json:"retry_count"`
	ScheduledEventID          string                                                              `json:"scheduled_event_id"`
	// Any of "pending", "in_progress", "completed", "failed".
	Status EventStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantID               respjson.Field
		ScheduledAtFixedDatetime  respjson.Field
		TelnyxAgentTarget         respjson.Field
		TelnyxConversationChannel respjson.Field
		TelnyxEndUserTarget       respjson.Field
		ConversationID            respjson.Field
		ConversationMetadata      respjson.Field
		CreatedAt                 respjson.Field
		Errors                    respjson.Field
		RetryAttempts             respjson.Field
		RetryCount                respjson.Field
		ScheduledEventID          respjson.Field
		Status                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledPhoneCallEventResponse) RawJSON() string { return r.JSON.raw }
func (r *ScheduledPhoneCallEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScheduledPhoneCallEventResponseConversationMetadataUnion contains all possible
// properties and values from [string], [int64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type ScheduledPhoneCallEventResponseConversationMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ScheduledPhoneCallEventResponseConversationMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScheduledPhoneCallEventResponseConversationMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScheduledPhoneCallEventResponseConversationMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScheduledPhoneCallEventResponseConversationMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ScheduledPhoneCallEventResponseConversationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduledSMSEventResponse struct {
	AssistantID              string    `json:"assistant_id,required"`
	ScheduledAtFixedDatetime time.Time `json:"scheduled_at_fixed_datetime,required" format:"date-time"`
	TelnyxAgentTarget        string    `json:"telnyx_agent_target,required"`
	// Any of "phone_call", "sms_chat".
	TelnyxConversationChannel ConversationChannelType                                       `json:"telnyx_conversation_channel,required"`
	TelnyxEndUserTarget       string                                                        `json:"telnyx_end_user_target,required"`
	Text                      string                                                        `json:"text,required"`
	ConversationID            string                                                        `json:"conversation_id"`
	ConversationMetadata      map[string]ScheduledSMSEventResponseConversationMetadataUnion `json:"conversation_metadata"`
	CreatedAt                 time.Time                                                     `json:"created_at" format:"date-time"`
	Errors                    []string                                                      `json:"errors"`
	RetryCount                int64                                                         `json:"retry_count"`
	ScheduledEventID          string                                                        `json:"scheduled_event_id"`
	// Any of "pending", "in_progress", "completed", "failed".
	Status EventStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantID               respjson.Field
		ScheduledAtFixedDatetime  respjson.Field
		TelnyxAgentTarget         respjson.Field
		TelnyxConversationChannel respjson.Field
		TelnyxEndUserTarget       respjson.Field
		Text                      respjson.Field
		ConversationID            respjson.Field
		ConversationMetadata      respjson.Field
		CreatedAt                 respjson.Field
		Errors                    respjson.Field
		RetryCount                respjson.Field
		ScheduledEventID          respjson.Field
		Status                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledSMSEventResponse) RawJSON() string { return r.JSON.raw }
func (r *ScheduledSMSEventResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScheduledSMSEventResponseConversationMetadataUnion contains all possible
// properties and values from [string], [int64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type ScheduledSMSEventResponseConversationMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ScheduledSMSEventResponseConversationMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScheduledSMSEventResponseConversationMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScheduledSMSEventResponseConversationMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScheduledSMSEventResponseConversationMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ScheduledSMSEventResponseConversationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantScheduledEventListResponse struct {
	Data []AIAssistantScheduledEventListResponseDataUnion `json:"data,required"`
	Meta Meta                                             `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantScheduledEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantScheduledEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AIAssistantScheduledEventListResponseDataUnion contains all possible properties
// and values from [ScheduledPhoneCallEventResponse], [ScheduledSMSEventResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AIAssistantScheduledEventListResponseDataUnion struct {
	AssistantID              string    `json:"assistant_id"`
	ScheduledAtFixedDatetime time.Time `json:"scheduled_at_fixed_datetime"`
	TelnyxAgentTarget        string    `json:"telnyx_agent_target"`
	// This field is from variant [ScheduledPhoneCallEventResponse].
	TelnyxConversationChannel ConversationChannelType `json:"telnyx_conversation_channel"`
	TelnyxEndUserTarget       string                  `json:"telnyx_end_user_target"`
	ConversationID            string                  `json:"conversation_id"`
	// This field is a union of
	// [map[string]ScheduledPhoneCallEventResponseConversationMetadataUnion],
	// [map[string]ScheduledSMSEventResponseConversationMetadataUnion]
	ConversationMetadata AIAssistantScheduledEventListResponseDataUnionConversationMetadata `json:"conversation_metadata"`
	CreatedAt            time.Time                                                          `json:"created_at"`
	Errors               []string                                                           `json:"errors"`
	// This field is from variant [ScheduledPhoneCallEventResponse].
	RetryAttempts    int64  `json:"retry_attempts"`
	RetryCount       int64  `json:"retry_count"`
	ScheduledEventID string `json:"scheduled_event_id"`
	// This field is from variant [ScheduledPhoneCallEventResponse].
	Status EventStatus `json:"status"`
	// This field is from variant [ScheduledSMSEventResponse].
	Text string `json:"text"`
	JSON struct {
		AssistantID               respjson.Field
		ScheduledAtFixedDatetime  respjson.Field
		TelnyxAgentTarget         respjson.Field
		TelnyxConversationChannel respjson.Field
		TelnyxEndUserTarget       respjson.Field
		ConversationID            respjson.Field
		ConversationMetadata      respjson.Field
		CreatedAt                 respjson.Field
		Errors                    respjson.Field
		RetryAttempts             respjson.Field
		RetryCount                respjson.Field
		ScheduledEventID          respjson.Field
		Status                    respjson.Field
		Text                      respjson.Field
		raw                       string
	} `json:"-"`
}

func (u AIAssistantScheduledEventListResponseDataUnion) AsScheduledPhoneCallEventResponse() (v ScheduledPhoneCallEventResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIAssistantScheduledEventListResponseDataUnion) AsScheduledSMSEventResponse() (v ScheduledSMSEventResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AIAssistantScheduledEventListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *AIAssistantScheduledEventListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AIAssistantScheduledEventListResponseDataUnionConversationMetadata is an
// implicit subunion of [AIAssistantScheduledEventListResponseDataUnion].
// AIAssistantScheduledEventListResponseDataUnionConversationMetadata provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AIAssistantScheduledEventListResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type AIAssistantScheduledEventListResponseDataUnionConversationMetadata struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (r *AIAssistantScheduledEventListResponseDataUnionConversationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantScheduledEventDeleteResponse = any

type AIAssistantScheduledEventNewParams struct {
	// The datetime at which the event should be scheduled. Formatted as ISO 8601.
	ScheduledAtFixedDatetime time.Time `json:"scheduled_at_fixed_datetime,required" format:"date-time"`
	// The phone number, SIP URI, to schedule the call or text from.
	TelnyxAgentTarget string `json:"telnyx_agent_target,required"`
	// Any of "phone_call", "sms_chat".
	TelnyxConversationChannel ConversationChannelType `json:"telnyx_conversation_channel,omitzero,required"`
	// The phone number, SIP URI, to schedule the call or text to.
	TelnyxEndUserTarget string `json:"telnyx_end_user_target,required"`
	// Required for sms scheduled events. The text to be sent to the end user.
	Text param.Opt[string] `json:"text,omitzero"`
	// Metadata associated with the conversation. Telnyx provides several pieces of
	// metadata, but customers can also add their own.
	ConversationMetadata map[string]AIAssistantScheduledEventNewParamsConversationMetadataUnion `json:"conversation_metadata,omitzero"`
	paramObj
}

func (r AIAssistantScheduledEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantScheduledEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantScheduledEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIAssistantScheduledEventNewParamsConversationMetadataUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u AIAssistantScheduledEventNewParamsConversationMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt, u.OfBool)
}
func (u *AIAssistantScheduledEventNewParamsConversationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIAssistantScheduledEventNewParamsConversationMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type AIAssistantScheduledEventGetParams struct {
	AssistantID string `path:"assistant_id,required" json:"-"`
	paramObj
}

type AIAssistantScheduledEventListParams struct {
	FromDate param.Opt[time.Time] `query:"from_date,omitzero" format:"date-time" json:"-"`
	ToDate   param.Opt[time.Time] `query:"to_date,omitzero" format:"date-time" json:"-"`
	// Any of "phone_call", "sms_chat".
	ConversationChannel ConversationChannelType `query:"conversation_channel,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page AIAssistantScheduledEventListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantScheduledEventListParams]'s query parameters as
// `url.Values`.
func (r AIAssistantScheduledEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type AIAssistantScheduledEventListParamsPage struct {
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	Size   param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantScheduledEventListParamsPage]'s query parameters
// as `url.Values`.
func (r AIAssistantScheduledEventListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAssistantScheduledEventDeleteParams struct {
	AssistantID string `path:"assistant_id,required" json:"-"`
	paramObj
}
