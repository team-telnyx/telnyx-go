# AdvancedOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** |  | [optional] [default to "US"]
**Comments** | Pointer to **string** |  | [optional] [default to ""]
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**AreaCode** | Pointer to **string** |  | [optional] [default to ""]
**PhoneNumberType** | Pointer to **string** |  | [optional] [default to "local"]
**Features** | Pointer to **[]string** |  | [optional] 
**CustomerReference** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewAdvancedOrderRequest

`func NewAdvancedOrderRequest() *AdvancedOrderRequest`

NewAdvancedOrderRequest instantiates a new AdvancedOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedOrderRequestWithDefaults

`func NewAdvancedOrderRequestWithDefaults() *AdvancedOrderRequest`

NewAdvancedOrderRequestWithDefaults instantiates a new AdvancedOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *AdvancedOrderRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AdvancedOrderRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AdvancedOrderRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AdvancedOrderRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetComments

`func (o *AdvancedOrderRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AdvancedOrderRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AdvancedOrderRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AdvancedOrderRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetQuantity

`func (o *AdvancedOrderRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AdvancedOrderRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AdvancedOrderRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AdvancedOrderRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAreaCode

`func (o *AdvancedOrderRequest) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *AdvancedOrderRequest) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCode

`func (o *AdvancedOrderRequest) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.

### HasAreaCode

`func (o *AdvancedOrderRequest) HasAreaCode() bool`

HasAreaCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *AdvancedOrderRequest) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *AdvancedOrderRequest) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *AdvancedOrderRequest) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *AdvancedOrderRequest) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetFeatures

`func (o *AdvancedOrderRequest) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AdvancedOrderRequest) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AdvancedOrderRequest) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AdvancedOrderRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetCustomerReference

`func (o *AdvancedOrderRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *AdvancedOrderRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *AdvancedOrderRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *AdvancedOrderRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


