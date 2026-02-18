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
)

// AIMissionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionService] method instead.
type AIMissionService struct {
	Options        []option.RequestOption
	Runs           AIMissionRunService
	KnowledgeBases AIMissionKnowledgeBaseService
	McpServers     AIMissionMcpServerService
	Tools          AIMissionToolService
}

// NewAIMissionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMissionService(opts ...option.RequestOption) (r AIMissionService) {
	r = AIMissionService{}
	r.Options = opts
	r.Runs = NewAIMissionRunService(opts...)
	r.KnowledgeBases = NewAIMissionKnowledgeBaseService(opts...)
	r.McpServers = NewAIMissionMcpServerService(opts...)
	r.Tools = NewAIMissionToolService(opts...)
	return
}

// Create a new mission definition
func (r *AIMissionService) New(ctx context.Context, body AIMissionNewParams, opts ...option.RequestOption) (res *AIMissionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/missions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a mission by ID (includes tools, knowledge_bases, mcp_servers)
func (r *AIMissionService) Get(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all missions for the organization
func (r *AIMissionService) List(ctx context.Context, query AIMissionListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIMissionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/missions"
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

// List all missions for the organization
func (r *AIMissionService) ListAutoPaging(ctx context.Context, query AIMissionListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIMissionListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Clone an existing mission
func (r *AIMissionService) CloneMission(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionCloneMissionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/clone", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Delete a mission
func (r *AIMissionService) DeleteMission(ctx context.Context, missionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// List recent events across all missions
func (r *AIMissionService) ListEvents(ctx context.Context, query AIMissionListEventsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIMissionListEventsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/missions/events"
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

// List recent events across all missions
func (r *AIMissionService) ListEventsAutoPaging(ctx context.Context, query AIMissionListEventsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIMissionListEventsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListEvents(ctx, query, opts...))
}

// Update a mission definition
func (r *AIMissionService) UpdateMission(ctx context.Context, missionID string, body AIMissionUpdateMissionParams, opts ...option.RequestOption) (res *AIMissionUpdateMissionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AIMissionNewResponse struct {
	Data AIMissionNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionNewResponseData struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "external", "managed".
	ExecutionMode string         `json:"execution_mode,required"`
	MissionID     string         `json:"mission_id,required" format:"uuid"`
	Name          string         `json:"name,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Description   string         `json:"description"`
	Instructions  string         `json:"instructions"`
	Metadata      map[string]any `json:"metadata"`
	Model         string         `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		ExecutionMode respjson.Field
		MissionID     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		Instructions  respjson.Field
		Metadata      respjson.Field
		Model         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionGetResponse struct {
	Data AIMissionGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionGetResponseData struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "external", "managed".
	ExecutionMode string         `json:"execution_mode,required"`
	MissionID     string         `json:"mission_id,required" format:"uuid"`
	Name          string         `json:"name,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Description   string         `json:"description"`
	Instructions  string         `json:"instructions"`
	Metadata      map[string]any `json:"metadata"`
	Model         string         `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		ExecutionMode respjson.Field
		MissionID     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		Instructions  respjson.Field
		Metadata      respjson.Field
		Model         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionListResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "external", "managed".
	ExecutionMode AIMissionListResponseExecutionMode `json:"execution_mode,required"`
	MissionID     string                             `json:"mission_id,required" format:"uuid"`
	Name          string                             `json:"name,required"`
	UpdatedAt     time.Time                          `json:"updated_at,required" format:"date-time"`
	Description   string                             `json:"description"`
	Instructions  string                             `json:"instructions"`
	Metadata      map[string]any                     `json:"metadata"`
	Model         string                             `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		ExecutionMode respjson.Field
		MissionID     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		Instructions  respjson.Field
		Metadata      respjson.Field
		Model         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionListResponseExecutionMode string

const (
	AIMissionListResponseExecutionModeExternal AIMissionListResponseExecutionMode = "external"
	AIMissionListResponseExecutionModeManaged  AIMissionListResponseExecutionMode = "managed"
)

type AIMissionCloneMissionResponse = any

type AIMissionListEventsResponse struct {
	EventID   string    `json:"event_id,required"`
	RunID     string    `json:"run_id,required"`
	Summary   string    `json:"summary,required"`
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Any of "status_change", "step_started", "step_completed", "step_failed",
	// "tool_call", "tool_result", "message", "error", "custom".
	Type           AIMissionListEventsResponseType `json:"type,required"`
	AgentID        string                          `json:"agent_id"`
	IdempotencyKey string                          `json:"idempotency_key"`
	Payload        map[string]any                  `json:"payload"`
	StepID         string                          `json:"step_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID        respjson.Field
		RunID          respjson.Field
		Summary        respjson.Field
		Timestamp      respjson.Field
		Type           respjson.Field
		AgentID        respjson.Field
		IdempotencyKey respjson.Field
		Payload        respjson.Field
		StepID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionListEventsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionListEventsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionListEventsResponseType string

const (
	AIMissionListEventsResponseTypeStatusChange  AIMissionListEventsResponseType = "status_change"
	AIMissionListEventsResponseTypeStepStarted   AIMissionListEventsResponseType = "step_started"
	AIMissionListEventsResponseTypeStepCompleted AIMissionListEventsResponseType = "step_completed"
	AIMissionListEventsResponseTypeStepFailed    AIMissionListEventsResponseType = "step_failed"
	AIMissionListEventsResponseTypeToolCall      AIMissionListEventsResponseType = "tool_call"
	AIMissionListEventsResponseTypeToolResult    AIMissionListEventsResponseType = "tool_result"
	AIMissionListEventsResponseTypeMessage       AIMissionListEventsResponseType = "message"
	AIMissionListEventsResponseTypeError         AIMissionListEventsResponseType = "error"
	AIMissionListEventsResponseTypeCustom        AIMissionListEventsResponseType = "custom"
)

type AIMissionUpdateMissionResponse struct {
	Data AIMissionUpdateMissionResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionUpdateMissionResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionUpdateMissionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionUpdateMissionResponseData struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Any of "external", "managed".
	ExecutionMode string         `json:"execution_mode,required"`
	MissionID     string         `json:"mission_id,required" format:"uuid"`
	Name          string         `json:"name,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Description   string         `json:"description"`
	Instructions  string         `json:"instructions"`
	Metadata      map[string]any `json:"metadata"`
	Model         string         `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		ExecutionMode respjson.Field
		MissionID     respjson.Field
		Name          respjson.Field
		UpdatedAt     respjson.Field
		Description   respjson.Field
		Instructions  respjson.Field
		Metadata      respjson.Field
		Model         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionUpdateMissionResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionUpdateMissionResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionNewParams struct {
	Name         string            `json:"name,required"`
	Description  param.Opt[string] `json:"description,omitzero"`
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	Model        param.Opt[string] `json:"model,omitzero"`
	// Any of "external", "managed".
	ExecutionMode AIMissionNewParamsExecutionMode `json:"execution_mode,omitzero"`
	Metadata      map[string]any                  `json:"metadata,omitzero"`
	paramObj
}

func (r AIMissionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionNewParamsExecutionMode string

const (
	AIMissionNewParamsExecutionModeExternal AIMissionNewParamsExecutionMode = "external"
	AIMissionNewParamsExecutionModeManaged  AIMissionNewParamsExecutionMode = "managed"
)

type AIMissionListParams struct {
	// Page number (1-based)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMissionListParams]'s query parameters as `url.Values`.
func (r AIMissionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMissionListEventsParams struct {
	// Page number (1-based)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	Type     param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMissionListEventsParams]'s query parameters as
// `url.Values`.
func (r AIMissionListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMissionUpdateMissionParams struct {
	Description  param.Opt[string] `json:"description,omitzero"`
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	Model        param.Opt[string] `json:"model,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	// Any of "external", "managed".
	ExecutionMode AIMissionUpdateMissionParamsExecutionMode `json:"execution_mode,omitzero"`
	Metadata      map[string]any                            `json:"metadata,omitzero"`
	paramObj
}

func (r AIMissionUpdateMissionParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionUpdateMissionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionUpdateMissionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionUpdateMissionParamsExecutionMode string

const (
	AIMissionUpdateMissionParamsExecutionModeExternal AIMissionUpdateMissionParamsExecutionMode = "external"
	AIMissionUpdateMissionParamsExecutionModeManaged  AIMissionUpdateMissionParamsExecutionMode = "managed"
)
