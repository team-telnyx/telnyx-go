// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Deep research with citations and async task polling.
//
// WebSearchResearchService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebSearchResearchService] method instead.
type WebSearchResearchService struct {
	Options []option.RequestOption
}

// NewWebSearchResearchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWebSearchResearchService(opts ...option.RequestOption) (r WebSearchResearchService) {
	r = WebSearchResearchService{}
	r.Options = opts
	return
}

// Starts a deep research task that runs multiple searches, reads sources, and
// synthesizes an answer with citations.
//
// ## Synchronous mode (default)
//
// When `background` is `false` or omitted, the request blocks until the research
// completes and returns the answer with citations. This can take up to 120 seconds
// depending on `research_effort`.
//
// ## Asynchronous mode
//
// When `background` is `true`, the request returns immediately with a `task_id`
// and `status: pending`. Poll `GET /web_search/research/{task_id}` to check when
// the research completes and retrieve the answer.
func (r *WebSearchResearchService) New(ctx context.Context, body WebSearchResearchNewParams, opts ...option.RequestOption) (res *WebSearchResearchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web_search/research"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Polls the status of a previously started asynchronous research task. When the
// status is `completed`, the response includes the answer and citations. When the
// status is `failed`, the response includes an error message.
func (r *WebSearchResearchService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *WebSearchResearchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("web_search/research/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ResearchCitation struct {
	// Title of the source page.
	Title string `json:"title" api:"required"`
	// Source URL.
	URL string `json:"url" api:"required" format:"uri"`
	// Relevant excerpt from the source (if available).
	Snippet string `json:"snippet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Title       respjson.Field
		URL         respjson.Field
		Snippet     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResearchCitation) RawJSON() string { return r.JSON.raw }
func (r *ResearchCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchResearchNewResponse struct {
	// Synchronous research response (when `background` is false or unset).
	Data WebSearchResearchNewResponseDataUnion `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResearchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResearchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebSearchResearchNewResponseDataUnion contains all possible properties and
// values from [WebSearchResearchNewResponseDataResearchResponseSync],
// [WebSearchResearchNewResponseDataResearchResponseAsync].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WebSearchResearchNewResponseDataUnion struct {
	// This field is from variant
	// [WebSearchResearchNewResponseDataResearchResponseSync].
	Answer string `json:"answer"`
	// This field is from variant
	// [WebSearchResearchNewResponseDataResearchResponseSync].
	Citations []ResearchCitation `json:"citations"`
	// This field is from variant
	// [WebSearchResearchNewResponseDataResearchResponseAsync].
	Status string `json:"status"`
	// This field is from variant
	// [WebSearchResearchNewResponseDataResearchResponseAsync].
	TaskID string `json:"task_id"`
	JSON   struct {
		Answer    respjson.Field
		Citations respjson.Field
		Status    respjson.Field
		TaskID    respjson.Field
		raw       string
	} `json:"-"`
}

func (u WebSearchResearchNewResponseDataUnion) AsWebSearchResearchNewResponseDataResearchResponseSync() (v WebSearchResearchNewResponseDataResearchResponseSync) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebSearchResearchNewResponseDataUnion) AsWebSearchResearchNewResponseDataResearchResponseAsync() (v WebSearchResearchNewResponseDataResearchResponseAsync) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebSearchResearchNewResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *WebSearchResearchNewResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Synchronous research response (when `background` is false or unset).
type WebSearchResearchNewResponseDataResearchResponseSync struct {
	// The synthesized research answer.
	Answer string `json:"answer" api:"required"`
	// Sources cited in the answer.
	Citations []ResearchCitation `json:"citations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		Citations   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResearchNewResponseDataResearchResponseSync) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResearchNewResponseDataResearchResponseSync) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Asynchronous research response (when `background` is true).
type WebSearchResearchNewResponseDataResearchResponseAsync struct {
	// Current status of the research task.
	//
	// Any of "pending", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// Unique identifier for the research task. Use this to poll the status.
	TaskID string `json:"task_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResearchNewResponseDataResearchResponseAsync) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResearchNewResponseDataResearchResponseAsync) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchResearchGetResponse struct {
	Data WebSearchResearchGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResearchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResearchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchResearchGetResponseData struct {
	// Current status of the research task.
	//
	// Any of "pending", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// The research task identifier.
	TaskID string `json:"task_id" api:"required"`
	// The synthesized research answer (present when status is `completed`).
	Answer string `json:"answer"`
	// Sources cited in the answer (present when status is `completed`).
	Citations []ResearchCitation `json:"citations"`
	// Always present in poll responses; `null` unless the task failed.
	Error string `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		TaskID      respjson.Field
		Answer      respjson.Field
		Citations   respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebSearchResearchGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WebSearchResearchGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebSearchResearchNewParams struct {
	// The research question or topic.
	Query string `json:"query" api:"required"`
	// When `true`, the research runs asynchronously. The response returns a `task_id`
	// immediately instead of waiting for the result. Poll
	// `GET /web_search/research/{task_id}` to check status.
	Background param.Opt[bool] `json:"background,omitzero"`
	// Maximum number of sources to use.
	MaxSources param.Opt[int64] `json:"max_sources,omitzero"`
	// Research depth level. `lite` is fastest, `deep` is most thorough.
	//
	// Any of "lite", "standard", "deep".
	ResearchEffort WebSearchResearchNewParamsResearchEffort `json:"research_effort,omitzero"`
	paramObj
}

func (r WebSearchResearchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebSearchResearchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebSearchResearchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Research depth level. `lite` is fastest, `deep` is most thorough.
type WebSearchResearchNewParamsResearchEffort string

const (
	WebSearchResearchNewParamsResearchEffortLite     WebSearchResearchNewParamsResearchEffort = "lite"
	WebSearchResearchNewParamsResearchEffortStandard WebSearchResearchNewParamsResearchEffort = "standard"
	WebSearchResearchNewParamsResearchEffortDeep     WebSearchResearchNewParamsResearchEffort = "deep"
)
