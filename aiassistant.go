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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared/constant"
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
// [create a conversation](https://developers.telnyx.com/api/inference/inference-embedding/create-new-conversation-public-conversations-post),
// [filter existing conversations](https://developers.telnyx.com/api/inference/inference-embedding/get-conversations-public-conversations-get),
// [fetch messages for a conversation](https://developers.telnyx.com/api/inference/inference-embedding/get-conversations-public-conversation-id-messages-get),
// and
// [manually add messages to a conversation](https://developers.telnyx.com/api/inference/inference-embedding/add-new-message).
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
	OfRetrieval         *RetrievalToolParam                  `json:",omitzero,inline"`
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
func (u AssistantToolUnionParam) GetWebhook() *InferenceEmbeddingWebhookToolParams {
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
func (u AssistantToolUnionParam) GetTransfer() *InferenceEmbeddingTransferToolParams {
	if vt := u.OfTransfer; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetRetrieval() *InferenceEmbeddingBucketIDsParam {
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
		apijson.Discriminator[RetrievalToolParam]("retrieval"),
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

// AssistantToolsItemsUnion contains all possible properties and values from
// [WebhookTool], [RetrievalTool], [AssistantToolHandoff], [HangupTool],
// [TransferTool], [AssistantToolRefer], [AssistantToolSendDtmf].
//
// Use the [AssistantToolsItemsUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantToolsItemsUnion struct {
	// Any of "webhook", "retrieval", "handoff", "hangup", "transfer", "refer",
	// "send_dtmf".
	Type string `json:"type"`
	// This field is from variant [WebhookTool].
	Webhook InferenceEmbeddingWebhookToolParamsResp `json:"webhook"`
	// This field is from variant [RetrievalTool].
	Retrieval InferenceEmbeddingBucketIDs `json:"retrieval"`
	// This field is from variant [AssistantToolHandoff].
	Handoff AssistantToolHandoffHandoff `json:"handoff"`
	// This field is from variant [HangupTool].
	Hangup HangupToolParamsResp `json:"hangup"`
	// This field is from variant [TransferTool].
	Transfer InferenceEmbeddingTransferToolParamsResp `json:"transfer"`
	// This field is from variant [AssistantToolRefer].
	Refer AssistantToolReferRefer `json:"refer"`
	// This field is from variant [AssistantToolSendDtmf].
	SendDtmf map[string]any `json:"send_dtmf"`
	JSON     struct {
		Type      respjson.Field
		Webhook   respjson.Field
		Retrieval respjson.Field
		Handoff   respjson.Field
		Hangup    respjson.Field
		Transfer  respjson.Field
		Refer     respjson.Field
		SendDtmf  respjson.Field
		raw       string
	} `json:"-"`
}

// anyAssistantToolsItems is implemented by each variant of
// [AssistantToolsItemsUnion] to add type safety for the return type of
// [AssistantToolsItemsUnion.AsAny]
type anyAssistantToolsItems interface {
	implAssistantToolsItemsUnion()
}

func (WebhookTool) implAssistantToolsItemsUnion()           {}
func (RetrievalTool) implAssistantToolsItemsUnion()         {}
func (AssistantToolHandoff) implAssistantToolsItemsUnion()  {}
func (HangupTool) implAssistantToolsItemsUnion()            {}
func (TransferTool) implAssistantToolsItemsUnion()          {}
func (AssistantToolRefer) implAssistantToolsItemsUnion()    {}
func (AssistantToolSendDtmf) implAssistantToolsItemsUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AssistantToolsItemsUnion.AsAny().(type) {
//	case telnyx.WebhookTool:
//	case telnyx.RetrievalTool:
//	case telnyx.AssistantToolHandoff:
//	case telnyx.HangupTool:
//	case telnyx.TransferTool:
//	case telnyx.AssistantToolRefer:
//	case telnyx.AssistantToolSendDtmf:
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
	}
	return nil
}

func (u AssistantToolsItemsUnion) AsWebhook() (v WebhookTool) {
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

func (u AssistantToolsItemsUnion) AsTransfer() (v TransferTool) {
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

func AssistantToolsItemsParamOfWebhook(webhook InferenceEmbeddingWebhookToolParams) AssistantToolsItemsUnionParam {
	var variant WebhookToolParam
	variant.Webhook = webhook
	return AssistantToolsItemsUnionParam{OfWebhook: &variant}
}

func AssistantToolsItemsParamOfRetrieval(retrieval InferenceEmbeddingBucketIDsParam) AssistantToolsItemsUnionParam {
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

func AssistantToolsItemsParamOfTransfer(transfer InferenceEmbeddingTransferToolParams) AssistantToolsItemsUnionParam {
	var variant TransferToolParam
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AssistantToolsItemsUnionParam struct {
	OfWebhook   *WebhookToolParam           `json:",omitzero,inline"`
	OfRetrieval *RetrievalToolParam         `json:",omitzero,inline"`
	OfHandoff   *AssistantToolHandoffParam  `json:",omitzero,inline"`
	OfHangup    *HangupToolParam            `json:",omitzero,inline"`
	OfTransfer  *TransferToolParam          `json:",omitzero,inline"`
	OfRefer     *AssistantToolReferParam    `json:",omitzero,inline"`
	OfSendDtmf  *AssistantToolSendDtmfParam `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolsItemsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhook,
		u.OfRetrieval,
		u.OfHandoff,
		u.OfHangup,
		u.OfTransfer,
		u.OfRefer,
		u.OfSendDtmf)
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
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetWebhook() *InferenceEmbeddingWebhookToolParams {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolsItemsUnionParam) GetRetrieval() *InferenceEmbeddingBucketIDsParam {
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
func (u AssistantToolsItemsUnionParam) GetTransfer() *InferenceEmbeddingTransferToolParams {
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
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AssistantToolsItemsUnionParam](
		"type",
		apijson.Discriminator[WebhookToolParam]("webhook"),
		apijson.Discriminator[RetrievalToolParam]("retrieval"),
		apijson.Discriminator[AssistantToolHandoffParam]("handoff"),
		apijson.Discriminator[HangupToolParam]("hangup"),
		apijson.Discriminator[TransferToolParam]("transfer"),
		apijson.Discriminator[AssistantToolReferParam]("refer"),
		apijson.Discriminator[AssistantToolSendDtmfParam]("send_dtmf"),
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
	// [Get models API](https://developers.telnyx.com/api/inference/inference-embedding/get-models-public-models-get)
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
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Greeting        string          `json:"greeting"`
	ImportMetadata  ImportMetadata  `json:"import_metadata"`
	InsightSettings InsightSettings `json:"insight_settings"`
	// This is only needed when using third-party inference providers. The `identifier`
	// for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
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
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbedding) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingBucketIDs struct {
	// List of
	// [embedded storage buckets](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding)
	// to use for retrieval-augmented generation.
	BucketIDs []string `json:"bucket_ids,required"`
	// The maximum number of results to retrieve as context for the language model.
	MaxNumResults int64 `json:"max_num_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketIDs     respjson.Field
		MaxNumResults respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingBucketIDs) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingBucketIDs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InferenceEmbeddingBucketIDs to a
// InferenceEmbeddingBucketIDsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InferenceEmbeddingBucketIDsParam.Overrides()
func (r InferenceEmbeddingBucketIDs) ToParam() InferenceEmbeddingBucketIDsParam {
	return param.Override[InferenceEmbeddingBucketIDsParam](json.RawMessage(r.RawJSON()))
}

// The property BucketIDs is required.
type InferenceEmbeddingBucketIDsParam struct {
	// List of
	// [embedded storage buckets](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding)
	// to use for retrieval-augmented generation.
	BucketIDs []string `json:"bucket_ids,omitzero,required"`
	// The maximum number of results to retrieve as context for the language model.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	paramObj
}

func (r InferenceEmbeddingBucketIDsParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingBucketIDsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingBucketIDsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingTransferToolParamsResp struct {
	// Number or SIP URI placing the call.
	From string `json:"from,required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to.
	Targets []InferenceEmbeddingTransferToolParamsTargetResp `json:"targets,required"`
	// Custom headers to be added to the SIP INVITE for the transfer command.
	CustomHeaders []InferenceEmbeddingTransferToolParamsCustomHeaderResp `json:"custom_headers"`
	// Natural language instructions for your agent for how to provide context for the
	// transfer recipient.
	WarmTransferInstructions string `json:"warm_transfer_instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From                     respjson.Field
		Targets                  respjson.Field
		CustomHeaders            respjson.Field
		WarmTransferInstructions respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceEmbeddingTransferToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingTransferToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InferenceEmbeddingTransferToolParamsResp to a
// InferenceEmbeddingTransferToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InferenceEmbeddingTransferToolParams.Overrides()
func (r InferenceEmbeddingTransferToolParamsResp) ToParam() InferenceEmbeddingTransferToolParams {
	return param.Override[InferenceEmbeddingTransferToolParams](json.RawMessage(r.RawJSON()))
}

type InferenceEmbeddingTransferToolParamsTargetResp struct {
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
func (r InferenceEmbeddingTransferToolParamsTargetResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingTransferToolParamsTargetResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingTransferToolParamsCustomHeaderResp struct {
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
func (r InferenceEmbeddingTransferToolParamsCustomHeaderResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingTransferToolParamsCustomHeaderResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, Targets are required.
type InferenceEmbeddingTransferToolParams struct {
	// Number or SIP URI placing the call.
	From string `json:"from,required"`
	// The different possible targets of the transfer. The assistant will be able to
	// choose one of the targets to transfer the call to.
	Targets []InferenceEmbeddingTransferToolParamsTarget `json:"targets,omitzero,required"`
	// Natural language instructions for your agent for how to provide context for the
	// transfer recipient.
	WarmTransferInstructions param.Opt[string] `json:"warm_transfer_instructions,omitzero"`
	// Custom headers to be added to the SIP INVITE for the transfer command.
	CustomHeaders []InferenceEmbeddingTransferToolParamsCustomHeader `json:"custom_headers,omitzero"`
	paramObj
}

func (r InferenceEmbeddingTransferToolParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingTransferToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingTransferToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingTransferToolParamsTarget struct {
	// The name of the target.
	Name param.Opt[string] `json:"name,omitzero"`
	// The destination number or SIP URI of the call.
	To param.Opt[string] `json:"to,omitzero"`
	paramObj
}

func (r InferenceEmbeddingTransferToolParamsTarget) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingTransferToolParamsTarget
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingTransferToolParamsTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingTransferToolParamsCustomHeader struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// The value of the header. Note that we support mustache templating for the value.
	// For example you can use
	// `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of
	// the integration secret.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r InferenceEmbeddingTransferToolParamsCustomHeader) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingTransferToolParamsCustomHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingTransferToolParamsCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingWebhookToolParamsResp struct {
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
	BodyParameters InferenceEmbeddingWebhookToolParamsBodyParametersResp `json:"body_parameters"`
	// The headers to be sent to the external tool.
	Headers []InferenceEmbeddingWebhookToolParamsHeaderResp `json:"headers"`
	// The HTTP method to be used when calling the external tool.
	//
	// Any of "GET", "POST", "PUT", "DELETE", "PATCH".
	Method InferenceEmbeddingWebhookToolParamsMethod `json:"method"`
	// The path parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the path of the request if the
	// URL contains a placeholder for a value. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	PathParameters InferenceEmbeddingWebhookToolParamsPathParametersResp `json:"path_parameters"`
	// The query parameters the webhook tool accepts, described as a JSON Schema
	// object. These parameters will be passed to the webhook as the query of the
	// request. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	QueryParameters InferenceEmbeddingWebhookToolParamsQueryParametersResp `json:"query_parameters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description     respjson.Field
		Name            respjson.Field
		URL             respjson.Field
		BodyParameters  respjson.Field
		Headers         respjson.Field
		Method          respjson.Field
		PathParameters  respjson.Field
		QueryParameters respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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

// The body parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the body of the request. See
// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
// for documentation about the format
type InferenceEmbeddingWebhookToolParamsBodyParametersResp struct {
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
func (r InferenceEmbeddingWebhookToolParamsBodyParametersResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsBodyParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingWebhookToolParamsHeaderResp struct {
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
func (r InferenceEmbeddingWebhookToolParamsHeaderResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsHeaderResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The HTTP method to be used when calling the external tool.
type InferenceEmbeddingWebhookToolParamsMethod string

const (
	InferenceEmbeddingWebhookToolParamsMethodGet    InferenceEmbeddingWebhookToolParamsMethod = "GET"
	InferenceEmbeddingWebhookToolParamsMethodPost   InferenceEmbeddingWebhookToolParamsMethod = "POST"
	InferenceEmbeddingWebhookToolParamsMethodPut    InferenceEmbeddingWebhookToolParamsMethod = "PUT"
	InferenceEmbeddingWebhookToolParamsMethodDelete InferenceEmbeddingWebhookToolParamsMethod = "DELETE"
	InferenceEmbeddingWebhookToolParamsMethodPatch  InferenceEmbeddingWebhookToolParamsMethod = "PATCH"
)

// The path parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the path of the request if the
// URL contains a placeholder for a value. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsPathParametersResp struct {
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
func (r InferenceEmbeddingWebhookToolParamsPathParametersResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsPathParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The query parameters the webhook tool accepts, described as a JSON Schema
// object. These parameters will be passed to the webhook as the query of the
// request. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsQueryParametersResp struct {
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
func (r InferenceEmbeddingWebhookToolParamsQueryParametersResp) RawJSON() string { return r.JSON.raw }
func (r *InferenceEmbeddingWebhookToolParamsQueryParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Name, URL are required.
type InferenceEmbeddingWebhookToolParams struct {
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
	BodyParameters InferenceEmbeddingWebhookToolParamsBodyParameters `json:"body_parameters,omitzero"`
	// The headers to be sent to the external tool.
	Headers []InferenceEmbeddingWebhookToolParamsHeader `json:"headers,omitzero"`
	// The HTTP method to be used when calling the external tool.
	//
	// Any of "GET", "POST", "PUT", "DELETE", "PATCH".
	Method InferenceEmbeddingWebhookToolParamsMethod `json:"method,omitzero"`
	// The path parameters the webhook tool accepts, described as a JSON Schema object.
	// These parameters will be passed to the webhook as the path of the request if the
	// URL contains a placeholder for a value. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	PathParameters InferenceEmbeddingWebhookToolParamsPathParameters `json:"path_parameters,omitzero"`
	// The query parameters the webhook tool accepts, described as a JSON Schema
	// object. These parameters will be passed to the webhook as the query of the
	// request. See the
	// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
	// documentation about the format
	QueryParameters InferenceEmbeddingWebhookToolParamsQueryParameters `json:"query_parameters,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The body parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the body of the request. See
// the [JSON Schema reference](https://json-schema.org/understanding-json-schema)
// for documentation about the format
type InferenceEmbeddingWebhookToolParamsBodyParameters struct {
	// The properties of the body parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the body parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsBodyParameters) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsBodyParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsBodyParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsBodyParameters](
		"type", "object",
	)
}

type InferenceEmbeddingWebhookToolParamsHeader struct {
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

func (r InferenceEmbeddingWebhookToolParamsHeader) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The path parameters the webhook tool accepts, described as a JSON Schema object.
// These parameters will be passed to the webhook as the path of the request if the
// URL contains a placeholder for a value. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsPathParameters struct {
	// The properties of the path parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the path parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsPathParameters) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsPathParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsPathParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsPathParameters](
		"type", "object",
	)
}

// The query parameters the webhook tool accepts, described as a JSON Schema
// object. These parameters will be passed to the webhook as the query of the
// request. See the
// [JSON Schema reference](https://json-schema.org/understanding-json-schema) for
// documentation about the format
type InferenceEmbeddingWebhookToolParamsQueryParameters struct {
	// The properties of the query parameters.
	Properties map[string]any `json:"properties,omitzero"`
	// The required properties of the query parameters.
	Required []string `json:"required,omitzero"`
	// Any of "object".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InferenceEmbeddingWebhookToolParamsQueryParameters) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingWebhookToolParamsQueryParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingWebhookToolParamsQueryParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceEmbeddingWebhookToolParamsQueryParameters](
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
	// Default Messaging Profile used for messaging exchanges with your assistant. This
	// will be created automatically on assistant creation.
	DefaultMessagingProfileID string `json:"default_messaging_profile_id"`
	// The URL where webhooks related to delivery statused for assistant messages will
	// be sent.
	DeliveryStatusWebhookURL string `json:"delivery_status_webhook_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultMessagingProfileID respjson.Field
		DeliveryStatusWebhookURL  respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
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
	Retrieval InferenceEmbeddingBucketIDs `json:"retrieval,required"`
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
	Retrieval InferenceEmbeddingBucketIDsParam `json:"retrieval,omitzero,required"`
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
	// When enabled, allows users to interact with your AI assistant directly from your
	// website without requiring authentication. This is required for FE widgets that
	// work with assistants that have telephony enabled.
	SupportsUnauthenticatedWebCalls bool `json:"supports_unauthenticated_web_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultTexmlAppID               respjson.Field
		SupportsUnauthenticatedWebCalls respjson.Field
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

type TelephonySettingsParam struct {
	// Default Texml App used for voice calls with your assistant. This will be created
	// automatically on assistant creation.
	DefaultTexmlAppID param.Opt[string] `json:"default_texml_app_id,omitzero"`
	// When enabled, allows users to interact with your AI assistant directly from your
	// website without requiring authentication. This is required for FE widgets that
	// work with assistants that have telephony enabled.
	SupportsUnauthenticatedWebCalls param.Opt[bool] `json:"supports_unauthenticated_web_calls,omitzero"`
	paramObj
}

func (r TelephonySettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TelephonySettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelephonySettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
		EotThreshold respjson.Field
		EotTimeoutMs respjson.Field
		Numerals     respjson.Field
		SmartFormat  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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

type TransferTool struct {
	Transfer InferenceEmbeddingTransferToolParamsResp `json:"transfer,required"`
	// Any of "transfer".
	Type TransferToolType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transfer    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferTool) RawJSON() string { return r.JSON.raw }
func (r *TransferTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TransferTool to a TransferToolParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TransferToolParam.Overrides()
func (r TransferTool) ToParam() TransferToolParam {
	return param.Override[TransferToolParam](json.RawMessage(r.RawJSON()))
}

type TransferToolType string

const (
	TransferToolTypeTransfer TransferToolType = "transfer"
)

// The properties Transfer, Type are required.
type TransferToolParam struct {
	Transfer InferenceEmbeddingTransferToolParams `json:"transfer,omitzero,required"`
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

type VoiceSettings struct {
	// The voice to be used by the voice assistant. Check the full list of
	// [available voices](https://developers.telnyx.com/api/call-control/list-text-to-speech-voices)
	// via our voices API. To use ElevenLabs, you must reference your ElevenLabs API
	// key as an integration secret under the `api_key_ref` field. See
	// [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g.
	// Telnyx.KokoroTTS.af_heart)
	Voice string `json:"voice,required"`
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef string `json:"api_key_ref"`
	// Optional background audio to play on the call. Use a predefined media bed, or
	// supply a looped MP3 URL. If a media URL is chosen in the portal, customers can
	// preview it before saving.
	BackgroundAudio VoiceSettingsBackgroundAudioUnion `json:"background_audio"`
	// The speed of the voice in the range [0.25, 2.0]. 1.0 is deafult speed. Larger
	// numbers make the voice faster, smaller numbers make it slower. This is only
	// applicable for Telnyx Natural voices.
	VoiceSpeed float64 `json:"voice_speed"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voice           respjson.Field
		APIKeyRef       respjson.Field
		BackgroundAudio respjson.Field
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

// The property Voice is required.
type VoiceSettingsParam struct {
	// The voice to be used by the voice assistant. Check the full list of
	// [available voices](https://developers.telnyx.com/api/call-control/list-text-to-speech-voices)
	// via our voices API. To use ElevenLabs, you must reference your ElevenLabs API
	// key as an integration secret under the `api_key_ref` field. See
	// [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// for details. For Telnyx voices, use `Telnyx.<model_id>.<voice_id>` (e.g.
	// Telnyx.KokoroTTS.af_heart)
	Voice string `json:"voice,required"`
	// The `identifier` for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work
	// with this integration.
	APIKeyRef param.Opt[string] `json:"api_key_ref,omitzero"`
	// The speed of the voice in the range [0.25, 2.0]. 1.0 is deafult speed. Larger
	// numbers make the voice faster, smaller numbers make it slower. This is only
	// applicable for Telnyx Natural voices.
	VoiceSpeed param.Opt[float64] `json:"voice_speed,omitzero"`
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

type WebhookTool struct {
	// Any of "webhook".
	Type    WebhookToolType                         `json:"type,required"`
	Webhook InferenceEmbeddingWebhookToolParamsResp `json:"webhook,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Webhook     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookTool) RawJSON() string { return r.JSON.raw }
func (r *WebhookTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebhookTool to a WebhookToolParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebhookToolParam.Overrides()
func (r WebhookTool) ToParam() WebhookToolParam {
	return param.Override[WebhookToolParam](json.RawMessage(r.RawJSON()))
}

type WebhookToolType string

const (
	WebhookToolTypeWebhook WebhookToolType = "webhook"
)

// The properties Type, Webhook are required.
type WebhookToolParam struct {
	// Any of "webhook".
	Type    WebhookToolType                     `json:"type,omitzero,required"`
	Webhook InferenceEmbeddingWebhookToolParams `json:"webhook,omitzero,required"`
	paramObj
}

func (r WebhookToolParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolParam) UnmarshalJSON(data []byte) error {
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
	// [Get models API](https://developers.telnyx.com/api/inference/inference-embedding/get-models-public-models-get)
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
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// This is only needed when using third-party inference providers. The `identifier`
	// for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
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
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Greeting param.Opt[string] `json:"greeting,omitzero"`
	// System instructions for the assistant. These may be templated with
	// [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables)
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	// This is only needed when using third-party inference providers. The `identifier`
	// for an integration secret
	// [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret)
	// that refers to your LLM provider's API key. Warning: Free plans are unlikely to
	// work with this integration.
	LlmAPIKeyRef param.Opt[string] `json:"llm_api_key_ref,omitzero"`
	// ID of the model to use. You can use the
	// [Get models API](https://developers.telnyx.com/api/inference/inference-embedding/get-models-public-models-get)
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
	Text                     string                                                       `json:"text,required"`
	To                       string                                                       `json:"to,required"`
	ShouldCreateConversation param.Opt[bool]                                              `json:"should_create_conversation,omitzero"`
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
