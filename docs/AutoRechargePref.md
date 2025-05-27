# AutoRechargePref

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the auto recharge preference. | [optional] 
**RecordType** | Pointer to **string** | The record type. | [optional] 
**ThresholdAmount** | Pointer to **float64** | The threshold amount at which the account will be recharged. | [optional] 
**RechargeAmount** | Pointer to **float64** | The amount to recharge the account, the actual recharge amount will be the amount necessary to reach the threshold amount plus the recharge amount. | [optional] 
**Enabled** | Pointer to **bool** | Whether auto recharge is enabled. | [optional] 
**InvoiceEnabled** | Pointer to **bool** |  | [optional] 
**Preference** | Pointer to **string** | The payment preference for auto recharge. | [optional] 

## Methods

### NewAutoRechargePref

`func NewAutoRechargePref() *AutoRechargePref`

NewAutoRechargePref instantiates a new AutoRechargePref object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRechargePrefWithDefaults

`func NewAutoRechargePrefWithDefaults() *AutoRechargePref`

NewAutoRechargePrefWithDefaults instantiates a new AutoRechargePref object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoRechargePref) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoRechargePref) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoRechargePref) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoRechargePref) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *AutoRechargePref) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AutoRechargePref) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AutoRechargePref) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AutoRechargePref) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetThresholdAmount

`func (o *AutoRechargePref) GetThresholdAmount() float64`

GetThresholdAmount returns the ThresholdAmount field if non-nil, zero value otherwise.

### GetThresholdAmountOk

`func (o *AutoRechargePref) GetThresholdAmountOk() (*float64, bool)`

GetThresholdAmountOk returns a tuple with the ThresholdAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdAmount

`func (o *AutoRechargePref) SetThresholdAmount(v float64)`

SetThresholdAmount sets ThresholdAmount field to given value.

### HasThresholdAmount

`func (o *AutoRechargePref) HasThresholdAmount() bool`

HasThresholdAmount returns a boolean if a field has been set.

### GetRechargeAmount

`func (o *AutoRechargePref) GetRechargeAmount() float64`

GetRechargeAmount returns the RechargeAmount field if non-nil, zero value otherwise.

### GetRechargeAmountOk

`func (o *AutoRechargePref) GetRechargeAmountOk() (*float64, bool)`

GetRechargeAmountOk returns a tuple with the RechargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeAmount

`func (o *AutoRechargePref) SetRechargeAmount(v float64)`

SetRechargeAmount sets RechargeAmount field to given value.

### HasRechargeAmount

`func (o *AutoRechargePref) HasRechargeAmount() bool`

HasRechargeAmount returns a boolean if a field has been set.

### GetEnabled

`func (o *AutoRechargePref) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoRechargePref) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoRechargePref) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoRechargePref) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInvoiceEnabled

`func (o *AutoRechargePref) GetInvoiceEnabled() bool`

GetInvoiceEnabled returns the InvoiceEnabled field if non-nil, zero value otherwise.

### GetInvoiceEnabledOk

`func (o *AutoRechargePref) GetInvoiceEnabledOk() (*bool, bool)`

GetInvoiceEnabledOk returns a tuple with the InvoiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceEnabled

`func (o *AutoRechargePref) SetInvoiceEnabled(v bool)`

SetInvoiceEnabled sets InvoiceEnabled field to given value.

### HasInvoiceEnabled

`func (o *AutoRechargePref) HasInvoiceEnabled() bool`

HasInvoiceEnabled returns a boolean if a field has been set.

### GetPreference

`func (o *AutoRechargePref) GetPreference() string`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *AutoRechargePref) GetPreferenceOk() (*string, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *AutoRechargePref) SetPreference(v string)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *AutoRechargePref) HasPreference() bool`

HasPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


