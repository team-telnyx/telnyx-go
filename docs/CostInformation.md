# CostInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpfrontCost** | Pointer to **string** |  | [optional] 
**MonthlyCost** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | The ISO 4217 code for the currency. | [optional] 

## Methods

### NewCostInformation

`func NewCostInformation() *CostInformation`

NewCostInformation instantiates a new CostInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostInformationWithDefaults

`func NewCostInformationWithDefaults() *CostInformation`

NewCostInformationWithDefaults instantiates a new CostInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpfrontCost

`func (o *CostInformation) GetUpfrontCost() string`

GetUpfrontCost returns the UpfrontCost field if non-nil, zero value otherwise.

### GetUpfrontCostOk

`func (o *CostInformation) GetUpfrontCostOk() (*string, bool)`

GetUpfrontCostOk returns a tuple with the UpfrontCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfrontCost

`func (o *CostInformation) SetUpfrontCost(v string)`

SetUpfrontCost sets UpfrontCost field to given value.

### HasUpfrontCost

`func (o *CostInformation) HasUpfrontCost() bool`

HasUpfrontCost returns a boolean if a field has been set.

### GetMonthlyCost

`func (o *CostInformation) GetMonthlyCost() string`

GetMonthlyCost returns the MonthlyCost field if non-nil, zero value otherwise.

### GetMonthlyCostOk

`func (o *CostInformation) GetMonthlyCostOk() (*string, bool)`

GetMonthlyCostOk returns a tuple with the MonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCost

`func (o *CostInformation) SetMonthlyCost(v string)`

SetMonthlyCost sets MonthlyCost field to given value.

### HasMonthlyCost

`func (o *CostInformation) HasMonthlyCost() bool`

HasMonthlyCost returns a boolean if a field has been set.

### GetCurrency

`func (o *CostInformation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CostInformation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CostInformation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CostInformation) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


