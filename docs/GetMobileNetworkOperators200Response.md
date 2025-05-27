# GetMobileNetworkOperators200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MobileNetworkOperator**](MobileNetworkOperator.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewGetMobileNetworkOperators200Response

`func NewGetMobileNetworkOperators200Response() *GetMobileNetworkOperators200Response`

NewGetMobileNetworkOperators200Response instantiates a new GetMobileNetworkOperators200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMobileNetworkOperators200ResponseWithDefaults

`func NewGetMobileNetworkOperators200ResponseWithDefaults() *GetMobileNetworkOperators200Response`

NewGetMobileNetworkOperators200ResponseWithDefaults instantiates a new GetMobileNetworkOperators200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMobileNetworkOperators200Response) GetData() []MobileNetworkOperator`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMobileNetworkOperators200Response) GetDataOk() (*[]MobileNetworkOperator, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMobileNetworkOperators200Response) SetData(v []MobileNetworkOperator)`

SetData sets Data field to given value.

### HasData

`func (o *GetMobileNetworkOperators200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetMobileNetworkOperators200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMobileNetworkOperators200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMobileNetworkOperators200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetMobileNetworkOperators200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


