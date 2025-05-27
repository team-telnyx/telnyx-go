# CustomerServiceRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this customer service record | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The phone number of the customer service record. | [optional] 
**Status** | Pointer to **string** | The status of the customer service record | [optional] 
**ErrorMessage** | Pointer to **string** | The error message in case status is &#x60;failed&#x60;. This field would be null in case of &#x60;pending&#x60; or &#x60;completed&#x60; status. | [optional] 
**Result** | Pointer to [**CustomerServiceRecordResult**](CustomerServiceRecordResult.md) |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 

## Methods

### NewCustomerServiceRecord

`func NewCustomerServiceRecord() *CustomerServiceRecord`

NewCustomerServiceRecord instantiates a new CustomerServiceRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceRecordWithDefaults

`func NewCustomerServiceRecordWithDefaults() *CustomerServiceRecord`

NewCustomerServiceRecordWithDefaults instantiates a new CustomerServiceRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerServiceRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerServiceRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerServiceRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerServiceRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerServiceRecord) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerServiceRecord) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerServiceRecord) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerServiceRecord) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerServiceRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerServiceRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerServiceRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerServiceRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CustomerServiceRecord) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CustomerServiceRecord) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CustomerServiceRecord) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CustomerServiceRecord) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetResult

`func (o *CustomerServiceRecord) GetResult() CustomerServiceRecordResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CustomerServiceRecord) GetResultOk() (*CustomerServiceRecordResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CustomerServiceRecord) SetResult(v CustomerServiceRecordResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CustomerServiceRecord) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRecordType

`func (o *CustomerServiceRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CustomerServiceRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CustomerServiceRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CustomerServiceRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerServiceRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerServiceRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerServiceRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CustomerServiceRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomerServiceRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerServiceRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerServiceRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerServiceRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


