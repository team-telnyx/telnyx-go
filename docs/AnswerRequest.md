# AnswerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingGroupId** | Pointer to **string** | Use this field to set the Billing Group ID for the call. Must be a valid and existing Billing Group ID. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**CustomHeaders** | Pointer to [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers to be added to the SIP INVITE response. | [optional] 
**PreferredCodecs** | Pointer to **string** | The list of comma-separated codecs in a preferred order for the forked media to be received. | [optional] 
**SipHeaders** | Pointer to [**[]SipHeader**](SipHeader.md) | SIP headers to be added to the SIP INVITE response. Currently only User-to-User header is supported. | [optional] 
**SoundModifications** | Pointer to [**SoundModifications**](SoundModifications.md) |  | [optional] 
**StreamUrl** | Pointer to **string** | The destination WebSocket address where the stream is going to be delivered. | [optional] 
**StreamTrack** | Pointer to **string** | Specifies which track should be streamed. | [optional] [default to "inbound_track"]
**StreamBidirectionalMode** | Pointer to [**StreamBidirectionalMode**](StreamBidirectionalMode.md) |  | [optional] [default to MP3]
**StreamBidirectionalCodec** | Pointer to [**StreamBidirectionalCodec**](StreamBidirectionalCodec.md) |  | [optional] [default to PCMU]
**StreamBidirectionalTargetLegs** | Pointer to [**StreamBidirectionalTargetLegs**](StreamBidirectionalTargetLegs.md) |  | [optional] [default to OPPOSITE]
**SendSilenceWhenIdle** | Pointer to **bool** | Generate silence RTP packets when no transmission available. | [optional] [default to false]
**WebhookUrl** | Pointer to **string** | Use this field to override the URL for which Telnyx will send subsequent webhooks to for this call. | [optional] 
**WebhookUrlMethod** | Pointer to **string** | HTTP request type used for &#x60;webhook_url&#x60;. | [optional] [default to "POST"]
**Transcription** | Pointer to **bool** | Enable transcription upon call answer. The default value is false. | [optional] [default to false]
**TranscriptionConfig** | Pointer to [**TranscriptionStartRequest**](TranscriptionStartRequest.md) |  | [optional] 
**Record** | Pointer to **string** | Start recording automatically after an event. Disabled by default. | [optional] 
**RecordChannels** | Pointer to **string** | Defines which channel should be recorded (&#39;single&#39; or &#39;dual&#39;) when &#x60;record&#x60; is specified. | [optional] [default to "dual"]
**RecordFormat** | Pointer to **string** | Defines the format of the recording (&#39;wav&#39; or &#39;mp3&#39;) when &#x60;record&#x60; is specified. | [optional] [default to "mp3"]
**RecordMaxLength** | Pointer to **int32** | Defines the maximum length for the recording in seconds when &#x60;record&#x60; is specified. The minimum value is 0. The maximum value is 43200. The default value is 0 (infinite). | [optional] [default to 0]
**RecordTimeoutSecs** | Pointer to **int32** | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected when &#x60;record&#x60; is specified. The timer only starts when the speech is detected. Please note that call transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite). | [optional] [default to 0]
**RecordTrack** | Pointer to **string** | The audio track to be recorded. Can be either &#x60;both&#x60;, &#x60;inbound&#x60; or &#x60;outbound&#x60;. If only single track is specified (&#x60;inbound&#x60;, &#x60;outbound&#x60;), &#x60;channels&#x60; configuration is ignored and it will be recorded as mono (single channel). | [optional] [default to "both"]
**RecordTrim** | Pointer to **string** | When set to &#x60;trim-silence&#x60;, silence will be removed from the beginning and end of the recording. | [optional] 
**RecordCustomFileName** | Pointer to **string** | The custom recording file name to be used instead of the default &#x60;call_leg_id&#x60;. Telnyx will still add a Unix timestamp suffix. | [optional] 

## Methods

### NewAnswerRequest

`func NewAnswerRequest() *AnswerRequest`

NewAnswerRequest instantiates a new AnswerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerRequestWithDefaults

`func NewAnswerRequestWithDefaults() *AnswerRequest`

NewAnswerRequestWithDefaults instantiates a new AnswerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingGroupId

`func (o *AnswerRequest) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *AnswerRequest) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *AnswerRequest) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *AnswerRequest) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetClientState

`func (o *AnswerRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *AnswerRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *AnswerRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *AnswerRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *AnswerRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *AnswerRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *AnswerRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *AnswerRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *AnswerRequest) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *AnswerRequest) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *AnswerRequest) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *AnswerRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetPreferredCodecs

`func (o *AnswerRequest) GetPreferredCodecs() string`

GetPreferredCodecs returns the PreferredCodecs field if non-nil, zero value otherwise.

### GetPreferredCodecsOk

`func (o *AnswerRequest) GetPreferredCodecsOk() (*string, bool)`

GetPreferredCodecsOk returns a tuple with the PreferredCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredCodecs

`func (o *AnswerRequest) SetPreferredCodecs(v string)`

SetPreferredCodecs sets PreferredCodecs field to given value.

### HasPreferredCodecs

`func (o *AnswerRequest) HasPreferredCodecs() bool`

HasPreferredCodecs returns a boolean if a field has been set.

### GetSipHeaders

`func (o *AnswerRequest) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *AnswerRequest) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *AnswerRequest) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *AnswerRequest) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.

### GetSoundModifications

`func (o *AnswerRequest) GetSoundModifications() SoundModifications`

GetSoundModifications returns the SoundModifications field if non-nil, zero value otherwise.

### GetSoundModificationsOk

`func (o *AnswerRequest) GetSoundModificationsOk() (*SoundModifications, bool)`

GetSoundModificationsOk returns a tuple with the SoundModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoundModifications

`func (o *AnswerRequest) SetSoundModifications(v SoundModifications)`

SetSoundModifications sets SoundModifications field to given value.

### HasSoundModifications

`func (o *AnswerRequest) HasSoundModifications() bool`

HasSoundModifications returns a boolean if a field has been set.

### GetStreamUrl

`func (o *AnswerRequest) GetStreamUrl() string`

GetStreamUrl returns the StreamUrl field if non-nil, zero value otherwise.

### GetStreamUrlOk

`func (o *AnswerRequest) GetStreamUrlOk() (*string, bool)`

GetStreamUrlOk returns a tuple with the StreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamUrl

`func (o *AnswerRequest) SetStreamUrl(v string)`

SetStreamUrl sets StreamUrl field to given value.

### HasStreamUrl

`func (o *AnswerRequest) HasStreamUrl() bool`

HasStreamUrl returns a boolean if a field has been set.

### GetStreamTrack

`func (o *AnswerRequest) GetStreamTrack() string`

GetStreamTrack returns the StreamTrack field if non-nil, zero value otherwise.

### GetStreamTrackOk

`func (o *AnswerRequest) GetStreamTrackOk() (*string, bool)`

GetStreamTrackOk returns a tuple with the StreamTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamTrack

`func (o *AnswerRequest) SetStreamTrack(v string)`

SetStreamTrack sets StreamTrack field to given value.

### HasStreamTrack

`func (o *AnswerRequest) HasStreamTrack() bool`

HasStreamTrack returns a boolean if a field has been set.

### GetStreamBidirectionalMode

`func (o *AnswerRequest) GetStreamBidirectionalMode() StreamBidirectionalMode`

GetStreamBidirectionalMode returns the StreamBidirectionalMode field if non-nil, zero value otherwise.

### GetStreamBidirectionalModeOk

`func (o *AnswerRequest) GetStreamBidirectionalModeOk() (*StreamBidirectionalMode, bool)`

GetStreamBidirectionalModeOk returns a tuple with the StreamBidirectionalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalMode

`func (o *AnswerRequest) SetStreamBidirectionalMode(v StreamBidirectionalMode)`

SetStreamBidirectionalMode sets StreamBidirectionalMode field to given value.

### HasStreamBidirectionalMode

`func (o *AnswerRequest) HasStreamBidirectionalMode() bool`

HasStreamBidirectionalMode returns a boolean if a field has been set.

### GetStreamBidirectionalCodec

`func (o *AnswerRequest) GetStreamBidirectionalCodec() StreamBidirectionalCodec`

GetStreamBidirectionalCodec returns the StreamBidirectionalCodec field if non-nil, zero value otherwise.

### GetStreamBidirectionalCodecOk

`func (o *AnswerRequest) GetStreamBidirectionalCodecOk() (*StreamBidirectionalCodec, bool)`

GetStreamBidirectionalCodecOk returns a tuple with the StreamBidirectionalCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalCodec

`func (o *AnswerRequest) SetStreamBidirectionalCodec(v StreamBidirectionalCodec)`

SetStreamBidirectionalCodec sets StreamBidirectionalCodec field to given value.

### HasStreamBidirectionalCodec

`func (o *AnswerRequest) HasStreamBidirectionalCodec() bool`

HasStreamBidirectionalCodec returns a boolean if a field has been set.

### GetStreamBidirectionalTargetLegs

`func (o *AnswerRequest) GetStreamBidirectionalTargetLegs() StreamBidirectionalTargetLegs`

GetStreamBidirectionalTargetLegs returns the StreamBidirectionalTargetLegs field if non-nil, zero value otherwise.

### GetStreamBidirectionalTargetLegsOk

`func (o *AnswerRequest) GetStreamBidirectionalTargetLegsOk() (*StreamBidirectionalTargetLegs, bool)`

GetStreamBidirectionalTargetLegsOk returns a tuple with the StreamBidirectionalTargetLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalTargetLegs

`func (o *AnswerRequest) SetStreamBidirectionalTargetLegs(v StreamBidirectionalTargetLegs)`

SetStreamBidirectionalTargetLegs sets StreamBidirectionalTargetLegs field to given value.

### HasStreamBidirectionalTargetLegs

`func (o *AnswerRequest) HasStreamBidirectionalTargetLegs() bool`

HasStreamBidirectionalTargetLegs returns a boolean if a field has been set.

### GetSendSilenceWhenIdle

`func (o *AnswerRequest) GetSendSilenceWhenIdle() bool`

GetSendSilenceWhenIdle returns the SendSilenceWhenIdle field if non-nil, zero value otherwise.

### GetSendSilenceWhenIdleOk

`func (o *AnswerRequest) GetSendSilenceWhenIdleOk() (*bool, bool)`

GetSendSilenceWhenIdleOk returns a tuple with the SendSilenceWhenIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendSilenceWhenIdle

`func (o *AnswerRequest) SetSendSilenceWhenIdle(v bool)`

SetSendSilenceWhenIdle sets SendSilenceWhenIdle field to given value.

### HasSendSilenceWhenIdle

`func (o *AnswerRequest) HasSendSilenceWhenIdle() bool`

HasSendSilenceWhenIdle returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *AnswerRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AnswerRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AnswerRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *AnswerRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetWebhookUrlMethod

`func (o *AnswerRequest) GetWebhookUrlMethod() string`

GetWebhookUrlMethod returns the WebhookUrlMethod field if non-nil, zero value otherwise.

### GetWebhookUrlMethodOk

`func (o *AnswerRequest) GetWebhookUrlMethodOk() (*string, bool)`

GetWebhookUrlMethodOk returns a tuple with the WebhookUrlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrlMethod

`func (o *AnswerRequest) SetWebhookUrlMethod(v string)`

SetWebhookUrlMethod sets WebhookUrlMethod field to given value.

### HasWebhookUrlMethod

`func (o *AnswerRequest) HasWebhookUrlMethod() bool`

HasWebhookUrlMethod returns a boolean if a field has been set.

### GetTranscription

`func (o *AnswerRequest) GetTranscription() bool`

GetTranscription returns the Transcription field if non-nil, zero value otherwise.

### GetTranscriptionOk

`func (o *AnswerRequest) GetTranscriptionOk() (*bool, bool)`

GetTranscriptionOk returns a tuple with the Transcription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscription

`func (o *AnswerRequest) SetTranscription(v bool)`

SetTranscription sets Transcription field to given value.

### HasTranscription

`func (o *AnswerRequest) HasTranscription() bool`

HasTranscription returns a boolean if a field has been set.

### GetTranscriptionConfig

`func (o *AnswerRequest) GetTranscriptionConfig() TranscriptionStartRequest`

GetTranscriptionConfig returns the TranscriptionConfig field if non-nil, zero value otherwise.

### GetTranscriptionConfigOk

`func (o *AnswerRequest) GetTranscriptionConfigOk() (*TranscriptionStartRequest, bool)`

GetTranscriptionConfigOk returns a tuple with the TranscriptionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionConfig

`func (o *AnswerRequest) SetTranscriptionConfig(v TranscriptionStartRequest)`

SetTranscriptionConfig sets TranscriptionConfig field to given value.

### HasTranscriptionConfig

`func (o *AnswerRequest) HasTranscriptionConfig() bool`

HasTranscriptionConfig returns a boolean if a field has been set.

### GetRecord

`func (o *AnswerRequest) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *AnswerRequest) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *AnswerRequest) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *AnswerRequest) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRecordChannels

`func (o *AnswerRequest) GetRecordChannels() string`

GetRecordChannels returns the RecordChannels field if non-nil, zero value otherwise.

### GetRecordChannelsOk

`func (o *AnswerRequest) GetRecordChannelsOk() (*string, bool)`

GetRecordChannelsOk returns a tuple with the RecordChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordChannels

`func (o *AnswerRequest) SetRecordChannels(v string)`

SetRecordChannels sets RecordChannels field to given value.

### HasRecordChannels

`func (o *AnswerRequest) HasRecordChannels() bool`

HasRecordChannels returns a boolean if a field has been set.

### GetRecordFormat

`func (o *AnswerRequest) GetRecordFormat() string`

GetRecordFormat returns the RecordFormat field if non-nil, zero value otherwise.

### GetRecordFormatOk

`func (o *AnswerRequest) GetRecordFormatOk() (*string, bool)`

GetRecordFormatOk returns a tuple with the RecordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFormat

`func (o *AnswerRequest) SetRecordFormat(v string)`

SetRecordFormat sets RecordFormat field to given value.

### HasRecordFormat

`func (o *AnswerRequest) HasRecordFormat() bool`

HasRecordFormat returns a boolean if a field has been set.

### GetRecordMaxLength

`func (o *AnswerRequest) GetRecordMaxLength() int32`

GetRecordMaxLength returns the RecordMaxLength field if non-nil, zero value otherwise.

### GetRecordMaxLengthOk

`func (o *AnswerRequest) GetRecordMaxLengthOk() (*int32, bool)`

GetRecordMaxLengthOk returns a tuple with the RecordMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordMaxLength

`func (o *AnswerRequest) SetRecordMaxLength(v int32)`

SetRecordMaxLength sets RecordMaxLength field to given value.

### HasRecordMaxLength

`func (o *AnswerRequest) HasRecordMaxLength() bool`

HasRecordMaxLength returns a boolean if a field has been set.

### GetRecordTimeoutSecs

`func (o *AnswerRequest) GetRecordTimeoutSecs() int32`

GetRecordTimeoutSecs returns the RecordTimeoutSecs field if non-nil, zero value otherwise.

### GetRecordTimeoutSecsOk

`func (o *AnswerRequest) GetRecordTimeoutSecsOk() (*int32, bool)`

GetRecordTimeoutSecsOk returns a tuple with the RecordTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTimeoutSecs

`func (o *AnswerRequest) SetRecordTimeoutSecs(v int32)`

SetRecordTimeoutSecs sets RecordTimeoutSecs field to given value.

### HasRecordTimeoutSecs

`func (o *AnswerRequest) HasRecordTimeoutSecs() bool`

HasRecordTimeoutSecs returns a boolean if a field has been set.

### GetRecordTrack

`func (o *AnswerRequest) GetRecordTrack() string`

GetRecordTrack returns the RecordTrack field if non-nil, zero value otherwise.

### GetRecordTrackOk

`func (o *AnswerRequest) GetRecordTrackOk() (*string, bool)`

GetRecordTrackOk returns a tuple with the RecordTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTrack

`func (o *AnswerRequest) SetRecordTrack(v string)`

SetRecordTrack sets RecordTrack field to given value.

### HasRecordTrack

`func (o *AnswerRequest) HasRecordTrack() bool`

HasRecordTrack returns a boolean if a field has been set.

### GetRecordTrim

`func (o *AnswerRequest) GetRecordTrim() string`

GetRecordTrim returns the RecordTrim field if non-nil, zero value otherwise.

### GetRecordTrimOk

`func (o *AnswerRequest) GetRecordTrimOk() (*string, bool)`

GetRecordTrimOk returns a tuple with the RecordTrim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTrim

`func (o *AnswerRequest) SetRecordTrim(v string)`

SetRecordTrim sets RecordTrim field to given value.

### HasRecordTrim

`func (o *AnswerRequest) HasRecordTrim() bool`

HasRecordTrim returns a boolean if a field has been set.

### GetRecordCustomFileName

`func (o *AnswerRequest) GetRecordCustomFileName() string`

GetRecordCustomFileName returns the RecordCustomFileName field if non-nil, zero value otherwise.

### GetRecordCustomFileNameOk

`func (o *AnswerRequest) GetRecordCustomFileNameOk() (*string, bool)`

GetRecordCustomFileNameOk returns a tuple with the RecordCustomFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCustomFileName

`func (o *AnswerRequest) SetRecordCustomFileName(v string)`

SetRecordCustomFileName sets RecordCustomFileName field to given value.

### HasRecordCustomFileName

`func (o *AnswerRequest) HasRecordCustomFileName() bool`

HasRecordCustomFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


