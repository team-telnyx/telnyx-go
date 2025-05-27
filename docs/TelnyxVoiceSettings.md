# TelnyxVoiceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoiceSpeed** | Pointer to **float32** | The voice speed to be used for the voice. The voice speed must be between 0.1 and 2.0. Default value is 1.0. | [optional] [default to 1.0]

## Methods

### NewTelnyxVoiceSettings

`func NewTelnyxVoiceSettings() *TelnyxVoiceSettings`

NewTelnyxVoiceSettings instantiates a new TelnyxVoiceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelnyxVoiceSettingsWithDefaults

`func NewTelnyxVoiceSettingsWithDefaults() *TelnyxVoiceSettings`

NewTelnyxVoiceSettingsWithDefaults instantiates a new TelnyxVoiceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoiceSpeed

`func (o *TelnyxVoiceSettings) GetVoiceSpeed() float32`

GetVoiceSpeed returns the VoiceSpeed field if non-nil, zero value otherwise.

### GetVoiceSpeedOk

`func (o *TelnyxVoiceSettings) GetVoiceSpeedOk() (*float32, bool)`

GetVoiceSpeedOk returns a tuple with the VoiceSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSpeed

`func (o *TelnyxVoiceSettings) SetVoiceSpeed(v float32)`

SetVoiceSpeed sets VoiceSpeed field to given value.

### HasVoiceSpeed

`func (o *TelnyxVoiceSettings) HasVoiceSpeed() bool`

HasVoiceSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


