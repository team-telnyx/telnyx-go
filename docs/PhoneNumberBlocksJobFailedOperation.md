# PhoneNumberBlocksJobFailedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | The phone number in e164 format. | [optional] 
**Id** | Pointer to **string** | The phone number&#39;s ID | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewPhoneNumberBlocksJobFailedOperation

`func NewPhoneNumberBlocksJobFailedOperation() *PhoneNumberBlocksJobFailedOperation`

NewPhoneNumberBlocksJobFailedOperation instantiates a new PhoneNumberBlocksJobFailedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberBlocksJobFailedOperationWithDefaults

`func NewPhoneNumberBlocksJobFailedOperationWithDefaults() *PhoneNumberBlocksJobFailedOperation`

NewPhoneNumberBlocksJobFailedOperationWithDefaults instantiates a new PhoneNumberBlocksJobFailedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *PhoneNumberBlocksJobFailedOperation) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberBlocksJobFailedOperation) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberBlocksJobFailedOperation) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumberBlocksJobFailedOperation) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetId

`func (o *PhoneNumberBlocksJobFailedOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneNumberBlocksJobFailedOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneNumberBlocksJobFailedOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneNumberBlocksJobFailedOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrors

`func (o *PhoneNumberBlocksJobFailedOperation) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PhoneNumberBlocksJobFailedOperation) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PhoneNumberBlocksJobFailedOperation) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PhoneNumberBlocksJobFailedOperation) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


