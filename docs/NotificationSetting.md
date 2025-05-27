# NotificationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A UUID. | [optional] [readonly] 
**NotificationEventConditionId** | Pointer to **string** | A UUID reference to the associated Notification Event Condition. | [optional] 
**NotificationProfileId** | Pointer to **string** | A UUID reference to the associated Notification Profile. | [optional] 
**AssociatedRecordType** | Pointer to **string** |  | [optional] [readonly] 
**AssociatedRecordTypeValue** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** | Most preferences apply immediately; however, other may needs to propagate. | [optional] [readonly] 
**NotificationChannelId** | Pointer to **string** | A UUID reference to the associated Notification Channel. | [optional] 
**Parameters** | Pointer to [**[]NotificationSettingParametersInner**](NotificationSettingParametersInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewNotificationSetting

`func NewNotificationSetting() *NotificationSetting`

NewNotificationSetting instantiates a new NotificationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingWithDefaults

`func NewNotificationSettingWithDefaults() *NotificationSetting`

NewNotificationSettingWithDefaults instantiates a new NotificationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationSetting) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationSetting) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationSetting) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotificationEventConditionId

`func (o *NotificationSetting) GetNotificationEventConditionId() string`

GetNotificationEventConditionId returns the NotificationEventConditionId field if non-nil, zero value otherwise.

### GetNotificationEventConditionIdOk

`func (o *NotificationSetting) GetNotificationEventConditionIdOk() (*string, bool)`

GetNotificationEventConditionIdOk returns a tuple with the NotificationEventConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEventConditionId

`func (o *NotificationSetting) SetNotificationEventConditionId(v string)`

SetNotificationEventConditionId sets NotificationEventConditionId field to given value.

### HasNotificationEventConditionId

`func (o *NotificationSetting) HasNotificationEventConditionId() bool`

HasNotificationEventConditionId returns a boolean if a field has been set.

### GetNotificationProfileId

`func (o *NotificationSetting) GetNotificationProfileId() string`

GetNotificationProfileId returns the NotificationProfileId field if non-nil, zero value otherwise.

### GetNotificationProfileIdOk

`func (o *NotificationSetting) GetNotificationProfileIdOk() (*string, bool)`

GetNotificationProfileIdOk returns a tuple with the NotificationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationProfileId

`func (o *NotificationSetting) SetNotificationProfileId(v string)`

SetNotificationProfileId sets NotificationProfileId field to given value.

### HasNotificationProfileId

`func (o *NotificationSetting) HasNotificationProfileId() bool`

HasNotificationProfileId returns a boolean if a field has been set.

### GetAssociatedRecordType

`func (o *NotificationSetting) GetAssociatedRecordType() string`

GetAssociatedRecordType returns the AssociatedRecordType field if non-nil, zero value otherwise.

### GetAssociatedRecordTypeOk

`func (o *NotificationSetting) GetAssociatedRecordTypeOk() (*string, bool)`

GetAssociatedRecordTypeOk returns a tuple with the AssociatedRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRecordType

`func (o *NotificationSetting) SetAssociatedRecordType(v string)`

SetAssociatedRecordType sets AssociatedRecordType field to given value.

### HasAssociatedRecordType

`func (o *NotificationSetting) HasAssociatedRecordType() bool`

HasAssociatedRecordType returns a boolean if a field has been set.

### GetAssociatedRecordTypeValue

`func (o *NotificationSetting) GetAssociatedRecordTypeValue() string`

GetAssociatedRecordTypeValue returns the AssociatedRecordTypeValue field if non-nil, zero value otherwise.

### GetAssociatedRecordTypeValueOk

`func (o *NotificationSetting) GetAssociatedRecordTypeValueOk() (*string, bool)`

GetAssociatedRecordTypeValueOk returns a tuple with the AssociatedRecordTypeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRecordTypeValue

`func (o *NotificationSetting) SetAssociatedRecordTypeValue(v string)`

SetAssociatedRecordTypeValue sets AssociatedRecordTypeValue field to given value.

### HasAssociatedRecordTypeValue

`func (o *NotificationSetting) HasAssociatedRecordTypeValue() bool`

HasAssociatedRecordTypeValue returns a boolean if a field has been set.

### GetStatus

`func (o *NotificationSetting) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationSetting) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationSetting) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotificationSetting) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotificationChannelId

`func (o *NotificationSetting) GetNotificationChannelId() string`

GetNotificationChannelId returns the NotificationChannelId field if non-nil, zero value otherwise.

### GetNotificationChannelIdOk

`func (o *NotificationSetting) GetNotificationChannelIdOk() (*string, bool)`

GetNotificationChannelIdOk returns a tuple with the NotificationChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationChannelId

`func (o *NotificationSetting) SetNotificationChannelId(v string)`

SetNotificationChannelId sets NotificationChannelId field to given value.

### HasNotificationChannelId

`func (o *NotificationSetting) HasNotificationChannelId() bool`

HasNotificationChannelId returns a boolean if a field has been set.

### GetParameters

`func (o *NotificationSetting) GetParameters() []NotificationSettingParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *NotificationSetting) GetParametersOk() (*[]NotificationSettingParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *NotificationSetting) SetParameters(v []NotificationSettingParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *NotificationSetting) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationSetting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationSetting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationSetting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationSetting) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationSetting) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationSetting) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationSetting) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationSetting) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


