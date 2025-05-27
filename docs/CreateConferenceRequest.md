# CreateConferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Unique identifier and token for controlling the call | 
**Name** | **string** | Name of the conference | 
**BeepEnabled** | Pointer to **string** | Whether a beep sound should be played when participants join and/or leave the conference. | [optional] [default to "never"]
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. The client_state will be updated for the creator call leg and will be used for all webhooks related to the created conference. | [optional] 
**ComfortNoise** | Pointer to **bool** | Toggle background comfort noise. | [optional] [default to true]
**CommandId** | Pointer to **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] 
**DurationMinutes** | Pointer to **int32** | Time length (minutes) after which the conference will end. | [optional] 
**HoldAudioUrl** | Pointer to **string** | The URL of a file to be played to participants joining the conference. The URL can point to either a WAV or MP3 file. hold_media_name and hold_audio_url cannot be used together in one request. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. | [optional] 
**HoldMediaName** | Pointer to **string** | The media_name of a file to be played to participants joining the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. Takes effect only when \&quot;start_conference_on_create\&quot; is set to \&quot;false\&quot;. | [optional] 
**MaxParticipants** | Pointer to **int32** | The maximum number of active conference participants to allow. Must be between 2 and 800. Defaults to 250 | [optional] 
**StartConferenceOnCreate** | Pointer to **bool** | Whether the conference should be started on creation. If the conference isn&#39;t started all participants that join are automatically put on hold. Defaults to \&quot;true\&quot;. | [optional] 

## Methods

### NewCreateConferenceRequest

`func NewCreateConferenceRequest(callControlId string, name string, ) *CreateConferenceRequest`

NewCreateConferenceRequest instantiates a new CreateConferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConferenceRequestWithDefaults

`func NewCreateConferenceRequestWithDefaults() *CreateConferenceRequest`

NewCreateConferenceRequestWithDefaults instantiates a new CreateConferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CreateConferenceRequest) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CreateConferenceRequest) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CreateConferenceRequest) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetName

`func (o *CreateConferenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConferenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConferenceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBeepEnabled

`func (o *CreateConferenceRequest) GetBeepEnabled() string`

GetBeepEnabled returns the BeepEnabled field if non-nil, zero value otherwise.

### GetBeepEnabledOk

`func (o *CreateConferenceRequest) GetBeepEnabledOk() (*string, bool)`

GetBeepEnabledOk returns a tuple with the BeepEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeepEnabled

`func (o *CreateConferenceRequest) SetBeepEnabled(v string)`

SetBeepEnabled sets BeepEnabled field to given value.

### HasBeepEnabled

`func (o *CreateConferenceRequest) HasBeepEnabled() bool`

HasBeepEnabled returns a boolean if a field has been set.

### GetClientState

`func (o *CreateConferenceRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CreateConferenceRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CreateConferenceRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CreateConferenceRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetComfortNoise

`func (o *CreateConferenceRequest) GetComfortNoise() bool`

GetComfortNoise returns the ComfortNoise field if non-nil, zero value otherwise.

### GetComfortNoiseOk

`func (o *CreateConferenceRequest) GetComfortNoiseOk() (*bool, bool)`

GetComfortNoiseOk returns a tuple with the ComfortNoise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComfortNoise

`func (o *CreateConferenceRequest) SetComfortNoise(v bool)`

SetComfortNoise sets ComfortNoise field to given value.

### HasComfortNoise

`func (o *CreateConferenceRequest) HasComfortNoise() bool`

HasComfortNoise returns a boolean if a field has been set.

### GetCommandId

`func (o *CreateConferenceRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *CreateConferenceRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *CreateConferenceRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *CreateConferenceRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetDurationMinutes

`func (o *CreateConferenceRequest) GetDurationMinutes() int32`

GetDurationMinutes returns the DurationMinutes field if non-nil, zero value otherwise.

### GetDurationMinutesOk

`func (o *CreateConferenceRequest) GetDurationMinutesOk() (*int32, bool)`

GetDurationMinutesOk returns a tuple with the DurationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMinutes

`func (o *CreateConferenceRequest) SetDurationMinutes(v int32)`

SetDurationMinutes sets DurationMinutes field to given value.

### HasDurationMinutes

`func (o *CreateConferenceRequest) HasDurationMinutes() bool`

HasDurationMinutes returns a boolean if a field has been set.

### GetHoldAudioUrl

`func (o *CreateConferenceRequest) GetHoldAudioUrl() string`

GetHoldAudioUrl returns the HoldAudioUrl field if non-nil, zero value otherwise.

### GetHoldAudioUrlOk

`func (o *CreateConferenceRequest) GetHoldAudioUrlOk() (*string, bool)`

GetHoldAudioUrlOk returns a tuple with the HoldAudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldAudioUrl

`func (o *CreateConferenceRequest) SetHoldAudioUrl(v string)`

SetHoldAudioUrl sets HoldAudioUrl field to given value.

### HasHoldAudioUrl

`func (o *CreateConferenceRequest) HasHoldAudioUrl() bool`

HasHoldAudioUrl returns a boolean if a field has been set.

### GetHoldMediaName

`func (o *CreateConferenceRequest) GetHoldMediaName() string`

GetHoldMediaName returns the HoldMediaName field if non-nil, zero value otherwise.

### GetHoldMediaNameOk

`func (o *CreateConferenceRequest) GetHoldMediaNameOk() (*string, bool)`

GetHoldMediaNameOk returns a tuple with the HoldMediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldMediaName

`func (o *CreateConferenceRequest) SetHoldMediaName(v string)`

SetHoldMediaName sets HoldMediaName field to given value.

### HasHoldMediaName

`func (o *CreateConferenceRequest) HasHoldMediaName() bool`

HasHoldMediaName returns a boolean if a field has been set.

### GetMaxParticipants

`func (o *CreateConferenceRequest) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *CreateConferenceRequest) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *CreateConferenceRequest) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *CreateConferenceRequest) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### GetStartConferenceOnCreate

`func (o *CreateConferenceRequest) GetStartConferenceOnCreate() bool`

GetStartConferenceOnCreate returns the StartConferenceOnCreate field if non-nil, zero value otherwise.

### GetStartConferenceOnCreateOk

`func (o *CreateConferenceRequest) GetStartConferenceOnCreateOk() (*bool, bool)`

GetStartConferenceOnCreateOk returns a tuple with the StartConferenceOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConferenceOnCreate

`func (o *CreateConferenceRequest) SetStartConferenceOnCreate(v bool)`

SetStartConferenceOnCreate sets StartConferenceOnCreate field to given value.

### HasStartConferenceOnCreate

`func (o *CreateConferenceRequest) HasStartConferenceOnCreate() bool`

HasStartConferenceOnCreate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


