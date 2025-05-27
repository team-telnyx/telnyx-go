# UpdatePhoneNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | A list of user-assigned tags to help organize phone numbers. | [optional] 
**ExternalPin** | Pointer to **string** | If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, we will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism. | [optional] 
**HdVoiceEnabled** | Pointer to **bool** | Indicates whether HD voice is enabled for this number. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with the phone number. | [optional] 
**BillingGroupId** | Pointer to **string** | Identifies the billing group associated with the phone number. | [optional] 

## Methods

### NewUpdatePhoneNumberRequest

`func NewUpdatePhoneNumberRequest() *UpdatePhoneNumberRequest`

NewUpdatePhoneNumberRequest instantiates a new UpdatePhoneNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePhoneNumberRequestWithDefaults

`func NewUpdatePhoneNumberRequestWithDefaults() *UpdatePhoneNumberRequest`

NewUpdatePhoneNumberRequestWithDefaults instantiates a new UpdatePhoneNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePhoneNumberRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePhoneNumberRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePhoneNumberRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatePhoneNumberRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePhoneNumberRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePhoneNumberRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePhoneNumberRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePhoneNumberRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetExternalPin

`func (o *UpdatePhoneNumberRequest) GetExternalPin() string`

GetExternalPin returns the ExternalPin field if non-nil, zero value otherwise.

### GetExternalPinOk

`func (o *UpdatePhoneNumberRequest) GetExternalPinOk() (*string, bool)`

GetExternalPinOk returns a tuple with the ExternalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPin

`func (o *UpdatePhoneNumberRequest) SetExternalPin(v string)`

SetExternalPin sets ExternalPin field to given value.

### HasExternalPin

`func (o *UpdatePhoneNumberRequest) HasExternalPin() bool`

HasExternalPin returns a boolean if a field has been set.

### GetHdVoiceEnabled

`func (o *UpdatePhoneNumberRequest) GetHdVoiceEnabled() bool`

GetHdVoiceEnabled returns the HdVoiceEnabled field if non-nil, zero value otherwise.

### GetHdVoiceEnabledOk

`func (o *UpdatePhoneNumberRequest) GetHdVoiceEnabledOk() (*bool, bool)`

GetHdVoiceEnabledOk returns a tuple with the HdVoiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdVoiceEnabled

`func (o *UpdatePhoneNumberRequest) SetHdVoiceEnabled(v bool)`

SetHdVoiceEnabled sets HdVoiceEnabled field to given value.

### HasHdVoiceEnabled

`func (o *UpdatePhoneNumberRequest) HasHdVoiceEnabled() bool`

HasHdVoiceEnabled returns a boolean if a field has been set.

### GetCustomerReference

`func (o *UpdatePhoneNumberRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *UpdatePhoneNumberRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *UpdatePhoneNumberRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *UpdatePhoneNumberRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetConnectionId

`func (o *UpdatePhoneNumberRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UpdatePhoneNumberRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UpdatePhoneNumberRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *UpdatePhoneNumberRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *UpdatePhoneNumberRequest) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *UpdatePhoneNumberRequest) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *UpdatePhoneNumberRequest) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *UpdatePhoneNumberRequest) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


