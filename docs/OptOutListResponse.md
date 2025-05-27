# OptOutListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]OptOutItem**](OptOutItem.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewOptOutListResponse

`func NewOptOutListResponse() *OptOutListResponse`

NewOptOutListResponse instantiates a new OptOutListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptOutListResponseWithDefaults

`func NewOptOutListResponseWithDefaults() *OptOutListResponse`

NewOptOutListResponseWithDefaults instantiates a new OptOutListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OptOutListResponse) GetData() []OptOutItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OptOutListResponse) GetDataOk() (*[]OptOutItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OptOutListResponse) SetData(v []OptOutItem)`

SetData sets Data field to given value.

### HasData

`func (o *OptOutListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *OptOutListResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OptOutListResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OptOutListResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OptOutListResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


