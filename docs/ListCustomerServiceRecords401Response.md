# ListCustomerServiceRecords401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]UnauthorizedError**](UnauthorizedError.md) |  | [optional] 

## Methods

### NewListCustomerServiceRecords401Response

`func NewListCustomerServiceRecords401Response() *ListCustomerServiceRecords401Response`

NewListCustomerServiceRecords401Response instantiates a new ListCustomerServiceRecords401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerServiceRecords401ResponseWithDefaults

`func NewListCustomerServiceRecords401ResponseWithDefaults() *ListCustomerServiceRecords401Response`

NewListCustomerServiceRecords401ResponseWithDefaults instantiates a new ListCustomerServiceRecords401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ListCustomerServiceRecords401Response) GetErrors() []UnauthorizedError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ListCustomerServiceRecords401Response) GetErrorsOk() (*[]UnauthorizedError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ListCustomerServiceRecords401Response) SetErrors(v []UnauthorizedError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ListCustomerServiceRecords401Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


