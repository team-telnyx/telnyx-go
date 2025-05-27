# InboundMessagePayloadCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | The amount deducted from your account. | [optional] 
**Currency** | Pointer to **string** | The ISO 4217 currency identifier. | [optional] 

## Methods

### NewInboundMessagePayloadCost

`func NewInboundMessagePayloadCost() *InboundMessagePayloadCost`

NewInboundMessagePayloadCost instantiates a new InboundMessagePayloadCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundMessagePayloadCostWithDefaults

`func NewInboundMessagePayloadCostWithDefaults() *InboundMessagePayloadCost`

NewInboundMessagePayloadCostWithDefaults instantiates a new InboundMessagePayloadCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InboundMessagePayloadCost) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InboundMessagePayloadCost) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InboundMessagePayloadCost) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InboundMessagePayloadCost) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *InboundMessagePayloadCost) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InboundMessagePayloadCost) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InboundMessagePayloadCost) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InboundMessagePayloadCost) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


