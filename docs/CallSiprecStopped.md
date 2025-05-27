# CallSiprecStopped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Payload** | Pointer to [**CallSiprecStoppedPayload**](CallSiprecStoppedPayload.md) |  | [optional] 

## Methods

### NewCallSiprecStopped

`func NewCallSiprecStopped() *CallSiprecStopped`

NewCallSiprecStopped instantiates a new CallSiprecStopped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallSiprecStoppedWithDefaults

`func NewCallSiprecStoppedWithDefaults() *CallSiprecStopped`

NewCallSiprecStoppedWithDefaults instantiates a new CallSiprecStopped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallSiprecStopped) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallSiprecStopped) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallSiprecStopped) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallSiprecStopped) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *CallSiprecStopped) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CallSiprecStopped) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CallSiprecStopped) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CallSiprecStopped) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CallSiprecStopped) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallSiprecStopped) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallSiprecStopped) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallSiprecStopped) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CallSiprecStopped) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CallSiprecStopped) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CallSiprecStopped) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CallSiprecStopped) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *CallSiprecStopped) GetPayload() CallSiprecStoppedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CallSiprecStopped) GetPayloadOk() (*CallSiprecStoppedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CallSiprecStopped) SetPayload(v CallSiprecStoppedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CallSiprecStopped) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


