# FaxQueued

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to [**RecordType**](RecordType.md) |  | [optional] 
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Payload** | Pointer to [**FaxQueuedPayload**](FaxQueuedPayload.md) |  | [optional] 

## Methods

### NewFaxQueued

`func NewFaxQueued() *FaxQueued`

NewFaxQueued instantiates a new FaxQueued object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxQueuedWithDefaults

`func NewFaxQueuedWithDefaults() *FaxQueued`

NewFaxQueuedWithDefaults instantiates a new FaxQueued object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *FaxQueued) GetRecordType() RecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *FaxQueued) GetRecordTypeOk() (*RecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *FaxQueued) SetRecordType(v RecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *FaxQueued) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *FaxQueued) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FaxQueued) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FaxQueued) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FaxQueued) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *FaxQueued) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *FaxQueued) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *FaxQueued) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *FaxQueued) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPayload

`func (o *FaxQueued) GetPayload() FaxQueuedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FaxQueued) GetPayloadOk() (*FaxQueuedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FaxQueued) SetPayload(v FaxQueuedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *FaxQueued) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


