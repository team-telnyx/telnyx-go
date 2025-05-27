# SIMCardGroupAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**SimCardGroupId** | Pointer to **string** | The SIM card group identification. | [optional] 
**Type** | Pointer to **string** | Represents the type of the operation requested. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**SIMCardGroupActionSettings**](SIMCardGroupActionSettings.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSIMCardGroupAction

`func NewSIMCardGroupAction() *SIMCardGroupAction`

NewSIMCardGroupAction instantiates a new SIMCardGroupAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardGroupActionWithDefaults

`func NewSIMCardGroupActionWithDefaults() *SIMCardGroupAction`

NewSIMCardGroupActionWithDefaults instantiates a new SIMCardGroupAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SIMCardGroupAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIMCardGroupAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIMCardGroupAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIMCardGroupAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SIMCardGroupAction) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardGroupAction) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardGroupAction) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardGroupAction) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSimCardGroupId

`func (o *SIMCardGroupAction) GetSimCardGroupId() string`

GetSimCardGroupId returns the SimCardGroupId field if non-nil, zero value otherwise.

### GetSimCardGroupIdOk

`func (o *SIMCardGroupAction) GetSimCardGroupIdOk() (*string, bool)`

GetSimCardGroupIdOk returns a tuple with the SimCardGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardGroupId

`func (o *SIMCardGroupAction) SetSimCardGroupId(v string)`

SetSimCardGroupId sets SimCardGroupId field to given value.

### HasSimCardGroupId

`func (o *SIMCardGroupAction) HasSimCardGroupId() bool`

HasSimCardGroupId returns a boolean if a field has been set.

### GetType

`func (o *SIMCardGroupAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SIMCardGroupAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SIMCardGroupAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SIMCardGroupAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SIMCardGroupAction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SIMCardGroupAction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SIMCardGroupAction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SIMCardGroupAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettings

`func (o *SIMCardGroupAction) GetSettings() SIMCardGroupActionSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SIMCardGroupAction) GetSettingsOk() (*SIMCardGroupActionSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SIMCardGroupAction) SetSettings(v SIMCardGroupActionSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SIMCardGroupAction) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SIMCardGroupAction) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SIMCardGroupAction) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SIMCardGroupAction) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SIMCardGroupAction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SIMCardGroupAction) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SIMCardGroupAction) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SIMCardGroupAction) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SIMCardGroupAction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


