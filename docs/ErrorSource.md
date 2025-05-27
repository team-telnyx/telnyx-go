# ErrorSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pointer** | Pointer to **string** | JSON pointer (RFC6901) to the offending entity. | [optional] 
**Parameter** | Pointer to **string** | Indicates which query parameter caused the error. | [optional] 

## Methods

### NewErrorSource

`func NewErrorSource() *ErrorSource`

NewErrorSource instantiates a new ErrorSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorSourceWithDefaults

`func NewErrorSourceWithDefaults() *ErrorSource`

NewErrorSourceWithDefaults instantiates a new ErrorSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointer

`func (o *ErrorSource) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorSource) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorSource) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ErrorSource) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetParameter

`func (o *ErrorSource) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ErrorSource) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *ErrorSource) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *ErrorSource) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


