# ListRequirementTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DocReqsRequirementType**](DocReqsRequirementType.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListRequirementTypes200Response

`func NewListRequirementTypes200Response() *ListRequirementTypes200Response`

NewListRequirementTypes200Response instantiates a new ListRequirementTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRequirementTypes200ResponseWithDefaults

`func NewListRequirementTypes200ResponseWithDefaults() *ListRequirementTypes200Response`

NewListRequirementTypes200ResponseWithDefaults instantiates a new ListRequirementTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRequirementTypes200Response) GetData() []DocReqsRequirementType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRequirementTypes200Response) GetDataOk() (*[]DocReqsRequirementType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRequirementTypes200Response) SetData(v []DocReqsRequirementType)`

SetData sets Data field to given value.

### HasData

`func (o *ListRequirementTypes200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListRequirementTypes200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRequirementTypes200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRequirementTypes200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRequirementTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


