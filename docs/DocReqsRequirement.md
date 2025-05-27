# DocReqsRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CountryCode** | Pointer to **string** | The 2-character (ISO 3166-1 alpha-2) country code where this requirement applies | [optional] 
**Locality** | Pointer to **string** | The locality where this requirement applies | [optional] 
**PhoneNumberType** | Pointer to **string** | Indicates the phone_number_type this requirement applies to. Leave blank if this requirement applies to all number_types. | [optional] 
**Action** | Pointer to **string** | Indicates whether this requirement applies to ordering, porting, or both | [optional] 
**RequirementsTypes** | Pointer to [**[]DocReqsRequirementType**](DocReqsRequirementType.md) | Lists the requirement types necessary to fulfill this requirement | [optional] [readonly] 
**Id** | Pointer to **string** | Identifies the associated document | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was last updated. | [optional] [readonly] 

## Methods

### NewDocReqsRequirement

`func NewDocReqsRequirement() *DocReqsRequirement`

NewDocReqsRequirement instantiates a new DocReqsRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocReqsRequirementWithDefaults

`func NewDocReqsRequirementWithDefaults() *DocReqsRequirement`

NewDocReqsRequirementWithDefaults instantiates a new DocReqsRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *DocReqsRequirement) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DocReqsRequirement) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DocReqsRequirement) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DocReqsRequirement) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCountryCode

`func (o *DocReqsRequirement) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *DocReqsRequirement) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *DocReqsRequirement) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *DocReqsRequirement) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocality

`func (o *DocReqsRequirement) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *DocReqsRequirement) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *DocReqsRequirement) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *DocReqsRequirement) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *DocReqsRequirement) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *DocReqsRequirement) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *DocReqsRequirement) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *DocReqsRequirement) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetAction

`func (o *DocReqsRequirement) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DocReqsRequirement) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DocReqsRequirement) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DocReqsRequirement) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRequirementsTypes

`func (o *DocReqsRequirement) GetRequirementsTypes() []DocReqsRequirementType`

GetRequirementsTypes returns the RequirementsTypes field if non-nil, zero value otherwise.

### GetRequirementsTypesOk

`func (o *DocReqsRequirement) GetRequirementsTypesOk() (*[]DocReqsRequirementType, bool)`

GetRequirementsTypesOk returns a tuple with the RequirementsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsTypes

`func (o *DocReqsRequirement) SetRequirementsTypes(v []DocReqsRequirementType)`

SetRequirementsTypes sets RequirementsTypes field to given value.

### HasRequirementsTypes

`func (o *DocReqsRequirement) HasRequirementsTypes() bool`

HasRequirementsTypes returns a boolean if a field has been set.

### GetId

`func (o *DocReqsRequirement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocReqsRequirement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocReqsRequirement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocReqsRequirement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocReqsRequirement) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocReqsRequirement) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocReqsRequirement) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocReqsRequirement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocReqsRequirement) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocReqsRequirement) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocReqsRequirement) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocReqsRequirement) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


