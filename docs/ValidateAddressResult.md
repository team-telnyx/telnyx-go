# ValidateAddressResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **string** | Indicates whether an address is valid or invalid. | 
**Suggested** | [**ValidateAddressField**](ValidateAddressField.md) |  | 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Errors** | Pointer to [**[]Error**](Error.md) |  | [optional] 

## Methods

### NewValidateAddressResult

`func NewValidateAddressResult(result string, suggested ValidateAddressField, ) *ValidateAddressResult`

NewValidateAddressResult instantiates a new ValidateAddressResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateAddressResultWithDefaults

`func NewValidateAddressResultWithDefaults() *ValidateAddressResult`

NewValidateAddressResultWithDefaults instantiates a new ValidateAddressResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ValidateAddressResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ValidateAddressResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ValidateAddressResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetSuggested

`func (o *ValidateAddressResult) GetSuggested() ValidateAddressField`

GetSuggested returns the Suggested field if non-nil, zero value otherwise.

### GetSuggestedOk

`func (o *ValidateAddressResult) GetSuggestedOk() (*ValidateAddressField, bool)`

GetSuggestedOk returns a tuple with the Suggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggested

`func (o *ValidateAddressResult) SetSuggested(v ValidateAddressField)`

SetSuggested sets Suggested field to given value.


### GetRecordType

`func (o *ValidateAddressResult) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ValidateAddressResult) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ValidateAddressResult) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ValidateAddressResult) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetErrors

`func (o *ValidateAddressResult) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidateAddressResult) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidateAddressResult) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidateAddressResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


