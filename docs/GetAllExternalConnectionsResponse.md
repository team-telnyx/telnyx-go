# GetAllExternalConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ExternalConnection**](ExternalConnection.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewGetAllExternalConnectionsResponse

`func NewGetAllExternalConnectionsResponse() *GetAllExternalConnectionsResponse`

NewGetAllExternalConnectionsResponse instantiates a new GetAllExternalConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllExternalConnectionsResponseWithDefaults

`func NewGetAllExternalConnectionsResponseWithDefaults() *GetAllExternalConnectionsResponse`

NewGetAllExternalConnectionsResponseWithDefaults instantiates a new GetAllExternalConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllExternalConnectionsResponse) GetData() []ExternalConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllExternalConnectionsResponse) GetDataOk() (*[]ExternalConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllExternalConnectionsResponse) SetData(v []ExternalConnection)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllExternalConnectionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetAllExternalConnectionsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllExternalConnectionsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllExternalConnectionsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAllExternalConnectionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


