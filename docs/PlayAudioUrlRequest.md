# PlayAudioUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | Pointer to **string** | The URL of a file to be played back on the call. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request. | [optional] 
**MediaName** | Pointer to **string** | The media_name of a file to be played back on the call. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 
**Loop** | Pointer to [**Loopcount**](Loopcount.md) |  | [optional] 
**Overlay** | Pointer to **bool** | When enabled, audio will be mixed on top of any other audio that is actively being played back. Note that &#x60;overlay: true&#x60; will only work if there is another audio file already being played on the call. | [optional] [default to false]
**Stop** | Pointer to **string** | When specified, it stops the current audio being played. Specify &#x60;current&#x60; to stop the current audio being played, and to play the next file in the queue. Specify &#x60;all&#x60; to stop the current audio file being played and to also clear all audio files from the queue. | [optional] 
**TargetLegs** | Pointer to **string** | Specifies the leg or legs on which audio will be played. If supplied, the value must be either &#x60;self&#x60;, &#x60;opposite&#x60; or &#x60;both&#x60;. | [optional] [default to "self"]
**CacheAudio** | Pointer to **bool** | Caches the audio file. Useful when playing the same audio file multiple times during the call. | [optional] [default to true]
**AudioType** | Pointer to **string** | Specifies the type of audio provided in &#x60;audio_url&#x60; or &#x60;playback_content&#x60;. | [optional] [default to "mp3"]
**PlaybackContent** | Pointer to **string** | Allows a user to provide base64 encoded mp3 or wav. Note: when using this parameter, &#x60;media_url&#x60; and &#x60;media_name&#x60; in the &#x60;playback_started&#x60; and &#x60;playback_ended&#x60; webhooks will be empty | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewPlayAudioUrlRequest

`func NewPlayAudioUrlRequest() *PlayAudioUrlRequest`

NewPlayAudioUrlRequest instantiates a new PlayAudioUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayAudioUrlRequestWithDefaults

`func NewPlayAudioUrlRequestWithDefaults() *PlayAudioUrlRequest`

NewPlayAudioUrlRequestWithDefaults instantiates a new PlayAudioUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *PlayAudioUrlRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *PlayAudioUrlRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *PlayAudioUrlRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *PlayAudioUrlRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *PlayAudioUrlRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *PlayAudioUrlRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *PlayAudioUrlRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *PlayAudioUrlRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetLoop

`func (o *PlayAudioUrlRequest) GetLoop() Loopcount`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *PlayAudioUrlRequest) GetLoopOk() (*Loopcount, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *PlayAudioUrlRequest) SetLoop(v Loopcount)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *PlayAudioUrlRequest) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetOverlay

`func (o *PlayAudioUrlRequest) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *PlayAudioUrlRequest) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *PlayAudioUrlRequest) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *PlayAudioUrlRequest) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetStop

`func (o *PlayAudioUrlRequest) GetStop() string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *PlayAudioUrlRequest) GetStopOk() (*string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *PlayAudioUrlRequest) SetStop(v string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *PlayAudioUrlRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetTargetLegs

`func (o *PlayAudioUrlRequest) GetTargetLegs() string`

GetTargetLegs returns the TargetLegs field if non-nil, zero value otherwise.

### GetTargetLegsOk

`func (o *PlayAudioUrlRequest) GetTargetLegsOk() (*string, bool)`

GetTargetLegsOk returns a tuple with the TargetLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLegs

`func (o *PlayAudioUrlRequest) SetTargetLegs(v string)`

SetTargetLegs sets TargetLegs field to given value.

### HasTargetLegs

`func (o *PlayAudioUrlRequest) HasTargetLegs() bool`

HasTargetLegs returns a boolean if a field has been set.

### GetCacheAudio

`func (o *PlayAudioUrlRequest) GetCacheAudio() bool`

GetCacheAudio returns the CacheAudio field if non-nil, zero value otherwise.

### GetCacheAudioOk

`func (o *PlayAudioUrlRequest) GetCacheAudioOk() (*bool, bool)`

GetCacheAudioOk returns a tuple with the CacheAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAudio

`func (o *PlayAudioUrlRequest) SetCacheAudio(v bool)`

SetCacheAudio sets CacheAudio field to given value.

### HasCacheAudio

`func (o *PlayAudioUrlRequest) HasCacheAudio() bool`

HasCacheAudio returns a boolean if a field has been set.

### GetAudioType

`func (o *PlayAudioUrlRequest) GetAudioType() string`

GetAudioType returns the AudioType field if non-nil, zero value otherwise.

### GetAudioTypeOk

`func (o *PlayAudioUrlRequest) GetAudioTypeOk() (*string, bool)`

GetAudioTypeOk returns a tuple with the AudioType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioType

`func (o *PlayAudioUrlRequest) SetAudioType(v string)`

SetAudioType sets AudioType field to given value.

### HasAudioType

`func (o *PlayAudioUrlRequest) HasAudioType() bool`

HasAudioType returns a boolean if a field has been set.

### GetPlaybackContent

`func (o *PlayAudioUrlRequest) GetPlaybackContent() string`

GetPlaybackContent returns the PlaybackContent field if non-nil, zero value otherwise.

### GetPlaybackContentOk

`func (o *PlayAudioUrlRequest) GetPlaybackContentOk() (*string, bool)`

GetPlaybackContentOk returns a tuple with the PlaybackContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackContent

`func (o *PlayAudioUrlRequest) SetPlaybackContent(v string)`

SetPlaybackContent sets PlaybackContent field to given value.

### HasPlaybackContent

`func (o *PlayAudioUrlRequest) HasPlaybackContent() bool`

HasPlaybackContent returns a boolean if a field has been set.

### GetClientState

`func (o *PlayAudioUrlRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *PlayAudioUrlRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *PlayAudioUrlRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *PlayAudioUrlRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *PlayAudioUrlRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *PlayAudioUrlRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *PlayAudioUrlRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *PlayAudioUrlRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


