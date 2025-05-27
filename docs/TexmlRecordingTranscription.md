# TexmlRecordingTranscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** |  | [optional] 
**CallSid** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** | The version of the API that was used to make the request. | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **NullableString** | The duration of this recording, given in seconds. | [optional] 
**Sid** | Pointer to **string** | Identifier of a resource. | [optional] 
**RecordingSid** | Pointer to **string** | Identifier of a resource. | [optional] 
**Status** | Pointer to **string** | The status of the recording transcriptions. The transcription text will be available only when the status is completed. | [optional] 
**TranscriptionText** | Pointer to **string** | The recording&#39;s transcribed text | [optional] 
**Uri** | Pointer to **string** | The relative URI for the recording transcription resource. | [optional] 

## Methods

### NewTexmlRecordingTranscription

`func NewTexmlRecordingTranscription() *TexmlRecordingTranscription`

NewTexmlRecordingTranscription instantiates a new TexmlRecordingTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTexmlRecordingTranscriptionWithDefaults

`func NewTexmlRecordingTranscriptionWithDefaults() *TexmlRecordingTranscription`

NewTexmlRecordingTranscriptionWithDefaults instantiates a new TexmlRecordingTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TexmlRecordingTranscription) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TexmlRecordingTranscription) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TexmlRecordingTranscription) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TexmlRecordingTranscription) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetCallSid

`func (o *TexmlRecordingTranscription) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *TexmlRecordingTranscription) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *TexmlRecordingTranscription) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *TexmlRecordingTranscription) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### GetApiVersion

`func (o *TexmlRecordingTranscription) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TexmlRecordingTranscription) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TexmlRecordingTranscription) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TexmlRecordingTranscription) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDateCreated

`func (o *TexmlRecordingTranscription) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TexmlRecordingTranscription) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TexmlRecordingTranscription) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TexmlRecordingTranscription) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TexmlRecordingTranscription) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TexmlRecordingTranscription) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TexmlRecordingTranscription) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TexmlRecordingTranscription) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDuration

`func (o *TexmlRecordingTranscription) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TexmlRecordingTranscription) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TexmlRecordingTranscription) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TexmlRecordingTranscription) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TexmlRecordingTranscription) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TexmlRecordingTranscription) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSid

`func (o *TexmlRecordingTranscription) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TexmlRecordingTranscription) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TexmlRecordingTranscription) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TexmlRecordingTranscription) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetRecordingSid

`func (o *TexmlRecordingTranscription) GetRecordingSid() string`

GetRecordingSid returns the RecordingSid field if non-nil, zero value otherwise.

### GetRecordingSidOk

`func (o *TexmlRecordingTranscription) GetRecordingSidOk() (*string, bool)`

GetRecordingSidOk returns a tuple with the RecordingSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingSid

`func (o *TexmlRecordingTranscription) SetRecordingSid(v string)`

SetRecordingSid sets RecordingSid field to given value.

### HasRecordingSid

`func (o *TexmlRecordingTranscription) HasRecordingSid() bool`

HasRecordingSid returns a boolean if a field has been set.

### GetStatus

`func (o *TexmlRecordingTranscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TexmlRecordingTranscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TexmlRecordingTranscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TexmlRecordingTranscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTranscriptionText

`func (o *TexmlRecordingTranscription) GetTranscriptionText() string`

GetTranscriptionText returns the TranscriptionText field if non-nil, zero value otherwise.

### GetTranscriptionTextOk

`func (o *TexmlRecordingTranscription) GetTranscriptionTextOk() (*string, bool)`

GetTranscriptionTextOk returns a tuple with the TranscriptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionText

`func (o *TexmlRecordingTranscription) SetTranscriptionText(v string)`

SetTranscriptionText sets TranscriptionText field to given value.

### HasTranscriptionText

`func (o *TexmlRecordingTranscription) HasTranscriptionText() bool`

HasTranscriptionText returns a boolean if a field has been set.

### GetUri

`func (o *TexmlRecordingTranscription) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TexmlRecordingTranscription) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TexmlRecordingTranscription) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TexmlRecordingTranscription) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


