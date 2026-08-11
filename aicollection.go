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

// Create and manage logical collections of your Telnyx data, tune retrieval
// settings, manage sources, and run collection-scoped semantic search.
//
// AICollectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAICollectionService] method instead.
type AICollectionService struct {
	Options []option.RequestOption
	// Create and manage logical collections of your Telnyx data, tune retrieval
	// settings, manage sources, and run collection-scoped semantic search.
	Settings AICollectionSettingService
	// Create and manage logical collections of your Telnyx data, tune retrieval
	// settings, manage sources, and run collection-scoped semantic search.
	Sources AICollectionSourceService
}

// NewAICollectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAICollectionService(opts ...option.RequestOption) (r AICollectionService) {
	r = AICollectionService{}
	r.Options = opts
	r.Settings = NewAICollectionSettingService(opts...)
	r.Sources = NewAICollectionSourceService(opts...)
	return
}

// Creates a new collection scoped to your organization. Optionally attach sources
// and retrieval settings at creation time. If `slug` is omitted, one is derived
// from `name` and must be unique within your organization.
func (r *AICollectionService) New(ctx context.Context, body AICollectionNewParams, opts ...option.RequestOption) (res *CollectionEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetches a single collection by its `slug`.
func (r *AICollectionService) Get(ctx context.Context, slug string, opts ...option.RequestOption) (res *CollectionEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/slug/%s", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a collection's metadata (`name` and/or `description`). Sources and
// settings are managed through their own sub-resources.
func (r *AICollectionService) Update(ctx context.Context, uuid string, body AICollectionUpdateParams, opts ...option.RequestOption) (res *CollectionEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of collections in your organization.
func (r *AICollectionService) List(ctx context.Context, query AICollectionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Collection], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/collections"
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

// Returns a paginated list of collections in your organization.
func (r *AICollectionService) ListAutoPaging(ctx context.Context, query AICollectionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Collection] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Soft-deletes a collection. Its `slug` is freed and may be reused by a new
// collection.
func (r *AICollectionService) Delete(ctx context.Context, uuid string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return err
	}
	path := fmt.Sprintf("ai/collections/%s", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetches a single collection by its `uuid`.
func (r *AICollectionService) GetByID(ctx context.Context, uuid string, opts ...option.RequestOption) (res *CollectionEnvelope, err error) {
	opts = slices.Concat(r.Options, opts)
	if uuid == "" {
		err = errors.New("missing required uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s", uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Runs search over the documents in a collection, ranked by relevance to `query`.
// The collection's `retrieval_type` setting selects the strategy: `vector`
// (semantic similarity), `hybrid` (vector similarity fused with keyword matching),
// or `keyword` (lexical BM25 matching). When `query` is omitted, returns a plain
// catalog listing of the collection's documents.
//
// **How it works:**
//
//  1. For `vector` and `hybrid`, the `query` text is embedded into a
//     1024-dimensional vector using the multilingual-e5-large model.
//  2. For `vector`, the embedding is compared against the collection's indexed
//     document chunks using semantic similarity; for `hybrid`, those similarity
//     scores are fused with keyword-match scores; for `keyword`, only lexical BM25
//     matching is applied.
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
// - `GET /v2/ai/collections/my-collection/documents?query=billing+issue&top_k=10`
// - `GET /v2/ai/collections/my-collection/documents?query=refund&sources=voice,message`
// - `GET /v2/ai/collections/my-collection/documents?query=outage&filter[record_created_at][gte]=2026-01-01T00:00:00Z`
func (r *AICollectionService) GetDocuments(ctx context.Context, slug string, query AICollectionGetDocumentsParams, opts ...option.RequestOption) (res *AICollectionGetDocumentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/collections/%s/documents", slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Collection struct {
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	// Identifies the record type. Always `ai_collection`.
	RecordType string                   `json:"record_type"`
	Settings   RetrievalSettingsWrapper `json:"settings"`
	Slug       string                   `json:"slug"`
	Sources    []Source                 `json:"sources"`
	Status     string                   `json:"status"`
	UpdatedAt  time.Time                `json:"updated_at" format:"date-time"`
	Uuid       string                   `json:"uuid" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		Settings    respjson.Field
		Slug        respjson.Field
		Sources     respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		Uuid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Collection) RawJSON() string { return r.JSON.raw }
func (r *Collection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionEnvelope struct {
	Data Collection `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionEnvelope) RawJSON() string { return r.JSON.raw }
func (r *CollectionEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionGetDocumentsResponse struct {
	Data []AICollectionGetDocumentsResponseData `json:"data"`
	Meta AICollectionGetDocumentsResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AICollectionGetDocumentsResponse) RawJSON() string { return r.JSON.raw }
func (r *AICollectionGetDocumentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionGetDocumentsResponseData struct {
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
func (r AICollectionGetDocumentsResponseData) RawJSON() string { return r.JSON.raw }
func (r *AICollectionGetDocumentsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionGetDocumentsResponseMeta struct {
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
func (r AICollectionGetDocumentsResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AICollectionGetDocumentsResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionNewParams struct {
	// Human-readable collection name.
	Name string `json:"name" api:"required"`
	// Optional description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional slug (unique per organization). Derived from `name` when omitted.
	Slug param.Opt[string] `json:"slug,omitzero"`
	// Optional retrieval settings.
	Settings RetrievalSettingsWrapperParam `json:"settings,omitzero"`
	// Optional sources to attach at creation time.
	Sources []SourceRequestParam `json:"sources,omitzero"`
	paramObj
}

func (r AICollectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AICollectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AICollectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AICollectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AICollectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AICollectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AICollectionListParams struct {
	// Page number to return (1-based). Defaults to 1.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page. Defaults to 20.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AICollectionListParams]'s query parameters as `url.Values`.
func (r AICollectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AICollectionGetDocumentsParams struct {
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
	// Override the collection's configured retrieval strategy for this request. Echoed
	// back in `meta.retrieval_type`.
	//
	// Any of "vector", "hybrid", "keyword".
	RetrievalType AICollectionGetDocumentsParamsRetrievalType `query:"retrieval_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AICollectionGetDocumentsParams]'s query parameters as
// `url.Values`.
func (r AICollectionGetDocumentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Override the collection's configured retrieval strategy for this request. Echoed
// back in `meta.retrieval_type`.
type AICollectionGetDocumentsParamsRetrievalType string

const (
	AICollectionGetDocumentsParamsRetrievalTypeVector  AICollectionGetDocumentsParamsRetrievalType = "vector"
	AICollectionGetDocumentsParamsRetrievalTypeHybrid  AICollectionGetDocumentsParamsRetrievalType = "hybrid"
	AICollectionGetDocumentsParamsRetrievalTypeKeyword AICollectionGetDocumentsParamsRetrievalType = "keyword"
)
