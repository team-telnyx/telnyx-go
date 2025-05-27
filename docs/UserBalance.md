# UserBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to **float64** | The account’s pending amount. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Balance** | Pointer to **float64** | The account&#39;s current balance. | [optional] 
**CreditLimit** | Pointer to **float64** | The account&#39;s credit limit. | [optional] 
**AvailableCredit** | Pointer to **float64** | Available amount to spend (balance + credit limit) | [optional] 
**Currency** | Pointer to **string** | The ISO 4217 currency identifier. | [optional] 

## Methods

### NewUserBalance

`func NewUserBalance() *UserBalance`

NewUserBalance instantiates a new UserBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBalanceWithDefaults

`func NewUserBalanceWithDefaults() *UserBalance`

NewUserBalanceWithDefaults instantiates a new UserBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *UserBalance) GetPending() float64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *UserBalance) GetPendingOk() (*float64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *UserBalance) SetPending(v float64)`

SetPending sets Pending field to given value.

### HasPending

`func (o *UserBalance) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetRecordType

`func (o *UserBalance) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *UserBalance) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *UserBalance) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *UserBalance) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetBalance

`func (o *UserBalance) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *UserBalance) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *UserBalance) SetBalance(v float64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *UserBalance) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCreditLimit

`func (o *UserBalance) GetCreditLimit() float64`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *UserBalance) GetCreditLimitOk() (*float64, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *UserBalance) SetCreditLimit(v float64)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *UserBalance) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetAvailableCredit

`func (o *UserBalance) GetAvailableCredit() float64`

GetAvailableCredit returns the AvailableCredit field if non-nil, zero value otherwise.

### GetAvailableCreditOk

`func (o *UserBalance) GetAvailableCreditOk() (*float64, bool)`

GetAvailableCreditOk returns a tuple with the AvailableCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCredit

`func (o *UserBalance) SetAvailableCredit(v float64)`

SetAvailableCredit sets AvailableCredit field to given value.

### HasAvailableCredit

`func (o *UserBalance) HasAvailableCredit() bool`

HasAvailableCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *UserBalance) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UserBalance) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UserBalance) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UserBalance) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


