# SearchedSIMCardGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Default** | Pointer to **bool** | Indicates whether the SIM card group is the users default group.&lt;br/&gt;The default group is created for the user and can not be removed. | [optional] [readonly] 
**Name** | Pointer to **string** | A user friendly name for the SIM card group. | [optional] 
**DataLimit** | Pointer to [**SIMCardGroupDataLimit**](SIMCardGroupDataLimit.md) |  | [optional] 
**ConsumedData** | Pointer to [**ConsumedData**](ConsumedData.md) |  | [optional] 
**SimCardCount** | Pointer to **int32** | The number of SIM cards associated with the group. | [optional] [default to 0]
**PrivateWirelessGatewayId** | Pointer to **string** | The identification of the related Private Wireless Gateway resource. | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSearchedSIMCardGroup

`func NewSearchedSIMCardGroup() *SearchedSIMCardGroup`

NewSearchedSIMCardGroup instantiates a new SearchedSIMCardGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchedSIMCardGroupWithDefaults

`func NewSearchedSIMCardGroupWithDefaults() *SearchedSIMCardGroup`

NewSearchedSIMCardGroupWithDefaults instantiates a new SearchedSIMCardGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchedSIMCardGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchedSIMCardGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchedSIMCardGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchedSIMCardGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SearchedSIMCardGroup) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SearchedSIMCardGroup) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SearchedSIMCardGroup) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SearchedSIMCardGroup) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDefault

`func (o *SearchedSIMCardGroup) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SearchedSIMCardGroup) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SearchedSIMCardGroup) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SearchedSIMCardGroup) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetName

`func (o *SearchedSIMCardGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchedSIMCardGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchedSIMCardGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchedSIMCardGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataLimit

`func (o *SearchedSIMCardGroup) GetDataLimit() SIMCardGroupDataLimit`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *SearchedSIMCardGroup) GetDataLimitOk() (*SIMCardGroupDataLimit, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *SearchedSIMCardGroup) SetDataLimit(v SIMCardGroupDataLimit)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *SearchedSIMCardGroup) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.

### GetConsumedData

`func (o *SearchedSIMCardGroup) GetConsumedData() ConsumedData`

GetConsumedData returns the ConsumedData field if non-nil, zero value otherwise.

### GetConsumedDataOk

`func (o *SearchedSIMCardGroup) GetConsumedDataOk() (*ConsumedData, bool)`

GetConsumedDataOk returns a tuple with the ConsumedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedData

`func (o *SearchedSIMCardGroup) SetConsumedData(v ConsumedData)`

SetConsumedData sets ConsumedData field to given value.

### HasConsumedData

`func (o *SearchedSIMCardGroup) HasConsumedData() bool`

HasConsumedData returns a boolean if a field has been set.

### GetSimCardCount

`func (o *SearchedSIMCardGroup) GetSimCardCount() int32`

GetSimCardCount returns the SimCardCount field if non-nil, zero value otherwise.

### GetSimCardCountOk

`func (o *SearchedSIMCardGroup) GetSimCardCountOk() (*int32, bool)`

GetSimCardCountOk returns a tuple with the SimCardCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardCount

`func (o *SearchedSIMCardGroup) SetSimCardCount(v int32)`

SetSimCardCount sets SimCardCount field to given value.

### HasSimCardCount

`func (o *SearchedSIMCardGroup) HasSimCardCount() bool`

HasSimCardCount returns a boolean if a field has been set.

### GetPrivateWirelessGatewayId

`func (o *SearchedSIMCardGroup) GetPrivateWirelessGatewayId() string`

GetPrivateWirelessGatewayId returns the PrivateWirelessGatewayId field if non-nil, zero value otherwise.

### GetPrivateWirelessGatewayIdOk

`func (o *SearchedSIMCardGroup) GetPrivateWirelessGatewayIdOk() (*string, bool)`

GetPrivateWirelessGatewayIdOk returns a tuple with the PrivateWirelessGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateWirelessGatewayId

`func (o *SearchedSIMCardGroup) SetPrivateWirelessGatewayId(v string)`

SetPrivateWirelessGatewayId sets PrivateWirelessGatewayId field to given value.

### HasPrivateWirelessGatewayId

`func (o *SearchedSIMCardGroup) HasPrivateWirelessGatewayId() bool`

HasPrivateWirelessGatewayId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SearchedSIMCardGroup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchedSIMCardGroup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchedSIMCardGroup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SearchedSIMCardGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SearchedSIMCardGroup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SearchedSIMCardGroup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SearchedSIMCardGroup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SearchedSIMCardGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


