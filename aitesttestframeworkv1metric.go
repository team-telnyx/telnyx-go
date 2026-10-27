// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Test agents, evaluator scenarios, metrics, and run results for the AI test
// framework.
//
// AITestTestFrameworkV1MetricService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestTestFrameworkV1MetricService] method instead.
type AITestTestFrameworkV1MetricService struct {
	Options []option.RequestOption
}

// NewAITestTestFrameworkV1MetricService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAITestTestFrameworkV1MetricService(opts ...option.RequestOption) (r AITestTestFrameworkV1MetricService) {
	r = AITestTestFrameworkV1MetricService{}
	r.Options = opts
	return
}

// Creates a metric in your organization.
func (r *AITestTestFrameworkV1MetricService) New(ctx context.Context, body AITestTestFrameworkV1MetricNewParams, opts ...option.RequestOption) (res *AITestMetric, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the specified metric.
func (r *AITestTestFrameworkV1MetricService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *AITestMetric, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/metrics/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the specified metric.
func (r *AITestTestFrameworkV1MetricService) Update(ctx context.Context, id int64, body AITestTestFrameworkV1MetricUpdateParams, opts ...option.RequestOption) (res *AITestMetric, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/metrics/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns the metrics available in your organization, optionally filtered by test
// agent. Each metric scores one aspect of a conversation or scenario run.
func (r *AITestTestFrameworkV1MetricService) List(ctx context.Context, query AITestTestFrameworkV1MetricListParams, opts ...option.RequestOption) (res *[]AITestMetric, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the specified metric.
func (r *AITestTestFrameworkV1MetricService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *AITestTestFrameworkV1MetricDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/metrics/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Creates several metrics in one request. Accepts an array of metric objects and
// returns the created metrics.
func (r *AITestTestFrameworkV1MetricService) BulkNew(ctx context.Context, body AITestTestFrameworkV1MetricBulkNewParams, opts ...option.RequestOption) (res *[]AITestMetric, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/metrics/bulk_create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Starts background generation of suggested metrics for a test agent. Returns a
// progress ID to poll with the get-test-metric-generation-progress endpoint.
func (r *AITestTestFrameworkV1MetricService) GenerateMetrics(ctx context.Context, body AITestTestFrameworkV1MetricGenerateMetricsParams, opts ...option.RequestOption) (res *ScenarioGenerationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/metrics/generate_metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the status of a background metric-generation job started with the
// generate-test-metrics endpoint.
func (r *AITestTestFrameworkV1MetricService) GetGenerateMetricsProgress(ctx context.Context, query AITestTestFrameworkV1MetricGetGenerateMetricsProgressParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/metrics/generate_metrics_progress"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AITestMetric struct {
	ID int64 `json:"id" api:"required"`
	// What the metric scores: numeric, continuous_qualitative, or
	// binary_workflow_adherence.
	EvalType string `json:"eval_type" api:"required"`
	Name     string `json:"name" api:"required"`
	// How the metric is evaluated: custom_code or llm_judge.
	Type string `json:"type" api:"required"`
	// Whether the metric is attached to newly created test agents by default.
	AddToNewAgents bool     `json:"add_to_new_agents" api:"nullable"`
	Agent          int64    `json:"agent" api:"nullable"`
	AgentNames     []string `json:"agent_names"`
	// Test agent IDs the metric is attached to.
	Agents        []int64        `json:"agents"`
	AlertFilters  map[string]any `json:"alert_filters" api:"nullable"`
	AlertType     string         `json:"alert_type"`
	AudioEnabled  bool           `json:"audio_enabled"`
	Configuration map[string]any `json:"configuration" api:"nullable"`
	Description   string         `json:"description"`
	DisplayOrder  int64          `json:"display_order"`
	// Allowed values for qualitative enum metrics.
	EnumValues                  []string `json:"enum_values"`
	EvaluationTrigger           string   `json:"evaluation_trigger"`
	EvaluationTriggerCustomCode string   `json:"evaluation_trigger_custom_code"`
	EvaluationTriggerPrompt     string   `json:"evaluation_trigger_prompt"`
	// Smoothing factor for the alert trend.
	EwmaAlpha float64 `json:"ewma_alpha"`
	// Built-in evaluation function used by custom_code metrics.
	FunctionName             string                        `json:"function_name" api:"nullable"`
	KBFileIDs                []int64                       `json:"kb_file_ids"`
	KnowledgeBaseFiles       []string                      `json:"knowledge_base_files"`
	MetricCost               float64                       `json:"metric_cost"`
	ObservabilityEnabled     bool                          `json:"observability_enabled"`
	OverallScore             AITestMetricOverallScoreUnion `json:"overall_score"`
	PriorityAssignmentPrompt string                        `json:"priority_assignment_prompt"`
	Project                  int64                         `json:"project" api:"nullable"`
	// Judge prompt used by llm_judge metrics.
	Prompt                          string  `json:"prompt"`
	SampleRate                      float64 `json:"sample_rate" api:"nullable"`
	SamplingEnabled                 bool    `json:"sampling_enabled"`
	SignificantChangeAlertDirection string  `json:"significant_change_alert_direction"`
	SignificantChangeAlertStatus    string  `json:"significant_change_alert_status"`
	SimulationEnabled               bool    `json:"simulation_enabled"`
	// Stable identifier of a built-in metric.
	Slug string `json:"slug" api:"nullable"`
	// Standard-deviation multiplier for significant-change alerts.
	StdMultiplier float64                     `json:"std_multiplier"`
	TotalScore    AITestMetricTotalScoreUnion `json:"total_score"`
	TriggerType   string                      `json:"trigger_type"`
	UserDefined   bool                        `json:"user_defined"`
	UsesLlmJudge  bool                        `json:"uses_llm_judge"`
	// Sliding window size the alert status is computed over.
	WindowSize int64 `json:"window_size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		EvalType                        respjson.Field
		Name                            respjson.Field
		Type                            respjson.Field
		AddToNewAgents                  respjson.Field
		Agent                           respjson.Field
		AgentNames                      respjson.Field
		Agents                          respjson.Field
		AlertFilters                    respjson.Field
		AlertType                       respjson.Field
		AudioEnabled                    respjson.Field
		Configuration                   respjson.Field
		Description                     respjson.Field
		DisplayOrder                    respjson.Field
		EnumValues                      respjson.Field
		EvaluationTrigger               respjson.Field
		EvaluationTriggerCustomCode     respjson.Field
		EvaluationTriggerPrompt         respjson.Field
		EwmaAlpha                       respjson.Field
		FunctionName                    respjson.Field
		KBFileIDs                       respjson.Field
		KnowledgeBaseFiles              respjson.Field
		MetricCost                      respjson.Field
		ObservabilityEnabled            respjson.Field
		OverallScore                    respjson.Field
		PriorityAssignmentPrompt        respjson.Field
		Project                         respjson.Field
		Prompt                          respjson.Field
		SampleRate                      respjson.Field
		SamplingEnabled                 respjson.Field
		SignificantChangeAlertDirection respjson.Field
		SignificantChangeAlertStatus    respjson.Field
		SimulationEnabled               respjson.Field
		Slug                            respjson.Field
		StdMultiplier                   respjson.Field
		TotalScore                      respjson.Field
		TriggerType                     respjson.Field
		UserDefined                     respjson.Field
		UsesLlmJudge                    respjson.Field
		WindowSize                      respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestMetric) RawJSON() string { return r.JSON.raw }
func (r *AITestMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestMetricOverallScoreUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type AITestMetricOverallScoreUnion struct {
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

func (u AITestMetricOverallScoreUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestMetricOverallScoreUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestMetricOverallScoreUnion) RawJSON() string { return u.JSON.raw }

func (r *AITestMetricOverallScoreUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestMetricTotalScoreUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type AITestMetricTotalScoreUnion struct {
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

func (u AITestMetricTotalScoreUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestMetricTotalScoreUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestMetricTotalScoreUnion) RawJSON() string { return u.JSON.raw }

func (r *AITestMetricTotalScoreUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type CreateAITestMetricRequestParam struct {
	Name           string            `json:"name" api:"required"`
	AddToNewAgents param.Opt[bool]   `json:"add_to_new_agents,omitzero"`
	AssistantID    param.Opt[string] `json:"assistant_id,omitzero"`
	AudioEnabled   param.Opt[bool]   `json:"audio_enabled,omitzero"`
	Description    param.Opt[string] `json:"description,omitzero"`
	DisplayOrder   param.Opt[int64]  `json:"display_order,omitzero"`
	// What the metric scores: numeric, continuous_qualitative, or
	// binary_workflow_adherence.
	EvalType             param.Opt[string] `json:"eval_type,omitzero"`
	ObservabilityEnabled param.Opt[bool]   `json:"observability_enabled,omitzero"`
	// Judge prompt used by llm_judge metrics.
	Prompt            param.Opt[string] `json:"prompt,omitzero"`
	SamplingEnabled   param.Opt[bool]   `json:"sampling_enabled,omitzero"`
	SimulationEnabled param.Opt[bool]   `json:"simulation_enabled,omitzero"`
	// How the metric is evaluated: custom_code or llm_judge.
	Type       param.Opt[string] `json:"type,omitzero"`
	Agents     []int64           `json:"agents,omitzero"`
	EnumValues []string          `json:"enum_values,omitzero"`
	paramObj
}

func (r CreateAITestMetricRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAITestMetricRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAITestMetricRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScenarioGenerationResponse struct {
	// Identifier to poll generation progress with.
	ProgressID string `json:"progress_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProgressID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScenarioGenerationResponse) RawJSON() string { return r.JSON.raw }
func (r *ScenarioGenerationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricDeleteResponse = any

type AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponse struct {
	AgentID     int64                                                                 `json:"agent_id" api:"required"`
	CurrentStep string                                                                `json:"current_step" api:"required"`
	Error       string                                                                `json:"error" api:"required"`
	Metrics     []AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseMetric `json:"metrics" api:"required"`
	// Any of "queued", "processing", "completed", "failed".
	Status AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		CurrentStep respjson.Field
		Error       respjson.Field
		Metrics     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseMetric struct {
	Agent       int64   `json:"agent"`
	Description string  `json:"description"`
	EvalType    string  `json:"eval_type"`
	Name        string  `json:"name"`
	Scenarios   []int64 `json:"scenarios"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		Description respjson.Field
		EvalType    respjson.Field
		Name        respjson.Field
		Scenarios   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseMetric) RawJSON() string {
	return r.JSON.raw
}
func (r *AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatus string

const (
	AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatusQueued     AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatus = "queued"
	AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatusProcessing AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatus = "processing"
	AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatusCompleted  AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatus = "completed"
	AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatusFailed     AITestTestFrameworkV1MetricGetGenerateMetricsProgressResponseStatus = "failed"
)

type AITestTestFrameworkV1MetricNewParams struct {
	CreateAITestMetricRequest CreateAITestMetricRequestParam
	paramObj
}

func (r AITestTestFrameworkV1MetricNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAITestMetricRequest)
}
func (r *AITestTestFrameworkV1MetricNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricUpdateParams struct {
	AddToNewAgents param.Opt[bool]   `json:"add_to_new_agents,omitzero"`
	AudioEnabled   param.Opt[bool]   `json:"audio_enabled,omitzero"`
	Description    param.Opt[string] `json:"description,omitzero"`
	DisplayOrder   param.Opt[int64]  `json:"display_order,omitzero"`
	// What the metric scores: numeric, continuous_qualitative, or
	// binary_workflow_adherence.
	EvalType             param.Opt[string] `json:"eval_type,omitzero"`
	Name                 param.Opt[string] `json:"name,omitzero"`
	ObservabilityEnabled param.Opt[bool]   `json:"observability_enabled,omitzero"`
	// Judge prompt used by llm_judge metrics.
	Prompt            param.Opt[string] `json:"prompt,omitzero"`
	SamplingEnabled   param.Opt[bool]   `json:"sampling_enabled,omitzero"`
	SimulationEnabled param.Opt[bool]   `json:"simulation_enabled,omitzero"`
	// How the metric is evaluated: custom_code or llm_judge.
	Type       param.Opt[string] `json:"type,omitzero"`
	Agents     []int64           `json:"agents,omitzero"`
	EnumValues []string          `json:"enum_values,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1MetricUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1MetricUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1MetricUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricListParams struct {
	// Filter metrics attached to the specified test agent.
	AgentID param.Opt[int64] `query:"agent_id,omitzero" json:"-"`
	// Filter metrics attached to the specified Telnyx assistant.
	AssistantID param.Opt[string] `query:"assistant_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITestTestFrameworkV1MetricListParams]'s query parameters
// as `url.Values`.
func (r AITestTestFrameworkV1MetricListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AITestTestFrameworkV1MetricBulkNewParams struct {
	Body []CreateAITestMetricRequestParam
	paramObj
}

func (r AITestTestFrameworkV1MetricBulkNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AITestTestFrameworkV1MetricBulkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricGenerateMetricsParams struct {
	// Test agent to suggest metrics for.
	AgentID int64 `json:"agent_id" api:"required"`
	// How many metrics to suggest.
	NumMetrics int64 `json:"num_metrics" api:"required"`
	// Scenario IDs to base the suggestions on.
	Scenarios []int64 `json:"scenarios,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1MetricGenerateMetricsParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1MetricGenerateMetricsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1MetricGenerateMetricsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1MetricGetGenerateMetricsProgressParams struct {
	// Progress ID returned by the generate-test-metrics endpoint.
	ProgressID string `query:"progress_id" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes
// [AITestTestFrameworkV1MetricGetGenerateMetricsProgressParams]'s query parameters
// as `url.Values`.
func (r AITestTestFrameworkV1MetricGetGenerateMetricsProgressParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
