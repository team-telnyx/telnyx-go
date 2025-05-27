# MessagingProfileHighLevelMetricsOutbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **float32** | The number of outbound messages sent. | [optional] [readonly] 
**Delivered** | Pointer to **float32** | The number of outbound messages successfully delivered. | [optional] [readonly] 
**ErrorRatio** | Pointer to **float32** | The ratio of messages sent that resulted in errors. | [optional] [readonly] 

## Methods

### NewMessagingProfileHighLevelMetricsOutbound

`func NewMessagingProfileHighLevelMetricsOutbound() *MessagingProfileHighLevelMetricsOutbound`

NewMessagingProfileHighLevelMetricsOutbound instantiates a new MessagingProfileHighLevelMetricsOutbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingProfileHighLevelMetricsOutboundWithDefaults

`func NewMessagingProfileHighLevelMetricsOutboundWithDefaults() *MessagingProfileHighLevelMetricsOutbound`

NewMessagingProfileHighLevelMetricsOutboundWithDefaults instantiates a new MessagingProfileHighLevelMetricsOutbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *MessagingProfileHighLevelMetricsOutbound) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MessagingProfileHighLevelMetricsOutbound) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MessagingProfileHighLevelMetricsOutbound) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MessagingProfileHighLevelMetricsOutbound) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *MessagingProfileHighLevelMetricsOutbound) GetDelivered() float32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *MessagingProfileHighLevelMetricsOutbound) GetDeliveredOk() (*float32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *MessagingProfileHighLevelMetricsOutbound) SetDelivered(v float32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *MessagingProfileHighLevelMetricsOutbound) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetErrorRatio

`func (o *MessagingProfileHighLevelMetricsOutbound) GetErrorRatio() float32`

GetErrorRatio returns the ErrorRatio field if non-nil, zero value otherwise.

### GetErrorRatioOk

`func (o *MessagingProfileHighLevelMetricsOutbound) GetErrorRatioOk() (*float32, bool)`

GetErrorRatioOk returns a tuple with the ErrorRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRatio

`func (o *MessagingProfileHighLevelMetricsOutbound) SetErrorRatio(v float32)`

SetErrorRatio sets ErrorRatio field to given value.

### HasErrorRatio

`func (o *MessagingProfileHighLevelMetricsOutbound) HasErrorRatio() bool`

HasErrorRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


