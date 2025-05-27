# ChatCompletionToolParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Function** | [**FunctionDefinition**](FunctionDefinition.md) |  | 

## Methods

### NewChatCompletionToolParam

`func NewChatCompletionToolParam(type_ string, function FunctionDefinition, ) *ChatCompletionToolParam`

NewChatCompletionToolParam instantiates a new ChatCompletionToolParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionToolParamWithDefaults

`func NewChatCompletionToolParamWithDefaults() *ChatCompletionToolParam`

NewChatCompletionToolParamWithDefaults instantiates a new ChatCompletionToolParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatCompletionToolParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatCompletionToolParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatCompletionToolParam) SetType(v string)`

SetType sets Type field to given value.


### GetFunction

`func (o *ChatCompletionToolParam) GetFunction() FunctionDefinition`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ChatCompletionToolParam) GetFunctionOk() (*FunctionDefinition, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ChatCompletionToolParam) SetFunction(v FunctionDefinition)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


