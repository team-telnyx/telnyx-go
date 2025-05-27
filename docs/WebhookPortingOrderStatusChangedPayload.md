# WebhookPortingOrderStatusChangedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the porting order that was moved. | [optional] 
**CustomerReference** | Pointer to **string** | Identifies the customer reference associated with the porting order. | [optional] 
**Status** | Pointer to [**PortingOrderStatus**](PortingOrderStatus.md) |  | [optional] 
**SupportKey** | Pointer to **string** | Identifies the support key associated with the porting order. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the porting order was moved. | [optional] 
**WebhookUrl** | Pointer to **string** | The URL to send the webhook to. | [optional] 

## Methods

### NewWebhookPortingOrderStatusChangedPayload

`func NewWebhookPortingOrderStatusChangedPayload() *WebhookPortingOrderStatusChangedPayload`

NewWebhookPortingOrderStatusChangedPayload instantiates a new WebhookPortingOrderStatusChangedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderStatusChangedPayloadWithDefaults

`func NewWebhookPortingOrderStatusChangedPayloadWithDefaults() *WebhookPortingOrderStatusChangedPayload`

NewWebhookPortingOrderStatusChangedPayloadWithDefaults instantiates a new WebhookPortingOrderStatusChangedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortingOrderStatusChangedPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortingOrderStatusChangedPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortingOrderStatusChangedPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortingOrderStatusChangedPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *WebhookPortingOrderStatusChangedPayload) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *WebhookPortingOrderStatusChangedPayload) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *WebhookPortingOrderStatusChangedPayload) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *WebhookPortingOrderStatusChangedPayload) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookPortingOrderStatusChangedPayload) GetStatus() PortingOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPortingOrderStatusChangedPayload) GetStatusOk() (*PortingOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPortingOrderStatusChangedPayload) SetStatus(v PortingOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookPortingOrderStatusChangedPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportKey

`func (o *WebhookPortingOrderStatusChangedPayload) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *WebhookPortingOrderStatusChangedPayload) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *WebhookPortingOrderStatusChangedPayload) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *WebhookPortingOrderStatusChangedPayload) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookPortingOrderStatusChangedPayload) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookPortingOrderStatusChangedPayload) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookPortingOrderStatusChangedPayload) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookPortingOrderStatusChangedPayload) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *WebhookPortingOrderStatusChangedPayload) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *WebhookPortingOrderStatusChangedPayload) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *WebhookPortingOrderStatusChangedPayload) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *WebhookPortingOrderStatusChangedPayload) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


