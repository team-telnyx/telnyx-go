# CallConversationInsightsGeneratedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**CallingPartyType** | Pointer to **string** | The type of calling party connection. | [optional] 
**InsightGroupId** | Pointer to **string** | ID that is unique to the insight group being generated for the call. | [optional] 
**Results** | Pointer to [**[]CallConversationInsightsGeneratedPayloadResultsInner**](CallConversationInsightsGeneratedPayloadResultsInner.md) | Array of insight results being generated for the call. | [optional] 

## Methods

### NewCallConversationInsightsGeneratedPayload

`func NewCallConversationInsightsGeneratedPayload() *CallConversationInsightsGeneratedPayload`

NewCallConversationInsightsGeneratedPayload instantiates a new CallConversationInsightsGeneratedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallConversationInsightsGeneratedPayloadWithDefaults

`func NewCallConversationInsightsGeneratedPayloadWithDefaults() *CallConversationInsightsGeneratedPayload`

NewCallConversationInsightsGeneratedPayloadWithDefaults instantiates a new CallConversationInsightsGeneratedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallConversationInsightsGeneratedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallConversationInsightsGeneratedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallConversationInsightsGeneratedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallConversationInsightsGeneratedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallConversationInsightsGeneratedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallConversationInsightsGeneratedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallConversationInsightsGeneratedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallConversationInsightsGeneratedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallConversationInsightsGeneratedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallConversationInsightsGeneratedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallConversationInsightsGeneratedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallConversationInsightsGeneratedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallConversationInsightsGeneratedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallConversationInsightsGeneratedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallConversationInsightsGeneratedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallConversationInsightsGeneratedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallConversationInsightsGeneratedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallConversationInsightsGeneratedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallConversationInsightsGeneratedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallConversationInsightsGeneratedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCallingPartyType

`func (o *CallConversationInsightsGeneratedPayload) GetCallingPartyType() string`

GetCallingPartyType returns the CallingPartyType field if non-nil, zero value otherwise.

### GetCallingPartyTypeOk

`func (o *CallConversationInsightsGeneratedPayload) GetCallingPartyTypeOk() (*string, bool)`

GetCallingPartyTypeOk returns a tuple with the CallingPartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingPartyType

`func (o *CallConversationInsightsGeneratedPayload) SetCallingPartyType(v string)`

SetCallingPartyType sets CallingPartyType field to given value.

### HasCallingPartyType

`func (o *CallConversationInsightsGeneratedPayload) HasCallingPartyType() bool`

HasCallingPartyType returns a boolean if a field has been set.

### GetInsightGroupId

`func (o *CallConversationInsightsGeneratedPayload) GetInsightGroupId() string`

GetInsightGroupId returns the InsightGroupId field if non-nil, zero value otherwise.

### GetInsightGroupIdOk

`func (o *CallConversationInsightsGeneratedPayload) GetInsightGroupIdOk() (*string, bool)`

GetInsightGroupIdOk returns a tuple with the InsightGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightGroupId

`func (o *CallConversationInsightsGeneratedPayload) SetInsightGroupId(v string)`

SetInsightGroupId sets InsightGroupId field to given value.

### HasInsightGroupId

`func (o *CallConversationInsightsGeneratedPayload) HasInsightGroupId() bool`

HasInsightGroupId returns a boolean if a field has been set.

### GetResults

`func (o *CallConversationInsightsGeneratedPayload) GetResults() []CallConversationInsightsGeneratedPayloadResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CallConversationInsightsGeneratedPayload) GetResultsOk() (*[]CallConversationInsightsGeneratedPayloadResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CallConversationInsightsGeneratedPayload) SetResults(v []CallConversationInsightsGeneratedPayloadResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *CallConversationInsightsGeneratedPayload) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


