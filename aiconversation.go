// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage historical AI assistant conversations
//
// AIConversationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationService] method instead.
type AIConversationService struct {
	Options []option.RequestOption
	// Manage historical AI assistant conversations
	InsightGroups AIConversationInsightGroupService
	// Manage historical AI assistant conversations
	Insights AIConversationInsightService
	// Manage historical AI assistant conversations
	Messages AIConversationMessageService
}

// NewAIConversationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIConversationService(opts ...option.RequestOption) (r AIConversationService) {
	r = AIConversationService{}
	r.Options = opts
	r.InsightGroups = NewAIConversationInsightGroupService(opts...)
	r.Insights = NewAIConversationInsightService(opts...)
	r.Messages = NewAIConversationMessageService(opts...)
	return
}

// Create a new AI Conversation.
func (r *AIConversationService) New(ctx context.Context, body AIConversationNewParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific AI conversation by its ID.
func (r *AIConversationService) Get(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *AIConversationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update metadata for a specific conversation.
func (r *AIConversationService) Update(ctx context.Context, conversationID string, body AIConversationUpdateParams, opts ...option.RequestOption) (res *AIConversationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of all AI conversations configured by the user. Supports
// [PostgREST-style query parameters](https://postgrest.org/en/stable/api.html#horizontal-filtering-rows)
// for filtering. Examples are included for the standard metadata fields, but you
// can filter on any field in the metadata JSON object. For example, to filter by a
// custom field `metadata->custom_field`, use `metadata->custom_field=eq.value`.
func (r *AIConversationService) List(ctx context.Context, query AIConversationListParams, opts ...option.RequestOption) (res *AIConversationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific conversation by its ID.
func (r *AIConversationService) Delete(ctx context.Context, conversationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Add a new message to the conversation. Used to insert a new messages to a
// conversation manually ( without using chat endpoint )
func (r *AIConversationService) AddMessage(ctx context.Context, conversationID string, body AIConversationAddMessageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s/message", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Retrieve insights for a specific conversation
func (r *AIConversationService) GetConversationsInsights(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *AIConversationGetConversationsInsightsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s/conversations-insights", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Conversation struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// The datetime the conversation was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The datetime of the latest message in the conversation.
	LastMessageAt time.Time `json:"last_message_at" api:"required" format:"date-time"`
	// Metadata associated with the conversation. Telnyx provides several pieces of
	// metadata, but customers can also add their own. The reserved field `ai_disabled`
	// (boolean) can be set to `true` to prevent AI-generated responses on this
	// conversation. When `ai_disabled` is `true`, calls to the chat endpoint will
	// return a 400 error. Set to `false` or remove the field to re-enable AI
	// responses. This is useful when a human agent needs to take over the conversation
	// mid-stream (e.g., a technician stepping in while AI was messaging a resident).
	Metadata map[string]string `json:"metadata" api:"required"`
	Name     string            `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		LastMessageAt respjson.Field
		Metadata      respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Conversation) RawJSON() string { return r.JSON.raw }
func (r *Conversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationGetResponse struct {
	Data Conversation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationUpdateResponse struct {
	Data Conversation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationListResponse struct {
	Data []Conversation `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationGetConversationsInsightsResponse struct {
	Data []AIConversationGetConversationsInsightsResponseData `json:"data" api:"required"`
	Meta Meta                                                 `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationGetConversationsInsightsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationGetConversationsInsightsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationGetConversationsInsightsResponseData struct {
	// Unique identifier for the conversation insight.
	ID string `json:"id" api:"required"`
	// List of insights extracted from the conversation.
	ConversationInsights []AIConversationGetConversationsInsightsResponseDataConversationInsight `json:"conversation_insights" api:"required"`
	// Timestamp of when the object was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Current status of the insight generation for the conversation.
	//
	// Any of "pending", "in_progress", "completed", "failed".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ConversationInsights respjson.Field
		CreatedAt            respjson.Field
		Status               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationGetConversationsInsightsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIConversationGetConversationsInsightsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationGetConversationsInsightsResponseDataConversationInsight struct {
	// Unique identifier for the insight configuration.
	InsightID string `json:"insight_id" api:"required"`
	// Insight result from the conversation. If the insight has a JSON schema, this
	// will be stringified JSON object.
	Result string `json:"result" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsightID   respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationGetConversationsInsightsResponseDataConversationInsight) RawJSON() string {
	return r.JSON.raw
}
func (r *AIConversationGetConversationsInsightsResponseDataConversationInsight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationNewParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Metadata associated with the conversation. Set `ai_disabled` to `true` to create
	// the conversation with AI message responses disabled.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r AIConversationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationUpdateParams struct {
	// Metadata associated with the conversation. Set `ai_disabled` to `true` to stop
	// AI from responding to messages (e.g., when a human agent takes over). Set to
	// `false` to re-enable AI responses.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r AIConversationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationListParams struct {
	// Filter by conversation ID (e.g. id=eq.123)
	ID param.Opt[string] `query:"id,omitzero" json:"-"`
	// Filter by creation datetime (e.g., `created_at=gte.2025-01-01`)
	CreatedAt param.Opt[string] `query:"created_at,omitzero" json:"-"`
	// Filter by last message datetime (e.g., `last_message_at=lte.2025-06-01`)
	LastMessageAt param.Opt[string] `query:"last_message_at,omitzero" json:"-"`
	// Limit the number of returned conversations (e.g., `limit=10`)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by assistant ID (e.g., `metadata->assistant_id=eq.assistant-123`)
	MetadataAssistantID param.Opt[string] `query:"metadata->assistant_id,omitzero" json:"-"`
	// Filter by call control ID (e.g., `metadata->call_control_id=eq.v3:123`)
	MetadataCallControlID param.Opt[string] `query:"metadata->call_control_id,omitzero" json:"-"`
	// Filter by the phone number, SIP URI, or other identifier for the agent (e.g.,
	// `metadata->telnyx_agent_target=eq.+13128675309`)
	MetadataTelnyxAgentTarget param.Opt[string] `query:"metadata->telnyx_agent_target,omitzero" json:"-"`
	// Filter by conversation channel (e.g.,
	// `metadata->telnyx_conversation_channel=eq.phone_call`)
	MetadataTelnyxConversationChannel param.Opt[string] `query:"metadata->telnyx_conversation_channel,omitzero" json:"-"`
	// Filter by the phone number, SIP URI, or other identifier for the end user (e.g.,
	// `metadata->telnyx_end_user_target=eq.+13128675309`)
	MetadataTelnyxEndUserTarget param.Opt[string] `query:"metadata->telnyx_end_user_target,omitzero" json:"-"`
	// Filter by conversation Name (e.g. `name=like.Voice%`)
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Apply OR conditions using PostgREST syntax (e.g.,
	// `or=(created_at.gte.2025-04-01,last_message_at.gte.2025-04-01)`)
	Or param.Opt[string] `query:"or,omitzero" json:"-"`
	// Order the results by specific fields (e.g., `order=created_at.desc` or
	// `order=last_message_at.asc`)
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIConversationListParams]'s query parameters as
// `url.Values`.
func (r AIConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIConversationAddMessageParams struct {
	Role       string                                                 `json:"role" api:"required"`
	Content    param.Opt[string]                                      `json:"content,omitzero"`
	Name       param.Opt[string]                                      `json:"name,omitzero"`
	SentAt     param.Opt[time.Time]                                   `json:"sent_at,omitzero" format:"date-time"`
	ToolCallID param.Opt[string]                                      `json:"tool_call_id,omitzero"`
	Metadata   map[string]AIConversationAddMessageParamsMetadataUnion `json:"metadata,omitzero"`
	ToolCalls  []map[string]any                                       `json:"tool_calls,omitzero"`
	ToolChoice AIConversationAddMessageParamsToolChoiceUnion          `json:"tool_choice,omitzero"`
	paramObj
}

func (r AIConversationAddMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow AIConversationAddMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIConversationAddMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIConversationAddMessageParamsMetadataUnion struct {
	OfString                                                   param.Opt[string]                                                   `json:",omitzero,inline"`
	OfInt                                                      param.Opt[int64]                                                    `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                     `json:",omitzero,inline"`
	OfAIConversationAddMessagesMetadataUnionArrayVariant3Array []AIConversationAddMessageParamsMetadataUnionArrayVariant3ItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AIConversationAddMessageParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt, u.OfBool, u.OfAIConversationAddMessagesMetadataUnionArrayVariant3Array)
}
func (u *AIConversationAddMessageParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIConversationAddMessageParamsMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAIConversationAddMessagesMetadataUnionArrayVariant3Array) {
		return &u.OfAIConversationAddMessagesMetadataUnionArrayVariant3Array
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIConversationAddMessageParamsMetadataUnionArrayVariant3ItemUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u AIConversationAddMessageParamsMetadataUnionArrayVariant3ItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt, u.OfBool)
}
func (u *AIConversationAddMessageParamsMetadataUnionArrayVariant3ItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIConversationAddMessageParamsMetadataUnionArrayVariant3ItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIConversationAddMessageParamsToolChoiceUnion struct {
	OfString           param.Opt[string] `json:",omitzero,inline"`
	OfToolChoiceObject map[string]any    `json:",omitzero,inline"`
	paramUnion
}

func (u AIConversationAddMessageParamsToolChoiceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfToolChoiceObject)
}
func (u *AIConversationAddMessageParamsToolChoiceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIConversationAddMessageParamsToolChoiceUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfToolChoiceObject) {
		return &u.OfToolChoiceObject
	}
	return nil
}
