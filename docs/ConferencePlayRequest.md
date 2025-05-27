# ConferencePlayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | Pointer to **string** | The URL of a file to be played back in the conference. media_name and audio_url cannot be used together in one request. | [optional] 
**MediaName** | Pointer to **string** | The media_name of a file to be played back in the conference. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 
**Loop** | Pointer to [**Loopcount**](Loopcount.md) |  | [optional] 
**CallControlIds** | Pointer to **[]string** | List of call control ids identifying participants the audio file should be played to. If not given, the audio file will be played to the entire conference. | [optional] 

## Methods

### NewConferencePlayRequest

`func NewConferencePlayRequest() *ConferencePlayRequest`

NewConferencePlayRequest instantiates a new ConferencePlayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferencePlayRequestWithDefaults

`func NewConferencePlayRequestWithDefaults() *ConferencePlayRequest`

NewConferencePlayRequestWithDefaults instantiates a new ConferencePlayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *ConferencePlayRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *ConferencePlayRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *ConferencePlayRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *ConferencePlayRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *ConferencePlayRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *ConferencePlayRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *ConferencePlayRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *ConferencePlayRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetLoop

`func (o *ConferencePlayRequest) GetLoop() Loopcount`

GetLoop returns the Loop field if non-nil, zero value otherwise.

### GetLoopOk

`func (o *ConferencePlayRequest) GetLoopOk() (*Loopcount, bool)`

GetLoopOk returns a tuple with the Loop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoop

`func (o *ConferencePlayRequest) SetLoop(v Loopcount)`

SetLoop sets Loop field to given value.

### HasLoop

`func (o *ConferencePlayRequest) HasLoop() bool`

HasLoop returns a boolean if a field has been set.

### GetCallControlIds

`func (o *ConferencePlayRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferencePlayRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferencePlayRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.

### HasCallControlIds

`func (o *ConferencePlayRequest) HasCallControlIds() bool`

HasCallControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


