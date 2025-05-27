# NotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A UUID. | [optional] [readonly] 
**NotificationProfileId** | Pointer to **string** | A UUID reference to the associated Notification Profile. | [optional] 
**ChannelTypeId** | Pointer to **string** | A Channel Type ID | [optional] 
**ChannelDestination** | Pointer to **string** | The destination associated with the channel type. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewNotificationChannel

`func NewNotificationChannel() *NotificationChannel`

NewNotificationChannel instantiates a new NotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelWithDefaults

`func NewNotificationChannelWithDefaults() *NotificationChannel`

NewNotificationChannelWithDefaults instantiates a new NotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotificationProfileId

`func (o *NotificationChannel) GetNotificationProfileId() string`

GetNotificationProfileId returns the NotificationProfileId field if non-nil, zero value otherwise.

### GetNotificationProfileIdOk

`func (o *NotificationChannel) GetNotificationProfileIdOk() (*string, bool)`

GetNotificationProfileIdOk returns a tuple with the NotificationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationProfileId

`func (o *NotificationChannel) SetNotificationProfileId(v string)`

SetNotificationProfileId sets NotificationProfileId field to given value.

### HasNotificationProfileId

`func (o *NotificationChannel) HasNotificationProfileId() bool`

HasNotificationProfileId returns a boolean if a field has been set.

### GetChannelTypeId

`func (o *NotificationChannel) GetChannelTypeId() string`

GetChannelTypeId returns the ChannelTypeId field if non-nil, zero value otherwise.

### GetChannelTypeIdOk

`func (o *NotificationChannel) GetChannelTypeIdOk() (*string, bool)`

GetChannelTypeIdOk returns a tuple with the ChannelTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelTypeId

`func (o *NotificationChannel) SetChannelTypeId(v string)`

SetChannelTypeId sets ChannelTypeId field to given value.

### HasChannelTypeId

`func (o *NotificationChannel) HasChannelTypeId() bool`

HasChannelTypeId returns a boolean if a field has been set.

### GetChannelDestination

`func (o *NotificationChannel) GetChannelDestination() string`

GetChannelDestination returns the ChannelDestination field if non-nil, zero value otherwise.

### GetChannelDestinationOk

`func (o *NotificationChannel) GetChannelDestinationOk() (*string, bool)`

GetChannelDestinationOk returns a tuple with the ChannelDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDestination

`func (o *NotificationChannel) SetChannelDestination(v string)`

SetChannelDestination sets ChannelDestination field to given value.

### HasChannelDestination

`func (o *NotificationChannel) HasChannelDestination() bool`

HasChannelDestination returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationChannel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationChannel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationChannel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationChannel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationChannel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationChannel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationChannel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


