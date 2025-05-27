# TranscriptionEngineAConfigSpeechContextInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phrases** | Pointer to **[]string** |  | [optional] [default to []]
**Boost** | Pointer to **float32** | Boost factor for the speech context. | [optional] [default to 1.0]

## Methods

### NewTranscriptionEngineAConfigSpeechContextInner

`func NewTranscriptionEngineAConfigSpeechContextInner() *TranscriptionEngineAConfigSpeechContextInner`

NewTranscriptionEngineAConfigSpeechContextInner instantiates a new TranscriptionEngineAConfigSpeechContextInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionEngineAConfigSpeechContextInnerWithDefaults

`func NewTranscriptionEngineAConfigSpeechContextInnerWithDefaults() *TranscriptionEngineAConfigSpeechContextInner`

NewTranscriptionEngineAConfigSpeechContextInnerWithDefaults instantiates a new TranscriptionEngineAConfigSpeechContextInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhrases

`func (o *TranscriptionEngineAConfigSpeechContextInner) GetPhrases() []string`

GetPhrases returns the Phrases field if non-nil, zero value otherwise.

### GetPhrasesOk

`func (o *TranscriptionEngineAConfigSpeechContextInner) GetPhrasesOk() (*[]string, bool)`

GetPhrasesOk returns a tuple with the Phrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhrases

`func (o *TranscriptionEngineAConfigSpeechContextInner) SetPhrases(v []string)`

SetPhrases sets Phrases field to given value.

### HasPhrases

`func (o *TranscriptionEngineAConfigSpeechContextInner) HasPhrases() bool`

HasPhrases returns a boolean if a field has been set.

### GetBoost

`func (o *TranscriptionEngineAConfigSpeechContextInner) GetBoost() float32`

GetBoost returns the Boost field if non-nil, zero value otherwise.

### GetBoostOk

`func (o *TranscriptionEngineAConfigSpeechContextInner) GetBoostOk() (*float32, bool)`

GetBoostOk returns a tuple with the Boost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoost

`func (o *TranscriptionEngineAConfigSpeechContextInner) SetBoost(v float32)`

SetBoost sets Boost field to given value.

### HasBoost

`func (o *TranscriptionEngineAConfigSpeechContextInner) HasBoost() bool`

HasBoost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


