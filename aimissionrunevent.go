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

// AIMissionRunEventService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionRunEventService] method instead.
type AIMissionRunEventService struct {
	Options []option.RequestOption
}

// NewAIMissionRunEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMissionRunEventService(opts ...option.RequestOption) (r AIMissionRunEventService) {
	r = AIMissionRunEventService{}
	r.Options = opts
	return
}

// List events for a run (paginated)
func (r *AIMissionRunEventService) List(ctx context.Context, runID string, params AIMissionRunEventListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[EventData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/events", params.MissionID, runID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List events for a run (paginated)
func (r *AIMissionRunEventService) ListAutoPaging(ctx context.Context, runID string, params AIMissionRunEventListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[EventData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, runID, params, opts...))
}

// Get details of a specific event
func (r *AIMissionRunEventService) GetEventDetails(ctx context.Context, eventID string, query AIMissionRunEventGetEventDetailsParams, opts ...option.RequestOption) (res *AIMissionRunEventGetEventDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if query.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/events/%s", query.MissionID, query.RunID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Log an event for a run
func (r *AIMissionRunEventService) Log(ctx context.Context, runID string, params AIMissionRunEventLogParams, opts ...option.RequestOption) (res *AIMissionRunEventLogResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/missions/%s/runs/%s/events", params.MissionID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type EventData struct {
	EventID   string    `json:"event_id" api:"required"`
	RunID     string    `json:"run_id" api:"required"`
	Summary   string    `json:"summary" api:"required"`
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Any of "status_change", "step_started", "step_completed", "step_failed",
	// "tool_call", "tool_result", "message", "error", "custom".
	Type           EventDataType  `json:"type" api:"required"`
	AgentID        string         `json:"agent_id"`
	IdempotencyKey string         `json:"idempotency_key"`
	Payload        map[string]any `json:"payload"`
	StepID         string         `json:"step_id"`
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
func (r EventData) RawJSON() string { return r.JSON.raw }
func (r *EventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventDataType string

const (
	EventDataTypeStatusChange  EventDataType = "status_change"
	EventDataTypeStepStarted   EventDataType = "step_started"
	EventDataTypeStepCompleted EventDataType = "step_completed"
	EventDataTypeStepFailed    EventDataType = "step_failed"
	EventDataTypeToolCall      EventDataType = "tool_call"
	EventDataTypeToolResult    EventDataType = "tool_result"
	EventDataTypeMessage       EventDataType = "message"
	EventDataTypeError         EventDataType = "error"
	EventDataTypeCustom        EventDataType = "custom"
)

type AIMissionRunEventGetEventDetailsResponse struct {
	Data EventData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunEventGetEventDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunEventGetEventDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunEventLogResponse struct {
	Data EventData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIMissionRunEventLogResponse) RawJSON() string { return r.JSON.raw }
func (r *AIMissionRunEventLogResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunEventListParams struct {
	MissionID string            `path:"mission_id" api:"required" format:"uuid" json:"-"`
	AgentID   param.Opt[string] `query:"agent_id,omitzero" json:"-"`
	// Page number (1-based)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	StepID   param.Opt[string] `query:"step_id,omitzero" json:"-"`
	Type     param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIMissionRunEventListParams]'s query parameters as
// `url.Values`.
func (r AIMissionRunEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIMissionRunEventGetEventDetailsParams struct {
	MissionID string `path:"mission_id" api:"required" format:"uuid" json:"-"`
	RunID     string `path:"run_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AIMissionRunEventLogParams struct {
	MissionID string `path:"mission_id" api:"required" format:"uuid" json:"-"`
	Summary   string `json:"summary" api:"required"`
	// Any of "status_change", "step_started", "step_completed", "step_failed",
	// "tool_call", "tool_result", "message", "error", "custom".
	Type    AIMissionRunEventLogParamsType `json:"type,omitzero" api:"required"`
	AgentID param.Opt[string]              `json:"agent_id,omitzero"`
	// Prevents duplicate events on retry
	IdempotencyKey param.Opt[string] `json:"idempotency_key,omitzero"`
	StepID         param.Opt[string] `json:"step_id,omitzero"`
	Payload        map[string]any    `json:"payload,omitzero"`
	paramObj
}

func (r AIMissionRunEventLogParams) MarshalJSON() (data []byte, err error) {
	type shadow AIMissionRunEventLogParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIMissionRunEventLogParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIMissionRunEventLogParamsType string

const (
	AIMissionRunEventLogParamsTypeStatusChange  AIMissionRunEventLogParamsType = "status_change"
	AIMissionRunEventLogParamsTypeStepStarted   AIMissionRunEventLogParamsType = "step_started"
	AIMissionRunEventLogParamsTypeStepCompleted AIMissionRunEventLogParamsType = "step_completed"
	AIMissionRunEventLogParamsTypeStepFailed    AIMissionRunEventLogParamsType = "step_failed"
	AIMissionRunEventLogParamsTypeToolCall      AIMissionRunEventLogParamsType = "tool_call"
	AIMissionRunEventLogParamsTypeToolResult    AIMissionRunEventLogParamsType = "tool_result"
	AIMissionRunEventLogParamsTypeMessage       AIMissionRunEventLogParamsType = "message"
	AIMissionRunEventLogParamsTypeError         AIMissionRunEventLogParamsType = "error"
	AIMissionRunEventLogParamsTypeCustom        AIMissionRunEventLogParamsType = "custom"
)
