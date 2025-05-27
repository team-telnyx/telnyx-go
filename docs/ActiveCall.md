# ActiveCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call | 
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events | 
**CallControlId** | **string** | Unique identifier and token for controlling the call. | 
**ClientState** | **string** | State received from a command. | 
**CallDuration** | **int32** | Indicates the duration of the call in seconds | 

## Methods

### NewActiveCall

`func NewActiveCall(recordType string, callSessionId string, callLegId string, callControlId string, clientState string, callDuration int32, ) *ActiveCall`

NewActiveCall instantiates a new ActiveCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveCallWithDefaults

`func NewActiveCallWithDefaults() *ActiveCall`

NewActiveCallWithDefaults instantiates a new ActiveCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ActiveCall) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ActiveCall) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ActiveCall) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetCallSessionId

`func (o *ActiveCall) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ActiveCall) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ActiveCall) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.


### GetCallLegId

`func (o *ActiveCall) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *ActiveCall) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *ActiveCall) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.


### GetCallControlId

`func (o *ActiveCall) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *ActiveCall) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *ActiveCall) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetClientState

`func (o *ActiveCall) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ActiveCall) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ActiveCall) SetClientState(v string)`

SetClientState sets ClientState field to given value.


### GetCallDuration

`func (o *ActiveCall) GetCallDuration() int32`

GetCallDuration returns the CallDuration field if non-nil, zero value otherwise.

### GetCallDurationOk

`func (o *ActiveCall) GetCallDurationOk() (*int32, bool)`

GetCallDurationOk returns a tuple with the CallDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDuration

`func (o *ActiveCall) SetCallDuration(v int32)`

SetCallDuration sets CallDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


