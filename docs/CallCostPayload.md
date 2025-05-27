# CallCostPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledDurationSecs** | Pointer to **int32** | The number of seconds for which this call will be billed | [optional] 
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the type of resource. | [optional] 
**CostParts** | Pointer to [**[]CallCostPayloadCostPartsInner**](CallCostPayloadCostPartsInner.md) |  | [optional] 
**TotalCost** | Pointer to **float32** | The billed cost of the call | [optional] 
**Status** | Pointer to **string** | Reflects how command ended. | [optional] 

## Methods

### NewCallCostPayload

`func NewCallCostPayload() *CallCostPayload`

NewCallCostPayload instantiates a new CallCostPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallCostPayloadWithDefaults

`func NewCallCostPayloadWithDefaults() *CallCostPayload`

NewCallCostPayloadWithDefaults instantiates a new CallCostPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilledDurationSecs

`func (o *CallCostPayload) GetBilledDurationSecs() int32`

GetBilledDurationSecs returns the BilledDurationSecs field if non-nil, zero value otherwise.

### GetBilledDurationSecsOk

`func (o *CallCostPayload) GetBilledDurationSecsOk() (*int32, bool)`

GetBilledDurationSecsOk returns a tuple with the BilledDurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledDurationSecs

`func (o *CallCostPayload) SetBilledDurationSecs(v int32)`

SetBilledDurationSecs sets BilledDurationSecs field to given value.

### HasBilledDurationSecs

`func (o *CallCostPayload) HasBilledDurationSecs() bool`

HasBilledDurationSecs returns a boolean if a field has been set.

### GetCallControlId

`func (o *CallCostPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallCostPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallCostPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallCostPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallCostPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallCostPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallCostPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallCostPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallCostPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallCostPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallCostPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallCostPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallCostPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallCostPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallCostPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallCostPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallCostPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallCostPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallCostPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallCostPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCostParts

`func (o *CallCostPayload) GetCostParts() []CallCostPayloadCostPartsInner`

GetCostParts returns the CostParts field if non-nil, zero value otherwise.

### GetCostPartsOk

`func (o *CallCostPayload) GetCostPartsOk() (*[]CallCostPayloadCostPartsInner, bool)`

GetCostPartsOk returns a tuple with the CostParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostParts

`func (o *CallCostPayload) SetCostParts(v []CallCostPayloadCostPartsInner)`

SetCostParts sets CostParts field to given value.

### HasCostParts

`func (o *CallCostPayload) HasCostParts() bool`

HasCostParts returns a boolean if a field has been set.

### GetTotalCost

`func (o *CallCostPayload) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *CallCostPayload) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *CallCostPayload) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *CallCostPayload) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetStatus

`func (o *CallCostPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallCostPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallCostPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallCostPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


