# GlobalIpAssignmentUsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The timestamp of the metric. | [optional] 
**GlobalIp** | Pointer to [**GlobalIpAssignmentUsageMetricGlobalIp**](GlobalIpAssignmentUsageMetricGlobalIp.md) |  | [optional] 
**GlobalIpAssignment** | Pointer to [**GlobalIpAssignmentUsageMetricGlobalIpAssignment**](GlobalIpAssignmentUsageMetricGlobalIpAssignment.md) |  | [optional] 
**Transmitted** | Pointer to [**GlobalIpAssignmentUsageMetricTransmitted**](GlobalIpAssignmentUsageMetricTransmitted.md) |  | [optional] 
**Received** | Pointer to [**GlobalIpAssignmentUsageMetricReceived**](GlobalIpAssignmentUsageMetricReceived.md) |  | [optional] 

## Methods

### NewGlobalIpAssignmentUsageMetric

`func NewGlobalIpAssignmentUsageMetric() *GlobalIpAssignmentUsageMetric`

NewGlobalIpAssignmentUsageMetric instantiates a new GlobalIpAssignmentUsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpAssignmentUsageMetricWithDefaults

`func NewGlobalIpAssignmentUsageMetricWithDefaults() *GlobalIpAssignmentUsageMetric`

NewGlobalIpAssignmentUsageMetricWithDefaults instantiates a new GlobalIpAssignmentUsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GlobalIpAssignmentUsageMetric) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GlobalIpAssignmentUsageMetric) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GlobalIpAssignmentUsageMetric) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GlobalIpAssignmentUsageMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGlobalIp

`func (o *GlobalIpAssignmentUsageMetric) GetGlobalIp() GlobalIpAssignmentUsageMetricGlobalIp`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *GlobalIpAssignmentUsageMetric) GetGlobalIpOk() (*GlobalIpAssignmentUsageMetricGlobalIp, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *GlobalIpAssignmentUsageMetric) SetGlobalIp(v GlobalIpAssignmentUsageMetricGlobalIp)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *GlobalIpAssignmentUsageMetric) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetGlobalIpAssignment

`func (o *GlobalIpAssignmentUsageMetric) GetGlobalIpAssignment() GlobalIpAssignmentUsageMetricGlobalIpAssignment`

GetGlobalIpAssignment returns the GlobalIpAssignment field if non-nil, zero value otherwise.

### GetGlobalIpAssignmentOk

`func (o *GlobalIpAssignmentUsageMetric) GetGlobalIpAssignmentOk() (*GlobalIpAssignmentUsageMetricGlobalIpAssignment, bool)`

GetGlobalIpAssignmentOk returns a tuple with the GlobalIpAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIpAssignment

`func (o *GlobalIpAssignmentUsageMetric) SetGlobalIpAssignment(v GlobalIpAssignmentUsageMetricGlobalIpAssignment)`

SetGlobalIpAssignment sets GlobalIpAssignment field to given value.

### HasGlobalIpAssignment

`func (o *GlobalIpAssignmentUsageMetric) HasGlobalIpAssignment() bool`

HasGlobalIpAssignment returns a boolean if a field has been set.

### GetTransmitted

`func (o *GlobalIpAssignmentUsageMetric) GetTransmitted() GlobalIpAssignmentUsageMetricTransmitted`

GetTransmitted returns the Transmitted field if non-nil, zero value otherwise.

### GetTransmittedOk

`func (o *GlobalIpAssignmentUsageMetric) GetTransmittedOk() (*GlobalIpAssignmentUsageMetricTransmitted, bool)`

GetTransmittedOk returns a tuple with the Transmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitted

`func (o *GlobalIpAssignmentUsageMetric) SetTransmitted(v GlobalIpAssignmentUsageMetricTransmitted)`

SetTransmitted sets Transmitted field to given value.

### HasTransmitted

`func (o *GlobalIpAssignmentUsageMetric) HasTransmitted() bool`

HasTransmitted returns a boolean if a field has been set.

### GetReceived

`func (o *GlobalIpAssignmentUsageMetric) GetReceived() GlobalIpAssignmentUsageMetricReceived`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GlobalIpAssignmentUsageMetric) GetReceivedOk() (*GlobalIpAssignmentUsageMetricReceived, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GlobalIpAssignmentUsageMetric) SetReceived(v GlobalIpAssignmentUsageMetricReceived)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *GlobalIpAssignmentUsageMetric) HasReceived() bool`

HasReceived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


