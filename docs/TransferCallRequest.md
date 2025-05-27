# TransferCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The DID or SIP URI to dial out to. | 
**From** | Pointer to **string** | The &#x60;from&#x60; number to be used as the caller id presented to the destination (&#x60;to&#x60; number). The number should be in +E164 format. This attribute will default to the &#x60;to&#x60; number of the original call if omitted. | [optional] 
**FromDisplayName** | Pointer to **string** | The &#x60;from_display_name&#x60; string to be used as the caller id name (SIP From Display Name) presented to the destination (&#x60;to&#x60; number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the &#x60;from&#x60; field. | [optional] 
**AudioUrl** | Pointer to **string** | The URL of a file to be played back when the transfer destination answers before bridging the call. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request. | [optional] 
**EarlyMedia** | Pointer to **bool** | If set to false, early media will not be passed to the originating leg. | [optional] [default to true]
**MediaName** | Pointer to **string** | The media_name of a file to be played back when the transfer destination answers before bridging the call. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 
**TimeoutSecs** | Pointer to **int32** | The number of seconds that Telnyx will wait for the call to be answered by the destination to which it is being transferred. If the timeout is reached before an answer is received, the call will hangup and a &#x60;call.hangup&#x60; webhook with a &#x60;hangup_cause&#x60; of &#x60;timeout&#x60; will be sent. Minimum value is 5 seconds. Maximum value is 600 seconds. | [optional] [default to 30]
**TimeLimitSecs** | Pointer to **int32** | Sets the maximum duration of a Call Control Leg in seconds. If the time limit is reached, the call will hangup and a &#x60;call.hangup&#x60; webhook with a &#x60;hangup_cause&#x60; of &#x60;time_limit&#x60; will be sent. For example, by setting a time limit of 120 seconds, a Call Leg will be automatically terminated two minutes after being answered. The default time limit is 14400 seconds or 4 hours and this is also the maximum allowed call length. | [optional] [default to 14400]
**AnsweringMachineDetection** | Pointer to **string** | Enables Answering Machine Detection. When a call is answered, Telnyx runs real-time detection to determine if it was picked up by a human or a machine and sends an &#x60;call.machine.detection.ended&#x60; webhook with the analysis result. If &#39;greeting_end&#39; or &#39;detect_words&#39; is used and a &#39;machine&#39; is detected, you will receive another &#39;call.machine.greeting.ended&#39; webhook when the answering machine greeting ends with a beep or silence. If &#x60;detect_beep&#x60; is used, you will only receive &#39;call.machine.greeting.ended&#39; if a beep is detected. | [optional] [default to "disabled"]
**AnsweringMachineDetectionConfig** | Pointer to [**CallRequestAnsweringMachineDetectionConfig**](CallRequestAnsweringMachineDetectionConfig.md) |  | [optional] 
**CustomHeaders** | Pointer to [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers to be added to the SIP INVITE. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**TargetLegClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook for the new leg. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**MediaEncryption** | Pointer to **string** | Defines whether media should be encrypted on the new call leg. | [optional] [default to "disabled"]
**SipAuthUsername** | Pointer to **string** | SIP Authentication username used for SIP challenges. | [optional] 
**SipAuthPassword** | Pointer to **string** | SIP Authentication password used for SIP challenges. | [optional] 
**SipHeaders** | Pointer to [**[]SipHeader**](SipHeader.md) | SIP headers to be added to the SIP INVITE. Currently only User-to-User header is supported. | [optional] 
**SipTransportProtocol** | Pointer to **string** | Defines SIP transport protocol to be used on the call. | [optional] [default to "UDP"]
**SoundModifications** | Pointer to [**SoundModifications**](SoundModifications.md) |  | [optional] 
**WebhookUrl** | Pointer to **string** | Use this field to override the URL for which Telnyx will send subsequent webhooks to for this call. | [optional] 
**WebhookUrlMethod** | Pointer to **string** | HTTP request type used for &#x60;webhook_url&#x60;. | [optional] [default to "POST"]
**MuteDtmf** | Pointer to **string** | When enabled, DTMF tones are not passed to the call participant. The webhooks containing the DTMF information will be sent. | [optional] [default to "none"]

## Methods

### NewTransferCallRequest

`func NewTransferCallRequest(to string, ) *TransferCallRequest`

NewTransferCallRequest instantiates a new TransferCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferCallRequestWithDefaults

`func NewTransferCallRequestWithDefaults() *TransferCallRequest`

NewTransferCallRequestWithDefaults instantiates a new TransferCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *TransferCallRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TransferCallRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TransferCallRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *TransferCallRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransferCallRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransferCallRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TransferCallRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromDisplayName

`func (o *TransferCallRequest) GetFromDisplayName() string`

GetFromDisplayName returns the FromDisplayName field if non-nil, zero value otherwise.

### GetFromDisplayNameOk

`func (o *TransferCallRequest) GetFromDisplayNameOk() (*string, bool)`

GetFromDisplayNameOk returns a tuple with the FromDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDisplayName

`func (o *TransferCallRequest) SetFromDisplayName(v string)`

SetFromDisplayName sets FromDisplayName field to given value.

### HasFromDisplayName

`func (o *TransferCallRequest) HasFromDisplayName() bool`

HasFromDisplayName returns a boolean if a field has been set.

### GetAudioUrl

`func (o *TransferCallRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *TransferCallRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *TransferCallRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *TransferCallRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetEarlyMedia

`func (o *TransferCallRequest) GetEarlyMedia() bool`

GetEarlyMedia returns the EarlyMedia field if non-nil, zero value otherwise.

### GetEarlyMediaOk

`func (o *TransferCallRequest) GetEarlyMediaOk() (*bool, bool)`

GetEarlyMediaOk returns a tuple with the EarlyMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyMedia

`func (o *TransferCallRequest) SetEarlyMedia(v bool)`

SetEarlyMedia sets EarlyMedia field to given value.

### HasEarlyMedia

`func (o *TransferCallRequest) HasEarlyMedia() bool`

HasEarlyMedia returns a boolean if a field has been set.

### GetMediaName

`func (o *TransferCallRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *TransferCallRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *TransferCallRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *TransferCallRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *TransferCallRequest) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *TransferCallRequest) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *TransferCallRequest) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *TransferCallRequest) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetTimeLimitSecs

`func (o *TransferCallRequest) GetTimeLimitSecs() int32`

GetTimeLimitSecs returns the TimeLimitSecs field if non-nil, zero value otherwise.

### GetTimeLimitSecsOk

`func (o *TransferCallRequest) GetTimeLimitSecsOk() (*int32, bool)`

GetTimeLimitSecsOk returns a tuple with the TimeLimitSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitSecs

`func (o *TransferCallRequest) SetTimeLimitSecs(v int32)`

SetTimeLimitSecs sets TimeLimitSecs field to given value.

### HasTimeLimitSecs

`func (o *TransferCallRequest) HasTimeLimitSecs() bool`

HasTimeLimitSecs returns a boolean if a field has been set.

### GetAnsweringMachineDetection

`func (o *TransferCallRequest) GetAnsweringMachineDetection() string`

GetAnsweringMachineDetection returns the AnsweringMachineDetection field if non-nil, zero value otherwise.

### GetAnsweringMachineDetectionOk

`func (o *TransferCallRequest) GetAnsweringMachineDetectionOk() (*string, bool)`

GetAnsweringMachineDetectionOk returns a tuple with the AnsweringMachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweringMachineDetection

`func (o *TransferCallRequest) SetAnsweringMachineDetection(v string)`

SetAnsweringMachineDetection sets AnsweringMachineDetection field to given value.

### HasAnsweringMachineDetection

`func (o *TransferCallRequest) HasAnsweringMachineDetection() bool`

HasAnsweringMachineDetection returns a boolean if a field has been set.

### GetAnsweringMachineDetectionConfig

`func (o *TransferCallRequest) GetAnsweringMachineDetectionConfig() CallRequestAnsweringMachineDetectionConfig`

GetAnsweringMachineDetectionConfig returns the AnsweringMachineDetectionConfig field if non-nil, zero value otherwise.

### GetAnsweringMachineDetectionConfigOk

`func (o *TransferCallRequest) GetAnsweringMachineDetectionConfigOk() (*CallRequestAnsweringMachineDetectionConfig, bool)`

GetAnsweringMachineDetectionConfigOk returns a tuple with the AnsweringMachineDetectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweringMachineDetectionConfig

`func (o *TransferCallRequest) SetAnsweringMachineDetectionConfig(v CallRequestAnsweringMachineDetectionConfig)`

SetAnsweringMachineDetectionConfig sets AnsweringMachineDetectionConfig field to given value.

### HasAnsweringMachineDetectionConfig

`func (o *TransferCallRequest) HasAnsweringMachineDetectionConfig() bool`

HasAnsweringMachineDetectionConfig returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *TransferCallRequest) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *TransferCallRequest) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *TransferCallRequest) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *TransferCallRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetClientState

`func (o *TransferCallRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *TransferCallRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *TransferCallRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *TransferCallRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetTargetLegClientState

`func (o *TransferCallRequest) GetTargetLegClientState() string`

GetTargetLegClientState returns the TargetLegClientState field if non-nil, zero value otherwise.

### GetTargetLegClientStateOk

`func (o *TransferCallRequest) GetTargetLegClientStateOk() (*string, bool)`

GetTargetLegClientStateOk returns a tuple with the TargetLegClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLegClientState

`func (o *TransferCallRequest) SetTargetLegClientState(v string)`

SetTargetLegClientState sets TargetLegClientState field to given value.

### HasTargetLegClientState

`func (o *TransferCallRequest) HasTargetLegClientState() bool`

HasTargetLegClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *TransferCallRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *TransferCallRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *TransferCallRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *TransferCallRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetMediaEncryption

`func (o *TransferCallRequest) GetMediaEncryption() string`

GetMediaEncryption returns the MediaEncryption field if non-nil, zero value otherwise.

### GetMediaEncryptionOk

`func (o *TransferCallRequest) GetMediaEncryptionOk() (*string, bool)`

GetMediaEncryptionOk returns a tuple with the MediaEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEncryption

`func (o *TransferCallRequest) SetMediaEncryption(v string)`

SetMediaEncryption sets MediaEncryption field to given value.

### HasMediaEncryption

`func (o *TransferCallRequest) HasMediaEncryption() bool`

HasMediaEncryption returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *TransferCallRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *TransferCallRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *TransferCallRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *TransferCallRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### GetSipAuthPassword

`func (o *TransferCallRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *TransferCallRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *TransferCallRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *TransferCallRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### GetSipHeaders

`func (o *TransferCallRequest) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *TransferCallRequest) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *TransferCallRequest) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *TransferCallRequest) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.

### GetSipTransportProtocol

`func (o *TransferCallRequest) GetSipTransportProtocol() string`

GetSipTransportProtocol returns the SipTransportProtocol field if non-nil, zero value otherwise.

### GetSipTransportProtocolOk

`func (o *TransferCallRequest) GetSipTransportProtocolOk() (*string, bool)`

GetSipTransportProtocolOk returns a tuple with the SipTransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTransportProtocol

`func (o *TransferCallRequest) SetSipTransportProtocol(v string)`

SetSipTransportProtocol sets SipTransportProtocol field to given value.

### HasSipTransportProtocol

`func (o *TransferCallRequest) HasSipTransportProtocol() bool`

HasSipTransportProtocol returns a boolean if a field has been set.

### GetSoundModifications

`func (o *TransferCallRequest) GetSoundModifications() SoundModifications`

GetSoundModifications returns the SoundModifications field if non-nil, zero value otherwise.

### GetSoundModificationsOk

`func (o *TransferCallRequest) GetSoundModificationsOk() (*SoundModifications, bool)`

GetSoundModificationsOk returns a tuple with the SoundModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundModifications

`func (o *TransferCallRequest) SetSoundModifications(v SoundModifications)`

SetSoundModifications sets SoundModifications field to given value.

### HasSoundModifications

`func (o *TransferCallRequest) HasSoundModifications() bool`

HasSoundModifications returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *TransferCallRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *TransferCallRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *TransferCallRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *TransferCallRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlMethod

`func (o *TransferCallRequest) GetWebhookUrlMethod() string`

GetWebhookUrlMethod returns the WebhookUrlMethod field if non-nil, zero value otherwise.

### GetWebhookUrlMethodOk

`func (o *TransferCallRequest) GetWebhookUrlMethodOk() (*string, bool)`

GetWebhookUrlMethodOk returns a tuple with the WebhookUrlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlMethod

`func (o *TransferCallRequest) SetWebhookUrlMethod(v string)`

SetWebhookUrlMethod sets WebhookUrlMethod field to given value.

### HasWebhookUrlMethod

`func (o *TransferCallRequest) HasWebhookUrlMethod() bool`

HasWebhookUrlMethod returns a boolean if a field has been set.

### GetMuteDtmf

`func (o *TransferCallRequest) GetMuteDtmf() string`

GetMuteDtmf returns the MuteDtmf field if non-nil, zero value otherwise.

### GetMuteDtmfOk

`func (o *TransferCallRequest) GetMuteDtmfOk() (*string, bool)`

GetMuteDtmfOk returns a tuple with the MuteDtmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteDtmf

`func (o *TransferCallRequest) SetMuteDtmf(v string)`

SetMuteDtmf sets MuteDtmf field to given value.

### HasMuteDtmf

`func (o *TransferCallRequest) HasMuteDtmf() bool`

HasMuteDtmf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


