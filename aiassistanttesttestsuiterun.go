// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/apiquery"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// AIAssistantTestTestSuiteRunService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantTestTestSuiteRunService] method instead.
type AIAssistantTestTestSuiteRunService struct {
	Options []option.RequestOption
}

// NewAIAssistantTestTestSuiteRunService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIAssistantTestTestSuiteRunService(opts ...option.RequestOption) (r AIAssistantTestTestSuiteRunService) {
	r = AIAssistantTestTestSuiteRunService{}
	r.Options = opts
	return
}

// Retrieves paginated history of test runs for a specific test suite with
// filtering options
func (r *AIAssistantTestTestSuiteRunService) List(ctx context.Context, suiteName string, query AIAssistantTestTestSuiteRunListParams, opts ...option.RequestOption) (res *PaginatedTestRunList, err error) {
	opts = append(r.Options[:], opts...)
	if suiteName == "" {
		err = errors.New("missing required suite_name parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/test-suites/%s/runs", suiteName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Executes all tests within a specific test suite as a batch operation
func (r *AIAssistantTestTestSuiteRunService) Trigger(ctx context.Context, suiteName string, body AIAssistantTestTestSuiteRunTriggerParams, opts ...option.RequestOption) (res *[]TestRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	if suiteName == "" {
		err = errors.New("missing required suite_name parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/test-suites/%s/runs", suiteName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Meta struct {
	PageNumber   int64 `json:"page_number,required"`
	PageSize     int64 `json:"page_size,required"`
	TotalPages   int64 `json:"total_pages,required"`
	TotalResults int64 `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Meta) RawJSON() string { return r.JSON.raw }
func (r *Meta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated list of test runs with metadata.
//
// Returns test run execution results with pagination support for handling large
// numbers of test executions.
type PaginatedTestRunList struct {
	// Array of test run objects for the current page.
	Data []TestRunResponse `json:"data,required"`
	// Pagination metadata including total counts and current page info.
	Meta Meta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaginatedTestRunList) RawJSON() string { return r.JSON.raw }
func (r *PaginatedTestRunList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTestTestSuiteRunListParams struct {
	// Filter runs by execution status (pending, running, completed, failed, timeout)
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	// Filter runs by specific suite execution batch ID
	TestSuiteRunID param.Opt[string] `query:"test_suite_run_id,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page AIAssistantTestTestSuiteRunListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantTestTestSuiteRunListParams]'s query parameters
// as `url.Values`.
func (r AIAssistantTestTestSuiteRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type AIAssistantTestTestSuiteRunListParamsPage struct {
	// Page number to retrieve (1-based indexing)
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Number of test runs to return per page (1-100)
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantTestTestSuiteRunListParamsPage]'s query
// parameters as `url.Values`.
func (r AIAssistantTestTestSuiteRunListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAssistantTestTestSuiteRunTriggerParams struct {
	// Optional assistant version ID to use for all test runs in this suite. If
	// provided, the version must exist or a 400 error will be returned. If not
	// provided, test will run on main version
	DestinationVersionID param.Opt[string] `json:"destination_version_id,omitzero"`
	paramObj
}

func (r AIAssistantTestTestSuiteRunTriggerParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTestTestSuiteRunTriggerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTestTestSuiteRunTriggerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
