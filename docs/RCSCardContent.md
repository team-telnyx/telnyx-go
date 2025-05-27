# RCSCardContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the card (at most 200 characters) | [optional] 
**Description** | Pointer to **string** | Description of the card (at most 2000 characters) | [optional] 
**Media** | Pointer to [**RCSMedia**](RCSMedia.md) |  | [optional] 
**Suggestions** | Pointer to [**[]RCSSuggestion**](RCSSuggestion.md) | List of suggestions to include in the card. Maximum 10 suggestions. | [optional] 

## Methods

### NewRCSCardContent

`func NewRCSCardContent() *RCSCardContent`

NewRCSCardContent instantiates a new RCSCardContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSCardContentWithDefaults

`func NewRCSCardContentWithDefaults() *RCSCardContent`

NewRCSCardContentWithDefaults instantiates a new RCSCardContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RCSCardContent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RCSCardContent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RCSCardContent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RCSCardContent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *RCSCardContent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RCSCardContent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RCSCardContent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RCSCardContent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMedia

`func (o *RCSCardContent) GetMedia() RCSMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *RCSCardContent) GetMediaOk() (*RCSMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *RCSCardContent) SetMedia(v RCSMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *RCSCardContent) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetSuggestions

`func (o *RCSCardContent) GetSuggestions() []RCSSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *RCSCardContent) GetSuggestionsOk() (*[]RCSSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *RCSCardContent) SetSuggestions(v []RCSSuggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *RCSCardContent) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


