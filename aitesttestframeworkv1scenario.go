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
// AITestTestFrameworkV1ScenarioService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestTestFrameworkV1ScenarioService] method instead.
type AITestTestFrameworkV1ScenarioService struct {
	Options []option.RequestOption
}

// NewAITestTestFrameworkV1ScenarioService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAITestTestFrameworkV1ScenarioService(opts ...option.RequestOption) (r AITestTestFrameworkV1ScenarioService) {
	r = AITestTestFrameworkV1ScenarioService{}
	r.Options = opts
	return
}

// Creates an evaluator scenario in your organization.
func (r *AITestTestFrameworkV1ScenarioService) New(ctx context.Context, body AITestTestFrameworkV1ScenarioNewParams, opts ...option.RequestOption) (res *EvaluatorScenarioDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the specified evaluator scenario with full detail.
func (r *AITestTestFrameworkV1ScenarioService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *EvaluatorScenarioDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/scenarios/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the specified evaluator scenario.
func (r *AITestTestFrameworkV1ScenarioService) Update(ctx context.Context, id int64, body AITestTestFrameworkV1ScenarioUpdateParams, opts ...option.RequestOption) (res *EvaluatorScenarioDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/scenarios/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of evaluator scenarios in your organization, optionally
// filtered by test agent or tags.
func (r *AITestTestFrameworkV1ScenarioService) List(ctx context.Context, query AITestTestFrameworkV1ScenarioListParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1ScenarioListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the specified evaluator scenario.
func (r *AITestTestFrameworkV1ScenarioService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *AITestTestFrameworkV1ScenarioDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/scenarios/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Starts background generation of evaluator scenarios for a test agent. Returns a
// progress ID to poll with the generate-scenarios-progress endpoint.
func (r *AITestTestFrameworkV1ScenarioService) GenerateBg(ctx context.Context, body AITestTestFrameworkV1ScenarioGenerateBgParams, opts ...option.RequestOption) (res *ScenarioGenerationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios/generate-bg"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Exports the selected evaluator scenarios as a CSV or XLSX file download.
func (r *AITestTestFrameworkV1ScenarioService) GetExportScenarios(ctx context.Context, query AITestTestFrameworkV1ScenarioGetExportScenariosParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")}, opts...)
	path := "ai/tests/test_framework/v1/scenarios/export_scenarios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the status of a background scenario-generation job started with the
// generate-evaluator-scenarios endpoint.
func (r *AITestTestFrameworkV1ScenarioService) GetGenerateProgress(ctx context.Context, query AITestTestFrameworkV1ScenarioGetGenerateProgressParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1ScenarioGetGenerateProgressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios/generate-progress"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the status of an improve-scenarios session started with the
// improve-evaluator-scenarios endpoint.
func (r *AITestTestFrameworkV1ScenarioService) GetScenarioAgentProgress(ctx context.Context, query AITestTestFrameworkV1ScenarioGetScenarioAgentProgressParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios/scenario-agent-progress"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Runs the selected evaluator scenarios against their test agents and creates a
// run batch. Returns the ID of the created batch, whose results can be listed with
// the test-run-results endpoints.
func (r *AITestTestFrameworkV1ScenarioService) RunScenarios(ctx context.Context, body AITestTestFrameworkV1ScenarioRunScenariosParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1ScenarioRunScenariosResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios/run_scenarios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Starts an interactive session that improves selected evaluator scenarios from a
// natural-language instruction. Returns a progress ID to poll with the
// improve-scenarios-progress endpoint.
func (r *AITestTestFrameworkV1ScenarioService) ScenarioAgent(ctx context.Context, body AITestTestFrameworkV1ScenarioScenarioAgentParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1ScenarioScenarioAgentResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/scenarios/scenario-agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EvaluatorScenarioDetail struct {
	ID                      int64                                          `json:"id" api:"required"`
	Agent                   int64                                          `json:"agent" api:"required"`
	Name                    string                                         `json:"name" api:"required"`
	DynamicVariableValues   map[string]string                              `json:"dynamic_variable_values" api:"nullable"`
	ExpectedOutcomePrompt   string                                         `json:"expected_outcome_prompt"`
	FirstMessage            string                                         `json:"first_message"`
	FolderPath              string                                         `json:"folder_path" api:"nullable"`
	InboundPhoneNumber      int64                                          `json:"inbound_phone_number" api:"nullable"`
	InboundPhoneNumberData  EvaluatorScenarioDetailInboundPhoneNumberData  `json:"inbound_phone_number_data" api:"nullable"`
	InformationFields       map[string]string                              `json:"information_fields" api:"nullable"`
	Instructions            string                                         `json:"instructions"`
	IsPredefined            bool                                           `json:"is_predefined"`
	MaxDuration             int64                                          `json:"max_duration" api:"nullable"`
	MetricNames             []string                                       `json:"metric_names"`
	Metrics                 []int64                                        `json:"metrics"`
	OutboundPhoneNumberData EvaluatorScenarioDetailOutboundPhoneNumberData `json:"outbound_phone_number_data" api:"nullable"`
	Personality             int64                                          `json:"personality" api:"nullable"`
	PersonalityName         string                                         `json:"personality_name"`
	PhoneNumber             string                                         `json:"phone_number"`
	Project                 int64                                          `json:"project" api:"nullable"`
	Runs                    []int64                                        `json:"runs"`
	ScenarioLanguage        string                                         `json:"scenario_language" api:"nullable"`
	ScenarioType            string                                         `json:"scenario_type"`
	SimulationDescription   string                                         `json:"simulation_description"`
	SuiteType               string                                         `json:"suite_type"`
	Tags                    []string                                       `json:"tags"`
	TestProfile             int64                                          `json:"test_profile" api:"nullable"`
	TestProfileData         EvaluatorScenarioDetailTestProfileData         `json:"test_profile_data" api:"nullable"`
	ToolIDs                 []string                                       `json:"tool_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Agent                   respjson.Field
		Name                    respjson.Field
		DynamicVariableValues   respjson.Field
		ExpectedOutcomePrompt   respjson.Field
		FirstMessage            respjson.Field
		FolderPath              respjson.Field
		InboundPhoneNumber      respjson.Field
		InboundPhoneNumberData  respjson.Field
		InformationFields       respjson.Field
		Instructions            respjson.Field
		IsPredefined            respjson.Field
		MaxDuration             respjson.Field
		MetricNames             respjson.Field
		Metrics                 respjson.Field
		OutboundPhoneNumberData respjson.Field
		Personality             respjson.Field
		PersonalityName         respjson.Field
		PhoneNumber             respjson.Field
		Project                 respjson.Field
		Runs                    respjson.Field
		ScenarioLanguage        respjson.Field
		ScenarioType            respjson.Field
		SimulationDescription   respjson.Field
		SuiteType               respjson.Field
		Tags                    respjson.Field
		TestProfile             respjson.Field
		TestProfileData         respjson.Field
		ToolIDs                 respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluatorScenarioDetail) RawJSON() string { return r.JSON.raw }
func (r *EvaluatorScenarioDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluatorScenarioDetailInboundPhoneNumberData struct {
	ID            int64  `json:"id"`
	Number        string `json:"number"`
	PhoneNumberID string `json:"phone_number_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Number        respjson.Field
		PhoneNumberID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluatorScenarioDetailInboundPhoneNumberData) RawJSON() string { return r.JSON.raw }
func (r *EvaluatorScenarioDetailInboundPhoneNumberData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluatorScenarioDetailOutboundPhoneNumberData struct {
	ID            int64  `json:"id"`
	Number        string `json:"number"`
	PhoneNumberID string `json:"phone_number_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Number        respjson.Field
		PhoneNumberID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluatorScenarioDetailOutboundPhoneNumberData) RawJSON() string { return r.JSON.raw }
func (r *EvaluatorScenarioDetailOutboundPhoneNumberData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluatorScenarioDetailTestProfileData struct {
	ID          int64             `json:"id"`
	Agent       int64             `json:"agent" api:"nullable"`
	Information map[string]string `json:"information" api:"nullable"`
	Name        string            `json:"name"`
	Project     int64             `json:"project" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Agent       respjson.Field
		Information respjson.Field
		Name        respjson.Field
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluatorScenarioDetailTestProfileData) RawJSON() string { return r.JSON.raw }
func (r *EvaluatorScenarioDetailTestProfileData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioListResponse struct {
	Count    int64                                             `json:"count" api:"required"`
	Next     string                                            `json:"next" api:"required" format:"uri"`
	Previous string                                            `json:"previous" api:"required" format:"uri"`
	Results  []AITestTestFrameworkV1ScenarioListResponseResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioListResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ScenarioListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioListResponseResult struct {
	ID int64 `json:"id" api:"required"`
	// ID of the test agent the scenario runs against.
	Agent                 int64             `json:"agent" api:"required"`
	CreatedAt             time.Time         `json:"created_at" api:"required" format:"date-time"`
	Name                  string            `json:"name" api:"required"`
	UpdatedAt             time.Time         `json:"updated_at" api:"required" format:"date-time"`
	DynamicVariableValues map[string]string `json:"dynamic_variable_values" api:"nullable"`
	// What a successful conversation looks like; scored by the expected-outcome
	// metric.
	ExpectedOutcomePrompt string `json:"expected_outcome_prompt"`
	// First message the simulated caller speaks or sends.
	FirstMessage           string `json:"first_message"`
	FolderPath             string `json:"folder_path" api:"nullable"`
	InboundPhoneNumber     int64  `json:"inbound_phone_number" api:"nullable"`
	InboundPhoneNumberData string `json:"inbound_phone_number_data" api:"nullable"`
	// What the simulated caller should do during the conversation.
	Instructions  string   `json:"instructions"`
	LastUpdatedBy int64    `json:"last_updated_by" api:"nullable"`
	MetricNames   []string `json:"metric_names"`
	// Metric IDs applied to the scenario.
	Metrics []int64 `json:"metrics"`
	// ID of the test personality the caller simulates.
	Personality     int64  `json:"personality" api:"nullable"`
	PersonalityName string `json:"personality_name"`
	// Outbound phone number used to place the simulated call.
	PhoneNumber      string `json:"phone_number"`
	ScenarioLanguage string `json:"scenario_language" api:"nullable"`
	// Type of scenario: workflow, red_teaming_voice, or red_teaming_text.
	ScenarioType    string                                                          `json:"scenario_type"`
	Tags            []string                                                        `json:"tags"`
	TestProfile     int64                                                           `json:"test_profile" api:"nullable"`
	TestProfileData AITestTestFrameworkV1ScenarioListResponseResultsTestProfileData `json:"test_profile_data" api:"nullable"`
	ToolIDs         []string                                                        `json:"tool_ids"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Agent                  respjson.Field
		CreatedAt              respjson.Field
		Name                   respjson.Field
		UpdatedAt              respjson.Field
		DynamicVariableValues  respjson.Field
		ExpectedOutcomePrompt  respjson.Field
		FirstMessage           respjson.Field
		FolderPath             respjson.Field
		InboundPhoneNumber     respjson.Field
		InboundPhoneNumberData respjson.Field
		Instructions           respjson.Field
		LastUpdatedBy          respjson.Field
		MetricNames            respjson.Field
		Metrics                respjson.Field
		Personality            respjson.Field
		PersonalityName        respjson.Field
		PhoneNumber            respjson.Field
		ScenarioLanguage       respjson.Field
		ScenarioType           respjson.Field
		Tags                   respjson.Field
		TestProfile            respjson.Field
		TestProfileData        respjson.Field
		ToolIDs                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ScenarioListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioListResponseResultsTestProfileData struct {
	ID          int64             `json:"id"`
	Agent       int64             `json:"agent" api:"nullable"`
	Information map[string]string `json:"information" api:"nullable"`
	Name        string            `json:"name"`
	Project     int64             `json:"project" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Agent       respjson.Field
		Information respjson.Field
		Name        respjson.Field
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioListResponseResultsTestProfileData) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1ScenarioListResponseResultsTestProfileData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioDeleteResponse = any

type AITestTestFrameworkV1ScenarioGetGenerateProgressResponse struct {
	CompletedScenarios int64            `json:"completed_scenarios"`
	FailedScenarios    int64            `json:"failed_scenarios"`
	ScenariosList      []map[string]any `json:"scenarios_list"`
	TotalScenarios     int64            `json:"total_scenarios"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedScenarios respjson.Field
		FailedScenarios    respjson.Field
		ScenariosList      respjson.Field
		TotalScenarios     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioGetGenerateProgressResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ScenarioGetGenerateProgressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponse struct {
	Error      string                                                              `json:"error" api:"nullable"`
	Mode       string                                                              `json:"mode"`
	ProgressID string                                                              `json:"progress_id"`
	Result     AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResult `json:"result" api:"nullable"`
	SessionID  int64                                                               `json:"session_id" api:"nullable"`
	Status     string                                                              `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Mode        respjson.Field
		ProgressID  respjson.Field
		Result      respjson.Field
		SessionID   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResult struct {
	Cancelled              bool                                                                                       `json:"cancelled"`
	ClarificationQuestions []AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultClarificationQuestion `json:"clarification_questions"`
	ConversationHistory    []AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultConversationHistory   `json:"conversation_history"`
	Message                string                                                                                     `json:"message"`
	NeedsClarification     bool                                                                                       `json:"needs_clarification"`
	SessionID              int64                                                                                      `json:"session_id"`
	Success                bool                                                                                       `json:"success"`
	UpdatedScenarios       []map[string]any                                                                           `json:"updated_scenarios"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled              respjson.Field
		ClarificationQuestions respjson.Field
		ConversationHistory    respjson.Field
		Message                respjson.Field
		NeedsClarification     respjson.Field
		SessionID              respjson.Field
		Success                respjson.Field
		UpdatedScenarios       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResult) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultClarificationQuestion struct {
	Question string `json:"question"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Question    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultClarificationQuestion) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultClarificationQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultConversationHistory struct {
	Content   string `json:"content"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultConversationHistory) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1ScenarioGetScenarioAgentProgressResponseResultConversationHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioRunScenariosResponse struct {
	// ID of the created test-run batch whose results can be polled.
	ID int64 `json:"id" api:"required"`
	// Initial status of the run batch.
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioRunScenariosResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ScenarioRunScenariosResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioScenarioAgentResponse struct {
	ProgressID string `json:"progress_id" api:"required"`
	SessionID  int64  `json:"session_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProgressID  respjson.Field
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1ScenarioScenarioAgentResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1ScenarioScenarioAgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioNewParams struct {
	Name string `json:"name" api:"required"`
	// ID of the personality the caller simulates.
	Personality           int64             `json:"personality" api:"required"`
	Agent                 param.Opt[int64]  `json:"agent,omitzero"`
	FolderPath            param.Opt[string] `json:"folder_path,omitzero"`
	ExpectedOutcomePrompt param.Opt[string] `json:"expected_outcome_prompt,omitzero"`
	FirstMessage          param.Opt[string] `json:"first_message,omitzero"`
	Instructions          param.Opt[string] `json:"instructions,omitzero"`
	TestProfile           param.Opt[int64]  `json:"test_profile,omitzero"`
	Metrics               []int64           `json:"metrics,omitzero"`
	Tags                  []string          `json:"tags,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1ScenarioNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1ScenarioNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1ScenarioNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioUpdateParams struct {
	Agent                 param.Opt[int64]  `json:"agent,omitzero"`
	FolderPath            param.Opt[string] `json:"folder_path,omitzero"`
	MaxDuration           param.Opt[int64]  `json:"max_duration,omitzero"`
	ExpectedOutcomePrompt param.Opt[string] `json:"expected_outcome_prompt,omitzero"`
	FirstMessage          param.Opt[string] `json:"first_message,omitzero"`
	InboundPhoneNumber    param.Opt[int64]  `json:"inbound_phone_number,omitzero"`
	Instructions          param.Opt[string] `json:"instructions,omitzero"`
	Name                  param.Opt[string] `json:"name,omitzero"`
	Personality           param.Opt[int64]  `json:"personality,omitzero"`
	PhoneNumber           param.Opt[int64]  `json:"phone_number,omitzero"`
	ScenarioLanguage      param.Opt[string] `json:"scenario_language,omitzero"`
	TestProfile           param.Opt[int64]  `json:"test_profile,omitzero"`
	Metrics               []int64           `json:"metrics,omitzero"`
	Tags                  []string          `json:"tags,omitzero"`
	ToolIDs               []string          `json:"tool_ids,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1ScenarioUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1ScenarioUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1ScenarioUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioListParams struct {
	// Filter scenarios belonging to the specified test agent.
	AgentID param.Opt[int64] `query:"agent_id,omitzero" json:"-"`
	// Filter scenarios belonging to the specified Telnyx assistant.
	AssistantID param.Opt[string] `query:"assistant_id,omitzero" json:"-"`
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter scenarios by a comma-separated list of tags.
	Tags param.Opt[string] `query:"tags,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITestTestFrameworkV1ScenarioListParams]'s query parameters
// as `url.Values`.
func (r AITestTestFrameworkV1ScenarioListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AITestTestFrameworkV1ScenarioGenerateBgParams struct {
	// How many scenarios to generate.
	NumScenarios       int64             `json:"num_scenarios" api:"required"`
	FolderPath         param.Opt[string] `json:"folder_path,omitzero"`
	InboundPhoneNumber param.Opt[int64]  `json:"inbound_phone_number,omitzero"`
	AgentID            param.Opt[int64]  `json:"agent_id,omitzero"`
	// Optional guidance steering scenario generation.
	ExtraInstructions param.Opt[string] `json:"extra_instructions,omitzero"`
	FirstMessage      param.Opt[string] `json:"first_message,omitzero"`
	// Whether to generate expected-outcome prompts for the scenarios.
	GenerateExpectedOutcomes param.Opt[bool]  `json:"generate_expected_outcomes,omitzero"`
	TestProfile              param.Opt[int64] `json:"test_profile,omitzero"`
	Personalities            []int64          `json:"personalities,omitzero"`
	// Type of scenario to generate.
	//
	// Any of "workflow", "red_teaming_voice", "red_teaming_text".
	ScenarioType AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioType `json:"scenario_type,omitzero"`
	Tags         []string                                                  `json:"tags,omitzero"`
	ToolIDs      []string                                                  `json:"tool_ids,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1ScenarioGenerateBgParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1ScenarioGenerateBgParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1ScenarioGenerateBgParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of scenario to generate.
type AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioType string

const (
	AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioTypeWorkflow        AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioType = "workflow"
	AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioTypeRedTeamingVoice AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioType = "red_teaming_voice"
	AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioTypeRedTeamingText  AITestTestFrameworkV1ScenarioGenerateBgParamsScenarioType = "red_teaming_text"
)

type AITestTestFrameworkV1ScenarioGetExportScenariosParams struct {
	// File format of the export.
	//
	// Any of "csv", "xlsx".
	ExportType AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportType `query:"export_type,omitzero" api:"required" json:"-"`
	// Comma-separated scenario IDs to export.
	ScenarioIDs string `query:"scenario_ids" api:"required" json:"-"`
	// Export scenarios belonging to the specified test agent.
	AgentID param.Opt[int64] `query:"agent_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITestTestFrameworkV1ScenarioGetExportScenariosParams]'s
// query parameters as `url.Values`.
func (r AITestTestFrameworkV1ScenarioGetExportScenariosParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// File format of the export.
type AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportType string

const (
	AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportTypeCsv  AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportType = "csv"
	AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportTypeXlsx AITestTestFrameworkV1ScenarioGetExportScenariosParamsExportType = "xlsx"
)

type AITestTestFrameworkV1ScenarioGetGenerateProgressParams struct {
	// Progress ID returned by the generate-evaluator-scenarios endpoint.
	ProgressID string `query:"progress_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AITestTestFrameworkV1ScenarioGetGenerateProgressParams]'s
// query parameters as `url.Values`.
func (r AITestTestFrameworkV1ScenarioGetGenerateProgressParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AITestTestFrameworkV1ScenarioGetScenarioAgentProgressParams struct {
	// Progress ID returned by the improve-evaluator-scenarios endpoint.
	ProgressID string `query:"progress_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes
// [AITestTestFrameworkV1ScenarioGetScenarioAgentProgressParams]'s query parameters
// as `url.Values`.
func (r AITestTestFrameworkV1ScenarioGetScenarioAgentProgressParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AITestTestFrameworkV1ScenarioRunScenariosParams struct {
	AgentID     param.Opt[int64]  `json:"agent_id,omitzero"`
	AgentNumber param.Opt[string] `json:"agent_number,omitzero"`
	AssistantID param.Opt[string] `json:"assistant_id,omitzero"`
	// Maximum number of concurrent simulated calls.
	ConcurrencyLimit    param.Opt[int64]  `json:"concurrency_limit,omitzero"`
	Frequency           param.Opt[int64]  `json:"frequency,omitzero"`
	Name                param.Opt[string] `json:"name,omitzero"`
	OutboundPhoneNumber param.Opt[string] `json:"outbound_phone_number,omitzero"`
	// Any of "same_number", "different_numbers".
	Mode           AITestTestFrameworkV1ScenarioRunScenariosParamsMode `json:"mode,omitzero"`
	PersonalityIDs []int64                                             `json:"personality_ids,omitzero"`
	// Scenario IDs to run, or the literal string 'all'.
	Scenarios      AITestTestFrameworkV1ScenarioRunScenariosParamsScenariosUnion `json:"scenarios,omitzero"`
	Tags           []string                                                      `json:"tags,omitzero"`
	TestProfileIDs []int64                                                       `json:"test_profile_ids,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1ScenarioRunScenariosParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1ScenarioRunScenariosParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1ScenarioRunScenariosParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1ScenarioRunScenariosParamsMode string

const (
	AITestTestFrameworkV1ScenarioRunScenariosParamsModeSameNumber       AITestTestFrameworkV1ScenarioRunScenariosParamsMode = "same_number"
	AITestTestFrameworkV1ScenarioRunScenariosParamsModeDifferentNumbers AITestTestFrameworkV1ScenarioRunScenariosParamsMode = "different_numbers"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITestTestFrameworkV1ScenarioRunScenariosParamsScenariosUnion struct {
	OfAnyArray []any             `json:",omitzero,inline"`
	OfString   param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u AITestTestFrameworkV1ScenarioRunScenariosParamsScenariosUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyArray, u.OfString)
}
func (u *AITestTestFrameworkV1ScenarioRunScenariosParamsScenariosUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITestTestFrameworkV1ScenarioRunScenariosParamsScenariosUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type AITestTestFrameworkV1ScenarioScenarioAgentParams struct {
	// Scenario IDs to improve, or the literal string 'all'.
	Scenarios AITestTestFrameworkV1ScenarioScenarioAgentParamsScenariosUnion `json:"scenarios,omitzero" api:"required"`
	// What the assistant should improve about the selected scenarios.
	UserMessage             string                                                    `json:"user_message" api:"required"`
	AgentID                 param.Opt[int64]                                          `json:"agent_id,omitzero"`
	SessionID               param.Opt[int64]                                          `json:"session_id,omitzero"`
	AgentResolutionResponse param.Opt[string]                                         `json:"agent_resolution_response,omitzero"`
	ApplyChanges            param.Opt[bool]                                           `json:"apply_changes,omitzero"`
	NumScenarios            param.Opt[int64]                                          `json:"num_scenarios,omitzero"`
	Context                 []AITestTestFrameworkV1ScenarioScenarioAgentParamsContext `json:"context,omitzero"`
	PendingToolCalls        []map[string]any                                          `json:"pending_tool_calls,omitzero"`
	KnowledgeBase           []int64                                                   `json:"knowledge_base,omitzero"`
	ScenariosNeedingAgent   []int64                                                   `json:"scenarios_needing_agent,omitzero"`
	UpdatedScenarios        []map[string]any                                          `json:"updated_scenarios,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1ScenarioScenarioAgentParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1ScenarioScenarioAgentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1ScenarioScenarioAgentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITestTestFrameworkV1ScenarioScenarioAgentParamsScenariosUnion struct {
	OfAnyArray []any             `json:",omitzero,inline"`
	OfString   param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u AITestTestFrameworkV1ScenarioScenarioAgentParamsScenariosUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyArray, u.OfString)
}
func (u *AITestTestFrameworkV1ScenarioScenarioAgentParamsScenariosUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITestTestFrameworkV1ScenarioScenarioAgentParamsScenariosUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type AITestTestFrameworkV1ScenarioScenarioAgentParamsContext struct {
	Content   param.Opt[string] `json:"content,omitzero"`
	Role      param.Opt[string] `json:"role,omitzero"`
	Timestamp param.Opt[string] `json:"timestamp,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1ScenarioScenarioAgentParamsContext) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1ScenarioScenarioAgentParamsContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1ScenarioScenarioAgentParamsContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
