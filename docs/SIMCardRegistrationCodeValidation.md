# SIMCardRegistrationCodeValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] 
**RegistrationCode** | Pointer to **string** | The 10-digit SIM card registration code | [optional] 
**Valid** | Pointer to **bool** | The attribute that denotes whether the code is valid or not | [optional] 
**InvalidDetail** | Pointer to **NullableString** | The validation message | [optional] 

## Methods

### NewSIMCardRegistrationCodeValidation

`func NewSIMCardRegistrationCodeValidation() *SIMCardRegistrationCodeValidation`

NewSIMCardRegistrationCodeValidation instantiates a new SIMCardRegistrationCodeValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardRegistrationCodeValidationWithDefaults

`func NewSIMCardRegistrationCodeValidationWithDefaults() *SIMCardRegistrationCodeValidation`

NewSIMCardRegistrationCodeValidationWithDefaults instantiates a new SIMCardRegistrationCodeValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *SIMCardRegistrationCodeValidation) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardRegistrationCodeValidation) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardRegistrationCodeValidation) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardRegistrationCodeValidation) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRegistrationCode

`func (o *SIMCardRegistrationCodeValidation) GetRegistrationCode() string`

GetRegistrationCode returns the RegistrationCode field if non-nil, zero value otherwise.

### GetRegistrationCodeOk

`func (o *SIMCardRegistrationCodeValidation) GetRegistrationCodeOk() (*string, bool)`

GetRegistrationCodeOk returns a tuple with the RegistrationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationCode

`func (o *SIMCardRegistrationCodeValidation) SetRegistrationCode(v string)`

SetRegistrationCode sets RegistrationCode field to given value.

### HasRegistrationCode

`func (o *SIMCardRegistrationCodeValidation) HasRegistrationCode() bool`

HasRegistrationCode returns a boolean if a field has been set.

### GetValid

`func (o *SIMCardRegistrationCodeValidation) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SIMCardRegistrationCodeValidation) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SIMCardRegistrationCodeValidation) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *SIMCardRegistrationCodeValidation) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetInvalidDetail

`func (o *SIMCardRegistrationCodeValidation) GetInvalidDetail() string`

GetInvalidDetail returns the InvalidDetail field if non-nil, zero value otherwise.

### GetInvalidDetailOk

`func (o *SIMCardRegistrationCodeValidation) GetInvalidDetailOk() (*string, bool)`

GetInvalidDetailOk returns a tuple with the InvalidDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidDetail

`func (o *SIMCardRegistrationCodeValidation) SetInvalidDetail(v string)`

SetInvalidDetail sets InvalidDetail field to given value.

### HasInvalidDetail

`func (o *SIMCardRegistrationCodeValidation) HasInvalidDetail() bool`

HasInvalidDetail returns a boolean if a field has been set.

### SetInvalidDetailNil

`func (o *SIMCardRegistrationCodeValidation) SetInvalidDetailNil(b bool)`

 SetInvalidDetailNil sets the value for InvalidDetail to be an explicit nil

### UnsetInvalidDetail
`func (o *SIMCardRegistrationCodeValidation) UnsetInvalidDetail()`

UnsetInvalidDetail ensures that no value is present for InvalidDetail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


