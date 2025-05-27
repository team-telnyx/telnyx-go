# AutorespConfigsResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AutorespConfigSchema**](AutorespConfigSchema.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewAutorespConfigsResponseSchema

`func NewAutorespConfigsResponseSchema(data []AutorespConfigSchema, meta PaginationMeta, ) *AutorespConfigsResponseSchema`

NewAutorespConfigsResponseSchema instantiates a new AutorespConfigsResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutorespConfigsResponseSchemaWithDefaults

`func NewAutorespConfigsResponseSchemaWithDefaults() *AutorespConfigsResponseSchema`

NewAutorespConfigsResponseSchemaWithDefaults instantiates a new AutorespConfigsResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AutorespConfigsResponseSchema) GetData() []AutorespConfigSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutorespConfigsResponseSchema) GetDataOk() (*[]AutorespConfigSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutorespConfigsResponseSchema) SetData(v []AutorespConfigSchema)`

SetData sets Data field to given value.


### GetMeta

`func (o *AutorespConfigsResponseSchema) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AutorespConfigsResponseSchema) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AutorespConfigsResponseSchema) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


