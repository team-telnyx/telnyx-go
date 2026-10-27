// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
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
// AITestTestFrameworkV1ResultService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestTestFrameworkV1ResultService] method instead.
type AITestTestFrameworkV1ResultService struct {
	Options []option.RequestOption
}

// NewAITestTestFrameworkV1ResultService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAITestTestFrameworkV1ResultService(opts ...option.RequestOption) (r AITestTestFrameworkV1ResultService) {
	r = AITestTestFrameworkV1ResultService{}
	r.Options = opts
	return
}

// Returns the specified test run batch, optionally with per-run detail.
func (r *AITestTestFrameworkV1ResultService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *TestRunResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/results/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of test run batches in your organization, optionally
// filtered by test agent, search text, or scenarios.
func (r *AITestTestFrameworkV1ResultService) List(ctx context.Context, query AITestTestFrameworkV1ResultListParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1ResultListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the specified test run batch.
func (r *AITestTestFrameworkV1ResultService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *AITestTestFrameworkV1ResultDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/results/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Re-executes the scenarios of the specified test run batch and creates a new run
// batch.
func (r *AITestTestFrameworkV1ResultService) Rerun(ctx context.Context, id int64, opts ...option.RequestOption) (res *TestRunResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/results/%v/rerun", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type TestRunResult struct {
	ID        int64     `json:"id" api:"required"`
	Agent     int64     `json:"agent" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Run batch status: running, completed, failed, pending, in_progress, evaluating,
	// in_queue, timeout, cancelled, or scaling_up.
	Status             string                               `json:"status" api:"required"`
	UpdatedAt          time.Time                            `json:"updated_at" api:"required" format:"date-time"`
	AgentInbound       bool                                 `json:"agent_inbound"`
	AgentName          string                               `json:"agent_name"`
	CompletedRunsCount TestRunResultCompletedRunsCountUnion `json:"completed_runs_count"`
	FailedRunsCount    TestRunResultFailedRunsCountUnion    `json:"failed_runs_count"`
	// Whether this batch was started by a scheduled job.
	IsCronjob               bool                                      `json:"is_cronjob"`
	MetExpectedOutcomeCount TestRunResultMetExpectedOutcomeCountUnion `json:"met_expected_outcome_count"`
	// Whether the batch ran as text chat rather than voice calls.
	RunAsText bool   `json:"run_as_text"`
	RunType   string `json:"run_type" api:"nullable"`
	// Per-run detail keyed by run ID.
	Runs      map[string]TestRunResultRun `json:"runs" api:"nullable"`
	Scenarios []TestRunResultScenario     `json:"scenarios" api:"nullable"`
	// Share of runs judged successful.
	SuccessRate               TestRunResultSuccessRateUnion               `json:"success_rate"`
	SuccessRunsCount          TestRunResultSuccessRunsCountUnion          `json:"success_runs_count"`
	TotalExpectedOutcomeCount TestRunResultTotalExpectedOutcomeCountUnion `json:"total_expected_outcome_count"`
	TotalRunsCount            TestRunResultTotalRunsCountUnion            `json:"total_runs_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Agent                     respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Status                    respjson.Field
		UpdatedAt                 respjson.Field
		AgentInbound              respjson.Field
		AgentName                 respjson.Field
		CompletedRunsCount        respjson.Field
		FailedRunsCount           respjson.Field
		IsCronjob                 respjson.Field
		MetExpectedOutcomeCount   respjson.Field
		RunAsText                 respjson.Field
		RunType                   respjson.Field
		Runs                      respjson.Field
		Scenarios                 respjson.Field
		SuccessRate               respjson.Field
		SuccessRunsCount          respjson.Field
		TotalExpectedOutcomeCount respjson.Field
		TotalRunsCount            respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestRunResult) RawJSON() string { return r.JSON.raw }
func (r *TestRunResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultCompletedRunsCountUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type TestRunResultCompletedRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultCompletedRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultCompletedRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultCompletedRunsCountUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultCompletedRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultFailedRunsCountUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type TestRunResultFailedRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultFailedRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultFailedRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultFailedRunsCountUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultFailedRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultMetExpectedOutcomeCountUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type TestRunResultMetExpectedOutcomeCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultMetExpectedOutcomeCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultMetExpectedOutcomeCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultMetExpectedOutcomeCountUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultMetExpectedOutcomeCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestRunResultRun struct {
	ID               int64                     `json:"id"`
	Duration         string                    `json:"duration"`
	ErrorMessage     string                    `json:"error_message"`
	EvaluationStatus string                    `json:"evaluation_status"`
	ExecutedAt       string                    `json:"executed_at"`
	PersonalityName  string                    `json:"personality_name"`
	Scenario         TestRunResultRunsScenario `json:"scenario"`
	ScenarioName     string                    `json:"scenario_name"`
	Status           string                    `json:"status"`
	Success          bool                      `json:"success" api:"nullable"`
	TestProfileName  string                    `json:"test_profile_name"`
	Timestamp        string                    `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Duration         respjson.Field
		ErrorMessage     respjson.Field
		EvaluationStatus respjson.Field
		ExecutedAt       respjson.Field
		PersonalityName  respjson.Field
		Scenario         respjson.Field
		ScenarioName     respjson.Field
		Status           respjson.Field
		Success          respjson.Field
		TestProfileName  respjson.Field
		Timestamp        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestRunResultRun) RawJSON() string { return r.JSON.raw }
func (r *TestRunResultRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestRunResultRunsScenario struct {
	ID                    int64  `json:"id"`
	ExpectedOutcomePrompt string `json:"expected_outcome_prompt"`
	Name                  string `json:"name"`
	PersonalityName       string `json:"personality_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ExpectedOutcomePrompt respjson.Field
		Name                  respjson.Field
		PersonalityName       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestRunResultRunsScenario) RawJSON() string { return r.JSON.raw }
func (r *TestRunResultRunsScenario) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestRunResultScenario struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestRunResultScenario) RawJSON() string { return r.JSON.raw }
func (r *TestRunResultScenario) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultSuccessRateUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TestRunResultSuccessRateUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultSuccessRateUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultSuccessRateUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultSuccessRateUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultSuccessRateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultSuccessRunsCountUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type TestRunResultSuccessRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultSuccessRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultSuccessRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultSuccessRunsCountUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultSuccessRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultTotalExpectedOutcomeCountUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type TestRunResultTotalExpectedOutcomeCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultTotalExpectedOutcomeCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultTotalExpectedOutcomeCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultTotalExpectedOutcomeCountUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultTotalExpectedOutcomeCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestRunResultTotalRunsCountUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type TestRunResultTotalRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TestRunResultTotalRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestRunResultTotalRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestRunResultTotalRunsCountUnion) RawJSON() string { return u.JSON.raw }

func (r *TestRunResultTotalRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ResultListResponse struct {
	Next     string                                          `json:"next" api:"required" format:"uri"`
	Previous string                                          `json:"previous" api:"required" format:"uri"`
	Results  []AITestTestFrameworkV1ResultListResponseResult `json:"results" api:"required"`
	Count    int64                                           `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ResultListResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ResultListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ResultListResponseResult struct {
	ID        int64     `json:"id" api:"required"`
	Agent     int64     `json:"agent" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Run batch status: running, completed, failed, pending, in_progress, evaluating,
	// in_queue, timeout, cancelled, or scaling_up.
	Status             string                                                                `json:"status" api:"required"`
	UpdatedAt          time.Time                                                             `json:"updated_at" api:"required" format:"date-time"`
	AgentInbound       bool                                                                  `json:"agent_inbound"`
	AgentName          string                                                                `json:"agent_name"`
	CompletedRunsCount AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion `json:"completed_runs_count"`
	FailedRunsCount    AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion    `json:"failed_runs_count"`
	// Whether this batch was started by a scheduled job.
	IsCronjob               bool                                                                       `json:"is_cronjob"`
	MetExpectedOutcomeCount AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion `json:"met_expected_outcome_count"`
	// Whether the batch ran as text chat rather than voice calls.
	RunAsText bool   `json:"run_as_text"`
	RunType   string `json:"run_type" api:"nullable"`
	// Run summaries in this batch. Retrieve the batch for detail keyed by run ID.
	Runs      []AITestTestFrameworkV1ResultListResponseResultsRun      `json:"runs" api:"nullable"`
	Scenarios []AITestTestFrameworkV1ResultListResponseResultsScenario `json:"scenarios" api:"nullable"`
	// Share of runs judged successful.
	SuccessRate               AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion               `json:"success_rate"`
	SuccessRunsCount          AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion          `json:"success_runs_count"`
	TotalExpectedOutcomeCount AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion `json:"total_expected_outcome_count"`
	TotalRunsCount            AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion            `json:"total_runs_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Agent                     respjson.Field
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Status                    respjson.Field
		UpdatedAt                 respjson.Field
		AgentInbound              respjson.Field
		AgentName                 respjson.Field
		CompletedRunsCount        respjson.Field
		FailedRunsCount           respjson.Field
		IsCronjob                 respjson.Field
		MetExpectedOutcomeCount   respjson.Field
		RunAsText                 respjson.Field
		RunType                   respjson.Field
		Runs                      respjson.Field
		Scenarios                 respjson.Field
		SuccessRate               respjson.Field
		SuccessRunsCount          respjson.Field
		TotalExpectedOutcomeCount respjson.Field
		TotalRunsCount            respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ResultListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ResultListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion contains
// all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsCompletedRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion contains all
// possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsFailedRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsMetExpectedOutcomeCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ResultListResponseResultsRun struct {
	ID            int64  `json:"id"`
	InboundNumber string `json:"inbound_number" api:"nullable"`
	Number        string `json:"number" api:"nullable"`
	Scenario      int64  `json:"scenario"`
	ScenarioName  string `json:"scenario_name"`
	Status        string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		InboundNumber respjson.Field
		Number        respjson.Field
		Scenario      respjson.Field
		ScenarioName  respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ResultListResponseResultsRun) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ResultListResponseResultsRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ResultListResponseResultsScenario struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ResultListResponseResultsScenario) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ResultListResponseResultsScenario) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion contains all
// possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsSuccessRateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion contains all
// possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsSuccessRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsTotalExpectedOutcomeCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion contains all
// possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1ResultListResponseResultsTotalRunsCountUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ResultDeleteResponse = any

type AITestTestFrameworkV1ResultListParams struct {
	// Filter results belonging to the specified test agent.
	AgentID param.Opt[int64] `query:"agent_id,omitzero" json:"-"`
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter results by a comma-separated list of scenario IDs.
	Scenarios param.Opt[string] `query:"scenarios,omitzero" json:"-"`
	// Filter results by a search string.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITestTestFrameworkV1ResultListParams]'s query parameters
// as `url.Values`.
func (r AITestTestFrameworkV1ResultListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
