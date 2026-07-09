// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage historical AI assistant conversations
//
// AIConversationConversationInsightService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationConversationInsightService] method instead.
type AIConversationConversationInsightService struct {
	Options []option.RequestOption
}

// NewAIConversationConversationInsightService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIConversationConversationInsightService(opts ...option.RequestOption) (r AIConversationConversationInsightService) {
	r = AIConversationConversationInsightService{}
	r.Options = opts
	return
}

// Aggregate conversation insights by specified fields
func (r *AIConversationConversationInsightService) GetAggregates(ctx context.Context, query AIConversationConversationInsightGetAggregatesParams, opts ...option.RequestOption) (res *AIConversationConversationInsightGetAggregatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/conversations/conversation-insights/aggregates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Aggregated conversation insight counts grouped by the specified fields. Each
// item in `data` contains the grouped field values and a `record_count` indicating
// how many conversation insights match that combination.
type AIConversationConversationInsightGetAggregatesResponse struct {
	// Aggregation result rows. Each row contains the grouped field values and a
	// `record_count`.
	Data []AIConversationConversationInsightGetAggregatesResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationConversationInsightGetAggregatesResponse) RawJSON() string { return r.JSON.raw }
func (r *AIConversationConversationInsightGetAggregatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation row. Contains the grouped field values (keyed by the group_by
// field names) and a `record_count` integer. For example, when grouping by
// `score`, each row has a `score` value and a `record_count` of conversations with
// that score. When also splitting by `metadata.assistant_version_id`, each row
// includes both `score` and `metadata.assistant_version_id` plus their combined
// `record_count`.
type AIConversationConversationInsightGetAggregatesResponseData struct {
	// Number of conversation insights that match this combination of grouped field
	// values.
	RecordCount int64          `json:"record_count" api:"required"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RecordCount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIConversationConversationInsightGetAggregatesResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *AIConversationConversationInsightGetAggregatesResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIConversationConversationInsightGetAggregatesParams struct {
	// Filter by creation datetime to scope the aggregation window. Supports range
	// operators (e.g., `created_at=gte.2025-01-01T00:00:00Z` for the start of the
	// range, `created_at=lt.2025-01-02T00:00:00Z` for the end). To build per-day time
	// series (as the portal does for the 'Insights Over Time' chart), issue one
	// request per day bounded by `created_at=gte.<day_start>` and
	// `created_at=lt.<next_day_start>`.
	CreatedAt param.Opt[string] `query:"created_at,omitzero" json:"-"`
	// Optional insight ID to filter conversation insights. Only insights matching this
	// ID will be included in the aggregation.
	InsightID param.Opt[string] `query:"insight_id,omitzero" format:"uuid" json:"-"`
	// Fields to group by (can be comma-separated or multiple parameters). Prefix a
	// field with 'metadata.' (e.g. 'metadata.assistant_id') to group by the
	// conversation's metadata instead of the insight result.
	//
	// Common fields used for over-time charts:
	//
	//   - `score` — Group by the insight's score value (e.g. for Agent Instruction
	//     Following, User Satisfaction).
	//   - `metadata.assistant_id` — Group by the assistant that handled the
	//     conversation.
	//   - `metadata.assistant_version_id` — Group by the assistant version, useful for
	//     comparing performance across versions in the portal's 'Insights Over Time'
	//     chart.
	//   - `metadata.telnyx_conversation_channel` — Group by conversation channel
	//     (phone_call, web_chat, etc.).
	GroupBy  []string                                                     `query:"group_by,omitzero" json:"-"`
	Metadata AIConversationConversationInsightGetAggregatesParamsMetadata `query:"metadata,omitzero" json:"-"`
	// Fields to include in the result (can be comma-separated or multiple parameters).
	// Supports the same 'metadata.<key>' prefix as group_by. Each returned row will
	// contain the grouped field values plus a `record_count` indicating how many
	// conversation insights match that combination.
	Show []string `query:"show,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIConversationConversationInsightGetAggregatesParams]'s
// query parameters as `url.Values`.
func (r AIConversationConversationInsightGetAggregatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIConversationConversationInsightGetAggregatesParamsMetadata struct {
	// Filter by assistant ID (e.g., `metadata.assistant_id=eq.<assistant_id>`). When
	// provided, only conversation insights for the specified assistant are aggregated.
	// Used by the portal to scope the 'Insights Over Time' chart to a single
	// assistant.
	AssistantID param.Opt[string] `query:"assistant_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes
// [AIConversationConversationInsightGetAggregatesParamsMetadata]'s query
// parameters as `url.Values`.
func (r AIConversationConversationInsightGetAggregatesParamsMetadata) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
