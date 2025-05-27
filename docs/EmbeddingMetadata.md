# EmbeddingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** |  | 
**Checksum** | **string** |  | 
**Embedding** | **string** |  | 
**Filename** | **string** |  | 
**Certainty** | Pointer to **float32** |  | [optional] 
**LoaderMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEmbeddingMetadata

`func NewEmbeddingMetadata(source string, checksum string, embedding string, filename string, ) *EmbeddingMetadata`

NewEmbeddingMetadata instantiates a new EmbeddingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingMetadataWithDefaults

`func NewEmbeddingMetadataWithDefaults() *EmbeddingMetadata`

NewEmbeddingMetadataWithDefaults instantiates a new EmbeddingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *EmbeddingMetadata) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EmbeddingMetadata) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EmbeddingMetadata) SetSource(v string)`

SetSource sets Source field to given value.


### GetChecksum

`func (o *EmbeddingMetadata) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *EmbeddingMetadata) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *EmbeddingMetadata) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetEmbedding

`func (o *EmbeddingMetadata) GetEmbedding() string`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *EmbeddingMetadata) GetEmbeddingOk() (*string, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *EmbeddingMetadata) SetEmbedding(v string)`

SetEmbedding sets Embedding field to given value.


### GetFilename

`func (o *EmbeddingMetadata) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EmbeddingMetadata) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EmbeddingMetadata) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetCertainty

`func (o *EmbeddingMetadata) GetCertainty() float32`

GetCertainty returns the Certainty field if non-nil, zero value otherwise.

### GetCertaintyOk

`func (o *EmbeddingMetadata) GetCertaintyOk() (*float32, bool)`

GetCertaintyOk returns a tuple with the Certainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertainty

`func (o *EmbeddingMetadata) SetCertainty(v float32)`

SetCertainty sets Certainty field to given value.

### HasCertainty

`func (o *EmbeddingMetadata) HasCertainty() bool`

HasCertainty returns a boolean if a field has been set.

### GetLoaderMetadata

`func (o *EmbeddingMetadata) GetLoaderMetadata() map[string]interface{}`

GetLoaderMetadata returns the LoaderMetadata field if non-nil, zero value otherwise.

### GetLoaderMetadataOk

`func (o *EmbeddingMetadata) GetLoaderMetadataOk() (*map[string]interface{}, bool)`

GetLoaderMetadataOk returns a tuple with the LoaderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoaderMetadata

`func (o *EmbeddingMetadata) SetLoaderMetadata(v map[string]interface{})`

SetLoaderMetadata sets LoaderMetadata field to given value.

### HasLoaderMetadata

`func (o *EmbeddingMetadata) HasLoaderMetadata() bool`

HasLoaderMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


