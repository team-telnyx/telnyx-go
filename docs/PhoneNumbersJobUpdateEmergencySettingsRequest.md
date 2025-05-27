# PhoneNumbersJobUpdateEmergencySettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | **[]string** |  | 
**EmergencyEnabled** | **bool** | Indicates whether to enable or disable emergency services on the numbers. | 
**EmergencyAddressId** | Pointer to **NullableString** | Identifies the address to be used with emergency services. Required if emergency_enabled is true, must be null or omitted if emergency_enabled is false. | [optional] 

## Methods

### NewPhoneNumbersJobUpdateEmergencySettingsRequest

`func NewPhoneNumbersJobUpdateEmergencySettingsRequest(phoneNumbers []string, emergencyEnabled bool, ) *PhoneNumbersJobUpdateEmergencySettingsRequest`

NewPhoneNumbersJobUpdateEmergencySettingsRequest instantiates a new PhoneNumbersJobUpdateEmergencySettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumbersJobUpdateEmergencySettingsRequestWithDefaults

`func NewPhoneNumbersJobUpdateEmergencySettingsRequestWithDefaults() *PhoneNumbersJobUpdateEmergencySettingsRequest`

NewPhoneNumbersJobUpdateEmergencySettingsRequestWithDefaults instantiates a new PhoneNumbersJobUpdateEmergencySettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetEmergencyEnabled

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.


### GetEmergencyAddressId

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.

### SetEmergencyAddressIdNil

`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) SetEmergencyAddressIdNil(b bool)`

 SetEmergencyAddressIdNil sets the value for EmergencyAddressId to be an explicit nil

### UnsetEmergencyAddressId
`func (o *PhoneNumbersJobUpdateEmergencySettingsRequest) UnsetEmergencyAddressId()`

UnsetEmergencyAddressId ensures that no value is present for EmergencyAddressId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


