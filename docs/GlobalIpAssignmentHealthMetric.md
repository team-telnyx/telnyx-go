# GlobalIpAssignmentHealthMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The timestamp of the metric. | [optional] 
**GlobalIp** | Pointer to [**GlobalIpAssignmentHealthMetricGlobalIp**](GlobalIpAssignmentHealthMetricGlobalIp.md) |  | [optional] 
**GlobalIpAssignment** | Pointer to [**GlobalIpAssignmentHealthMetricGlobalIpAssignment**](GlobalIpAssignmentHealthMetricGlobalIpAssignment.md) |  | [optional] 
**Health** | Pointer to [**GlobalIpAssignmentHealthMetricHealth**](GlobalIpAssignmentHealthMetricHealth.md) |  | [optional] 

## Methods

### NewGlobalIpAssignmentHealthMetric

`func NewGlobalIpAssignmentHealthMetric() *GlobalIpAssignmentHealthMetric`

NewGlobalIpAssignmentHealthMetric instantiates a new GlobalIpAssignmentHealthMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpAssignmentHealthMetricWithDefaults

`func NewGlobalIpAssignmentHealthMetricWithDefaults() *GlobalIpAssignmentHealthMetric`

NewGlobalIpAssignmentHealthMetricWithDefaults instantiates a new GlobalIpAssignmentHealthMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *GlobalIpAssignmentHealthMetric) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GlobalIpAssignmentHealthMetric) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GlobalIpAssignmentHealthMetric) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GlobalIpAssignmentHealthMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetGlobalIp

`func (o *GlobalIpAssignmentHealthMetric) GetGlobalIp() GlobalIpAssignmentHealthMetricGlobalIp`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *GlobalIpAssignmentHealthMetric) GetGlobalIpOk() (*GlobalIpAssignmentHealthMetricGlobalIp, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *GlobalIpAssignmentHealthMetric) SetGlobalIp(v GlobalIpAssignmentHealthMetricGlobalIp)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *GlobalIpAssignmentHealthMetric) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetGlobalIpAssignment

`func (o *GlobalIpAssignmentHealthMetric) GetGlobalIpAssignment() GlobalIpAssignmentHealthMetricGlobalIpAssignment`

GetGlobalIpAssignment returns the GlobalIpAssignment field if non-nil, zero value otherwise.

### GetGlobalIpAssignmentOk

`func (o *GlobalIpAssignmentHealthMetric) GetGlobalIpAssignmentOk() (*GlobalIpAssignmentHealthMetricGlobalIpAssignment, bool)`

GetGlobalIpAssignmentOk returns a tuple with the GlobalIpAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIpAssignment

`func (o *GlobalIpAssignmentHealthMetric) SetGlobalIpAssignment(v GlobalIpAssignmentHealthMetricGlobalIpAssignment)`

SetGlobalIpAssignment sets GlobalIpAssignment field to given value.

### HasGlobalIpAssignment

`func (o *GlobalIpAssignmentHealthMetric) HasGlobalIpAssignment() bool`

HasGlobalIpAssignment returns a boolean if a field has been set.

### GetHealth

`func (o *GlobalIpAssignmentHealthMetric) GetHealth() GlobalIpAssignmentHealthMetricHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *GlobalIpAssignmentHealthMetric) GetHealthOk() (*GlobalIpAssignmentHealthMetricHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *GlobalIpAssignmentHealthMetric) SetHealth(v GlobalIpAssignmentHealthMetricHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *GlobalIpAssignmentHealthMetric) HasHealth() bool`

HasHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


