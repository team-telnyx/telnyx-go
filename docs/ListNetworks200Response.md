# ListNetworks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListNetworks200Response

`func NewListNetworks200Response() *ListNetworks200Response`

NewListNetworks200Response instantiates a new ListNetworks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworks200ResponseWithDefaults

`func NewListNetworks200ResponseWithDefaults() *ListNetworks200Response`

NewListNetworks200ResponseWithDefaults instantiates a new ListNetworks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListNetworks200Response) GetData() []Network`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListNetworks200Response) GetDataOk() (*[]Network, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListNetworks200Response) SetData(v []Network)`

SetData sets Data field to given value.

### HasData

`func (o *ListNetworks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworks200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworks200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworks200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


