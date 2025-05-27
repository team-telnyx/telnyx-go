# SimplifiedOTAUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**SimCardId** | Pointer to **string** | The identification UUID of the related SIM card resource. | [optional] 
**Type** | Pointer to **string** | Represents the type of the operation requested. This will relate directly to the source of the request. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSimplifiedOTAUpdate

`func NewSimplifiedOTAUpdate() *SimplifiedOTAUpdate`

NewSimplifiedOTAUpdate instantiates a new SimplifiedOTAUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedOTAUpdateWithDefaults

`func NewSimplifiedOTAUpdateWithDefaults() *SimplifiedOTAUpdate`

NewSimplifiedOTAUpdateWithDefaults instantiates a new SimplifiedOTAUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimplifiedOTAUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedOTAUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedOTAUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimplifiedOTAUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SimplifiedOTAUpdate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SimplifiedOTAUpdate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SimplifiedOTAUpdate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SimplifiedOTAUpdate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSimCardId

`func (o *SimplifiedOTAUpdate) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *SimplifiedOTAUpdate) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *SimplifiedOTAUpdate) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *SimplifiedOTAUpdate) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedOTAUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedOTAUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedOTAUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedOTAUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SimplifiedOTAUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimplifiedOTAUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimplifiedOTAUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimplifiedOTAUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SimplifiedOTAUpdate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SimplifiedOTAUpdate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SimplifiedOTAUpdate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SimplifiedOTAUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SimplifiedOTAUpdate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SimplifiedOTAUpdate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SimplifiedOTAUpdate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SimplifiedOTAUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


