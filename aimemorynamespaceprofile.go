// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// AIMemoryNamespaceProfileService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemoryNamespaceProfileService] method instead.
type AIMemoryNamespaceProfileService struct {
	Options []option.RequestOption
	// What a namespace and a profile hold.
	Memories AIMemoryNamespaceProfileMemoryService
	// What a profile stored, and what its memories came from.
	Sources AIMemoryNamespaceProfileSourceService
}

// NewAIMemoryNamespaceProfileService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIMemoryNamespaceProfileService(opts ...option.RequestOption) (r AIMemoryNamespaceProfileService) {
	r = AIMemoryNamespaceProfileService{}
	r.Options = opts
	r.Memories = NewAIMemoryNamespaceProfileMemoryService(opts...)
	r.Sources = NewAIMemoryNamespaceProfileSourceService(opts...)
	return
}

// Profiles are never created, only written to, so this lists the ones that hold a
// memory. A profile whose first ingest is still running is not here yet. Ordered
// by memory count, largest first, so a profile written to while the listing is
// paged can move between pages and be repeated or missed.
func (r *AIMemoryNamespaceProfileService) List(ctx context.Context, namespace string, query AIMemoryNamespaceProfileListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIMemoryNamespaceProfileListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles", url.PathEscape(namespace))
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

// Profiles are never created, only written to, so this lists the ones that hold a
// memory. A profile whose first ingest is still running is not here yet. Ordered
// by memory count, largest first, so a profile written to while the listing is
// paged can move between pages and be repeated or missed.
func (r *AIMemoryNamespaceProfileService) ListAutoPaging(ctx context.Context, namespace string, query AIMemoryNamespaceProfileListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIMemoryNamespaceProfileListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, namespace, query, opts...))
}

// Delete everything held about one profile. A 2xx means none of its memories are
// left, and its summary goes with them. There is no undo.
func (r *AIMemoryNamespaceProfileService) Delete(ctx context.Context, profileID string, body AIMemoryNamespaceProfileDeleteParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s", url.PathEscape(body.Namespace), url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Store a session. Facts are extracted from whatever you send — the body is taken
// as any JSON value and stored whole, so a framework's own transcript shape works
// unchanged. `messages` of `role`/`content` is the conventional shape, not a
// requirement. An empty object or a null body is refused. Carry a `session_id` to
// name the session: re-ingesting the same one replaces what it held. Omit it and a
// session is opened and returned. Extraction runs asynchronously — poll the
// returned operation.
func (r *AIMemoryNamespaceProfileService) Ingest(ctx context.Context, profileID string, params AIMemoryNamespaceProfileIngestParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileIngestResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/ingest", url.PathEscape(params.Namespace), url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Ranked memories for a question. Matching runs over the profile's memories and
// returns them in rank order with a relevance `score`; the score is null where the
// deployment's reranker is a passthrough, in which case order is the only signal.
// No model runs in this path — recall returns facts, it does not compose an
// answer.
func (r *AIMemoryNamespaceProfileService) Recall(ctx context.Context, profileID string, params AIMemoryNamespaceProfileRecallParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileRecallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/recall", url.PathEscape(params.Namespace), url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// For a fact the agent has already distilled: `text` is stored as given, with
// nothing extracted from it. Send a transcript to `ingest` instead. Remembering
// the same text again writes the same memory rather than a second copy of it, so a
// retry is safe. The write runs asynchronously -- poll the returned operation.
func (r *AIMemoryNamespaceProfileService) Remember(ctx context.Context, profileID string, params AIMemoryNamespaceProfileRememberParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileRememberResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/remember", url.PathEscape(params.Namespace), url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// The whole profile as one card, precomputed, with no query. Built for the start
// of a session, where there is no question to ask yet.
//
// A summary is generated in the background. `is_stale` tells you newer memories
// have arrived since it was written; that is ordinary and the card is still
// usable.
func (r *AIMemoryNamespaceProfileService) GetSummary(ctx context.Context, profileID string, query AIMemoryNamespaceProfileGetSummaryParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileGetSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/summary", url.PathEscape(query.Namespace), url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Where a listing's page sits in the whole.
//
// A page is a snapshot: the counts it reports and the order it is drawn in both
// move as writes land, so paging through a busy namespace can repeat or miss an
// entry at a page boundary.
type PageMeta struct {
	// The page returned, counting from 1.
	PageNumber int64 `json:"page_number" api:"required"`
	// How many results a page holds.
	PageSize int64 `json:"page_size" api:"required"`
	// Pages that can be requested; 0 when nothing matched. Page until `page_number`
	// reaches it rather than until a page comes back short: a page can hold fewer than
	// `page_size` results without being the last. Capped at the deepest page served,
	// so on a very large listing it covers fewer results than `total_results`.
	TotalPages int64 `json:"total_pages" api:"required"`
	// Results the request matched, including any past the deepest page.
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
func (r PageMeta) RawJSON() string { return r.JSON.raw }
func (r *PageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileListResponse struct {
	// Memories stored under this profile, including the consolidated ones that
	// paraphrase others. Listings are ordered by it.
	MemoryCount int64  `json:"memory_count" api:"required"`
	ProfileID   string `json:"profile_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryCount respjson.Field
		ProfileID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileDeleteResponse struct {
	Data AIMemoryNamespaceProfileDeleteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileDeleteResponseData struct {
	// Memories the profile held and no longer does, counted before and after. A report
	// rather than an audit: memory moves in the background between the two counts. The
	// status carries the outcome.
	MemoriesDeleted int64  `json:"memories_deleted" api:"required"`
	ProfileID       string `json:"profile_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoriesDeleted respjson.Field
		ProfileID       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileIngestResponse struct {
	Data AIMemoryNamespaceProfileIngestResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileIngestResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileIngestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileIngestResponseData struct {
	OperationID string `json:"operation_id" api:"required"`
	ProfileID   string `json:"profile_id" api:"required"`
	SessionID   string `json:"session_id" api:"required"`
	// Identifies one source within its profile: an ingested session, or one remembered
	// fact. Returned by `ingest` and `remember` when the write is accepted.
	// Re-ingesting a session keeps its source id.
	SourceID string `json:"source_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OperationID respjson.Field
		ProfileID   respjson.Field
		SessionID   respjson.Field
		SourceID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileIngestResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileIngestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileRecallResponse struct {
	Data []AIMemoryNamespaceProfileRecallResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileRecallResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileRecallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileRecallResponseData struct {
	ID         string `json:"id" api:"required"`
	Text       string `json:"text" api:"required"`
	RecordedAt string `json:"recorded_at" api:"nullable"`
	// Relevance, 0-1. Null where the deployment's reranker is a passthrough; results
	// are in rank order either way.
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Text        respjson.Field
		RecordedAt  respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileRecallResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileRecallResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileRememberResponse struct {
	Data AIMemoryNamespaceProfileRememberResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileRememberResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileRememberResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileRememberResponseData struct {
	OperationID string `json:"operation_id" api:"required"`
	ProfileID   string `json:"profile_id" api:"required"`
	// Identifies one source within its profile: an ingested session, or one remembered
	// fact. Returned by `ingest` and `remember` when the write is accepted.
	// Re-ingesting a session keeps its source id.
	SourceID string `json:"source_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OperationID respjson.Field
		ProfileID   respjson.Field
		SourceID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileRememberResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileRememberResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileGetSummaryResponse struct {
	Data AIMemoryNamespaceProfileGetSummaryResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileGetSummaryResponseData struct {
	// Whether newer memories have arrived since the summary was generated. The summary
	// is regenerated in the background, so a true here is ordinary and the summary is
	// still usable.
	IsStale   bool   `json:"is_stale" api:"required"`
	ProfileID string `json:"profile_id" api:"required"`
	// When the summary was last generated. Null while none is ready.
	GeneratedAt string `json:"generated_at" api:"nullable"`
	// The precomputed summary, ready to place in an assistant's context at the start
	// of a session.
	Text string `json:"text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsStale     respjson.Field
		ProfileID   respjson.Field
		GeneratedAt respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileGetSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileGetSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileListParams struct {
	// The page to return, counting from 1. Bounded in depth: (page[number] - 1) \*
	// page[size] may be at most 10000.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// How many results a page holds.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMemoryNamespaceProfileListParams]'s query parameters as
// `url.Values`.
func (r AIMemoryNamespaceProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMemoryNamespaceProfileDeleteParams struct {
	// The namespace. `default` exists for every organization.
	Namespace string `path:"namespace" api:"required" json:"-"`
	paramObj
}

type AIMemoryNamespaceProfileIngestParams struct {
	Namespace string `path:"namespace" api:"required" json:"-"`

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfAnyMap map[string]any `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	Of *AIMemoryNamespaceProfileIngestParamsBodyUnionMember1Union `json:",inline"`

	// Names the session. Re-ingesting the same session replaces what it held and keeps
	// its `source_id`. Omit it to have one derived from the content and returned. No
	// whitespace, control characters, or any of / \ # ? %.
	SessionID param.Opt[string] `query:"session_id,omitzero" json:"-"`
	paramObj
}

func (u AIMemoryNamespaceProfileIngestParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.Of)
}
func (r *AIMemoryNamespaceProfileIngestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AIMemoryNamespaceProfileIngestParams]'s query parameters as
// `url.Values`.
func (r AIMemoryNamespaceProfileIngestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIMemoryNamespaceProfileIngestParamsBodyUnionMember1Union struct {
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u AIMemoryNamespaceProfileIngestParamsBodyUnionMember1Union) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyArray, u.OfString, u.OfFloat, u.OfBool)
}
func (u *AIMemoryNamespaceProfileIngestParamsBodyUnionMember1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIMemoryNamespaceProfileIngestParamsBodyUnionMember1Union) asAny() any {
	if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type AIMemoryNamespaceProfileRecallParams struct {
	Namespace string           `path:"namespace" api:"required" json:"-"`
	Query     string           `json:"query" api:"required"`
	TopK      param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r AIMemoryNamespaceProfileRecallParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMemoryNamespaceProfileRecallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMemoryNamespaceProfileRecallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileRememberParams struct {
	Namespace string `path:"namespace" api:"required" json:"-"`
	Text      string `json:"text" api:"required"`
	paramObj
}

func (r AIMemoryNamespaceProfileRememberParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMemoryNamespaceProfileRememberParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMemoryNamespaceProfileRememberParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileGetSummaryParams struct {
	// The namespace. `default` exists for every organization.
	Namespace string `path:"namespace" api:"required" json:"-"`
	paramObj
}
