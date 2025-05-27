# RCSContentMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggestions** | Pointer to [**[]RCSSuggestion**](RCSSuggestion.md) | List of suggested actions and replies | [optional] 
**Text** | Pointer to **string** | Text (maximum 3072 characters) | [optional] 
**RichCard** | Pointer to [**RCSRichCard**](RCSRichCard.md) |  | [optional] 
**ContentInfo** | Pointer to [**RCSContentInfo**](RCSContentInfo.md) |  | [optional] 

## Methods

### NewRCSContentMessage

`func NewRCSContentMessage() *RCSContentMessage`

NewRCSContentMessage instantiates a new RCSContentMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSContentMessageWithDefaults

`func NewRCSContentMessageWithDefaults() *RCSContentMessage`

NewRCSContentMessageWithDefaults instantiates a new RCSContentMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestions

`func (o *RCSContentMessage) GetSuggestions() []RCSSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *RCSContentMessage) GetSuggestionsOk() (*[]RCSSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *RCSContentMessage) SetSuggestions(v []RCSSuggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *RCSContentMessage) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetText

`func (o *RCSContentMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RCSContentMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RCSContentMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RCSContentMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetRichCard

`func (o *RCSContentMessage) GetRichCard() RCSRichCard`

GetRichCard returns the RichCard field if non-nil, zero value otherwise.

### GetRichCardOk

`func (o *RCSContentMessage) GetRichCardOk() (*RCSRichCard, bool)`

GetRichCardOk returns a tuple with the RichCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichCard

`func (o *RCSContentMessage) SetRichCard(v RCSRichCard)`

SetRichCard sets RichCard field to given value.

### HasRichCard

`func (o *RCSContentMessage) HasRichCard() bool`

HasRichCard returns a boolean if a field has been set.

### GetContentInfo

`func (o *RCSContentMessage) GetContentInfo() RCSContentInfo`

GetContentInfo returns the ContentInfo field if non-nil, zero value otherwise.

### GetContentInfoOk

`func (o *RCSContentMessage) GetContentInfoOk() (*RCSContentInfo, bool)`

GetContentInfoOk returns a tuple with the ContentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentInfo

`func (o *RCSContentMessage) SetContentInfo(v RCSContentInfo)`

SetContentInfo sets ContentInfo field to given value.

### HasContentInfo

`func (o *RCSContentMessage) HasContentInfo() bool`

HasContentInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


