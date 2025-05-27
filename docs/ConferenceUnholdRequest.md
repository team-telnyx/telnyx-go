# ConferenceUnholdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlIds** | **[]string** | List of unique identifiers and tokens for controlling the call. Enter each call control ID to be unheld. | 

## Methods

### NewConferenceUnholdRequest

`func NewConferenceUnholdRequest(callControlIds []string, ) *ConferenceUnholdRequest`

NewConferenceUnholdRequest instantiates a new ConferenceUnholdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceUnholdRequestWithDefaults

`func NewConferenceUnholdRequestWithDefaults() *ConferenceUnholdRequest`

NewConferenceUnholdRequestWithDefaults instantiates a new ConferenceUnholdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlIds

`func (o *ConferenceUnholdRequest) GetCallControlIds() []string`

GetCallControlIds returns the CallControlIds field if non-nil, zero value otherwise.

### GetCallControlIdsOk

`func (o *ConferenceUnholdRequest) GetCallControlIdsOk() (*[]string, bool)`

GetCallControlIdsOk returns a tuple with the CallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlIds

`func (o *ConferenceUnholdRequest) SetCallControlIds(v []string)`

SetCallControlIds sets CallControlIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


