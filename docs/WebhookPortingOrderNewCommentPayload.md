# WebhookPortingOrderNewCommentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortingOrderId** | Pointer to **string** | Identifies the porting order that the comment was added to. | [optional] 
**SupportKey** | Pointer to **string** | Identifies the support key associated with the porting order. | [optional] 
**Comment** | Pointer to [**WebhookPortingOrderNewCommentPayloadComment**](WebhookPortingOrderNewCommentPayloadComment.md) |  | [optional] 

## Methods

### NewWebhookPortingOrderNewCommentPayload

`func NewWebhookPortingOrderNewCommentPayload() *WebhookPortingOrderNewCommentPayload`

NewWebhookPortingOrderNewCommentPayload instantiates a new WebhookPortingOrderNewCommentPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderNewCommentPayloadWithDefaults

`func NewWebhookPortingOrderNewCommentPayloadWithDefaults() *WebhookPortingOrderNewCommentPayload`

NewWebhookPortingOrderNewCommentPayloadWithDefaults instantiates a new WebhookPortingOrderNewCommentPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortingOrderId

`func (o *WebhookPortingOrderNewCommentPayload) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *WebhookPortingOrderNewCommentPayload) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *WebhookPortingOrderNewCommentPayload) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *WebhookPortingOrderNewCommentPayload) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetSupportKey

`func (o *WebhookPortingOrderNewCommentPayload) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *WebhookPortingOrderNewCommentPayload) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *WebhookPortingOrderNewCommentPayload) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *WebhookPortingOrderNewCommentPayload) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetComment

`func (o *WebhookPortingOrderNewCommentPayload) GetComment() WebhookPortingOrderNewCommentPayloadComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WebhookPortingOrderNewCommentPayload) GetCommentOk() (*WebhookPortingOrderNewCommentPayloadComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WebhookPortingOrderNewCommentPayload) SetComment(v WebhookPortingOrderNewCommentPayloadComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WebhookPortingOrderNewCommentPayload) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


