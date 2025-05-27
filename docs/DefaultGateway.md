# DefaultGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**NetworkId** | Pointer to **string** | Network ID. | [optional] [readonly] 
**WireguardPeerId** | Pointer to **string** | Wireguard peer ID. | [optional] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 

## Methods

### NewDefaultGateway

`func NewDefaultGateway() *DefaultGateway`

NewDefaultGateway instantiates a new DefaultGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultGatewayWithDefaults

`func NewDefaultGatewayWithDefaults() *DefaultGateway`

NewDefaultGatewayWithDefaults instantiates a new DefaultGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DefaultGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefaultGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefaultGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DefaultGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *DefaultGateway) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DefaultGateway) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DefaultGateway) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DefaultGateway) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DefaultGateway) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DefaultGateway) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DefaultGateway) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DefaultGateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DefaultGateway) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DefaultGateway) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DefaultGateway) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DefaultGateway) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *DefaultGateway) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *DefaultGateway) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *DefaultGateway) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *DefaultGateway) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetWireguardPeerId

`func (o *DefaultGateway) GetWireguardPeerId() string`

GetWireguardPeerId returns the WireguardPeerId field if non-nil, zero value otherwise.

### GetWireguardPeerIdOk

`func (o *DefaultGateway) GetWireguardPeerIdOk() (*string, bool)`

GetWireguardPeerIdOk returns a tuple with the WireguardPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPeerId

`func (o *DefaultGateway) SetWireguardPeerId(v string)`

SetWireguardPeerId sets WireguardPeerId field to given value.

### HasWireguardPeerId

`func (o *DefaultGateway) HasWireguardPeerId() bool`

HasWireguardPeerId returns a boolean if a field has been set.

### GetStatus

`func (o *DefaultGateway) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DefaultGateway) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DefaultGateway) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DefaultGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


