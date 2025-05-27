# InterruptionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | When true, allows users to interrupt the assistant while speaking | [optional] [default to true]

## Methods

### NewInterruptionSettings

`func NewInterruptionSettings() *InterruptionSettings`

NewInterruptionSettings instantiates a new InterruptionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterruptionSettingsWithDefaults

`func NewInterruptionSettingsWithDefaults() *InterruptionSettings`

NewInterruptionSettingsWithDefaults instantiates a new InterruptionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *InterruptionSettings) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *InterruptionSettings) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *InterruptionSettings) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *InterruptionSettings) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


