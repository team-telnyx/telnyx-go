# ManagedAccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Balance** | Pointer to **float64** | The account&#39;s current balance. | [optional] 
**CreditLimit** | Pointer to **float64** | The account&#39;s credit limit. | [optional] 
**AvailableCredit** | Pointer to **float64** | Available amount to spend (balance + credit limit) | [optional] 
**Currency** | Pointer to **string** | The ISO 4217 currency identifier. | [optional] 

## Methods

### NewManagedAccountBalance

`func NewManagedAccountBalance() *ManagedAccountBalance`

NewManagedAccountBalance instantiates a new ManagedAccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountBalanceWithDefaults

`func NewManagedAccountBalanceWithDefaults() *ManagedAccountBalance`

NewManagedAccountBalanceWithDefaults instantiates a new ManagedAccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ManagedAccountBalance) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ManagedAccountBalance) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ManagedAccountBalance) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ManagedAccountBalance) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetBalance

`func (o *ManagedAccountBalance) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *ManagedAccountBalance) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *ManagedAccountBalance) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *ManagedAccountBalance) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreditLimit

`func (o *ManagedAccountBalance) GetCreditLimit() float64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *ManagedAccountBalance) GetCreditLimitOk() (*float64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *ManagedAccountBalance) SetCreditLimit(v float64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *ManagedAccountBalance) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetAvailableCredit

`func (o *ManagedAccountBalance) GetAvailableCredit() float64`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *ManagedAccountBalance) GetAvailableCreditOk() (*float64, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *ManagedAccountBalance) SetAvailableCredit(v float64)`

SetAvailableCredit sets AvailableCredit field to given value.

### HasAvailableCredit

`func (o *ManagedAccountBalance) HasAvailableCredit() bool`

HasAvailableCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *ManagedAccountBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ManagedAccountBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ManagedAccountBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ManagedAccountBalance) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


