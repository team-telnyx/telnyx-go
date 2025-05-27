# OutboundCallRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallRecordingType** | Pointer to **string** | Specifies which calls are recorded. | [optional] 
**CallRecordingCallerPhoneNumbers** | Pointer to **[]string** | When call_recording_type is &#39;by_caller_phone_number&#39;, only outbound calls using one of these numbers will be recorded. Numbers must be specified in E164 format. | [optional] 
**CallRecordingChannels** | Pointer to **string** | When using &#39;dual&#39; channels, the final audio file will be a stereo recording with the first leg on channel A, and the rest on channel B. | [optional] [default to "single"]
**CallRecordingFormat** | Pointer to **string** | The audio file format for calls being recorded. | [optional] [default to "wav"]

## Methods

### NewOutboundCallRecording

`func NewOutboundCallRecording() *OutboundCallRecording`

NewOutboundCallRecording instantiates a new OutboundCallRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundCallRecordingWithDefaults

`func NewOutboundCallRecordingWithDefaults() *OutboundCallRecording`

NewOutboundCallRecordingWithDefaults instantiates a new OutboundCallRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallRecordingType

`func (o *OutboundCallRecording) GetCallRecordingType() string`

GetCallRecordingType returns the CallRecordingType field if non-nil, zero value otherwise.

### GetCallRecordingTypeOk

`func (o *OutboundCallRecording) GetCallRecordingTypeOk() (*string, bool)`

GetCallRecordingTypeOk returns a tuple with the CallRecordingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingType

`func (o *OutboundCallRecording) SetCallRecordingType(v string)`

SetCallRecordingType sets CallRecordingType field to given value.

### HasCallRecordingType

`func (o *OutboundCallRecording) HasCallRecordingType() bool`

HasCallRecordingType returns a boolean if a field has been set.

### GetCallRecordingCallerPhoneNumbers

`func (o *OutboundCallRecording) GetCallRecordingCallerPhoneNumbers() []string`

GetCallRecordingCallerPhoneNumbers returns the CallRecordingCallerPhoneNumbers field if non-nil, zero value otherwise.

### GetCallRecordingCallerPhoneNumbersOk

`func (o *OutboundCallRecording) GetCallRecordingCallerPhoneNumbersOk() (*[]string, bool)`

GetCallRecordingCallerPhoneNumbersOk returns a tuple with the CallRecordingCallerPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingCallerPhoneNumbers

`func (o *OutboundCallRecording) SetCallRecordingCallerPhoneNumbers(v []string)`

SetCallRecordingCallerPhoneNumbers sets CallRecordingCallerPhoneNumbers field to given value.

### HasCallRecordingCallerPhoneNumbers

`func (o *OutboundCallRecording) HasCallRecordingCallerPhoneNumbers() bool`

HasCallRecordingCallerPhoneNumbers returns a boolean if a field has been set.

### GetCallRecordingChannels

`func (o *OutboundCallRecording) GetCallRecordingChannels() string`

GetCallRecordingChannels returns the CallRecordingChannels field if non-nil, zero value otherwise.

### GetCallRecordingChannelsOk

`func (o *OutboundCallRecording) GetCallRecordingChannelsOk() (*string, bool)`

GetCallRecordingChannelsOk returns a tuple with the CallRecordingChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingChannels

`func (o *OutboundCallRecording) SetCallRecordingChannels(v string)`

SetCallRecordingChannels sets CallRecordingChannels field to given value.

### HasCallRecordingChannels

`func (o *OutboundCallRecording) HasCallRecordingChannels() bool`

HasCallRecordingChannels returns a boolean if a field has been set.

### GetCallRecordingFormat

`func (o *OutboundCallRecording) GetCallRecordingFormat() string`

GetCallRecordingFormat returns the CallRecordingFormat field if non-nil, zero value otherwise.

### GetCallRecordingFormatOk

`func (o *OutboundCallRecording) GetCallRecordingFormatOk() (*string, bool)`

GetCallRecordingFormatOk returns a tuple with the CallRecordingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingFormat

`func (o *OutboundCallRecording) SetCallRecordingFormat(v string)`

SetCallRecordingFormat sets CallRecordingFormat field to given value.

### HasCallRecordingFormat

`func (o *OutboundCallRecording) HasCallRecordingFormat() bool`

HasCallRecordingFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


