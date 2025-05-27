# CallCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of resource. | [optional] 
**EventType** | Pointer to **string** | The type of the event being delivered | [optional] 
**Id** | Pointer to **string** | Uniquely identifies a webhook | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Meta** | Pointer to [**CallCostMeta**](CallCostMeta.md) |  | [optional] 
**Payload** | Pointer to [**CallCostPayload**](CallCostPayload.md) |  | [optional] 

## Methods

### NewCallCost

`func NewCallCost() *CallCost`

NewCallCost instantiates a new CallCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallCostWithDefaults

`func NewCallCostWithDefaults() *CallCost`

NewCallCostWithDefaults instantiates a new CallCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallCost) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallCost) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallCost) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallCost) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *CallCost) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CallCost) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CallCost) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CallCost) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CallCost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallCost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallCost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallCost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CallCost) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CallCost) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CallCost) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CallCost) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetMeta

`func (o *CallCost) GetMeta() CallCostMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CallCost) GetMetaOk() (*CallCostMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CallCost) SetMeta(v CallCostMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CallCost) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPayload

`func (o *CallCost) GetPayload() CallCostPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CallCost) GetPayloadOk() (*CallCostPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CallCost) SetPayload(v CallCostPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CallCost) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


