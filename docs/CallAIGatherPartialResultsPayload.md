# CallAIGatherPartialResultsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Telnyx connection ID used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**From** | Pointer to **string** | Number or SIP URI placing the call. | [optional] 
**To** | Pointer to **string** | Destination number or SIP URI of the call. | [optional] 
**MessageHistory** | Pointer to [**[]CallAIGatherEndedPayloadMessageHistoryInner**](CallAIGatherEndedPayloadMessageHistoryInner.md) | The history of the messages exchanged during the AI gather | [optional] 
**PartialResults** | Pointer to **map[string]interface{}** | The partial result of the AI gather, its type depends of the &#x60;parameters&#x60; provided in the command | [optional] 

## Methods

### NewCallAIGatherPartialResultsPayload

`func NewCallAIGatherPartialResultsPayload() *CallAIGatherPartialResultsPayload`

NewCallAIGatherPartialResultsPayload instantiates a new CallAIGatherPartialResultsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallAIGatherPartialResultsPayloadWithDefaults

`func NewCallAIGatherPartialResultsPayloadWithDefaults() *CallAIGatherPartialResultsPayload`

NewCallAIGatherPartialResultsPayloadWithDefaults instantiates a new CallAIGatherPartialResultsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallAIGatherPartialResultsPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallAIGatherPartialResultsPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallAIGatherPartialResultsPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallAIGatherPartialResultsPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallAIGatherPartialResultsPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallAIGatherPartialResultsPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallAIGatherPartialResultsPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallAIGatherPartialResultsPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallAIGatherPartialResultsPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallAIGatherPartialResultsPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallAIGatherPartialResultsPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallAIGatherPartialResultsPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallAIGatherPartialResultsPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallAIGatherPartialResultsPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallAIGatherPartialResultsPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallAIGatherPartialResultsPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallAIGatherPartialResultsPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallAIGatherPartialResultsPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallAIGatherPartialResultsPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallAIGatherPartialResultsPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFrom

`func (o *CallAIGatherPartialResultsPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallAIGatherPartialResultsPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallAIGatherPartialResultsPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallAIGatherPartialResultsPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallAIGatherPartialResultsPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallAIGatherPartialResultsPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallAIGatherPartialResultsPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallAIGatherPartialResultsPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessageHistory

`func (o *CallAIGatherPartialResultsPayload) GetMessageHistory() []CallAIGatherEndedPayloadMessageHistoryInner`

GetMessageHistory returns the MessageHistory field if non-nil, zero value otherwise.

### GetMessageHistoryOk

`func (o *CallAIGatherPartialResultsPayload) GetMessageHistoryOk() (*[]CallAIGatherEndedPayloadMessageHistoryInner, bool)`

GetMessageHistoryOk returns a tuple with the MessageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHistory

`func (o *CallAIGatherPartialResultsPayload) SetMessageHistory(v []CallAIGatherEndedPayloadMessageHistoryInner)`

SetMessageHistory sets MessageHistory field to given value.

### HasMessageHistory

`func (o *CallAIGatherPartialResultsPayload) HasMessageHistory() bool`

HasMessageHistory returns a boolean if a field has been set.

### GetPartialResults

`func (o *CallAIGatherPartialResultsPayload) GetPartialResults() map[string]interface{}`

GetPartialResults returns the PartialResults field if non-nil, zero value otherwise.

### GetPartialResultsOk

`func (o *CallAIGatherPartialResultsPayload) GetPartialResultsOk() (*map[string]interface{}, bool)`

GetPartialResultsOk returns a tuple with the PartialResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialResults

`func (o *CallAIGatherPartialResultsPayload) SetPartialResults(v map[string]interface{})`

SetPartialResults sets PartialResults field to given value.

### HasPartialResults

`func (o *CallAIGatherPartialResultsPayload) HasPartialResults() bool`

HasPartialResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


