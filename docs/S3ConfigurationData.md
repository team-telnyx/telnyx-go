# S3ConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Name of the bucket to be used to store recording files. | [optional] 
**Region** | Pointer to [**Region**](Region.md) |  | [optional] 
**AwsAccessKeyId** | Pointer to **string** | AWS credentials access key id. | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | AWS secret access key. | [optional] 

## Methods

### NewS3ConfigurationData

`func NewS3ConfigurationData() *S3ConfigurationData`

NewS3ConfigurationData instantiates a new S3ConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ConfigurationDataWithDefaults

`func NewS3ConfigurationDataWithDefaults() *S3ConfigurationData`

NewS3ConfigurationDataWithDefaults instantiates a new S3ConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *S3ConfigurationData) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *S3ConfigurationData) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *S3ConfigurationData) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *S3ConfigurationData) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetRegion

`func (o *S3ConfigurationData) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3ConfigurationData) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3ConfigurationData) SetRegion(v Region)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *S3ConfigurationData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetAwsAccessKeyId

`func (o *S3ConfigurationData) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *S3ConfigurationData) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *S3ConfigurationData) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *S3ConfigurationData) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *S3ConfigurationData) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *S3ConfigurationData) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *S3ConfigurationData) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *S3ConfigurationData) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


