# GatherUsingSpeakRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **string** | The text or SSML to be converted into speech. There is a 3,000 character limit. | 
**InvalidPayload** | Pointer to **string** | The text or SSML to be converted into speech when digits don&#39;t match the &#x60;valid_digits&#x60; parameter or the number of digits is not between &#x60;min&#x60; and &#x60;max&#x60;. There is a 3,000 character limit. | [optional] 
**PayloadType** | Pointer to **string** | The type of the provided payload. The payload can either be plain text, or Speech Synthesis Markup Language (SSML). | [optional] [default to "text"]
**ServiceLevel** | Pointer to **string** | This parameter impacts speech quality, language options and payload types. When using &#x60;basic&#x60;, only the &#x60;en-US&#x60; language and payload type &#x60;text&#x60; are allowed. | [optional] [default to "premium"]
**Voice** | **string** | Specifies the voice used in speech synthesis.  - Define voices using the format &#x60;&lt;Provider&gt;.&lt;Model&gt;.&lt;VoiceId&gt;&#x60;. Specifying only the provider will give default values for voice_id and model_id.   **Supported Providers:** - **AWS:** Use &#x60;AWS.Polly.&lt;VoiceId&gt;&#x60; (e.g., &#x60;AWS.Polly.Joanna&#x60;). For neural voices, which provide more realistic, human-like speech, append &#x60;-Neural&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Joanna-Neural&#x60;). For long-form engine, append &#x60;-LongForm&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-LongForm&#x60;). For generative engine, the latest provided by AWS Polly, append &#x60;-Generative&#x60; to the &#x60;VoiceId&#x60; (e.g., &#x60;AWS.Polly.Danielle-Generative&#x60;). Check the [available voices](https://docs.aws.amazon.com/polly/latest/dg/available-voices.html) for compatibility. - **Azure:** Use &#x60;Azure.&lt;VoiceId&gt;. (e.g. Azure.en-CA-ClaraNeural, Azure.en-CA-LiamNeural, Azure.en-US-BrianMultilingualNeural, Azure.en-US-Ava:DragonHDLatestNeural. For a complete list of voices, go to [Azure Voice Gallery](https://speech.microsoft.com/portal/voicegallery).) - **ElevenLabs:** Use &#x60;ElevenLabs.&lt;ModelId&gt;.&lt;VoiceId&gt;&#x60; (e.g., &#x60;ElevenLabs.eleven_multilingual_v2.21m00Tcm4TlvDq8ikWAM&#x60;). The &#x60;ModelId&#x60; part is optional. To use ElevenLabs, you must provide your ElevenLabs API key as an integration identifier secret in &#x60;\&quot;voice_settings\&quot;: {\&quot;api_key_ref\&quot;: \&quot;&lt;secret_identifier&gt;\&quot;}&#x60;. See [integration secrets documentation](https://developers.telnyx.com/api/secrets-manager/integration-secrets/create-integration-secret) for details. Check [available voices](https://elevenlabs.io/docs/api-reference/get-voices).  - **Telnyx:** Use &#x60;Telnyx.&lt;model_id&gt;.&lt;voice_id&gt;&#x60;  For service_level basic, you may define the gender of the speaker (male or female). | 
**VoiceSettings** | Pointer to [**ElevenLabsVoiceSettings**](ElevenLabsVoiceSettings.md) |  | [optional] 
**Language** | Pointer to **string** | The language you want spoken. This parameter is ignored when a &#x60;Polly.*&#x60; voice is specified. | [optional] 
**MinimumDigits** | Pointer to **int32** | The minimum number of digits to fetch. This parameter has a minimum value of 1. | [optional] [default to 1]
**MaximumDigits** | Pointer to **int32** | The maximum number of digits to fetch. This parameter has a maximum value of 128. | [optional] [default to 128]
**MaximumTries** | Pointer to **int32** | The maximum number of times that a file should be played back if there is no input from the user on the call. | [optional] [default to 3]
**TimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait for a DTMF response after speak ends before a replaying the sound file. | [optional] [default to 60000]
**TerminatingDigit** | Pointer to **string** | The digit used to terminate input if fewer than &#x60;maximum_digits&#x60; digits have been gathered. | [optional] [default to "#"]
**ValidDigits** | Pointer to **string** | A list of all digits accepted as valid. | [optional] [default to "0123456789#*"]
**InterDigitTimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait for input between digits. | [optional] [default to 5000]
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewGatherUsingSpeakRequest

`func NewGatherUsingSpeakRequest(payload string, voice string, ) *GatherUsingSpeakRequest`

NewGatherUsingSpeakRequest instantiates a new GatherUsingSpeakRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatherUsingSpeakRequestWithDefaults

`func NewGatherUsingSpeakRequestWithDefaults() *GatherUsingSpeakRequest`

NewGatherUsingSpeakRequestWithDefaults instantiates a new GatherUsingSpeakRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *GatherUsingSpeakRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GatherUsingSpeakRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GatherUsingSpeakRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetInvalidPayload

`func (o *GatherUsingSpeakRequest) GetInvalidPayload() string`

GetInvalidPayload returns the InvalidPayload field if non-nil, zero value otherwise.

### GetInvalidPayloadOk

`func (o *GatherUsingSpeakRequest) GetInvalidPayloadOk() (*string, bool)`

GetInvalidPayloadOk returns a tuple with the InvalidPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPayload

`func (o *GatherUsingSpeakRequest) SetInvalidPayload(v string)`

SetInvalidPayload sets InvalidPayload field to given value.

### HasInvalidPayload

`func (o *GatherUsingSpeakRequest) HasInvalidPayload() bool`

HasInvalidPayload returns a boolean if a field has been set.

### GetPayloadType

`func (o *GatherUsingSpeakRequest) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *GatherUsingSpeakRequest) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *GatherUsingSpeakRequest) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *GatherUsingSpeakRequest) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### GetServiceLevel

`func (o *GatherUsingSpeakRequest) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *GatherUsingSpeakRequest) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *GatherUsingSpeakRequest) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *GatherUsingSpeakRequest) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetVoice

`func (o *GatherUsingSpeakRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *GatherUsingSpeakRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *GatherUsingSpeakRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.


### GetVoiceSettings

`func (o *GatherUsingSpeakRequest) GetVoiceSettings() ElevenLabsVoiceSettings`

GetVoiceSettings returns the VoiceSettings field if non-nil, zero value otherwise.

### GetVoiceSettingsOk

`func (o *GatherUsingSpeakRequest) GetVoiceSettingsOk() (*ElevenLabsVoiceSettings, bool)`

GetVoiceSettingsOk returns a tuple with the VoiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceSettings

`func (o *GatherUsingSpeakRequest) SetVoiceSettings(v ElevenLabsVoiceSettings)`

SetVoiceSettings sets VoiceSettings field to given value.

### HasVoiceSettings

`func (o *GatherUsingSpeakRequest) HasVoiceSettings() bool`

HasVoiceSettings returns a boolean if a field has been set.

### GetLanguage

`func (o *GatherUsingSpeakRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GatherUsingSpeakRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GatherUsingSpeakRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GatherUsingSpeakRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMinimumDigits

`func (o *GatherUsingSpeakRequest) GetMinimumDigits() int32`

GetMinimumDigits returns the MinimumDigits field if non-nil, zero value otherwise.

### GetMinimumDigitsOk

`func (o *GatherUsingSpeakRequest) GetMinimumDigitsOk() (*int32, bool)`

GetMinimumDigitsOk returns a tuple with the MinimumDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDigits

`func (o *GatherUsingSpeakRequest) SetMinimumDigits(v int32)`

SetMinimumDigits sets MinimumDigits field to given value.

### HasMinimumDigits

`func (o *GatherUsingSpeakRequest) HasMinimumDigits() bool`

HasMinimumDigits returns a boolean if a field has been set.

### GetMaximumDigits

`func (o *GatherUsingSpeakRequest) GetMaximumDigits() int32`

GetMaximumDigits returns the MaximumDigits field if non-nil, zero value otherwise.

### GetMaximumDigitsOk

`func (o *GatherUsingSpeakRequest) GetMaximumDigitsOk() (*int32, bool)`

GetMaximumDigitsOk returns a tuple with the MaximumDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDigits

`func (o *GatherUsingSpeakRequest) SetMaximumDigits(v int32)`

SetMaximumDigits sets MaximumDigits field to given value.

### HasMaximumDigits

`func (o *GatherUsingSpeakRequest) HasMaximumDigits() bool`

HasMaximumDigits returns a boolean if a field has been set.

### GetMaximumTries

`func (o *GatherUsingSpeakRequest) GetMaximumTries() int32`

GetMaximumTries returns the MaximumTries field if non-nil, zero value otherwise.

### GetMaximumTriesOk

`func (o *GatherUsingSpeakRequest) GetMaximumTriesOk() (*int32, bool)`

GetMaximumTriesOk returns a tuple with the MaximumTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTries

`func (o *GatherUsingSpeakRequest) SetMaximumTries(v int32)`

SetMaximumTries sets MaximumTries field to given value.

### HasMaximumTries

`func (o *GatherUsingSpeakRequest) HasMaximumTries() bool`

HasMaximumTries returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *GatherUsingSpeakRequest) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *GatherUsingSpeakRequest) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *GatherUsingSpeakRequest) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *GatherUsingSpeakRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetTerminatingDigit

`func (o *GatherUsingSpeakRequest) GetTerminatingDigit() string`

GetTerminatingDigit returns the TerminatingDigit field if non-nil, zero value otherwise.

### GetTerminatingDigitOk

`func (o *GatherUsingSpeakRequest) GetTerminatingDigitOk() (*string, bool)`

GetTerminatingDigitOk returns a tuple with the TerminatingDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingDigit

`func (o *GatherUsingSpeakRequest) SetTerminatingDigit(v string)`

SetTerminatingDigit sets TerminatingDigit field to given value.

### HasTerminatingDigit

`func (o *GatherUsingSpeakRequest) HasTerminatingDigit() bool`

HasTerminatingDigit returns a boolean if a field has been set.

### GetValidDigits

`func (o *GatherUsingSpeakRequest) GetValidDigits() string`

GetValidDigits returns the ValidDigits field if non-nil, zero value otherwise.

### GetValidDigitsOk

`func (o *GatherUsingSpeakRequest) GetValidDigitsOk() (*string, bool)`

GetValidDigitsOk returns a tuple with the ValidDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDigits

`func (o *GatherUsingSpeakRequest) SetValidDigits(v string)`

SetValidDigits sets ValidDigits field to given value.

### HasValidDigits

`func (o *GatherUsingSpeakRequest) HasValidDigits() bool`

HasValidDigits returns a boolean if a field has been set.

### GetInterDigitTimeoutMillis

`func (o *GatherUsingSpeakRequest) GetInterDigitTimeoutMillis() int32`

GetInterDigitTimeoutMillis returns the InterDigitTimeoutMillis field if non-nil, zero value otherwise.

### GetInterDigitTimeoutMillisOk

`func (o *GatherUsingSpeakRequest) GetInterDigitTimeoutMillisOk() (*int32, bool)`

GetInterDigitTimeoutMillisOk returns a tuple with the InterDigitTimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterDigitTimeoutMillis

`func (o *GatherUsingSpeakRequest) SetInterDigitTimeoutMillis(v int32)`

SetInterDigitTimeoutMillis sets InterDigitTimeoutMillis field to given value.

### HasInterDigitTimeoutMillis

`func (o *GatherUsingSpeakRequest) HasInterDigitTimeoutMillis() bool`

HasInterDigitTimeoutMillis returns a boolean if a field has been set.

### GetClientState

`func (o *GatherUsingSpeakRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *GatherUsingSpeakRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *GatherUsingSpeakRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *GatherUsingSpeakRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *GatherUsingSpeakRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *GatherUsingSpeakRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *GatherUsingSpeakRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *GatherUsingSpeakRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


