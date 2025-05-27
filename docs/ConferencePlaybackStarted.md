# ConferencePlaybackStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**Payload** | Pointer to [**ConferencePlaybackEndedPayload**](ConferencePlaybackEndedPayload.md) |  | [optional] 

## Methods

### NewConferencePlaybackStarted

`func NewConferencePlaybackStarted() *ConferencePlaybackStarted`

NewConferencePlaybackStarted instantiates a new ConferencePlaybackStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferencePlaybackStartedWithDefaults

`func NewConferencePlaybackStartedWithDefaults() *ConferencePlaybackStarted`

NewConferencePlaybackStartedWithDefaults instantiates a new ConferencePlaybackStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ConferencePlaybackStarted) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferencePlaybackStarted) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferencePlaybackStarted) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ConferencePlaybackStarted) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *ConferencePlaybackStarted) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ConferencePlaybackStarted) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ConferencePlaybackStarted) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ConferencePlaybackStarted) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *ConferencePlaybackStarted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferencePlaybackStarted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferencePlaybackStarted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferencePlaybackStarted) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *ConferencePlaybackStarted) GetPayload() ConferencePlaybackEndedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ConferencePlaybackStarted) GetPayloadOk() (*ConferencePlaybackEndedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ConferencePlaybackStarted) SetPayload(v ConferencePlaybackEndedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ConferencePlaybackStarted) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


