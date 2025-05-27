# ListFQDNConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FqdnConnection**](FqdnConnection.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListFQDNConnectionsResponse

`func NewListFQDNConnectionsResponse() *ListFQDNConnectionsResponse`

NewListFQDNConnectionsResponse instantiates a new ListFQDNConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFQDNConnectionsResponseWithDefaults

`func NewListFQDNConnectionsResponseWithDefaults() *ListFQDNConnectionsResponse`

NewListFQDNConnectionsResponseWithDefaults instantiates a new ListFQDNConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListFQDNConnectionsResponse) GetData() []FqdnConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListFQDNConnectionsResponse) GetDataOk() (*[]FqdnConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListFQDNConnectionsResponse) SetData(v []FqdnConnection)`

SetData sets Data field to given value.

### HasData

`func (o *ListFQDNConnectionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListFQDNConnectionsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFQDNConnectionsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFQDNConnectionsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFQDNConnectionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


