# GatherRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinimumDigits** | Pointer to **int32** | The minimum number of digits to fetch. This parameter has a minimum value of 1. | [optional] [default to 1]
**MaximumDigits** | Pointer to **int32** | The maximum number of digits to fetch. This parameter has a maximum value of 128. | [optional] [default to 128]
**TimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait to complete the request. | [optional] [default to 60000]
**InterDigitTimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait for input between digits. | [optional] [default to 5000]
**InitialTimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait for the first DTMF. | [optional] [default to 5000]
**TerminatingDigit** | Pointer to **string** | The digit used to terminate input if fewer than &#x60;maximum_digits&#x60; digits have been gathered. | [optional] [default to "#"]
**ValidDigits** | Pointer to **string** | A list of all digits accepted as valid. | [optional] [default to "0123456789#*"]
**GatherId** | Pointer to **string** | An id that will be sent back in the corresponding &#x60;call.gather.ended&#x60; webhook. Will be randomly generated if not specified. | [optional] 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewGatherRequest

`func NewGatherRequest() *GatherRequest`

NewGatherRequest instantiates a new GatherRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatherRequestWithDefaults

`func NewGatherRequestWithDefaults() *GatherRequest`

NewGatherRequestWithDefaults instantiates a new GatherRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimumDigits

`func (o *GatherRequest) GetMinimumDigits() int32`

GetMinimumDigits returns the MinimumDigits field if non-nil, zero value otherwise.

### GetMinimumDigitsOk

`func (o *GatherRequest) GetMinimumDigitsOk() (*int32, bool)`

GetMinimumDigitsOk returns a tuple with the MinimumDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDigits

`func (o *GatherRequest) SetMinimumDigits(v int32)`

SetMinimumDigits sets MinimumDigits field to given value.

### HasMinimumDigits

`func (o *GatherRequest) HasMinimumDigits() bool`

HasMinimumDigits returns a boolean if a field has been set.

### GetMaximumDigits

`func (o *GatherRequest) GetMaximumDigits() int32`

GetMaximumDigits returns the MaximumDigits field if non-nil, zero value otherwise.

### GetMaximumDigitsOk

`func (o *GatherRequest) GetMaximumDigitsOk() (*int32, bool)`

GetMaximumDigitsOk returns a tuple with the MaximumDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDigits

`func (o *GatherRequest) SetMaximumDigits(v int32)`

SetMaximumDigits sets MaximumDigits field to given value.

### HasMaximumDigits

`func (o *GatherRequest) HasMaximumDigits() bool`

HasMaximumDigits returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *GatherRequest) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *GatherRequest) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *GatherRequest) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *GatherRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetInterDigitTimeoutMillis

`func (o *GatherRequest) GetInterDigitTimeoutMillis() int32`

GetInterDigitTimeoutMillis returns the InterDigitTimeoutMillis field if non-nil, zero value otherwise.

### GetInterDigitTimeoutMillisOk

`func (o *GatherRequest) GetInterDigitTimeoutMillisOk() (*int32, bool)`

GetInterDigitTimeoutMillisOk returns a tuple with the InterDigitTimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterDigitTimeoutMillis

`func (o *GatherRequest) SetInterDigitTimeoutMillis(v int32)`

SetInterDigitTimeoutMillis sets InterDigitTimeoutMillis field to given value.

### HasInterDigitTimeoutMillis

`func (o *GatherRequest) HasInterDigitTimeoutMillis() bool`

HasInterDigitTimeoutMillis returns a boolean if a field has been set.

### GetInitialTimeoutMillis

`func (o *GatherRequest) GetInitialTimeoutMillis() int32`

GetInitialTimeoutMillis returns the InitialTimeoutMillis field if non-nil, zero value otherwise.

### GetInitialTimeoutMillisOk

`func (o *GatherRequest) GetInitialTimeoutMillisOk() (*int32, bool)`

GetInitialTimeoutMillisOk returns a tuple with the InitialTimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialTimeoutMillis

`func (o *GatherRequest) SetInitialTimeoutMillis(v int32)`

SetInitialTimeoutMillis sets InitialTimeoutMillis field to given value.

### HasInitialTimeoutMillis

`func (o *GatherRequest) HasInitialTimeoutMillis() bool`

HasInitialTimeoutMillis returns a boolean if a field has been set.

### GetTerminatingDigit

`func (o *GatherRequest) GetTerminatingDigit() string`

GetTerminatingDigit returns the TerminatingDigit field if non-nil, zero value otherwise.

### GetTerminatingDigitOk

`func (o *GatherRequest) GetTerminatingDigitOk() (*string, bool)`

GetTerminatingDigitOk returns a tuple with the TerminatingDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingDigit

`func (o *GatherRequest) SetTerminatingDigit(v string)`

SetTerminatingDigit sets TerminatingDigit field to given value.

### HasTerminatingDigit

`func (o *GatherRequest) HasTerminatingDigit() bool`

HasTerminatingDigit returns a boolean if a field has been set.

### GetValidDigits

`func (o *GatherRequest) GetValidDigits() string`

GetValidDigits returns the ValidDigits field if non-nil, zero value otherwise.

### GetValidDigitsOk

`func (o *GatherRequest) GetValidDigitsOk() (*string, bool)`

GetValidDigitsOk returns a tuple with the ValidDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDigits

`func (o *GatherRequest) SetValidDigits(v string)`

SetValidDigits sets ValidDigits field to given value.

### HasValidDigits

`func (o *GatherRequest) HasValidDigits() bool`

HasValidDigits returns a boolean if a field has been set.

### GetGatherId

`func (o *GatherRequest) GetGatherId() string`

GetGatherId returns the GatherId field if non-nil, zero value otherwise.

### GetGatherIdOk

`func (o *GatherRequest) GetGatherIdOk() (*string, bool)`

GetGatherIdOk returns a tuple with the GatherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatherId

`func (o *GatherRequest) SetGatherId(v string)`

SetGatherId sets GatherId field to given value.

### HasGatherId

`func (o *GatherRequest) HasGatherId() bool`

HasGatherId returns a boolean if a field has been set.

### GetClientState

`func (o *GatherRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *GatherRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *GatherRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *GatherRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *GatherRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *GatherRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *GatherRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *GatherRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


