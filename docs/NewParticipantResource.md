# NewParticipantResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The id of the account the resource belongs to. | [optional] 
**CallSid** | Pointer to **string** | The identifier of this participant&#39;s call. | [optional] 
**Coaching** | Pointer to **bool** | Whether the participant is coaching another call. | [optional] 
**CoachingCallSid** | Pointer to **string** | The identifier of the coached participant&#39;s call. | [optional] 
**EndConferenceOnExit** | Pointer to **bool** | Whether the conference ends when the participant leaves. | [optional] 
**Hold** | Pointer to **bool** | Whether the participant is on hold. | [optional] 
**Muted** | Pointer to **bool** | Whether the participant is muted. | [optional] 
**Status** | Pointer to **string** | The status of the participant&#39;s call in the conference. | [optional] 
**Uri** | Pointer to **string** | The relative URI for this participant. | [optional] 

## Methods

### NewNewParticipantResource

`func NewNewParticipantResource() *NewParticipantResource`

NewNewParticipantResource instantiates a new NewParticipantResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewParticipantResourceWithDefaults

`func NewNewParticipantResourceWithDefaults() *NewParticipantResource`

NewNewParticipantResourceWithDefaults instantiates a new NewParticipantResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NewParticipantResource) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NewParticipantResource) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NewParticipantResource) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NewParticipantResource) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetCallSid

`func (o *NewParticipantResource) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *NewParticipantResource) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *NewParticipantResource) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *NewParticipantResource) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### GetCoaching

`func (o *NewParticipantResource) GetCoaching() bool`

GetCoaching returns the Coaching field if non-nil, zero value otherwise.

### GetCoachingOk

`func (o *NewParticipantResource) GetCoachingOk() (*bool, bool)`

GetCoachingOk returns a tuple with the Coaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaching

`func (o *NewParticipantResource) SetCoaching(v bool)`

SetCoaching sets Coaching field to given value.

### HasCoaching

`func (o *NewParticipantResource) HasCoaching() bool`

HasCoaching returns a boolean if a field has been set.

### GetCoachingCallSid

`func (o *NewParticipantResource) GetCoachingCallSid() string`

GetCoachingCallSid returns the CoachingCallSid field if non-nil, zero value otherwise.

### GetCoachingCallSidOk

`func (o *NewParticipantResource) GetCoachingCallSidOk() (*string, bool)`

GetCoachingCallSidOk returns a tuple with the CoachingCallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoachingCallSid

`func (o *NewParticipantResource) SetCoachingCallSid(v string)`

SetCoachingCallSid sets CoachingCallSid field to given value.

### HasCoachingCallSid

`func (o *NewParticipantResource) HasCoachingCallSid() bool`

HasCoachingCallSid returns a boolean if a field has been set.

### GetEndConferenceOnExit

`func (o *NewParticipantResource) GetEndConferenceOnExit() bool`

GetEndConferenceOnExit returns the EndConferenceOnExit field if non-nil, zero value otherwise.

### GetEndConferenceOnExitOk

`func (o *NewParticipantResource) GetEndConferenceOnExitOk() (*bool, bool)`

GetEndConferenceOnExitOk returns a tuple with the EndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndConferenceOnExit

`func (o *NewParticipantResource) SetEndConferenceOnExit(v bool)`

SetEndConferenceOnExit sets EndConferenceOnExit field to given value.

### HasEndConferenceOnExit

`func (o *NewParticipantResource) HasEndConferenceOnExit() bool`

HasEndConferenceOnExit returns a boolean if a field has been set.

### GetHold

`func (o *NewParticipantResource) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *NewParticipantResource) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *NewParticipantResource) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *NewParticipantResource) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetMuted

`func (o *NewParticipantResource) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *NewParticipantResource) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *NewParticipantResource) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *NewParticipantResource) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetStatus

`func (o *NewParticipantResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewParticipantResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewParticipantResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewParticipantResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUri

`func (o *NewParticipantResource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NewParticipantResource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NewParticipantResource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NewParticipantResource) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


