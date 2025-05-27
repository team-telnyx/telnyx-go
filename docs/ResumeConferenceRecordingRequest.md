# ResumeConferenceRecordingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**RecordingId** | Pointer to **string** | Use this field to resume specific recording. | [optional] 

## Methods

### NewResumeConferenceRecordingRequest

`func NewResumeConferenceRecordingRequest() *ResumeConferenceRecordingRequest`

NewResumeConferenceRecordingRequest instantiates a new ResumeConferenceRecordingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeConferenceRecordingRequestWithDefaults

`func NewResumeConferenceRecordingRequestWithDefaults() *ResumeConferenceRecordingRequest`

NewResumeConferenceRecordingRequestWithDefaults instantiates a new ResumeConferenceRecordingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandId

`func (o *ResumeConferenceRecordingRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *ResumeConferenceRecordingRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *ResumeConferenceRecordingRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *ResumeConferenceRecordingRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetRecordingId

`func (o *ResumeConferenceRecordingRequest) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *ResumeConferenceRecordingRequest) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *ResumeConferenceRecordingRequest) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *ResumeConferenceRecordingRequest) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


