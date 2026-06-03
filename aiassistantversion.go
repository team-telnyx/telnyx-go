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
	paramUnion
}

func (u UpdateAssistantConversationFlowNodeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPrompt, u.OfTool)
}
func (u *UpdateAssistantConversationFlowNodeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowNodeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPrompt) {
		return u.OfPrompt
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
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
func (u UpdateAssistantConversationFlowNodeUnionParam) GetID() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.ID)
	} else if vt := u.OfTool; vt != nil {
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowNodeUnionParam) GetType() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTool; vt != nil {
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
	}
	return
}

// Can have the runtime types
// [*UpdateAssistantConversationFlowNodePromptPositionParam],
// [*UpdateAssistantConversationFlowNodeToolPositionParam]
type updateAssistantConversationFlowNodeUnionParamPosition struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.UpdateAssistantConversationFlowNodePromptPositionParam:
//	case *telnyx.UpdateAssistantConversationFlowNodeToolPositionParam:
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
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UpdateAssistantConversationFlowNodeUnionParam](
		"type",
		apijson.Discriminator[UpdateAssistantConversationFlowNodePromptParam]("prompt"),
		apijson.Discriminator[UpdateAssistantConversationFlowNodeToolParam]("tool"),
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
	// `expression`, or `tool_result`. A `tool_result` condition is only valid on an
	// edge leaving a tool node.
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
	OfToolResult *UpdateAssistantConversationFlowEdgeConditionToolResultParam `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlm, u.OfExpression, u.OfToolResult)
}
func (u *UpdateAssistantConversationFlowEdgeConditionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLlm) {
		return u.OfLlm
	} else if !param.IsOmitted(u.OfExpression) {
		return u.OfExpression
	} else if !param.IsOmitted(u.OfToolResult) {
		return u.OfToolResult
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
func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) GetOutcome() *string {
	if vt := u.OfToolResult; vt != nil {
		return &vt.Outcome
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionUnionParam) GetType() *string {
	if vt := u.OfLlm; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExpression; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfToolResult; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UpdateAssistantConversationFlowEdgeConditionUnionParam](
		"type",
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionLlmParam]("llm"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionParam]("expression"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionToolResultParam]("tool_result"),
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
	// Root of the expression AST. Must evaluate to a boolean.
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
	OfComparison    *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam    `json:",omitzero,inline"`
	OfBoolOp        *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam        `json:",omitzero,inline"`
	OfArithmetic    *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam    `json:",omitzero,inline"`
	OfVariable      *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionVariableParam      `json:",omitzero,inline"`
	OfStringLiteral *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralParam `json:",omitzero,inline"`
	OfNumberLiteral *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralParam `json:",omitzero,inline"`
	OfBoolLiteral   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolLiteralParam   `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfComparison,
		u.OfBoolOp,
		u.OfArithmetic,
		u.OfVariable,
		u.OfStringLiteral,
		u.OfNumberLiteral,
		u.OfBoolLiteral)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfComparison) {
		return u.OfComparison
	} else if !param.IsOmitted(u.OfBoolOp) {
		return u.OfBoolOp
	} else if !param.IsOmitted(u.OfArithmetic) {
		return u.OfArithmetic
	} else if !param.IsOmitted(u.OfVariable) {
		return u.OfVariable
	} else if !param.IsOmitted(u.OfStringLiteral) {
		return u.OfStringLiteral
	} else if !param.IsOmitted(u.OfNumberLiteral) {
		return u.OfNumberLiteral
	} else if !param.IsOmitted(u.OfBoolLiteral) {
		return u.OfBoolLiteral
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetOperands() []UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam {
	if vt := u.OfBoolOp; vt != nil {
		return vt.Operands
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetName() *string {
	if vt := u.OfVariable; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetOp() *string {
	if vt := u.OfComparison; vt != nil {
		return (*string)(&vt.Op)
	} else if vt := u.OfBoolOp; vt != nil {
		return (*string)(&vt.Op)
	} else if vt := u.OfArithmetic; vt != nil {
		return (*string)(&vt.Op)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetType() *string {
	if vt := u.OfComparison; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBoolOp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfArithmetic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariable; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStringLiteral; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNumberLiteral; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBoolLiteral; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetLeft() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeft) {
	if vt := u.OfComparison; vt != nil {
		res.any = vt.Left.asAny()
	} else if vt := u.OfArithmetic; vt != nil {
		res.any = vt.Left.asAny()
	}
	return
}

// Can have the runtime types
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam]
type updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeft struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeft) AsAny() any {
	return u.any
}

// Returns a pointer to the underlying variant's property, if present.
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeft) GetName() *string {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam:
		return vt.GetName()
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam:
		return vt.GetName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeft) GetType() *string {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam:
		return vt.GetType()
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeft) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeftValue) {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam:
		res.any = vt.GetValue()
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam:
		res.any = vt.GetValue()
	}
	return res
}

// Can have the runtime types [*string], [*float64], [*bool]
type updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeftValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamLeftValue) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetRight() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRight) {
	if vt := u.OfComparison; vt != nil {
		res.any = vt.Right.asAny()
	} else if vt := u.OfArithmetic; vt != nil {
		res.any = vt.Right.asAny()
	}
	return
}

// Can have the runtime types
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam],
// [*UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam]
type updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRight struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam:
//	case *telnyx.UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRight) AsAny() any {
	return u.any
}

// Returns a pointer to the underlying variant's property, if present.
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRight) GetName() *string {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam:
		return vt.GetName()
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam:
		return vt.GetName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRight) GetType() *string {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam:
		return vt.GetType()
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam:
		return vt.GetType()
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRight) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRightValue) {
	switch vt := u.any.(type) {
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam:
		res.any = vt.GetValue()
	case *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam:
		res.any = vt.GetValue()
	}
	return res
}

// Can have the runtime types [*string], [*float64], [*bool]
type updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRightValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamRightValue) AsAny() any {
	return u.any
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParamValue) {
	if vt := u.OfStringLiteral; vt != nil {
		res.any = &vt.Value
	} else if vt := u.OfNumberLiteral; vt != nil {
		res.any = &vt.Value
	} else if vt := u.OfBoolLiteral; vt != nil {
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

func init() {
	apijson.RegisterUnion[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionUnionParam](
		"type",
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam]("comparison"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam]("bool_op"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam]("arithmetic"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionVariableParam]("variable"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralParam]("string_literal"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralParam]("number_literal"),
		apijson.Discriminator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolLiteralParam]("bool_literal"),
	)
}

// Compare two sub-expressions with a relational or membership operator.
//
// Evaluates to a boolean. Used in edge conditions to gate transitions on runtime
// values, e.g. `user_age >= 18` or `tier == "gold"`.
//
// The properties Left, Op, Right, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam struct {
	// Left-hand operand sub-expression.
	Left UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam `json:"left,omitzero" api:"required"`
	// Relational/membership operator. `contains` / `not_contains` apply to strings
	// (substring) and arrays (membership).
	//
	// Any of "==", "!=", "<", "<=", ">", ">=", "contains", "not_contains".
	Op string `json:"op,omitzero" api:"required"`
	// Right-hand operand sub-expression.
	Right UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam `json:"right,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "comparison".
	Type constant.Comparison `json:"type" default:"comparison"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonParam](
		"op", "==", "!=", "<", "<=", ">", ">=", "contains", "not_contains",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam struct {
	OfDynamicVariableExpression *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam `json:",omitzero,inline"`
	OfStringLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam   `json:",omitzero,inline"`
	OfNumberLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam   `json:",omitzero,inline"`
	OfBooleanLiteralExpression  *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDynamicVariableExpression, u.OfStringLiteralExpression, u.OfNumberLiteralExpression, u.OfBooleanLiteralExpression)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam) asAny() any {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam) GetName() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam) GetType() *string {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParamValue) {
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
type updateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParamValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftUnionParamValue) AsAny() any {
	return u.any
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftDynamicVariableExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftStringLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftNumberLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonLeftBooleanLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam struct {
	OfDynamicVariableExpression *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam `json:",omitzero,inline"`
	OfStringLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam   `json:",omitzero,inline"`
	OfNumberLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam   `json:",omitzero,inline"`
	OfBooleanLiteralExpression  *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDynamicVariableExpression, u.OfStringLiteralExpression, u.OfNumberLiteralExpression, u.OfBooleanLiteralExpression)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam) asAny() any {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam) GetName() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam) GetType() *string {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParamValue) {
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
type updateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParamValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightUnionParamValue) AsAny() any {
	return u.any
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightDynamicVariableExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightStringLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightNumberLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionComparisonRightBooleanLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Combine sub-expressions with a logical operator (`and` / `or` / `not`).
//
// `and` and `or` accept two or more operands; `not` accepts exactly one.
//
// The properties Op, Operands, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam struct {
	// Logical operator. `not` is unary; `and`/`or` are n-ary (>=2).
	//
	// Any of "and", "or", "not".
	Op string `json:"op,omitzero" api:"required"`
	// Operand sub-expressions. Length must be exactly 1 for `not` and >= 2 for
	// `and`/`or`.
	Operands []UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam `json:"operands,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_op".
	Type constant.BoolOp `json:"type" default:"bool_op"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpParam](
		"op", "and", "or", "not",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam struct {
	OfDynamicVariableExpression *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandDynamicVariableExpressionParam `json:",omitzero,inline"`
	OfStringLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandStringLiteralExpressionParam   `json:",omitzero,inline"`
	OfNumberLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandNumberLiteralExpressionParam   `json:",omitzero,inline"`
	OfBooleanLiteralExpression  *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandBooleanLiteralExpressionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDynamicVariableExpression, u.OfStringLiteralExpression, u.OfNumberLiteralExpression, u.OfBooleanLiteralExpression)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam) asAny() any {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam) GetName() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam) GetType() *string {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParamValue) {
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
type updateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParamValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandUnionParamValue) AsAny() any {
	return u.any
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandDynamicVariableExpressionParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandDynamicVariableExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandDynamicVariableExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandDynamicVariableExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandStringLiteralExpressionParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandStringLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandStringLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandStringLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandNumberLiteralExpressionParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandNumberLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandNumberLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandNumberLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandBooleanLiteralExpressionParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandBooleanLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandBooleanLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolOpOperandBooleanLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Numeric expression: applies an arithmetic operator to two sub-expressions.
//
// Useful for derived numeric checks, e.g. `cart_total + shipping > 50`. Both
// operands should resolve to numbers at runtime.
//
// The properties Left, Op, Right, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam struct {
	// Left-hand operand sub-expression.
	Left UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam `json:"left,omitzero" api:"required"`
	// Arithmetic operator applied to `left` and `right`.
	//
	// Any of "+", "-", "\*", "/", "%".
	Op string `json:"op,omitzero" api:"required"`
	// Right-hand operand sub-expression.
	Right UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam `json:"right,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "arithmetic".
	Type constant.Arithmetic `json:"type" default:"arithmetic"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticParam](
		"op", "+", "-", "*", "/", "%",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam struct {
	OfDynamicVariableExpression *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam `json:",omitzero,inline"`
	OfStringLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam   `json:",omitzero,inline"`
	OfNumberLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam   `json:",omitzero,inline"`
	OfBooleanLiteralExpression  *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDynamicVariableExpression, u.OfStringLiteralExpression, u.OfNumberLiteralExpression, u.OfBooleanLiteralExpression)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam) asAny() any {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam) GetName() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam) GetType() *string {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParamValue) {
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
type updateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParamValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftUnionParamValue) AsAny() any {
	return u.any
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftDynamicVariableExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftStringLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftNumberLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticLeftBooleanLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam struct {
	OfDynamicVariableExpression *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam `json:",omitzero,inline"`
	OfStringLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam   `json:",omitzero,inline"`
	OfNumberLiteralExpression   *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam   `json:",omitzero,inline"`
	OfBooleanLiteralExpression  *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam  `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDynamicVariableExpression, u.OfStringLiteralExpression, u.OfNumberLiteralExpression, u.OfBooleanLiteralExpression)
}
func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam) asAny() any {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam) GetName() *string {
	if vt := u.OfDynamicVariableExpression; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam) GetType() *string {
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
func (u UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParam) GetValue() (res updateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParamValue) {
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
type updateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParamValue struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *float64:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u updateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightUnionParamValue) AsAny() any {
	return u.any
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightDynamicVariableExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightStringLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightNumberLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionArithmeticRightBooleanLiteralExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference a dynamic variable by name.
//
// Resolved at runtime from the assistant's dynamic-variables context (see
// `Assistant.dynamic_variables` and the dynamic-variables webhook).
//
// The properties Name, Type are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionVariableParam struct {
	// Variable name to look up in the runtime context.
	Name string `json:"name" api:"required"`
	// This field can be elided, and will marshal its zero value as "variable".
	Type constant.Variable `json:"type" default:"variable"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant string value.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralParam struct {
	// Literal string value.
	Value string `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "string_literal".
	Type constant.StringLiteral `json:"type" default:"string_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionStringLiteralParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant numeric value (float; integers are accepted and stored as float).
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralParam struct {
	// Literal numeric value.
	Value float64 `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "number_literal".
	Type constant.NumberLiteral `json:"type" default:"number_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionNumberLiteralParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Constant boolean value. Useful for unconditional ('always') edges.
//
// The properties Type, Value are required.
type UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolLiteralParam struct {
	// Literal boolean value.
	Value bool `json:"value" api:"required"`
	// This field can be elided, and will marshal its zero value as "bool_literal".
	Type constant.BoolLiteral `json:"type" default:"bool_literal"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolLiteralParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolLiteralParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionExpressionExpressionBoolLiteralParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge condition that fires on the outcome of a tool node's execution.
//
// Only valid on edges leaving a tool node (`type == "tool"`). A tool node runs
// exactly one tool as a deliberate flow step; this condition routes on whether
// that tool reported `success` or `failure`. Use it to split the happy path from
// the error path after a tool runs (e.g. payment succeeded vs. declined). There is
// no `tool_id` field — the tool node has a single tool, so the outcome is
// unambiguous.
//
// The properties Outcome, Type are required.
type UpdateAssistantConversationFlowEdgeConditionToolResultParam struct {
	// Match either the tool node's success or failure outcome.
	//
	// Any of "success", "failure".
	Outcome string `json:"outcome,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "tool_result".
	Type constant.ToolResult `json:"type" default:"tool_result"`
	paramObj
}

func (r UpdateAssistantConversationFlowEdgeConditionToolResultParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantConversationFlowEdgeConditionToolResultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantConversationFlowEdgeConditionToolResultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateAssistantConversationFlowEdgeConditionToolResultParam](
		"outcome", "success", "failure",
	)
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
