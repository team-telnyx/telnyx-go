# ConferenceResourceIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conferences** | Pointer to [**[]ConferenceResource**](ConferenceResource.md) |  | [optional] 
**End** | Pointer to **int32** | The number of the last element on the page, zero-indexed. | [optional] 
**FirstPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences.json?Page&#x3D;0&amp;PageSize&#x3D;1 | [optional] 
**NextPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences.json?Page&#x3D;1&amp;PageSize&#x3D;1&amp;PageToken&#x3D;MTY4AjgyNDkwNzIxMQ | [optional] 
**Page** | Pointer to **int32** | Current page number, zero-indexed. | [optional] 
**PageSize** | Pointer to **int32** | The number of items on the page | [optional] 
**Start** | Pointer to **int32** | The number of the first element on the page, zero-indexed. | [optional] 
**Uri** | Pointer to **string** | The URI of the current page. | [optional] 

## Methods

### NewConferenceResourceIndex

`func NewConferenceResourceIndex() *ConferenceResourceIndex`

NewConferenceResourceIndex instantiates a new ConferenceResourceIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceResourceIndexWithDefaults

`func NewConferenceResourceIndexWithDefaults() *ConferenceResourceIndex`

NewConferenceResourceIndexWithDefaults instantiates a new ConferenceResourceIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConferences

`func (o *ConferenceResourceIndex) GetConferences() []ConferenceResource`

GetConferences returns the Conferences field if non-nil, zero value otherwise.

### GetConferencesOk

`func (o *ConferenceResourceIndex) GetConferencesOk() (*[]ConferenceResource, bool)`

GetConferencesOk returns a tuple with the Conferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferences

`func (o *ConferenceResourceIndex) SetConferences(v []ConferenceResource)`

SetConferences sets Conferences field to given value.

### HasConferences

`func (o *ConferenceResourceIndex) HasConferences() bool`

HasConferences returns a boolean if a field has been set.

### GetEnd

`func (o *ConferenceResourceIndex) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ConferenceResourceIndex) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ConferenceResourceIndex) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ConferenceResourceIndex) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ConferenceResourceIndex) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ConferenceResourceIndex) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ConferenceResourceIndex) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ConferenceResourceIndex) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ConferenceResourceIndex) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ConferenceResourceIndex) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ConferenceResourceIndex) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ConferenceResourceIndex) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ConferenceResourceIndex) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ConferenceResourceIndex) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ConferenceResourceIndex) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ConferenceResourceIndex) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ConferenceResourceIndex) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ConferenceResourceIndex) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ConferenceResourceIndex) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ConferenceResourceIndex) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStart

`func (o *ConferenceResourceIndex) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ConferenceResourceIndex) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ConferenceResourceIndex) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ConferenceResourceIndex) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ConferenceResourceIndex) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ConferenceResourceIndex) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ConferenceResourceIndex) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ConferenceResourceIndex) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


