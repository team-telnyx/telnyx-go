# CallSpeakEndedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**Status** | Pointer to **string** | Reflects how the command ended. | [optional] 

## Methods

### NewCallSpeakEndedPayload

`func NewCallSpeakEndedPayload() *CallSpeakEndedPayload`

NewCallSpeakEndedPayload instantiates a new CallSpeakEndedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallSpeakEndedPayloadWithDefaults

`func NewCallSpeakEndedPayloadWithDefaults() *CallSpeakEndedPayload`

NewCallSpeakEndedPayloadWithDefaults instantiates a new CallSpeakEndedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallSpeakEndedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallSpeakEndedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallSpeakEndedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallSpeakEndedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallSpeakEndedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallSpeakEndedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallSpeakEndedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallSpeakEndedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallSpeakEndedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallSpeakEndedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallSpeakEndedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallSpeakEndedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallSpeakEndedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallSpeakEndedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallSpeakEndedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallSpeakEndedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallSpeakEndedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallSpeakEndedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallSpeakEndedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallSpeakEndedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetStatus

`func (o *CallSpeakEndedPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallSpeakEndedPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallSpeakEndedPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallSpeakEndedPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


