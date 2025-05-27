# SubNumberOrderRequirementGroupResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**OrderRequestId** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**PhoneNumbersCount** | Pointer to **int32** |  | [optional] 
**RequirementsMet** | Pointer to **bool** |  | [optional] 
**IsBlockSubNumberOrder** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CustomerReference** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]RegulatoryRequirement**](RegulatoryRequirement.md) |  | [optional] 
**PhoneNumbers** | Pointer to [**[]SubNumberOrderPhoneNumber**](SubNumberOrderPhoneNumber.md) |  | [optional] 

## Methods

### NewSubNumberOrderRequirementGroupResponseData

`func NewSubNumberOrderRequirementGroupResponseData() *SubNumberOrderRequirementGroupResponseData`

NewSubNumberOrderRequirementGroupResponseData instantiates a new SubNumberOrderRequirementGroupResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNumberOrderRequirementGroupResponseDataWithDefaults

`func NewSubNumberOrderRequirementGroupResponseDataWithDefaults() *SubNumberOrderRequirementGroupResponseData`

NewSubNumberOrderRequirementGroupResponseDataWithDefaults instantiates a new SubNumberOrderRequirementGroupResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNumberOrderRequirementGroupResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNumberOrderRequirementGroupResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubNumberOrderRequirementGroupResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderRequestId

`func (o *SubNumberOrderRequirementGroupResponseData) GetOrderRequestId() string`

GetOrderRequestId returns the OrderRequestId field if non-nil, zero value otherwise.

### GetOrderRequestIdOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetOrderRequestIdOk() (*string, bool)`

GetOrderRequestIdOk returns a tuple with the OrderRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRequestId

`func (o *SubNumberOrderRequirementGroupResponseData) SetOrderRequestId(v string)`

SetOrderRequestId sets OrderRequestId field to given value.

### HasOrderRequestId

`func (o *SubNumberOrderRequirementGroupResponseData) HasOrderRequestId() bool`

HasOrderRequestId returns a boolean if a field has been set.

### GetCountryCode

`func (o *SubNumberOrderRequirementGroupResponseData) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SubNumberOrderRequirementGroupResponseData) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SubNumberOrderRequirementGroupResponseData) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *SubNumberOrderRequirementGroupResponseData) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *SubNumberOrderRequirementGroupResponseData) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *SubNumberOrderRequirementGroupResponseData) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetPhoneNumbersCount

`func (o *SubNumberOrderRequirementGroupResponseData) GetPhoneNumbersCount() int32`

GetPhoneNumbersCount returns the PhoneNumbersCount field if non-nil, zero value otherwise.

### GetPhoneNumbersCountOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetPhoneNumbersCountOk() (*int32, bool)`

GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersCount

`func (o *SubNumberOrderRequirementGroupResponseData) SetPhoneNumbersCount(v int32)`

SetPhoneNumbersCount sets PhoneNumbersCount field to given value.

### HasPhoneNumbersCount

`func (o *SubNumberOrderRequirementGroupResponseData) HasPhoneNumbersCount() bool`

HasPhoneNumbersCount returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *SubNumberOrderRequirementGroupResponseData) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *SubNumberOrderRequirementGroupResponseData) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *SubNumberOrderRequirementGroupResponseData) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetIsBlockSubNumberOrder

`func (o *SubNumberOrderRequirementGroupResponseData) GetIsBlockSubNumberOrder() bool`

GetIsBlockSubNumberOrder returns the IsBlockSubNumberOrder field if non-nil, zero value otherwise.

### GetIsBlockSubNumberOrderOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetIsBlockSubNumberOrderOk() (*bool, bool)`

GetIsBlockSubNumberOrderOk returns a tuple with the IsBlockSubNumberOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlockSubNumberOrder

`func (o *SubNumberOrderRequirementGroupResponseData) SetIsBlockSubNumberOrder(v bool)`

SetIsBlockSubNumberOrder sets IsBlockSubNumberOrder field to given value.

### HasIsBlockSubNumberOrder

`func (o *SubNumberOrderRequirementGroupResponseData) HasIsBlockSubNumberOrder() bool`

HasIsBlockSubNumberOrder returns a boolean if a field has been set.

### GetStatus

`func (o *SubNumberOrderRequirementGroupResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubNumberOrderRequirementGroupResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubNumberOrderRequirementGroupResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *SubNumberOrderRequirementGroupResponseData) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *SubNumberOrderRequirementGroupResponseData) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *SubNumberOrderRequirementGroupResponseData) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubNumberOrderRequirementGroupResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubNumberOrderRequirementGroupResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubNumberOrderRequirementGroupResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubNumberOrderRequirementGroupResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubNumberOrderRequirementGroupResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubNumberOrderRequirementGroupResponseData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRecordType

`func (o *SubNumberOrderRequirementGroupResponseData) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SubNumberOrderRequirementGroupResponseData) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SubNumberOrderRequirementGroupResponseData) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *SubNumberOrderRequirementGroupResponseData) GetRegulatoryRequirements() []RegulatoryRequirement`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetRegulatoryRequirementsOk() (*[]RegulatoryRequirement, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *SubNumberOrderRequirementGroupResponseData) SetRegulatoryRequirements(v []RegulatoryRequirement)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *SubNumberOrderRequirementGroupResponseData) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *SubNumberOrderRequirementGroupResponseData) GetPhoneNumbers() []SubNumberOrderPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *SubNumberOrderRequirementGroupResponseData) GetPhoneNumbersOk() (*[]SubNumberOrderPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *SubNumberOrderRequirementGroupResponseData) SetPhoneNumbers(v []SubNumberOrderPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *SubNumberOrderRequirementGroupResponseData) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


