# ListReleasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Release**](Release.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListReleasesResponse

`func NewListReleasesResponse() *ListReleasesResponse`

NewListReleasesResponse instantiates a new ListReleasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReleasesResponseWithDefaults

`func NewListReleasesResponseWithDefaults() *ListReleasesResponse`

NewListReleasesResponseWithDefaults instantiates a new ListReleasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListReleasesResponse) GetData() []Release`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListReleasesResponse) GetDataOk() (*[]Release, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListReleasesResponse) SetData(v []Release)`

SetData sets Data field to given value.

### HasData

`func (o *ListReleasesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListReleasesResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListReleasesResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListReleasesResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListReleasesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


