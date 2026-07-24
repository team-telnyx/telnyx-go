// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
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
	"github.com/team-telnyx/telnyx-go/v4/shared"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Configure AI assistant specifications
//
// AIAssistantService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantService] method instead.
type AIAssistantService struct {
	Options []option.RequestOption
	// Configure AI assistant specifications
	Tests AIAssistantTestService
	// Configure AI assistant specifications
	CanaryDeploys AIAssistantCanaryDeployService
	// Configure AI assistant specifications
	ScheduledEvents AIAssistantScheduledEventService
	// Configure AI assistant specifications
	Tools AIAssistantToolService
	// Configure AI assistant specifications
	Versions AIAssistantVersionService
	// Configure AI assistant specifications
	Tags AIAssistantTagService
	// Configure AI assistant specifications
	Instructions AIAssistantInstructionService
}

// NewAIAssistantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAssistantService(opts ...option.RequestOption) (r AIAssistantService) {
	r = AIAssistantService{}
	r.Options = opts
	r.Tests = NewAIAssistantTestService(opts...)
	r.CanaryDeploys = NewAIAssistantCanaryDeployService(opts...)
	r.ScheduledEvents = NewAIAssistantScheduledEventService(opts...)
	r.Tools = NewAIAssistantToolService(opts...)
	r.Versions = NewAIAssistantVersionService(opts...)
	r.Tags = NewAIAssistantTagService(opts...)
	r.Instructions = NewAIAssistantInstructionService(opts...)
	return
}

// Create a new AI Assistant.
func (r *AIAssistantService) New(ctx context.Context, body AIAssistantNewParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve an AI Assistant configuration by `assistant_id`.
func (r *AIAssistantService) Get(ctx context.Context, assistantID string, query AIAssistantGetParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update an AI Assistant's attributes.
func (r *AIAssistantService) Update(ctx context.Context, assistantID string, body AIAssistantUpdateParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a list of all AI Assistants configured by the user.
func (r *AIAssistantService) List(ctx context.Context, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an AI Assistant by `assistant_id`.
func (r *AIAssistantService) Delete(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *AIAssistantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows a client to send a chat message to a specific AI Assistant.
// The assistant processes the message and returns a relevant reply based on the
// current conversation context. Refer to the Conversation API to
// [create a conversation](https://developers.telnyx.com/api-reference/conversations/create-a-conversation),
// [filter existing conversations](https://developers.telnyx.com/api-reference/conversations/list-conversations),
// [fetch messages for a conversation](https://developers.telnyx.com/api-reference/conversations/get-conversation-messages),
// and
// [manually add messages to a conversation](https://developers.telnyx.com/api-reference/conversations/create-message).
func (r *AIAssistantService) Chat(ctx context.Context, assistantID string, body AIAssistantChatParams, opts ...option.RequestOption) (res *AIAssistantChatResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/chat", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Clone an existing assistant, excluding telephony and messaging settings.
func (r *AIAssistantService) Clone(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/clone", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get an assistant texml by `assistant_id`.
func (r *AIAssistantService) GetTexml(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/texml", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Import assistants from external providers. Any assistant that has already been
// imported will be overwritten with its latest version from the importing
// provider.
func (r *AIAssistantService) Imports(ctx context.Context, body AIAssistantImportsParams, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send an SMS message for an assistant. This endpoint:
//
//  1. Validates the assistant exists and has messaging profile configured
//  2. If should_create_conversation is true, creates a new conversation with
//     metadata
//  3. Sends the SMS message (If `text` is set, this will be sent. Otherwise, if
//     this is the first message in the conversation and the assistant has a
//     `greeting` configured, this will be sent. Otherwise the assistant will
//     generate the text to send.)
//  4. Updates conversation metadata if provided
//  5. Returns the conversation ID
func (r *AIAssistantService) SendSMS(ctx context.Context, assistantID string, body AIAssistantSendSMSParams, opts ...option.RequestOption) (res *AIAssistantSendSMSResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/chat/sms", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Assistant configuration including choice of LLM, custom instructions, and tools.
type AssistantParam struct {
	// The system instructions that the voice assistant uses during the gather command
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// The model to be used by the voice assistant.
	Model param.Opt[string] `json:"model,omitzero"`
	// This is necessary only if the model selected is from OpenAI. You would pass the
	// `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your OpenAI API Key. Warning: Free plans are unlikely to work
	// with this integration.
	OpenAIAPIKeyRef param.Opt[string] `json:"openai_api_key_ref,omitzero"`
	// The tools that the voice assistant can use.
	Tools []AssistantToolsUnionParam `json:"tools,omitzero"`
	paramObj
}

func (r AssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantToolsUnionParam struct {
	OfBookAppointment   *shared.BookAppointmentToolParam      `json:",omitzero,inline"`
	OfCheckAvailability *shared.CheckAvailabilityToolParam    `json:",omitzero,inline"`
	OfWebhook           *WebhookToolParam                     `json:",omitzero,inline"`
	OfHangup            *HangupToolParam                      `json:",omitzero,inline"`
	OfTransfer          *TransferToolParam                    `json:",omitzero,inline"`
	OfRetrieval         *shared.CallControlRetrievalToolParam `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBookAppointment,
		u.OfCheckAvailability,
		u.OfWebhook,
		u.OfHangup,
		u.OfTransfer,
		u.OfRetrieval)
}
func (u *AssistantToolsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBookAppointment) {
		return u.OfBookAppointment
	} else if !param.IsOmitted(u.OfCheckAvailability) {
		return u.OfCheckAvailability
	} else if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfHangup) {
		return u.OfHangup
	} else if !param.IsOmitted(u.OfTransfer) {
		return u.OfTransfer
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetBookAppointment() *shared.BookAppointmentToolParams {
	if vt := u.OfBookAppointment; vt != nil {
		return &vt.BookAppointment
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetCheckAvailability() *shared.CheckAvailabilityToolParams {
	if vt := u.OfCheckAvailability; vt != nil {
		return &vt.CheckAvailability
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetWebhook() *WebhookToolWebhookParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetHangup() *HangupToolParams {
	if vt := u.OfHangup; vt != nil {
		return &vt.Hangup
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetTransfer() *TransferToolTransferParam {
	if vt := u.OfTransfer; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetRetrieval() *shared.CallControlBucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsUnionParam) GetType() *string {
	if vt := u.OfBookAppointment; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCheckAvailability; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHangup; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTransfer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AssistantToolsUnionParam](
		"type",
		apijson.Discriminator[shared.BookAppointmentToolParam]("book_appointment"),
		apijson.Discriminator[shared.CheckAvailabilityToolParam]("check_availability"),
		apijson.Discriminator[WebhookToolParam]("webhook"),
		apijson.Discriminator[HangupToolParam]("hangup"),
		apijson.Discriminator[TransferToolParam]("transfer"),
		apijson.Discriminator[shared.CallControlRetrievalToolParam]("retrieval"),
	)
}

// Reference to a connected integration attached to an assistant. Discover
// available integrations with `/ai/integrations` and connected integrations with
// `/ai/integrations/connections`.
type AssistantIntegration struct {
	// Catalog integration ID to attach. This is the `id` from the integrations catalog
	// at `/ai/integrations` (the same value also appears as `integration_id` on
	// entries returned by `/ai/integrations/connections`). It is **not** the
	// connection-level `id` from `/ai/integrations/connections`.
	IntegrationID string `json:"integration_id" api:"required"`
	// Optional per-assistant allowlist of integration tool names. When omitted or
	// empty, all tools allowed by the connected integration are available to the
	// assistant.
	AllowedList []string `json:"allowed_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntegrationID respjson.Field
		AllowedList   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantIntegration) RawJSON() string { return r.JSON.raw }
func (r *AssistantIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AssistantIntegration to a AssistantIntegrationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AssistantIntegrationParam.Overrides()
func (r AssistantIntegration) ToParam() AssistantIntegrationParam {
	return param.Override[AssistantIntegrationParam](json.RawMessage(r.RawJSON()))
}

// Reference to a connected integration attached to an assistant. Discover
// available integrations with `/ai/integrations` and connected integrations with
// `/ai/integrations/connections`.
//
// The property IntegrationID is required.
type AssistantIntegrationParam struct {
	// Catalog integration ID to attach. This is the `id` from the integrations catalog
	// at `/ai/integrations` (the same value also appears as `integration_id` on
	// entries returned by `/ai/integrations/connections`). It is **not** the
	// connection-level `id` from `/ai/integrations/connections`.
	IntegrationID string `json:"integration_id" api:"required"`
	// Optional per-assistant allowlist of integration tool names. When omitted or
	// empty, all tools allowed by the connected integration are available to the
	// assistant.
	AllowedList []string `json:"allowed_list,omitzero"`
	paramObj
}

func (r AssistantIntegrationParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantIntegrationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantIntegrationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to an MCP server attached to an assistant. Create and manage MCP
// servers with the `/ai/mcp_servers` endpoints, then attach them to assistants by
// ID.
type AssistantMcpServer struct {
	// ID of the MCP server to attach. This must be the `id` of an MCP server returned
	// by the `/ai/mcp_servers` endpoints.
	ID string `json:"id" api:"required"`
	// Optional per-assistant allowlist of MCP tool names. When omitted, the assistant
	// uses the MCP server's configured `allowed_tools`.
	AllowedTools []string `json:"allowed_tools"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AllowedTools respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMcpServer) RawJSON() string { return r.JSON.raw }
func (r *AssistantMcpServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AssistantMcpServer to a AssistantMcpServerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AssistantMcpServerParam.Overrides()
func (r AssistantMcpServer) ToParam() AssistantMcpServerParam {
	return param.Override[AssistantMcpServerParam](json.RawMessage(r.RawJSON()))
}

// Reference to an MCP server attached to an assistant. Create and manage MCP
// servers with the `/ai/mcp_servers` endpoints, then attach them to assistants by
// ID.
//
// The property ID is required.
type AssistantMcpServerParam struct {
	// ID of the MCP server to attach. This must be the `id` of an MCP server returned
	// by the `/ai/mcp_servers` endpoints.
	ID string `json:"id" api:"required"`
	// Optional per-assistant allowlist of MCP tool names. When omitted, the assistant
	// uses the MCP server's configured `allowed_tools`.
	AllowedTools []string `json:"allowed_tools,omitzero"`
	paramObj
}

func (r AssistantMcpServerParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantMcpServerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantMcpServerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantToolUnion contains all possible properties and values from
// [InferenceEmbeddingWebhookToolParamsResp], [AssistantToolClientSideTool],
// [RetrievalTool], [AssistantToolHandoff], [HangupTool], [AssistantToolTransfer],
// [AssistantToolInvite], [AssistantToolRefer], [AssistantToolSendDtmf],
// [AssistantToolSendMessage], [AssistantToolSkipTurn], [AssistantToolPay].
//
// Use the [AssistantToolUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantToolUnion struct {
	// Any of "webhook", "client_side_tool", "retrieval", "handoff", "hangup",
	// "transfer", "invite", "refer", "send_dtmf", "send_message", "skip_turn", "pay".
	Type string `json:"type"`
	// This field is from variant [InferenceEmbeddingWebhookToolParamsResp].
	Webhook InferenceEmbeddingWebhookToolParamsWebhookResp `json:"webhook"`
	// This field is from variant [AssistantToolClientSideTool].
	ClientSideTool AssistantToolClientSideToolClientSideTool `json:"client_side_tool"`
	// This field is from variant [RetrievalTool].
	Retrieval BucketIDs `json:"retrieval"`
	// This field is from variant [AssistantToolHandoff].
	Handoff AssistantToolHandoffHandoff `json:"handoff"`
	// This field is from variant [HangupTool].
	Hangup HangupToolParamsResp `json:"hangup"`
	// This field is from variant [AssistantToolTransfer].
	Transfer AssistantToolTransferTransfer `json:"transfer"`
	// This field is from variant [AssistantToolInvite].
	Invite AssistantToolInviteInvite `json:"invite"`
	// This field is from variant [AssistantToolRefer].
	Refer AssistantToolReferRefer `json:"refer"`
	// This field is from variant [AssistantToolSendDtmf].
	SendDtmf map[string]any `json:"send_dtmf"`
	// This field is from variant [AssistantToolSendMessage].
	SendMessage AssistantToolSendMessageSendMessage `json:"send_message"`
	// This field is from variant [AssistantToolSkipTurn].
	SkipTurn AssistantToolSkipTurnSkipTurn `json:"skip_turn"`
	// This field is from variant [AssistantToolPay].
	Pay  PayToolParamsResp `json:"pay"`
	JSON struct {
		Type           respjson.Field
		Webhook        respjson.Field
		ClientSideTool respjson.Field
		Retrieval      respjson.Field
		Handoff        respjson.Field
		Hangup         respjson.Field
		Transfer       respjson.Field
		Invite         respjson.Field
		Refer          respjson.Field
		SendDtmf       respjson.Field
		SendMessage    respjson.Field
		SkipTurn       respjson.Field
		Pay            respjson.Field
		raw            string
	} `json:"-"`
}

// anyAssistantTool is implemented by each variant of [AssistantToolUnion] to add
// type safety for the return type of [AssistantToolUnion.AsAny]
type anyAssistantTool interface {
	implAssistantToolUnion()
}

func (InferenceEmbeddingWebhookToolParamsResp) implAssistantToolUnion() {}
func (AssistantToolClientSideTool) implAssistantToolUnion()             {}
func (RetrievalTool) implAssistantToolUnion()                           {}
func (AssistantToolHandoff) implAssistantToolUnion()                    {}
func (HangupTool) implAssistantToolUnion()                              {}
func (AssistantToolTransfer) implAssistantToolUnion()                   {}
func (AssistantToolInvite) implAssistantToolUnion()                     {}
func (AssistantToolRefer) implAssistantToolUnion()                      {}
func (AssistantToolSendDtmf) implAssistantToolUnion()                   {}
func (AssistantToolSendMessage) implAssistantToolUnion()                {}
func (AssistantToolSkipTurn) implAssistantToolUnion()                   {}
func (AssistantToolPay) implAssistantToolUnion()                        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssistantToolUnion.AsAny().(type) {
//	case telnyx.InferenceEmbeddingWebhookToolParamsResp:
//	case telnyx.AssistantToolClientSideTool:
//	case telnyx.RetrievalTool:
//	case telnyx.AssistantToolHandoff:
//	case telnyx.HangupTool:
//	case telnyx.AssistantToolTransfer:
//	case telnyx.AssistantToolInvite:
//	case telnyx.AssistantToolRefer:
//	case telnyx.AssistantToolSendDtmf:
//	case telnyx.AssistantToolSendMessage:
//	case telnyx.AssistantToolSkipTurn:
//	case telnyx.AssistantToolPay:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AssistantToolUnion) AsAny() anyAssistantTool {
	switch u.Type {
	case "webhook":
		return u.AsWebhook()
	case "client_side_tool":
		return u.AsClientSideTool()
	case "retrieval":
		return u.AsRetrieval()
	case "handoff":
		return u.AsHandoff()
	case "hangup":
		return u.AsHangup()
	case "transfer":
		return u.AsTransfer()
	case "invite":
		return u.AsInvite()
	case "refer":
		return u.AsRefer()
	case "send_dtmf":
		return u.AsSendDtmf()
	case "send_message":
		return u.AsSendMessage()
	case "skip_turn":
		return u.AsSkipTurn()
	case "pay":
		return u.AsPay()
	}
	return nil
}

func (u AssistantToolUnion) AsWebhook() (v InferenceEmbeddingWebhookToolParamsResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsClientSideTool() (v AssistantToolClientSideTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsRetrieval() (v RetrievalTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsHandoff() (v AssistantToolHandoff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsHangup() (v HangupTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsTransfer() (v AssistantToolTransfer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsInvite() (v AssistantToolInvite) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsRefer() (v AssistantToolRefer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsSendDtmf() (v AssistantToolSendDtmf) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsSendMessage() (v AssistantToolSendMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsSkipTurn() (v AssistantToolSkipTurn) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsPay() (v AssistantToolPay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantToolUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AssistantToolUnion to a AssistantToolUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AssistantToolUnionParam.Overrides()
func (r AssistantToolUnion) ToParam() AssistantToolUnionParam {
	return param.Override[AssistantToolUnionParam](json.RawMessage(r.RawJSON()))
}

type AssistantToolClientSideTool struct {
	ClientSideTool AssistantToolClientSideToolClientSideTool `json:"client_side_tool" api:"required"`
	Type           constant.ClientSideTool                   `json:"type" default:"client_side_tool"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientSideTool respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolClientSideTool) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolClientSideTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolClientSideToolClientSideTool struct {
	// The description of the tool.
	Description string `json:"description" api:"required"`
	// The name of the tool.
	Name string `json:"name" api:"required"`
	// The parameters the tool accepts, described as a JSON Schema object. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	Parameters AssistantToolClientSideToolClientSideToolParameters `json:"parameters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolClientSideToolClientSideTool) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolClientSideToolClientSideTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters the tool accepts, described as a JSON Schema object. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type AssistantToolClientSideToolClientSideToolParameters struct {
	// The properties of the parameters.
	Properties map[string]any `json:"properties"`
	// The required properties of the parameters.
	Required []string `json:"required"`
	// Any of "object".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolClientSideToolClientSideToolParameters) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolClientSideToolClientSideToolParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The handoff tool allows the assistant to hand off control of the conversation to
// another AI assistant. By default, this will happen transparently to the end
// user.
type AssistantToolHandoff struct {
	Handoff AssistantToolHandoffHandoff `json:"handoff" api:"required"`
	Type    constant.Handoff            `json:"type" default:"handoff"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handoff     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolHandoff) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolHandoff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolHandoffHandoff struct {
	// List of possible assistants that can receive a handoff.
	AIAssistants []AssistantToolHandoffHandoffAIAssistant `json:"ai_assistants" api:"required"`
	// With the unified voice mode all assistants share the same voice, making the
	// handoff transparent to the user. With the distinct voice mode all assistants
	// retain their voice configuration, providing the experience of a conference call
	// with a team of assistants.
	//
	// Any of "unified", "distinct".
	VoiceMode string `json:"voice_mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIAssistants respjson.Field
		VoiceMode    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolHandoffHandoff) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolHandoffHandoff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolHandoffHandoffAIAssistant struct {
	// The ID of the assistant to hand off to.
	ID string `json:"id" api:"required"`
	// Helpful name for giving context on when to handoff to the assistant.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolHandoffHandoffAIAssistant) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolHandoffHandoffAIAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolTransfer struct {
	Transfer AssistantToolTransferTransfer `json:"transfer" api:"required"`
	Type     constant.Transfer             `json:"type" default:"transfer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transfer    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransfer) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolTransferTransfer struct {
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to. This can also be a dynamic
	// variable string like `{{ targets }}` where `targets` is returned by the dynamic
	// variables webhook and resolves to an array of target objects at runtime.
	Targets AssistantToolTransferTransferTargetsUnion `json:"targets" api:"required"`
	// Custom headers to be added to the SIP INVITE for the transfer command.
	CustomHeaders []AssistantToolTransferTransferCustomHeader `json:"custom_headers"`
	// A description of the transfer tool. By default, Telnyx generates this
	// automatically based on the configured targets. Typically only set when importing
	// an assistant from another provider that allowed a custom description; in that
	// case the provided value is preserved. Most users should leave this empty and let
	// Telnyx manage it.
	Description string `json:"description"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
	// transferred call. Allows the assistant to detect when a voicemail system answers
	// the transferred call and take appropriate action.
	VoicemailDetection AssistantToolTransferTransferVoicemailDetection `json:"voicemail_detection"`
	// Optional delay in milliseconds before playing the warm message audio when the
	// transferred call is answered. When set, the audio_url is not included in the
	// dial command; instead, playback starts after the specified delay. When not set,
	// existing behavior (audio_url in dial) is preserved.
	WarmMessageDelayMs int64 `json:"warm_message_delay_ms" api:"nullable"`
	// Natural language instructions for your agent for how to provide context for the
	// transfer recipient.
	WarmTransferInstructions string `json:"warm_transfer_instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From                     respjson.Field
		Targets                  respjson.Field
		CustomHeaders            respjson.Field
		Description              respjson.Field
		VoicemailDetection       respjson.Field
		WarmMessageDelayMs       respjson.Field
		WarmTransferInstructions respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransfer) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolTransferTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantToolTransferTransferTargetsUnion contains all possible properties and
// values from [[]AssistantToolTransferTransferTargetsTargetsListItem], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTargetsList OfString]
type AssistantToolTransferTransferTargetsUnion struct {
	// This field will be present if the value is a
	// [[]AssistantToolTransferTransferTargetsTargetsListItem] instead of an object.
	OfTargetsList []AssistantToolTransferTransferTargetsTargetsListItem `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTargetsList respjson.Field
		OfString      respjson.Field
		raw           string
	} `json:"-"`
}

func (u AssistantToolTransferTransferTargetsUnion) AsTargetsList() (v []AssistantToolTransferTransferTargetsTargetsListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolTransferTransferTargetsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantToolTransferTransferTargetsUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantToolTransferTransferTargetsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolTransferTransferTargetsTargetsListItem struct {
	// The destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The name of the target.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		To          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferTargetsTargetsListItem) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolTransferTransferTargetsTargetsListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolTransferTransferCustomHeader struct {
	Name string `json:"name"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferCustomHeader) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolTransferTransferCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
// transferred call. Allows the assistant to detect when a voicemail system answers
// the transferred call and take appropriate action.
type AssistantToolTransferTransferVoicemailDetection struct {
	// Advanced AMD detection configuration parameters. All values are optional -
	// Telnyx will use defaults if not specified.
	DetectionConfig AssistantToolTransferTransferVoicemailDetectionDetectionConfig `json:"detection_config"`
	// The AMD detection mode to use. 'premium' enables premium answering machine
	// detection. 'disabled' turns off AMD detection.
	//
	// Any of "disabled", "premium".
	DetectionMode string `json:"detection_mode"`
	// Action to take when voicemail is detected on the transferred call.
	OnVoicemailDetected AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetected `json:"on_voicemail_detected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionConfig     respjson.Field
		DetectionMode       respjson.Field
		OnVoicemailDetected respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferVoicemailDetection) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolTransferTransferVoicemailDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced AMD detection configuration parameters. All values are optional -
// Telnyx will use defaults if not specified.
type AssistantToolTransferTransferVoicemailDetectionDetectionConfig struct {
	// Duration of silence after greeting detection before finalizing the result.
	AfterGreetingSilenceMillis int64 `json:"after_greeting_silence_millis"`
	// Maximum silence duration between words during greeting.
	BetweenWordsSilenceMillis int64 `json:"between_words_silence_millis"`
	// Expected duration of greeting speech.
	GreetingDurationMillis int64 `json:"greeting_duration_millis"`
	// Duration of silence after the greeting to wait before considering the greeting
	// complete.
	GreetingSilenceDurationMillis int64 `json:"greeting_silence_duration_millis"`
	// Maximum time to spend analyzing the greeting.
	GreetingTotalAnalysisTimeMillis int64 `json:"greeting_total_analysis_time_millis"`
	// Maximum silence duration at the start of the call before speech.
	InitialSilenceMillis int64 `json:"initial_silence_millis"`
	// Maximum number of words expected in a human greeting.
	MaximumNumberOfWords int64 `json:"maximum_number_of_words"`
	// Maximum duration of a single word.
	MaximumWordLengthMillis int64 `json:"maximum_word_length_millis"`
	// Minimum duration for audio to be considered a word.
	MinWordLengthMillis int64 `json:"min_word_length_millis"`
	// Audio level threshold for silence detection.
	SilenceThreshold int64 `json:"silence_threshold"`
	// Total time allowed for AMD analysis.
	TotalAnalysisTimeMillis int64 `json:"total_analysis_time_millis"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AfterGreetingSilenceMillis      respjson.Field
		BetweenWordsSilenceMillis       respjson.Field
		GreetingDurationMillis          respjson.Field
		GreetingSilenceDurationMillis   respjson.Field
		GreetingTotalAnalysisTimeMillis respjson.Field
		InitialSilenceMillis            respjson.Field
		MaximumNumberOfWords            respjson.Field
		MaximumWordLengthMillis         respjson.Field
		MinWordLengthMillis             respjson.Field
		SilenceThreshold                respjson.Field
		TotalAnalysisTimeMillis         respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferVoicemailDetectionDetectionConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *AssistantToolTransferTransferVoicemailDetectionDetectionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take when voicemail is detected on the transferred call.
type AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetected struct {
	// The action to take when voicemail is detected. 'stop_transfer' hangs up
	// immediately. 'leave_message_and_stop_transfer' leaves a message then hangs up.
	//
	// Any of "stop_transfer", "leave_message_and_stop_transfer".
	Action string `json:"action"`
	// Configuration for the voicemail message to leave. Only applicable when action is
	// 'leave_message_and_stop_transfer'.
	VoicemailMessage AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessage `json:"voicemail_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action           respjson.Field
		VoicemailMessage respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetected) RawJSON() string {
	return r.JSON.raw
}
func (r *AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the voicemail message to leave. Only applicable when action is
// 'leave_message_and_stop_transfer'.
type AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessage struct {
	// The specific message to leave as voicemail (converted to speech). Only
	// applicable when type is 'message'.
	Message string `json:"message"`
	// The type of voicemail message. Use 'message' to leave a specific TTS message, or
	// 'warm_transfer_instructions' to play the warm transfer audio.
	//
	// Any of "message", "warm_transfer_instructions".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolInvite struct {
	Invite AssistantToolInviteInvite `json:"invite" api:"required"`
	Type   constant.Invite           `json:"type" default:"invite"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invite      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolInvite) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolInviteInvite struct {
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// Custom headers to be added to the SIP INVITE for the invite command.
	CustomHeaders []AssistantToolInviteInviteCustomHeader `json:"custom_headers"`
	// The different possible targets of the invite. The assistant will be able to
	// choose one of the targets to invite to the call. This can also be a dynamic
	// variable string like `{{ targets }}` where `targets` is returned by the dynamic
	// variables webhook and resolves to an array of target objects at runtime. If
	// omitted or null, the invite tool can still be configured and targets may be
	// supplied dynamically at runtime.
	Targets AssistantToolInviteInviteTargetsUnion `json:"targets" api:"nullable"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
	// invited call.
	VoicemailDetection AssistantToolInviteInviteVoicemailDetection `json:"voicemail_detection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From               respjson.Field
		CustomHeaders      respjson.Field
		Targets            respjson.Field
		VoicemailDetection respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolInviteInvite) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolInviteInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolInviteInviteCustomHeader struct {
	Name string `json:"name"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolInviteInviteCustomHeader) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolInviteInviteCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantToolInviteInviteTargetsUnion contains all possible properties and
// values from [[]AssistantToolInviteInviteTargetsTargetsListItem], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTargetsList OfString]
type AssistantToolInviteInviteTargetsUnion struct {
	// This field will be present if the value is a
	// [[]AssistantToolInviteInviteTargetsTargetsListItem] instead of an object.
	OfTargetsList []AssistantToolInviteInviteTargetsTargetsListItem `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTargetsList respjson.Field
		OfString      respjson.Field
		raw           string
	} `json:"-"`
}

func (u AssistantToolInviteInviteTargetsUnion) AsTargetsList() (v []AssistantToolInviteInviteTargetsTargetsListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolInviteInviteTargetsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantToolInviteInviteTargetsUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantToolInviteInviteTargetsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolInviteInviteTargetsTargetsListItem struct {
	// The destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The name of the target.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		To          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolInviteInviteTargetsTargetsListItem) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolInviteInviteTargetsTargetsListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
// invited call.
type AssistantToolInviteInviteVoicemailDetection struct {
	// The AMD detection mode to use. 'premium' enables premium answering machine
	// detection. 'disabled' turns off AMD detection.
	//
	// Any of "disabled", "premium".
	DetectionMode string `json:"detection_mode"`
	// Action to take when voicemail is detected on the invited call.
	OnVoicemailDetected AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetected `json:"on_voicemail_detected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionMode       respjson.Field
		OnVoicemailDetected respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolInviteInviteVoicemailDetection) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolInviteInviteVoicemailDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take when voicemail is detected on the invited call.
type AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetected struct {
	// The action to take when voicemail is detected.
	//
	// Any of "stop_invite".
	Action string `json:"action"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetected) RawJSON() string {
	return r.JSON.raw
}
func (r *AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolRefer struct {
	Refer AssistantToolReferRefer `json:"refer" api:"required"`
	Type  constant.Refer          `json:"type" default:"refer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refer       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolRefer) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolRefer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolReferRefer struct {
	// The different possible targets of the SIP refer. The assistant will be able to
	// choose one of the targets to refer the call to.
	Targets []AssistantToolReferReferTarget `json:"targets" api:"required"`
	// Custom headers to be added to the SIP REFER.
	CustomHeaders []AssistantToolReferReferCustomHeader `json:"custom_headers"`
	// SIP headers to be added to the SIP REFER. Currently only User-to-User and
	// Diversion headers are supported.
	SipHeaders []AssistantToolReferReferSipHeader `json:"sip_headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Targets       respjson.Field
		CustomHeaders respjson.Field
		SipHeaders    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolReferRefer) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolReferRefer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolReferReferTarget struct {
	// The name of the target.
	Name string `json:"name" api:"required"`
	// The SIP URI to which the call will be referred.
	SipAddress string `json:"sip_address" api:"required"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword string `json:"sip_auth_password"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername string `json:"sip_auth_username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		SipAddress      respjson.Field
		SipAuthPassword respjson.Field
		SipAuthUsername respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolReferReferTarget) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolReferReferTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolReferReferCustomHeader struct {
	Name string `json:"name"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolReferReferCustomHeader) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolReferReferCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolReferReferSipHeader struct {
	// Any of "User-to-User", "Diversion".
	Name string `json:"name"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolReferReferSipHeader) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolReferReferSipHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSendDtmf struct {
	SendDtmf map[string]any    `json:"send_dtmf" api:"required"`
	Type     constant.SendDtmf `json:"type" default:"send_dtmf"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SendDtmf    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolSendDtmf) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSendDtmf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The send_message tool allows the assistant to send SMS or MMS messages to the
// end user. The 'to' and 'from' addresses are automatically determined from the
// conversation context, and the message text is generated by the assistant unless
// a message_template is provided for runtime variable substitution.
type AssistantToolSendMessage struct {
	SendMessage AssistantToolSendMessageSendMessage `json:"send_message" api:"required"`
	Type        constant.SendMessage                `json:"type" default:"send_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SendMessage respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolSendMessage) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSendMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSendMessageSendMessage struct {
	// Optional message template with dynamic variable support using mustache syntax
	// (e.g., {{variable_name}}). When set, the assistant will use this template for
	// the SMS body instead of generating one. Dynamic variables like
	// {{telnyx_end_user_target}}, {{telnyx_agent_target}}, and custom webhook-provided
	// variables will be resolved at runtime.
	MessageTemplate string         `json:"message_template" api:"nullable"`
	ExtraFields     map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageTemplate respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolSendMessageSendMessage) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSendMessageSendMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSkipTurn struct {
	SkipTurn AssistantToolSkipTurnSkipTurn `json:"skip_turn" api:"required"`
	Type     constant.SkipTurn             `json:"type" default:"skip_turn"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SkipTurn    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolSkipTurn) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSkipTurn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSkipTurnSkipTurn struct {
	// The description of the function that will be passed to the assistant.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolSkipTurnSkipTurn) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSkipTurnSkipTurn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (BETA) The pay tool allows the assistant to collect card payments from the
// caller via DTMF during the conversation. Recording is automatically paused while
// the pay tool is active and resumes when the payment flow completes. The
// connector_name must reference a pay connector configured in the Telnyx API.
type AssistantToolPay struct {
	Pay  PayToolParamsResp `json:"pay" api:"required"`
	Type constant.Pay      `json:"type" default:"pay"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pay         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolPay) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolPay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func AssistantToolParamOfWebhook(webhook InferenceEmbeddingWebhookToolParamsWebhook) AssistantToolUnionParam {
	var variant InferenceEmbeddingWebhookToolParams
	variant.Webhook = webhook
	return AssistantToolUnionParam{OfWebhook: &variant}
}

func AssistantToolParamOfClientSideTool(clientSideTool AssistantToolClientSideToolClientSideToolParam) AssistantToolUnionParam {
	var variant AssistantToolClientSideToolParam
	variant.ClientSideTool = clientSideTool
	return AssistantToolUnionParam{OfClientSideTool: &variant}
}

func AssistantToolParamOfRetrieval(retrieval BucketIDsParam) AssistantToolUnionParam {
	var variant RetrievalToolParam
	variant.Retrieval = retrieval
	return AssistantToolUnionParam{OfRetrieval: &variant}
}

func AssistantToolParamOfHandoff(handoff AssistantToolHandoffHandoffParam) AssistantToolUnionParam {
	var variant AssistantToolHandoffParam
	variant.Handoff = handoff
	return AssistantToolUnionParam{OfHandoff: &variant}
}

func AssistantToolParamOfHangup(hangup HangupToolParams) AssistantToolUnionParam {
	var variant HangupToolParam
	variant.Hangup = hangup
	return AssistantToolUnionParam{OfHangup: &variant}
}

func AssistantToolParamOfTransfer(transfer AssistantToolTransferTransferParam) AssistantToolUnionParam {
	var variant AssistantToolTransferParam
	variant.Transfer = transfer
	return AssistantToolUnionParam{OfTransfer: &variant}
}

func AssistantToolParamOfInvite(invite AssistantToolInviteInviteParam) AssistantToolUnionParam {
	var variant AssistantToolInviteParam
	variant.Invite = invite
	return AssistantToolUnionParam{OfInvite: &variant}
}

func AssistantToolParamOfRefer(refer AssistantToolReferReferParam) AssistantToolUnionParam {
	var variant AssistantToolReferParam
	variant.Refer = refer
	return AssistantToolUnionParam{OfRefer: &variant}
}

func AssistantToolParamOfSendDtmf(sendDtmf map[string]any) AssistantToolUnionParam {
	var variant AssistantToolSendDtmfParam
	variant.SendDtmf = sendDtmf
	return AssistantToolUnionParam{OfSendDtmf: &variant}
}

func AssistantToolParamOfSendMessage(sendMessage AssistantToolSendMessageSendMessageParam) AssistantToolUnionParam {
	var variant AssistantToolSendMessageParam
	variant.SendMessage = sendMessage
	return AssistantToolUnionParam{OfSendMessage: &variant}
}

func AssistantToolParamOfSkipTurn(skipTurn AssistantToolSkipTurnSkipTurnParam) AssistantToolUnionParam {
	var variant AssistantToolSkipTurnParam
	variant.SkipTurn = skipTurn
	return AssistantToolUnionParam{OfSkipTurn: &variant}
}

func AssistantToolParamOfPay(pay PayToolParams) AssistantToolUnionParam {
	var variant AssistantToolPayParam
	variant.Pay = pay
	return AssistantToolUnionParam{OfPay: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantToolUnionParam struct {
	OfWebhook        *InferenceEmbeddingWebhookToolParams `json:",omitzero,inline"`
	OfClientSideTool *AssistantToolClientSideToolParam    `json:",omitzero,inline"`
	OfRetrieval      *RetrievalToolParam                  `json:",omitzero,inline"`
	OfHandoff        *AssistantToolHandoffParam           `json:",omitzero,inline"`
	OfHangup         *HangupToolParam                     `json:",omitzero,inline"`
	OfTransfer       *AssistantToolTransferParam          `json:",omitzero,inline"`
	OfInvite         *AssistantToolInviteParam            `json:",omitzero,inline"`
	OfRefer          *AssistantToolReferParam             `json:",omitzero,inline"`
	OfSendDtmf       *AssistantToolSendDtmfParam          `json:",omitzero,inline"`
	OfSendMessage    *AssistantToolSendMessageParam       `json:",omitzero,inline"`
	OfSkipTurn       *AssistantToolSkipTurnParam          `json:",omitzero,inline"`
	OfPay            *AssistantToolPayParam               `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhook,
		u.OfClientSideTool,
		u.OfRetrieval,
		u.OfHandoff,
		u.OfHangup,
		u.OfTransfer,
		u.OfInvite,
		u.OfRefer,
		u.OfSendDtmf,
		u.OfSendMessage,
		u.OfSkipTurn,
		u.OfPay)
}
func (u *AssistantToolUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfClientSideTool) {
		return u.OfClientSideTool
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	} else if !param.IsOmitted(u.OfHandoff) {
		return u.OfHandoff
	} else if !param.IsOmitted(u.OfHangup) {
		return u.OfHangup
	} else if !param.IsOmitted(u.OfTransfer) {
		return u.OfTransfer
	} else if !param.IsOmitted(u.OfInvite) {
		return u.OfInvite
	} else if !param.IsOmitted(u.OfRefer) {
		return u.OfRefer
	} else if !param.IsOmitted(u.OfSendDtmf) {
		return u.OfSendDtmf
	} else if !param.IsOmitted(u.OfSendMessage) {
		return u.OfSendMessage
	} else if !param.IsOmitted(u.OfSkipTurn) {
		return u.OfSkipTurn
	} else if !param.IsOmitted(u.OfPay) {
		return u.OfPay
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetWebhook() *InferenceEmbeddingWebhookToolParamsWebhook {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetClientSideTool() *AssistantToolClientSideToolClientSideToolParam {
	if vt := u.OfClientSideTool; vt != nil {
		return &vt.ClientSideTool
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetRetrieval() *BucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetHandoff() *AssistantToolHandoffHandoffParam {
	if vt := u.OfHandoff; vt != nil {
		return &vt.Handoff
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetHangup() *HangupToolParams {
	if vt := u.OfHangup; vt != nil {
		return &vt.Hangup
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetTransfer() *AssistantToolTransferTransferParam {
	if vt := u.OfTransfer; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetInvite() *AssistantToolInviteInviteParam {
	if vt := u.OfInvite; vt != nil {
		return &vt.Invite
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetRefer() *AssistantToolReferReferParam {
	if vt := u.OfRefer; vt != nil {
		return &vt.Refer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetSendDtmf() map[string]any {
	if vt := u.OfSendDtmf; vt != nil {
		return vt.SendDtmf
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetSendMessage() *AssistantToolSendMessageSendMessageParam {
	if vt := u.OfSendMessage; vt != nil {
		return &vt.SendMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetSkipTurn() *AssistantToolSkipTurnSkipTurnParam {
	if vt := u.OfSkipTurn; vt != nil {
		return &vt.SkipTurn
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetPay() *PayToolParams {
	if vt := u.OfPay; vt != nil {
		return &vt.Pay
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetType() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfClientSideTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHandoff; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHangup; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTransfer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfInvite; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRefer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSendDtmf; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSendMessage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSkipTurn; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPay; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AssistantToolUnionParam](
		"type",
		apijson.Discriminator[InferenceEmbeddingWebhookToolParams]("webhook"),
		apijson.Discriminator[AssistantToolClientSideToolParam]("client_side_tool"),
		apijson.Discriminator[RetrievalToolParam]("retrieval"),
		apijson.Discriminator[AssistantToolHandoffParam]("handoff"),
		apijson.Discriminator[HangupToolParam]("hangup"),
		apijson.Discriminator[AssistantToolTransferParam]("transfer"),
		apijson.Discriminator[AssistantToolInviteParam]("invite"),
		apijson.Discriminator[AssistantToolReferParam]("refer"),
		apijson.Discriminator[AssistantToolSendDtmfParam]("send_dtmf"),
		apijson.Discriminator[AssistantToolSendMessageParam]("send_message"),
		apijson.Discriminator[AssistantToolSkipTurnParam]("skip_turn"),
		apijson.Discriminator[AssistantToolPayParam]("pay"),
	)
}

// The properties ClientSideTool, Type are required.
type AssistantToolClientSideToolParam struct {
	ClientSideTool AssistantToolClientSideToolClientSideToolParam `json:"client_side_tool,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "client_side_tool".
	Type constant.ClientSideTool `json:"type" default:"client_side_tool"`
	paramObj
}

func (r AssistantToolClientSideToolParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolClientSideToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolClientSideToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name, Parameters are required.
type AssistantToolClientSideToolClientSideToolParam struct {
	// The description of the tool.
	Description string `json:"description" api:"required"`
	// The name of the tool.
	Name string `json:"name" api:"required"`
	// The parameters the tool accepts, described as a JSON Schema object. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	Parameters AssistantToolClientSideToolClientSideToolParametersParam `json:"parameters,omitzero" api:"required"`
	paramObj
}

func (r AssistantToolClientSideToolClientSideToolParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolClientSideToolClientSideToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolClientSideToolClientSideToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters the tool accepts, described as a JSON Schema object. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type AssistantToolClientSideToolClientSideToolParametersParam struct {
	// The properties of the parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r AssistantToolClientSideToolClientSideToolParametersParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolClientSideToolClientSideToolParametersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolClientSideToolClientSideToolParametersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolClientSideToolClientSideToolParametersParam](
		"type", "object",
	)
}

// The handoff tool allows the assistant to hand off control of the conversation to
// another AI assistant. By default, this will happen transparently to the end
// user.
//
// The properties Handoff, Type are required.
type AssistantToolHandoffParam struct {
	Handoff AssistantToolHandoffHandoffParam `json:"handoff,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "handoff".
	Type constant.Handoff `json:"type" default:"handoff"`
	paramObj
}

func (r AssistantToolHandoffParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolHandoffParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolHandoffParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AIAssistants is required.
type AssistantToolHandoffHandoffParam struct {
	// List of possible assistants that can receive a handoff.
	AIAssistants []AssistantToolHandoffHandoffAIAssistantParam `json:"ai_assistants,omitzero" api:"required"`
	// With the unified voice mode all assistants share the same voice, making the
	// handoff transparent to the user. With the distinct voice mode all assistants
	// retain their voice configuration, providing the experience of a conference call
	// with a team of assistants.
	//
	// Any of "unified", "distinct".
	VoiceMode string `json:"voice_mode,omitzero"`
	paramObj
}

func (r AssistantToolHandoffHandoffParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolHandoffHandoffParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolHandoffHandoffParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolHandoffHandoffParam](
		"voice_mode", "unified", "distinct",
	)
}

// The properties ID, Name are required.
type AssistantToolHandoffHandoffAIAssistantParam struct {
	// The ID of the assistant to hand off to.
	ID string `json:"id" api:"required"`
	// Helpful name for giving context on when to handoff to the assistant.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r AssistantToolHandoffHandoffAIAssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolHandoffHandoffAIAssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolHandoffHandoffAIAssistantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Transfer, Type are required.
type AssistantToolTransferParam struct {
	Transfer AssistantToolTransferTransferParam `json:"transfer,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "transfer".
	Type constant.Transfer `json:"type" default:"transfer"`
	paramObj
}

func (r AssistantToolTransferParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, Targets are required.
type AssistantToolTransferTransferParam struct {
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to. This can also be a dynamic
	// variable string like `{{ targets }}` where `targets` is returned by the dynamic
	// variables webhook and resolves to an array of target objects at runtime.
	Targets AssistantToolTransferTransferTargetsUnionParam `json:"targets,omitzero" api:"required"`
	// Optional delay in milliseconds before playing the warm message audio when the
	// transferred call is answered. When set, the audio_url is not included in the
	// dial command; instead, playback starts after the specified delay. When not set,
	// existing behavior (audio_url in dial) is preserved.
	WarmMessageDelayMs param.Opt[int64] `json:"warm_message_delay_ms,omitzero"`
	// A description of the transfer tool. By default, Telnyx generates this
	// automatically based on the configured targets. Typically only set when importing
	// an assistant from another provider that allowed a custom description; in that
	// case the provided value is preserved. Most users should leave this empty and let
	// Telnyx manage it.
	Description param.Opt[string] `json:"description,omitzero"`
	// Natural language instructions for your agent for how to provide context for the
	// transfer recipient.
	WarmTransferInstructions param.Opt[string] `json:"warm_transfer_instructions,omitzero"`
	// Custom headers to be added to the SIP INVITE for the transfer command.
	CustomHeaders []AssistantToolTransferTransferCustomHeaderParam `json:"custom_headers,omitzero"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
	// transferred call. Allows the assistant to detect when a voicemail system answers
	// the transferred call and take appropriate action.
	VoicemailDetection AssistantToolTransferTransferVoicemailDetectionParam `json:"voicemail_detection,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantToolTransferTransferTargetsUnionParam struct {
	OfTargetsList []AssistantToolTransferTransferTargetsTargetsListItemParam `json:",omitzero,inline"`
	OfString      param.Opt[string]                                          `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolTransferTransferTargetsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTargetsList, u.OfString)
}
func (u *AssistantToolTransferTransferTargetsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolTransferTransferTargetsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTargetsList) {
		return &u.OfTargetsList
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property To is required.
type AssistantToolTransferTransferTargetsTargetsListItemParam struct {
	// The destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferTargetsTargetsListItemParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferTargetsTargetsListItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferTargetsTargetsListItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolTransferTransferCustomHeaderParam struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferCustomHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferCustomHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferCustomHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
// transferred call. Allows the assistant to detect when a voicemail system answers
// the transferred call and take appropriate action.
type AssistantToolTransferTransferVoicemailDetectionParam struct {
	// Advanced AMD detection configuration parameters. All values are optional -
	// Telnyx will use defaults if not specified.
	DetectionConfig AssistantToolTransferTransferVoicemailDetectionDetectionConfigParam `json:"detection_config,omitzero"`
	// The AMD detection mode to use. 'premium' enables premium answering machine
	// detection. 'disabled' turns off AMD detection.
	//
	// Any of "disabled", "premium".
	DetectionMode string `json:"detection_mode,omitzero"`
	// Action to take when voicemail is detected on the transferred call.
	OnVoicemailDetected AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedParam `json:"on_voicemail_detected,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferVoicemailDetectionParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferVoicemailDetectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferVoicemailDetectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolTransferTransferVoicemailDetectionParam](
		"detection_mode", "disabled", "premium",
	)
}

// Advanced AMD detection configuration parameters. All values are optional -
// Telnyx will use defaults if not specified.
type AssistantToolTransferTransferVoicemailDetectionDetectionConfigParam struct {
	// Duration of silence after greeting detection before finalizing the result.
	AfterGreetingSilenceMillis param.Opt[int64] `json:"after_greeting_silence_millis,omitzero"`
	// Maximum silence duration between words during greeting.
	BetweenWordsSilenceMillis param.Opt[int64] `json:"between_words_silence_millis,omitzero"`
	// Expected duration of greeting speech.
	GreetingDurationMillis param.Opt[int64] `json:"greeting_duration_millis,omitzero"`
	// Duration of silence after the greeting to wait before considering the greeting
	// complete.
	GreetingSilenceDurationMillis param.Opt[int64] `json:"greeting_silence_duration_millis,omitzero"`
	// Maximum time to spend analyzing the greeting.
	GreetingTotalAnalysisTimeMillis param.Opt[int64] `json:"greeting_total_analysis_time_millis,omitzero"`
	// Maximum silence duration at the start of the call before speech.
	InitialSilenceMillis param.Opt[int64] `json:"initial_silence_millis,omitzero"`
	// Maximum number of words expected in a human greeting.
	MaximumNumberOfWords param.Opt[int64] `json:"maximum_number_of_words,omitzero"`
	// Maximum duration of a single word.
	MaximumWordLengthMillis param.Opt[int64] `json:"maximum_word_length_millis,omitzero"`
	// Minimum duration for audio to be considered a word.
	MinWordLengthMillis param.Opt[int64] `json:"min_word_length_millis,omitzero"`
	// Audio level threshold for silence detection.
	SilenceThreshold param.Opt[int64] `json:"silence_threshold,omitzero"`
	// Total time allowed for AMD analysis.
	TotalAnalysisTimeMillis param.Opt[int64] `json:"total_analysis_time_millis,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferVoicemailDetectionDetectionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferVoicemailDetectionDetectionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferVoicemailDetectionDetectionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take when voicemail is detected on the transferred call.
type AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedParam struct {
	// The action to take when voicemail is detected. 'stop_transfer' hangs up
	// immediately. 'leave_message_and_stop_transfer' leaves a message then hangs up.
	//
	// Any of "stop_transfer", "leave_message_and_stop_transfer".
	Action string `json:"action,omitzero"`
	// Configuration for the voicemail message to leave. Only applicable when action is
	// 'leave_message_and_stop_transfer'.
	VoicemailMessage AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam `json:"voicemail_message,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedParam](
		"action", "stop_transfer", "leave_message_and_stop_transfer",
	)
}

// Configuration for the voicemail message to leave. Only applicable when action is
// 'leave_message_and_stop_transfer'.
type AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam struct {
	// The specific message to leave as voicemail (converted to speech). Only
	// applicable when type is 'message'.
	Message param.Opt[string] `json:"message,omitzero"`
	// The type of voicemail message. Use 'message' to leave a specific TTS message, or
	// 'warm_transfer_instructions' to play the warm transfer audio.
	//
	// Any of "message", "warm_transfer_instructions".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolTransferTransferVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam](
		"type", "message", "warm_transfer_instructions",
	)
}

// The properties Invite, Type are required.
type AssistantToolInviteParam struct {
	Invite AssistantToolInviteInviteParam `json:"invite,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "invite".
	Type constant.Invite `json:"type" default:"invite"`
	paramObj
}

func (r AssistantToolInviteParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolInviteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolInviteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property From is required.
type AssistantToolInviteInviteParam struct {
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// The different possible targets of the invite. The assistant will be able to
	// choose one of the targets to invite to the call. This can also be a dynamic
	// variable string like `{{ targets }}` where `targets` is returned by the dynamic
	// variables webhook and resolves to an array of target objects at runtime. If
	// omitted or null, the invite tool can still be configured and targets may be
	// supplied dynamically at runtime.
	Targets AssistantToolInviteInviteTargetsUnionParam `json:"targets,omitzero"`
	// Custom headers to be added to the SIP INVITE for the invite command.
	CustomHeaders []AssistantToolInviteInviteCustomHeaderParam `json:"custom_headers,omitzero"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
	// invited call.
	VoicemailDetection AssistantToolInviteInviteVoicemailDetectionParam `json:"voicemail_detection,omitzero"`
	paramObj
}

func (r AssistantToolInviteInviteParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolInviteInviteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolInviteInviteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolInviteInviteCustomHeaderParam struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r AssistantToolInviteInviteCustomHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolInviteInviteCustomHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolInviteInviteCustomHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantToolInviteInviteTargetsUnionParam struct {
	OfTargetsList []AssistantToolInviteInviteTargetsTargetsListItemParam `json:",omitzero,inline"`
	OfString      param.Opt[string]                                      `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolInviteInviteTargetsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTargetsList, u.OfString)
}
func (u *AssistantToolInviteInviteTargetsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolInviteInviteTargetsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTargetsList) {
		return &u.OfTargetsList
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property To is required.
type AssistantToolInviteInviteTargetsTargetsListItemParam struct {
	// The destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AssistantToolInviteInviteTargetsTargetsListItemParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolInviteInviteTargetsTargetsListItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolInviteInviteTargetsTargetsListItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
// invited call.
type AssistantToolInviteInviteVoicemailDetectionParam struct {
	// The AMD detection mode to use. 'premium' enables premium answering machine
	// detection. 'disabled' turns off AMD detection.
	//
	// Any of "disabled", "premium".
	DetectionMode string `json:"detection_mode,omitzero"`
	// Action to take when voicemail is detected on the invited call.
	OnVoicemailDetected AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetectedParam `json:"on_voicemail_detected,omitzero"`
	paramObj
}

func (r AssistantToolInviteInviteVoicemailDetectionParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolInviteInviteVoicemailDetectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolInviteInviteVoicemailDetectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolInviteInviteVoicemailDetectionParam](
		"detection_mode", "disabled", "premium",
	)
}

// Action to take when voicemail is detected on the invited call.
type AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetectedParam struct {
	// The action to take when voicemail is detected.
	//
	// Any of "stop_invite".
	Action string `json:"action,omitzero"`
	paramObj
}

func (r AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetectedParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetectedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetectedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolInviteInviteVoicemailDetectionOnVoicemailDetectedParam](
		"action", "stop_invite",
	)
}

// The properties Refer, Type are required.
type AssistantToolReferParam struct {
	Refer AssistantToolReferReferParam `json:"refer,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "refer".
	Type constant.Refer `json:"type" default:"refer"`
	paramObj
}

func (r AssistantToolReferParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolReferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolReferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Targets is required.
type AssistantToolReferReferParam struct {
	// The different possible targets of the SIP refer. The assistant will be able to
	// choose one of the targets to refer the call to.
	Targets []AssistantToolReferReferTargetParam `json:"targets,omitzero" api:"required"`
	// Custom headers to be added to the SIP REFER.
	CustomHeaders []AssistantToolReferReferCustomHeaderParam `json:"custom_headers,omitzero"`
	// SIP headers to be added to the SIP REFER. Currently only User-to-User and
	// Diversion headers are supported.
	SipHeaders []AssistantToolReferReferSipHeaderParam `json:"sip_headers,omitzero"`
	paramObj
}

func (r AssistantToolReferReferParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolReferReferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolReferReferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, SipAddress are required.
type AssistantToolReferReferTargetParam struct {
	// The name of the target.
	Name string `json:"name" api:"required"`
	// The SIP URI to which the call will be referred.
	SipAddress string `json:"sip_address" api:"required"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword param.Opt[string] `json:"sip_auth_password,omitzero"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername param.Opt[string] `json:"sip_auth_username,omitzero"`
	paramObj
}

func (r AssistantToolReferReferTargetParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolReferReferTargetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolReferReferTargetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolReferReferCustomHeaderParam struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r AssistantToolReferReferCustomHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolReferReferCustomHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolReferReferCustomHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolReferReferSipHeaderParam struct {
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value param.Opt[string] `json:"value,omitzero"`
	// Any of "User-to-User", "Diversion".
	Name string `json:"name,omitzero"`
	paramObj
}

func (r AssistantToolReferReferSipHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolReferReferSipHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolReferReferSipHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolReferReferSipHeaderParam](
		"name", "User-to-User", "Diversion",
	)
}

// The properties SendDtmf, Type are required.
type AssistantToolSendDtmfParam struct {
	SendDtmf map[string]any `json:"send_dtmf,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "send_dtmf".
	Type constant.SendDtmf `json:"type" default:"send_dtmf"`
	paramObj
}

func (r AssistantToolSendDtmfParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolSendDtmfParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolSendDtmfParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The send_message tool allows the assistant to send SMS or MMS messages to the
// end user. The 'to' and 'from' addresses are automatically determined from the
// conversation context, and the message text is generated by the assistant unless
// a message_template is provided for runtime variable substitution.
//
// The properties SendMessage, Type are required.
type AssistantToolSendMessageParam struct {
	SendMessage AssistantToolSendMessageSendMessageParam `json:"send_message,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "send_message".
	Type constant.SendMessage `json:"type" default:"send_message"`
	paramObj
}

func (r AssistantToolSendMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolSendMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolSendMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSendMessageSendMessageParam struct {
	// Optional message template with dynamic variable support using mustache syntax
	// (e.g., {{variable_name}}). When set, the assistant will use this template for
	// the SMS body instead of generating one. Dynamic variables like
	// {{telnyx_end_user_target}}, {{telnyx_agent_target}}, and custom webhook-provided
	// variables will be resolved at runtime.
	MessageTemplate param.Opt[string] `json:"message_template,omitzero"`
	ExtraFields     map[string]any    `json:"-"`
	paramObj
}

func (r AssistantToolSendMessageSendMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolSendMessageSendMessageParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *AssistantToolSendMessageSendMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SkipTurn, Type are required.
type AssistantToolSkipTurnParam struct {
	SkipTurn AssistantToolSkipTurnSkipTurnParam `json:"skip_turn,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "skip_turn".
	Type constant.SkipTurn `json:"type" default:"skip_turn"`
	paramObj
}

func (r AssistantToolSkipTurnParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolSkipTurnParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolSkipTurnParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSkipTurnSkipTurnParam struct {
	// The description of the function that will be passed to the assistant.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AssistantToolSkipTurnSkipTurnParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolSkipTurnSkipTurnParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolSkipTurnSkipTurnParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (BETA) The pay tool allows the assistant to collect card payments from the
// caller via DTMF during the conversation. Recording is automatically paused while
// the pay tool is active and resumes when the payment flow completes. The
// connector_name must reference a pay connector configured in the Telnyx API.
//
// The properties Pay, Type are required.
type AssistantToolPayParam struct {
	Pay PayToolParams `json:"pay,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "pay".
	Type constant.Pay `json:"type" default:"pay"`
	paramObj
}

func (r AssistantToolPayParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolPayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolPayParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantsList struct {
	Data []InferenceEmbedding `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantsList) RawJSON() string { return r.JSON.raw }
func (r *AssistantsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioVisualizerConfig struct {
	// The color theme for the audio visualizer.
	//
	// Any of "verdant", "twilight", "bloom", "mystic", "flare", "glacier".
	Color AudioVisualizerConfigColor `json:"color"`
	// The preset style for the audio visualizer.
	Preset string `json:"preset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Preset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioVisualizerConfig) RawJSON() string { return r.JSON.raw }
func (r *AudioVisualizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AudioVisualizerConfig to a AudioVisualizerConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AudioVisualizerConfigParam.Overrides()
func (r AudioVisualizerConfig) ToParam() AudioVisualizerConfigParam {
	return param.Override[AudioVisualizerConfigParam](json.RawMessage(r.RawJSON()))
}

// The color theme for the audio visualizer.
type AudioVisualizerConfigColor string

const (
	AudioVisualizerConfigColorVerdant  AudioVisualizerConfigColor = "verdant"
	AudioVisualizerConfigColorTwilight AudioVisualizerConfigColor = "twilight"
	AudioVisualizerConfigColorBloom    AudioVisualizerConfigColor = "bloom"
	AudioVisualizerConfigColorMystic   AudioVisualizerConfigColor = "mystic"
	AudioVisualizerConfigColorFlare    AudioVisualizerConfigColor = "flare"
	AudioVisualizerConfigColorGlacier  AudioVisualizerConfigColor = "glacier"
)

type AudioVisualizerConfigParam struct {
	// The preset style for the audio visualizer.
	Preset param.Opt[string] `json:"preset,omitzero"`
	// The color theme for the audio visualizer.
	//
	// Any of "verdant", "twilight", "bloom", "mystic", "flare", "glacier".
	Color AudioVisualizerConfigColor `json:"color,omitzero"`
	paramObj
}

func (r AudioVisualizerConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AudioVisualizerConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudioVisualizerConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Authentication method used when connecting to the external LLM endpoint.
type AuthenticationMethod string

const (
	AuthenticationMethodToken       AuthenticationMethod = "token"
	AuthenticationMethodCertificate AuthenticationMethod = "certificate"
)

// Conversation flow as supplied by API clients (create / update).
//
// A directed graph of `FlowNodeReq` connected by `FlowEdge`s. Validation enforces
// unique node/edge IDs, that `start_node_id` references a real node, and that
// every edge's endpoints reference real nodes.
//
// The properties Nodes, StartNodeID are required.
type ConversationFlowReqParam struct {
	// All nodes in the flow. Must contain `start_node_id`. Each node is a prompt node
	// (`type: prompt`) or a tool node (`type: tool`).
	Nodes []ConversationFlowReqNodesUnionParam `json:"nodes,omitzero" api:"required"`
	// ID of the node where the conversation begins.
	StartNodeID string `json:"start_node_id" api:"required"`
	// Directed transitions between nodes. May be empty for a single-node flow.
	Edges []FlowEdgeParam `json:"edges,omitzero"`
	paramObj
}

func (r ConversationFlowReqParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationFlowReqParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationFlowReqParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationFlowReqNodesUnionParam struct {
	OfPrompt *ConversationFlowReqNodesPromptParam `json:",omitzero,inline"`
	OfTool   *ConversationFlowReqNodesToolParam   `json:",omitzero,inline"`
	OfSpeak  *ConversationFlowReqNodesSpeakParam  `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationFlowReqNodesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPrompt, u.OfTool, u.OfSpeak)
}
func (u *ConversationFlowReqNodesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationFlowReqNodesUnionParam) asAny() any {
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
func (u ConversationFlowReqNodesUnionParam) GetInstructions() *string {
	if vt := u.OfPrompt; vt != nil {
		return &vt.Instructions
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetExternalLlm() *ExternalLlmReqParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.ExternalLlm
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetInstructionsMode() *string {
	if vt := u.OfPrompt; vt != nil {
		return &vt.InstructionsMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetLlmAPIKeyRef() *string {
	if vt := u.OfPrompt; vt != nil && vt.LlmAPIKeyRef.Valid() {
		return &vt.LlmAPIKeyRef.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetModel() *string {
	if vt := u.OfPrompt; vt != nil && vt.Model.Valid() {
		return &vt.Model.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetSharedToolIDs() []string {
	if vt := u.OfPrompt; vt != nil {
		return vt.SharedToolIDs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetToolsMode() *string {
	if vt := u.OfPrompt; vt != nil {
		return &vt.ToolsMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetTranscription() *TranscriptionSettingsParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.Transcription
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetVoiceSettings() *VoiceSettingsParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.VoiceSettings
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetSharedToolID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.SharedToolID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetMessage() *string {
	if vt := u.OfSpeak; vt != nil {
		return &vt.Message
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationFlowReqNodesUnionParam) GetID() *string {
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
func (u ConversationFlowReqNodesUnionParam) GetName() *string {
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
func (u ConversationFlowReqNodesUnionParam) GetType() *string {
	if vt := u.OfPrompt; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSpeak; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Position property, if present.
func (u ConversationFlowReqNodesUnionParam) GetPosition() *NodePositionParam {
	if vt := u.OfPrompt; vt != nil {
		return &vt.Position
	} else if vt := u.OfTool; vt != nil {
		return &vt.Position
	} else if vt := u.OfSpeak; vt != nil {
		return &vt.Position
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationFlowReqNodesUnionParam](
		"type",
		apijson.Discriminator[ConversationFlowReqNodesPromptParam]("prompt"),
		apijson.Discriminator[ConversationFlowReqNodesToolParam]("tool"),
		apijson.Discriminator[ConversationFlowReqNodesSpeakParam]("speak"),
	)
}

// One step in a conversation flow, as supplied by API clients.
//
// Each node carries the prompt, tool scope, and optional overrides for
// model/voice/transcription. Unset overrides cascade from the assistant.
//
// The properties ID, Instructions are required.
type ConversationFlowReqNodesPromptParam struct {
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
	Position NodePositionParam `json:"position,omitzero"`
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

func (r ConversationFlowReqNodesPromptParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationFlowReqNodesPromptParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationFlowReqNodesPromptParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationFlowReqNodesPromptParam](
		"instructions_mode", "replace", "append",
	)
	apijson.RegisterFieldValidator[ConversationFlowReqNodesPromptParam](
		"tools_mode", "replace", "append",
	)
	apijson.RegisterFieldValidator[ConversationFlowReqNodesPromptParam](
		"type", "prompt",
	)
}

// A standalone tool step in a conversation flow, as supplied by clients.
//
// Unlike a prompt node, a tool node has no instructions or model — it isn't an LLM
// turn. Reaching it deterministically runs one shared tool (arguments filled from
// matching dynamic variables by name), then routes on the result via outgoing
// `tool_result` edges.
//
// The properties ID, SharedToolID are required.
type ConversationFlowReqNodesToolParam struct {
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
	Position NodePositionParam `json:"position,omitzero"`
	// Node kind discriminator. Always `tool` for a tool node.
	//
	// Any of "tool".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationFlowReqNodesToolParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationFlowReqNodesToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationFlowReqNodesToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationFlowReqNodesToolParam](
		"type", "tool",
	)
}

// A standalone scripted-message step in a flow, as supplied by clients.
//
// Unlike a prompt node, a speak node has no instructions or model — it isn't an
// LLM turn. Reaching it delivers `message` to the user verbatim (with
// `{{variable}}` interpolation), then routes via outgoing `llm` / `expression`
// edges.
//
// The properties ID, Message are required.
type ConversationFlowReqNodesSpeakParam struct {
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
	Position NodePositionParam `json:"position,omitzero"`
	// Node kind discriminator. Always `speak` for a speak node.
	//
	// Any of "speak".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ConversationFlowReqNodesSpeakParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationFlowReqNodesSpeakParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationFlowReqNodesSpeakParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationFlowReqNodesSpeakParam](
		"type", "speak",
	)
}

// If `telephony` is enabled, the assistant will be able to make and receive calls.
// If `messaging` is enabled, the assistant will be able to send and receive
// messages.
type EnabledFeatures string

const (
	EnabledFeaturesTelephony EnabledFeatures = "telephony"
	EnabledFeaturesMessaging EnabledFeatures = "messaging"
)

type ExternalLlm struct {
	// Base URL for the external LLM endpoint.
	BaseURL string `json:"base_url" api:"required"`
	// Model identifier to use with the external LLM endpoint.
	Model string `json:"model" api:"required"`
	// Authentication method used when connecting to the external LLM endpoint.
	//
	// Any of "token", "certificate".
	AuthenticationMethod AuthenticationMethod `json:"authentication_method"`
	// Integration secret identifier for the client certificate used with certificate
	// authentication.
	CertificateRef string `json:"certificate_ref"`
	// When `true`, Telnyx forwards the assistant's dynamic variables to the external
	// LLM endpoint as a top-level `extra_metadata` object on the chat completion
	// request body. Defaults to `false`. Example payload sent to the external
	// endpoint:
	// `{"extra_metadata": {"customer_name": "Jane", "account_id": "acct_789", "telnyx_agent_target": "+13125550100", "telnyx_end_user_target": "+13125550123"}}`.
	// Distinct from OpenAI's native `metadata` field, which has its own size and type
	// limits.
	ForwardMetadata bool `json:"forward_metadata"`
	// Integration secret identifier for the external LLM API key.
	LlmAPIKeyRef string `json:"llm_api_key_ref"`
	// URL used to retrieve an access token when certificate authentication is enabled.
	TokenRetrievalURL string `json:"token_retrieval_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL              respjson.Field
		Model                respjson.Field
		AuthenticationMethod respjson.Field
		CertificateRef       respjson.Field
		ForwardMetadata      respjson.Field
		LlmAPIKeyRef         respjson.Field
		TokenRetrievalURL    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalLlm) RawJSON() string { return r.JSON.raw }
func (r *ExternalLlm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Model are required.
type ExternalLlmReqParam struct {
	// Base URL for the external LLM endpoint.
	BaseURL string `json:"base_url" api:"required"`
	// Model identifier to use with the external LLM endpoint.
	Model string `json:"model" api:"required"`
	// Integration secret identifier for the client certificate used with certificate
	// authentication.
	CertificateRef param.Opt[string] `json:"certificate_ref,omitzero"`
	// When `true`, Telnyx forwards the assistant's dynamic variables to the external
	// LLM endpoint as a top-level `extra_metadata` object on the chat completion
	// request body. Defaults to `false`. Example payload sent to the external
	// endpoint:
	// `{"extra_metadata": {"customer_name": "Jane", "account_id": "acct_789", "telnyx_agent_target": "+13125550100", "telnyx_end_user_target": "+13125550123"}}`.
	// Distinct from OpenAI's native `metadata` field, which has its own size and type
	// limits.
	ForwardMetadata param.Opt[bool] `json:"forward_metadata,omitzero"`
	// Integration secret identifier for the external LLM API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// URL used to retrieve an access token when certificate authentication is enabled.
	TokenRetrievalURL param.Opt[string] `json:"token_retrieval_url,omitzero"`
	// Authentication method used when connecting to the external LLM endpoint.
	//
	// Any of "token", "certificate".
	AuthenticationMethod AuthenticationMethod `json:"authentication_method,omitzero"`
	paramObj
}

func (r ExternalLlmReqParam) MarshalJSON() (data []byte, err error) {
	type shadow ExternalLlmReqParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalLlmReqParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackConfig struct {
	ExternalLlm ExternalLlm `json:"external_llm"`
	// Integration secret identifier for the fallback model API key.
	LlmAPIKeyRef string `json:"llm_api_key_ref"`
	// Fallback Telnyx-hosted model to use when the primary LLM provider is
	// unavailable.
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExternalLlm  respjson.Field
		LlmAPIKeyRef respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FallbackConfig) RawJSON() string { return r.JSON.raw }
func (r *FallbackConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FallbackConfigReqParam struct {
	// Integration secret identifier for the fallback model API key.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Fallback Telnyx-hosted model to use when the primary LLM provider is
	// unavailable.
	Model       param.Opt[string]   `json:"model,omitzero"`
	ExternalLlm ExternalLlmReqParam `json:"external_llm,omitzero"`
	paramObj
}

func (r FallbackConfigReqParam) MarshalJSON() (data []byte, err error) {
	type shadow FallbackConfigReqParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FallbackConfigReqParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Directed transition from one node to a target, gated by a condition.
//
// The target is either another node in the same flow (`NodeTarget`) or a different
// assistant (`AssistantTarget`). Multiple edges may share a `start_node_id`; the
// runtime evaluates them in the order they're declared and takes the first whose
// condition is true.
type FlowEdge struct {
	// Caller-supplied unique identifier for this edge within the flow.
	ID string `json:"id" api:"required"`
	// Condition that gates the transition. Discriminated by `type`: `llm`,
	// `expression`.
	Condition FlowEdgeConditionUnion `json:"condition" api:"required"`
	// ID of the node this edge transitions away from.
	StartNodeID string `json:"start_node_id" api:"required"`
	// Destination of the transition. Discriminated by `type`: `node` (jump to another
	// node in this flow) or `assistant` (hand off to a different assistant).
	Target FlowEdgeTargetUnion `json:"target" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Condition   respjson.Field
		StartNodeID respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowEdge) RawJSON() string { return r.JSON.raw }
func (r *FlowEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FlowEdge to a FlowEdgeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FlowEdgeParam.Overrides()
func (r FlowEdge) ToParam() FlowEdgeParam {
	return param.Override[FlowEdgeParam](json.RawMessage(r.RawJSON()))
}

// FlowEdgeConditionUnion contains all possible properties and values from
// [FlowEdgeConditionLlm], [FlowEdgeConditionExpression],
// [FlowEdgeConditionDefault].
//
// Use the [FlowEdgeConditionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FlowEdgeConditionUnion struct {
	// This field is from variant [FlowEdgeConditionLlm].
	Prompt string `json:"prompt"`
	// Any of "llm", "expression", "default".
	Type string `json:"type"`
	// This field is from variant [FlowEdgeConditionExpression].
	Expression any `json:"expression"`
	JSON       struct {
		Prompt     respjson.Field
		Type       respjson.Field
		Expression respjson.Field
		raw        string
	} `json:"-"`
}

// anyFlowEdgeCondition is implemented by each variant of [FlowEdgeConditionUnion]
// to add type safety for the return type of [FlowEdgeConditionUnion.AsAny]
type anyFlowEdgeCondition interface {
	implFlowEdgeConditionUnion()
}

func (FlowEdgeConditionLlm) implFlowEdgeConditionUnion()        {}
func (FlowEdgeConditionExpression) implFlowEdgeConditionUnion() {}
func (FlowEdgeConditionDefault) implFlowEdgeConditionUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FlowEdgeConditionUnion.AsAny().(type) {
//	case telnyx.FlowEdgeConditionLlm:
//	case telnyx.FlowEdgeConditionExpression:
//	case telnyx.FlowEdgeConditionDefault:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FlowEdgeConditionUnion) AsAny() anyFlowEdgeCondition {
	switch u.Type {
	case "llm":
		return u.AsLlm()
	case "expression":
		return u.AsExpression()
	case "default":
		return u.AsDefault()
	}
	return nil
}

func (u FlowEdgeConditionUnion) AsLlm() (v FlowEdgeConditionLlm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FlowEdgeConditionUnion) AsExpression() (v FlowEdgeConditionExpression) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FlowEdgeConditionUnion) AsDefault() (v FlowEdgeConditionDefault) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FlowEdgeConditionUnion) RawJSON() string { return u.JSON.raw }

func (r *FlowEdgeConditionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge condition evaluated by the LLM from a natural-language prompt.
//
// The model is asked to judge the prompt against conversation context and returns
// true/false. Use this for fuzzy intents that aren't expressible as a
// deterministic expression (e.g. 'user wants to escalate to a human').
type FlowEdgeConditionLlm struct {
	// Natural-language criterion the LLM judges as true/false.
	Prompt string       `json:"prompt" api:"required"`
	Type   constant.Llm `json:"type" default:"llm"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompt      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowEdgeConditionLlm) RawJSON() string { return r.JSON.raw }
func (r *FlowEdgeConditionLlm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge condition evaluated as a deterministic expression AST.
//
// The expression is computed against runtime dynamic variables and must evaluate
// to a boolean. Prefer this over `LLMCondition` when the rule is a clean function
// of known variables — it's cheaper and predictable.
type FlowEdgeConditionExpression struct {
	// Root of the expression AST; evaluates to a boolean. Typed as free-form JSON to
	// avoid an uncompilable by-value self-reference; see the Expression schema for the
	// variant structure.
	Expression any                 `json:"expression" api:"required"`
	Type       constant.Expression `json:"type" default:"expression"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Expression  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowEdgeConditionExpression) RawJSON() string { return r.JSON.raw }
func (r *FlowEdgeConditionExpression) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
type FlowEdgeConditionDefault struct {
	Type constant.Default `json:"type" default:"default"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowEdgeConditionDefault) RawJSON() string { return r.JSON.raw }
func (r *FlowEdgeConditionDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FlowEdgeTargetUnion contains all possible properties and values from
// [FlowEdgeTargetNode], [FlowEdgeTargetAssistant].
//
// Use the [FlowEdgeTargetUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FlowEdgeTargetUnion struct {
	// This field is from variant [FlowEdgeTargetNode].
	NodeID string `json:"node_id"`
	// Any of "node", "assistant".
	Type string `json:"type"`
	// This field is from variant [FlowEdgeTargetAssistant].
	AssistantID string `json:"assistant_id"`
	// This field is from variant [FlowEdgeTargetAssistant].
	Position NodePosition `json:"position"`
	// This field is from variant [FlowEdgeTargetAssistant].
	VoiceMode string `json:"voice_mode"`
	JSON      struct {
		NodeID      respjson.Field
		Type        respjson.Field
		AssistantID respjson.Field
		Position    respjson.Field
		VoiceMode   respjson.Field
		raw         string
	} `json:"-"`
}

// anyFlowEdgeTarget is implemented by each variant of [FlowEdgeTargetUnion] to add
// type safety for the return type of [FlowEdgeTargetUnion.AsAny]
type anyFlowEdgeTarget interface {
	implFlowEdgeTargetUnion()
}

func (FlowEdgeTargetNode) implFlowEdgeTargetUnion()      {}
func (FlowEdgeTargetAssistant) implFlowEdgeTargetUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FlowEdgeTargetUnion.AsAny().(type) {
//	case telnyx.FlowEdgeTargetNode:
//	case telnyx.FlowEdgeTargetAssistant:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FlowEdgeTargetUnion) AsAny() anyFlowEdgeTarget {
	switch u.Type {
	case "node":
		return u.AsNode()
	case "assistant":
		return u.AsAssistant()
	}
	return nil
}

func (u FlowEdgeTargetUnion) AsNode() (v FlowEdgeTargetNode) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FlowEdgeTargetUnion) AsAssistant() (v FlowEdgeTargetAssistant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FlowEdgeTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *FlowEdgeTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge target referencing another node within the same flow.
//
// The runtime transitions the active node to `node_id` and continues processing
// within the current assistant's flow.
type FlowEdgeTargetNode struct {
	// ID of the node this edge transitions into.
	NodeID string        `json:"node_id" api:"required"`
	Type   constant.Node `json:"type" default:"node"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NodeID      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowEdgeTargetNode) RawJSON() string { return r.JSON.raw }
func (r *FlowEdgeTargetNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge target referencing a different assistant.
//
// When the edge fires, the conversation hands off to `assistant_id`: the active
// assistant on the conversation row is rewritten and the new assistant's flow
// starts at its own `start_node_id`. The current turn's LLM response is delivered
// to the user as-is; subsequent turns route to the new assistant.
type FlowEdgeTargetAssistant struct {
	// ID of the assistant the conversation transitions to.
	AssistantID string             `json:"assistant_id" api:"required"`
	Type        constant.Assistant `json:"type" default:"assistant"`
	// Optional canvas coordinates for rendering the target assistant as a node in
	// authoring UIs. Pure presentation — the runtime ignores it; round-trips so
	// frontends can persist graph layout across reloads. When multiple edges target
	// the same assistant, each edge's `position` is independent (frontends typically
	// use the first non-null one).
	Position NodePosition `json:"position"`
	// Voice behavior when handing off to the target assistant, mirroring the handoff
	// tool's `voice_mode`. `unified` (default) keeps the current voice across the
	// handoff; `distinct` lets the target assistant speak with its own configured
	// voice. Only applies to assistant targets — node targets override voice via the
	// node's own `voice_settings`.
	//
	// Any of "unified", "distinct".
	VoiceMode string `json:"voice_mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantID respjson.Field
		Type        respjson.Field
		Position    respjson.Field
		VoiceMode   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FlowEdgeTargetAssistant) RawJSON() string { return r.JSON.raw }
func (r *FlowEdgeTargetAssistant) UnmarshalJSON(data []byte) error {
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
type FlowEdgeParam struct {
	// Caller-supplied unique identifier for this edge within the flow.
	ID string `json:"id" api:"required"`
	// Condition that gates the transition. Discriminated by `type`: `llm`,
	// `expression`.
	Condition FlowEdgeConditionUnionParam `json:"condition,omitzero" api:"required"`
	// ID of the node this edge transitions away from.
	StartNodeID string `json:"start_node_id" api:"required"`
	// Destination of the transition. Discriminated by `type`: `node` (jump to another
	// node in this flow) or `assistant` (hand off to a different assistant).
	Target FlowEdgeTargetUnionParam `json:"target,omitzero" api:"required"`
	paramObj
}

func (r FlowEdgeParam) MarshalJSON() (data []byte, err error) {
	type shadow FlowEdgeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlowEdgeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FlowEdgeConditionUnionParam struct {
	OfLlm        *FlowEdgeConditionLlmParam        `json:",omitzero,inline"`
	OfExpression *FlowEdgeConditionExpressionParam `json:",omitzero,inline"`
	OfDefault    *FlowEdgeConditionDefaultParam    `json:",omitzero,inline"`
	paramUnion
}

func (u FlowEdgeConditionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlm, u.OfExpression, u.OfDefault)
}
func (u *FlowEdgeConditionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FlowEdgeConditionUnionParam) asAny() any {
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
func (u FlowEdgeConditionUnionParam) GetPrompt() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Prompt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeConditionUnionParam) GetExpression() *any {
	if vt := u.OfExpression; vt != nil {
		return &vt.Expression
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeConditionUnionParam) GetType() *string {
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
	apijson.RegisterUnion[FlowEdgeConditionUnionParam](
		"type",
		apijson.Discriminator[FlowEdgeConditionLlmParam]("llm"),
		apijson.Discriminator[FlowEdgeConditionExpressionParam]("expression"),
		apijson.Discriminator[FlowEdgeConditionDefaultParam]("default"),
	)
}

// Edge condition evaluated by the LLM from a natural-language prompt.
//
// The model is asked to judge the prompt against conversation context and returns
// true/false. Use this for fuzzy intents that aren't expressible as a
// deterministic expression (e.g. 'user wants to escalate to a human').
//
// The properties Prompt, Type are required.
type FlowEdgeConditionLlmParam struct {
	// Natural-language criterion the LLM judges as true/false.
	Prompt string `json:"prompt" api:"required"`
	// This field can be elided, and will marshal its zero value as "llm".
	Type constant.Llm `json:"type" default:"llm"`
	paramObj
}

func (r FlowEdgeConditionLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow FlowEdgeConditionLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlowEdgeConditionLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Edge condition evaluated as a deterministic expression AST.
//
// The expression is computed against runtime dynamic variables and must evaluate
// to a boolean. Prefer this over `LLMCondition` when the rule is a clean function
// of known variables — it's cheaper and predictable.
//
// The properties Expression, Type are required.
type FlowEdgeConditionExpressionParam struct {
	// Root of the expression AST; evaluates to a boolean. Typed as free-form JSON to
	// avoid an uncompilable by-value self-reference; see the Expression schema for the
	// variant structure.
	Expression any `json:"expression,omitzero" api:"required"`
	// This field can be elided, and will marshal its zero value as "expression".
	Type constant.Expression `json:"type" default:"expression"`
	paramObj
}

func (r FlowEdgeConditionExpressionParam) MarshalJSON() (data []byte, err error) {
	type shadow FlowEdgeConditionExpressionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlowEdgeConditionExpressionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewFlowEdgeConditionDefaultParam() FlowEdgeConditionDefaultParam {
	return FlowEdgeConditionDefaultParam{
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
// [NewFlowEdgeConditionDefaultParam].
type FlowEdgeConditionDefaultParam struct {
	Type constant.Default `json:"type" default:"default"`
	paramObj
}

func (r FlowEdgeConditionDefaultParam) MarshalJSON() (data []byte, err error) {
	type shadow FlowEdgeConditionDefaultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlowEdgeConditionDefaultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FlowEdgeTargetUnionParam struct {
	OfNode      *FlowEdgeTargetNodeParam      `json:",omitzero,inline"`
	OfAssistant *FlowEdgeTargetAssistantParam `json:",omitzero,inline"`
	paramUnion
}

func (u FlowEdgeTargetUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfNode, u.OfAssistant)
}
func (u *FlowEdgeTargetUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FlowEdgeTargetUnionParam) asAny() any {
	if !param.IsOmitted(u.OfNode) {
		return u.OfNode
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeTargetUnionParam) GetNodeID() *string {
	if vt := u.OfNode; vt != nil {
		return &vt.NodeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeTargetUnionParam) GetAssistantID() *string {
	if vt := u.OfAssistant; vt != nil {
		return &vt.AssistantID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeTargetUnionParam) GetPosition() *NodePositionParam {
	if vt := u.OfAssistant; vt != nil {
		return &vt.Position
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeTargetUnionParam) GetVoiceMode() *string {
	if vt := u.OfAssistant; vt != nil {
		return &vt.VoiceMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FlowEdgeTargetUnionParam) GetType() *string {
	if vt := u.OfNode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[FlowEdgeTargetUnionParam](
		"type",
		apijson.Discriminator[FlowEdgeTargetNodeParam]("node"),
		apijson.Discriminator[FlowEdgeTargetAssistantParam]("assistant"),
	)
}

// Edge target referencing another node within the same flow.
//
// The runtime transitions the active node to `node_id` and continues processing
// within the current assistant's flow.
//
// The properties NodeID, Type are required.
type FlowEdgeTargetNodeParam struct {
	// ID of the node this edge transitions into.
	NodeID string `json:"node_id" api:"required"`
	// This field can be elided, and will marshal its zero value as "node".
	Type constant.Node `json:"type" default:"node"`
	paramObj
}

func (r FlowEdgeTargetNodeParam) MarshalJSON() (data []byte, err error) {
	type shadow FlowEdgeTargetNodeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlowEdgeTargetNodeParam) UnmarshalJSON(data []byte) error {
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
type FlowEdgeTargetAssistantParam struct {
	// ID of the assistant the conversation transitions to.
	AssistantID string `json:"assistant_id" api:"required"`
	// Optional canvas coordinates for rendering the target assistant as a node in
	// authoring UIs. Pure presentation — the runtime ignores it; round-trips so
	// frontends can persist graph layout across reloads. When multiple edges target
	// the same assistant, each edge's `position` is independent (frontends typically
	// use the first non-null one).
	Position NodePositionParam `json:"position,omitzero"`
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

func (r FlowEdgeTargetAssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow FlowEdgeTargetAssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlowEdgeTargetAssistantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FlowEdgeTargetAssistantParam](
		"voice_mode", "unified", "distinct",
	)
}

type HangupTool struct {
	Hangup HangupToolParamsResp `json:"hangup" api:"required"`
	// Any of "hangup".
	Type HangupToolType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hangup      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HangupTool) RawJSON() string { return r.JSON.raw }
func (r *HangupTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HangupTool to a HangupToolParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HangupToolParam.Overrides()
func (r HangupTool) ToParam() HangupToolParam {
	return param.Override[HangupToolParam](json.RawMessage(r.RawJSON()))
}

type HangupToolType string

const (
	HangupToolTypeHangup HangupToolType = "hangup"
)

// The properties Hangup, Type are required.
type HangupToolParam struct {
	Hangup HangupToolParams `json:"hangup,omitzero" api:"required"`
	// Any of "hangup".
	Type HangupToolType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r HangupToolParam) MarshalJSON() (data []byte, err error) {
	type shadow HangupToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HangupToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HangupToolParamsResp struct {
	// The description of the function that will be passed to the assistant.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HangupToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *HangupToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this HangupToolParamsResp to a HangupToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// HangupToolParams.Overrides()
func (r HangupToolParamsResp) ToParam() HangupToolParams {
	return param.Override[HangupToolParams](json.RawMessage(r.RawJSON()))
}

type HangupToolParams struct {
	// The description of the function that will be passed to the assistant.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r HangupToolParams) MarshalJSON() (data []byte, err error) {
	type shadow HangupToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HangupToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImportMetadata struct {
	// ID of the assistant in the provider's system.
	ImportID string `json:"import_id"`
	// Provider the assistant was imported from.
	//
	// Any of "elevenlabs", "vapi", "retell".
	ImportProvider ImportMetadataImportProvider `json:"import_provider"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImportID       respjson.Field
		ImportProvider respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImportMetadata) RawJSON() string { return r.JSON.raw }
func (r *ImportMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provider the assistant was imported from.
type ImportMetadataImportProvider string

const (
	ImportMetadataImportProviderElevenlabs ImportMetadataImportProvider = "elevenlabs"
	ImportMetadataImportProviderVapi       ImportMetadataImportProvider = "vapi"
	ImportMetadataImportProviderRetell     ImportMetadataImportProvider = "retell"
)

type InferenceEmbedding struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// System instructions for the assistant. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Instructions string `json:"instructions" api:"required"`
	// ID of the model to use when `external_llm` is not set. You can use the
	// [Get models API](https://developers.telnyx.com/api-reference/openai-chat/get-available-models-openai-compatible)
	// to see available models. If `external_llm` is provided, the assistant uses
	// `external_llm` instead of this field. If neither `model` nor `external_llm` is
	// provided, Telnyx applies the default model.
	Model string `json:"model" api:"required"`
	Name  string `json:"name" api:"required"`
	// Conversation flow as returned by the API.
	ConversationFlow InferenceEmbeddingConversationFlow `json:"conversation_flow"`
	Description      string                             `json:"description"`
	// Map of dynamic variables and their values
	DynamicVariables map[string]any `json:"dynamic_variables"`
	// Timeout in milliseconds for the dynamic variables webhook. Must be between 1 and
	// 10000 ms. If the webhook does not respond within this timeout, the call proceeds
	// with default values. See the
	// [dynamic variables guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables).
	DynamicVariablesWebhookTimeoutMs int64 `json:"dynamic_variables_webhook_timeout_ms"`
	// If `dynamic_variables_webhook_url` is set, Telnyx sends a POST request to this
	// URL at the start of the conversation to resolve dynamic variables. **Gotcha:**
	// the webhook response must wrap variables under a top-level `dynamic_variables`
	// object, e.g. `{"dynamic_variables": {"customer_name": "Jane"}}`. Returning a
	// flat object will be ignored and variables will fall back to their defaults. See
	// the
	// [dynamic variables guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// for the full request/response format and timeout behavior.
	DynamicVariablesWebhookURL string            `json:"dynamic_variables_webhook_url"`
	EnabledFeatures            []EnabledFeatures `json:"enabled_features"`
	ExternalLlm                ExternalLlm       `json:"external_llm"`
	FallbackConfig             FallbackConfig    `json:"fallback_config"`
	// Text that the assistant will use to start the conversation. This may be
	// templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables).
	// Use an empty string to have the assistant wait for the user to speak first. Use
	// the special value `<assistant-speaks-first-with-model-generated-message>` to
	// have the assistant generate the greeting based on the system instructions.
	Greeting        string          `json:"greeting"`
	ImportMetadata  ImportMetadata  `json:"import_metadata"`
	InsightSettings InsightSettings `json:"insight_settings"`
	// Connected integrations attached to the assistant. The catalog of available
	// integrations is at `/ai/integrations`; the user's connected integrations are at
	// `/ai/integrations/connections`. Each item references a catalog integration by
	// `integration_id`.
	Integrations []AssistantIntegration `json:"integrations"`
	// Settings for interruptions and how the assistant decides the user has finished
	// speaking. These timings are most relevant when using non turn-taking
	// transcription models. For turn-taking models like `deepgram/flux`, end-of-turn
	// behavior is controlled by the transcription end-of-turn settings under
	// `transcription.settings` (`eot_threshold`, `eot_timeout_ms`,
	// `eager_eot_threshold`).
	InterruptionSettings InferenceEmbeddingInterruptionSettings `json:"interruption_settings"`
	// This is only needed when using third-party inference providers selected by
	// `model`. The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your LLM provider's API key. For bring-your-own endpoint
	// authentication, use `external_llm.llm_api_key_ref` instead. Warning: Free plans
	// are unlikely to work with this integration.
	LlmAPIKeyRef string `json:"llm_api_key_ref"`
	// MCP servers attached to the assistant. Create MCP servers with
	// `/ai/mcp_servers`, then reference them by `id` here.
	McpServers            []AssistantMcpServer `json:"mcp_servers"`
	MessagingSettings     MessagingSettings    `json:"messaging_settings"`
	ObservabilitySettings Observability        `json:"observability_settings"`
	// Configuration for post-conversation processing. When enabled, the assistant
	// receives one additional LLM turn after the conversation ends, allowing it to
	// execute tool calls such as logging to a CRM or sending a summary. The assistant
	// can execute multiple parallel or sequential tools during this phase.
	// Telephony-control tools (e.g. hangup, transfer) are unavailable
	// post-conversation. Beta feature.
	PostConversationSettings PostConversationSettings `json:"post_conversation_settings"`
	PrivacySettings          PrivacySettings          `json:"privacy_settings"`
	// IDs of missions related to this assistant.
	RelatedMissionIDs []string `json:"related_mission_ids"`
	// Tags associated with the assistant. Tags can also be managed with the assistant
	// tag endpoints.
	Tags              []string          `json:"tags"`
	TelephonySettings TelephonySettings `json:"telephony_settings"`
	// Deprecated for new integrations. Inline tool definitions available to the
	// assistant. Prefer `tool_ids` to attach shared tools created with the AI Tools
	// endpoints.
	Tools         []AssistantToolUnion  `json:"tools"`
	Transcription TranscriptionSettings `json:"transcription"`
	// Timestamp when this assistant version was created.
	VersionCreatedAt time.Time `json:"version_created_at" format:"date-time"`
	// Identifier for the assistant version returned by version-aware assistant
	// endpoints.
	VersionID string `json:"version_id"`
	// Human-readable name for the assistant version.
	VersionName   string        `json:"version_name"`
	VoiceSettings VoiceSettings `json:"voice_settings"`
	// Configuration settings for the assistant's web widget.
	WidgetSettings WidgetSettings `json:"widget_settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		CreatedAt                        respjson.Field
		Instructions                     respjson.Field
		Model                            respjson.Field
		Name                             respjson.Field
		ConversationFlow                 respjson.Field
		Description                      respjson.Field
		DynamicVariables                 respjson.Field
		DynamicVariablesWebhookTimeoutMs respjson.Field
		DynamicVariablesWebhookURL       respjson.Field
		EnabledFeatures                  respjson.Field
		ExternalLlm                      respjson.Field
		FallbackConfig                   respjson.Field
		Greeting                         respjson.Field
		ImportMetadata                   respjson.Field
		InsightSettings                  respjson.Field
		Integrations                     respjson.Field
		InterruptionSettings             respjson.Field
		LlmAPIKeyRef                     respjson.Field
		McpServers                       respjson.Field
		MessagingSettings                respjson.Field
		ObservabilitySettings            respjson.Field
		PostConversationSettings         respjson.Field
		PrivacySettings                  respjson.Field
		RelatedMissionIDs                respjson.Field
		Tags                             respjson.Field
		TelephonySettings                respjson.Field
		Tools                            respjson.Field
		Transcription                    respjson.Field
		VersionCreatedAt                 respjson.Field
		VersionID                        respjson.Field
		VersionName                      respjson.Field
		VoiceSettings                    respjson.Field
		WidgetSettings                   respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbedding) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Conversation flow as returned by the API.
type InferenceEmbeddingConversationFlow struct {
	// All nodes in the flow.
	Nodes []InferenceEmbeddingConversationFlowNodesUnion `json:"nodes" api:"required"`
	// ID of the node where the conversation begins.
	StartNodeID string `json:"start_node_id" api:"required"`
	// Directed transitions between nodes.
	Edges []FlowEdge `json:"edges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Nodes       respjson.Field
		StartNodeID respjson.Field
		Edges       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingConversationFlow) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingConversationFlow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InferenceEmbeddingConversationFlowNodesUnion contains all possible properties
// and values from [InferenceEmbeddingConversationFlowNodesPrompt],
// [InferenceEmbeddingConversationFlowNodesTool],
// [InferenceEmbeddingConversationFlowNodesSpeak].
//
// Use the [InferenceEmbeddingConversationFlowNodesUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InferenceEmbeddingConversationFlowNodesUnion struct {
	ID string `json:"id"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	Instructions string `json:"instructions"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	ExternalLlm ExternalLlm `json:"external_llm"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	InstructionsMode string `json:"instructions_mode"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	LlmAPIKeyRef string `json:"llm_api_key_ref"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	Model string `json:"model"`
	Name  string `json:"name"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	Position NodePosition `json:"position"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	SharedToolIDs []string `json:"shared_tool_ids"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	Tools [][]AssistantToolUnion `json:"tools"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	ToolsMode string `json:"tools_mode"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	Transcription TranscriptionSettings `json:"transcription"`
	// Any of "prompt", "tool", "speak".
	Type string `json:"type"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesPrompt].
	VoiceSettings VoiceSettings `json:"voice_settings"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesTool].
	SharedToolID string `json:"shared_tool_id"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesTool].
	Tool []AssistantToolUnion `json:"tool"`
	// This field is from variant [InferenceEmbeddingConversationFlowNodesSpeak].
	Message string `json:"message"`
	JSON    struct {
		ID               respjson.Field
		Instructions     respjson.Field
		ExternalLlm      respjson.Field
		InstructionsMode respjson.Field
		LlmAPIKeyRef     respjson.Field
		Model            respjson.Field
		Name             respjson.Field
		Position         respjson.Field
		SharedToolIDs    respjson.Field
		Tools            respjson.Field
		ToolsMode        respjson.Field
		Transcription    respjson.Field
		Type             respjson.Field
		VoiceSettings    respjson.Field
		SharedToolID     respjson.Field
		Tool             respjson.Field
		Message          respjson.Field
		raw              string
	} `json:"-"`
}

// anyInferenceEmbeddingConversationFlowNode is implemented by each variant of
// [InferenceEmbeddingConversationFlowNodesUnion] to add type safety for the return
// type of [InferenceEmbeddingConversationFlowNodesUnion.AsAny]
type anyInferenceEmbeddingConversationFlowNode interface {
	implInferenceEmbeddingConversationFlowNodesUnion()
}

func (InferenceEmbeddingConversationFlowNodesPrompt) implInferenceEmbeddingConversationFlowNodesUnion() {
}
func (InferenceEmbeddingConversationFlowNodesTool) implInferenceEmbeddingConversationFlowNodesUnion() {
}
func (InferenceEmbeddingConversationFlowNodesSpeak) implInferenceEmbeddingConversationFlowNodesUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := InferenceEmbeddingConversationFlowNodesUnion.AsAny().(type) {
//	case telnyx.InferenceEmbeddingConversationFlowNodesPrompt:
//	case telnyx.InferenceEmbeddingConversationFlowNodesTool:
//	case telnyx.InferenceEmbeddingConversationFlowNodesSpeak:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InferenceEmbeddingConversationFlowNodesUnion) AsAny() anyInferenceEmbeddingConversationFlowNode {
	switch u.Type {
	case "prompt":
		return u.AsPrompt()
	case "tool":
		return u.AsTool()
	case "speak":
		return u.AsSpeak()
	}
	return nil
}

func (u InferenceEmbeddingConversationFlowNodesUnion) AsPrompt() (v InferenceEmbeddingConversationFlowNodesPrompt) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InferenceEmbeddingConversationFlowNodesUnion) AsTool() (v InferenceEmbeddingConversationFlowNodesTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InferenceEmbeddingConversationFlowNodesUnion) AsSpeak() (v InferenceEmbeddingConversationFlowNodesSpeak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InferenceEmbeddingConversationFlowNodesUnion) RawJSON() string { return u.JSON.raw }

func (r *InferenceEmbeddingConversationFlowNodesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One step in a conversation flow, as returned by the API.
type InferenceEmbeddingConversationFlowNodesPrompt struct {
	// Caller-supplied unique identifier for this node within the flow.
	ID string `json:"id" api:"required"`
	// Prompt that drives the LLM while this node is active. Required.
	Instructions string `json:"instructions" api:"required"`
	// Override for `Assistant.external_llm` while this node is active. Use this to
	// route a node's turns to a different external LLM (different `model`, `base_url`,
	// credentials). Part of the LLM bundle — see `model` for cascade semantics.
	// Mutually exclusive with `model` on the node (a single LLM identity per node).
	ExternalLlm ExternalLlm `json:"external_llm"`
	// How `instructions` combine with the assistant-level instructions. `replace`
	// (default): the node's instructions are used alone. `append`: the node's
	// instructions are concatenated after the assistant's instructions.
	//
	// Any of "replace", "append".
	InstructionsMode string `json:"instructions_mode"`
	// Override for `Assistant.llm_api_key_ref` while this node is active. Part of the
	// LLM bundle — see `model` for cascade semantics.
	LlmAPIKeyRef string `json:"llm_api_key_ref"`
	// Override for `Assistant.model` while this node is active. Part of the LLM bundle
	// (`model` + `llm_api_key_ref` + `external_llm`): when any of the three is set on
	// the node, all three are taken from the node and the assistant-level LLM identity
	// is not consulted. When none of the three is set, the assistant's bundle cascades
	// unchanged.
	Model string `json:"model"`
	// Optional human-readable label, displayed in authoring UIs.
	Name string `json:"name"`
	// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
	// by the runtime; round-trips so frontends can persist graph layout across
	// reloads.
	Position NodePosition `json:"position"`
	// IDs of shared (org-level) tools available at this node. Knowledge bases are
	// attached the same way — via a shared retrieval tool. Tools not listed here are
	// not callable while this node is active.
	SharedToolIDs []string `json:"shared_tool_ids"`
	// Full tool definitions for this node, resolved from `shared_tool_ids`
	// server-side. Populated on responses so clients can render the flow without a
	// follow-up fetch per shared tool. Ignored on input — set `shared_tool_ids` to
	// configure a node's tools.
	Tools [][]AssistantToolUnion `json:"tools"`
	// How `shared_tool_ids` combine with the assistant-level tool set. `replace`
	// (default): only the node's tools are callable. `append`: the node's tools are
	// added to the assistant's tools. Ignored when `shared_tool_ids` is null.
	//
	// Any of "replace", "append".
	ToolsMode string `json:"tools_mode"`
	// Per-node transcription override (response form).
	Transcription TranscriptionSettings `json:"transcription"`
	// Node kind discriminator. `prompt` is an LLM-driven step.
	//
	// Any of "prompt".
	Type string `json:"type"`
	// Per-node voice override (response form).
	VoiceSettings VoiceSettings `json:"voice_settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Instructions     respjson.Field
		ExternalLlm      respjson.Field
		InstructionsMode respjson.Field
		LlmAPIKeyRef     respjson.Field
		Model            respjson.Field
		Name             respjson.Field
		Position         respjson.Field
		SharedToolIDs    respjson.Field
		Tools            respjson.Field
		ToolsMode        respjson.Field
		Transcription    respjson.Field
		Type             respjson.Field
		VoiceSettings    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingConversationFlowNodesPrompt) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingConversationFlowNodesPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A standalone tool step in a conversation flow, as returned by the API.
type InferenceEmbeddingConversationFlowNodesTool struct {
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
	Name string `json:"name"`
	// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
	// by the runtime; round-trips so frontends can persist graph layout across
	// reloads.
	Position NodePosition `json:"position"`
	// Full tool definition resolved from `shared_tool_id` server-side. Populated on
	// responses so clients can render the node without a follow-up fetch. Ignored on
	// input — set `shared_tool_id`.
	Tool []AssistantToolUnion `json:"tool"`
	// Node kind discriminator. Always `tool` for a tool node.
	//
	// Any of "tool".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		SharedToolID respjson.Field
		Name         respjson.Field
		Position     respjson.Field
		Tool         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingConversationFlowNodesTool) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingConversationFlowNodesTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A standalone scripted-message step in a flow, as returned by the API.
type InferenceEmbeddingConversationFlowNodesSpeak struct {
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
	Name string `json:"name"`
	// Optional canvas coordinates used by authoring UIs to lay out the graph. Ignored
	// by the runtime; round-trips so frontends can persist graph layout across
	// reloads.
	Position NodePosition `json:"position"`
	// Node kind discriminator. Always `speak` for a speak node.
	//
	// Any of "speak".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Message     respjson.Field
		Name        respjson.Field
		Position    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingConversationFlowNodesSpeak) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingConversationFlowNodesSpeak) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Settings for interruptions and how the assistant decides the user has finished
// speaking. These timings are most relevant when using non turn-taking
// transcription models. For turn-taking models like `deepgram/flux`, end-of-turn
// behavior is controlled by the transcription end-of-turn settings under
// `transcription.settings` (`eot_threshold`, `eot_timeout_ms`,
// `eager_eot_threshold`).
type InferenceEmbeddingInterruptionSettings struct {
	// When true, disables user interruptions while the assistant greeting is playing.
	DisableGreetingInterruption bool `json:"disable_greeting_interruption"`
	// Whether users can interrupt the assistant while it is speaking.
	Enable bool `json:"enable"`
	// Controls when the assistant starts speaking after the user stops. These
	// thresholds primarily apply to non turn-taking transcription models. For
	// turn-taking models like `deepgram/flux`, end-of-turn detection is driven by the
	// transcription end-of-turn settings under `transcription.settings` instead.
	StartSpeakingPlan StartSpeakingPlan `json:"start_speaking_plan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisableGreetingInterruption respjson.Field
		Enable                      respjson.Field
		StartSpeakingPlan           respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingInterruptionSettings) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingInterruptionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InferenceEmbeddingInterruptionSettings to a
// InferenceEmbeddingInterruptionSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InferenceEmbeddingInterruptionSettingsParam.Overrides()
func (r InferenceEmbeddingInterruptionSettings) ToParam() InferenceEmbeddingInterruptionSettingsParam {
	return param.Override[InferenceEmbeddingInterruptionSettingsParam](json.RawMessage(r.RawJSON()))
}

// Settings for interruptions and how the assistant decides the user has finished
// speaking. These timings are most relevant when using non turn-taking
// transcription models. For turn-taking models like `deepgram/flux`, end-of-turn
// behavior is controlled by the transcription end-of-turn settings under
// `transcription.settings` (`eot_threshold`, `eot_timeout_ms`,
// `eager_eot_threshold`).
type InferenceEmbeddingInterruptionSettingsParam struct {
	// When true, disables user interruptions while the assistant greeting is playing.
	DisableGreetingInterruption param.Opt[bool] `json:"disable_greeting_interruption,omitzero"`
	// Whether users can interrupt the assistant while it is speaking.
	Enable param.Opt[bool] `json:"enable,omitzero"`
	// Controls when the assistant starts speaking after the user stops. These
	// thresholds primarily apply to non turn-taking transcription models. For
	// turn-taking models like `deepgram/flux`, end-of-turn detection is driven by the
	// transcription end-of-turn settings under `transcription.settings` instead.
	StartSpeakingPlan StartSpeakingPlanParam `json:"start_speaking_plan,omitzero"`
	paramObj
}

func (r InferenceEmbeddingInterruptionSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingInterruptionSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingInterruptionSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingWebhookToolParamsResp struct {
	// Any of "webhook".
	Type    InferenceEmbeddingWebhookToolParamsType        `json:"type" api:"required"`
	Webhook InferenceEmbeddingWebhookToolParamsWebhookResp `json:"webhook" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InferenceEmbeddingWebhookToolParamsResp to a
// InferenceEmbeddingWebhookToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InferenceEmbeddingWebhookToolParams.Overrides()
func (r InferenceEmbeddingWebhookToolParamsResp) ToParam() InferenceEmbeddingWebhookToolParams {
	return param.Override[InferenceEmbeddingWebhookToolParams](json.RawMessage(r.RawJSON()))
}

type InferenceEmbeddingWebhookToolParamsType string

const (
	InferenceEmbeddingWebhookToolParamsTypeWebhook InferenceEmbeddingWebhookToolParamsType = "webhook"
)

type InferenceEmbeddingWebhookToolParamsWebhookResp struct {
	// The description of the tool.
	Description string `json:"description" api:"required"`
	// The name of the tool.
	Name string `json:"name" api:"required"`
	// The URL of the external tool to be called. This URL is going to be used by the
	// assistant. The URL can be templated like: `https://example.com/api/v1/{id}`,
	// where `{id}` is a placeholder for a value that will be provided by the assistant
	// if `path_parameters` are provided with the `id` attribute.
	URL string `json:"url" api:"required"`
	// If async, the assistant will move forward without waiting for your server to
	// respond.
	Async bool `json:"async"`
	// Maximum time in milliseconds that the conversation worker waits for an async
	// webhook response before returning "Submitted" to the LLM. If unset, the platform
	// default (currently 300ms) is used.
	AsyncTimeoutMs int64 `json:"async_timeout_ms"`
	// The body parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the body of the request. See
	// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
	// for documentation about the format
	BodyParameters InferenceEmbeddingWebhookToolParamsWebhookBodyParametersResp `json:"body_parameters"`
	// The headers to be sent to the external tool.
	Headers []InferenceEmbeddingWebhookToolParamsWebhookHeaderResp `json:"headers"`
	// The HTTP method to be used when calling the external tool.
	//
	// Any of "GET", "POST", "PUT", "DELETE", "PATCH".
	Method string `json:"method"`
	// The path parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the path of the request if the
	// URL contains a placeholder for a value. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	PathParameters InferenceEmbeddingWebhookToolParamsWebhookPathParametersResp `json:"path_parameters"`
	// The query parameters the webhook tool accepts, described as a JSON Schema
	// object. These parameters will be passed to the webhook as the query of the
	// request. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	QueryParameters InferenceEmbeddingWebhookToolParamsWebhookQueryParametersResp `json:"query_parameters"`
	// A list of mappings that extract values from the webhook response and store them
	// as dynamic variables. Each mapping specifies a dynamic variable name and a
	// dot-notation path to the value in the response body.
	StoreFieldsAsVariables []InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariableResp `json:"store_fields_as_variables"`
	// The maximum number of milliseconds to wait for the webhook to respond. Only
	// applicable when async is false.
	TimeoutMs int64 `json:"timeout_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description            respjson.Field
		Name                   respjson.Field
		URL                    respjson.Field
		Async                  respjson.Field
		AsyncTimeoutMs         respjson.Field
		BodyParameters         respjson.Field
		Headers                respjson.Field
		Method                 respjson.Field
		PathParameters         respjson.Field
		QueryParameters        respjson.Field
		StoreFieldsAsVariables respjson.Field
		TimeoutMs              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsWebhookResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsWebhookResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The body parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the body of the request. See
// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
// for documentation about the format
type InferenceEmbeddingWebhookToolParamsWebhookBodyParametersResp struct {
	// The properties of the body parameters.
	Properties map[string]any `json:"properties"`
	// The required properties of the body parameters.
	Required []string `json:"required"`
	// Any of "object".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsWebhookBodyParametersResp) RawJSON() string {
	return r.JSON.raw
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookBodyParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingWebhookToolParamsWebhookHeaderResp struct {
	Name string `json:"name"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `Bearer {{#integration_secret}}test-secret{{/integration_secret}}` to pass the
	// value of the integration secret as the bearer token.
	// [Telnyx signature headers](https://developers.telnyx.com/docs/voice/programmable-voice/voice-api-webhooks)
	// will be automatically added to the request.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsWebhookHeaderResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsWebhookHeaderResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The path parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the path of the request if the
// URL contains a placeholder for a value. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsWebhookPathParametersResp struct {
	// The properties of the path parameters.
	Properties map[string]any `json:"properties"`
	// The required properties of the path parameters.
	Required []string `json:"required"`
	// Any of "object".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsWebhookPathParametersResp) RawJSON() string {
	return r.JSON.raw
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookPathParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The query parameters the webhook tool accepts, described as a JSON Schema
// object. These parameters will be passed to the webhook as the query of the
// request. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsWebhookQueryParametersResp struct {
	// The properties of the query parameters.
	Properties map[string]any `json:"properties"`
	// The required properties of the query parameters.
	Required []string `json:"required"`
	// Any of "object".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsWebhookQueryParametersResp) RawJSON() string {
	return r.JSON.raw
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookQueryParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariableResp struct {
	// The name of the dynamic variable to store the extracted value in.
	Name string `json:"name" api:"required"`
	// A dot-notation path to the value in the webhook response body (e.g.
	// 'customer.name' or 'id').
	ValuePath string `json:"value_path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ValuePath   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariableResp) RawJSON() string {
	return r.JSON.raw
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariableResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Webhook are required.
type InferenceEmbeddingWebhookToolParams struct {
	// Any of "webhook".
	Type    InferenceEmbeddingWebhookToolParamsType    `json:"type,omitzero" api:"required"`
	Webhook InferenceEmbeddingWebhookToolParamsWebhook `json:"webhook,omitzero" api:"required"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name, URL are required.
type InferenceEmbeddingWebhookToolParamsWebhook struct {
	// The description of the tool.
	Description string `json:"description" api:"required"`
	// The name of the tool.
	Name string `json:"name" api:"required"`
	// The URL of the external tool to be called. This URL is going to be used by the
	// assistant. The URL can be templated like: `https://example.com/api/v1/{id}`,
	// where `{id}` is a placeholder for a value that will be provided by the assistant
	// if `path_parameters` are provided with the `id` attribute.
	URL string `json:"url" api:"required"`
	// If async, the assistant will move forward without waiting for your server to
	// respond.
	Async param.Opt[bool] `json:"async,omitzero"`
	// Maximum time in milliseconds that the conversation worker waits for an async
	// webhook response before returning "Submitted" to the LLM. If unset, the platform
	// default (currently 300ms) is used.
	AsyncTimeoutMs param.Opt[int64] `json:"async_timeout_ms,omitzero"`
	// The maximum number of milliseconds to wait for the webhook to respond. Only
	// applicable when async is false.
	TimeoutMs param.Opt[int64] `json:"timeout_ms,omitzero"`
	// The body parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the body of the request. See
	// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
	// for documentation about the format
	BodyParameters InferenceEmbeddingWebhookToolParamsWebhookBodyParameters `json:"body_parameters,omitzero"`
	// The headers to be sent to the external tool.
	Headers []InferenceEmbeddingWebhookToolParamsWebhookHeader `json:"headers,omitzero"`
	// The HTTP method to be used when calling the external tool.
	//
	// Any of "GET", "POST", "PUT", "DELETE", "PATCH".
	Method string `json:"method,omitzero"`
	// The path parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the path of the request if the
	// URL contains a placeholder for a value. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	PathParameters InferenceEmbeddingWebhookToolParamsWebhookPathParameters `json:"path_parameters,omitzero"`
	// The query parameters the webhook tool accepts, described as a JSON Schema
	// object. These parameters will be passed to the webhook as the query of the
	// request. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	QueryParameters InferenceEmbeddingWebhookToolParamsWebhookQueryParameters `json:"query_parameters,omitzero"`
	// A list of mappings that extract values from the webhook response and store them
	// as dynamic variables. Each mapping specifies a dynamic variable name and a
	// dot-notation path to the value in the response body.
	StoreFieldsAsVariables []InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariable `json:"store_fields_as_variables,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsWebhook) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsWebhook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsWebhook](
		"method", "GET", "POST", "PUT", "DELETE", "PATCH",
	)
}

// The body parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the body of the request. See
// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
// for documentation about the format
type InferenceEmbeddingWebhookToolParamsWebhookBodyParameters struct {
	// The properties of the body parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the body parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsWebhookBodyParameters) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsWebhookBodyParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookBodyParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsWebhookBodyParameters](
		"type", "object",
	)
}

type InferenceEmbeddingWebhookToolParamsWebhookHeader struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `Bearer {{#integration_secret}}test-secret{{/integration_secret}}` to pass the
	// value of the integration secret as the bearer token.
	// [Telnyx signature headers](https://developers.telnyx.com/docs/voice/programmable-voice/voice-api-webhooks)
	// will be automatically added to the request.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsWebhookHeader) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsWebhookHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The path parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the path of the request if the
// URL contains a placeholder for a value. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsWebhookPathParameters struct {
	// The properties of the path parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the path parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsWebhookPathParameters) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsWebhookPathParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookPathParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsWebhookPathParameters](
		"type", "object",
	)
}

// The query parameters the webhook tool accepts, described as a JSON Schema
// object. These parameters will be passed to the webhook as the query of the
// request. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsWebhookQueryParameters struct {
	// The properties of the query parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the query parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsWebhookQueryParameters) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsWebhookQueryParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookQueryParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsWebhookQueryParameters](
		"type", "object",
	)
}

// The properties Name, ValuePath are required.
type InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariable struct {
	// The name of the dynamic variable to store the extracted value in.
	Name string `json:"name" api:"required"`
	// A dot-notation path to the value in the webhook response body (e.g.
	// 'customer.name' or 'id').
	ValuePath string `json:"value_path" api:"required"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariable) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsWebhookStoreFieldsAsVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightSettings struct {
	// Reference to an Insight Group. Insights in this group will be run automatically
	// for all the assistant's conversations.
	InsightGroupID string `json:"insight_group_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsightGroupID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightSettings) RawJSON() string { return r.JSON.raw }
func (r *InsightSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InsightSettings to a InsightSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InsightSettingsParam.Overrides()
func (r InsightSettings) ToParam() InsightSettingsParam {
	return param.Override[InsightSettingsParam](json.RawMessage(r.RawJSON()))
}

type InsightSettingsParam struct {
	// Reference to an Insight Group. Insights in this group will be run automatically
	// for all the assistant's conversations.
	InsightGroupID param.Opt[string] `json:"insight_group_id,omitzero"`
	paramObj
}

func (r InsightSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow InsightSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingSettings struct {
	// If more than this many minutes have passed since the last message, the assistant
	// will start a new conversation instead of continuing the existing one.
	ConversationInactivityMinutes int64 `json:"conversation_inactivity_minutes"`
	// Default Messaging Profile used for messaging exchanges with your assistant. This
	// will be created automatically on assistant creation.
	DefaultMessagingProfileID string `json:"default_messaging_profile_id"`
	// The URL where webhooks related to delivery statused for assistant messages will
	// be sent.
	DeliveryStatusWebhookURL string `json:"delivery_status_webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationInactivityMinutes respjson.Field
		DefaultMessagingProfileID     respjson.Field
		DeliveryStatusWebhookURL      respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingSettings) RawJSON() string { return r.JSON.raw }
func (r *MessagingSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MessagingSettings to a MessagingSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MessagingSettingsParam.Overrides()
func (r MessagingSettings) ToParam() MessagingSettingsParam {
	return param.Override[MessagingSettingsParam](json.RawMessage(r.RawJSON()))
}

type MessagingSettingsParam struct {
	// If more than this many minutes have passed since the last message, the assistant
	// will start a new conversation instead of continuing the existing one.
	ConversationInactivityMinutes param.Opt[int64] `json:"conversation_inactivity_minutes,omitzero"`
	// Default Messaging Profile used for messaging exchanges with your assistant. This
	// will be created automatically on assistant creation.
	DefaultMessagingProfileID param.Opt[string] `json:"default_messaging_profile_id,omitzero"`
	// The URL where webhooks related to delivery statused for assistant messages will
	// be sent.
	DeliveryStatusWebhookURL param.Opt[string] `json:"delivery_status_webhook_url,omitzero"`
	paramObj
}

func (r MessagingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow MessagingSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// 2D coordinates for a node, used by authoring UIs to lay out the graph.
//
// Purely a presentation aid. The runtime ignores `position`; it round-trips
// through the API so frontends can persist the graph layout customers arrange in
// the editor.
type NodePosition struct {
	// Horizontal coordinate in the authoring canvas.
	X float64 `json:"x" api:"required"`
	// Vertical coordinate in the authoring canvas.
	Y float64 `json:"y" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NodePosition) RawJSON() string { return r.JSON.raw }
func (r *NodePosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NodePosition to a NodePositionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NodePositionParam.Overrides()
func (r NodePosition) ToParam() NodePositionParam {
	return param.Override[NodePositionParam](json.RawMessage(r.RawJSON()))
}

// 2D coordinates for a node, used by authoring UIs to lay out the graph.
//
// Purely a presentation aid. The runtime ignores `position`; it round-trips
// through the API so frontends can persist the graph layout customers arrange in
// the editor.
//
// The properties X, Y are required.
type NodePositionParam struct {
	// Horizontal coordinate in the authoring canvas.
	X float64 `json:"x" api:"required"`
	// Vertical coordinate in the authoring canvas.
	Y float64 `json:"y" api:"required"`
	paramObj
}

func (r NodePositionParam) MarshalJSON() (data []byte, err error) {
	type shadow NodePositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NodePositionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Observability struct {
	Host        string `json:"host"`
	PromptLabel string `json:"prompt_label"`
	PromptName  string `json:"prompt_name"`
	// Whether to auto-publish the assistant's instructions as a Langfuse prompt.
	//
	// When ENABLED + prompt_name set, every assistant create/update pushes
	// `instructions` to Langfuse via create_prompt and stores the returned version in
	// prompt_version.
	//
	// Any of "enabled", "disabled".
	PromptSync    PromptSyncStatus `json:"prompt_sync"`
	PromptVersion int64            `json:"prompt_version"`
	PublicKeyRef  string           `json:"public_key_ref"`
	SecretKeyRef  string           `json:"secret_key_ref"`
	// Any of "enabled", "disabled".
	Status ObservabilityStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Host          respjson.Field
		PromptLabel   respjson.Field
		PromptName    respjson.Field
		PromptSync    respjson.Field
		PromptVersion respjson.Field
		PublicKeyRef  respjson.Field
		SecretKeyRef  respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Observability) RawJSON() string { return r.JSON.raw }
func (r *Observability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObservabilityReqParam struct {
	Host          param.Opt[string] `json:"host,omitzero"`
	PromptLabel   param.Opt[string] `json:"prompt_label,omitzero"`
	PromptName    param.Opt[string] `json:"prompt_name,omitzero"`
	PromptVersion param.Opt[int64]  `json:"prompt_version,omitzero"`
	PublicKeyRef  param.Opt[string] `json:"public_key_ref,omitzero"`
	SecretKeyRef  param.Opt[string] `json:"secret_key_ref,omitzero"`
	// Whether to auto-publish the assistant's instructions as a Langfuse prompt.
	//
	// When ENABLED + prompt_name set, every assistant create/update pushes
	// `instructions` to Langfuse via create_prompt and stores the returned version in
	// prompt_version.
	//
	// Any of "enabled", "disabled".
	PromptSync PromptSyncStatus `json:"prompt_sync,omitzero"`
	// Any of "enabled", "disabled".
	Status ObservabilityStatus `json:"status,omitzero"`
	paramObj
}

func (r ObservabilityReqParam) MarshalJSON() (data []byte, err error) {
	type shadow ObservabilityReqParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObservabilityReqParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObservabilityStatus string

const (
	ObservabilityStatusEnabled  ObservabilityStatus = "enabled"
	ObservabilityStatusDisabled ObservabilityStatus = "disabled"
)

// Configuration for post-conversation processing. When enabled, the assistant
// receives one additional LLM turn after the conversation ends, allowing it to
// execute tool calls such as logging to a CRM or sending a summary. The assistant
// can execute multiple parallel or sequential tools during this phase.
// Telephony-control tools (e.g. hangup, transfer) are unavailable
// post-conversation. Beta feature.
type PostConversationSettings struct {
	// Whether post-conversation processing is enabled. When true, the assistant will
	// be invoked after the conversation ends to perform any final tool calls. Defaults
	// to false.
	Enabled bool `json:"enabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostConversationSettings) RawJSON() string { return r.JSON.raw }
func (r *PostConversationSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for post-conversation processing. When enabled, the assistant
// receives one additional LLM turn after the conversation ends, allowing it to
// execute tool calls such as logging to a CRM or sending a summary. The assistant
// can execute multiple parallel or sequential tools during this phase.
// Telephony-control tools (e.g. hangup, transfer) are unavailable
// post-conversation. Beta feature.
type PostConversationSettingsReqParam struct {
	// Whether post-conversation processing is enabled. When true, the assistant will
	// be invoked after the conversation ends to perform any final tool calls. Defaults
	// to false.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r PostConversationSettingsReqParam) MarshalJSON() (data []byte, err error) {
	type shadow PostConversationSettingsReqParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostConversationSettingsReqParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PrivacySettings struct {
	// If true, conversation history and insights will be stored. If false, they will
	// not be stored. This in‑tool toggle governs solely the retention of conversation
	// history and insights via the AI assistant. It has no effect on any separate
	// recording, transcription, or storage configuration that you have set at the
	// account, number, or application level. All such external settings remain in
	// force regardless of your selection here.
	DataRetention bool `json:"data_retention"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataRetention respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PrivacySettings) RawJSON() string { return r.JSON.raw }
func (r *PrivacySettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PrivacySettings to a PrivacySettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PrivacySettingsParam.Overrides()
func (r PrivacySettings) ToParam() PrivacySettingsParam {
	return param.Override[PrivacySettingsParam](json.RawMessage(r.RawJSON()))
}

type PrivacySettingsParam struct {
	// If true, conversation history and insights will be stored. If false, they will
	// not be stored. This in‑tool toggle governs solely the retention of conversation
	// history and insights via the AI assistant. It has no effect on any separate
	// recording, transcription, or storage configuration that you have set at the
	// account, number, or application level. All such external settings remain in
	// force regardless of your selection here.
	DataRetention param.Opt[bool] `json:"data_retention,omitzero"`
	paramObj
}

func (r PrivacySettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PrivacySettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PrivacySettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether to auto-publish the assistant's instructions as a Langfuse prompt.
//
// When ENABLED + prompt_name set, every assistant create/update pushes
// `instructions` to Langfuse via create_prompt and stores the returned version in
// prompt_version.
type PromptSyncStatus string

const (
	PromptSyncStatusEnabled  PromptSyncStatus = "enabled"
	PromptSyncStatusDisabled PromptSyncStatus = "disabled"
)

type RetrievalTool struct {
	Retrieval BucketIDs `json:"retrieval" api:"required"`
	// Any of "retrieval".
	Type RetrievalToolType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Retrieval   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RetrievalTool) RawJSON() string { return r.JSON.raw }
func (r *RetrievalTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RetrievalTool to a RetrievalToolParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RetrievalToolParam.Overrides()
func (r RetrievalTool) ToParam() RetrievalToolParam {
	return param.Override[RetrievalToolParam](json.RawMessage(r.RawJSON()))
}

type RetrievalToolType string

const (
	RetrievalToolTypeRetrieval RetrievalToolType = "retrieval"
)

// The properties Retrieval, Type are required.
type RetrievalToolParam struct {
	Retrieval BucketIDsParam `json:"retrieval,omitzero" api:"required"`
	// Any of "retrieval".
	Type RetrievalToolType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r RetrievalToolParam) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls when the assistant starts speaking after the user stops. These
// thresholds primarily apply to non turn-taking transcription models. For
// turn-taking models like `deepgram/flux`, end-of-turn detection is driven by the
// transcription end-of-turn settings under `transcription.settings` instead.
type StartSpeakingPlan struct {
	// Endpointing thresholds used to decide when the user has finished speaking.
	// Applies to non turn-taking transcription models. For `deepgram/flux`, use
	// `transcription.settings.eot_threshold` / `eot_timeout_ms` /
	// `eager_eot_threshold`.
	TranscriptionEndpointingPlan TranscriptionEndpointingPlan `json:"transcription_endpointing_plan"`
	// Minimum seconds to wait before the assistant starts speaking.
	WaitSeconds float64 `json:"wait_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TranscriptionEndpointingPlan respjson.Field
		WaitSeconds                  respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StartSpeakingPlan) RawJSON() string { return r.JSON.raw }
func (r *StartSpeakingPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StartSpeakingPlan to a StartSpeakingPlanParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StartSpeakingPlanParam.Overrides()
func (r StartSpeakingPlan) ToParam() StartSpeakingPlanParam {
	return param.Override[StartSpeakingPlanParam](json.RawMessage(r.RawJSON()))
}

// Controls when the assistant starts speaking after the user stops. These
// thresholds primarily apply to non turn-taking transcription models. For
// turn-taking models like `deepgram/flux`, end-of-turn detection is driven by the
// transcription end-of-turn settings under `transcription.settings` instead.
type StartSpeakingPlanParam struct {
	// Minimum seconds to wait before the assistant starts speaking.
	WaitSeconds param.Opt[float64] `json:"wait_seconds,omitzero"`
	// Endpointing thresholds used to decide when the user has finished speaking.
	// Applies to non turn-taking transcription models. For `deepgram/flux`, use
	// `transcription.settings.eot_threshold` / `eot_timeout_ms` /
	// `eager_eot_threshold`.
	TranscriptionEndpointingPlan TranscriptionEndpointingPlanParam `json:"transcription_endpointing_plan,omitzero"`
	paramObj
}

func (r StartSpeakingPlanParam) MarshalJSON() (data []byte, err error) {
	type shadow StartSpeakingPlanParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StartSpeakingPlanParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonySettings struct {
	// Default Texml App used for voice calls with your assistant. This will be created
	// automatically on assistant creation.
	DefaultTexmlAppID string `json:"default_texml_app_id"`
	// The noise suppression engine to use. Use 'disabled' to turn off noise
	// suppression.
	//
	// Any of "krisp", "deepfilternet", "disabled".
	NoiseSuppression TelephonySettingsNoiseSuppression `json:"noise_suppression"`
	// Configuration for noise suppression. Only applicable when noise_suppression is
	// 'deepfilternet'.
	NoiseSuppressionConfig TelephonySettingsNoiseSuppressionConfig `json:"noise_suppression_config"`
	// Configuration for call recording format and channel settings.
	RecordingSettings TelephonySettingsRecordingSettings `json:"recording_settings"`
	// When enabled, allows users to interact with your AI assistant directly from your
	// website without requiring authentication. This is required for FE widgets that
	// work with assistants that have telephony enabled.
	SupportsUnauthenticatedWebCalls bool `json:"supports_unauthenticated_web_calls"`
	// Maximum duration in seconds for the AI assistant to participate on the call.
	// When this limit is reached the assistant will be stopped. This limit does not
	// apply to portions of a call without an active assistant (for instance, a call
	// transferred to a human representative).
	TimeLimitSecs int64 `json:"time_limit_secs"`
	// Duration in seconds of end user silence before the assistant checks in on the
	// user. When this limit is reached the assistant will prompt the user to respond.
	// This is distinct from user_idle_timeout_secs which stops the assistant entirely.
	UserIdleReplySecs int64 `json:"user_idle_reply_secs"`
	// Maximum duration in seconds of end user silence on the call. When this limit is
	// reached the assistant will be stopped. This limit does not apply to portions of
	// a call without an active assistant (for instance, a call transferred to a human
	// representative).
	UserIdleTimeoutSecs int64 `json:"user_idle_timeout_secs"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on
	// outgoing calls. These settings only apply if AMD is enabled on the Dial command.
	// See
	// [TeXML Dial documentation](https://developers.telnyx.com/api-reference/texml-rest-commands/initiate-an-outbound-call)
	// for enabling AMD. Recommended settings: MachineDetection=Enable, AsyncAmd=true,
	// DetectionMode=Premium.
	VoicemailDetection TelephonySettingsVoicemailDetection `json:"voicemail_detection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultTexmlAppID               respjson.Field
		NoiseSuppression                respjson.Field
		NoiseSuppressionConfig          respjson.Field
		RecordingSettings               respjson.Field
		SupportsUnauthenticatedWebCalls respjson.Field
		TimeLimitSecs                   respjson.Field
		UserIdleReplySecs               respjson.Field
		UserIdleTimeoutSecs             respjson.Field
		VoicemailDetection              respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonySettings) RawJSON() string { return r.JSON.raw }
func (r *TelephonySettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TelephonySettings to a TelephonySettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TelephonySettingsParam.Overrides()
func (r TelephonySettings) ToParam() TelephonySettingsParam {
	return param.Override[TelephonySettingsParam](json.RawMessage(r.RawJSON()))
}

// The noise suppression engine to use. Use 'disabled' to turn off noise
// suppression.
type TelephonySettingsNoiseSuppression string

const (
	TelephonySettingsNoiseSuppressionKrisp         TelephonySettingsNoiseSuppression = "krisp"
	TelephonySettingsNoiseSuppressionDeepfilternet TelephonySettingsNoiseSuppression = "deepfilternet"
	TelephonySettingsNoiseSuppressionDisabled      TelephonySettingsNoiseSuppression = "disabled"
)

// Configuration for noise suppression. Only applicable when noise_suppression is
// 'deepfilternet'.
type TelephonySettingsNoiseSuppressionConfig struct {
	// Attenuation limit for noise suppression. Range: 0-100.
	AttenuationLimit int64 `json:"attenuation_limit"`
	// Mode for noise suppression configuration.
	//
	// Any of "advanced".
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttenuationLimit respjson.Field
		Mode             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonySettingsNoiseSuppressionConfig) RawJSON() string { return r.JSON.raw }
func (r *TelephonySettingsNoiseSuppressionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for call recording format and channel settings.
type TelephonySettingsRecordingSettings struct {
	// The number of channels for the recording. 'single' for mono, 'dual' for stereo.
	//
	// Any of "single", "dual".
	Channels string `json:"channels"`
	// Whether call recording is enabled. When set to false, calls will not be recorded
	// regardless of other recording configuration.
	Enabled bool `json:"enabled"`
	// The format of the recording file.
	//
	// Any of "wav", "mp3".
	Format string `json:"format"`
	// When enabled, the call recording will stop when the conversation ends (for
	// example, when the assistant hangs up or the call is transferred). When disabled,
	// recording continues until the call itself ends.
	StopOnConversationEnd bool `json:"stop_on_conversation_end"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels              respjson.Field
		Enabled               respjson.Field
		Format                respjson.Field
		StopOnConversationEnd respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonySettingsRecordingSettings) RawJSON() string { return r.JSON.raw }
func (r *TelephonySettingsRecordingSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for voicemail detection (AMD - Answering Machine Detection) on
// outgoing calls. These settings only apply if AMD is enabled on the Dial command.
// See
// [TeXML Dial documentation](https://developers.telnyx.com/api-reference/texml-rest-commands/initiate-an-outbound-call)
// for enabling AMD. Recommended settings: MachineDetection=Enable, AsyncAmd=true,
// DetectionMode=Premium.
type TelephonySettingsVoicemailDetection struct {
	// Action to take when voicemail is detected.
	OnVoicemailDetected TelephonySettingsVoicemailDetectionOnVoicemailDetected `json:"on_voicemail_detected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OnVoicemailDetected respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonySettingsVoicemailDetection) RawJSON() string { return r.JSON.raw }
func (r *TelephonySettingsVoicemailDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take when voicemail is detected.
type TelephonySettingsVoicemailDetectionOnVoicemailDetected struct {
	// The action to take when voicemail is detected.
	//
	// Any of "stop_assistant", "leave_message_and_stop_assistant",
	// "continue_assistant".
	Action string `json:"action"`
	// Configuration for the voicemail message to leave. Only applicable when action is
	// 'leave_message_and_stop_assistant'.
	VoicemailMessage TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessage `json:"voicemail_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action           respjson.Field
		VoicemailMessage respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonySettingsVoicemailDetectionOnVoicemailDetected) RawJSON() string { return r.JSON.raw }
func (r *TelephonySettingsVoicemailDetectionOnVoicemailDetected) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the voicemail message to leave. Only applicable when action is
// 'leave_message_and_stop_assistant'.
type TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessage struct {
	// The specific message to leave as voicemail. Only applicable when type is
	// 'message'.
	Message string `json:"message"`
	// The prompt to use for generating the voicemail message. Only applicable when
	// type is 'prompt'.
	Prompt string `json:"prompt"`
	// The type of voicemail message. Use 'prompt' to have the assistant generate a
	// message based on a prompt, or 'message' to leave a specific message.
	//
	// Any of "prompt", "message".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Prompt      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessage) RawJSON() string {
	return r.JSON.raw
}
func (r *TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonySettingsParam struct {
	// Default Texml App used for voice calls with your assistant. This will be created
	// automatically on assistant creation.
	DefaultTexmlAppID param.Opt[string] `json:"default_texml_app_id,omitzero"`
	// When enabled, allows users to interact with your AI assistant directly from your
	// website without requiring authentication. This is required for FE widgets that
	// work with assistants that have telephony enabled.
	SupportsUnauthenticatedWebCalls param.Opt[bool] `json:"supports_unauthenticated_web_calls,omitzero"`
	// Maximum duration in seconds for the AI assistant to participate on the call.
	// When this limit is reached the assistant will be stopped. This limit does not
	// apply to portions of a call without an active assistant (for instance, a call
	// transferred to a human representative).
	TimeLimitSecs param.Opt[int64] `json:"time_limit_secs,omitzero"`
	// Duration in seconds of end user silence before the assistant checks in on the
	// user. When this limit is reached the assistant will prompt the user to respond.
	// This is distinct from user_idle_timeout_secs which stops the assistant entirely.
	UserIdleReplySecs param.Opt[int64] `json:"user_idle_reply_secs,omitzero"`
	// Maximum duration in seconds of end user silence on the call. When this limit is
	// reached the assistant will be stopped. This limit does not apply to portions of
	// a call without an active assistant (for instance, a call transferred to a human
	// representative).
	UserIdleTimeoutSecs param.Opt[int64] `json:"user_idle_timeout_secs,omitzero"`
	// The noise suppression engine to use. Use 'disabled' to turn off noise
	// suppression.
	//
	// Any of "krisp", "deepfilternet", "disabled".
	NoiseSuppression TelephonySettingsNoiseSuppression `json:"noise_suppression,omitzero"`
	// Configuration for noise suppression. Only applicable when noise_suppression is
	// 'deepfilternet'.
	NoiseSuppressionConfig TelephonySettingsNoiseSuppressionConfigParam `json:"noise_suppression_config,omitzero"`
	// Configuration for call recording format and channel settings.
	RecordingSettings TelephonySettingsRecordingSettingsParam `json:"recording_settings,omitzero"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on
	// outgoing calls. These settings only apply if AMD is enabled on the Dial command.
	// See
	// [TeXML Dial documentation](https://developers.telnyx.com/api-reference/texml-rest-commands/initiate-an-outbound-call)
	// for enabling AMD. Recommended settings: MachineDetection=Enable, AsyncAmd=true,
	// DetectionMode=Premium.
	VoicemailDetection TelephonySettingsVoicemailDetectionParam `json:"voicemail_detection,omitzero"`
	paramObj
}

func (r TelephonySettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for noise suppression. Only applicable when noise_suppression is
// 'deepfilternet'.
type TelephonySettingsNoiseSuppressionConfigParam struct {
	// Attenuation limit for noise suppression. Range: 0-100.
	AttenuationLimit param.Opt[int64] `json:"attenuation_limit,omitzero"`
	// Mode for noise suppression configuration.
	//
	// Any of "advanced".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r TelephonySettingsNoiseSuppressionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsNoiseSuppressionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsNoiseSuppressionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelephonySettingsNoiseSuppressionConfigParam](
		"mode", "advanced",
	)
}

// Configuration for call recording format and channel settings.
type TelephonySettingsRecordingSettingsParam struct {
	// Whether call recording is enabled. When set to false, calls will not be recorded
	// regardless of other recording configuration.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// When enabled, the call recording will stop when the conversation ends (for
	// example, when the assistant hangs up or the call is transferred). When disabled,
	// recording continues until the call itself ends.
	StopOnConversationEnd param.Opt[bool] `json:"stop_on_conversation_end,omitzero"`
	// The number of channels for the recording. 'single' for mono, 'dual' for stereo.
	//
	// Any of "single", "dual".
	Channels string `json:"channels,omitzero"`
	// The format of the recording file.
	//
	// Any of "wav", "mp3".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r TelephonySettingsRecordingSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsRecordingSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsRecordingSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelephonySettingsRecordingSettingsParam](
		"channels", "single", "dual",
	)
	apijson.RegisterFieldValidator[TelephonySettingsRecordingSettingsParam](
		"format", "wav", "mp3",
	)
}

// Configuration for voicemail detection (AMD - Answering Machine Detection) on
// outgoing calls. These settings only apply if AMD is enabled on the Dial command.
// See
// [TeXML Dial documentation](https://developers.telnyx.com/api-reference/texml-rest-commands/initiate-an-outbound-call)
// for enabling AMD. Recommended settings: MachineDetection=Enable, AsyncAmd=true,
// DetectionMode=Premium.
type TelephonySettingsVoicemailDetectionParam struct {
	// Action to take when voicemail is detected.
	OnVoicemailDetected TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam `json:"on_voicemail_detected,omitzero"`
	paramObj
}

func (r TelephonySettingsVoicemailDetectionParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsVoicemailDetectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsVoicemailDetectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action to take when voicemail is detected.
type TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam struct {
	// The action to take when voicemail is detected.
	//
	// Any of "stop_assistant", "leave_message_and_stop_assistant",
	// "continue_assistant".
	Action string `json:"action,omitzero"`
	// Configuration for the voicemail message to leave. Only applicable when action is
	// 'leave_message_and_stop_assistant'.
	VoicemailMessage TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam `json:"voicemail_message,omitzero"`
	paramObj
}

func (r TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelephonySettingsVoicemailDetectionOnVoicemailDetectedParam](
		"action", "stop_assistant", "leave_message_and_stop_assistant", "continue_assistant",
	)
}

// Configuration for the voicemail message to leave. Only applicable when action is
// 'leave_message_and_stop_assistant'.
type TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam struct {
	// The specific message to leave as voicemail. Only applicable when type is
	// 'message'.
	Message param.Opt[string] `json:"message,omitzero"`
	// The prompt to use for generating the voicemail message. Only applicable when
	// type is 'prompt'.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// The type of voicemail message. Use 'prompt' to have the assistant generate a
	// message based on a prompt, or 'message' to leave a specific message.
	//
	// Any of "prompt", "message".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelephonySettingsVoicemailDetectionOnVoicemailDetectedVoicemailMessageParam](
		"type", "prompt", "message",
	)
}

// Endpointing thresholds used to decide when the user has finished speaking.
// Applies to non turn-taking transcription models. For `deepgram/flux`, use
// `transcription.settings.eot_threshold` / `eot_timeout_ms` /
// `eager_eot_threshold`.
type TranscriptionEndpointingPlan struct {
	// Seconds to wait after the transcript ends without punctuation.
	OnNoPunctuationSeconds float64 `json:"on_no_punctuation_seconds"`
	// Seconds to wait after the transcript ends with a number.
	OnNumberSeconds float64 `json:"on_number_seconds"`
	// Seconds to wait after the transcript ends with punctuation.
	OnPunctuationSeconds float64 `json:"on_punctuation_seconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OnNoPunctuationSeconds respjson.Field
		OnNumberSeconds        respjson.Field
		OnPunctuationSeconds   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionEndpointingPlan) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionEndpointingPlan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TranscriptionEndpointingPlan to a
// TranscriptionEndpointingPlanParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TranscriptionEndpointingPlanParam.Overrides()
func (r TranscriptionEndpointingPlan) ToParam() TranscriptionEndpointingPlanParam {
	return param.Override[TranscriptionEndpointingPlanParam](json.RawMessage(r.RawJSON()))
}

// Endpointing thresholds used to decide when the user has finished speaking.
// Applies to non turn-taking transcription models. For `deepgram/flux`, use
// `transcription.settings.eot_threshold` / `eot_timeout_ms` /
// `eager_eot_threshold`.
type TranscriptionEndpointingPlanParam struct {
	// Seconds to wait after the transcript ends without punctuation.
	OnNoPunctuationSeconds param.Opt[float64] `json:"on_no_punctuation_seconds,omitzero"`
	// Seconds to wait after the transcript ends with a number.
	OnNumberSeconds param.Opt[float64] `json:"on_number_seconds,omitzero"`
	// Seconds to wait after the transcript ends with punctuation.
	OnPunctuationSeconds param.Opt[float64] `json:"on_punctuation_seconds,omitzero"`
	paramObj
}

func (r TranscriptionEndpointingPlanParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionEndpointingPlanParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionEndpointingPlanParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionSettings struct {
	// Integration secret identifier for the transcription provider API key. Currently
	// used for Azure transcription regions that require a customer-provided API key.
	APIKeyRef string `json:"api_key_ref"`
	// The language of the audio to be transcribed. If not set, or if set to `auto`,
	// supported models will automatically detect the language. For `deepgram/flux`,
	// supported values are: `auto` (Telnyx language detection controls the language
	// hint), `multi` (no language hint), and language-specific hints `en`, `es`, `fr`,
	// `de`, `hi`, `ru`, `pt`, `ja`, `it`, and `nl`. For `soniox/stt-rt-v4`, `auto`
	// omits the language hint and lets Soniox auto-detect; ISO 639-1 codes (e.g. `en`,
	// `es`) bias detection toward that language. For `humain/realtime`, supported
	// values are `ar`, `en`, `codeswitch` (Arabic/English code-switching), and `auto`
	// (resolves server-side to code-switching). Unlike other models, `humain/realtime`
	// does not fall back to `auto` when `language` is omitted — omitting it applies
	// `en` instead.
	Language string `json:"language"`
	// The speech to text model to be used by the voice assistant. All Deepgram models
	// are run on-premise.
	//
	//   - `deepgram/flux` is optimized for turn-taking with multilingual language hints.
	//   - `deepgram/nova-3` is multilingual with automatic language detection.
	//   - `deepgram/nova-2` is Deepgram's previous-generation multilingual model.
	//   - `azure/fast` is a multilingual Azure transcription model.
	//   - `assemblyai/universal-streaming` is a multilingual streaming model with
	//     configurable turn detection.
	//   - `xai/grok-stt` is a multilingual Grok STT model.
	//   - `soniox/stt-rt-v4` is a multilingual streaming model with automatic language
	//     detection and configurable endpointing.
	//   - `nvidia/parakeet-v3` is a multilingual transcription model with automatic
	//     language detection.
	//   - `humain/realtime` is a streaming model with native Arabic and Arabic/English
	//     code-switching support.
	//
	// Any of "deepgram/flux", "deepgram/nova-3", "deepgram/nova-2", "azure/fast",
	// "assemblyai/universal-streaming", "xai/grok-stt", "soniox/stt-rt-v4",
	// "nvidia/parakeet-v3", "humain/realtime", "distil-whisper/distil-large-v2",
	// "openai/whisper-large-v3-turbo".
	Model TranscriptionSettingsModel `json:"model"`
	// Region on third party cloud providers (currently Azure) if using one of their
	// models. Some regions require `api_key_ref`.
	Region   string                      `json:"region"`
	Settings TranscriptionSettingsConfig `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyRef   respjson.Field
		Language    respjson.Field
		Model       respjson.Field
		Region      respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionSettings) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TranscriptionSettings to a TranscriptionSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TranscriptionSettingsParam.Overrides()
func (r TranscriptionSettings) ToParam() TranscriptionSettingsParam {
	return param.Override[TranscriptionSettingsParam](json.RawMessage(r.RawJSON()))
}

// The speech to text model to be used by the voice assistant. All Deepgram models
// are run on-premise.
//
//   - `deepgram/flux` is optimized for turn-taking with multilingual language hints.
//   - `deepgram/nova-3` is multilingual with automatic language detection.
//   - `deepgram/nova-2` is Deepgram's previous-generation multilingual model.
//   - `azure/fast` is a multilingual Azure transcription model.
//   - `assemblyai/universal-streaming` is a multilingual streaming model with
//     configurable turn detection.
//   - `xai/grok-stt` is a multilingual Grok STT model.
//   - `soniox/stt-rt-v4` is a multilingual streaming model with automatic language
//     detection and configurable endpointing.
//   - `nvidia/parakeet-v3` is a multilingual transcription model with automatic
//     language detection.
//   - `humain/realtime` is a streaming model with native Arabic and Arabic/English
//     code-switching support.
type TranscriptionSettingsModel string

const (
	TranscriptionSettingsModelDeepgramFlux                 TranscriptionSettingsModel = "deepgram/flux"
	TranscriptionSettingsModelDeepgramNova3                TranscriptionSettingsModel = "deepgram/nova-3"
	TranscriptionSettingsModelDeepgramNova2                TranscriptionSettingsModel = "deepgram/nova-2"
	TranscriptionSettingsModelAzureFast                    TranscriptionSettingsModel = "azure/fast"
	TranscriptionSettingsModelAssemblyaiUniversalStreaming TranscriptionSettingsModel = "assemblyai/universal-streaming"
	TranscriptionSettingsModelXaiGrokStt                   TranscriptionSettingsModel = "xai/grok-stt"
	TranscriptionSettingsModelSonioxSttRtV4                TranscriptionSettingsModel = "soniox/stt-rt-v4"
	TranscriptionSettingsModelNvidiaParakeetV3             TranscriptionSettingsModel = "nvidia/parakeet-v3"
	TranscriptionSettingsModelHumainRealtime               TranscriptionSettingsModel = "humain/realtime"
	TranscriptionSettingsModelDistilWhisperDistilLargeV2   TranscriptionSettingsModel = "distil-whisper/distil-large-v2"
	TranscriptionSettingsModelOpenAIWhisperLargeV3Turbo    TranscriptionSettingsModel = "openai/whisper-large-v3-turbo"
)

type TranscriptionSettingsParam struct {
	// Integration secret identifier for the transcription provider API key. Currently
	// used for Azure transcription regions that require a customer-provided API key.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	// The language of the audio to be transcribed. If not set, or if set to `auto`,
	// supported models will automatically detect the language. For `deepgram/flux`,
	// supported values are: `auto` (Telnyx language detection controls the language
	// hint), `multi` (no language hint), and language-specific hints `en`, `es`, `fr`,
	// `de`, `hi`, `ru`, `pt`, `ja`, `it`, and `nl`. For `soniox/stt-rt-v4`, `auto`
	// omits the language hint and lets Soniox auto-detect; ISO 639-1 codes (e.g. `en`,
	// `es`) bias detection toward that language. For `humain/realtime`, supported
	// values are `ar`, `en`, `codeswitch` (Arabic/English code-switching), and `auto`
	// (resolves server-side to code-switching). Unlike other models, `humain/realtime`
	// does not fall back to `auto` when `language` is omitted — omitting it applies
	// `en` instead.
	Language param.Opt[string] `json:"language,omitzero"`
	// Region on third party cloud providers (currently Azure) if using one of their
	// models. Some regions require `api_key_ref`.
	Region param.Opt[string] `json:"region,omitzero"`
	// The speech to text model to be used by the voice assistant. All Deepgram models
	// are run on-premise.
	//
	//   - `deepgram/flux` is optimized for turn-taking with multilingual language hints.
	//   - `deepgram/nova-3` is multilingual with automatic language detection.
	//   - `deepgram/nova-2` is Deepgram's previous-generation multilingual model.
	//   - `azure/fast` is a multilingual Azure transcription model.
	//   - `assemblyai/universal-streaming` is a multilingual streaming model with
	//     configurable turn detection.
	//   - `xai/grok-stt` is a multilingual Grok STT model.
	//   - `soniox/stt-rt-v4` is a multilingual streaming model with automatic language
	//     detection and configurable endpointing.
	//   - `nvidia/parakeet-v3` is a multilingual transcription model with automatic
	//     language detection.
	//   - `humain/realtime` is a streaming model with native Arabic and Arabic/English
	//     code-switching support.
	//
	// Any of "deepgram/flux", "deepgram/nova-3", "deepgram/nova-2", "azure/fast",
	// "assemblyai/universal-streaming", "xai/grok-stt", "soniox/stt-rt-v4",
	// "nvidia/parakeet-v3", "humain/realtime", "distil-whisper/distil-large-v2",
	// "openai/whisper-large-v3-turbo".
	Model    TranscriptionSettingsModel       `json:"model,omitzero"`
	Settings TranscriptionSettingsConfigParam `json:"settings,omitzero"`
	paramObj
}

func (r TranscriptionSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranscriptionSettingsConfig struct {
	// Available only for deepgram/flux. Confidence threshold for eager end of turn
	// detection. Must be lower than or equal to eot_threshold. Setting this equal to
	// eot_threshold effectively disables eager end of turn.
	EagerEotThreshold float64 `json:"eager_eot_threshold"`
	// Available only for soniox/stt-rt-v4. When true, Soniox emits end-of-utterance
	// events at the cadence configured by `max_endpoint_delay_ms`.
	EnableEndpointDetection bool `json:"enable_endpoint_detection"`
	// Available only for assemblyai/universal-streaming. Confidence level required to
	// trigger an end of turn. Higher values require more certainty before ending a
	// turn.
	EndOfTurnConfidenceThreshold float64 `json:"end_of_turn_confidence_threshold"`
	// Available only for deepgram/flux. Confidence required to trigger an end of turn.
	// Higher values = more reliable turn detection but slightly increased latency.
	EotThreshold float64 `json:"eot_threshold"`
	// Available only for deepgram/flux. Maximum milliseconds of silence before forcing
	// an end of turn, regardless of confidence.
	EotTimeoutMs int64 `json:"eot_timeout_ms"`
	// Available only for soniox/stt-rt-v4. When true, Soniox streams interim
	// (non-final) results in addition to finalized transcripts.
	InterimResults bool `json:"interim_results"`
	// Available only for deepgram/nova-3 and deepgram/flux. A comma-separated list of
	// key terms to boost for recognition during transcription. Helps improve accuracy
	// for domain-specific terminology, proper nouns, or uncommon words. This field may
	// be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// using mustache syntax (e.g. `Telnyx,{{customer_name}},VoIP`). Variables are
	// resolved at call time before the value is sent to the speech-to-text engine.
	Keyterm string `json:"keyterm"`
	// Available only for soniox/stt-rt-v4. Maximum silence (in milliseconds) before
	// Soniox emits an end-of-utterance event. Only honored when
	// `enable_endpoint_detection` is true.
	MaxEndpointDelayMs int64 `json:"max_endpoint_delay_ms"`
	// Available only for assemblyai/universal-streaming. Maximum duration of silence
	// in milliseconds before forcing an end of turn.
	MaxTurnSilence int64 `json:"max_turn_silence"`
	// Available only for assemblyai/universal-streaming. Minimum duration of silence
	// in milliseconds before a turn can end. Must be less than or equal to
	// max_turn_silence.
	MinTurnSilence int64 `json:"min_turn_silence"`
	Numerals       bool  `json:"numerals"`
	SmartFormat    bool  `json:"smart_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EagerEotThreshold            respjson.Field
		EnableEndpointDetection      respjson.Field
		EndOfTurnConfidenceThreshold respjson.Field
		EotThreshold                 respjson.Field
		EotTimeoutMs                 respjson.Field
		InterimResults               respjson.Field
		Keyterm                      respjson.Field
		MaxEndpointDelayMs           respjson.Field
		MaxTurnSilence               respjson.Field
		MinTurnSilence               respjson.Field
		Numerals                     respjson.Field
		SmartFormat                  respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranscriptionSettingsConfig) RawJSON() string { return r.JSON.raw }
func (r *TranscriptionSettingsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TranscriptionSettingsConfig to a
// TranscriptionSettingsConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TranscriptionSettingsConfigParam.Overrides()
func (r TranscriptionSettingsConfig) ToParam() TranscriptionSettingsConfigParam {
	return param.Override[TranscriptionSettingsConfigParam](json.RawMessage(r.RawJSON()))
}

type TranscriptionSettingsConfigParam struct {
	// Available only for deepgram/flux. Confidence threshold for eager end of turn
	// detection. Must be lower than or equal to eot_threshold. Setting this equal to
	// eot_threshold effectively disables eager end of turn.
	EagerEotThreshold param.Opt[float64] `json:"eager_eot_threshold,omitzero"`
	// Available only for soniox/stt-rt-v4. When true, Soniox emits end-of-utterance
	// events at the cadence configured by `max_endpoint_delay_ms`.
	EnableEndpointDetection param.Opt[bool] `json:"enable_endpoint_detection,omitzero"`
	// Available only for assemblyai/universal-streaming. Confidence level required to
	// trigger an end of turn. Higher values require more certainty before ending a
	// turn.
	EndOfTurnConfidenceThreshold param.Opt[float64] `json:"end_of_turn_confidence_threshold,omitzero"`
	// Available only for deepgram/flux. Confidence required to trigger an end of turn.
	// Higher values = more reliable turn detection but slightly increased latency.
	EotThreshold param.Opt[float64] `json:"eot_threshold,omitzero"`
	// Available only for deepgram/flux. Maximum milliseconds of silence before forcing
	// an end of turn, regardless of confidence.
	EotTimeoutMs param.Opt[int64] `json:"eot_timeout_ms,omitzero"`
	// Available only for soniox/stt-rt-v4. When true, Soniox streams interim
	// (non-final) results in addition to finalized transcripts.
	InterimResults param.Opt[bool] `json:"interim_results,omitzero"`
	// Available only for deepgram/nova-3 and deepgram/flux. A comma-separated list of
	// key terms to boost for recognition during transcription. Helps improve accuracy
	// for domain-specific terminology, proper nouns, or uncommon words. This field may
	// be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// using mustache syntax (e.g. `Telnyx,{{customer_name}},VoIP`). Variables are
	// resolved at call time before the value is sent to the speech-to-text engine.
	Keyterm param.Opt[string] `json:"keyterm,omitzero"`
	// Available only for soniox/stt-rt-v4. Maximum silence (in milliseconds) before
	// Soniox emits an end-of-utterance event. Only honored when
	// `enable_endpoint_detection` is true.
	MaxEndpointDelayMs param.Opt[int64] `json:"max_endpoint_delay_ms,omitzero"`
	// Available only for assemblyai/universal-streaming. Maximum duration of silence
	// in milliseconds before forcing an end of turn.
	MaxTurnSilence param.Opt[int64] `json:"max_turn_silence,omitzero"`
	// Available only for assemblyai/universal-streaming. Minimum duration of silence
	// in milliseconds before a turn can end. Must be less than or equal to
	// max_turn_silence.
	MinTurnSilence param.Opt[int64] `json:"min_turn_silence,omitzero"`
	Numerals       param.Opt[bool]  `json:"numerals,omitzero"`
	SmartFormat    param.Opt[bool]  `json:"smart_format,omitzero"`
	paramObj
}

func (r TranscriptionSettingsConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionSettingsConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionSettingsConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Transfer, Type are required.
type TransferToolParam struct {
	Transfer TransferToolTransferParam `json:"transfer,omitzero" api:"required"`
	// Any of "transfer".
	Type TransferToolType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r TransferToolParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, Targets are required.
type TransferToolTransferParam struct {
	// Number or SIP URI placing the call.
	From string `json:"from" api:"required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to. This can also be a dynamic
	// variable string like `{{ targets }}` where `targets` is returned by the dynamic
	// variables webhook and resolves to an array of target objects at runtime.
	Targets TransferToolTransferTargetsUnionParam `json:"targets,omitzero" api:"required"`
	paramObj
}

func (r TransferToolTransferParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToolTransferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToolTransferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransferToolTransferTargetsUnionParam struct {
	OfTargetsList []TransferToolTransferTargetsTargetsListItemParam `json:",omitzero,inline"`
	OfString      param.Opt[string]                                 `json:",omitzero,inline"`
	paramUnion
}

func (u TransferToolTransferTargetsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTargetsList, u.OfString)
}
func (u *TransferToolTransferTargetsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransferToolTransferTargetsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTargetsList) {
		return &u.OfTargetsList
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property To is required.
type TransferToolTransferTargetsTargetsListItemParam struct {
	// The destination number or SIP URI of the call.
	To string `json:"to" api:"required"`
	// The name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TransferToolTransferTargetsTargetsListItemParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToolTransferTargetsTargetsListItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToolTransferTargetsTargetsListItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferToolType string

const (
	TransferToolTypeTransfer TransferToolType = "transfer"
)

type VoiceSettings struct {
	// The voice to be used by the voice assistant. Check the full list of
	// [available voices](https://developers.telnyx.com/docs/tts-stt/tts-available-voices)
	// via our voices API. To use ElevenLabs, you must reference your ElevenLabs API
	// key as an integration secret under the `api_key_ref` field. See
	// [integration secrets documentation](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g.
	// Telnyx.KokoroTTS.af_heart). The voice portion of the identifier supports
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// using mustache syntax (e.g. `Telnyx.Ultra.{{voice_id}}`). The variable is
	// resolved at call time from your dynamic variables webhook, allowing you to
	// select the voice dynamically per call.
	Voice string `json:"voice" api:"required"`
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef string `json:"api_key_ref"`
	// Optional background audio to play on the call. Use a predefined media bed, or
	// supply a looped MP3 URL. If a media URL is chosen in the portal, customers can
	// preview it before saving.
	BackgroundAudio VoiceSettingsBackgroundAudioUnion `json:"background_audio"`
	// Enables emotionally expressive speech using SSML emotion tags. When enabled, the
	// assistant uses audio tags like angry, excited, content, and sad to add emotional
	// nuance. Only supported for Telnyx Ultra voices.
	ExpressiveMode bool `json:"expressive_mode"`
	// Enhances recognition for specific languages and dialects during MiniMax TTS
	// synthesis. Default is null (no boost). Set to 'auto' for automatic language
	// detection. Only applicable when using MiniMax voices.
	//
	// Any of "auto", "Chinese", "Chinese,Yue", "English", "Arabic", "Russian",
	// "Spanish", "French", "Portuguese", "German", "Turkish", "Dutch", "Ukrainian",
	// "Vietnamese", "Indonesian", "Japanese", "Italian", "Korean", "Thai", "Polish",
	// "Romanian", "Greek", "Czech", "Finnish", "Hindi", "Bulgarian", "Danish",
	// "Hebrew", "Malay", "Persian", "Slovak", "Swedish", "Croatian", "Filipino",
	// "Hungarian", "Norwegian", "Slovenian", "Catalan", "Nynorsk", "Tamil",
	// "Afrikaans".
	LanguageBoost VoiceSettingsLanguageBoost `json:"language_boost" api:"nullable"`
	// Determines how closely the AI should adhere to the original voice when
	// attempting to replicate it. Only applicable when using ElevenLabs.
	SimilarityBoost float64 `json:"similarity_boost"`
	// Adjusts speech velocity. 1.0 is default speed; values less than 1.0 slow speech;
	// values greater than 1.0 accelerate it. Only applicable when using ElevenLabs.
	Speed float64 `json:"speed"`
	// Determines the style exaggeration of the voice. Amplifies speaker style but
	// consumes additional resources when set above 0. Only applicable when using
	// ElevenLabs.
	Style float64 `json:"style"`
	// Determines how stable the voice is and the randomness between each generation.
	// Lower values create a broader emotional range; higher values produce more
	// consistent, monotonous output. Only applicable when using ElevenLabs.
	Temperature float64 `json:"temperature"`
	// Amplifies similarity to the original speaker voice. Increases computational load
	// and latency slightly. Only applicable when using ElevenLabs.
	UseSpeakerBoost bool `json:"use_speaker_boost"`
	// The speed of the voice in the range [0.25, 2.0]. 1.0 is deafult speed. Larger
	// numbers make the voice faster, smaller numbers make it slower. This is only
	// applicable for Telnyx Natural voices.
	VoiceSpeed float64 `json:"voice_speed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voice           respjson.Field
		APIKeyRef       respjson.Field
		BackgroundAudio respjson.Field
		ExpressiveMode  respjson.Field
		LanguageBoost   respjson.Field
		SimilarityBoost respjson.Field
		Speed           respjson.Field
		Style           respjson.Field
		Temperature     respjson.Field
		UseSpeakerBoost respjson.Field
		VoiceSpeed      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSettings) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VoiceSettings to a VoiceSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VoiceSettingsParam.Overrides()
func (r VoiceSettings) ToParam() VoiceSettingsParam {
	return param.Override[VoiceSettingsParam](json.RawMessage(r.RawJSON()))
}

// VoiceSettingsBackgroundAudioUnion contains all possible properties and values
// from [VoiceSettingsBackgroundAudioPredefinedMedia],
// [VoiceSettingsBackgroundAudioMediaURL], [VoiceSettingsBackgroundAudioMediaName].
//
// Use the [VoiceSettingsBackgroundAudioUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type VoiceSettingsBackgroundAudioUnion struct {
	// Any of "predefined_media", "media_url", "media_name".
	Type  string `json:"type"`
	Value string `json:"value"`
	// This field is from variant [VoiceSettingsBackgroundAudioPredefinedMedia].
	Volume float64 `json:"volume"`
	JSON   struct {
		Type   respjson.Field
		Value  respjson.Field
		Volume respjson.Field
		raw    string
	} `json:"-"`
}

// anyVoiceSettingsBackgroundAudio is implemented by each variant of
// [VoiceSettingsBackgroundAudioUnion] to add type safety for the return type of
// [VoiceSettingsBackgroundAudioUnion.AsAny]
type anyVoiceSettingsBackgroundAudio interface {
	implVoiceSettingsBackgroundAudioUnion()
}

func (VoiceSettingsBackgroundAudioPredefinedMedia) implVoiceSettingsBackgroundAudioUnion() {}
func (VoiceSettingsBackgroundAudioMediaURL) implVoiceSettingsBackgroundAudioUnion()        {}
func (VoiceSettingsBackgroundAudioMediaName) implVoiceSettingsBackgroundAudioUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := VoiceSettingsBackgroundAudioUnion.AsAny().(type) {
//	case telnyx.VoiceSettingsBackgroundAudioPredefinedMedia:
//	case telnyx.VoiceSettingsBackgroundAudioMediaURL:
//	case telnyx.VoiceSettingsBackgroundAudioMediaName:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u VoiceSettingsBackgroundAudioUnion) AsAny() anyVoiceSettingsBackgroundAudio {
	switch u.Type {
	case "predefined_media":
		return u.AsPredefinedMedia()
	case "media_url":
		return u.AsMediaURL()
	case "media_name":
		return u.AsMediaName()
	}
	return nil
}

func (u VoiceSettingsBackgroundAudioUnion) AsPredefinedMedia() (v VoiceSettingsBackgroundAudioPredefinedMedia) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSettingsBackgroundAudioUnion) AsMediaURL() (v VoiceSettingsBackgroundAudioMediaURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSettingsBackgroundAudioUnion) AsMediaName() (v VoiceSettingsBackgroundAudioMediaName) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSettingsBackgroundAudioUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSettingsBackgroundAudioUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSettingsBackgroundAudioPredefinedMedia struct {
	// Select from predefined media options.
	Type constant.PredefinedMedia `json:"type" default:"predefined_media"`
	// The predefined media to use. `silence` disables background audio.
	//
	// Any of "silence", "office".
	Value string `json:"value" api:"required"`
	// Volume level for the predefined background audio. Supports values from 0.1 to
	// 1.0 in 0.1 increments.
	Volume float64 `json:"volume"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		Volume      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSettingsBackgroundAudioPredefinedMedia) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettingsBackgroundAudioPredefinedMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSettingsBackgroundAudioMediaURL struct {
	// Provide a direct URL to an MP3 file. The audio will loop during the call.
	Type constant.MediaURL `json:"type" default:"media_url"`
	// HTTPS URL to an MP3 file.
	Value string `json:"value" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSettingsBackgroundAudioMediaURL) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettingsBackgroundAudioMediaURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSettingsBackgroundAudioMediaName struct {
	// Reference a previously uploaded media by its name from Telnyx Media Storage.
	Type constant.MediaName `json:"type" default:"media_name"`
	// The `name` of a media asset created via
	// [Media Storage API](https://developers.telnyx.com/api/media-storage/create-media-storage).
	// The audio will loop during the call.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSettingsBackgroundAudioMediaName) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettingsBackgroundAudioMediaName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enhances recognition for specific languages and dialects during MiniMax TTS
// synthesis. Default is null (no boost). Set to 'auto' for automatic language
// detection. Only applicable when using MiniMax voices.
type VoiceSettingsLanguageBoost string

const (
	VoiceSettingsLanguageBoostAuto       VoiceSettingsLanguageBoost = "auto"
	VoiceSettingsLanguageBoostChinese    VoiceSettingsLanguageBoost = "Chinese"
	VoiceSettingsLanguageBoostChineseYue VoiceSettingsLanguageBoost = "Chinese,Yue"
	VoiceSettingsLanguageBoostEnglish    VoiceSettingsLanguageBoost = "English"
	VoiceSettingsLanguageBoostArabic     VoiceSettingsLanguageBoost = "Arabic"
	VoiceSettingsLanguageBoostRussian    VoiceSettingsLanguageBoost = "Russian"
	VoiceSettingsLanguageBoostSpanish    VoiceSettingsLanguageBoost = "Spanish"
	VoiceSettingsLanguageBoostFrench     VoiceSettingsLanguageBoost = "French"
	VoiceSettingsLanguageBoostPortuguese VoiceSettingsLanguageBoost = "Portuguese"
	VoiceSettingsLanguageBoostGerman     VoiceSettingsLanguageBoost = "German"
	VoiceSettingsLanguageBoostTurkish    VoiceSettingsLanguageBoost = "Turkish"
	VoiceSettingsLanguageBoostDutch      VoiceSettingsLanguageBoost = "Dutch"
	VoiceSettingsLanguageBoostUkrainian  VoiceSettingsLanguageBoost = "Ukrainian"
	VoiceSettingsLanguageBoostVietnamese VoiceSettingsLanguageBoost = "Vietnamese"
	VoiceSettingsLanguageBoostIndonesian VoiceSettingsLanguageBoost = "Indonesian"
	VoiceSettingsLanguageBoostJapanese   VoiceSettingsLanguageBoost = "Japanese"
	VoiceSettingsLanguageBoostItalian    VoiceSettingsLanguageBoost = "Italian"
	VoiceSettingsLanguageBoostKorean     VoiceSettingsLanguageBoost = "Korean"
	VoiceSettingsLanguageBoostThai       VoiceSettingsLanguageBoost = "Thai"
	VoiceSettingsLanguageBoostPolish     VoiceSettingsLanguageBoost = "Polish"
	VoiceSettingsLanguageBoostRomanian   VoiceSettingsLanguageBoost = "Romanian"
	VoiceSettingsLanguageBoostGreek      VoiceSettingsLanguageBoost = "Greek"
	VoiceSettingsLanguageBoostCzech      VoiceSettingsLanguageBoost = "Czech"
	VoiceSettingsLanguageBoostFinnish    VoiceSettingsLanguageBoost = "Finnish"
	VoiceSettingsLanguageBoostHindi      VoiceSettingsLanguageBoost = "Hindi"
	VoiceSettingsLanguageBoostBulgarian  VoiceSettingsLanguageBoost = "Bulgarian"
	VoiceSettingsLanguageBoostDanish     VoiceSettingsLanguageBoost = "Danish"
	VoiceSettingsLanguageBoostHebrew     VoiceSettingsLanguageBoost = "Hebrew"
	VoiceSettingsLanguageBoostMalay      VoiceSettingsLanguageBoost = "Malay"
	VoiceSettingsLanguageBoostPersian    VoiceSettingsLanguageBoost = "Persian"
	VoiceSettingsLanguageBoostSlovak     VoiceSettingsLanguageBoost = "Slovak"
	VoiceSettingsLanguageBoostSwedish    VoiceSettingsLanguageBoost = "Swedish"
	VoiceSettingsLanguageBoostCroatian   VoiceSettingsLanguageBoost = "Croatian"
	VoiceSettingsLanguageBoostFilipino   VoiceSettingsLanguageBoost = "Filipino"
	VoiceSettingsLanguageBoostHungarian  VoiceSettingsLanguageBoost = "Hungarian"
	VoiceSettingsLanguageBoostNorwegian  VoiceSettingsLanguageBoost = "Norwegian"
	VoiceSettingsLanguageBoostSlovenian  VoiceSettingsLanguageBoost = "Slovenian"
	VoiceSettingsLanguageBoostCatalan    VoiceSettingsLanguageBoost = "Catalan"
	VoiceSettingsLanguageBoostNynorsk    VoiceSettingsLanguageBoost = "Nynorsk"
	VoiceSettingsLanguageBoostTamil      VoiceSettingsLanguageBoost = "Tamil"
	VoiceSettingsLanguageBoostAfrikaans  VoiceSettingsLanguageBoost = "Afrikaans"
)

// The property Voice is required.
type VoiceSettingsParam struct {
	// The voice to be used by the voice assistant. Check the full list of
	// [available voices](https://developers.telnyx.com/docs/tts-stt/tts-available-voices)
	// via our voices API. To use ElevenLabs, you must reference your ElevenLabs API
	// key as an integration secret under the `api_key_ref` field. See
	// [integration secrets documentation](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g.
	// Telnyx.KokoroTTS.af_heart). The voice portion of the identifier supports
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// using mustache syntax (e.g. `Telnyx.Ultra.{{voice_id}}`). The variable is
	// resolved at call time from your dynamic variables webhook, allowing you to
	// select the voice dynamically per call.
	Voice string `json:"voice" api:"required"`
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	// Enables emotionally expressive speech using SSML emotion tags. When enabled, the
	// assistant uses audio tags like angry, excited, content, and sad to add emotional
	// nuance. Only supported for Telnyx Ultra voices.
	ExpressiveMode param.Opt[bool] `json:"expressive_mode,omitzero"`
	// Determines how closely the AI should adhere to the original voice when
	// attempting to replicate it. Only applicable when using ElevenLabs.
	SimilarityBoost param.Opt[float64] `json:"similarity_boost,omitzero"`
	// Adjusts speech velocity. 1.0 is default speed; values less than 1.0 slow speech;
	// values greater than 1.0 accelerate it. Only applicable when using ElevenLabs.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Determines the style exaggeration of the voice. Amplifies speaker style but
	// consumes additional resources when set above 0. Only applicable when using
	// ElevenLabs.
	Style param.Opt[float64] `json:"style,omitzero"`
	// Determines how stable the voice is and the randomness between each generation.
	// Lower values create a broader emotional range; higher values produce more
	// consistent, monotonous output. Only applicable when using ElevenLabs.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Amplifies similarity to the original speaker voice. Increases computational load
	// and latency slightly. Only applicable when using ElevenLabs.
	UseSpeakerBoost param.Opt[bool] `json:"use_speaker_boost,omitzero"`
	// The speed of the voice in the range [0.25, 2.0]. 1.0 is deafult speed. Larger
	// numbers make the voice faster, smaller numbers make it slower. This is only
	// applicable for Telnyx Natural voices.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
	// Enhances recognition for specific languages and dialects during MiniMax TTS
	// synthesis. Default is null (no boost). Set to 'auto' for automatic language
	// detection. Only applicable when using MiniMax voices.
	//
	// Any of "auto", "Chinese", "Chinese,Yue", "English", "Arabic", "Russian",
	// "Spanish", "French", "Portuguese", "German", "Turkish", "Dutch", "Ukrainian",
	// "Vietnamese", "Indonesian", "Japanese", "Italian", "Korean", "Thai", "Polish",
	// "Romanian", "Greek", "Czech", "Finnish", "Hindi", "Bulgarian", "Danish",
	// "Hebrew", "Malay", "Persian", "Slovak", "Swedish", "Croatian", "Filipino",
	// "Hungarian", "Norwegian", "Slovenian", "Catalan", "Nynorsk", "Tamil",
	// "Afrikaans".
	LanguageBoost VoiceSettingsLanguageBoost `json:"language_boost,omitzero"`
	// Optional background audio to play on the call. Use a predefined media bed, or
	// supply a looped MP3 URL. If a media URL is chosen in the portal, customers can
	// preview it before saving.
	BackgroundAudio VoiceSettingsBackgroundAudioUnionParam `json:"background_audio,omitzero"`
	paramObj
}

func (r VoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VoiceSettingsBackgroundAudioUnionParam struct {
	OfPredefinedMedia *VoiceSettingsBackgroundAudioPredefinedMediaParam `json:",omitzero,inline"`
	OfMediaURL        *VoiceSettingsBackgroundAudioMediaURLParam        `json:",omitzero,inline"`
	OfMediaName       *VoiceSettingsBackgroundAudioMediaNameParam       `json:",omitzero,inline"`
	paramUnion
}

func (u VoiceSettingsBackgroundAudioUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPredefinedMedia, u.OfMediaURL, u.OfMediaName)
}
func (u *VoiceSettingsBackgroundAudioUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VoiceSettingsBackgroundAudioUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPredefinedMedia) {
		return u.OfPredefinedMedia
	} else if !param.IsOmitted(u.OfMediaURL) {
		return u.OfMediaURL
	} else if !param.IsOmitted(u.OfMediaName) {
		return u.OfMediaName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VoiceSettingsBackgroundAudioUnionParam) GetVolume() *float64 {
	if vt := u.OfPredefinedMedia; vt != nil && vt.Volume.Valid() {
		return &vt.Volume.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VoiceSettingsBackgroundAudioUnionParam) GetType() *string {
	if vt := u.OfPredefinedMedia; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMediaURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMediaName; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VoiceSettingsBackgroundAudioUnionParam) GetValue() *string {
	if vt := u.OfPredefinedMedia; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfMediaURL; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfMediaName; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[VoiceSettingsBackgroundAudioUnionParam](
		"type",
		apijson.Discriminator[VoiceSettingsBackgroundAudioPredefinedMediaParam]("predefined_media"),
		apijson.Discriminator[VoiceSettingsBackgroundAudioMediaURLParam]("media_url"),
		apijson.Discriminator[VoiceSettingsBackgroundAudioMediaNameParam]("media_name"),
	)
}

// The properties Type, Value are required.
type VoiceSettingsBackgroundAudioPredefinedMediaParam struct {
	// The predefined media to use. `silence` disables background audio.
	//
	// Any of "silence", "office".
	Value string `json:"value,omitzero" api:"required"`
	// Volume level for the predefined background audio. Supports values from 0.1 to
	// 1.0 in 0.1 increments.
	Volume param.Opt[float64] `json:"volume,omitzero"`
	// Select from predefined media options.
	//
	// This field can be elided, and will marshal its zero value as "predefined_media".
	Type constant.PredefinedMedia `json:"type" default:"predefined_media"`
	paramObj
}

func (r VoiceSettingsBackgroundAudioPredefinedMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceSettingsBackgroundAudioPredefinedMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceSettingsBackgroundAudioPredefinedMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceSettingsBackgroundAudioPredefinedMediaParam](
		"value", "silence", "office",
	)
}

// The properties Type, Value are required.
type VoiceSettingsBackgroundAudioMediaURLParam struct {
	// HTTPS URL to an MP3 file.
	Value string `json:"value" api:"required" format:"uri"`
	// Provide a direct URL to an MP3 file. The audio will loop during the call.
	//
	// This field can be elided, and will marshal its zero value as "media_url".
	Type constant.MediaURL `json:"type" default:"media_url"`
	paramObj
}

func (r VoiceSettingsBackgroundAudioMediaURLParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceSettingsBackgroundAudioMediaURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceSettingsBackgroundAudioMediaURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Value are required.
type VoiceSettingsBackgroundAudioMediaNameParam struct {
	// The `name` of a media asset created via
	// [Media Storage API](https://developers.telnyx.com/api/media-storage/create-media-storage).
	// The audio will loop during the call.
	Value string `json:"value" api:"required"`
	// Reference a previously uploaded media by its name from Telnyx Media Storage.
	//
	// This field can be elided, and will marshal its zero value as "media_name".
	Type constant.MediaName `json:"type" default:"media_name"`
	paramObj
}

func (r VoiceSettingsBackgroundAudioMediaNameParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceSettingsBackgroundAudioMediaNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceSettingsBackgroundAudioMediaNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, Webhook are required.
type WebhookToolParam struct {
	// Any of "webhook".
	Type    WebhookToolType         `json:"type,omitzero" api:"required"`
	Webhook WebhookToolWebhookParam `json:"webhook,omitzero" api:"required"`
	paramObj
}

func (r WebhookToolParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookToolType string

const (
	WebhookToolTypeWebhook WebhookToolType = "webhook"
)

// The properties Description, Name, URL are required.
type WebhookToolWebhookParam struct {
	// The description of the tool.
	Description string `json:"description" api:"required"`
	// The name of the tool.
	Name string `json:"name" api:"required"`
	// The URL of the external tool to be called. This URL is going to be used by the
	// assistant. The URL can be templated like: `https://example.com/api/v1/{id}`,
	// where `{id}` is a placeholder for a value that will be provided by the assistant
	// if `path_parameters` are provided with the `id` attribute.
	URL string `json:"url" api:"required"`
	// The body parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the body of the request. See
	// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
	// for documentation about the format
	BodyParameters WebhookToolWebhookBodyParametersParam `json:"body_parameters,omitzero"`
	// The headers to be sent to the external tool.
	Headers []WebhookToolWebhookHeaderParam `json:"headers,omitzero"`
	// The HTTP method to be used when calling the external tool.
	//
	// Any of "GET", "POST", "PUT", "DELETE", "PATCH".
	Method string `json:"method,omitzero"`
	// The path parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the path of the request if the
	// URL contains a placeholder for a value. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	PathParameters WebhookToolWebhookPathParametersParam `json:"path_parameters,omitzero"`
	// The query parameters the webhook tool accepts, described as a JSON Schema
	// object. These parameters will be passed to the webhook as the query of the
	// request. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	QueryParameters WebhookToolWebhookQueryParametersParam `json:"query_parameters,omitzero"`
	paramObj
}

func (r WebhookToolWebhookParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolWebhookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolWebhookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookToolWebhookParam](
		"method", "GET", "POST", "PUT", "DELETE", "PATCH",
	)
}

// The body parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the body of the request. See
// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
// for documentation about the format
type WebhookToolWebhookBodyParametersParam struct {
	// The properties of the body parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the body parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WebhookToolWebhookBodyParametersParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolWebhookBodyParametersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolWebhookBodyParametersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookToolWebhookBodyParametersParam](
		"type", "object",
	)
}

type WebhookToolWebhookHeaderParam struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `Bearer {{#integration_secret}}test-secret{{/integration_secret}}` to pass the
	// value of the integration secret as the bearer token.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r WebhookToolWebhookHeaderParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolWebhookHeaderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolWebhookHeaderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The path parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the path of the request if the
// URL contains a placeholder for a value. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type WebhookToolWebhookPathParametersParam struct {
	// The properties of the path parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the path parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WebhookToolWebhookPathParametersParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolWebhookPathParametersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolWebhookPathParametersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookToolWebhookPathParametersParam](
		"type", "object",
	)
}

// The query parameters the webhook tool accepts, described as a JSON Schema
// object. These parameters will be passed to the webhook as the query of the
// request. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type WebhookToolWebhookQueryParametersParam struct {
	// The properties of the query parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the query parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WebhookToolWebhookQueryParametersParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolWebhookQueryParametersParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolWebhookQueryParametersParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookToolWebhookQueryParametersParam](
		"type", "object",
	)
}

// Configuration settings for the assistant's web widget.
type WidgetSettings struct {
	// Text displayed while the agent is processing.
	AgentThinkingText     string                `json:"agent_thinking_text"`
	AudioVisualizerConfig AudioVisualizerConfig `json:"audio_visualizer_config"`
	// The default state of the widget.
	//
	// Any of "expanded", "collapsed".
	DefaultState WidgetSettingsDefaultState `json:"default_state"`
	// URL for users to give feedback.
	GiveFeedbackURL string `json:"give_feedback_url" api:"nullable"`
	// URL to a custom logo icon for the widget.
	LogoIconURL string `json:"logo_icon_url" api:"nullable"`
	// The positioning style for the widget.
	//
	// Any of "fixed", "static".
	Position WidgetSettingsPosition `json:"position"`
	// URL for users to report issues.
	ReportIssueURL string `json:"report_issue_url" api:"nullable"`
	// Text prompting users to speak to interrupt.
	SpeakToInterruptText string `json:"speak_to_interrupt_text"`
	// Custom text displayed on the start call button.
	StartCallText string `json:"start_call_text"`
	// The visual theme for the widget.
	//
	// Any of "light", "dark".
	Theme WidgetSettingsTheme `json:"theme"`
	// URL to view conversation history.
	ViewHistoryURL string `json:"view_history_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentThinkingText     respjson.Field
		AudioVisualizerConfig respjson.Field
		DefaultState          respjson.Field
		GiveFeedbackURL       respjson.Field
		LogoIconURL           respjson.Field
		Position              respjson.Field
		ReportIssueURL        respjson.Field
		SpeakToInterruptText  respjson.Field
		StartCallText         respjson.Field
		Theme                 respjson.Field
		ViewHistoryURL        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WidgetSettings) RawJSON() string { return r.JSON.raw }
func (r *WidgetSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WidgetSettings to a WidgetSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WidgetSettingsParam.Overrides()
func (r WidgetSettings) ToParam() WidgetSettingsParam {
	return param.Override[WidgetSettingsParam](json.RawMessage(r.RawJSON()))
}

// The default state of the widget.
type WidgetSettingsDefaultState string

const (
	WidgetSettingsDefaultStateExpanded  WidgetSettingsDefaultState = "expanded"
	WidgetSettingsDefaultStateCollapsed WidgetSettingsDefaultState = "collapsed"
)

// The positioning style for the widget.
type WidgetSettingsPosition string

const (
	WidgetSettingsPositionFixed  WidgetSettingsPosition = "fixed"
	WidgetSettingsPositionStatic WidgetSettingsPosition = "static"
)

// The visual theme for the widget.
type WidgetSettingsTheme string

const (
	WidgetSettingsThemeLight WidgetSettingsTheme = "light"
	WidgetSettingsThemeDark  WidgetSettingsTheme = "dark"
)

// Configuration settings for the assistant's web widget.
type WidgetSettingsParam struct {
	// URL for users to give feedback.
	GiveFeedbackURL param.Opt[string] `json:"give_feedback_url,omitzero"`
	// URL to a custom logo icon for the widget.
	LogoIconURL param.Opt[string] `json:"logo_icon_url,omitzero"`
	// URL for users to report issues.
	ReportIssueURL param.Opt[string] `json:"report_issue_url,omitzero"`
	// URL to view conversation history.
	ViewHistoryURL param.Opt[string] `json:"view_history_url,omitzero"`
	// Text displayed while the agent is processing.
	AgentThinkingText param.Opt[string] `json:"agent_thinking_text,omitzero"`
	// Text prompting users to speak to interrupt.
	SpeakToInterruptText param.Opt[string] `json:"speak_to_interrupt_text,omitzero"`
	// Custom text displayed on the start call button.
	StartCallText         param.Opt[string]          `json:"start_call_text,omitzero"`
	AudioVisualizerConfig AudioVisualizerConfigParam `json:"audio_visualizer_config,omitzero"`
	// The default state of the widget.
	//
	// Any of "expanded", "collapsed".
	DefaultState WidgetSettingsDefaultState `json:"default_state,omitzero"`
	// The positioning style for the widget.
	//
	// Any of "fixed", "static".
	Position WidgetSettingsPosition `json:"position,omitzero"`
	// The visual theme for the widget.
	//
	// Any of "light", "dark".
	Theme WidgetSettingsTheme `json:"theme,omitzero"`
	paramObj
}

func (r WidgetSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow WidgetSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WidgetSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aligns with the OpenAI API:
// https://platform.openai.com/docs/api-reference/assistants/deleteAssistant
type AIAssistantDeleteResponse struct {
	ID      string `json:"id" api:"required"`
	Deleted bool   `json:"deleted" api:"required"`
	Object  string `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantChatResponse struct {
	// The assistant's generated response based on the input message and context.
	Content string `json:"content" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantChatResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantChatResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantSendSMSResponse struct {
	ConversationID string `json:"conversation_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAssistantSendSMSResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantSendSMSResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantNewParams struct {
	// System instructions for the assistant. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Instructions string            `json:"instructions" api:"required"`
	Name         string            `json:"name" api:"required"`
	Description  param.Opt[string] `json:"description,omitzero"`
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
	// Conversation flow as supplied by API clients (create / update).
	//
	// A directed graph of `FlowNodeReq` connected by `FlowEdge`s. Validation enforces
	// unique node/edge IDs, that `start_node_id` references a real node, and that
	// every edge's endpoints reference real nodes.
	ConversationFlow ConversationFlowReqParam `json:"conversation_flow,omitzero"`
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
	Tools         []AssistantToolUnionParam  `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam         `json:"voice_settings,omitzero"`
	// Configuration settings for the assistant's web widget.
	WidgetSettings WidgetSettingsParam `json:"widget_settings,omitzero"`
	paramObj
}

func (r AIAssistantNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantGetParams struct {
	// Filter results by call control id.
	CallControlID param.Opt[string] `query:"call_control_id,omitzero" json:"-"`
	// Whether to fetch dynamic variables from the configured webhook.
	FetchDynamicVariablesFromWebhook param.Opt[bool] `query:"fetch_dynamic_variables_from_webhook,omitzero" json:"-"`
	// Start of the filter range.
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	// End of the filter range.
	To param.Opt[string] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAssistantGetParams]'s query parameters as `url.Values`.
func (r AIAssistantGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAssistantUpdateParams struct {
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
	// Indicates whether the assistant should be promoted to the main version. Defaults
	// to true.
	PromoteToMain param.Opt[bool] `json:"promote_to_main,omitzero"`
	// Human-readable name for the assistant version.
	VersionName param.Opt[string] `json:"version_name,omitzero"`
	// Conversation flow as supplied by API clients (create / update).
	//
	// A directed graph of `FlowNodeReq` connected by `FlowEdge`s. Validation enforces
	// unique node/edge IDs, that `start_node_id` references a real node, and that
	// every edge's endpoints reference real nodes.
	ConversationFlow ConversationFlowReqParam `json:"conversation_flow,omitzero"`
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
	Tools         []AssistantToolUnionParam  `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam         `json:"voice_settings,omitzero"`
	// Configuration settings for the assistant's web widget.
	WidgetSettings WidgetSettingsParam `json:"widget_settings,omitzero"`
	paramObj
}

func (r AIAssistantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantChatParams struct {
	// The message content sent by the client to the assistant
	Content string `json:"content" api:"required"`
	// A unique identifier for the conversation thread, used to maintain context
	ConversationID string `json:"conversation_id" api:"required"`
	// The optional display name of the user sending the message
	Name param.Opt[string] `json:"name,omitzero"`
	// When true, the response is streamed as Server-Sent Events (`text/event-stream`):
	// `delta` events carry content fragments as they are generated, a final `done`
	// event carries the full content plus `whatsapp_template`, and a terminal `error`
	// event reports failures that happen after streaming started. When false
	// (default), the response is a single JSON object.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	paramObj
}

func (r AIAssistantChatParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantChatParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantChatParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantImportsParams struct {
	// Integration secret pointer that refers to the API key for the external provider.
	// This should be an identifier for an integration secret created via
	// /v2/integration_secrets.
	APIKeyRef string `json:"api_key_ref" api:"required"`
	// The external provider to import assistants from.
	//
	// Any of "elevenlabs", "vapi", "retell".
	Provider AIAssistantImportsParamsProvider `json:"provider,omitzero" api:"required"`
	// Optional list of assistant IDs to import from the external provider. If not
	// provided, all assistants will be imported.
	ImportIDs []string `json:"import_ids,omitzero"`
	paramObj
}

func (r AIAssistantImportsParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantImportsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantImportsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The external provider to import assistants from.
type AIAssistantImportsParamsProvider string

const (
	AIAssistantImportsParamsProviderElevenlabs AIAssistantImportsParamsProvider = "elevenlabs"
	AIAssistantImportsParamsProviderVapi       AIAssistantImportsParamsProvider = "vapi"
	AIAssistantImportsParamsProviderRetell     AIAssistantImportsParamsProvider = "retell"
)

type AIAssistantSendSMSParams struct {
	From                     string                                                       `json:"from" api:"required"`
	To                       string                                                       `json:"to" api:"required"`
	ShouldCreateConversation param.Opt[bool]                                              `json:"should_create_conversation,omitzero"`
	Text                     param.Opt[string]                                            `json:"text,omitzero"`
	ConversationMetadata     map[string]AIAssistantSendSMSParamsConversationMetadataUnion `json:"conversation_metadata,omitzero"`
	paramObj
}

func (r AIAssistantSendSMSParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantSendSMSParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantSendSMSParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AIAssistantSendSMSParamsConversationMetadataUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u AIAssistantSendSMSParamsConversationMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt, u.OfBool)
}
func (u *AIAssistantSendSMSParamsConversationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AIAssistantSendSMSParamsConversationMetadataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}
