# ParticipantResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The id of the account the resource belongs to. | [optional] 
**ApiVersion** | Pointer to **string** | The version of the API that was used to make the request. | [optional] 
**CallSid** | Pointer to **string** | The identifier of this participant&#39;s call. | [optional] 
**CallSidLegacy** | Pointer to **string** | The identifier of this participant&#39;s call. | [optional] 
**Coaching** | Pointer to **bool** | Whether the participant is coaching another call. | [optional] 
**CoachingCallSid** | Pointer to **string** | The identifier of the coached participant&#39;s call. | [optional] 
**CoachingCallSidLegacy** | Pointer to **string** | The identifier of the coached participant&#39;s call. | [optional] 
**DateCreated** | Pointer to **string** | The timestamp of when the resource was created. | [optional] 
**DateUpdated** | Pointer to **string** | The timestamp of when the resource was last updated. | [optional] 
**EndConferenceOnExit** | Pointer to **bool** | Whether the conference ends when the participant leaves. | [optional] 
**Hold** | Pointer to **bool** | Whether the participant is on hold. | [optional] 
**Muted** | Pointer to **bool** | Whether the participant is muted. | [optional] 
**Status** | Pointer to **string** | The status of the participant&#39;s call in the conference. | [optional] 
**Uri** | Pointer to **string** | The relative URI for this participant. | [optional] 

## Methods

### NewParticipantResource

`func NewParticipantResource() *ParticipantResource`

NewParticipantResource instantiates a new ParticipantResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantResourceWithDefaults

`func NewParticipantResourceWithDefaults() *ParticipantResource`

NewParticipantResourceWithDefaults instantiates a new ParticipantResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ParticipantResource) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ParticipantResource) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ParticipantResource) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ParticipantResource) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetApiVersion

`func (o *ParticipantResource) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ParticipantResource) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ParticipantResource) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ParticipantResource) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCallSid

`func (o *ParticipantResource) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ParticipantResource) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ParticipantResource) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ParticipantResource) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### GetCallSidLegacy

`func (o *ParticipantResource) GetCallSidLegacy() string`

GetCallSidLegacy returns the CallSidLegacy field if non-nil, zero value otherwise.

### GetCallSidLegacyOk

`func (o *ParticipantResource) GetCallSidLegacyOk() (*string, bool)`

GetCallSidLegacyOk returns a tuple with the CallSidLegacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSidLegacy

`func (o *ParticipantResource) SetCallSidLegacy(v string)`

SetCallSidLegacy sets CallSidLegacy field to given value.

### HasCallSidLegacy

`func (o *ParticipantResource) HasCallSidLegacy() bool`

HasCallSidLegacy returns a boolean if a field has been set.

### GetCoaching

`func (o *ParticipantResource) GetCoaching() bool`

GetCoaching returns the Coaching field if non-nil, zero value otherwise.

### GetCoachingOk

`func (o *ParticipantResource) GetCoachingOk() (*bool, bool)`

GetCoachingOk returns a tuple with the Coaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaching

`func (o *ParticipantResource) SetCoaching(v bool)`

SetCoaching sets Coaching field to given value.

### HasCoaching

`func (o *ParticipantResource) HasCoaching() bool`

HasCoaching returns a boolean if a field has been set.

### GetCoachingCallSid

`func (o *ParticipantResource) GetCoachingCallSid() string`

GetCoachingCallSid returns the CoachingCallSid field if non-nil, zero value otherwise.

### GetCoachingCallSidOk

`func (o *ParticipantResource) GetCoachingCallSidOk() (*string, bool)`

GetCoachingCallSidOk returns a tuple with the CoachingCallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoachingCallSid

`func (o *ParticipantResource) SetCoachingCallSid(v string)`

SetCoachingCallSid sets CoachingCallSid field to given value.

### HasCoachingCallSid

`func (o *ParticipantResource) HasCoachingCallSid() bool`

HasCoachingCallSid returns a boolean if a field has been set.

### GetCoachingCallSidLegacy

`func (o *ParticipantResource) GetCoachingCallSidLegacy() string`

GetCoachingCallSidLegacy returns the CoachingCallSidLegacy field if non-nil, zero value otherwise.

### GetCoachingCallSidLegacyOk

`func (o *ParticipantResource) GetCoachingCallSidLegacyOk() (*string, bool)`

GetCoachingCallSidLegacyOk returns a tuple with the CoachingCallSidLegacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoachingCallSidLegacy

`func (o *ParticipantResource) SetCoachingCallSidLegacy(v string)`

SetCoachingCallSidLegacy sets CoachingCallSidLegacy field to given value.

### HasCoachingCallSidLegacy

`func (o *ParticipantResource) HasCoachingCallSidLegacy() bool`

HasCoachingCallSidLegacy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ParticipantResource) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ParticipantResource) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ParticipantResource) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ParticipantResource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *ParticipantResource) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ParticipantResource) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ParticipantResource) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ParticipantResource) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetEndConferenceOnExit

`func (o *ParticipantResource) GetEndConferenceOnExit() bool`

GetEndConferenceOnExit returns the EndConferenceOnExit field if non-nil, zero value otherwise.

### GetEndConferenceOnExitOk

`func (o *ParticipantResource) GetEndConferenceOnExitOk() (*bool, bool)`

GetEndConferenceOnExitOk returns a tuple with the EndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndConferenceOnExit

`func (o *ParticipantResource) SetEndConferenceOnExit(v bool)`

SetEndConferenceOnExit sets EndConferenceOnExit field to given value.

### HasEndConferenceOnExit

`func (o *ParticipantResource) HasEndConferenceOnExit() bool`

HasEndConferenceOnExit returns a boolean if a field has been set.

### GetHold

`func (o *ParticipantResource) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *ParticipantResource) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *ParticipantResource) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *ParticipantResource) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetMuted

`func (o *ParticipantResource) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *ParticipantResource) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *ParticipantResource) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *ParticipantResource) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetStatus

`func (o *ParticipantResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ParticipantResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ParticipantResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ParticipantResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUri

`func (o *ParticipantResource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ParticipantResource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ParticipantResource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ParticipantResource) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


