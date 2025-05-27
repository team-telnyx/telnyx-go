# NoiseSuppressionStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**Direction** | Pointer to [**NoiseSuppressionDirection**](NoiseSuppressionDirection.md) |  | [optional] [default to INBOUND]

## Methods

### NewNoiseSuppressionStart

`func NewNoiseSuppressionStart() *NoiseSuppressionStart`

NewNoiseSuppressionStart instantiates a new NoiseSuppressionStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoiseSuppressionStartWithDefaults

`func NewNoiseSuppressionStartWithDefaults() *NoiseSuppressionStart`

NewNoiseSuppressionStartWithDefaults instantiates a new NoiseSuppressionStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientState

`func (o *NoiseSuppressionStart) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *NoiseSuppressionStart) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *NoiseSuppressionStart) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *NoiseSuppressionStart) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *NoiseSuppressionStart) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *NoiseSuppressionStart) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *NoiseSuppressionStart) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *NoiseSuppressionStart) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetDirection

`func (o *NoiseSuppressionStart) GetDirection() NoiseSuppressionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NoiseSuppressionStart) GetDirectionOk() (*NoiseSuppressionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NoiseSuppressionStart) SetDirection(v NoiseSuppressionDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NoiseSuppressionStart) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


