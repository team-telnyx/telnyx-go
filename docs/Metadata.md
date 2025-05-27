# Metadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPages** | Pointer to **float64** | Total number of pages based on pagination settings | [optional] 
**TotalResults** | Pointer to **float64** | Total number of results | [optional] 
**PageNumber** | Pointer to **float64** | Current Page based on pagination settings (included when defaults are used.) | [optional] 
**PageSize** | Pointer to **float64** | Number of results to return per page based on pagination settings (included when defaults are used.) | [optional] 

## Methods

### NewMetadata

`func NewMetadata() *Metadata`

NewMetadata instantiates a new Metadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithDefaults

`func NewMetadataWithDefaults() *Metadata`

NewMetadataWithDefaults instantiates a new Metadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPages

`func (o *Metadata) GetTotalPages() float64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *Metadata) GetTotalPagesOk() (*float64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *Metadata) SetTotalPages(v float64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *Metadata) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *Metadata) GetTotalResults() float64`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *Metadata) GetTotalResultsOk() (*float64, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *Metadata) SetTotalResults(v float64)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *Metadata) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetPageNumber

`func (o *Metadata) GetPageNumber() float64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *Metadata) GetPageNumberOk() (*float64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *Metadata) SetPageNumber(v float64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *Metadata) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *Metadata) GetPageSize() float64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Metadata) GetPageSizeOk() (*float64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Metadata) SetPageSize(v float64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Metadata) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


