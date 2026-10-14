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
// LlmTokenGatewayTokenKeyService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLlmTokenGatewayTokenKeyService] method instead.
type LlmTokenGatewayTokenKeyService struct {
	Options []option.RequestOption
}

// NewLlmTokenGatewayTokenKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLlmTokenGatewayTokenKeyService(opts ...option.RequestOption) (r LlmTokenGatewayTokenKeyService) {
	r = LlmTokenGatewayTokenKeyService{}
	r.Options = opts
	return
}

// Create gateway token key within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response. The token secret is returned
// only on the original create response. A replay returns metadata without the
// secret. Provisioning is advisory; ready is not a guarantee of inference
// availability.
func (r *LlmTokenGatewayTokenKeyService) New(ctx context.Context, params LlmTokenGatewayTokenKeyNewParams, opts ...option.RequestOption) (res *TokenKeyResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "llm_token_gateway/token_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get gateway token key within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenKeyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TokenKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("llm_token_gateway/token_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update gateway token key within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response. Omitted fields are preserved.
// Null clears max_parallel_requests; null resets budget and rate limits to their
// maximum. A stale If-Match returns 412.
func (r *LlmTokenGatewayTokenKeyService) Update(ctx context.Context, id string, params LlmTokenGatewayTokenKeyUpdateParams, opts ...option.RequestOption) (res *TokenKeyResponse, err error) {
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
	path := fmt.Sprintf("llm_token_gateway/token_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// List gateway token keys within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenKeyService) List(ctx context.Context, query LlmTokenGatewayTokenKeyListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[TokenKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "llm_token_gateway/token_keys"
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

// List gateway token keys within the authenticated Telnyx account. Foreign and
// unknown resource IDs return the same 404 response.
func (r *LlmTokenGatewayTokenKeyService) ListAutoPaging(ctx context.Context, query LlmTokenGatewayTokenKeyListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[TokenKey] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Revoke this token key. Acknowledged revocation blocks new admissions with the
// key; requests already admitted may complete. Historical usage attribution is
// retained. Other keys in the group are unchanged.
func (r *LlmTokenGatewayTokenKeyService) Delete(ctx context.Context, id string, body LlmTokenGatewayTokenKeyDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("llm_token_gateway/token_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type TokenKey struct {
	ID            string   `json:"id" api:"required" format:"uuid"`
	AllowedModels []string `json:"allowed_models" api:"required"`
	Blocked       bool     `json:"blocked" api:"required"`
	// Any of "1d", "7d", "30d".
	BudgetDuration  BudgetDuration `json:"budget_duration" api:"required"`
	BudgetStartedAt time.Time      `json:"budget_started_at" api:"required" format:"date-time"`
	CreatedAt       time.Time      `json:"created_at" api:"required" format:"date-time"`
	ExpiresAt       time.Time      `json:"expires_at" api:"required" format:"date-time"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	MaxBudget  float64           `json:"max_budget" api:"required"`
	Name       string            `json:"name" api:"required"`
	RecordType constant.TokenKey `json:"record_type" default:"token_key"`
	// Require a nonempty caller-asserted end-user ID; not proof of identity.
	RequiredEndUserID bool `json:"required_end_user_id" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	ReservedSpend float64   `json:"reserved_spend" api:"required"`
	ResetsAt      time.Time `json:"resets_at" api:"required" format:"date-time"`
	// Stored key limit; set to the maximum when the key was written without one (on a
	// blocked key, null is kept until the key is unblocked).
	RpmLimit int64 `json:"rpm_limit" api:"required"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	Spend        float64 `json:"spend" api:"required"`
	TokenGroupID string  `json:"token_group_id" api:"required" format:"uuid"`
	TokenUserID  string  `json:"token_user_id" api:"required" format:"uuid"`
	// Stored key limit; set to the maximum when the key was written without one (on a
	// blocked key, null is kept until the key is unblocked).
	TpmLimit  int64     `json:"tpm_limit" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	Version   int64     `json:"version" api:"required"`
	// Optional maximum concurrent reserved/dispatched cache misses across replicas.
	// Null means no concurrency cap. Cache hits do not consume slots; unknown usage
	// retains its budget hold without a slot.
	MaxParallelRequests int64 `json:"max_parallel_requests" api:"nullable"`
	// Provisioning state of the key's current policy. `pending`: retry readiness with
	// bounded backoff (treat `unknown` the same). `inactive` and `unsupported` are
	// terminal until the key, group, user or catalog changes. `unavailable` means
	// provisioning is stuck and is retried slowly; surface it. Absent while the
	// gateway runs in compatibility mode.
	//
	// Any of "ready", "pending", "unavailable", "inactive", "unsupported", "unknown".
	Provisioning TokenKeyProvisioning `json:"provisioning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AllowedModels       respjson.Field
		Blocked             respjson.Field
		BudgetDuration      respjson.Field
		BudgetStartedAt     respjson.Field
		CreatedAt           respjson.Field
		ExpiresAt           respjson.Field
		MaxBudget           respjson.Field
		Name                respjson.Field
		RecordType          respjson.Field
		RequiredEndUserID   respjson.Field
		ReservedSpend       respjson.Field
		ResetsAt            respjson.Field
		RpmLimit            respjson.Field
		Spend               respjson.Field
		TokenGroupID        respjson.Field
		TokenUserID         respjson.Field
		TpmLimit            respjson.Field
		UpdatedAt           respjson.Field
		Version             respjson.Field
		MaxParallelRequests respjson.Field
		Provisioning        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenKey) RawJSON() string { return r.JSON.raw }
func (r *TokenKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning state of the key's current policy. `pending`: retry readiness with
// bounded backoff (treat `unknown` the same). `inactive` and `unsupported` are
// terminal until the key, group, user or catalog changes. `unavailable` means
// provisioning is stuck and is retried slowly; surface it. Absent while the
// gateway runs in compatibility mode.
type TokenKeyProvisioning string

const (
	TokenKeyProvisioningReady       TokenKeyProvisioning = "ready"
	TokenKeyProvisioningPending     TokenKeyProvisioning = "pending"
	TokenKeyProvisioningUnavailable TokenKeyProvisioning = "unavailable"
	TokenKeyProvisioningInactive    TokenKeyProvisioning = "inactive"
	TokenKeyProvisioningUnsupported TokenKeyProvisioning = "unsupported"
	TokenKeyProvisioningUnknown     TokenKeyProvisioning = "unknown"
)

type TokenKeyResponse struct {
	Data TokenKey `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayTokenKeyNewParams struct {
	Name           string               `json:"name" api:"required"`
	TokenGroupID   string               `json:"token_group_id" api:"required" format:"uuid"`
	IdempotencyKey string               `header:"Idempotency-Key" api:"required" json:"-"`
	ExpiresAt      param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// Optional maximum concurrent reserved/dispatched cache misses across replicas.
	// Null means no concurrency cap. Cache hits do not consume slots; unknown usage
	// retains its budget hold without a slot.
	MaxParallelRequests param.Opt[int64] `json:"max_parallel_requests,omitzero"`
	// Omitted or null defaults to the maximum (6000) (on a blocked key, null is kept
	// until the key is unblocked); a value below the minimum (1) or above the maximum
	// is rejected with 400 limit_out_of_range; a malformed value (negative, wrong type
	// or precision) is 400 invalid_request. Block the key (blocked: true) to deny it;
	// blocking is always accepted.
	RpmLimit    param.Opt[int64]  `json:"rpm_limit,omitzero"`
	TokenUserID param.Opt[string] `json:"token_user_id,omitzero" format:"uuid"`
	// Omitted or null defaults to the maximum (10000000) (on a blocked key, null is
	// kept until the key is unblocked); a value below the minimum (1) or above the
	// maximum is rejected with 400 limit_out_of_range; a malformed value (negative,
	// wrong type or precision) is 400 invalid_request. Block the key (blocked: true)
	// to deny it; blocking is always accepted.
	TpmLimit param.Opt[int64] `json:"tpm_limit,omitzero"`
	Blocked  param.Opt[bool]  `json:"blocked,omitzero"`
	// Require a nonempty caller-asserted end-user ID; not proof of identity.
	RequiredEndUserID param.Opt[bool] `json:"required_end_user_id,omitzero"`
	AllowedModels     []string        `json:"allowed_models,omitzero"`
	// Any of "1d", "7d", "30d".
	BudgetDuration BudgetDuration `json:"budget_duration,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayTokenKeyUpdateParams struct {
	IdempotencyKey string               `header:"Idempotency-Key" api:"required" json:"-"`
	IfMatch        string               `header:"If-Match" api:"required" json:"-"`
	ExpiresAt      param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// USD, decimal precision at most 6 places; parse exactly to integer micro-USD,
	// never binary float accounting.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// Optional maximum concurrent reserved/dispatched cache misses across replicas.
	// Null means no concurrency cap. Cache hits do not consume slots; unknown usage
	// retains its budget hold without a slot.
	MaxParallelRequests param.Opt[int64] `json:"max_parallel_requests,omitzero"`
	// Omitted or null defaults to the maximum (6000) (on a blocked key, null is kept
	// until the key is unblocked); a value below the minimum (1) or above the maximum
	// is rejected with 400 limit_out_of_range; a malformed value (negative, wrong type
	// or precision) is 400 invalid_request. Block the key (blocked: true) to deny it;
	// blocking is always accepted.
	RpmLimit param.Opt[int64] `json:"rpm_limit,omitzero"`
	// Omitted or null defaults to the maximum (10000000) (on a blocked key, null is
	// kept until the key is unblocked); a value below the minimum (1) or above the
	// maximum is rejected with 400 limit_out_of_range; a malformed value (negative,
	// wrong type or precision) is 400 invalid_request. Block the key (blocked: true)
	// to deny it; blocking is always accepted.
	TpmLimit param.Opt[int64]  `json:"tpm_limit,omitzero"`
	Blocked  param.Opt[bool]   `json:"blocked,omitzero"`
	Name     param.Opt[string] `json:"name,omitzero"`
	// Require a nonempty caller-asserted end-user ID; not proof of identity. Omission
	// preserves the current requirement.
	RequiredEndUserID param.Opt[bool] `json:"required_end_user_id,omitzero"`
	AllowedModels     []string        `json:"allowed_models,omitzero"`
	// Any of "1d", "7d", "30d".
	BudgetDuration BudgetDuration `json:"budget_duration,omitzero"`
	paramObj
}

func (r LlmTokenGatewayTokenKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LlmTokenGatewayTokenKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LlmTokenGatewayTokenKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmTokenGatewayTokenKeyListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Required after page 1; account/filter-bound 15-minute stable snapshot;
	// expired/mismatch 409.
	PageSnapshot param.Opt[string] `query:"page[snapshot],omitzero" json:"-"`
	TokenGroupID param.Opt[string] `query:"token_group_id,omitzero" format:"uuid" json:"-"`
	TokenUserID  param.Opt[string] `query:"token_user_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [LlmTokenGatewayTokenKeyListParams]'s query parameters as
// `url.Values`.
func (r LlmTokenGatewayTokenKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LlmTokenGatewayTokenKeyDeleteParams struct {
	IdempotencyKey string `header:"Idempotency-Key" api:"required" json:"-"`
	IfMatch        string `header:"If-Match" api:"required" json:"-"`
	paramObj
}
