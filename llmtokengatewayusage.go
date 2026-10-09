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
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Manage and report AI Gateway traffic.
//
// LlmTokenGatewayUsageService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmTokenGatewayUsageService] method instead.
type LlmTokenGatewayUsageService struct {
	Options []option.RequestOption
}

// NewLlmTokenGatewayUsageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLlmTokenGatewayUsageService(opts ...option.RequestOption) (r LlmTokenGatewayUsageService) {
	r = LlmTokenGatewayUsageService{}
	r.Options = opts
	return
}

// Return complete usage totals, UTC daily and model breakdowns, and guardrail
// event counts for one token group owned by the authenticated account. Requires
// the llm_token_gateway.usage.read permission; spend and guardrail read
// permissions do not grant this combined report. All sections share one database
// snapshot and include the latest usage corrections. Dates use an inclusive start
// and exclusive end spanning 1 to 31 days. Only token_group_id, start_date and
// end_date are accepted; pagination, group_by and other filters are rejected.
// Spend is reference/enforcement USD, not invoice truth or BYOK provider charges.
// Unknown cost is excluded from spend and reported through unknown_requests and
// reserved_spend. Daily rows include zero-activity days. Model rows are ordered by
// request count descending, then model name, and are limited to 1,000. Guardrail
// counts count events, not distinct requests; recent_events contains at most 20
// newest events. A report that exceeds model or query limits returns 503 rather
// than a truncated success.
func (r *LlmTokenGatewayUsageService) GetSummary(ctx context.Context, query LlmTokenGatewayUsageGetSummaryParams, opts ...option.RequestOption) (res *LlmTokenGatewayUsageGetSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/usage/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type LlmTokenGatewayUsageGetSummaryResponse struct {
	Data LlmTokenGatewayUsageGetSummaryResponseData `json:"data" api:"required"`
	Meta LlmTokenGatewayUsageGetSummaryResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseData struct {
	// One row per UTC day, including zero-activity days.
	ByDay []LlmTokenGatewayUsageGetSummaryResponseDataByDay `json:"by_day" api:"required"`
	// One row per model, ordered by request count descending then model name.
	ByModel []LlmTokenGatewayUsageGetSummaryResponseDataByModel `json:"by_model" api:"required"`
	// Complete guardrail event counts and bounded recent findings for the same group
	// and range.
	Guardrails LlmTokenGatewayUsageGetSummaryResponseDataGuardrails `json:"guardrails" api:"required"`
	// Metrics for all matching requests.
	Totals LlmTokenGatewayUsageGetSummaryResponseDataTotals `json:"totals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByDay       respjson.Field
		ByModel     respjson.Field
		Guardrails  respjson.Field
		Totals      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseData) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseDataByDay struct {
	// Requests served from the gateway cache.
	CacheHits int64 `json:"cache_hits" api:"required"`
	// UTC day.
	Date time.Time `json:"date" api:"required" format:"date"`
	// Requests classified as failed.
	FailedRequests int64 `json:"failed_requests" api:"required"`
	// Independently known input tokens across attempts, including corrected usage.
	InputTokens int64 `json:"input_tokens" api:"required"`
	// Independently known output tokens across attempts, including corrected usage.
	OutputTokens int64 `json:"output_tokens" api:"required"`
	// Requests classified as partial after streaming began.
	PartialRequests int64 `json:"partial_requests" api:"required"`
	// Number of matching requests.
	Requests int64 `json:"requests" api:"required"`
	// Unresolved budget reservations in USD.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// Sum of known reference/enforcement cost in USD.
	Spend float64 `json:"spend" api:"required"`
	// Requests classified as succeeded.
	SucceededRequests int64 `json:"succeeded_requests" api:"required"`
	// Requests whose cost remains unresolved; unknown cost is excluded from spend.
	UnknownRequests int64 `json:"unknown_requests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHits         respjson.Field
		Date              respjson.Field
		FailedRequests    respjson.Field
		InputTokens       respjson.Field
		OutputTokens      respjson.Field
		PartialRequests   respjson.Field
		Requests          respjson.Field
		ReservedSpend     respjson.Field
		Spend             respjson.Field
		SucceededRequests respjson.Field
		UnknownRequests   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseDataByDay) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponseDataByDay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseDataByModel struct {
	// Requests served from the gateway cache.
	CacheHits int64 `json:"cache_hits" api:"required"`
	// Requests classified as failed.
	FailedRequests int64 `json:"failed_requests" api:"required"`
	// Independently known input tokens across attempts, including corrected usage.
	InputTokens int64 `json:"input_tokens" api:"required"`
	// Model identifier.
	Model string `json:"model" api:"required"`
	// Independently known output tokens across attempts, including corrected usage.
	OutputTokens int64 `json:"output_tokens" api:"required"`
	// Requests classified as partial after streaming began.
	PartialRequests int64 `json:"partial_requests" api:"required"`
	// Number of matching requests.
	Requests int64 `json:"requests" api:"required"`
	// Unresolved budget reservations in USD.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// Sum of known reference/enforcement cost in USD.
	Spend float64 `json:"spend" api:"required"`
	// Requests classified as succeeded.
	SucceededRequests int64 `json:"succeeded_requests" api:"required"`
	// Requests whose cost remains unresolved; unknown cost is excluded from spend.
	UnknownRequests int64 `json:"unknown_requests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHits         respjson.Field
		FailedRequests    respjson.Field
		InputTokens       respjson.Field
		Model             respjson.Field
		OutputTokens      respjson.Field
		PartialRequests   respjson.Field
		Requests          respjson.Field
		ReservedSpend     respjson.Field
		Spend             respjson.Field
		SucceededRequests respjson.Field
		UnknownRequests   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseDataByModel) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponseDataByModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete guardrail event counts and bounded recent findings for the same group
// and range.
type LlmTokenGatewayUsageGetSummaryResponseDataGuardrails struct {
	// Total blocked guardrail events, not distinct requests.
	BlockedEvents int64 `json:"blocked_events" api:"required"`
	// Total flagged guardrail events, not distinct requests.
	FlaggedEvents int64 `json:"flagged_events" api:"required"`
	// Up to 20 newest privacy-safe guardrail events, ordered by creation time
	// descending and event ID.
	RecentEvents []LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEvent `json:"recent_events" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedEvents respjson.Field
		FlaggedEvents respjson.Field
		RecentEvents  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseDataGuardrails) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponseDataGuardrails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEvent struct {
	ID                     string                                                                    `json:"id" api:"required" format:"uuid"`
	CreatedAt              time.Time                                                                 `json:"created_at" api:"required" format:"date-time"`
	EndUserID              string                                                                    `json:"end_user_id" api:"required"`
	EvaluationInputTokens  int64                                                                     `json:"evaluation_input_tokens" api:"required"`
	EvaluationOutputTokens int64                                                                     `json:"evaluation_output_tokens" api:"required"`
	Findings               []LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsFinding `json:"findings" api:"required"`
	// Model identifier.
	Model string `json:"model" api:"required"`
	// Any of "evaluated", "flagged", "blocked", "unevaluated".
	Outcome    string                  `json:"outcome" api:"required"`
	RecordType constant.GuardrailEvent `json:"record_type" default:"guardrail_event"`
	RequestID  string                  `json:"request_id" api:"required" format:"uuid"`
	// Any of "prompt", "response".
	Stage        string `json:"stage" api:"required"`
	TokenGroupID string `json:"token_group_id" api:"required" format:"uuid"`
	TokenKeyID   string `json:"token_key_id" api:"required" format:"uuid"`
	TokenUserID  string `json:"token_user_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		EndUserID              respjson.Field
		EvaluationInputTokens  respjson.Field
		EvaluationOutputTokens respjson.Field
		Findings               respjson.Field
		Model                  respjson.Field
		Outcome                respjson.Field
		RecordType             respjson.Field
		RequestID              respjson.Field
		Stage                  respjson.Field
		TokenGroupID           respjson.Field
		TokenKeyID             respjson.Field
		TokenUserID            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEvent) RawJSON() string {
	return r.JSON.raw
}
func (r *LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsFinding struct {
	// Any of "flag", "block".
	Action string `json:"action" api:"required"`
	Code   string `json:"code" api:"required"`
	Count  int64  `json:"count" api:"required"`
	// Any of "secrets", "dlp", "safety".
	Detector string `json:"detector" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Code        respjson.Field
		Count       respjson.Field
		Detector    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsFinding) RawJSON() string {
	return r.JSON.raw
}
func (r *LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsFinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metrics for all matching requests.
type LlmTokenGatewayUsageGetSummaryResponseDataTotals struct {
	// Requests served from the gateway cache.
	CacheHits int64 `json:"cache_hits" api:"required"`
	// Requests classified as failed.
	FailedRequests int64 `json:"failed_requests" api:"required"`
	// Independently known input tokens across attempts, including corrected usage.
	InputTokens int64 `json:"input_tokens" api:"required"`
	// Independently known output tokens across attempts, including corrected usage.
	OutputTokens int64 `json:"output_tokens" api:"required"`
	// Requests classified as partial after streaming began.
	PartialRequests int64 `json:"partial_requests" api:"required"`
	// Number of matching requests.
	Requests int64 `json:"requests" api:"required"`
	// Unresolved budget reservations in USD.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// Sum of known reference/enforcement cost in USD.
	Spend float64 `json:"spend" api:"required"`
	// Requests classified as succeeded.
	SucceededRequests int64 `json:"succeeded_requests" api:"required"`
	// Requests whose cost remains unresolved; unknown cost is excluded from spend.
	UnknownRequests int64 `json:"unknown_requests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CacheHits         respjson.Field
		FailedRequests    respjson.Field
		InputTokens       respjson.Field
		OutputTokens      respjson.Field
		PartialRequests   respjson.Field
		Requests          respjson.Field
		ReservedSpend     respjson.Field
		Spend             respjson.Field
		SucceededRequests respjson.Field
		UnknownRequests   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseDataTotals) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponseDataTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseMeta struct {
	EndDate      time.Time `json:"end_date" api:"required" format:"date"`
	StartDate    time.Time `json:"start_date" api:"required" format:"date"`
	TokenGroupID string    `json:"token_group_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate      respjson.Field
		StartDate    respjson.Field
		TokenGroupID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetSummaryResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetSummaryResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryParams struct {
	// Exclusive UTC date in YYYY-MM-DD format. Must follow start_date by 1 to 31 days.
	EndDate time.Time `query:"end_date" api:"required" format:"date" json:"-"`
	// Inclusive UTC date in YYYY-MM-DD format. Must precede end_date by 1 to 31 days.
	StartDate time.Time `query:"start_date" api:"required" format:"date" json:"-"`
	// ID of a token group owned by the authenticated account.
	TokenGroupID string `query:"token_group_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [LlmTokenGatewayUsageGetSummaryParams]'s query parameters as
// `url.Values`.
func (r LlmTokenGatewayUsageGetSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
