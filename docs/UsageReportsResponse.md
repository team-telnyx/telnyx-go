# UsageReportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**PaginationData**](PaginationData.md) |  | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewUsageReportsResponse

`func NewUsageReportsResponse() *UsageReportsResponse`

NewUsageReportsResponse instantiates a new UsageReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportsResponseWithDefaults

`func NewUsageReportsResponseWithDefaults() *UsageReportsResponse`

NewUsageReportsResponseWithDefaults instantiates a new UsageReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *UsageReportsResponse) GetMeta() PaginationData`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UsageReportsResponse) GetMetaOk() (*PaginationData, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UsageReportsResponse) SetMeta(v PaginationData)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UsageReportsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *UsageReportsResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageReportsResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageReportsResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UsageReportsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


