# JoinConferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Unique identifier and token for controlling the call | 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. Please note that the client_state will be updated for the participient call leg and the change will not affect conferencing webhooks unless the participient is the owner of the conference. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] 
**EndConferenceOnExit** | Pointer to **bool** | Whether the conference should end and all remaining participants be hung up after the participant leaves the conference. Defaults to \&quot;false\&quot;. | [optional] 
**SoftEndConferenceOnExit** | Pointer to **bool** | Whether the conference should end after the participant leaves the conference. NOTE this doesn&#39;t hang up the other participants. Defaults to \&quot;false\&quot;. | [optional] 
**Hold** | Pointer to **bool** | Whether the participant should be put on hold immediately after joining the conference. Defaults to \&quot;false\&quot;. | [optional] 
**HoldAudioUrl** | Pointer to **string** | The URL of a file to be played to the participant when they are put on hold after joining the conference. hold_media_name and hold_audio_url cannot be used together in one request. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. This property takes effect only if \&quot;hold\&quot; is set to \&quot;true\&quot;. | [optional] 
**HoldMediaName** | Pointer to **string** | The media_name of a file to be played to the participant when they are put on hold after joining the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. This property takes effect only if \&quot;hold\&quot; is set to \&quot;true\&quot;. | [optional] 
**Mute** | Pointer to **bool** | Whether the participant should be muted immediately after joining the conference. Defaults to \&quot;false\&quot;. | [optional] 
**StartConferenceOnEnter** | Pointer to **bool** | Whether the conference should be started after the participant joins the conference. Defaults to \&quot;false\&quot;. | [optional] 
**SupervisorRole** | Pointer to **string** | Sets the joining participant as a supervisor for the conference. A conference can have multiple supervisors. \&quot;barge\&quot; means the supervisor enters the conference as a normal participant. This is the same as \&quot;none\&quot;. \&quot;monitor\&quot; means the supervisor is muted but can hear all participants. \&quot;whisper\&quot; means that only the specified \&quot;whisper_call_control_ids\&quot; can hear the supervisor. Defaults to \&quot;none\&quot;. | [optional] 
**WhisperCallControlIds** | Pointer to **[]string** | Array of unique call_control_ids the joining supervisor can whisper to. If none provided, the supervisor will join the conference as a monitoring participant only. | [optional] 
**BeepEnabled** | Pointer to **string** | Whether a beep sound should be played when the participant joins and/or leaves the conference. Can be used to override the conference-level setting. | [optional] 

## Methods

### NewJoinConferenceRequest

`func NewJoinConferenceRequest(callControlId string, ) *JoinConferenceRequest`

NewJoinConferenceRequest instantiates a new JoinConferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinConferenceRequestWithDefaults

`func NewJoinConferenceRequestWithDefaults() *JoinConferenceRequest`

NewJoinConferenceRequestWithDefaults instantiates a new JoinConferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *JoinConferenceRequest) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *JoinConferenceRequest) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *JoinConferenceRequest) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetClientState

`func (o *JoinConferenceRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *JoinConferenceRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *JoinConferenceRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *JoinConferenceRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *JoinConferenceRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *JoinConferenceRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *JoinConferenceRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *JoinConferenceRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetEndConferenceOnExit

`func (o *JoinConferenceRequest) GetEndConferenceOnExit() bool`

GetEndConferenceOnExit returns the EndConferenceOnExit field if non-nil, zero value otherwise.

### GetEndConferenceOnExitOk

`func (o *JoinConferenceRequest) GetEndConferenceOnExitOk() (*bool, bool)`

GetEndConferenceOnExitOk returns a tuple with the EndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndConferenceOnExit

`func (o *JoinConferenceRequest) SetEndConferenceOnExit(v bool)`

SetEndConferenceOnExit sets EndConferenceOnExit field to given value.

### HasEndConferenceOnExit

`func (o *JoinConferenceRequest) HasEndConferenceOnExit() bool`

HasEndConferenceOnExit returns a boolean if a field has been set.

### GetSoftEndConferenceOnExit

`func (o *JoinConferenceRequest) GetSoftEndConferenceOnExit() bool`

GetSoftEndConferenceOnExit returns the SoftEndConferenceOnExit field if non-nil, zero value otherwise.

### GetSoftEndConferenceOnExitOk

`func (o *JoinConferenceRequest) GetSoftEndConferenceOnExitOk() (*bool, bool)`

GetSoftEndConferenceOnExitOk returns a tuple with the SoftEndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftEndConferenceOnExit

`func (o *JoinConferenceRequest) SetSoftEndConferenceOnExit(v bool)`

SetSoftEndConferenceOnExit sets SoftEndConferenceOnExit field to given value.

### HasSoftEndConferenceOnExit

`func (o *JoinConferenceRequest) HasSoftEndConferenceOnExit() bool`

HasSoftEndConferenceOnExit returns a boolean if a field has been set.

### GetHold

`func (o *JoinConferenceRequest) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *JoinConferenceRequest) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *JoinConferenceRequest) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *JoinConferenceRequest) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetHoldAudioUrl

`func (o *JoinConferenceRequest) GetHoldAudioUrl() string`

GetHoldAudioUrl returns the HoldAudioUrl field if non-nil, zero value otherwise.

### GetHoldAudioUrlOk

`func (o *JoinConferenceRequest) GetHoldAudioUrlOk() (*string, bool)`

GetHoldAudioUrlOk returns a tuple with the HoldAudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldAudioUrl

`func (o *JoinConferenceRequest) SetHoldAudioUrl(v string)`

SetHoldAudioUrl sets HoldAudioUrl field to given value.

### HasHoldAudioUrl

`func (o *JoinConferenceRequest) HasHoldAudioUrl() bool`

HasHoldAudioUrl returns a boolean if a field has been set.

### GetHoldMediaName

`func (o *JoinConferenceRequest) GetHoldMediaName() string`

GetHoldMediaName returns the HoldMediaName field if non-nil, zero value otherwise.

### GetHoldMediaNameOk

`func (o *JoinConferenceRequest) GetHoldMediaNameOk() (*string, bool)`

GetHoldMediaNameOk returns a tuple with the HoldMediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldMediaName

`func (o *JoinConferenceRequest) SetHoldMediaName(v string)`

SetHoldMediaName sets HoldMediaName field to given value.

### HasHoldMediaName

`func (o *JoinConferenceRequest) HasHoldMediaName() bool`

HasHoldMediaName returns a boolean if a field has been set.

### GetMute

`func (o *JoinConferenceRequest) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *JoinConferenceRequest) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *JoinConferenceRequest) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *JoinConferenceRequest) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetStartConferenceOnEnter

`func (o *JoinConferenceRequest) GetStartConferenceOnEnter() bool`

GetStartConferenceOnEnter returns the StartConferenceOnEnter field if non-nil, zero value otherwise.

### GetStartConferenceOnEnterOk

`func (o *JoinConferenceRequest) GetStartConferenceOnEnterOk() (*bool, bool)`

GetStartConferenceOnEnterOk returns a tuple with the StartConferenceOnEnter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConferenceOnEnter

`func (o *JoinConferenceRequest) SetStartConferenceOnEnter(v bool)`

SetStartConferenceOnEnter sets StartConferenceOnEnter field to given value.

### HasStartConferenceOnEnter

`func (o *JoinConferenceRequest) HasStartConferenceOnEnter() bool`

HasStartConferenceOnEnter returns a boolean if a field has been set.

### GetSupervisorRole

`func (o *JoinConferenceRequest) GetSupervisorRole() string`

GetSupervisorRole returns the SupervisorRole field if non-nil, zero value otherwise.

### GetSupervisorRoleOk

`func (o *JoinConferenceRequest) GetSupervisorRoleOk() (*string, bool)`

GetSupervisorRoleOk returns a tuple with the SupervisorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorRole

`func (o *JoinConferenceRequest) SetSupervisorRole(v string)`

SetSupervisorRole sets SupervisorRole field to given value.

### HasSupervisorRole

`func (o *JoinConferenceRequest) HasSupervisorRole() bool`

HasSupervisorRole returns a boolean if a field has been set.

### GetWhisperCallControlIds

`func (o *JoinConferenceRequest) GetWhisperCallControlIds() []string`

GetWhisperCallControlIds returns the WhisperCallControlIds field if non-nil, zero value otherwise.

### GetWhisperCallControlIdsOk

`func (o *JoinConferenceRequest) GetWhisperCallControlIdsOk() (*[]string, bool)`

GetWhisperCallControlIdsOk returns a tuple with the WhisperCallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhisperCallControlIds

`func (o *JoinConferenceRequest) SetWhisperCallControlIds(v []string)`

SetWhisperCallControlIds sets WhisperCallControlIds field to given value.

### HasWhisperCallControlIds

`func (o *JoinConferenceRequest) HasWhisperCallControlIds() bool`

HasWhisperCallControlIds returns a boolean if a field has been set.

### GetBeepEnabled

`func (o *JoinConferenceRequest) GetBeepEnabled() string`

GetBeepEnabled returns the BeepEnabled field if non-nil, zero value otherwise.

### GetBeepEnabledOk

`func (o *JoinConferenceRequest) GetBeepEnabledOk() (*string, bool)`

GetBeepEnabledOk returns a tuple with the BeepEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeepEnabled

`func (o *JoinConferenceRequest) SetBeepEnabled(v string)`

SetBeepEnabled sets BeepEnabled field to given value.

### HasBeepEnabled

`func (o *JoinConferenceRequest) HasBeepEnabled() bool`

HasBeepEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


