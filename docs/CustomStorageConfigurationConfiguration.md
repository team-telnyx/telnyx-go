# CustomStorageConfigurationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | Opaque credential token used to authenticate and authorize with storage provider. | [optional] 
**Bucket** | Pointer to **string** | Name of the bucket to be used to store recording files. | [optional] 
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**AwsAccessKeyId** | Pointer to **string** | AWS credentials access key id. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | AWS secret access key. | [optional] 
**AccountName** | Pointer to **string** | Azure Blob Storage account name. | [optional] 
**AccountKey** | Pointer to **string** | Azure Blob Storage account key. | [optional] 

## Methods

### NewCustomStorageConfigurationConfiguration

`func NewCustomStorageConfigurationConfiguration() *CustomStorageConfigurationConfiguration`

NewCustomStorageConfigurationConfiguration instantiates a new CustomStorageConfigurationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomStorageConfigurationConfigurationWithDefaults

`func NewCustomStorageConfigurationConfigurationWithDefaults() *CustomStorageConfigurationConfiguration`

NewCustomStorageConfigurationConfigurationWithDefaults instantiates a new CustomStorageConfigurationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CustomStorageConfigurationConfiguration) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CustomStorageConfigurationConfiguration) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CustomStorageConfigurationConfiguration) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CustomStorageConfigurationConfiguration) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBucket

`func (o *CustomStorageConfigurationConfiguration) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *CustomStorageConfigurationConfiguration) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *CustomStorageConfigurationConfiguration) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *CustomStorageConfigurationConfiguration) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetRegion

`func (o *CustomStorageConfigurationConfiguration) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CustomStorageConfigurationConfiguration) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CustomStorageConfigurationConfiguration) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CustomStorageConfigurationConfiguration) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *CustomStorageConfigurationConfiguration) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *CustomStorageConfigurationConfiguration) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *CustomStorageConfigurationConfiguration) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *CustomStorageConfigurationConfiguration) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *CustomStorageConfigurationConfiguration) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *CustomStorageConfigurationConfiguration) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *CustomStorageConfigurationConfiguration) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *CustomStorageConfigurationConfiguration) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAccountName

`func (o *CustomStorageConfigurationConfiguration) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CustomStorageConfigurationConfiguration) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CustomStorageConfigurationConfiguration) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *CustomStorageConfigurationConfiguration) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountKey

`func (o *CustomStorageConfigurationConfiguration) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *CustomStorageConfigurationConfiguration) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *CustomStorageConfigurationConfiguration) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *CustomStorageConfigurationConfiguration) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


