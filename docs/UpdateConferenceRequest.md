# UpdateConferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The new status of the resource. Specifying &#x60;completed&#x60; will end the conference and hang up all participants. | [optional] 
**AnnounceUrl** | Pointer to **string** | The URL we should call to announce something into the conference. The URL may return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60;, &#x60;&lt;Pause&gt;&#x60;, or &#x60;&lt;Redirect&gt;&#x60; verbs. | [optional] 
**AnnounceMethod** | Pointer to **string** | The HTTP method used to call the &#x60;AnnounceUrl&#x60;. Defaults to &#x60;POST&#x60;. | [optional] 

## Methods

### NewUpdateConferenceRequest

`func NewUpdateConferenceRequest() *UpdateConferenceRequest`

NewUpdateConferenceRequest instantiates a new UpdateConferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConferenceRequestWithDefaults

`func NewUpdateConferenceRequestWithDefaults() *UpdateConferenceRequest`

NewUpdateConferenceRequestWithDefaults instantiates a new UpdateConferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateConferenceRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateConferenceRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateConferenceRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateConferenceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAnnounceUrl

`func (o *UpdateConferenceRequest) GetAnnounceUrl() string`

GetAnnounceUrl returns the AnnounceUrl field if non-nil, zero value otherwise.

### GetAnnounceUrlOk

`func (o *UpdateConferenceRequest) GetAnnounceUrlOk() (*string, bool)`

GetAnnounceUrlOk returns a tuple with the AnnounceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceUrl

`func (o *UpdateConferenceRequest) SetAnnounceUrl(v string)`

SetAnnounceUrl sets AnnounceUrl field to given value.

### HasAnnounceUrl

`func (o *UpdateConferenceRequest) HasAnnounceUrl() bool`

HasAnnounceUrl returns a boolean if a field has been set.

### GetAnnounceMethod

`func (o *UpdateConferenceRequest) GetAnnounceMethod() string`

GetAnnounceMethod returns the AnnounceMethod field if non-nil, zero value otherwise.

### GetAnnounceMethodOk

`func (o *UpdateConferenceRequest) GetAnnounceMethodOk() (*string, bool)`

GetAnnounceMethodOk returns a tuple with the AnnounceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceMethod

`func (o *UpdateConferenceRequest) SetAnnounceMethod(v string)`

SetAnnounceMethod sets AnnounceMethod field to given value.

### HasAnnounceMethod

`func (o *UpdateConferenceRequest) HasAnnounceMethod() bool`

HasAnnounceMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


