# AIAssistantStartRequestVoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyRef** | Pointer to **string** | The &#x60;identifier&#x60; for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work with this integration. | [optional] 
**VoiceSpeed** | Pointer to **float32** | The voice speed to be used for the voice. The voice speed must be between 0.1 and 2.0. Default value is 1.0. | [optional] [default to 1.0]

## Methods

### NewAIAssistantStartRequestVoiceSettings

`func NewAIAssistantStartRequestVoiceSettings() *AIAssistantStartRequestVoiceSettings`

NewAIAssistantStartRequestVoiceSettings instantiates a new AIAssistantStartRequestVoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAssistantStartRequestVoiceSettingsWithDefaults

`func NewAIAssistantStartRequestVoiceSettingsWithDefaults() *AIAssistantStartRequestVoiceSettings`

NewAIAssistantStartRequestVoiceSettingsWithDefaults instantiates a new AIAssistantStartRequestVoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyRef

`func (o *AIAssistantStartRequestVoiceSettings) GetApiKeyRef() string`

GetApiKeyRef returns the ApiKeyRef field if non-nil, zero value otherwise.

### GetApiKeyRefOk

`func (o *AIAssistantStartRequestVoiceSettings) GetApiKeyRefOk() (*string, bool)`

GetApiKeyRefOk returns a tuple with the ApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyRef

`func (o *AIAssistantStartRequestVoiceSettings) SetApiKeyRef(v string)`

SetApiKeyRef sets ApiKeyRef field to given value.

### HasApiKeyRef

`func (o *AIAssistantStartRequestVoiceSettings) HasApiKeyRef() bool`

HasApiKeyRef returns a boolean if a field has been set.

### GetVoiceSpeed

`func (o *AIAssistantStartRequestVoiceSettings) GetVoiceSpeed() float32`

GetVoiceSpeed returns the VoiceSpeed field if non-nil, zero value otherwise.

### GetVoiceSpeedOk

`func (o *AIAssistantStartRequestVoiceSettings) GetVoiceSpeedOk() (*float32, bool)`

GetVoiceSpeedOk returns a tuple with the VoiceSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSpeed

`func (o *AIAssistantStartRequestVoiceSettings) SetVoiceSpeed(v float32)`

SetVoiceSpeed sets VoiceSpeed field to given value.

### HasVoiceSpeed

`func (o *AIAssistantStartRequestVoiceSettings) HasVoiceSpeed() bool`

HasVoiceSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


