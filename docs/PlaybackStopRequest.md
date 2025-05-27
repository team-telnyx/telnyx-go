# PlaybackStopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overlay** | Pointer to **bool** | When enabled, it stops the audio being played in the overlay queue. | [optional] [default to false]
**Stop** | Pointer to **string** | Use &#x60;current&#x60; to stop the current audio being played. Use &#x60;all&#x60; to stop the current audio file being played and clear all audio files from the queue. | [optional] [default to "all"]
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewPlaybackStopRequest

`func NewPlaybackStopRequest() *PlaybackStopRequest`

NewPlaybackStopRequest instantiates a new PlaybackStopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybackStopRequestWithDefaults

`func NewPlaybackStopRequestWithDefaults() *PlaybackStopRequest`

NewPlaybackStopRequestWithDefaults instantiates a new PlaybackStopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverlay

`func (o *PlaybackStopRequest) GetOverlay() bool`

GetOverlay returns the Overlay field if non-nil, zero value otherwise.

### GetOverlayOk

`func (o *PlaybackStopRequest) GetOverlayOk() (*bool, bool)`

GetOverlayOk returns a tuple with the Overlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlay

`func (o *PlaybackStopRequest) SetOverlay(v bool)`

SetOverlay sets Overlay field to given value.

### HasOverlay

`func (o *PlaybackStopRequest) HasOverlay() bool`

HasOverlay returns a boolean if a field has been set.

### GetStop

`func (o *PlaybackStopRequest) GetStop() string`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *PlaybackStopRequest) GetStopOk() (*string, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *PlaybackStopRequest) SetStop(v string)`

SetStop sets Stop field to given value.

### HasStop

`func (o *PlaybackStopRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### GetClientState

`func (o *PlaybackStopRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *PlaybackStopRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *PlaybackStopRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *PlaybackStopRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *PlaybackStopRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *PlaybackStopRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *PlaybackStopRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *PlaybackStopRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


