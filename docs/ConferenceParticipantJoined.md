# ConferenceParticipantJoined

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Payload** | Pointer to [**ConferenceParticipantJoinedPayload**](ConferenceParticipantJoinedPayload.md) |  | [optional] 

## Methods

### NewConferenceParticipantJoined

`func NewConferenceParticipantJoined() *ConferenceParticipantJoined`

NewConferenceParticipantJoined instantiates a new ConferenceParticipantJoined object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceParticipantJoinedWithDefaults

`func NewConferenceParticipantJoinedWithDefaults() *ConferenceParticipantJoined`

NewConferenceParticipantJoinedWithDefaults instantiates a new ConferenceParticipantJoined object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ConferenceParticipantJoined) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferenceParticipantJoined) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferenceParticipantJoined) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ConferenceParticipantJoined) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *ConferenceParticipantJoined) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ConferenceParticipantJoined) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ConferenceParticipantJoined) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ConferenceParticipantJoined) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *ConferenceParticipantJoined) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceParticipantJoined) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceParticipantJoined) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceParticipantJoined) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *ConferenceParticipantJoined) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ConferenceParticipantJoined) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ConferenceParticipantJoined) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *ConferenceParticipantJoined) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *ConferenceParticipantJoined) GetPayload() ConferenceParticipantJoinedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ConferenceParticipantJoined) GetPayloadOk() (*ConferenceParticipantJoinedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ConferenceParticipantJoined) SetPayload(v ConferenceParticipantJoinedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ConferenceParticipantJoined) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


