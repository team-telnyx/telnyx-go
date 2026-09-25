// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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

// What a profile stored, and what its memories came from.
//
// AIMemoryNamespaceProfileSourceService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemoryNamespaceProfileSourceService] method instead.
type AIMemoryNamespaceProfileSourceService struct {
	Options []option.RequestOption
}

// NewAIMemoryNamespaceProfileSourceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIMemoryNamespaceProfileSourceService(opts ...option.RequestOption) (r AIMemoryNamespaceProfileSourceService) {
	r = AIMemoryNamespaceProfileSourceService{}
	r.Options = opts
	return
}

// One source and its content, as it was stored: an ingested session's payload or a
// remembered fact. A source whose ingest is still queued answers 404 until it has
// been stored.
func (r *AIMemoryNamespaceProfileSourceService) Get(ctx context.Context, sourceID string, query AIMemoryNamespaceProfileSourceGetParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileSourceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if query.ProfileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/sources/%s", url.PathEscape(query.Namespace), url.PathEscape(query.ProfileID), url.PathEscape(sourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Everything a profile has stored and extracts memories from: each ingested
// session, and each remembered fact, which has no session. Content is not listed;
// read one source for it. A source whose ingest is still queued is not here yet.
// Re-ingesting a session moves it to the front, so a listing paged while sessions
// are written can repeat or miss one at a page boundary. A `session_id` narrows
// the listing to the source that session was stored as: one source or none, and
// none -- an empty page, not a 404 -- for a session never ingested, still queued,
// or another profile's.
func (r *AIMemoryNamespaceProfileSourceService) List(ctx context.Context, profileID string, params AIMemoryNamespaceProfileSourceListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Source], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/sources", url.PathEscape(params.Namespace), url.PathEscape(profileID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Everything a profile has stored and extracts memories from: each ingested
// session, and each remembered fact, which has no session. Content is not listed;
// read one source for it. A source whose ingest is still queued is not here yet.
// Re-ingesting a session moves it to the front, so a listing paged while sessions
// are written can repeat or miss one at a page boundary. A `session_id` narrows
// the listing to the source that session was stored as: one source or none, and
// none -- an empty page, not a 404 -- for a session never ingested, still queued,
// or another profile's.
func (r *AIMemoryNamespaceProfileSourceService) ListAutoPaging(ctx context.Context, profileID string, params AIMemoryNamespaceProfileSourceListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Source] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, profileID, params, opts...))
}

// Deletes one source -- an ingested session or a remembered fact -- together with
// the memories derived from it. A memory derived from this source and others is
// deleted too, and derived again from what remains in the background. It answers
// only once the source is gone. A source that is not there -- never stored,
// another profile's, or already deleted -- answers 404, so on a `502` or a `504`
// repeat the identical request and read a 404 as done. An ingest of the same
// session that is still queued is not cancelled, and stores the session again when
// it runs. Nothing here can be undone.
func (r *AIMemoryNamespaceProfileSourceService) Delete(ctx context.Context, sourceID string, body AIMemoryNamespaceProfileSourceDeleteParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileSourceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if body.ProfileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	if sourceID == "" {
		err = errors.New("missing required source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/sources/%s", url.PathEscape(body.Namespace), url.PathEscape(body.ProfileID), url.PathEscape(sourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AIMemoryNamespaceProfileSourceGetResponse struct {
	Data AIMemoryNamespaceProfileSourceGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileSourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileSourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileSourceGetResponseData struct {
	// Identifies one source within its profile: an ingested session, or one remembered
	// fact. Returned by `ingest` and `remember` when the write is accepted.
	// Re-ingesting a session keeps its source id.
	ID string `json:"id" api:"required"`
	// What was stored, in the shape it was sent: an ingested JSON body as JSON, a
	// string body or a remembered fact as a string. A session ingested before formats
	// were recorded is returned as the text it was stored as.
	Content AIMemoryNamespaceProfileSourceGetResponseDataContentUnion `json:"content" api:"required"`
	// Memories extracted from this source. A memory derived from several sources is
	// not counted here.
	MemoryCount int64 `json:"memory_count" api:"required"`
	// The session this source was ingested as. Null for a remembered fact.
	SessionID string `json:"session_id" api:"required"`
	// When the source was first stored.
	CreatedAt string `json:"created_at" api:"nullable"`
	// When the source was last written; re-ingesting moves it.
	UpdatedAt string `json:"updated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		MemoryCount respjson.Field
		SessionID   respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileSourceGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileSourceGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AIMemoryNamespaceProfileSourceGetResponseDataContentUnion contains all possible
// properties and values from [map[string]any],
// [AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfAIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember0 OfAnyArray
// OfString OfFloat OfBool]
type AIMemoryNamespaceProfileSourceGetResponseDataContentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember0 any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfAIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember0 respjson.Field
		OfAnyArray                                                         respjson.Field
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnion) AsAIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1() (v AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AIMemoryNamespaceProfileSourceGetResponseDataContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union contains
// all possible properties and values from [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfAnyArray respjson.Field
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		raw        string
	} `json:"-"`
}

func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) RawJSON() string {
	return u.JSON.raw
}

func (r *AIMemoryNamespaceProfileSourceGetResponseDataContentUnionMember1Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileSourceDeleteResponse struct {
	Data AIMemoryNamespaceProfileSourceDeleteResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileSourceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileSourceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileSourceDeleteResponseData struct {
	// Memories the profile held and no longer does, counted before and after across
	// the whole profile: it includes memories derived from this source together with
	// others, and anything else the profile lost in between. A report rather than an
	// audit. The status carries the outcome.
	MemoriesDeleted int64  `json:"memories_deleted" api:"required"`
	ProfileID       string `json:"profile_id" api:"required"`
	// Identifies one source within its profile: an ingested session, or one remembered
	// fact. Returned by `ingest` and `remember` when the write is accepted.
	// Re-ingesting a session keeps its source id.
	SourceID string `json:"source_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoriesDeleted respjson.Field
		ProfileID       respjson.Field
		SourceID        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileSourceDeleteResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileSourceDeleteResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileSourceGetParams struct {
	// The namespace. `default` exists for every organization.
	Namespace string `path:"namespace" api:"required" json:"-"`
	// The profile: your identifier for the user, caller or agent this memory is about.
	ProfileID string `path:"profile_id" api:"required" json:"-"`
	paramObj
}

type AIMemoryNamespaceProfileSourceListParams struct {
	Namespace string `path:"namespace" api:"required" json:"-"`
	// An ingested session, by the `session_id` it was ingested with. Narrows the
	// request to the source that session was stored as.
	SessionID param.Opt[string] `query:"session_id,omitzero" json:"-"`
	// The page to return, counting from 1. Bounded in depth: (page[number] - 1) \*
	// page[size] may be at most 10000.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// How many results a page holds.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMemoryNamespaceProfileSourceListParams]'s query
// parameters as `url.Values`.
func (r AIMemoryNamespaceProfileSourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMemoryNamespaceProfileSourceDeleteParams struct {
	// The namespace. `default` exists for every organization.
	Namespace string `path:"namespace" api:"required" json:"-"`
	// The profile: your identifier for the user, caller or agent this memory is about.
	ProfileID string `path:"profile_id" api:"required" json:"-"`
	paramObj
}
