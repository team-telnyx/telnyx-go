# ConversationMessageListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ConversationMessage**](ConversationMessage.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewConversationMessageListData

`func NewConversationMessageListData(data []ConversationMessage, meta Meta, ) *ConversationMessageListData`

NewConversationMessageListData instantiates a new ConversationMessageListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationMessageListDataWithDefaults

`func NewConversationMessageListDataWithDefaults() *ConversationMessageListData`

NewConversationMessageListDataWithDefaults instantiates a new ConversationMessageListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConversationMessageListData) GetData() []ConversationMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConversationMessageListData) GetDataOk() (*[]ConversationMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConversationMessageListData) SetData(v []ConversationMessage)`

SetData sets Data field to given value.


### GetMeta

`func (o *ConversationMessageListData) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ConversationMessageListData) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ConversationMessageListData) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


