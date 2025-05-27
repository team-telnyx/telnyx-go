# WebhookPortingOrderDeletedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the porting order that was deleted. | [optional] 
**CustomerReference** | Pointer to **string** | Identifies the customer reference associated with the porting order. | [optional] 
**DeletedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the porting order was deleted. | [optional] 

## Methods

### NewWebhookPortingOrderDeletedPayload

`func NewWebhookPortingOrderDeletedPayload() *WebhookPortingOrderDeletedPayload`

NewWebhookPortingOrderDeletedPayload instantiates a new WebhookPortingOrderDeletedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderDeletedPayloadWithDefaults

`func NewWebhookPortingOrderDeletedPayloadWithDefaults() *WebhookPortingOrderDeletedPayload`

NewWebhookPortingOrderDeletedPayloadWithDefaults instantiates a new WebhookPortingOrderDeletedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortingOrderDeletedPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortingOrderDeletedPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortingOrderDeletedPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortingOrderDeletedPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *WebhookPortingOrderDeletedPayload) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *WebhookPortingOrderDeletedPayload) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *WebhookPortingOrderDeletedPayload) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *WebhookPortingOrderDeletedPayload) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetDeletedAt

`func (o *WebhookPortingOrderDeletedPayload) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WebhookPortingOrderDeletedPayload) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WebhookPortingOrderDeletedPayload) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WebhookPortingOrderDeletedPayload) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


