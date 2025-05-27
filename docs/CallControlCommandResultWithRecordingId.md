# CallControlCommandResultWithRecordingId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** |  | [optional] 
**RecordingId** | Pointer to **string** | The ID of the recording. Only present when the record parameter is set to record-from-answer. | [optional] 

## Methods

### NewCallControlCommandResultWithRecordingId

`func NewCallControlCommandResultWithRecordingId() *CallControlCommandResultWithRecordingId`

NewCallControlCommandResultWithRecordingId instantiates a new CallControlCommandResultWithRecordingId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallControlCommandResultWithRecordingIdWithDefaults

`func NewCallControlCommandResultWithRecordingIdWithDefaults() *CallControlCommandResultWithRecordingId`

NewCallControlCommandResultWithRecordingIdWithDefaults instantiates a new CallControlCommandResultWithRecordingId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CallControlCommandResultWithRecordingId) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CallControlCommandResultWithRecordingId) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CallControlCommandResultWithRecordingId) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CallControlCommandResultWithRecordingId) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRecordingId

`func (o *CallControlCommandResultWithRecordingId) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *CallControlCommandResultWithRecordingId) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *CallControlCommandResultWithRecordingId) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *CallControlCommandResultWithRecordingId) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


