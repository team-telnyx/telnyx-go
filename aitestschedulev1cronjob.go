// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Test agents, evaluator scenarios, metrics, and run results for the AI test
// framework.
//
// AITestScheduleV1CronJobService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestScheduleV1CronJobService] method instead.
type AITestScheduleV1CronJobService struct {
	Options []option.RequestOption
}

// NewAITestScheduleV1CronJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITestScheduleV1CronJobService(opts ...option.RequestOption) (r AITestScheduleV1CronJobService) {
	r = AITestScheduleV1CronJobService{}
	r.Options = opts
	return
}

// Returns the specified scheduled test job.
func (r *AITestScheduleV1CronJobService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ScheduledTestJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/schedules/v1/cron-jobs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the specified scheduled test job.
func (r *AITestScheduleV1CronJobService) Update(ctx context.Context, id int64, body AITestScheduleV1CronJobUpdateParams, opts ...option.RequestOption) (res *ScheduledTestJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/schedules/v1/cron-jobs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Deletes the specified scheduled test job.
func (r *AITestScheduleV1CronJobService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *AITestScheduleV1CronJobDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/schedules/v1/cron-jobs/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Creates a scheduled test job that runs the selected scenarios on a cron
// expression.
func (r *AITestScheduleV1CronJobService) CronJobs(ctx context.Context, body AITestScheduleV1CronJobCronJobsParams, opts ...option.RequestOption) (res *ScheduledTestJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/schedules/v1/cron-jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the scheduled test jobs in your organization, optionally filtered by
// test agent. A scheduled job runs selected scenarios on a cron expression.
func (r *AITestScheduleV1CronJobService) GetCronJobs(ctx context.Context, query AITestScheduleV1CronJobGetCronJobsParams, opts ...option.RequestOption) (res *[]ScheduledTestJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/schedules/v1/cron-jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ScheduledTestJob struct {
	ID int64 `json:"id" api:"required"`
	// ID of the test agent the job belongs to.
	Agent     int64     `json:"agent" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Cron expression controlling when the job runs, in the job's timezone.
	CrontabExpression string    `json:"crontab_expression" api:"required"`
	IsActive          bool      `json:"is_active" api:"required"`
	Name              string    `json:"name" api:"required"`
	UpdatedAt         time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Execution mode for the test runs, such as voice, text, or sip.
	ExecutionMode string `json:"execution_mode"`
	Frequency     int64  `json:"frequency"`
	// When the job sends a notification.
	//
	// Any of "never", "success", "failure", "both".
	NotifyOn      ScheduledTestJobNotifyOn `json:"notify_on"`
	Personalities []int64                  `json:"personalities"`
	Personality   int64                    `json:"personality" api:"nullable"`
	// Whether the job runs scenarios as text chat rather than voice calls.
	RunAsText bool `json:"run_as_text"`
	// Scenario IDs the job executes.
	ScenarioIDs   []int64                         `json:"scenario_ids"`
	ScenariosData []ScheduledTestJobScenariosData `json:"scenarios_data"`
	Tags          []string                        `json:"tags"`
	Timezone      string                          `json:"timezone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Agent             respjson.Field
		CreatedAt         respjson.Field
		CrontabExpression respjson.Field
		IsActive          respjson.Field
		Name              respjson.Field
		UpdatedAt         respjson.Field
		ExecutionMode     respjson.Field
		Frequency         respjson.Field
		NotifyOn          respjson.Field
		Personalities     respjson.Field
		Personality       respjson.Field
		RunAsText         respjson.Field
		ScenarioIDs       respjson.Field
		ScenariosData     respjson.Field
		Tags              respjson.Field
		Timezone          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledTestJob) RawJSON() string { return r.JSON.raw }
func (r *ScheduledTestJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the job sends a notification.
type ScheduledTestJobNotifyOn string

const (
	ScheduledTestJobNotifyOnNever   ScheduledTestJobNotifyOn = "never"
	ScheduledTestJobNotifyOnSuccess ScheduledTestJobNotifyOn = "success"
	ScheduledTestJobNotifyOnFailure ScheduledTestJobNotifyOn = "failure"
	ScheduledTestJobNotifyOnBoth    ScheduledTestJobNotifyOn = "both"
)

type ScheduledTestJobScenariosData struct {
	ID              int64  `json:"id"`
	Agent           int64  `json:"agent"`
	Name            string `json:"name"`
	Personality     int64  `json:"personality" api:"nullable"`
	PersonalityName string `json:"personality_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Agent           respjson.Field
		Name            respjson.Field
		Personality     respjson.Field
		PersonalityName respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduledTestJobScenariosData) RawJSON() string { return r.JSON.raw }
func (r *ScheduledTestJobScenariosData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestScheduleV1CronJobDeleteResponse = any

type AITestScheduleV1CronJobUpdateParams struct {
	Agent param.Opt[int64] `json:"agent,omitzero"`
	// Cron expression controlling when the job runs.
	CrontabExpression param.Opt[string] `json:"crontab_expression,omitzero"`
	// Execution mode for the test runs, such as voice, text, or sip.
	ExecutionMode param.Opt[string] `json:"execution_mode,omitzero"`
	Frequency     param.Opt[int64]  `json:"frequency,omitzero"`
	IsActive      param.Opt[bool]   `json:"is_active,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	RunAsText     param.Opt[bool]   `json:"run_as_text,omitzero"`
	Timezone      param.Opt[string] `json:"timezone,omitzero"`
	// Any of "never", "success", "failure", "both".
	NotifyOn      AITestScheduleV1CronJobUpdateParamsNotifyOn `json:"notify_on,omitzero"`
	Personalities []int64                                     `json:"personalities,omitzero"`
	Scenarios     []int64                                     `json:"scenarios,omitzero"`
	Tags          []string                                    `json:"tags,omitzero"`
	paramObj
}

func (r AITestScheduleV1CronJobUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestScheduleV1CronJobUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestScheduleV1CronJobUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestScheduleV1CronJobUpdateParamsNotifyOn string

const (
	AITestScheduleV1CronJobUpdateParamsNotifyOnNever   AITestScheduleV1CronJobUpdateParamsNotifyOn = "never"
	AITestScheduleV1CronJobUpdateParamsNotifyOnSuccess AITestScheduleV1CronJobUpdateParamsNotifyOn = "success"
	AITestScheduleV1CronJobUpdateParamsNotifyOnFailure AITestScheduleV1CronJobUpdateParamsNotifyOn = "failure"
	AITestScheduleV1CronJobUpdateParamsNotifyOnBoth    AITestScheduleV1CronJobUpdateParamsNotifyOn = "both"
)

type AITestScheduleV1CronJobCronJobsParams struct {
	// Cron expression controlling when the job runs.
	CrontabExpression string           `json:"crontab_expression" api:"required"`
	Agent             param.Opt[int64] `json:"agent,omitzero"`
	// Execution mode for the test runs, such as voice, text, or sip.
	ExecutionMode param.Opt[string] `json:"execution_mode,omitzero"`
	Frequency     param.Opt[int64]  `json:"frequency,omitzero"`
	IsActive      param.Opt[bool]   `json:"is_active,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	RunAsText     param.Opt[bool]   `json:"run_as_text,omitzero"`
	Timezone      param.Opt[string] `json:"timezone,omitzero"`
	// Any of "never", "success", "failure", "both".
	NotifyOn      AITestScheduleV1CronJobCronJobsParamsNotifyOn `json:"notify_on,omitzero"`
	Personalities []int64                                       `json:"personalities,omitzero"`
	Scenarios     []int64                                       `json:"scenarios,omitzero"`
	Tags          []string                                      `json:"tags,omitzero"`
	paramObj
}

func (r AITestScheduleV1CronJobCronJobsParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestScheduleV1CronJobCronJobsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestScheduleV1CronJobCronJobsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestScheduleV1CronJobCronJobsParamsNotifyOn string

const (
	AITestScheduleV1CronJobCronJobsParamsNotifyOnNever   AITestScheduleV1CronJobCronJobsParamsNotifyOn = "never"
	AITestScheduleV1CronJobCronJobsParamsNotifyOnSuccess AITestScheduleV1CronJobCronJobsParamsNotifyOn = "success"
	AITestScheduleV1CronJobCronJobsParamsNotifyOnFailure AITestScheduleV1CronJobCronJobsParamsNotifyOn = "failure"
	AITestScheduleV1CronJobCronJobsParamsNotifyOnBoth    AITestScheduleV1CronJobCronJobsParamsNotifyOn = "both"
)

type AITestScheduleV1CronJobGetCronJobsParams struct {
	// Filter jobs belonging to the specified test agent.
	AgentID param.Opt[int64] `query:"agent_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITestScheduleV1CronJobGetCronJobsParams]'s query
// parameters as `url.Values`.
func (r AITestScheduleV1CronJobGetCronJobsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
