# WebhookPortingOrderNewCommentPayloadComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the comment. | [optional] 
**Body** | Pointer to **string** | The body of the comment. | [optional] 
**UserId** | Pointer to **string** | Identifies the user that create the comment. | [optional] 
**UserType** | Pointer to **string** | Identifies the type of the user that created the comment. | [optional] 
**InsertedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the comment was created. | [optional] 

## Methods

### NewWebhookPortingOrderNewCommentPayloadComment

`func NewWebhookPortingOrderNewCommentPayloadComment() *WebhookPortingOrderNewCommentPayloadComment`

NewWebhookPortingOrderNewCommentPayloadComment instantiates a new WebhookPortingOrderNewCommentPayloadComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderNewCommentPayloadCommentWithDefaults

`func NewWebhookPortingOrderNewCommentPayloadCommentWithDefaults() *WebhookPortingOrderNewCommentPayloadComment`

NewWebhookPortingOrderNewCommentPayloadCommentWithDefaults instantiates a new WebhookPortingOrderNewCommentPayloadComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortingOrderNewCommentPayloadComment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortingOrderNewCommentPayloadComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBody

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WebhookPortingOrderNewCommentPayloadComment) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WebhookPortingOrderNewCommentPayloadComment) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetUserId

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookPortingOrderNewCommentPayloadComment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookPortingOrderNewCommentPayloadComment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserType

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *WebhookPortingOrderNewCommentPayloadComment) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *WebhookPortingOrderNewCommentPayloadComment) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetInsertedAt

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetInsertedAt() time.Time`

GetInsertedAt returns the InsertedAt field if non-nil, zero value otherwise.

### GetInsertedAtOk

`func (o *WebhookPortingOrderNewCommentPayloadComment) GetInsertedAtOk() (*time.Time, bool)`

GetInsertedAtOk returns a tuple with the InsertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedAt

`func (o *WebhookPortingOrderNewCommentPayloadComment) SetInsertedAt(v time.Time)`

SetInsertedAt sets InsertedAt field to given value.

### HasInsertedAt

`func (o *WebhookPortingOrderNewCommentPayloadComment) HasInsertedAt() bool`

HasInsertedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


