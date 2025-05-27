# PaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResults** | **int32** | Total number of results returned. | 
**TotalPages** | **int32** | Total number of pages from the results. | 
**PageNumber** | **int32** | The current page number. | 
**PageSize** | **int32** | The number of results per page. | 

## Methods

### NewPaginationResponse

`func NewPaginationResponse(totalResults int32, totalPages int32, pageNumber int32, pageSize int32, ) *PaginationResponse`

NewPaginationResponse instantiates a new PaginationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseWithDefaults

`func NewPaginationResponseWithDefaults() *PaginationResponse`

NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResults

`func (o *PaginationResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PaginationResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PaginationResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetTotalPages

`func (o *PaginationResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PaginationResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PaginationResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.


### GetPageNumber

`func (o *PaginationResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *PaginationResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *PaginationResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *PaginationResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginationResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginationResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


