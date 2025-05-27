# MessagingFeatureSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomesticTwoWay** | **bool** | Send messages to and receive messages from numbers in the same country. | 
**InternationalInbound** | **bool** | Receive messages from numbers in other countries. | 
**InternationalOutbound** | **bool** | Send messages to numbers in other countries. | 

## Methods

### NewMessagingFeatureSet

`func NewMessagingFeatureSet(domesticTwoWay bool, internationalInbound bool, internationalOutbound bool, ) *MessagingFeatureSet`

NewMessagingFeatureSet instantiates a new MessagingFeatureSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingFeatureSetWithDefaults

`func NewMessagingFeatureSetWithDefaults() *MessagingFeatureSet`

NewMessagingFeatureSetWithDefaults instantiates a new MessagingFeatureSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomesticTwoWay

`func (o *MessagingFeatureSet) GetDomesticTwoWay() bool`

GetDomesticTwoWay returns the DomesticTwoWay field if non-nil, zero value otherwise.

### GetDomesticTwoWayOk

`func (o *MessagingFeatureSet) GetDomesticTwoWayOk() (*bool, bool)`

GetDomesticTwoWayOk returns a tuple with the DomesticTwoWay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticTwoWay

`func (o *MessagingFeatureSet) SetDomesticTwoWay(v bool)`

SetDomesticTwoWay sets DomesticTwoWay field to given value.


### GetInternationalInbound

`func (o *MessagingFeatureSet) GetInternationalInbound() bool`

GetInternationalInbound returns the InternationalInbound field if non-nil, zero value otherwise.

### GetInternationalInboundOk

`func (o *MessagingFeatureSet) GetInternationalInboundOk() (*bool, bool)`

GetInternationalInboundOk returns a tuple with the InternationalInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalInbound

`func (o *MessagingFeatureSet) SetInternationalInbound(v bool)`

SetInternationalInbound sets InternationalInbound field to given value.


### GetInternationalOutbound

`func (o *MessagingFeatureSet) GetInternationalOutbound() bool`

GetInternationalOutbound returns the InternationalOutbound field if non-nil, zero value otherwise.

### GetInternationalOutboundOk

`func (o *MessagingFeatureSet) GetInternationalOutboundOk() (*bool, bool)`

GetInternationalOutboundOk returns a tuple with the InternationalOutbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalOutbound

`func (o *MessagingFeatureSet) SetInternationalOutbound(v bool)`

SetInternationalOutbound sets InternationalOutbound field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


