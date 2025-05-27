# CallRequestAnsweringMachineDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAnalysisTimeMillis** | Pointer to **int32** | Maximum timeout threshold for overall detection. | [optional] [default to 3500]
**AfterGreetingSilenceMillis** | Pointer to **int32** | Silence duration threshold after a greeting message or voice for it be considered human. | [optional] [default to 800]
**BetweenWordsSilenceMillis** | Pointer to **int32** | Maximum threshold for silence between words. | [optional] [default to 50]
**GreetingDurationMillis** | Pointer to **int32** | Maximum threshold of a human greeting. If greeting longer than this value, considered machine. | [optional] [default to 3500]
**InitialSilenceMillis** | Pointer to **int32** | If initial silence duration is greater than this value, consider it a machine. | [optional] [default to 3500]
**MaximumNumberOfWords** | Pointer to **int32** | If number of detected words is greater than this value, consder it a machine. | [optional] [default to 5]
**MaximumWordLengthMillis** | Pointer to **int32** | If a single word lasts longer than this threshold, consider it a machine. | [optional] [default to 3500]
**SilenceThreshold** | Pointer to **int32** | Minimum noise threshold for any analysis. | [optional] [default to 256]
**GreetingTotalAnalysisTimeMillis** | Pointer to **int32** | If machine already detected, maximum timeout threshold to determine the end of the machine greeting. | [optional] [default to 5000]
**GreetingSilenceDurationMillis** | Pointer to **int32** | If machine already detected, maximum threshold for silence between words. If exceeded, the greeting is considered ended. | [optional] [default to 1500]

## Methods

### NewCallRequestAnsweringMachineDetectionConfig

`func NewCallRequestAnsweringMachineDetectionConfig() *CallRequestAnsweringMachineDetectionConfig`

NewCallRequestAnsweringMachineDetectionConfig instantiates a new CallRequestAnsweringMachineDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestAnsweringMachineDetectionConfigWithDefaults

`func NewCallRequestAnsweringMachineDetectionConfigWithDefaults() *CallRequestAnsweringMachineDetectionConfig`

NewCallRequestAnsweringMachineDetectionConfigWithDefaults instantiates a new CallRequestAnsweringMachineDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAnalysisTimeMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetTotalAnalysisTimeMillis() int32`

GetTotalAnalysisTimeMillis returns the TotalAnalysisTimeMillis field if non-nil, zero value otherwise.

### GetTotalAnalysisTimeMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetTotalAnalysisTimeMillisOk() (*int32, bool)`

GetTotalAnalysisTimeMillisOk returns a tuple with the TotalAnalysisTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAnalysisTimeMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetTotalAnalysisTimeMillis(v int32)`

SetTotalAnalysisTimeMillis sets TotalAnalysisTimeMillis field to given value.

### HasTotalAnalysisTimeMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasTotalAnalysisTimeMillis() bool`

HasTotalAnalysisTimeMillis returns a boolean if a field has been set.

### GetAfterGreetingSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetAfterGreetingSilenceMillis() int32`

GetAfterGreetingSilenceMillis returns the AfterGreetingSilenceMillis field if non-nil, zero value otherwise.

### GetAfterGreetingSilenceMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetAfterGreetingSilenceMillisOk() (*int32, bool)`

GetAfterGreetingSilenceMillisOk returns a tuple with the AfterGreetingSilenceMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterGreetingSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetAfterGreetingSilenceMillis(v int32)`

SetAfterGreetingSilenceMillis sets AfterGreetingSilenceMillis field to given value.

### HasAfterGreetingSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasAfterGreetingSilenceMillis() bool`

HasAfterGreetingSilenceMillis returns a boolean if a field has been set.

### GetBetweenWordsSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetBetweenWordsSilenceMillis() int32`

GetBetweenWordsSilenceMillis returns the BetweenWordsSilenceMillis field if non-nil, zero value otherwise.

### GetBetweenWordsSilenceMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetBetweenWordsSilenceMillisOk() (*int32, bool)`

GetBetweenWordsSilenceMillisOk returns a tuple with the BetweenWordsSilenceMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetweenWordsSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetBetweenWordsSilenceMillis(v int32)`

SetBetweenWordsSilenceMillis sets BetweenWordsSilenceMillis field to given value.

### HasBetweenWordsSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasBetweenWordsSilenceMillis() bool`

HasBetweenWordsSilenceMillis returns a boolean if a field has been set.

### GetGreetingDurationMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetGreetingDurationMillis() int32`

GetGreetingDurationMillis returns the GreetingDurationMillis field if non-nil, zero value otherwise.

### GetGreetingDurationMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetGreetingDurationMillisOk() (*int32, bool)`

GetGreetingDurationMillisOk returns a tuple with the GreetingDurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingDurationMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetGreetingDurationMillis(v int32)`

SetGreetingDurationMillis sets GreetingDurationMillis field to given value.

### HasGreetingDurationMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasGreetingDurationMillis() bool`

HasGreetingDurationMillis returns a boolean if a field has been set.

### GetInitialSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetInitialSilenceMillis() int32`

GetInitialSilenceMillis returns the InitialSilenceMillis field if non-nil, zero value otherwise.

### GetInitialSilenceMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetInitialSilenceMillisOk() (*int32, bool)`

GetInitialSilenceMillisOk returns a tuple with the InitialSilenceMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetInitialSilenceMillis(v int32)`

SetInitialSilenceMillis sets InitialSilenceMillis field to given value.

### HasInitialSilenceMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasInitialSilenceMillis() bool`

HasInitialSilenceMillis returns a boolean if a field has been set.

### GetMaximumNumberOfWords

`func (o *CallRequestAnsweringMachineDetectionConfig) GetMaximumNumberOfWords() int32`

GetMaximumNumberOfWords returns the MaximumNumberOfWords field if non-nil, zero value otherwise.

### GetMaximumNumberOfWordsOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetMaximumNumberOfWordsOk() (*int32, bool)`

GetMaximumNumberOfWordsOk returns a tuple with the MaximumNumberOfWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNumberOfWords

`func (o *CallRequestAnsweringMachineDetectionConfig) SetMaximumNumberOfWords(v int32)`

SetMaximumNumberOfWords sets MaximumNumberOfWords field to given value.

### HasMaximumNumberOfWords

`func (o *CallRequestAnsweringMachineDetectionConfig) HasMaximumNumberOfWords() bool`

HasMaximumNumberOfWords returns a boolean if a field has been set.

### GetMaximumWordLengthMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetMaximumWordLengthMillis() int32`

GetMaximumWordLengthMillis returns the MaximumWordLengthMillis field if non-nil, zero value otherwise.

### GetMaximumWordLengthMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetMaximumWordLengthMillisOk() (*int32, bool)`

GetMaximumWordLengthMillisOk returns a tuple with the MaximumWordLengthMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumWordLengthMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetMaximumWordLengthMillis(v int32)`

SetMaximumWordLengthMillis sets MaximumWordLengthMillis field to given value.

### HasMaximumWordLengthMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasMaximumWordLengthMillis() bool`

HasMaximumWordLengthMillis returns a boolean if a field has been set.

### GetSilenceThreshold

`func (o *CallRequestAnsweringMachineDetectionConfig) GetSilenceThreshold() int32`

GetSilenceThreshold returns the SilenceThreshold field if non-nil, zero value otherwise.

### GetSilenceThresholdOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetSilenceThresholdOk() (*int32, bool)`

GetSilenceThresholdOk returns a tuple with the SilenceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilenceThreshold

`func (o *CallRequestAnsweringMachineDetectionConfig) SetSilenceThreshold(v int32)`

SetSilenceThreshold sets SilenceThreshold field to given value.

### HasSilenceThreshold

`func (o *CallRequestAnsweringMachineDetectionConfig) HasSilenceThreshold() bool`

HasSilenceThreshold returns a boolean if a field has been set.

### GetGreetingTotalAnalysisTimeMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetGreetingTotalAnalysisTimeMillis() int32`

GetGreetingTotalAnalysisTimeMillis returns the GreetingTotalAnalysisTimeMillis field if non-nil, zero value otherwise.

### GetGreetingTotalAnalysisTimeMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetGreetingTotalAnalysisTimeMillisOk() (*int32, bool)`

GetGreetingTotalAnalysisTimeMillisOk returns a tuple with the GreetingTotalAnalysisTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingTotalAnalysisTimeMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetGreetingTotalAnalysisTimeMillis(v int32)`

SetGreetingTotalAnalysisTimeMillis sets GreetingTotalAnalysisTimeMillis field to given value.

### HasGreetingTotalAnalysisTimeMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasGreetingTotalAnalysisTimeMillis() bool`

HasGreetingTotalAnalysisTimeMillis returns a boolean if a field has been set.

### GetGreetingSilenceDurationMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) GetGreetingSilenceDurationMillis() int32`

GetGreetingSilenceDurationMillis returns the GreetingSilenceDurationMillis field if non-nil, zero value otherwise.

### GetGreetingSilenceDurationMillisOk

`func (o *CallRequestAnsweringMachineDetectionConfig) GetGreetingSilenceDurationMillisOk() (*int32, bool)`

GetGreetingSilenceDurationMillisOk returns a tuple with the GreetingSilenceDurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingSilenceDurationMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) SetGreetingSilenceDurationMillis(v int32)`

SetGreetingSilenceDurationMillis sets GreetingSilenceDurationMillis field to given value.

### HasGreetingSilenceDurationMillis

`func (o *CallRequestAnsweringMachineDetectionConfig) HasGreetingSilenceDurationMillis() bool`

HasGreetingSilenceDurationMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


