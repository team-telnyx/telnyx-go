# ListUploadsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Upload**](Upload.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListUploadsResponse

`func NewListUploadsResponse() *ListUploadsResponse`

NewListUploadsResponse instantiates a new ListUploadsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUploadsResponseWithDefaults

`func NewListUploadsResponseWithDefaults() *ListUploadsResponse`

NewListUploadsResponseWithDefaults instantiates a new ListUploadsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListUploadsResponse) GetData() []Upload`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUploadsResponse) GetDataOk() (*[]Upload, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUploadsResponse) SetData(v []Upload)`

SetData sets Data field to given value.

### HasData

`func (o *ListUploadsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListUploadsResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUploadsResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUploadsResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUploadsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


