# PhoneNumberEnableEmergencyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmergencyEnabled** | **bool** | Indicates whether to enable emergency services on this number. | 
**EmergencyAddressId** | **string** | Identifies the address to be used with emergency services. | 

## Methods

### NewPhoneNumberEnableEmergencyRequest

`func NewPhoneNumberEnableEmergencyRequest(emergencyEnabled bool, emergencyAddressId string, ) *PhoneNumberEnableEmergencyRequest`

NewPhoneNumberEnableEmergencyRequest instantiates a new PhoneNumberEnableEmergencyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberEnableEmergencyRequestWithDefaults

`func NewPhoneNumberEnableEmergencyRequestWithDefaults() *PhoneNumberEnableEmergencyRequest`

NewPhoneNumberEnableEmergencyRequestWithDefaults instantiates a new PhoneNumberEnableEmergencyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmergencyEnabled

`func (o *PhoneNumberEnableEmergencyRequest) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *PhoneNumberEnableEmergencyRequest) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *PhoneNumberEnableEmergencyRequest) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.


### GetEmergencyAddressId

`func (o *PhoneNumberEnableEmergencyRequest) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *PhoneNumberEnableEmergencyRequest) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *PhoneNumberEnableEmergencyRequest) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


