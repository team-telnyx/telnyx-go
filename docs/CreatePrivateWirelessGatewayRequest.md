# CreatePrivateWirelessGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | The identification of the related network resource. | 
**Name** | **string** | The private wireless gateway name. | 
**RegionCode** | Pointer to **string** | The code of the region where the private wireless gateway will be assigned. A list of available regions can be found at the regions endpoint | [optional] 

## Methods

### NewCreatePrivateWirelessGatewayRequest

`func NewCreatePrivateWirelessGatewayRequest(networkId string, name string, ) *CreatePrivateWirelessGatewayRequest`

NewCreatePrivateWirelessGatewayRequest instantiates a new CreatePrivateWirelessGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrivateWirelessGatewayRequestWithDefaults

`func NewCreatePrivateWirelessGatewayRequestWithDefaults() *CreatePrivateWirelessGatewayRequest`

NewCreatePrivateWirelessGatewayRequestWithDefaults instantiates a new CreatePrivateWirelessGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *CreatePrivateWirelessGatewayRequest) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreatePrivateWirelessGatewayRequest) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreatePrivateWirelessGatewayRequest) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *CreatePrivateWirelessGatewayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePrivateWirelessGatewayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePrivateWirelessGatewayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegionCode

`func (o *CreatePrivateWirelessGatewayRequest) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *CreatePrivateWirelessGatewayRequest) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *CreatePrivateWirelessGatewayRequest) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *CreatePrivateWirelessGatewayRequest) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


