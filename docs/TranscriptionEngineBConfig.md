# TranscriptionEngineBConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**TelnyxTranscriptionLanguage**](TelnyxTranscriptionLanguage.md) |  | [optional] [default to EN]
**TranscriptionModel** | Pointer to **string** | The model to use for transcription. | [optional] [default to "openai/whisper-tiny"]

## Methods

### NewTranscriptionEngineBConfig

`func NewTranscriptionEngineBConfig() *TranscriptionEngineBConfig`

NewTranscriptionEngineBConfig instantiates a new TranscriptionEngineBConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionEngineBConfigWithDefaults

`func NewTranscriptionEngineBConfigWithDefaults() *TranscriptionEngineBConfig`

NewTranscriptionEngineBConfigWithDefaults instantiates a new TranscriptionEngineBConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *TranscriptionEngineBConfig) GetLanguage() TelnyxTranscriptionLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TranscriptionEngineBConfig) GetLanguageOk() (*TelnyxTranscriptionLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TranscriptionEngineBConfig) SetLanguage(v TelnyxTranscriptionLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *TranscriptionEngineBConfig) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTranscriptionModel

`func (o *TranscriptionEngineBConfig) GetTranscriptionModel() string`

GetTranscriptionModel returns the TranscriptionModel field if non-nil, zero value otherwise.

### GetTranscriptionModelOk

`func (o *TranscriptionEngineBConfig) GetTranscriptionModelOk() (*string, bool)`

GetTranscriptionModelOk returns a tuple with the TranscriptionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionModel

`func (o *TranscriptionEngineBConfig) SetTranscriptionModel(v string)`

SetTranscriptionModel sets TranscriptionModel field to given value.

### HasTranscriptionModel

`func (o *TranscriptionEngineBConfig) HasTranscriptionModel() bool`

HasTranscriptionModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


