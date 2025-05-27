# NotificationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A UUID. | [optional] [readonly] 
**Name** | Pointer to **string** | A human readable name. | [optional] 
**NotificationCategory** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewNotificationEvent

`func NewNotificationEvent() *NotificationEvent`

NewNotificationEvent instantiates a new NotificationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventWithDefaults

`func NewNotificationEventWithDefaults() *NotificationEvent`

NewNotificationEventWithDefaults instantiates a new NotificationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationCategory

`func (o *NotificationEvent) GetNotificationCategory() string`

GetNotificationCategory returns the NotificationCategory field if non-nil, zero value otherwise.

### GetNotificationCategoryOk

`func (o *NotificationEvent) GetNotificationCategoryOk() (*string, bool)`

GetNotificationCategoryOk returns a tuple with the NotificationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCategory

`func (o *NotificationEvent) SetNotificationCategory(v string)`

SetNotificationCategory sets NotificationCategory field to given value.

### HasNotificationCategory

`func (o *NotificationEvent) HasNotificationCategory() bool`

HasNotificationCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *NotificationEvent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationEvent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationEvent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationEvent) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


