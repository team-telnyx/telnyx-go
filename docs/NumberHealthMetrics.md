# NumberHealthMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageCount** | **int32** | The number of messages analyzed for the health metrics. | 
**InboundOutboundRatio** | **float32** | The ratio of messages received to the number of messages sent. | 
**SuccessRatio** | **float32** | The ratio of messages sucessfully delivered to the number of messages attempted. | 
**SpamRatio** | **float32** | The ratio of messages blocked for spam to the number of messages attempted. | 

## Methods

### NewNumberHealthMetrics

`func NewNumberHealthMetrics(messageCount int32, inboundOutboundRatio float32, successRatio float32, spamRatio float32, ) *NumberHealthMetrics`

NewNumberHealthMetrics instantiates a new NumberHealthMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberHealthMetricsWithDefaults

`func NewNumberHealthMetricsWithDefaults() *NumberHealthMetrics`

NewNumberHealthMetricsWithDefaults instantiates a new NumberHealthMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageCount

`func (o *NumberHealthMetrics) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *NumberHealthMetrics) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *NumberHealthMetrics) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.


### GetInboundOutboundRatio

`func (o *NumberHealthMetrics) GetInboundOutboundRatio() float32`

GetInboundOutboundRatio returns the InboundOutboundRatio field if non-nil, zero value otherwise.

### GetInboundOutboundRatioOk

`func (o *NumberHealthMetrics) GetInboundOutboundRatioOk() (*float32, bool)`

GetInboundOutboundRatioOk returns a tuple with the InboundOutboundRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundOutboundRatio

`func (o *NumberHealthMetrics) SetInboundOutboundRatio(v float32)`

SetInboundOutboundRatio sets InboundOutboundRatio field to given value.


### GetSuccessRatio

`func (o *NumberHealthMetrics) GetSuccessRatio() float32`

GetSuccessRatio returns the SuccessRatio field if non-nil, zero value otherwise.

### GetSuccessRatioOk

`func (o *NumberHealthMetrics) GetSuccessRatioOk() (*float32, bool)`

GetSuccessRatioOk returns a tuple with the SuccessRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRatio

`func (o *NumberHealthMetrics) SetSuccessRatio(v float32)`

SetSuccessRatio sets SuccessRatio field to given value.


### GetSpamRatio

`func (o *NumberHealthMetrics) GetSpamRatio() float32`

GetSpamRatio returns the SpamRatio field if non-nil, zero value otherwise.

### GetSpamRatioOk

`func (o *NumberHealthMetrics) GetSpamRatioOk() (*float32, bool)`

GetSpamRatioOk returns a tuple with the SpamRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamRatio

`func (o *NumberHealthMetrics) SetSpamRatio(v float32)`

SetSpamRatio sets SpamRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


