# ConversationInsightConversationInsightsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Insight result from the conversation. If the insight has a JSON schema, this will be stringified JSON object. | 
**InsightId** | **string** | Unique identifier for the insight configuration. | 

## Methods

### NewConversationInsightConversationInsightsInner

`func NewConversationInsightConversationInsightsInner(result string, insightId string, ) *ConversationInsightConversationInsightsInner`

NewConversationInsightConversationInsightsInner instantiates a new ConversationInsightConversationInsightsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationInsightConversationInsightsInnerWithDefaults

`func NewConversationInsightConversationInsightsInnerWithDefaults() *ConversationInsightConversationInsightsInner`

NewConversationInsightConversationInsightsInnerWithDefaults instantiates a new ConversationInsightConversationInsightsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ConversationInsightConversationInsightsInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConversationInsightConversationInsightsInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConversationInsightConversationInsightsInner) SetResult(v string)`

SetResult sets Result field to given value.


### GetInsightId

`func (o *ConversationInsightConversationInsightsInner) GetInsightId() string`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *ConversationInsightConversationInsightsInner) GetInsightIdOk() (*string, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *ConversationInsightConversationInsightsInner) SetInsightId(v string)`

SetInsightId sets InsightId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


