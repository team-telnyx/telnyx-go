// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

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
	Data   []AIOpenAIListModelsResponseData `json:"data" api:"required"`
	Object string                           `json:"object"`
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

// Metadata for a model available on Telnyx Inference. Returned by
// `GET /v2/ai/openai/models` (and the deprecated `GET /v2/ai/models`). Open-source
// models live under their Hugging Face organization (e.g. `moonshotai/Kimi-K2.6`,
// `zai-org/GLM-5.1-FP8`, `MiniMaxAI/MiniMax-M2.7`); fine-tuned models are owned by
// the Telnyx organization that trained them.
type AIOpenAIListModelsResponseData struct {
	// Model identifier. For open-source models, follows the
	// `{organization}/{model_name}` convention from Hugging Face (e.g.
	// `moonshotai/Kimi-K2.6`).
	ID string `json:"id" api:"required"`
	// Maximum total tokens (prompt + completion) supported by the model in a single
	// request.
	ContextLength int64 `json:"context_length" api:"required"`
	// Timestamp at which the model was registered on Telnyx Inference (ISO 8601).
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// ISO language codes the model supports (e.g. `en`, `es`).
	Languages []string `json:"languages" api:"required"`
	// License the model is distributed under, e.g. `Apache 2.0`, `MIT`,
	// `Llama 3 Community License`.
	License string `json:"license" api:"required"`
	// Organization that originally published the model, matching the prefix of `id`
	// for open-source models.
	Organization string `json:"organization" api:"required"`
	// Owner of the model. `Telnyx` for Telnyx-hosted open-source models, the upstream
	// provider name for proxied models, or the Telnyx organization id for fine-tuned
	// models.
	OwnedBy string `json:"owned_by" api:"required"`
	// Total parameter count of the model.
	Parameters int64 `json:"parameters" api:"required"`
	// Billing tier the model belongs to. Used together with `pricing` to determine
	// cost per 1M tokens.
	//
	// Any of "small", "medium", "large", "unlisted".
	Tier string `json:"tier" api:"required"`
	// Base model the fine-tuned model was trained from. Only set for fine-tuned
	// models.
	BaseModel string `json:"base_model" api:"nullable"`
	// Short, human-readable summary of what the model is best suited for.
	Description string `json:"description" api:"nullable"`
	// Whether the model can be used as a base for a fine-tuning job via
	// `POST /v2/ai/fine_tuning/jobs`.
	IsFineTunable bool `json:"is_fine_tunable"`
	// Whether the model accepts image inputs in chat completions (multimodal vision
	// support).
	IsVisionSupported bool `json:"is_vision_supported"`
	// Maximum number of completion (output) tokens the model will generate per
	// request. `null` if unconstrained beyond `context_length`.
	MaxCompletionTokens int64 `json:"max_completion_tokens" api:"nullable"`
	// Object type. Always `model`.
	Object string `json:"object"`
	// Human-readable parameter count, e.g. `1.0T`, `753.9B`, `8B`.
	ParametersStr string `json:"parameters_str" api:"nullable"`
	// Mapping of token kind to price, as strings to preserve precision. Typical keys
	// are `prompt`, `cached_prompt`, and `completion`. When pricing is available the
	// block also includes `currency` (ISO 4217 code matching the account's configured
	// billing currency) and `unit` (the denomination the prices are quoted in,
	// currently always `1M_tokens` for token-priced models).
	Pricing map[string]string `json:"pricing"`
	// Whether Telnyx currently recommends this model as the LLM powering a Telnyx AI
	// Assistant.
	RecommendedForAssistants bool `json:"recommended_for_assistants"`
	// Public region names where the model is currently deployed (e.g. `us-central-1`,
	// `eu-central-1`).
	Regions []string `json:"regions"`
	// Primary task the model is intended for, e.g. `text-generation`,
	// `audio-text-to-text`, `feature-extraction` (embeddings).
	Task string `json:"task"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ContextLength            respjson.Field
		Created                  respjson.Field
		Languages                respjson.Field
		License                  respjson.Field
		Organization             respjson.Field
		OwnedBy                  respjson.Field
		Parameters               respjson.Field
		Tier                     respjson.Field
		BaseModel                respjson.Field
		Description              respjson.Field
		IsFineTunable            respjson.Field
		IsVisionSupported        respjson.Field
		MaxCompletionTokens      respjson.Field
		Object                   respjson.Field
		ParametersStr            respjson.Field
		Pricing                  respjson.Field
		RecommendedForAssistants respjson.Field
		Regions                  respjson.Field
		Task                     respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIOpenAIListModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIOpenAIListModelsResponseData) UnmarshalJSON(data []byte) error {
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
