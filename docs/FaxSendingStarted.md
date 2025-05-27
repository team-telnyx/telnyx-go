# FaxSendingStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to [**RecordType**](RecordType.md) |  | [optional] 
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Payload** | Pointer to [**FaxSendingStartedPayload**](FaxSendingStartedPayload.md) |  | [optional] 

## Methods

### NewFaxSendingStarted

`func NewFaxSendingStarted() *FaxSendingStarted`

NewFaxSendingStarted instantiates a new FaxSendingStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxSendingStartedWithDefaults

`func NewFaxSendingStartedWithDefaults() *FaxSendingStarted`

NewFaxSendingStartedWithDefaults instantiates a new FaxSendingStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *FaxSendingStarted) GetRecordType() RecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *FaxSendingStarted) GetRecordTypeOk() (*RecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *FaxSendingStarted) SetRecordType(v RecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *FaxSendingStarted) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *FaxSendingStarted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FaxSendingStarted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FaxSendingStarted) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FaxSendingStarted) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *FaxSendingStarted) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *FaxSendingStarted) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *FaxSendingStarted) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *FaxSendingStarted) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPayload

`func (o *FaxSendingStarted) GetPayload() FaxSendingStartedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FaxSendingStarted) GetPayloadOk() (*FaxSendingStartedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FaxSendingStarted) SetPayload(v FaxSendingStartedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *FaxSendingStarted) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


