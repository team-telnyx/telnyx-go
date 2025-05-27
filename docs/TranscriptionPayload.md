# TranscriptionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Unique identifier and token for controlling the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**TranscriptionData** | Pointer to [**TranscriptionPayloadTranscriptionData**](TranscriptionPayloadTranscriptionData.md) |  | [optional] 

## Methods

### NewTranscriptionPayload

`func NewTranscriptionPayload() *TranscriptionPayload`

NewTranscriptionPayload instantiates a new TranscriptionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionPayloadWithDefaults

`func NewTranscriptionPayloadWithDefaults() *TranscriptionPayload`

NewTranscriptionPayloadWithDefaults instantiates a new TranscriptionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *TranscriptionPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *TranscriptionPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *TranscriptionPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *TranscriptionPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *TranscriptionPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *TranscriptionPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *TranscriptionPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *TranscriptionPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *TranscriptionPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *TranscriptionPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *TranscriptionPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *TranscriptionPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *TranscriptionPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *TranscriptionPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *TranscriptionPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *TranscriptionPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetConnectionId

`func (o *TranscriptionPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *TranscriptionPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *TranscriptionPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *TranscriptionPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetTranscriptionData

`func (o *TranscriptionPayload) GetTranscriptionData() TranscriptionPayloadTranscriptionData`

GetTranscriptionData returns the TranscriptionData field if non-nil, zero value otherwise.

### GetTranscriptionDataOk

`func (o *TranscriptionPayload) GetTranscriptionDataOk() (*TranscriptionPayloadTranscriptionData, bool)`

GetTranscriptionDataOk returns a tuple with the TranscriptionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionData

`func (o *TranscriptionPayload) SetTranscriptionData(v TranscriptionPayloadTranscriptionData)`

SetTranscriptionData sets TranscriptionData field to given value.

### HasTranscriptionData

`func (o *TranscriptionPayload) HasTranscriptionData() bool`

HasTranscriptionData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


