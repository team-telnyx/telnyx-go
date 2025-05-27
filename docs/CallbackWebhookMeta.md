# CallbackWebhookMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int32** | The number of times the callback webhook has been attempted. | [optional] 
**DeliveredTo** | Pointer to **string** | The URL that the callback webhook was delivered to. | [optional] 

## Methods

### NewCallbackWebhookMeta

`func NewCallbackWebhookMeta() *CallbackWebhookMeta`

NewCallbackWebhookMeta instantiates a new CallbackWebhookMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackWebhookMetaWithDefaults

`func NewCallbackWebhookMetaWithDefaults() *CallbackWebhookMeta`

NewCallbackWebhookMetaWithDefaults instantiates a new CallbackWebhookMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *CallbackWebhookMeta) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *CallbackWebhookMeta) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *CallbackWebhookMeta) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *CallbackWebhookMeta) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetDeliveredTo

`func (o *CallbackWebhookMeta) GetDeliveredTo() string`

GetDeliveredTo returns the DeliveredTo field if non-nil, zero value otherwise.

### GetDeliveredToOk

`func (o *CallbackWebhookMeta) GetDeliveredToOk() (*string, bool)`

GetDeliveredToOk returns a tuple with the DeliveredTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredTo

`func (o *CallbackWebhookMeta) SetDeliveredTo(v string)`

SetDeliveredTo sets DeliveredTo field to given value.

### HasDeliveredTo

`func (o *CallbackWebhookMeta) HasDeliveredTo() bool`

HasDeliveredTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


