# PaginatedVerificationRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]VerificationRequestStatus**](VerificationRequestStatus.md) | The records yielded by this request | [optional] [default to []]
**TotalRecords** | Pointer to **int32** | The total amount of records for these query parameters | [optional] [default to 0]

## Methods

### NewPaginatedVerificationRequestStatus

`func NewPaginatedVerificationRequestStatus() *PaginatedVerificationRequestStatus`

NewPaginatedVerificationRequestStatus instantiates a new PaginatedVerificationRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVerificationRequestStatusWithDefaults

`func NewPaginatedVerificationRequestStatusWithDefaults() *PaginatedVerificationRequestStatus`

NewPaginatedVerificationRequestStatusWithDefaults instantiates a new PaginatedVerificationRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *PaginatedVerificationRequestStatus) GetRecords() []VerificationRequestStatus`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PaginatedVerificationRequestStatus) GetRecordsOk() (*[]VerificationRequestStatus, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PaginatedVerificationRequestStatus) SetRecords(v []VerificationRequestStatus)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *PaginatedVerificationRequestStatus) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotalRecords

`func (o *PaginatedVerificationRequestStatus) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *PaginatedVerificationRequestStatus) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *PaginatedVerificationRequestStatus) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *PaginatedVerificationRequestStatus) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


