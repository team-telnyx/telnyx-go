# CallWithRecordingId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call | 
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events | 
**CallControlId** | **string** | Unique identifier and token for controlling the call. | 
**IsAlive** | **bool** | Indicates whether the call is alive or not. For Dial command it will always be &#x60;false&#x60; (dialing is asynchronous). | 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**CallDuration** | Pointer to **int32** | Indicates the duration of the call in seconds | [optional] 
**RecordingId** | Pointer to **string** | The ID of the recording. Only present when the record parameter is set to record-from-answer. | [optional] 

## Methods

### NewCallWithRecordingId

`func NewCallWithRecordingId(recordType string, callSessionId string, callLegId string, callControlId string, isAlive bool, ) *CallWithRecordingId`

NewCallWithRecordingId instantiates a new CallWithRecordingId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithRecordingIdWithDefaults

`func NewCallWithRecordingIdWithDefaults() *CallWithRecordingId`

NewCallWithRecordingIdWithDefaults instantiates a new CallWithRecordingId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallWithRecordingId) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallWithRecordingId) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallWithRecordingId) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetCallSessionId

`func (o *CallWithRecordingId) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallWithRecordingId) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallWithRecordingId) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.


### GetCallLegId

`func (o *CallWithRecordingId) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallWithRecordingId) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallWithRecordingId) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.


### GetCallControlId

`func (o *CallWithRecordingId) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallWithRecordingId) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallWithRecordingId) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetIsAlive

`func (o *CallWithRecordingId) GetIsAlive() bool`

GetIsAlive returns the IsAlive field if non-nil, zero value otherwise.

### GetIsAliveOk

`func (o *CallWithRecordingId) GetIsAliveOk() (*bool, bool)`

GetIsAliveOk returns a tuple with the IsAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlive

`func (o *CallWithRecordingId) SetIsAlive(v bool)`

SetIsAlive sets IsAlive field to given value.


### GetClientState

`func (o *CallWithRecordingId) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallWithRecordingId) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallWithRecordingId) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallWithRecordingId) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCallDuration

`func (o *CallWithRecordingId) GetCallDuration() int32`

GetCallDuration returns the CallDuration field if non-nil, zero value otherwise.

### GetCallDurationOk

`func (o *CallWithRecordingId) GetCallDurationOk() (*int32, bool)`

GetCallDurationOk returns a tuple with the CallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDuration

`func (o *CallWithRecordingId) SetCallDuration(v int32)`

SetCallDuration sets CallDuration field to given value.

### HasCallDuration

`func (o *CallWithRecordingId) HasCallDuration() bool`

HasCallDuration returns a boolean if a field has been set.

### GetRecordingId

`func (o *CallWithRecordingId) GetRecordingId() string`

GetRecordingId returns the RecordingId field if non-nil, zero value otherwise.

### GetRecordingIdOk

`func (o *CallWithRecordingId) GetRecordingIdOk() (*string, bool)`

GetRecordingIdOk returns a tuple with the RecordingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingId

`func (o *CallWithRecordingId) SetRecordingId(v string)`

SetRecordingId sets RecordingId field to given value.

### HasRecordingId

`func (o *CallWithRecordingId) HasRecordingId() bool`

HasRecordingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


