# WireguardInterfaceCreate

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
**Endpoint** | Pointer to **string** | The Telnyx WireGuard peers &#x60;Peer.endpoint&#x60; value. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The Telnyx WireGuard peers &#x60;Peer.PublicKey&#x60;. | [optional] [readonly] 
**EnableSipTrunking** | Pointer to **bool** | Enable SIP traffic forwarding over VPN interface. | [optional] 
**RegionCode** | **string** | The region the interface should be deployed to. | 

## Methods

### NewWireguardInterfaceCreate

`func NewWireguardInterfaceCreate(networkId string, regionCode string, ) *WireguardInterfaceCreate`

NewWireguardInterfaceCreate instantiates a new WireguardInterfaceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardInterfaceCreateWithDefaults

`func NewWireguardInterfaceCreateWithDefaults() *WireguardInterfaceCreate`

NewWireguardInterfaceCreateWithDefaults instantiates a new WireguardInterfaceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardInterfaceCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardInterfaceCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardInterfaceCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireguardInterfaceCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *WireguardInterfaceCreate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WireguardInterfaceCreate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WireguardInterfaceCreate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WireguardInterfaceCreate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WireguardInterfaceCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WireguardInterfaceCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WireguardInterfaceCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WireguardInterfaceCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WireguardInterfaceCreate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WireguardInterfaceCreate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WireguardInterfaceCreate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WireguardInterfaceCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *WireguardInterfaceCreate) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *WireguardInterfaceCreate) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *WireguardInterfaceCreate) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *WireguardInterfaceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardInterfaceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardInterfaceCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WireguardInterfaceCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WireguardInterfaceCreate) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardInterfaceCreate) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardInterfaceCreate) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WireguardInterfaceCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndpoint

`func (o *WireguardInterfaceCreate) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WireguardInterfaceCreate) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WireguardInterfaceCreate) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WireguardInterfaceCreate) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardInterfaceCreate) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardInterfaceCreate) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardInterfaceCreate) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardInterfaceCreate) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetEnableSipTrunking

`func (o *WireguardInterfaceCreate) GetEnableSipTrunking() bool`

GetEnableSipTrunking returns the EnableSipTrunking field if non-nil, zero value otherwise.

### GetEnableSipTrunkingOk

`func (o *WireguardInterfaceCreate) GetEnableSipTrunkingOk() (*bool, bool)`

GetEnableSipTrunkingOk returns a tuple with the EnableSipTrunking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSipTrunking

`func (o *WireguardInterfaceCreate) SetEnableSipTrunking(v bool)`

SetEnableSipTrunking sets EnableSipTrunking field to given value.

### HasEnableSipTrunking

`func (o *WireguardInterfaceCreate) HasEnableSipTrunking() bool`

HasEnableSipTrunking returns a boolean if a field has been set.

### GetRegionCode

`func (o *WireguardInterfaceCreate) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *WireguardInterfaceCreate) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *WireguardInterfaceCreate) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


