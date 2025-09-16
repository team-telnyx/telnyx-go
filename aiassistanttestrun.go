// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// AIAssistantTestRunService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantTestRunService] method instead.
type AIAssistantTestRunService struct {
	Options []option.RequestOption
}

// NewAIAssistantTestRunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIAssistantTestRunService(opts ...option.RequestOption) (r AIAssistantTestRunService) {
	r = AIAssistantTestRunService{}
	r.Options = opts
	return
}

// Retrieves detailed information about a specific test run execution
func (r *AIAssistantTestRunService) Get(ctx context.Context, runID string, query AIAssistantTestRunGetParams, opts ...option.RequestOption) (res *TestRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.TestID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/%s/runs/%s", query.TestID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves paginated execution history for a specific assistant test with
// filtering options
func (r *AIAssistantTestRunService) List(ctx context.Context, testID string, query AIAssistantTestRunListParams, opts ...option.RequestOption) (res *PaginatedTestRunList, err error) {
	opts = append(r.Options[:], opts...)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/%s/runs", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiates immediate execution of a specific assistant test
func (r *AIAssistantTestRunService) Trigger(ctx context.Context, testID string, body AIAssistantTestRunTriggerParams, opts ...option.RequestOption) (res *TestRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/%s/runs", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response model containing test run execution details and results.
//
// Provides comprehensive information about a test execution including status,
// timing, logs, and detailed evaluation results.
type TestRunResponse struct {
	// Timestamp when the test run was created and queued.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Unique identifier for this specific test run execution.
	RunID string `json:"run_id,required" format:"uuid"`
	// Represents the lifecycle of a test:
	//
	// - 'pending': Test is waiting to be executed.
	// - 'starting': Test execution is initializing.
	// - 'running': Test is currently executing.
	// - 'passed': Test completed successfully.
	// - 'failed': Test executed but did not pass.
	// - 'error': An error occurred during test execution.
	//
	// Any of "pending", "starting", "running", "passed", "failed", "error".
	Status TestStatus `json:"status,required"`
	// Identifier of the assistant test that was executed.
	TestID string `json:"test_id,required" format:"uuid"`
	// How this test run was initiated (manual, scheduled, or API).
	TriggeredBy string `json:"triggered_by,required"`
	// Timestamp when the test run finished execution.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// Identifier of the conversation created during test execution.
	ConversationID string `json:"conversation_id"`
	// Identifier for conversation analysis and insights data.
	ConversationInsightsID string `json:"conversation_insights_id"`
	// Detailed evaluation results for each rubric criteria. Name is name of the
	// criteria from the rubric and status is the result of the evaluation. This list
	// will have a result for every criteria in the rubric section.
	DetailStatus []TestRunResponseDetailStatus `json:"detail_status"`
	// Detailed execution logs and debug information.
	Logs string `json:"logs"`
	// Identifier linking this run to a test suite execution batch.
	TestSuiteRunID string `json:"test_suite_run_id" format:"uuid"`
	// Timestamp of the last update to this test run.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt              respjson.Field
		RunID                  respjson.Field
		Status                 respjson.Field
		TestID                 respjson.Field
		TriggeredBy            respjson.Field
		CompletedAt            respjson.Field
		ConversationID         respjson.Field
		ConversationInsightsID respjson.Field
		DetailStatus           respjson.Field
		Logs                   respjson.Field
		TestSuiteRunID         respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestRunResponse) RawJSON() string { return r.JSON.raw }
func (r *TestRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestRunResponseDetailStatus struct {
	Name string `json:"name,required"`
	// Represents the lifecycle of a test:
	//
	// - 'pending': Test is waiting to be executed.
	// - 'starting': Test execution is initializing.
	// - 'running': Test is currently executing.
	// - 'passed': Test completed successfully.
	// - 'failed': Test executed but did not pass.
	// - 'error': An error occurred during test execution.
	//
	// Any of "pending", "starting", "running", "passed", "failed", "error".
	Status TestStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestRunResponseDetailStatus) RawJSON() string { return r.JSON.raw }
func (r *TestRunResponseDetailStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the lifecycle of a test:
//
// - 'pending': Test is waiting to be executed.
// - 'starting': Test execution is initializing.
// - 'running': Test is currently executing.
// - 'passed': Test completed successfully.
// - 'failed': Test executed but did not pass.
// - 'error': An error occurred during test execution.
type TestStatus string

const (
	TestStatusPending  TestStatus = "pending"
	TestStatusStarting TestStatus = "starting"
	TestStatusRunning  TestStatus = "running"
	TestStatusPassed   TestStatus = "passed"
	TestStatusFailed   TestStatus = "failed"
	TestStatusError    TestStatus = "error"
)

type AIAssistantTestRunGetParams struct {
	TestID string `path:"test_id,required" json:"-"`
	paramObj
}

type AIAssistantTestRunListParams struct {
	// Filter runs by execution status (pending, running, completed, failed, timeout)
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page AIAssistantTestRunListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantTestRunListParams]'s query parameters as
// `url.Values`.
func (r AIAssistantTestRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type AIAssistantTestRunListParamsPage struct {
	// Page number to retrieve (1-based indexing)
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Number of test runs to return per page (1-100)
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantTestRunListParamsPage]'s query parameters as
// `url.Values`.
func (r AIAssistantTestRunListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAssistantTestRunTriggerParams struct {
	// Optional assistant version ID to use for this test run. If provided, the version
	// must exist or a 400 error will be returned. If not provided, test will run on
	// main version
	DestinationVersionID param.Opt[string] `json:"destination_version_id,omitzero"`
	paramObj
}

func (r AIAssistantTestRunTriggerParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTestRunTriggerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTestRunTriggerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
