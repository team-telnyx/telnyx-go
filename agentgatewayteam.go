// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Teams on an instance: shared model lists, budgets, and rate limits.
//
// AgentGatewayTeamService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayTeamService] method instead.
type AgentGatewayTeamService struct {
	Options []option.RequestOption
	// The MCP servers and tools a key or team may use.
	ToolPermissions AgentGatewayTeamToolPermissionService
}

// NewAgentGatewayTeamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentGatewayTeamService(opts ...option.RequestOption) (r AgentGatewayTeamService) {
	r = AgentGatewayTeamService{}
	r.Options = opts
	r.ToolPermissions = NewAgentGatewayTeamToolPermissionService(opts...)
	return
}

// Creates a team on the instance. A team is a model list, a budget, and rate
// limits that its keys share. Names need not be unique; the `id` identifies the
// team. The instance must be `ready`; otherwise the request answers `409`.
func (r *AgentGatewayTeamService) New(ctx context.Context, instanceID string, body AgentGatewayTeamNewParams, opts ...option.RequestOption) (res *AgentGatewayTeamResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/teams", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Changes a team's settings. A field left out stays as it is, and `null` clears
// it. `models: []` lifts the model restriction. Setting `blocked` to `true` makes
// the instance refuse every key in the team with `401`. The instance must be
// `ready`; otherwise the request answers `409`.
func (r *AgentGatewayTeamService) Update(ctx context.Context, teamID string, params AgentGatewayTeamUpdateParams, opts ...option.RequestOption) (res *AgentGatewayTeamResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/teams/%s", params.InstanceID, teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the instance's teams. The instance must be `ready`; otherwise the request
// answers `409`.
func (r *AgentGatewayTeamService) List(ctx context.Context, instanceID string, opts ...option.RequestOption) (res *AgentGatewayTeamListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/teams", instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a team that has no keys and returns it. A team that still has keys
// answers `409`; delete or move its keys first. The instance must be `ready`;
// otherwise the request answers `409`.
func (r *AgentGatewayTeamService) Delete(ctx context.Context, teamID string, body AgentGatewayTeamDeleteParams, opts ...option.RequestOption) (res *AgentGatewayTeamResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/teams/%s", body.InstanceID, teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A team on an instance: a model list, a budget, and rate limits that its keys
// share.
type AgentGatewayTeam struct {
	// Identifies the team.
	ID string `json:"id" api:"required"`
	// Whether the instance refuses every key in the team.
	Blocked bool `json:"blocked" api:"required"`
	// When the team's spend next resets.
	BudgetResetsAt time.Time `json:"budget_resets_at" api:"required" format:"date-time"`
	// How often the team's spend resets, such as `30d`.
	BudgetResetsEvery string `json:"budget_resets_every" api:"required"`
	// When the team was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The most the team's keys may spend together, in USD, until the budget resets.
	MaxBudget float64 `json:"max_budget" api:"required"`
	// The models the team's keys may use. Empty means every model the instance serves.
	Models []string `json:"models" api:"required"`
	// A label for people. Names need not be unique.
	Name string `json:"name" api:"required"`
	// The most requests the team's keys may make per minute, together.
	RequestsPerMinute int64 `json:"requests_per_minute" api:"required"`
	// What the team's keys have spent since the budget last reset, in USD.
	Spend float64 `json:"spend" api:"required"`
	// The most tokens the team's keys may use per minute, together.
	TokensPerMinute int64 `json:"tokens_per_minute" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Blocked           respjson.Field
		BudgetResetsAt    respjson.Field
		BudgetResetsEvery respjson.Field
		CreatedAt         respjson.Field
		MaxBudget         respjson.Field
		Models            respjson.Field
		Name              respjson.Field
		RequestsPerMinute respjson.Field
		Spend             respjson.Field
		TokensPerMinute   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayTeam) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayTeam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayTeamResponse struct {
	// A team on an instance: a model list, a budget, and rate limits that its keys
	// share.
	Data AgentGatewayTeam `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayTeamResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayTeamResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayTeamListResponse struct {
	Data []AgentGatewayTeam `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayTeamListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayTeamListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayTeamNewParams struct {
	// A label for people.
	Name string `json:"name" api:"required"`
	// How often the spend resets: a number and a unit, `s`, `m`, `h`, or `d`.
	BudgetResetsEvery param.Opt[string] `json:"budget_resets_every,omitzero"`
	// The most the team's keys may spend together, in USD.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// The most requests the team's keys may make per minute, together.
	RequestsPerMinute param.Opt[int64] `json:"requests_per_minute,omitzero"`
	// The most tokens the team's keys may use per minute, together.
	TokensPerMinute param.Opt[int64] `json:"tokens_per_minute,omitzero"`
	// The models the team's keys may use. Leave it out to allow every model.
	Models []string `json:"models,omitzero"`
	paramObj
}

func (r AgentGatewayTeamNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayTeamNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayTeamNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayTeamUpdateParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	// How often the spend resets.
	BudgetResetsEvery param.Opt[string] `json:"budget_resets_every,omitzero"`
	// The most the team's keys may spend together, in USD.
	MaxBudget param.Opt[float64] `json:"max_budget,omitzero"`
	// The most requests the team's keys may make per minute, together.
	RequestsPerMinute param.Opt[int64] `json:"requests_per_minute,omitzero"`
	// The most tokens the team's keys may use per minute, together.
	TokensPerMinute param.Opt[int64] `json:"tokens_per_minute,omitzero"`
	// `true` makes the instance refuse every key in the team; `false` lets them
	// through again.
	Blocked param.Opt[bool] `json:"blocked,omitzero"`
	// A label for people.
	Name param.Opt[string] `json:"name,omitzero"`
	// The models the team's keys may use. `[]` allows every model.
	Models []string `json:"models,omitzero"`
	paramObj
}

func (r AgentGatewayTeamUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayTeamUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayTeamUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayTeamDeleteParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
