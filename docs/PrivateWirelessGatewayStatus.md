# PrivateWirelessGatewayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The current status or failure details of the Private Wireless Gateway. &lt;ul&gt;  &lt;li&gt;&lt;code&gt;provisioning&lt;/code&gt; - the Private Wireless Gateway is being provisioned.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;provisioned&lt;/code&gt; - the Private Wireless Gateway was provisioned and able to receive connections.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;failed&lt;/code&gt; - the provisioning had failed for a reason and it requires an intervention.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;decommissioning&lt;/code&gt; - the Private Wireless Gateway is being removed from the network.&lt;/li&gt;  &lt;/ul&gt;  Transitioning between the provisioning and provisioned states may take some time. | [optional] [readonly] [default to "provisioning"]
**ErrorDescription** | Pointer to **string** | This attribute provides a human-readable explanation of why a failure happened. | [optional] [readonly] 
**ErrorCode** | Pointer to **string** | This attribute is an [error code](https://developers.telnyx.com/api/errors) related to the failure reason. | [optional] [readonly] 

## Methods

### NewPrivateWirelessGatewayStatus

`func NewPrivateWirelessGatewayStatus() *PrivateWirelessGatewayStatus`

NewPrivateWirelessGatewayStatus instantiates a new PrivateWirelessGatewayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateWirelessGatewayStatusWithDefaults

`func NewPrivateWirelessGatewayStatusWithDefaults() *PrivateWirelessGatewayStatus`

NewPrivateWirelessGatewayStatusWithDefaults instantiates a new PrivateWirelessGatewayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PrivateWirelessGatewayStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PrivateWirelessGatewayStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PrivateWirelessGatewayStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PrivateWirelessGatewayStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetErrorDescription

`func (o *PrivateWirelessGatewayStatus) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *PrivateWirelessGatewayStatus) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *PrivateWirelessGatewayStatus) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *PrivateWirelessGatewayStatus) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorCode

`func (o *PrivateWirelessGatewayStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PrivateWirelessGatewayStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PrivateWirelessGatewayStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PrivateWirelessGatewayStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


