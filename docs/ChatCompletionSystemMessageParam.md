# ChatCompletionSystemMessageParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**ChatCompletionSystemMessageParamContent**](ChatCompletionSystemMessageParamContent.md) |  | 
**Role** | **string** |  | 

## Methods

### NewChatCompletionSystemMessageParam

`func NewChatCompletionSystemMessageParam(content ChatCompletionSystemMessageParamContent, role string, ) *ChatCompletionSystemMessageParam`

NewChatCompletionSystemMessageParam instantiates a new ChatCompletionSystemMessageParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionSystemMessageParamWithDefaults

`func NewChatCompletionSystemMessageParamWithDefaults() *ChatCompletionSystemMessageParam`

NewChatCompletionSystemMessageParamWithDefaults instantiates a new ChatCompletionSystemMessageParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ChatCompletionSystemMessageParam) GetContent() ChatCompletionSystemMessageParamContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatCompletionSystemMessageParam) GetContentOk() (*ChatCompletionSystemMessageParamContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatCompletionSystemMessageParam) SetContent(v ChatCompletionSystemMessageParamContent)`

SetContent sets Content field to given value.


### GetRole

`func (o *ChatCompletionSystemMessageParam) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatCompletionSystemMessageParam) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatCompletionSystemMessageParam) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


