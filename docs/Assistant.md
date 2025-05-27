# Assistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Model** | **string** | ID of the model to use. You can use the [Get models API](https://developers.telnyx.com/api/inference/inference-embedding/get-models-public-models-get) to see all of your available models, | 
**Instructions** | **string** | System instructions for the assistant. These may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) | 
**Tools** | Pointer to [**[]AssistantToolsInner**](AssistantToolsInner.md) | The tools that the assistant can use. These may be templated with [dynamic variables](https://developers.telnyx.com/docs/inference/ai-assistants/dynamic-variables) | [optional] 
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
**DynamicVariables** | Pointer to **map[string]interface{}** | Map of dynamic variables and their values | [optional] 

## Methods

### NewAssistant

`func NewAssistant(id string, name string, createdAt time.Time, model string, instructions string, ) *Assistant`

NewAssistant instantiates a new Assistant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantWithDefaults

`func NewAssistantWithDefaults() *Assistant`

NewAssistantWithDefaults instantiates a new Assistant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Assistant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Assistant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Assistant) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Assistant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Assistant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Assistant) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Assistant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Assistant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Assistant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Assistant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Assistant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Assistant) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Assistant) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModel

`func (o *Assistant) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Assistant) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Assistant) SetModel(v string)`

SetModel sets Model field to given value.


### GetInstructions

`func (o *Assistant) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Assistant) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Assistant) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.


### GetTools

`func (o *Assistant) GetTools() []AssistantToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Assistant) GetToolsOk() (*[]AssistantToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Assistant) SetTools(v []AssistantToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *Assistant) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetGreeting

`func (o *Assistant) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *Assistant) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *Assistant) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.

### HasGreeting

`func (o *Assistant) HasGreeting() bool`

HasGreeting returns a boolean if a field has been set.

### GetLlmApiKeyRef

`func (o *Assistant) GetLlmApiKeyRef() string`

GetLlmApiKeyRef returns the LlmApiKeyRef field if non-nil, zero value otherwise.

### GetLlmApiKeyRefOk

`func (o *Assistant) GetLlmApiKeyRefOk() (*string, bool)`

GetLlmApiKeyRefOk returns a tuple with the LlmApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmApiKeyRef

`func (o *Assistant) SetLlmApiKeyRef(v string)`

SetLlmApiKeyRef sets LlmApiKeyRef field to given value.

### HasLlmApiKeyRef

`func (o *Assistant) HasLlmApiKeyRef() bool`

HasLlmApiKeyRef returns a boolean if a field has been set.

### GetVoiceSettings

`func (o *Assistant) GetVoiceSettings() VoiceSettings`

GetVoiceSettings returns the VoiceSettings field if non-nil, zero value otherwise.

### GetVoiceSettingsOk

`func (o *Assistant) GetVoiceSettingsOk() (*VoiceSettings, bool)`

GetVoiceSettingsOk returns a tuple with the VoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSettings

`func (o *Assistant) SetVoiceSettings(v VoiceSettings)`

SetVoiceSettings sets VoiceSettings field to given value.

### HasVoiceSettings

`func (o *Assistant) HasVoiceSettings() bool`

HasVoiceSettings returns a boolean if a field has been set.

### GetTranscription

`func (o *Assistant) GetTranscription() TranscriptionSettings`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *Assistant) GetTranscriptionOk() (*TranscriptionSettings, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *Assistant) SetTranscription(v TranscriptionSettings)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *Assistant) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetTelephonySettings

`func (o *Assistant) GetTelephonySettings() TelephonySettings`

GetTelephonySettings returns the TelephonySettings field if non-nil, zero value otherwise.

### GetTelephonySettingsOk

`func (o *Assistant) GetTelephonySettingsOk() (*TelephonySettings, bool)`

GetTelephonySettingsOk returns a tuple with the TelephonySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephonySettings

`func (o *Assistant) SetTelephonySettings(v TelephonySettings)`

SetTelephonySettings sets TelephonySettings field to given value.

### HasTelephonySettings

`func (o *Assistant) HasTelephonySettings() bool`

HasTelephonySettings returns a boolean if a field has been set.

### GetMessagingSettings

`func (o *Assistant) GetMessagingSettings() MessagingSettings`

GetMessagingSettings returns the MessagingSettings field if non-nil, zero value otherwise.

### GetMessagingSettingsOk

`func (o *Assistant) GetMessagingSettingsOk() (*MessagingSettings, bool)`

GetMessagingSettingsOk returns a tuple with the MessagingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingSettings

`func (o *Assistant) SetMessagingSettings(v MessagingSettings)`

SetMessagingSettings sets MessagingSettings field to given value.

### HasMessagingSettings

`func (o *Assistant) HasMessagingSettings() bool`

HasMessagingSettings returns a boolean if a field has been set.

### GetEnabledFeatures

`func (o *Assistant) GetEnabledFeatures() []EnabledFeatures`

GetEnabledFeatures returns the EnabledFeatures field if non-nil, zero value otherwise.

### GetEnabledFeaturesOk

`func (o *Assistant) GetEnabledFeaturesOk() (*[]EnabledFeatures, bool)`

GetEnabledFeaturesOk returns a tuple with the EnabledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFeatures

`func (o *Assistant) SetEnabledFeatures(v []EnabledFeatures)`

SetEnabledFeatures sets EnabledFeatures field to given value.

### HasEnabledFeatures

`func (o *Assistant) HasEnabledFeatures() bool`

HasEnabledFeatures returns a boolean if a field has been set.

### GetInsightSettings

`func (o *Assistant) GetInsightSettings() InsightSettings`

GetInsightSettings returns the InsightSettings field if non-nil, zero value otherwise.

### GetInsightSettingsOk

`func (o *Assistant) GetInsightSettingsOk() (*InsightSettings, bool)`

GetInsightSettingsOk returns a tuple with the InsightSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSettings

`func (o *Assistant) SetInsightSettings(v InsightSettings)`

SetInsightSettings sets InsightSettings field to given value.

### HasInsightSettings

`func (o *Assistant) HasInsightSettings() bool`

HasInsightSettings returns a boolean if a field has been set.

### GetPrivacySettings

`func (o *Assistant) GetPrivacySettings() PrivacySettings`

GetPrivacySettings returns the PrivacySettings field if non-nil, zero value otherwise.

### GetPrivacySettingsOk

`func (o *Assistant) GetPrivacySettingsOk() (*PrivacySettings, bool)`

GetPrivacySettingsOk returns a tuple with the PrivacySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacySettings

`func (o *Assistant) SetPrivacySettings(v PrivacySettings)`

SetPrivacySettings sets PrivacySettings field to given value.

### HasPrivacySettings

`func (o *Assistant) HasPrivacySettings() bool`

HasPrivacySettings returns a boolean if a field has been set.

### GetDynamicVariablesWebhookUrl

`func (o *Assistant) GetDynamicVariablesWebhookUrl() string`

GetDynamicVariablesWebhookUrl returns the DynamicVariablesWebhookUrl field if non-nil, zero value otherwise.

### GetDynamicVariablesWebhookUrlOk

`func (o *Assistant) GetDynamicVariablesWebhookUrlOk() (*string, bool)`

GetDynamicVariablesWebhookUrlOk returns a tuple with the DynamicVariablesWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVariablesWebhookUrl

`func (o *Assistant) SetDynamicVariablesWebhookUrl(v string)`

SetDynamicVariablesWebhookUrl sets DynamicVariablesWebhookUrl field to given value.

### HasDynamicVariablesWebhookUrl

`func (o *Assistant) HasDynamicVariablesWebhookUrl() bool`

HasDynamicVariablesWebhookUrl returns a boolean if a field has been set.

### GetDynamicVariables

`func (o *Assistant) GetDynamicVariables() map[string]interface{}`

GetDynamicVariables returns the DynamicVariables field if non-nil, zero value otherwise.

### GetDynamicVariablesOk

`func (o *Assistant) GetDynamicVariablesOk() (*map[string]interface{}, bool)`

GetDynamicVariablesOk returns a tuple with the DynamicVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicVariables

`func (o *Assistant) SetDynamicVariables(v map[string]interface{})`

SetDynamicVariables sets DynamicVariables field to given value.

### HasDynamicVariables

`func (o *Assistant) HasDynamicVariables() bool`

HasDynamicVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


