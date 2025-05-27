# CallAnsweredPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**CustomHeaders** | Pointer to [**[]CustomSipHeader**](CustomSipHeader.md) | Custom headers set on answer command | [optional] 
**SipHeaders** | Pointer to [**[]SipHeader**](SipHeader.md) | User-to-User and Diversion headers from sip invite. | [optional] 
**From** | Pointer to **string** | Number or SIP URI placing the call. | [optional] 
**To** | Pointer to **string** | Destination number or SIP URI of the call. | [optional] 
**StartTime** | Pointer to **time.Time** | ISO 8601 datetime of when the call started. | [optional] 
**State** | Pointer to **string** | State received from a command. | [optional] 
**Tags** | Pointer to **[]string** | Array of tags associated to number. | [optional] 

## Methods

### NewCallAnsweredPayload

`func NewCallAnsweredPayload() *CallAnsweredPayload`

NewCallAnsweredPayload instantiates a new CallAnsweredPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallAnsweredPayloadWithDefaults

`func NewCallAnsweredPayloadWithDefaults() *CallAnsweredPayload`

NewCallAnsweredPayloadWithDefaults instantiates a new CallAnsweredPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallAnsweredPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallAnsweredPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallAnsweredPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallAnsweredPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallAnsweredPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallAnsweredPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallAnsweredPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallAnsweredPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallAnsweredPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallAnsweredPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallAnsweredPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallAnsweredPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallAnsweredPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallAnsweredPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallAnsweredPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallAnsweredPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallAnsweredPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallAnsweredPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallAnsweredPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallAnsweredPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *CallAnsweredPayload) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CallAnsweredPayload) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CallAnsweredPayload) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CallAnsweredPayload) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetSipHeaders

`func (o *CallAnsweredPayload) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *CallAnsweredPayload) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *CallAnsweredPayload) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *CallAnsweredPayload) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.

### GetFrom

`func (o *CallAnsweredPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallAnsweredPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallAnsweredPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallAnsweredPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallAnsweredPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallAnsweredPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallAnsweredPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallAnsweredPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetStartTime

`func (o *CallAnsweredPayload) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallAnsweredPayload) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallAnsweredPayload) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallAnsweredPayload) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *CallAnsweredPayload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CallAnsweredPayload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CallAnsweredPayload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CallAnsweredPayload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *CallAnsweredPayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CallAnsweredPayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CallAnsweredPayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CallAnsweredPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


