# AccessIPAddressListResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccessIPAddressResponseSchema**](AccessIPAddressResponseSchema.md) |  | 
**Meta** | [**PaginationMeta**](PaginationMeta.md) |  | 

## Methods

### NewAccessIPAddressListResponseSchema

`func NewAccessIPAddressListResponseSchema(data []AccessIPAddressResponseSchema, meta PaginationMeta, ) *AccessIPAddressListResponseSchema`

NewAccessIPAddressListResponseSchema instantiates a new AccessIPAddressListResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessIPAddressListResponseSchemaWithDefaults

`func NewAccessIPAddressListResponseSchemaWithDefaults() *AccessIPAddressListResponseSchema`

NewAccessIPAddressListResponseSchemaWithDefaults instantiates a new AccessIPAddressListResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccessIPAddressListResponseSchema) GetData() []AccessIPAddressResponseSchema`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessIPAddressListResponseSchema) GetDataOk() (*[]AccessIPAddressResponseSchema, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessIPAddressListResponseSchema) SetData(v []AccessIPAddressResponseSchema)`

SetData sets Data field to given value.


### GetMeta

`func (o *AccessIPAddressListResponseSchema) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AccessIPAddressListResponseSchema) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AccessIPAddressListResponseSchema) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


