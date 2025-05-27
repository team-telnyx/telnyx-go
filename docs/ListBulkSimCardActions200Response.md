# ListBulkSimCardActions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BulkSIMCardActionDetailed**](BulkSIMCardActionDetailed.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListBulkSimCardActions200Response

`func NewListBulkSimCardActions200Response() *ListBulkSimCardActions200Response`

NewListBulkSimCardActions200Response instantiates a new ListBulkSimCardActions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBulkSimCardActions200ResponseWithDefaults

`func NewListBulkSimCardActions200ResponseWithDefaults() *ListBulkSimCardActions200Response`

NewListBulkSimCardActions200ResponseWithDefaults instantiates a new ListBulkSimCardActions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListBulkSimCardActions200Response) GetData() []BulkSIMCardActionDetailed`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListBulkSimCardActions200Response) GetDataOk() (*[]BulkSIMCardActionDetailed, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListBulkSimCardActions200Response) SetData(v []BulkSIMCardActionDetailed)`

SetData sets Data field to given value.

### HasData

`func (o *ListBulkSimCardActions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListBulkSimCardActions200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBulkSimCardActions200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBulkSimCardActions200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBulkSimCardActions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


