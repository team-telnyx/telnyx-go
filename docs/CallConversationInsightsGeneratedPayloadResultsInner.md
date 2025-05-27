# CallConversationInsightsGeneratedPayloadResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsightId** | Pointer to **string** | ID that is unique to the insight result being generated for the call. | [optional] 
**Result** | Pointer to [**CallConversationInsightsGeneratedPayloadResultsInnerResult**](CallConversationInsightsGeneratedPayloadResultsInnerResult.md) |  | [optional] 

## Methods

### NewCallConversationInsightsGeneratedPayloadResultsInner

`func NewCallConversationInsightsGeneratedPayloadResultsInner() *CallConversationInsightsGeneratedPayloadResultsInner`

NewCallConversationInsightsGeneratedPayloadResultsInner instantiates a new CallConversationInsightsGeneratedPayloadResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallConversationInsightsGeneratedPayloadResultsInnerWithDefaults

`func NewCallConversationInsightsGeneratedPayloadResultsInnerWithDefaults() *CallConversationInsightsGeneratedPayloadResultsInner`

NewCallConversationInsightsGeneratedPayloadResultsInnerWithDefaults instantiates a new CallConversationInsightsGeneratedPayloadResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsightId

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) GetInsightId() string`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) GetInsightIdOk() (*string, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) SetInsightId(v string)`

SetInsightId sets InsightId field to given value.

### HasInsightId

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) HasInsightId() bool`

HasInsightId returns a boolean if a field has been set.

### GetResult

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) GetResult() CallConversationInsightsGeneratedPayloadResultsInnerResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) GetResultOk() (*CallConversationInsightsGeneratedPayloadResultsInnerResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) SetResult(v CallConversationInsightsGeneratedPayloadResultsInnerResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CallConversationInsightsGeneratedPayloadResultsInner) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


