# PublicInternetGatewayCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**NetworkId** | **string** | The id of the network associated with the interface. | 
**Name** | Pointer to **string** | A user specified name for the interface. | [optional] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 
**PublicIp** | Pointer to **string** | The publically accessible ip for this interface. | [optional] [readonly] 
**RegionCode** | Pointer to **string** | The region the interface should be deployed to. | [optional] 

## Methods

### NewPublicInternetGatewayCreate

`func NewPublicInternetGatewayCreate(networkId string, ) *PublicInternetGatewayCreate`

NewPublicInternetGatewayCreate instantiates a new PublicInternetGatewayCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicInternetGatewayCreateWithDefaults

`func NewPublicInternetGatewayCreateWithDefaults() *PublicInternetGatewayCreate`

NewPublicInternetGatewayCreateWithDefaults instantiates a new PublicInternetGatewayCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublicInternetGatewayCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicInternetGatewayCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicInternetGatewayCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicInternetGatewayCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *PublicInternetGatewayCreate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PublicInternetGatewayCreate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PublicInternetGatewayCreate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PublicInternetGatewayCreate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PublicInternetGatewayCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicInternetGatewayCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicInternetGatewayCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicInternetGatewayCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicInternetGatewayCreate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicInternetGatewayCreate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicInternetGatewayCreate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicInternetGatewayCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *PublicInternetGatewayCreate) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *PublicInternetGatewayCreate) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *PublicInternetGatewayCreate) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *PublicInternetGatewayCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicInternetGatewayCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicInternetGatewayCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicInternetGatewayCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PublicInternetGatewayCreate) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicInternetGatewayCreate) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicInternetGatewayCreate) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PublicInternetGatewayCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPublicIp

`func (o *PublicInternetGatewayCreate) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *PublicInternetGatewayCreate) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *PublicInternetGatewayCreate) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *PublicInternetGatewayCreate) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetRegionCode

`func (o *PublicInternetGatewayCreate) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *PublicInternetGatewayCreate) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *PublicInternetGatewayCreate) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *PublicInternetGatewayCreate) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


