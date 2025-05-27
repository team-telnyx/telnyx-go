# CreateRequirementGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** | ISO alpha 2 country code | 
**PhoneNumberType** | **string** |  | 
**Action** | **string** |  | 
**CustomerReference** | Pointer to **string** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]CreateRequirementGroupRequestRegulatoryRequirementsInner**](CreateRequirementGroupRequestRegulatoryRequirementsInner.md) |  | [optional] 

## Methods

### NewCreateRequirementGroupRequest

`func NewCreateRequirementGroupRequest(countryCode string, phoneNumberType string, action string, ) *CreateRequirementGroupRequest`

NewCreateRequirementGroupRequest instantiates a new CreateRequirementGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequirementGroupRequestWithDefaults

`func NewCreateRequirementGroupRequestWithDefaults() *CreateRequirementGroupRequest`

NewCreateRequirementGroupRequestWithDefaults instantiates a new CreateRequirementGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CreateRequirementGroupRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CreateRequirementGroupRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CreateRequirementGroupRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetPhoneNumberType

`func (o *CreateRequirementGroupRequest) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *CreateRequirementGroupRequest) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *CreateRequirementGroupRequest) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.


### GetAction

`func (o *CreateRequirementGroupRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateRequirementGroupRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateRequirementGroupRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetCustomerReference

`func (o *CreateRequirementGroupRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreateRequirementGroupRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreateRequirementGroupRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreateRequirementGroupRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *CreateRequirementGroupRequest) GetRegulatoryRequirements() []CreateRequirementGroupRequestRegulatoryRequirementsInner`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *CreateRequirementGroupRequest) GetRegulatoryRequirementsOk() (*[]CreateRequirementGroupRequestRegulatoryRequirementsInner, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *CreateRequirementGroupRequest) SetRegulatoryRequirements(v []CreateRequirementGroupRequestRegulatoryRequirementsInner)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *CreateRequirementGroupRequest) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


