# CreateVerificationRequestSMS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | +E164 formatted phone number. | 
**VerifyProfileId** | **string** | The identifier of the associated Verify profile. | 
**CustomCode** | Pointer to **NullableString** | Send a self-generated numeric code to the end-user | [optional] 
**TimeoutSecs** | Pointer to **int32** | The number of seconds the verification code is valid for. | [optional] 

## Methods

### NewCreateVerificationRequestSMS

`func NewCreateVerificationRequestSMS(phoneNumber string, verifyProfileId string, ) *CreateVerificationRequestSMS`

NewCreateVerificationRequestSMS instantiates a new CreateVerificationRequestSMS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerificationRequestSMSWithDefaults

`func NewCreateVerificationRequestSMSWithDefaults() *CreateVerificationRequestSMS`

NewCreateVerificationRequestSMSWithDefaults instantiates a new CreateVerificationRequestSMS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CreateVerificationRequestSMS) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateVerificationRequestSMS) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateVerificationRequestSMS) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetVerifyProfileId

`func (o *CreateVerificationRequestSMS) GetVerifyProfileId() string`

GetVerifyProfileId returns the VerifyProfileId field if non-nil, zero value otherwise.

### GetVerifyProfileIdOk

`func (o *CreateVerificationRequestSMS) GetVerifyProfileIdOk() (*string, bool)`

GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProfileId

`func (o *CreateVerificationRequestSMS) SetVerifyProfileId(v string)`

SetVerifyProfileId sets VerifyProfileId field to given value.


### GetCustomCode

`func (o *CreateVerificationRequestSMS) GetCustomCode() string`

GetCustomCode returns the CustomCode field if non-nil, zero value otherwise.

### GetCustomCodeOk

`func (o *CreateVerificationRequestSMS) GetCustomCodeOk() (*string, bool)`

GetCustomCodeOk returns a tuple with the CustomCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCode

`func (o *CreateVerificationRequestSMS) SetCustomCode(v string)`

SetCustomCode sets CustomCode field to given value.

### HasCustomCode

`func (o *CreateVerificationRequestSMS) HasCustomCode() bool`

HasCustomCode returns a boolean if a field has been set.

### SetCustomCodeNil

`func (o *CreateVerificationRequestSMS) SetCustomCodeNil(b bool)`

 SetCustomCodeNil sets the value for CustomCode to be an explicit nil

### UnsetCustomCode
`func (o *CreateVerificationRequestSMS) UnsetCustomCode()`

UnsetCustomCode ensures that no value is present for CustomCode, not even an explicit nil
### GetTimeoutSecs

`func (o *CreateVerificationRequestSMS) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CreateVerificationRequestSMS) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CreateVerificationRequestSMS) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CreateVerificationRequestSMS) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


