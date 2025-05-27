# AudioTranscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The transcribed text for the audio file. | 
**Duration** | Pointer to **float32** | The duration of the audio file in seconds. This is only included if &#x60;response_format&#x60; is set to &#x60;verbose_json&#x60;. | [optional] 
**Segments** | Pointer to [**[]AudioTranscriptionResponseSegments**](AudioTranscriptionResponseSegments.md) | Segments of the transcribed text and their corresponding details. This is only included if &#x60;response_format&#x60; is set to &#x60;verbose_json&#x60;. | [optional] 

## Methods

### NewAudioTranscriptionResponse

`func NewAudioTranscriptionResponse(text string, ) *AudioTranscriptionResponse`

NewAudioTranscriptionResponse instantiates a new AudioTranscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioTranscriptionResponseWithDefaults

`func NewAudioTranscriptionResponseWithDefaults() *AudioTranscriptionResponse`

NewAudioTranscriptionResponseWithDefaults instantiates a new AudioTranscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *AudioTranscriptionResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AudioTranscriptionResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AudioTranscriptionResponse) SetText(v string)`

SetText sets Text field to given value.


### GetDuration

`func (o *AudioTranscriptionResponse) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AudioTranscriptionResponse) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AudioTranscriptionResponse) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AudioTranscriptionResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSegments

`func (o *AudioTranscriptionResponse) GetSegments() []AudioTranscriptionResponseSegments`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *AudioTranscriptionResponse) GetSegmentsOk() (*[]AudioTranscriptionResponseSegments, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *AudioTranscriptionResponse) SetSegments(v []AudioTranscriptionResponseSegments)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *AudioTranscriptionResponse) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


