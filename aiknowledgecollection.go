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

// Create and manage logical collections of your Telnyx data, tune retrieval
// settings, manage sources, and run collection-scoped semantic search.
//
// AIKnowledgeCollectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIKnowledgeCollectionService] method instead.
type AIKnowledgeCollectionService struct {
	Options []option.RequestOption
}

// NewAIKnowledgeCollectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIKnowledgeCollectionService(opts ...option.RequestOption) (r AIKnowledgeCollectionService) {
	r = AIKnowledgeCollectionService{}
	r.Options = opts
	return
}

// Runs search over the documents in a collection, ranked by relevance to `query`.
// Searches currently run `vector` retrieval (semantic similarity). The
// collection's `retrieval_type` setting is the forward-compatible selector:
// `hybrid` (vector similarity fused with keyword matching) can be set but cannot
// be searched yet, and `keyword` (lexical BM25 matching) is not accepted yet --
// setting it returns 422 `unsupported_retrieval_type`. A per-request
// `retrieval_type` is accepted but ignored; `meta.retrieval_type` echoes the mode
// that actually ran. When `query` is omitted, returns a plain catalog listing of
// the collection's documents.
//
// **How it works:**
//
//  1. The `query` text is embedded into a 1024-dimensional vector using the
//     multilingual-e5-large model.
//  2. The embedding is compared against the collection's indexed document chunks
//     using semantic similarity. When `hybrid` and `keyword` execution ship, those
//     scores will be fused with, or replaced by, lexical BM25 matching.
//  3. Results are ranked by `score` (descending) and paginated via `page[number]` /
//     `page[size]`.
//
// **Authentication:** Requires a Telnyx API key via `Authorization: Bearer <key>`.
// Results are automatically scoped to your organization and cannot be overridden.
//
// **Filtering:** Use `filter[field][operator]=value` query parameters to narrow
// results before search. Supported operators: `eq` (default), `in`, `gte`, `gt`,
// `lte`, `lt`, `contains`. Metadata fields resolve to `metadata.<field>`.
//
// **Examples:**
//
// - `GET /v2/ai/knowledge/collections/my-collection/documents?query=billing+issue&top_k=10`
// - `GET /v2/ai/knowledge/collections/my-collection/documents?query=refund&sources=voice,message`
// - `GET /v2/ai/knowledge/collections/my-collection/documents?query=outage&filter[record_created_at][gte]=2026-01-01T00:00:00Z`
func (r *AIKnowledgeCollectionService) GetDocuments(ctx context.Context, slug string, query AIKnowledgeCollectionGetDocumentsParams, opts ...option.RequestOption) (res *AIKnowledgeCollectionGetDocumentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/knowledge/collections/%s/documents", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AIKnowledgeCollectionGetDocumentsResponse struct {
	Data []AIKnowledgeCollectionGetDocumentsResponseData `json:"data"`
	Meta AIKnowledgeCollectionGetDocumentsResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIKnowledgeCollectionGetDocumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIKnowledgeCollectionGetDocumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIKnowledgeCollectionGetDocumentsResponseData struct {
	ID              string         `json:"id"`
	ChunkIndex      int64          `json:"chunk_index"`
	ChunkTotal      int64          `json:"chunk_total"`
	IngestedAt      time.Time      `json:"ingested_at" format:"date-time"`
	Metadata        map[string]any `json:"metadata"`
	OrganizationID  string         `json:"organization_id"`
	RecordCreatedAt time.Time      `json:"record_created_at" format:"date-time"`
	RecordID        string         `json:"record_id"`
	// The source record kind this chunk came from (e.g. `voice`, `meeting_bot`,
	// `message`).
	RecordType string `json:"record_type"`
	Region     string `json:"region"`
	// Relevance score (higher = more relevant) for ranked search. `0.0` for plain
	// catalog listings (when `query` is omitted).
	Score  float64 `json:"score"`
	Text   string  `json:"text"`
	UserID string  `json:"user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChunkIndex      respjson.Field
		ChunkTotal      respjson.Field
		IngestedAt      respjson.Field
		Metadata        respjson.Field
		OrganizationID  respjson.Field
		RecordCreatedAt respjson.Field
		RecordID        respjson.Field
		RecordType      respjson.Field
		Region          respjson.Field
		Score           respjson.Field
		Text            respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIKnowledgeCollectionGetDocumentsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIKnowledgeCollectionGetDocumentsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIKnowledgeCollectionGetDocumentsResponseMeta struct {
	CollectionSlug  string   `json:"collection_slug"`
	PageNumber      int64    `json:"page_number"`
	PageSize        int64    `json:"page_size"`
	RetrievalType   string   `json:"retrieval_type"`
	SearchedSources []string `json:"searched_sources"`
	TopK            int64    `json:"top_k"`
	TotalPages      int64    `json:"total_pages"`
	TotalResults    int64    `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionSlug  respjson.Field
		PageNumber      respjson.Field
		PageSize        respjson.Field
		RetrievalType   respjson.Field
		SearchedSources respjson.Field
		TopK            respjson.Field
		TotalPages      respjson.Field
		TotalResults    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIKnowledgeCollectionGetDocumentsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AIKnowledgeCollectionGetDocumentsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIKnowledgeCollectionGetDocumentsParams struct {
	// Page number to return (1-based). Defaults to 1.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page. Defaults to 20.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Natural-language search query. When provided, the text is matched against the
	// collection's document chunks using the collection's `retrieval_type` (vector or
	// hybrid). When omitted, documents are returned as a plain catalog listing.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Comma-separated list of source types to restrict the search to. When omitted,
	// all of the collection's sources are searched.
	Sources param.Opt[string] `query:"sources,omitzero" json:"-"`
	// Maximum number of ranked results to consider. When omitted, the collection's
	// configured `top_k` setting is used.
	TopK param.Opt[int64] `query:"top_k,omitzero" json:"-"`
	// Field filters applied before ranking, using `filter[field][operator]=value`.
	// Supported operators: `eq` (default), `in`, `gte`, `gt`, `lte`, `lt`, `contains`.
	// Known fields: `record_type`, `record_id`, `user_id`, `record_created_at`,
	// `ingested_at`; any other name resolves to a `metadata.<field>` filter. Example:
	// `filter[record_id][eq]=rec_123`.
	Filter map[string]any `query:"filter,omitzero" json:"-"`
	// Reserved; not yet functional. A value supplied here is accepted but ignored — it
	// does not override the collection's configured strategy, and it is not echoed
	// back. Searches run `vector` retrieval, and `meta.retrieval_type` reports the
	// mode that actually ran. To change retrieval strategy, set it on the collection's
	// settings subresource.
	//
	// Any of "vector", "hybrid", "keyword".
	RetrievalType AIKnowledgeCollectionGetDocumentsParamsRetrievalType `query:"retrieval_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIKnowledgeCollectionGetDocumentsParams]'s query parameters
// as `url.Values`.
func (r AIKnowledgeCollectionGetDocumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Reserved; not yet functional. A value supplied here is accepted but ignored — it
// does not override the collection's configured strategy, and it is not echoed
// back. Searches run `vector` retrieval, and `meta.retrieval_type` reports the
// mode that actually ran. To change retrieval strategy, set it on the collection's
// settings subresource.
type AIKnowledgeCollectionGetDocumentsParamsRetrievalType string

const (
	AIKnowledgeCollectionGetDocumentsParamsRetrievalTypeVector  AIKnowledgeCollectionGetDocumentsParamsRetrievalType = "vector"
	AIKnowledgeCollectionGetDocumentsParamsRetrievalTypeHybrid  AIKnowledgeCollectionGetDocumentsParamsRetrievalType = "hybrid"
	AIKnowledgeCollectionGetDocumentsParamsRetrievalTypeKeyword AIKnowledgeCollectionGetDocumentsParamsRetrievalType = "keyword"
)
