# VerifyVerificationCodeResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | +E164 formatted phone number. | 
**ResponseCode** | **string** | Identifies if the verification code has been accepted or rejected. | 

## Methods

### NewVerifyVerificationCodeResponseData

`func NewVerifyVerificationCodeResponseData(phoneNumber string, responseCode string, ) *VerifyVerificationCodeResponseData`

NewVerifyVerificationCodeResponseData instantiates a new VerifyVerificationCodeResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyVerificationCodeResponseDataWithDefaults

`func NewVerifyVerificationCodeResponseDataWithDefaults() *VerifyVerificationCodeResponseData`

NewVerifyVerificationCodeResponseDataWithDefaults instantiates a new VerifyVerificationCodeResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *VerifyVerificationCodeResponseData) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *VerifyVerificationCodeResponseData) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *VerifyVerificationCodeResponseData) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetResponseCode

`func (o *VerifyVerificationCodeResponseData) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *VerifyVerificationCodeResponseData) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *VerifyVerificationCodeResponseData) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


