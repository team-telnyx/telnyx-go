# CreateVerificationRequestFlashcall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | +E164 formatted phone number. | 
**VerifyProfileId** | **string** | The identifier of the associated Verify profile. | 
**TimeoutSecs** | Pointer to **int32** | The number of seconds the verification code is valid for. | [optional] 

## Methods

### NewCreateVerificationRequestFlashcall

`func NewCreateVerificationRequestFlashcall(phoneNumber string, verifyProfileId string, ) *CreateVerificationRequestFlashcall`

NewCreateVerificationRequestFlashcall instantiates a new CreateVerificationRequestFlashcall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVerificationRequestFlashcallWithDefaults

`func NewCreateVerificationRequestFlashcallWithDefaults() *CreateVerificationRequestFlashcall`

NewCreateVerificationRequestFlashcallWithDefaults instantiates a new CreateVerificationRequestFlashcall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CreateVerificationRequestFlashcall) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateVerificationRequestFlashcall) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateVerificationRequestFlashcall) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetVerifyProfileId

`func (o *CreateVerificationRequestFlashcall) GetVerifyProfileId() string`

GetVerifyProfileId returns the VerifyProfileId field if non-nil, zero value otherwise.

### GetVerifyProfileIdOk

`func (o *CreateVerificationRequestFlashcall) GetVerifyProfileIdOk() (*string, bool)`

GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProfileId

`func (o *CreateVerificationRequestFlashcall) SetVerifyProfileId(v string)`

SetVerifyProfileId sets VerifyProfileId field to given value.


### GetTimeoutSecs

`func (o *CreateVerificationRequestFlashcall) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *CreateVerificationRequestFlashcall) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *CreateVerificationRequestFlashcall) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *CreateVerificationRequestFlashcall) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


