# RCSSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reply** | Pointer to [**RCSReply**](RCSReply.md) |  | [optional] 
**Action** | Pointer to [**RCSAction**](RCSAction.md) |  | [optional] 

## Methods

### NewRCSSuggestion

`func NewRCSSuggestion() *RCSSuggestion`

NewRCSSuggestion instantiates a new RCSSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSSuggestionWithDefaults

`func NewRCSSuggestionWithDefaults() *RCSSuggestion`

NewRCSSuggestionWithDefaults instantiates a new RCSSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReply

`func (o *RCSSuggestion) GetReply() RCSReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *RCSSuggestion) GetReplyOk() (*RCSReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *RCSSuggestion) SetReply(v RCSReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *RCSSuggestion) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetAction

`func (o *RCSSuggestion) GetAction() RCSAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RCSSuggestion) GetActionOk() (*RCSAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RCSSuggestion) SetAction(v RCSAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *RCSSuggestion) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


