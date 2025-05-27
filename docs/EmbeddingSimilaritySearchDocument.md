# EmbeddingSimilaritySearchDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentChunk** | **string** |  | 
**Distance** | **float32** |  | 
**Metadata** | [**EmbeddingMetadata**](EmbeddingMetadata.md) |  | 

## Methods

### NewEmbeddingSimilaritySearchDocument

`func NewEmbeddingSimilaritySearchDocument(documentChunk string, distance float32, metadata EmbeddingMetadata, ) *EmbeddingSimilaritySearchDocument`

NewEmbeddingSimilaritySearchDocument instantiates a new EmbeddingSimilaritySearchDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingSimilaritySearchDocumentWithDefaults

`func NewEmbeddingSimilaritySearchDocumentWithDefaults() *EmbeddingSimilaritySearchDocument`

NewEmbeddingSimilaritySearchDocumentWithDefaults instantiates a new EmbeddingSimilaritySearchDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentChunk

`func (o *EmbeddingSimilaritySearchDocument) GetDocumentChunk() string`

GetDocumentChunk returns the DocumentChunk field if non-nil, zero value otherwise.

### GetDocumentChunkOk

`func (o *EmbeddingSimilaritySearchDocument) GetDocumentChunkOk() (*string, bool)`

GetDocumentChunkOk returns a tuple with the DocumentChunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentChunk

`func (o *EmbeddingSimilaritySearchDocument) SetDocumentChunk(v string)`

SetDocumentChunk sets DocumentChunk field to given value.


### GetDistance

`func (o *EmbeddingSimilaritySearchDocument) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *EmbeddingSimilaritySearchDocument) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *EmbeddingSimilaritySearchDocument) SetDistance(v float32)`

SetDistance sets Distance field to given value.


### GetMetadata

`func (o *EmbeddingSimilaritySearchDocument) GetMetadata() EmbeddingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EmbeddingSimilaritySearchDocument) GetMetadataOk() (*EmbeddingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EmbeddingSimilaritySearchDocument) SetMetadata(v EmbeddingMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


