# CallDtmfReceived

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Payload** | Pointer to [**CallDtmfReceivedPayload**](CallDtmfReceivedPayload.md) |  | [optional] 

## Methods

### NewCallDtmfReceived

`func NewCallDtmfReceived() *CallDtmfReceived`

NewCallDtmfReceived instantiates a new CallDtmfReceived object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallDtmfReceivedWithDefaults

`func NewCallDtmfReceivedWithDefaults() *CallDtmfReceived`

NewCallDtmfReceivedWithDefaults instantiates a new CallDtmfReceived object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallDtmfReceived) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallDtmfReceived) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallDtmfReceived) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallDtmfReceived) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *CallDtmfReceived) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CallDtmfReceived) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CallDtmfReceived) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CallDtmfReceived) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CallDtmfReceived) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallDtmfReceived) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallDtmfReceived) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallDtmfReceived) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CallDtmfReceived) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CallDtmfReceived) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CallDtmfReceived) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CallDtmfReceived) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *CallDtmfReceived) GetPayload() CallDtmfReceivedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CallDtmfReceived) GetPayloadOk() (*CallDtmfReceivedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CallDtmfReceived) SetPayload(v CallDtmfReceivedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CallDtmfReceived) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


