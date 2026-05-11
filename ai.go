// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Generate text with LLMs
//
// AIService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	Options []option.RequestOption
	// Configure AI assistant specifications
	Assistants AIAssistantService
	Audio      AIAudioService
	// Generate text with LLMs
	Chat AIChatService
	// Identify common themes and patterns in your embedded documents
	Clusters AIClusterService
	// Manage historical AI assistant conversations
	Conversations AIConversationService
	// Embed documents and perform text searches
	Embeddings   AIEmbeddingService
	FineTuning   AIFineTuningService
	Integrations AIIntegrationService
	McpServers   AIMcpServerService
	Missions     AIMissionService
	OpenAI       AIOpenAIService
	// Configure AI assistant specifications
	Tools AIToolService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r AIService) {
	r = AIService{}
	r.Options = opts
	r.Assistants = NewAIAssistantService(opts...)
	r.Audio = NewAIAudioService(opts...)
	r.Chat = NewAIChatService(opts...)
	r.Clusters = NewAIClusterService(opts...)
	r.Conversations = NewAIConversationService(opts...)
	r.Embeddings = NewAIEmbeddingService(opts...)
	r.FineTuning = NewAIFineTuningService(opts...)
	r.Integrations = NewAIIntegrationService(opts...)
	r.McpServers = NewAIMcpServerService(opts...)
	r.Missions = NewAIMissionService(opts...)
	r.OpenAI = NewAIOpenAIService(opts...)
	r.Tools = NewAIToolService(opts...)
	return
}

// **Deprecated**: Use `POST /v2/ai/openai/responses` instead. Chat with a language
// model. This endpoint is consistent with the
// [OpenAI Responses API](https://platform.openai.com/docs/api-reference/responses)
// and may be used with the OpenAI JS or Python SDK. Response id parameter is not
// supported at the moment. Use 'conversation' parameter to leverage persistent
// conversations feature.
//
// Deprecated: deprecated
func (r *AIService) NewResponse(ctx context.Context, body AINewResponseParams, opts ...option.RequestOption) (res *AINewResponseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Deprecated**: Use `GET /v2/ai/openai/models` instead.
//
// Returns the same `ModelsResponse` payload as the OpenAI-compatible endpoint —
// open-source LLMs hosted on Telnyx (e.g. `moonshotai/Kimi-K2.6`,
// `zai-org/GLM-5.1-FP8`, `MiniMaxAI/MiniMax-M2.7`), embedding models, and
// fine-tuned models — kept around for backwards compatibility. New integrations
// should use `/v2/ai/openai/models`.
//
// Model ids follow the `{organization}/{model_name}` convention from Hugging Face.
//
// Deprecated: deprecated
func (r *AIService) GetModels(ctx context.Context, opts ...option.RequestOption) (res *AIGetModelsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Generate a summary of a file's contents.
//
// Supports the following text formats:
//
// - PDF, HTML, txt, json, csv
//
// Supports the following media formats (billed for both the transcription and
// summary):
//
// - flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm
// - Up to 100 MB
func (r *AIService) Summarize(ctx context.Context, body AISummarizeParams, opts ...option.RequestOption) (res *AISummarizeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/summarize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Metadata for a model available on Telnyx Inference. Returned by
// `GET /v2/ai/openai/models` (and the deprecated `GET /v2/ai/models`). Open-source
// models live under their Hugging Face organization (e.g. `moonshotai/Kimi-K2.6`,
// `zai-org/GLM-5.1-FP8`, `MiniMaxAI/MiniMax-M2.7`); fine-tuned models are owned by
// the Telnyx organization that trained them.
type ModelMetadata struct {
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
	Tier ModelMetadataTier `json:"tier" api:"required"`
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
	// Mapping of token kind to price in USD per 1M tokens, as a string. Typical keys
	// are `input` and `output`; embedding models expose `embedding`. Empty object when
	// pricing is not yet published for the model.
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
func (r ModelMetadata) RawJSON() string { return r.JSON.raw }
func (r *ModelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Billing tier the model belongs to. Used together with `pricing` to determine
// cost per 1M tokens.
type ModelMetadataTier string

const (
	ModelMetadataTierSmall    ModelMetadataTier = "small"
	ModelMetadataTierMedium   ModelMetadataTier = "medium"
	ModelMetadataTierLarge    ModelMetadataTier = "large"
	ModelMetadataTierUnlisted ModelMetadataTier = "unlisted"
)

type AINewResponseResponse map[string]any

type AIGetModelsResponse struct {
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
func (r AIGetModelsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIGetModelsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISummarizeResponse struct {
	Data AISummarizeResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISummarizeResponse) RawJSON() string { return r.JSON.raw }
func (r *AISummarizeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISummarizeResponseData struct {
	Summary string `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISummarizeResponseData) RawJSON() string { return r.JSON.raw }
func (r *AISummarizeResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AINewResponseParams struct {
	Body map[string]any
	paramObj
}

func (r AINewResponseParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AINewResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISummarizeParams struct {
	// The name of the bucket that contains the file to be summarized.
	Bucket string `json:"bucket" api:"required"`
	// The name of the file to be summarized.
	Filename string `json:"filename" api:"required"`
	// A system prompt to guide the summary generation.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	paramObj
}

func (r AISummarizeParams) MarshalJSON() (data []byte, err error) {
	type shadow AISummarizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AISummarizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
