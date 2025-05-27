# ModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "list"]
**Data** | [**[]ModelMetadata**](ModelMetadata.md) |  | 

## Methods

### NewModelsResponse

`func NewModelsResponse(data []ModelMetadata, ) *ModelsResponse`

NewModelsResponse instantiates a new ModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsResponseWithDefaults

`func NewModelsResponseWithDefaults() *ModelsResponse`

NewModelsResponseWithDefaults instantiates a new ModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *ModelsResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelsResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelsResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelsResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *ModelsResponse) GetData() []ModelMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModelsResponse) GetDataOk() (*[]ModelMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModelsResponse) SetData(v []ModelMetadata)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


