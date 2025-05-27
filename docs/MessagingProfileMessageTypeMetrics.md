# MessagingProfileMessageTypeMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The metric type. | [optional] [readonly] 
**Sent** | Pointer to **float32** | The number of outbound messages sent. | [optional] [readonly] 
**Delivered** | Pointer to **float32** | The number of outbound messages successfully delivered. | [optional] [readonly] 
**OutboundErrorRatio** | Pointer to **float32** | The ratio of outbound messages sent that resulted in errors. | [optional] [readonly] 
**Received** | Pointer to **float32** | The number of inbound messages received. | [optional] [readonly] 

## Methods

### NewMessagingProfileMessageTypeMetrics

`func NewMessagingProfileMessageTypeMetrics() *MessagingProfileMessageTypeMetrics`

NewMessagingProfileMessageTypeMetrics instantiates a new MessagingProfileMessageTypeMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingProfileMessageTypeMetricsWithDefaults

`func NewMessagingProfileMessageTypeMetricsWithDefaults() *MessagingProfileMessageTypeMetrics`

NewMessagingProfileMessageTypeMetricsWithDefaults instantiates a new MessagingProfileMessageTypeMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *MessagingProfileMessageTypeMetrics) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MessagingProfileMessageTypeMetrics) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MessagingProfileMessageTypeMetrics) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MessagingProfileMessageTypeMetrics) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSent

`func (o *MessagingProfileMessageTypeMetrics) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MessagingProfileMessageTypeMetrics) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MessagingProfileMessageTypeMetrics) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MessagingProfileMessageTypeMetrics) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *MessagingProfileMessageTypeMetrics) GetDelivered() float32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *MessagingProfileMessageTypeMetrics) GetDeliveredOk() (*float32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *MessagingProfileMessageTypeMetrics) SetDelivered(v float32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *MessagingProfileMessageTypeMetrics) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetOutboundErrorRatio

`func (o *MessagingProfileMessageTypeMetrics) GetOutboundErrorRatio() float32`

GetOutboundErrorRatio returns the OutboundErrorRatio field if non-nil, zero value otherwise.

### GetOutboundErrorRatioOk

`func (o *MessagingProfileMessageTypeMetrics) GetOutboundErrorRatioOk() (*float32, bool)`

GetOutboundErrorRatioOk returns a tuple with the OutboundErrorRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundErrorRatio

`func (o *MessagingProfileMessageTypeMetrics) SetOutboundErrorRatio(v float32)`

SetOutboundErrorRatio sets OutboundErrorRatio field to given value.

### HasOutboundErrorRatio

`func (o *MessagingProfileMessageTypeMetrics) HasOutboundErrorRatio() bool`

HasOutboundErrorRatio returns a boolean if a field has been set.

### GetReceived

`func (o *MessagingProfileMessageTypeMetrics) GetReceived() float32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *MessagingProfileMessageTypeMetrics) GetReceivedOk() (*float32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *MessagingProfileMessageTypeMetrics) SetReceived(v float32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *MessagingProfileMessageTypeMetrics) HasReceived() bool`

HasReceived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


