# TranscriptionPayloadTranscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidence** | Pointer to **float64** | Speech recognition confidence level. | [optional] 
**IsFinal** | Pointer to **bool** | When false, it means that this is an interim result. | [optional] 
**Transcript** | Pointer to **string** | Recognized text. | [optional] 
**TranscriptionTrack** | Pointer to **string** | Indicates which leg of the call has been transcribed. This is only available when &#x60;transcription_engine&#x60; is set to &#x60;B&#x60;. | [optional] 

## Methods

### NewTranscriptionPayloadTranscriptionData

`func NewTranscriptionPayloadTranscriptionData() *TranscriptionPayloadTranscriptionData`

NewTranscriptionPayloadTranscriptionData instantiates a new TranscriptionPayloadTranscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionPayloadTranscriptionDataWithDefaults

`func NewTranscriptionPayloadTranscriptionDataWithDefaults() *TranscriptionPayloadTranscriptionData`

NewTranscriptionPayloadTranscriptionDataWithDefaults instantiates a new TranscriptionPayloadTranscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidence

`func (o *TranscriptionPayloadTranscriptionData) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *TranscriptionPayloadTranscriptionData) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *TranscriptionPayloadTranscriptionData) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *TranscriptionPayloadTranscriptionData) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetIsFinal

`func (o *TranscriptionPayloadTranscriptionData) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *TranscriptionPayloadTranscriptionData) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *TranscriptionPayloadTranscriptionData) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *TranscriptionPayloadTranscriptionData) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetTranscript

`func (o *TranscriptionPayloadTranscriptionData) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *TranscriptionPayloadTranscriptionData) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *TranscriptionPayloadTranscriptionData) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *TranscriptionPayloadTranscriptionData) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### GetTranscriptionTrack

`func (o *TranscriptionPayloadTranscriptionData) GetTranscriptionTrack() string`

GetTranscriptionTrack returns the TranscriptionTrack field if non-nil, zero value otherwise.

### GetTranscriptionTrackOk

`func (o *TranscriptionPayloadTranscriptionData) GetTranscriptionTrackOk() (*string, bool)`

GetTranscriptionTrackOk returns a tuple with the TranscriptionTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionTrack

`func (o *TranscriptionPayloadTranscriptionData) SetTranscriptionTrack(v string)`

SetTranscriptionTrack sets TranscriptionTrack field to given value.

### HasTranscriptionTrack

`func (o *TranscriptionPayloadTranscriptionData) HasTranscriptionTrack() bool`

HasTranscriptionTrack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


