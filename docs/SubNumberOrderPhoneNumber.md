# SubNumberOrderPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**RequirementsMet** | Pointer to **bool** |  | [optional] 
**RequirementsStatus** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**BundleId** | Pointer to **string** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]SubNumberOrderPhoneNumberRegulatoryRequirementsInner**](SubNumberOrderPhoneNumberRegulatoryRequirementsInner.md) |  | [optional] 

## Methods

### NewSubNumberOrderPhoneNumber

`func NewSubNumberOrderPhoneNumber() *SubNumberOrderPhoneNumber`

NewSubNumberOrderPhoneNumber instantiates a new SubNumberOrderPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNumberOrderPhoneNumberWithDefaults

`func NewSubNumberOrderPhoneNumberWithDefaults() *SubNumberOrderPhoneNumber`

NewSubNumberOrderPhoneNumberWithDefaults instantiates a new SubNumberOrderPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNumberOrderPhoneNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNumberOrderPhoneNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNumberOrderPhoneNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubNumberOrderPhoneNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *SubNumberOrderPhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SubNumberOrderPhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SubNumberOrderPhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SubNumberOrderPhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *SubNumberOrderPhoneNumber) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SubNumberOrderPhoneNumber) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SubNumberOrderPhoneNumber) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SubNumberOrderPhoneNumber) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *SubNumberOrderPhoneNumber) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *SubNumberOrderPhoneNumber) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *SubNumberOrderPhoneNumber) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *SubNumberOrderPhoneNumber) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *SubNumberOrderPhoneNumber) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *SubNumberOrderPhoneNumber) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *SubNumberOrderPhoneNumber) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *SubNumberOrderPhoneNumber) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetRequirementsStatus

`func (o *SubNumberOrderPhoneNumber) GetRequirementsStatus() string`

GetRequirementsStatus returns the RequirementsStatus field if non-nil, zero value otherwise.

### GetRequirementsStatusOk

`func (o *SubNumberOrderPhoneNumber) GetRequirementsStatusOk() (*string, bool)`

GetRequirementsStatusOk returns a tuple with the RequirementsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsStatus

`func (o *SubNumberOrderPhoneNumber) SetRequirementsStatus(v string)`

SetRequirementsStatus sets RequirementsStatus field to given value.

### HasRequirementsStatus

`func (o *SubNumberOrderPhoneNumber) HasRequirementsStatus() bool`

HasRequirementsStatus returns a boolean if a field has been set.

### GetStatus

`func (o *SubNumberOrderPhoneNumber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubNumberOrderPhoneNumber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubNumberOrderPhoneNumber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubNumberOrderPhoneNumber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRecordType

`func (o *SubNumberOrderPhoneNumber) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SubNumberOrderPhoneNumber) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SubNumberOrderPhoneNumber) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SubNumberOrderPhoneNumber) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetBundleId

`func (o *SubNumberOrderPhoneNumber) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *SubNumberOrderPhoneNumber) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *SubNumberOrderPhoneNumber) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *SubNumberOrderPhoneNumber) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *SubNumberOrderPhoneNumber) GetRegulatoryRequirements() []SubNumberOrderPhoneNumberRegulatoryRequirementsInner`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *SubNumberOrderPhoneNumber) GetRegulatoryRequirementsOk() (*[]SubNumberOrderPhoneNumberRegulatoryRequirementsInner, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *SubNumberOrderPhoneNumber) SetRegulatoryRequirements(v []SubNumberOrderPhoneNumberRegulatoryRequirementsInner)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *SubNumberOrderPhoneNumber) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


