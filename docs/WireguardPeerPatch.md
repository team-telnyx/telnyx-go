# WireguardPeerPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | The WireGuard &#x60;PublicKey&#x60;.&lt;br /&gt;&lt;br /&gt;If you do not provide a Public Key, a new Public and Private key pair will be generated for you. | [optional] 

## Methods

### NewWireguardPeerPatch

`func NewWireguardPeerPatch() *WireguardPeerPatch`

NewWireguardPeerPatch instantiates a new WireguardPeerPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerPatchWithDefaults

`func NewWireguardPeerPatchWithDefaults() *WireguardPeerPatch`

NewWireguardPeerPatchWithDefaults instantiates a new WireguardPeerPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *WireguardPeerPatch) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardPeerPatch) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardPeerPatch) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardPeerPatch) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


