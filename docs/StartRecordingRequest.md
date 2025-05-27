# StartRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** | The audio file format used when storing the call recording. Can be either &#x60;mp3&#x60; or &#x60;wav&#x60;. | 
**Channels** | **string** | When &#x60;dual&#x60;, final audio file will be stereo recorded with the first leg on channel A, and the rest on channel B. | 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**PlayBeep** | Pointer to **bool** | If enabled, a beep sound will be played at the start of a recording. | [optional] 
**MaxLength** | Pointer to **int32** | Defines the maximum length for the recording in seconds. The minimum value is 0. The maximum value is 14400. The default value is 0 (infinite) | [optional] [default to 0]
**TimeoutSecs** | Pointer to **int32** | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that call transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite) | [optional] [default to 0]
**RecordingTrack** | Pointer to **string** | The audio track to be recorded. Can be either &#x60;both&#x60;, &#x60;inbound&#x60; or &#x60;outbound&#x60;. If only single track is specified (&#x60;inbound&#x60;, &#x60;outbound&#x60;), &#x60;channels&#x60; configuration is ignored and it will be recorded as mono (single channel). | [optional] [default to "both"]
**Trim** | Pointer to **string** | When set to &#x60;trim-silence&#x60;, silence will be removed from the beginning and end of the recording. | [optional] 
**CustomFileName** | Pointer to **string** | The custom recording file name to be used instead of the default &#x60;call_leg_id&#x60;. Telnyx will still add a Unix timestamp suffix. | [optional] 
**Transcription** | Pointer to **bool** | Enable post recording transcription. The default value is false. | [optional] [default to false]
**TranscriptionEngine** | Pointer to **string** | Engine to use for speech recognition. &#x60;A&#x60; - &#x60;Google&#x60; | [optional] [default to "A"]
**TranscriptionLanguage** | Pointer to [**GoogleTranscriptionLanguageLong**](GoogleTranscriptionLanguageLong.md) |  | [optional] [default to EN_US]
**TranscriptionProfanityFilter** | Pointer to **bool** | Enables profanity_filter. Applies to &#x60;google&#x60; engine only. | [optional] [default to false]
**TranscriptionSpeakerDiarization** | Pointer to **bool** | Enables speaker diarization. Applies to &#x60;google&#x60; engine only. | [optional] [default to false]
**TranscriptionMinSpeakerCount** | Pointer to **int32** | Defines minimum number of speakers in the conversation. Applies to &#x60;google&#x60; engine only. | [optional] [default to 2]
**TranscriptionMaxSpeakerCount** | Pointer to **int32** | Defines maximum number of speakers in the conversation. Applies to &#x60;google&#x60; engine only. | [optional] [default to 6]

## Methods

### NewStartRecordingRequest

`func NewStartRecordingRequest(format string, channels string, ) *StartRecordingRequest`

NewStartRecordingRequest instantiates a new StartRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartRecordingRequestWithDefaults

`func NewStartRecordingRequestWithDefaults() *StartRecordingRequest`

NewStartRecordingRequestWithDefaults instantiates a new StartRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *StartRecordingRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StartRecordingRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StartRecordingRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetChannels

`func (o *StartRecordingRequest) GetChannels() string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *StartRecordingRequest) GetChannelsOk() (*string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *StartRecordingRequest) SetChannels(v string)`

SetChannels sets Channels field to given value.


### GetClientState

`func (o *StartRecordingRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *StartRecordingRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *StartRecordingRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *StartRecordingRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *StartRecordingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *StartRecordingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *StartRecordingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *StartRecordingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetPlayBeep

`func (o *StartRecordingRequest) GetPlayBeep() bool`

GetPlayBeep returns the PlayBeep field if non-nil, zero value otherwise.

### GetPlayBeepOk

`func (o *StartRecordingRequest) GetPlayBeepOk() (*bool, bool)`

GetPlayBeepOk returns a tuple with the PlayBeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayBeep

`func (o *StartRecordingRequest) SetPlayBeep(v bool)`

SetPlayBeep sets PlayBeep field to given value.

### HasPlayBeep

`func (o *StartRecordingRequest) HasPlayBeep() bool`

HasPlayBeep returns a boolean if a field has been set.

### GetMaxLength

`func (o *StartRecordingRequest) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *StartRecordingRequest) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *StartRecordingRequest) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *StartRecordingRequest) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *StartRecordingRequest) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *StartRecordingRequest) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *StartRecordingRequest) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *StartRecordingRequest) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetRecordingTrack

`func (o *StartRecordingRequest) GetRecordingTrack() string`

GetRecordingTrack returns the RecordingTrack field if non-nil, zero value otherwise.

### GetRecordingTrackOk

`func (o *StartRecordingRequest) GetRecordingTrackOk() (*string, bool)`

GetRecordingTrackOk returns a tuple with the RecordingTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTrack

`func (o *StartRecordingRequest) SetRecordingTrack(v string)`

SetRecordingTrack sets RecordingTrack field to given value.

### HasRecordingTrack

`func (o *StartRecordingRequest) HasRecordingTrack() bool`

HasRecordingTrack returns a boolean if a field has been set.

### GetTrim

`func (o *StartRecordingRequest) GetTrim() string`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *StartRecordingRequest) GetTrimOk() (*string, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *StartRecordingRequest) SetTrim(v string)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *StartRecordingRequest) HasTrim() bool`

HasTrim returns a boolean if a field has been set.

### GetCustomFileName

`func (o *StartRecordingRequest) GetCustomFileName() string`

GetCustomFileName returns the CustomFileName field if non-nil, zero value otherwise.

### GetCustomFileNameOk

`func (o *StartRecordingRequest) GetCustomFileNameOk() (*string, bool)`

GetCustomFileNameOk returns a tuple with the CustomFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFileName

`func (o *StartRecordingRequest) SetCustomFileName(v string)`

SetCustomFileName sets CustomFileName field to given value.

### HasCustomFileName

`func (o *StartRecordingRequest) HasCustomFileName() bool`

HasCustomFileName returns a boolean if a field has been set.

### GetTranscription

`func (o *StartRecordingRequest) GetTranscription() bool`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *StartRecordingRequest) GetTranscriptionOk() (*bool, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *StartRecordingRequest) SetTranscription(v bool)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *StartRecordingRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetTranscriptionEngine

`func (o *StartRecordingRequest) GetTranscriptionEngine() string`

GetTranscriptionEngine returns the TranscriptionEngine field if non-nil, zero value otherwise.

### GetTranscriptionEngineOk

`func (o *StartRecordingRequest) GetTranscriptionEngineOk() (*string, bool)`

GetTranscriptionEngineOk returns a tuple with the TranscriptionEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionEngine

`func (o *StartRecordingRequest) SetTranscriptionEngine(v string)`

SetTranscriptionEngine sets TranscriptionEngine field to given value.

### HasTranscriptionEngine

`func (o *StartRecordingRequest) HasTranscriptionEngine() bool`

HasTranscriptionEngine returns a boolean if a field has been set.

### GetTranscriptionLanguage

`func (o *StartRecordingRequest) GetTranscriptionLanguage() GoogleTranscriptionLanguageLong`

GetTranscriptionLanguage returns the TranscriptionLanguage field if non-nil, zero value otherwise.

### GetTranscriptionLanguageOk

`func (o *StartRecordingRequest) GetTranscriptionLanguageOk() (*GoogleTranscriptionLanguageLong, bool)`

GetTranscriptionLanguageOk returns a tuple with the TranscriptionLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionLanguage

`func (o *StartRecordingRequest) SetTranscriptionLanguage(v GoogleTranscriptionLanguageLong)`

SetTranscriptionLanguage sets TranscriptionLanguage field to given value.

### HasTranscriptionLanguage

`func (o *StartRecordingRequest) HasTranscriptionLanguage() bool`

HasTranscriptionLanguage returns a boolean if a field has been set.

### GetTranscriptionProfanityFilter

`func (o *StartRecordingRequest) GetTranscriptionProfanityFilter() bool`

GetTranscriptionProfanityFilter returns the TranscriptionProfanityFilter field if non-nil, zero value otherwise.

### GetTranscriptionProfanityFilterOk

`func (o *StartRecordingRequest) GetTranscriptionProfanityFilterOk() (*bool, bool)`

GetTranscriptionProfanityFilterOk returns a tuple with the TranscriptionProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionProfanityFilter

`func (o *StartRecordingRequest) SetTranscriptionProfanityFilter(v bool)`

SetTranscriptionProfanityFilter sets TranscriptionProfanityFilter field to given value.

### HasTranscriptionProfanityFilter

`func (o *StartRecordingRequest) HasTranscriptionProfanityFilter() bool`

HasTranscriptionProfanityFilter returns a boolean if a field has been set.

### GetTranscriptionSpeakerDiarization

`func (o *StartRecordingRequest) GetTranscriptionSpeakerDiarization() bool`

GetTranscriptionSpeakerDiarization returns the TranscriptionSpeakerDiarization field if non-nil, zero value otherwise.

### GetTranscriptionSpeakerDiarizationOk

`func (o *StartRecordingRequest) GetTranscriptionSpeakerDiarizationOk() (*bool, bool)`

GetTranscriptionSpeakerDiarizationOk returns a tuple with the TranscriptionSpeakerDiarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionSpeakerDiarization

`func (o *StartRecordingRequest) SetTranscriptionSpeakerDiarization(v bool)`

SetTranscriptionSpeakerDiarization sets TranscriptionSpeakerDiarization field to given value.

### HasTranscriptionSpeakerDiarization

`func (o *StartRecordingRequest) HasTranscriptionSpeakerDiarization() bool`

HasTranscriptionSpeakerDiarization returns a boolean if a field has been set.

### GetTranscriptionMinSpeakerCount

`func (o *StartRecordingRequest) GetTranscriptionMinSpeakerCount() int32`

GetTranscriptionMinSpeakerCount returns the TranscriptionMinSpeakerCount field if non-nil, zero value otherwise.

### GetTranscriptionMinSpeakerCountOk

`func (o *StartRecordingRequest) GetTranscriptionMinSpeakerCountOk() (*int32, bool)`

GetTranscriptionMinSpeakerCountOk returns a tuple with the TranscriptionMinSpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionMinSpeakerCount

`func (o *StartRecordingRequest) SetTranscriptionMinSpeakerCount(v int32)`

SetTranscriptionMinSpeakerCount sets TranscriptionMinSpeakerCount field to given value.

### HasTranscriptionMinSpeakerCount

`func (o *StartRecordingRequest) HasTranscriptionMinSpeakerCount() bool`

HasTranscriptionMinSpeakerCount returns a boolean if a field has been set.

### GetTranscriptionMaxSpeakerCount

`func (o *StartRecordingRequest) GetTranscriptionMaxSpeakerCount() int32`

GetTranscriptionMaxSpeakerCount returns the TranscriptionMaxSpeakerCount field if non-nil, zero value otherwise.

### GetTranscriptionMaxSpeakerCountOk

`func (o *StartRecordingRequest) GetTranscriptionMaxSpeakerCountOk() (*int32, bool)`

GetTranscriptionMaxSpeakerCountOk returns a tuple with the TranscriptionMaxSpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionMaxSpeakerCount

`func (o *StartRecordingRequest) SetTranscriptionMaxSpeakerCount(v int32)`

SetTranscriptionMaxSpeakerCount sets TranscriptionMaxSpeakerCount field to given value.

### HasTranscriptionMaxSpeakerCount

`func (o *StartRecordingRequest) HasTranscriptionMaxSpeakerCount() bool`

HasTranscriptionMaxSpeakerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


