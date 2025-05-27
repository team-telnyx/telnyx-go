# GatherUsingAudioRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | Pointer to **string** | The URL of a file to be played back at the beginning of each prompt. The URL can point to either a WAV or MP3 file. media_name and audio_url cannot be used together in one request. | [optional] 
**MediaName** | Pointer to **string** | The media_name of a file to be played back at the beginning of each prompt. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 
**InvalidAudioUrl** | Pointer to **string** | The URL of a file to play when digits don&#39;t match the &#x60;valid_digits&#x60; parameter or the number of digits is not between &#x60;min&#x60; and &#x60;max&#x60;. The URL can point to either a WAV or MP3 file. invalid_media_name and invalid_audio_url cannot be used together in one request. | [optional] 
**InvalidMediaName** | Pointer to **string** | The media_name of a file to be played back when digits don&#39;t match the &#x60;valid_digits&#x60; parameter or the number of digits is not between &#x60;min&#x60; and &#x60;max&#x60;. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 
**MinimumDigits** | Pointer to **int32** | The minimum number of digits to fetch. This parameter has a minimum value of 1. | [optional] [default to 1]
**MaximumDigits** | Pointer to **int32** | The maximum number of digits to fetch. This parameter has a maximum value of 128. | [optional] [default to 128]
**MaximumTries** | Pointer to **int32** | The maximum number of times the file should be played if there is no input from the user on the call. | [optional] [default to 3]
**TimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait for a DTMF response after file playback ends before a replaying the sound file. | [optional] [default to 60000]
**TerminatingDigit** | Pointer to **string** | The digit used to terminate input if fewer than &#x60;maximum_digits&#x60; digits have been gathered. | [optional] [default to "#"]
**ValidDigits** | Pointer to **string** | A list of all digits accepted as valid. | [optional] [default to "0123456789#*"]
**InterDigitTimeoutMillis** | Pointer to **int32** | The number of milliseconds to wait for input between digits. | [optional] [default to 5000]
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 

## Methods

### NewGatherUsingAudioRequest

`func NewGatherUsingAudioRequest() *GatherUsingAudioRequest`

NewGatherUsingAudioRequest instantiates a new GatherUsingAudioRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatherUsingAudioRequestWithDefaults

`func NewGatherUsingAudioRequestWithDefaults() *GatherUsingAudioRequest`

NewGatherUsingAudioRequestWithDefaults instantiates a new GatherUsingAudioRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *GatherUsingAudioRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *GatherUsingAudioRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *GatherUsingAudioRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *GatherUsingAudioRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *GatherUsingAudioRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *GatherUsingAudioRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *GatherUsingAudioRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *GatherUsingAudioRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetInvalidAudioUrl

`func (o *GatherUsingAudioRequest) GetInvalidAudioUrl() string`

GetInvalidAudioUrl returns the InvalidAudioUrl field if non-nil, zero value otherwise.

### GetInvalidAudioUrlOk

`func (o *GatherUsingAudioRequest) GetInvalidAudioUrlOk() (*string, bool)`

GetInvalidAudioUrlOk returns a tuple with the InvalidAudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidAudioUrl

`func (o *GatherUsingAudioRequest) SetInvalidAudioUrl(v string)`

SetInvalidAudioUrl sets InvalidAudioUrl field to given value.

### HasInvalidAudioUrl

`func (o *GatherUsingAudioRequest) HasInvalidAudioUrl() bool`

HasInvalidAudioUrl returns a boolean if a field has been set.

### GetInvalidMediaName

`func (o *GatherUsingAudioRequest) GetInvalidMediaName() string`

GetInvalidMediaName returns the InvalidMediaName field if non-nil, zero value otherwise.

### GetInvalidMediaNameOk

`func (o *GatherUsingAudioRequest) GetInvalidMediaNameOk() (*string, bool)`

GetInvalidMediaNameOk returns a tuple with the InvalidMediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidMediaName

`func (o *GatherUsingAudioRequest) SetInvalidMediaName(v string)`

SetInvalidMediaName sets InvalidMediaName field to given value.

### HasInvalidMediaName

`func (o *GatherUsingAudioRequest) HasInvalidMediaName() bool`

HasInvalidMediaName returns a boolean if a field has been set.

### GetMinimumDigits

`func (o *GatherUsingAudioRequest) GetMinimumDigits() int32`

GetMinimumDigits returns the MinimumDigits field if non-nil, zero value otherwise.

### GetMinimumDigitsOk

`func (o *GatherUsingAudioRequest) GetMinimumDigitsOk() (*int32, bool)`

GetMinimumDigitsOk returns a tuple with the MinimumDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDigits

`func (o *GatherUsingAudioRequest) SetMinimumDigits(v int32)`

SetMinimumDigits sets MinimumDigits field to given value.

### HasMinimumDigits

`func (o *GatherUsingAudioRequest) HasMinimumDigits() bool`

HasMinimumDigits returns a boolean if a field has been set.

### GetMaximumDigits

`func (o *GatherUsingAudioRequest) GetMaximumDigits() int32`

GetMaximumDigits returns the MaximumDigits field if non-nil, zero value otherwise.

### GetMaximumDigitsOk

`func (o *GatherUsingAudioRequest) GetMaximumDigitsOk() (*int32, bool)`

GetMaximumDigitsOk returns a tuple with the MaximumDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDigits

`func (o *GatherUsingAudioRequest) SetMaximumDigits(v int32)`

SetMaximumDigits sets MaximumDigits field to given value.

### HasMaximumDigits

`func (o *GatherUsingAudioRequest) HasMaximumDigits() bool`

HasMaximumDigits returns a boolean if a field has been set.

### GetMaximumTries

`func (o *GatherUsingAudioRequest) GetMaximumTries() int32`

GetMaximumTries returns the MaximumTries field if non-nil, zero value otherwise.

### GetMaximumTriesOk

`func (o *GatherUsingAudioRequest) GetMaximumTriesOk() (*int32, bool)`

GetMaximumTriesOk returns a tuple with the MaximumTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTries

`func (o *GatherUsingAudioRequest) SetMaximumTries(v int32)`

SetMaximumTries sets MaximumTries field to given value.

### HasMaximumTries

`func (o *GatherUsingAudioRequest) HasMaximumTries() bool`

HasMaximumTries returns a boolean if a field has been set.

### GetTimeoutMillis

`func (o *GatherUsingAudioRequest) GetTimeoutMillis() int32`

GetTimeoutMillis returns the TimeoutMillis field if non-nil, zero value otherwise.

### GetTimeoutMillisOk

`func (o *GatherUsingAudioRequest) GetTimeoutMillisOk() (*int32, bool)`

GetTimeoutMillisOk returns a tuple with the TimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMillis

`func (o *GatherUsingAudioRequest) SetTimeoutMillis(v int32)`

SetTimeoutMillis sets TimeoutMillis field to given value.

### HasTimeoutMillis

`func (o *GatherUsingAudioRequest) HasTimeoutMillis() bool`

HasTimeoutMillis returns a boolean if a field has been set.

### GetTerminatingDigit

`func (o *GatherUsingAudioRequest) GetTerminatingDigit() string`

GetTerminatingDigit returns the TerminatingDigit field if non-nil, zero value otherwise.

### GetTerminatingDigitOk

`func (o *GatherUsingAudioRequest) GetTerminatingDigitOk() (*string, bool)`

GetTerminatingDigitOk returns a tuple with the TerminatingDigit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatingDigit

`func (o *GatherUsingAudioRequest) SetTerminatingDigit(v string)`

SetTerminatingDigit sets TerminatingDigit field to given value.

### HasTerminatingDigit

`func (o *GatherUsingAudioRequest) HasTerminatingDigit() bool`

HasTerminatingDigit returns a boolean if a field has been set.

### GetValidDigits

`func (o *GatherUsingAudioRequest) GetValidDigits() string`

GetValidDigits returns the ValidDigits field if non-nil, zero value otherwise.

### GetValidDigitsOk

`func (o *GatherUsingAudioRequest) GetValidDigitsOk() (*string, bool)`

GetValidDigitsOk returns a tuple with the ValidDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDigits

`func (o *GatherUsingAudioRequest) SetValidDigits(v string)`

SetValidDigits sets ValidDigits field to given value.

### HasValidDigits

`func (o *GatherUsingAudioRequest) HasValidDigits() bool`

HasValidDigits returns a boolean if a field has been set.

### GetInterDigitTimeoutMillis

`func (o *GatherUsingAudioRequest) GetInterDigitTimeoutMillis() int32`

GetInterDigitTimeoutMillis returns the InterDigitTimeoutMillis field if non-nil, zero value otherwise.

### GetInterDigitTimeoutMillisOk

`func (o *GatherUsingAudioRequest) GetInterDigitTimeoutMillisOk() (*int32, bool)`

GetInterDigitTimeoutMillisOk returns a tuple with the InterDigitTimeoutMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterDigitTimeoutMillis

`func (o *GatherUsingAudioRequest) SetInterDigitTimeoutMillis(v int32)`

SetInterDigitTimeoutMillis sets InterDigitTimeoutMillis field to given value.

### HasInterDigitTimeoutMillis

`func (o *GatherUsingAudioRequest) HasInterDigitTimeoutMillis() bool`

HasInterDigitTimeoutMillis returns a boolean if a field has been set.

### GetClientState

`func (o *GatherUsingAudioRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *GatherUsingAudioRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *GatherUsingAudioRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *GatherUsingAudioRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *GatherUsingAudioRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *GatherUsingAudioRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *GatherUsingAudioRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *GatherUsingAudioRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


