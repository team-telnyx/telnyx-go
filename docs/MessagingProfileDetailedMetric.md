# MessagingProfileDetailedMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** | The timestamp of the aggregated data. | [optional] [readonly] 
**Metrics** | Pointer to [**[]MessagingProfileMessageTypeMetrics**](MessagingProfileMessageTypeMetrics.md) |  | [optional] 

## Methods

### NewMessagingProfileDetailedMetric

`func NewMessagingProfileDetailedMetric() *MessagingProfileDetailedMetric`

NewMessagingProfileDetailedMetric instantiates a new MessagingProfileDetailedMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingProfileDetailedMetricWithDefaults

`func NewMessagingProfileDetailedMetricWithDefaults() *MessagingProfileDetailedMetric`

NewMessagingProfileDetailedMetricWithDefaults instantiates a new MessagingProfileDetailedMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MessagingProfileDetailedMetric) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessagingProfileDetailedMetric) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessagingProfileDetailedMetric) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessagingProfileDetailedMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMetrics

`func (o *MessagingProfileDetailedMetric) GetMetrics() []MessagingProfileMessageTypeMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MessagingProfileDetailedMetric) GetMetricsOk() (*[]MessagingProfileMessageTypeMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MessagingProfileDetailedMetric) SetMetrics(v []MessagingProfileMessageTypeMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MessagingProfileDetailedMetric) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


