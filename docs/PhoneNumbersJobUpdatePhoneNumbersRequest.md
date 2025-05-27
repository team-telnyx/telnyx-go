# PhoneNumbersJobUpdatePhoneNumbersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | **[]string** | Array of phone number ids and/or phone numbers in E164 format to update. This parameter is required if no filter parameters are provided. If you want to update specific numbers rather than all numbers matching a filter, you must use this parameter. Each item must be either a valid phone number ID or a phone number in E164 format (e.g., &#39;+13127367254&#39;). | 
**Tags** | Pointer to **[]string** | A list of user-assigned tags to help organize phone numbers. | [optional] 
**ExternalPin** | Pointer to **string** | If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, we will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with the phone number. | [optional] 
**BillingGroupId** | Pointer to **string** | Identifies the billing group associated with the phone number. | [optional] 
**HdVoiceEnabled** | Pointer to **bool** | Indicates whether to enable or disable HD Voice on each phone number. HD Voice is a paid feature and may not be available for all phone numbers, more details about it can be found in the Telnyx support documentation. | [optional] 
**Voice** | Pointer to [**UpdatePhoneNumberVoiceSettingsRequest**](UpdatePhoneNumberVoiceSettingsRequest.md) |  | [optional] 

## Methods

### NewPhoneNumbersJobUpdatePhoneNumbersRequest

`func NewPhoneNumbersJobUpdatePhoneNumbersRequest(phoneNumbers []string, ) *PhoneNumbersJobUpdatePhoneNumbersRequest`

NewPhoneNumbersJobUpdatePhoneNumbersRequest instantiates a new PhoneNumbersJobUpdatePhoneNumbersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumbersJobUpdatePhoneNumbersRequestWithDefaults

`func NewPhoneNumbersJobUpdatePhoneNumbersRequestWithDefaults() *PhoneNumbersJobUpdatePhoneNumbersRequest`

NewPhoneNumbersJobUpdatePhoneNumbersRequestWithDefaults instantiates a new PhoneNumbersJobUpdatePhoneNumbersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetTags

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetExternalPin

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetExternalPin() string`

GetExternalPin returns the ExternalPin field if non-nil, zero value otherwise.

### GetExternalPinOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetExternalPinOk() (*string, bool)`

GetExternalPinOk returns a tuple with the ExternalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPin

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetExternalPin(v string)`

SetExternalPin sets ExternalPin field to given value.

### HasExternalPin

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasExternalPin() bool`

HasExternalPin returns a boolean if a field has been set.

### GetCustomerReference

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetConnectionId

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetHdVoiceEnabled

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetHdVoiceEnabled() bool`

GetHdVoiceEnabled returns the HdVoiceEnabled field if non-nil, zero value otherwise.

### GetHdVoiceEnabledOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetHdVoiceEnabledOk() (*bool, bool)`

GetHdVoiceEnabledOk returns a tuple with the HdVoiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdVoiceEnabled

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetHdVoiceEnabled(v bool)`

SetHdVoiceEnabled sets HdVoiceEnabled field to given value.

### HasHdVoiceEnabled

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasHdVoiceEnabled() bool`

HasHdVoiceEnabled returns a boolean if a field has been set.

### GetVoice

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetVoice() UpdatePhoneNumberVoiceSettingsRequest`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) GetVoiceOk() (*UpdatePhoneNumberVoiceSettingsRequest, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) SetVoice(v UpdatePhoneNumberVoiceSettingsRequest)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *PhoneNumbersJobUpdatePhoneNumbersRequest) HasVoice() bool`

HasVoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


