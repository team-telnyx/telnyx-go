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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Manage and report AI Gateway usage.
//
// LlmTokenGatewayTokenGroupService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmTokenGatewayTokenGroupService] method instead.
type LlmTokenGatewayTokenGroupService struct {
	Options []option.RequestOption
}

// NewLlmTokenGatewayTokenGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLlmTokenGatewayTokenGroupService(opts ...option.RequestOption) (r LlmTokenGatewayTokenGroupService) {
	r = LlmTokenGatewayTokenGroupService{}
	r.Options = opts
	return
}

// Create gateway token group within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenGroupService) New(ctx context.Context, params LlmTokenGatewayTokenGroupNewParams, opts ...option.RequestOption) (res *TokenGroupResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/token_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get gateway token group within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenGroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TokenGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("llm_token_gateway/token_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update gateway token group within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response. Omitted fields are preserved.
// Null clears the optional concurrency, fallback and retry settings. A stale
// If-Match returns 412; read the resource again before updating.
func (r *LlmTokenGatewayTokenGroupService) Update(ctx context.Context, id string, params LlmTokenGatewayTokenGroupUpdateParams, opts ...option.RequestOption) (res *TokenGroupResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	if !param.IsOmitted(params.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", params.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("llm_token_gateway/token_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List gateway token groups within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenGroupService) List(ctx context.Context, query LlmTokenGatewayTokenGroupListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[TokenGroup], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "llm_token_gateway/token_groups"
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

// List gateway token groups within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenGroupService) ListAutoPaging(ctx context.Context, query LlmTokenGatewayTokenGroupListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[TokenGroup] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Revoke resource and its keys atomically; retain historical attribution, never
// detach into an ungoverned service key.
func (r *LlmTokenGatewayTokenGroupService) Delete(ctx context.Context, id string, body LlmTokenGatewayTokenGroupDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey)))
	}
	if !param.IsOmitted(body.IfMatch) {
		opts = append(opts, option.WithHeader("If-Match", fmt.Sprintf("%v", body.IfMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("llm_token_gateway/token_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type AIGatewayResourcesPageMeta struct {
	HasMore    bool   `json:"has_more" api:"required"`
	PageNumber int64  `json:"page_number" api:"required"`
	PageSize   int64  `json:"page_size" api:"required"`
	Snapshot   string `json:"snapshot" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		PageNumber  respjson.Field
		PageSize    respjson.Field
		Snapshot    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIGatewayResourcesPageMeta) RawJSON() string { return r.JSON.raw }
func (r *AIGatewayResourcesPageMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetDuration string

const (
	BudgetDuration1d  BudgetDuration = "1d"
	BudgetDuration7d  BudgetDuration = "7d"
	BudgetDuration30d BudgetDuration = "30d"
)

type TokenGroup struct {
	ID            string   `json:"id" api:"required" format:"uuid"`
	AllowedModels []string `json:"allowed_models" api:"required"`
	Blocked       bool     `json:"blocked" api:"required"`
	// Any of "1d", "7d", "30d".
	BudgetDuration  BudgetDuration `json:"budget_duration" api:"required"`
	BudgetStartedAt time.Time      `json:"budget_started_at" api:"required" format:"date-time"`
	CacheTtl        int64          `json:"cache_ttl" api:"required"`
	CreatedAt       time.Time      `json:"created_at" api:"required" format:"date-time"`
	// Per-group model fallbacks: primary alias -> ordered list of 1..3 other aliases.
	// Every key and listed alias must be in allowed_models (a PATCH removing a
	// referenced alias is 400); a chain never lists its own key; wildcards are
	// rejected. Chains are one level deep (a step's own chain is not expanded).
	// Fallback is attempted only for managed routes, only before the first provider
	// event is forwarded, and only to steps the key may use; each attempt is reserved
	// and recorded. {} and null both disable fallbacks.
	Fallbacks map[string][]string `json:"fallbacks" api:"required"`
	// Per-group guardrail policy. Missing actions default to ignore. profiles may be
	// listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp
	// action requires at least one profile. streaming defaults to buffered when any
	// response action is block and passthrough otherwise; passthrough with a response
	// block is 400. Findings are codes and counts only; content is never stored. The
	// stored form is echoed back canonicalised (all actions explicit, profiles
	// sorted); {} is stored as an all-ignore policy, null disables.
	Guardrails TokenGroupGuardrails `json:"guardrails" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	MaxBudget      float64             `json:"max_budget" api:"required"`
	Name           string              `json:"name" api:"required"`
	ProviderKeyIDs []string            `json:"provider_key_ids" api:"required" format:"uuid"`
	RecordType     constant.TokenGroup `json:"record_type" default:"token_group"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	ReservedSpend float64   `json:"reserved_spend" api:"required"`
	ResetsAt      time.Time `json:"resets_at" api:"required" format:"date-time"`
	// Per-group same-route retries. max_attempts counts the first try of each step
	// (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts
	// request header. A step is retried only when its attempt definitely did no
	// provider work (connection failure before dispatch, or 429/529) before the first
	// provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the
	// next fallback step instead. max_attempts x (1 + the longest fallback chain) must
	// not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant),
	// retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s
	// and never past the request deadline, then jittered uniformly down to half that
	// value so concurrent retries spread out; a provider Retry-After of at most 5 s is
	// honoured and a longer one moves to the next step. request_timeout_ms applies to
	// streaming requests only: it bounds each attempt's time to its first provider
	// event and never applies after it. Non-stream requests keep the per-attempt hop
	// and idle bounds (their headers arrive only with the finished completion).
	// retry_delay_ms and backoff are floors for the request knobs. Each planned
	// attempt is reserved (at most 8 per request); unused reservation is released at
	// settlement. Managed routes only. The stored form is echoed back with defaults
	// filled in; null disables retries.
	Retries  TokenGroupRetries `json:"retries" api:"required"`
	RpmLimit int64             `json:"rpm_limit" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Spend     float64   `json:"spend" api:"required"`
	TpmLimit  int64     `json:"tpm_limit" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	Version   int64     `json:"version" api:"required"`
	// Optional maximum concurrent reserved/dispatched cache misses across replicas.
	// Null means no concurrency cap. Cache hits do not consume slots; unknown usage
	// retains its budget hold without a slot.
	MaxParallelRequests int64 `json:"max_parallel_requests" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AllowedModels       respjson.Field
		Blocked             respjson.Field
		BudgetDuration      respjson.Field
		BudgetStartedAt     respjson.Field
		CacheTtl            respjson.Field
		CreatedAt           respjson.Field
		Fallbacks           respjson.Field
		Guardrails          respjson.Field
		MaxBudget           respjson.Field
		Name                respjson.Field
		ProviderKeyIDs      respjson.Field
		RecordType          respjson.Field
		ReservedSpend       respjson.Field
		ResetsAt            respjson.Field
		Retries             respjson.Field
		RpmLimit            respjson.Field
		Spend               respjson.Field
		TpmLimit            respjson.Field
		UpdatedAt           respjson.Field
		Version             respjson.Field
		MaxParallelRequests respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGroup) RawJSON() string { return r.JSON.raw }
func (r *TokenGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-group guardrail policy. Missing actions default to ignore. profiles may be
// listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp
// action requires at least one profile. streaming defaults to buffered when any
// response action is block and passthrough otherwise; passthrough with a response
// block is 400. Findings are codes and counts only; content is never stored. The
// stored form is echoed back canonicalised (all actions explicit, profiles
// sorted); {} is stored as an all-ignore policy, null disables.
type TokenGroupGuardrails struct {
	Dlp     TokenGroupGuardrailsDlp     `json:"dlp"`
	Secrets TokenGroupGuardrailsSecrets `json:"secrets"`
	// Any of "buffered", "passthrough".
	Streaming string `json:"streaming"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dlp         respjson.Field
		Secrets     respjson.Field
		Streaming   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGroupGuardrails) RawJSON() string { return r.JSON.raw }
func (r *TokenGroupGuardrails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGroupGuardrailsDlp struct {
	// Any of "financial", "government_id", "contact".
	Profiles []string `json:"profiles"`
	// Any of "ignore", "flag", "block".
	Prompt string `json:"prompt"`
	// Any of "ignore", "flag", "block".
	Response string `json:"response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Profiles    respjson.Field
		Prompt      respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGroupGuardrailsDlp) RawJSON() string { return r.JSON.raw }
func (r *TokenGroupGuardrailsDlp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGroupGuardrailsSecrets struct {
	// Any of "ignore", "flag", "block".
	Prompt string `json:"prompt"`
	// Any of "ignore", "flag", "block".
	Response string `json:"response"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompt      respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGroupGuardrailsSecrets) RawJSON() string { return r.JSON.raw }
func (r *TokenGroupGuardrailsSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-group same-route retries. max_attempts counts the first try of each step
// (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts
// request header. A step is retried only when its attempt definitely did no
// provider work (connection failure before dispatch, or 429/529) before the first
// provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the
// next fallback step instead. max_attempts x (1 + the longest fallback chain) must
// not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant),
// retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s
// and never past the request deadline, then jittered uniformly down to half that
// value so concurrent retries spread out; a provider Retry-After of at most 5 s is
// honoured and a longer one moves to the next step. request_timeout_ms applies to
// streaming requests only: it bounds each attempt's time to its first provider
// event and never applies after it. Non-stream requests keep the per-attempt hop
// and idle bounds (their headers arrive only with the finished completion).
// retry_delay_ms and backoff are floors for the request knobs. Each planned
// attempt is reserved (at most 8 per request); unused reservation is released at
// settlement. Managed routes only. The stored form is echoed back with defaults
// filled in; null disables retries.
type TokenGroupRetries struct {
	MaxAttempts int64 `json:"max_attempts" api:"required"`
	// Any of "constant", "linear", "exponential".
	Backoff          string `json:"backoff"`
	RequestTimeoutMs int64  `json:"request_timeout_ms" api:"nullable"`
	RetryDelayMs     int64  `json:"retry_delay_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxAttempts      respjson.Field
		Backoff          respjson.Field
		RequestTimeoutMs respjson.Field
		RetryDelayMs     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGroupRetries) RawJSON() string { return r.JSON.raw }
func (r *TokenGroupRetries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGroupResponse struct {
	Data TokenGroup `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayTokenGroupNewParams struct {
	AllowedModels  []string         `json:"allowed_models,omitzero" api:"required"`
	Name           string           `json:"name" api:"required"`
	IdempotencyKey string           `header:"Idempotency-Key" api:"required" json:"-"`
	CacheTtl       param.Opt[int64] `json:"cache_ttl,omitzero"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// Optional maximum concurrent reserved/dispatched cache misses across replicas.
	// Null means no concurrency cap. Cache hits do not consume slots; unknown usage
	// retains its budget hold without a slot.
	MaxParallelRequests param.Opt[int64] `json:"max_parallel_requests,omitzero"`
	RpmLimit            param.Opt[int64] `json:"rpm_limit,omitzero"`
	TpmLimit            param.Opt[int64] `json:"tpm_limit,omitzero"`
	Blocked             param.Opt[bool]  `json:"blocked,omitzero"`
	// Any of "1d", "7d", "30d".
	BudgetDuration BudgetDuration `json:"budget_duration,omitzero"`
	// Per-group model fallbacks: primary alias -> ordered list of 1..3 other aliases.
	// Every key and listed alias must be in allowed_models (a PATCH removing a
	// referenced alias is 400); a chain never lists its own key; wildcards are
	// rejected. Chains are one level deep (a step's own chain is not expanded).
	// Fallback is attempted only for managed routes, only before the first provider
	// event is forwarded, and only to steps the key may use; each attempt is reserved
	// and recorded. {} and null both disable fallbacks.
	Fallbacks map[string][]string `json:"fallbacks,omitzero"`
	// Per-group guardrail policy. Missing actions default to ignore. profiles may be
	// listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp
	// action requires at least one profile. streaming defaults to buffered when any
	// response action is block and passthrough otherwise; passthrough with a response
	// block is 400. Findings are codes and counts only; content is never stored. The
	// stored form is echoed back canonicalised (all actions explicit, profiles
	// sorted); {} is stored as an all-ignore policy, null disables.
	Guardrails LlmTokenGatewayTokenGroupNewParamsGuardrails `json:"guardrails,omitzero"`
	// Per-group same-route retries. max_attempts counts the first try of each step
	// (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts
	// request header. A step is retried only when its attempt definitely did no
	// provider work (connection failure before dispatch, or 429/529) before the first
	// provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the
	// next fallback step instead. max_attempts x (1 + the longest fallback chain) must
	// not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant),
	// retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s
	// and never past the request deadline, then jittered uniformly down to half that
	// value so concurrent retries spread out; a provider Retry-After of at most 5 s is
	// honoured and a longer one moves to the next step. request_timeout_ms applies to
	// streaming requests only: it bounds each attempt's time to its first provider
	// event and never applies after it. Non-stream requests keep the per-attempt hop
	// and idle bounds (their headers arrive only with the finished completion).
	// retry_delay_ms and backoff are floors for the request knobs. Each planned
	// attempt is reserved (at most 8 per request); unused reservation is released at
	// settlement. Managed routes only. The stored form is echoed back with defaults
	// filled in; null disables retries.
	Retries        LlmTokenGatewayTokenGroupNewParamsRetries `json:"retries,omitzero"`
	ProviderKeyIDs []string                                  `json:"provider_key_ids,omitzero" format:"uuid"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-group guardrail policy. Missing actions default to ignore. profiles may be
// listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp
// action requires at least one profile. streaming defaults to buffered when any
// response action is block and passthrough otherwise; passthrough with a response
// block is 400. Findings are codes and counts only; content is never stored. The
// stored form is echoed back canonicalised (all actions explicit, profiles
// sorted); {} is stored as an all-ignore policy, null disables.
type LlmTokenGatewayTokenGroupNewParamsGuardrails struct {
	Dlp     LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp     `json:"dlp,omitzero"`
	Secrets LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets `json:"secrets,omitzero"`
	// Any of "buffered", "passthrough".
	Streaming string `json:"streaming,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupNewParamsGuardrails) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupNewParamsGuardrails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupNewParamsGuardrails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupNewParamsGuardrails](
		"streaming", "buffered", "passthrough",
	)
}

type LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp struct {
	// Any of "financial", "government_id", "contact".
	Profiles []string `json:"profiles,omitzero"`
	// Any of "ignore", "flag", "block".
	Prompt string `json:"prompt,omitzero"`
	// Any of "ignore", "flag", "block".
	Response string `json:"response,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp](
		"prompt", "ignore", "flag", "block",
	)
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupNewParamsGuardrailsDlp](
		"response", "ignore", "flag", "block",
	)
}

type LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets struct {
	// Any of "ignore", "flag", "block".
	Prompt string `json:"prompt,omitzero"`
	// Any of "ignore", "flag", "block".
	Response string `json:"response,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets](
		"prompt", "ignore", "flag", "block",
	)
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupNewParamsGuardrailsSecrets](
		"response", "ignore", "flag", "block",
	)
}

// Per-group same-route retries. max_attempts counts the first try of each step
// (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts
// request header. A step is retried only when its attempt definitely did no
// provider work (connection failure before dispatch, or 429/529) before the first
// provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the
// next fallback step instead. max_attempts x (1 + the longest fallback chain) must
// not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant),
// retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s
// and never past the request deadline, then jittered uniformly down to half that
// value so concurrent retries spread out; a provider Retry-After of at most 5 s is
// honoured and a longer one moves to the next step. request_timeout_ms applies to
// streaming requests only: it bounds each attempt's time to its first provider
// event and never applies after it. Non-stream requests keep the per-attempt hop
// and idle bounds (their headers arrive only with the finished completion).
// retry_delay_ms and backoff are floors for the request knobs. Each planned
// attempt is reserved (at most 8 per request); unused reservation is released at
// settlement. Managed routes only. The stored form is echoed back with defaults
// filled in; null disables retries.
//
// The property MaxAttempts is required.
type LlmTokenGatewayTokenGroupNewParamsRetries struct {
	MaxAttempts      int64            `json:"max_attempts" api:"required"`
	RequestTimeoutMs param.Opt[int64] `json:"request_timeout_ms,omitzero"`
	RetryDelayMs     param.Opt[int64] `json:"retry_delay_ms,omitzero"`
	// Any of "constant", "linear", "exponential".
	Backoff string `json:"backoff,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupNewParamsRetries) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupNewParamsRetries
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupNewParamsRetries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupNewParamsRetries](
		"backoff", "constant", "linear", "exponential",
	)
}

type LlmTokenGatewayTokenGroupUpdateParams struct {
	IdempotencyKey string           `header:"Idempotency-Key" api:"required" json:"-"`
	IfMatch        string           `header:"If-Match" api:"required" json:"-"`
	CacheTtl       param.Opt[int64] `json:"cache_ttl,omitzero"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// Optional maximum concurrent reserved/dispatched cache misses across replicas.
	// Null means no concurrency cap. Cache hits do not consume slots; unknown usage
	// retains its budget hold without a slot.
	MaxParallelRequests param.Opt[int64]  `json:"max_parallel_requests,omitzero"`
	RpmLimit            param.Opt[int64]  `json:"rpm_limit,omitzero"`
	TpmLimit            param.Opt[int64]  `json:"tpm_limit,omitzero"`
	Blocked             param.Opt[bool]   `json:"blocked,omitzero"`
	Name                param.Opt[string] `json:"name,omitzero"`
	// Any of "1d", "7d", "30d".
	BudgetDuration BudgetDuration `json:"budget_duration,omitzero"`
	// Per-group model fallbacks: primary alias -> ordered list of 1..3 other aliases.
	// Every key and listed alias must be in allowed_models (a PATCH removing a
	// referenced alias is 400); a chain never lists its own key; wildcards are
	// rejected. Chains are one level deep (a step's own chain is not expanded).
	// Fallback is attempted only for managed routes, only before the first provider
	// event is forwarded, and only to steps the key may use; each attempt is reserved
	// and recorded. {} and null both disable fallbacks.
	Fallbacks map[string][]string `json:"fallbacks,omitzero"`
	// Per-group guardrail policy. Missing actions default to ignore. profiles may be
	// listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp
	// action requires at least one profile. streaming defaults to buffered when any
	// response action is block and passthrough otherwise; passthrough with a response
	// block is 400. Findings are codes and counts only; content is never stored. The
	// stored form is echoed back canonicalised (all actions explicit, profiles
	// sorted); {} is stored as an all-ignore policy, null disables.
	Guardrails LlmTokenGatewayTokenGroupUpdateParamsGuardrails `json:"guardrails,omitzero"`
	// Per-group same-route retries. max_attempts counts the first try of each step
	// (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts
	// request header. A step is retried only when its attempt definitely did no
	// provider work (connection failure before dispatch, or 429/529) before the first
	// provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the
	// next fallback step instead. max_attempts x (1 + the longest fallback chain) must
	// not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant),
	// retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s
	// and never past the request deadline, then jittered uniformly down to half that
	// value so concurrent retries spread out; a provider Retry-After of at most 5 s is
	// honoured and a longer one moves to the next step. request_timeout_ms applies to
	// streaming requests only: it bounds each attempt's time to its first provider
	// event and never applies after it. Non-stream requests keep the per-attempt hop
	// and idle bounds (their headers arrive only with the finished completion).
	// retry_delay_ms and backoff are floors for the request knobs. Each planned
	// attempt is reserved (at most 8 per request); unused reservation is released at
	// settlement. Managed routes only. The stored form is echoed back with defaults
	// filled in; null disables retries.
	Retries        LlmTokenGatewayTokenGroupUpdateParamsRetries `json:"retries,omitzero"`
	AllowedModels  []string                                     `json:"allowed_models,omitzero"`
	ProviderKeyIDs []string                                     `json:"provider_key_ids,omitzero" format:"uuid"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-group guardrail policy. Missing actions default to ignore. profiles may be
// listed with both dlp actions ignore (staged, inactive) but a non-ignore dlp
// action requires at least one profile. streaming defaults to buffered when any
// response action is block and passthrough otherwise; passthrough with a response
// block is 400. Findings are codes and counts only; content is never stored. The
// stored form is echoed back canonicalised (all actions explicit, profiles
// sorted); {} is stored as an all-ignore policy, null disables.
type LlmTokenGatewayTokenGroupUpdateParamsGuardrails struct {
	Dlp     LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp     `json:"dlp,omitzero"`
	Secrets LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets `json:"secrets,omitzero"`
	// Any of "buffered", "passthrough".
	Streaming string `json:"streaming,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupUpdateParamsGuardrails) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupUpdateParamsGuardrails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupUpdateParamsGuardrails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupUpdateParamsGuardrails](
		"streaming", "buffered", "passthrough",
	)
}

type LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp struct {
	// Any of "financial", "government_id", "contact".
	Profiles []string `json:"profiles,omitzero"`
	// Any of "ignore", "flag", "block".
	Prompt string `json:"prompt,omitzero"`
	// Any of "ignore", "flag", "block".
	Response string `json:"response,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp](
		"prompt", "ignore", "flag", "block",
	)
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupUpdateParamsGuardrailsDlp](
		"response", "ignore", "flag", "block",
	)
}

type LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets struct {
	// Any of "ignore", "flag", "block".
	Prompt string `json:"prompt,omitzero"`
	// Any of "ignore", "flag", "block".
	Response string `json:"response,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets](
		"prompt", "ignore", "flag", "block",
	)
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupUpdateParamsGuardrailsSecrets](
		"response", "ignore", "flag", "block",
	)
}

// Per-group same-route retries. max_attempts counts the first try of each step
// (the primary and each fallback) and is the ceiling for the x-ltg-max-attempts
// request header. A step is retried only when its attempt definitely did no
// provider work (connection failure before dispatch, or 429/529) before the first
// provider event; an uncertain failure (timeout, 408/500/502/503/504) moves to the
// next fallback step instead. max_attempts x (1 + the longest fallback chain) must
// not exceed 8, else 400. The n-th retry waits retry_delay_ms (constant),
// retry_delay_ms*n (linear) or retry_delay_ms*2^(n-1) (exponential), at most 5 s
// and never past the request deadline, then jittered uniformly down to half that
// value so concurrent retries spread out; a provider Retry-After of at most 5 s is
// honoured and a longer one moves to the next step. request_timeout_ms applies to
// streaming requests only: it bounds each attempt's time to its first provider
// event and never applies after it. Non-stream requests keep the per-attempt hop
// and idle bounds (their headers arrive only with the finished completion).
// retry_delay_ms and backoff are floors for the request knobs. Each planned
// attempt is reserved (at most 8 per request); unused reservation is released at
// settlement. Managed routes only. The stored form is echoed back with defaults
// filled in; null disables retries.
//
// The property MaxAttempts is required.
type LlmTokenGatewayTokenGroupUpdateParamsRetries struct {
	MaxAttempts      int64            `json:"max_attempts" api:"required"`
	RequestTimeoutMs param.Opt[int64] `json:"request_timeout_ms,omitzero"`
	RetryDelayMs     param.Opt[int64] `json:"retry_delay_ms,omitzero"`
	// Any of "constant", "linear", "exponential".
	Backoff string `json:"backoff,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenGroupUpdateParamsRetries) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenGroupUpdateParamsRetries
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenGroupUpdateParamsRetries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[LlmTokenGatewayTokenGroupUpdateParamsRetries](
		"backoff", "constant", "linear", "exponential",
	)
}

type LlmTokenGatewayTokenGroupListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Required after page 1; account/filter-bound 15-minute stable snapshot;
	// expired/mismatch 409.
	PageSnapshot param.Opt[string] `query:"page[snapshot],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LlmTokenGatewayTokenGroupListParams]'s query parameters as
// `url.Values`.
func (r LlmTokenGatewayTokenGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LlmTokenGatewayTokenGroupDeleteParams struct {
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	IfMatch        string `header:"If-Match" api:"required" json:"-"`
	paramObj
}
