# RCSReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Text that is shown in the suggested reply (maximum 25 characters) | [optional] 
**PostbackData** | Pointer to **string** | Payload (base64 encoded) that will be sent to the agent in the user event that results when the user taps the suggested action. Maximum 2048 characters. | [optional] 

## Methods

### NewRCSReply

`func NewRCSReply() *RCSReply`

NewRCSReply instantiates a new RCSReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSReplyWithDefaults

`func NewRCSReplyWithDefaults() *RCSReply`

NewRCSReplyWithDefaults instantiates a new RCSReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *RCSReply) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RCSReply) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RCSReply) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RCSReply) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPostbackData

`func (o *RCSReply) GetPostbackData() string`

GetPostbackData returns the PostbackData field if non-nil, zero value otherwise.

### GetPostbackDataOk

`func (o *RCSReply) GetPostbackDataOk() (*string, bool)`

GetPostbackDataOk returns a tuple with the PostbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackData

`func (o *RCSReply) SetPostbackData(v string)`

SetPostbackData sets PostbackData field to given value.

### HasPostbackData

`func (o *RCSReply) HasPostbackData() bool`

HasPostbackData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


