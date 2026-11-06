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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Keys on an instance.
//
// AgentGatewayKeyService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayKeyService] method instead.
type AgentGatewayKeyService struct {
	Options []option.RequestOption
	// The MCP servers and tools a key or team may use.
	ToolPermissions AgentGatewayKeyToolPermissionService
}

// NewAgentGatewayKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentGatewayKeyService(opts ...option.RequestOption) (r AgentGatewayKeyService) {
	r = AgentGatewayKeyService{}
	r.Options = opts
	r.ToolPermissions = NewAgentGatewayKeyToolPermissionService(opts...)
	return
}

// Issues a key on the instance. The key's value is in this response only. Every
// field is optional; a key without `models` may use every model the instance
// serves. The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayKeyService) New(ctx context.Context, instanceID string, body AgentGatewayKeyNewParams, opts ...option.RequestOption) (res *AgentGatewayKeyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/keys", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Changes a key's settings without reissuing it. A field left out stays as it is,
// and `null` clears it. `models: []` lifts the model restriction. `expires_in`
// counts from now. Setting `blocked` to `true` makes the instance refuse the key
// with `401` while keeping its settings and spend; `false` lets it through again.
// The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayKeyService) Update(ctx context.Context, keyID string, params AgentGatewayKeyUpdateParams, opts ...option.RequestOption) (res *AgentGatewayKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/keys/%s", params.InstanceID, url.PathEscape(keyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the instance's keys, without their values. Returns the first 100. The
// instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayKeyService) List(ctx context.Context, instanceID string, query AgentGatewayKeyListParams, opts ...option.RequestOption) (res *AgentGatewayKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/keys", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Revokes a key and returns it as it was, without its value. The instance must be
// `ready`; otherwise the request answers `409`.
func (r *AgentGatewayKeyService) Delete(ctx context.Context, keyID string, body AgentGatewayKeyDeleteParams, opts ...option.RequestOption) (res *AgentGatewayKeyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/keys/%s", body.InstanceID, url.PathEscape(keyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A key on an instance, without its value.
type AgentGatewayKey struct {
	// Identifies the key. It is a 64-character hash the instance stores in place of
	// the key's value.
	ID string `json:"id" api:"required"`
	// Whether the instance refuses the key.
	Blocked bool `json:"blocked" api:"required"`
	// When the key's spend next resets.
	BudgetResetsAt time.Time `json:"budget_resets_at" api:"required" format:"date-time"`
	// How often the key's spend resets, such as `30d`.
	BudgetResetsEvery string `json:"budget_resets_every" api:"required"`
	// When the key was issued.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key stops working; `null` when it does not expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The most the key may spend, in USD, until its budget resets.
	MaxBudget float64 `json:"max_budget" api:"required"`
	// The most requests the key may have in flight at once.
	MaxParallelRequests int64 `json:"max_parallel_requests" api:"required"`
	// The models the key may use. Empty means every model the instance serves.
	Models []string `json:"models" api:"required"`
	// A label for people; `null` when none was given.
	Name string `json:"name" api:"required"`
	// The most requests the key may make per minute.
	RequestsPerMinute int64 `json:"requests_per_minute" api:"required"`
	// What the key has spent since its budget last reset, in USD.
	Spend float64 `json:"spend" api:"required"`
	// The team the key belongs to, if any.
	TeamID string `json:"team_id" api:"required"`
	// The most tokens the key may use per minute.
	TokensPerMinute int64 `json:"tokens_per_minute" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Blocked             respjson.Field
		BudgetResetsAt      respjson.Field
		BudgetResetsEvery   respjson.Field
		CreatedAt           respjson.Field
		ExpiresAt           respjson.Field
		MaxBudget           respjson.Field
		MaxParallelRequests respjson.Field
		Models              respjson.Field
		Name                respjson.Field
		RequestsPerMinute   respjson.Field
		Spend               respjson.Field
		TeamID              respjson.Field
		TokensPerMinute     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayKey) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyResponse struct {
	// A key on an instance, without its value.
	Data AgentGatewayKey `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyNewResponse struct {
	// A key that was just issued: the only response that carries its value.
	Data AgentGatewayKeyNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A key that was just issued: the only response that carries its value.
type AgentGatewayKeyNewResponseData struct {
	// Identifies the key. It is a 64-character hash the instance stores in place of
	// the key's value.
	ID string `json:"id" api:"required"`
	// Whether the instance refuses the key.
	Blocked bool `json:"blocked" api:"required"`
	// When the key's spend next resets.
	BudgetResetsAt time.Time `json:"budget_resets_at" api:"required" format:"date-time"`
	// How often the key's spend resets, such as `30d`.
	BudgetResetsEvery string `json:"budget_resets_every" api:"required"`
	// When the key was issued.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key stops working; `null` when it does not expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// The key's value, sent as `Authorization: Bearer <key>` to the instance. Shown
	// only in this response.
	Key string `json:"key" api:"required"`
	// The most the key may spend, in USD, until its budget resets.
	MaxBudget float64 `json:"max_budget" api:"required"`
	// The most requests the key may have in flight at once.
	MaxParallelRequests int64 `json:"max_parallel_requests" api:"required"`
	// The models the key may use. Empty means every model the instance serves.
	Models []string `json:"models" api:"required"`
	// A label for people; `null` when none was given.
	Name string `json:"name" api:"required"`
	// The most requests the key may make per minute.
	RequestsPerMinute int64 `json:"requests_per_minute" api:"required"`
	// What the key has spent since its budget last reset, in USD.
	Spend float64 `json:"spend" api:"required"`
	// The team the key belongs to, if any.
	TeamID string `json:"team_id" api:"required"`
	// The most tokens the key may use per minute.
	TokensPerMinute int64 `json:"tokens_per_minute" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Blocked             respjson.Field
		BudgetResetsAt      respjson.Field
		BudgetResetsEvery   respjson.Field
		CreatedAt           respjson.Field
		ExpiresAt           respjson.Field
		Key                 respjson.Field
		MaxBudget           respjson.Field
		MaxParallelRequests respjson.Field
		Models              respjson.Field
		Name                respjson.Field
		RequestsPerMinute   respjson.Field
		Spend               respjson.Field
		TeamID              respjson.Field
		TokensPerMinute     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayKeyNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayKeyNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyListResponse struct {
	Data []AgentGatewayKey `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyNewParams struct {
	// How often the spend resets: a number and a unit, `s`, `m`, `h`, or `d`.
	BudgetResetsEvery param.Opt[string] `json:"budget_resets_every,omitzero"`
	// How long until the key stops working: a number and a unit, `s`, `m`, `h`, or
	// `d`.
	ExpiresIn param.Opt[string] `json:"expires_in,omitzero"`
	// The most the key may spend, in USD.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// The most requests the key may have in flight at once.
	MaxParallelRequests param.Opt[int64] `json:"max_parallel_requests,omitzero"`
	// A label for people, unique on the instance.
	Name param.Opt[string] `json:"name,omitzero"`
	// The most requests the key may make per minute.
	RequestsPerMinute param.Opt[int64] `json:"requests_per_minute,omitzero"`
	// A team on the instance. The key shares the team's model list, budget, and rate
	// limits.
	TeamID param.Opt[string] `json:"team_id,omitzero" format:"uuid"`
	// The most tokens the key may use per minute.
	TokensPerMinute param.Opt[int64] `json:"tokens_per_minute,omitzero"`
	// The models the key may use. Leave it out to allow every model.
	Models []string `json:"models,omitzero"`
	paramObj
}

func (r AgentGatewayKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyUpdateParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	// How often the spend resets.
	BudgetResetsEvery param.Opt[string] `json:"budget_resets_every,omitzero"`
	// How long from now until the key stops working.
	ExpiresIn param.Opt[string] `json:"expires_in,omitzero"`
	// The most the key may spend, in USD.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// The most requests the key may have in flight at once.
	MaxParallelRequests param.Opt[int64] `json:"max_parallel_requests,omitzero"`
	// The most requests the key may make per minute.
	RequestsPerMinute param.Opt[int64] `json:"requests_per_minute,omitzero"`
	// Move the key to this team, or `null` to leave its team.
	TeamID param.Opt[string] `json:"team_id,omitzero" format:"uuid"`
	// The most tokens the key may use per minute.
	TokensPerMinute param.Opt[int64] `json:"tokens_per_minute,omitzero"`
	// `true` makes the instance refuse the key; `false` lets it through again.
	Blocked param.Opt[bool] `json:"blocked,omitzero"`
	// A label for people.
	Name param.Opt[string] `json:"name,omitzero"`
	// The models the key may use. `[]` allows every model.
	Models []string `json:"models,omitzero"`
	paramObj
}

func (r AgentGatewayKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyListParams struct {
	// Only the keys of this team.
	TeamID param.Opt[string] `query:"team_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AgentGatewayKeyListParams]'s query parameters as
// `url.Values`.
func (r AgentGatewayKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentGatewayKeyDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
