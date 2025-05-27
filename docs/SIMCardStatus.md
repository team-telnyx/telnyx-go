# SIMCardStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The current status of the SIM card. It will be one of the following: &lt;br/&gt; &lt;ul&gt;  &lt;li&gt;&lt;code&gt;registering&lt;/code&gt; - the card is being registered&lt;/li&gt;  &lt;li&gt;&lt;code&gt;enabling&lt;/code&gt; - the card is being enabled&lt;/li&gt;  &lt;li&gt;&lt;code&gt;enabled&lt;/code&gt; - the card is enabled and ready for use&lt;/li&gt;  &lt;li&gt;&lt;code&gt;disabling&lt;/code&gt; - the card is being disabled&lt;/li&gt;  &lt;li&gt;&lt;code&gt;disabled&lt;/code&gt; - the card has been disabled and cannot be used&lt;/li&gt;  &lt;li&gt;&lt;code&gt;data_limit_exceeded&lt;/code&gt; - the card has exceeded its data consumption limit&lt;/li&gt;  &lt;li&gt;&lt;code&gt;setting_standby&lt;/code&gt; - the process to set the card in stand by is in progress&lt;/li&gt;  &lt;li&gt;&lt;code&gt;standby&lt;/code&gt; - the card is in stand by&lt;/li&gt; &lt;/ul&gt; Transitioning between the enabled and disabled states may take a period of time.  | [optional] [readonly] 
**Reason** | Pointer to **string** | It describes why the SIM card is in the current status. | [optional] [readonly] 

## Methods

### NewSIMCardStatus

`func NewSIMCardStatus() *SIMCardStatus`

NewSIMCardStatus instantiates a new SIMCardStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardStatusWithDefaults

`func NewSIMCardStatusWithDefaults() *SIMCardStatus`

NewSIMCardStatusWithDefaults instantiates a new SIMCardStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SIMCardStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SIMCardStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SIMCardStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SIMCardStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetReason

`func (o *SIMCardStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SIMCardStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SIMCardStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SIMCardStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


