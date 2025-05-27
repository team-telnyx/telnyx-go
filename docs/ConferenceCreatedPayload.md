# ConferenceCreatedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**ConferenceId** | Pointer to **string** | Conference ID that the participant joined. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 

## Methods

### NewConferenceCreatedPayload

`func NewConferenceCreatedPayload() *ConferenceCreatedPayload`

NewConferenceCreatedPayload instantiates a new ConferenceCreatedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceCreatedPayloadWithDefaults

`func NewConferenceCreatedPayloadWithDefaults() *ConferenceCreatedPayload`

NewConferenceCreatedPayloadWithDefaults instantiates a new ConferenceCreatedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *ConferenceCreatedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *ConferenceCreatedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *ConferenceCreatedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *ConferenceCreatedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *ConferenceCreatedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferenceCreatedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferenceCreatedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferenceCreatedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *ConferenceCreatedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *ConferenceCreatedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *ConferenceCreatedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *ConferenceCreatedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *ConferenceCreatedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ConferenceCreatedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ConferenceCreatedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *ConferenceCreatedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *ConferenceCreatedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ConferenceCreatedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ConferenceCreatedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *ConferenceCreatedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceCreatedPayload) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceCreatedPayload) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceCreatedPayload) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceCreatedPayload) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *ConferenceCreatedPayload) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *ConferenceCreatedPayload) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *ConferenceCreatedPayload) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *ConferenceCreatedPayload) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


