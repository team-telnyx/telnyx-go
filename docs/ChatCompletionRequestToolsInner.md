# ChatCompletionRequestToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Function** | [**FunctionDefinition**](FunctionDefinition.md) |  | 
**Retrieval** | [**BucketIds**](BucketIds.md) |  | 

## Methods

### NewChatCompletionRequestToolsInner

`func NewChatCompletionRequestToolsInner(type_ string, function FunctionDefinition, retrieval BucketIds, ) *ChatCompletionRequestToolsInner`

NewChatCompletionRequestToolsInner instantiates a new ChatCompletionRequestToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestToolsInnerWithDefaults

`func NewChatCompletionRequestToolsInnerWithDefaults() *ChatCompletionRequestToolsInner`

NewChatCompletionRequestToolsInnerWithDefaults instantiates a new ChatCompletionRequestToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatCompletionRequestToolsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatCompletionRequestToolsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatCompletionRequestToolsInner) SetType(v string)`

SetType sets Type field to given value.


### GetFunction

`func (o *ChatCompletionRequestToolsInner) GetFunction() FunctionDefinition`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ChatCompletionRequestToolsInner) GetFunctionOk() (*FunctionDefinition, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ChatCompletionRequestToolsInner) SetFunction(v FunctionDefinition)`

SetFunction sets Function field to given value.


### GetRetrieval

`func (o *ChatCompletionRequestToolsInner) GetRetrieval() BucketIds`

GetRetrieval returns the Retrieval field if non-nil, zero value otherwise.

### GetRetrievalOk

`func (o *ChatCompletionRequestToolsInner) GetRetrievalOk() (*BucketIds, bool)`

GetRetrievalOk returns a tuple with the Retrieval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieval

`func (o *ChatCompletionRequestToolsInner) SetRetrieval(v BucketIds)`

SetRetrieval sets Retrieval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


