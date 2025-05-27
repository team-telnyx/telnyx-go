# SMSFallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Phone number in +E.164 format | [optional] 
**Text** | Pointer to **string** | Text (maximum 3072 characters) | [optional] 

## Methods

### NewSMSFallback

`func NewSMSFallback() *SMSFallback`

NewSMSFallback instantiates a new SMSFallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSFallbackWithDefaults

`func NewSMSFallbackWithDefaults() *SMSFallback`

NewSMSFallbackWithDefaults instantiates a new SMSFallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SMSFallback) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SMSFallback) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SMSFallback) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SMSFallback) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetText

`func (o *SMSFallback) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SMSFallback) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SMSFallback) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SMSFallback) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


