# ConferenceParticipantPlaybackStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**Payload** | Pointer to [**ConferenceParticipantPlaybackEndedPayload**](ConferenceParticipantPlaybackEndedPayload.md) |  | [optional] 

## Methods

### NewConferenceParticipantPlaybackStarted

`func NewConferenceParticipantPlaybackStarted() *ConferenceParticipantPlaybackStarted`

NewConferenceParticipantPlaybackStarted instantiates a new ConferenceParticipantPlaybackStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceParticipantPlaybackStartedWithDefaults

`func NewConferenceParticipantPlaybackStartedWithDefaults() *ConferenceParticipantPlaybackStarted`

NewConferenceParticipantPlaybackStartedWithDefaults instantiates a new ConferenceParticipantPlaybackStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ConferenceParticipantPlaybackStarted) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferenceParticipantPlaybackStarted) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferenceParticipantPlaybackStarted) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ConferenceParticipantPlaybackStarted) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *ConferenceParticipantPlaybackStarted) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ConferenceParticipantPlaybackStarted) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ConferenceParticipantPlaybackStarted) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ConferenceParticipantPlaybackStarted) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *ConferenceParticipantPlaybackStarted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceParticipantPlaybackStarted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceParticipantPlaybackStarted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceParticipantPlaybackStarted) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *ConferenceParticipantPlaybackStarted) GetPayload() ConferenceParticipantPlaybackEndedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ConferenceParticipantPlaybackStarted) GetPayloadOk() (*ConferenceParticipantPlaybackEndedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ConferenceParticipantPlaybackStarted) SetPayload(v ConferenceParticipantPlaybackEndedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ConferenceParticipantPlaybackStarted) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


