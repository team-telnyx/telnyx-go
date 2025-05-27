# ListConferencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Conference**](Conference.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListConferencesResponse

`func NewListConferencesResponse() *ListConferencesResponse`

NewListConferencesResponse instantiates a new ListConferencesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConferencesResponseWithDefaults

`func NewListConferencesResponseWithDefaults() *ListConferencesResponse`

NewListConferencesResponseWithDefaults instantiates a new ListConferencesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListConferencesResponse) GetData() []Conference`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListConferencesResponse) GetDataOk() (*[]Conference, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListConferencesResponse) SetData(v []Conference)`

SetData sets Data field to given value.

### HasData

`func (o *ListConferencesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListConferencesResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConferencesResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConferencesResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConferencesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


