# SoundModifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pitch** | Pointer to **float64** | Set the pitch directly, value should be &gt; 0, default 1 (lower &#x3D; lower tone) | [optional] 
**Semitone** | Pointer to **float64** | Adjust the pitch in semitones, values should be between -14 and 14, default 0 | [optional] 
**Octaves** | Pointer to **float64** | Adjust the pitch in octaves, values should be between -1 and 1, default 0 | [optional] 
**Track** | Pointer to **string** | The track to which the sound modifications will be applied. Accepted values are &#x60;inbound&#x60; or &#x60;outbound&#x60; | [optional] [default to "outbound"]

## Methods

### NewSoundModifications

`func NewSoundModifications() *SoundModifications`

NewSoundModifications instantiates a new SoundModifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundModificationsWithDefaults

`func NewSoundModificationsWithDefaults() *SoundModifications`

NewSoundModificationsWithDefaults instantiates a new SoundModifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPitch

`func (o *SoundModifications) GetPitch() float64`

GetPitch returns the Pitch field if non-nil, zero value otherwise.

### GetPitchOk

`func (o *SoundModifications) GetPitchOk() (*float64, bool)`

GetPitchOk returns a tuple with the Pitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitch

`func (o *SoundModifications) SetPitch(v float64)`

SetPitch sets Pitch field to given value.

### HasPitch

`func (o *SoundModifications) HasPitch() bool`

HasPitch returns a boolean if a field has been set.

### GetSemitone

`func (o *SoundModifications) GetSemitone() float64`

GetSemitone returns the Semitone field if non-nil, zero value otherwise.

### GetSemitoneOk

`func (o *SoundModifications) GetSemitoneOk() (*float64, bool)`

GetSemitoneOk returns a tuple with the Semitone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemitone

`func (o *SoundModifications) SetSemitone(v float64)`

SetSemitone sets Semitone field to given value.

### HasSemitone

`func (o *SoundModifications) HasSemitone() bool`

HasSemitone returns a boolean if a field has been set.

### GetOctaves

`func (o *SoundModifications) GetOctaves() float64`

GetOctaves returns the Octaves field if non-nil, zero value otherwise.

### GetOctavesOk

`func (o *SoundModifications) GetOctavesOk() (*float64, bool)`

GetOctavesOk returns a tuple with the Octaves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctaves

`func (o *SoundModifications) SetOctaves(v float64)`

SetOctaves sets Octaves field to given value.

### HasOctaves

`func (o *SoundModifications) HasOctaves() bool`

HasOctaves returns a boolean if a field has been set.

### GetTrack

`func (o *SoundModifications) GetTrack() string`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *SoundModifications) GetTrackOk() (*string, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *SoundModifications) SetTrack(v string)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *SoundModifications) HasTrack() bool`

HasTrack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


