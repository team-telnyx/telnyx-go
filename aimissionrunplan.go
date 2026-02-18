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

// AIMissionRunPlanService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionRunPlanService] method instead.
type AIMissionRunPlanService struct {
	Options []option.RequestOption
}

// NewAIMissionRunPlanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMissionRunPlanService(opts ...option.RequestOption) (r AIMissionRunPlanService) {
	r = AIMissionRunPlanService{}
	r.Options = opts
	return
}

// Create the initial plan for a run
func (r *AIMissionRunPlanService) New(ctx context.Context, runID string, params AIMissionRunPlanNewParams, opts ...option.RequestOption) (res *AIMissionRunPlanNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/plan", params.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get the plan (all steps) for a run
func (r *AIMissionRunPlanService) Get(ctx context.Context, runID string, query AIMissionRunPlanGetParams, opts ...option.RequestOption) (res *AIMissionRunPlanGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/plan", query.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add one or more steps to an existing plan
func (r *AIMissionRunPlanService) AddStepsToPlan(ctx context.Context, runID string, params AIMissionRunPlanAddStepsToPlanParams, opts ...option.RequestOption) (res *AIMissionRunPlanAddStepsToPlanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/plan/steps", params.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details of a specific plan step
func (r *AIMissionRunPlanService) GetStepDetails(ctx context.Context, stepID string, query AIMissionRunPlanGetStepDetailsParams, opts ...option.RequestOption) (res *AIMissionRunPlanGetStepDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if query.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	if stepID == "" {
		err = errors.New("missing required step_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/plan/steps/%s", query.MissionID, query.RunID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the status of a plan step
func (r *AIMissionRunPlanService) UpdateStep(ctx context.Context, stepID string, params AIMissionRunPlanUpdateStepParams, opts ...option.RequestOption) (res *AIMissionRunPlanUpdateStepResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if params.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	if stepID == "" {
		err = errors.New("missing required step_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/plan/steps/%s", params.MissionID, params.RunID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type PlanStepData struct {
	Description string `json:"description,required"`
	RunID       string `json:"run_id,required" format:"uuid"`
	Sequence    int64  `json:"sequence,required"`
	// Any of "pending", "in_progress", "completed", "skipped", "failed".
	Status       PlanStepDataStatus `json:"status,required"`
	StepID       string             `json:"step_id,required"`
	CompletedAt  time.Time          `json:"completed_at" format:"date-time"`
	Metadata     map[string]any     `json:"metadata"`
	ParentStepID string             `json:"parent_step_id"`
	StartedAt    time.Time          `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		RunID        respjson.Field
		Sequence     respjson.Field
		Status       respjson.Field
		StepID       respjson.Field
		CompletedAt  respjson.Field
		Metadata     respjson.Field
		ParentStepID respjson.Field
		StartedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlanStepData) RawJSON() string { return r.JSON.raw }
func (r *PlanStepData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlanStepDataStatus string

const (
	PlanStepDataStatusPending    PlanStepDataStatus = "pending"
	PlanStepDataStatusInProgress PlanStepDataStatus = "in_progress"
	PlanStepDataStatusCompleted  PlanStepDataStatus = "completed"
	PlanStepDataStatusSkipped    PlanStepDataStatus = "skipped"
	PlanStepDataStatusFailed     PlanStepDataStatus = "failed"
)

type AIMissionRunPlanNewResponse struct {
	Data []PlanStepData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPlanNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPlanNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanGetResponse struct {
	Data []PlanStepData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPlanGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPlanGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanAddStepsToPlanResponse struct {
	Data []PlanStepData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPlanAddStepsToPlanResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPlanAddStepsToPlanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanGetStepDetailsResponse struct {
	Data PlanStepData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPlanGetStepDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPlanGetStepDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanUpdateStepResponse struct {
	Data PlanStepData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunPlanUpdateStepResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunPlanUpdateStepResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanNewParams struct {
	MissionID string                          `path:"mission_id,required" format:"uuid" json:"-"`
	Steps     []AIMissionRunPlanNewParamsStep `json:"steps,omitzero,required"`
	paramObj
}

func (r AIMissionRunPlanNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunPlanNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunPlanNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Sequence, StepID are required.
type AIMissionRunPlanNewParamsStep struct {
	Description  string            `json:"description,required"`
	Sequence     int64             `json:"sequence,required"`
	StepID       string            `json:"step_id,required"`
	ParentStepID param.Opt[string] `json:"parent_step_id,omitzero"`
	Metadata     map[string]any    `json:"metadata,omitzero"`
	paramObj
}

func (r AIMissionRunPlanNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunPlanNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunPlanNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanGetParams struct {
	MissionID string `path:"mission_id,required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunPlanAddStepsToPlanParams struct {
	MissionID string                                     `path:"mission_id,required" format:"uuid" json:"-"`
	Steps     []AIMissionRunPlanAddStepsToPlanParamsStep `json:"steps,omitzero,required"`
	paramObj
}

func (r AIMissionRunPlanAddStepsToPlanParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunPlanAddStepsToPlanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunPlanAddStepsToPlanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Sequence, StepID are required.
type AIMissionRunPlanAddStepsToPlanParamsStep struct {
	Description  string            `json:"description,required"`
	Sequence     int64             `json:"sequence,required"`
	StepID       string            `json:"step_id,required"`
	ParentStepID param.Opt[string] `json:"parent_step_id,omitzero"`
	Metadata     map[string]any    `json:"metadata,omitzero"`
	paramObj
}

func (r AIMissionRunPlanAddStepsToPlanParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunPlanAddStepsToPlanParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunPlanAddStepsToPlanParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanGetStepDetailsParams struct {
	MissionID string `path:"mission_id,required" format:"uuid" json:"-"`
	RunID     string `path:"run_id,required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunPlanUpdateStepParams struct {
	MissionID string         `path:"mission_id,required" format:"uuid" json:"-"`
	RunID     string         `path:"run_id,required" format:"uuid" json:"-"`
	Metadata  map[string]any `json:"metadata,omitzero"`
	// Any of "pending", "in_progress", "completed", "skipped", "failed".
	Status AIMissionRunPlanUpdateStepParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r AIMissionRunPlanUpdateStepParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunPlanUpdateStepParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunPlanUpdateStepParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunPlanUpdateStepParamsStatus string

const (
	AIMissionRunPlanUpdateStepParamsStatusPending    AIMissionRunPlanUpdateStepParamsStatus = "pending"
	AIMissionRunPlanUpdateStepParamsStatusInProgress AIMissionRunPlanUpdateStepParamsStatus = "in_progress"
	AIMissionRunPlanUpdateStepParamsStatusCompleted  AIMissionRunPlanUpdateStepParamsStatus = "completed"
	AIMissionRunPlanUpdateStepParamsStatusSkipped    AIMissionRunPlanUpdateStepParamsStatus = "skipped"
	AIMissionRunPlanUpdateStepParamsStatusFailed     AIMissionRunPlanUpdateStepParamsStatus = "failed"
)
