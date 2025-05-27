# TranscriptionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** | The speech to text model to be used by the voice assistant.  - &#x60;distil-whisper/distil-large-v2&#x60; is lower latency but English-only. - &#x60;openai/whisper-large-v3-turbo&#x60; is multi-lingual with automatic language detection but slightly higher latency. - &#x60;google&#x60; is a multi-lingual option, please describe the language in the &#x60;language&#x60; field. | [optional] [default to "distil-whisper/distil-large-v2"]

## Methods

### NewTranscriptionConfig

`func NewTranscriptionConfig() *TranscriptionConfig`

NewTranscriptionConfig instantiates a new TranscriptionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionConfigWithDefaults

`func NewTranscriptionConfigWithDefaults() *TranscriptionConfig`

NewTranscriptionConfigWithDefaults instantiates a new TranscriptionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TranscriptionConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TranscriptionConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TranscriptionConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TranscriptionConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


