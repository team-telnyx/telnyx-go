# NumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumbersCount** | Pointer to **int32** | The count of phone numbers in the number order. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with this phone number. | [optional] 
**MessagingProfileId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**BillingGroupId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**PhoneNumbers** | Pointer to [**[]PhoneNumbersJobPhoneNumber**](PhoneNumbersJobPhoneNumber.md) |  | [optional] 
**SubNumberOrdersIds** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** | The status of the order. | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the number order was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the number order was updated. | [optional] [readonly] 
**RequirementsMet** | Pointer to **bool** | True if all requirements are met for every phone number, false otherwise. | [optional] [readonly] 

## Methods

### NewNumberOrder

`func NewNumberOrder() *NumberOrder`

NewNumberOrder instantiates a new NumberOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberOrderWithDefaults

`func NewNumberOrderWithDefaults() *NumberOrder`

NewNumberOrderWithDefaults instantiates a new NumberOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NumberOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NumberOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *NumberOrder) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberOrder) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberOrder) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberOrder) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumbersCount

`func (o *NumberOrder) GetPhoneNumbersCount() int32`

GetPhoneNumbersCount returns the PhoneNumbersCount field if non-nil, zero value otherwise.

### GetPhoneNumbersCountOk

`func (o *NumberOrder) GetPhoneNumbersCountOk() (*int32, bool)`

GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersCount

`func (o *NumberOrder) SetPhoneNumbersCount(v int32)`

SetPhoneNumbersCount sets PhoneNumbersCount field to given value.

### HasPhoneNumbersCount

`func (o *NumberOrder) HasPhoneNumbersCount() bool`

HasPhoneNumbersCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *NumberOrder) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *NumberOrder) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *NumberOrder) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *NumberOrder) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *NumberOrder) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *NumberOrder) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *NumberOrder) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *NumberOrder) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *NumberOrder) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *NumberOrder) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *NumberOrder) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *NumberOrder) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *NumberOrder) GetPhoneNumbers() []PhoneNumbersJobPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *NumberOrder) GetPhoneNumbersOk() (*[]PhoneNumbersJobPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *NumberOrder) SetPhoneNumbers(v []PhoneNumbersJobPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *NumberOrder) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetSubNumberOrdersIds

`func (o *NumberOrder) GetSubNumberOrdersIds() []string`

GetSubNumberOrdersIds returns the SubNumberOrdersIds field if non-nil, zero value otherwise.

### GetSubNumberOrdersIdsOk

`func (o *NumberOrder) GetSubNumberOrdersIdsOk() (*[]string, bool)`

GetSubNumberOrdersIdsOk returns a tuple with the SubNumberOrdersIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNumberOrdersIds

`func (o *NumberOrder) SetSubNumberOrdersIds(v []string)`

SetSubNumberOrdersIds sets SubNumberOrdersIds field to given value.

### HasSubNumberOrdersIds

`func (o *NumberOrder) HasSubNumberOrdersIds() bool`

HasSubNumberOrdersIds returns a boolean if a field has been set.

### GetStatus

`func (o *NumberOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumberOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumberOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumberOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *NumberOrder) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *NumberOrder) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *NumberOrder) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *NumberOrder) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NumberOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NumberOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NumberOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NumberOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NumberOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NumberOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NumberOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NumberOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *NumberOrder) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *NumberOrder) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *NumberOrder) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *NumberOrder) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


