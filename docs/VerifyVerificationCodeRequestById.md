# VerifyVerificationCodeRequestById

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | This is the code the user submits for verification. | [optional] 
**Status** | Pointer to **string** | Identifies if the verification code has been accepted or rejected. Only permitted if custom_code was used for the verification. | [optional] 

## Methods

### NewVerifyVerificationCodeRequestById

`func NewVerifyVerificationCodeRequestById() *VerifyVerificationCodeRequestById`

NewVerifyVerificationCodeRequestById instantiates a new VerifyVerificationCodeRequestById object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyVerificationCodeRequestByIdWithDefaults

`func NewVerifyVerificationCodeRequestByIdWithDefaults() *VerifyVerificationCodeRequestById`

NewVerifyVerificationCodeRequestByIdWithDefaults instantiates a new VerifyVerificationCodeRequestById object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *VerifyVerificationCodeRequestById) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VerifyVerificationCodeRequestById) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VerifyVerificationCodeRequestById) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VerifyVerificationCodeRequestById) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *VerifyVerificationCodeRequestById) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyVerificationCodeRequestById) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyVerificationCodeRequestById) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyVerificationCodeRequestById) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


