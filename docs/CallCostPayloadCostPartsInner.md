# CallCostPayloadCostPartsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledDurationSecs** | Pointer to **int32** | The number of seconds for which this item will be billed | [optional] 
**CallPart** | Pointer to **string** | The service incurring a charge | [optional] 
**Cost** | Pointer to **float32** | The billed cost of the item, in currency shown in the &#x60;currency&#x60; field | [optional] 
**Currency** | Pointer to **string** | The currency in which &#x60;cost&#x60; is measured | [optional] 
**Rate** | Pointer to **float32** | The cost per unit of the item incurring a charge | [optional] 

## Methods

### NewCallCostPayloadCostPartsInner

`func NewCallCostPayloadCostPartsInner() *CallCostPayloadCostPartsInner`

NewCallCostPayloadCostPartsInner instantiates a new CallCostPayloadCostPartsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallCostPayloadCostPartsInnerWithDefaults

`func NewCallCostPayloadCostPartsInnerWithDefaults() *CallCostPayloadCostPartsInner`

NewCallCostPayloadCostPartsInnerWithDefaults instantiates a new CallCostPayloadCostPartsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilledDurationSecs

`func (o *CallCostPayloadCostPartsInner) GetBilledDurationSecs() int32`

GetBilledDurationSecs returns the BilledDurationSecs field if non-nil, zero value otherwise.

### GetBilledDurationSecsOk

`func (o *CallCostPayloadCostPartsInner) GetBilledDurationSecsOk() (*int32, bool)`

GetBilledDurationSecsOk returns a tuple with the BilledDurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledDurationSecs

`func (o *CallCostPayloadCostPartsInner) SetBilledDurationSecs(v int32)`

SetBilledDurationSecs sets BilledDurationSecs field to given value.

### HasBilledDurationSecs

`func (o *CallCostPayloadCostPartsInner) HasBilledDurationSecs() bool`

HasBilledDurationSecs returns a boolean if a field has been set.

### GetCallPart

`func (o *CallCostPayloadCostPartsInner) GetCallPart() string`

GetCallPart returns the CallPart field if non-nil, zero value otherwise.

### GetCallPartOk

`func (o *CallCostPayloadCostPartsInner) GetCallPartOk() (*string, bool)`

GetCallPartOk returns a tuple with the CallPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPart

`func (o *CallCostPayloadCostPartsInner) SetCallPart(v string)`

SetCallPart sets CallPart field to given value.

### HasCallPart

`func (o *CallCostPayloadCostPartsInner) HasCallPart() bool`

HasCallPart returns a boolean if a field has been set.

### GetCost

`func (o *CallCostPayloadCostPartsInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CallCostPayloadCostPartsInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CallCostPayloadCostPartsInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CallCostPayloadCostPartsInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *CallCostPayloadCostPartsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CallCostPayloadCostPartsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CallCostPayloadCostPartsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CallCostPayloadCostPartsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRate

`func (o *CallCostPayloadCostPartsInner) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CallCostPayloadCostPartsInner) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CallCostPayloadCostPartsInner) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *CallCostPayloadCostPartsInner) HasRate() bool`

HasRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


