# CreateVerificationRequestCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | +E164 formatted phone number. | 
**VerifyProfileId** | **string** | The identifier of the associated Verify profile. | 
**CustomCode** | Pointer to **NullableString** | Send a self-generated numeric code to the end-user | [optional] 
**TimeoutSecs** | Pointer to **int32** | The number of seconds the verification code is valid for. | [optional] 

## Methods

### NewCreateVerificationRequestCall

`func NewCreateVerificationRequestCall(phoneNumber string, verifyProfileId string, ) *CreateVerificationRequestCall`

NewCreateVerificationRequestCall instantiates a new CreateVerificationRequestCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerificationRequestCallWithDefaults

`func NewCreateVerificationRequestCallWithDefaults() *CreateVerificationRequestCall`

NewCreateVerificationRequestCallWithDefaults instantiates a new CreateVerificationRequestCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CreateVerificationRequestCall) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateVerificationRequestCall) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateVerificationRequestCall) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetVerifyProfileId

`func (o *CreateVerificationRequestCall) GetVerifyProfileId() string`

GetVerifyProfileId returns the VerifyProfileId field if non-nil, zero value otherwise.

### GetVerifyProfileIdOk

`func (o *CreateVerificationRequestCall) GetVerifyProfileIdOk() (*string, bool)`

GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProfileId

`func (o *CreateVerificationRequestCall) SetVerifyProfileId(v string)`

SetVerifyProfileId sets VerifyProfileId field to given value.


### GetCustomCode

`func (o *CreateVerificationRequestCall) GetCustomCode() string`

GetCustomCode returns the CustomCode field if non-nil, zero value otherwise.

### GetCustomCodeOk

`func (o *CreateVerificationRequestCall) GetCustomCodeOk() (*string, bool)`

GetCustomCodeOk returns a tuple with the CustomCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCode

`func (o *CreateVerificationRequestCall) SetCustomCode(v string)`

SetCustomCode sets CustomCode field to given value.

### HasCustomCode

`func (o *CreateVerificationRequestCall) HasCustomCode() bool`

HasCustomCode returns a boolean if a field has been set.

### SetCustomCodeNil

`func (o *CreateVerificationRequestCall) SetCustomCodeNil(b bool)`

 SetCustomCodeNil sets the value for CustomCode to be an explicit nil

### UnsetCustomCode
`func (o *CreateVerificationRequestCall) UnsetCustomCode()`

UnsetCustomCode ensures that no value is present for CustomCode, not even an explicit nil
### GetTimeoutSecs

`func (o *CreateVerificationRequestCall) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CreateVerificationRequestCall) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CreateVerificationRequestCall) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CreateVerificationRequestCall) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


