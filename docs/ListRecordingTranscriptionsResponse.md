# ListRecordingTranscriptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transcriptions** | Pointer to [**[]TexmlRecordingTranscription**](TexmlRecordingTranscription.md) |  | [optional] 
**End** | Pointer to **int32** | The number of the last element on the page, zero-indexed | [optional] 
**FirstPageUri** | Pointer to **string** | Relative uri to the first page of the query results | [optional] 
**PreviousPageUri** | Pointer to **string** | Relative uri to the previous page of the query results | [optional] 
**NextPageUri** | Pointer to **string** | Relative uri to the next page of the query results | [optional] 
**Page** | Pointer to **int32** | Current page number, zero-indexed. | [optional] 
**PageSize** | Pointer to **int32** | The number of items on the page | [optional] 
**Start** | Pointer to **int32** | The number of the first element on the page, zero-indexed. | [optional] 
**Uri** | Pointer to **string** | The URI of the current page. | [optional] 

## Methods

### NewListRecordingTranscriptionsResponse

`func NewListRecordingTranscriptionsResponse() *ListRecordingTranscriptionsResponse`

NewListRecordingTranscriptionsResponse instantiates a new ListRecordingTranscriptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRecordingTranscriptionsResponseWithDefaults

`func NewListRecordingTranscriptionsResponseWithDefaults() *ListRecordingTranscriptionsResponse`

NewListRecordingTranscriptionsResponseWithDefaults instantiates a new ListRecordingTranscriptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranscriptions

`func (o *ListRecordingTranscriptionsResponse) GetTranscriptions() []TexmlRecordingTranscription`

GetTranscriptions returns the Transcriptions field if non-nil, zero value otherwise.

### GetTranscriptionsOk

`func (o *ListRecordingTranscriptionsResponse) GetTranscriptionsOk() (*[]TexmlRecordingTranscription, bool)`

GetTranscriptionsOk returns a tuple with the Transcriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptions

`func (o *ListRecordingTranscriptionsResponse) SetTranscriptions(v []TexmlRecordingTranscription)`

SetTranscriptions sets Transcriptions field to given value.

### HasTranscriptions

`func (o *ListRecordingTranscriptionsResponse) HasTranscriptions() bool`

HasTranscriptions returns a boolean if a field has been set.

### GetEnd

`func (o *ListRecordingTranscriptionsResponse) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ListRecordingTranscriptionsResponse) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ListRecordingTranscriptionsResponse) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ListRecordingTranscriptionsResponse) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ListRecordingTranscriptionsResponse) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ListRecordingTranscriptionsResponse) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ListRecordingTranscriptionsResponse) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ListRecordingTranscriptionsResponse) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetPreviousPageUri

`func (o *ListRecordingTranscriptionsResponse) GetPreviousPageUri() string`

GetPreviousPageUri returns the PreviousPageUri field if non-nil, zero value otherwise.

### GetPreviousPageUriOk

`func (o *ListRecordingTranscriptionsResponse) GetPreviousPageUriOk() (*string, bool)`

GetPreviousPageUriOk returns a tuple with the PreviousPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPageUri

`func (o *ListRecordingTranscriptionsResponse) SetPreviousPageUri(v string)`

SetPreviousPageUri sets PreviousPageUri field to given value.

### HasPreviousPageUri

`func (o *ListRecordingTranscriptionsResponse) HasPreviousPageUri() bool`

HasPreviousPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ListRecordingTranscriptionsResponse) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ListRecordingTranscriptionsResponse) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ListRecordingTranscriptionsResponse) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ListRecordingTranscriptionsResponse) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ListRecordingTranscriptionsResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListRecordingTranscriptionsResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListRecordingTranscriptionsResponse) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListRecordingTranscriptionsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ListRecordingTranscriptionsResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListRecordingTranscriptionsResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListRecordingTranscriptionsResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListRecordingTranscriptionsResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStart

`func (o *ListRecordingTranscriptionsResponse) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ListRecordingTranscriptionsResponse) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ListRecordingTranscriptionsResponse) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ListRecordingTranscriptionsResponse) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ListRecordingTranscriptionsResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ListRecordingTranscriptionsResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ListRecordingTranscriptionsResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ListRecordingTranscriptionsResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


