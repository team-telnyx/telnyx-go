# SIMCardOrderCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | A string representing the cost amount. | [optional] 
**Currency** | Pointer to **string** | Filter by ISO 4217 currency string. | [optional] 

## Methods

### NewSIMCardOrderCost

`func NewSIMCardOrderCost() *SIMCardOrderCost`

NewSIMCardOrderCost instantiates a new SIMCardOrderCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardOrderCostWithDefaults

`func NewSIMCardOrderCostWithDefaults() *SIMCardOrderCost`

NewSIMCardOrderCostWithDefaults instantiates a new SIMCardOrderCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SIMCardOrderCost) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SIMCardOrderCost) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SIMCardOrderCost) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SIMCardOrderCost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *SIMCardOrderCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SIMCardOrderCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SIMCardOrderCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SIMCardOrderCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


