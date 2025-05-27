# CallAIGatherEnded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Payload** | Pointer to [**CallAIGatherEndedPayload**](CallAIGatherEndedPayload.md) |  | [optional] 

## Methods

### NewCallAIGatherEnded

`func NewCallAIGatherEnded() *CallAIGatherEnded`

NewCallAIGatherEnded instantiates a new CallAIGatherEnded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallAIGatherEndedWithDefaults

`func NewCallAIGatherEndedWithDefaults() *CallAIGatherEnded`

NewCallAIGatherEndedWithDefaults instantiates a new CallAIGatherEnded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallAIGatherEnded) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallAIGatherEnded) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallAIGatherEnded) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallAIGatherEnded) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *CallAIGatherEnded) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CallAIGatherEnded) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CallAIGatherEnded) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CallAIGatherEnded) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CallAIGatherEnded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallAIGatherEnded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallAIGatherEnded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallAIGatherEnded) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CallAIGatherEnded) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CallAIGatherEnded) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CallAIGatherEnded) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CallAIGatherEnded) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *CallAIGatherEnded) GetPayload() CallAIGatherEndedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CallAIGatherEnded) GetPayloadOk() (*CallAIGatherEndedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CallAIGatherEnded) SetPayload(v CallAIGatherEndedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CallAIGatherEnded) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


