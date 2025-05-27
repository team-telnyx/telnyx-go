# PortingOrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]PortingOrdersExceptionType**](PortingOrdersExceptionType.md) | A list of 0 or more details about this porting order&#39;s status | [optional] 
**Value** | Pointer to **string** | The current status of the porting order | [optional] 

## Methods

### NewPortingOrderStatus

`func NewPortingOrderStatus() *PortingOrderStatus`

NewPortingOrderStatus instantiates a new PortingOrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderStatusWithDefaults

`func NewPortingOrderStatusWithDefaults() *PortingOrderStatus`

NewPortingOrderStatusWithDefaults instantiates a new PortingOrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *PortingOrderStatus) GetDetails() []PortingOrdersExceptionType`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PortingOrderStatus) GetDetailsOk() (*[]PortingOrdersExceptionType, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PortingOrderStatus) SetDetails(v []PortingOrdersExceptionType)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PortingOrderStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetValue

`func (o *PortingOrderStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PortingOrderStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PortingOrderStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PortingOrderStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


