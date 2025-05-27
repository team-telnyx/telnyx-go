# SubNumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**OrderRequestId** | Pointer to **string** |  | [optional] [readonly] 
**CountryCode** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumberType** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**RegulatoryRequirements** | Pointer to [**[]SubNumberOrderRegulatoryRequirement**](SubNumberOrderRegulatoryRequirement.md) |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumbersCount** | Pointer to **int32** | The count of phone numbers in the number order. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the number order was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the number order was updated. | [optional] [readonly] 
**RequirementsMet** | Pointer to **bool** | True if all requirements are met for every phone number, false otherwise. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the order. | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**IsBlockSubNumberOrder** | Pointer to **bool** | True if the sub number order is a block sub number order | [optional] [readonly] 

## Methods

### NewSubNumberOrder

`func NewSubNumberOrder() *SubNumberOrder`

NewSubNumberOrder instantiates a new SubNumberOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNumberOrderWithDefaults

`func NewSubNumberOrderWithDefaults() *SubNumberOrder`

NewSubNumberOrderWithDefaults instantiates a new SubNumberOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubNumberOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubNumberOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubNumberOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubNumberOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderRequestId

`func (o *SubNumberOrder) GetOrderRequestId() string`

GetOrderRequestId returns the OrderRequestId field if non-nil, zero value otherwise.

### GetOrderRequestIdOk

`func (o *SubNumberOrder) GetOrderRequestIdOk() (*string, bool)`

GetOrderRequestIdOk returns a tuple with the OrderRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRequestId

`func (o *SubNumberOrder) SetOrderRequestId(v string)`

SetOrderRequestId sets OrderRequestId field to given value.

### HasOrderRequestId

`func (o *SubNumberOrder) HasOrderRequestId() bool`

HasOrderRequestId returns a boolean if a field has been set.

### GetCountryCode

`func (o *SubNumberOrder) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SubNumberOrder) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SubNumberOrder) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SubNumberOrder) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *SubNumberOrder) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *SubNumberOrder) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *SubNumberOrder) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *SubNumberOrder) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetUserId

`func (o *SubNumberOrder) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubNumberOrder) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubNumberOrder) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SubNumberOrder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRegulatoryRequirements

`func (o *SubNumberOrder) GetRegulatoryRequirements() []SubNumberOrderRegulatoryRequirement`

GetRegulatoryRequirements returns the RegulatoryRequirements field if non-nil, zero value otherwise.

### GetRegulatoryRequirementsOk

`func (o *SubNumberOrder) GetRegulatoryRequirementsOk() (*[]SubNumberOrderRegulatoryRequirement, bool)`

GetRegulatoryRequirementsOk returns a tuple with the RegulatoryRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryRequirements

`func (o *SubNumberOrder) SetRegulatoryRequirements(v []SubNumberOrderRegulatoryRequirement)`

SetRegulatoryRequirements sets RegulatoryRequirements field to given value.

### HasRegulatoryRequirements

`func (o *SubNumberOrder) HasRegulatoryRequirements() bool`

HasRegulatoryRequirements returns a boolean if a field has been set.

### GetRecordType

`func (o *SubNumberOrder) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SubNumberOrder) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SubNumberOrder) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SubNumberOrder) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumbersCount

`func (o *SubNumberOrder) GetPhoneNumbersCount() int32`

GetPhoneNumbersCount returns the PhoneNumbersCount field if non-nil, zero value otherwise.

### GetPhoneNumbersCountOk

`func (o *SubNumberOrder) GetPhoneNumbersCountOk() (*int32, bool)`

GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersCount

`func (o *SubNumberOrder) SetPhoneNumbersCount(v int32)`

SetPhoneNumbersCount sets PhoneNumbersCount field to given value.

### HasPhoneNumbersCount

`func (o *SubNumberOrder) HasPhoneNumbersCount() bool`

HasPhoneNumbersCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubNumberOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubNumberOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubNumberOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubNumberOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubNumberOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubNumberOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubNumberOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubNumberOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *SubNumberOrder) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *SubNumberOrder) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *SubNumberOrder) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *SubNumberOrder) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetStatus

`func (o *SubNumberOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubNumberOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubNumberOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubNumberOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *SubNumberOrder) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *SubNumberOrder) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *SubNumberOrder) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *SubNumberOrder) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetIsBlockSubNumberOrder

`func (o *SubNumberOrder) GetIsBlockSubNumberOrder() bool`

GetIsBlockSubNumberOrder returns the IsBlockSubNumberOrder field if non-nil, zero value otherwise.

### GetIsBlockSubNumberOrderOk

`func (o *SubNumberOrder) GetIsBlockSubNumberOrderOk() (*bool, bool)`

GetIsBlockSubNumberOrderOk returns a tuple with the IsBlockSubNumberOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlockSubNumberOrder

`func (o *SubNumberOrder) SetIsBlockSubNumberOrder(v bool)`

SetIsBlockSubNumberOrder sets IsBlockSubNumberOrder field to given value.

### HasIsBlockSubNumberOrder

`func (o *SubNumberOrder) HasIsBlockSubNumberOrder() bool`

HasIsBlockSubNumberOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


