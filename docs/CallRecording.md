# CallRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InboundCallRecordingEnabled** | Pointer to **bool** | When enabled, any inbound call to this number will be recorded. | [optional] [default to false]
**InboundCallRecordingFormat** | Pointer to **string** | The audio file format for calls being recorded. | [optional] [default to "wav"]
**InboundCallRecordingChannels** | Pointer to **string** | When using &#39;dual&#39; channels, final audio file will be stereo recorded with the first leg on channel A, and the rest on channel B. | [optional] [default to "single"]

## Methods

### NewCallRecording

`func NewCallRecording() *CallRecording`

NewCallRecording instantiates a new CallRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordingWithDefaults

`func NewCallRecordingWithDefaults() *CallRecording`

NewCallRecordingWithDefaults instantiates a new CallRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInboundCallRecordingEnabled

`func (o *CallRecording) GetInboundCallRecordingEnabled() bool`

GetInboundCallRecordingEnabled returns the InboundCallRecordingEnabled field if non-nil, zero value otherwise.

### GetInboundCallRecordingEnabledOk

`func (o *CallRecording) GetInboundCallRecordingEnabledOk() (*bool, bool)`

GetInboundCallRecordingEnabledOk returns a tuple with the InboundCallRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallRecordingEnabled

`func (o *CallRecording) SetInboundCallRecordingEnabled(v bool)`

SetInboundCallRecordingEnabled sets InboundCallRecordingEnabled field to given value.

### HasInboundCallRecordingEnabled

`func (o *CallRecording) HasInboundCallRecordingEnabled() bool`

HasInboundCallRecordingEnabled returns a boolean if a field has been set.

### GetInboundCallRecordingFormat

`func (o *CallRecording) GetInboundCallRecordingFormat() string`

GetInboundCallRecordingFormat returns the InboundCallRecordingFormat field if non-nil, zero value otherwise.

### GetInboundCallRecordingFormatOk

`func (o *CallRecording) GetInboundCallRecordingFormatOk() (*string, bool)`

GetInboundCallRecordingFormatOk returns a tuple with the InboundCallRecordingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallRecordingFormat

`func (o *CallRecording) SetInboundCallRecordingFormat(v string)`

SetInboundCallRecordingFormat sets InboundCallRecordingFormat field to given value.

### HasInboundCallRecordingFormat

`func (o *CallRecording) HasInboundCallRecordingFormat() bool`

HasInboundCallRecordingFormat returns a boolean if a field has been set.

### GetInboundCallRecordingChannels

`func (o *CallRecording) GetInboundCallRecordingChannels() string`

GetInboundCallRecordingChannels returns the InboundCallRecordingChannels field if non-nil, zero value otherwise.

### GetInboundCallRecordingChannelsOk

`func (o *CallRecording) GetInboundCallRecordingChannelsOk() (*string, bool)`

GetInboundCallRecordingChannelsOk returns a tuple with the InboundCallRecordingChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallRecordingChannels

`func (o *CallRecording) SetInboundCallRecordingChannels(v string)`

SetInboundCallRecordingChannels sets InboundCallRecordingChannels field to given value.

### HasInboundCallRecordingChannels

`func (o *CallRecording) HasInboundCallRecordingChannels() bool`

HasInboundCallRecordingChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


