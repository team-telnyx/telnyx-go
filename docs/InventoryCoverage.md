# InventoryCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** |  | [optional] 
**GroupType** | Pointer to **string** |  | [optional] 
**NumberRange** | Pointer to **int32** |  | [optional] 
**NumberType** | Pointer to **string** |  | [optional] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**CoverageType** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**AdministrativeArea** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**AdvanceRequirements** | Pointer to **bool** | Indicates if the phone number requires advance requirements. | [optional] 

## Methods

### NewInventoryCoverage

`func NewInventoryCoverage() *InventoryCoverage`

NewInventoryCoverage instantiates a new InventoryCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryCoverageWithDefaults

`func NewInventoryCoverageWithDefaults() *InventoryCoverage`

NewInventoryCoverageWithDefaults instantiates a new InventoryCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InventoryCoverage) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InventoryCoverage) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InventoryCoverage) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InventoryCoverage) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupType

`func (o *InventoryCoverage) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *InventoryCoverage) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *InventoryCoverage) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *InventoryCoverage) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetNumberRange

`func (o *InventoryCoverage) GetNumberRange() int32`

GetNumberRange returns the NumberRange field if non-nil, zero value otherwise.

### GetNumberRangeOk

`func (o *InventoryCoverage) GetNumberRangeOk() (*int32, bool)`

GetNumberRangeOk returns a tuple with the NumberRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberRange

`func (o *InventoryCoverage) SetNumberRange(v int32)`

SetNumberRange sets NumberRange field to given value.

### HasNumberRange

`func (o *InventoryCoverage) HasNumberRange() bool`

HasNumberRange returns a boolean if a field has been set.

### GetNumberType

`func (o *InventoryCoverage) GetNumberType() string`

GetNumberType returns the NumberType field if non-nil, zero value otherwise.

### GetNumberTypeOk

`func (o *InventoryCoverage) GetNumberTypeOk() (*string, bool)`

GetNumberTypeOk returns a tuple with the NumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberType

`func (o *InventoryCoverage) SetNumberType(v string)`

SetNumberType sets NumberType field to given value.

### HasNumberType

`func (o *InventoryCoverage) HasNumberType() bool`

HasNumberType returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *InventoryCoverage) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *InventoryCoverage) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *InventoryCoverage) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *InventoryCoverage) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetCoverageType

`func (o *InventoryCoverage) GetCoverageType() string`

GetCoverageType returns the CoverageType field if non-nil, zero value otherwise.

### GetCoverageTypeOk

`func (o *InventoryCoverage) GetCoverageTypeOk() (*string, bool)`

GetCoverageTypeOk returns a tuple with the CoverageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageType

`func (o *InventoryCoverage) SetCoverageType(v string)`

SetCoverageType sets CoverageType field to given value.

### HasCoverageType

`func (o *InventoryCoverage) HasCoverageType() bool`

HasCoverageType returns a boolean if a field has been set.

### GetRecordType

`func (o *InventoryCoverage) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *InventoryCoverage) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *InventoryCoverage) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *InventoryCoverage) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetAdministrativeArea

`func (o *InventoryCoverage) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *InventoryCoverage) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *InventoryCoverage) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *InventoryCoverage) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetCount

`func (o *InventoryCoverage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InventoryCoverage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InventoryCoverage) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *InventoryCoverage) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetAdvanceRequirements

`func (o *InventoryCoverage) GetAdvanceRequirements() bool`

GetAdvanceRequirements returns the AdvanceRequirements field if non-nil, zero value otherwise.

### GetAdvanceRequirementsOk

`func (o *InventoryCoverage) GetAdvanceRequirementsOk() (*bool, bool)`

GetAdvanceRequirementsOk returns a tuple with the AdvanceRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceRequirements

`func (o *InventoryCoverage) SetAdvanceRequirements(v bool)`

SetAdvanceRequirements sets AdvanceRequirements field to given value.

### HasAdvanceRequirements

`func (o *InventoryCoverage) HasAdvanceRequirements() bool`

HasAdvanceRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


