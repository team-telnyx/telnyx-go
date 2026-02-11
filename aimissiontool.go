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

// AIMissionToolService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionToolService] method instead.
type AIMissionToolService struct {
	Options []option.RequestOption
}

// NewAIMissionToolService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMissionToolService(opts ...option.RequestOption) (r AIMissionToolService) {
	r = AIMissionToolService{}
	r.Options = opts
	return
}

// Create a new tool for a mission
func (r *AIMissionToolService) NewTool(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionToolNewToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/tools", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Delete a tool from a mission
func (r *AIMissionToolService) DeleteTool(ctx context.Context, toolID string, body AIMissionToolDeleteToolParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/tools/%s", body.MissionID, toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get a specific tool by ID
func (r *AIMissionToolService) GetTool(ctx context.Context, toolID string, query AIMissionToolGetToolParams, opts ...option.RequestOption) (res *AIMissionToolGetToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/tools/%s", query.MissionID, toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all tools for a mission
func (r *AIMissionToolService) ListTools(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionToolListToolsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/tools", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a tool definition
func (r *AIMissionToolService) UpdateTool(ctx context.Context, toolID string, body AIMissionToolUpdateToolParams, opts ...option.RequestOption) (res *AIMissionToolUpdateToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/tools/%s", body.MissionID, toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type AIMissionToolNewToolResponse = any

type AIMissionToolGetToolResponse = any

type AIMissionToolListToolsResponse = any

type AIMissionToolUpdateToolResponse = any

type AIMissionToolDeleteToolParams struct {
	MissionID string `path:"mission_id,required" json:"-"`
	paramObj
}

type AIMissionToolGetToolParams struct {
	MissionID string `path:"mission_id,required" json:"-"`
	paramObj
}

type AIMissionToolUpdateToolParams struct {
	MissionID string `path:"mission_id,required" json:"-"`
	paramObj
}
