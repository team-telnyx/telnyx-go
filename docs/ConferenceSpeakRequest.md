# ConferenceSpeakRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | Pointer to **[]string** | Call Control IDs of participants who will hear the spoken text. When empty all participants will hear the spoken text. | [optional] 
**Payload** | **string** | The text or SSML to be converted into speech. There is a 3,000 character limit. | 
**PayloadType** | Pointer to **string** | The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML). | [optional] [default to "text"]
**Voice** | **string** | Specifies the voice used in speech synthesis.  - Define voices using the format &#x60;&lt;Provider&gt;.&lt;Model&gt;.&lt;VoiceId&gt;&#x60;. Specifying only the provider will give default values for voice_id and model_id.   **Supported Providers:** - **AWS:** Use &#x60;AWS.Polly.&lt;VoiceId&gt;&#x60; (e.g., &#x60;AWS.Polly.Joanna&#x60;). For neural voices, which provide more realistic, human-like speech, append &#x60;-Neural&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Joanna-Neural&#x60;). For long-form engine, append &#x60;-LongForm&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-LongForm&#x60;). For generative engine, the latest provided by AWS Polly, append &#x60;-Generative&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-Generative&#x60;). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility. - **Azure:** Use &#x60;Azure.&lt;VoiceId&gt;. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).) - **ElevenLabs:** Use &#x60;ElevenLabs.&lt;ModelId&gt;.&lt;VoiceId&gt;&#x60; (e.g., &#x60;ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM&#x60;). The &#x60;ModelId&#x60; part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration identifier secret in &#x60;\&quot;voice_settings\&quot;: {\&quot;api_key_ref\&quot;: \&quot;&lt;secret_identifier&gt;\&quot;}&#x60;. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).  - **Telnyx:** Use &#x60;Telnyx.&lt;model_id&gt;.&lt;voice_id&gt;&#x60;  For service_level basic, you may define the gender of the speaker (male or female). | 
**VoiceSettings** | Pointer to [**ElevenLabsVoiceSettings**](ElevenLabsVoiceSettings.md) |  | [optional] 
**Language** | Pointer to **string** | The language you want spoken. This parameter is ignored when a &#x60;Polly.*&#x60; voice is specified. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] 

## Methods

### NewConferenceSpeakRequest

`func NewConferenceSpeakRequest(payload string, voice string, ) *ConferenceSpeakRequest`

NewConferenceSpeakRequest instantiates a new ConferenceSpeakRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceSpeakRequestWithDefaults

`func NewConferenceSpeakRequestWithDefaults() *ConferenceSpeakRequest`

NewConferenceSpeakRequestWithDefaults instantiates a new ConferenceSpeakRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlIds

`func (o *ConferenceSpeakRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferenceSpeakRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferenceSpeakRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.

### HasCallControlIds

`func (o *ConferenceSpeakRequest) HasCallControlIds() bool`

HasCallControlIds returns a boolean if a field has been set.

### GetPayload

`func (o *ConferenceSpeakRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ConferenceSpeakRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ConferenceSpeakRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetPayloadType

`func (o *ConferenceSpeakRequest) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *ConferenceSpeakRequest) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *ConferenceSpeakRequest) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *ConferenceSpeakRequest) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### GetVoice

`func (o *ConferenceSpeakRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *ConferenceSpeakRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *ConferenceSpeakRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.


### GetVoiceSettings

`func (o *ConferenceSpeakRequest) GetVoiceSettings() ElevenLabsVoiceSettings`

GetVoiceSettings returns the VoiceSettings field if non-nil, zero value otherwise.

### GetVoiceSettingsOk

`func (o *ConferenceSpeakRequest) GetVoiceSettingsOk() (*ElevenLabsVoiceSettings, bool)`

GetVoiceSettingsOk returns a tuple with the VoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSettings

`func (o *ConferenceSpeakRequest) SetVoiceSettings(v ElevenLabsVoiceSettings)`

SetVoiceSettings sets VoiceSettings field to given value.

### HasVoiceSettings

`func (o *ConferenceSpeakRequest) HasVoiceSettings() bool`

HasVoiceSettings returns a boolean if a field has been set.

### GetLanguage

`func (o *ConferenceSpeakRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ConferenceSpeakRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ConferenceSpeakRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ConferenceSpeakRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCommandId

`func (o *ConferenceSpeakRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *ConferenceSpeakRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *ConferenceSpeakRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *ConferenceSpeakRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


