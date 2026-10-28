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
// AITestTestFrameworkV1AiagentService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestTestFrameworkV1AiagentService] method instead.
type AITestTestFrameworkV1AiagentService struct {
	Options []option.RequestOption
}

// NewAITestTestFrameworkV1AiagentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAITestTestFrameworkV1AiagentService(opts ...option.RequestOption) (r AITestTestFrameworkV1AiagentService) {
	r = AITestTestFrameworkV1AiagentService{}
	r.Options = opts
	return
}

// Creates a test agent in your organization. The agent wraps a Telnyx AI assistant
// so simulated callers can be sent against it.
func (r *AITestTestFrameworkV1AiagentService) New(ctx context.Context, body AITestTestFrameworkV1AiagentNewParams, opts ...option.RequestOption) (res *TestAgent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/aiagents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns the specified test agent.
func (r *AITestTestFrameworkV1AiagentService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *TestAgent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/aiagents/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the specified test agent.
func (r *AITestTestFrameworkV1AiagentService) Update(ctx context.Context, id int64, body AITestTestFrameworkV1AiagentUpdateParams, opts ...option.RequestOption) (res *TestAgent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/aiagents/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of test agents in your organization. A test agent wraps
// the Telnyx AI assistant under test and drives simulated callers against it.
func (r *AITestTestFrameworkV1AiagentService) List(ctx context.Context, query AITestTestFrameworkV1AiagentListParams, opts ...option.RequestOption) (res *AITestTestFrameworkV1AiagentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/aiagents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes the specified test agent.
func (r *AITestTestFrameworkV1AiagentService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *AITestTestFrameworkV1AiagentDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/tests/test_framework/v1/aiagents/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type TestAgent struct {
	ID int64 `json:"id" api:"required"`
	// Name of the test agent.
	AgentName string `json:"agent_name" api:"required"`
	// ID of the Telnyx AI assistant under test.
	AssistantID            string    `json:"assistant_id" api:"required"`
	CreatedAt              time.Time `json:"created_at" api:"required" format:"date-time"`
	UpdatedAt              time.Time `json:"updated_at" api:"required" format:"date-time"`
	AutoUpdateDropoffNodes bool      `json:"auto_update_dropoff_nodes"`
	AutoUpdateTopicNodes   bool      `json:"auto_update_topic_nodes"`
	// ID of the chat assistant under test, when testing text channels.
	ChatAssistantID string `json:"chat_assistant_id"`
	// Contact phone number for the test agent.
	ContactNumber string `json:"contact_number"`
	Description   string `json:"description"`
	// IDs of the test personalities enabled for this agent.
	EnabledPersonalities       []int64 `json:"enabled_personalities"`
	HallucinationMetricKBFiles []int64 `json:"hallucination_metric_kb_files"`
	// Whether the test agent accepts inbound calls.
	Inbound bool `json:"inbound"`
	// Knowledge-base files associated with the test agent.
	KnowledgeBaseFiles []any `json:"knowledge_base_files"`
	// ISO language code the test agent converses in.
	Language     string `json:"language"`
	LlmMaxTokens int64  `json:"llm_max_tokens"`
	// Model used to drive the test agent's side of the conversation.
	LlmModel         string  `json:"llm_model"`
	LlmSystemPrompt  string  `json:"llm_system_prompt"`
	LlmTemperature   float64 `json:"llm_temperature"`
	OutboundAutoCall bool    `json:"outbound_auto_call"`
	// Outbound numbers associated with the test agent.
	OutboundNumbers []any             `json:"outbound_numbers" api:"nullable"`
	SipAuth         map[string]string `json:"sip_auth" api:"nullable"`
	SipEndpoint     string            `json:"sip_endpoint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AgentName                  respjson.Field
		AssistantID                respjson.Field
		CreatedAt                  respjson.Field
		UpdatedAt                  respjson.Field
		AutoUpdateDropoffNodes     respjson.Field
		AutoUpdateTopicNodes       respjson.Field
		ChatAssistantID            respjson.Field
		ContactNumber              respjson.Field
		Description                respjson.Field
		EnabledPersonalities       respjson.Field
		HallucinationMetricKBFiles respjson.Field
		Inbound                    respjson.Field
		KnowledgeBaseFiles         respjson.Field
		Language                   respjson.Field
		LlmMaxTokens               respjson.Field
		LlmModel                   respjson.Field
		LlmSystemPrompt            respjson.Field
		LlmTemperature             respjson.Field
		OutboundAutoCall           respjson.Field
		OutboundNumbers            respjson.Field
		SipAuth                    respjson.Field
		SipEndpoint                respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestAgent) RawJSON() string { return r.JSON.raw }
func (r *TestAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1AiagentListResponse struct {
	Count int64 `json:"count" api:"required"`
	// Link to the next page of results, or null.
	Next string `json:"next" api:"required" format:"uri"`
	// Link to the previous page of results, or null.
	Previous string      `json:"previous" api:"required" format:"uri"`
	Results  []TestAgent `json:"results" api:"required"`
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
func (r AITestTestFrameworkV1AiagentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1AiagentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1AiagentDeleteResponse = any

type AITestTestFrameworkV1AiagentNewParams struct {
	// Name of the test agent.
	AgentName string `json:"agent_name" api:"required"`
	// ID of the Telnyx AI assistant under test.
	AssistantID            string            `json:"assistant_id" api:"required"`
	AutoFetchCallsEnabled  param.Opt[bool]   `json:"auto_fetch_calls_enabled,omitzero"`
	AutoSyncPromptEnabled  param.Opt[bool]   `json:"auto_sync_prompt_enabled,omitzero"`
	AutoUpdateDropoffNodes param.Opt[bool]   `json:"auto_update_dropoff_nodes,omitzero"`
	AutoUpdateTopicNodes   param.Opt[bool]   `json:"auto_update_topic_nodes,omitzero"`
	ContactNumber          param.Opt[string] `json:"contact_number,omitzero"`
	Description            param.Opt[string] `json:"description,omitzero"`
	Inbound                param.Opt[bool]   `json:"inbound,omitzero"`
	// ISO language code the test agent converses in.
	Language         param.Opt[string]  `json:"language,omitzero"`
	LlmMaxTokens     param.Opt[int64]   `json:"llm_max_tokens,omitzero"`
	LlmModel         param.Opt[string]  `json:"llm_model,omitzero"`
	LlmSystemPrompt  param.Opt[string]  `json:"llm_system_prompt,omitzero"`
	LlmTemperature   param.Opt[float64] `json:"llm_temperature,omitzero"`
	OutboundAutoCall param.Opt[bool]    `json:"outbound_auto_call,omitzero"`
	SipEndpoint      param.Opt[string]  `json:"sip_endpoint,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1AiagentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1AiagentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1AiagentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1AiagentUpdateParams struct {
	// Name of the test agent.
	AgentName param.Opt[string] `json:"agent_name,omitzero"`
	// ID of the Telnyx AI assistant under test.
	AssistantID            param.Opt[string] `json:"assistant_id,omitzero"`
	AutoFetchCallsEnabled  param.Opt[bool]   `json:"auto_fetch_calls_enabled,omitzero"`
	AutoSyncPromptEnabled  param.Opt[bool]   `json:"auto_sync_prompt_enabled,omitzero"`
	AutoUpdateDropoffNodes param.Opt[bool]   `json:"auto_update_dropoff_nodes,omitzero"`
	AutoUpdateTopicNodes   param.Opt[bool]   `json:"auto_update_topic_nodes,omitzero"`
	ContactNumber          param.Opt[string] `json:"contact_number,omitzero"`
	Description            param.Opt[string] `json:"description,omitzero"`
	Inbound                param.Opt[bool]   `json:"inbound,omitzero"`
	// ISO language code the test agent converses in.
	Language         param.Opt[string]  `json:"language,omitzero"`
	LlmMaxTokens     param.Opt[int64]   `json:"llm_max_tokens,omitzero"`
	LlmModel         param.Opt[string]  `json:"llm_model,omitzero"`
	LlmSystemPrompt  param.Opt[string]  `json:"llm_system_prompt,omitzero"`
	LlmTemperature   param.Opt[float64] `json:"llm_temperature,omitzero"`
	OutboundAutoCall param.Opt[bool]    `json:"outbound_auto_call,omitzero"`
	SipEndpoint      param.Opt[string]  `json:"sip_endpoint,omitzero"`
	paramObj
}

func (r AITestTestFrameworkV1AiagentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AITestTestFrameworkV1AiagentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITestTestFrameworkV1AiagentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITestTestFrameworkV1AiagentListParams struct {
	// Page number (1-based).
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITestTestFrameworkV1AiagentListParams]'s query parameters
// as `url.Values`.
func (r AITestTestFrameworkV1AiagentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
