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

// AIMissionRunTelnyxAgentService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionRunTelnyxAgentService] method instead.
type AIMissionRunTelnyxAgentService struct {
	Options []option.RequestOption
}

// NewAIMissionRunTelnyxAgentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMissionRunTelnyxAgentService(opts ...option.RequestOption) (r AIMissionRunTelnyxAgentService) {
	r = AIMissionRunTelnyxAgentService{}
	r.Options = opts
	return
}

// List all Telnyx agents linked to a run
func (r *AIMissionRunTelnyxAgentService) List(ctx context.Context, runID string, query AIMissionRunTelnyxAgentListParams, opts ...option.RequestOption) (res *AIMissionRunTelnyxAgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/telnyx-agents", query.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Link a Telnyx AI agent (voice/messaging) to a run
func (r *AIMissionRunTelnyxAgentService) Link(ctx context.Context, runID string, params AIMissionRunTelnyxAgentLinkParams, opts ...option.RequestOption) (res *AIMissionRunTelnyxAgentLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/telnyx-agents", params.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Unlink a Telnyx agent from a run
func (r *AIMissionRunTelnyxAgentService) Unlink(ctx context.Context, telnyxAgentID string, body AIMissionRunTelnyxAgentUnlinkParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if body.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	if telnyxAgentID == "" {
		err = errors.New("missing required telnyx_agent_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/telnyx-agents/%s", body.MissionID, body.RunID, telnyxAgentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AIMissionRunTelnyxAgentListResponse struct {
	Data []AIMissionRunTelnyxAgentListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunTelnyxAgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunTelnyxAgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunTelnyxAgentListResponseData struct {
	CreatedAt     time.Time `json:"created_at" api:"required" format:"date-time"`
	RunID         string    `json:"run_id" api:"required"`
	TelnyxAgentID string    `json:"telnyx_agent_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		RunID         respjson.Field
		TelnyxAgentID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunTelnyxAgentListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunTelnyxAgentListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunTelnyxAgentLinkResponse struct {
	Data AIMissionRunTelnyxAgentLinkResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunTelnyxAgentLinkResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunTelnyxAgentLinkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunTelnyxAgentLinkResponseData struct {
	CreatedAt     time.Time `json:"created_at" api:"required" format:"date-time"`
	RunID         string    `json:"run_id" api:"required"`
	TelnyxAgentID string    `json:"telnyx_agent_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		RunID         respjson.Field
		TelnyxAgentID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunTelnyxAgentLinkResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunTelnyxAgentLinkResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunTelnyxAgentListParams struct {
	MissionID string `path:"mission_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunTelnyxAgentLinkParams struct {
	MissionID string `path:"mission_id" api:"required" format:"uuid" json:"-"`
	// The Telnyx AI agent ID to link
	TelnyxAgentID string `json:"telnyx_agent_id" api:"required"`
	paramObj
}

func (r AIMissionRunTelnyxAgentLinkParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunTelnyxAgentLinkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunTelnyxAgentLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunTelnyxAgentUnlinkParams struct {
	MissionID string `path:"mission_id" api:"required" format:"uuid" json:"-"`
	RunID     string `path:"run_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
