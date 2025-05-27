# CreateNumberOrderRequestPhoneNumbersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | e164_phone_number | 
**RequirementGroupId** | Pointer to **string** | ID of requirement group to use to satisfy number requirements | [optional] 
**BundleId** | Pointer to **string** | ID of bundle to associate the number to | [optional] 

## Methods

### NewCreateNumberOrderRequestPhoneNumbersInner

`func NewCreateNumberOrderRequestPhoneNumbersInner(phoneNumber string, ) *CreateNumberOrderRequestPhoneNumbersInner`

NewCreateNumberOrderRequestPhoneNumbersInner instantiates a new CreateNumberOrderRequestPhoneNumbersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNumberOrderRequestPhoneNumbersInnerWithDefaults

`func NewCreateNumberOrderRequestPhoneNumbersInnerWithDefaults() *CreateNumberOrderRequestPhoneNumbersInner`

NewCreateNumberOrderRequestPhoneNumbersInnerWithDefaults instantiates a new CreateNumberOrderRequestPhoneNumbersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CreateNumberOrderRequestPhoneNumbersInner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateNumberOrderRequestPhoneNumbersInner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateNumberOrderRequestPhoneNumbersInner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetRequirementGroupId

`func (o *CreateNumberOrderRequestPhoneNumbersInner) GetRequirementGroupId() string`

GetRequirementGroupId returns the RequirementGroupId field if non-nil, zero value otherwise.

### GetRequirementGroupIdOk

`func (o *CreateNumberOrderRequestPhoneNumbersInner) GetRequirementGroupIdOk() (*string, bool)`

GetRequirementGroupIdOk returns a tuple with the RequirementGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementGroupId

`func (o *CreateNumberOrderRequestPhoneNumbersInner) SetRequirementGroupId(v string)`

SetRequirementGroupId sets RequirementGroupId field to given value.

### HasRequirementGroupId

`func (o *CreateNumberOrderRequestPhoneNumbersInner) HasRequirementGroupId() bool`

HasRequirementGroupId returns a boolean if a field has been set.

### GetBundleId

`func (o *CreateNumberOrderRequestPhoneNumbersInner) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *CreateNumberOrderRequestPhoneNumbersInner) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *CreateNumberOrderRequestPhoneNumbersInner) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *CreateNumberOrderRequestPhoneNumbersInner) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


