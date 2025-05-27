# ConferenceHoldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | Pointer to **[]string** | List of unique identifiers and tokens for controlling the call. When empty all participants will be placed on hold. | [optional] 
**AudioUrl** | Pointer to **string** | The URL of a file to be played to the participants when they are put on hold. media_name and audio_url cannot be used together in one request. | [optional] 
**MediaName** | Pointer to **string** | The media_name of a file to be played to the participants when they are put on hold. The media_name must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. The file must either be a WAV or MP3 file. | [optional] 

## Methods

### NewConferenceHoldRequest

`func NewConferenceHoldRequest() *ConferenceHoldRequest`

NewConferenceHoldRequest instantiates a new ConferenceHoldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceHoldRequestWithDefaults

`func NewConferenceHoldRequestWithDefaults() *ConferenceHoldRequest`

NewConferenceHoldRequestWithDefaults instantiates a new ConferenceHoldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlIds

`func (o *ConferenceHoldRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferenceHoldRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferenceHoldRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.

### HasCallControlIds

`func (o *ConferenceHoldRequest) HasCallControlIds() bool`

HasCallControlIds returns a boolean if a field has been set.

### GetAudioUrl

`func (o *ConferenceHoldRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *ConferenceHoldRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *ConferenceHoldRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *ConferenceHoldRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *ConferenceHoldRequest) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *ConferenceHoldRequest) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *ConferenceHoldRequest) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *ConferenceHoldRequest) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


