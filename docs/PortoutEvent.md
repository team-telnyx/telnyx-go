# PortoutEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the event. | [optional] 
**EventType** | Pointer to **string** | Identifies the event type | [optional] 
**PortoutId** | Pointer to **string** | Identifies the port-out order associated with the event. | [optional] 
**AvailableNotificationMethods** | Pointer to **[]string** | Indicates the notification methods used. | [optional] 
**PayloadStatus** | Pointer to **string** | The status of the payload generation. | [optional] 
**Payload** | Pointer to [**PortoutEventPayload**](PortoutEventPayload.md) |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewPortoutEvent

`func NewPortoutEvent() *PortoutEvent`

NewPortoutEvent instantiates a new PortoutEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortoutEventWithDefaults

`func NewPortoutEventWithDefaults() *PortoutEvent`

NewPortoutEventWithDefaults instantiates a new PortoutEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortoutEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortoutEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortoutEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortoutEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEventType

`func (o *PortoutEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PortoutEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PortoutEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PortoutEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetPortoutId

`func (o *PortoutEvent) GetPortoutId() string`

GetPortoutId returns the PortoutId field if non-nil, zero value otherwise.

### GetPortoutIdOk

`func (o *PortoutEvent) GetPortoutIdOk() (*string, bool)`

GetPortoutIdOk returns a tuple with the PortoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortoutId

`func (o *PortoutEvent) SetPortoutId(v string)`

SetPortoutId sets PortoutId field to given value.

### HasPortoutId

`func (o *PortoutEvent) HasPortoutId() bool`

HasPortoutId returns a boolean if a field has been set.

### GetAvailableNotificationMethods

`func (o *PortoutEvent) GetAvailableNotificationMethods() []string`

GetAvailableNotificationMethods returns the AvailableNotificationMethods field if non-nil, zero value otherwise.

### GetAvailableNotificationMethodsOk

`func (o *PortoutEvent) GetAvailableNotificationMethodsOk() (*[]string, bool)`

GetAvailableNotificationMethodsOk returns a tuple with the AvailableNotificationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNotificationMethods

`func (o *PortoutEvent) SetAvailableNotificationMethods(v []string)`

SetAvailableNotificationMethods sets AvailableNotificationMethods field to given value.

### HasAvailableNotificationMethods

`func (o *PortoutEvent) HasAvailableNotificationMethods() bool`

HasAvailableNotificationMethods returns a boolean if a field has been set.

### GetPayloadStatus

`func (o *PortoutEvent) GetPayloadStatus() string`

GetPayloadStatus returns the PayloadStatus field if non-nil, zero value otherwise.

### GetPayloadStatusOk

`func (o *PortoutEvent) GetPayloadStatusOk() (*string, bool)`

GetPayloadStatusOk returns a tuple with the PayloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadStatus

`func (o *PortoutEvent) SetPayloadStatus(v string)`

SetPayloadStatus sets PayloadStatus field to given value.

### HasPayloadStatus

`func (o *PortoutEvent) HasPayloadStatus() bool`

HasPayloadStatus returns a boolean if a field has been set.

### GetPayload

`func (o *PortoutEvent) GetPayload() PortoutEventPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *PortoutEvent) GetPayloadOk() (*PortoutEventPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *PortoutEvent) SetPayload(v PortoutEventPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *PortoutEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRecordType

`func (o *PortoutEvent) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortoutEvent) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortoutEvent) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortoutEvent) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortoutEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortoutEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortoutEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortoutEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortoutEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortoutEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortoutEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortoutEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


