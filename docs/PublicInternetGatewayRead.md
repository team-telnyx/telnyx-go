# PublicInternetGatewayRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**NetworkId** | Pointer to **string** | The id of the network associated with the interface. | [optional] 
**Name** | Pointer to **string** | A user specified name for the interface. | [optional] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 
**PublicIp** | Pointer to **string** | The publically accessible ip for this interface. | [optional] [readonly] 
**RegionCode** | Pointer to **string** | The region interface is deployed to. | [optional] 
**Region** | Pointer to [**RegionOutRegion**](RegionOutRegion.md) |  | [optional] 

## Methods

### NewPublicInternetGatewayRead

`func NewPublicInternetGatewayRead() *PublicInternetGatewayRead`

NewPublicInternetGatewayRead instantiates a new PublicInternetGatewayRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicInternetGatewayReadWithDefaults

`func NewPublicInternetGatewayReadWithDefaults() *PublicInternetGatewayRead`

NewPublicInternetGatewayReadWithDefaults instantiates a new PublicInternetGatewayRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicInternetGatewayRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicInternetGatewayRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicInternetGatewayRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicInternetGatewayRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *PublicInternetGatewayRead) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PublicInternetGatewayRead) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PublicInternetGatewayRead) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PublicInternetGatewayRead) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicInternetGatewayRead) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicInternetGatewayRead) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicInternetGatewayRead) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicInternetGatewayRead) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicInternetGatewayRead) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicInternetGatewayRead) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicInternetGatewayRead) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicInternetGatewayRead) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *PublicInternetGatewayRead) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *PublicInternetGatewayRead) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *PublicInternetGatewayRead) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *PublicInternetGatewayRead) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *PublicInternetGatewayRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicInternetGatewayRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicInternetGatewayRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicInternetGatewayRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PublicInternetGatewayRead) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicInternetGatewayRead) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicInternetGatewayRead) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicInternetGatewayRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPublicIp

`func (o *PublicInternetGatewayRead) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *PublicInternetGatewayRead) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *PublicInternetGatewayRead) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *PublicInternetGatewayRead) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegionCode

`func (o *PublicInternetGatewayRead) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *PublicInternetGatewayRead) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *PublicInternetGatewayRead) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *PublicInternetGatewayRead) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegion

`func (o *PublicInternetGatewayRead) GetRegion() RegionOutRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PublicInternetGatewayRead) GetRegionOk() (*RegionOutRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PublicInternetGatewayRead) SetRegion(v RegionOutRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PublicInternetGatewayRead) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


