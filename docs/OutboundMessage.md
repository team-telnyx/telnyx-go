# OutboundMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**Payload** | Pointer to [**OutboundMessagePayload**](OutboundMessagePayload.md) |  | [optional] 

## Methods

### NewOutboundMessage

`func NewOutboundMessage() *OutboundMessage`

NewOutboundMessage instantiates a new OutboundMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessageWithDefaults

`func NewOutboundMessageWithDefaults() *OutboundMessage`

NewOutboundMessageWithDefaults instantiates a new OutboundMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *OutboundMessage) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *OutboundMessage) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *OutboundMessage) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *OutboundMessage) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *OutboundMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboundMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboundMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutboundMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *OutboundMessage) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OutboundMessage) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OutboundMessage) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OutboundMessage) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetOccurredAt

`func (o *OutboundMessage) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *OutboundMessage) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *OutboundMessage) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *OutboundMessage) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *OutboundMessage) GetPayload() OutboundMessagePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OutboundMessage) GetPayloadOk() (*OutboundMessagePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OutboundMessage) SetPayload(v OutboundMessagePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *OutboundMessage) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


