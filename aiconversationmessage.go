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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage historical AI assistant conversations
//
// AIConversationMessageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationMessageService] method instead.
type AIConversationMessageService struct {
	Options []option.RequestOption
}

// NewAIConversationMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIConversationMessageService(opts ...option.RequestOption) (r AIConversationMessageService) {
	r = AIConversationMessageService{}
	r.Options = opts
	return
}

// Retrieve messages for a specific conversation, including tool calls made by the
// assistant.
func (r *AIConversationMessageService) List(ctx context.Context, conversationID string, query AIConversationMessageListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIConversationMessageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/conversations/%s/messages", conversationID)
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

// Retrieve messages for a specific conversation, including tool calls made by the
// assistant.
func (r *AIConversationMessageService) ListAutoPaging(ctx context.Context, conversationID string, query AIConversationMessageListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIConversationMessageListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, conversationID, query, opts...))
}

type AIConversationMessageListResponse struct {
	// The role of the message sender.
	//
	// Any of "user", "assistant", "tool".
	Role AIConversationMessageListResponseRole `json:"role" api:"required"`
	// The message content. Can be null for tool calls.
	Text string `json:"text" api:"required"`
	// The datetime the message was created on the conversation. This does not
	// necesarily correspond to the time the message was sent. The best field to use to
	// determine the time the end user experienced the message is `sent_at`.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The datetime the message was sent to the end user.
	SentAt time.Time `json:"sent_at" format:"date-time"`
	// Optional tool calls made by the assistant.
	ToolCalls []AIConversationMessageListResponseToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role        respjson.Field
		Text        respjson.Field
		CreatedAt   respjson.Field
		SentAt      respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The role of the message sender.
type AIConversationMessageListResponseRole string

const (
	AIConversationMessageListResponseRoleUser      AIConversationMessageListResponseRole = "user"
	AIConversationMessageListResponseRoleAssistant AIConversationMessageListResponseRole = "assistant"
	AIConversationMessageListResponseRoleTool      AIConversationMessageListResponseRole = "tool"
)

type AIConversationMessageListResponseToolCall struct {
	// Unique identifier for the tool call.
	ID       string                                            `json:"id" api:"required"`
	Function AIConversationMessageListResponseToolCallFunction `json:"function" api:"required"`
	// Type of the tool call.
	//
	// Any of "function".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Function    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationMessageListResponseToolCall) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponseToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationMessageListResponseToolCallFunction struct {
	// JSON-formatted arguments to pass to the function.
	Arguments string `json:"arguments" api:"required"`
	// Name of the function to call.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationMessageListResponseToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponseToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationMessageListParams struct {
	// The page number to retrieve.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The number of messages to return per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIConversationMessageListParams]'s query parameters as
// `url.Values`.
func (r AIConversationMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
