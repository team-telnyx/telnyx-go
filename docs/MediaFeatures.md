# MediaFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RtpAutoAdjustEnabled** | Pointer to **bool** | When RTP Auto-Adjust is enabled, the destination RTP address port will be automatically changed to match the source of the incoming RTP packets. | [optional] [default to true]
**AcceptAnyRtpPacketsEnabled** | Pointer to **bool** | When enabled, Telnyx will accept RTP packets from any customer-side IP address and port, not just those to which Telnyx is sending RTP. | [optional] [default to false]
**T38FaxGatewayEnabled** | Pointer to **bool** | Controls whether Telnyx will accept a T.38 re-INVITE for this phone number. Note that Telnyx will not send a T.38 re-INVITE; this option only controls whether one will be accepted. | [optional] [default to false]

## Methods

### NewMediaFeatures

`func NewMediaFeatures() *MediaFeatures`

NewMediaFeatures instantiates a new MediaFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaFeaturesWithDefaults

`func NewMediaFeaturesWithDefaults() *MediaFeatures`

NewMediaFeaturesWithDefaults instantiates a new MediaFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRtpAutoAdjustEnabled

`func (o *MediaFeatures) GetRtpAutoAdjustEnabled() bool`

GetRtpAutoAdjustEnabled returns the RtpAutoAdjustEnabled field if non-nil, zero value otherwise.

### GetRtpAutoAdjustEnabledOk

`func (o *MediaFeatures) GetRtpAutoAdjustEnabledOk() (*bool, bool)`

GetRtpAutoAdjustEnabledOk returns a tuple with the RtpAutoAdjustEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtpAutoAdjustEnabled

`func (o *MediaFeatures) SetRtpAutoAdjustEnabled(v bool)`

SetRtpAutoAdjustEnabled sets RtpAutoAdjustEnabled field to given value.

### HasRtpAutoAdjustEnabled

`func (o *MediaFeatures) HasRtpAutoAdjustEnabled() bool`

HasRtpAutoAdjustEnabled returns a boolean if a field has been set.

### GetAcceptAnyRtpPacketsEnabled

`func (o *MediaFeatures) GetAcceptAnyRtpPacketsEnabled() bool`

GetAcceptAnyRtpPacketsEnabled returns the AcceptAnyRtpPacketsEnabled field if non-nil, zero value otherwise.

### GetAcceptAnyRtpPacketsEnabledOk

`func (o *MediaFeatures) GetAcceptAnyRtpPacketsEnabledOk() (*bool, bool)`

GetAcceptAnyRtpPacketsEnabledOk returns a tuple with the AcceptAnyRtpPacketsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyRtpPacketsEnabled

`func (o *MediaFeatures) SetAcceptAnyRtpPacketsEnabled(v bool)`

SetAcceptAnyRtpPacketsEnabled sets AcceptAnyRtpPacketsEnabled field to given value.

### HasAcceptAnyRtpPacketsEnabled

`func (o *MediaFeatures) HasAcceptAnyRtpPacketsEnabled() bool`

HasAcceptAnyRtpPacketsEnabled returns a boolean if a field has been set.

### GetT38FaxGatewayEnabled

`func (o *MediaFeatures) GetT38FaxGatewayEnabled() bool`

GetT38FaxGatewayEnabled returns the T38FaxGatewayEnabled field if non-nil, zero value otherwise.

### GetT38FaxGatewayEnabledOk

`func (o *MediaFeatures) GetT38FaxGatewayEnabledOk() (*bool, bool)`

GetT38FaxGatewayEnabledOk returns a tuple with the T38FaxGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38FaxGatewayEnabled

`func (o *MediaFeatures) SetT38FaxGatewayEnabled(v bool)`

SetT38FaxGatewayEnabled sets T38FaxGatewayEnabled field to given value.

### HasT38FaxGatewayEnabled

`func (o *MediaFeatures) HasT38FaxGatewayEnabled() bool`

HasT38FaxGatewayEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


