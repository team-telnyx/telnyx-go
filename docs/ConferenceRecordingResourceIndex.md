# ConferenceRecordingResourceIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recordings** | Pointer to [**[]ConferenceRecordingResource**](ConferenceRecordingResource.md) |  | [optional] 
**End** | Pointer to **int32** | The number of the last element on the page, zero-indexed. | [optional] 
**FirstPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Recordings.json?page&#x3D;0&amp;pagesize&#x3D;20 | [optional] 
**NextPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Recordings.json?Page&#x3D;1&amp;PageSize&#x3D;1&amp;PageToken&#x3D;MTY4AjgyNDkwNzIxMQ | [optional] 
**Page** | Pointer to **int32** | Current page number, zero-indexed. | [optional] 
**PageSize** | Pointer to **int32** | The number of items on the page | [optional] 
**Start** | Pointer to **int32** | The number of the first element on the page, zero-indexed. | [optional] 
**Uri** | Pointer to **string** | The URI of the current page. | [optional] 

## Methods

### NewConferenceRecordingResourceIndex

`func NewConferenceRecordingResourceIndex() *ConferenceRecordingResourceIndex`

NewConferenceRecordingResourceIndex instantiates a new ConferenceRecordingResourceIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceRecordingResourceIndexWithDefaults

`func NewConferenceRecordingResourceIndexWithDefaults() *ConferenceRecordingResourceIndex`

NewConferenceRecordingResourceIndexWithDefaults instantiates a new ConferenceRecordingResourceIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordings

`func (o *ConferenceRecordingResourceIndex) GetRecordings() []ConferenceRecordingResource`

GetRecordings returns the Recordings field if non-nil, zero value otherwise.

### GetRecordingsOk

`func (o *ConferenceRecordingResourceIndex) GetRecordingsOk() (*[]ConferenceRecordingResource, bool)`

GetRecordingsOk returns a tuple with the Recordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordings

`func (o *ConferenceRecordingResourceIndex) SetRecordings(v []ConferenceRecordingResource)`

SetRecordings sets Recordings field to given value.

### HasRecordings

`func (o *ConferenceRecordingResourceIndex) HasRecordings() bool`

HasRecordings returns a boolean if a field has been set.

### GetEnd

`func (o *ConferenceRecordingResourceIndex) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ConferenceRecordingResourceIndex) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ConferenceRecordingResourceIndex) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ConferenceRecordingResourceIndex) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ConferenceRecordingResourceIndex) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ConferenceRecordingResourceIndex) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ConferenceRecordingResourceIndex) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ConferenceRecordingResourceIndex) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ConferenceRecordingResourceIndex) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ConferenceRecordingResourceIndex) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ConferenceRecordingResourceIndex) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ConferenceRecordingResourceIndex) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ConferenceRecordingResourceIndex) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ConferenceRecordingResourceIndex) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ConferenceRecordingResourceIndex) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ConferenceRecordingResourceIndex) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ConferenceRecordingResourceIndex) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ConferenceRecordingResourceIndex) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ConferenceRecordingResourceIndex) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ConferenceRecordingResourceIndex) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStart

`func (o *ConferenceRecordingResourceIndex) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ConferenceRecordingResourceIndex) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ConferenceRecordingResourceIndex) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ConferenceRecordingResourceIndex) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ConferenceRecordingResourceIndex) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ConferenceRecordingResourceIndex) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ConferenceRecordingResourceIndex) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ConferenceRecordingResourceIndex) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


