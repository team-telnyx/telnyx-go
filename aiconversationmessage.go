// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

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
func (r *AIConversationMessageService) List(ctx context.Context, conversationID string, opts ...option.RequestOption) (res *AIConversationMessageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/%s/messages", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIConversationMessageListResponse struct {
	Data []AIConversationMessageListResponseData `json:"data,required"`
	Meta Meta                                    `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationMessageListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationMessageListResponseData struct {
	// The role of the message sender.
	//
	// Any of "user", "assistant", "tool".
	Role string `json:"role,required"`
	// The message content. Can be null for tool calls.
	Text string `json:"text,required"`
	// The datetime the message was created on the conversation. This does not
	// necesarily correspond to the time the message was sent. The best field to use to
	// determine the time the end user experienced the message is `sent_at`.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The datetime the message was sent to the end user.
	SentAt time.Time `json:"sent_at" format:"date-time"`
	// Optional tool calls made by the assistant.
	ToolCalls []AIConversationMessageListResponseDataToolCall `json:"tool_calls"`
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
func (r AIConversationMessageListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationMessageListResponseDataToolCall struct {
	// Unique identifier for the tool call.
	ID       string                                                `json:"id,required"`
	Function AIConversationMessageListResponseDataToolCallFunction `json:"function,required"`
	// Type of the tool call.
	//
	// Any of "function".
	Type string `json:"type,required"`
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
func (r AIConversationMessageListResponseDataToolCall) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponseDataToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationMessageListResponseDataToolCallFunction struct {
	// JSON-formatted arguments to pass to the function.
	Arguments string `json:"arguments,required"`
	// Name of the function to call.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationMessageListResponseDataToolCallFunction) RawJSON() string { return r.JSON.raw }
func (r *AIConversationMessageListResponseDataToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
