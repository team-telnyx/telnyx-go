# FaxFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to [**RecordType**](RecordType.md) |  | [optional] 
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Payload** | Pointer to [**FaxFailedPayload**](FaxFailedPayload.md) |  | [optional] 

## Methods

### NewFaxFailed

`func NewFaxFailed() *FaxFailed`

NewFaxFailed instantiates a new FaxFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxFailedWithDefaults

`func NewFaxFailedWithDefaults() *FaxFailed`

NewFaxFailedWithDefaults instantiates a new FaxFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *FaxFailed) GetRecordType() RecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *FaxFailed) GetRecordTypeOk() (*RecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *FaxFailed) SetRecordType(v RecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *FaxFailed) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *FaxFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FaxFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FaxFailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FaxFailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *FaxFailed) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *FaxFailed) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *FaxFailed) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *FaxFailed) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPayload

`func (o *FaxFailed) GetPayload() FaxFailedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FaxFailed) GetPayloadOk() (*FaxFailedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FaxFailed) SetPayload(v FaxFailedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *FaxFailed) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


