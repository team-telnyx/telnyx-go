# SendDTMFRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digits** | **string** | DTMF digits to send. Valid digits are 0-9, A-D, *, and #. Pauses can be added using w (0.5s) and W (1s). | 
**DurationMillis** | Pointer to **int32** | Specifies for how many milliseconds each digit will be played in the audio stream. Ranges from 100 to 500ms | [optional] [default to 250]
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewSendDTMFRequest

`func NewSendDTMFRequest(digits string, ) *SendDTMFRequest`

NewSendDTMFRequest instantiates a new SendDTMFRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDTMFRequestWithDefaults

`func NewSendDTMFRequestWithDefaults() *SendDTMFRequest`

NewSendDTMFRequestWithDefaults instantiates a new SendDTMFRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigits

`func (o *SendDTMFRequest) GetDigits() string`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *SendDTMFRequest) GetDigitsOk() (*string, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *SendDTMFRequest) SetDigits(v string)`

SetDigits sets Digits field to given value.


### GetDurationMillis

`func (o *SendDTMFRequest) GetDurationMillis() int32`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *SendDTMFRequest) GetDurationMillisOk() (*int32, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *SendDTMFRequest) SetDurationMillis(v int32)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *SendDTMFRequest) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetClientState

`func (o *SendDTMFRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *SendDTMFRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *SendDTMFRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *SendDTMFRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *SendDTMFRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *SendDTMFRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *SendDTMFRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *SendDTMFRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


