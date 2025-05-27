# WebhookPortoutNewCommentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the comment that was added to the port-out order. | [optional] 
**PortoutId** | Pointer to **string** | Identifies the port-out order that the comment was added to. | [optional] 
**UserId** | Pointer to **string** | Identifies the user that added the comment. | [optional] 
**Comment** | Pointer to **string** | The body of the comment. | [optional] 

## Methods

### NewWebhookPortoutNewCommentPayload

`func NewWebhookPortoutNewCommentPayload() *WebhookPortoutNewCommentPayload`

NewWebhookPortoutNewCommentPayload instantiates a new WebhookPortoutNewCommentPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortoutNewCommentPayloadWithDefaults

`func NewWebhookPortoutNewCommentPayloadWithDefaults() *WebhookPortoutNewCommentPayload`

NewWebhookPortoutNewCommentPayloadWithDefaults instantiates a new WebhookPortoutNewCommentPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortoutNewCommentPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortoutNewCommentPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortoutNewCommentPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortoutNewCommentPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortoutId

`func (o *WebhookPortoutNewCommentPayload) GetPortoutId() string`

GetPortoutId returns the PortoutId field if non-nil, zero value otherwise.

### GetPortoutIdOk

`func (o *WebhookPortoutNewCommentPayload) GetPortoutIdOk() (*string, bool)`

GetPortoutIdOk returns a tuple with the PortoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortoutId

`func (o *WebhookPortoutNewCommentPayload) SetPortoutId(v string)`

SetPortoutId sets PortoutId field to given value.

### HasPortoutId

`func (o *WebhookPortoutNewCommentPayload) HasPortoutId() bool`

HasPortoutId returns a boolean if a field has been set.

### GetUserId

`func (o *WebhookPortoutNewCommentPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookPortoutNewCommentPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookPortoutNewCommentPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookPortoutNewCommentPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetComment

`func (o *WebhookPortoutNewCommentPayload) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WebhookPortoutNewCommentPayload) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WebhookPortoutNewCommentPayload) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WebhookPortoutNewCommentPayload) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


