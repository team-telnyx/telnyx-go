# WirelessCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Final cost. Cost is calculated as rate * unit | [optional] 
**Currency** | Pointer to **string** | Currency of the rate and cost | [optional] 

## Methods

### NewWirelessCost

`func NewWirelessCost() *WirelessCost`

NewWirelessCost instantiates a new WirelessCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessCostWithDefaults

`func NewWirelessCostWithDefaults() *WirelessCost`

NewWirelessCostWithDefaults instantiates a new WirelessCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WirelessCost) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WirelessCost) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WirelessCost) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WirelessCost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *WirelessCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WirelessCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WirelessCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WirelessCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


