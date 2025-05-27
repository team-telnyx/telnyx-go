# RecordingTranscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**DurationMillis** | Pointer to **int32** | The duration of the recording transcription in milliseconds. | [optional] 
**Id** | Pointer to **string** | Uniquely identifies the recording transcription. | [optional] 
**RecordingId** | Pointer to **string** | Uniquely identifies the recording associated with this transcription. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | The status of the recording transcription. Only &#x60;completed&#x60; has transcription text available. | [optional] 
**TranscriptionText** | Pointer to **string** | The recording&#39;s transcribed text. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewRecordingTranscription

`func NewRecordingTranscription() *RecordingTranscription`

NewRecordingTranscription instantiates a new RecordingTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingTranscriptionWithDefaults

`func NewRecordingTranscriptionWithDefaults() *RecordingTranscription`

NewRecordingTranscriptionWithDefaults instantiates a new RecordingTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RecordingTranscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RecordingTranscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RecordingTranscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RecordingTranscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationMillis

`func (o *RecordingTranscription) GetDurationMillis() int32`

GetDurationMillis returns the DurationMillis field if non-nil, zero value otherwise.

### GetDurationMillisOk

`func (o *RecordingTranscription) GetDurationMillisOk() (*int32, bool)`

GetDurationMillisOk returns a tuple with the DurationMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMillis

`func (o *RecordingTranscription) SetDurationMillis(v int32)`

SetDurationMillis sets DurationMillis field to given value.

### HasDurationMillis

`func (o *RecordingTranscription) HasDurationMillis() bool`

HasDurationMillis returns a boolean if a field has been set.

### GetId

`func (o *RecordingTranscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordingTranscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordingTranscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecordingTranscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordingId

`func (o *RecordingTranscription) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *RecordingTranscription) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *RecordingTranscription) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *RecordingTranscription) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.

### GetRecordType

`func (o *RecordingTranscription) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RecordingTranscription) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RecordingTranscription) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RecordingTranscription) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStatus

`func (o *RecordingTranscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RecordingTranscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RecordingTranscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RecordingTranscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTranscriptionText

`func (o *RecordingTranscription) GetTranscriptionText() string`

GetTranscriptionText returns the TranscriptionText field if non-nil, zero value otherwise.

### GetTranscriptionTextOk

`func (o *RecordingTranscription) GetTranscriptionTextOk() (*string, bool)`

GetTranscriptionTextOk returns a tuple with the TranscriptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionText

`func (o *RecordingTranscription) SetTranscriptionText(v string)`

SetTranscriptionText sets TranscriptionText field to given value.

### HasTranscriptionText

`func (o *RecordingTranscription) HasTranscriptionText() bool`

HasTranscriptionText returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RecordingTranscription) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RecordingTranscription) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RecordingTranscription) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RecordingTranscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


