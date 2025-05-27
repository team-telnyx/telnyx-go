# AzureConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Name of the bucket to be used to store recording files. | [optional] 
**AccountName** | Pointer to **string** | Azure Blob Storage account name. | [optional] 
**AccountKey** | Pointer to **string** | Azure Blob Storage account key. | [optional] 

## Methods

### NewAzureConfigurationData

`func NewAzureConfigurationData() *AzureConfigurationData`

NewAzureConfigurationData instantiates a new AzureConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConfigurationDataWithDefaults

`func NewAzureConfigurationDataWithDefaults() *AzureConfigurationData`

NewAzureConfigurationDataWithDefaults instantiates a new AzureConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *AzureConfigurationData) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *AzureConfigurationData) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *AzureConfigurationData) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *AzureConfigurationData) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetAccountName

`func (o *AzureConfigurationData) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AzureConfigurationData) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AzureConfigurationData) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AzureConfigurationData) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountKey

`func (o *AzureConfigurationData) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AzureConfigurationData) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AzureConfigurationData) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *AzureConfigurationData) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


