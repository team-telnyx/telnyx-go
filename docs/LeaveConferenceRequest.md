# LeaveConferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | **string** | Unique identifier and token for controlling the call | 
**CommandId** | Pointer to **string** | Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same &#x60;command_id&#x60; as one that has already been executed. | [optional] 
**BeepEnabled** | Pointer to **string** | Whether a beep sound should be played when the participant leaves the conference. Can be used to override the conference-level setting. | [optional] 

## Methods

### NewLeaveConferenceRequest

`func NewLeaveConferenceRequest(callControlId string, ) *LeaveConferenceRequest`

NewLeaveConferenceRequest instantiates a new LeaveConferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaveConferenceRequestWithDefaults

`func NewLeaveConferenceRequestWithDefaults() *LeaveConferenceRequest`

NewLeaveConferenceRequestWithDefaults instantiates a new LeaveConferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *LeaveConferenceRequest) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *LeaveConferenceRequest) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *LeaveConferenceRequest) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetCommandId

`func (o *LeaveConferenceRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *LeaveConferenceRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *LeaveConferenceRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *LeaveConferenceRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetBeepEnabled

`func (o *LeaveConferenceRequest) GetBeepEnabled() string`

GetBeepEnabled returns the BeepEnabled field if non-nil, zero value otherwise.

### GetBeepEnabledOk

`func (o *LeaveConferenceRequest) GetBeepEnabledOk() (*string, bool)`

GetBeepEnabledOk returns a tuple with the BeepEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeepEnabled

`func (o *LeaveConferenceRequest) SetBeepEnabled(v string)`

SetBeepEnabled sets BeepEnabled field to given value.

### HasBeepEnabled

`func (o *LeaveConferenceRequest) HasBeepEnabled() bool`

HasBeepEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


