# EmbeddingSimilaritySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** |  | 
**Query** | **string** |  | 
**NumOfDocs** | Pointer to **int32** |  | [optional] [default to 3]

## Methods

### NewEmbeddingSimilaritySearchRequest

`func NewEmbeddingSimilaritySearchRequest(bucketName string, query string, ) *EmbeddingSimilaritySearchRequest`

NewEmbeddingSimilaritySearchRequest instantiates a new EmbeddingSimilaritySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingSimilaritySearchRequestWithDefaults

`func NewEmbeddingSimilaritySearchRequestWithDefaults() *EmbeddingSimilaritySearchRequest`

NewEmbeddingSimilaritySearchRequestWithDefaults instantiates a new EmbeddingSimilaritySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *EmbeddingSimilaritySearchRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *EmbeddingSimilaritySearchRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *EmbeddingSimilaritySearchRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetQuery

`func (o *EmbeddingSimilaritySearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EmbeddingSimilaritySearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EmbeddingSimilaritySearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetNumOfDocs

`func (o *EmbeddingSimilaritySearchRequest) GetNumOfDocs() int32`

GetNumOfDocs returns the NumOfDocs field if non-nil, zero value otherwise.

### GetNumOfDocsOk

`func (o *EmbeddingSimilaritySearchRequest) GetNumOfDocsOk() (*int32, bool)`

GetNumOfDocsOk returns a tuple with the NumOfDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfDocs

`func (o *EmbeddingSimilaritySearchRequest) SetNumOfDocs(v int32)`

SetNumOfDocs sets NumOfDocs field to given value.

### HasNumOfDocs

`func (o *EmbeddingSimilaritySearchRequest) HasNumOfDocs() bool`

HasNumOfDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


