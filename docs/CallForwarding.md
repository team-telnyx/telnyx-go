# CallForwarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallForwardingEnabled** | Pointer to **bool** | Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints. | [optional] [default to true]
**ForwardsTo** | Pointer to **string** | The phone number to which inbound calls to this number are forwarded. Inbound calls will not be forwarded if this field is left blank. If set, must be a +E.164-formatted phone number. | [optional] 
**ForwardingType** | Pointer to **string** | Call forwarding type. &#39;forwards_to&#39; must be set for this to have an effect. | [optional] 

## Methods

### NewCallForwarding

`func NewCallForwarding() *CallForwarding`

NewCallForwarding instantiates a new CallForwarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallForwardingWithDefaults

`func NewCallForwardingWithDefaults() *CallForwarding`

NewCallForwardingWithDefaults instantiates a new CallForwarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallForwardingEnabled

`func (o *CallForwarding) GetCallForwardingEnabled() bool`

GetCallForwardingEnabled returns the CallForwardingEnabled field if non-nil, zero value otherwise.

### GetCallForwardingEnabledOk

`func (o *CallForwarding) GetCallForwardingEnabledOk() (*bool, bool)`

GetCallForwardingEnabledOk returns a tuple with the CallForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForwardingEnabled

`func (o *CallForwarding) SetCallForwardingEnabled(v bool)`

SetCallForwardingEnabled sets CallForwardingEnabled field to given value.

### HasCallForwardingEnabled

`func (o *CallForwarding) HasCallForwardingEnabled() bool`

HasCallForwardingEnabled returns a boolean if a field has been set.

### GetForwardsTo

`func (o *CallForwarding) GetForwardsTo() string`

GetForwardsTo returns the ForwardsTo field if non-nil, zero value otherwise.

### GetForwardsToOk

`func (o *CallForwarding) GetForwardsToOk() (*string, bool)`

GetForwardsToOk returns a tuple with the ForwardsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardsTo

`func (o *CallForwarding) SetForwardsTo(v string)`

SetForwardsTo sets ForwardsTo field to given value.

### HasForwardsTo

`func (o *CallForwarding) HasForwardsTo() bool`

HasForwardsTo returns a boolean if a field has been set.

### GetForwardingType

`func (o *CallForwarding) GetForwardingType() string`

GetForwardingType returns the ForwardingType field if non-nil, zero value otherwise.

### GetForwardingTypeOk

`func (o *CallForwarding) GetForwardingTypeOk() (*string, bool)`

GetForwardingTypeOk returns a tuple with the ForwardingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingType

`func (o *CallForwarding) SetForwardingType(v string)`

SetForwardingType sets ForwardingType field to given value.

### HasForwardingType

`func (o *CallForwarding) HasForwardingType() bool`

HasForwardingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


