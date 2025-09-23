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

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

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
func (r *AIAssistantVersionService) Get(ctx context.Context, versionID string, params AIAssistantVersionGetParams, opts ...option.RequestOption) (res *AIAssistantVersionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s", params.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Updates the configuration of a specific assistant version. Can not update main
// version
func (r *AIAssistantVersionService) Update(ctx context.Context, versionID string, params AIAssistantVersionUpdateParams, opts ...option.RequestOption) (res *AIAssistantVersionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s", params.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves all versions of a specific assistant with complete configuration and
// metadata
func (r *AIAssistantVersionService) List(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *AssistantsList, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/versions", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently removes a specific version of an assistant. Can not delete main
// version
func (r *AIAssistantVersionService) Delete(ctx context.Context, versionID string, body AIAssistantVersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s", body.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Promotes a specific version to be the main/current version of the assistant.
// This will delete any existing canary deploy configuration and send all live
// production traffic to this version.
func (r *AIAssistantVersionService) Promote(ctx context.Context, versionID string, body AIAssistantVersionPromoteParams, opts ...option.RequestOption) (res *AIAssistantVersionPromoteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AssistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("ai/assistants/%s/versions/%s/promote", body.AssistantID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
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

func (r UpdateAssistantParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAssistantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAssistantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantVersionGetResponse struct {
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
func (r AIAssistantVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantVersionUpdateResponse struct {
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
func (r AIAssistantVersionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantVersionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantVersionPromoteResponse struct {
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
func (r AIAssistantVersionPromoteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAssistantVersionPromoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantVersionGetParams struct {
	AssistantID       string          `path:"assistant_id,required" json:"-"`
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
	AssistantID     string `path:"assistant_id,required" json:"-"`
	UpdateAssistant UpdateAssistantParam
	paramObj
}

func (r AIAssistantVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAssistant)
}
func (r *AIAssistantVersionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UpdateAssistant)
}

type AIAssistantVersionDeleteParams struct {
	AssistantID string `path:"assistant_id,required" json:"-"`
	paramObj
}

type AIAssistantVersionPromoteParams struct {
	AssistantID string `path:"assistant_id,required" json:"-"`
	paramObj
}
