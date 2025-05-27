# DeprecatedInitiateCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The phone number of the called party. Phone numbers are formatted with a &#x60;+&#x60; and country code. | 
**From** | **string** | The phone number of the party that initiated the call. Phone numbers are formatted with a &#x60;+&#x60; and country code. | 
**CallerId** | Pointer to **string** | To be used as the caller id name (SIP From Display Name) presented to the destination (&#x60;To&#x60; number). The string should have a maximum of 128 characters, containing only letters, numbers, spaces, and &#x60;-_~!.+&#x60; special characters. If ommited, the display name will be the same as the number in the &#x60;From&#x60; field. | [optional] 
**Url** | Pointer to **string** | The URL from which Telnyx will retrieve the TeXML call instructions. | [optional] 
**UrlMethod** | Pointer to **string** | HTTP request type used for &#x60;Url&#x60;. The default value is inherited from TeXML Application setting. | [optional] [default to "POST"]
**FallbackUrl** | Pointer to **string** | A failover URL for which Telnyx will retrieve the TeXML call instructions if the &#x60;Url&#x60; is not responding. | [optional] 
**StatusCallback** | Pointer to **string** | URL destination for Telnyx to send status callback events to for the call. | [optional] 
**StatusCallbackMethod** | Pointer to **string** | HTTP request type used for &#x60;StatusCallback&#x60;. | [optional] [default to "POST"]
**StatusCallbackEvent** | Pointer to **string** | The call events for which Telnyx should send a webhook. Multiple events can be defined when separated by a space. | [optional] [default to "completed"]
**MachineDetection** | Pointer to **string** | Enables Answering Machine Detection. | [optional] [default to "Disable"]
**DetectionMode** | Pointer to **string** | Allows you to chose between Premium and Standard detections. | [optional] [default to "Regular"]
**AsyncAmd** | Pointer to **bool** | Select whether to perform answering machine detection in the background. By default execution is blocked until Answering Machine Detection is completed. | [optional] [default to false]
**AsyncAmdStatusCallback** | Pointer to **string** | URL destination for Telnyx to send AMD callback events to for the call. | [optional] 
**AsyncAmdStatusCallbackMethod** | Pointer to **string** | HTTP request type used for &#x60;AsyncAmdStatusCallback&#x60;. The default value is inherited from TeXML Application setting. | [optional] [default to "POST"]
**MachineDetectionTimeout** | Pointer to **int32** | Maximum timeout threshold in milliseconds for overall detection. | [optional] [default to 30000]
**MachineDetectionSpeechThreshold** | Pointer to **int32** | Maximum threshold of a human greeting. If greeting longer than this value, considered machine. Ignored when &#x60;premium&#x60; detection is used. | [optional] [default to 3500]
**MachineDetectionSpeechEndThreshold** | Pointer to **int32** | Silence duration threshold after a greeting message or voice for it be considered human. Ignored when &#x60;premium&#x60; detection is used. | [optional] [default to 800]
**MachineDetectionSilenceTimeout** | Pointer to **int32** | If initial silence duration is greater than this value, consider it a machine. Ignored when &#x60;premium&#x60; detection is used. | [optional] [default to 3500]
**CancelPlaybackOnMachineDetection** | Pointer to **bool** | Whether to cancel ongoing playback on &#x60;machine&#x60; detection. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**CancelPlaybackOnDetectMessageEnd** | Pointer to **bool** | Whether to cancel ongoing playback on &#x60;greeting ended&#x60; detection. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**PreferredCodecs** | Pointer to **string** | The list of comma-separated codecs to be offered on a call. | [optional] 
**Record** | Pointer to **bool** | Whether to record the entire participant&#39;s call leg. Defaults to &#x60;false&#x60;. | [optional] 
**RecordingChannels** | Pointer to **string** | The number of channels in the final recording. Defaults to &#x60;mono&#x60;. | [optional] 
**RecordingStatusCallback** | Pointer to **string** | The URL the recording callbacks will be sent to. | [optional] 
**RecordingStatusCallbackMethod** | Pointer to **string** | HTTP request type used for &#x60;RecordingStatusCallback&#x60;. Defaults to &#x60;POST&#x60;. | [optional] 
**RecordingStatusCallbackEvent** | Pointer to **string** | The changes to the recording&#39;s state that should generate a call to &#x60;RecoridngStatusCallback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Separate multiple values with a space. Defaults to &#x60;completed&#x60;. | [optional] 
**RecordingTimeout** | Pointer to **int32** | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected. The timer only starts when the speech is detected. Please note that the transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite) | [optional] [default to 0]
**RecordingTrack** | Pointer to **string** | The audio track to record for the call. The default is &#x60;both&#x60;. | [optional] 
**SipAuthPassword** | Pointer to **string** | The password to use for SIP authentication. | [optional] 
**SipAuthUsername** | Pointer to **string** | The username to use for SIP authentication. | [optional] 
**Trim** | Pointer to **string** | Whether to trim any leading and trailing silence from the recording. Defaults to &#x60;trim-silence&#x60;. | [optional] 

## Methods

### NewDeprecatedInitiateCallRequest

`func NewDeprecatedInitiateCallRequest(to string, from string, ) *DeprecatedInitiateCallRequest`

NewDeprecatedInitiateCallRequest instantiates a new DeprecatedInitiateCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedInitiateCallRequestWithDefaults

`func NewDeprecatedInitiateCallRequestWithDefaults() *DeprecatedInitiateCallRequest`

NewDeprecatedInitiateCallRequestWithDefaults instantiates a new DeprecatedInitiateCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *DeprecatedInitiateCallRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DeprecatedInitiateCallRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DeprecatedInitiateCallRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *DeprecatedInitiateCallRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DeprecatedInitiateCallRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DeprecatedInitiateCallRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetCallerId

`func (o *DeprecatedInitiateCallRequest) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *DeprecatedInitiateCallRequest) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *DeprecatedInitiateCallRequest) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *DeprecatedInitiateCallRequest) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetUrl

`func (o *DeprecatedInitiateCallRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeprecatedInitiateCallRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeprecatedInitiateCallRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeprecatedInitiateCallRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlMethod

`func (o *DeprecatedInitiateCallRequest) GetUrlMethod() string`

GetUrlMethod returns the UrlMethod field if non-nil, zero value otherwise.

### GetUrlMethodOk

`func (o *DeprecatedInitiateCallRequest) GetUrlMethodOk() (*string, bool)`

GetUrlMethodOk returns a tuple with the UrlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlMethod

`func (o *DeprecatedInitiateCallRequest) SetUrlMethod(v string)`

SetUrlMethod sets UrlMethod field to given value.

### HasUrlMethod

`func (o *DeprecatedInitiateCallRequest) HasUrlMethod() bool`

HasUrlMethod returns a boolean if a field has been set.

### GetFallbackUrl

`func (o *DeprecatedInitiateCallRequest) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *DeprecatedInitiateCallRequest) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *DeprecatedInitiateCallRequest) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *DeprecatedInitiateCallRequest) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### GetStatusCallback

`func (o *DeprecatedInitiateCallRequest) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *DeprecatedInitiateCallRequest) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *DeprecatedInitiateCallRequest) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *DeprecatedInitiateCallRequest) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### GetStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *DeprecatedInitiateCallRequest) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### GetStatusCallbackEvent

`func (o *DeprecatedInitiateCallRequest) GetStatusCallbackEvent() string`

GetStatusCallbackEvent returns the StatusCallbackEvent field if non-nil, zero value otherwise.

### GetStatusCallbackEventOk

`func (o *DeprecatedInitiateCallRequest) GetStatusCallbackEventOk() (*string, bool)`

GetStatusCallbackEventOk returns a tuple with the StatusCallbackEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackEvent

`func (o *DeprecatedInitiateCallRequest) SetStatusCallbackEvent(v string)`

SetStatusCallbackEvent sets StatusCallbackEvent field to given value.

### HasStatusCallbackEvent

`func (o *DeprecatedInitiateCallRequest) HasStatusCallbackEvent() bool`

HasStatusCallbackEvent returns a boolean if a field has been set.

### GetMachineDetection

`func (o *DeprecatedInitiateCallRequest) GetMachineDetection() string`

GetMachineDetection returns the MachineDetection field if non-nil, zero value otherwise.

### GetMachineDetectionOk

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionOk() (*string, bool)`

GetMachineDetectionOk returns a tuple with the MachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetection

`func (o *DeprecatedInitiateCallRequest) SetMachineDetection(v string)`

SetMachineDetection sets MachineDetection field to given value.

### HasMachineDetection

`func (o *DeprecatedInitiateCallRequest) HasMachineDetection() bool`

HasMachineDetection returns a boolean if a field has been set.

### GetDetectionMode

`func (o *DeprecatedInitiateCallRequest) GetDetectionMode() string`

GetDetectionMode returns the DetectionMode field if non-nil, zero value otherwise.

### GetDetectionModeOk

`func (o *DeprecatedInitiateCallRequest) GetDetectionModeOk() (*string, bool)`

GetDetectionModeOk returns a tuple with the DetectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMode

`func (o *DeprecatedInitiateCallRequest) SetDetectionMode(v string)`

SetDetectionMode sets DetectionMode field to given value.

### HasDetectionMode

`func (o *DeprecatedInitiateCallRequest) HasDetectionMode() bool`

HasDetectionMode returns a boolean if a field has been set.

### GetAsyncAmd

`func (o *DeprecatedInitiateCallRequest) GetAsyncAmd() bool`

GetAsyncAmd returns the AsyncAmd field if non-nil, zero value otherwise.

### GetAsyncAmdOk

`func (o *DeprecatedInitiateCallRequest) GetAsyncAmdOk() (*bool, bool)`

GetAsyncAmdOk returns a tuple with the AsyncAmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncAmd

`func (o *DeprecatedInitiateCallRequest) SetAsyncAmd(v bool)`

SetAsyncAmd sets AsyncAmd field to given value.

### HasAsyncAmd

`func (o *DeprecatedInitiateCallRequest) HasAsyncAmd() bool`

HasAsyncAmd returns a boolean if a field has been set.

### GetAsyncAmdStatusCallback

`func (o *DeprecatedInitiateCallRequest) GetAsyncAmdStatusCallback() string`

GetAsyncAmdStatusCallback returns the AsyncAmdStatusCallback field if non-nil, zero value otherwise.

### GetAsyncAmdStatusCallbackOk

`func (o *DeprecatedInitiateCallRequest) GetAsyncAmdStatusCallbackOk() (*string, bool)`

GetAsyncAmdStatusCallbackOk returns a tuple with the AsyncAmdStatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncAmdStatusCallback

`func (o *DeprecatedInitiateCallRequest) SetAsyncAmdStatusCallback(v string)`

SetAsyncAmdStatusCallback sets AsyncAmdStatusCallback field to given value.

### HasAsyncAmdStatusCallback

`func (o *DeprecatedInitiateCallRequest) HasAsyncAmdStatusCallback() bool`

HasAsyncAmdStatusCallback returns a boolean if a field has been set.

### GetAsyncAmdStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) GetAsyncAmdStatusCallbackMethod() string`

GetAsyncAmdStatusCallbackMethod returns the AsyncAmdStatusCallbackMethod field if non-nil, zero value otherwise.

### GetAsyncAmdStatusCallbackMethodOk

`func (o *DeprecatedInitiateCallRequest) GetAsyncAmdStatusCallbackMethodOk() (*string, bool)`

GetAsyncAmdStatusCallbackMethodOk returns a tuple with the AsyncAmdStatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncAmdStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) SetAsyncAmdStatusCallbackMethod(v string)`

SetAsyncAmdStatusCallbackMethod sets AsyncAmdStatusCallbackMethod field to given value.

### HasAsyncAmdStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) HasAsyncAmdStatusCallbackMethod() bool`

HasAsyncAmdStatusCallbackMethod returns a boolean if a field has been set.

### GetMachineDetectionTimeout

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionTimeout() int32`

GetMachineDetectionTimeout returns the MachineDetectionTimeout field if non-nil, zero value otherwise.

### GetMachineDetectionTimeoutOk

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionTimeoutOk() (*int32, bool)`

GetMachineDetectionTimeoutOk returns a tuple with the MachineDetectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionTimeout

`func (o *DeprecatedInitiateCallRequest) SetMachineDetectionTimeout(v int32)`

SetMachineDetectionTimeout sets MachineDetectionTimeout field to given value.

### HasMachineDetectionTimeout

`func (o *DeprecatedInitiateCallRequest) HasMachineDetectionTimeout() bool`

HasMachineDetectionTimeout returns a boolean if a field has been set.

### GetMachineDetectionSpeechThreshold

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionSpeechThreshold() int32`

GetMachineDetectionSpeechThreshold returns the MachineDetectionSpeechThreshold field if non-nil, zero value otherwise.

### GetMachineDetectionSpeechThresholdOk

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionSpeechThresholdOk() (*int32, bool)`

GetMachineDetectionSpeechThresholdOk returns a tuple with the MachineDetectionSpeechThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionSpeechThreshold

`func (o *DeprecatedInitiateCallRequest) SetMachineDetectionSpeechThreshold(v int32)`

SetMachineDetectionSpeechThreshold sets MachineDetectionSpeechThreshold field to given value.

### HasMachineDetectionSpeechThreshold

`func (o *DeprecatedInitiateCallRequest) HasMachineDetectionSpeechThreshold() bool`

HasMachineDetectionSpeechThreshold returns a boolean if a field has been set.

### GetMachineDetectionSpeechEndThreshold

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionSpeechEndThreshold() int32`

GetMachineDetectionSpeechEndThreshold returns the MachineDetectionSpeechEndThreshold field if non-nil, zero value otherwise.

### GetMachineDetectionSpeechEndThresholdOk

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionSpeechEndThresholdOk() (*int32, bool)`

GetMachineDetectionSpeechEndThresholdOk returns a tuple with the MachineDetectionSpeechEndThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionSpeechEndThreshold

`func (o *DeprecatedInitiateCallRequest) SetMachineDetectionSpeechEndThreshold(v int32)`

SetMachineDetectionSpeechEndThreshold sets MachineDetectionSpeechEndThreshold field to given value.

### HasMachineDetectionSpeechEndThreshold

`func (o *DeprecatedInitiateCallRequest) HasMachineDetectionSpeechEndThreshold() bool`

HasMachineDetectionSpeechEndThreshold returns a boolean if a field has been set.

### GetMachineDetectionSilenceTimeout

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionSilenceTimeout() int32`

GetMachineDetectionSilenceTimeout returns the MachineDetectionSilenceTimeout field if non-nil, zero value otherwise.

### GetMachineDetectionSilenceTimeoutOk

`func (o *DeprecatedInitiateCallRequest) GetMachineDetectionSilenceTimeoutOk() (*int32, bool)`

GetMachineDetectionSilenceTimeoutOk returns a tuple with the MachineDetectionSilenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionSilenceTimeout

`func (o *DeprecatedInitiateCallRequest) SetMachineDetectionSilenceTimeout(v int32)`

SetMachineDetectionSilenceTimeout sets MachineDetectionSilenceTimeout field to given value.

### HasMachineDetectionSilenceTimeout

`func (o *DeprecatedInitiateCallRequest) HasMachineDetectionSilenceTimeout() bool`

HasMachineDetectionSilenceTimeout returns a boolean if a field has been set.

### GetCancelPlaybackOnMachineDetection

`func (o *DeprecatedInitiateCallRequest) GetCancelPlaybackOnMachineDetection() bool`

GetCancelPlaybackOnMachineDetection returns the CancelPlaybackOnMachineDetection field if non-nil, zero value otherwise.

### GetCancelPlaybackOnMachineDetectionOk

`func (o *DeprecatedInitiateCallRequest) GetCancelPlaybackOnMachineDetectionOk() (*bool, bool)`

GetCancelPlaybackOnMachineDetectionOk returns a tuple with the CancelPlaybackOnMachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelPlaybackOnMachineDetection

`func (o *DeprecatedInitiateCallRequest) SetCancelPlaybackOnMachineDetection(v bool)`

SetCancelPlaybackOnMachineDetection sets CancelPlaybackOnMachineDetection field to given value.

### HasCancelPlaybackOnMachineDetection

`func (o *DeprecatedInitiateCallRequest) HasCancelPlaybackOnMachineDetection() bool`

HasCancelPlaybackOnMachineDetection returns a boolean if a field has been set.

### GetCancelPlaybackOnDetectMessageEnd

`func (o *DeprecatedInitiateCallRequest) GetCancelPlaybackOnDetectMessageEnd() bool`

GetCancelPlaybackOnDetectMessageEnd returns the CancelPlaybackOnDetectMessageEnd field if non-nil, zero value otherwise.

### GetCancelPlaybackOnDetectMessageEndOk

`func (o *DeprecatedInitiateCallRequest) GetCancelPlaybackOnDetectMessageEndOk() (*bool, bool)`

GetCancelPlaybackOnDetectMessageEndOk returns a tuple with the CancelPlaybackOnDetectMessageEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelPlaybackOnDetectMessageEnd

`func (o *DeprecatedInitiateCallRequest) SetCancelPlaybackOnDetectMessageEnd(v bool)`

SetCancelPlaybackOnDetectMessageEnd sets CancelPlaybackOnDetectMessageEnd field to given value.

### HasCancelPlaybackOnDetectMessageEnd

`func (o *DeprecatedInitiateCallRequest) HasCancelPlaybackOnDetectMessageEnd() bool`

HasCancelPlaybackOnDetectMessageEnd returns a boolean if a field has been set.

### GetPreferredCodecs

`func (o *DeprecatedInitiateCallRequest) GetPreferredCodecs() string`

GetPreferredCodecs returns the PreferredCodecs field if non-nil, zero value otherwise.

### GetPreferredCodecsOk

`func (o *DeprecatedInitiateCallRequest) GetPreferredCodecsOk() (*string, bool)`

GetPreferredCodecsOk returns a tuple with the PreferredCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredCodecs

`func (o *DeprecatedInitiateCallRequest) SetPreferredCodecs(v string)`

SetPreferredCodecs sets PreferredCodecs field to given value.

### HasPreferredCodecs

`func (o *DeprecatedInitiateCallRequest) HasPreferredCodecs() bool`

HasPreferredCodecs returns a boolean if a field has been set.

### GetRecord

`func (o *DeprecatedInitiateCallRequest) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DeprecatedInitiateCallRequest) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DeprecatedInitiateCallRequest) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DeprecatedInitiateCallRequest) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRecordingChannels

`func (o *DeprecatedInitiateCallRequest) GetRecordingChannels() string`

GetRecordingChannels returns the RecordingChannels field if non-nil, zero value otherwise.

### GetRecordingChannelsOk

`func (o *DeprecatedInitiateCallRequest) GetRecordingChannelsOk() (*string, bool)`

GetRecordingChannelsOk returns a tuple with the RecordingChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingChannels

`func (o *DeprecatedInitiateCallRequest) SetRecordingChannels(v string)`

SetRecordingChannels sets RecordingChannels field to given value.

### HasRecordingChannels

`func (o *DeprecatedInitiateCallRequest) HasRecordingChannels() bool`

HasRecordingChannels returns a boolean if a field has been set.

### GetRecordingStatusCallback

`func (o *DeprecatedInitiateCallRequest) GetRecordingStatusCallback() string`

GetRecordingStatusCallback returns the RecordingStatusCallback field if non-nil, zero value otherwise.

### GetRecordingStatusCallbackOk

`func (o *DeprecatedInitiateCallRequest) GetRecordingStatusCallbackOk() (*string, bool)`

GetRecordingStatusCallbackOk returns a tuple with the RecordingStatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatusCallback

`func (o *DeprecatedInitiateCallRequest) SetRecordingStatusCallback(v string)`

SetRecordingStatusCallback sets RecordingStatusCallback field to given value.

### HasRecordingStatusCallback

`func (o *DeprecatedInitiateCallRequest) HasRecordingStatusCallback() bool`

HasRecordingStatusCallback returns a boolean if a field has been set.

### GetRecordingStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) GetRecordingStatusCallbackMethod() string`

GetRecordingStatusCallbackMethod returns the RecordingStatusCallbackMethod field if non-nil, zero value otherwise.

### GetRecordingStatusCallbackMethodOk

`func (o *DeprecatedInitiateCallRequest) GetRecordingStatusCallbackMethodOk() (*string, bool)`

GetRecordingStatusCallbackMethodOk returns a tuple with the RecordingStatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) SetRecordingStatusCallbackMethod(v string)`

SetRecordingStatusCallbackMethod sets RecordingStatusCallbackMethod field to given value.

### HasRecordingStatusCallbackMethod

`func (o *DeprecatedInitiateCallRequest) HasRecordingStatusCallbackMethod() bool`

HasRecordingStatusCallbackMethod returns a boolean if a field has been set.

### GetRecordingStatusCallbackEvent

`func (o *DeprecatedInitiateCallRequest) GetRecordingStatusCallbackEvent() string`

GetRecordingStatusCallbackEvent returns the RecordingStatusCallbackEvent field if non-nil, zero value otherwise.

### GetRecordingStatusCallbackEventOk

`func (o *DeprecatedInitiateCallRequest) GetRecordingStatusCallbackEventOk() (*string, bool)`

GetRecordingStatusCallbackEventOk returns a tuple with the RecordingStatusCallbackEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatusCallbackEvent

`func (o *DeprecatedInitiateCallRequest) SetRecordingStatusCallbackEvent(v string)`

SetRecordingStatusCallbackEvent sets RecordingStatusCallbackEvent field to given value.

### HasRecordingStatusCallbackEvent

`func (o *DeprecatedInitiateCallRequest) HasRecordingStatusCallbackEvent() bool`

HasRecordingStatusCallbackEvent returns a boolean if a field has been set.

### GetRecordingTimeout

`func (o *DeprecatedInitiateCallRequest) GetRecordingTimeout() int32`

GetRecordingTimeout returns the RecordingTimeout field if non-nil, zero value otherwise.

### GetRecordingTimeoutOk

`func (o *DeprecatedInitiateCallRequest) GetRecordingTimeoutOk() (*int32, bool)`

GetRecordingTimeoutOk returns a tuple with the RecordingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTimeout

`func (o *DeprecatedInitiateCallRequest) SetRecordingTimeout(v int32)`

SetRecordingTimeout sets RecordingTimeout field to given value.

### HasRecordingTimeout

`func (o *DeprecatedInitiateCallRequest) HasRecordingTimeout() bool`

HasRecordingTimeout returns a boolean if a field has been set.

### GetRecordingTrack

`func (o *DeprecatedInitiateCallRequest) GetRecordingTrack() string`

GetRecordingTrack returns the RecordingTrack field if non-nil, zero value otherwise.

### GetRecordingTrackOk

`func (o *DeprecatedInitiateCallRequest) GetRecordingTrackOk() (*string, bool)`

GetRecordingTrackOk returns a tuple with the RecordingTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTrack

`func (o *DeprecatedInitiateCallRequest) SetRecordingTrack(v string)`

SetRecordingTrack sets RecordingTrack field to given value.

### HasRecordingTrack

`func (o *DeprecatedInitiateCallRequest) HasRecordingTrack() bool`

HasRecordingTrack returns a boolean if a field has been set.

### GetSipAuthPassword

`func (o *DeprecatedInitiateCallRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *DeprecatedInitiateCallRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *DeprecatedInitiateCallRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *DeprecatedInitiateCallRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *DeprecatedInitiateCallRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *DeprecatedInitiateCallRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *DeprecatedInitiateCallRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *DeprecatedInitiateCallRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### GetTrim

`func (o *DeprecatedInitiateCallRequest) GetTrim() string`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *DeprecatedInitiateCallRequest) GetTrimOk() (*string, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *DeprecatedInitiateCallRequest) SetTrim(v string)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *DeprecatedInitiateCallRequest) HasTrim() bool`

HasTrim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


