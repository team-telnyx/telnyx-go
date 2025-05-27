# ConferenceSpeakStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**Payload** | Pointer to [**ConferenceSpeakEndedPayload**](ConferenceSpeakEndedPayload.md) |  | [optional] 

## Methods

### NewConferenceSpeakStarted

`func NewConferenceSpeakStarted() *ConferenceSpeakStarted`

NewConferenceSpeakStarted instantiates a new ConferenceSpeakStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceSpeakStartedWithDefaults

`func NewConferenceSpeakStartedWithDefaults() *ConferenceSpeakStarted`

NewConferenceSpeakStartedWithDefaults instantiates a new ConferenceSpeakStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ConferenceSpeakStarted) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferenceSpeakStarted) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferenceSpeakStarted) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ConferenceSpeakStarted) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *ConferenceSpeakStarted) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ConferenceSpeakStarted) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ConferenceSpeakStarted) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ConferenceSpeakStarted) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *ConferenceSpeakStarted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceSpeakStarted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceSpeakStarted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceSpeakStarted) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *ConferenceSpeakStarted) GetPayload() ConferenceSpeakEndedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ConferenceSpeakStarted) GetPayloadOk() (*ConferenceSpeakEndedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ConferenceSpeakStarted) SetPayload(v ConferenceSpeakEndedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ConferenceSpeakStarted) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


