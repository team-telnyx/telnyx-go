# GlobalIpUsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The timestamp of the metric. | [optional] 
**GlobalIp** | Pointer to [**GlobalIpUsageMetricGlobalIp**](GlobalIpUsageMetricGlobalIp.md) |  | [optional] 
**Transmitted** | Pointer to [**GlobalIpAssignmentUsageMetricTransmitted**](GlobalIpAssignmentUsageMetricTransmitted.md) |  | [optional] 
**Received** | Pointer to [**GlobalIpAssignmentUsageMetricReceived**](GlobalIpAssignmentUsageMetricReceived.md) |  | [optional] 

## Methods

### NewGlobalIpUsageMetric

`func NewGlobalIpUsageMetric() *GlobalIpUsageMetric`

NewGlobalIpUsageMetric instantiates a new GlobalIpUsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpUsageMetricWithDefaults

`func NewGlobalIpUsageMetricWithDefaults() *GlobalIpUsageMetric`

NewGlobalIpUsageMetricWithDefaults instantiates a new GlobalIpUsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GlobalIpUsageMetric) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GlobalIpUsageMetric) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GlobalIpUsageMetric) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GlobalIpUsageMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGlobalIp

`func (o *GlobalIpUsageMetric) GetGlobalIp() GlobalIpUsageMetricGlobalIp`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *GlobalIpUsageMetric) GetGlobalIpOk() (*GlobalIpUsageMetricGlobalIp, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *GlobalIpUsageMetric) SetGlobalIp(v GlobalIpUsageMetricGlobalIp)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *GlobalIpUsageMetric) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetTransmitted

`func (o *GlobalIpUsageMetric) GetTransmitted() GlobalIpAssignmentUsageMetricTransmitted`

GetTransmitted returns the Transmitted field if non-nil, zero value otherwise.

### GetTransmittedOk

`func (o *GlobalIpUsageMetric) GetTransmittedOk() (*GlobalIpAssignmentUsageMetricTransmitted, bool)`

GetTransmittedOk returns a tuple with the Transmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitted

`func (o *GlobalIpUsageMetric) SetTransmitted(v GlobalIpAssignmentUsageMetricTransmitted)`

SetTransmitted sets Transmitted field to given value.

### HasTransmitted

`func (o *GlobalIpUsageMetric) HasTransmitted() bool`

HasTransmitted returns a boolean if a field has been set.

### GetReceived

`func (o *GlobalIpUsageMetric) GetReceived() GlobalIpAssignmentUsageMetricReceived`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GlobalIpUsageMetric) GetReceivedOk() (*GlobalIpAssignmentUsageMetricReceived, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GlobalIpUsageMetric) SetReceived(v GlobalIpAssignmentUsageMetricReceived)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *GlobalIpUsageMetric) HasReceived() bool`

HasReceived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


