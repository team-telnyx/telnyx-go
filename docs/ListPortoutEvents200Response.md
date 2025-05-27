# ListPortoutEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortoutEvent**](PortoutEvent.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListPortoutEvents200Response

`func NewListPortoutEvents200Response() *ListPortoutEvents200Response`

NewListPortoutEvents200Response instantiates a new ListPortoutEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPortoutEvents200ResponseWithDefaults

`func NewListPortoutEvents200ResponseWithDefaults() *ListPortoutEvents200Response`

NewListPortoutEvents200ResponseWithDefaults instantiates a new ListPortoutEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPortoutEvents200Response) GetData() []PortoutEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPortoutEvents200Response) GetDataOk() (*[]PortoutEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPortoutEvents200Response) SetData(v []PortoutEvent)`

SetData sets Data field to given value.

### HasData

`func (o *ListPortoutEvents200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPortoutEvents200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPortoutEvents200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPortoutEvents200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPortoutEvents200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


