# ConferenceSpeakEndedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CreatorCallSessionId** | Pointer to **string** | ID that is unique to the call session that started the conference. | [optional] 
**ConferenceId** | Pointer to **string** | ID of the conference the text was spoken in. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 

## Methods

### NewConferenceSpeakEndedPayload

`func NewConferenceSpeakEndedPayload() *ConferenceSpeakEndedPayload`

NewConferenceSpeakEndedPayload instantiates a new ConferenceSpeakEndedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceSpeakEndedPayloadWithDefaults

`func NewConferenceSpeakEndedPayloadWithDefaults() *ConferenceSpeakEndedPayload`

NewConferenceSpeakEndedPayloadWithDefaults instantiates a new ConferenceSpeakEndedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConferenceSpeakEndedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferenceSpeakEndedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferenceSpeakEndedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferenceSpeakEndedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatorCallSessionId

`func (o *ConferenceSpeakEndedPayload) GetCreatorCallSessionId() string`

GetCreatorCallSessionId returns the CreatorCallSessionId field if non-nil, zero value otherwise.

### GetCreatorCallSessionIdOk

`func (o *ConferenceSpeakEndedPayload) GetCreatorCallSessionIdOk() (*string, bool)`

GetCreatorCallSessionIdOk returns a tuple with the CreatorCallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorCallSessionId

`func (o *ConferenceSpeakEndedPayload) SetCreatorCallSessionId(v string)`

SetCreatorCallSessionId sets CreatorCallSessionId field to given value.

### HasCreatorCallSessionId

`func (o *ConferenceSpeakEndedPayload) HasCreatorCallSessionId() bool`

HasCreatorCallSessionId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceSpeakEndedPayload) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceSpeakEndedPayload) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceSpeakEndedPayload) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceSpeakEndedPayload) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *ConferenceSpeakEndedPayload) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ConferenceSpeakEndedPayload) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ConferenceSpeakEndedPayload) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *ConferenceSpeakEndedPayload) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


