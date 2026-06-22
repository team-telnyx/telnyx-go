// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

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

// **Deprecated**: Use `POST /v2/ai/openai/responses` instead. This endpoint is
// compatible with the
// [OpenAI Responses API](https://developers.openai.com/api/reference/responses/overview)
// and may be used with the OpenAI JS or Python SDK. Response id parameter is not
// supported at the moment. Use the `conversation` parameter with a Telnyx
// Conversation ID to leverage persistent conversations.
//
// Deprecated: deprecated
func (r *AIService) NewResponseDeprecated(ctx context.Context, body AINewResponseDeprecatedParams, opts ...option.RequestOption) (res *AINewResponseDeprecatedResponse, err error) {
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

// Performs semantic vector search across conversation history records.
//
// **How it works:**
//
//  1. The query text is embedded into a 1024-dimensional vector using the
//     multilingual-e5-large model.
//  2. The vector is sent to regional OpenSearch clusters for kNN search using HNSW
//     cosine similarity.
//  3. When no region is specified, all regions are queried in parallel (fan-out)
//     and results are merged by score.
//  4. Results are ranked by cosine similarity score (descending) and truncated to
//     `top_k`.
//
// **Authentication:** Requires a Telnyx API key via `Authorization: Bearer <key>`.
// Results are automatically scoped to the caller's organization —
// `organization_id` is injected from the auth token and cannot be overridden.
//
// **Chunking:** Records are split into chunks of up to 480 tokens with 64-token
// overlap at ingestion time. Each search result represents a single chunk, with
// `chunk_index` and `chunk_total` indicating its position within the original
// record.
//
// **Filtering:** Use `filter[field][operator]=value` query parameters to narrow
// results before vector search.
//
// Top-level filterable fields: `user_id`, `record_type`, `region`, `document_id`,
// `record_id`, `record_created_at`, `ingested_at`, `retention`
//
// Note: `retention` is filter-only — it can be used to narrow results but is not
// returned in the response body.
//
// Metadata fields: any field not in the list above is resolved to
// `data.metadata.<field>` in OpenSearch (e.g., `filter[language]=en` →
// `data.metadata.language`).
//
// Supported filter operators:
//
// - `eq` — exact match (default when no operator specified)
// - `in` — match any of comma-separated values
// - `gte`, `gt`, `lte`, `lt` — range comparisons (useful for date filtering)
// - `contains` — wildcard substring match
//
// **Examples:**
//
// ```
// GET /v2/ai/conversation_histories?q=billing+issue&record_type=voice&top_k=10
// GET /v2/ai/conversation_histories?q=setup+guide&record_type=knowledge_base&region=USA&min_score=0.5
// GET /v2/ai/conversation_histories?q=refund&record_type=voice&filter[record_created_at][gte]=2026-01-01T00:00:00Z
// GET /v2/ai/conversation_histories?q=outage&record_type=voice&filter[region][in]=USA,DEU
// GET /v2/ai/conversation_histories?q=hold+time&record_type=voice&filter[language]=en
// ```
func (r *AIService) SearchConversationHistories(ctx context.Context, query AISearchConversationHistoriesParams, opts ...option.RequestOption) (res *AISearchConversationHistoriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/conversation_histories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type AINewResponseDeprecatedResponse map[string]any

type AIGetModelsResponse struct {
	Data   []AIGetModelsResponseData `json:"data" api:"required"`
	Object string                    `json:"object"`
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

// Metadata for a model available on Telnyx Inference. Returned by
// `GET /v2/ai/openai/models` (and the deprecated `GET /v2/ai/models`). Open-source
// models live under their Hugging Face organization (e.g. `moonshotai/Kimi-K2.6`,
// `zai-org/GLM-5.1-FP8`, `MiniMaxAI/MiniMax-M2.7`); fine-tuned models are owned by
// the Telnyx organization that trained them.
type AIGetModelsResponseData struct {
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
func (r AIGetModelsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIGetModelsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search response following the standard Telnyx V2 API format.
type AISearchConversationHistoriesResponse struct {
	// Ranked list of matching text chunks, sorted by cosine similarity score
	// descending.
	Data []AISearchConversationHistoriesResponseData `json:"data" api:"required"`
	// Pagination metadata following the standard Telnyx V2 API format.
	Meta AISearchConversationHistoriesResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISearchConversationHistoriesResponse) RawJSON() string { return r.JSON.raw }
func (r *AISearchConversationHistoriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single search result representing one chunk of a conversation history record.
// Records are split into chunks of up to 480 tokens with 64-token overlap at
// ingestion time.
type AISearchConversationHistoriesResponseData struct {
	// Unique chunk identifier.
	ID string `json:"id" api:"required"`
	// Zero-based index of this chunk within the parent record.
	ChunkIndex int64 `json:"chunk_index" api:"required"`
	// Total number of chunks the parent record was split into.
	ChunkTotal int64 `json:"chunk_total" api:"required"`
	// Document identifier. Present only for knowledge_base records; null for all other
	// record types.
	DocumentID string `json:"document_id" api:"required"`
	// When the record was chunked, embedded, and indexed (ISO 8601).
	IngestedAt time.Time `json:"ingested_at" api:"required" format:"date-time"`
	// Identifier of the organization that owns this record.
	OrganizationID string `json:"organization_id" api:"required"`
	// When the original record was created (ISO 8601).
	RecordCreatedAt time.Time `json:"record_created_at" api:"required" format:"date-time"`
	// Identifier of the parent record. Multiple chunks from the same record share this
	// ID.
	RecordID string `json:"record_id" api:"required"`
	// Type of the record.
	//
	// Any of "voice", "message", "ai_pipeline_storage", "knowledge_base".
	RecordType string `json:"record_type" api:"required"`
	// The region where this record is stored.
	//
	// Any of "USA", "DEU", "AUS", "UAE".
	Region string `json:"region" api:"required"`
	// Cosine similarity score between the query vector and this chunk's vector. Higher
	// values indicate greater semantic relevance.
	Score float64 `json:"score" api:"required"`
	// The text content of this chunk (up to 480 tokens).
	Text string `json:"text" api:"required"`
	// Identifier of the user who owns this record.
	UserID string `json:"user_id" api:"required"`
	// Arbitrary metadata attached to the record at ingestion time. Stored as a
	// flat_object in OpenSearch and filterable via filter[field]=value query
	// parameters.
	Metadata map[string]any `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChunkIndex      respjson.Field
		ChunkTotal      respjson.Field
		DocumentID      respjson.Field
		IngestedAt      respjson.Field
		OrganizationID  respjson.Field
		RecordCreatedAt respjson.Field
		RecordID        respjson.Field
		RecordType      respjson.Field
		Region          respjson.Field
		Score           respjson.Field
		Text            respjson.Field
		UserID          respjson.Field
		Metadata        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISearchConversationHistoriesResponseData) RawJSON() string { return r.JSON.raw }
func (r *AISearchConversationHistoriesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata following the standard Telnyx V2 API format.
type AISearchConversationHistoriesResponseMeta struct {
	// Current page number (always 1 — this API does not support pagination, use top_k
	// instead).
	PageNumber int64 `json:"page_number" api:"required"`
	// Number of results per page (equals the effective top_k value).
	PageSize int64 `json:"page_size" api:"required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages" api:"required"`
	// Total number of matching results across all queried regions (before top_k
	// truncation).
	TotalResults int64 `json:"total_results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AISearchConversationHistoriesResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AISearchConversationHistoriesResponseMeta) UnmarshalJSON(data []byte) error {
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

type AINewResponseDeprecatedParams struct {
	Body map[string]any
	paramObj
}

func (r AINewResponseDeprecatedParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AINewResponseDeprecatedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AISearchConversationHistoriesParams struct {
	// Natural language search query. The text is embedded into a 1024-dimensional
	// vector and compared against indexed record chunks using kNN cosine similarity.
	Q string `query:"q" api:"required" json:"-"`
	// The type of records to search. Each record type is stored in a separate vector
	// index.
	//
	// Any of "voice", "message", "ai_pipeline_storage", "knowledge_base".
	RecordType AISearchConversationHistoriesParamsRecordType `query:"record_type,omitzero" api:"required" json:"-"`
	// Filter by document identifier (exact match). Populated for knowledge_base
	// records.
	FilterDocumentID param.Opt[string] `query:"filter[document_id],omitzero" json:"-"`
	// Only include records ingested (chunked, embedded, and indexed) on or after this
	// ISO 8601 timestamp.
	FilterIngestedAtGte param.Opt[time.Time] `query:"filter[ingested_at][gte],omitzero" format:"date-time" json:"-"`
	// Only include records ingested (chunked, embedded, and indexed) on or before this
	// ISO 8601 timestamp.
	FilterIngestedAtLte param.Opt[time.Time] `query:"filter[ingested_at][lte],omitzero" format:"date-time" json:"-"`
	// Only include records whose original creation time is on or after this ISO 8601
	// timestamp.
	FilterRecordCreatedAtGte param.Opt[time.Time] `query:"filter[record_created_at][gte],omitzero" format:"date-time" json:"-"`
	// Only include records whose original creation time is on or before this ISO 8601
	// timestamp.
	FilterRecordCreatedAtLte param.Opt[time.Time] `query:"filter[record_created_at][lte],omitzero" format:"date-time" json:"-"`
	// Filter to chunks belonging to a specific parent record (exact match).
	FilterRecordID param.Opt[string] `query:"filter[record_id],omitzero" json:"-"`
	// Filter by the region stored on the record. Comma-separated to match multiple
	// regions (USA, DEU, AUS, UAE). Distinct from the `region` parameter, which
	// selects which cluster(s) are queried.
	FilterRegionIn param.Opt[string] `query:"filter[region][in],omitzero" json:"-"`
	// Filter by retention policy (exact match). Filter-only: not returned in the
	// response body.
	FilterRetention param.Opt[string] `query:"filter[retention],omitzero" json:"-"`
	// Filter to records owned by a specific user (exact match).
	FilterUserID param.Opt[string] `query:"filter[user_id],omitzero" json:"-"`
	// Minimum cosine similarity score threshold (0.0 to 1.0). Results below this
	// threshold are excluded.
	MinScore param.Opt[float64] `query:"min_score,omitzero" json:"-"`
	// Maximum number of results to return. Defaults to 20, maximum 100.
	TopK param.Opt[int64] `query:"top_k,omitzero" json:"-"`
	// Restrict search to a specific region's OpenSearch cluster. When omitted, all
	// regions are queried in parallel (fan-out) and results are merged by cosine
	// similarity score.
	//
	// Any of "USA", "DEU", "AUS", "UAE".
	Region AISearchConversationHistoriesParamsRegion `query:"region,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AISearchConversationHistoriesParams]'s query parameters as
// `url.Values`.
func (r AISearchConversationHistoriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of records to search. Each record type is stored in a separate vector
// index.
type AISearchConversationHistoriesParamsRecordType string

const (
	AISearchConversationHistoriesParamsRecordTypeVoice             AISearchConversationHistoriesParamsRecordType = "voice"
	AISearchConversationHistoriesParamsRecordTypeMessage           AISearchConversationHistoriesParamsRecordType = "message"
	AISearchConversationHistoriesParamsRecordTypeAIPipelineStorage AISearchConversationHistoriesParamsRecordType = "ai_pipeline_storage"
	AISearchConversationHistoriesParamsRecordTypeKnowledgeBase     AISearchConversationHistoriesParamsRecordType = "knowledge_base"
)

// Restrict search to a specific region's OpenSearch cluster. When omitted, all
// regions are queried in parallel (fan-out) and results are merged by cosine
// similarity score.
type AISearchConversationHistoriesParamsRegion string

const (
	AISearchConversationHistoriesParamsRegionUsa AISearchConversationHistoriesParamsRegion = "USA"
	AISearchConversationHistoriesParamsRegionDeu AISearchConversationHistoriesParamsRegion = "DEU"
	AISearchConversationHistoriesParamsRegionAus AISearchConversationHistoriesParamsRegion = "AUS"
	AISearchConversationHistoriesParamsRegionUae AISearchConversationHistoriesParamsRegion = "UAE"
)

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
