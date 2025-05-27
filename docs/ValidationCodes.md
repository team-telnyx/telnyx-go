# ValidationCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | [**[]ValidationCodesPhoneNumbersInner**](ValidationCodesPhoneNumbersInner.md) |  | 
**OrderId** | **string** |  | 

## Methods

### NewValidationCodes

`func NewValidationCodes(phoneNumbers []ValidationCodesPhoneNumbersInner, orderId string, ) *ValidationCodes`

NewValidationCodes instantiates a new ValidationCodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationCodesWithDefaults

`func NewValidationCodesWithDefaults() *ValidationCodes`

NewValidationCodesWithDefaults instantiates a new ValidationCodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *ValidationCodes) GetPhoneNumbers() []ValidationCodesPhoneNumbersInner`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *ValidationCodes) GetPhoneNumbersOk() (*[]ValidationCodesPhoneNumbersInner, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *ValidationCodes) SetPhoneNumbers(v []ValidationCodesPhoneNumbersInner)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetOrderId

`func (o *ValidationCodes) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ValidationCodes) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ValidationCodes) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


