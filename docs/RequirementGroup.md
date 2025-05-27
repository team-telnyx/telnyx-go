# RequirementGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**CustomerReference** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]UserRequirement**](UserRequirement.md) |  | [optional] 

## Methods

### NewRequirementGroup

`func NewRequirementGroup() *RequirementGroup`

NewRequirementGroup instantiates a new RequirementGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequirementGroupWithDefaults

`func NewRequirementGroupWithDefaults() *RequirementGroup`

NewRequirementGroupWithDefaults instantiates a new RequirementGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequirementGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequirementGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequirementGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequirementGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCountryCode

`func (o *RequirementGroup) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RequirementGroup) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RequirementGroup) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *RequirementGroup) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *RequirementGroup) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *RequirementGroup) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *RequirementGroup) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *RequirementGroup) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetStatus

`func (o *RequirementGroup) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequirementGroup) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequirementGroup) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequirementGroup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAction

`func (o *RequirementGroup) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequirementGroup) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequirementGroup) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RequirementGroup) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCustomerReference

`func (o *RequirementGroup) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *RequirementGroup) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *RequirementGroup) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *RequirementGroup) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RequirementGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequirementGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequirementGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequirementGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RequirementGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RequirementGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RequirementGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RequirementGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRecordType

`func (o *RequirementGroup) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RequirementGroup) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RequirementGroup) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RequirementGroup) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *RequirementGroup) GetRegulatoryRequirements() []UserRequirement`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *RequirementGroup) GetRegulatoryRequirementsOk() (*[]UserRequirement, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *RequirementGroup) SetRegulatoryRequirements(v []UserRequirement)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *RequirementGroup) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


