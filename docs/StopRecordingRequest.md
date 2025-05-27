# StopRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**RecordingId** | Pointer to **string** | Uniquely identifies the resource. | [optional] 

## Methods

### NewStopRecordingRequest

`func NewStopRecordingRequest() *StopRecordingRequest`

NewStopRecordingRequest instantiates a new StopRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopRecordingRequestWithDefaults

`func NewStopRecordingRequestWithDefaults() *StopRecordingRequest`

NewStopRecordingRequestWithDefaults instantiates a new StopRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientState

`func (o *StopRecordingRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *StopRecordingRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *StopRecordingRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *StopRecordingRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *StopRecordingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *StopRecordingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *StopRecordingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *StopRecordingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetRecordingId

`func (o *StopRecordingRequest) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *StopRecordingRequest) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *StopRecordingRequest) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *StopRecordingRequest) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


