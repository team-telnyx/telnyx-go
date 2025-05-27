# CallSiprecFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Payload** | Pointer to [**CallSiprecFailedPayload**](CallSiprecFailedPayload.md) |  | [optional] 

## Methods

### NewCallSiprecFailed

`func NewCallSiprecFailed() *CallSiprecFailed`

NewCallSiprecFailed instantiates a new CallSiprecFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallSiprecFailedWithDefaults

`func NewCallSiprecFailedWithDefaults() *CallSiprecFailed`

NewCallSiprecFailedWithDefaults instantiates a new CallSiprecFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallSiprecFailed) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallSiprecFailed) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallSiprecFailed) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallSiprecFailed) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *CallSiprecFailed) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CallSiprecFailed) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CallSiprecFailed) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CallSiprecFailed) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CallSiprecFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallSiprecFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallSiprecFailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallSiprecFailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CallSiprecFailed) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CallSiprecFailed) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CallSiprecFailed) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CallSiprecFailed) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *CallSiprecFailed) GetPayload() CallSiprecFailedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CallSiprecFailed) GetPayloadOk() (*CallSiprecFailedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CallSiprecFailed) SetPayload(v CallSiprecFailedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CallSiprecFailed) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


