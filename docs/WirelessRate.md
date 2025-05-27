# WirelessRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Rate from which cost is calculated | [optional] 
**Currency** | Pointer to **string** | Currency of the rate and cost | [optional] 

## Methods

### NewWirelessRate

`func NewWirelessRate() *WirelessRate`

NewWirelessRate instantiates a new WirelessRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessRateWithDefaults

`func NewWirelessRateWithDefaults() *WirelessRate`

NewWirelessRateWithDefaults instantiates a new WirelessRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WirelessRate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WirelessRate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WirelessRate) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WirelessRate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *WirelessRate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WirelessRate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WirelessRate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WirelessRate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


