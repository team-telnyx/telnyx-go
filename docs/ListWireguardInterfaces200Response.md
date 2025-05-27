# ListWireguardInterfaces200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WireguardInterfaceRead**](WireguardInterfaceRead.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListWireguardInterfaces200Response

`func NewListWireguardInterfaces200Response() *ListWireguardInterfaces200Response`

NewListWireguardInterfaces200Response instantiates a new ListWireguardInterfaces200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWireguardInterfaces200ResponseWithDefaults

`func NewListWireguardInterfaces200ResponseWithDefaults() *ListWireguardInterfaces200Response`

NewListWireguardInterfaces200ResponseWithDefaults instantiates a new ListWireguardInterfaces200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListWireguardInterfaces200Response) GetData() []WireguardInterfaceRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWireguardInterfaces200Response) GetDataOk() (*[]WireguardInterfaceRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWireguardInterfaces200Response) SetData(v []WireguardInterfaceRead)`

SetData sets Data field to given value.

### HasData

`func (o *ListWireguardInterfaces200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListWireguardInterfaces200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWireguardInterfaces200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWireguardInterfaces200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWireguardInterfaces200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


