# ValidationCodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationCodes** | [**[]ValidationCodesRequestVerificationCodesInner**](ValidationCodesRequestVerificationCodesInner.md) |  | 

## Methods

### NewValidationCodesRequest

`func NewValidationCodesRequest(verificationCodes []ValidationCodesRequestVerificationCodesInner, ) *ValidationCodesRequest`

NewValidationCodesRequest instantiates a new ValidationCodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationCodesRequestWithDefaults

`func NewValidationCodesRequestWithDefaults() *ValidationCodesRequest`

NewValidationCodesRequestWithDefaults instantiates a new ValidationCodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationCodes

`func (o *ValidationCodesRequest) GetVerificationCodes() []ValidationCodesRequestVerificationCodesInner`

GetVerificationCodes returns the VerificationCodes field if non-nil, zero value otherwise.

### GetVerificationCodesOk

`func (o *ValidationCodesRequest) GetVerificationCodesOk() (*[]ValidationCodesRequestVerificationCodesInner, bool)`

GetVerificationCodesOk returns a tuple with the VerificationCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCodes

`func (o *ValidationCodesRequest) SetVerificationCodes(v []ValidationCodesRequestVerificationCodesInner)`

SetVerificationCodes sets VerificationCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


