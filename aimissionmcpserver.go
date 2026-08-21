// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AIMissionMcpServerService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionMcpServerService] method instead.
type AIMissionMcpServerService struct {
	Options []option.RequestOption
}

// NewAIMissionMcpServerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMissionMcpServerService(opts ...option.RequestOption) (r AIMissionMcpServerService) {
	r = AIMissionMcpServerService{}
	r.Options = opts
	return
}

// Adds an MCP server to the specified mission, making the server's tools available
// to agents during runs of this mission.
func (r *AIMissionMcpServerService) NewMcpServer(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionMcpServerNewMcpServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/mcp-servers", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Removes the specified MCP server from the mission, revoking agent access to its
// tools in subsequent runs.
func (r *AIMissionMcpServerService) DeleteMcpServer(ctx context.Context, mcpServerID string, body AIMissionMcpServerDeleteMcpServerParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return err
	}
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return err
	}
	path := fmt.Sprintf("ai/missions/%s/mcp-servers/%s", body.MissionID, mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the configuration of a single MCP server attached to the specified
// mission.
func (r *AIMissionMcpServerService) GetMcpServer(ctx context.Context, mcpServerID string, query AIMissionMcpServerGetMcpServerParams, opts ...option.RequestOption) (res *AIMissionMcpServerGetMcpServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/mcp-servers/%s", query.MissionID, mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the MCP servers configured on the specified mission. MCP servers expose
// external tools and data sources agents can use during runs.
func (r *AIMissionMcpServerService) ListMcpServers(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionMcpServerListMcpServersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/mcp-servers", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replaces the configuration of the specified MCP server on this mission.
func (r *AIMissionMcpServerService) UpdateMcpServer(ctx context.Context, mcpServerID string, body AIMissionMcpServerUpdateMcpServerParams, opts ...option.RequestOption) (res *AIMissionMcpServerUpdateMcpServerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/mcp-servers/%s", body.MissionID, mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type AIMissionMcpServerNewMcpServerResponse = any

type AIMissionMcpServerGetMcpServerResponse = any

type AIMissionMcpServerListMcpServersResponse = any

type AIMissionMcpServerUpdateMcpServerResponse = any

type AIMissionMcpServerDeleteMcpServerParams struct {
	MissionID string `path:"mission_id" api:"required" json:"-"`
	paramObj
}

type AIMissionMcpServerGetMcpServerParams struct {
	MissionID string `path:"mission_id" api:"required" json:"-"`
	paramObj
}

type AIMissionMcpServerUpdateMcpServerParams struct {
	MissionID string `path:"mission_id" api:"required" json:"-"`
	paramObj
}
