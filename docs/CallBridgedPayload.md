# CallBridgedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**From** | Pointer to **string** | Number or SIP URI placing the call. | [optional] 
**To** | Pointer to **string** | Destination number or SIP URI of the call. | [optional] 

## Methods

### NewCallBridgedPayload

`func NewCallBridgedPayload() *CallBridgedPayload`

NewCallBridgedPayload instantiates a new CallBridgedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallBridgedPayloadWithDefaults

`func NewCallBridgedPayloadWithDefaults() *CallBridgedPayload`

NewCallBridgedPayloadWithDefaults instantiates a new CallBridgedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallBridgedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallBridgedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallBridgedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallBridgedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallBridgedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallBridgedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallBridgedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallBridgedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallBridgedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallBridgedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallBridgedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallBridgedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallBridgedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallBridgedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallBridgedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallBridgedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallBridgedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallBridgedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallBridgedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallBridgedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFrom

`func (o *CallBridgedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallBridgedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallBridgedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallBridgedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallBridgedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallBridgedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallBridgedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallBridgedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


