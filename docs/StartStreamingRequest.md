# StartStreamingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamUrl** | Pointer to **string** | The destination WebSocket address where the stream is going to be delivered. | [optional] 
**StreamTrack** | Pointer to **string** | Specifies which track should be streamed. | [optional] [default to "inbound_track"]
**StreamBidirectionalMode** | Pointer to [**StreamBidirectionalMode**](StreamBidirectionalMode.md) |  | [optional] [default to MP3]
**StreamBidirectionalCodec** | Pointer to [**StreamBidirectionalCodec**](StreamBidirectionalCodec.md) |  | [optional] [default to PCMU]
**StreamBidirectionalTargetLegs** | Pointer to [**StreamBidirectionalTargetLegs**](StreamBidirectionalTargetLegs.md) |  | [optional] [default to OPPOSITE]
**EnableDialogflow** | Pointer to **bool** | Enables Dialogflow for the current call. The default value is false. | [optional] [default to false]
**DialogflowConfig** | Pointer to [**DialogflowConfig**](DialogflowConfig.md) |  | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewStartStreamingRequest

`func NewStartStreamingRequest() *StartStreamingRequest`

NewStartStreamingRequest instantiates a new StartStreamingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartStreamingRequestWithDefaults

`func NewStartStreamingRequestWithDefaults() *StartStreamingRequest`

NewStartStreamingRequestWithDefaults instantiates a new StartStreamingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamUrl

`func (o *StartStreamingRequest) GetStreamUrl() string`

GetStreamUrl returns the StreamUrl field if non-nil, zero value otherwise.

### GetStreamUrlOk

`func (o *StartStreamingRequest) GetStreamUrlOk() (*string, bool)`

GetStreamUrlOk returns a tuple with the StreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamUrl

`func (o *StartStreamingRequest) SetStreamUrl(v string)`

SetStreamUrl sets StreamUrl field to given value.

### HasStreamUrl

`func (o *StartStreamingRequest) HasStreamUrl() bool`

HasStreamUrl returns a boolean if a field has been set.

### GetStreamTrack

`func (o *StartStreamingRequest) GetStreamTrack() string`

GetStreamTrack returns the StreamTrack field if non-nil, zero value otherwise.

### GetStreamTrackOk

`func (o *StartStreamingRequest) GetStreamTrackOk() (*string, bool)`

GetStreamTrackOk returns a tuple with the StreamTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamTrack

`func (o *StartStreamingRequest) SetStreamTrack(v string)`

SetStreamTrack sets StreamTrack field to given value.

### HasStreamTrack

`func (o *StartStreamingRequest) HasStreamTrack() bool`

HasStreamTrack returns a boolean if a field has been set.

### GetStreamBidirectionalMode

`func (o *StartStreamingRequest) GetStreamBidirectionalMode() StreamBidirectionalMode`

GetStreamBidirectionalMode returns the StreamBidirectionalMode field if non-nil, zero value otherwise.

### GetStreamBidirectionalModeOk

`func (o *StartStreamingRequest) GetStreamBidirectionalModeOk() (*StreamBidirectionalMode, bool)`

GetStreamBidirectionalModeOk returns a tuple with the StreamBidirectionalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalMode

`func (o *StartStreamingRequest) SetStreamBidirectionalMode(v StreamBidirectionalMode)`

SetStreamBidirectionalMode sets StreamBidirectionalMode field to given value.

### HasStreamBidirectionalMode

`func (o *StartStreamingRequest) HasStreamBidirectionalMode() bool`

HasStreamBidirectionalMode returns a boolean if a field has been set.

### GetStreamBidirectionalCodec

`func (o *StartStreamingRequest) GetStreamBidirectionalCodec() StreamBidirectionalCodec`

GetStreamBidirectionalCodec returns the StreamBidirectionalCodec field if non-nil, zero value otherwise.

### GetStreamBidirectionalCodecOk

`func (o *StartStreamingRequest) GetStreamBidirectionalCodecOk() (*StreamBidirectionalCodec, bool)`

GetStreamBidirectionalCodecOk returns a tuple with the StreamBidirectionalCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalCodec

`func (o *StartStreamingRequest) SetStreamBidirectionalCodec(v StreamBidirectionalCodec)`

SetStreamBidirectionalCodec sets StreamBidirectionalCodec field to given value.

### HasStreamBidirectionalCodec

`func (o *StartStreamingRequest) HasStreamBidirectionalCodec() bool`

HasStreamBidirectionalCodec returns a boolean if a field has been set.

### GetStreamBidirectionalTargetLegs

`func (o *StartStreamingRequest) GetStreamBidirectionalTargetLegs() StreamBidirectionalTargetLegs`

GetStreamBidirectionalTargetLegs returns the StreamBidirectionalTargetLegs field if non-nil, zero value otherwise.

### GetStreamBidirectionalTargetLegsOk

`func (o *StartStreamingRequest) GetStreamBidirectionalTargetLegsOk() (*StreamBidirectionalTargetLegs, bool)`

GetStreamBidirectionalTargetLegsOk returns a tuple with the StreamBidirectionalTargetLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamBidirectionalTargetLegs

`func (o *StartStreamingRequest) SetStreamBidirectionalTargetLegs(v StreamBidirectionalTargetLegs)`

SetStreamBidirectionalTargetLegs sets StreamBidirectionalTargetLegs field to given value.

### HasStreamBidirectionalTargetLegs

`func (o *StartStreamingRequest) HasStreamBidirectionalTargetLegs() bool`

HasStreamBidirectionalTargetLegs returns a boolean if a field has been set.

### GetEnableDialogflow

`func (o *StartStreamingRequest) GetEnableDialogflow() bool`

GetEnableDialogflow returns the EnableDialogflow field if non-nil, zero value otherwise.

### GetEnableDialogflowOk

`func (o *StartStreamingRequest) GetEnableDialogflowOk() (*bool, bool)`

GetEnableDialogflowOk returns a tuple with the EnableDialogflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDialogflow

`func (o *StartStreamingRequest) SetEnableDialogflow(v bool)`

SetEnableDialogflow sets EnableDialogflow field to given value.

### HasEnableDialogflow

`func (o *StartStreamingRequest) HasEnableDialogflow() bool`

HasEnableDialogflow returns a boolean if a field has been set.

### GetDialogflowConfig

`func (o *StartStreamingRequest) GetDialogflowConfig() DialogflowConfig`

GetDialogflowConfig returns the DialogflowConfig field if non-nil, zero value otherwise.

### GetDialogflowConfigOk

`func (o *StartStreamingRequest) GetDialogflowConfigOk() (*DialogflowConfig, bool)`

GetDialogflowConfigOk returns a tuple with the DialogflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialogflowConfig

`func (o *StartStreamingRequest) SetDialogflowConfig(v DialogflowConfig)`

SetDialogflowConfig sets DialogflowConfig field to given value.

### HasDialogflowConfig

`func (o *StartStreamingRequest) HasDialogflowConfig() bool`

HasDialogflowConfig returns a boolean if a field has been set.

### GetClientState

`func (o *StartStreamingRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *StartStreamingRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *StartStreamingRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *StartStreamingRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *StartStreamingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *StartStreamingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *StartStreamingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *StartStreamingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


