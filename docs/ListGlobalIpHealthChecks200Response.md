# ListGlobalIpHealthChecks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GlobalIPHealthCheck**](GlobalIPHealthCheck.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListGlobalIpHealthChecks200Response

`func NewListGlobalIpHealthChecks200Response() *ListGlobalIpHealthChecks200Response`

NewListGlobalIpHealthChecks200Response instantiates a new ListGlobalIpHealthChecks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGlobalIpHealthChecks200ResponseWithDefaults

`func NewListGlobalIpHealthChecks200ResponseWithDefaults() *ListGlobalIpHealthChecks200Response`

NewListGlobalIpHealthChecks200ResponseWithDefaults instantiates a new ListGlobalIpHealthChecks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListGlobalIpHealthChecks200Response) GetData() []GlobalIPHealthCheck`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListGlobalIpHealthChecks200Response) GetDataOk() (*[]GlobalIPHealthCheck, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListGlobalIpHealthChecks200Response) SetData(v []GlobalIPHealthCheck)`

SetData sets Data field to given value.

### HasData

`func (o *ListGlobalIpHealthChecks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListGlobalIpHealthChecks200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListGlobalIpHealthChecks200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListGlobalIpHealthChecks200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListGlobalIpHealthChecks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


