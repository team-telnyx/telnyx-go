# BridgeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | The Call Control ID of the call you want to bridge with, can&#39;t be used together with queue parameter or video_room_id parameter. | 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**Queue** | Pointer to **string** | The name of the queue you want to bridge with, can&#39;t be used together with call_control_id parameter or video_room_id parameter. Bridging with a queue means bridging with the first call in the queue. The call will always be removed from the queue regardless of whether bridging succeeds. Returns an error when the queue is empty. | [optional] 
**VideoRoomId** | Pointer to **string** | The ID of the video room you want to bridge with, can&#39;t be used together with call_control_id parameter or queue parameter. | [optional] 
**VideoRoomContext** | Pointer to **string** | The additional parameter that will be passed to the video conference. It is a text field and the user can decide how to use it. For example, you can set the participant name or pass JSON text. It can be used only with video_room_id parameter. | [optional] 
**ParkAfterUnbridge** | Pointer to **string** | Specifies behavior after the bridge ends (i.e. the opposite leg either hangs up or is transferred). If supplied with the value &#x60;self&#x60;, the current leg will be parked after unbridge. If not set, the default behavior is to hang up the leg. | [optional] 
**PlayRingtone** | Pointer to **bool** | Specifies whether to play a ringtone if the call you want to bridge with has not yet been answered. | [optional] [default to false]
**Ringtone** | Pointer to **string** | Specifies which country ringtone to play when &#x60;play_ringtone&#x60; is set to &#x60;true&#x60;. If not set, the US ringtone will be played. | [optional] [default to "us"]
**Record** | Pointer to **string** | Start recording automatically after an event. Disabled by default. | [optional] 
**RecordChannels** | Pointer to **string** | Defines which channel should be recorded (&#39;single&#39; or &#39;dual&#39;) when &#x60;record&#x60; is specified. | [optional] [default to "dual"]
**RecordFormat** | Pointer to **string** | Defines the format of the recording (&#39;wav&#39; or &#39;mp3&#39;) when &#x60;record&#x60; is specified. | [optional] [default to "mp3"]
**RecordMaxLength** | Pointer to **int32** | Defines the maximum length for the recording in seconds when &#x60;record&#x60; is specified. The minimum value is 0. The maximum value is 43200. The default value is 0 (infinite). | [optional] [default to 0]
**RecordTimeoutSecs** | Pointer to **int32** | The number of seconds that Telnyx will wait for the recording to be stopped if silence is detected when &#x60;record&#x60; is specified. The timer only starts when the speech is detected. Please note that call transcription is used to detect silence and the related charge will be applied. The minimum value is 0. The default value is 0 (infinite). | [optional] [default to 0]
**RecordTrack** | Pointer to **string** | The audio track to be recorded. Can be either &#x60;both&#x60;, &#x60;inbound&#x60; or &#x60;outbound&#x60;. If only single track is specified (&#x60;inbound&#x60;, &#x60;outbound&#x60;), &#x60;channels&#x60; configuration is ignored and it will be recorded as mono (single channel). | [optional] [default to "both"]
**RecordTrim** | Pointer to **string** | When set to &#x60;trim-silence&#x60;, silence will be removed from the beginning and end of the recording. | [optional] 
**RecordCustomFileName** | Pointer to **string** | The custom recording file name to be used instead of the default &#x60;call_leg_id&#x60;. Telnyx will still add a Unix timestamp suffix. | [optional] 
**MuteDtmf** | Pointer to **string** | When enabled, DTMF tones are not passed to the call participant. The webhooks containing the DTMF information will be sent. | [optional] [default to "none"]

## Methods

### NewBridgeRequest

`func NewBridgeRequest(callControlId string, ) *BridgeRequest`

NewBridgeRequest instantiates a new BridgeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeRequestWithDefaults

`func NewBridgeRequestWithDefaults() *BridgeRequest`

NewBridgeRequestWithDefaults instantiates a new BridgeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *BridgeRequest) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *BridgeRequest) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *BridgeRequest) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetClientState

`func (o *BridgeRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *BridgeRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *BridgeRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *BridgeRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *BridgeRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *BridgeRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *BridgeRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *BridgeRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetQueue

`func (o *BridgeRequest) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *BridgeRequest) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *BridgeRequest) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *BridgeRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetVideoRoomId

`func (o *BridgeRequest) GetVideoRoomId() string`

GetVideoRoomId returns the VideoRoomId field if non-nil, zero value otherwise.

### GetVideoRoomIdOk

`func (o *BridgeRequest) GetVideoRoomIdOk() (*string, bool)`

GetVideoRoomIdOk returns a tuple with the VideoRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRoomId

`func (o *BridgeRequest) SetVideoRoomId(v string)`

SetVideoRoomId sets VideoRoomId field to given value.

### HasVideoRoomId

`func (o *BridgeRequest) HasVideoRoomId() bool`

HasVideoRoomId returns a boolean if a field has been set.

### GetVideoRoomContext

`func (o *BridgeRequest) GetVideoRoomContext() string`

GetVideoRoomContext returns the VideoRoomContext field if non-nil, zero value otherwise.

### GetVideoRoomContextOk

`func (o *BridgeRequest) GetVideoRoomContextOk() (*string, bool)`

GetVideoRoomContextOk returns a tuple with the VideoRoomContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRoomContext

`func (o *BridgeRequest) SetVideoRoomContext(v string)`

SetVideoRoomContext sets VideoRoomContext field to given value.

### HasVideoRoomContext

`func (o *BridgeRequest) HasVideoRoomContext() bool`

HasVideoRoomContext returns a boolean if a field has been set.

### GetParkAfterUnbridge

`func (o *BridgeRequest) GetParkAfterUnbridge() string`

GetParkAfterUnbridge returns the ParkAfterUnbridge field if non-nil, zero value otherwise.

### GetParkAfterUnbridgeOk

`func (o *BridgeRequest) GetParkAfterUnbridgeOk() (*string, bool)`

GetParkAfterUnbridgeOk returns a tuple with the ParkAfterUnbridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkAfterUnbridge

`func (o *BridgeRequest) SetParkAfterUnbridge(v string)`

SetParkAfterUnbridge sets ParkAfterUnbridge field to given value.

### HasParkAfterUnbridge

`func (o *BridgeRequest) HasParkAfterUnbridge() bool`

HasParkAfterUnbridge returns a boolean if a field has been set.

### GetPlayRingtone

`func (o *BridgeRequest) GetPlayRingtone() bool`

GetPlayRingtone returns the PlayRingtone field if non-nil, zero value otherwise.

### GetPlayRingtoneOk

`func (o *BridgeRequest) GetPlayRingtoneOk() (*bool, bool)`

GetPlayRingtoneOk returns a tuple with the PlayRingtone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayRingtone

`func (o *BridgeRequest) SetPlayRingtone(v bool)`

SetPlayRingtone sets PlayRingtone field to given value.

### HasPlayRingtone

`func (o *BridgeRequest) HasPlayRingtone() bool`

HasPlayRingtone returns a boolean if a field has been set.

### GetRingtone

`func (o *BridgeRequest) GetRingtone() string`

GetRingtone returns the Ringtone field if non-nil, zero value otherwise.

### GetRingtoneOk

`func (o *BridgeRequest) GetRingtoneOk() (*string, bool)`

GetRingtoneOk returns a tuple with the Ringtone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingtone

`func (o *BridgeRequest) SetRingtone(v string)`

SetRingtone sets Ringtone field to given value.

### HasRingtone

`func (o *BridgeRequest) HasRingtone() bool`

HasRingtone returns a boolean if a field has been set.

### GetRecord

`func (o *BridgeRequest) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *BridgeRequest) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *BridgeRequest) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *BridgeRequest) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRecordChannels

`func (o *BridgeRequest) GetRecordChannels() string`

GetRecordChannels returns the RecordChannels field if non-nil, zero value otherwise.

### GetRecordChannelsOk

`func (o *BridgeRequest) GetRecordChannelsOk() (*string, bool)`

GetRecordChannelsOk returns a tuple with the RecordChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordChannels

`func (o *BridgeRequest) SetRecordChannels(v string)`

SetRecordChannels sets RecordChannels field to given value.

### HasRecordChannels

`func (o *BridgeRequest) HasRecordChannels() bool`

HasRecordChannels returns a boolean if a field has been set.

### GetRecordFormat

`func (o *BridgeRequest) GetRecordFormat() string`

GetRecordFormat returns the RecordFormat field if non-nil, zero value otherwise.

### GetRecordFormatOk

`func (o *BridgeRequest) GetRecordFormatOk() (*string, bool)`

GetRecordFormatOk returns a tuple with the RecordFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordFormat

`func (o *BridgeRequest) SetRecordFormat(v string)`

SetRecordFormat sets RecordFormat field to given value.

### HasRecordFormat

`func (o *BridgeRequest) HasRecordFormat() bool`

HasRecordFormat returns a boolean if a field has been set.

### GetRecordMaxLength

`func (o *BridgeRequest) GetRecordMaxLength() int32`

GetRecordMaxLength returns the RecordMaxLength field if non-nil, zero value otherwise.

### GetRecordMaxLengthOk

`func (o *BridgeRequest) GetRecordMaxLengthOk() (*int32, bool)`

GetRecordMaxLengthOk returns a tuple with the RecordMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordMaxLength

`func (o *BridgeRequest) SetRecordMaxLength(v int32)`

SetRecordMaxLength sets RecordMaxLength field to given value.

### HasRecordMaxLength

`func (o *BridgeRequest) HasRecordMaxLength() bool`

HasRecordMaxLength returns a boolean if a field has been set.

### GetRecordTimeoutSecs

`func (o *BridgeRequest) GetRecordTimeoutSecs() int32`

GetRecordTimeoutSecs returns the RecordTimeoutSecs field if non-nil, zero value otherwise.

### GetRecordTimeoutSecsOk

`func (o *BridgeRequest) GetRecordTimeoutSecsOk() (*int32, bool)`

GetRecordTimeoutSecsOk returns a tuple with the RecordTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTimeoutSecs

`func (o *BridgeRequest) SetRecordTimeoutSecs(v int32)`

SetRecordTimeoutSecs sets RecordTimeoutSecs field to given value.

### HasRecordTimeoutSecs

`func (o *BridgeRequest) HasRecordTimeoutSecs() bool`

HasRecordTimeoutSecs returns a boolean if a field has been set.

### GetRecordTrack

`func (o *BridgeRequest) GetRecordTrack() string`

GetRecordTrack returns the RecordTrack field if non-nil, zero value otherwise.

### GetRecordTrackOk

`func (o *BridgeRequest) GetRecordTrackOk() (*string, bool)`

GetRecordTrackOk returns a tuple with the RecordTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTrack

`func (o *BridgeRequest) SetRecordTrack(v string)`

SetRecordTrack sets RecordTrack field to given value.

### HasRecordTrack

`func (o *BridgeRequest) HasRecordTrack() bool`

HasRecordTrack returns a boolean if a field has been set.

### GetRecordTrim

`func (o *BridgeRequest) GetRecordTrim() string`

GetRecordTrim returns the RecordTrim field if non-nil, zero value otherwise.

### GetRecordTrimOk

`func (o *BridgeRequest) GetRecordTrimOk() (*string, bool)`

GetRecordTrimOk returns a tuple with the RecordTrim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTrim

`func (o *BridgeRequest) SetRecordTrim(v string)`

SetRecordTrim sets RecordTrim field to given value.

### HasRecordTrim

`func (o *BridgeRequest) HasRecordTrim() bool`

HasRecordTrim returns a boolean if a field has been set.

### GetRecordCustomFileName

`func (o *BridgeRequest) GetRecordCustomFileName() string`

GetRecordCustomFileName returns the RecordCustomFileName field if non-nil, zero value otherwise.

### GetRecordCustomFileNameOk

`func (o *BridgeRequest) GetRecordCustomFileNameOk() (*string, bool)`

GetRecordCustomFileNameOk returns a tuple with the RecordCustomFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCustomFileName

`func (o *BridgeRequest) SetRecordCustomFileName(v string)`

SetRecordCustomFileName sets RecordCustomFileName field to given value.

### HasRecordCustomFileName

`func (o *BridgeRequest) HasRecordCustomFileName() bool`

HasRecordCustomFileName returns a boolean if a field has been set.

### GetMuteDtmf

`func (o *BridgeRequest) GetMuteDtmf() string`

GetMuteDtmf returns the MuteDtmf field if non-nil, zero value otherwise.

### GetMuteDtmfOk

`func (o *BridgeRequest) GetMuteDtmfOk() (*string, bool)`

GetMuteDtmfOk returns a tuple with the MuteDtmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteDtmf

`func (o *BridgeRequest) SetMuteDtmf(v string)`

SetMuteDtmf sets MuteDtmf field to given value.

### HasMuteDtmf

`func (o *BridgeRequest) HasMuteDtmf() bool`

HasMuteDtmf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


