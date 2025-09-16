// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
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

// AIConversationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationService] method instead.
type AIConversationService struct {
	Options       []option.RequestOption
	InsightGroups AIConversationInsightGroupService
	Insights      AIConversationInsightService
	Messages      AIConversationMessageService
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
	opts = append(r.Options[:], opts...)
	path := "ai/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific AI conversation by its ID.
func (r *AIConversationService) Get(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *AIConversationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	path := "ai/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific conversation by its ID.
func (r *AIConversationService) Delete(ctx context.Context, conversationID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve insights for a specific conversation
func (r *AIConversationService) GetConversationsInsights(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *AIConversationGetConversationsInsightsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s/conversations-insights", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Conversation struct {
	ID string `json:"id,required" format:"uuid"`
	// The datetime the conversation was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The datetime of the latest message in the conversation.
	LastMessageAt time.Time `json:"last_message_at,required" format:"date-time"`
	// Metadata associated with the conversation. Telnyx provides several pieces of
	// metadata, but customers can also add their own.
	Metadata map[string]string `json:"metadata,required"`
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
	Data []Conversation `json:"data,required"`
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
	Data []AIConversationGetConversationsInsightsResponseData `json:"data,required"`
	Meta Meta                                                 `json:"meta,required"`
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
	ID string `json:"id,required"`
	// List of insights extracted from the conversation.
	ConversationInsights []AIConversationGetConversationsInsightsResponseDataConversationInsight `json:"conversation_insights,required"`
	// Timestamp of when the object was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Current status of the insight generation for the conversation.
	//
	// Any of "pending", "in_progress", "completed", "failed".
	Status string `json:"status,required"`
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
	InsightID string `json:"insight_id,required"`
	// Insight result from the conversation. If the insight has a JSON schema, this
	// will be stringified JSON object.
	Result string `json:"result,required"`
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
	// Metadata associated with the conversation.
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
	// Metadata associated with the conversation.
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
