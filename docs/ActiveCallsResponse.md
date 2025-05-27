# ActiveCallsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ActiveCall**](ActiveCall.md) |  | [optional] 
**Meta** | Pointer to [**CursorPaginationMeta**](CursorPaginationMeta.md) |  | [optional] 

## Methods

### NewActiveCallsResponse

`func NewActiveCallsResponse() *ActiveCallsResponse`

NewActiveCallsResponse instantiates a new ActiveCallsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveCallsResponseWithDefaults

`func NewActiveCallsResponseWithDefaults() *ActiveCallsResponse`

NewActiveCallsResponseWithDefaults instantiates a new ActiveCallsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ActiveCallsResponse) GetData() []ActiveCall`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActiveCallsResponse) GetDataOk() (*[]ActiveCall, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActiveCallsResponse) SetData(v []ActiveCall)`

SetData sets Data field to given value.

### HasData

`func (o *ActiveCallsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ActiveCallsResponse) GetMeta() CursorPaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ActiveCallsResponse) GetMetaOk() (*CursorPaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ActiveCallsResponse) SetMeta(v CursorPaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ActiveCallsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


