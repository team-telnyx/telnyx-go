// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
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
)

// Configure AI assistant specifications
//
// AIAssistantVersionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantVersionService] method instead.
type AIAssistantVersionService struct {
	Options []option.RequestOption
}

// NewAIAssistantVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIAssistantVersionService(opts ...option.RequestOption) (r AIAssistantVersionService) {
	r = AIAssistantVersionService{}
	r.Options = opts
	return
}

// Retrieves a specific version of an assistant by assistant_id and version_id
func (r *AIAssistantVersionService) Get(ctx context.Context, versionID string, params AIAssistantVersionGetParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s", params.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Updates the configuration of a specific assistant version. Can not update main
// version
func (r *AIAssistantVersionService) Update(ctx context.Context, versionID string, params AIAssistantVersionUpdateParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s", params.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves all versions of a specific assistant with complete configuration and
// metadata
func (r *AIAssistantVersionService) List(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/versions", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Permanently removes a specific version of an assistant. Can not delete main
// version
func (r *AIAssistantVersionService) Delete(ctx context.Context, versionID string, body AIAssistantVersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return err
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s", body.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Promotes a specific version to be the main/current version of the assistant.
// This will delete any existing canary deploy configuration and send all live
// production traffic to this version.
func (r *AIAssistantVersionService) Promote(ctx context.Context, versionID string, body AIAssistantVersionPromoteParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s/promote", body.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type UpdateAssistantParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	// If the dynamic_variables_webhook_url is set for the assistant, we will send a
	// request at the start of the conversation. See our
	// [guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// for more information.
	DynamicVariablesWebhookURL param.Opt[string] `json:"dynamic_variables_webhook_url,omitzero"`
	// Text that the assistant will use to start the conversation. This may be
	// templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables).
	// Use an empty string to have the assistant wait for the user to speak first. Use
	// the special value `<assistant-speaks-first-with-model-generated-message>` to
	// have the assistant generate the greeting based on the system instructions.
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// System instructions for the assistant. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// This is only needed when using third-party inference providers. The `identifier`
	// for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your LLM provider's API key. Warning: Free plans are unlikely to
	// work with this integration.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// ID of the model to use. You can use the
	// [Get models API](https://developers.telnyx.com/api-reference/chat/get-available-models)
	// to see all of your available models,
	Model param.Opt[string] `json:"model,omitzero"`
	Name  param.Opt[string] `json:"name,omitzero"`
	// Map of dynamic variables and their default values
	DynamicVariables      map[string]any                     `json:"dynamic_variables,omitzero"`
	EnabledFeatures       []EnabledFeatures                  `json:"enabled_features,omitzero"`
	ExternalLlm           UpdateAssistantExternalLlmParam    `json:"external_llm,omitzero"`
	FallbackConfig        UpdateAssistantFallbackConfigParam `json:"fallback_config,omitzero"`
	InsightSettings       InsightSettingsParam               `json:"insight_settings,omitzero"`
	MessagingSettings     MessagingSettingsParam             `json:"messaging_settings,omitzero"`
	ObservabilitySettings ObservabilityReqParam              `json:"observability_settings,omitzero"`
	// Configuration for post-conversation processing. When enabled, the assistant
	// receives one additional LLM turn after the conversation ends, allowing it to
	// execute tool calls such as logging to a CRM or sending a summary. The assistant
	// can execute multiple parallel or sequential tools during this phase.
	// Telephony-control tools (e.g. hangup, transfer) are unavailable
	// post-conversation. Beta feature.
	PostConversationSettings UpdateAssistantPostConversationSettingsParam `json:"post_conversation_settings,omitzero"`
	PrivacySettings          PrivacySettingsParam                         `json:"privacy_settings,omitzero"`
	TelephonySettings        TelephonySettingsParam                       `json:"telephony_settings,omitzero"`
	ToolIDs                  []string                                     `json:"tool_ids,omitzero"`
	// The tools that the assistant can use. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Tools         []AssistantToolsItemsUnionParam `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam      `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam              `json:"voice_settings,omitzero"`
	// Configuration settings for the assistant's web widget.
	WidgetSettings WidgetSettingsParam `json:"widget_settings,omitzero"`
	paramObj
}

func (r UpdateAssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Model are required.
type UpdateAssistantExternalLlmParam struct {
	// Base URL for the external LLM endpoint.
	BaseURL string `json:"base_url" api:"required"`
	// Model identifier to use with the external LLM endpoint.
	Model string `json:"model" api:"required"`
	// Integration secret identifier for the client certificate used with certificate
	// authentication.
	CertificateRef param.Opt[string] `json:"certificate_ref,omitzero"`
	// When enabled, Telnyx forwards conversation metadata and dynamic variables to the
	// external LLM endpoint. Defaults to false. The external endpoint receives the
	// standard chat completions payload with top-level `metadata` and
	// `dynamic_variables` objects when values are available. For example:
	// `{"metadata":{"conversation_id":"conv_123","assistant_id":"assistant_456","call_control_id":"v3:abc123","telnyx_conversation_channel":"phone_call"},"dynamic_variables":{"customer_name":"Jane","account_id":"acct_789","telnyx_agent_target":"+13125550100","telnyx_end_user_target":"+13125550123"}}`.
	ForwardMetadata param.Opt[bool] `json:"forward_metadata,omitzero"`
	// Integration secret identifier for the external LLM API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// URL used to retrieve an access token when certificate authentication is enabled.
	TokenRetrievalURL param.Opt[string] `json:"token_retrieval_url,omitzero"`
	// Authentication method used when connecting to the external LLM endpoint.
	//
	// Any of "token", "certificate".
	AuthenticationMethod string `json:"authentication_method,omitzero"`
	paramObj
}

func (r UpdateAssistantExternalLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantExternalLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantExternalLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantExternalLlmParam](
		"authentication_method", "token", "certificate",
	)
}

type UpdateAssistantFallbackConfigParam struct {
	// Integration secret identifier for the fallback model API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Fallback Telnyx-hosted model to use when the primary LLM provider is
	// unavailable.
	Model       param.Opt[string]                             `json:"model,omitzero"`
	ExternalLlm UpdateAssistantFallbackConfigExternalLlmParam `json:"external_llm,omitzero"`
	paramObj
}

func (r UpdateAssistantFallbackConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantFallbackConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantFallbackConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Model are required.
type UpdateAssistantFallbackConfigExternalLlmParam struct {
	// Base URL for the external LLM endpoint.
	BaseURL string `json:"base_url" api:"required"`
	// Model identifier to use with the external LLM endpoint.
	Model string `json:"model" api:"required"`
	// Integration secret identifier for the client certificate used with certificate
	// authentication.
	CertificateRef param.Opt[string] `json:"certificate_ref,omitzero"`
	// When enabled, Telnyx forwards conversation metadata and dynamic variables to the
	// external LLM endpoint. Defaults to false. The external endpoint receives the
	// standard chat completions payload with top-level `metadata` and
	// `dynamic_variables` objects when values are available. For example:
	// `{"metadata":{"conversation_id":"conv_123","assistant_id":"assistant_456","call_control_id":"v3:abc123","telnyx_conversation_channel":"phone_call"},"dynamic_variables":{"customer_name":"Jane","account_id":"acct_789","telnyx_agent_target":"+13125550100","telnyx_end_user_target":"+13125550123"}}`.
	ForwardMetadata param.Opt[bool] `json:"forward_metadata,omitzero"`
	// Integration secret identifier for the external LLM API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// URL used to retrieve an access token when certificate authentication is enabled.
	TokenRetrievalURL param.Opt[string] `json:"token_retrieval_url,omitzero"`
	// Authentication method used when connecting to the external LLM endpoint.
	//
	// Any of "token", "certificate".
	AuthenticationMethod string `json:"authentication_method,omitzero"`
	paramObj
}

func (r UpdateAssistantFallbackConfigExternalLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantFallbackConfigExternalLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantFallbackConfigExternalLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantFallbackConfigExternalLlmParam](
		"authentication_method", "token", "certificate",
	)
}

// Configuration for post-conversation processing. When enabled, the assistant
// receives one additional LLM turn after the conversation ends, allowing it to
// execute tool calls such as logging to a CRM or sending a summary. The assistant
// can execute multiple parallel or sequential tools during this phase.
// Telephony-control tools (e.g. hangup, transfer) are unavailable
// post-conversation. Beta feature.
type UpdateAssistantPostConversationSettingsParam struct {
	// Whether post-conversation processing is enabled. When true, the assistant will
	// be invoked after the conversation ends to perform any final tool calls. Defaults
	// to false.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r UpdateAssistantPostConversationSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantPostConversationSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantPostConversationSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantVersionGetParams struct {
	AssistantID       string          `path:"assistant_id" api:"required" json:"-"`
	IncludeMcpServers param.Opt[bool] `query:"include_mcp_servers,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantVersionGetParams]'s query parameters as
// `url.Values`.
func (r AIAssistantVersionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAssistantVersionUpdateParams struct {
	AssistantID     string `path:"assistant_id" api:"required" json:"-"`
	UpdateAssistant UpdateAssistantParam
	paramObj
}

func (r AIAssistantVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAssistant)
}
func (r *AIAssistantVersionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantVersionDeleteParams struct {
	AssistantID string `path:"assistant_id" api:"required" json:"-"`
	paramObj
}

type AIAssistantVersionPromoteParams struct {
	AssistantID string `path:"assistant_id" api:"required" json:"-"`
	paramObj
}
