# PaginatedUserBundlesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**PaginationResponse**](PaginationResponse.md) |  | 
**Data** | [**[]UserBundle**](UserBundle.md) |  | 

## Methods

### NewPaginatedUserBundlesResponse

`func NewPaginatedUserBundlesResponse(meta PaginationResponse, data []UserBundle, ) *PaginatedUserBundlesResponse`

NewPaginatedUserBundlesResponse instantiates a new PaginatedUserBundlesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedUserBundlesResponseWithDefaults

`func NewPaginatedUserBundlesResponseWithDefaults() *PaginatedUserBundlesResponse`

NewPaginatedUserBundlesResponseWithDefaults instantiates a new PaginatedUserBundlesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PaginatedUserBundlesResponse) GetMeta() PaginationResponse`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaginatedUserBundlesResponse) GetMetaOk() (*PaginationResponse, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaginatedUserBundlesResponse) SetMeta(v PaginationResponse)`

SetMeta sets Meta field to given value.


### GetData

`func (o *PaginatedUserBundlesResponse) GetData() []UserBundle`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedUserBundlesResponse) GetDataOk() (*[]UserBundle, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedUserBundlesResponse) SetData(v []UserBundle)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


