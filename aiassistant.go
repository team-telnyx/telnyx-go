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
func (r *AIAssistantService) New(ctx context.Context, body AIAssistantNewParams, opts ...option.RequestOption) (res *AIAssistantNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an AI Assistant configuration by `assistant_id`.
func (r *AIAssistantService) Get(ctx context.Context, assistantID string, query AIAssistantGetParams, opts ...option.RequestOption) (res *AIAssistantGetResponse, err error) {
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
func (r *AIAssistantService) Update(ctx context.Context, assistantID string, body AIAssistantUpdateParams, opts ...option.RequestOption) (res *AIAssistantUpdateResponse, err error) {
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
func (r *AIAssistantService) Clone(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *AIAssistantCloneResponse, err error) {
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
func (r *AIAssistantService) Import(ctx context.Context, body AIAssistantImportParams, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/assistants/import"
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
	OfBookAppointmentTool   *AssistantToolBookAppointmentToolParam   `json:",omitzero,inline"`
	OfCheckAvailabilityTool *AssistantToolCheckAvailabilityToolParam `json:",omitzero,inline"`
	OfWebhookTool           *WebhookToolParam                        `json:",omitzero,inline"`
	OfHangupTool            *HangupToolParam                         `json:",omitzero,inline"`
	OfTransferTool          *TransferToolParam                       `json:",omitzero,inline"`
	OfRetrievalTool         *RetrievalToolParam                      `json:",omitzero,inline"`
	paramUnion
}

func (u AssistantToolUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBookAppointmentTool,
		u.OfCheckAvailabilityTool,
		u.OfWebhookTool,
		u.OfHangupTool,
		u.OfTransferTool,
		u.OfRetrievalTool)
}
func (u *AssistantToolUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AssistantToolUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBookAppointmentTool) {
		return u.OfBookAppointmentTool
	} else if !param.IsOmitted(u.OfCheckAvailabilityTool) {
		return u.OfCheckAvailabilityTool
	} else if !param.IsOmitted(u.OfWebhookTool) {
		return u.OfWebhookTool
	} else if !param.IsOmitted(u.OfHangupTool) {
		return u.OfHangupTool
	} else if !param.IsOmitted(u.OfTransferTool) {
		return u.OfTransferTool
	} else if !param.IsOmitted(u.OfRetrievalTool) {
		return u.OfRetrievalTool
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetBookAppointment() *AssistantToolBookAppointmentToolBookAppointmentParam {
	if vt := u.OfBookAppointmentTool; vt != nil {
		return &vt.BookAppointment
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetCheckAvailability() *AssistantToolCheckAvailabilityToolCheckAvailabilityParam {
	if vt := u.OfCheckAvailabilityTool; vt != nil {
		return &vt.CheckAvailability
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetWebhook() *InferenceEmbeddingWebhookToolParams {
	if vt := u.OfWebhookTool; vt != nil {
		return &vt.Webhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetHangup() *HangupToolParams {
	if vt := u.OfHangupTool; vt != nil {
		return &vt.Hangup
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetTransfer() *InferenceEmbeddingTransferToolParams {
	if vt := u.OfTransferTool; vt != nil {
		return &vt.Transfer
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetRetrieval() *InferenceEmbeddingBucketIDsParam {
	if vt := u.OfRetrievalTool; vt != nil {
		return &vt.Retrieval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AssistantToolUnionParam) GetType() *string {
	if vt := u.OfBookAppointmentTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCheckAvailabilityTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebhookTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfHangupTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTransferTool; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRetrievalTool; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties BookAppointment, Type are required.
type AssistantToolBookAppointmentToolParam struct {
	BookAppointment AssistantToolBookAppointmentToolBookAppointmentParam `json:"book_appointment,omitzero,required"`
	// Any of "book_appointment".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r AssistantToolBookAppointmentToolParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolBookAppointmentToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolBookAppointmentToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolBookAppointmentToolParam](
		"type", "book_appointment",
	)
}

// The properties APIKeyRef, EventTypeID are required.
type AssistantToolBookAppointmentToolBookAppointmentParam struct {
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

func (r AssistantToolBookAppointmentToolBookAppointmentParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolBookAppointmentToolBookAppointmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolBookAppointmentToolBookAppointmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CheckAvailability, Type are required.
type AssistantToolCheckAvailabilityToolParam struct {
	CheckAvailability AssistantToolCheckAvailabilityToolCheckAvailabilityParam `json:"check_availability,omitzero,required"`
	// Any of "check_availability".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r AssistantToolCheckAvailabilityToolParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolCheckAvailabilityToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolCheckAvailabilityToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AssistantToolCheckAvailabilityToolParam](
		"type", "check_availability",
	)
}

// The properties APIKeyRef, EventTypeID are required.
type AssistantToolCheckAvailabilityToolCheckAvailabilityParam struct {
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

func (r AssistantToolCheckAvailabilityToolCheckAvailabilityParam) MarshalJSON() (data []byte, err error) {
	type shadow AssistantToolCheckAvailabilityToolCheckAvailabilityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AssistantToolCheckAvailabilityToolCheckAvailabilityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AssistantToolUnion contains all possible properties and values from
// [WebhookTool], [RetrievalTool], [AssistantToolHandoffTool], [HangupTool],
// [TransferTool], [AssistantToolSipReferTool], [AssistantToolDtmfTool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssistantToolUnion struct {
	Type string `json:"type"`
	// This field is from variant [WebhookTool].
	Webhook InferenceEmbeddingWebhookToolParamsResp `json:"webhook"`
	// This field is from variant [RetrievalTool].
	Retrieval InferenceEmbeddingBucketIDs `json:"retrieval"`
	// This field is from variant [AssistantToolHandoffTool].
	Handoff AssistantToolHandoffToolHandoff `json:"handoff"`
	// This field is from variant [HangupTool].
	Hangup HangupToolParamsResp `json:"hangup"`
	// This field is from variant [TransferTool].
	Transfer InferenceEmbeddingTransferToolParamsResp `json:"transfer"`
	// This field is from variant [AssistantToolSipReferTool].
	Refer AssistantToolSipReferToolRefer `json:"refer"`
	// This field is from variant [AssistantToolDtmfTool].
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

func (u AssistantToolUnion) AsWebhookTool() (v WebhookTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsRetrievalTool() (v RetrievalTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsHandoffTool() (v AssistantToolHandoffTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsHangupTool() (v HangupTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsTransferTool() (v TransferTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsSipReferTool() (v AssistantToolSipReferTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssistantToolUnion) AsDtmfTool() (v AssistantToolDtmfTool) {
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

// The handoff tool allows the assistant to hand off control of the conversation to
// another AI assistant. By default, this will happen transparently to the end
// user.
type AssistantToolHandoffTool struct {
	Handoff AssistantToolHandoffToolHandoff `json:"handoff,required"`
	// Any of "handoff".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Handoff     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolHandoffTool) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolHandoffTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolHandoffToolHandoff struct {
	// List of possible assistants that can receive a handoff.
	AIAssistants []AssistantToolHandoffToolHandoffAIAssistant `json:"ai_assistants,required"`
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
func (r AssistantToolHandoffToolHandoff) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolHandoffToolHandoff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolHandoffToolHandoffAIAssistant struct {
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
func (r AssistantToolHandoffToolHandoffAIAssistant) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolHandoffToolHandoffAIAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSipReferTool struct {
	Refer AssistantToolSipReferToolRefer `json:"refer,required"`
	// Any of "refer".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Refer       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolSipReferTool) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSipReferTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSipReferToolRefer struct {
	// The different possible targets of the SIP refer. The assistant will be able to
	// choose one of the targets to refer the call to.
	Targets []AssistantToolSipReferToolReferTarget `json:"targets,required"`
	// Custom headers to be added to the SIP REFER.
	CustomHeaders []AssistantToolSipReferToolReferCustomHeader `json:"custom_headers"`
	// SIP headers to be added to the SIP REFER. Currently only User-to-User and
	// Diversion headers are supported.
	SipHeaders []AssistantToolSipReferToolReferSipHeader `json:"sip_headers"`
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
func (r AssistantToolSipReferToolRefer) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSipReferToolRefer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSipReferToolReferTarget struct {
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
func (r AssistantToolSipReferToolReferTarget) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSipReferToolReferTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSipReferToolReferCustomHeader struct {
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
func (r AssistantToolSipReferToolReferCustomHeader) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSipReferToolReferCustomHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolSipReferToolReferSipHeader struct {
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
func (r AssistantToolSipReferToolReferSipHeader) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolSipReferToolReferSipHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolDtmfTool struct {
	SendDtmf map[string]any `json:"send_dtmf,required"`
	// Any of "send_dtmf".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SendDtmf    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolDtmfTool) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolDtmfTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantsList struct {
	Data []AssistantsListData `json:"data,required"`
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

type AssistantsListData struct {
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
	Tools         []AssistantToolUnion  `json:"tools"`
	Transcription TranscriptionSettings `json:"transcription"`
	VoiceSettings VoiceSettings         `json:"voice_settings"`
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
func (r AssistantsListData) RawJSON() string { return r.JSON.raw }
func (r *AssistantsListData) UnmarshalJSON(data []byte) error {
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
	// Any of "elevenlabs", "vapi".
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
)

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
	// The language of the audio to be transcribed. This is only applicable for
	// `openai/whisper-large-v3-turbo` model. If not set, of if set to `auto`, the
	// model will automatically detect the language. For the full list of supported
	// languages, see the
	// [whisper tokenizer](https://github.com/openai/whisper/blob/main/whisper/tokenizer.py).
	Language string `json:"language"`
	// The speech to text model to be used by the voice assistant.
	//
	//   - `distil-whisper/distil-large-v2` is lower latency but English-only.
	//   - `openai/whisper-large-v3-turbo` is multi-lingual with automatic language
	//     detection but slightly higher latency.
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Model       respjson.Field
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

type TranscriptionSettingsParam struct {
	// The language of the audio to be transcribed. This is only applicable for
	// `openai/whisper-large-v3-turbo` model. If not set, of if set to `auto`, the
	// model will automatically detect the language. For the full list of supported
	// languages, see the
	// [whisper tokenizer](https://github.com/openai/whisper/blob/main/whisper/tokenizer.py).
	Language param.Opt[string] `json:"language,omitzero"`
	// The speech to text model to be used by the voice assistant.
	//
	//   - `distil-whisper/distil-large-v2` is lower latency but English-only.
	//   - `openai/whisper-large-v3-turbo` is multi-lingual with automatic language
	//     detection but slightly higher latency.
	Model param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r TranscriptionSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow TranscriptionSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranscriptionSettingsParam) UnmarshalJSON(data []byte) error {
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
// from [VoiceSettingsBackgroundAudioObject], [VoiceSettingsBackgroundAudioObject],
// [VoiceSettingsBackgroundAudioObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type VoiceSettingsBackgroundAudioUnion struct {
	// This field is from variant [VoiceSettingsBackgroundAudioObject].
	Type string `json:"type"`
	// This field is from variant [VoiceSettingsBackgroundAudioObject].
	Value string `json:"value"`
	JSON  struct {
		Type  respjson.Field
		Value respjson.Field
		raw   string
	} `json:"-"`
}

func (u VoiceSettingsBackgroundAudioUnion) AsVoiceSettingsBackgroundAudioObject() (v VoiceSettingsBackgroundAudioObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSettingsBackgroundAudioUnion) AsVariant2() (v VoiceSettingsBackgroundAudioObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VoiceSettingsBackgroundAudioUnion) AsVariant3() (v VoiceSettingsBackgroundAudioObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VoiceSettingsBackgroundAudioUnion) RawJSON() string { return u.JSON.raw }

func (r *VoiceSettingsBackgroundAudioUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSettingsBackgroundAudioObject struct {
	// Select from predefined media options.
	//
	// Any of "predefined_media".
	Type string `json:"type,required"`
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
func (r VoiceSettingsBackgroundAudioObject) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettingsBackgroundAudioObject) UnmarshalJSON(data []byte) error {
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
	OfVoiceSettingsBackgroundAudioObject *VoiceSettingsBackgroundAudioObjectParam `json:",omitzero,inline"`
	OfVariant2                           *VoiceSettingsBackgroundAudioObjectParam `json:",omitzero,inline"`
	OfVariant3                           *VoiceSettingsBackgroundAudioObjectParam `json:",omitzero,inline"`
	paramUnion
}

func (u VoiceSettingsBackgroundAudioUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfVoiceSettingsBackgroundAudioObject, u.OfVariant2, u.OfVariant3)
}
func (u *VoiceSettingsBackgroundAudioUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VoiceSettingsBackgroundAudioUnionParam) asAny() any {
	if !param.IsOmitted(u.OfVoiceSettingsBackgroundAudioObject) {
		return u.OfVoiceSettingsBackgroundAudioObject
	} else if !param.IsOmitted(u.OfVariant2) {
		return u.OfVariant2
	} else if !param.IsOmitted(u.OfVariant3) {
		return u.OfVariant3
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VoiceSettingsBackgroundAudioUnionParam) GetType() *string {
	if vt := u.OfVoiceSettingsBackgroundAudioObject; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant2; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u VoiceSettingsBackgroundAudioUnionParam) GetValue() *string {
	if vt := u.OfVoiceSettingsBackgroundAudioObject; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfVariant2; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfVariant3; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// The properties Type, Value are required.
type VoiceSettingsBackgroundAudioObjectParam struct {
	// Select from predefined media options.
	//
	// Any of "predefined_media".
	Type string `json:"type,omitzero,required"`
	// The predefined media to use. `silence` disables background audio.
	//
	// Any of "silence", "office".
	Value string `json:"value,omitzero,required"`
	paramObj
}

func (r VoiceSettingsBackgroundAudioObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceSettingsBackgroundAudioObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceSettingsBackgroundAudioObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VoiceSettingsBackgroundAudioObjectParam](
		"type", "predefined_media",
	)
	apijson.RegisterFieldValidator[VoiceSettingsBackgroundAudioObjectParam](
		"value", "silence", "office",
	)
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

type AIAssistantNewResponse struct {
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
	Tools         []AssistantToolUnion  `json:"tools"`
	Transcription TranscriptionSettings `json:"transcription"`
	VoiceSettings VoiceSettings         `json:"voice_settings"`
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
func (r AIAssistantNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantGetResponse struct {
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
	Tools         []AssistantToolUnion  `json:"tools"`
	Transcription TranscriptionSettings `json:"transcription"`
	VoiceSettings VoiceSettings         `json:"voice_settings"`
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
func (r AIAssistantGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantUpdateResponse = any

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

type AIAssistantCloneResponse struct {
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
	Tools         []AssistantToolUnion  `json:"tools"`
	Transcription TranscriptionSettings `json:"transcription"`
	VoiceSettings VoiceSettings         `json:"voice_settings"`
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
func (r AIAssistantCloneResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantCloneResponse) UnmarshalJSON(data []byte) error {
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
	Tools         []AssistantToolUnionParam  `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam         `json:"voice_settings,omitzero"`
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
	Tools         []AssistantToolUnionParam  `json:"tools,omitzero"`
	Transcription TranscriptionSettingsParam `json:"transcription,omitzero"`
	VoiceSettings VoiceSettingsParam         `json:"voice_settings,omitzero"`
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

type AIAssistantImportParams struct {
	// Integration secret pointer that refers to the API key for the external provider.
	// This should be an identifier for an integration secret created via
	// /v2/integration_secrets.
	APIKeyRef string `json:"api_key_ref,required"`
	// The external provider to import assistants from.
	//
	// Any of "elevenlabs", "vapi".
	Provider AIAssistantImportParamsProvider `json:"provider,omitzero,required"`
	paramObj
}

func (r AIAssistantImportParams) MarshalJSON() (data []byte, err error) {
	type shadow AIAssistantImportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIAssistantImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The external provider to import assistants from.
type AIAssistantImportParamsProvider string

const (
	AIAssistantImportParamsProviderElevenlabs AIAssistantImportParamsProvider = "elevenlabs"
	AIAssistantImportParamsProviderVapi       AIAssistantImportParamsProvider = "vapi"
)
