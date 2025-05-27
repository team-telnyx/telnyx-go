# CallReferCompletedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Unique ID for controlling the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**From** | Pointer to **string** | Number or SIP URI placing the call. | [optional] 
**SipNotifyResponse** | Pointer to **int32** | SIP NOTIFY event status for tracking the REFER attempt. | [optional] 
**To** | Pointer to **string** | Destination number or SIP URI of the call. | [optional] 

## Methods

### NewCallReferCompletedPayload

`func NewCallReferCompletedPayload() *CallReferCompletedPayload`

NewCallReferCompletedPayload instantiates a new CallReferCompletedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallReferCompletedPayloadWithDefaults

`func NewCallReferCompletedPayloadWithDefaults() *CallReferCompletedPayload`

NewCallReferCompletedPayloadWithDefaults instantiates a new CallReferCompletedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallReferCompletedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallReferCompletedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallReferCompletedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallReferCompletedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallReferCompletedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallReferCompletedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallReferCompletedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallReferCompletedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallReferCompletedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallReferCompletedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallReferCompletedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallReferCompletedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallReferCompletedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallReferCompletedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallReferCompletedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallReferCompletedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallReferCompletedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallReferCompletedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallReferCompletedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallReferCompletedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFrom

`func (o *CallReferCompletedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallReferCompletedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallReferCompletedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallReferCompletedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSipNotifyResponse

`func (o *CallReferCompletedPayload) GetSipNotifyResponse() int32`

GetSipNotifyResponse returns the SipNotifyResponse field if non-nil, zero value otherwise.

### GetSipNotifyResponseOk

`func (o *CallReferCompletedPayload) GetSipNotifyResponseOk() (*int32, bool)`

GetSipNotifyResponseOk returns a tuple with the SipNotifyResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipNotifyResponse

`func (o *CallReferCompletedPayload) SetSipNotifyResponse(v int32)`

SetSipNotifyResponse sets SipNotifyResponse field to given value.

### HasSipNotifyResponse

`func (o *CallReferCompletedPayload) HasSipNotifyResponse() bool`

HasSipNotifyResponse returns a boolean if a field has been set.

### GetTo

`func (o *CallReferCompletedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallReferCompletedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallReferCompletedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallReferCompletedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


