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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// AIAssistantService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantService] method instead.
type AIAssistantService struct {
	Options         []option.RequestOption
	Tests           AIAssistantTestService
	CanaryDeploys   AIAssistantCanaryDeployService
	ScheduledEvents AIAssistantScheduledEventService
	Tools           AIAssistantToolService
	Versions        AIAssistantVersionService
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
	return
}

// Create a new AI Assistant.
func (r *AIAssistantService) New(ctx context.Context, body AIAssistantNewParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an AI Assistant configuration by `assistant_id`.
func (r *AIAssistantService) Get(ctx context.Context, assistantID string, query AIAssistantGetParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update an AI Assistant's attributes.
func (r *AIAssistantService) Update(ctx context.Context, assistantID string, body AIAssistantUpdateParams, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a list of all AI Assistants configured by the user.
func (r *AIAssistantService) List(ctx context.Context, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an AI Assistant by `assistant_id`.
func (r *AIAssistantService) Delete(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *AIAssistantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/chat", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Clone an existing assistant, excluding telephony and messaging settings.
func (r *AIAssistantService) Clone(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *InferenceEmbedding, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/clone", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get an assistant texml by `assistant_id`.
func (r *AIAssistantService) GetTexml(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/texml", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Import assistants from external providers. Any assistant that has already been
// imported will be overwritten with its latest version from the importing
// provider.
func (r *AIAssistantService) Imports(ctx context.Context, body AIAssistantImportsParams, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/chat/sms", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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
	Tools []AssistantToolUnionParam `json:"tools,omitzero"`
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
type AssistantToolUnionParam struct {
	OfBookAppointment   *AssistantToolBookAppointmentParam   `json:",omitzero,inline"`
	OfCheckAvailability *AssistantToolCheckAvailabilityParam `json:",omitzero,inline"`
	OfWebhook           *WebhookToolParam                    `json:",omitzero,inline"`
	OfHangup            *HangupToolParam                     `json:",omitzero,inline"`
	OfTransfer          *TransferToolParam                   `json:",omitzero,inline"`
	OfRetrieval         *AssistantToolRetrievalParam         `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBookAppointment,
		u.OfCheckAvailability,
		u.OfWebhook,
		u.OfHangup,
		u.OfTransfer,
		u.OfRetrieval)
}
func (u *AssistantToolUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolUnionParam) asAny() any {
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
func (u AssistantToolUnionParam) GetBookAppointment() *AssistantToolBookAppointmentBookAppointmentParam {
	if vt := u.OfBookAppointment; vt != nil {
		return &vt.BookAppointment
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetCheckAvailability() *AssistantToolCheckAvailabilityCheckAvailabilityParam {
	if vt := u.OfCheckAvailability; vt != nil {
		return &vt.CheckAvailability
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetWebhook() *WebhookToolWebhookParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Webhook
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
func (u AssistantToolUnionParam) GetTransfer() *TransferToolTransferParam {
	if vt := u.OfTransfer; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetRetrieval() *AssistantToolRetrievalRetrievalParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetType() *string {
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
	apijson.RegisterUnion[AssistantToolUnionParam](
		"type",
		apijson.Discriminator[AssistantToolBookAppointmentParam]("book_appointment"),
		apijson.Discriminator[AssistantToolCheckAvailabilityParam]("check_availability"),
		apijson.Discriminator[WebhookToolParam]("webhook"),
		apijson.Discriminator[HangupToolParam]("hangup"),
		apijson.Discriminator[TransferToolParam]("transfer"),
		apijson.Discriminator[AssistantToolRetrievalParam]("retrieval"),
	)
}

// The properties BookAppointment, Type are required.
type AssistantToolBookAppointmentParam struct {
	BookAppointment AssistantToolBookAppointmentBookAppointmentParam `json:"book_appointment,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "book_appointment".
	Type constant.BookAppointment `json:"type,required"`
	paramObj
}

func (r AssistantToolBookAppointmentParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolBookAppointmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolBookAppointmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties APIKeyRef, EventTypeID are required.
type AssistantToolBookAppointmentBookAppointmentParam struct {
	// Reference to an integration secret that contains your Cal.com API key. You would
	// pass the `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your Cal.com API key.
	APIKeyRef string `json:"api_key_ref,required"`
	// Event Type ID for which slots are being fetched.
	// [cal.com](https://cal.com/docs/api-reference/v2/bookings/create-a-booking#body-event-type-id)
	EventTypeID int64 `json:"event_type_id,required"`
	// The name of the attendee
	// [cal.com](https://cal.com/docs/api-reference/v2/bookings/create-a-booking#body-attendee-name).
	// If not provided, the assistant will ask for the attendee's name.
	AttendeeName param.Opt[string] `json:"attendee_name,omitzero"`
	// The timezone of the attendee
	// [cal.com](https://cal.com/docs/api-reference/v2/bookings/create-a-booking#body-attendee-timezone).
	// If not provided, the assistant will ask for the attendee's timezone.
	AttendeeTimezone param.Opt[string] `json:"attendee_timezone,omitzero"`
	paramObj
}

func (r AssistantToolBookAppointmentBookAppointmentParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolBookAppointmentBookAppointmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolBookAppointmentBookAppointmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CheckAvailability, Type are required.
type AssistantToolCheckAvailabilityParam struct {
	CheckAvailability AssistantToolCheckAvailabilityCheckAvailabilityParam `json:"check_availability,omitzero,required"`
	// This field can be elided, and will marshal its zero value as
	// "check_availability".
	Type constant.CheckAvailability `json:"type,required"`
	paramObj
}

func (r AssistantToolCheckAvailabilityParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolCheckAvailabilityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolCheckAvailabilityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties APIKeyRef, EventTypeID are required.
type AssistantToolCheckAvailabilityCheckAvailabilityParam struct {
	// Reference to an integration secret that contains your Cal.com API key. You would
	// pass the `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your Cal.com API key.
	APIKeyRef string `json:"api_key_ref,required"`
	// Event Type ID for which slots are being fetched.
	// [cal.com](https://cal.com/docs/api-reference/v2/slots/get-available-slots#parameter-event-type-id)
	EventTypeID int64 `json:"event_type_id,required"`
	paramObj
}

func (r AssistantToolCheckAvailabilityCheckAvailabilityParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolCheckAvailabilityCheckAvailabilityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolCheckAvailabilityCheckAvailabilityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Retrieval, Type are required.
type AssistantToolRetrievalParam struct {
	Retrieval AssistantToolRetrievalRetrievalParam `json:"retrieval,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "retrieval".
	Type constant.Retrieval `json:"type,required"`
	paramObj
}

func (r AssistantToolRetrievalParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolRetrievalParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolRetrievalParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property BucketIDs is required.
type AssistantToolRetrievalRetrievalParam struct {
	BucketIDs []string `json:"bucket_ids,omitzero,required"`
	// The maximum number of results to retrieve as context for the language model.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	paramObj
}

func (r AssistantToolRetrievalRetrievalParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolRetrievalRetrievalParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolRetrievalRetrievalParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantToolsItemsUnion contains all possible properties and values from
// [InferenceEmbeddingWebhookToolParamsResp], [RetrievalTool],
// [AssistantToolHandoff], [HangupTool], [AssistantToolTransfer],
// [AssistantToolRefer], [AssistantToolSendDtmf], [AssistantToolSendMessage],
// [AssistantToolSkipTurn].
//
// Use the [AssistantToolsItemsUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantToolsItemsUnion struct {
	// Any of "webhook", "retrieval", "handoff", "hangup", "transfer", "refer",
	// "send_dtmf", "send_message", "skip_turn".
	Type string `json:"type"`
	// This field is from variant [InferenceEmbeddingWebhookToolParamsResp].
	Webhook InferenceEmbeddingWebhookToolParamsWebhookResp `json:"webhook"`
	// This field is from variant [RetrievalTool].
	Retrieval BucketIDs `json:"retrieval"`
	// This field is from variant [AssistantToolHandoff].
	Handoff AssistantToolHandoffHandoff `json:"handoff"`
	// This field is from variant [HangupTool].
	Hangup HangupToolParamsResp `json:"hangup"`
	// This field is from variant [AssistantToolTransfer].
	Transfer AssistantToolTransferTransfer `json:"transfer"`
	// This field is from variant [AssistantToolRefer].
	Refer AssistantToolReferRefer `json:"refer"`
	// This field is from variant [AssistantToolSendDtmf].
	SendDtmf map[string]any `json:"send_dtmf"`
	// This field is from variant [AssistantToolSendMessage].
	SendMessage map[string]any `json:"send_message"`
	// This field is from variant [AssistantToolSkipTurn].
	SkipTurn AssistantToolSkipTurnSkipTurn `json:"skip_turn"`
	JSON     struct {
		Type        respjson.Field
		Webhook     respjson.Field
		Retrieval   respjson.Field
		Handoff     respjson.Field
		Hangup      respjson.Field
		Transfer    respjson.Field
		Refer       respjson.Field
		SendDtmf    respjson.Field
		SendMessage respjson.Field
		SkipTurn    respjson.Field
		raw         string
	} `json:"-"`
}

// anyAssistantToolsItems is implemented by each variant of
// [AssistantToolsItemsUnion] to add type safety for the return type of
// [AssistantToolsItemsUnion.AsAny]
type anyAssistantToolsItems interface {
	implAssistantToolsItemsUnion()
}

func (InferenceEmbeddingWebhookToolParamsResp) implAssistantToolsItemsUnion() {}
func (RetrievalTool) implAssistantToolsItemsUnion()                           {}
func (AssistantToolHandoff) implAssistantToolsItemsUnion()                    {}
func (HangupTool) implAssistantToolsItemsUnion()                              {}
func (AssistantToolTransfer) implAssistantToolsItemsUnion()                   {}
func (AssistantToolRefer) implAssistantToolsItemsUnion()                      {}
func (AssistantToolSendDtmf) implAssistantToolsItemsUnion()                   {}
func (AssistantToolSendMessage) implAssistantToolsItemsUnion()                {}
func (AssistantToolSkipTurn) implAssistantToolsItemsUnion()                   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssistantToolsItemsUnion.AsAny().(type) {
//	case telnyx.InferenceEmbeddingWebhookToolParamsResp:
//	case telnyx.RetrievalTool:
//	case telnyx.AssistantToolHandoff:
//	case telnyx.HangupTool:
//	case telnyx.AssistantToolTransfer:
//	case telnyx.AssistantToolRefer:
//	case telnyx.AssistantToolSendDtmf:
//	case telnyx.AssistantToolSendMessage:
//	case telnyx.AssistantToolSkipTurn:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AssistantToolsItemsUnion) AsAny() anyAssistantToolsItems {
	switch u.Type {
	case "webhook":
		return u.AsWebhook()
	case "retrieval":
		return u.AsRetrieval()
	case "handoff":
		return u.AsHandoff()
	case "hangup":
		return u.AsHangup()
	case "transfer":
		return u.AsTransfer()
	case "refer":
		return u.AsRefer()
	case "send_dtmf":
		return u.AsSendDtmf()
	case "send_message":
		return u.AsSendMessage()
	case "skip_turn":
		return u.AsSkipTurn()
	}
	return nil
}

func (u AssistantToolsItemsUnion) AsWebhook() (v InferenceEmbeddingWebhookToolParamsResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsRetrieval() (v RetrievalTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsHandoff() (v AssistantToolHandoff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsHangup() (v HangupTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsTransfer() (v AssistantToolTransfer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsRefer() (v AssistantToolRefer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsSendDtmf() (v AssistantToolSendDtmf) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsSendMessage() (v AssistantToolSendMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolsItemsUnion) AsSkipTurn() (v AssistantToolSkipTurn) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssistantToolsItemsUnion) RawJSON() string { return u.JSON.raw }

func (r *AssistantToolsItemsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AssistantToolsItemsUnion to a
// AssistantToolsItemsUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AssistantToolsItemsUnionParam.Overrides()
func (r AssistantToolsItemsUnion) ToParam() AssistantToolsItemsUnionParam {
	return param.Override[AssistantToolsItemsUnionParam](json.RawMessage(r.RawJSON()))
}

// The handoff tool allows the assistant to hand off control of the conversation to
// another AI assistant. By default, this will happen transparently to the end
// user.
type AssistantToolHandoff struct {
	Handoff AssistantToolHandoffHandoff `json:"handoff,required"`
	Type    constant.Handoff            `json:"type,required"`
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
	AIAssistants []AssistantToolHandoffHandoffAIAssistant `json:"ai_assistants,required"`
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
	ID string `json:"id,required"`
	// Helpful name for giving context on when to handoff to the assistant.
	Name string `json:"name,required"`
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
	Transfer AssistantToolTransferTransfer `json:"transfer,required"`
	Type     constant.Transfer             `json:"type,required"`
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
	From string `json:"from,required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to.
	Targets []AssistantToolTransferTransferTarget `json:"targets,required"`
	// Custom headers to be added to the SIP INVITE for the transfer command.
	CustomHeaders []AssistantToolTransferTransferCustomHeader `json:"custom_headers"`
	// Configuration for voicemail detection (AMD - Answering Machine Detection) on the
	// transferred call. Allows the assistant to detect when a voicemail system answers
	// the transferred call and take appropriate action.
	VoicemailDetection AssistantToolTransferTransferVoicemailDetection `json:"voicemail_detection"`
	// Natural language instructions for your agent for how to provide context for the
	// transfer recipient.
	WarmTransferInstructions string `json:"warm_transfer_instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From                     respjson.Field
		Targets                  respjson.Field
		CustomHeaders            respjson.Field
		VoicemailDetection       respjson.Field
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

type AssistantToolTransferTransferTarget struct {
	// The name of the target.
	Name string `json:"name"`
	// The destination number or SIP URI of the call.
	To string `json:"to"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolTransferTransferTarget) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolTransferTransferTarget) UnmarshalJSON(data []byte) error {
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

type AssistantToolRefer struct {
	Refer AssistantToolReferRefer `json:"refer,required"`
	Type  constant.Refer          `json:"type,required"`
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
	Targets []AssistantToolReferReferTarget `json:"targets,required"`
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
	Name string `json:"name,required"`
	// The SIP URI to which the call will be referred.
	SipAddress string `json:"sip_address,required"`
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
	SendDtmf map[string]any    `json:"send_dtmf,required"`
	Type     constant.SendDtmf `json:"type,required"`
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
// conversation context, and the message text is generated by the assistant.
type AssistantToolSendMessage struct {
	SendMessage map[string]any       `json:"send_message,required"`
	Type        constant.SendMessage `json:"type,required"`
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

type AssistantToolSkipTurn struct {
	SkipTurn AssistantToolSkipTurnSkipTurn `json:"skip_turn,required"`
	Type     constant.SkipTurn             `json:"type,required"`
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

func AssistantToolsItemsParamOfWebhook(webhook InferenceEmbeddingWebhookToolParamsWebhook) AssistantToolsItemsUnionParam {
	var variant InferenceEmbeddingWebhookToolParams
	variant.Webhook = webhook
	return AssistantToolsItemsUnionParam{OfWebhook: &variant}
}

func AssistantToolsItemsParamOfRetrieval(retrieval BucketIDsParam) AssistantToolsItemsUnionParam {
	var variant RetrievalToolParam
	variant.Retrieval = retrieval
	return AssistantToolsItemsUnionParam{OfRetrieval: &variant}
}

func AssistantToolsItemsParamOfHandoff(handoff AssistantToolHandoffHandoffParam) AssistantToolsItemsUnionParam {
	var variant AssistantToolHandoffParam
	variant.Handoff = handoff
	return AssistantToolsItemsUnionParam{OfHandoff: &variant}
}

func AssistantToolsItemsParamOfHangup(hangup HangupToolParams) AssistantToolsItemsUnionParam {
	var variant HangupToolParam
	variant.Hangup = hangup
	return AssistantToolsItemsUnionParam{OfHangup: &variant}
}

func AssistantToolsItemsParamOfTransfer(transfer AssistantToolTransferTransferParam) AssistantToolsItemsUnionParam {
	var variant AssistantToolTransferParam
	variant.Transfer = transfer
	return AssistantToolsItemsUnionParam{OfTransfer: &variant}
}

func AssistantToolsItemsParamOfRefer(refer AssistantToolReferReferParam) AssistantToolsItemsUnionParam {
	var variant AssistantToolReferParam
	variant.Refer = refer
	return AssistantToolsItemsUnionParam{OfRefer: &variant}
}

func AssistantToolsItemsParamOfSendDtmf(sendDtmf map[string]any) AssistantToolsItemsUnionParam {
	var variant AssistantToolSendDtmfParam
	variant.SendDtmf = sendDtmf
	return AssistantToolsItemsUnionParam{OfSendDtmf: &variant}
}

func AssistantToolsItemsParamOfSendMessage(sendMessage map[string]any) AssistantToolsItemsUnionParam {
	var variant AssistantToolSendMessageParam
	variant.SendMessage = sendMessage
	return AssistantToolsItemsUnionParam{OfSendMessage: &variant}
}

func AssistantToolsItemsParamOfSkipTurn(skipTurn AssistantToolSkipTurnSkipTurnParam) AssistantToolsItemsUnionParam {
	var variant AssistantToolSkipTurnParam
	variant.SkipTurn = skipTurn
	return AssistantToolsItemsUnionParam{OfSkipTurn: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantToolsItemsUnionParam struct {
	OfWebhook     *InferenceEmbeddingWebhookToolParams `json:",omitzero,inline"`
	OfRetrieval   *RetrievalToolParam                  `json:",omitzero,inline"`
	OfHandoff     *AssistantToolHandoffParam           `json:",omitzero,inline"`
	OfHangup      *HangupToolParam                     `json:",omitzero,inline"`
	OfTransfer    *AssistantToolTransferParam          `json:",omitzero,inline"`
	OfRefer       *AssistantToolReferParam             `json:",omitzero,inline"`
	OfSendDtmf    *AssistantToolSendDtmfParam          `json:",omitzero,inline"`
	OfSendMessage *AssistantToolSendMessageParam       `json:",omitzero,inline"`
	OfSkipTurn    *AssistantToolSkipTurnParam          `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolsItemsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhook,
		u.OfRetrieval,
		u.OfHandoff,
		u.OfHangup,
		u.OfTransfer,
		u.OfRefer,
		u.OfSendDtmf,
		u.OfSendMessage,
		u.OfSkipTurn)
}
func (u *AssistantToolsItemsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolsItemsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfRetrieval) {
		return u.OfRetrieval
	} else if !param.IsOmitted(u.OfHandoff) {
		return u.OfHandoff
	} else if !param.IsOmitted(u.OfHangup) {
		return u.OfHangup
	} else if !param.IsOmitted(u.OfTransfer) {
		return u.OfTransfer
	} else if !param.IsOmitted(u.OfRefer) {
		return u.OfRefer
	} else if !param.IsOmitted(u.OfSendDtmf) {
		return u.OfSendDtmf
	} else if !param.IsOmitted(u.OfSendMessage) {
		return u.OfSendMessage
	} else if !param.IsOmitted(u.OfSkipTurn) {
		return u.OfSkipTurn
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetWebhook() *InferenceEmbeddingWebhookToolParamsWebhook {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetRetrieval() *BucketIDsParam {
	if vt := u.OfRetrieval; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetHandoff() *AssistantToolHandoffHandoffParam {
	if vt := u.OfHandoff; vt != nil {
		return &vt.Handoff
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetHangup() *HangupToolParams {
	if vt := u.OfHangup; vt != nil {
		return &vt.Hangup
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetTransfer() *AssistantToolTransferTransferParam {
	if vt := u.OfTransfer; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetRefer() *AssistantToolReferReferParam {
	if vt := u.OfRefer; vt != nil {
		return &vt.Refer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetSendDtmf() map[string]any {
	if vt := u.OfSendDtmf; vt != nil {
		return vt.SendDtmf
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetSendMessage() map[string]any {
	if vt := u.OfSendMessage; vt != nil {
		return vt.SendMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetSkipTurn() *AssistantToolSkipTurnSkipTurnParam {
	if vt := u.OfSkipTurn; vt != nil {
		return &vt.SkipTurn
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetType() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrieval; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHandoff; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHangup; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTransfer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRefer; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSendDtmf; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSendMessage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSkipTurn; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AssistantToolsItemsUnionParam](
		"type",
		apijson.Discriminator[InferenceEmbeddingWebhookToolParams]("webhook"),
		apijson.Discriminator[RetrievalToolParam]("retrieval"),
		apijson.Discriminator[AssistantToolHandoffParam]("handoff"),
		apijson.Discriminator[HangupToolParam]("hangup"),
		apijson.Discriminator[AssistantToolTransferParam]("transfer"),
		apijson.Discriminator[AssistantToolReferParam]("refer"),
		apijson.Discriminator[AssistantToolSendDtmfParam]("send_dtmf"),
		apijson.Discriminator[AssistantToolSendMessageParam]("send_message"),
		apijson.Discriminator[AssistantToolSkipTurnParam]("skip_turn"),
	)
}

// The handoff tool allows the assistant to hand off control of the conversation to
// another AI assistant. By default, this will happen transparently to the end
// user.
//
// The properties Handoff, Type are required.
type AssistantToolHandoffParam struct {
	Handoff AssistantToolHandoffHandoffParam `json:"handoff,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "handoff".
	Type constant.Handoff `json:"type,required"`
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
	AIAssistants []AssistantToolHandoffHandoffAIAssistantParam `json:"ai_assistants,omitzero,required"`
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
	ID string `json:"id,required"`
	// Helpful name for giving context on when to handoff to the assistant.
	Name string `json:"name,required"`
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
	Transfer AssistantToolTransferTransferParam `json:"transfer,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "transfer".
	Type constant.Transfer `json:"type,required"`
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
	From string `json:"from,required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to.
	Targets []AssistantToolTransferTransferTargetParam `json:"targets,omitzero,required"`
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

type AssistantToolTransferTransferTargetParam struct {
	// The name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// The destination number or SIP URI of the call.
	To param.Opt[string] `json:"to,omitzero"`
	paramObj
}

func (r AssistantToolTransferTransferTargetParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolTransferTransferTargetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolTransferTransferTargetParam) UnmarshalJSON(data []byte) error {
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

// The properties Refer, Type are required.
type AssistantToolReferParam struct {
	Refer AssistantToolReferReferParam `json:"refer,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "refer".
	Type constant.Refer `json:"type,required"`
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
	Targets []AssistantToolReferReferTargetParam `json:"targets,omitzero,required"`
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
	Name string `json:"name,required"`
	// The SIP URI to which the call will be referred.
	SipAddress string `json:"sip_address,required"`
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
	SendDtmf map[string]any `json:"send_dtmf,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "send_dtmf".
	Type constant.SendDtmf `json:"type,required"`
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
// conversation context, and the message text is generated by the assistant.
//
// The properties SendMessage, Type are required.
type AssistantToolSendMessageParam struct {
	SendMessage map[string]any `json:"send_message,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "send_message".
	Type constant.SendMessage `json:"type,required"`
	paramObj
}

func (r AssistantToolSendMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolSendMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolSendMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties SkipTurn, Type are required.
type AssistantToolSkipTurnParam struct {
	SkipTurn AssistantToolSkipTurnSkipTurnParam `json:"skip_turn,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "skip_turn".
	Type constant.SkipTurn `json:"type,required"`
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

type AssistantsList struct {
	Data []InferenceEmbedding `json:"data,required"`
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

// If `telephony` is enabled, the assistant will be able to make and receive calls.
// If `messaging` is enabled, the assistant will be able to send and receive
// messages.
type EnabledFeatures string

const (
	EnabledFeaturesTelephony EnabledFeatures = "telephony"
	EnabledFeaturesMessaging EnabledFeatures = "messaging"
)

type HangupTool struct {
	Hangup HangupToolParamsResp `json:"hangup,required"`
	// Any of "hangup".
	Type HangupToolType `json:"type,required"`
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
	Hangup HangupToolParams `json:"hangup,omitzero,required"`
	// Any of "hangup".
	Type HangupToolType `json:"type,omitzero,required"`
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
	ID        string    `json:"id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// System instructions for the assistant. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Instructions string `json:"instructions,required"`
	// ID of the model to use. You can use the
	// [Get models API](https://developers.telnyx.com/api-reference/chat/get-available-models)
	// to see all of your available models,
	Model       string `json:"model,required"`
	Name        string `json:"name,required"`
	Description string `json:"description"`
	// Map of dynamic variables and their values
	DynamicVariables map[string]any `json:"dynamic_variables"`
	// If the dynamic_variables_webhook_url is set for the assistant, we will send a
	// request at the start of the conversation. See our
	// [guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	// for more information.
	DynamicVariablesWebhookURL string            `json:"dynamic_variables_webhook_url"`
	EnabledFeatures            []EnabledFeatures `json:"enabled_features"`
	// Text that the assistant will use to start the conversation. This may be
	// templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables).
	// Use an empty string to have the assistant wait for the user to speak first. Use
	// the special value `<assistant-speaks-first-with-model-generated-message>` to
	// have the assistant generate the greeting based on the system instructions.
	Greeting        string          `json:"greeting"`
	ImportMetadata  ImportMetadata  `json:"import_metadata"`
	InsightSettings InsightSettings `json:"insight_settings"`
	// This is only needed when using third-party inference providers. The `identifier`
	// for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your LLM provider's API key. Warning: Free plans are unlikely to
	// work with this integration.
	LlmAPIKeyRef      string            `json:"llm_api_key_ref"`
	MessagingSettings MessagingSettings `json:"messaging_settings"`
	PrivacySettings   PrivacySettings   `json:"privacy_settings"`
	TelephonySettings TelephonySettings `json:"telephony_settings"`
	// The tools that the assistant can use. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Tools         []AssistantToolsItemsUnion `json:"tools"`
	Transcription TranscriptionSettings      `json:"transcription"`
	VoiceSettings VoiceSettings              `json:"voice_settings"`
	// Configuration settings for the assistant's web widget.
	WidgetSettings WidgetSettings `json:"widget_settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		CreatedAt                  respjson.Field
		Instructions               respjson.Field
		Model                      respjson.Field
		Name                       respjson.Field
		Description                respjson.Field
		DynamicVariables           respjson.Field
		DynamicVariablesWebhookURL respjson.Field
		EnabledFeatures            respjson.Field
		Greeting                   respjson.Field
		ImportMetadata             respjson.Field
		InsightSettings            respjson.Field
		LlmAPIKeyRef               respjson.Field
		MessagingSettings          respjson.Field
		PrivacySettings            respjson.Field
		TelephonySettings          respjson.Field
		Tools                      respjson.Field
		Transcription              respjson.Field
		VoiceSettings              respjson.Field
		WidgetSettings             respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbedding) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingWebhookToolParamsResp struct {
	// Any of "webhook".
	Type    InferenceEmbeddingWebhookToolParamsType        `json:"type,required"`
	Webhook InferenceEmbeddingWebhookToolParamsWebhookResp `json:"webhook,required"`
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
	Description string `json:"description,required"`
	// The name of the tool.
	Name string `json:"name,required"`
	// The URL of the external tool to be called. This URL is going to be used by the
	// assistant. The URL can be templated like: `https://example.com/api/v1/{id}`,
	// where `{id}` is a placeholder for a value that will be provided by the assistant
	// if `path_parameters` are provided with the `id` attribute.
	URL string `json:"url,required"`
	// If async, the assistant will move forward without waiting for your server to
	// respond.
	Async bool `json:"async"`
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
	// The maximum number of milliseconds to wait for the webhook to respond. Only
	// applicable when async is false.
	TimeoutMs int64 `json:"timeout_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description     respjson.Field
		Name            respjson.Field
		URL             respjson.Field
		Async           respjson.Field
		BodyParameters  respjson.Field
		Headers         respjson.Field
		Method          respjson.Field
		PathParameters  respjson.Field
		QueryParameters respjson.Field
		TimeoutMs       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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

// The properties Type, Webhook are required.
type InferenceEmbeddingWebhookToolParams struct {
	// Any of "webhook".
	Type    InferenceEmbeddingWebhookToolParamsType    `json:"type,omitzero,required"`
	Webhook InferenceEmbeddingWebhookToolParamsWebhook `json:"webhook,omitzero,required"`
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
	Description string `json:"description,required"`
	// The name of the tool.
	Name string `json:"name,required"`
	// The URL of the external tool to be called. This URL is going to be used by the
	// assistant. The URL can be templated like: `https://example.com/api/v1/{id}`,
	// where `{id}` is a placeholder for a value that will be provided by the assistant
	// if `path_parameters` are provided with the `id` attribute.
	URL string `json:"url,required"`
	// If async, the assistant will move forward without waiting for your server to
	// respond.
	Async param.Opt[bool] `json:"async,omitzero"`
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

type RetrievalTool struct {
	Retrieval BucketIDs `json:"retrieval,required"`
	// Any of "retrieval".
	Type RetrievalToolType `json:"type,required"`
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
	Retrieval BucketIDsParam `json:"retrieval,omitzero,required"`
	// Any of "retrieval".
	Type RetrievalToolType `json:"type,omitzero,required"`
	paramObj
}

func (r RetrievalToolParam) MarshalJSON() (data []byte, err error) {
	type shadow RetrievalToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RetrievalToolParam) UnmarshalJSON(data []byte) error {
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
	// The format of the recording file.
	//
	// Any of "wav", "mp3".
	Format string `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Channels    respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type TranscriptionSettings struct {
	// The language of the audio to be transcribed. If not set, of if set to `auto`,
	// the model will automatically detect the language.
	Language string `json:"language"`
	// The speech to text model to be used by the voice assistant. All the deepgram
	// models are run on-premise.
	//
	//   - `deepgram/flux` is optimized for turn-taking but is English-only.
	//   - `deepgram/nova-3` is multi-lingual with automatic language detection but
	//     slightly higher latency.
	//
	// Any of "deepgram/flux", "deepgram/nova-3", "deepgram/nova-2", "azure/fast",
	// "distil-whisper/distil-large-v2", "openai/whisper-large-v3-turbo".
	Model TranscriptionSettingsModel `json:"model"`
	// Region on third party cloud providers (currently Azure) if using one of their
	// models
	Region   string                      `json:"region"`
	Settings TranscriptionSettingsConfig `json:"settings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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

// The speech to text model to be used by the voice assistant. All the deepgram
// models are run on-premise.
//
//   - `deepgram/flux` is optimized for turn-taking but is English-only.
//   - `deepgram/nova-3` is multi-lingual with automatic language detection but
//     slightly higher latency.
type TranscriptionSettingsModel string

const (
	TranscriptionSettingsModelDeepgramFlux               TranscriptionSettingsModel = "deepgram/flux"
	TranscriptionSettingsModelDeepgramNova3              TranscriptionSettingsModel = "deepgram/nova-3"
	TranscriptionSettingsModelDeepgramNova2              TranscriptionSettingsModel = "deepgram/nova-2"
	TranscriptionSettingsModelAzureFast                  TranscriptionSettingsModel = "azure/fast"
	TranscriptionSettingsModelDistilWhisperDistilLargeV2 TranscriptionSettingsModel = "distil-whisper/distil-large-v2"
	TranscriptionSettingsModelOpenAIWhisperLargeV3Turbo  TranscriptionSettingsModel = "openai/whisper-large-v3-turbo"
)

type TranscriptionSettingsParam struct {
	// The language of the audio to be transcribed. If not set, of if set to `auto`,
	// the model will automatically detect the language.
	Language param.Opt[string] `json:"language,omitzero"`
	// Region on third party cloud providers (currently Azure) if using one of their
	// models
	Region param.Opt[string] `json:"region,omitzero"`
	// The speech to text model to be used by the voice assistant. All the deepgram
	// models are run on-premise.
	//
	//   - `deepgram/flux` is optimized for turn-taking but is English-only.
	//   - `deepgram/nova-3` is multi-lingual with automatic language detection but
	//     slightly higher latency.
	//
	// Any of "deepgram/flux", "deepgram/nova-3", "deepgram/nova-2", "azure/fast",
	// "distil-whisper/distil-large-v2", "openai/whisper-large-v3-turbo".
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
	// Available only for deepgram/flux. Confidence required to trigger an end of turn.
	// Higher values = more reliable turn detection but slightly increased latency.
	EotThreshold float64 `json:"eot_threshold"`
	// Available only for deepgram/flux. Maximum milliseconds of silence before forcing
	// an end of turn, regardless of confidence.
	EotTimeoutMs int64 `json:"eot_timeout_ms"`
	Numerals     bool  `json:"numerals"`
	SmartFormat  bool  `json:"smart_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EagerEotThreshold respjson.Field
		EotThreshold      respjson.Field
		EotTimeoutMs      respjson.Field
		Numerals          respjson.Field
		SmartFormat       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
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
	// Available only for deepgram/flux. Confidence required to trigger an end of turn.
	// Higher values = more reliable turn detection but slightly increased latency.
	EotThreshold param.Opt[float64] `json:"eot_threshold,omitzero"`
	// Available only for deepgram/flux. Maximum milliseconds of silence before forcing
	// an end of turn, regardless of confidence.
	EotTimeoutMs param.Opt[int64] `json:"eot_timeout_ms,omitzero"`
	Numerals     param.Opt[bool]  `json:"numerals,omitzero"`
	SmartFormat  param.Opt[bool]  `json:"smart_format,omitzero"`
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
	Transfer TransferToolTransferParam `json:"transfer,omitzero,required"`
	// Any of "transfer".
	Type TransferToolType `json:"type,omitzero,required"`
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
	From string `json:"from,required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to.
	Targets []TransferToolTransferTargetParam `json:"targets,omitzero,required"`
	paramObj
}

func (r TransferToolTransferParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToolTransferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToolTransferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferToolTransferTargetParam struct {
	// The name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// The destination number or SIP URI of the call.
	To param.Opt[string] `json:"to,omitzero"`
	paramObj
}

func (r TransferToolTransferTargetParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToolTransferTargetParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToolTransferTargetParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferToolType string

const (
	TransferToolTypeTransfer TransferToolType = "transfer"
)

type VoiceSettings struct {
	// The voice to be used by the voice assistant. Check the full list of
	// [available voices](https://developers.telnyx.com/api/call-control/list-text-to-speech-voices)
	// via our voices API. To use ElevenLabs, you must reference your ElevenLabs API
	// key as an integration secret under the `api_key_ref` field. See
	// [integration secrets documentation](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g.
	// Telnyx.KokoroTTS.af_heart)
	Voice string `json:"voice,required"`
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef string `json:"api_key_ref"`
	// Optional background audio to play on the call. Use a predefined media bed, or
	// supply a looped MP3 URL. If a media URL is chosen in the portal, customers can
	// preview it before saving.
	BackgroundAudio VoiceSettingsBackgroundAudioUnion `json:"background_audio"`
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
	LanguageBoost VoiceSettingsLanguageBoost `json:"language_boost,nullable"`
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
	JSON  struct {
		Type  respjson.Field
		Value respjson.Field
		raw   string
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
	Type constant.PredefinedMedia `json:"type,required"`
	// The predefined media to use. `silence` disables background audio.
	//
	// Any of "silence", "office".
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
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
	Type constant.MediaURL `json:"type,required"`
	// HTTPS URL to an MP3 file.
	Value string `json:"value,required" format:"uri"`
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
	Type constant.MediaName `json:"type,required"`
	// The `name` of a media asset created via
	// [Media Storage API](https://developers.telnyx.com/api/media-storage/create-media-storage).
	// The audio will loop during the call.
	Value string `json:"value,required"`
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
	// [available voices](https://developers.telnyx.com/api/call-control/list-text-to-speech-voices)
	// via our voices API. To use ElevenLabs, you must reference your ElevenLabs API
	// key as an integration secret under the `api_key_ref` field. See
	// [integration secrets documentation](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g.
	// Telnyx.KokoroTTS.af_heart)
	Voice string `json:"voice,required"`
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
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
	Value string `json:"value,omitzero,required"`
	// Select from predefined media options.
	//
	// This field can be elided, and will marshal its zero value as "predefined_media".
	Type constant.PredefinedMedia `json:"type,required"`
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
	Value string `json:"value,required" format:"uri"`
	// Provide a direct URL to an MP3 file. The audio will loop during the call.
	//
	// This field can be elided, and will marshal its zero value as "media_url".
	Type constant.MediaURL `json:"type,required"`
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
	Value string `json:"value,required"`
	// Reference a previously uploaded media by its name from Telnyx Media Storage.
	//
	// This field can be elided, and will marshal its zero value as "media_name".
	Type constant.MediaName `json:"type,required"`
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
	Type    WebhookToolType         `json:"type,omitzero,required"`
	Webhook WebhookToolWebhookParam `json:"webhook,omitzero,required"`
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
	Description string `json:"description,required"`
	// The name of the tool.
	Name string `json:"name,required"`
	// The URL of the external tool to be called. This URL is going to be used by the
	// assistant. The URL can be templated like: `https://example.com/api/v1/{id}`,
	// where `{id}` is a placeholder for a value that will be provided by the assistant
	// if `path_parameters` are provided with the `id` attribute.
	URL string `json:"url,required"`
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
	GiveFeedbackURL string `json:"give_feedback_url,nullable"`
	// URL to a custom logo icon for the widget.
	LogoIconURL string `json:"logo_icon_url,nullable"`
	// The positioning style for the widget.
	//
	// Any of "fixed", "static".
	Position WidgetSettingsPosition `json:"position"`
	// URL for users to report issues.
	ReportIssueURL string `json:"report_issue_url,nullable"`
	// Text prompting users to speak to interrupt.
	SpeakToInterruptText string `json:"speak_to_interrupt_text"`
	// Custom text displayed on the start call button.
	StartCallText string `json:"start_call_text"`
	// The visual theme for the widget.
	//
	// Any of "light", "dark".
	Theme WidgetSettingsTheme `json:"theme"`
	// URL to view conversation history.
	ViewHistoryURL string `json:"view_history_url,nullable"`
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
	ID      string `json:"id,required"`
	Deleted bool   `json:"deleted,required"`
	Object  string `json:"object,required"`
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
	Content string `json:"content,required"`
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
	Instructions string `json:"instructions,required"`
	// ID of the model to use. You can use the
	// [Get models API](https://developers.telnyx.com/api-reference/chat/get-available-models)
	// to see all of your available models,
	Model       string            `json:"model,required"`
	Name        string            `json:"name,required"`
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
	// This is only needed when using third-party inference providers. The `identifier`
	// for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api-reference/integration-secrets/create-a-secret)
	// that refers to your LLM provider's API key. Warning: Free plans are unlikely to
	// work with this integration.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// Map of dynamic variables and their default values
	DynamicVariables  map[string]any         `json:"dynamic_variables,omitzero"`
	EnabledFeatures   []EnabledFeatures      `json:"enabled_features,omitzero"`
	InsightSettings   InsightSettingsParam   `json:"insight_settings,omitzero"`
	MessagingSettings MessagingSettingsParam `json:"messaging_settings,omitzero"`
	PrivacySettings   PrivacySettingsParam   `json:"privacy_settings,omitzero"`
	TelephonySettings TelephonySettingsParam `json:"telephony_settings,omitzero"`
	// The tools that the assistant can use. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Tools         []AssistantToolsItemsUnionParam `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam      `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam              `json:"voice_settings,omitzero"`
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
	CallControlID                    param.Opt[string] `query:"call_control_id,omitzero" json:"-"`
	FetchDynamicVariablesFromWebhook param.Opt[bool]   `query:"fetch_dynamic_variables_from_webhook,omitzero" json:"-"`
	From                             param.Opt[string] `query:"from,omitzero" json:"-"`
	To                               param.Opt[string] `query:"to,omitzero" json:"-"`
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
	// Indicates whether the assistant should be promoted to the main version. Defaults
	// to true.
	PromoteToMain param.Opt[bool] `json:"promote_to_main,omitzero"`
	// Map of dynamic variables and their default values
	DynamicVariables  map[string]any         `json:"dynamic_variables,omitzero"`
	EnabledFeatures   []EnabledFeatures      `json:"enabled_features,omitzero"`
	InsightSettings   InsightSettingsParam   `json:"insight_settings,omitzero"`
	MessagingSettings MessagingSettingsParam `json:"messaging_settings,omitzero"`
	PrivacySettings   PrivacySettingsParam   `json:"privacy_settings,omitzero"`
	TelephonySettings TelephonySettingsParam `json:"telephony_settings,omitzero"`
	// The tools that the assistant can use. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Tools         []AssistantToolsItemsUnionParam `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam      `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam              `json:"voice_settings,omitzero"`
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
	Content string `json:"content,required"`
	// A unique identifier for the conversation thread, used to maintain context
	ConversationID string `json:"conversation_id,required"`
	// The optional display name of the user sending the message
	Name param.Opt[string] `json:"name,omitzero"`
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
	APIKeyRef string `json:"api_key_ref,required"`
	// The external provider to import assistants from.
	//
	// Any of "elevenlabs", "vapi", "retell".
	Provider AIAssistantImportsParamsProvider `json:"provider,omitzero,required"`
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
	From                     string                                                       `json:"from,required"`
	To                       string                                                       `json:"to,required"`
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
