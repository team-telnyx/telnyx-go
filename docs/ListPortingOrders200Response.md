# ListPortingOrders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PortingOrder**](PortingOrder.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListPortingOrders200Response

`func NewListPortingOrders200Response() *ListPortingOrders200Response`

NewListPortingOrders200Response instantiates a new ListPortingOrders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPortingOrders200ResponseWithDefaults

`func NewListPortingOrders200ResponseWithDefaults() *ListPortingOrders200Response`

NewListPortingOrders200ResponseWithDefaults instantiates a new ListPortingOrders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPortingOrders200Response) GetData() []PortingOrder`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPortingOrders200Response) GetDataOk() (*[]PortingOrder, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPortingOrders200Response) SetData(v []PortingOrder)`

SetData sets Data field to given value.

### HasData

`func (o *ListPortingOrders200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPortingOrders200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPortingOrders200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPortingOrders200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPortingOrders200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


