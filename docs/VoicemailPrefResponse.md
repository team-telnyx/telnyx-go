# VoicemailPrefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether voicemail is enabled. | [optional] 
**Pin** | Pointer to **string** | The pin used for the voicemail. | [optional] 

## Methods

### NewVoicemailPrefResponse

`func NewVoicemailPrefResponse() *VoicemailPrefResponse`

NewVoicemailPrefResponse instantiates a new VoicemailPrefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailPrefResponseWithDefaults

`func NewVoicemailPrefResponseWithDefaults() *VoicemailPrefResponse`

NewVoicemailPrefResponseWithDefaults instantiates a new VoicemailPrefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VoicemailPrefResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VoicemailPrefResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VoicemailPrefResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VoicemailPrefResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPin

`func (o *VoicemailPrefResponse) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *VoicemailPrefResponse) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *VoicemailPrefResponse) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *VoicemailPrefResponse) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


