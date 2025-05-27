# CallRecordingTranscriptionSaved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**OccurredAt** | Pointer to **time.Time** | ISO 8601 datetime of when the event occurred. | [optional] 
**Payload** | Pointer to [**CallRecordingTranscriptionSavedPayload**](CallRecordingTranscriptionSavedPayload.md) |  | [optional] 

## Methods

### NewCallRecordingTranscriptionSaved

`func NewCallRecordingTranscriptionSaved() *CallRecordingTranscriptionSaved`

NewCallRecordingTranscriptionSaved instantiates a new CallRecordingTranscriptionSaved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRecordingTranscriptionSavedWithDefaults

`func NewCallRecordingTranscriptionSavedWithDefaults() *CallRecordingTranscriptionSaved`

NewCallRecordingTranscriptionSavedWithDefaults instantiates a new CallRecordingTranscriptionSaved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallRecordingTranscriptionSaved) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallRecordingTranscriptionSaved) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallRecordingTranscriptionSaved) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CallRecordingTranscriptionSaved) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *CallRecordingTranscriptionSaved) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CallRecordingTranscriptionSaved) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CallRecordingTranscriptionSaved) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CallRecordingTranscriptionSaved) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *CallRecordingTranscriptionSaved) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallRecordingTranscriptionSaved) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallRecordingTranscriptionSaved) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CallRecordingTranscriptionSaved) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOccurredAt

`func (o *CallRecordingTranscriptionSaved) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *CallRecordingTranscriptionSaved) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *CallRecordingTranscriptionSaved) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *CallRecordingTranscriptionSaved) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetPayload

`func (o *CallRecordingTranscriptionSaved) GetPayload() CallRecordingTranscriptionSavedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CallRecordingTranscriptionSaved) GetPayloadOk() (*CallRecordingTranscriptionSavedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CallRecordingTranscriptionSaved) SetPayload(v CallRecordingTranscriptionSavedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CallRecordingTranscriptionSaved) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


