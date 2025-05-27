# GetSimCardGroupActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SIMCardGroupAction**](SIMCardGroupAction.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewGetSimCardGroupActions200Response

`func NewGetSimCardGroupActions200Response() *GetSimCardGroupActions200Response`

NewGetSimCardGroupActions200Response instantiates a new GetSimCardGroupActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSimCardGroupActions200ResponseWithDefaults

`func NewGetSimCardGroupActions200ResponseWithDefaults() *GetSimCardGroupActions200Response`

NewGetSimCardGroupActions200ResponseWithDefaults instantiates a new GetSimCardGroupActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSimCardGroupActions200Response) GetData() []SIMCardGroupAction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSimCardGroupActions200Response) GetDataOk() (*[]SIMCardGroupAction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSimCardGroupActions200Response) SetData(v []SIMCardGroupAction)`

SetData sets Data field to given value.

### HasData

`func (o *GetSimCardGroupActions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetSimCardGroupActions200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSimCardGroupActions200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSimCardGroupActions200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetSimCardGroupActions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


