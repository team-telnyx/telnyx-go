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

// WebSearchService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebSearchService] method instead.
type WebSearchService struct {
	Options []option.RequestOption
	// Deep research with citations and async task polling.
	Research WebSearchResearchService
}

// NewWebSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebSearchService(opts ...option.RequestOption) (r WebSearchService) {
	r = WebSearchService{}
	r.Options = opts
	r.Research = NewWebSearchResearchService(opts...)
	return
}

// Performs a real-time web search and returns structured, LLM-ready JSON results
// with titles, URLs, descriptions, and snippets. Supports filtering by domain,
// country, safe search, freshness, and live crawl.
//
// **Note:** `include_domains` and `exclude_domains` cannot be used in the same
// request. Use one or the other.
func (r *WebSearchService) New(ctx context.Context, body WebSearchNewParams, opts ...option.RequestOption) (res *WebSearchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web_search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves clean HTML or Markdown content from a list of URLs. Supports up to 20
// URLs per request (public API limit). Specify which formats to return: `html`,
// `markdown`, `metadata`.
func (r *WebSearchService) Contents(ctx context.Context, body WebSearchContentsParams, opts ...option.RequestOption) (res *WebSearchContentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web_search/contents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type WebSearchResult struct {
	// Short description or excerpt.
	Description string `json:"description" api:"required"`
	// Relevant text snippets from the page.
	Snippets []string `json:"snippets" api:"required"`
	// Result title.
	Title string `json:"title" api:"required"`
	// Result URL.
	URL string `json:"url" api:"required" format:"uri"`
	// Favicon URL (if available).
	FaviconURL string `json:"favicon_url" format:"uri"`
	// Thumbnail image URL (if available).
	ThumbnailURL string `json:"thumbnail_url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		Snippets     respjson.Field
		Title        respjson.Field
		URL          respjson.Field
		FaviconURL   respjson.Field
		ThumbnailURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResult) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchNewResponse struct {
	Data WebSearchNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebSearchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchNewResponseData struct {
	Results WebSearchNewResponseDataResults `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebSearchNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchNewResponseDataResults struct {
	// Web search results.
	Web []WebSearchResult `json:"web" api:"required"`
	// News search results. Present only when the query surfaces news results.
	News []WebSearchResult `json:"news"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Web         respjson.Field
		News        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchNewResponseDataResults) RawJSON() string { return r.JSON.raw }
func (r *WebSearchNewResponseDataResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchContentsResponse struct {
	Data WebSearchContentsResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchContentsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebSearchContentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchContentsResponseData struct {
	Results []WebSearchContentsResponseDataResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchContentsResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebSearchContentsResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchContentsResponseDataResult struct {
	// The source URL.
	URL string `json:"url" api:"required" format:"uri"`
	// Cleaned HTML content (if `html` format requested; may also be present on freshly
	// crawled pages).
	HTML string `json:"html"`
	// Markdown content (if `markdown` format requested).
	Markdown string `json:"markdown"`
	// Page metadata (if `metadata` format requested).
	Metadata WebSearchContentsResponseDataResultsMetadata `json:"metadata"`
	// Page title (if available).
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		HTML        respjson.Field
		Markdown    respjson.Field
		Metadata    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchContentsResponseDataResult) RawJSON() string { return r.JSON.raw }
func (r *WebSearchContentsResponseDataResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page metadata (if `metadata` format requested).
type WebSearchContentsResponseDataResultsMetadata struct {
	// Favicon URL (if available).
	FaviconURL string `json:"favicon_url" format:"uri"`
	// Site name. Often empty.
	SiteName    string         `json:"site_name"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FaviconURL  respjson.Field
		SiteName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchContentsResponseDataResultsMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebSearchContentsResponseDataResultsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchNewParams struct {
	// The search query text.
	Query string `json:"query" api:"required"`
	// Number of results to return (1-100).
	Count param.Opt[int64] `json:"count,omitzero"`
	// Two-letter country code (ISO 3166-1 alpha-2) to bias results.
	Country param.Opt[string] `json:"country,omitzero"`
	// Time-based filter for results. Common values: `day`, `week`, `month`, `year`.
	Freshness param.Opt[string] `json:"freshness,omitzero"`
	// When true, the provider crawls pages in real-time for fresh content. The boolean
	// is translated to the provider's internal enum internally; callers always pass
	// `true` or `false`.
	Livecrawl param.Opt[bool] `json:"livecrawl,omitzero"`
	// Exclude results from these domains (bare hostnames, e.g. `pinterest.com`).
	ExcludeDomains []string `json:"exclude_domains,omitzero"`
	// Restrict results to these domains (bare hostnames, e.g. `arxiv.org`).
	IncludeDomains []string `json:"include_domains,omitzero"`
	// Safe search filter level.
	//
	// Any of "off", "moderate", "strict".
	Safesearch WebSearchNewParamsSafesearch `json:"safesearch,omitzero"`
	paramObj
}

func (r WebSearchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Safe search filter level.
type WebSearchNewParamsSafesearch string

const (
	WebSearchNewParamsSafesearchOff      WebSearchNewParamsSafesearch = "off"
	WebSearchNewParamsSafesearchModerate WebSearchNewParamsSafesearch = "moderate"
	WebSearchNewParamsSafesearchStrict   WebSearchNewParamsSafesearch = "strict"
)

type WebSearchContentsParams struct {
	// List of URLs to retrieve content from (max 20 for public API).
	URLs []string `json:"urls,omitzero" api:"required" format:"uri"`
	// Maximum age of cached content in seconds. `null` means no limit.
	MaxAge param.Opt[int64] `json:"max_age,omitzero"`
	// Timeout for crawling each URL, in seconds (1-60).
	CrawlTimeout param.Opt[int64] `json:"crawl_timeout,omitzero"`
	// Content formats to return. If omitted, `html` and `metadata` are returned by
	// default. Retrieval is best-effort per URL: a format field appears only when that
	// content could be produced, and a freshly crawled page may also include `html`
	// even when not requested.
	//
	// Any of "html", "markdown", "metadata".
	Formats []string `json:"formats,omitzero"`
	paramObj
}

func (r WebSearchContentsParams) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchContentsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchContentsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
