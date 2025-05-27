# TexmlGetCallRecordingsResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recordings** | Pointer to [**[]TexmlGetCallRecordingResponseBody**](TexmlGetCallRecordingResponseBody.md) |  | [optional] 
**End** | Pointer to **int32** | The number of the last element on the page, zero-indexed. | [optional] 
**FirstPageUri** | Pointer to **string** | Relative uri to the first page of the query results | [optional] 
**PreviousPageUri** | Pointer to **string** | Relative uri to the previous page of the query results | [optional] 
**NextPageUri** | Pointer to **string** | Relative uri to the next page of the query results | [optional] 
**Page** | Pointer to **int32** | Current page number, zero-indexed. | [optional] 
**PageSize** | Pointer to **int32** | The number of items on the page | [optional] 
**Start** | Pointer to **int32** | The number of the first element on the page, zero-indexed. | [optional] 
**Uri** | Pointer to **string** | The URI of the current page. | [optional] 

## Methods

### NewTexmlGetCallRecordingsResponseBody

`func NewTexmlGetCallRecordingsResponseBody() *TexmlGetCallRecordingsResponseBody`

NewTexmlGetCallRecordingsResponseBody instantiates a new TexmlGetCallRecordingsResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTexmlGetCallRecordingsResponseBodyWithDefaults

`func NewTexmlGetCallRecordingsResponseBodyWithDefaults() *TexmlGetCallRecordingsResponseBody`

NewTexmlGetCallRecordingsResponseBodyWithDefaults instantiates a new TexmlGetCallRecordingsResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordings

`func (o *TexmlGetCallRecordingsResponseBody) GetRecordings() []TexmlGetCallRecordingResponseBody`

GetRecordings returns the Recordings field if non-nil, zero value otherwise.

### GetRecordingsOk

`func (o *TexmlGetCallRecordingsResponseBody) GetRecordingsOk() (*[]TexmlGetCallRecordingResponseBody, bool)`

GetRecordingsOk returns a tuple with the Recordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordings

`func (o *TexmlGetCallRecordingsResponseBody) SetRecordings(v []TexmlGetCallRecordingResponseBody)`

SetRecordings sets Recordings field to given value.

### HasRecordings

`func (o *TexmlGetCallRecordingsResponseBody) HasRecordings() bool`

HasRecordings returns a boolean if a field has been set.

### GetEnd

`func (o *TexmlGetCallRecordingsResponseBody) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TexmlGetCallRecordingsResponseBody) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TexmlGetCallRecordingsResponseBody) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TexmlGetCallRecordingsResponseBody) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *TexmlGetCallRecordingsResponseBody) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *TexmlGetCallRecordingsResponseBody) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *TexmlGetCallRecordingsResponseBody) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *TexmlGetCallRecordingsResponseBody) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *TexmlGetCallRecordingsResponseBody) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *TexmlGetCallRecordingsResponseBody) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *TexmlGetCallRecordingsResponseBody) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *TexmlGetCallRecordingsResponseBody) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *TexmlGetCallRecordingsResponseBody) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *TexmlGetCallRecordingsResponseBody) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *TexmlGetCallRecordingsResponseBody) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *TexmlGetCallRecordingsResponseBody) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *TexmlGetCallRecordingsResponseBody) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TexmlGetCallRecordingsResponseBody) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TexmlGetCallRecordingsResponseBody) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *TexmlGetCallRecordingsResponseBody) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *TexmlGetCallRecordingsResponseBody) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TexmlGetCallRecordingsResponseBody) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TexmlGetCallRecordingsResponseBody) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TexmlGetCallRecordingsResponseBody) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStart

`func (o *TexmlGetCallRecordingsResponseBody) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TexmlGetCallRecordingsResponseBody) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TexmlGetCallRecordingsResponseBody) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *TexmlGetCallRecordingsResponseBody) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *TexmlGetCallRecordingsResponseBody) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TexmlGetCallRecordingsResponseBody) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TexmlGetCallRecordingsResponseBody) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TexmlGetCallRecordingsResponseBody) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


