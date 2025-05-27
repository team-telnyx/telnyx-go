# SIMCardRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SimCardGroupId** | Pointer to **string** | The group SIMCardGroup identification. This attribute can be &lt;code&gt;null&lt;/code&gt; when it&#39;s present in an associated resource. | [optional] 
**Tags** | Pointer to **[]string** | Searchable tags associated with the SIM card | [optional] 
**RegistrationCodes** | **[]string** |  | 
**Status** | Pointer to **string** | Status on which the SIM card will be set after being successful registered. | [optional] [default to "enabled"]

## Methods

### NewSIMCardRegistration

`func NewSIMCardRegistration(registrationCodes []string, ) *SIMCardRegistration`

NewSIMCardRegistration instantiates a new SIMCardRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardRegistrationWithDefaults

`func NewSIMCardRegistrationWithDefaults() *SIMCardRegistration`

NewSIMCardRegistrationWithDefaults instantiates a new SIMCardRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSimCardGroupId

`func (o *SIMCardRegistration) GetSimCardGroupId() string`

GetSimCardGroupId returns the SimCardGroupId field if non-nil, zero value otherwise.

### GetSimCardGroupIdOk

`func (o *SIMCardRegistration) GetSimCardGroupIdOk() (*string, bool)`

GetSimCardGroupIdOk returns a tuple with the SimCardGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardGroupId

`func (o *SIMCardRegistration) SetSimCardGroupId(v string)`

SetSimCardGroupId sets SimCardGroupId field to given value.

### HasSimCardGroupId

`func (o *SIMCardRegistration) HasSimCardGroupId() bool`

HasSimCardGroupId returns a boolean if a field has been set.

### GetTags

`func (o *SIMCardRegistration) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SIMCardRegistration) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SIMCardRegistration) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SIMCardRegistration) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRegistrationCodes

`func (o *SIMCardRegistration) GetRegistrationCodes() []string`

GetRegistrationCodes returns the RegistrationCodes field if non-nil, zero value otherwise.

### GetRegistrationCodesOk

`func (o *SIMCardRegistration) GetRegistrationCodesOk() (*[]string, bool)`

GetRegistrationCodesOk returns a tuple with the RegistrationCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationCodes

`func (o *SIMCardRegistration) SetRegistrationCodes(v []string)`

SetRegistrationCodes sets RegistrationCodes field to given value.


### GetStatus

`func (o *SIMCardRegistration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SIMCardRegistration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SIMCardRegistration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SIMCardRegistration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


