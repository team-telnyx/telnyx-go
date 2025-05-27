# GetCustomerServiceRecord404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ResourceNotFoundError**](ResourceNotFoundError.md) |  | [optional] 

## Methods

### NewGetCustomerServiceRecord404Response

`func NewGetCustomerServiceRecord404Response() *GetCustomerServiceRecord404Response`

NewGetCustomerServiceRecord404Response instantiates a new GetCustomerServiceRecord404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerServiceRecord404ResponseWithDefaults

`func NewGetCustomerServiceRecord404ResponseWithDefaults() *GetCustomerServiceRecord404Response`

NewGetCustomerServiceRecord404ResponseWithDefaults instantiates a new GetCustomerServiceRecord404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *GetCustomerServiceRecord404Response) GetErrors() []ResourceNotFoundError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetCustomerServiceRecord404Response) GetErrorsOk() (*[]ResourceNotFoundError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetCustomerServiceRecord404Response) SetErrors(v []ResourceNotFoundError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetCustomerServiceRecord404Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


