# MigrationSourceParamsProviderAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AWS Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key here. | [optional] 
**SecretAccessKey** | Pointer to **string** | AWS Secret Access Key. For Telnyx-to-Telnyx migrations, use your Telnyx API key here as well. | [optional] 

## Methods

### NewMigrationSourceParamsProviderAuth

`func NewMigrationSourceParamsProviderAuth() *MigrationSourceParamsProviderAuth`

NewMigrationSourceParamsProviderAuth instantiates a new MigrationSourceParamsProviderAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationSourceParamsProviderAuthWithDefaults

`func NewMigrationSourceParamsProviderAuthWithDefaults() *MigrationSourceParamsProviderAuth`

NewMigrationSourceParamsProviderAuthWithDefaults instantiates a new MigrationSourceParamsProviderAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *MigrationSourceParamsProviderAuth) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *MigrationSourceParamsProviderAuth) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *MigrationSourceParamsProviderAuth) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *MigrationSourceParamsProviderAuth) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *MigrationSourceParamsProviderAuth) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *MigrationSourceParamsProviderAuth) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *MigrationSourceParamsProviderAuth) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *MigrationSourceParamsProviderAuth) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


