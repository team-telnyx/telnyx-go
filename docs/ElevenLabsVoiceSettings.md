# ElevenLabsVoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyRef** | Pointer to **string** | The &#x60;identifier&#x60; for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work with this integration. | [optional] 

## Methods

### NewElevenLabsVoiceSettings

`func NewElevenLabsVoiceSettings() *ElevenLabsVoiceSettings`

NewElevenLabsVoiceSettings instantiates a new ElevenLabsVoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElevenLabsVoiceSettingsWithDefaults

`func NewElevenLabsVoiceSettingsWithDefaults() *ElevenLabsVoiceSettings`

NewElevenLabsVoiceSettingsWithDefaults instantiates a new ElevenLabsVoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyRef

`func (o *ElevenLabsVoiceSettings) GetApiKeyRef() string`

GetApiKeyRef returns the ApiKeyRef field if non-nil, zero value otherwise.

### GetApiKeyRefOk

`func (o *ElevenLabsVoiceSettings) GetApiKeyRefOk() (*string, bool)`

GetApiKeyRefOk returns a tuple with the ApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyRef

`func (o *ElevenLabsVoiceSettings) SetApiKeyRef(v string)`

SetApiKeyRef sets ApiKeyRef field to given value.

### HasApiKeyRef

`func (o *ElevenLabsVoiceSettings) HasApiKeyRef() bool`

HasApiKeyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


