# EmergencySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmergencyEnabled** | Pointer to **bool** | Allows you to enable or disable emergency services on the phone number. In order to enable emergency services, you must also set an emergency_address_id. | [optional] [default to false]
**EmergencyAddressId** | Pointer to **string** | Identifies the address to be used with emergency services. | [optional] 
**EmergencyStatus** | Pointer to **string** | Represents the state of the number regarding emergency activation. | [optional] [default to "disabled"]

## Methods

### NewEmergencySettings

`func NewEmergencySettings() *EmergencySettings`

NewEmergencySettings instantiates a new EmergencySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmergencySettingsWithDefaults

`func NewEmergencySettingsWithDefaults() *EmergencySettings`

NewEmergencySettingsWithDefaults instantiates a new EmergencySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmergencyEnabled

`func (o *EmergencySettings) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *EmergencySettings) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *EmergencySettings) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.

### HasEmergencyEnabled

`func (o *EmergencySettings) HasEmergencyEnabled() bool`

HasEmergencyEnabled returns a boolean if a field has been set.

### GetEmergencyAddressId

`func (o *EmergencySettings) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *EmergencySettings) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *EmergencySettings) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *EmergencySettings) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.

### GetEmergencyStatus

`func (o *EmergencySettings) GetEmergencyStatus() string`

GetEmergencyStatus returns the EmergencyStatus field if non-nil, zero value otherwise.

### GetEmergencyStatusOk

`func (o *EmergencySettings) GetEmergencyStatusOk() (*string, bool)`

GetEmergencyStatusOk returns a tuple with the EmergencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyStatus

`func (o *EmergencySettings) SetEmergencyStatus(v string)`

SetEmergencyStatus sets EmergencyStatus field to given value.

### HasEmergencyStatus

`func (o *EmergencySettings) HasEmergencyStatus() bool`

HasEmergencyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


