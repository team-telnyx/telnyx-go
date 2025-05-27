# SIMCardGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Default** | Pointer to **bool** | Indicates whether the SIM card group is the users default group.&lt;br/&gt;The default group is created for the user and can not be removed. | [optional] [readonly] 
**Name** | Pointer to **string** | A user friendly name for the SIM card group. | [optional] 
**DataLimit** | Pointer to [**SIMCardGroupDataLimit**](SIMCardGroupDataLimit.md) |  | [optional] 
**ConsumedData** | Pointer to [**ConsumedData**](ConsumedData.md) |  | [optional] 
**PrivateWirelessGatewayId** | Pointer to **string** | The identification of the related Private Wireless Gateway resource. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSIMCardGroup

`func NewSIMCardGroup() *SIMCardGroup`

NewSIMCardGroup instantiates a new SIMCardGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardGroupWithDefaults

`func NewSIMCardGroupWithDefaults() *SIMCardGroup`

NewSIMCardGroupWithDefaults instantiates a new SIMCardGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SIMCardGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIMCardGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIMCardGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIMCardGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SIMCardGroup) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardGroup) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardGroup) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardGroup) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDefault

`func (o *SIMCardGroup) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SIMCardGroup) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SIMCardGroup) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SIMCardGroup) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *SIMCardGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SIMCardGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SIMCardGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SIMCardGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataLimit

`func (o *SIMCardGroup) GetDataLimit() SIMCardGroupDataLimit`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SIMCardGroup) GetDataLimitOk() (*SIMCardGroupDataLimit, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SIMCardGroup) SetDataLimit(v SIMCardGroupDataLimit)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SIMCardGroup) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.

### GetConsumedData

`func (o *SIMCardGroup) GetConsumedData() ConsumedData`

GetConsumedData returns the ConsumedData field if non-nil, zero value otherwise.

### GetConsumedDataOk

`func (o *SIMCardGroup) GetConsumedDataOk() (*ConsumedData, bool)`

GetConsumedDataOk returns a tuple with the ConsumedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedData

`func (o *SIMCardGroup) SetConsumedData(v ConsumedData)`

SetConsumedData sets ConsumedData field to given value.

### HasConsumedData

`func (o *SIMCardGroup) HasConsumedData() bool`

HasConsumedData returns a boolean if a field has been set.

### GetPrivateWirelessGatewayId

`func (o *SIMCardGroup) GetPrivateWirelessGatewayId() string`

GetPrivateWirelessGatewayId returns the PrivateWirelessGatewayId field if non-nil, zero value otherwise.

### GetPrivateWirelessGatewayIdOk

`func (o *SIMCardGroup) GetPrivateWirelessGatewayIdOk() (*string, bool)`

GetPrivateWirelessGatewayIdOk returns a tuple with the PrivateWirelessGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateWirelessGatewayId

`func (o *SIMCardGroup) SetPrivateWirelessGatewayId(v string)`

SetPrivateWirelessGatewayId sets PrivateWirelessGatewayId field to given value.

### HasPrivateWirelessGatewayId

`func (o *SIMCardGroup) HasPrivateWirelessGatewayId() bool`

HasPrivateWirelessGatewayId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SIMCardGroup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SIMCardGroup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SIMCardGroup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SIMCardGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SIMCardGroup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SIMCardGroup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SIMCardGroup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SIMCardGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


