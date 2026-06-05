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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
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
	// Timeout in milliseconds for the dynamic variables webhook. Must be between 1 and
	// 10000 ms. If the webhook does not respond within this timeout, the call proceeds
	// with default values. See the
	// [dynamic variables guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables).
	DynamicVariablesWebhookTimeoutMs param.Opt[int64] `json:"dynamic_variables_webhook_timeout_ms,omitzero"`
	// If `dynamic_variables_webhook_url` is set, Telnyx sends a POST request to this
	// URL at the start of the conversation to resolve dynamic variables. **Gotcha:**
	// the webhook response must wrap variables under a top-level `dynamic_variables`
	// object, e.g. `{"dynamic_variables": {"customer_name": "Jane"}}`. Returning a
	// flat object will be ignored and variables will fall back to their defaults. See
	// the
	// [dynamic variables guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// for the full request/response format and timeout behavior.
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
	// This is only needed when using third-party inference providers selected by
	// `model`. The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your LLM provider's API key. For bring-your-own endpoint
	// authentication, use `external_llm.llm_api_key_ref` instead. Warning: Free plans
	// are unlikely to work with this integration.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// ID of the model to use when `external_llm` is not set. You can use the
	// [Get models API](https://developers.telnyx.com/api-reference/openai-chat/get-available-models-openai-compatible)
	// to see available models. If `external_llm` is provided, the assistant uses
	// `external_llm` instead of this field. If neither `model` nor `external_llm` is
	// provided, Telnyx applies the default model.
	Model param.Opt[string] `json:"model,omitzero"`
	Name  param.Opt[string] `json:"name,omitzero"`
	// Human-readable name for the assistant version.
	VersionName param.Opt[string] `json:"version_name,omitzero"`
	// Conversation flow as supplied by API clients (create / update).
	//
	// A directed graph of `FlowNodeReq` connected by `FlowEdge`s. Validation enforces
	// unique node/edge IDs, that `start_node_id` references a real node, and that
	// every edge's endpoints reference real nodes.
	ConversationFlow UpdateAssistantConversationFlowParam `json:"conversation_flow,omitzero"`
	// Map of dynamic variables and their default values
	DynamicVariables map[string]any         `json:"dynamic_variables,omitzero"`
	EnabledFeatures  []EnabledFeatures      `json:"enabled_features,omitzero"`
	ExternalLlm      ExternalLlmReqParam    `json:"external_llm,omitzero"`
	FallbackConfig   FallbackConfigReqParam `json:"fallback_config,omitzero"`
	InsightSettings  InsightSettingsParam   `json:"insight_settings,omitzero"`
	// Connected integrations attached to the assistant. The catalog of available
	// integrations is at `/ai/integrations`; the user's connected integrations are at
	// `/ai/integrations/connections`. Each item references a catalog integration by
	// `integration_id`.
	Integrations []AssistantIntegrationParam `json:"integrations,omitzero"`
	// Settings for interruptions and how the assistant decides the user has finished
	// speaking. These timings are most relevant when using non turn-taking
	// transcription models. For turn-taking models like `deepgram/flux`, end-of-turn
	// behavior is controlled by the transcription end-of-turn settings under
	// `transcription.settings` (`eot_threshold`, `eot_timeout_ms`,
	// `eager_eot_threshold`).
	InterruptionSettings InferenceEmbeddingInterruptionSettingsParam `json:"interruption_settings,omitzero"`
	// MCP servers attached to the assistant. Create MCP servers with
	// `/ai/mcp_servers`, then reference them by `id` here.
	McpServers            []AssistantMcpServerParam `json:"mcp_servers,omitzero"`
	MessagingSettings     MessagingSettingsParam    `json:"messaging_settings,omitzero"`
	ObservabilitySettings ObservabilityReqParam     `json:"observability_settings,omitzero"`
	// Configuration for post-conversation processing. When enabled, the assistant
	// receives one additional LLM turn after the conversation ends, allowing it to
	// execute tool calls such as logging to a CRM or sending a summary. The assistant
	// can execute multiple parallel or sequential tools during this phase.
	// Telephony-control tools (e.g. hangup, transfer) are unavailable
	// post-conversation. Beta feature.
	PostConversationSettings PostConversationSettingsReqParam `json:"post_conversation_settings,omitzero"`
	PrivacySettings          PrivacySettingsParam             `json:"privacy_settings,omitzero"`
	// Tags associated with the assistant. Tags can also be managed with the assistant
	// tag endpoints.
	Tags              []string               `json:"tags,omitzero"`
	TelephonySettings TelephonySettingsParam `json:"telephony_settings,omitzero"`
	// IDs of shared tools to attach to the assistant. New integrations should prefer
	// `tool_ids` over inline `tools`.
	ToolIDs []string `json:"tool_ids,omitzero"`
	// Deprecated for new integrations. Inline tool definitions available to the
	// assistant. Prefer `tool_ids` to attach shared tools created with the AI Tools
	// endpoints.
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

// Conversation flow as supplied by API clients (create / update).
//
// A directed graph of `FlowNodeReq` connected by `FlowEdge`s. Validation enforces
// unique node/edge IDs, that `start_node_id` references a real node, and that
// every edge's endpoints reference real nodes.
//
// The properties Nodes, StartNodeID are required.
type UpdateAssistantConversationFlowParam struct {
	// All nodes in the flow. Must contain `start_node_id`. Each node is a prompt node
	// (`type: prompt`) or a tool node (`type: tool`).
	Nodes []UpdateAssistantConversationFlowNodeUnionParam `json:"nodes,omitzero" api:"required"`
	// ID of the node where the conversation begins.
	StartNodeID string `json:"start_node_id" api:"required"`
	// Directed transitions between nodes. May be empty for a single-node flow.
	Edges []UpdateAssistantConversationFlowEdgeParam `json:"edges,omitzero"`
	paramObj
}

func (r UpdateAssistantConversationFlowParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowNodeUnionParam struct {
	OfPrompt *UpdateAssistantConversationFlowNodePromptParam `json:",omitzero,inline"`
	OfTool   *UpdateAssistantConversationFlowNodeToolParam   `json:",omitzero,inline"`
	OfSpeak  *UpdateAssistantConversationFlowNodeSpeakParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPrompt, u.OfTool, u.OfSpeak)
}
func (u *UpdateAssistantConversationFlowNodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowNodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPrompt) {
		return u.OfPrompt
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfSpeak) {
		return u.OfSpeak
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetInstructions() *string {
	if vt := u.OfPrompt; vt != nil {
		return &vt.Instructions
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetExternalLlm() *ExternalLlmReqParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.ExternalLlm
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetInstructionsMode() *string {
	if vt := u.OfPrompt; vt != nil {
		return &vt.InstructionsMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetLlmAPIKeyRef() *string {
	if vt := u.OfPrompt; vt != nil && vt.LlmAPIKeyRef.Valid() {
		return &vt.LlmAPIKeyRef.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetModel() *string {
	if vt := u.OfPrompt; vt != nil && vt.Model.Valid() {
		return &vt.Model.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetSharedToolIDs() []string {
	if vt := u.OfPrompt; vt != nil {
		return vt.SharedToolIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetToolsMode() *string {
	if vt := u.OfPrompt; vt != nil {
		return &vt.ToolsMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetTranscription() *TranscriptionSettingsParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.Transcription
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetVoiceSettings() *VoiceSettingsParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.VoiceSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetSharedToolID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.SharedToolID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetMessage() *string {
	if vt := u.OfSpeak; vt != nil {
		return &vt.Message
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetID() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfSpeak; vt != nil {
		return (*string)(&vt.ID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetName() *string {
	if vt := u.OfPrompt; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfTool; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSpeak; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetType() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSpeak; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u UpdateAssistantConversationFlowNodeUnionParam) GetPosition() (res updateAssistantConversationFlowNodeUnionParamPosition) {
	if vt := u.OfPrompt; vt != nil {
		res.any = &vt.Position
	} else if vt := u.OfTool; vt != nil {
		res.any = &vt.Position
	} else if vt := u.OfSpeak; vt != nil {
		res.any = &vt.Position
	}
	return
}

// Can have the runtime types
// [*UpdateAssistantConversationFlowNodePromptPositionParam],
// [*UpdateAssistantConversationFlowNodeToolPositionParam],
// [*UpdateAssistantConversationFlowNodeSpeakPositionParam]
type updateAssistantConversationFlowNodeUnionParamPosition struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.UpdateAssistantConversationFlowNodePromptPositionParam:
//	case *telnyx.UpdateAssistantConversationFlowNodeToolPositionParam:
//	case *telnyx.UpdateAssistantConversationFlowNodeSpeakPositionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowNodeUnionParamPosition) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u updateAssistantConversationFlowNodeUnionParamPosition) GetX() *float64 {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowNodePromptPositionParam:
		return (*float64)(&vt.X)
	case *UpdateAssistantConversationFlowNodeToolPositionParam:
		return (*float64)(&vt.X)
	case *UpdateAssistantConversationFlowNodeSpeakPositionParam:
		return (*float64)(&vt.X)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u updateAssistantConversationFlowNodeUnionParamPosition) GetY() *float64 {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowNodePromptPositionParam:
		return (*float64)(&vt.Y)
	case *UpdateAssistantConversationFlowNodeToolPositionParam:
		return (*float64)(&vt.Y)
	case *UpdateAssistantConversationFlowNodeSpeakPositionParam:
		return (*float64)(&vt.Y)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UpdateAssistantConversationFlowNodeUnionParam](
		"type",
		apijson.Discriminator[UpdateAssistantConversationFlowNodePromptParam]("prompt"),
		apijson.Discriminator[UpdateAssistantConversationFlowNodeToolParam]("tool"),
		apijson.Discriminator[UpdateAssistantConversationFlowNodeSpeakParam]("speak"),
	)
}

// One step in a conversation flow, as supplied by API clients.
//
// Each node carries the prompt, tool scope, and optional overrides for
// model/voice/transcription. Unset overrides cascade from the assistant.
//
// The properties ID, Instructions are required.
type UpdateAssistantConversationFlowNodePromptParam struct {
	// Caller-supplied unique identifier for this node within the flow.
	ID string `json:"id" api:"required"`
	// Prompt that drives the LLM while this node is active. Required.
	Instructions string `json:"instructions" api:"required"`
	// Override for `Assistant.llm_api_key_ref` while this node is active. Part of the
	// LLM bundle — see `model` for cascade semantics.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Override for `Assistant.model` while this node is active. Part of the LLM bundle
	// (`model` + `llm_api_key_ref` + `external_llm`): when any of the three is set on
	// the node, all three are taken from the node and the assistant-level LLM identity
	// is not consulted. When none of the three is set, the assistant's bundle cascades
	// unchanged.
	Model param.Opt[string] `json:"model,omitzero"`
	// Optional human-readable label, displayed in authoring UIs.
	Name param.Opt[string] `json:"name,omitzero"`
	// Override for `Assistant.external_llm` while this node is active. Use this to
	// route a node's turns to a different external LLM (different `model`, `base_url`,
	// credentials). Part of the LLM bundle — see `model` for cascade semantics.
	// Mutually exclusive with `model` on the node (a single LLM identity per node).
	ExternalLlm ExternalLlmReqParam `json:"external_llm,omitzero"`
	// How `instructions` combine with the assistant-level instructions. `replace`
	// (default): the node's instructions are used alone. `append`: the node's
	// instructions are concatenated after the assistant's instructions.
	//
	// Any of "replace", "append".
	InstructionsMode string `json:"instructions_mode,omitzero"`
	// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
	// by the runtime; round-trips so frontends can persist graph layout across
	// reloads.
	Position UpdateAssistantConversationFlowNodePromptPositionParam `json:"position,omitzero"`
	// IDs of shared (org-level) tools available at this node. Knowledge bases are
	// attached the same way — via a shared retrieval tool. Tools not listed here are
	// not callable while this node is active.
	SharedToolIDs []string `json:"shared_tool_ids,omitzero"`
	// How `shared_tool_ids` combine with the assistant-level tool set. `replace`
	// (default): only the node's tools are callable. `append`: the node's tools are
	// added to the assistant's tools. Ignored when `shared_tool_ids` is null.
	//
	// Any of "replace", "append".
	ToolsMode string `json:"tools_mode,omitzero"`
	// Per-node transcription override (model/language/region). Unset fields cascade
	// from the assistant-level transcription.
	Transcription TranscriptionSettingsParam `json:"transcription,omitzero"`
	// Node kind discriminator. `prompt` (default) is an LLM-driven step; `tool` is a
	// standalone tool execution (see `ToolNodeReq`).
	//
	// Any of "prompt".
	Type string `json:"type,omitzero"`
	// Per-node voice override. Only fields set here override the assistant-level voice
	// settings; unset fields cascade.
	VoiceSettings VoiceSettingsParam `json:"voice_settings,omitzero"`
	paramObj
}

func (r UpdateAssistantConversationFlowNodePromptParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowNodePromptParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowNodePromptParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowNodePromptParam](
		"instructions_mode", "replace", "append",
	)
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowNodePromptParam](
		"tools_mode", "replace", "append",
	)
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowNodePromptParam](
		"type", "prompt",
	)
}

// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
// by the runtime; round-trips so frontends can persist graph layout across
// reloads.
//
// The properties X, Y are required.
type UpdateAssistantConversationFlowNodePromptPositionParam struct {
	// Horizontal coordinate in the authoring canvas.
	X float64 `json:"x" api:"required"`
	// Vertical coordinate in the authoring canvas.
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r UpdateAssistantConversationFlowNodePromptPositionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowNodePromptPositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowNodePromptPositionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A standalone tool step in a conversation flow, as supplied by clients.
//
// Unlike a prompt node, a tool node has no instructions or model — it isn't an LLM
// turn. Reaching it deterministically runs one shared tool (arguments filled from
// matching dynamic variables by name), then routes on the result via outgoing
// `tool_result` edges.
//
// The properties ID, SharedToolID are required.
type UpdateAssistantConversationFlowNodeToolParam struct {
	// Caller-supplied unique identifier for this node within the flow.
	ID string `json:"id" api:"required"`
	// ID of the single shared (org-level) tool this node executes. When the flow
	// reaches this node the tool runs as a deliberate step (no LLM turn); its outgoing
	// `tool_result` edges then route on the outcome. Arguments are filled from the
	// conversation's dynamic variables by name — a dynamic variable whose name matches
	// one of the tool's parameters supplies that argument. Cross-validated against the
	// org's shared tools on write.
	SharedToolID string `json:"shared_tool_id" api:"required"`
	// Optional human-readable label, displayed in authoring UIs.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
	// by the runtime; round-trips so frontends can persist graph layout across
	// reloads.
	Position UpdateAssistantConversationFlowNodeToolPositionParam `json:"position,omitzero"`
	// Node kind discriminator. Always `tool` for a tool node.
	//
	// Any of "tool".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r UpdateAssistantConversationFlowNodeToolParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowNodeToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowNodeToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowNodeToolParam](
		"type", "tool",
	)
}

// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
// by the runtime; round-trips so frontends can persist graph layout across
// reloads.
//
// The properties X, Y are required.
type UpdateAssistantConversationFlowNodeToolPositionParam struct {
	// Horizontal coordinate in the authoring canvas.
	X float64 `json:"x" api:"required"`
	// Vertical coordinate in the authoring canvas.
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r UpdateAssistantConversationFlowNodeToolPositionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowNodeToolPositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowNodeToolPositionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A standalone scripted-message step in a flow, as supplied by clients.
//
// Unlike a prompt node, a speak node has no instructions or model — it isn't an
// LLM turn. Reaching it delivers `message` to the user verbatim (with
// `{{variable}}` interpolation), then routes via outgoing `llm` / `expression`
// edges.
//
// The properties ID, Message are required.
type UpdateAssistantConversationFlowNodeSpeakParam struct {
	// Caller-supplied unique identifier for this node within the flow.
	ID string `json:"id" api:"required"`
	// Message delivered to the user verbatim when the flow reaches this node. No LLM
	// turn — the text is spoken/sent exactly as written. `{{variable}}` placeholders
	// are interpolated from the conversation's dynamic variables; an unresolved
	// placeholder renders as an empty string. After delivering, the flow routes via
	// the node's outgoing `llm` / `expression` edges (commonly a single unconditional
	// edge).
	Message string `json:"message" api:"required"`
	// Optional human-readable label, displayed in authoring UIs.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
	// by the runtime; round-trips so frontends can persist graph layout across
	// reloads.
	Position UpdateAssistantConversationFlowNodeSpeakPositionParam `json:"position,omitzero"`
	// Node kind discriminator. Always `speak` for a speak node.
	//
	// Any of "speak".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r UpdateAssistantConversationFlowNodeSpeakParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowNodeSpeakParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowNodeSpeakParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowNodeSpeakParam](
		"type", "speak",
	)
}

// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
// by the runtime; round-trips so frontends can persist graph layout across
// reloads.
//
// The properties X, Y are required.
type UpdateAssistantConversationFlowNodeSpeakPositionParam struct {
	// Horizontal coordinate in the authoring canvas.
	X float64 `json:"x" api:"required"`
	// Vertical coordinate in the authoring canvas.
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r UpdateAssistantConversationFlowNodeSpeakPositionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowNodeSpeakPositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowNodeSpeakPositionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directed transition from one node to a target, gated by a condition.
//
// The target is either another node in the same flow (`NodeTarget`) or a different
// assistant (`AssistantTarget`). Multiple edges may share a `start_node_id`; the
// runtime evaluates them in the order they're declared and takes the first whose
// condition is true.
//
// The properties ID, Condition, StartNodeID, Target are required.
type UpdateAssistantConversationFlowEdgeParam struct {
	// Caller-supplied unique identifier for this edge within the flow.
	ID string `json:"id" api:"required"`
	// Condition that gates the transition. Discriminated by `type`: `llm`,
	// `expression`.
	Condition UpdateAssistantConversationFlowEdgeConditionUnionParam `json:"condition,omitzero" api:"required"`
	// ID of the node this edge transitions away from.
	StartNodeID string `json:"start_node_id" api:"required"`
	// Destination of the transition. Discriminated by `type`: `node` (jump to another
	// node in this flow) or `assistant` (hand off to a different assistant).
	Target UpdateAssistantConversationFlowEdgeTargetUnionParam `json:"target,omitzero" api:"required"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionUnionParam struct {
	OfLlm        *UpdateAssistantConversationFlowEdgeConditionLlmParam        `json:",omitzero,inline"`
	OfExpression *UpdateAssistantConversationFlowEdgeConditionExpressionParam `json:",omitzero,inline"`
	OfDefault    *UpdateAssistantConversationFlowEdgeConditionDefaultParam    `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlm, u.OfExpression, u.OfDefault)
}
func (u *UpdateAssistantConversationFlowEdgeConditionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLlm) {
		return u.OfLlm
	} else if !param.IsOmitted(u.OfExpression) {
		return u.OfExpression
	} else if !param.IsOmitted(u.OfDefault) {
		return u.OfDefault
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) GetPrompt() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Prompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) GetExpression() *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam {
	if vt := u.OfExpression; vt != nil {
		return &vt.Expression
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) GetType() *string {
	if vt := u.OfLlm; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExpression; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDefault; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UpdateAssistantConversationFlowEdgeConditionUnionParam](
		"type",
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionLlmParam]("llm"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionParam]("expression"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionDefaultParam]("default"),
	)
}

// Edge condition evaluated by the LLM from a natural-language prompt.
//
// The model is asked to judge the prompt against conversation context and returns
// true/false. Use this for fuzzy intents that aren't expressible as a
// deterministic expression (e.g. 'user wants to escalate to a human').
//
// The properties Prompt, Type are required.
type UpdateAssistantConversationFlowEdgeConditionLlmParam struct {
	// Natural-language criterion the LLM judges as true/false.
	Prompt string `json:"prompt" api:"required"`
	// This field can be elided, and will marshal its zero value as "llm".
	Type constant.Llm `json:"type" default:"llm"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge condition evaluated as a deterministic expression AST.
//
// The expression is computed against runtime dynamic variables and must evaluate
// to a boolean. Prefer this over `LLMCondition` when the rule is a clean function
// of known variables — it's cheaper and predictable.
//
// The properties Expression, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionParam struct {
	// A node in a deterministic expression AST. Exactly one variant is selected by the
	// `type` discriminator. Terminal variants (`number_literal`, `string_literal`,
	// `bool_literal`, `variable`) bottom out the recursion; `arithmetic`, `bool_op`,
	// and `comparison` nest further sub-expressions.
	//
	// Extracted into a single named schema so the recursive union is defined once (was
	// previously inlined at every operand site).
	Expression UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam `json:"expression,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "expression".
	Type constant.Expression `json:"type" default:"expression"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam struct {
	OfDynamicVariableExpression *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionDynamicVariableExpressionParam `json:",omitzero,inline"`
	OfStringLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralExpressionParam   `json:",omitzero,inline"`
	OfNumberLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralExpressionParam   `json:",omitzero,inline"`
	OfBooleanLiteralExpression  *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBooleanLiteralExpressionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDynamicVariableExpression, u.OfStringLiteralExpression, u.OfNumberLiteralExpression, u.OfBooleanLiteralExpression)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDynamicVariableExpression) {
		return u.OfDynamicVariableExpression
	} else if !param.IsOmitted(u.OfStringLiteralExpression) {
		return u.OfStringLiteralExpression
	} else if !param.IsOmitted(u.OfNumberLiteralExpression) {
		return u.OfNumberLiteralExpression
	} else if !param.IsOmitted(u.OfBooleanLiteralExpression) {
		return u.OfBooleanLiteralExpression
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetName() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetType() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStringLiteralExpression; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNumberLiteralExpression; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBooleanLiteralExpression; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamValue) {
	if vt := u.OfStringLiteralExpression; vt != nil {
		res.any = &vt.Value
	} else if vt := u.OfNumberLiteralExpression; vt != nil {
		res.any = &vt.Value
	} else if vt := u.OfBooleanLiteralExpression; vt != nil {
		res.any = &vt.Value
	}
	return
}

// Can have the runtime types [*string], [*float64], [*bool]
type updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamValue) AsAny() any {
	return u.any
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionDynamicVariableExpressionParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionDynamicVariableExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionDynamicVariableExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionDynamicVariableExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralExpressionParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralExpressionParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBooleanLiteralExpressionParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBooleanLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBooleanLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBooleanLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewUpdateAssistantConversationFlowEdgeConditionDefaultParam() UpdateAssistantConversationFlowEdgeConditionDefaultParam {
	return UpdateAssistantConversationFlowEdgeConditionDefaultParam{
		Type: "default",
	}
}

// Fallback edge condition: fires only when no other edge's condition is true.
//
// Evaluated after every conditioned (`llm` / `expression`) edge regardless of
// declaration order, so it routes the flow whenever none of the node's other
// outgoing edges match. Valid **only** on edges leaving a `tool` or `speak` node,
// where the deterministic step auto-advances and must always have somewhere to go.
// A tool/speak node with any outgoing edge is required to carry exactly one
// `default` edge so it never dead-ends; a tool/speak node with no outgoing edges
// is a valid terminal step. Carries no parameters.
//
// This struct has a constant value, construct it with
// [NewUpdateAssistantConversationFlowEdgeConditionDefaultParam].
type UpdateAssistantConversationFlowEdgeConditionDefaultParam struct {
	Type constant.Default `json:"type" default:"default"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionDefaultParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionDefaultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionDefaultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeTargetUnionParam struct {
	OfNode      *UpdateAssistantConversationFlowEdgeTargetNodeParam      `json:",omitzero,inline"`
	OfAssistant *UpdateAssistantConversationFlowEdgeTargetAssistantParam `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeTargetUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNode, u.OfAssistant)
}
func (u *UpdateAssistantConversationFlowEdgeTargetUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeTargetUnionParam) asAny() any {
	if !param.IsOmitted(u.OfNode) {
		return u.OfNode
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeTargetUnionParam) GetNodeID() *string {
	if vt := u.OfNode; vt != nil {
		return &vt.NodeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeTargetUnionParam) GetAssistantID() *string {
	if vt := u.OfAssistant; vt != nil {
		return &vt.AssistantID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeTargetUnionParam) GetPosition() *UpdateAssistantConversationFlowEdgeTargetAssistantPositionParam {
	if vt := u.OfAssistant; vt != nil {
		return &vt.Position
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeTargetUnionParam) GetVoiceMode() *string {
	if vt := u.OfAssistant; vt != nil {
		return &vt.VoiceMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeTargetUnionParam) GetType() *string {
	if vt := u.OfNode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UpdateAssistantConversationFlowEdgeTargetUnionParam](
		"type",
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeTargetNodeParam]("node"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeTargetAssistantParam]("assistant"),
	)
}

// Edge target referencing another node within the same flow.
//
// The runtime transitions the active node to `node_id` and continues processing
// within the current assistant's flow.
//
// The properties NodeID, Type are required.
type UpdateAssistantConversationFlowEdgeTargetNodeParam struct {
	// ID of the node this edge transitions into.
	NodeID string `json:"node_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "node".
	Type constant.Node `json:"type" default:"node"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeTargetNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeTargetNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeTargetNodeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge target referencing a different assistant.
//
// When the edge fires, the conversation hands off to `assistant_id`: the active
// assistant on the conversation row is rewritten and the new assistant's flow
// starts at its own `start_node_id`. The current turn's LLM response is delivered
// to the user as-is; subsequent turns route to the new assistant.
//
// The properties AssistantID, Type are required.
type UpdateAssistantConversationFlowEdgeTargetAssistantParam struct {
	// ID of the assistant the conversation transitions to.
	AssistantID string `json:"assistant_id" api:"required"`
	// Optional canvas coordinates for rendering the target assistant as a node in
	// authoring UIs. Pure presentation — the runtime ignores it; round-trips so
	// frontends can persist graph layout across reloads. When multiple edges target
	// the same assistant, each edge's `position` is independent (frontends typically
	// use the first non-null one).
	Position UpdateAssistantConversationFlowEdgeTargetAssistantPositionParam `json:"position,omitzero"`
	// Voice behavior when handing off to the target assistant, mirroring the handoff
	// tool's `voice_mode`. `unified` (default) keeps the current voice across the
	// handoff; `distinct` lets the target assistant speak with its own configured
	// voice. Only applies to assistant targets — node targets override voice via the
	// node's own `voice_settings`.
	//
	// Any of "unified", "distinct".
	VoiceMode string `json:"voice_mode,omitzero"`
	// This field can be elided, and will marshal its zero value as "assistant".
	Type constant.Assistant `json:"type" default:"assistant"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeTargetAssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeTargetAssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeTargetAssistantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowEdgeTargetAssistantParam](
		"voice_mode", "unified", "distinct",
	)
}

// Optional canvas coordinates for rendering the target assistant as a node in
// authoring UIs. Pure presentation — the runtime ignores it; round-trips so
// frontends can persist graph layout across reloads. When multiple edges target
// the same assistant, each edge's `position` is independent (frontends typically
// use the first non-null one).
//
// The properties X, Y are required.
type UpdateAssistantConversationFlowEdgeTargetAssistantPositionParam struct {
	// Horizontal coordinate in the authoring canvas.
	X float64 `json:"x" api:"required"`
	// Vertical coordinate in the authoring canvas.
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeTargetAssistantPositionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeTargetAssistantPositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeTargetAssistantPositionParam) UnmarshalJSON(data []byte) error {
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
