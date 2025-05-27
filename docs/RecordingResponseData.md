# RecordingResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Unique identifier and token for controlling the call. | [optional] 
**CallLegId** | Pointer to **string** | ID unique to the call leg (used to correlate webhook events). | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**Channels** | Pointer to **string** | When &#x60;dual&#x60;, the final audio file has the first leg on channel A, and the rest on channel B. | [optional] 
**ConferenceId** | Pointer to **string** | Uniquely identifies the conference. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**DownloadUrls** | Pointer to [**RecordingResponseDataDownloadUrls**](RecordingResponseDataDownloadUrls.md) |  | [optional] 
**DurationMillis** | Pointer to **int32** | The duration of the recording in milliseconds. | [optional] 
**Id** | Pointer to **string** | Uniquely identifies the recording. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**RecordingStartedAt** | Pointer to **string** | ISO 8601 formatted date of when the recording started. | [optional] 
**RecordingEndedAt** | Pointer to **string** | ISO 8601 formatted date of when the recording ended. | [optional] 
**Source** | Pointer to **string** | The kind of event that led to this recording being created. | [optional] 
**Status** | Pointer to **string** | The status of the recording. Only &#x60;completed&#x60; recordings are currently supported. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewRecordingResponseData

`func NewRecordingResponseData() *RecordingResponseData`

NewRecordingResponseData instantiates a new RecordingResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingResponseDataWithDefaults

`func NewRecordingResponseDataWithDefaults() *RecordingResponseData`

NewRecordingResponseDataWithDefaults instantiates a new RecordingResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *RecordingResponseData) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *RecordingResponseData) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *RecordingResponseData) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *RecordingResponseData) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *RecordingResponseData) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *RecordingResponseData) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *RecordingResponseData) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *RecordingResponseData) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *RecordingResponseData) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *RecordingResponseData) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *RecordingResponseData) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *RecordingResponseData) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetChannels

`func (o *RecordingResponseData) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *RecordingResponseData) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *RecordingResponseData) SetChannels(v string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *RecordingResponseData) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetConferenceId

`func (o *RecordingResponseData) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *RecordingResponseData) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *RecordingResponseData) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *RecordingResponseData) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RecordingResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RecordingResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RecordingResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RecordingResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDownloadUrls

`func (o *RecordingResponseData) GetDownloadUrls() RecordingResponseDataDownloadUrls`

GetDownloadUrls returns the DownloadUrls field if non-nil, zero value otherwise.

### GetDownloadUrlsOk

`func (o *RecordingResponseData) GetDownloadUrlsOk() (*RecordingResponseDataDownloadUrls, bool)`

GetDownloadUrlsOk returns a tuple with the DownloadUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrls

`func (o *RecordingResponseData) SetDownloadUrls(v RecordingResponseDataDownloadUrls)`

SetDownloadUrls sets DownloadUrls field to given value.

### HasDownloadUrls

`func (o *RecordingResponseData) HasDownloadUrls() bool`

HasDownloadUrls returns a boolean if a field has been set.

### GetDurationMillis

`func (o *RecordingResponseData) GetDurationMillis() int32`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *RecordingResponseData) GetDurationMillisOk() (*int32, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *RecordingResponseData) SetDurationMillis(v int32)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *RecordingResponseData) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetId

`func (o *RecordingResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordingResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordingResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecordingResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *RecordingResponseData) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RecordingResponseData) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RecordingResponseData) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RecordingResponseData) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordingStartedAt

`func (o *RecordingResponseData) GetRecordingStartedAt() string`

GetRecordingStartedAt returns the RecordingStartedAt field if non-nil, zero value otherwise.

### GetRecordingStartedAtOk

`func (o *RecordingResponseData) GetRecordingStartedAtOk() (*string, bool)`

GetRecordingStartedAtOk returns a tuple with the RecordingStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStartedAt

`func (o *RecordingResponseData) SetRecordingStartedAt(v string)`

SetRecordingStartedAt sets RecordingStartedAt field to given value.

### HasRecordingStartedAt

`func (o *RecordingResponseData) HasRecordingStartedAt() bool`

HasRecordingStartedAt returns a boolean if a field has been set.

### GetRecordingEndedAt

`func (o *RecordingResponseData) GetRecordingEndedAt() string`

GetRecordingEndedAt returns the RecordingEndedAt field if non-nil, zero value otherwise.

### GetRecordingEndedAtOk

`func (o *RecordingResponseData) GetRecordingEndedAtOk() (*string, bool)`

GetRecordingEndedAtOk returns a tuple with the RecordingEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEndedAt

`func (o *RecordingResponseData) SetRecordingEndedAt(v string)`

SetRecordingEndedAt sets RecordingEndedAt field to given value.

### HasRecordingEndedAt

`func (o *RecordingResponseData) HasRecordingEndedAt() bool`

HasRecordingEndedAt returns a boolean if a field has been set.

### GetSource

`func (o *RecordingResponseData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingResponseData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingResponseData) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingResponseData) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *RecordingResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecordingResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecordingResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecordingResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RecordingResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RecordingResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RecordingResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RecordingResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


