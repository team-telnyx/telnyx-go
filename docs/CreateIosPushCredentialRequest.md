# CreateIosPushCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of mobile push credential. Should be &lt;code&gt;ios&lt;/code&gt; here | 
**Certificate** | **string** | Certificate as received from APNs | 
**PrivateKey** | **string** | Corresponding private key to the certificate as received from APNs | 
**Alias** | **string** | Alias to uniquely identify the credential | 

## Methods

### NewCreateIosPushCredentialRequest

`func NewCreateIosPushCredentialRequest(type_ string, certificate string, privateKey string, alias string, ) *CreateIosPushCredentialRequest`

NewCreateIosPushCredentialRequest instantiates a new CreateIosPushCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIosPushCredentialRequestWithDefaults

`func NewCreateIosPushCredentialRequestWithDefaults() *CreateIosPushCredentialRequest`

NewCreateIosPushCredentialRequestWithDefaults instantiates a new CreateIosPushCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateIosPushCredentialRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIosPushCredentialRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIosPushCredentialRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCertificate

`func (o *CreateIosPushCredentialRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateIosPushCredentialRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateIosPushCredentialRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *CreateIosPushCredentialRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateIosPushCredentialRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateIosPushCredentialRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetAlias

`func (o *CreateIosPushCredentialRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateIosPushCredentialRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateIosPushCredentialRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


