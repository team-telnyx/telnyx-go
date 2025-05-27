# RCSComposeRecordingMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | phone number in +E.164 format | 
**Type** | **string** | The type of the recording action. | 

## Methods

### NewRCSComposeRecordingMessage

`func NewRCSComposeRecordingMessage(phoneNumber string, type_ string, ) *RCSComposeRecordingMessage`

NewRCSComposeRecordingMessage instantiates a new RCSComposeRecordingMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSComposeRecordingMessageWithDefaults

`func NewRCSComposeRecordingMessageWithDefaults() *RCSComposeRecordingMessage`

NewRCSComposeRecordingMessageWithDefaults instantiates a new RCSComposeRecordingMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *RCSComposeRecordingMessage) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *RCSComposeRecordingMessage) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *RCSComposeRecordingMessage) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetType

`func (o *RCSComposeRecordingMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RCSComposeRecordingMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RCSComposeRecordingMessage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


