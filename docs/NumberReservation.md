# NumberReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**PhoneNumbers** | Pointer to [**[]ReservedPhoneNumber**](ReservedPhoneNumber.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the entire reservation. | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**CreatedAt** | Pointer to **string** | An ISO 8901 datetime string denoting when the numbers reservation was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | An ISO 8901 datetime string for when the number reservation was updated. | [optional] [readonly] 

## Methods

### NewNumberReservation

`func NewNumberReservation() *NumberReservation`

NewNumberReservation instantiates a new NumberReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberReservationWithDefaults

`func NewNumberReservationWithDefaults() *NumberReservation`

NewNumberReservationWithDefaults instantiates a new NumberReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NumberReservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumberReservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumberReservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NumberReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *NumberReservation) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NumberReservation) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NumberReservation) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NumberReservation) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *NumberReservation) GetPhoneNumbers() []ReservedPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *NumberReservation) GetPhoneNumbersOk() (*[]ReservedPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *NumberReservation) SetPhoneNumbers(v []ReservedPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *NumberReservation) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetStatus

`func (o *NumberReservation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumberReservation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumberReservation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumberReservation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerReference

`func (o *NumberReservation) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *NumberReservation) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *NumberReservation) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *NumberReservation) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NumberReservation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NumberReservation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NumberReservation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NumberReservation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NumberReservation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NumberReservation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NumberReservation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NumberReservation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


