# SIMCardActionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The current status of the SIM card action. | [optional] [readonly] 
**Reason** | Pointer to **string** | It describes why the SIM card action is in the current status. This will be &lt;code&gt;null&lt;/code&gt; for self-explanatory statuses, such as &lt;code&gt;in-progress&lt;/code&gt; and &lt;code&gt;completed&lt;/code&gt; but will include further information on statuses like &lt;code&gt;interrupted&lt;/code&gt; and &lt;code&gt;failed&lt;/code&gt;. | [optional] [readonly] 

## Methods

### NewSIMCardActionStatus

`func NewSIMCardActionStatus() *SIMCardActionStatus`

NewSIMCardActionStatus instantiates a new SIMCardActionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardActionStatusWithDefaults

`func NewSIMCardActionStatusWithDefaults() *SIMCardActionStatus`

NewSIMCardActionStatusWithDefaults instantiates a new SIMCardActionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SIMCardActionStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SIMCardActionStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SIMCardActionStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SIMCardActionStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetReason

`func (o *SIMCardActionStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SIMCardActionStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SIMCardActionStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SIMCardActionStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


