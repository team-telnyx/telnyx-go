# CustomerServiceRecordPhoneNumberCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number that is being verified. | [optional] 
**HasCsrCoverage** | Pointer to **bool** | Indicates whether the phone number is covered or not. | [optional] 
**Reason** | Pointer to **string** | The reason why the phone number is not covered. Only returned if &#x60;has_csr_coverage&#x60; is false. | [optional] 
**AdditionalDataRequired** | Pointer to **[]string** | Additional data required to perform CSR for the phone number. Only returned if &#x60;has_csr_coverage&#x60; is true. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 

## Methods

### NewCustomerServiceRecordPhoneNumberCoverage

`func NewCustomerServiceRecordPhoneNumberCoverage() *CustomerServiceRecordPhoneNumberCoverage`

NewCustomerServiceRecordPhoneNumberCoverage instantiates a new CustomerServiceRecordPhoneNumberCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceRecordPhoneNumberCoverageWithDefaults

`func NewCustomerServiceRecordPhoneNumberCoverageWithDefaults() *CustomerServiceRecordPhoneNumberCoverage`

NewCustomerServiceRecordPhoneNumberCoverageWithDefaults instantiates a new CustomerServiceRecordPhoneNumberCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerServiceRecordPhoneNumberCoverage) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerServiceRecordPhoneNumberCoverage) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetHasCsrCoverage

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetHasCsrCoverage() bool`

GetHasCsrCoverage returns the HasCsrCoverage field if non-nil, zero value otherwise.

### GetHasCsrCoverageOk

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetHasCsrCoverageOk() (*bool, bool)`

GetHasCsrCoverageOk returns a tuple with the HasCsrCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCsrCoverage

`func (o *CustomerServiceRecordPhoneNumberCoverage) SetHasCsrCoverage(v bool)`

SetHasCsrCoverage sets HasCsrCoverage field to given value.

### HasHasCsrCoverage

`func (o *CustomerServiceRecordPhoneNumberCoverage) HasHasCsrCoverage() bool`

HasHasCsrCoverage returns a boolean if a field has been set.

### GetReason

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CustomerServiceRecordPhoneNumberCoverage) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CustomerServiceRecordPhoneNumberCoverage) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAdditionalDataRequired

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetAdditionalDataRequired() []string`

GetAdditionalDataRequired returns the AdditionalDataRequired field if non-nil, zero value otherwise.

### GetAdditionalDataRequiredOk

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetAdditionalDataRequiredOk() (*[]string, bool)`

GetAdditionalDataRequiredOk returns a tuple with the AdditionalDataRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDataRequired

`func (o *CustomerServiceRecordPhoneNumberCoverage) SetAdditionalDataRequired(v []string)`

SetAdditionalDataRequired sets AdditionalDataRequired field to given value.

### HasAdditionalDataRequired

`func (o *CustomerServiceRecordPhoneNumberCoverage) HasAdditionalDataRequired() bool`

HasAdditionalDataRequired returns a boolean if a field has been set.

### GetRecordType

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CustomerServiceRecordPhoneNumberCoverage) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CustomerServiceRecordPhoneNumberCoverage) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CustomerServiceRecordPhoneNumberCoverage) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


