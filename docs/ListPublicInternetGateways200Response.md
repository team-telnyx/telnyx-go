# ListPublicInternetGateways200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PublicInternetGatewayRead**](PublicInternetGatewayRead.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListPublicInternetGateways200Response

`func NewListPublicInternetGateways200Response() *ListPublicInternetGateways200Response`

NewListPublicInternetGateways200Response instantiates a new ListPublicInternetGateways200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPublicInternetGateways200ResponseWithDefaults

`func NewListPublicInternetGateways200ResponseWithDefaults() *ListPublicInternetGateways200Response`

NewListPublicInternetGateways200ResponseWithDefaults instantiates a new ListPublicInternetGateways200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPublicInternetGateways200Response) GetData() []PublicInternetGatewayRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPublicInternetGateways200Response) GetDataOk() (*[]PublicInternetGatewayRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPublicInternetGateways200Response) SetData(v []PublicInternetGatewayRead)`

SetData sets Data field to given value.

### HasData

`func (o *ListPublicInternetGateways200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListPublicInternetGateways200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPublicInternetGateways200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPublicInternetGateways200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPublicInternetGateways200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


