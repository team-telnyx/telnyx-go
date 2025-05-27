# CreateIntegrationSecretRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The unique identifier of the secret. | 
**Type** | **string** | The type of secret. | 
**Token** | Pointer to **string** | The token for the secret. Required for bearer type secrets, ignored otherwise. | [optional] 
**Username** | Pointer to **string** | The username for the secret. Required for basic type secrets, ignored otherwise. | [optional] 
**Password** | Pointer to **string** | The password for the secret. Required for basic type secrets, ignored otherwise. | [optional] 

## Methods

### NewCreateIntegrationSecretRequest

`func NewCreateIntegrationSecretRequest(identifier string, type_ string, ) *CreateIntegrationSecretRequest`

NewCreateIntegrationSecretRequest instantiates a new CreateIntegrationSecretRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationSecretRequestWithDefaults

`func NewCreateIntegrationSecretRequestWithDefaults() *CreateIntegrationSecretRequest`

NewCreateIntegrationSecretRequestWithDefaults instantiates a new CreateIntegrationSecretRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *CreateIntegrationSecretRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateIntegrationSecretRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateIntegrationSecretRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetType

`func (o *CreateIntegrationSecretRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIntegrationSecretRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIntegrationSecretRequest) SetType(v string)`

SetType sets Type field to given value.


### GetToken

`func (o *CreateIntegrationSecretRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateIntegrationSecretRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateIntegrationSecretRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateIntegrationSecretRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateIntegrationSecretRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateIntegrationSecretRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateIntegrationSecretRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateIntegrationSecretRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CreateIntegrationSecretRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateIntegrationSecretRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateIntegrationSecretRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateIntegrationSecretRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


