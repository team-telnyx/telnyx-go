# CreatePortingOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | **[]string** | The list of +E.164 formatted phone numbers | 
**CustomerReference** | Pointer to **string** | A customer-specified reference number for customer bookkeeping purposes | [optional] 

## Methods

### NewCreatePortingOrder

`func NewCreatePortingOrder(phoneNumbers []string, ) *CreatePortingOrder`

NewCreatePortingOrder instantiates a new CreatePortingOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePortingOrderWithDefaults

`func NewCreatePortingOrderWithDefaults() *CreatePortingOrder`

NewCreatePortingOrderWithDefaults instantiates a new CreatePortingOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *CreatePortingOrder) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *CreatePortingOrder) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *CreatePortingOrder) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetCustomerReference

`func (o *CreatePortingOrder) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreatePortingOrder) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreatePortingOrder) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreatePortingOrder) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


