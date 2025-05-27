# ConferenceUnmuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | Pointer to **[]string** | List of unique identifiers and tokens for controlling the call. Enter each call control ID to be unmuted. When empty all participants will be unmuted. | [optional] 

## Methods

### NewConferenceUnmuteRequest

`func NewConferenceUnmuteRequest() *ConferenceUnmuteRequest`

NewConferenceUnmuteRequest instantiates a new ConferenceUnmuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceUnmuteRequestWithDefaults

`func NewConferenceUnmuteRequestWithDefaults() *ConferenceUnmuteRequest`

NewConferenceUnmuteRequestWithDefaults instantiates a new ConferenceUnmuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlIds

`func (o *ConferenceUnmuteRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferenceUnmuteRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferenceUnmuteRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.

### HasCallControlIds

`func (o *ConferenceUnmuteRequest) HasCallControlIds() bool`

HasCallControlIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


