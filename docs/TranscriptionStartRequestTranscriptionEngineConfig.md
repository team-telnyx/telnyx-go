# TranscriptionStartRequestTranscriptionEngineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**TelnyxTranscriptionLanguage**](TelnyxTranscriptionLanguage.md) |  | [optional] [default to EN]
**InterimResults** | Pointer to **bool** | Whether to send also interim results. If set to false, only final results will be sent. | [optional] [default to false]
**EnableSpeakerDiarization** | Pointer to **bool** | Enables speaker diarization. | [optional] [default to false]
**MinSpeakerCount** | Pointer to **int32** | Defines minimum number of speakers in the conversation. | [optional] [default to 2]
**MaxSpeakerCount** | Pointer to **int32** | Defines maximum number of speakers in the conversation. | [optional] [default to 6]
**ProfanityFilter** | Pointer to **bool** | Enables profanity_filter. | [optional] [default to false]
**UseEnhanced** | Pointer to **bool** | Enables enhanced transcription, this works for models &#x60;phone_call&#x60; and &#x60;video&#x60;. | [optional] [default to false]
**Model** | Pointer to **string** | The model to use for transcription. | [optional] 
**Hints** | Pointer to **[]string** | Hints to improve transcription accuracy. | [optional] [default to []]
**SpeechContext** | Pointer to [**[]TranscriptionEngineAConfigSpeechContextInner**](TranscriptionEngineAConfigSpeechContextInner.md) | Speech context to improve transcription accuracy. | [optional] 
**TranscriptionModel** | Pointer to **string** | The model to use for transcription. | [optional] [default to "openai/whisper-tiny"]

## Methods

### NewTranscriptionStartRequestTranscriptionEngineConfig

`func NewTranscriptionStartRequestTranscriptionEngineConfig() *TranscriptionStartRequestTranscriptionEngineConfig`

NewTranscriptionStartRequestTranscriptionEngineConfig instantiates a new TranscriptionStartRequestTranscriptionEngineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionStartRequestTranscriptionEngineConfigWithDefaults

`func NewTranscriptionStartRequestTranscriptionEngineConfigWithDefaults() *TranscriptionStartRequestTranscriptionEngineConfig`

NewTranscriptionStartRequestTranscriptionEngineConfigWithDefaults instantiates a new TranscriptionStartRequestTranscriptionEngineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetLanguage() TelnyxTranscriptionLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetLanguageOk() (*TelnyxTranscriptionLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetLanguage(v TelnyxTranscriptionLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetInterimResults

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetInterimResults() bool`

GetInterimResults returns the InterimResults field if non-nil, zero value otherwise.

### GetInterimResultsOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetInterimResultsOk() (*bool, bool)`

GetInterimResultsOk returns a tuple with the InterimResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimResults

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetInterimResults(v bool)`

SetInterimResults sets InterimResults field to given value.

### HasInterimResults

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasInterimResults() bool`

HasInterimResults returns a boolean if a field has been set.

### GetEnableSpeakerDiarization

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetEnableSpeakerDiarization() bool`

GetEnableSpeakerDiarization returns the EnableSpeakerDiarization field if non-nil, zero value otherwise.

### GetEnableSpeakerDiarizationOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetEnableSpeakerDiarizationOk() (*bool, bool)`

GetEnableSpeakerDiarizationOk returns a tuple with the EnableSpeakerDiarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpeakerDiarization

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetEnableSpeakerDiarization(v bool)`

SetEnableSpeakerDiarization sets EnableSpeakerDiarization field to given value.

### HasEnableSpeakerDiarization

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasEnableSpeakerDiarization() bool`

HasEnableSpeakerDiarization returns a boolean if a field has been set.

### GetMinSpeakerCount

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetMinSpeakerCount() int32`

GetMinSpeakerCount returns the MinSpeakerCount field if non-nil, zero value otherwise.

### GetMinSpeakerCountOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetMinSpeakerCountOk() (*int32, bool)`

GetMinSpeakerCountOk returns a tuple with the MinSpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpeakerCount

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetMinSpeakerCount(v int32)`

SetMinSpeakerCount sets MinSpeakerCount field to given value.

### HasMinSpeakerCount

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasMinSpeakerCount() bool`

HasMinSpeakerCount returns a boolean if a field has been set.

### GetMaxSpeakerCount

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetMaxSpeakerCount() int32`

GetMaxSpeakerCount returns the MaxSpeakerCount field if non-nil, zero value otherwise.

### GetMaxSpeakerCountOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetMaxSpeakerCountOk() (*int32, bool)`

GetMaxSpeakerCountOk returns a tuple with the MaxSpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeakerCount

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetMaxSpeakerCount(v int32)`

SetMaxSpeakerCount sets MaxSpeakerCount field to given value.

### HasMaxSpeakerCount

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasMaxSpeakerCount() bool`

HasMaxSpeakerCount returns a boolean if a field has been set.

### GetProfanityFilter

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetProfanityFilter() bool`

GetProfanityFilter returns the ProfanityFilter field if non-nil, zero value otherwise.

### GetProfanityFilterOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetProfanityFilterOk() (*bool, bool)`

GetProfanityFilterOk returns a tuple with the ProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfanityFilter

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetProfanityFilter(v bool)`

SetProfanityFilter sets ProfanityFilter field to given value.

### HasProfanityFilter

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasProfanityFilter() bool`

HasProfanityFilter returns a boolean if a field has been set.

### GetUseEnhanced

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetUseEnhanced() bool`

GetUseEnhanced returns the UseEnhanced field if non-nil, zero value otherwise.

### GetUseEnhancedOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetUseEnhancedOk() (*bool, bool)`

GetUseEnhancedOk returns a tuple with the UseEnhanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnhanced

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetUseEnhanced(v bool)`

SetUseEnhanced sets UseEnhanced field to given value.

### HasUseEnhanced

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasUseEnhanced() bool`

HasUseEnhanced returns a boolean if a field has been set.

### GetModel

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetHints

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetHints() []string`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetHintsOk() (*[]string, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetHints(v []string)`

SetHints sets Hints field to given value.

### HasHints

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetSpeechContext

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetSpeechContext() []TranscriptionEngineAConfigSpeechContextInner`

GetSpeechContext returns the SpeechContext field if non-nil, zero value otherwise.

### GetSpeechContextOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetSpeechContextOk() (*[]TranscriptionEngineAConfigSpeechContextInner, bool)`

GetSpeechContextOk returns a tuple with the SpeechContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechContext

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetSpeechContext(v []TranscriptionEngineAConfigSpeechContextInner)`

SetSpeechContext sets SpeechContext field to given value.

### HasSpeechContext

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasSpeechContext() bool`

HasSpeechContext returns a boolean if a field has been set.

### GetTranscriptionModel

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetTranscriptionModel() string`

GetTranscriptionModel returns the TranscriptionModel field if non-nil, zero value otherwise.

### GetTranscriptionModelOk

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) GetTranscriptionModelOk() (*string, bool)`

GetTranscriptionModelOk returns a tuple with the TranscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionModel

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) SetTranscriptionModel(v string)`

SetTranscriptionModel sets TranscriptionModel field to given value.

### HasTranscriptionModel

`func (o *TranscriptionStartRequestTranscriptionEngineConfig) HasTranscriptionModel() bool`

HasTranscriptionModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


