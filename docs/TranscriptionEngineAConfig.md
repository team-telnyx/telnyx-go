# TranscriptionEngineAConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**GoogleTranscriptionLanguage**](GoogleTranscriptionLanguage.md) |  | [optional] [default to EN]
**InterimResults** | Pointer to **bool** | Whether to send also interim results. If set to false, only final results will be sent. | [optional] [default to false]
**EnableSpeakerDiarization** | Pointer to **bool** | Enables speaker diarization. | [optional] [default to false]
**MinSpeakerCount** | Pointer to **int32** | Defines minimum number of speakers in the conversation. | [optional] [default to 2]
**MaxSpeakerCount** | Pointer to **int32** | Defines maximum number of speakers in the conversation. | [optional] [default to 6]
**ProfanityFilter** | Pointer to **bool** | Enables profanity_filter. | [optional] [default to false]
**UseEnhanced** | Pointer to **bool** | Enables enhanced transcription, this works for models &#x60;phone_call&#x60; and &#x60;video&#x60;. | [optional] [default to false]
**Model** | Pointer to **string** | The model to use for transcription. | [optional] 
**Hints** | Pointer to **[]string** | Hints to improve transcription accuracy. | [optional] [default to []]
**SpeechContext** | Pointer to [**[]TranscriptionEngineAConfigSpeechContextInner**](TranscriptionEngineAConfigSpeechContextInner.md) | Speech context to improve transcription accuracy. | [optional] 

## Methods

### NewTranscriptionEngineAConfig

`func NewTranscriptionEngineAConfig() *TranscriptionEngineAConfig`

NewTranscriptionEngineAConfig instantiates a new TranscriptionEngineAConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionEngineAConfigWithDefaults

`func NewTranscriptionEngineAConfigWithDefaults() *TranscriptionEngineAConfig`

NewTranscriptionEngineAConfigWithDefaults instantiates a new TranscriptionEngineAConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *TranscriptionEngineAConfig) GetLanguage() GoogleTranscriptionLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TranscriptionEngineAConfig) GetLanguageOk() (*GoogleTranscriptionLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TranscriptionEngineAConfig) SetLanguage(v GoogleTranscriptionLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TranscriptionEngineAConfig) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetInterimResults

`func (o *TranscriptionEngineAConfig) GetInterimResults() bool`

GetInterimResults returns the InterimResults field if non-nil, zero value otherwise.

### GetInterimResultsOk

`func (o *TranscriptionEngineAConfig) GetInterimResultsOk() (*bool, bool)`

GetInterimResultsOk returns a tuple with the InterimResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimResults

`func (o *TranscriptionEngineAConfig) SetInterimResults(v bool)`

SetInterimResults sets InterimResults field to given value.

### HasInterimResults

`func (o *TranscriptionEngineAConfig) HasInterimResults() bool`

HasInterimResults returns a boolean if a field has been set.

### GetEnableSpeakerDiarization

`func (o *TranscriptionEngineAConfig) GetEnableSpeakerDiarization() bool`

GetEnableSpeakerDiarization returns the EnableSpeakerDiarization field if non-nil, zero value otherwise.

### GetEnableSpeakerDiarizationOk

`func (o *TranscriptionEngineAConfig) GetEnableSpeakerDiarizationOk() (*bool, bool)`

GetEnableSpeakerDiarizationOk returns a tuple with the EnableSpeakerDiarization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpeakerDiarization

`func (o *TranscriptionEngineAConfig) SetEnableSpeakerDiarization(v bool)`

SetEnableSpeakerDiarization sets EnableSpeakerDiarization field to given value.

### HasEnableSpeakerDiarization

`func (o *TranscriptionEngineAConfig) HasEnableSpeakerDiarization() bool`

HasEnableSpeakerDiarization returns a boolean if a field has been set.

### GetMinSpeakerCount

`func (o *TranscriptionEngineAConfig) GetMinSpeakerCount() int32`

GetMinSpeakerCount returns the MinSpeakerCount field if non-nil, zero value otherwise.

### GetMinSpeakerCountOk

`func (o *TranscriptionEngineAConfig) GetMinSpeakerCountOk() (*int32, bool)`

GetMinSpeakerCountOk returns a tuple with the MinSpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSpeakerCount

`func (o *TranscriptionEngineAConfig) SetMinSpeakerCount(v int32)`

SetMinSpeakerCount sets MinSpeakerCount field to given value.

### HasMinSpeakerCount

`func (o *TranscriptionEngineAConfig) HasMinSpeakerCount() bool`

HasMinSpeakerCount returns a boolean if a field has been set.

### GetMaxSpeakerCount

`func (o *TranscriptionEngineAConfig) GetMaxSpeakerCount() int32`

GetMaxSpeakerCount returns the MaxSpeakerCount field if non-nil, zero value otherwise.

### GetMaxSpeakerCountOk

`func (o *TranscriptionEngineAConfig) GetMaxSpeakerCountOk() (*int32, bool)`

GetMaxSpeakerCountOk returns a tuple with the MaxSpeakerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeakerCount

`func (o *TranscriptionEngineAConfig) SetMaxSpeakerCount(v int32)`

SetMaxSpeakerCount sets MaxSpeakerCount field to given value.

### HasMaxSpeakerCount

`func (o *TranscriptionEngineAConfig) HasMaxSpeakerCount() bool`

HasMaxSpeakerCount returns a boolean if a field has been set.

### GetProfanityFilter

`func (o *TranscriptionEngineAConfig) GetProfanityFilter() bool`

GetProfanityFilter returns the ProfanityFilter field if non-nil, zero value otherwise.

### GetProfanityFilterOk

`func (o *TranscriptionEngineAConfig) GetProfanityFilterOk() (*bool, bool)`

GetProfanityFilterOk returns a tuple with the ProfanityFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfanityFilter

`func (o *TranscriptionEngineAConfig) SetProfanityFilter(v bool)`

SetProfanityFilter sets ProfanityFilter field to given value.

### HasProfanityFilter

`func (o *TranscriptionEngineAConfig) HasProfanityFilter() bool`

HasProfanityFilter returns a boolean if a field has been set.

### GetUseEnhanced

`func (o *TranscriptionEngineAConfig) GetUseEnhanced() bool`

GetUseEnhanced returns the UseEnhanced field if non-nil, zero value otherwise.

### GetUseEnhancedOk

`func (o *TranscriptionEngineAConfig) GetUseEnhancedOk() (*bool, bool)`

GetUseEnhancedOk returns a tuple with the UseEnhanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnhanced

`func (o *TranscriptionEngineAConfig) SetUseEnhanced(v bool)`

SetUseEnhanced sets UseEnhanced field to given value.

### HasUseEnhanced

`func (o *TranscriptionEngineAConfig) HasUseEnhanced() bool`

HasUseEnhanced returns a boolean if a field has been set.

### GetModel

`func (o *TranscriptionEngineAConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TranscriptionEngineAConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TranscriptionEngineAConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TranscriptionEngineAConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetHints

`func (o *TranscriptionEngineAConfig) GetHints() []string`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *TranscriptionEngineAConfig) GetHintsOk() (*[]string, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *TranscriptionEngineAConfig) SetHints(v []string)`

SetHints sets Hints field to given value.

### HasHints

`func (o *TranscriptionEngineAConfig) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetSpeechContext

`func (o *TranscriptionEngineAConfig) GetSpeechContext() []TranscriptionEngineAConfigSpeechContextInner`

GetSpeechContext returns the SpeechContext field if non-nil, zero value otherwise.

### GetSpeechContextOk

`func (o *TranscriptionEngineAConfig) GetSpeechContextOk() (*[]TranscriptionEngineAConfigSpeechContextInner, bool)`

GetSpeechContextOk returns a tuple with the SpeechContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechContext

`func (o *TranscriptionEngineAConfig) SetSpeechContext(v []TranscriptionEngineAConfigSpeechContextInner)`

SetSpeechContext sets SpeechContext field to given value.

### HasSpeechContext

`func (o *TranscriptionEngineAConfig) HasSpeechContext() bool`

HasSpeechContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


