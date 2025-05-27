# ParticipantResourceIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participants** | Pointer to [**[]ParticipantResource**](ParticipantResource.md) |  | [optional] 
**End** | Pointer to **int32** | The number of the last element on the page, zero-indexed. | [optional] 
**FirstPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Participants.json?page&#x3D;0&amp;pagesize&#x3D;20 | [optional] 
**NextPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Conferences/6dc6cc1a-1ba1-4351-86b8-4c22c95cd98f/Participants.json?Page&#x3D;1&amp;PageSize&#x3D;1&amp;PageToken&#x3D;MTY4AjgyNDkwNzIxMQ | [optional] 
**Page** | Pointer to **int32** | Current page number, zero-indexed. | [optional] 
**PageSize** | Pointer to **int32** | The number of items on the page | [optional] 
**Start** | Pointer to **int32** | The number of the first element on the page, zero-indexed. | [optional] 
**Uri** | Pointer to **string** | The URI of the current page. | [optional] 

## Methods

### NewParticipantResourceIndex

`func NewParticipantResourceIndex() *ParticipantResourceIndex`

NewParticipantResourceIndex instantiates a new ParticipantResourceIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantResourceIndexWithDefaults

`func NewParticipantResourceIndexWithDefaults() *ParticipantResourceIndex`

NewParticipantResourceIndexWithDefaults instantiates a new ParticipantResourceIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipants

`func (o *ParticipantResourceIndex) GetParticipants() []ParticipantResource`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ParticipantResourceIndex) GetParticipantsOk() (*[]ParticipantResource, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ParticipantResourceIndex) SetParticipants(v []ParticipantResource)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ParticipantResourceIndex) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetEnd

`func (o *ParticipantResourceIndex) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ParticipantResourceIndex) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ParticipantResourceIndex) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ParticipantResourceIndex) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *ParticipantResourceIndex) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *ParticipantResourceIndex) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *ParticipantResourceIndex) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *ParticipantResourceIndex) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *ParticipantResourceIndex) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *ParticipantResourceIndex) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *ParticipantResourceIndex) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *ParticipantResourceIndex) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *ParticipantResourceIndex) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ParticipantResourceIndex) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ParticipantResourceIndex) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ParticipantResourceIndex) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ParticipantResourceIndex) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ParticipantResourceIndex) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ParticipantResourceIndex) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ParticipantResourceIndex) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStart

`func (o *ParticipantResourceIndex) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ParticipantResourceIndex) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ParticipantResourceIndex) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *ParticipantResourceIndex) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *ParticipantResourceIndex) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ParticipantResourceIndex) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ParticipantResourceIndex) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ParticipantResourceIndex) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


