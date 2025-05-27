# ConferencePlaybackEndedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CreatorCallSessionId** | Pointer to **string** | ID that is unique to the call session that started the conference. | [optional] 
**ConferenceId** | Pointer to **string** | ID of the conference the text was spoken in. | [optional] 
**MediaUrl** | Pointer to **string** | The audio URL being played back, if audio_url has been used to start. | [optional] 
**MediaName** | Pointer to **string** | The name of the audio media file being played back, if media_name has been used to start. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 

## Methods

### NewConferencePlaybackEndedPayload

`func NewConferencePlaybackEndedPayload() *ConferencePlaybackEndedPayload`

NewConferencePlaybackEndedPayload instantiates a new ConferencePlaybackEndedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferencePlaybackEndedPayloadWithDefaults

`func NewConferencePlaybackEndedPayloadWithDefaults() *ConferencePlaybackEndedPayload`

NewConferencePlaybackEndedPayloadWithDefaults instantiates a new ConferencePlaybackEndedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConferencePlaybackEndedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferencePlaybackEndedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferencePlaybackEndedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferencePlaybackEndedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatorCallSessionId

`func (o *ConferencePlaybackEndedPayload) GetCreatorCallSessionId() string`

GetCreatorCallSessionId returns the CreatorCallSessionId field if non-nil, zero value otherwise.

### GetCreatorCallSessionIdOk

`func (o *ConferencePlaybackEndedPayload) GetCreatorCallSessionIdOk() (*string, bool)`

GetCreatorCallSessionIdOk returns a tuple with the CreatorCallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCallSessionId

`func (o *ConferencePlaybackEndedPayload) SetCreatorCallSessionId(v string)`

SetCreatorCallSessionId sets CreatorCallSessionId field to given value.

### HasCreatorCallSessionId

`func (o *ConferencePlaybackEndedPayload) HasCreatorCallSessionId() bool`

HasCreatorCallSessionId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferencePlaybackEndedPayload) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferencePlaybackEndedPayload) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferencePlaybackEndedPayload) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferencePlaybackEndedPayload) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetMediaUrl

`func (o *ConferencePlaybackEndedPayload) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *ConferencePlaybackEndedPayload) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *ConferencePlaybackEndedPayload) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *ConferencePlaybackEndedPayload) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *ConferencePlaybackEndedPayload) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *ConferencePlaybackEndedPayload) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *ConferencePlaybackEndedPayload) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *ConferencePlaybackEndedPayload) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetOccurredAt

`func (o *ConferencePlaybackEndedPayload) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ConferencePlaybackEndedPayload) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ConferencePlaybackEndedPayload) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *ConferencePlaybackEndedPayload) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


