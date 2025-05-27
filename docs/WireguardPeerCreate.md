# WireguardPeerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The WireGuard &#x60;PublicKey&#x60;.&lt;br /&gt;&lt;br /&gt;If you do not provide a Public Key, a new Public and Private key pair will be generated for you. | [optional] 
**LastSeen** | Pointer to **string** | ISO 8601 formatted date-time indicating when peer sent traffic last time. | [optional] [readonly] 
**WireguardInterfaceId** | **string** | The id of the wireguard interface associated with the peer. | 
**PrivateKey** | Pointer to **string** | Your WireGuard &#x60;Interface.PrivateKey&#x60;.&lt;br /&gt;&lt;br /&gt;This attribute is only ever utlised if, on POST, you do NOT provide your own &#x60;public_key&#x60;. In which case, a new Public and Private key pair will be generated for you. When your &#x60;private_key&#x60; is returned, you must save this immediately as we do not save it within Telnyx. If you lose your Private Key, it can not be recovered. | [optional] [readonly] 

## Methods

### NewWireguardPeerCreate

`func NewWireguardPeerCreate(wireguardInterfaceId string, ) *WireguardPeerCreate`

NewWireguardPeerCreate instantiates a new WireguardPeerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerCreateWithDefaults

`func NewWireguardPeerCreateWithDefaults() *WireguardPeerCreate`

NewWireguardPeerCreateWithDefaults instantiates a new WireguardPeerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardPeerCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardPeerCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardPeerCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireguardPeerCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *WireguardPeerCreate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WireguardPeerCreate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WireguardPeerCreate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WireguardPeerCreate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WireguardPeerCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WireguardPeerCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WireguardPeerCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WireguardPeerCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WireguardPeerCreate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WireguardPeerCreate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WireguardPeerCreate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WireguardPeerCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardPeerCreate) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardPeerCreate) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardPeerCreate) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardPeerCreate) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetLastSeen

`func (o *WireguardPeerCreate) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *WireguardPeerCreate) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *WireguardPeerCreate) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *WireguardPeerCreate) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetWireguardInterfaceId

`func (o *WireguardPeerCreate) GetWireguardInterfaceId() string`

GetWireguardInterfaceId returns the WireguardInterfaceId field if non-nil, zero value otherwise.

### GetWireguardInterfaceIdOk

`func (o *WireguardPeerCreate) GetWireguardInterfaceIdOk() (*string, bool)`

GetWireguardInterfaceIdOk returns a tuple with the WireguardInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardInterfaceId

`func (o *WireguardPeerCreate) SetWireguardInterfaceId(v string)`

SetWireguardInterfaceId sets WireguardInterfaceId field to given value.


### GetPrivateKey

`func (o *WireguardPeerCreate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WireguardPeerCreate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WireguardPeerCreate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WireguardPeerCreate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


