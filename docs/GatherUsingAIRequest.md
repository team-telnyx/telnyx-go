# GatherUsingAIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | **map[string]interface{}** | The parameters described as a JSON Schema object that needs to be gathered by the voice assistant. See the [JSON Schema reference](https://json-schema.org/understanding-json-schema) for documentation about the format | 
**Assistant** | Pointer to [**Assistant**](Assistant.md) |  | [optional] 
**Transcription** | Pointer to [**TranscriptionConfig**](TranscriptionConfig.md) |  | [optional] 
**Language** | Pointer to [**GoogleTranscriptionLanguage**](GoogleTranscriptionLanguage.md) |  | [optional] [default to EN]
**Voice** | Pointer to **string** | The voice to be used by the voice assistant. Currently we support ElevenLabs, Telnyx and AWS voices.   **Supported Providers:** - **AWS:** Use &#x60;AWS.Polly.&lt;VoiceId&gt;&#x60; (e.g., &#x60;AWS.Polly.Joanna&#x60;). For neural voices, which provide more realistic, human-like speech, append &#x60;-Neural&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Joanna-Neural&#x60;). For long-form engine, append &#x60;-LongForm&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-LongForm&#x60;). For generative engine, the latest provided by AWS Polly, append &#x60;-Generative&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-Generative&#x60;). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility. - **Azure:** Use &#x60;Azure.&lt;VoiceId&gt;. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).) - **ElevenLabs:** Use &#x60;ElevenLabs.&lt;ModelId&gt;.&lt;VoiceId&gt;&#x60; (e.g., &#x60;ElevenLabs.BaseModel.John&#x60;). The &#x60;ModelId&#x60; part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration secret under &#x60;\&quot;voice_settings\&quot;: {\&quot;api_key_ref\&quot;: \&quot;&lt;secret_id&gt;\&quot;}&#x60;. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).  - **Telnyx:** Use &#x60;Telnyx.&lt;model_id&gt;.&lt;voice_id&gt;&#x60; | [optional] [default to "Telnyx.KokoroTTS.af"]
**VoiceSettings** | Pointer to [**AIAssistantStartRequestVoiceSettings**](AIAssistantStartRequestVoiceSettings.md) |  | [optional] 
**Greeting** | Pointer to **string** | Text that will be played when the gathering starts, if none then nothing will be played when the gathering starts. The greeting can be text for any voice or SSML for &#x60;AWS.Polly.&lt;voice_id&gt;&#x60; voices. There is a 3,000 character limit. | [optional] 
**SendPartialResults** | Pointer to **bool** | Default is &#x60;false&#x60;. If set to &#x60;true&#x60;, the voice assistant will send partial results via the &#x60;call.ai_gather.partial_results&#x60; [callback](https://developers.telnyx.com/api/call-control/call-gather-using-ai#callbacks) in real time as individual fields are gathered. If set to &#x60;false&#x60;, the voice assistant will only send the final result via the &#x60;call.ai_gather.ended&#x60; callback. | [optional] 
**SendMessageHistoryUpdates** | Pointer to **bool** | Default is &#x60;false&#x60;. If set to &#x60;true&#x60;, the voice assistant will send updates to the message history via the &#x60;call.ai_gather.message_history_updated&#x60; [callback](https://developers.telnyx.com/api/call-control/call-gather-using-ai#callbacks) in real time as the message history is updated. | [optional] 
**MessageHistory** | Pointer to [**[]CallAIGatherEndedPayloadMessageHistoryInner**](CallAIGatherEndedPayloadMessageHistoryInner.md) | The message history you want the voice assistant to be aware of, this can be useful to keep the context of the conversation, or to pass additional information to the voice assistant. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**InterruptionSettings** | Pointer to [**InterruptionSettings**](InterruptionSettings.md) |  | [optional] 
**UserResponseTimeoutMs** | Pointer to **int32** | The number of milliseconds to wait for a user response before the voice assistant times out and check if the user is still there. | [optional] [default to 10000]

## Methods

### NewGatherUsingAIRequest

`func NewGatherUsingAIRequest(parameters map[string]interface{}, ) *GatherUsingAIRequest`

NewGatherUsingAIRequest instantiates a new GatherUsingAIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatherUsingAIRequestWithDefaults

`func NewGatherUsingAIRequestWithDefaults() *GatherUsingAIRequest`

NewGatherUsingAIRequestWithDefaults instantiates a new GatherUsingAIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *GatherUsingAIRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GatherUsingAIRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GatherUsingAIRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.


### GetAssistant

`func (o *GatherUsingAIRequest) GetAssistant() Assistant`

GetAssistant returns the Assistant field if non-nil, zero value otherwise.

### GetAssistantOk

`func (o *GatherUsingAIRequest) GetAssistantOk() (*Assistant, bool)`

GetAssistantOk returns a tuple with the Assistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistant

`func (o *GatherUsingAIRequest) SetAssistant(v Assistant)`

SetAssistant sets Assistant field to given value.

### HasAssistant

`func (o *GatherUsingAIRequest) HasAssistant() bool`

HasAssistant returns a boolean if a field has been set.

### GetTranscription

`func (o *GatherUsingAIRequest) GetTranscription() TranscriptionConfig`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *GatherUsingAIRequest) GetTranscriptionOk() (*TranscriptionConfig, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *GatherUsingAIRequest) SetTranscription(v TranscriptionConfig)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *GatherUsingAIRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetLanguage

`func (o *GatherUsingAIRequest) GetLanguage() GoogleTranscriptionLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GatherUsingAIRequest) GetLanguageOk() (*GoogleTranscriptionLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GatherUsingAIRequest) SetLanguage(v GoogleTranscriptionLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GatherUsingAIRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVoice

`func (o *GatherUsingAIRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *GatherUsingAIRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *GatherUsingAIRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *GatherUsingAIRequest) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetVoiceSettings

`func (o *GatherUsingAIRequest) GetVoiceSettings() AIAssistantStartRequestVoiceSettings`

GetVoiceSettings returns the VoiceSettings field if non-nil, zero value otherwise.

### GetVoiceSettingsOk

`func (o *GatherUsingAIRequest) GetVoiceSettingsOk() (*AIAssistantStartRequestVoiceSettings, bool)`

GetVoiceSettingsOk returns a tuple with the VoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSettings

`func (o *GatherUsingAIRequest) SetVoiceSettings(v AIAssistantStartRequestVoiceSettings)`

SetVoiceSettings sets VoiceSettings field to given value.

### HasVoiceSettings

`func (o *GatherUsingAIRequest) HasVoiceSettings() bool`

HasVoiceSettings returns a boolean if a field has been set.

### GetGreeting

`func (o *GatherUsingAIRequest) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *GatherUsingAIRequest) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *GatherUsingAIRequest) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.

### HasGreeting

`func (o *GatherUsingAIRequest) HasGreeting() bool`

HasGreeting returns a boolean if a field has been set.

### GetSendPartialResults

`func (o *GatherUsingAIRequest) GetSendPartialResults() bool`

GetSendPartialResults returns the SendPartialResults field if non-nil, zero value otherwise.

### GetSendPartialResultsOk

`func (o *GatherUsingAIRequest) GetSendPartialResultsOk() (*bool, bool)`

GetSendPartialResultsOk returns a tuple with the SendPartialResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPartialResults

`func (o *GatherUsingAIRequest) SetSendPartialResults(v bool)`

SetSendPartialResults sets SendPartialResults field to given value.

### HasSendPartialResults

`func (o *GatherUsingAIRequest) HasSendPartialResults() bool`

HasSendPartialResults returns a boolean if a field has been set.

### GetSendMessageHistoryUpdates

`func (o *GatherUsingAIRequest) GetSendMessageHistoryUpdates() bool`

GetSendMessageHistoryUpdates returns the SendMessageHistoryUpdates field if non-nil, zero value otherwise.

### GetSendMessageHistoryUpdatesOk

`func (o *GatherUsingAIRequest) GetSendMessageHistoryUpdatesOk() (*bool, bool)`

GetSendMessageHistoryUpdatesOk returns a tuple with the SendMessageHistoryUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMessageHistoryUpdates

`func (o *GatherUsingAIRequest) SetSendMessageHistoryUpdates(v bool)`

SetSendMessageHistoryUpdates sets SendMessageHistoryUpdates field to given value.

### HasSendMessageHistoryUpdates

`func (o *GatherUsingAIRequest) HasSendMessageHistoryUpdates() bool`

HasSendMessageHistoryUpdates returns a boolean if a field has been set.

### GetMessageHistory

`func (o *GatherUsingAIRequest) GetMessageHistory() []CallAIGatherEndedPayloadMessageHistoryInner`

GetMessageHistory returns the MessageHistory field if non-nil, zero value otherwise.

### GetMessageHistoryOk

`func (o *GatherUsingAIRequest) GetMessageHistoryOk() (*[]CallAIGatherEndedPayloadMessageHistoryInner, bool)`

GetMessageHistoryOk returns a tuple with the MessageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHistory

`func (o *GatherUsingAIRequest) SetMessageHistory(v []CallAIGatherEndedPayloadMessageHistoryInner)`

SetMessageHistory sets MessageHistory field to given value.

### HasMessageHistory

`func (o *GatherUsingAIRequest) HasMessageHistory() bool`

HasMessageHistory returns a boolean if a field has been set.

### GetClientState

`func (o *GatherUsingAIRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *GatherUsingAIRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *GatherUsingAIRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *GatherUsingAIRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *GatherUsingAIRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *GatherUsingAIRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *GatherUsingAIRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *GatherUsingAIRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetInterruptionSettings

`func (o *GatherUsingAIRequest) GetInterruptionSettings() InterruptionSettings`

GetInterruptionSettings returns the InterruptionSettings field if non-nil, zero value otherwise.

### GetInterruptionSettingsOk

`func (o *GatherUsingAIRequest) GetInterruptionSettingsOk() (*InterruptionSettings, bool)`

GetInterruptionSettingsOk returns a tuple with the InterruptionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptionSettings

`func (o *GatherUsingAIRequest) SetInterruptionSettings(v InterruptionSettings)`

SetInterruptionSettings sets InterruptionSettings field to given value.

### HasInterruptionSettings

`func (o *GatherUsingAIRequest) HasInterruptionSettings() bool`

HasInterruptionSettings returns a boolean if a field has been set.

### GetUserResponseTimeoutMs

`func (o *GatherUsingAIRequest) GetUserResponseTimeoutMs() int32`

GetUserResponseTimeoutMs returns the UserResponseTimeoutMs field if non-nil, zero value otherwise.

### GetUserResponseTimeoutMsOk

`func (o *GatherUsingAIRequest) GetUserResponseTimeoutMsOk() (*int32, bool)`

GetUserResponseTimeoutMsOk returns a tuple with the UserResponseTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserResponseTimeoutMs

`func (o *GatherUsingAIRequest) SetUserResponseTimeoutMs(v int32)`

SetUserResponseTimeoutMs sets UserResponseTimeoutMs field to given value.

### HasUserResponseTimeoutMs

`func (o *GatherUsingAIRequest) HasUserResponseTimeoutMs() bool`

HasUserResponseTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


