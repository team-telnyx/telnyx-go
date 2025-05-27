# CallResourceIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to [**[]CallResource**](CallResource.md) |  | [optional] 
**End** | Pointer to **int32** | The number of the last element on the page, zero-indexed. | [optional] 
**FirstPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Calls.json?Page&#x3D;0&amp;PageSize&#x3D;1 | [optional] 
**NextPageUri** | Pointer to **string** | /v2/texml/Accounts/61bf923e-5e4d-4595-a110-56190ea18a1b/Calls.json?Page&#x3D;1&amp;PageSize&#x3D;1&amp;PageToken&#x3D;MTY4AjgyNDkwNzIxMQ | [optional] 
**Page** | Pointer to **int32** | Current page number, zero-indexed. | [optional] 
**PageSize** | Pointer to **int32** | The number of items on the page | [optional] 
**Start** | Pointer to **int32** | The number of the first element on the page, zero-indexed. | [optional] 
**Uri** | Pointer to **string** | The URI of the current page. | [optional] 

## Methods

### NewCallResourceIndex

`func NewCallResourceIndex() *CallResourceIndex`

NewCallResourceIndex instantiates a new CallResourceIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallResourceIndexWithDefaults

`func NewCallResourceIndexWithDefaults() *CallResourceIndex`

NewCallResourceIndexWithDefaults instantiates a new CallResourceIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *CallResourceIndex) GetCalls() []CallResource`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *CallResourceIndex) GetCallsOk() (*[]CallResource, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *CallResourceIndex) SetCalls(v []CallResource)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *CallResourceIndex) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetEnd

`func (o *CallResourceIndex) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CallResourceIndex) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CallResourceIndex) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CallResourceIndex) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFirstPageUri

`func (o *CallResourceIndex) GetFirstPageUri() string`

GetFirstPageUri returns the FirstPageUri field if non-nil, zero value otherwise.

### GetFirstPageUriOk

`func (o *CallResourceIndex) GetFirstPageUriOk() (*string, bool)`

GetFirstPageUriOk returns a tuple with the FirstPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUri

`func (o *CallResourceIndex) SetFirstPageUri(v string)`

SetFirstPageUri sets FirstPageUri field to given value.

### HasFirstPageUri

`func (o *CallResourceIndex) HasFirstPageUri() bool`

HasFirstPageUri returns a boolean if a field has been set.

### GetNextPageUri

`func (o *CallResourceIndex) GetNextPageUri() string`

GetNextPageUri returns the NextPageUri field if non-nil, zero value otherwise.

### GetNextPageUriOk

`func (o *CallResourceIndex) GetNextPageUriOk() (*string, bool)`

GetNextPageUriOk returns a tuple with the NextPageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUri

`func (o *CallResourceIndex) SetNextPageUri(v string)`

SetNextPageUri sets NextPageUri field to given value.

### HasNextPageUri

`func (o *CallResourceIndex) HasNextPageUri() bool`

HasNextPageUri returns a boolean if a field has been set.

### GetPage

`func (o *CallResourceIndex) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CallResourceIndex) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CallResourceIndex) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CallResourceIndex) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *CallResourceIndex) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *CallResourceIndex) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *CallResourceIndex) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *CallResourceIndex) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStart

`func (o *CallResourceIndex) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CallResourceIndex) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CallResourceIndex) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *CallResourceIndex) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUri

`func (o *CallResourceIndex) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CallResourceIndex) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CallResourceIndex) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CallResourceIndex) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


