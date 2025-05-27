# TranscriptionStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranscriptionEngine** | Pointer to **string** | Engine to use for speech recognition. &#x60;A&#x60; - &#x60;Google&#x60;, &#x60;B&#x60; - &#x60;Telnyx&#x60;. | [optional] [default to "A"]
**TranscriptionEngineConfig** | Pointer to [**TranscriptionStartRequestTranscriptionEngineConfig**](TranscriptionStartRequestTranscriptionEngineConfig.md) |  | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**TranscriptionTracks** | Pointer to **string** | Indicates which leg of the call will be transcribed. Use &#x60;inbound&#x60; for the leg that requested the transcription, &#x60;outbound&#x60; for the other leg, and &#x60;both&#x60; for both legs of the call. Will default to &#x60;inbound&#x60;. | [optional] [default to "inbound"]
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewTranscriptionStartRequest

`func NewTranscriptionStartRequest() *TranscriptionStartRequest`

NewTranscriptionStartRequest instantiates a new TranscriptionStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionStartRequestWithDefaults

`func NewTranscriptionStartRequestWithDefaults() *TranscriptionStartRequest`

NewTranscriptionStartRequestWithDefaults instantiates a new TranscriptionStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranscriptionEngine

`func (o *TranscriptionStartRequest) GetTranscriptionEngine() string`

GetTranscriptionEngine returns the TranscriptionEngine field if non-nil, zero value otherwise.

### GetTranscriptionEngineOk

`func (o *TranscriptionStartRequest) GetTranscriptionEngineOk() (*string, bool)`

GetTranscriptionEngineOk returns a tuple with the TranscriptionEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionEngine

`func (o *TranscriptionStartRequest) SetTranscriptionEngine(v string)`

SetTranscriptionEngine sets TranscriptionEngine field to given value.

### HasTranscriptionEngine

`func (o *TranscriptionStartRequest) HasTranscriptionEngine() bool`

HasTranscriptionEngine returns a boolean if a field has been set.

### GetTranscriptionEngineConfig

`func (o *TranscriptionStartRequest) GetTranscriptionEngineConfig() TranscriptionStartRequestTranscriptionEngineConfig`

GetTranscriptionEngineConfig returns the TranscriptionEngineConfig field if non-nil, zero value otherwise.

### GetTranscriptionEngineConfigOk

`func (o *TranscriptionStartRequest) GetTranscriptionEngineConfigOk() (*TranscriptionStartRequestTranscriptionEngineConfig, bool)`

GetTranscriptionEngineConfigOk returns a tuple with the TranscriptionEngineConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionEngineConfig

`func (o *TranscriptionStartRequest) SetTranscriptionEngineConfig(v TranscriptionStartRequestTranscriptionEngineConfig)`

SetTranscriptionEngineConfig sets TranscriptionEngineConfig field to given value.

### HasTranscriptionEngineConfig

`func (o *TranscriptionStartRequest) HasTranscriptionEngineConfig() bool`

HasTranscriptionEngineConfig returns a boolean if a field has been set.

### GetClientState

`func (o *TranscriptionStartRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *TranscriptionStartRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *TranscriptionStartRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *TranscriptionStartRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetTranscriptionTracks

`func (o *TranscriptionStartRequest) GetTranscriptionTracks() string`

GetTranscriptionTracks returns the TranscriptionTracks field if non-nil, zero value otherwise.

### GetTranscriptionTracksOk

`func (o *TranscriptionStartRequest) GetTranscriptionTracksOk() (*string, bool)`

GetTranscriptionTracksOk returns a tuple with the TranscriptionTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionTracks

`func (o *TranscriptionStartRequest) SetTranscriptionTracks(v string)`

SetTranscriptionTracks sets TranscriptionTracks field to given value.

### HasTranscriptionTracks

`func (o *TranscriptionStartRequest) HasTranscriptionTracks() bool`

HasTranscriptionTracks returns a boolean if a field has been set.

### GetCommandId

`func (o *TranscriptionStartRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *TranscriptionStartRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *TranscriptionStartRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *TranscriptionStartRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


