# AIAssistantStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assistant** | Pointer to [**AIAssistantStartRequestAssistant**](AIAssistantStartRequestAssistant.md) |  | [optional] 
**Voice** | Pointer to **string** | The voice to be used by the voice assistant. Currently we support ElevenLabs, Telnyx and AWS voices.   **Supported Providers:** - **AWS:** Use &#x60;AWS.Polly.&lt;VoiceId&gt;&#x60; (e.g., &#x60;AWS.Polly.Joanna&#x60;). For neural voices, which provide more realistic, human-like speech, append &#x60;-Neural&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Joanna-Neural&#x60;). For long-form engine, append &#x60;-LongForm&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-LongForm&#x60;). For generative engine, the latest provided by AWS Polly, append &#x60;-Generative&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-Generative&#x60;). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility. - **Azure:** Use &#x60;Azure.&lt;VoiceId&gt;. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).) - **ElevenLabs:** Use &#x60;ElevenLabs.&lt;ModelId&gt;.&lt;VoiceId&gt;&#x60; (e.g., &#x60;ElevenLabs.BaseModel.John&#x60;). The &#x60;ModelId&#x60; part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration secret under &#x60;\&quot;voice_settings\&quot;: {\&quot;api_key_ref\&quot;: \&quot;&lt;secret_id&gt;\&quot;}&#x60;. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).  - **Telnyx:** Use &#x60;Telnyx.&lt;model_id&gt;.&lt;voice_id&gt;&#x60; | [optional] [default to "Telnyx.KokoroTTS.af"]
**VoiceSettings** | Pointer to [**AIAssistantStartRequestVoiceSettings**](AIAssistantStartRequestVoiceSettings.md) |  | [optional] 
**Greeting** | Pointer to **string** | Text that will be played when the assistant starts, if none then nothing will be played when the assistant starts. The greeting can be text for any voice or SSML for &#x60;AWS.Polly.&lt;voice_id&gt;&#x60; voices. There is a 3,000 character limit. | [optional] 
**InterruptionSettings** | Pointer to [**InterruptionSettings**](InterruptionSettings.md) |  | [optional] 
**Transcription** | Pointer to [**TranscriptionConfig**](TranscriptionConfig.md) |  | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewAIAssistantStartRequest

`func NewAIAssistantStartRequest() *AIAssistantStartRequest`

NewAIAssistantStartRequest instantiates a new AIAssistantStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAssistantStartRequestWithDefaults

`func NewAIAssistantStartRequestWithDefaults() *AIAssistantStartRequest`

NewAIAssistantStartRequestWithDefaults instantiates a new AIAssistantStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistant

`func (o *AIAssistantStartRequest) GetAssistant() AIAssistantStartRequestAssistant`

GetAssistant returns the Assistant field if non-nil, zero value otherwise.

### GetAssistantOk

`func (o *AIAssistantStartRequest) GetAssistantOk() (*AIAssistantStartRequestAssistant, bool)`

GetAssistantOk returns a tuple with the Assistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistant

`func (o *AIAssistantStartRequest) SetAssistant(v AIAssistantStartRequestAssistant)`

SetAssistant sets Assistant field to given value.

### HasAssistant

`func (o *AIAssistantStartRequest) HasAssistant() bool`

HasAssistant returns a boolean if a field has been set.

### GetVoice

`func (o *AIAssistantStartRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *AIAssistantStartRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *AIAssistantStartRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *AIAssistantStartRequest) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetVoiceSettings

`func (o *AIAssistantStartRequest) GetVoiceSettings() AIAssistantStartRequestVoiceSettings`

GetVoiceSettings returns the VoiceSettings field if non-nil, zero value otherwise.

### GetVoiceSettingsOk

`func (o *AIAssistantStartRequest) GetVoiceSettingsOk() (*AIAssistantStartRequestVoiceSettings, bool)`

GetVoiceSettingsOk returns a tuple with the VoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSettings

`func (o *AIAssistantStartRequest) SetVoiceSettings(v AIAssistantStartRequestVoiceSettings)`

SetVoiceSettings sets VoiceSettings field to given value.

### HasVoiceSettings

`func (o *AIAssistantStartRequest) HasVoiceSettings() bool`

HasVoiceSettings returns a boolean if a field has been set.

### GetGreeting

`func (o *AIAssistantStartRequest) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *AIAssistantStartRequest) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *AIAssistantStartRequest) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.

### HasGreeting

`func (o *AIAssistantStartRequest) HasGreeting() bool`

HasGreeting returns a boolean if a field has been set.

### GetInterruptionSettings

`func (o *AIAssistantStartRequest) GetInterruptionSettings() InterruptionSettings`

GetInterruptionSettings returns the InterruptionSettings field if non-nil, zero value otherwise.

### GetInterruptionSettingsOk

`func (o *AIAssistantStartRequest) GetInterruptionSettingsOk() (*InterruptionSettings, bool)`

GetInterruptionSettingsOk returns a tuple with the InterruptionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptionSettings

`func (o *AIAssistantStartRequest) SetInterruptionSettings(v InterruptionSettings)`

SetInterruptionSettings sets InterruptionSettings field to given value.

### HasInterruptionSettings

`func (o *AIAssistantStartRequest) HasInterruptionSettings() bool`

HasInterruptionSettings returns a boolean if a field has been set.

### GetTranscription

`func (o *AIAssistantStartRequest) GetTranscription() TranscriptionConfig`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *AIAssistantStartRequest) GetTranscriptionOk() (*TranscriptionConfig, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *AIAssistantStartRequest) SetTranscription(v TranscriptionConfig)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *AIAssistantStartRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetClientState

`func (o *AIAssistantStartRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *AIAssistantStartRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *AIAssistantStartRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *AIAssistantStartRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *AIAssistantStartRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *AIAssistantStartRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *AIAssistantStartRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *AIAssistantStartRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


