# ConversationMessageToolCallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the tool call. | 
**Type** | **string** | Type of the tool call. | 
**Function** | [**ConversationMessageToolCallsInnerFunction**](ConversationMessageToolCallsInnerFunction.md) |  | 

## Methods

### NewConversationMessageToolCallsInner

`func NewConversationMessageToolCallsInner(id string, type_ string, function ConversationMessageToolCallsInnerFunction, ) *ConversationMessageToolCallsInner`

NewConversationMessageToolCallsInner instantiates a new ConversationMessageToolCallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationMessageToolCallsInnerWithDefaults

`func NewConversationMessageToolCallsInnerWithDefaults() *ConversationMessageToolCallsInner`

NewConversationMessageToolCallsInnerWithDefaults instantiates a new ConversationMessageToolCallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConversationMessageToolCallsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationMessageToolCallsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationMessageToolCallsInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ConversationMessageToolCallsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConversationMessageToolCallsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConversationMessageToolCallsInner) SetType(v string)`

SetType sets Type field to given value.


### GetFunction

`func (o *ConversationMessageToolCallsInner) GetFunction() ConversationMessageToolCallsInnerFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ConversationMessageToolCallsInner) GetFunctionOk() (*ConversationMessageToolCallsInnerFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ConversationMessageToolCallsInner) SetFunction(v ConversationMessageToolCallsInnerFunction)`

SetFunction sets Function field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


