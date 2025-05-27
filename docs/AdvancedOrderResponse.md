# AdvancedOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** |  | [optional] [default to "US"]
**Comments** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**AreaCode** | Pointer to **string** |  | [optional] [default to ""]
**PhoneNumberType** | Pointer to **string** |  | [optional] [default to ""]
**Features** | Pointer to **[]string** |  | [optional] 
**CustomerReference** | Pointer to **string** |  | [optional] [default to ""]
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Orders** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvancedOrderResponse

`func NewAdvancedOrderResponse() *AdvancedOrderResponse`

NewAdvancedOrderResponse instantiates a new AdvancedOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedOrderResponseWithDefaults

`func NewAdvancedOrderResponseWithDefaults() *AdvancedOrderResponse`

NewAdvancedOrderResponseWithDefaults instantiates a new AdvancedOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *AdvancedOrderResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AdvancedOrderResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AdvancedOrderResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AdvancedOrderResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetComments

`func (o *AdvancedOrderResponse) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AdvancedOrderResponse) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AdvancedOrderResponse) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AdvancedOrderResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetQuantity

`func (o *AdvancedOrderResponse) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AdvancedOrderResponse) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AdvancedOrderResponse) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AdvancedOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAreaCode

`func (o *AdvancedOrderResponse) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *AdvancedOrderResponse) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCode

`func (o *AdvancedOrderResponse) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.

### HasAreaCode

`func (o *AdvancedOrderResponse) HasAreaCode() bool`

HasAreaCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *AdvancedOrderResponse) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *AdvancedOrderResponse) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *AdvancedOrderResponse) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *AdvancedOrderResponse) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetFeatures

`func (o *AdvancedOrderResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AdvancedOrderResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AdvancedOrderResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AdvancedOrderResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCustomerReference

`func (o *AdvancedOrderResponse) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *AdvancedOrderResponse) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *AdvancedOrderResponse) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *AdvancedOrderResponse) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetId

`func (o *AdvancedOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvancedOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvancedOrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvancedOrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AdvancedOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvancedOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvancedOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvancedOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrders

`func (o *AdvancedOrderResponse) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *AdvancedOrderResponse) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *AdvancedOrderResponse) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *AdvancedOrderResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


