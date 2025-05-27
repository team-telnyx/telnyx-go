# ListCustomerServiceRecords200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CustomerServiceRecord**](CustomerServiceRecord.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewListCustomerServiceRecords200Response

`func NewListCustomerServiceRecords200Response() *ListCustomerServiceRecords200Response`

NewListCustomerServiceRecords200Response instantiates a new ListCustomerServiceRecords200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerServiceRecords200ResponseWithDefaults

`func NewListCustomerServiceRecords200ResponseWithDefaults() *ListCustomerServiceRecords200Response`

NewListCustomerServiceRecords200ResponseWithDefaults instantiates a new ListCustomerServiceRecords200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCustomerServiceRecords200Response) GetData() []CustomerServiceRecord`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCustomerServiceRecords200Response) GetDataOk() (*[]CustomerServiceRecord, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCustomerServiceRecords200Response) SetData(v []CustomerServiceRecord)`

SetData sets Data field to given value.

### HasData

`func (o *ListCustomerServiceRecords200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *ListCustomerServiceRecords200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCustomerServiceRecords200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCustomerServiceRecords200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCustomerServiceRecords200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


