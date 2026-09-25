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

// What a namespace and a profile hold.
//
// AIMemoryNamespaceProfileMemoryService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMemoryNamespaceProfileMemoryService] method instead.
type AIMemoryNamespaceProfileMemoryService struct {
	Options []option.RequestOption
}

// NewAIMemoryNamespaceProfileMemoryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIMemoryNamespaceProfileMemoryService(opts ...option.RequestOption) (r AIMemoryNamespaceProfileMemoryService) {
	r = AIMemoryNamespaceProfileMemoryService{}
	r.Options = opts
	return
}

// One memory by its id, as `recall` and the listing return it, together with what
// it came from. A fact names its `source_id`: read it with
// `GET .../sources/{source_id}` to see what was stored. A memory derived from
// other memories names them in `derived_from` instead; read each of those to reach
// its source.
func (r *AIMemoryNamespaceProfileMemoryService) Get(ctx context.Context, memoryID string, query AIMemoryNamespaceProfileMemoryGetParams, opts ...option.RequestOption) (res *AIMemoryNamespaceProfileMemoryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.Namespace == "" {
		err = errors.New("missing required namespace parameter")
		return nil, err
	}
	if query.ProfileID == "" {
		err = errors.New("missing required profile_id parameter")
		return nil, err
	}
	if memoryID == "" {
		err = errors.New("missing required memory_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/memories/%s", url.PathEscape(query.Namespace), url.PathEscape(query.ProfileID), url.PathEscape(memoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Everything stored under one profile, unranked -- ask `recall` for the memories
// that answer a question. A profile that holds nothing is an empty page rather
// than a 404: profiles exist by being written to. Each memory names the
// `source_id` it was extracted from, or null for a memory derived from other
// memories -- which can read almost the same as the fact it restates. A
// `source_id` narrows the listing to the memories extracted from that source, and
// a `session_id` to those extracted from the session, which is the same thing
// named another way; pass one or the other. Neither is everything the source led
// to: a memory derived from several sources belongs to no single one and appears
// only in the unfiltered listing. A memory written while the listing is paged
// shifts the pages after it, so an entry can be repeated or missed at a page
// boundary.
func (r *AIMemoryNamespaceProfileMemoryService) List(ctx context.Context, profileID string, params AIMemoryNamespaceProfileMemoryListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIMemoryNamespaceProfileMemoryListResponse], err error) {
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
	path := fmt.Sprintf("ai/memory/namespaces/%s/profiles/%s/memories", url.PathEscape(params.Namespace), url.PathEscape(profileID))
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

// Everything stored under one profile, unranked -- ask `recall` for the memories
// that answer a question. A profile that holds nothing is an empty page rather
// than a 404: profiles exist by being written to. Each memory names the
// `source_id` it was extracted from, or null for a memory derived from other
// memories -- which can read almost the same as the fact it restates. A
// `source_id` narrows the listing to the memories extracted from that source, and
// a `session_id` to those extracted from the session, which is the same thing
// named another way; pass one or the other. Neither is everything the source led
// to: a memory derived from several sources belongs to no single one and appears
// only in the unfiltered listing. A memory written while the listing is paged
// shifts the pages after it, so an entry can be repeated or missed at a page
// boundary.
func (r *AIMemoryNamespaceProfileMemoryService) ListAutoPaging(ctx context.Context, profileID string, params AIMemoryNamespaceProfileMemoryListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIMemoryNamespaceProfileMemoryListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, profileID, params, opts...))
}

type AIMemoryNamespaceProfileMemoryGetResponse struct {
	Data AIMemoryNamespaceProfileMemoryGetResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileMemoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileMemoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileMemoryGetResponseData struct {
	ID string `json:"id" api:"required"`
	// The ids of the memories this one was derived from. Read each with
	// `GET .../memories/{memory_id}` to reach its `source_id`. Set for a derived
	// memory; null for a fact.
	DerivedFrom []string `json:"derived_from" api:"required"`
	// The source this memory was extracted from. Set for a fact, which comes from
	// exactly one source; null for a memory derived from other memories. Read it with
	// `GET .../sources/{source_id}`. A source deleted a moment ago can still be named
	// here, and then answers 404.
	SourceID   string `json:"source_id" api:"required"`
	Text       string `json:"text" api:"required"`
	RecordedAt string `json:"recorded_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		DerivedFrom respjson.Field
		SourceID    respjson.Field
		Text        respjson.Field
		RecordedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileMemoryGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileMemoryGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileMemoryListResponse struct {
	ID string `json:"id" api:"required"`
	// The source this memory was extracted from. Set for a fact, which comes from
	// exactly one source; null for a memory derived from other memories. Read it with
	// `GET .../sources/{source_id}`. A source deleted a moment ago can still be named
	// here, and then answers 404.
	SourceID   string `json:"source_id" api:"required"`
	Text       string `json:"text" api:"required"`
	RecordedAt string `json:"recorded_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		SourceID    respjson.Field
		Text        respjson.Field
		RecordedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMemoryNamespaceProfileMemoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMemoryNamespaceProfileMemoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMemoryNamespaceProfileMemoryGetParams struct {
	// The namespace. `default` exists for every organization.
	Namespace string `path:"namespace" api:"required" json:"-"`
	// The profile: your identifier for the user, caller or agent this memory is about.
	ProfileID string `path:"profile_id" api:"required" json:"-"`
	paramObj
}

type AIMemoryNamespaceProfileMemoryListParams struct {
	Namespace string `path:"namespace" api:"required" json:"-"`
	// An ingested session, by the `session_id` it was ingested with. Narrows the
	// request to the source that session was stored as.
	SessionID param.Opt[string] `query:"session_id,omitzero" json:"-"`
	// Narrows the listing to the memories extracted from one source, a remembered fact
	// as well as a session. Pass this or `session_id`, not both.
	SourceID param.Opt[string] `query:"source_id,omitzero" json:"-"`
	// The page to return, counting from 1. Bounded in depth: (page[number] - 1) \*
	// page[size] may be at most 10000.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// How many results a page holds.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMemoryNamespaceProfileMemoryListParams]'s query
// parameters as `url.Values`.
func (r AIMemoryNamespaceProfileMemoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
