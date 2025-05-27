# GenerateTextToSpeechRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voice** | **string** | The voice ID in the format Provider.ModelId.VoiceId.  Examples: - AWS.Polly.Joanna-Neural - Azure.en-US-AvaMultilingualNeural - ElevenLabs.eleven_multilingual_v2.Rachel - Telnyx.KokoroTTS.af  Use the &#x60;GET /text-to-speech/voices&#x60; endpoint to get a complete list of available voices. | 
**Text** | **string** | The text to convert to speech | 

## Methods

### NewGenerateTextToSpeechRequest

`func NewGenerateTextToSpeechRequest(voice string, text string, ) *GenerateTextToSpeechRequest`

NewGenerateTextToSpeechRequest instantiates a new GenerateTextToSpeechRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateTextToSpeechRequestWithDefaults

`func NewGenerateTextToSpeechRequestWithDefaults() *GenerateTextToSpeechRequest`

NewGenerateTextToSpeechRequestWithDefaults instantiates a new GenerateTextToSpeechRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoice

`func (o *GenerateTextToSpeechRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *GenerateTextToSpeechRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *GenerateTextToSpeechRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.


### GetText

`func (o *GenerateTextToSpeechRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GenerateTextToSpeechRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GenerateTextToSpeechRequest) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


