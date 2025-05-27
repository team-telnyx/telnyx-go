# CompleteOTAUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**SimCardId** | Pointer to **string** | The identification UUID of the related SIM card resource. | [optional] 
**Type** | Pointer to **string** | Represents the type of the operation requested. This will relate directly to the source of the request. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**CompleteOTAUpdateSettings**](CompleteOTAUpdateSettings.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewCompleteOTAUpdate

`func NewCompleteOTAUpdate() *CompleteOTAUpdate`

NewCompleteOTAUpdate instantiates a new CompleteOTAUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteOTAUpdateWithDefaults

`func NewCompleteOTAUpdateWithDefaults() *CompleteOTAUpdate`

NewCompleteOTAUpdateWithDefaults instantiates a new CompleteOTAUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompleteOTAUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompleteOTAUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompleteOTAUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompleteOTAUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *CompleteOTAUpdate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CompleteOTAUpdate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CompleteOTAUpdate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CompleteOTAUpdate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSimCardId

`func (o *CompleteOTAUpdate) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *CompleteOTAUpdate) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *CompleteOTAUpdate) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *CompleteOTAUpdate) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetType

`func (o *CompleteOTAUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompleteOTAUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompleteOTAUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompleteOTAUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CompleteOTAUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompleteOTAUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompleteOTAUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompleteOTAUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSettings

`func (o *CompleteOTAUpdate) GetSettings() CompleteOTAUpdateSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CompleteOTAUpdate) GetSettingsOk() (*CompleteOTAUpdateSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CompleteOTAUpdate) SetSettings(v CompleteOTAUpdateSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CompleteOTAUpdate) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CompleteOTAUpdate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompleteOTAUpdate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompleteOTAUpdate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CompleteOTAUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompleteOTAUpdate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompleteOTAUpdate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompleteOTAUpdate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CompleteOTAUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


