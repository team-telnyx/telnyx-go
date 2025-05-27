# ConferenceParticipantSpeakEndedPayload

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
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 

## Methods

### NewConferenceParticipantSpeakEndedPayload

`func NewConferenceParticipantSpeakEndedPayload() *ConferenceParticipantSpeakEndedPayload`

NewConferenceParticipantSpeakEndedPayload instantiates a new ConferenceParticipantSpeakEndedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceParticipantSpeakEndedPayloadWithDefaults

`func NewConferenceParticipantSpeakEndedPayloadWithDefaults() *ConferenceParticipantSpeakEndedPayload`

NewConferenceParticipantSpeakEndedPayloadWithDefaults instantiates a new ConferenceParticipantSpeakEndedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *ConferenceParticipantSpeakEndedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *ConferenceParticipantSpeakEndedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *ConferenceParticipantSpeakEndedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *ConferenceParticipantSpeakEndedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *ConferenceParticipantSpeakEndedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *ConferenceParticipantSpeakEndedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *ConferenceParticipantSpeakEndedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ConferenceParticipantSpeakEndedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *ConferenceParticipantSpeakEndedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *ConferenceParticipantSpeakEndedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ConferenceParticipantSpeakEndedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *ConferenceParticipantSpeakEndedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetConnectionId

`func (o *ConferenceParticipantSpeakEndedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferenceParticipantSpeakEndedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferenceParticipantSpeakEndedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatorCallSessionId

`func (o *ConferenceParticipantSpeakEndedPayload) GetCreatorCallSessionId() string`

GetCreatorCallSessionId returns the CreatorCallSessionId field if non-nil, zero value otherwise.

### GetCreatorCallSessionIdOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetCreatorCallSessionIdOk() (*string, bool)`

GetCreatorCallSessionIdOk returns a tuple with the CreatorCallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCallSessionId

`func (o *ConferenceParticipantSpeakEndedPayload) SetCreatorCallSessionId(v string)`

SetCreatorCallSessionId sets CreatorCallSessionId field to given value.

### HasCreatorCallSessionId

`func (o *ConferenceParticipantSpeakEndedPayload) HasCreatorCallSessionId() bool`

HasCreatorCallSessionId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceParticipantSpeakEndedPayload) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceParticipantSpeakEndedPayload) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceParticipantSpeakEndedPayload) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *ConferenceParticipantSpeakEndedPayload) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ConferenceParticipantSpeakEndedPayload) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ConferenceParticipantSpeakEndedPayload) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *ConferenceParticipantSpeakEndedPayload) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


