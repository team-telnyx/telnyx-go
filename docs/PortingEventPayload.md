# PortingEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the porting order that was moved. | [optional] 
**CustomerReference** | Pointer to **string** | Identifies the customer reference associated with the porting order. | [optional] 
**DeletedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the porting order was deleted. | [optional] 
**SupportKey** | Pointer to **string** | Identifies the support key associated with the porting order. | [optional] 
**Messaging** | Pointer to [**WebhookPortingOrderMessagingChangedPayloadMessaging**](WebhookPortingOrderMessagingChangedPayloadMessaging.md) |  | [optional] 
**Status** | Pointer to [**PortingOrderStatus**](PortingOrderStatus.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the porting order was moved. | [optional] 
**WebhookUrl** | Pointer to **string** | The URL to send the webhook to. | [optional] 
**PortingOrderId** | Pointer to **string** | Identifies the porting order that the comment was added to. | [optional] 
**Comment** | Pointer to [**WebhookPortingOrderNewCommentPayloadComment**](WebhookPortingOrderNewCommentPayloadComment.md) |  | [optional] 
**From** | Pointer to [**WebhookPortingOrderSplitPayloadFrom**](WebhookPortingOrderSplitPayloadFrom.md) |  | [optional] 
**To** | Pointer to [**WebhookPortingOrderSplitPayloadTo**](WebhookPortingOrderSplitPayloadTo.md) |  | [optional] 
**PortingPhoneNumbers** | Pointer to [**[]WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner**](WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner.md) | The list of porting phone numbers that were moved to the new porting order. | [optional] 

## Methods

### NewPortingEventPayload

`func NewPortingEventPayload() *PortingEventPayload`

NewPortingEventPayload instantiates a new PortingEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingEventPayloadWithDefaults

`func NewPortingEventPayloadWithDefaults() *PortingEventPayload`

NewPortingEventPayloadWithDefaults instantiates a new PortingEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingEventPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingEventPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingEventPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingEventPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *PortingEventPayload) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *PortingEventPayload) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *PortingEventPayload) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *PortingEventPayload) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetDeletedAt

`func (o *PortingEventPayload) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PortingEventPayload) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PortingEventPayload) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PortingEventPayload) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetSupportKey

`func (o *PortingEventPayload) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *PortingEventPayload) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *PortingEventPayload) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *PortingEventPayload) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetMessaging

`func (o *PortingEventPayload) GetMessaging() WebhookPortingOrderMessagingChangedPayloadMessaging`

GetMessaging returns the Messaging field if non-nil, zero value otherwise.

### GetMessagingOk

`func (o *PortingEventPayload) GetMessagingOk() (*WebhookPortingOrderMessagingChangedPayloadMessaging, bool)`

GetMessagingOk returns a tuple with the Messaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessaging

`func (o *PortingEventPayload) SetMessaging(v WebhookPortingOrderMessagingChangedPayloadMessaging)`

SetMessaging sets Messaging field to given value.

### HasMessaging

`func (o *PortingEventPayload) HasMessaging() bool`

HasMessaging returns a boolean if a field has been set.

### GetStatus

`func (o *PortingEventPayload) GetStatus() PortingOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortingEventPayload) GetStatusOk() (*PortingOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortingEventPayload) SetStatus(v PortingOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortingEventPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingEventPayload) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingEventPayload) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingEventPayload) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingEventPayload) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PortingEventPayload) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PortingEventPayload) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PortingEventPayload) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PortingEventPayload) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetPortingOrderId

`func (o *PortingEventPayload) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *PortingEventPayload) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *PortingEventPayload) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *PortingEventPayload) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetComment

`func (o *PortingEventPayload) GetComment() WebhookPortingOrderNewCommentPayloadComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PortingEventPayload) GetCommentOk() (*WebhookPortingOrderNewCommentPayloadComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PortingEventPayload) SetComment(v WebhookPortingOrderNewCommentPayloadComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PortingEventPayload) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetFrom

`func (o *PortingEventPayload) GetFrom() WebhookPortingOrderSplitPayloadFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PortingEventPayload) GetFromOk() (*WebhookPortingOrderSplitPayloadFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PortingEventPayload) SetFrom(v WebhookPortingOrderSplitPayloadFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PortingEventPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PortingEventPayload) GetTo() WebhookPortingOrderSplitPayloadTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PortingEventPayload) GetToOk() (*WebhookPortingOrderSplitPayloadTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PortingEventPayload) SetTo(v WebhookPortingOrderSplitPayloadTo)`

SetTo sets To field to given value.

### HasTo

`func (o *PortingEventPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetPortingPhoneNumbers

`func (o *PortingEventPayload) GetPortingPhoneNumbers() []WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner`

GetPortingPhoneNumbers returns the PortingPhoneNumbers field if non-nil, zero value otherwise.

### GetPortingPhoneNumbersOk

`func (o *PortingEventPayload) GetPortingPhoneNumbersOk() (*[]WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner, bool)`

GetPortingPhoneNumbersOk returns a tuple with the PortingPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingPhoneNumbers

`func (o *PortingEventPayload) SetPortingPhoneNumbers(v []WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner)`

SetPortingPhoneNumbers sets PortingPhoneNumbers field to given value.

### HasPortingPhoneNumbers

`func (o *PortingEventPayload) HasPortingPhoneNumbers() bool`

HasPortingPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


