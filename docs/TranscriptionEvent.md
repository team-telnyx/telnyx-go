# TranscriptionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Transcription**](Transcription.md) |  | [optional] 

## Methods

### NewTranscriptionEvent

`func NewTranscriptionEvent() *TranscriptionEvent`

NewTranscriptionEvent instantiates a new TranscriptionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionEventWithDefaults

`func NewTranscriptionEventWithDefaults() *TranscriptionEvent`

NewTranscriptionEventWithDefaults instantiates a new TranscriptionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TranscriptionEvent) GetData() Transcription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TranscriptionEvent) GetDataOk() (*Transcription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TranscriptionEvent) SetData(v Transcription)`

SetData sets Data field to given value.

### HasData

`func (o *TranscriptionEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


