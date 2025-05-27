# WebhookDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the webhook_delivery record. | [optional] 
**UserId** | Pointer to **string** | Uniquely identifies the user that owns the webhook_delivery record. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Status** | Pointer to **string** | Delivery status: &#39;delivered&#39; when successfuly delivered or &#39;failed&#39; if all attempts have failed. | [optional] 
**Webhook** | Pointer to [**WebhookDeliveryWebhook**](WebhookDeliveryWebhook.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** | ISO 8601 timestamp indicating when the first request attempt was initiated. | [optional] 
**FinishedAt** | Pointer to **time.Time** | ISO 8601 timestamp indicating when the last webhook response has been received. | [optional] 
**Attempts** | Pointer to [**[]Attempt**](Attempt.md) | Detailed delivery attempts, ordered by most recent. | [optional] 

## Methods

### NewWebhookDelivery

`func NewWebhookDelivery() *WebhookDelivery`

NewWebhookDelivery instantiates a new WebhookDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryWithDefaults

`func NewWebhookDeliveryWithDefaults() *WebhookDelivery`

NewWebhookDeliveryWithDefaults instantiates a new WebhookDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDelivery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDelivery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDelivery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookDelivery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *WebhookDelivery) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookDelivery) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookDelivery) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookDelivery) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRecordType

`func (o *WebhookDelivery) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *WebhookDelivery) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *WebhookDelivery) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *WebhookDelivery) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookDelivery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWebhook

`func (o *WebhookDelivery) GetWebhook() WebhookDeliveryWebhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *WebhookDelivery) GetWebhookOk() (*WebhookDeliveryWebhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *WebhookDelivery) SetWebhook(v WebhookDeliveryWebhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *WebhookDelivery) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetStartedAt

`func (o *WebhookDelivery) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WebhookDelivery) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WebhookDelivery) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WebhookDelivery) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *WebhookDelivery) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *WebhookDelivery) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *WebhookDelivery) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *WebhookDelivery) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetAttempts

`func (o *WebhookDelivery) GetAttempts() []Attempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *WebhookDelivery) GetAttemptsOk() (*[]Attempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *WebhookDelivery) SetAttempts(v []Attempt)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *WebhookDelivery) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


