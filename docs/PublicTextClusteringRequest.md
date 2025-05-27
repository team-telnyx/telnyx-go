# PublicTextClusteringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | The embedded storage bucket to compute the clusters from. The bucket must already be [embedded](https://developers.telnyx.com/api/inference/inference-embedding/post-embedding). | 
**Prefix** | Pointer to **string** | Prefix to filter whcih files in the buckets are included. | [optional] 
**Files** | Pointer to **[]string** | Array of files to filter which are included. | [optional] 
**MinClusterSize** | Pointer to **int32** | Smallest number of related text chunks to qualify as a cluster. Top-level clusters should be thought of as identifying broad themes in your data. | [optional] [default to 25]
**MinSubclusterSize** | Pointer to **int32** | Smallest number of related text chunks to qualify as a sub-cluster. Sub-clusters should be thought of as identifying more specific topics within a broader theme. | [optional] [default to 5]

## Methods

### NewPublicTextClusteringRequest

`func NewPublicTextClusteringRequest(bucket string, ) *PublicTextClusteringRequest`

NewPublicTextClusteringRequest instantiates a new PublicTextClusteringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTextClusteringRequestWithDefaults

`func NewPublicTextClusteringRequestWithDefaults() *PublicTextClusteringRequest`

NewPublicTextClusteringRequestWithDefaults instantiates a new PublicTextClusteringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *PublicTextClusteringRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *PublicTextClusteringRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *PublicTextClusteringRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPrefix

`func (o *PublicTextClusteringRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PublicTextClusteringRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PublicTextClusteringRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PublicTextClusteringRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetFiles

`func (o *PublicTextClusteringRequest) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PublicTextClusteringRequest) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PublicTextClusteringRequest) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *PublicTextClusteringRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetMinClusterSize

`func (o *PublicTextClusteringRequest) GetMinClusterSize() int32`

GetMinClusterSize returns the MinClusterSize field if non-nil, zero value otherwise.

### GetMinClusterSizeOk

`func (o *PublicTextClusteringRequest) GetMinClusterSizeOk() (*int32, bool)`

GetMinClusterSizeOk returns a tuple with the MinClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinClusterSize

`func (o *PublicTextClusteringRequest) SetMinClusterSize(v int32)`

SetMinClusterSize sets MinClusterSize field to given value.

### HasMinClusterSize

`func (o *PublicTextClusteringRequest) HasMinClusterSize() bool`

HasMinClusterSize returns a boolean if a field has been set.

### GetMinSubclusterSize

`func (o *PublicTextClusteringRequest) GetMinSubclusterSize() int32`

GetMinSubclusterSize returns the MinSubclusterSize field if non-nil, zero value otherwise.

### GetMinSubclusterSizeOk

`func (o *PublicTextClusteringRequest) GetMinSubclusterSizeOk() (*int32, bool)`

GetMinSubclusterSizeOk returns a tuple with the MinSubclusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSubclusterSize

`func (o *PublicTextClusteringRequest) SetMinSubclusterSize(v int32)`

SetMinSubclusterSize sets MinSubclusterSize field to given value.

### HasMinSubclusterSize

`func (o *PublicTextClusteringRequest) HasMinSubclusterSize() bool`

HasMinSubclusterSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


