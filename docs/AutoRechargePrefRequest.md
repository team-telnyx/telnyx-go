# AutoRechargePrefRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThresholdAmount** | Pointer to **float64** | The threshold amount at which the account will be recharged. | [optional] 
**RechargeAmount** | Pointer to **float64** | The amount to recharge the account, the actual recharge amount will be the amount necessary to reach the threshold amount plus the recharge amount. | [optional] 
**Enabled** | Pointer to **bool** | Whether auto recharge is enabled. | [optional] 
**InvoiceEnabled** | Pointer to **bool** |  | [optional] 
**Preference** | Pointer to **string** | The payment preference for auto recharge. | [optional] 

## Methods

### NewAutoRechargePrefRequest

`func NewAutoRechargePrefRequest() *AutoRechargePrefRequest`

NewAutoRechargePrefRequest instantiates a new AutoRechargePrefRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRechargePrefRequestWithDefaults

`func NewAutoRechargePrefRequestWithDefaults() *AutoRechargePrefRequest`

NewAutoRechargePrefRequestWithDefaults instantiates a new AutoRechargePrefRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThresholdAmount

`func (o *AutoRechargePrefRequest) GetThresholdAmount() float64`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *AutoRechargePrefRequest) GetThresholdAmountOk() (*float64, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *AutoRechargePrefRequest) SetThresholdAmount(v float64)`

SetThresholdAmount sets ThresholdAmount field to given value.

### HasThresholdAmount

`func (o *AutoRechargePrefRequest) HasThresholdAmount() bool`

HasThresholdAmount returns a boolean if a field has been set.

### GetRechargeAmount

`func (o *AutoRechargePrefRequest) GetRechargeAmount() float64`

GetRechargeAmount returns the RechargeAmount field if non-nil, zero value otherwise.

### GetRechargeAmountOk

`func (o *AutoRechargePrefRequest) GetRechargeAmountOk() (*float64, bool)`

GetRechargeAmountOk returns a tuple with the RechargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeAmount

`func (o *AutoRechargePrefRequest) SetRechargeAmount(v float64)`

SetRechargeAmount sets RechargeAmount field to given value.

### HasRechargeAmount

`func (o *AutoRechargePrefRequest) HasRechargeAmount() bool`

HasRechargeAmount returns a boolean if a field has been set.

### GetEnabled

`func (o *AutoRechargePrefRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoRechargePrefRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoRechargePrefRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoRechargePrefRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInvoiceEnabled

`func (o *AutoRechargePrefRequest) GetInvoiceEnabled() bool`

GetInvoiceEnabled returns the InvoiceEnabled field if non-nil, zero value otherwise.

### GetInvoiceEnabledOk

`func (o *AutoRechargePrefRequest) GetInvoiceEnabledOk() (*bool, bool)`

GetInvoiceEnabledOk returns a tuple with the InvoiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceEnabled

`func (o *AutoRechargePrefRequest) SetInvoiceEnabled(v bool)`

SetInvoiceEnabled sets InvoiceEnabled field to given value.

### HasInvoiceEnabled

`func (o *AutoRechargePrefRequest) HasInvoiceEnabled() bool`

HasInvoiceEnabled returns a boolean if a field has been set.

### GetPreference

`func (o *AutoRechargePrefRequest) GetPreference() string`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *AutoRechargePrefRequest) GetPreferenceOk() (*string, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *AutoRechargePrefRequest) SetPreference(v string)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *AutoRechargePrefRequest) HasPreference() bool`

HasPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


