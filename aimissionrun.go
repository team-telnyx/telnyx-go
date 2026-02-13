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

// AIMissionRunService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionRunService] method instead.
type AIMissionRunService struct {
	Options      []option.RequestOption
	Events       AIMissionRunEventService
	Plan         AIMissionRunPlanService
	TelnyxAgents AIMissionRunTelnyxAgentService
}

// NewAIMissionRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMissionRunService(opts ...option.RequestOption) (r AIMissionRunService) {
	r = AIMissionRunService{}
	r.Options = opts
	r.Events = NewAIMissionRunEventService(opts...)
	r.Plan = NewAIMissionRunPlanService(opts...)
	r.TelnyxAgents = NewAIMissionRunTelnyxAgentService(opts...)
	return
}

// Start a new run for a mission
func (r *AIMissionRunService) New(ctx context.Context, missionID string, body AIMissionRunNewParams, opts ...option.RequestOption) (res *AIMissionRunNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details of a specific run
func (r *AIMissionRunService) Get(ctx context.Context, runID string, query AIMissionRunGetParams, opts ...option.RequestOption) (res *AIMissionRunGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s", query.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update run status and/or result
func (r *AIMissionRunService) Update(ctx context.Context, runID string, params AIMissionRunUpdateParams, opts ...option.RequestOption) (res *AIMissionRunUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s", params.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all runs for a specific mission
func (r *AIMissionRunService) List(ctx context.Context, missionID string, query AIMissionRunListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIMissionRunListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs", missionID)
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

// List all runs for a specific mission
func (r *AIMissionRunService) ListAutoPaging(ctx context.Context, missionID string, query AIMissionRunListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIMissionRunListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, missionID, query, opts...))
}

// Cancel a running or paused run
func (r *AIMissionRunService) CancelRun(ctx context.Context, runID string, body AIMissionRunCancelRunParams, opts ...option.RequestOption) (res *AIMissionRunCancelRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/cancel", body.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List recent runs across all missions
func (r *AIMissionRunService) ListRuns(ctx context.Context, query AIMissionRunListRunsParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AIMissionRunListRunsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/missions/runs"
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

// List recent runs across all missions
func (r *AIMissionRunService) ListRunsAutoPaging(ctx context.Context, query AIMissionRunListRunsParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AIMissionRunListRunsResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.ListRuns(ctx, query, opts...))
}

// Pause a running run
func (r *AIMissionRunService) PauseRun(ctx context.Context, runID string, body AIMissionRunPauseRunParams, opts ...option.RequestOption) (res *AIMissionRunPauseRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/pause", body.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Resume a paused run
func (r *AIMissionRunService) ResumeRun(ctx context.Context, runID string, body AIMissionRunResumeRunParams, opts ...option.RequestOption) (res *AIMissionRunResumeRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/resume", body.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AIMissionRunNewResponse struct {
	Data AIMissionRunNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunNewResponseData struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        string         `json:"status,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Error         string         `json:"error"`
	FinishedAt    time.Time      `json:"finished_at" format:"date-time"`
	Input         map[string]any `json:"input"`
	Metadata      map[string]any `json:"metadata"`
	ResultPayload map[string]any `json:"result_payload"`
	ResultSummary string         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunGetResponse struct {
	Data AIMissionRunGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunGetResponseData struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        string         `json:"status,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Error         string         `json:"error"`
	FinishedAt    time.Time      `json:"finished_at" format:"date-time"`
	Input         map[string]any `json:"input"`
	Metadata      map[string]any `json:"metadata"`
	ResultPayload map[string]any `json:"result_payload"`
	ResultSummary string         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunUpdateResponse struct {
	Data AIMissionRunUpdateResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunUpdateResponseData struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        string         `json:"status,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Error         string         `json:"error"`
	FinishedAt    time.Time      `json:"finished_at" format:"date-time"`
	Input         map[string]any `json:"input"`
	Metadata      map[string]any `json:"metadata"`
	ResultPayload map[string]any `json:"result_payload"`
	ResultSummary string         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunListResponse struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        AIMissionRunListResponseStatus `json:"status,required"`
	UpdatedAt     time.Time                      `json:"updated_at,required" format:"date-time"`
	Error         string                         `json:"error"`
	FinishedAt    time.Time                      `json:"finished_at" format:"date-time"`
	Input         map[string]any                 `json:"input"`
	Metadata      map[string]any                 `json:"metadata"`
	ResultPayload map[string]any                 `json:"result_payload"`
	ResultSummary string                         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunListResponseStatus string

const (
	AIMissionRunListResponseStatusPending   AIMissionRunListResponseStatus = "pending"
	AIMissionRunListResponseStatusRunning   AIMissionRunListResponseStatus = "running"
	AIMissionRunListResponseStatusPaused    AIMissionRunListResponseStatus = "paused"
	AIMissionRunListResponseStatusSucceeded AIMissionRunListResponseStatus = "succeeded"
	AIMissionRunListResponseStatusFailed    AIMissionRunListResponseStatus = "failed"
	AIMissionRunListResponseStatusCancelled AIMissionRunListResponseStatus = "cancelled"
)

type AIMissionRunCancelRunResponse struct {
	Data AIMissionRunCancelRunResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunCancelRunResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunCancelRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunCancelRunResponseData struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        string         `json:"status,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Error         string         `json:"error"`
	FinishedAt    time.Time      `json:"finished_at" format:"date-time"`
	Input         map[string]any `json:"input"`
	Metadata      map[string]any `json:"metadata"`
	ResultPayload map[string]any `json:"result_payload"`
	ResultSummary string         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunCancelRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunCancelRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunListRunsResponse struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        AIMissionRunListRunsResponseStatus `json:"status,required"`
	UpdatedAt     time.Time                          `json:"updated_at,required" format:"date-time"`
	Error         string                             `json:"error"`
	FinishedAt    time.Time                          `json:"finished_at" format:"date-time"`
	Input         map[string]any                     `json:"input"`
	Metadata      map[string]any                     `json:"metadata"`
	ResultPayload map[string]any                     `json:"result_payload"`
	ResultSummary string                             `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunListRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunListRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunListRunsResponseStatus string

const (
	AIMissionRunListRunsResponseStatusPending   AIMissionRunListRunsResponseStatus = "pending"
	AIMissionRunListRunsResponseStatusRunning   AIMissionRunListRunsResponseStatus = "running"
	AIMissionRunListRunsResponseStatusPaused    AIMissionRunListRunsResponseStatus = "paused"
	AIMissionRunListRunsResponseStatusSucceeded AIMissionRunListRunsResponseStatus = "succeeded"
	AIMissionRunListRunsResponseStatusFailed    AIMissionRunListRunsResponseStatus = "failed"
	AIMissionRunListRunsResponseStatusCancelled AIMissionRunListRunsResponseStatus = "cancelled"
)

type AIMissionRunPauseRunResponse struct {
	Data AIMissionRunPauseRunResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPauseRunResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPauseRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPauseRunResponseData struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        string         `json:"status,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Error         string         `json:"error"`
	FinishedAt    time.Time      `json:"finished_at" format:"date-time"`
	Input         map[string]any `json:"input"`
	Metadata      map[string]any `json:"metadata"`
	ResultPayload map[string]any `json:"result_payload"`
	ResultSummary string         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPauseRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPauseRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunResumeRunResponse struct {
	Data AIMissionRunResumeRunResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunResumeRunResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunResumeRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunResumeRunResponseData struct {
	MissionID string    `json:"mission_id,required" format:"uuid"`
	RunID     string    `json:"run_id,required" format:"uuid"`
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status        string         `json:"status,required"`
	UpdatedAt     time.Time      `json:"updated_at,required" format:"date-time"`
	Error         string         `json:"error"`
	FinishedAt    time.Time      `json:"finished_at" format:"date-time"`
	Input         map[string]any `json:"input"`
	Metadata      map[string]any `json:"metadata"`
	ResultPayload map[string]any `json:"result_payload"`
	ResultSummary string         `json:"result_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MissionID     respjson.Field
		RunID         respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpdatedAt     respjson.Field
		Error         respjson.Field
		FinishedAt    respjson.Field
		Input         respjson.Field
		Metadata      respjson.Field
		ResultPayload respjson.Field
		ResultSummary respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunResumeRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunResumeRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunNewParams struct {
	Input    map[string]any `json:"input,omitzero"`
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r AIMissionRunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunGetParams struct {
	MissionID string `path:"mission_id,required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunUpdateParams struct {
	MissionID     string            `path:"mission_id,required" format:"uuid" json:"-"`
	Error         param.Opt[string] `json:"error,omitzero"`
	ResultSummary param.Opt[string] `json:"result_summary,omitzero"`
	Metadata      map[string]any    `json:"metadata,omitzero"`
	ResultPayload map[string]any    `json:"result_payload,omitzero"`
	// Any of "pending", "running", "paused", "succeeded", "failed", "cancelled".
	Status AIMissionRunUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r AIMissionRunUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunUpdateParamsStatus string

const (
	AIMissionRunUpdateParamsStatusPending   AIMissionRunUpdateParamsStatus = "pending"
	AIMissionRunUpdateParamsStatusRunning   AIMissionRunUpdateParamsStatus = "running"
	AIMissionRunUpdateParamsStatusPaused    AIMissionRunUpdateParamsStatus = "paused"
	AIMissionRunUpdateParamsStatusSucceeded AIMissionRunUpdateParamsStatus = "succeeded"
	AIMissionRunUpdateParamsStatusFailed    AIMissionRunUpdateParamsStatus = "failed"
	AIMissionRunUpdateParamsStatusCancelled AIMissionRunUpdateParamsStatus = "cancelled"
)

type AIMissionRunListParams struct {
	// Page number (1-based)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	Status   param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMissionRunListParams]'s query parameters as `url.Values`.
func (r AIMissionRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMissionRunCancelRunParams struct {
	MissionID string `path:"mission_id,required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunListRunsParams struct {
	// Page number (1-based)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	Status   param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMissionRunListRunsParams]'s query parameters as
// `url.Values`.
func (r AIMissionRunListRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMissionRunPauseRunParams struct {
	MissionID string `path:"mission_id,required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunResumeRunParams struct {
	MissionID string `path:"mission_id,required" format:"uuid" json:"-"`
	paramObj
}
