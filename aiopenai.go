// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIOpenAIService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOpenAIService] method instead.
type AIOpenAIService struct {
	Options []option.RequestOption
	// OpenAI-compatible embeddings endpoints for generating vector representations of
	// text
	Embeddings AIOpenAIEmbeddingService
	Chat       AIOpenAIChatService
}

// NewAIOpenAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOpenAIService(opts ...option.RequestOption) (r AIOpenAIService) {
	r = AIOpenAIService{}
	r.Options = opts
	r.Embeddings = NewAIOpenAIEmbeddingService(opts...)
	r.Chat = NewAIOpenAIChatService(opts...)
	return
}

// Create a response using Telnyx's OpenAI-compatible Responses API. This endpoint
// is compatible with the
// [OpenAI Responses API](https://developers.openai.com/api/reference/responses/overview)
// and may be used with the OpenAI JS or Python SDK by setting the base URL to
// `https://api.telnyx.com/v2/ai/openai`.
//
// The `conversation` parameter refers to a Telnyx Conversation rather than an
// OpenAI-hosted conversation object. To persist a thread across turns, first
// [create a conversation](https://developers.telnyx.com/api-reference/conversations/create-a-conversation)
// with `POST /ai/conversations`, then pass that conversation's `id` in the
// Responses request as `conversation`. The endpoint appends the new input,
// assistant output, reasoning, and tool-call messages to that conversation. Reuse
// the same `conversation` id on subsequent Responses requests, including
// tool-result followups, so the model receives the prior context.
//
// If `conversation` is omitted, the request is processed without persisting
// messages to a Telnyx conversation. Use the Conversations API to manage history:
// [list conversations](https://developers.telnyx.com/api-reference/conversations/list-conversations)
// (optionally filtered by metadata),
// [fetch messages](https://developers.telnyx.com/api-reference/conversations/get-conversation-messages)
// for a conversation, and optionally
// [add messages](https://developers.telnyx.com/api-reference/conversations/create-message)
// outside the Responses flow.
//
// You can attach arbitrary metadata when creating a conversation (for example to
// tag the conversation's source, channel, or user) and later filter by it when
// listing conversations.
func (r *AIOpenAIService) NewResponse(ctx context.Context, body AIOpenAINewResponseParams, opts ...option.RequestOption) (res *AIOpenAINewResponseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists every model currently available to your account on Telnyx Inference,
// including SOTA open-source LLMs hosted on Telnyx GPUs (for example
// `moonshotai/Kimi-K2.6`, `zai-org/GLM-5.1-FP8`, and `MiniMaxAI/MiniMax-M2.7`),
// embedding models, and any fine-tuned models you have created.
//
// Each entry is a `ModelMetadata` object describing the model id, owner, task,
// context length, supported languages, billing tier, pricing per 1M tokens,
// deployment regions, and whether the model supports vision or fine-tuning. Use
// this endpoint to discover model ids you can pass to
// `POST /v2/ai/openai/chat/completions`.
//
// Model ids follow the `{organization}/{model_name}` convention from Hugging Face
// (for example `moonshotai/Kimi-K2.6`). This endpoint is OpenAI-compatible:
// clients pointed at `https://api.telnyx.com/v2/ai/openai` can call
// `client.models.list()` to retrieve the same payload.
func (r *AIOpenAIService) ListModels(ctx context.Context, opts ...option.RequestOption) (res *AIOpenAIListModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/openai/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AIOpenAINewResponseResponse map[string]any

type AIOpenAIListModelsResponse struct {
	Data   []ModelMetadata `json:"data" api:"required"`
	Object string          `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIListModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIListModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIOpenAINewResponseParams struct {
	// Optional Telnyx Conversation ID from `POST /ai/conversations`. When provided,
	// Telnyx stores this turn on that conversation and uses the conversation's prior
	// messages as context. Reuse the same ID for subsequent turns and tool-result
	// followups. Omit it for a non-persisted, stateless response.
	Conversation param.Opt[string] `json:"conversation,omitzero" format:"uuid"`
	// Optional system/developer instructions for the model. When used with a persisted
	// `conversation`, send these on the first request that creates the thread;
	// subsequent turns can rely on the stored history.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Model identifier to use for the response, for example `zai-org/GLM-5.1-FP8` or
	// another model available from the Telnyx OpenAI-compatible models endpoint.
	Model param.Opt[string] `json:"model,omitzero"`
	// Set to `true` to stream Server-Sent Events, matching OpenAI's Responses
	// streaming format.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// The input items for this turn, using the OpenAI Responses API input format.
	Input any `json:"input,omitzero"`
	paramObj
}

func (r AIOpenAINewResponseParams) MarshalJSON() (data []byte, err error) {
	type shadow AIOpenAINewResponseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIOpenAINewResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
