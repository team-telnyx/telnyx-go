# ListCustomerServiceRecords500Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]UnexpectedError**](UnexpectedError.md) |  | [optional] 

## Methods

### NewListCustomerServiceRecords500Response

`func NewListCustomerServiceRecords500Response() *ListCustomerServiceRecords500Response`

NewListCustomerServiceRecords500Response instantiates a new ListCustomerServiceRecords500Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerServiceRecords500ResponseWithDefaults

`func NewListCustomerServiceRecords500ResponseWithDefaults() *ListCustomerServiceRecords500Response`

NewListCustomerServiceRecords500ResponseWithDefaults instantiates a new ListCustomerServiceRecords500Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ListCustomerServiceRecords500Response) GetErrors() []UnexpectedError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ListCustomerServiceRecords500Response) GetErrorsOk() (*[]UnexpectedError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ListCustomerServiceRecords500Response) SetErrors(v []UnexpectedError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ListCustomerServiceRecords500Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


