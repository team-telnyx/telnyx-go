# ConferenceStopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | Pointer to **[]string** | List of call control ids identifying participants the audio file should stop be played to. If not given, the audio will be stoped to the entire conference. | [optional] 

## Methods

### NewConferenceStopRequest

`func NewConferenceStopRequest() *ConferenceStopRequest`

NewConferenceStopRequest instantiates a new ConferenceStopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceStopRequestWithDefaults

`func NewConferenceStopRequestWithDefaults() *ConferenceStopRequest`

NewConferenceStopRequestWithDefaults instantiates a new ConferenceStopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlIds

`func (o *ConferenceStopRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferenceStopRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferenceStopRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.

### HasCallControlIds

`func (o *ConferenceStopRequest) HasCallControlIds() bool`

HasCallControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


