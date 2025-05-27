# EmbeddingBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** |  | 
**DocumentChunkSize** | Pointer to **int32** |  | [optional] [default to 1024]
**DocumentChunkOverlapSize** | Pointer to **int32** |  | [optional] [default to 512]
**EmbeddingModel** | Pointer to [**SupportedEmbeddingModels**](SupportedEmbeddingModels.md) |  | [optional] [default to THENLPER_GTE_LARGE]
**Loader** | Pointer to [**SupportedEmbeddingLoaders**](SupportedEmbeddingLoaders.md) |  | [optional] [default to DEFAULT]

## Methods

### NewEmbeddingBucketRequest

`func NewEmbeddingBucketRequest(bucketName string, ) *EmbeddingBucketRequest`

NewEmbeddingBucketRequest instantiates a new EmbeddingBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingBucketRequestWithDefaults

`func NewEmbeddingBucketRequestWithDefaults() *EmbeddingBucketRequest`

NewEmbeddingBucketRequestWithDefaults instantiates a new EmbeddingBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *EmbeddingBucketRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *EmbeddingBucketRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *EmbeddingBucketRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetDocumentChunkSize

`func (o *EmbeddingBucketRequest) GetDocumentChunkSize() int32`

GetDocumentChunkSize returns the DocumentChunkSize field if non-nil, zero value otherwise.

### GetDocumentChunkSizeOk

`func (o *EmbeddingBucketRequest) GetDocumentChunkSizeOk() (*int32, bool)`

GetDocumentChunkSizeOk returns a tuple with the DocumentChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentChunkSize

`func (o *EmbeddingBucketRequest) SetDocumentChunkSize(v int32)`

SetDocumentChunkSize sets DocumentChunkSize field to given value.

### HasDocumentChunkSize

`func (o *EmbeddingBucketRequest) HasDocumentChunkSize() bool`

HasDocumentChunkSize returns a boolean if a field has been set.

### GetDocumentChunkOverlapSize

`func (o *EmbeddingBucketRequest) GetDocumentChunkOverlapSize() int32`

GetDocumentChunkOverlapSize returns the DocumentChunkOverlapSize field if non-nil, zero value otherwise.

### GetDocumentChunkOverlapSizeOk

`func (o *EmbeddingBucketRequest) GetDocumentChunkOverlapSizeOk() (*int32, bool)`

GetDocumentChunkOverlapSizeOk returns a tuple with the DocumentChunkOverlapSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentChunkOverlapSize

`func (o *EmbeddingBucketRequest) SetDocumentChunkOverlapSize(v int32)`

SetDocumentChunkOverlapSize sets DocumentChunkOverlapSize field to given value.

### HasDocumentChunkOverlapSize

`func (o *EmbeddingBucketRequest) HasDocumentChunkOverlapSize() bool`

HasDocumentChunkOverlapSize returns a boolean if a field has been set.

### GetEmbeddingModel

`func (o *EmbeddingBucketRequest) GetEmbeddingModel() SupportedEmbeddingModels`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *EmbeddingBucketRequest) GetEmbeddingModelOk() (*SupportedEmbeddingModels, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *EmbeddingBucketRequest) SetEmbeddingModel(v SupportedEmbeddingModels)`

SetEmbeddingModel sets EmbeddingModel field to given value.

### HasEmbeddingModel

`func (o *EmbeddingBucketRequest) HasEmbeddingModel() bool`

HasEmbeddingModel returns a boolean if a field has been set.

### GetLoader

`func (o *EmbeddingBucketRequest) GetLoader() SupportedEmbeddingLoaders`

GetLoader returns the Loader field if non-nil, zero value otherwise.

### GetLoaderOk

`func (o *EmbeddingBucketRequest) GetLoaderOk() (*SupportedEmbeddingLoaders, bool)`

GetLoaderOk returns a tuple with the Loader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoader

`func (o *EmbeddingBucketRequest) SetLoader(v SupportedEmbeddingLoaders)`

SetLoader sets Loader field to given value.

### HasLoader

`func (o *EmbeddingBucketRequest) HasLoader() bool`

HasLoader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


