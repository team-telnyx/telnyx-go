// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Manage and report AI Gateway usage.
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

// Durable, replayable, content-free usage export for customer pipelines. Requires
// the dedicated llm_token_gateway.usage_export.read V2 action; spend.read and
// usage.read do not grant it. Records are inference.request (one per immutable
// ledger revision; a correction adds a new id with supersedes, keep the highest
// revision per request_id) and guardrail.evaluation (one per guardrail event,
// including blocked prompts with no inference record; correlate on request_id).
// Delivery is at least once in commit order: deduplicate on id. Records appear
// after settlement and after every older database transaction has finished
// (meta.horizon_lag_seconds); keep polling with next_cursor, which is always
// returned and is reusable after has_more is false. After a logical database
// restore, restart from start_time and deduplicate on id; a cursor from before it
// may return 409 cursor_expired (no Retry-After). Retention promise 30 days. No
// prompts, responses, provider credentials, supplier or customer rates. Gateway
// reference/enforcement usage, not invoice truth or BYOK provider charges.
func (r *LlmTokenGatewayUsageService) GetExport(ctx context.Context, query LlmTokenGatewayUsageGetExportParams, opts ...option.RequestOption) (res *LlmTokenGatewayUsageGetExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/usage/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Gateway reference/enforcement usage, not invoice truth or BYOK provider charges.
// Requires the dedicated llm_token_gateway.usage.read V2 action for combined spend
// and guardrail reporting; existing spend.read and guardrail_events.read grants do
// not grant this endpoint. All totals, UTC daily buckets, models and latest 20
// guardrail events share one database snapshot. UTC half-open dates, at most 31
// days. Zero-filled days; at most 1000 models (503 rather than truncation on
// overflow/timeout). Known tokens sum independently known categories across latest
// attempt evidence, including corrections; unknown spend is excluded and
// unknown_requests/reserved_spend remain explicit. Guardrail counts count events,
// not distinct requests. No raw spend events, prompts or secrets. This read-only
// report creates no pagination snapshot and accepts no pagination parameters.
func (r *LlmTokenGatewayUsageService) GetSummary(ctx context.Context, query LlmTokenGatewayUsageGetSummaryParams, opts ...option.RequestOption) (res *LlmTokenGatewayUsageGetSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/usage/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GuardrailFinding struct {
	// Any of "flag", "block".
	Action GuardrailFindingAction `json:"action" api:"required"`
	Code   string                 `json:"code" api:"required"`
	Count  int64                  `json:"count" api:"required"`
	// Any of "secrets", "dlp", "safety".
	Detector GuardrailFindingDetector `json:"detector" api:"required"`
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
func (r GuardrailFinding) RawJSON() string { return r.JSON.raw }
func (r *GuardrailFinding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuardrailFindingAction string

const (
	GuardrailFindingActionFlag  GuardrailFindingAction = "flag"
	GuardrailFindingActionBlock GuardrailFindingAction = "block"
)

type GuardrailFindingDetector string

const (
	GuardrailFindingDetectorSecrets GuardrailFindingDetector = "secrets"
	GuardrailFindingDetectorDlp     GuardrailFindingDetector = "dlp"
	GuardrailFindingDetectorSafety  GuardrailFindingDetector = "safety"
)

type LlmTokenGatewayUsageGetExportResponse struct {
	Data []LlmTokenGatewayUsageGetExportResponseDataUnion `json:"data" api:"required"`
	Meta LlmTokenGatewayUsageGetExportResponseMeta        `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetExportResponse) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LlmTokenGatewayUsageGetExportResponseDataUnion contains all possible properties
// and values from [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest],
// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
//
// Use the [LlmTokenGatewayUsageGetExportResponseDataUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LlmTokenGatewayUsageGetExportResponseDataUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Attempts int64 `json:"attempts"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Biller string `json:"biller"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	CacheHit bool `json:"cache_hit"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	ConfigurationVersion int64 `json:"configuration_version"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Cost      string `json:"cost"`
	EndUserID string `json:"end_user_id"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	InputTokens int64 `json:"input_tokens"`
	// This field is a union of
	// [map[string]LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion],
	// [map[string]LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion]
	Metadata LlmTokenGatewayUsageGetExportResponseDataUnionMetadata `json:"metadata"`
	Model    string                                                 `json:"model"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	OutputTokens int64 `json:"output_tokens"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Protocol string `json:"protocol"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Provider string `json:"provider"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	ProviderKeyID string `json:"provider_key_id"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	ProviderModel string `json:"provider_model"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	RateVersion string `json:"rate_version"`
	RequestID   string `json:"request_id"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Revision int64 `json:"revision"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	RouteMode     string  `json:"route_mode"`
	SchemaVersion float64 `json:"schema_version"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Status string `json:"status"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Stream bool `json:"stream"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Supersedes string `json:"supersedes"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	Timings      LlmTokenGatewayUsageGetExportResponseDataInferenceRequestTimings `json:"timings"`
	TokenGroupID string                                                           `json:"token_group_id"`
	TokenKeyID   string                                                           `json:"token_key_id"`
	TokenUserID  string                                                           `json:"token_user_id"`
	// Any of "inference.request", "guardrail.evaluation".
	Type string `json:"type"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	UsageSource string `json:"usage_source"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataInferenceRequest].
	UsageStatus string `json:"usage_status"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
	EvaluationInputTokens int64 `json:"evaluation_input_tokens"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
	EvaluationOutputTokens int64 `json:"evaluation_output_tokens"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
	Findings []GuardrailFinding `json:"findings"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
	Outcome string `json:"outcome"`
	// This field is from variant
	// [LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation].
	Stage string `json:"stage"`
	JSON  struct {
		ID                     respjson.Field
		Attempts               respjson.Field
		Biller                 respjson.Field
		CacheHit               respjson.Field
		ConfigurationVersion   respjson.Field
		Cost                   respjson.Field
		EndUserID              respjson.Field
		InputTokens            respjson.Field
		Metadata               respjson.Field
		Model                  respjson.Field
		OutputTokens           respjson.Field
		Protocol               respjson.Field
		Provider               respjson.Field
		ProviderKeyID          respjson.Field
		ProviderModel          respjson.Field
		RateVersion            respjson.Field
		RequestID              respjson.Field
		Revision               respjson.Field
		RouteMode              respjson.Field
		SchemaVersion          respjson.Field
		Status                 respjson.Field
		Stream                 respjson.Field
		Supersedes             respjson.Field
		Timings                respjson.Field
		TokenGroupID           respjson.Field
		TokenKeyID             respjson.Field
		TokenUserID            respjson.Field
		Type                   respjson.Field
		UsageSource            respjson.Field
		UsageStatus            respjson.Field
		CreatedAt              respjson.Field
		EvaluationInputTokens  respjson.Field
		EvaluationOutputTokens respjson.Field
		Findings               respjson.Field
		Outcome                respjson.Field
		Stage                  respjson.Field
		raw                    string
	} `json:"-"`
}

// anyLlmTokenGatewayUsageGetExportResponseData is implemented by each variant of
// [LlmTokenGatewayUsageGetExportResponseDataUnion] to add type safety for the
// return type of [LlmTokenGatewayUsageGetExportResponseDataUnion.AsAny]
type anyLlmTokenGatewayUsageGetExportResponseData interface {
	implLlmTokenGatewayUsageGetExportResponseDataUnion()
}

func (LlmTokenGatewayUsageGetExportResponseDataInferenceRequest) implLlmTokenGatewayUsageGetExportResponseDataUnion() {
}
func (LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation) implLlmTokenGatewayUsageGetExportResponseDataUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := LlmTokenGatewayUsageGetExportResponseDataUnion.AsAny().(type) {
//	case telnyx.LlmTokenGatewayUsageGetExportResponseDataInferenceRequest:
//	case telnyx.LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LlmTokenGatewayUsageGetExportResponseDataUnion) AsAny() anyLlmTokenGatewayUsageGetExportResponseData {
	switch u.Type {
	case "inference.request":
		return u.AsInferenceRequest()
	case "guardrail.evaluation":
		return u.AsGuardrailEvaluation()
	}
	return nil
}

func (u LlmTokenGatewayUsageGetExportResponseDataUnion) AsInferenceRequest() (v LlmTokenGatewayUsageGetExportResponseDataInferenceRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetExportResponseDataUnion) AsGuardrailEvaluation() (v LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LlmTokenGatewayUsageGetExportResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *LlmTokenGatewayUsageGetExportResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LlmTokenGatewayUsageGetExportResponseDataUnionMetadata is an implicit subunion
// of [LlmTokenGatewayUsageGetExportResponseDataUnion].
// LlmTokenGatewayUsageGetExportResponseDataUnionMetadata provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LlmTokenGatewayUsageGetExportResponseDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type LlmTokenGatewayUsageGetExportResponseDataUnionMetadata struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (r *LlmTokenGatewayUsageGetExportResponseDataUnionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetExportResponseDataInferenceRequest struct {
	// Deterministic: req\_{request_id}\_r{revision}. Deduplicate on id.
	ID       string `json:"id" api:"required"`
	Attempts int64  `json:"attempts" api:"required"`
	// Any of "telnyx", "provider", "none".
	Biller               string `json:"biller" api:"required"`
	CacheHit             bool   `json:"cache_hit" api:"required"`
	ConfigurationVersion int64  `json:"configuration_version" api:"required"`
	// USD decimal string with exactly 6 places, 0 to 1000000000.000000 (the Money
	// bound); enforcement basis, not invoice truth.
	Cost        string `json:"cost" api:"required"`
	EndUserID   string `json:"end_user_id" api:"required"`
	InputTokens int64  `json:"input_tokens" api:"required"`
	// Customer request tags captured at admission; {} when untagged.
	Metadata     map[string]LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion `json:"metadata" api:"required"`
	Model        string                                                                            `json:"model" api:"required"`
	OutputTokens int64                                                                             `json:"output_tokens" api:"required"`
	// Any of "openai", "anthropic".
	Protocol      string `json:"protocol" api:"required"`
	Provider      string `json:"provider" api:"required"`
	ProviderKeyID string `json:"provider_key_id" api:"required" format:"uuid"`
	ProviderModel string `json:"provider_model" api:"required"`
	RateVersion   string `json:"rate_version" api:"required"`
	RequestID     string `json:"request_id" api:"required" format:"uuid"`
	Revision      int64  `json:"revision" api:"required"`
	// Any of "managed", "byok", "cache".
	RouteMode     string  `json:"route_mode" api:"required"`
	SchemaVersion float64 `json:"schema_version" api:"required"`
	// Any of "succeeded", "failed", "partial", "unknown".
	Status string `json:"status" api:"required"`
	Stream bool   `json:"stream" api:"required"`
	// Previous revision id; keep the highest revision per request_id.
	Supersedes   string                                                           `json:"supersedes" api:"required"`
	Timings      LlmTokenGatewayUsageGetExportResponseDataInferenceRequestTimings `json:"timings" api:"required"`
	TokenGroupID string                                                           `json:"token_group_id" api:"required" format:"uuid"`
	TokenKeyID   string                                                           `json:"token_key_id" api:"required" format:"uuid"`
	TokenUserID  string                                                           `json:"token_user_id" api:"required" format:"uuid"`
	Type         constant.InferenceRequest                                        `json:"type" default:"inference.request"`
	// Any of "provider", "gateway_tokenizer", "operator".
	UsageSource string `json:"usage_source" api:"required"`
	// Any of "known", "unknown", "reconciled".
	UsageStatus string `json:"usage_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Attempts             respjson.Field
		Biller               respjson.Field
		CacheHit             respjson.Field
		ConfigurationVersion respjson.Field
		Cost                 respjson.Field
		EndUserID            respjson.Field
		InputTokens          respjson.Field
		Metadata             respjson.Field
		Model                respjson.Field
		OutputTokens         respjson.Field
		Protocol             respjson.Field
		Provider             respjson.Field
		ProviderKeyID        respjson.Field
		ProviderModel        respjson.Field
		RateVersion          respjson.Field
		RequestID            respjson.Field
		Revision             respjson.Field
		RouteMode            respjson.Field
		SchemaVersion        respjson.Field
		Status               respjson.Field
		Stream               respjson.Field
		Supersedes           respjson.Field
		Timings              respjson.Field
		TokenGroupID         respjson.Field
		TokenKeyID           respjson.Field
		TokenUserID          respjson.Field
		Type                 respjson.Field
		UsageSource          respjson.Field
		UsageStatus          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetExportResponseDataInferenceRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *LlmTokenGatewayUsageGetExportResponseDataInferenceRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion contains
// all possible properties and values from [string], [int64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *LlmTokenGatewayUsageGetExportResponseDataInferenceRequestMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetExportResponseDataInferenceRequestTimings struct {
	AdmittedAt     time.Time `json:"admitted_at" api:"required" format:"date-time"`
	CompletedAt    time.Time `json:"completed_at" api:"required" format:"date-time"`
	FirstAttemptAt time.Time `json:"first_attempt_at" api:"required" format:"date-time"`
	// completed_at - admitted_at as recorded by the ledger; not time to first token.
	GatewayDurationMs int64     `json:"gateway_duration_ms" api:"required"`
	SettledAt         time.Time `json:"settled_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdmittedAt        respjson.Field
		CompletedAt       respjson.Field
		FirstAttemptAt    respjson.Field
		GatewayDurationMs respjson.Field
		SettledAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetExportResponseDataInferenceRequestTimings) RawJSON() string {
	return r.JSON.raw
}
func (r *LlmTokenGatewayUsageGetExportResponseDataInferenceRequestTimings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation struct {
	ID                     string             `json:"id" api:"required"`
	CreatedAt              time.Time          `json:"created_at" api:"required" format:"date-time"`
	EndUserID              string             `json:"end_user_id" api:"required"`
	EvaluationInputTokens  int64              `json:"evaluation_input_tokens" api:"required"`
	EvaluationOutputTokens int64              `json:"evaluation_output_tokens" api:"required"`
	Findings               []GuardrailFinding `json:"findings" api:"required"`
	// Customer request tags from the x-ltg-metadata header: up to 5 entries. Keys are
	// case-sensitive; keys starting with ltg*, telnyx*, litellm* or cf* (any case) are
	// reserved. Values are strings (1-256 characters, no control or format
	// characters), integers within +/-9007199254740991, or booleans; floats, null,
	// objects and arrays are rejected. Values that look like credentials are rejected.
	// Tags are attribution only: never forwarded upstream, never part of the response
	// cache key, never used for budgets or routing. An untagged request reports {}.
	Metadata map[string]LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion `json:"metadata" api:"required"`
	Model    string                                                                               `json:"model" api:"required"`
	// Any of "evaluated", "flagged", "blocked", "unevaluated".
	Outcome       string  `json:"outcome" api:"required"`
	RequestID     string  `json:"request_id" api:"required" format:"uuid"`
	SchemaVersion float64 `json:"schema_version" api:"required"`
	// Any of "prompt", "response".
	Stage        string                       `json:"stage" api:"required"`
	TokenGroupID string                       `json:"token_group_id" api:"required" format:"uuid"`
	TokenKeyID   string                       `json:"token_key_id" api:"required" format:"uuid"`
	TokenUserID  string                       `json:"token_user_id" api:"required" format:"uuid"`
	Type         constant.GuardrailEvaluation `json:"type" default:"guardrail.evaluation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		CreatedAt              respjson.Field
		EndUserID              respjson.Field
		EvaluationInputTokens  respjson.Field
		EvaluationOutputTokens respjson.Field
		Findings               respjson.Field
		Metadata               respjson.Field
		Model                  respjson.Field
		Outcome                respjson.Field
		RequestID              respjson.Field
		SchemaVersion          respjson.Field
		Stage                  respjson.Field
		TokenGroupID           respjson.Field
		TokenKeyID             respjson.Field
		TokenUserID            respjson.Field
		Type                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation) RawJSON() string {
	return r.JSON.raw
}
func (r *LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion
// contains all possible properties and values from [string], [int64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *LlmTokenGatewayUsageGetExportResponseDataGuardrailEvaluationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetExportResponseMeta struct {
	HasMore           bool   `json:"has_more" api:"required"`
	HorizonLagSeconds int64  `json:"horizon_lag_seconds" api:"required"`
	NextCursor        string `json:"next_cursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore           respjson.Field
		HorizonLagSeconds respjson.Field
		NextCursor        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmTokenGatewayUsageGetExportResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *LlmTokenGatewayUsageGetExportResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ByDay      []LlmTokenGatewayUsageGetSummaryResponseDataByDay    `json:"by_day" api:"required"`
	ByModel    []LlmTokenGatewayUsageGetSummaryResponseDataByModel  `json:"by_model" api:"required"`
	Guardrails LlmTokenGatewayUsageGetSummaryResponseDataGuardrails `json:"guardrails" api:"required"`
	Totals     LlmTokenGatewayUsageGetSummaryResponseDataTotals     `json:"totals" api:"required"`
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
	CacheHits       int64     `json:"cache_hits" api:"required"`
	Date            time.Time `json:"date" api:"required" format:"date"`
	FailedRequests  int64     `json:"failed_requests" api:"required"`
	InputTokens     int64     `json:"input_tokens" api:"required"`
	OutputTokens    int64     `json:"output_tokens" api:"required"`
	PartialRequests int64     `json:"partial_requests" api:"required"`
	Requests        int64     `json:"requests" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Spend             float64 `json:"spend" api:"required"`
	SucceededRequests int64   `json:"succeeded_requests" api:"required"`
	UnknownRequests   int64   `json:"unknown_requests" api:"required"`
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
	CacheHits       int64  `json:"cache_hits" api:"required"`
	FailedRequests  int64  `json:"failed_requests" api:"required"`
	InputTokens     int64  `json:"input_tokens" api:"required"`
	Model           string `json:"model" api:"required"`
	OutputTokens    int64  `json:"output_tokens" api:"required"`
	PartialRequests int64  `json:"partial_requests" api:"required"`
	Requests        int64  `json:"requests" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Spend             float64 `json:"spend" api:"required"`
	SucceededRequests int64   `json:"succeeded_requests" api:"required"`
	UnknownRequests   int64   `json:"unknown_requests" api:"required"`
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

type LlmTokenGatewayUsageGetSummaryResponseDataGuardrails struct {
	BlockedEvents int64                                                             `json:"blocked_events" api:"required"`
	FlaggedEvents int64                                                             `json:"flagged_events" api:"required"`
	RecentEvents  []LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEvent `json:"recent_events" api:"required"`
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
	ID                     string             `json:"id" api:"required" format:"uuid"`
	CreatedAt              time.Time          `json:"created_at" api:"required" format:"date-time"`
	EndUserID              string             `json:"end_user_id" api:"required"`
	EvaluationInputTokens  int64              `json:"evaluation_input_tokens" api:"required"`
	EvaluationOutputTokens int64              `json:"evaluation_output_tokens" api:"required"`
	Findings               []GuardrailFinding `json:"findings" api:"required"`
	// Customer request tags from the x-ltg-metadata header: up to 5 entries. Keys are
	// case-sensitive; keys starting with ltg*, telnyx*, litellm* or cf* (any case) are
	// reserved. Values are strings (1-256 characters, no control or format
	// characters), integers within +/-9007199254740991, or booleans; floats, null,
	// objects and arrays are rejected. Values that look like credentials are rejected.
	// Tags are attribution only: never forwarded upstream, never part of the response
	// cache key, never used for budgets or routing. An untagged request reports {}.
	Metadata map[string]LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion `json:"metadata" api:"required"`
	Model    string                                                                                   `json:"model" api:"required"`
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
		Metadata               respjson.Field
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

// LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion
// contains all possible properties and values from [string], [int64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt OfBool]
type LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfInt    respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *LlmTokenGatewayUsageGetSummaryResponseDataGuardrailsRecentEventsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayUsageGetSummaryResponseDataTotals struct {
	CacheHits       int64 `json:"cache_hits" api:"required"`
	FailedRequests  int64 `json:"failed_requests" api:"required"`
	InputTokens     int64 `json:"input_tokens" api:"required"`
	OutputTokens    int64 `json:"output_tokens" api:"required"`
	PartialRequests int64 `json:"partial_requests" api:"required"`
	Requests        int64 `json:"requests" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	ReservedSpend float64 `json:"reserved_spend" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Spend             float64 `json:"spend" api:"required"`
	SucceededRequests int64   `json:"succeeded_requests" api:"required"`
	UnknownRequests   int64   `json:"unknown_requests" api:"required"`
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
	EndDate time.Time `json:"end_date" api:"required" format:"date"`
	// The metadata[<key>] filters applied to every section; {} when unfiltered.
	Metadata     map[string]string `json:"metadata" api:"required"`
	StartDate    time.Time         `json:"start_date" api:"required" format:"date"`
	TokenGroupID string            `json:"token_group_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate      respjson.Field
		Metadata     respjson.Field
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

type LlmTokenGatewayUsageGetExportParams struct {
	// Opaque next_cursor from a previous page; not combinable with start_time. Bound
	// to the requested types; the account always comes from the credential.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// First page only: RFC 3339 with offset, at most 30 days ago and not in the
	// future; default now - 24h.
	StartTime param.Opt[time.Time] `query:"start_time,omitzero" format:"date-time" json:"-"`
	// Comma-separated subset of inference.request,guardrail.evaluation; default both.
	Types param.Opt[string] `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LlmTokenGatewayUsageGetExportParams]'s query parameters as
// `url.Values`.
func (r LlmTokenGatewayUsageGetExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LlmTokenGatewayUsageGetSummaryParams struct {
	EndDate      time.Time `query:"end_date" api:"required" format:"date" json:"-"`
	StartDate    time.Time `query:"start_date" api:"required" format:"date" json:"-"`
	TokenGroupID string    `query:"token_group_id" api:"required" format:"uuid" json:"-"`
	// Request tag filters as metadata[<key>]=<value>, at most 5, combined with AND. A
	// value matches the tag's text form exactly (booleans as true/false, integers in
	// decimal), so metadata[flag]=true matches both boolean true and the string
	// "true". Keys are case-sensitive; reserved-prefix keys are rejected. Unindexed:
	// bounded by the account, the date range and the report timeout.
	Metadata map[string]string `query:"metadata,omitzero" json:"-"`
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
