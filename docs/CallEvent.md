# CallEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**CallLegId** | **string** | Uniquely identifies an individual call leg. | 
**CallSessionId** | **string** | Uniquely identifies the call control session. A session may include multiple call leg events. | 
**EventTimestamp** | **string** | Event timestamp | 
**Name** | **string** | Event name | 
**Type** | **string** | Event type | 
**Metadata** | **map[string]interface{}** | Event metadata, which includes raw event, and extra information based on event type | 

## Methods

### NewCallEvent

`func NewCallEvent(recordType string, callLegId string, callSessionId string, eventTimestamp string, name string, type_ string, metadata map[string]interface{}, ) *CallEvent`

NewCallEvent instantiates a new CallEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallEventWithDefaults

`func NewCallEventWithDefaults() *CallEvent`

NewCallEventWithDefaults instantiates a new CallEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *CallEvent) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CallEvent) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CallEvent) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetCallLegId

`func (o *CallEvent) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallEvent) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallEvent) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.


### GetCallSessionId

`func (o *CallEvent) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallEvent) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallEvent) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.


### GetEventTimestamp

`func (o *CallEvent) GetEventTimestamp() string`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CallEvent) GetEventTimestampOk() (*string, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CallEvent) SetEventTimestamp(v string)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetName

`func (o *CallEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallEvent) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CallEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CallEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CallEvent) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *CallEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CallEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CallEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


