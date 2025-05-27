# CallReferFailedPayload

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

### NewCallReferFailedPayload

`func NewCallReferFailedPayload() *CallReferFailedPayload`

NewCallReferFailedPayload instantiates a new CallReferFailedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallReferFailedPayloadWithDefaults

`func NewCallReferFailedPayloadWithDefaults() *CallReferFailedPayload`

NewCallReferFailedPayloadWithDefaults instantiates a new CallReferFailedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallReferFailedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallReferFailedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallReferFailedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallReferFailedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallReferFailedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallReferFailedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallReferFailedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallReferFailedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallReferFailedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallReferFailedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallReferFailedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallReferFailedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallReferFailedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallReferFailedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallReferFailedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallReferFailedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallReferFailedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallReferFailedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallReferFailedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallReferFailedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFrom

`func (o *CallReferFailedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallReferFailedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallReferFailedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallReferFailedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSipNotifyResponse

`func (o *CallReferFailedPayload) GetSipNotifyResponse() int32`

GetSipNotifyResponse returns the SipNotifyResponse field if non-nil, zero value otherwise.

### GetSipNotifyResponseOk

`func (o *CallReferFailedPayload) GetSipNotifyResponseOk() (*int32, bool)`

GetSipNotifyResponseOk returns a tuple with the SipNotifyResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipNotifyResponse

`func (o *CallReferFailedPayload) SetSipNotifyResponse(v int32)`

SetSipNotifyResponse sets SipNotifyResponse field to given value.

### HasSipNotifyResponse

`func (o *CallReferFailedPayload) HasSipNotifyResponse() bool`

HasSipNotifyResponse returns a boolean if a field has been set.

### GetTo

`func (o *CallReferFailedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallReferFailedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallReferFailedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallReferFailedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


