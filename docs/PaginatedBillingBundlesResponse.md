# PaginatedBillingBundlesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**PaginationResponse**](PaginationResponse.md) |  | 
**Data** | [**[]BillingBundleSummary**](BillingBundleSummary.md) |  | 

## Methods

### NewPaginatedBillingBundlesResponse

`func NewPaginatedBillingBundlesResponse(meta PaginationResponse, data []BillingBundleSummary, ) *PaginatedBillingBundlesResponse`

NewPaginatedBillingBundlesResponse instantiates a new PaginatedBillingBundlesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBillingBundlesResponseWithDefaults

`func NewPaginatedBillingBundlesResponseWithDefaults() *PaginatedBillingBundlesResponse`

NewPaginatedBillingBundlesResponseWithDefaults instantiates a new PaginatedBillingBundlesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PaginatedBillingBundlesResponse) GetMeta() PaginationResponse`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedBillingBundlesResponse) GetMetaOk() (*PaginationResponse, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedBillingBundlesResponse) SetMeta(v PaginationResponse)`

SetMeta sets Meta field to given value.


### GetData

`func (o *PaginatedBillingBundlesResponse) GetData() []BillingBundleSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedBillingBundlesResponse) GetDataOk() (*[]BillingBundleSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedBillingBundlesResponse) SetData(v []BillingBundleSummary)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


