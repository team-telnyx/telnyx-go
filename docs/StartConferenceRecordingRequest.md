# StartConferenceRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The audio file format used when storing the conference recording. Can be either &#x60;mp3&#x60; or &#x60;wav&#x60;. | 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;conference_id&#x60;. | [optional] 
**PlayBeep** | Pointer to **bool** | If enabled, a beep sound will be played at the start of a recording. | [optional] 
**Trim** | Pointer to **string** | When set to &#x60;trim-silence&#x60;, silence will be removed from the beginning and end of the recording. | [optional] 
**CustomFileName** | Pointer to **string** | The custom recording file name to be used instead of the default &#x60;call_leg_id&#x60;. Telnyx will still add a Unix timestamp suffix. | [optional] 

## Methods

### NewStartConferenceRecordingRequest

`func NewStartConferenceRecordingRequest(format string, ) *StartConferenceRecordingRequest`

NewStartConferenceRecordingRequest instantiates a new StartConferenceRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartConferenceRecordingRequestWithDefaults

`func NewStartConferenceRecordingRequestWithDefaults() *StartConferenceRecordingRequest`

NewStartConferenceRecordingRequestWithDefaults instantiates a new StartConferenceRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *StartConferenceRecordingRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StartConferenceRecordingRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StartConferenceRecordingRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetCommandId

`func (o *StartConferenceRecordingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *StartConferenceRecordingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *StartConferenceRecordingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *StartConferenceRecordingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetPlayBeep

`func (o *StartConferenceRecordingRequest) GetPlayBeep() bool`

GetPlayBeep returns the PlayBeep field if non-nil, zero value otherwise.

### GetPlayBeepOk

`func (o *StartConferenceRecordingRequest) GetPlayBeepOk() (*bool, bool)`

GetPlayBeepOk returns a tuple with the PlayBeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayBeep

`func (o *StartConferenceRecordingRequest) SetPlayBeep(v bool)`

SetPlayBeep sets PlayBeep field to given value.

### HasPlayBeep

`func (o *StartConferenceRecordingRequest) HasPlayBeep() bool`

HasPlayBeep returns a boolean if a field has been set.

### GetTrim

`func (o *StartConferenceRecordingRequest) GetTrim() string`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *StartConferenceRecordingRequest) GetTrimOk() (*string, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *StartConferenceRecordingRequest) SetTrim(v string)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *StartConferenceRecordingRequest) HasTrim() bool`

HasTrim returns a boolean if a field has been set.

### GetCustomFileName

`func (o *StartConferenceRecordingRequest) GetCustomFileName() string`

GetCustomFileName returns the CustomFileName field if non-nil, zero value otherwise.

### GetCustomFileNameOk

`func (o *StartConferenceRecordingRequest) GetCustomFileNameOk() (*string, bool)`

GetCustomFileNameOk returns a tuple with the CustomFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFileName

`func (o *StartConferenceRecordingRequest) SetCustomFileName(v string)`

SetCustomFileName sets CustomFileName field to given value.

### HasCustomFileName

`func (o *StartConferenceRecordingRequest) HasCustomFileName() bool`

HasCustomFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


