# CallCostMetaMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** | The number of attempts made to deliver the webhook | [optional] 
**DeliveredTo** | Pointer to **string** | The URL where webhook was sent | [optional] 

## Methods

### NewCallCostMetaMeta

`func NewCallCostMetaMeta() *CallCostMetaMeta`

NewCallCostMetaMeta instantiates a new CallCostMetaMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallCostMetaMetaWithDefaults

`func NewCallCostMetaMetaWithDefaults() *CallCostMetaMeta`

NewCallCostMetaMetaWithDefaults instantiates a new CallCostMetaMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *CallCostMetaMeta) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *CallCostMetaMeta) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *CallCostMetaMeta) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *CallCostMetaMeta) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetDeliveredTo

`func (o *CallCostMetaMeta) GetDeliveredTo() string`

GetDeliveredTo returns the DeliveredTo field if non-nil, zero value otherwise.

### GetDeliveredToOk

`func (o *CallCostMetaMeta) GetDeliveredToOk() (*string, bool)`

GetDeliveredToOk returns a tuple with the DeliveredTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredTo

`func (o *CallCostMetaMeta) SetDeliveredTo(v string)`

SetDeliveredTo sets DeliveredTo field to given value.

### HasDeliveredTo

`func (o *CallCostMetaMeta) HasDeliveredTo() bool`

HasDeliveredTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


