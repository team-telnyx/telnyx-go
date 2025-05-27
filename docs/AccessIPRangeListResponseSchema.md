# AccessIPRangeListResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccessIPRangeResponseSchema**](AccessIPRangeResponseSchema.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewAccessIPRangeListResponseSchema

`func NewAccessIPRangeListResponseSchema(data []AccessIPRangeResponseSchema, meta PaginationMeta, ) *AccessIPRangeListResponseSchema`

NewAccessIPRangeListResponseSchema instantiates a new AccessIPRangeListResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessIPRangeListResponseSchemaWithDefaults

`func NewAccessIPRangeListResponseSchemaWithDefaults() *AccessIPRangeListResponseSchema`

NewAccessIPRangeListResponseSchemaWithDefaults instantiates a new AccessIPRangeListResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccessIPRangeListResponseSchema) GetData() []AccessIPRangeResponseSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessIPRangeListResponseSchema) GetDataOk() (*[]AccessIPRangeResponseSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessIPRangeListResponseSchema) SetData(v []AccessIPRangeResponseSchema)`

SetData sets Data field to given value.


### GetMeta

`func (o *AccessIPRangeListResponseSchema) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AccessIPRangeListResponseSchema) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AccessIPRangeListResponseSchema) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


