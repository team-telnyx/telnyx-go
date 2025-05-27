# CustomerServiceRecordStatusChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the callback event. | [optional] 
**EventType** | Pointer to **string** | The type of the callback event. | [optional] 
**Payload** | Pointer to [**CustomerServiceRecordStatusChangedEventPayload**](CustomerServiceRecordStatusChangedEventPayload.md) |  | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the callback event occurred. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 

## Methods

### NewCustomerServiceRecordStatusChangedEvent

`func NewCustomerServiceRecordStatusChangedEvent() *CustomerServiceRecordStatusChangedEvent`

NewCustomerServiceRecordStatusChangedEvent instantiates a new CustomerServiceRecordStatusChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceRecordStatusChangedEventWithDefaults

`func NewCustomerServiceRecordStatusChangedEventWithDefaults() *CustomerServiceRecordStatusChangedEvent`

NewCustomerServiceRecordStatusChangedEventWithDefaults instantiates a new CustomerServiceRecordStatusChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerServiceRecordStatusChangedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerServiceRecordStatusChangedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerServiceRecordStatusChangedEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerServiceRecordStatusChangedEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *CustomerServiceRecordStatusChangedEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CustomerServiceRecordStatusChangedEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CustomerServiceRecordStatusChangedEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CustomerServiceRecordStatusChangedEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPayload

`func (o *CustomerServiceRecordStatusChangedEvent) GetPayload() CustomerServiceRecordStatusChangedEventPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomerServiceRecordStatusChangedEvent) GetPayloadOk() (*CustomerServiceRecordStatusChangedEventPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomerServiceRecordStatusChangedEvent) SetPayload(v CustomerServiceRecordStatusChangedEventPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CustomerServiceRecordStatusChangedEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CustomerServiceRecordStatusChangedEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CustomerServiceRecordStatusChangedEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CustomerServiceRecordStatusChangedEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CustomerServiceRecordStatusChangedEvent) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetRecordType

`func (o *CustomerServiceRecordStatusChangedEvent) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CustomerServiceRecordStatusChangedEvent) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CustomerServiceRecordStatusChangedEvent) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CustomerServiceRecordStatusChangedEvent) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


