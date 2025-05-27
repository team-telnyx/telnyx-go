# ConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The role of the message sender. | 
**Text** | **string** | The message content. Can be null for tool calls. | 
**ToolCalls** | Pointer to [**[]ConversationMessageToolCallsInner**](ConversationMessageToolCallsInner.md) | Optional tool calls made by the assistant. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The datetime the message was created on the conversation. This does not necesarily correspond to the time the message was sent. The best field to use to determine the time the end user experienced the message is &#x60;sent_at&#x60;. | [optional] 
**SentAt** | Pointer to **time.Time** | The datetime the message was sent to the end user. | [optional] 

## Methods

### NewConversationMessage

`func NewConversationMessage(role string, text string, ) *ConversationMessage`

NewConversationMessage instantiates a new ConversationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationMessageWithDefaults

`func NewConversationMessageWithDefaults() *ConversationMessage`

NewConversationMessageWithDefaults instantiates a new ConversationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ConversationMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConversationMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConversationMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetText

`func (o *ConversationMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConversationMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConversationMessage) SetText(v string)`

SetText sets Text field to given value.


### GetToolCalls

`func (o *ConversationMessage) GetToolCalls() []ConversationMessageToolCallsInner`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *ConversationMessage) GetToolCallsOk() (*[]ConversationMessageToolCallsInner, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *ConversationMessage) SetToolCalls(v []ConversationMessageToolCallsInner)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *ConversationMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConversationMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConversationMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSentAt

`func (o *ConversationMessage) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *ConversationMessage) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *ConversationMessage) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *ConversationMessage) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


