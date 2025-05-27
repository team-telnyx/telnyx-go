# SendPortingVerificationCodesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | Pointer to **[]string** |  | [optional] 
**VerificationMethod** | Pointer to **string** |  | [optional] 

## Methods

### NewSendPortingVerificationCodesRequest

`func NewSendPortingVerificationCodesRequest() *SendPortingVerificationCodesRequest`

NewSendPortingVerificationCodesRequest instantiates a new SendPortingVerificationCodesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPortingVerificationCodesRequestWithDefaults

`func NewSendPortingVerificationCodesRequestWithDefaults() *SendPortingVerificationCodesRequest`

NewSendPortingVerificationCodesRequestWithDefaults instantiates a new SendPortingVerificationCodesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *SendPortingVerificationCodesRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *SendPortingVerificationCodesRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *SendPortingVerificationCodesRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *SendPortingVerificationCodesRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetVerificationMethod

`func (o *SendPortingVerificationCodesRequest) GetVerificationMethod() string`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *SendPortingVerificationCodesRequest) GetVerificationMethodOk() (*string, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *SendPortingVerificationCodesRequest) SetVerificationMethod(v string)`

SetVerificationMethod sets VerificationMethod field to given value.

### HasVerificationMethod

`func (o *SendPortingVerificationCodesRequest) HasVerificationMethod() bool`

HasVerificationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


