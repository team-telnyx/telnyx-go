# InitiateCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationSid** | **string** | The ID of the TeXML Application. | 
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

### NewInitiateCallRequest

`func NewInitiateCallRequest(applicationSid string, to string, from string, ) *InitiateCallRequest`

NewInitiateCallRequest instantiates a new InitiateCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateCallRequestWithDefaults

`func NewInitiateCallRequestWithDefaults() *InitiateCallRequest`

NewInitiateCallRequestWithDefaults instantiates a new InitiateCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationSid

`func (o *InitiateCallRequest) GetApplicationSid() string`

GetApplicationSid returns the ApplicationSid field if non-nil, zero value otherwise.

### GetApplicationSidOk

`func (o *InitiateCallRequest) GetApplicationSidOk() (*string, bool)`

GetApplicationSidOk returns a tuple with the ApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSid

`func (o *InitiateCallRequest) SetApplicationSid(v string)`

SetApplicationSid sets ApplicationSid field to given value.


### GetTo

`func (o *InitiateCallRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InitiateCallRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InitiateCallRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFrom

`func (o *InitiateCallRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InitiateCallRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InitiateCallRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetCallerId

`func (o *InitiateCallRequest) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *InitiateCallRequest) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *InitiateCallRequest) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *InitiateCallRequest) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetUrl

`func (o *InitiateCallRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InitiateCallRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InitiateCallRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InitiateCallRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlMethod

`func (o *InitiateCallRequest) GetUrlMethod() string`

GetUrlMethod returns the UrlMethod field if non-nil, zero value otherwise.

### GetUrlMethodOk

`func (o *InitiateCallRequest) GetUrlMethodOk() (*string, bool)`

GetUrlMethodOk returns a tuple with the UrlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlMethod

`func (o *InitiateCallRequest) SetUrlMethod(v string)`

SetUrlMethod sets UrlMethod field to given value.

### HasUrlMethod

`func (o *InitiateCallRequest) HasUrlMethod() bool`

HasUrlMethod returns a boolean if a field has been set.

### GetFallbackUrl

`func (o *InitiateCallRequest) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *InitiateCallRequest) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *InitiateCallRequest) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *InitiateCallRequest) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### GetStatusCallback

`func (o *InitiateCallRequest) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *InitiateCallRequest) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *InitiateCallRequest) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *InitiateCallRequest) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### GetStatusCallbackMethod

`func (o *InitiateCallRequest) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *InitiateCallRequest) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *InitiateCallRequest) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *InitiateCallRequest) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### GetStatusCallbackEvent

`func (o *InitiateCallRequest) GetStatusCallbackEvent() string`

GetStatusCallbackEvent returns the StatusCallbackEvent field if non-nil, zero value otherwise.

### GetStatusCallbackEventOk

`func (o *InitiateCallRequest) GetStatusCallbackEventOk() (*string, bool)`

GetStatusCallbackEventOk returns a tuple with the StatusCallbackEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackEvent

`func (o *InitiateCallRequest) SetStatusCallbackEvent(v string)`

SetStatusCallbackEvent sets StatusCallbackEvent field to given value.

### HasStatusCallbackEvent

`func (o *InitiateCallRequest) HasStatusCallbackEvent() bool`

HasStatusCallbackEvent returns a boolean if a field has been set.

### GetMachineDetection

`func (o *InitiateCallRequest) GetMachineDetection() string`

GetMachineDetection returns the MachineDetection field if non-nil, zero value otherwise.

### GetMachineDetectionOk

`func (o *InitiateCallRequest) GetMachineDetectionOk() (*string, bool)`

GetMachineDetectionOk returns a tuple with the MachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetection

`func (o *InitiateCallRequest) SetMachineDetection(v string)`

SetMachineDetection sets MachineDetection field to given value.

### HasMachineDetection

`func (o *InitiateCallRequest) HasMachineDetection() bool`

HasMachineDetection returns a boolean if a field has been set.

### GetDetectionMode

`func (o *InitiateCallRequest) GetDetectionMode() string`

GetDetectionMode returns the DetectionMode field if non-nil, zero value otherwise.

### GetDetectionModeOk

`func (o *InitiateCallRequest) GetDetectionModeOk() (*string, bool)`

GetDetectionModeOk returns a tuple with the DetectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMode

`func (o *InitiateCallRequest) SetDetectionMode(v string)`

SetDetectionMode sets DetectionMode field to given value.

### HasDetectionMode

`func (o *InitiateCallRequest) HasDetectionMode() bool`

HasDetectionMode returns a boolean if a field has been set.

### GetAsyncAmd

`func (o *InitiateCallRequest) GetAsyncAmd() bool`

GetAsyncAmd returns the AsyncAmd field if non-nil, zero value otherwise.

### GetAsyncAmdOk

`func (o *InitiateCallRequest) GetAsyncAmdOk() (*bool, bool)`

GetAsyncAmdOk returns a tuple with the AsyncAmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncAmd

`func (o *InitiateCallRequest) SetAsyncAmd(v bool)`

SetAsyncAmd sets AsyncAmd field to given value.

### HasAsyncAmd

`func (o *InitiateCallRequest) HasAsyncAmd() bool`

HasAsyncAmd returns a boolean if a field has been set.

### GetAsyncAmdStatusCallback

`func (o *InitiateCallRequest) GetAsyncAmdStatusCallback() string`

GetAsyncAmdStatusCallback returns the AsyncAmdStatusCallback field if non-nil, zero value otherwise.

### GetAsyncAmdStatusCallbackOk

`func (o *InitiateCallRequest) GetAsyncAmdStatusCallbackOk() (*string, bool)`

GetAsyncAmdStatusCallbackOk returns a tuple with the AsyncAmdStatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncAmdStatusCallback

`func (o *InitiateCallRequest) SetAsyncAmdStatusCallback(v string)`

SetAsyncAmdStatusCallback sets AsyncAmdStatusCallback field to given value.

### HasAsyncAmdStatusCallback

`func (o *InitiateCallRequest) HasAsyncAmdStatusCallback() bool`

HasAsyncAmdStatusCallback returns a boolean if a field has been set.

### GetAsyncAmdStatusCallbackMethod

`func (o *InitiateCallRequest) GetAsyncAmdStatusCallbackMethod() string`

GetAsyncAmdStatusCallbackMethod returns the AsyncAmdStatusCallbackMethod field if non-nil, zero value otherwise.

### GetAsyncAmdStatusCallbackMethodOk

`func (o *InitiateCallRequest) GetAsyncAmdStatusCallbackMethodOk() (*string, bool)`

GetAsyncAmdStatusCallbackMethodOk returns a tuple with the AsyncAmdStatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncAmdStatusCallbackMethod

`func (o *InitiateCallRequest) SetAsyncAmdStatusCallbackMethod(v string)`

SetAsyncAmdStatusCallbackMethod sets AsyncAmdStatusCallbackMethod field to given value.

### HasAsyncAmdStatusCallbackMethod

`func (o *InitiateCallRequest) HasAsyncAmdStatusCallbackMethod() bool`

HasAsyncAmdStatusCallbackMethod returns a boolean if a field has been set.

### GetMachineDetectionTimeout

`func (o *InitiateCallRequest) GetMachineDetectionTimeout() int32`

GetMachineDetectionTimeout returns the MachineDetectionTimeout field if non-nil, zero value otherwise.

### GetMachineDetectionTimeoutOk

`func (o *InitiateCallRequest) GetMachineDetectionTimeoutOk() (*int32, bool)`

GetMachineDetectionTimeoutOk returns a tuple with the MachineDetectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionTimeout

`func (o *InitiateCallRequest) SetMachineDetectionTimeout(v int32)`

SetMachineDetectionTimeout sets MachineDetectionTimeout field to given value.

### HasMachineDetectionTimeout

`func (o *InitiateCallRequest) HasMachineDetectionTimeout() bool`

HasMachineDetectionTimeout returns a boolean if a field has been set.

### GetMachineDetectionSpeechThreshold

`func (o *InitiateCallRequest) GetMachineDetectionSpeechThreshold() int32`

GetMachineDetectionSpeechThreshold returns the MachineDetectionSpeechThreshold field if non-nil, zero value otherwise.

### GetMachineDetectionSpeechThresholdOk

`func (o *InitiateCallRequest) GetMachineDetectionSpeechThresholdOk() (*int32, bool)`

GetMachineDetectionSpeechThresholdOk returns a tuple with the MachineDetectionSpeechThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionSpeechThreshold

`func (o *InitiateCallRequest) SetMachineDetectionSpeechThreshold(v int32)`

SetMachineDetectionSpeechThreshold sets MachineDetectionSpeechThreshold field to given value.

### HasMachineDetectionSpeechThreshold

`func (o *InitiateCallRequest) HasMachineDetectionSpeechThreshold() bool`

HasMachineDetectionSpeechThreshold returns a boolean if a field has been set.

### GetMachineDetectionSpeechEndThreshold

`func (o *InitiateCallRequest) GetMachineDetectionSpeechEndThreshold() int32`

GetMachineDetectionSpeechEndThreshold returns the MachineDetectionSpeechEndThreshold field if non-nil, zero value otherwise.

### GetMachineDetectionSpeechEndThresholdOk

`func (o *InitiateCallRequest) GetMachineDetectionSpeechEndThresholdOk() (*int32, bool)`

GetMachineDetectionSpeechEndThresholdOk returns a tuple with the MachineDetectionSpeechEndThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionSpeechEndThreshold

`func (o *InitiateCallRequest) SetMachineDetectionSpeechEndThreshold(v int32)`

SetMachineDetectionSpeechEndThreshold sets MachineDetectionSpeechEndThreshold field to given value.

### HasMachineDetectionSpeechEndThreshold

`func (o *InitiateCallRequest) HasMachineDetectionSpeechEndThreshold() bool`

HasMachineDetectionSpeechEndThreshold returns a boolean if a field has been set.

### GetMachineDetectionSilenceTimeout

`func (o *InitiateCallRequest) GetMachineDetectionSilenceTimeout() int32`

GetMachineDetectionSilenceTimeout returns the MachineDetectionSilenceTimeout field if non-nil, zero value otherwise.

### GetMachineDetectionSilenceTimeoutOk

`func (o *InitiateCallRequest) GetMachineDetectionSilenceTimeoutOk() (*int32, bool)`

GetMachineDetectionSilenceTimeoutOk returns a tuple with the MachineDetectionSilenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDetectionSilenceTimeout

`func (o *InitiateCallRequest) SetMachineDetectionSilenceTimeout(v int32)`

SetMachineDetectionSilenceTimeout sets MachineDetectionSilenceTimeout field to given value.

### HasMachineDetectionSilenceTimeout

`func (o *InitiateCallRequest) HasMachineDetectionSilenceTimeout() bool`

HasMachineDetectionSilenceTimeout returns a boolean if a field has been set.

### GetCancelPlaybackOnMachineDetection

`func (o *InitiateCallRequest) GetCancelPlaybackOnMachineDetection() bool`

GetCancelPlaybackOnMachineDetection returns the CancelPlaybackOnMachineDetection field if non-nil, zero value otherwise.

### GetCancelPlaybackOnMachineDetectionOk

`func (o *InitiateCallRequest) GetCancelPlaybackOnMachineDetectionOk() (*bool, bool)`

GetCancelPlaybackOnMachineDetectionOk returns a tuple with the CancelPlaybackOnMachineDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelPlaybackOnMachineDetection

`func (o *InitiateCallRequest) SetCancelPlaybackOnMachineDetection(v bool)`

SetCancelPlaybackOnMachineDetection sets CancelPlaybackOnMachineDetection field to given value.

### HasCancelPlaybackOnMachineDetection

`func (o *InitiateCallRequest) HasCancelPlaybackOnMachineDetection() bool`

HasCancelPlaybackOnMachineDetection returns a boolean if a field has been set.

### GetCancelPlaybackOnDetectMessageEnd

`func (o *InitiateCallRequest) GetCancelPlaybackOnDetectMessageEnd() bool`

GetCancelPlaybackOnDetectMessageEnd returns the CancelPlaybackOnDetectMessageEnd field if non-nil, zero value otherwise.

### GetCancelPlaybackOnDetectMessageEndOk

`func (o *InitiateCallRequest) GetCancelPlaybackOnDetectMessageEndOk() (*bool, bool)`

GetCancelPlaybackOnDetectMessageEndOk returns a tuple with the CancelPlaybackOnDetectMessageEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelPlaybackOnDetectMessageEnd

`func (o *InitiateCallRequest) SetCancelPlaybackOnDetectMessageEnd(v bool)`

SetCancelPlaybackOnDetectMessageEnd sets CancelPlaybackOnDetectMessageEnd field to given value.

### HasCancelPlaybackOnDetectMessageEnd

`func (o *InitiateCallRequest) HasCancelPlaybackOnDetectMessageEnd() bool`

HasCancelPlaybackOnDetectMessageEnd returns a boolean if a field has been set.

### GetPreferredCodecs

`func (o *InitiateCallRequest) GetPreferredCodecs() string`

GetPreferredCodecs returns the PreferredCodecs field if non-nil, zero value otherwise.

### GetPreferredCodecsOk

`func (o *InitiateCallRequest) GetPreferredCodecsOk() (*string, bool)`

GetPreferredCodecsOk returns a tuple with the PreferredCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredCodecs

`func (o *InitiateCallRequest) SetPreferredCodecs(v string)`

SetPreferredCodecs sets PreferredCodecs field to given value.

### HasPreferredCodecs

`func (o *InitiateCallRequest) HasPreferredCodecs() bool`

HasPreferredCodecs returns a boolean if a field has been set.

### GetRecord

`func (o *InitiateCallRequest) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *InitiateCallRequest) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *InitiateCallRequest) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *InitiateCallRequest) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRecordingChannels

`func (o *InitiateCallRequest) GetRecordingChannels() string`

GetRecordingChannels returns the RecordingChannels field if non-nil, zero value otherwise.

### GetRecordingChannelsOk

`func (o *InitiateCallRequest) GetRecordingChannelsOk() (*string, bool)`

GetRecordingChannelsOk returns a tuple with the RecordingChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingChannels

`func (o *InitiateCallRequest) SetRecordingChannels(v string)`

SetRecordingChannels sets RecordingChannels field to given value.

### HasRecordingChannels

`func (o *InitiateCallRequest) HasRecordingChannels() bool`

HasRecordingChannels returns a boolean if a field has been set.

### GetRecordingStatusCallback

`func (o *InitiateCallRequest) GetRecordingStatusCallback() string`

GetRecordingStatusCallback returns the RecordingStatusCallback field if non-nil, zero value otherwise.

### GetRecordingStatusCallbackOk

`func (o *InitiateCallRequest) GetRecordingStatusCallbackOk() (*string, bool)`

GetRecordingStatusCallbackOk returns a tuple with the RecordingStatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatusCallback

`func (o *InitiateCallRequest) SetRecordingStatusCallback(v string)`

SetRecordingStatusCallback sets RecordingStatusCallback field to given value.

### HasRecordingStatusCallback

`func (o *InitiateCallRequest) HasRecordingStatusCallback() bool`

HasRecordingStatusCallback returns a boolean if a field has been set.

### GetRecordingStatusCallbackMethod

`func (o *InitiateCallRequest) GetRecordingStatusCallbackMethod() string`

GetRecordingStatusCallbackMethod returns the RecordingStatusCallbackMethod field if non-nil, zero value otherwise.

### GetRecordingStatusCallbackMethodOk

`func (o *InitiateCallRequest) GetRecordingStatusCallbackMethodOk() (*string, bool)`

GetRecordingStatusCallbackMethodOk returns a tuple with the RecordingStatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatusCallbackMethod

`func (o *InitiateCallRequest) SetRecordingStatusCallbackMethod(v string)`

SetRecordingStatusCallbackMethod sets RecordingStatusCallbackMethod field to given value.

### HasRecordingStatusCallbackMethod

`func (o *InitiateCallRequest) HasRecordingStatusCallbackMethod() bool`

HasRecordingStatusCallbackMethod returns a boolean if a field has been set.

### GetRecordingStatusCallbackEvent

`func (o *InitiateCallRequest) GetRecordingStatusCallbackEvent() string`

GetRecordingStatusCallbackEvent returns the RecordingStatusCallbackEvent field if non-nil, zero value otherwise.

### GetRecordingStatusCallbackEventOk

`func (o *InitiateCallRequest) GetRecordingStatusCallbackEventOk() (*string, bool)`

GetRecordingStatusCallbackEventOk returns a tuple with the RecordingStatusCallbackEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingStatusCallbackEvent

`func (o *InitiateCallRequest) SetRecordingStatusCallbackEvent(v string)`

SetRecordingStatusCallbackEvent sets RecordingStatusCallbackEvent field to given value.

### HasRecordingStatusCallbackEvent

`func (o *InitiateCallRequest) HasRecordingStatusCallbackEvent() bool`

HasRecordingStatusCallbackEvent returns a boolean if a field has been set.

### GetRecordingTimeout

`func (o *InitiateCallRequest) GetRecordingTimeout() int32`

GetRecordingTimeout returns the RecordingTimeout field if non-nil, zero value otherwise.

### GetRecordingTimeoutOk

`func (o *InitiateCallRequest) GetRecordingTimeoutOk() (*int32, bool)`

GetRecordingTimeoutOk returns a tuple with the RecordingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTimeout

`func (o *InitiateCallRequest) SetRecordingTimeout(v int32)`

SetRecordingTimeout sets RecordingTimeout field to given value.

### HasRecordingTimeout

`func (o *InitiateCallRequest) HasRecordingTimeout() bool`

HasRecordingTimeout returns a boolean if a field has been set.

### GetRecordingTrack

`func (o *InitiateCallRequest) GetRecordingTrack() string`

GetRecordingTrack returns the RecordingTrack field if non-nil, zero value otherwise.

### GetRecordingTrackOk

`func (o *InitiateCallRequest) GetRecordingTrackOk() (*string, bool)`

GetRecordingTrackOk returns a tuple with the RecordingTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTrack

`func (o *InitiateCallRequest) SetRecordingTrack(v string)`

SetRecordingTrack sets RecordingTrack field to given value.

### HasRecordingTrack

`func (o *InitiateCallRequest) HasRecordingTrack() bool`

HasRecordingTrack returns a boolean if a field has been set.

### GetSipAuthPassword

`func (o *InitiateCallRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *InitiateCallRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *InitiateCallRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *InitiateCallRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *InitiateCallRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *InitiateCallRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *InitiateCallRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *InitiateCallRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### GetTrim

`func (o *InitiateCallRequest) GetTrim() string`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *InitiateCallRequest) GetTrimOk() (*string, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *InitiateCallRequest) SetTrim(v string)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *InitiateCallRequest) HasTrim() bool`

HasTrim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


