# DialogflowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzeSentiment** | Pointer to **bool** | Enable sentiment analysis from Dialogflow. | [optional] [default to false]
**PartialAutomatedAgentReply** | Pointer to **bool** | Enable partial automated agent reply from Dialogflow. | [optional] [default to false]

## Methods

### NewDialogflowConfig

`func NewDialogflowConfig() *DialogflowConfig`

NewDialogflowConfig instantiates a new DialogflowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDialogflowConfigWithDefaults

`func NewDialogflowConfigWithDefaults() *DialogflowConfig`

NewDialogflowConfigWithDefaults instantiates a new DialogflowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzeSentiment

`func (o *DialogflowConfig) GetAnalyzeSentiment() bool`

GetAnalyzeSentiment returns the AnalyzeSentiment field if non-nil, zero value otherwise.

### GetAnalyzeSentimentOk

`func (o *DialogflowConfig) GetAnalyzeSentimentOk() (*bool, bool)`

GetAnalyzeSentimentOk returns a tuple with the AnalyzeSentiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeSentiment

`func (o *DialogflowConfig) SetAnalyzeSentiment(v bool)`

SetAnalyzeSentiment sets AnalyzeSentiment field to given value.

### HasAnalyzeSentiment

`func (o *DialogflowConfig) HasAnalyzeSentiment() bool`

HasAnalyzeSentiment returns a boolean if a field has been set.

### GetPartialAutomatedAgentReply

`func (o *DialogflowConfig) GetPartialAutomatedAgentReply() bool`

GetPartialAutomatedAgentReply returns the PartialAutomatedAgentReply field if non-nil, zero value otherwise.

### GetPartialAutomatedAgentReplyOk

`func (o *DialogflowConfig) GetPartialAutomatedAgentReplyOk() (*bool, bool)`

GetPartialAutomatedAgentReplyOk returns a tuple with the PartialAutomatedAgentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialAutomatedAgentReply

`func (o *DialogflowConfig) SetPartialAutomatedAgentReply(v bool)`

SetPartialAutomatedAgentReply sets PartialAutomatedAgentReply field to given value.

### HasPartialAutomatedAgentReply

`func (o *DialogflowConfig) HasPartialAutomatedAgentReply() bool`

HasPartialAutomatedAgentReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


