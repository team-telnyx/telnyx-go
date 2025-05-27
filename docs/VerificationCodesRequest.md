# VerificationCodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | **[]string** |  | 
**VerificationMethod** | **string** |  | 

## Methods

### NewVerificationCodesRequest

`func NewVerificationCodesRequest(phoneNumbers []string, verificationMethod string, ) *VerificationCodesRequest`

NewVerificationCodesRequest instantiates a new VerificationCodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationCodesRequestWithDefaults

`func NewVerificationCodesRequestWithDefaults() *VerificationCodesRequest`

NewVerificationCodesRequestWithDefaults instantiates a new VerificationCodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *VerificationCodesRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *VerificationCodesRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *VerificationCodesRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetVerificationMethod

`func (o *VerificationCodesRequest) GetVerificationMethod() string`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *VerificationCodesRequest) GetVerificationMethodOk() (*string, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *VerificationCodesRequest) SetVerificationMethod(v string)`

SetVerificationMethod sets VerificationMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


