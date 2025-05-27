# ListRequirements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DocReqsRequirement**](DocReqsRequirement.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListRequirements200Response

`func NewListRequirements200Response() *ListRequirements200Response`

NewListRequirements200Response instantiates a new ListRequirements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRequirements200ResponseWithDefaults

`func NewListRequirements200ResponseWithDefaults() *ListRequirements200Response`

NewListRequirements200ResponseWithDefaults instantiates a new ListRequirements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRequirements200Response) GetData() []DocReqsRequirement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRequirements200Response) GetDataOk() (*[]DocReqsRequirement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRequirements200Response) SetData(v []DocReqsRequirement)`

SetData sets Data field to given value.

### HasData

`func (o *ListRequirements200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListRequirements200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRequirements200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRequirements200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRequirements200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


