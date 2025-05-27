# CallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**CallRequestTo**](CallRequestTo.md) |  | 
**From** | **string** | The &#x60;from&#x60; number to be used as the caller id presented to the destination (&#x60;to&#x60; number). The number should be in +E164 format. | 
**FromDisplayName** | Pointer to **string** | The &#x60;from_display_name&#x60; string to be used as the caller id name (SIP From Display Name) presented to the destination (&#x60;to&#x60; number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and -_~!.+ special characters. If ommited, the display name will be the same as the number in the &#x60;from&#x60; field. | [optional] 
**ConnectionId** | **string** | The ID of the Call Control App (formerly ID of the connection) to be used when dialing the destination. | 
**AudioUrl** | Pointer to **string** | The URL of a file to be played back to the callee when the call is answered. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request. | [optional] 
**MediaName** | Pointer to **string** | The media_name of a file to be played back to the callee when the call is answered. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 
**PreferredCodecs** | Pointer to **string** | The list of comma-separated codecs in a preferred order for the forked media to be received. | [optional] 
**TimeoutSecs** | Pointer to **int32** | The number of seconds that Telnyx will wait for the call to be answered by the destination to which it is being called. If the timeout is reached before an answer is received, the call will hangup and a &#x60;call.hangup&#x60; webhook with a &#x60;hangup_cause&#x60; of &#x60;timeout&#x60; will be sent. Minimum value is 5 seconds. Maximum value is 600 seconds. | [optional] [default to 30]
**TimeLimitSecs** | Pointer to **int32** | Sets the maximum duration of a Call Control Leg in seconds. If the time limit is reached, the call will hangup and a &#x60;call.hangup&#x60; webhook with a &#x60;hangup_cause&#x60; of &#x60;time_limit&#x60; will be sent. For example, by setting a time limit of 120 seconds, a Call Leg will be automatically terminated two minutes after being answered. The default time limit is 14400 seconds or 4 hours and this is also the maximum allowed call length. | [optional] [default to 14400]
**AnsweringMachineDetection** | Pointer to **string** | Enables Answering Machine Detection. Telnyx offers Premium and Standard detections. With Premium detection, when a call is answered, Telnyx runs real-time detection and sends a &#x60;call.machine.premium.detection.ended&#x60; webhook with one of the following results: &#x60;human_residence&#x60;, &#x60;human_business&#x60;, &#x60;machine&#x60;, &#x60;silence&#x60; or &#x60;fax_detected&#x60;. If we detect a beep, we also send a &#x60;call.machine.premium.greeting.ended&#x60; webhook with the result of &#x60;beep_detected&#x60;. If we detect a beep before &#x60;call.machine.premium.detection.ended&#x60; we only send &#x60;call.machine.premium.greeting.ended&#x60;, and if we detect a beep after &#x60;call.machine.premium.detection.ended&#x60;, we send both webhooks. With Standard detection, when a call is answered, Telnyx runs real-time detection to determine if it was picked up by a human or a machine and sends an &#x60;call.machine.detection.ended&#x60; webhook with the analysis result. If &#x60;greeting_end&#x60; or &#x60;detect_words&#x60; is used and a &#x60;machine&#x60; is detected, you will receive another &#x60;call.machine.greeting.ended&#x60; webhook when the answering machine greeting ends with a beep or silence. If &#x60;detect_beep&#x60; is used, you will only receive &#x60;call.machine.greeting.ended&#x60; if a beep is detected. | [optional] [default to "disabled"]
**AnsweringMachineDetectionConfig** | Pointer to [**CallRequestAnsweringMachineDetectionConfig**](CallRequestAnsweringMachineDetectionConfig.md) |  | [optional] 
**ConferenceConfig** | Pointer to [**CallRequestConferenceConfig**](CallRequestConferenceConfig.md) |  | [optional] 
**CustomHeaders** | Pointer to [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers to be added to the SIP INVITE. | [optional] 
**BillingGroupId** | Pointer to **string** | Use this field to set the Billing Group ID for the call. Must be a valid and existing Billing Group ID. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore others Dial commands with the same &#x60;command_id&#x60;. | [optional] 
**LinkTo** | Pointer to **string** | Use another call&#39;s control id for sharing the same call session id | [optional] 
**MediaEncryption** | Pointer to **string** | Defines whether media should be encrypted on the call. | [optional] [default to "disabled"]
**SipAuthUsername** | Pointer to **string** | SIP Authentication username used for SIP challenges. | [optional] 
**SipAuthPassword** | Pointer to **string** | SIP Authentication password used for SIP challenges. | [optional] 
**SipHeaders** | Pointer to [**[]SipHeader**](SipHeader.md) | SIP headers to be added to the SIP INVITE request. Currently only User-to-User header is supported. | [optional] 
**SipTransportProtocol** | Pointer to **string** | Defines SIP transport protocol to be used on the call. | [optional] [default to "UDP"]
**SoundModifications** | Pointer to [**SoundModifications**](SoundModifications.md) |  | [optional] 
**StreamUrl** | Pointer to **string** | The destination WebSocket address where the stream is going to be delivered. | [optional] 
**StreamTrack** | Pointer to **string** | Specifies which track should be streamed. | [optional] [default to "inbound_track"]
**StreamBidirectionalMode** | Pointer to [**StreamBidirectionalMode**](StreamBidirectionalMode.md) |  | [optional] [default to MP3]
**StreamBidirectionalCodec** | Pointer to [**StreamBidirectionalCodec**](StreamBidirectionalCodec.md) |  | [optional] [default to PCMU]
**StreamBidirectionalTargetLegs** | Pointer to [**StreamBidirectionalTargetLegs**](StreamBidirectionalTargetLegs.md) |  | [optional] [default to OPPOSITE]
**StreamBidirectionalSamplingRate** | Pointer to [**StreamBidirectionalSamplingRate**](StreamBidirectionalSamplingRate.md) |  | [optional] [default to _8000]
**StreamEstablishBeforeCallOriginate** | Pointer to **bool** | Establish websocket connection before dialing the destination. This is useful for cases where the websocket connection takes a long time to establish. | [optional] [default to false]
**SendSilenceWhenIdle** | Pointer to **bool** | Generate silence RTP packets when no transmission available. | [optional] [default to false]
**WebhookUrl** | Pointer to **string** | Use this field to override the URL for which Telnyx will send subsequent webhooks to for this call. | [optional] 
**WebhookUrlMethod** | Pointer to **string** | HTTP request type used for &#x60;webhook_url&#x60;. | [optional] [default to "POST"]
**Record** | Pointer to **string** | Start recording automatically after an event. Disabled by default. | [optional] 
**RecordChannels** | Pointer to **string** | Defines which channel should be recorded (&#39;single&#39; or &#39;dual&#39;) when &#x60;record&#x60; is specified. | [optional] [default to "dual"]
**RecordFormat** | Pointer to **string** | Defines the format of the recording (&#39;wav&#39; or &#39;mp3&#39;) when &#x60;record&#x60; is specified. | [optional] [default to "mp3"]
**RecordMaxLength** | Pointer to **int32** | Defines the maximum length for the recording in seconds when &#x60;record&#x60; is specified. The minimum value is 0. The maximum value is 43200. The default value is 0 (infinite). | [optional] [default to 0]
**RecordTimeoutSecs** | Pointer to **int32** | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected when &#x60;record&#x60; is specified. The timer only starts when the speech is detected. Please note that call transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite). | [optional] [default to 0]
**RecordTrack** | Pointer to **string** | The audio track to be recorded. Can be either &#x60;both&#x60;, &#x60;inbound&#x60; or &#x60;outbound&#x60;. If only single track is specified (&#x60;inbound&#x60;, &#x60;outbound&#x60;), &#x60;channels&#x60; configuration is ignored and it will be recorded as mono (single channel). | [optional] [default to "both"]
**RecordTrim** | Pointer to **string** | When set to &#x60;trim-silence&#x60;, silence will be removed from the beginning and end of the recording. | [optional] 
**RecordCustomFileName** | Pointer to **string** | The custom recording file name to be used instead of the default &#x60;call_leg_id&#x60;. Telnyx will still add a Unix timestamp suffix. | [optional] 
**SuperviseCallControlId** | Pointer to **string** | The call leg which will be supervised by the new call. | [optional] 
**SupervisorRole** | Pointer to **string** | The role of the supervisor call. &#39;barge&#39; means that supervisor call hears and is being heard by both ends of the call (caller &amp; callee). &#39;whisper&#39; means that only supervised_call_control_id hears supervisor but supervisor can hear everything. &#39;monitor&#39; means that nobody can hear supervisor call, but supervisor can hear everything on the call. | [optional] [default to "barge"]
**EnableDialogflow** | Pointer to **bool** | Enables Dialogflow for the current call. The default value is false. | [optional] [default to false]
**DialogflowConfig** | Pointer to [**DialogflowConfig**](DialogflowConfig.md) |  | [optional] 
**Transcription** | Pointer to **bool** | Enable transcription upon call answer. The default value is false. | [optional] [default to false]
**TranscriptionConfig** | Pointer to [**TranscriptionStartRequest**](TranscriptionStartRequest.md) |  | [optional] 

## Methods

### NewCallRequest

`func NewCallRequest(to CallRequestTo, from string, connectionId string, ) *CallRequest`

NewCallRequest instantiates a new CallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestWithDefaults

`func NewCallRequestWithDefaults() *CallRequest`

NewCallRequestWithDefaults instantiates a new CallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *CallRequest) GetTo() CallRequestTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallRequest) GetToOk() (*CallRequestTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallRequest) SetTo(v CallRequestTo)`

SetTo sets To field to given value.


### GetFrom

`func (o *CallRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetFromDisplayName

`func (o *CallRequest) GetFromDisplayName() string`

GetFromDisplayName returns the FromDisplayName field if non-nil, zero value otherwise.

### GetFromDisplayNameOk

`func (o *CallRequest) GetFromDisplayNameOk() (*string, bool)`

GetFromDisplayNameOk returns a tuple with the FromDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDisplayName

`func (o *CallRequest) SetFromDisplayName(v string)`

SetFromDisplayName sets FromDisplayName field to given value.

### HasFromDisplayName

`func (o *CallRequest) HasFromDisplayName() bool`

HasFromDisplayName returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetAudioUrl

`func (o *CallRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *CallRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *CallRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *CallRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *CallRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *CallRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *CallRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *CallRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetPreferredCodecs

`func (o *CallRequest) GetPreferredCodecs() string`

GetPreferredCodecs returns the PreferredCodecs field if non-nil, zero value otherwise.

### GetPreferredCodecsOk

`func (o *CallRequest) GetPreferredCodecsOk() (*string, bool)`

GetPreferredCodecsOk returns a tuple with the PreferredCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredCodecs

`func (o *CallRequest) SetPreferredCodecs(v string)`

SetPreferredCodecs sets PreferredCodecs field to given value.

### HasPreferredCodecs

`func (o *CallRequest) HasPreferredCodecs() bool`

HasPreferredCodecs returns a boolean if a field has been set.

### GetTimeoutSecs

`func (o *CallRequest) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CallRequest) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CallRequest) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CallRequest) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetTimeLimitSecs

`func (o *CallRequest) GetTimeLimitSecs() int32`

GetTimeLimitSecs returns the TimeLimitSecs field if non-nil, zero value otherwise.

### GetTimeLimitSecsOk

`func (o *CallRequest) GetTimeLimitSecsOk() (*int32, bool)`

GetTimeLimitSecsOk returns a tuple with the TimeLimitSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitSecs

`func (o *CallRequest) SetTimeLimitSecs(v int32)`

SetTimeLimitSecs sets TimeLimitSecs field to given value.

### HasTimeLimitSecs

`func (o *CallRequest) HasTimeLimitSecs() bool`

HasTimeLimitSecs returns a boolean if a field has been set.

### GetAnsweringMachineDetection

`func (o *CallRequest) GetAnsweringMachineDetection() string`

GetAnsweringMachineDetection returns the AnsweringMachineDetection field if non-nil, zero value otherwise.

### GetAnsweringMachineDetectionOk

`func (o *CallRequest) GetAnsweringMachineDetectionOk() (*string, bool)`

GetAnsweringMachineDetectionOk returns a tuple with the AnsweringMachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweringMachineDetection

`func (o *CallRequest) SetAnsweringMachineDetection(v string)`

SetAnsweringMachineDetection sets AnsweringMachineDetection field to given value.

### HasAnsweringMachineDetection

`func (o *CallRequest) HasAnsweringMachineDetection() bool`

HasAnsweringMachineDetection returns a boolean if a field has been set.

### GetAnsweringMachineDetectionConfig

`func (o *CallRequest) GetAnsweringMachineDetectionConfig() CallRequestAnsweringMachineDetectionConfig`

GetAnsweringMachineDetectionConfig returns the AnsweringMachineDetectionConfig field if non-nil, zero value otherwise.

### GetAnsweringMachineDetectionConfigOk

`func (o *CallRequest) GetAnsweringMachineDetectionConfigOk() (*CallRequestAnsweringMachineDetectionConfig, bool)`

GetAnsweringMachineDetectionConfigOk returns a tuple with the AnsweringMachineDetectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweringMachineDetectionConfig

`func (o *CallRequest) SetAnsweringMachineDetectionConfig(v CallRequestAnsweringMachineDetectionConfig)`

SetAnsweringMachineDetectionConfig sets AnsweringMachineDetectionConfig field to given value.

### HasAnsweringMachineDetectionConfig

`func (o *CallRequest) HasAnsweringMachineDetectionConfig() bool`

HasAnsweringMachineDetectionConfig returns a boolean if a field has been set.

### GetConferenceConfig

`func (o *CallRequest) GetConferenceConfig() CallRequestConferenceConfig`

GetConferenceConfig returns the ConferenceConfig field if non-nil, zero value otherwise.

### GetConferenceConfigOk

`func (o *CallRequest) GetConferenceConfigOk() (*CallRequestConferenceConfig, bool)`

GetConferenceConfigOk returns a tuple with the ConferenceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceConfig

`func (o *CallRequest) SetConferenceConfig(v CallRequestConferenceConfig)`

SetConferenceConfig sets ConferenceConfig field to given value.

### HasConferenceConfig

`func (o *CallRequest) HasConferenceConfig() bool`

HasConferenceConfig returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *CallRequest) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CallRequest) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CallRequest) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CallRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *CallRequest) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *CallRequest) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *CallRequest) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *CallRequest) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetClientState

`func (o *CallRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *CallRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *CallRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *CallRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *CallRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetLinkTo

`func (o *CallRequest) GetLinkTo() string`

GetLinkTo returns the LinkTo field if non-nil, zero value otherwise.

### GetLinkToOk

`func (o *CallRequest) GetLinkToOk() (*string, bool)`

GetLinkToOk returns a tuple with the LinkTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTo

`func (o *CallRequest) SetLinkTo(v string)`

SetLinkTo sets LinkTo field to given value.

### HasLinkTo

`func (o *CallRequest) HasLinkTo() bool`

HasLinkTo returns a boolean if a field has been set.

### GetMediaEncryption

`func (o *CallRequest) GetMediaEncryption() string`

GetMediaEncryption returns the MediaEncryption field if non-nil, zero value otherwise.

### GetMediaEncryptionOk

`func (o *CallRequest) GetMediaEncryptionOk() (*string, bool)`

GetMediaEncryptionOk returns a tuple with the MediaEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEncryption

`func (o *CallRequest) SetMediaEncryption(v string)`

SetMediaEncryption sets MediaEncryption field to given value.

### HasMediaEncryption

`func (o *CallRequest) HasMediaEncryption() bool`

HasMediaEncryption returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *CallRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *CallRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *CallRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *CallRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### GetSipAuthPassword

`func (o *CallRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *CallRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *CallRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *CallRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### GetSipHeaders

`func (o *CallRequest) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *CallRequest) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *CallRequest) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *CallRequest) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.

### GetSipTransportProtocol

`func (o *CallRequest) GetSipTransportProtocol() string`

GetSipTransportProtocol returns the SipTransportProtocol field if non-nil, zero value otherwise.

### GetSipTransportProtocolOk

`func (o *CallRequest) GetSipTransportProtocolOk() (*string, bool)`

GetSipTransportProtocolOk returns a tuple with the SipTransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTransportProtocol

`func (o *CallRequest) SetSipTransportProtocol(v string)`

SetSipTransportProtocol sets SipTransportProtocol field to given value.

### HasSipTransportProtocol

`func (o *CallRequest) HasSipTransportProtocol() bool`

HasSipTransportProtocol returns a boolean if a field has been set.

### GetSoundModifications

`func (o *CallRequest) GetSoundModifications() SoundModifications`

GetSoundModifications returns the SoundModifications field if non-nil, zero value otherwise.

### GetSoundModificationsOk

`func (o *CallRequest) GetSoundModificationsOk() (*SoundModifications, bool)`

GetSoundModificationsOk returns a tuple with the SoundModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundModifications

`func (o *CallRequest) SetSoundModifications(v SoundModifications)`

SetSoundModifications sets SoundModifications field to given value.

### HasSoundModifications

`func (o *CallRequest) HasSoundModifications() bool`

HasSoundModifications returns a boolean if a field has been set.

### GetStreamUrl

`func (o *CallRequest) GetStreamUrl() string`

GetStreamUrl returns the StreamUrl field if non-nil, zero value otherwise.

### GetStreamUrlOk

`func (o *CallRequest) GetStreamUrlOk() (*string, bool)`

GetStreamUrlOk returns a tuple with the StreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamUrl

`func (o *CallRequest) SetStreamUrl(v string)`

SetStreamUrl sets StreamUrl field to given value.

### HasStreamUrl

`func (o *CallRequest) HasStreamUrl() bool`

HasStreamUrl returns a boolean if a field has been set.

### GetStreamTrack

`func (o *CallRequest) GetStreamTrack() string`

GetStreamTrack returns the StreamTrack field if non-nil, zero value otherwise.

### GetStreamTrackOk

`func (o *CallRequest) GetStreamTrackOk() (*string, bool)`

GetStreamTrackOk returns a tuple with the StreamTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamTrack

`func (o *CallRequest) SetStreamTrack(v string)`

SetStreamTrack sets StreamTrack field to given value.

### HasStreamTrack

`func (o *CallRequest) HasStreamTrack() bool`

HasStreamTrack returns a boolean if a field has been set.

### GetStreamBidirectionalMode

`func (o *CallRequest) GetStreamBidirectionalMode() StreamBidirectionalMode`

GetStreamBidirectionalMode returns the StreamBidirectionalMode field if non-nil, zero value otherwise.

### GetStreamBidirectionalModeOk

`func (o *CallRequest) GetStreamBidirectionalModeOk() (*StreamBidirectionalMode, bool)`

GetStreamBidirectionalModeOk returns a tuple with the StreamBidirectionalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalMode

`func (o *CallRequest) SetStreamBidirectionalMode(v StreamBidirectionalMode)`

SetStreamBidirectionalMode sets StreamBidirectionalMode field to given value.

### HasStreamBidirectionalMode

`func (o *CallRequest) HasStreamBidirectionalMode() bool`

HasStreamBidirectionalMode returns a boolean if a field has been set.

### GetStreamBidirectionalCodec

`func (o *CallRequest) GetStreamBidirectionalCodec() StreamBidirectionalCodec`

GetStreamBidirectionalCodec returns the StreamBidirectionalCodec field if non-nil, zero value otherwise.

### GetStreamBidirectionalCodecOk

`func (o *CallRequest) GetStreamBidirectionalCodecOk() (*StreamBidirectionalCodec, bool)`

GetStreamBidirectionalCodecOk returns a tuple with the StreamBidirectionalCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalCodec

`func (o *CallRequest) SetStreamBidirectionalCodec(v StreamBidirectionalCodec)`

SetStreamBidirectionalCodec sets StreamBidirectionalCodec field to given value.

### HasStreamBidirectionalCodec

`func (o *CallRequest) HasStreamBidirectionalCodec() bool`

HasStreamBidirectionalCodec returns a boolean if a field has been set.

### GetStreamBidirectionalTargetLegs

`func (o *CallRequest) GetStreamBidirectionalTargetLegs() StreamBidirectionalTargetLegs`

GetStreamBidirectionalTargetLegs returns the StreamBidirectionalTargetLegs field if non-nil, zero value otherwise.

### GetStreamBidirectionalTargetLegsOk

`func (o *CallRequest) GetStreamBidirectionalTargetLegsOk() (*StreamBidirectionalTargetLegs, bool)`

GetStreamBidirectionalTargetLegsOk returns a tuple with the StreamBidirectionalTargetLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalTargetLegs

`func (o *CallRequest) SetStreamBidirectionalTargetLegs(v StreamBidirectionalTargetLegs)`

SetStreamBidirectionalTargetLegs sets StreamBidirectionalTargetLegs field to given value.

### HasStreamBidirectionalTargetLegs

`func (o *CallRequest) HasStreamBidirectionalTargetLegs() bool`

HasStreamBidirectionalTargetLegs returns a boolean if a field has been set.

### GetStreamBidirectionalSamplingRate

`func (o *CallRequest) GetStreamBidirectionalSamplingRate() StreamBidirectionalSamplingRate`

GetStreamBidirectionalSamplingRate returns the StreamBidirectionalSamplingRate field if non-nil, zero value otherwise.

### GetStreamBidirectionalSamplingRateOk

`func (o *CallRequest) GetStreamBidirectionalSamplingRateOk() (*StreamBidirectionalSamplingRate, bool)`

GetStreamBidirectionalSamplingRateOk returns a tuple with the StreamBidirectionalSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalSamplingRate

`func (o *CallRequest) SetStreamBidirectionalSamplingRate(v StreamBidirectionalSamplingRate)`

SetStreamBidirectionalSamplingRate sets StreamBidirectionalSamplingRate field to given value.

### HasStreamBidirectionalSamplingRate

`func (o *CallRequest) HasStreamBidirectionalSamplingRate() bool`

HasStreamBidirectionalSamplingRate returns a boolean if a field has been set.

### GetStreamEstablishBeforeCallOriginate

`func (o *CallRequest) GetStreamEstablishBeforeCallOriginate() bool`

GetStreamEstablishBeforeCallOriginate returns the StreamEstablishBeforeCallOriginate field if non-nil, zero value otherwise.

### GetStreamEstablishBeforeCallOriginateOk

`func (o *CallRequest) GetStreamEstablishBeforeCallOriginateOk() (*bool, bool)`

GetStreamEstablishBeforeCallOriginateOk returns a tuple with the StreamEstablishBeforeCallOriginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamEstablishBeforeCallOriginate

`func (o *CallRequest) SetStreamEstablishBeforeCallOriginate(v bool)`

SetStreamEstablishBeforeCallOriginate sets StreamEstablishBeforeCallOriginate field to given value.

### HasStreamEstablishBeforeCallOriginate

`func (o *CallRequest) HasStreamEstablishBeforeCallOriginate() bool`

HasStreamEstablishBeforeCallOriginate returns a boolean if a field has been set.

### GetSendSilenceWhenIdle

`func (o *CallRequest) GetSendSilenceWhenIdle() bool`

GetSendSilenceWhenIdle returns the SendSilenceWhenIdle field if non-nil, zero value otherwise.

### GetSendSilenceWhenIdleOk

`func (o *CallRequest) GetSendSilenceWhenIdleOk() (*bool, bool)`

GetSendSilenceWhenIdleOk returns a tuple with the SendSilenceWhenIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendSilenceWhenIdle

`func (o *CallRequest) SetSendSilenceWhenIdle(v bool)`

SetSendSilenceWhenIdle sets SendSilenceWhenIdle field to given value.

### HasSendSilenceWhenIdle

`func (o *CallRequest) HasSendSilenceWhenIdle() bool`

HasSendSilenceWhenIdle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CallRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CallRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CallRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CallRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlMethod

`func (o *CallRequest) GetWebhookUrlMethod() string`

GetWebhookUrlMethod returns the WebhookUrlMethod field if non-nil, zero value otherwise.

### GetWebhookUrlMethodOk

`func (o *CallRequest) GetWebhookUrlMethodOk() (*string, bool)`

GetWebhookUrlMethodOk returns a tuple with the WebhookUrlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlMethod

`func (o *CallRequest) SetWebhookUrlMethod(v string)`

SetWebhookUrlMethod sets WebhookUrlMethod field to given value.

### HasWebhookUrlMethod

`func (o *CallRequest) HasWebhookUrlMethod() bool`

HasWebhookUrlMethod returns a boolean if a field has been set.

### GetRecord

`func (o *CallRequest) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *CallRequest) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *CallRequest) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *CallRequest) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRecordChannels

`func (o *CallRequest) GetRecordChannels() string`

GetRecordChannels returns the RecordChannels field if non-nil, zero value otherwise.

### GetRecordChannelsOk

`func (o *CallRequest) GetRecordChannelsOk() (*string, bool)`

GetRecordChannelsOk returns a tuple with the RecordChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordChannels

`func (o *CallRequest) SetRecordChannels(v string)`

SetRecordChannels sets RecordChannels field to given value.

### HasRecordChannels

`func (o *CallRequest) HasRecordChannels() bool`

HasRecordChannels returns a boolean if a field has been set.

### GetRecordFormat

`func (o *CallRequest) GetRecordFormat() string`

GetRecordFormat returns the RecordFormat field if non-nil, zero value otherwise.

### GetRecordFormatOk

`func (o *CallRequest) GetRecordFormatOk() (*string, bool)`

GetRecordFormatOk returns a tuple with the RecordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFormat

`func (o *CallRequest) SetRecordFormat(v string)`

SetRecordFormat sets RecordFormat field to given value.

### HasRecordFormat

`func (o *CallRequest) HasRecordFormat() bool`

HasRecordFormat returns a boolean if a field has been set.

### GetRecordMaxLength

`func (o *CallRequest) GetRecordMaxLength() int32`

GetRecordMaxLength returns the RecordMaxLength field if non-nil, zero value otherwise.

### GetRecordMaxLengthOk

`func (o *CallRequest) GetRecordMaxLengthOk() (*int32, bool)`

GetRecordMaxLengthOk returns a tuple with the RecordMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordMaxLength

`func (o *CallRequest) SetRecordMaxLength(v int32)`

SetRecordMaxLength sets RecordMaxLength field to given value.

### HasRecordMaxLength

`func (o *CallRequest) HasRecordMaxLength() bool`

HasRecordMaxLength returns a boolean if a field has been set.

### GetRecordTimeoutSecs

`func (o *CallRequest) GetRecordTimeoutSecs() int32`

GetRecordTimeoutSecs returns the RecordTimeoutSecs field if non-nil, zero value otherwise.

### GetRecordTimeoutSecsOk

`func (o *CallRequest) GetRecordTimeoutSecsOk() (*int32, bool)`

GetRecordTimeoutSecsOk returns a tuple with the RecordTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTimeoutSecs

`func (o *CallRequest) SetRecordTimeoutSecs(v int32)`

SetRecordTimeoutSecs sets RecordTimeoutSecs field to given value.

### HasRecordTimeoutSecs

`func (o *CallRequest) HasRecordTimeoutSecs() bool`

HasRecordTimeoutSecs returns a boolean if a field has been set.

### GetRecordTrack

`func (o *CallRequest) GetRecordTrack() string`

GetRecordTrack returns the RecordTrack field if non-nil, zero value otherwise.

### GetRecordTrackOk

`func (o *CallRequest) GetRecordTrackOk() (*string, bool)`

GetRecordTrackOk returns a tuple with the RecordTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTrack

`func (o *CallRequest) SetRecordTrack(v string)`

SetRecordTrack sets RecordTrack field to given value.

### HasRecordTrack

`func (o *CallRequest) HasRecordTrack() bool`

HasRecordTrack returns a boolean if a field has been set.

### GetRecordTrim

`func (o *CallRequest) GetRecordTrim() string`

GetRecordTrim returns the RecordTrim field if non-nil, zero value otherwise.

### GetRecordTrimOk

`func (o *CallRequest) GetRecordTrimOk() (*string, bool)`

GetRecordTrimOk returns a tuple with the RecordTrim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTrim

`func (o *CallRequest) SetRecordTrim(v string)`

SetRecordTrim sets RecordTrim field to given value.

### HasRecordTrim

`func (o *CallRequest) HasRecordTrim() bool`

HasRecordTrim returns a boolean if a field has been set.

### GetRecordCustomFileName

`func (o *CallRequest) GetRecordCustomFileName() string`

GetRecordCustomFileName returns the RecordCustomFileName field if non-nil, zero value otherwise.

### GetRecordCustomFileNameOk

`func (o *CallRequest) GetRecordCustomFileNameOk() (*string, bool)`

GetRecordCustomFileNameOk returns a tuple with the RecordCustomFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCustomFileName

`func (o *CallRequest) SetRecordCustomFileName(v string)`

SetRecordCustomFileName sets RecordCustomFileName field to given value.

### HasRecordCustomFileName

`func (o *CallRequest) HasRecordCustomFileName() bool`

HasRecordCustomFileName returns a boolean if a field has been set.

### GetSuperviseCallControlId

`func (o *CallRequest) GetSuperviseCallControlId() string`

GetSuperviseCallControlId returns the SuperviseCallControlId field if non-nil, zero value otherwise.

### GetSuperviseCallControlIdOk

`func (o *CallRequest) GetSuperviseCallControlIdOk() (*string, bool)`

GetSuperviseCallControlIdOk returns a tuple with the SuperviseCallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperviseCallControlId

`func (o *CallRequest) SetSuperviseCallControlId(v string)`

SetSuperviseCallControlId sets SuperviseCallControlId field to given value.

### HasSuperviseCallControlId

`func (o *CallRequest) HasSuperviseCallControlId() bool`

HasSuperviseCallControlId returns a boolean if a field has been set.

### GetSupervisorRole

`func (o *CallRequest) GetSupervisorRole() string`

GetSupervisorRole returns the SupervisorRole field if non-nil, zero value otherwise.

### GetSupervisorRoleOk

`func (o *CallRequest) GetSupervisorRoleOk() (*string, bool)`

GetSupervisorRoleOk returns a tuple with the SupervisorRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisorRole

`func (o *CallRequest) SetSupervisorRole(v string)`

SetSupervisorRole sets SupervisorRole field to given value.

### HasSupervisorRole

`func (o *CallRequest) HasSupervisorRole() bool`

HasSupervisorRole returns a boolean if a field has been set.

### GetEnableDialogflow

`func (o *CallRequest) GetEnableDialogflow() bool`

GetEnableDialogflow returns the EnableDialogflow field if non-nil, zero value otherwise.

### GetEnableDialogflowOk

`func (o *CallRequest) GetEnableDialogflowOk() (*bool, bool)`

GetEnableDialogflowOk returns a tuple with the EnableDialogflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDialogflow

`func (o *CallRequest) SetEnableDialogflow(v bool)`

SetEnableDialogflow sets EnableDialogflow field to given value.

### HasEnableDialogflow

`func (o *CallRequest) HasEnableDialogflow() bool`

HasEnableDialogflow returns a boolean if a field has been set.

### GetDialogflowConfig

`func (o *CallRequest) GetDialogflowConfig() DialogflowConfig`

GetDialogflowConfig returns the DialogflowConfig field if non-nil, zero value otherwise.

### GetDialogflowConfigOk

`func (o *CallRequest) GetDialogflowConfigOk() (*DialogflowConfig, bool)`

GetDialogflowConfigOk returns a tuple with the DialogflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialogflowConfig

`func (o *CallRequest) SetDialogflowConfig(v DialogflowConfig)`

SetDialogflowConfig sets DialogflowConfig field to given value.

### HasDialogflowConfig

`func (o *CallRequest) HasDialogflowConfig() bool`

HasDialogflowConfig returns a boolean if a field has been set.

### GetTranscription

`func (o *CallRequest) GetTranscription() bool`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *CallRequest) GetTranscriptionOk() (*bool, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *CallRequest) SetTranscription(v bool)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *CallRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetTranscriptionConfig

`func (o *CallRequest) GetTranscriptionConfig() TranscriptionStartRequest`

GetTranscriptionConfig returns the TranscriptionConfig field if non-nil, zero value otherwise.

### GetTranscriptionConfigOk

`func (o *CallRequest) GetTranscriptionConfigOk() (*TranscriptionStartRequest, bool)`

GetTranscriptionConfigOk returns a tuple with the TranscriptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionConfig

`func (o *CallRequest) SetTranscriptionConfig(v TranscriptionStartRequest)`

SetTranscriptionConfig sets TranscriptionConfig field to given value.

### HasTranscriptionConfig

`func (o *CallRequest) HasTranscriptionConfig() bool`

HasTranscriptionConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


