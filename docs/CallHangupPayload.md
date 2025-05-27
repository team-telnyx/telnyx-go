# CallHangupPayload

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
**HangupCause** | Pointer to **string** | The reason the call was ended (&#x60;call_rejected&#x60;, &#x60;normal_clearing&#x60;, &#x60;originator_cancel&#x60;, &#x60;timeout&#x60;, &#x60;time_limit&#x60;, &#x60;user_busy&#x60;, &#x60;not_found&#x60; or &#x60;unspecified&#x60;). | [optional] 
**HangupSource** | Pointer to **string** | The party who ended the call (&#x60;callee&#x60;, &#x60;caller&#x60;, &#x60;unknown&#x60;). | [optional] 
**SipHangupCause** | Pointer to **string** | The reason the call was ended (SIP response code). If the SIP response is unavailable (in inbound calls for example) this is set to &#x60;unspecified&#x60;. | [optional] 

## Methods

### NewCallHangupPayload

`func NewCallHangupPayload() *CallHangupPayload`

NewCallHangupPayload instantiates a new CallHangupPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallHangupPayloadWithDefaults

`func NewCallHangupPayloadWithDefaults() *CallHangupPayload`

NewCallHangupPayloadWithDefaults instantiates a new CallHangupPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallHangupPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallHangupPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallHangupPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallHangupPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallHangupPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallHangupPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallHangupPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallHangupPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallHangupPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallHangupPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallHangupPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallHangupPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallHangupPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallHangupPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallHangupPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallHangupPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallHangupPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallHangupPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallHangupPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallHangupPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *CallHangupPayload) GetCustomHeaders() []CustomSipHeader`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CallHangupPayload) GetCustomHeadersOk() (*[]CustomSipHeader, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CallHangupPayload) SetCustomHeaders(v []CustomSipHeader)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CallHangupPayload) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.

### GetSipHeaders

`func (o *CallHangupPayload) GetSipHeaders() []SipHeader`

GetSipHeaders returns the SipHeaders field if non-nil, zero value otherwise.

### GetSipHeadersOk

`func (o *CallHangupPayload) GetSipHeadersOk() (*[]SipHeader, bool)`

GetSipHeadersOk returns a tuple with the SipHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHeaders

`func (o *CallHangupPayload) SetSipHeaders(v []SipHeader)`

SetSipHeaders sets SipHeaders field to given value.

### HasSipHeaders

`func (o *CallHangupPayload) HasSipHeaders() bool`

HasSipHeaders returns a boolean if a field has been set.

### GetFrom

`func (o *CallHangupPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallHangupPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallHangupPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallHangupPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallHangupPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallHangupPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallHangupPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallHangupPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetStartTime

`func (o *CallHangupPayload) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallHangupPayload) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallHangupPayload) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallHangupPayload) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetState

`func (o *CallHangupPayload) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CallHangupPayload) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CallHangupPayload) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CallHangupPayload) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *CallHangupPayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CallHangupPayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CallHangupPayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CallHangupPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHangupCause

`func (o *CallHangupPayload) GetHangupCause() string`

GetHangupCause returns the HangupCause field if non-nil, zero value otherwise.

### GetHangupCauseOk

`func (o *CallHangupPayload) GetHangupCauseOk() (*string, bool)`

GetHangupCauseOk returns a tuple with the HangupCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangupCause

`func (o *CallHangupPayload) SetHangupCause(v string)`

SetHangupCause sets HangupCause field to given value.

### HasHangupCause

`func (o *CallHangupPayload) HasHangupCause() bool`

HasHangupCause returns a boolean if a field has been set.

### GetHangupSource

`func (o *CallHangupPayload) GetHangupSource() string`

GetHangupSource returns the HangupSource field if non-nil, zero value otherwise.

### GetHangupSourceOk

`func (o *CallHangupPayload) GetHangupSourceOk() (*string, bool)`

GetHangupSourceOk returns a tuple with the HangupSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangupSource

`func (o *CallHangupPayload) SetHangupSource(v string)`

SetHangupSource sets HangupSource field to given value.

### HasHangupSource

`func (o *CallHangupPayload) HasHangupSource() bool`

HasHangupSource returns a boolean if a field has been set.

### GetSipHangupCause

`func (o *CallHangupPayload) GetSipHangupCause() string`

GetSipHangupCause returns the SipHangupCause field if non-nil, zero value otherwise.

### GetSipHangupCauseOk

`func (o *CallHangupPayload) GetSipHangupCauseOk() (*string, bool)`

GetSipHangupCauseOk returns a tuple with the SipHangupCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipHangupCause

`func (o *CallHangupPayload) SetSipHangupCause(v string)`

SetSipHangupCause sets SipHangupCause field to given value.

### HasSipHangupCause

`func (o *CallHangupPayload) HasSipHangupCause() bool`

HasSipHangupCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


