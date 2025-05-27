# CreateAssistantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Model** | **string** | ID of the model to use. You can use the [Get models API](https://developers.telnyx.com/api/inference/inference-embedding/get-models-public-models-get) to see all of your available models, | 
**Instructions** | **string** | System instructions for the assistant. These may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) | 
**Tools** | Pointer to [**[]AssistantToolsInner**](AssistantToolsInner.md) | The tools that the assistant can use. These may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Greeting** | Pointer to **string** | Text that the assistant will use to start the conversation. This may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) | [optional] 
**LlmApiKeyRef** | Pointer to **string** | This is only needed when using third-party inference providers. The &#x60;identifier&#x60; for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your LLM provider&#39;s API key. Warning: Free plans are unlikely to work with this integration. | [optional] 
**VoiceSettings** | Pointer to [**VoiceSettings**](VoiceSettings.md) |  | [optional] 
**Transcription** | Pointer to [**TranscriptionSettings**](TranscriptionSettings.md) |  | [optional] 
**TelephonySettings** | Pointer to [**TelephonySettings**](TelephonySettings.md) |  | [optional] 
**MessagingSettings** | Pointer to [**MessagingSettings**](MessagingSettings.md) |  | [optional] 
**EnabledFeatures** | Pointer to [**[]EnabledFeatures**](EnabledFeatures.md) |  | [optional] 
**InsightSettings** | Pointer to [**InsightSettings**](InsightSettings.md) |  | [optional] 
**PrivacySettings** | Pointer to [**PrivacySettings**](PrivacySettings.md) |  | [optional] 
**DynamicVariablesWebhookUrl** | Pointer to **string** | If the dynamic_variables_webhook_url is set for the assistant, we will send a request at the start of the conversation. See our [guide](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) for more information. | [optional] 
**DynamicVariables** | Pointer to **map[string]interface{}** | Map of dynamic variables and their default values | [optional] 

## Methods

### NewCreateAssistantRequest

`func NewCreateAssistantRequest(name string, model string, instructions string, ) *CreateAssistantRequest`

NewCreateAssistantRequest instantiates a new CreateAssistantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAssistantRequestWithDefaults

`func NewCreateAssistantRequestWithDefaults() *CreateAssistantRequest`

NewCreateAssistantRequestWithDefaults instantiates a new CreateAssistantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAssistantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAssistantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAssistantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModel

`func (o *CreateAssistantRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateAssistantRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateAssistantRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInstructions

`func (o *CreateAssistantRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateAssistantRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateAssistantRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetTools

`func (o *CreateAssistantRequest) GetTools() []AssistantToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CreateAssistantRequest) GetToolsOk() (*[]AssistantToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CreateAssistantRequest) SetTools(v []AssistantToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CreateAssistantRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAssistantRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAssistantRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAssistantRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAssistantRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGreeting

`func (o *CreateAssistantRequest) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *CreateAssistantRequest) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *CreateAssistantRequest) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.

### HasGreeting

`func (o *CreateAssistantRequest) HasGreeting() bool`

HasGreeting returns a boolean if a field has been set.

### GetLlmApiKeyRef

`func (o *CreateAssistantRequest) GetLlmApiKeyRef() string`

GetLlmApiKeyRef returns the LlmApiKeyRef field if non-nil, zero value otherwise.

### GetLlmApiKeyRefOk

`func (o *CreateAssistantRequest) GetLlmApiKeyRefOk() (*string, bool)`

GetLlmApiKeyRefOk returns a tuple with the LlmApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmApiKeyRef

`func (o *CreateAssistantRequest) SetLlmApiKeyRef(v string)`

SetLlmApiKeyRef sets LlmApiKeyRef field to given value.

### HasLlmApiKeyRef

`func (o *CreateAssistantRequest) HasLlmApiKeyRef() bool`

HasLlmApiKeyRef returns a boolean if a field has been set.

### GetVoiceSettings

`func (o *CreateAssistantRequest) GetVoiceSettings() VoiceSettings`

GetVoiceSettings returns the VoiceSettings field if non-nil, zero value otherwise.

### GetVoiceSettingsOk

`func (o *CreateAssistantRequest) GetVoiceSettingsOk() (*VoiceSettings, bool)`

GetVoiceSettingsOk returns a tuple with the VoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSettings

`func (o *CreateAssistantRequest) SetVoiceSettings(v VoiceSettings)`

SetVoiceSettings sets VoiceSettings field to given value.

### HasVoiceSettings

`func (o *CreateAssistantRequest) HasVoiceSettings() bool`

HasVoiceSettings returns a boolean if a field has been set.

### GetTranscription

`func (o *CreateAssistantRequest) GetTranscription() TranscriptionSettings`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *CreateAssistantRequest) GetTranscriptionOk() (*TranscriptionSettings, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *CreateAssistantRequest) SetTranscription(v TranscriptionSettings)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *CreateAssistantRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetTelephonySettings

`func (o *CreateAssistantRequest) GetTelephonySettings() TelephonySettings`

GetTelephonySettings returns the TelephonySettings field if non-nil, zero value otherwise.

### GetTelephonySettingsOk

`func (o *CreateAssistantRequest) GetTelephonySettingsOk() (*TelephonySettings, bool)`

GetTelephonySettingsOk returns a tuple with the TelephonySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephonySettings

`func (o *CreateAssistantRequest) SetTelephonySettings(v TelephonySettings)`

SetTelephonySettings sets TelephonySettings field to given value.

### HasTelephonySettings

`func (o *CreateAssistantRequest) HasTelephonySettings() bool`

HasTelephonySettings returns a boolean if a field has been set.

### GetMessagingSettings

`func (o *CreateAssistantRequest) GetMessagingSettings() MessagingSettings`

GetMessagingSettings returns the MessagingSettings field if non-nil, zero value otherwise.

### GetMessagingSettingsOk

`func (o *CreateAssistantRequest) GetMessagingSettingsOk() (*MessagingSettings, bool)`

GetMessagingSettingsOk returns a tuple with the MessagingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingSettings

`func (o *CreateAssistantRequest) SetMessagingSettings(v MessagingSettings)`

SetMessagingSettings sets MessagingSettings field to given value.

### HasMessagingSettings

`func (o *CreateAssistantRequest) HasMessagingSettings() bool`

HasMessagingSettings returns a boolean if a field has been set.

### GetEnabledFeatures

`func (o *CreateAssistantRequest) GetEnabledFeatures() []EnabledFeatures`

GetEnabledFeatures returns the EnabledFeatures field if non-nil, zero value otherwise.

### GetEnabledFeaturesOk

`func (o *CreateAssistantRequest) GetEnabledFeaturesOk() (*[]EnabledFeatures, bool)`

GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatures

`func (o *CreateAssistantRequest) SetEnabledFeatures(v []EnabledFeatures)`

SetEnabledFeatures sets EnabledFeatures field to given value.

### HasEnabledFeatures

`func (o *CreateAssistantRequest) HasEnabledFeatures() bool`

HasEnabledFeatures returns a boolean if a field has been set.

### GetInsightSettings

`func (o *CreateAssistantRequest) GetInsightSettings() InsightSettings`

GetInsightSettings returns the InsightSettings field if non-nil, zero value otherwise.

### GetInsightSettingsOk

`func (o *CreateAssistantRequest) GetInsightSettingsOk() (*InsightSettings, bool)`

GetInsightSettingsOk returns a tuple with the InsightSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSettings

`func (o *CreateAssistantRequest) SetInsightSettings(v InsightSettings)`

SetInsightSettings sets InsightSettings field to given value.

### HasInsightSettings

`func (o *CreateAssistantRequest) HasInsightSettings() bool`

HasInsightSettings returns a boolean if a field has been set.

### GetPrivacySettings

`func (o *CreateAssistantRequest) GetPrivacySettings() PrivacySettings`

GetPrivacySettings returns the PrivacySettings field if non-nil, zero value otherwise.

### GetPrivacySettingsOk

`func (o *CreateAssistantRequest) GetPrivacySettingsOk() (*PrivacySettings, bool)`

GetPrivacySettingsOk returns a tuple with the PrivacySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacySettings

`func (o *CreateAssistantRequest) SetPrivacySettings(v PrivacySettings)`

SetPrivacySettings sets PrivacySettings field to given value.

### HasPrivacySettings

`func (o *CreateAssistantRequest) HasPrivacySettings() bool`

HasPrivacySettings returns a boolean if a field has been set.

### GetDynamicVariablesWebhookUrl

`func (o *CreateAssistantRequest) GetDynamicVariablesWebhookUrl() string`

GetDynamicVariablesWebhookUrl returns the DynamicVariablesWebhookUrl field if non-nil, zero value otherwise.

### GetDynamicVariablesWebhookUrlOk

`func (o *CreateAssistantRequest) GetDynamicVariablesWebhookUrlOk() (*string, bool)`

GetDynamicVariablesWebhookUrlOk returns a tuple with the DynamicVariablesWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVariablesWebhookUrl

`func (o *CreateAssistantRequest) SetDynamicVariablesWebhookUrl(v string)`

SetDynamicVariablesWebhookUrl sets DynamicVariablesWebhookUrl field to given value.

### HasDynamicVariablesWebhookUrl

`func (o *CreateAssistantRequest) HasDynamicVariablesWebhookUrl() bool`

HasDynamicVariablesWebhookUrl returns a boolean if a field has been set.

### GetDynamicVariables

`func (o *CreateAssistantRequest) GetDynamicVariables() map[string]interface{}`

GetDynamicVariables returns the DynamicVariables field if non-nil, zero value otherwise.

### GetDynamicVariablesOk

`func (o *CreateAssistantRequest) GetDynamicVariablesOk() (*map[string]interface{}, bool)`

GetDynamicVariablesOk returns a tuple with the DynamicVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVariables

`func (o *CreateAssistantRequest) SetDynamicVariables(v map[string]interface{})`

SetDynamicVariables sets DynamicVariables field to given value.

### HasDynamicVariables

`func (o *CreateAssistantRequest) HasDynamicVariables() bool`

HasDynamicVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


