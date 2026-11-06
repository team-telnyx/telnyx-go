// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// The MCP servers and tools a key or team may use.
//
// AgentGatewayTeamToolPermissionService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayTeamToolPermissionService] method instead.
type AgentGatewayTeamToolPermissionService struct {
	Options []option.RequestOption
}

// NewAgentGatewayTeamToolPermissionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAgentGatewayTeamToolPermissionService(opts ...option.RequestOption) (r AgentGatewayTeamToolPermissionService) {
	r = AgentGatewayTeamToolPermissionService{}
	r.Options = opts
	return
}

// Replaces the grant whole and returns it. Every server must be one of the
// instance's MCP servers or APIs, named once. A key reaches the servers granted to
// it or to its team, and the servers open to every key; when both grant, only the
// servers in both. The instance must be `ready`; otherwise the request answers
// `409`.
func (r *AgentGatewayTeamToolPermissionService) New(ctx context.Context, teamID string, params AgentGatewayTeamToolPermissionNewParams, opts ...option.RequestOption) (res *AgentGatewayToolPermissionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/teams/%s/tool_permissions", params.InstanceID, teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the MCP servers and APIs the team's keys may use, and the tools within
// each. `servers: []` means none. The instance must be `ready`; otherwise the
// request answers `409`.
func (r *AgentGatewayTeamToolPermissionService) List(ctx context.Context, teamID string, query AgentGatewayTeamToolPermissionListParams, opts ...option.RequestOption) (res *AgentGatewayToolPermissionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/teams/%s/tool_permissions", query.InstanceID, teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentGatewayTeamToolPermissionNewParams struct {
	InstanceID                  string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	AgentGatewayToolPermissions AgentGatewayToolPermissionsParam
	paramObj
}

func (r AgentGatewayTeamToolPermissionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AgentGatewayToolPermissions)
}
func (r *AgentGatewayTeamToolPermissionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayTeamToolPermissionListParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
