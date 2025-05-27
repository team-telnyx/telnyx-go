# NotificationEventCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A UUID. | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**NotificationEventId** | Pointer to **string** |  | [optional] 
**AssociatedRecordType** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to [**[]NotificationEventConditionParametersInner**](NotificationEventConditionParametersInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AllowMultipleChannels** | Pointer to **bool** | Dictates whether a notification channel id needs to be provided when creating a notficiation setting. | [optional] 
**Asynchronous** | Pointer to **bool** | Dictates whether a notification setting will take effect immediately. | [optional] 
**SupportedChannels** | Pointer to **[]string** | Dictates the supported notification channel types that can be emitted. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewNotificationEventCondition

`func NewNotificationEventCondition() *NotificationEventCondition`

NewNotificationEventCondition instantiates a new NotificationEventCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationEventConditionWithDefaults

`func NewNotificationEventConditionWithDefaults() *NotificationEventCondition`

NewNotificationEventConditionWithDefaults instantiates a new NotificationEventCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationEventCondition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationEventCondition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationEventCondition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationEventCondition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationEventCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationEventCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationEventCondition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationEventCondition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationEventCondition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationEventCondition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationEventCondition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationEventCondition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotificationEventId

`func (o *NotificationEventCondition) GetNotificationEventId() string`

GetNotificationEventId returns the NotificationEventId field if non-nil, zero value otherwise.

### GetNotificationEventIdOk

`func (o *NotificationEventCondition) GetNotificationEventIdOk() (*string, bool)`

GetNotificationEventIdOk returns a tuple with the NotificationEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEventId

`func (o *NotificationEventCondition) SetNotificationEventId(v string)`

SetNotificationEventId sets NotificationEventId field to given value.

### HasNotificationEventId

`func (o *NotificationEventCondition) HasNotificationEventId() bool`

HasNotificationEventId returns a boolean if a field has been set.

### GetAssociatedRecordType

`func (o *NotificationEventCondition) GetAssociatedRecordType() string`

GetAssociatedRecordType returns the AssociatedRecordType field if non-nil, zero value otherwise.

### GetAssociatedRecordTypeOk

`func (o *NotificationEventCondition) GetAssociatedRecordTypeOk() (*string, bool)`

GetAssociatedRecordTypeOk returns a tuple with the AssociatedRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRecordType

`func (o *NotificationEventCondition) SetAssociatedRecordType(v string)`

SetAssociatedRecordType sets AssociatedRecordType field to given value.

### HasAssociatedRecordType

`func (o *NotificationEventCondition) HasAssociatedRecordType() bool`

HasAssociatedRecordType returns a boolean if a field has been set.

### GetParameters

`func (o *NotificationEventCondition) GetParameters() []NotificationEventConditionParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *NotificationEventCondition) GetParametersOk() (*[]NotificationEventConditionParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *NotificationEventCondition) SetParameters(v []NotificationEventConditionParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *NotificationEventCondition) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEnabled

`func (o *NotificationEventCondition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationEventCondition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationEventCondition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationEventCondition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowMultipleChannels

`func (o *NotificationEventCondition) GetAllowMultipleChannels() bool`

GetAllowMultipleChannels returns the AllowMultipleChannels field if non-nil, zero value otherwise.

### GetAllowMultipleChannelsOk

`func (o *NotificationEventCondition) GetAllowMultipleChannelsOk() (*bool, bool)`

GetAllowMultipleChannelsOk returns a tuple with the AllowMultipleChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleChannels

`func (o *NotificationEventCondition) SetAllowMultipleChannels(v bool)`

SetAllowMultipleChannels sets AllowMultipleChannels field to given value.

### HasAllowMultipleChannels

`func (o *NotificationEventCondition) HasAllowMultipleChannels() bool`

HasAllowMultipleChannels returns a boolean if a field has been set.

### GetAsynchronous

`func (o *NotificationEventCondition) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *NotificationEventCondition) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *NotificationEventCondition) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *NotificationEventCondition) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetSupportedChannels

`func (o *NotificationEventCondition) GetSupportedChannels() []string`

GetSupportedChannels returns the SupportedChannels field if non-nil, zero value otherwise.

### GetSupportedChannelsOk

`func (o *NotificationEventCondition) GetSupportedChannelsOk() (*[]string, bool)`

GetSupportedChannelsOk returns a tuple with the SupportedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedChannels

`func (o *NotificationEventCondition) SetSupportedChannels(v []string)`

SetSupportedChannels sets SupportedChannels field to given value.

### HasSupportedChannels

`func (o *NotificationEventCondition) HasSupportedChannels() bool`

HasSupportedChannels returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationEventCondition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationEventCondition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationEventCondition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationEventCondition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationEventCondition) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationEventCondition) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationEventCondition) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationEventCondition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


