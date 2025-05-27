# WireguardPeer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The WireGuard &#x60;PublicKey&#x60;.&lt;br /&gt;&lt;br /&gt;If you do not provide a Public Key, a new Public and Private key pair will be generated for you. | [optional] 
**LastSeen** | Pointer to **string** | ISO 8601 formatted date-time indicating when peer sent traffic last time. | [optional] [readonly] 
**WireguardInterfaceId** | Pointer to **string** | The id of the wireguard interface associated with the peer. | [optional] 
**PrivateKey** | Pointer to **string** | Your WireGuard &#x60;Interface.PrivateKey&#x60;.&lt;br /&gt;&lt;br /&gt;This attribute is only ever utlised if, on POST, you do NOT provide your own &#x60;public_key&#x60;. In which case, a new Public and Private key pair will be generated for you. When your &#x60;private_key&#x60; is returned, you must save this immediately as we do not save it within Telnyx. If you lose your Private Key, it can not be recovered. | [optional] [readonly] 

## Methods

### NewWireguardPeer

`func NewWireguardPeer() *WireguardPeer`

NewWireguardPeer instantiates a new WireguardPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerWithDefaults

`func NewWireguardPeerWithDefaults() *WireguardPeer`

NewWireguardPeerWithDefaults instantiates a new WireguardPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardPeer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardPeer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardPeer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireguardPeer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *WireguardPeer) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WireguardPeer) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WireguardPeer) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WireguardPeer) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WireguardPeer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WireguardPeer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WireguardPeer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WireguardPeer) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WireguardPeer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WireguardPeer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WireguardPeer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WireguardPeer) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardPeer) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardPeer) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardPeer) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardPeer) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetLastSeen

`func (o *WireguardPeer) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *WireguardPeer) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *WireguardPeer) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *WireguardPeer) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetWireguardInterfaceId

`func (o *WireguardPeer) GetWireguardInterfaceId() string`

GetWireguardInterfaceId returns the WireguardInterfaceId field if non-nil, zero value otherwise.

### GetWireguardInterfaceIdOk

`func (o *WireguardPeer) GetWireguardInterfaceIdOk() (*string, bool)`

GetWireguardInterfaceIdOk returns a tuple with the WireguardInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardInterfaceId

`func (o *WireguardPeer) SetWireguardInterfaceId(v string)`

SetWireguardInterfaceId sets WireguardInterfaceId field to given value.

### HasWireguardInterfaceId

`func (o *WireguardPeer) HasWireguardInterfaceId() bool`

HasWireguardInterfaceId returns a boolean if a field has been set.

### GetPrivateKey

`func (o *WireguardPeer) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WireguardPeer) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WireguardPeer) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WireguardPeer) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


