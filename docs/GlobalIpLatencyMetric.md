# GlobalIpLatencyMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The timestamp of the metric. | [optional] 
**GlobalIp** | Pointer to [**GlobalIpAssignmentHealthMetricGlobalIp**](GlobalIpAssignmentHealthMetricGlobalIp.md) |  | [optional] 
**ProberLocation** | Pointer to [**GlobalIpLatencyMetricProberLocation**](GlobalIpLatencyMetricProberLocation.md) |  | [optional] 
**MeanLatency** | Pointer to [**GlobalIpLatencyMetricMeanLatency**](GlobalIpLatencyMetricMeanLatency.md) |  | [optional] 
**PercentileLatency** | Pointer to [**GlobalIpLatencyMetricPercentileLatency**](GlobalIpLatencyMetricPercentileLatency.md) |  | [optional] 

## Methods

### NewGlobalIpLatencyMetric

`func NewGlobalIpLatencyMetric() *GlobalIpLatencyMetric`

NewGlobalIpLatencyMetric instantiates a new GlobalIpLatencyMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpLatencyMetricWithDefaults

`func NewGlobalIpLatencyMetricWithDefaults() *GlobalIpLatencyMetric`

NewGlobalIpLatencyMetricWithDefaults instantiates a new GlobalIpLatencyMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GlobalIpLatencyMetric) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GlobalIpLatencyMetric) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GlobalIpLatencyMetric) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GlobalIpLatencyMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGlobalIp

`func (o *GlobalIpLatencyMetric) GetGlobalIp() GlobalIpAssignmentHealthMetricGlobalIp`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *GlobalIpLatencyMetric) GetGlobalIpOk() (*GlobalIpAssignmentHealthMetricGlobalIp, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *GlobalIpLatencyMetric) SetGlobalIp(v GlobalIpAssignmentHealthMetricGlobalIp)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *GlobalIpLatencyMetric) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetProberLocation

`func (o *GlobalIpLatencyMetric) GetProberLocation() GlobalIpLatencyMetricProberLocation`

GetProberLocation returns the ProberLocation field if non-nil, zero value otherwise.

### GetProberLocationOk

`func (o *GlobalIpLatencyMetric) GetProberLocationOk() (*GlobalIpLatencyMetricProberLocation, bool)`

GetProberLocationOk returns a tuple with the ProberLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProberLocation

`func (o *GlobalIpLatencyMetric) SetProberLocation(v GlobalIpLatencyMetricProberLocation)`

SetProberLocation sets ProberLocation field to given value.

### HasProberLocation

`func (o *GlobalIpLatencyMetric) HasProberLocation() bool`

HasProberLocation returns a boolean if a field has been set.

### GetMeanLatency

`func (o *GlobalIpLatencyMetric) GetMeanLatency() GlobalIpLatencyMetricMeanLatency`

GetMeanLatency returns the MeanLatency field if non-nil, zero value otherwise.

### GetMeanLatencyOk

`func (o *GlobalIpLatencyMetric) GetMeanLatencyOk() (*GlobalIpLatencyMetricMeanLatency, bool)`

GetMeanLatencyOk returns a tuple with the MeanLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeanLatency

`func (o *GlobalIpLatencyMetric) SetMeanLatency(v GlobalIpLatencyMetricMeanLatency)`

SetMeanLatency sets MeanLatency field to given value.

### HasMeanLatency

`func (o *GlobalIpLatencyMetric) HasMeanLatency() bool`

HasMeanLatency returns a boolean if a field has been set.

### GetPercentileLatency

`func (o *GlobalIpLatencyMetric) GetPercentileLatency() GlobalIpLatencyMetricPercentileLatency`

GetPercentileLatency returns the PercentileLatency field if non-nil, zero value otherwise.

### GetPercentileLatencyOk

`func (o *GlobalIpLatencyMetric) GetPercentileLatencyOk() (*GlobalIpLatencyMetricPercentileLatency, bool)`

GetPercentileLatencyOk returns a tuple with the PercentileLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentileLatency

`func (o *GlobalIpLatencyMetric) SetPercentileLatency(v GlobalIpLatencyMetricPercentileLatency)`

SetPercentileLatency sets PercentileLatency field to given value.

### HasPercentileLatency

`func (o *GlobalIpLatencyMetric) HasPercentileLatency() bool`

HasPercentileLatency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


