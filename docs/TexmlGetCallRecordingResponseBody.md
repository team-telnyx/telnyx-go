# TexmlGetCallRecordingResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** |  | [optional] 
**CallSid** | Pointer to **string** |  | [optional] 
**ConferenceSid** | Pointer to **NullableString** |  | [optional] 
**Channels** | Pointer to [**TwimlRecordingChannels**](TwimlRecordingChannels.md) |  | [optional] [default to _2]
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **NullableString** | The duration of this recording, given in seconds. | [optional] 
**Sid** | Pointer to **string** | Identifier of a resource. | [optional] 
**Source** | Pointer to [**RecordingSource**](RecordingSource.md) |  | [optional] 
**Status** | Pointer to [**TexmlRecordingStatus**](TexmlRecordingStatus.md) |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**SubresourcesUris** | Pointer to [**TexmlRecordingSubresourcesUris**](TexmlRecordingSubresourcesUris.md) |  | [optional] 
**Uri** | Pointer to **string** | The relative URI for this recording resource. | [optional] 
**MediaUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewTexmlGetCallRecordingResponseBody

`func NewTexmlGetCallRecordingResponseBody() *TexmlGetCallRecordingResponseBody`

NewTexmlGetCallRecordingResponseBody instantiates a new TexmlGetCallRecordingResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTexmlGetCallRecordingResponseBodyWithDefaults

`func NewTexmlGetCallRecordingResponseBodyWithDefaults() *TexmlGetCallRecordingResponseBody`

NewTexmlGetCallRecordingResponseBodyWithDefaults instantiates a new TexmlGetCallRecordingResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TexmlGetCallRecordingResponseBody) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TexmlGetCallRecordingResponseBody) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TexmlGetCallRecordingResponseBody) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TexmlGetCallRecordingResponseBody) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetCallSid

`func (o *TexmlGetCallRecordingResponseBody) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *TexmlGetCallRecordingResponseBody) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *TexmlGetCallRecordingResponseBody) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *TexmlGetCallRecordingResponseBody) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### GetConferenceSid

`func (o *TexmlGetCallRecordingResponseBody) GetConferenceSid() string`

GetConferenceSid returns the ConferenceSid field if non-nil, zero value otherwise.

### GetConferenceSidOk

`func (o *TexmlGetCallRecordingResponseBody) GetConferenceSidOk() (*string, bool)`

GetConferenceSidOk returns a tuple with the ConferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceSid

`func (o *TexmlGetCallRecordingResponseBody) SetConferenceSid(v string)`

SetConferenceSid sets ConferenceSid field to given value.

### HasConferenceSid

`func (o *TexmlGetCallRecordingResponseBody) HasConferenceSid() bool`

HasConferenceSid returns a boolean if a field has been set.

### SetConferenceSidNil

`func (o *TexmlGetCallRecordingResponseBody) SetConferenceSidNil(b bool)`

 SetConferenceSidNil sets the value for ConferenceSid to be an explicit nil

### UnsetConferenceSid
`func (o *TexmlGetCallRecordingResponseBody) UnsetConferenceSid()`

UnsetConferenceSid ensures that no value is present for ConferenceSid, not even an explicit nil
### GetChannels

`func (o *TexmlGetCallRecordingResponseBody) GetChannels() TwimlRecordingChannels`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TexmlGetCallRecordingResponseBody) GetChannelsOk() (*TwimlRecordingChannels, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TexmlGetCallRecordingResponseBody) SetChannels(v TwimlRecordingChannels)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *TexmlGetCallRecordingResponseBody) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDateCreated

`func (o *TexmlGetCallRecordingResponseBody) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TexmlGetCallRecordingResponseBody) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TexmlGetCallRecordingResponseBody) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TexmlGetCallRecordingResponseBody) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TexmlGetCallRecordingResponseBody) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TexmlGetCallRecordingResponseBody) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TexmlGetCallRecordingResponseBody) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TexmlGetCallRecordingResponseBody) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetStartTime

`func (o *TexmlGetCallRecordingResponseBody) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TexmlGetCallRecordingResponseBody) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TexmlGetCallRecordingResponseBody) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TexmlGetCallRecordingResponseBody) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDuration

`func (o *TexmlGetCallRecordingResponseBody) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TexmlGetCallRecordingResponseBody) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TexmlGetCallRecordingResponseBody) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TexmlGetCallRecordingResponseBody) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TexmlGetCallRecordingResponseBody) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TexmlGetCallRecordingResponseBody) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSid

`func (o *TexmlGetCallRecordingResponseBody) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TexmlGetCallRecordingResponseBody) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TexmlGetCallRecordingResponseBody) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TexmlGetCallRecordingResponseBody) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetSource

`func (o *TexmlGetCallRecordingResponseBody) GetSource() RecordingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TexmlGetCallRecordingResponseBody) GetSourceOk() (*RecordingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TexmlGetCallRecordingResponseBody) SetSource(v RecordingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TexmlGetCallRecordingResponseBody) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *TexmlGetCallRecordingResponseBody) GetStatus() TexmlRecordingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TexmlGetCallRecordingResponseBody) GetStatusOk() (*TexmlRecordingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TexmlGetCallRecordingResponseBody) SetStatus(v TexmlRecordingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TexmlGetCallRecordingResponseBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *TexmlGetCallRecordingResponseBody) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TexmlGetCallRecordingResponseBody) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TexmlGetCallRecordingResponseBody) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TexmlGetCallRecordingResponseBody) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *TexmlGetCallRecordingResponseBody) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *TexmlGetCallRecordingResponseBody) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetSubresourcesUris

`func (o *TexmlGetCallRecordingResponseBody) GetSubresourcesUris() TexmlRecordingSubresourcesUris`

GetSubresourcesUris returns the SubresourcesUris field if non-nil, zero value otherwise.

### GetSubresourcesUrisOk

`func (o *TexmlGetCallRecordingResponseBody) GetSubresourcesUrisOk() (*TexmlRecordingSubresourcesUris, bool)`

GetSubresourcesUrisOk returns a tuple with the SubresourcesUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourcesUris

`func (o *TexmlGetCallRecordingResponseBody) SetSubresourcesUris(v TexmlRecordingSubresourcesUris)`

SetSubresourcesUris sets SubresourcesUris field to given value.

### HasSubresourcesUris

`func (o *TexmlGetCallRecordingResponseBody) HasSubresourcesUris() bool`

HasSubresourcesUris returns a boolean if a field has been set.

### GetUri

`func (o *TexmlGetCallRecordingResponseBody) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TexmlGetCallRecordingResponseBody) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TexmlGetCallRecordingResponseBody) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TexmlGetCallRecordingResponseBody) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetMediaUrl

`func (o *TexmlGetCallRecordingResponseBody) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *TexmlGetCallRecordingResponseBody) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *TexmlGetCallRecordingResponseBody) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *TexmlGetCallRecordingResponseBody) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


