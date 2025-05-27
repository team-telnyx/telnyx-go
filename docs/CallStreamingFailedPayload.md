# CallStreamingFailedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**FailureReason** | Pointer to **string** | A short description explaning why the media streaming failed. | [optional] 
**StreamId** | Pointer to **string** | Identifies the streaming. | [optional] 
**StreamParams** | Pointer to [**CallStreamingFailedPayloadStreamParams**](CallStreamingFailedPayloadStreamParams.md) |  | [optional] 
**StreamType** | Pointer to **string** | The type of stream connection the stream is performing. | [optional] 

## Methods

### NewCallStreamingFailedPayload

`func NewCallStreamingFailedPayload() *CallStreamingFailedPayload`

NewCallStreamingFailedPayload instantiates a new CallStreamingFailedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallStreamingFailedPayloadWithDefaults

`func NewCallStreamingFailedPayloadWithDefaults() *CallStreamingFailedPayload`

NewCallStreamingFailedPayloadWithDefaults instantiates a new CallStreamingFailedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallStreamingFailedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallStreamingFailedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallStreamingFailedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallStreamingFailedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallStreamingFailedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallStreamingFailedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallStreamingFailedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallStreamingFailedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallStreamingFailedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallStreamingFailedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallStreamingFailedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallStreamingFailedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallStreamingFailedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallStreamingFailedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallStreamingFailedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallStreamingFailedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallStreamingFailedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallStreamingFailedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallStreamingFailedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallStreamingFailedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFailureReason

`func (o *CallStreamingFailedPayload) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *CallStreamingFailedPayload) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *CallStreamingFailedPayload) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *CallStreamingFailedPayload) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetStreamId

`func (o *CallStreamingFailedPayload) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *CallStreamingFailedPayload) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *CallStreamingFailedPayload) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *CallStreamingFailedPayload) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetStreamParams

`func (o *CallStreamingFailedPayload) GetStreamParams() CallStreamingFailedPayloadStreamParams`

GetStreamParams returns the StreamParams field if non-nil, zero value otherwise.

### GetStreamParamsOk

`func (o *CallStreamingFailedPayload) GetStreamParamsOk() (*CallStreamingFailedPayloadStreamParams, bool)`

GetStreamParamsOk returns a tuple with the StreamParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamParams

`func (o *CallStreamingFailedPayload) SetStreamParams(v CallStreamingFailedPayloadStreamParams)`

SetStreamParams sets StreamParams field to given value.

### HasStreamParams

`func (o *CallStreamingFailedPayload) HasStreamParams() bool`

HasStreamParams returns a boolean if a field has been set.

### GetStreamType

`func (o *CallStreamingFailedPayload) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *CallStreamingFailedPayload) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *CallStreamingFailedPayload) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *CallStreamingFailedPayload) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


