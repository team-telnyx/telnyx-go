# NumberOrderWithPhoneNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumbersCount** | Pointer to **int32** | The count of phone numbers in the number order. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with this phone number. | [optional] 
**MessagingProfileId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**BillingGroupId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**PhoneNumbers** | Pointer to [**[]NumberOrderPhoneNumber**](NumberOrderPhoneNumber.md) |  | [optional] 
**SubNumberOrdersIds** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** | The status of the order. | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the number order was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the number order was updated. | [optional] [readonly] 
**RequirementsMet** | Pointer to **bool** | True if all requirements are met for every phone number, false otherwise. | [optional] [readonly] 

## Methods

### NewNumberOrderWithPhoneNumbers

`func NewNumberOrderWithPhoneNumbers() *NumberOrderWithPhoneNumbers`

NewNumberOrderWithPhoneNumbers instantiates a new NumberOrderWithPhoneNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberOrderWithPhoneNumbersWithDefaults

`func NewNumberOrderWithPhoneNumbersWithDefaults() *NumberOrderWithPhoneNumbers`

NewNumberOrderWithPhoneNumbersWithDefaults instantiates a new NumberOrderWithPhoneNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NumberOrderWithPhoneNumbers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberOrderWithPhoneNumbers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberOrderWithPhoneNumbers) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NumberOrderWithPhoneNumbers) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *NumberOrderWithPhoneNumbers) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberOrderWithPhoneNumbers) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberOrderWithPhoneNumbers) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberOrderWithPhoneNumbers) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumbersCount

`func (o *NumberOrderWithPhoneNumbers) GetPhoneNumbersCount() int32`

GetPhoneNumbersCount returns the PhoneNumbersCount field if non-nil, zero value otherwise.

### GetPhoneNumbersCountOk

`func (o *NumberOrderWithPhoneNumbers) GetPhoneNumbersCountOk() (*int32, bool)`

GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersCount

`func (o *NumberOrderWithPhoneNumbers) SetPhoneNumbersCount(v int32)`

SetPhoneNumbersCount sets PhoneNumbersCount field to given value.

### HasPhoneNumbersCount

`func (o *NumberOrderWithPhoneNumbers) HasPhoneNumbersCount() bool`

HasPhoneNumbersCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *NumberOrderWithPhoneNumbers) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *NumberOrderWithPhoneNumbers) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *NumberOrderWithPhoneNumbers) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *NumberOrderWithPhoneNumbers) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *NumberOrderWithPhoneNumbers) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *NumberOrderWithPhoneNumbers) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *NumberOrderWithPhoneNumbers) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *NumberOrderWithPhoneNumbers) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *NumberOrderWithPhoneNumbers) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *NumberOrderWithPhoneNumbers) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *NumberOrderWithPhoneNumbers) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *NumberOrderWithPhoneNumbers) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *NumberOrderWithPhoneNumbers) GetPhoneNumbers() []NumberOrderPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *NumberOrderWithPhoneNumbers) GetPhoneNumbersOk() (*[]NumberOrderPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *NumberOrderWithPhoneNumbers) SetPhoneNumbers(v []NumberOrderPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *NumberOrderWithPhoneNumbers) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetSubNumberOrdersIds

`func (o *NumberOrderWithPhoneNumbers) GetSubNumberOrdersIds() []string`

GetSubNumberOrdersIds returns the SubNumberOrdersIds field if non-nil, zero value otherwise.

### GetSubNumberOrdersIdsOk

`func (o *NumberOrderWithPhoneNumbers) GetSubNumberOrdersIdsOk() (*[]string, bool)`

GetSubNumberOrdersIdsOk returns a tuple with the SubNumberOrdersIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNumberOrdersIds

`func (o *NumberOrderWithPhoneNumbers) SetSubNumberOrdersIds(v []string)`

SetSubNumberOrdersIds sets SubNumberOrdersIds field to given value.

### HasSubNumberOrdersIds

`func (o *NumberOrderWithPhoneNumbers) HasSubNumberOrdersIds() bool`

HasSubNumberOrdersIds returns a boolean if a field has been set.

### GetStatus

`func (o *NumberOrderWithPhoneNumbers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumberOrderWithPhoneNumbers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumberOrderWithPhoneNumbers) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumberOrderWithPhoneNumbers) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *NumberOrderWithPhoneNumbers) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *NumberOrderWithPhoneNumbers) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *NumberOrderWithPhoneNumbers) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *NumberOrderWithPhoneNumbers) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NumberOrderWithPhoneNumbers) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NumberOrderWithPhoneNumbers) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NumberOrderWithPhoneNumbers) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NumberOrderWithPhoneNumbers) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NumberOrderWithPhoneNumbers) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NumberOrderWithPhoneNumbers) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NumberOrderWithPhoneNumbers) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NumberOrderWithPhoneNumbers) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *NumberOrderWithPhoneNumbers) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *NumberOrderWithPhoneNumbers) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *NumberOrderWithPhoneNumbers) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *NumberOrderWithPhoneNumbers) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


