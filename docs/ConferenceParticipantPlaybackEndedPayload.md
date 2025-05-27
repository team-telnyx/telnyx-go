# ConferenceParticipantPlaybackEndedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Participant&#39;s call ID used to issue commands via Call Control API. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CreatorCallSessionId** | Pointer to **string** | ID that is unique to the call session that started the conference. | [optional] 
**ConferenceId** | Pointer to **string** | ID of the conference the text was spoken in. | [optional] 
**MediaUrl** | Pointer to **string** | The audio URL being played back, if audio_url has been used to start. | [optional] 
**MediaName** | Pointer to **string** | The name of the audio media file being played back, if media_name has been used to start. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 

## Methods

### NewConferenceParticipantPlaybackEndedPayload

`func NewConferenceParticipantPlaybackEndedPayload() *ConferenceParticipantPlaybackEndedPayload`

NewConferenceParticipantPlaybackEndedPayload instantiates a new ConferenceParticipantPlaybackEndedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceParticipantPlaybackEndedPayloadWithDefaults

`func NewConferenceParticipantPlaybackEndedPayloadWithDefaults() *ConferenceParticipantPlaybackEndedPayload`

NewConferenceParticipantPlaybackEndedPayloadWithDefaults instantiates a new ConferenceParticipantPlaybackEndedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *ConferenceParticipantPlaybackEndedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *ConferenceParticipantPlaybackEndedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *ConferenceParticipantPlaybackEndedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *ConferenceParticipantPlaybackEndedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ConferenceParticipantPlaybackEndedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *ConferenceParticipantPlaybackEndedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *ConferenceParticipantPlaybackEndedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ConferenceParticipantPlaybackEndedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *ConferenceParticipantPlaybackEndedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetConnectionId

`func (o *ConferenceParticipantPlaybackEndedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferenceParticipantPlaybackEndedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferenceParticipantPlaybackEndedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatorCallSessionId

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCreatorCallSessionId() string`

GetCreatorCallSessionId returns the CreatorCallSessionId field if non-nil, zero value otherwise.

### GetCreatorCallSessionIdOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetCreatorCallSessionIdOk() (*string, bool)`

GetCreatorCallSessionIdOk returns a tuple with the CreatorCallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCallSessionId

`func (o *ConferenceParticipantPlaybackEndedPayload) SetCreatorCallSessionId(v string)`

SetCreatorCallSessionId sets CreatorCallSessionId field to given value.

### HasCreatorCallSessionId

`func (o *ConferenceParticipantPlaybackEndedPayload) HasCreatorCallSessionId() bool`

HasCreatorCallSessionId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceParticipantPlaybackEndedPayload) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceParticipantPlaybackEndedPayload) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceParticipantPlaybackEndedPayload) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetMediaUrl

`func (o *ConferenceParticipantPlaybackEndedPayload) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *ConferenceParticipantPlaybackEndedPayload) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *ConferenceParticipantPlaybackEndedPayload) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *ConferenceParticipantPlaybackEndedPayload) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *ConferenceParticipantPlaybackEndedPayload) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *ConferenceParticipantPlaybackEndedPayload) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetOccurredAt

`func (o *ConferenceParticipantPlaybackEndedPayload) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ConferenceParticipantPlaybackEndedPayload) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ConferenceParticipantPlaybackEndedPayload) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *ConferenceParticipantPlaybackEndedPayload) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


