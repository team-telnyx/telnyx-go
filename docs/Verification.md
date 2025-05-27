# Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**VerificationType**](VerificationType.md) |  | [optional] 
**RecordType** | Pointer to [**VerificationRecordType**](VerificationRecordType.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** | +E164 formatted phone number. | [optional] 
**VerifyProfileId** | Pointer to **string** | The identifier of the associated Verify profile. | [optional] 
**CustomCode** | Pointer to **NullableString** | Send a self-generated numeric code to the end-user | [optional] 
**TimeoutSecs** | Pointer to **int32** | This is the number of seconds before the code of the request is expired. Once this request has expired, the code will no longer verify the user. Note: this will override the &#x60;default_verification_timeout_secs&#x60; on the Verify profile. | [optional] 
**Status** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewVerification

`func NewVerification() *Verification`

NewVerification instantiates a new Verification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationWithDefaults

`func NewVerificationWithDefaults() *Verification`

NewVerificationWithDefaults instantiates a new Verification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Verification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Verification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Verification) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Verification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Verification) GetType() VerificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Verification) GetTypeOk() (*VerificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Verification) SetType(v VerificationType)`

SetType sets Type field to given value.

### HasType

`func (o *Verification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRecordType

`func (o *Verification) GetRecordType() VerificationRecordType`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Verification) GetRecordTypeOk() (*VerificationRecordType, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Verification) SetRecordType(v VerificationRecordType)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Verification) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Verification) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Verification) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Verification) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Verification) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetVerifyProfileId

`func (o *Verification) GetVerifyProfileId() string`

GetVerifyProfileId returns the VerifyProfileId field if non-nil, zero value otherwise.

### GetVerifyProfileIdOk

`func (o *Verification) GetVerifyProfileIdOk() (*string, bool)`

GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyProfileId

`func (o *Verification) SetVerifyProfileId(v string)`

SetVerifyProfileId sets VerifyProfileId field to given value.

### HasVerifyProfileId

`func (o *Verification) HasVerifyProfileId() bool`

HasVerifyProfileId returns a boolean if a field has been set.

### GetCustomCode

`func (o *Verification) GetCustomCode() string`

GetCustomCode returns the CustomCode field if non-nil, zero value otherwise.

### GetCustomCodeOk

`func (o *Verification) GetCustomCodeOk() (*string, bool)`

GetCustomCodeOk returns a tuple with the CustomCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCode

`func (o *Verification) SetCustomCode(v string)`

SetCustomCode sets CustomCode field to given value.

### HasCustomCode

`func (o *Verification) HasCustomCode() bool`

HasCustomCode returns a boolean if a field has been set.

### SetCustomCodeNil

`func (o *Verification) SetCustomCodeNil(b bool)`

 SetCustomCodeNil sets the value for CustomCode to be an explicit nil

### UnsetCustomCode
`func (o *Verification) UnsetCustomCode()`

UnsetCustomCode ensures that no value is present for CustomCode, not even an explicit nil
### GetTimeoutSecs

`func (o *Verification) GetTimeoutSecs() int32`

GetTimeoutSecs returns the TimeoutSecs field if non-nil, zero value otherwise.

### GetTimeoutSecsOk

`func (o *Verification) GetTimeoutSecsOk() (*int32, bool)`

GetTimeoutSecsOk returns a tuple with the TimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSecs

`func (o *Verification) SetTimeoutSecs(v int32)`

SetTimeoutSecs sets TimeoutSecs field to given value.

### HasTimeoutSecs

`func (o *Verification) HasTimeoutSecs() bool`

HasTimeoutSecs returns a boolean if a field has been set.

### GetStatus

`func (o *Verification) GetStatus() VerificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Verification) GetStatusOk() (*VerificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Verification) SetStatus(v VerificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Verification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Verification) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Verification) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Verification) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Verification) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Verification) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Verification) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Verification) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Verification) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


