# SourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pointer** | Pointer to **string** | A JSON Pointer [RFC6901] to the associated entity in the request document | [optional] 
**Parameter** | Pointer to **string** | a string indicating which URI query parameter caused the error | [optional] 

## Methods

### NewSourceResponse

`func NewSourceResponse() *SourceResponse`

NewSourceResponse instantiates a new SourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceResponseWithDefaults

`func NewSourceResponseWithDefaults() *SourceResponse`

NewSourceResponseWithDefaults instantiates a new SourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPointer

`func (o *SourceResponse) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *SourceResponse) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *SourceResponse) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *SourceResponse) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetParameter

`func (o *SourceResponse) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *SourceResponse) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *SourceResponse) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *SourceResponse) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


