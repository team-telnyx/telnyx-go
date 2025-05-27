# CreatePushCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of mobile push credential. Should be &lt;code&gt;ios&lt;/code&gt; here | 
**Certificate** | **string** | Certificate as received from APNs | 
**PrivateKey** | **string** | Corresponding private key to the certificate as received from APNs | 
**Alias** | **string** | Alias to uniquely identify the credential | 
**ProjectAccountJsonFile** | **map[string]interface{}** | Private key file in JSON format | 

## Methods

### NewCreatePushCredentialRequest

`func NewCreatePushCredentialRequest(type_ string, certificate string, privateKey string, alias string, projectAccountJsonFile map[string]interface{}, ) *CreatePushCredentialRequest`

NewCreatePushCredentialRequest instantiates a new CreatePushCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePushCredentialRequestWithDefaults

`func NewCreatePushCredentialRequestWithDefaults() *CreatePushCredentialRequest`

NewCreatePushCredentialRequestWithDefaults instantiates a new CreatePushCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreatePushCredentialRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePushCredentialRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePushCredentialRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCertificate

`func (o *CreatePushCredentialRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreatePushCredentialRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreatePushCredentialRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *CreatePushCredentialRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreatePushCredentialRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreatePushCredentialRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetAlias

`func (o *CreatePushCredentialRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreatePushCredentialRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreatePushCredentialRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetProjectAccountJsonFile

`func (o *CreatePushCredentialRequest) GetProjectAccountJsonFile() map[string]interface{}`

GetProjectAccountJsonFile returns the ProjectAccountJsonFile field if non-nil, zero value otherwise.

### GetProjectAccountJsonFileOk

`func (o *CreatePushCredentialRequest) GetProjectAccountJsonFileOk() (*map[string]interface{}, bool)`

GetProjectAccountJsonFileOk returns a tuple with the ProjectAccountJsonFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAccountJsonFile

`func (o *CreatePushCredentialRequest) SetProjectAccountJsonFile(v map[string]interface{})`

SetProjectAccountJsonFile sets ProjectAccountJsonFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


