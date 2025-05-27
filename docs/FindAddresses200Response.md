# FindAddresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Address**](Address.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewFindAddresses200Response

`func NewFindAddresses200Response() *FindAddresses200Response`

NewFindAddresses200Response instantiates a new FindAddresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAddresses200ResponseWithDefaults

`func NewFindAddresses200ResponseWithDefaults() *FindAddresses200Response`

NewFindAddresses200ResponseWithDefaults instantiates a new FindAddresses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAddresses200Response) GetData() []Address`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAddresses200Response) GetDataOk() (*[]Address, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAddresses200Response) SetData(v []Address)`

SetData sets Data field to given value.

### HasData

`func (o *FindAddresses200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *FindAddresses200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAddresses200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAddresses200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FindAddresses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


