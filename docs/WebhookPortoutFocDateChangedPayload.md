# WebhookPortoutFocDateChangedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the port-out order that have the FOC date changed. | [optional] 
**UserId** | Pointer to **string** | Identifies the organization that port-out order belongs to. | [optional] 
**FocDate** | Pointer to **time.Time** | ISO 8601 formatted date indicating the new FOC date. | [optional] 

## Methods

### NewWebhookPortoutFocDateChangedPayload

`func NewWebhookPortoutFocDateChangedPayload() *WebhookPortoutFocDateChangedPayload`

NewWebhookPortoutFocDateChangedPayload instantiates a new WebhookPortoutFocDateChangedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortoutFocDateChangedPayloadWithDefaults

`func NewWebhookPortoutFocDateChangedPayloadWithDefaults() *WebhookPortoutFocDateChangedPayload`

NewWebhookPortoutFocDateChangedPayloadWithDefaults instantiates a new WebhookPortoutFocDateChangedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortoutFocDateChangedPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortoutFocDateChangedPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortoutFocDateChangedPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortoutFocDateChangedPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *WebhookPortoutFocDateChangedPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookPortoutFocDateChangedPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookPortoutFocDateChangedPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookPortoutFocDateChangedPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetFocDate

`func (o *WebhookPortoutFocDateChangedPayload) GetFocDate() time.Time`

GetFocDate returns the FocDate field if non-nil, zero value otherwise.

### GetFocDateOk

`func (o *WebhookPortoutFocDateChangedPayload) GetFocDateOk() (*time.Time, bool)`

GetFocDateOk returns a tuple with the FocDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocDate

`func (o *WebhookPortoutFocDateChangedPayload) SetFocDate(v time.Time)`

SetFocDate sets FocDate field to given value.

### HasFocDate

`func (o *WebhookPortoutFocDateChangedPayload) HasFocDate() bool`

HasFocDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


