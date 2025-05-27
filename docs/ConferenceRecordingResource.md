# ConferenceRecordingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The id of the account the resource belongs to. | [optional] 
**CallSid** | Pointer to **string** | The identifier of the related participant&#39;s call. | [optional] 
**Channels** | Pointer to **int32** | The number of channels in the recording. | [optional] 
**ConferenceSid** | Pointer to **string** | The identifier of the related conference. | [optional] 
**DateCreated** | Pointer to **string** | The timestamp of when the resource was created. | [optional] 
**DateUpdated** | Pointer to **string** | The timestamp of when the resource was last updated. | [optional] 
**Duration** | Pointer to **int32** | Duratin of the recording in seconds. | [optional] 
**ErrorCode** | Pointer to **string** | The recording error, if any. | [optional] 
**MediaUrl** | Pointer to **string** | The URL to use to download the recording. | [optional] 
**Sid** | Pointer to **string** | The unique identifier of the recording. | [optional] 
**Source** | Pointer to **string** | How the recording was started. | [optional] 
**StartTime** | Pointer to **string** | The timestamp of when the recording was started. | [optional] 
**Status** | Pointer to **string** | The status of the recording. | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs. | [optional] 
**Uri** | Pointer to **string** | The relative URI for this recording. | [optional] 

## Methods

### NewConferenceRecordingResource

`func NewConferenceRecordingResource() *ConferenceRecordingResource`

NewConferenceRecordingResource instantiates a new ConferenceRecordingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceRecordingResourceWithDefaults

`func NewConferenceRecordingResourceWithDefaults() *ConferenceRecordingResource`

NewConferenceRecordingResourceWithDefaults instantiates a new ConferenceRecordingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConferenceRecordingResource) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConferenceRecordingResource) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConferenceRecordingResource) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConferenceRecordingResource) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetCallSid

`func (o *ConferenceRecordingResource) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ConferenceRecordingResource) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ConferenceRecordingResource) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ConferenceRecordingResource) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### GetChannels

`func (o *ConferenceRecordingResource) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ConferenceRecordingResource) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ConferenceRecordingResource) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ConferenceRecordingResource) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetConferenceSid

`func (o *ConferenceRecordingResource) GetConferenceSid() string`

GetConferenceSid returns the ConferenceSid field if non-nil, zero value otherwise.

### GetConferenceSidOk

`func (o *ConferenceRecordingResource) GetConferenceSidOk() (*string, bool)`

GetConferenceSidOk returns a tuple with the ConferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceSid

`func (o *ConferenceRecordingResource) SetConferenceSid(v string)`

SetConferenceSid sets ConferenceSid field to given value.

### HasConferenceSid

`func (o *ConferenceRecordingResource) HasConferenceSid() bool`

HasConferenceSid returns a boolean if a field has been set.

### GetDateCreated

`func (o *ConferenceRecordingResource) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConferenceRecordingResource) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConferenceRecordingResource) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConferenceRecordingResource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ConferenceRecordingResource) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConferenceRecordingResource) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConferenceRecordingResource) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConferenceRecordingResource) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDuration

`func (o *ConferenceRecordingResource) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ConferenceRecordingResource) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ConferenceRecordingResource) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ConferenceRecordingResource) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetErrorCode

`func (o *ConferenceRecordingResource) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ConferenceRecordingResource) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ConferenceRecordingResource) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ConferenceRecordingResource) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMediaUrl

`func (o *ConferenceRecordingResource) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *ConferenceRecordingResource) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *ConferenceRecordingResource) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *ConferenceRecordingResource) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetSid

`func (o *ConferenceRecordingResource) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConferenceRecordingResource) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConferenceRecordingResource) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConferenceRecordingResource) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetSource

`func (o *ConferenceRecordingResource) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConferenceRecordingResource) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConferenceRecordingResource) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConferenceRecordingResource) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStartTime

`func (o *ConferenceRecordingResource) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ConferenceRecordingResource) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ConferenceRecordingResource) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ConferenceRecordingResource) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ConferenceRecordingResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConferenceRecordingResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConferenceRecordingResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConferenceRecordingResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubresourceUris

`func (o *ConferenceRecordingResource) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ConferenceRecordingResource) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ConferenceRecordingResource) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ConferenceRecordingResource) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### GetUri

`func (o *ConferenceRecordingResource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ConferenceRecordingResource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ConferenceRecordingResource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ConferenceRecordingResource) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


