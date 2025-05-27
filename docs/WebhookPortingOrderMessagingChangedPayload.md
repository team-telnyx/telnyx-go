# WebhookPortingOrderMessagingChangedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the porting order that was moved. | [optional] 
**CustomerReference** | Pointer to **string** | Identifies the customer reference associated with the porting order. | [optional] 
**SupportKey** | Pointer to **string** | Identifies the support key associated with the porting order. | [optional] 
**Messaging** | Pointer to [**WebhookPortingOrderMessagingChangedPayloadMessaging**](WebhookPortingOrderMessagingChangedPayloadMessaging.md) |  | [optional] 

## Methods

### NewWebhookPortingOrderMessagingChangedPayload

`func NewWebhookPortingOrderMessagingChangedPayload() *WebhookPortingOrderMessagingChangedPayload`

NewWebhookPortingOrderMessagingChangedPayload instantiates a new WebhookPortingOrderMessagingChangedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderMessagingChangedPayloadWithDefaults

`func NewWebhookPortingOrderMessagingChangedPayloadWithDefaults() *WebhookPortingOrderMessagingChangedPayload`

NewWebhookPortingOrderMessagingChangedPayloadWithDefaults instantiates a new WebhookPortingOrderMessagingChangedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortingOrderMessagingChangedPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortingOrderMessagingChangedPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortingOrderMessagingChangedPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortingOrderMessagingChangedPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *WebhookPortingOrderMessagingChangedPayload) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *WebhookPortingOrderMessagingChangedPayload) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *WebhookPortingOrderMessagingChangedPayload) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *WebhookPortingOrderMessagingChangedPayload) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetSupportKey

`func (o *WebhookPortingOrderMessagingChangedPayload) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *WebhookPortingOrderMessagingChangedPayload) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *WebhookPortingOrderMessagingChangedPayload) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *WebhookPortingOrderMessagingChangedPayload) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetMessaging

`func (o *WebhookPortingOrderMessagingChangedPayload) GetMessaging() WebhookPortingOrderMessagingChangedPayloadMessaging`

GetMessaging returns the Messaging field if non-nil, zero value otherwise.

### GetMessagingOk

`func (o *WebhookPortingOrderMessagingChangedPayload) GetMessagingOk() (*WebhookPortingOrderMessagingChangedPayloadMessaging, bool)`

GetMessagingOk returns a tuple with the Messaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessaging

`func (o *WebhookPortingOrderMessagingChangedPayload) SetMessaging(v WebhookPortingOrderMessagingChangedPayloadMessaging)`

SetMessaging sets Messaging field to given value.

### HasMessaging

`func (o *WebhookPortingOrderMessagingChangedPayload) HasMessaging() bool`

HasMessaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


