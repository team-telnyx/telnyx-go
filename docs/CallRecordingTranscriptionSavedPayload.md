# CallRecordingTranscriptionSavedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**CallingPartyType** | Pointer to **string** | The type of calling party connection. | [optional] 
**RecordingId** | Pointer to **string** | ID that is unique to the recording session and can be used to correlate webhook events. | [optional] 
**RecordingTranscriptionId** | Pointer to **string** | ID that is unique to the transcription process and can be used to correlate webhook events. | [optional] 
**Status** | Pointer to **string** | The transcription status. | [optional] 
**TranscriptionText** | Pointer to **string** | The transcribed text | [optional] 

## Methods

### NewCallRecordingTranscriptionSavedPayload

`func NewCallRecordingTranscriptionSavedPayload() *CallRecordingTranscriptionSavedPayload`

NewCallRecordingTranscriptionSavedPayload instantiates a new CallRecordingTranscriptionSavedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordingTranscriptionSavedPayloadWithDefaults

`func NewCallRecordingTranscriptionSavedPayloadWithDefaults() *CallRecordingTranscriptionSavedPayload`

NewCallRecordingTranscriptionSavedPayloadWithDefaults instantiates a new CallRecordingTranscriptionSavedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallRecordingTranscriptionSavedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallRecordingTranscriptionSavedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallRecordingTranscriptionSavedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallRecordingTranscriptionSavedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallRecordingTranscriptionSavedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallRecordingTranscriptionSavedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallRecordingTranscriptionSavedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallRecordingTranscriptionSavedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallRecordingTranscriptionSavedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallRecordingTranscriptionSavedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallRecordingTranscriptionSavedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallRecordingTranscriptionSavedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallRecordingTranscriptionSavedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallRecordingTranscriptionSavedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallRecordingTranscriptionSavedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallRecordingTranscriptionSavedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallRecordingTranscriptionSavedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallRecordingTranscriptionSavedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallRecordingTranscriptionSavedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallRecordingTranscriptionSavedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCallingPartyType

`func (o *CallRecordingTranscriptionSavedPayload) GetCallingPartyType() string`

GetCallingPartyType returns the CallingPartyType field if non-nil, zero value otherwise.

### GetCallingPartyTypeOk

`func (o *CallRecordingTranscriptionSavedPayload) GetCallingPartyTypeOk() (*string, bool)`

GetCallingPartyTypeOk returns a tuple with the CallingPartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingPartyType

`func (o *CallRecordingTranscriptionSavedPayload) SetCallingPartyType(v string)`

SetCallingPartyType sets CallingPartyType field to given value.

### HasCallingPartyType

`func (o *CallRecordingTranscriptionSavedPayload) HasCallingPartyType() bool`

HasCallingPartyType returns a boolean if a field has been set.

### GetRecordingId

`func (o *CallRecordingTranscriptionSavedPayload) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *CallRecordingTranscriptionSavedPayload) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *CallRecordingTranscriptionSavedPayload) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *CallRecordingTranscriptionSavedPayload) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.

### GetRecordingTranscriptionId

`func (o *CallRecordingTranscriptionSavedPayload) GetRecordingTranscriptionId() string`

GetRecordingTranscriptionId returns the RecordingTranscriptionId field if non-nil, zero value otherwise.

### GetRecordingTranscriptionIdOk

`func (o *CallRecordingTranscriptionSavedPayload) GetRecordingTranscriptionIdOk() (*string, bool)`

GetRecordingTranscriptionIdOk returns a tuple with the RecordingTranscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingTranscriptionId

`func (o *CallRecordingTranscriptionSavedPayload) SetRecordingTranscriptionId(v string)`

SetRecordingTranscriptionId sets RecordingTranscriptionId field to given value.

### HasRecordingTranscriptionId

`func (o *CallRecordingTranscriptionSavedPayload) HasRecordingTranscriptionId() bool`

HasRecordingTranscriptionId returns a boolean if a field has been set.

### GetStatus

`func (o *CallRecordingTranscriptionSavedPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallRecordingTranscriptionSavedPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallRecordingTranscriptionSavedPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallRecordingTranscriptionSavedPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTranscriptionText

`func (o *CallRecordingTranscriptionSavedPayload) GetTranscriptionText() string`

GetTranscriptionText returns the TranscriptionText field if non-nil, zero value otherwise.

### GetTranscriptionTextOk

`func (o *CallRecordingTranscriptionSavedPayload) GetTranscriptionTextOk() (*string, bool)`

GetTranscriptionTextOk returns a tuple with the TranscriptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionText

`func (o *CallRecordingTranscriptionSavedPayload) SetTranscriptionText(v string)`

SetTranscriptionText sets TranscriptionText field to given value.

### HasTranscriptionText

`func (o *CallRecordingTranscriptionSavedPayload) HasTranscriptionText() bool`

HasTranscriptionText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


