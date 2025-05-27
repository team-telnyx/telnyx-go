# StopStreamingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**StreamId** | Pointer to **string** | Identifies the stream. If the &#x60;stream_id&#x60; is not provided the command stops all streams associated with a given &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewStopStreamingRequest

`func NewStopStreamingRequest() *StopStreamingRequest`

NewStopStreamingRequest instantiates a new StopStreamingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopStreamingRequestWithDefaults

`func NewStopStreamingRequestWithDefaults() *StopStreamingRequest`

NewStopStreamingRequestWithDefaults instantiates a new StopStreamingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientState

`func (o *StopStreamingRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *StopStreamingRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *StopStreamingRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *StopStreamingRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *StopStreamingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *StopStreamingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *StopStreamingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *StopStreamingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetStreamId

`func (o *StopStreamingRequest) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StopStreamingRequest) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StopStreamingRequest) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *StopStreamingRequest) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


