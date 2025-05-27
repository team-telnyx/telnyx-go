# PostPortabilityCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | Pointer to **[]string** | The list of +E.164 formatted phone numbers to check for portability | [optional] 

## Methods

### NewPostPortabilityCheckRequest

`func NewPostPortabilityCheckRequest() *PostPortabilityCheckRequest`

NewPostPortabilityCheckRequest instantiates a new PostPortabilityCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPortabilityCheckRequestWithDefaults

`func NewPostPortabilityCheckRequestWithDefaults() *PostPortabilityCheckRequest`

NewPostPortabilityCheckRequestWithDefaults instantiates a new PostPortabilityCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *PostPortabilityCheckRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *PostPortabilityCheckRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *PostPortabilityCheckRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *PostPortabilityCheckRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


