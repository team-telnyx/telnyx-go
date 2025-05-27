# ConferenceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The id of the account the resource belongs to. | [optional] 
**ApiVersion** | Pointer to **string** | The version of the API that was used to make the request. | [optional] 
**CallSidEndingConference** | Pointer to **string** | Caller ID, if present. | [optional] 
**DateCreated** | Pointer to **string** | The timestamp of when the resource was created. | [optional] 
**DateUpdated** | Pointer to **string** | The timestamp of when the resource was last updated. | [optional] 
**FriendlyName** | Pointer to **string** | A string that you assigned to describe this conference room. | [optional] 
**ReasonConferenceEnded** | Pointer to **string** | The reason why a conference ended. When a conference is in progress, will be null. | [optional] 
**Region** | Pointer to **string** | A string representing the region where the conference is hosted. | [optional] 
**Sid** | Pointer to **string** | The unique identifier of the conference. | [optional] 
**Status** | Pointer to **string** | The status of this conference. | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs. | [optional] 
**Uri** | Pointer to **string** | The relative URI for this conference. | [optional] 

## Methods

### NewConferenceResource

`func NewConferenceResource() *ConferenceResource`

NewConferenceResource instantiates a new ConferenceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceResourceWithDefaults

`func NewConferenceResourceWithDefaults() *ConferenceResource`

NewConferenceResourceWithDefaults instantiates a new ConferenceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConferenceResource) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConferenceResource) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConferenceResource) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConferenceResource) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetApiVersion

`func (o *ConferenceResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ConferenceResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ConferenceResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ConferenceResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCallSidEndingConference

`func (o *ConferenceResource) GetCallSidEndingConference() string`

GetCallSidEndingConference returns the CallSidEndingConference field if non-nil, zero value otherwise.

### GetCallSidEndingConferenceOk

`func (o *ConferenceResource) GetCallSidEndingConferenceOk() (*string, bool)`

GetCallSidEndingConferenceOk returns a tuple with the CallSidEndingConference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSidEndingConference

`func (o *ConferenceResource) SetCallSidEndingConference(v string)`

SetCallSidEndingConference sets CallSidEndingConference field to given value.

### HasCallSidEndingConference

`func (o *ConferenceResource) HasCallSidEndingConference() bool`

HasCallSidEndingConference returns a boolean if a field has been set.

### GetDateCreated

`func (o *ConferenceResource) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConferenceResource) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConferenceResource) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConferenceResource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ConferenceResource) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConferenceResource) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConferenceResource) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConferenceResource) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetFriendlyName

`func (o *ConferenceResource) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ConferenceResource) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ConferenceResource) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ConferenceResource) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetReasonConferenceEnded

`func (o *ConferenceResource) GetReasonConferenceEnded() string`

GetReasonConferenceEnded returns the ReasonConferenceEnded field if non-nil, zero value otherwise.

### GetReasonConferenceEndedOk

`func (o *ConferenceResource) GetReasonConferenceEndedOk() (*string, bool)`

GetReasonConferenceEndedOk returns a tuple with the ReasonConferenceEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonConferenceEnded

`func (o *ConferenceResource) SetReasonConferenceEnded(v string)`

SetReasonConferenceEnded sets ReasonConferenceEnded field to given value.

### HasReasonConferenceEnded

`func (o *ConferenceResource) HasReasonConferenceEnded() bool`

HasReasonConferenceEnded returns a boolean if a field has been set.

### GetRegion

`func (o *ConferenceResource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ConferenceResource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ConferenceResource) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ConferenceResource) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSid

`func (o *ConferenceResource) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConferenceResource) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConferenceResource) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConferenceResource) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetStatus

`func (o *ConferenceResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConferenceResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConferenceResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConferenceResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubresourceUris

`func (o *ConferenceResource) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ConferenceResource) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ConferenceResource) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ConferenceResource) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### GetUri

`func (o *ConferenceResource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ConferenceResource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ConferenceResource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ConferenceResource) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


