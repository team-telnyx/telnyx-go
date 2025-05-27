# CreateVerifiedNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** |  | 
**VerificationMethod** | **string** | Verification method. | 

## Methods

### NewCreateVerifiedNumberRequest

`func NewCreateVerifiedNumberRequest(phoneNumber string, verificationMethod string, ) *CreateVerifiedNumberRequest`

NewCreateVerifiedNumberRequest instantiates a new CreateVerifiedNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerifiedNumberRequestWithDefaults

`func NewCreateVerifiedNumberRequestWithDefaults() *CreateVerifiedNumberRequest`

NewCreateVerifiedNumberRequestWithDefaults instantiates a new CreateVerifiedNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CreateVerifiedNumberRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateVerifiedNumberRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateVerifiedNumberRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetVerificationMethod

`func (o *CreateVerifiedNumberRequest) GetVerificationMethod() string`

GetVerificationMethod returns the VerificationMethod field if non-nil, zero value otherwise.

### GetVerificationMethodOk

`func (o *CreateVerifiedNumberRequest) GetVerificationMethodOk() (*string, bool)`

GetVerificationMethodOk returns a tuple with the VerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationMethod

`func (o *CreateVerifiedNumberRequest) SetVerificationMethod(v string)`

SetVerificationMethod sets VerificationMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


