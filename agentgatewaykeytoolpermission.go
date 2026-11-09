// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// The MCP servers and tools a key or team may use.
//
// AgentGatewayKeyToolPermissionService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGatewayKeyToolPermissionService] method instead.
type AgentGatewayKeyToolPermissionService struct {
	Options []option.RequestOption
}

// NewAgentGatewayKeyToolPermissionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAgentGatewayKeyToolPermissionService(opts ...option.RequestOption) (r AgentGatewayKeyToolPermissionService) {
	r = AgentGatewayKeyToolPermissionService{}
	r.Options = opts
	return
}

// Replaces the grant whole and returns it. Every server must be one of the
// instance's MCP servers or APIs, named once. A key reaches the servers granted to
// it or to its team, and the servers open to every key; when both grant, only the
// servers in both. The instance must be `ready`; otherwise the request answers
// `409`.
func (r *AgentGatewayKeyToolPermissionService) New(ctx context.Context, keyID string, params AgentGatewayKeyToolPermissionNewParams, opts ...option.RequestOption) (res *AgentGatewayToolPermissionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/keys/%s/tool_permissions", params.InstanceID, url.PathEscape(keyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the MCP servers and APIs the key may use, and the tools within each.
// `servers: []` means none. The instance must be `ready`; otherwise the request
// answers `409`.
func (r *AgentGatewayKeyToolPermissionService) List(ctx context.Context, keyID string, query AgentGatewayKeyToolPermissionListParams, opts ...option.RequestOption) (res *AgentGatewayToolPermissionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.InstanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	if keyID == "" {
		err = errors.New("missing required key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("agent_gateways/%s/keys/%s/tool_permissions", query.InstanceID, url.PathEscape(keyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentGatewayToolPermissions struct {
	// The servers granted, each with its tools. `[]` grants none.
	Servers []AgentGatewayToolPermissionsServer `json:"servers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Servers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayToolPermissions) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayToolPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentGatewayToolPermissions to a
// AgentGatewayToolPermissionsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentGatewayToolPermissionsParam.Overrides()
func (r AgentGatewayToolPermissions) ToParam() AgentGatewayToolPermissionsParam {
	return param.Override[AgentGatewayToolPermissionsParam](json.RawMessage(r.RawJSON()))
}

type AgentGatewayToolPermissionsServer struct {
	// An MCP server's or API's `id` on the instance.
	ID string `json:"id" api:"required" format:"uuid"`
	// The tools the grant covers, by name without the server prefix. `[]` covers every
	// tool the server exposes.
	Tools []string `json:"tools"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Tools       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayToolPermissionsServer) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayToolPermissionsServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Servers is required.
type AgentGatewayToolPermissionsParam struct {
	// The servers granted, each with its tools. `[]` grants none.
	Servers []AgentGatewayToolPermissionsServerParam `json:"servers,omitzero" api:"required"`
	paramObj
}

func (r AgentGatewayToolPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayToolPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayToolPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type AgentGatewayToolPermissionsServerParam struct {
	// An MCP server's or API's `id` on the instance.
	ID string `json:"id" api:"required" format:"uuid"`
	// The tools the grant covers, by name without the server prefix. `[]` covers every
	// tool the server exposes.
	Tools []string `json:"tools,omitzero"`
	paramObj
}

func (r AgentGatewayToolPermissionsServerParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentGatewayToolPermissionsServerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGatewayToolPermissionsServerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayToolPermissionsResponse struct {
	Data AgentGatewayToolPermissions `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGatewayToolPermissionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGatewayToolPermissionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyToolPermissionNewParams struct {
	InstanceID                  string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	AgentGatewayToolPermissions AgentGatewayToolPermissionsParam
	paramObj
}

func (r AgentGatewayKeyToolPermissionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AgentGatewayToolPermissions)
}
func (r *AgentGatewayKeyToolPermissionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGatewayKeyToolPermissionListParams struct {
	InstanceID string `path:"instance_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
