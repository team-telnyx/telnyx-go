# CallRequestConferenceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Conference ID to be joined | [optional] 
**ConferenceName** | Pointer to **string** | Conference name to be joined | [optional] 
**EarlyMedia** | Pointer to **bool** | Controls the moment when dialled call is joined into conference. If set to &#x60;true&#x60; user will be joined as soon as media is available (ringback). If &#x60;false&#x60; user will be joined when call is answered. Defaults to &#x60;true&#x60; | [optional] [default to true]
**EndConferenceOnExit** | Pointer to **bool** | Whether the conference should end and all remaining participants be hung up after the participant leaves the conference. Defaults to \&quot;false\&quot;. | [optional] 
**SoftEndConferenceOnExit** | Pointer to **bool** | Whether the conference should end after the participant leaves the conference. NOTE this doesn&#39;t hang up the other participants. Defaults to \&quot;false\&quot;. | [optional] 
**Hold** | Pointer to **bool** | Whether the participant should be put on hold immediately after joining the conference. Defaults to \&quot;false\&quot;. | [optional] 
**HoldAudioUrl** | Pointer to **string** | The URL of a file to be played to the participant when they are put on hold after joining the conference. hold_media_name and hold_audio_url cannot be used together in one request. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. This property takes effect only if \&quot;hold\&quot; is set to \&quot;true\&quot;. | [optional] 
**HoldMediaName** | Pointer to **string** | The media_name of a file to be played to the participant when they are put on hold after joining the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. This property takes effect only if \&quot;hold\&quot; is set to \&quot;true\&quot;. | [optional] 
**Mute** | Pointer to **bool** | Whether the participant should be muted immediately after joining the conference. Defaults to \&quot;false\&quot;. | [optional] 
**StartConferenceOnEnter** | Pointer to **bool** | Whether the conference should be started after the participant joins the conference. Defaults to \&quot;false\&quot;. | [optional] 
**StartConferenceOnCreate** | Pointer to **bool** | Whether the conference should be started on creation. If the conference isn&#39;t started all participants that join are automatically put on hold. Defaults to \&quot;true\&quot;. | [optional] 
**SupervisorRole** | Pointer to **string** | Sets the joining participant as a supervisor for the conference. A conference can have multiple supervisors. \&quot;barge\&quot; means the supervisor enters the conference as a normal participant. This is the same as \&quot;none\&quot;. \&quot;monitor\&quot; means the supervisor is muted but can hear all participants. \&quot;whisper\&quot; means that only the specified \&quot;whisper_call_control_ids\&quot; can hear the supervisor. Defaults to \&quot;none\&quot;. | [optional] 
**WhisperCallControlIds** | Pointer to **[]string** | Array of unique call_control_ids the joining supervisor can whisper to. If none provided, the supervisor will join the conference as a monitoring participant only. | [optional] 
**BeepEnabled** | Pointer to **string** | Whether a beep sound should be played when the participant joins and/or leaves the conference. Can be used to override the conference-level setting. | [optional] 

## Methods

### NewCallRequestConferenceConfig

`func NewCallRequestConferenceConfig() *CallRequestConferenceConfig`

NewCallRequestConferenceConfig instantiates a new CallRequestConferenceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestConferenceConfigWithDefaults

`func NewCallRequestConferenceConfigWithDefaults() *CallRequestConferenceConfig`

NewCallRequestConferenceConfigWithDefaults instantiates a new CallRequestConferenceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CallRequestConferenceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallRequestConferenceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallRequestConferenceConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallRequestConferenceConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConferenceName

`func (o *CallRequestConferenceConfig) GetConferenceName() string`

GetConferenceName returns the ConferenceName field if non-nil, zero value otherwise.

### GetConferenceNameOk

`func (o *CallRequestConferenceConfig) GetConferenceNameOk() (*string, bool)`

GetConferenceNameOk returns a tuple with the ConferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceName

`func (o *CallRequestConferenceConfig) SetConferenceName(v string)`

SetConferenceName sets ConferenceName field to given value.

### HasConferenceName

`func (o *CallRequestConferenceConfig) HasConferenceName() bool`

HasConferenceName returns a boolean if a field has been set.

### GetEarlyMedia

`func (o *CallRequestConferenceConfig) GetEarlyMedia() bool`

GetEarlyMedia returns the EarlyMedia field if non-nil, zero value otherwise.

### GetEarlyMediaOk

`func (o *CallRequestConferenceConfig) GetEarlyMediaOk() (*bool, bool)`

GetEarlyMediaOk returns a tuple with the EarlyMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyMedia

`func (o *CallRequestConferenceConfig) SetEarlyMedia(v bool)`

SetEarlyMedia sets EarlyMedia field to given value.

### HasEarlyMedia

`func (o *CallRequestConferenceConfig) HasEarlyMedia() bool`

HasEarlyMedia returns a boolean if a field has been set.

### GetEndConferenceOnExit

`func (o *CallRequestConferenceConfig) GetEndConferenceOnExit() bool`

GetEndConferenceOnExit returns the EndConferenceOnExit field if non-nil, zero value otherwise.

### GetEndConferenceOnExitOk

`func (o *CallRequestConferenceConfig) GetEndConferenceOnExitOk() (*bool, bool)`

GetEndConferenceOnExitOk returns a tuple with the EndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndConferenceOnExit

`func (o *CallRequestConferenceConfig) SetEndConferenceOnExit(v bool)`

SetEndConferenceOnExit sets EndConferenceOnExit field to given value.

### HasEndConferenceOnExit

`func (o *CallRequestConferenceConfig) HasEndConferenceOnExit() bool`

HasEndConferenceOnExit returns a boolean if a field has been set.

### GetSoftEndConferenceOnExit

`func (o *CallRequestConferenceConfig) GetSoftEndConferenceOnExit() bool`

GetSoftEndConferenceOnExit returns the SoftEndConferenceOnExit field if non-nil, zero value otherwise.

### GetSoftEndConferenceOnExitOk

`func (o *CallRequestConferenceConfig) GetSoftEndConferenceOnExitOk() (*bool, bool)`

GetSoftEndConferenceOnExitOk returns a tuple with the SoftEndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftEndConferenceOnExit

`func (o *CallRequestConferenceConfig) SetSoftEndConferenceOnExit(v bool)`

SetSoftEndConferenceOnExit sets SoftEndConferenceOnExit field to given value.

### HasSoftEndConferenceOnExit

`func (o *CallRequestConferenceConfig) HasSoftEndConferenceOnExit() bool`

HasSoftEndConferenceOnExit returns a boolean if a field has been set.

### GetHold

`func (o *CallRequestConferenceConfig) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *CallRequestConferenceConfig) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *CallRequestConferenceConfig) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *CallRequestConferenceConfig) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetHoldAudioUrl

`func (o *CallRequestConferenceConfig) GetHoldAudioUrl() string`

GetHoldAudioUrl returns the HoldAudioUrl field if non-nil, zero value otherwise.

### GetHoldAudioUrlOk

`func (o *CallRequestConferenceConfig) GetHoldAudioUrlOk() (*string, bool)`

GetHoldAudioUrlOk returns a tuple with the HoldAudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldAudioUrl

`func (o *CallRequestConferenceConfig) SetHoldAudioUrl(v string)`

SetHoldAudioUrl sets HoldAudioUrl field to given value.

### HasHoldAudioUrl

`func (o *CallRequestConferenceConfig) HasHoldAudioUrl() bool`

HasHoldAudioUrl returns a boolean if a field has been set.

### GetHoldMediaName

`func (o *CallRequestConferenceConfig) GetHoldMediaName() string`

GetHoldMediaName returns the HoldMediaName field if non-nil, zero value otherwise.

### GetHoldMediaNameOk

`func (o *CallRequestConferenceConfig) GetHoldMediaNameOk() (*string, bool)`

GetHoldMediaNameOk returns a tuple with the HoldMediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldMediaName

`func (o *CallRequestConferenceConfig) SetHoldMediaName(v string)`

SetHoldMediaName sets HoldMediaName field to given value.

### HasHoldMediaName

`func (o *CallRequestConferenceConfig) HasHoldMediaName() bool`

HasHoldMediaName returns a boolean if a field has been set.

### GetMute

`func (o *CallRequestConferenceConfig) GetMute() bool`

GetMute returns the Mute field if non-nil, zero value otherwise.

### GetMuteOk

`func (o *CallRequestConferenceConfig) GetMuteOk() (*bool, bool)`

GetMuteOk returns a tuple with the Mute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMute

`func (o *CallRequestConferenceConfig) SetMute(v bool)`

SetMute sets Mute field to given value.

### HasMute

`func (o *CallRequestConferenceConfig) HasMute() bool`

HasMute returns a boolean if a field has been set.

### GetStartConferenceOnEnter

`func (o *CallRequestConferenceConfig) GetStartConferenceOnEnter() bool`

GetStartConferenceOnEnter returns the StartConferenceOnEnter field if non-nil, zero value otherwise.

### GetStartConferenceOnEnterOk

`func (o *CallRequestConferenceConfig) GetStartConferenceOnEnterOk() (*bool, bool)`

GetStartConferenceOnEnterOk returns a tuple with the StartConferenceOnEnter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConferenceOnEnter

`func (o *CallRequestConferenceConfig) SetStartConferenceOnEnter(v bool)`

SetStartConferenceOnEnter sets StartConferenceOnEnter field to given value.

### HasStartConferenceOnEnter

`func (o *CallRequestConferenceConfig) HasStartConferenceOnEnter() bool`

HasStartConferenceOnEnter returns a boolean if a field has been set.

### GetStartConferenceOnCreate

`func (o *CallRequestConferenceConfig) GetStartConferenceOnCreate() bool`

GetStartConferenceOnCreate returns the StartConferenceOnCreate field if non-nil, zero value otherwise.

### GetStartConferenceOnCreateOk

`func (o *CallRequestConferenceConfig) GetStartConferenceOnCreateOk() (*bool, bool)`

GetStartConferenceOnCreateOk returns a tuple with the StartConferenceOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConferenceOnCreate

`func (o *CallRequestConferenceConfig) SetStartConferenceOnCreate(v bool)`

SetStartConferenceOnCreate sets StartConferenceOnCreate field to given value.

### HasStartConferenceOnCreate

`func (o *CallRequestConferenceConfig) HasStartConferenceOnCreate() bool`

HasStartConferenceOnCreate returns a boolean if a field has been set.

### GetSupervisorRole

`func (o *CallRequestConferenceConfig) GetSupervisorRole() string`

GetSupervisorRole returns the SupervisorRole field if non-nil, zero value otherwise.

### GetSupervisorRoleOk

`func (o *CallRequestConferenceConfig) GetSupervisorRoleOk() (*string, bool)`

GetSupervisorRoleOk returns a tuple with the SupervisorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorRole

`func (o *CallRequestConferenceConfig) SetSupervisorRole(v string)`

SetSupervisorRole sets SupervisorRole field to given value.

### HasSupervisorRole

`func (o *CallRequestConferenceConfig) HasSupervisorRole() bool`

HasSupervisorRole returns a boolean if a field has been set.

### GetWhisperCallControlIds

`func (o *CallRequestConferenceConfig) GetWhisperCallControlIds() []string`

GetWhisperCallControlIds returns the WhisperCallControlIds field if non-nil, zero value otherwise.

### GetWhisperCallControlIdsOk

`func (o *CallRequestConferenceConfig) GetWhisperCallControlIdsOk() (*[]string, bool)`

GetWhisperCallControlIdsOk returns a tuple with the WhisperCallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhisperCallControlIds

`func (o *CallRequestConferenceConfig) SetWhisperCallControlIds(v []string)`

SetWhisperCallControlIds sets WhisperCallControlIds field to given value.

### HasWhisperCallControlIds

`func (o *CallRequestConferenceConfig) HasWhisperCallControlIds() bool`

HasWhisperCallControlIds returns a boolean if a field has been set.

### GetBeepEnabled

`func (o *CallRequestConferenceConfig) GetBeepEnabled() string`

GetBeepEnabled returns the BeepEnabled field if non-nil, zero value otherwise.

### GetBeepEnabledOk

`func (o *CallRequestConferenceConfig) GetBeepEnabledOk() (*string, bool)`

GetBeepEnabledOk returns a tuple with the BeepEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeepEnabled

`func (o *CallRequestConferenceConfig) SetBeepEnabled(v string)`

SetBeepEnabled sets BeepEnabled field to given value.

### HasBeepEnabled

`func (o *CallRequestConferenceConfig) HasBeepEnabled() bool`

HasBeepEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


