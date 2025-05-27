# RCSComposeTextMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | phone number in +E.164 format | 
**Text** | Pointer to **string** | Draft to go into the send message text field. | [optional] 

## Methods

### NewRCSComposeTextMessage

`func NewRCSComposeTextMessage(phoneNumber string, ) *RCSComposeTextMessage`

NewRCSComposeTextMessage instantiates a new RCSComposeTextMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSComposeTextMessageWithDefaults

`func NewRCSComposeTextMessageWithDefaults() *RCSComposeTextMessage`

NewRCSComposeTextMessageWithDefaults instantiates a new RCSComposeTextMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *RCSComposeTextMessage) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RCSComposeTextMessage) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RCSComposeTextMessage) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetText

`func (o *RCSComposeTextMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RCSComposeTextMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RCSComposeTextMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RCSComposeTextMessage) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


