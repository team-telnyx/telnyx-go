# ConversationInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the conversation insight. | 
**Status** | **string** | Current status of the insight generation for the conversation. | 
**CreatedAt** | **time.Time** | Timestamp of when the object was created. | 
**ConversationInsights** | [**[]ConversationInsightConversationInsightsInner**](ConversationInsightConversationInsightsInner.md) | List of insights extracted from the conversation. | 

## Methods

### NewConversationInsight

`func NewConversationInsight(id string, status string, createdAt time.Time, conversationInsights []ConversationInsightConversationInsightsInner, ) *ConversationInsight`

NewConversationInsight instantiates a new ConversationInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationInsightWithDefaults

`func NewConversationInsightWithDefaults() *ConversationInsight`

NewConversationInsightWithDefaults instantiates a new ConversationInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConversationInsight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationInsight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationInsight) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ConversationInsight) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversationInsight) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversationInsight) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ConversationInsight) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationInsight) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationInsight) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetConversationInsights

`func (o *ConversationInsight) GetConversationInsights() []ConversationInsightConversationInsightsInner`

GetConversationInsights returns the ConversationInsights field if non-nil, zero value otherwise.

### GetConversationInsightsOk

`func (o *ConversationInsight) GetConversationInsightsOk() (*[]ConversationInsightConversationInsightsInner, bool)`

GetConversationInsightsOk returns a tuple with the ConversationInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationInsights

`func (o *ConversationInsight) SetConversationInsights(v []ConversationInsightConversationInsightsInner)`

SetConversationInsights sets ConversationInsights field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


