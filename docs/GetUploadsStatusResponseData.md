# GetUploadsStatusResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingNumbersCount** | Pointer to **int32** | The count of phone numbers that are pending assignment to the external connection. | [optional] 
**PendingOrdersCount** | Pointer to **int32** | The count of number uploads that have not yet been uploaded to Microsoft. | [optional] 

## Methods

### NewGetUploadsStatusResponseData

`func NewGetUploadsStatusResponseData() *GetUploadsStatusResponseData`

NewGetUploadsStatusResponseData instantiates a new GetUploadsStatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUploadsStatusResponseDataWithDefaults

`func NewGetUploadsStatusResponseDataWithDefaults() *GetUploadsStatusResponseData`

NewGetUploadsStatusResponseDataWithDefaults instantiates a new GetUploadsStatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingNumbersCount

`func (o *GetUploadsStatusResponseData) GetPendingNumbersCount() int32`

GetPendingNumbersCount returns the PendingNumbersCount field if non-nil, zero value otherwise.

### GetPendingNumbersCountOk

`func (o *GetUploadsStatusResponseData) GetPendingNumbersCountOk() (*int32, bool)`

GetPendingNumbersCountOk returns a tuple with the PendingNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingNumbersCount

`func (o *GetUploadsStatusResponseData) SetPendingNumbersCount(v int32)`

SetPendingNumbersCount sets PendingNumbersCount field to given value.

### HasPendingNumbersCount

`func (o *GetUploadsStatusResponseData) HasPendingNumbersCount() bool`

HasPendingNumbersCount returns a boolean if a field has been set.

### GetPendingOrdersCount

`func (o *GetUploadsStatusResponseData) GetPendingOrdersCount() int32`

GetPendingOrdersCount returns the PendingOrdersCount field if non-nil, zero value otherwise.

### GetPendingOrdersCountOk

`func (o *GetUploadsStatusResponseData) GetPendingOrdersCountOk() (*int32, bool)`

GetPendingOrdersCountOk returns a tuple with the PendingOrdersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOrdersCount

`func (o *GetUploadsStatusResponseData) SetPendingOrdersCount(v int32)`

SetPendingOrdersCount sets PendingOrdersCount field to given value.

### HasPendingOrdersCount

`func (o *GetUploadsStatusResponseData) HasPendingOrdersCount() bool`

HasPendingOrdersCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


