# VoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voice** | **string** | The voice to be used by the voice assistant. Check the full list of [available voices](https://developers.telnyx.com/api/call-control/list-text-to-speech-voices) via our voices API. To use ElevenLabs, you must reference your ElevenLabs API key as an integration secret under the &#x60;api_key_ref&#x60; field. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. For Telnyx voices, use &#x60;Telnyx.&lt;model_id&gt;.&lt;voice_id&gt;&#x60; (e.g. Telnyx.KokoroTTS.af_heart) | 
**ApiKeyRef** | Pointer to **string** | The &#x60;identifier&#x60; for an integration secret [/v2/integration_secrets](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) that refers to your ElevenLabs API key. Warning: Free plans are unlikely to work with this integration. | [optional] 

## Methods

### NewVoiceSettings

`func NewVoiceSettings(voice string, ) *VoiceSettings`

NewVoiceSettings instantiates a new VoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceSettingsWithDefaults

`func NewVoiceSettingsWithDefaults() *VoiceSettings`

NewVoiceSettingsWithDefaults instantiates a new VoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoice

`func (o *VoiceSettings) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *VoiceSettings) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *VoiceSettings) SetVoice(v string)`

SetVoice sets Voice field to given value.


### GetApiKeyRef

`func (o *VoiceSettings) GetApiKeyRef() string`

GetApiKeyRef returns the ApiKeyRef field if non-nil, zero value otherwise.

### GetApiKeyRefOk

`func (o *VoiceSettings) GetApiKeyRefOk() (*string, bool)`

GetApiKeyRefOk returns a tuple with the ApiKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyRef

`func (o *VoiceSettings) SetApiKeyRef(v string)`

SetApiKeyRef sets ApiKeyRef field to given value.

### HasApiKeyRef

`func (o *VoiceSettings) HasApiKeyRef() bool`

HasApiKeyRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


