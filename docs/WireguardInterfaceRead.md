# WireguardInterfaceRead

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
**Endpoint** | Pointer to **string** | The Telnyx WireGuard peers &#x60;Peer.endpoint&#x60; value. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The Telnyx WireGuard peers &#x60;Peer.PublicKey&#x60;. | [optional] [readonly] 
**EnableSipTrunking** | Pointer to **bool** | Enable SIP traffic forwarding over VPN interface. | [optional] 
**RegionCode** | Pointer to **string** | The region interface is deployed to. | [optional] 
**Region** | Pointer to [**RegionOutRegion**](RegionOutRegion.md) |  | [optional] 

## Methods

### NewWireguardInterfaceRead

`func NewWireguardInterfaceRead() *WireguardInterfaceRead`

NewWireguardInterfaceRead instantiates a new WireguardInterfaceRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardInterfaceReadWithDefaults

`func NewWireguardInterfaceReadWithDefaults() *WireguardInterfaceRead`

NewWireguardInterfaceReadWithDefaults instantiates a new WireguardInterfaceRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardInterfaceRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardInterfaceRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardInterfaceRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireguardInterfaceRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *WireguardInterfaceRead) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WireguardInterfaceRead) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WireguardInterfaceRead) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WireguardInterfaceRead) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WireguardInterfaceRead) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WireguardInterfaceRead) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WireguardInterfaceRead) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WireguardInterfaceRead) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WireguardInterfaceRead) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WireguardInterfaceRead) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WireguardInterfaceRead) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WireguardInterfaceRead) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *WireguardInterfaceRead) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *WireguardInterfaceRead) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *WireguardInterfaceRead) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *WireguardInterfaceRead) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *WireguardInterfaceRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardInterfaceRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardInterfaceRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WireguardInterfaceRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WireguardInterfaceRead) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardInterfaceRead) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardInterfaceRead) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WireguardInterfaceRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEndpoint

`func (o *WireguardInterfaceRead) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WireguardInterfaceRead) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WireguardInterfaceRead) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WireguardInterfaceRead) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardInterfaceRead) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardInterfaceRead) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardInterfaceRead) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardInterfaceRead) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetEnableSipTrunking

`func (o *WireguardInterfaceRead) GetEnableSipTrunking() bool`

GetEnableSipTrunking returns the EnableSipTrunking field if non-nil, zero value otherwise.

### GetEnableSipTrunkingOk

`func (o *WireguardInterfaceRead) GetEnableSipTrunkingOk() (*bool, bool)`

GetEnableSipTrunkingOk returns a tuple with the EnableSipTrunking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSipTrunking

`func (o *WireguardInterfaceRead) SetEnableSipTrunking(v bool)`

SetEnableSipTrunking sets EnableSipTrunking field to given value.

### HasEnableSipTrunking

`func (o *WireguardInterfaceRead) HasEnableSipTrunking() bool`

HasEnableSipTrunking returns a boolean if a field has been set.

### GetRegionCode

`func (o *WireguardInterfaceRead) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *WireguardInterfaceRead) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *WireguardInterfaceRead) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *WireguardInterfaceRead) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetRegion

`func (o *WireguardInterfaceRead) GetRegion() RegionOutRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *WireguardInterfaceRead) GetRegionOk() (*RegionOutRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *WireguardInterfaceRead) SetRegion(v RegionOutRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *WireguardInterfaceRead) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


