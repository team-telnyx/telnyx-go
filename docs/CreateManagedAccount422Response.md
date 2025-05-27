# CreateManagedAccount422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]UnprocessableEntityError**](UnprocessableEntityError.md) |  | [optional] 

## Methods

### NewCreateManagedAccount422Response

`func NewCreateManagedAccount422Response() *CreateManagedAccount422Response`

NewCreateManagedAccount422Response instantiates a new CreateManagedAccount422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateManagedAccount422ResponseWithDefaults

`func NewCreateManagedAccount422ResponseWithDefaults() *CreateManagedAccount422Response`

NewCreateManagedAccount422ResponseWithDefaults instantiates a new CreateManagedAccount422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CreateManagedAccount422Response) GetErrors() []UnprocessableEntityError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateManagedAccount422Response) GetErrorsOk() (*[]UnprocessableEntityError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateManagedAccount422Response) SetErrors(v []UnprocessableEntityError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateManagedAccount422Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


