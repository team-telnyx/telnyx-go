# GCSConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to **string** | Opaque credential token used to authenticate and authorize with storage provider. | [optional] 
**Bucket** | Pointer to **string** | Name of the bucket to be used to store recording files. | [optional] 

## Methods

### NewGCSConfigurationData

`func NewGCSConfigurationData() *GCSConfigurationData`

NewGCSConfigurationData instantiates a new GCSConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCSConfigurationDataWithDefaults

`func NewGCSConfigurationDataWithDefaults() *GCSConfigurationData`

NewGCSConfigurationDataWithDefaults instantiates a new GCSConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *GCSConfigurationData) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GCSConfigurationData) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GCSConfigurationData) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GCSConfigurationData) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetBucket

`func (o *GCSConfigurationData) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *GCSConfigurationData) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *GCSConfigurationData) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *GCSConfigurationData) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


