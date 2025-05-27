# NumberBlockOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**StartingNumber** | Pointer to **string** | Starting phone number block | [optional] 
**Range** | Pointer to **int32** | The phone number range included in the block. | [optional] 
**PhoneNumbersCount** | Pointer to **int32** | The count of phone numbers in the number order. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated to all numbers in the phone number block. | [optional] 
**MessagingProfileId** | Pointer to **string** | Identifies the messaging profile associated to all numbers in the phone number block. | [optional] 
**Status** | Pointer to **string** | The status of the order. | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the number order was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the number order was updated. | [optional] [readonly] 
**RequirementsMet** | Pointer to **bool** | True if all requirements are met for every phone number, false otherwise. | [optional] [readonly] 

## Methods

### NewNumberBlockOrder

`func NewNumberBlockOrder() *NumberBlockOrder`

NewNumberBlockOrder instantiates a new NumberBlockOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberBlockOrderWithDefaults

`func NewNumberBlockOrderWithDefaults() *NumberBlockOrder`

NewNumberBlockOrderWithDefaults instantiates a new NumberBlockOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NumberBlockOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberBlockOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberBlockOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NumberBlockOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *NumberBlockOrder) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberBlockOrder) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberBlockOrder) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberBlockOrder) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStartingNumber

`func (o *NumberBlockOrder) GetStartingNumber() string`

GetStartingNumber returns the StartingNumber field if non-nil, zero value otherwise.

### GetStartingNumberOk

`func (o *NumberBlockOrder) GetStartingNumberOk() (*string, bool)`

GetStartingNumberOk returns a tuple with the StartingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingNumber

`func (o *NumberBlockOrder) SetStartingNumber(v string)`

SetStartingNumber sets StartingNumber field to given value.

### HasStartingNumber

`func (o *NumberBlockOrder) HasStartingNumber() bool`

HasStartingNumber returns a boolean if a field has been set.

### GetRange

`func (o *NumberBlockOrder) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *NumberBlockOrder) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *NumberBlockOrder) SetRange(v int32)`

SetRange sets Range field to given value.

### HasRange

`func (o *NumberBlockOrder) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetPhoneNumbersCount

`func (o *NumberBlockOrder) GetPhoneNumbersCount() int32`

GetPhoneNumbersCount returns the PhoneNumbersCount field if non-nil, zero value otherwise.

### GetPhoneNumbersCountOk

`func (o *NumberBlockOrder) GetPhoneNumbersCountOk() (*int32, bool)`

GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersCount

`func (o *NumberBlockOrder) SetPhoneNumbersCount(v int32)`

SetPhoneNumbersCount sets PhoneNumbersCount field to given value.

### HasPhoneNumbersCount

`func (o *NumberBlockOrder) HasPhoneNumbersCount() bool`

HasPhoneNumbersCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *NumberBlockOrder) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *NumberBlockOrder) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *NumberBlockOrder) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *NumberBlockOrder) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *NumberBlockOrder) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *NumberBlockOrder) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *NumberBlockOrder) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *NumberBlockOrder) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetStatus

`func (o *NumberBlockOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumberBlockOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumberBlockOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumberBlockOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *NumberBlockOrder) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *NumberBlockOrder) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *NumberBlockOrder) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *NumberBlockOrder) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NumberBlockOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NumberBlockOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NumberBlockOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NumberBlockOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NumberBlockOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NumberBlockOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NumberBlockOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NumberBlockOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *NumberBlockOrder) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *NumberBlockOrder) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *NumberBlockOrder) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *NumberBlockOrder) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


