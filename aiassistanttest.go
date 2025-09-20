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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// AIAssistantTestService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantTestService] method instead.
type AIAssistantTestService struct {
	Options    []option.RequestOption
	TestSuites AIAssistantTestTestSuiteService
	Runs       AIAssistantTestRunService
}

// NewAIAssistantTestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAssistantTestService(opts ...option.RequestOption) (r AIAssistantTestService) {
	r = AIAssistantTestService{}
	r.Options = opts
	r.TestSuites = NewAIAssistantTestTestSuiteService(opts...)
	r.Runs = NewAIAssistantTestRunService(opts...)
	return
}

// Creates a comprehensive test configuration for evaluating AI assistant
// performance
func (r *AIAssistantTestService) New(ctx context.Context, body AIAssistantTestNewParams, opts ...option.RequestOption) (res *AssistantTest, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/tests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information about a specific assistant test
func (r *AIAssistantTestService) Get(ctx context.Context, testID string, opts ...option.RequestOption) (res *AssistantTest, err error) {
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/%s", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing assistant test configuration with new settings
func (r *AIAssistantTestService) Update(ctx context.Context, testID string, body AIAssistantTestUpdateParams, opts ...option.RequestOption) (res *AssistantTest, err error) {
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/%s", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of assistant tests with optional filtering
// capabilities
func (r *AIAssistantTestService) List(ctx context.Context, query AIAssistantTestListParams, opts ...option.RequestOption) (res *AIAssistantTestListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/tests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Permanently removes an assistant test and all associated data
func (r *AIAssistantTestService) Delete(ctx context.Context, testID string, opts ...option.RequestOption) (res *AIAssistantTestDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/tests/%s", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response model containing complete assistant test information.
//
// Returns all test configuration details including evaluation criteria,
// scheduling, and metadata. Used when retrieving individual tests or after
// creating/updating tests.
type AssistantTest struct {
	// Timestamp when the test was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Human-readable name of the test.
	Name string `json:"name,required"`
	// Evaluation criteria used to assess test performance.
	Rubric []AssistantTestRubric `json:"rubric,required"`
	// Communication channel used for test execution.
	//
	// Any of "phone_call", "web_call", "sms_chat", "web_chat".
	TelnyxConversationChannel TelnyxConversationChannel `json:"telnyx_conversation_channel,required"`
	// Unique identifier for the assistant test.
	TestID string `json:"test_id,required" format:"uuid"`
	// Detailed description of the test's purpose and scope.
	Description string `json:"description"`
	// Target destination for test conversations.
	Destination string `json:"destination"`
	// Detailed test scenario instructions and objectives.
	Instructions string `json:"instructions"`
	// Maximum allowed duration for test execution in seconds.
	MaxDurationSeconds int64 `json:"max_duration_seconds"`
	// Test suite grouping for organizational purposes.
	TestSuite string `json:"test_suite"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt                 respjson.Field
		Name                      respjson.Field
		Rubric                    respjson.Field
		TelnyxConversationChannel respjson.Field
		TestID                    respjson.Field
		Description               respjson.Field
		Destination               respjson.Field
		Instructions              respjson.Field
		MaxDurationSeconds        respjson.Field
		TestSuite                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantTest) RawJSON() string { return r.JSON.raw }
func (r *AssistantTest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantTestRubric struct {
	// Specific guidance on how to assess the assistant’s performance for this rubric
	// item.
	Criteria string `json:"criteria,required"`
	// Label for the evaluation criterion, e.g., Empathy, Accuracy, Clarity.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Criteria    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantTestRubric) RawJSON() string { return r.JSON.raw }
func (r *AssistantTestRubric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelnyxConversationChannel string

const (
	TelnyxConversationChannelPhoneCall TelnyxConversationChannel = "phone_call"
	TelnyxConversationChannelWebCall   TelnyxConversationChannel = "web_call"
	TelnyxConversationChannelSMSChat   TelnyxConversationChannel = "sms_chat"
	TelnyxConversationChannelWebChat   TelnyxConversationChannel = "web_chat"
)

// Paginated list of assistant tests with metadata.
//
// Returns a subset of tests based on pagination parameters along with metadata for
// implementing pagination controls in the UI.
type AIAssistantTestListResponse struct {
	// Array of assistant test objects for the current page.
	Data []AssistantTest `json:"data,required"`
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
func (r AIAssistantTestListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantTestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTestDeleteResponse = any

type AIAssistantTestNewParams struct {
	// The target destination for the test conversation. Format depends on the channel:
	// phone number for SMS/voice, webhook URL for web chat, etc.
	Destination string `json:"destination,required"`
	// Detailed instructions that define the test scenario and what the assistant
	// should accomplish. This guides the test execution and evaluation.
	Instructions string `json:"instructions,required"`
	// A descriptive name for the assistant test. This will be used to identify the
	// test in the UI and reports.
	Name string `json:"name,required"`
	// Evaluation criteria used to assess the assistant's performance. Each rubric item
	// contains a name and specific criteria for evaluation.
	Rubric []AIAssistantTestNewParamsRubric `json:"rubric,omitzero,required"`
	// Optional detailed description of what this test evaluates and its purpose. Helps
	// team members understand the test's objectives.
	Description param.Opt[string] `json:"description,omitzero"`
	// Maximum duration in seconds that the test conversation should run before timing
	// out. If not specified, uses system default timeout.
	MaxDurationSeconds param.Opt[int64] `json:"max_duration_seconds,omitzero"`
	// Optional test suite name to group related tests together. Useful for organizing
	// tests by feature, team, or release cycle.
	TestSuite param.Opt[string] `json:"test_suite,omitzero"`
	// The communication channel through which the test will be conducted. Determines
	// how the assistant will receive and respond to test messages.
	//
	// Any of "phone_call", "web_call", "sms_chat", "web_chat".
	TelnyxConversationChannel TelnyxConversationChannel `json:"telnyx_conversation_channel,omitzero"`
	paramObj
}

func (r AIAssistantTestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTestNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Criteria, Name are required.
type AIAssistantTestNewParamsRubric struct {
	// Specific guidance on how to assess the assistant’s performance for this rubric
	// item.
	Criteria string `json:"criteria,required"`
	// Label for the evaluation criterion, e.g., Empathy, Accuracy, Clarity.
	Name string `json:"name,required"`
	paramObj
}

func (r AIAssistantTestNewParamsRubric) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTestNewParamsRubric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTestNewParamsRubric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTestUpdateParams struct {
	// Updated description of the test's purpose and evaluation criteria.
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated target destination for test conversations.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// Updated test scenario instructions and objectives.
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// Updated maximum test duration in seconds.
	MaxDurationSeconds param.Opt[int64] `json:"max_duration_seconds,omitzero"`
	// Updated name for the assistant test. Must be unique and descriptive.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated test suite assignment for better organization.
	TestSuite param.Opt[string] `json:"test_suite,omitzero"`
	// Updated evaluation criteria for assessing assistant performance.
	Rubric []AIAssistantTestUpdateParamsRubric `json:"rubric,omitzero"`
	// Updated communication channel for the test execution.
	//
	// Any of "phone_call", "web_call", "sms_chat", "web_chat".
	TelnyxConversationChannel TelnyxConversationChannel `json:"telnyx_conversation_channel,omitzero"`
	paramObj
}

func (r AIAssistantTestUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTestUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTestUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Criteria, Name are required.
type AIAssistantTestUpdateParamsRubric struct {
	// Specific guidance on how to assess the assistant’s performance for this rubric
	// item.
	Criteria string `json:"criteria,required"`
	// Label for the evaluation criterion, e.g., Empathy, Accuracy, Clarity.
	Name string `json:"name,required"`
	paramObj
}

func (r AIAssistantTestUpdateParamsRubric) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantTestUpdateParamsRubric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantTestUpdateParamsRubric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantTestListParams struct {
	// Filter tests by destination (phone number, webhook URL, etc.)
	Destination param.Opt[string] `query:"destination,omitzero" json:"-"`
	// Filter tests by communication channel (e.g., 'web_chat', 'sms')
	TelnyxConversationChannel param.Opt[string] `query:"telnyx_conversation_channel,omitzero" json:"-"`
	// Filter tests by test suite name
	TestSuite param.Opt[string] `query:"test_suite,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[size],
	// page[number]
	Page AIAssistantTestListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantTestListParams]'s query parameters as
// `url.Values`.
func (r AIAssistantTestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[size],
// page[number]
type AIAssistantTestListParamsPage struct {
	// Page number to retrieve (1-based indexing)
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Number of tests to return per page (1-100)
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantTestListParamsPage]'s query parameters as
// `url.Values`.
func (r AIAssistantTestListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
