# PortingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the event. | [optional] 
**EventType** | Pointer to **string** | Identifies the event type | [optional] 
**PortingOrderId** | Pointer to **string** | Identifies the porting order associated with the event. | [optional] 
**AvailableNotificationMethods** | Pointer to **[]string** | Indicates the notification methods used. | [optional] 
**PayloadStatus** | Pointer to **string** | The status of the payload generation. | [optional] 
**Payload** | Pointer to [**PortingEventPayload**](PortingEventPayload.md) |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewPortingEvent

`func NewPortingEvent() *PortingEvent`

NewPortingEvent instantiates a new PortingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingEventWithDefaults

`func NewPortingEventWithDefaults() *PortingEvent`

NewPortingEventWithDefaults instantiates a new PortingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *PortingEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PortingEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PortingEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PortingEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPortingOrderId

`func (o *PortingEvent) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *PortingEvent) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *PortingEvent) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *PortingEvent) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetAvailableNotificationMethods

`func (o *PortingEvent) GetAvailableNotificationMethods() []string`

GetAvailableNotificationMethods returns the AvailableNotificationMethods field if non-nil, zero value otherwise.

### GetAvailableNotificationMethodsOk

`func (o *PortingEvent) GetAvailableNotificationMethodsOk() (*[]string, bool)`

GetAvailableNotificationMethodsOk returns a tuple with the AvailableNotificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNotificationMethods

`func (o *PortingEvent) SetAvailableNotificationMethods(v []string)`

SetAvailableNotificationMethods sets AvailableNotificationMethods field to given value.

### HasAvailableNotificationMethods

`func (o *PortingEvent) HasAvailableNotificationMethods() bool`

HasAvailableNotificationMethods returns a boolean if a field has been set.

### GetPayloadStatus

`func (o *PortingEvent) GetPayloadStatus() string`

GetPayloadStatus returns the PayloadStatus field if non-nil, zero value otherwise.

### GetPayloadStatusOk

`func (o *PortingEvent) GetPayloadStatusOk() (*string, bool)`

GetPayloadStatusOk returns a tuple with the PayloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadStatus

`func (o *PortingEvent) SetPayloadStatus(v string)`

SetPayloadStatus sets PayloadStatus field to given value.

### HasPayloadStatus

`func (o *PortingEvent) HasPayloadStatus() bool`

HasPayloadStatus returns a boolean if a field has been set.

### GetPayload

`func (o *PortingEvent) GetPayload() PortingEventPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PortingEvent) GetPayloadOk() (*PortingEventPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PortingEvent) SetPayload(v PortingEventPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PortingEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingEvent) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingEvent) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingEvent) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingEvent) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


