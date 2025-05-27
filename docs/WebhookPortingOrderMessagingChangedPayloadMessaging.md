# WebhookPortingOrderMessagingChangedPayloadMessaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingCapable** | Pointer to **bool** | Indicates whether the porting order is messaging capable. | [optional] 
**EnableMessaging** | Pointer to **bool** | Indicates whether Telnyx will port messaging capabilities from the losing carrier. If false, any messaging capabilities will stay with their current provider. | [optional] 
**MessagingPortStatus** | Pointer to **string** | Indicates the messaging port status of the porting order. | [optional] 
**MessagingPortCompleted** | Pointer to **bool** | Indicates whether the messaging port is completed. | [optional] 

## Methods

### NewWebhookPortingOrderMessagingChangedPayloadMessaging

`func NewWebhookPortingOrderMessagingChangedPayloadMessaging() *WebhookPortingOrderMessagingChangedPayloadMessaging`

NewWebhookPortingOrderMessagingChangedPayloadMessaging instantiates a new WebhookPortingOrderMessagingChangedPayloadMessaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderMessagingChangedPayloadMessagingWithDefaults

`func NewWebhookPortingOrderMessagingChangedPayloadMessagingWithDefaults() *WebhookPortingOrderMessagingChangedPayloadMessaging`

NewWebhookPortingOrderMessagingChangedPayloadMessagingWithDefaults instantiates a new WebhookPortingOrderMessagingChangedPayloadMessaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingCapable

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetMessagingCapable() bool`

GetMessagingCapable returns the MessagingCapable field if non-nil, zero value otherwise.

### GetMessagingCapableOk

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetMessagingCapableOk() (*bool, bool)`

GetMessagingCapableOk returns a tuple with the MessagingCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingCapable

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) SetMessagingCapable(v bool)`

SetMessagingCapable sets MessagingCapable field to given value.

### HasMessagingCapable

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) HasMessagingCapable() bool`

HasMessagingCapable returns a boolean if a field has been set.

### GetEnableMessaging

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetEnableMessaging() bool`

GetEnableMessaging returns the EnableMessaging field if non-nil, zero value otherwise.

### GetEnableMessagingOk

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetEnableMessagingOk() (*bool, bool)`

GetEnableMessagingOk returns a tuple with the EnableMessaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMessaging

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) SetEnableMessaging(v bool)`

SetEnableMessaging sets EnableMessaging field to given value.

### HasEnableMessaging

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) HasEnableMessaging() bool`

HasEnableMessaging returns a boolean if a field has been set.

### GetMessagingPortStatus

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetMessagingPortStatus() string`

GetMessagingPortStatus returns the MessagingPortStatus field if non-nil, zero value otherwise.

### GetMessagingPortStatusOk

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetMessagingPortStatusOk() (*string, bool)`

GetMessagingPortStatusOk returns a tuple with the MessagingPortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingPortStatus

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) SetMessagingPortStatus(v string)`

SetMessagingPortStatus sets MessagingPortStatus field to given value.

### HasMessagingPortStatus

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) HasMessagingPortStatus() bool`

HasMessagingPortStatus returns a boolean if a field has been set.

### GetMessagingPortCompleted

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetMessagingPortCompleted() bool`

GetMessagingPortCompleted returns the MessagingPortCompleted field if non-nil, zero value otherwise.

### GetMessagingPortCompletedOk

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) GetMessagingPortCompletedOk() (*bool, bool)`

GetMessagingPortCompletedOk returns a tuple with the MessagingPortCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingPortCompleted

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) SetMessagingPortCompleted(v bool)`

SetMessagingPortCompleted sets MessagingPortCompleted field to given value.

### HasMessagingPortCompleted

`func (o *WebhookPortingOrderMessagingChangedPayloadMessaging) HasMessagingPortCompleted() bool`

HasMessagingPortCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


