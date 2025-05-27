# EmbeddingUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL of the webpage to embed | 
**BucketName** | **string** | Name of the bucket to store the embeddings. This bucket must already exist. | 

## Methods

### NewEmbeddingUrlRequest

`func NewEmbeddingUrlRequest(url string, bucketName string, ) *EmbeddingUrlRequest`

NewEmbeddingUrlRequest instantiates a new EmbeddingUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingUrlRequestWithDefaults

`func NewEmbeddingUrlRequestWithDefaults() *EmbeddingUrlRequest`

NewEmbeddingUrlRequestWithDefaults instantiates a new EmbeddingUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *EmbeddingUrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmbeddingUrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmbeddingUrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBucketName

`func (o *EmbeddingUrlRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *EmbeddingUrlRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *EmbeddingUrlRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


