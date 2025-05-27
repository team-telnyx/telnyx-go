# CreateAndroidPushCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of mobile push credential. Should be &lt;code&gt;android&lt;/code&gt; here | 
**ProjectAccountJsonFile** | **map[string]interface{}** | Private key file in JSON format | 
**Alias** | **string** | Alias to uniquely identify the credential | 

## Methods

### NewCreateAndroidPushCredentialRequest

`func NewCreateAndroidPushCredentialRequest(type_ string, projectAccountJsonFile map[string]interface{}, alias string, ) *CreateAndroidPushCredentialRequest`

NewCreateAndroidPushCredentialRequest instantiates a new CreateAndroidPushCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAndroidPushCredentialRequestWithDefaults

`func NewCreateAndroidPushCredentialRequestWithDefaults() *CreateAndroidPushCredentialRequest`

NewCreateAndroidPushCredentialRequestWithDefaults instantiates a new CreateAndroidPushCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateAndroidPushCredentialRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAndroidPushCredentialRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAndroidPushCredentialRequest) SetType(v string)`

SetType sets Type field to given value.


### GetProjectAccountJsonFile

`func (o *CreateAndroidPushCredentialRequest) GetProjectAccountJsonFile() map[string]interface{}`

GetProjectAccountJsonFile returns the ProjectAccountJsonFile field if non-nil, zero value otherwise.

### GetProjectAccountJsonFileOk

`func (o *CreateAndroidPushCredentialRequest) GetProjectAccountJsonFileOk() (*map[string]interface{}, bool)`

GetProjectAccountJsonFileOk returns a tuple with the ProjectAccountJsonFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAccountJsonFile

`func (o *CreateAndroidPushCredentialRequest) SetProjectAccountJsonFile(v map[string]interface{})`

SetProjectAccountJsonFile sets ProjectAccountJsonFile field to given value.


### GetAlias

`func (o *CreateAndroidPushCredentialRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateAndroidPushCredentialRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateAndroidPushCredentialRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


