# CreateNumberBlockOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**StartingNumber** | **string** | Starting phone number block | 
**Range** | **int32** | The phone number range included in the block. | 
**PhoneNumbersCount** | Pointer to **int32** | The count of phone numbers in the number order. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with this phone number. | [optional] 
**MessagingProfileId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**Status** | Pointer to **string** | The status of the order. | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the number order was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the number order was updated. | [optional] [readonly] 
**RequirementsMet** | Pointer to **bool** | True if all requirements are met for every phone number, false otherwise. | [optional] [readonly] 
**Errors** | Pointer to **string** | Errors the reservation could happen upon | [optional] [readonly] 

## Methods

### NewCreateNumberBlockOrderRequest

`func NewCreateNumberBlockOrderRequest(startingNumber string, range_ int32, ) *CreateNumberBlockOrderRequest`

NewCreateNumberBlockOrderRequest instantiates a new CreateNumberBlockOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNumberBlockOrderRequestWithDefaults

`func NewCreateNumberBlockOrderRequestWithDefaults() *CreateNumberBlockOrderRequest`

NewCreateNumberBlockOrderRequestWithDefaults instantiates a new CreateNumberBlockOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNumberBlockOrderRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNumberBlockOrderRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNumberBlockOrderRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNumberBlockOrderRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *CreateNumberBlockOrderRequest) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CreateNumberBlockOrderRequest) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CreateNumberBlockOrderRequest) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CreateNumberBlockOrderRequest) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStartingNumber

`func (o *CreateNumberBlockOrderRequest) GetStartingNumber() string`

GetStartingNumber returns the StartingNumber field if non-nil, zero value otherwise.

### GetStartingNumberOk

`func (o *CreateNumberBlockOrderRequest) GetStartingNumberOk() (*string, bool)`

GetStartingNumberOk returns a tuple with the StartingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingNumber

`func (o *CreateNumberBlockOrderRequest) SetStartingNumber(v string)`

SetStartingNumber sets StartingNumber field to given value.


### GetRange

`func (o *CreateNumberBlockOrderRequest) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CreateNumberBlockOrderRequest) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CreateNumberBlockOrderRequest) SetRange(v int32)`

SetRange sets Range field to given value.


### GetPhoneNumbersCount

`func (o *CreateNumberBlockOrderRequest) GetPhoneNumbersCount() int32`

GetPhoneNumbersCount returns the PhoneNumbersCount field if non-nil, zero value otherwise.

### GetPhoneNumbersCountOk

`func (o *CreateNumberBlockOrderRequest) GetPhoneNumbersCountOk() (*int32, bool)`

GetPhoneNumbersCountOk returns a tuple with the PhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbersCount

`func (o *CreateNumberBlockOrderRequest) SetPhoneNumbersCount(v int32)`

SetPhoneNumbersCount sets PhoneNumbersCount field to given value.

### HasPhoneNumbersCount

`func (o *CreateNumberBlockOrderRequest) HasPhoneNumbersCount() bool`

HasPhoneNumbersCount returns a boolean if a field has been set.

### GetConnectionId

`func (o *CreateNumberBlockOrderRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateNumberBlockOrderRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateNumberBlockOrderRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CreateNumberBlockOrderRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *CreateNumberBlockOrderRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *CreateNumberBlockOrderRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *CreateNumberBlockOrderRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *CreateNumberBlockOrderRequest) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateNumberBlockOrderRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNumberBlockOrderRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNumberBlockOrderRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateNumberBlockOrderRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *CreateNumberBlockOrderRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreateNumberBlockOrderRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreateNumberBlockOrderRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreateNumberBlockOrderRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateNumberBlockOrderRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateNumberBlockOrderRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateNumberBlockOrderRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateNumberBlockOrderRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateNumberBlockOrderRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateNumberBlockOrderRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateNumberBlockOrderRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateNumberBlockOrderRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *CreateNumberBlockOrderRequest) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *CreateNumberBlockOrderRequest) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *CreateNumberBlockOrderRequest) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *CreateNumberBlockOrderRequest) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetErrors

`func (o *CreateNumberBlockOrderRequest) GetErrors() string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateNumberBlockOrderRequest) GetErrorsOk() (*string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateNumberBlockOrderRequest) SetErrors(v string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateNumberBlockOrderRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


