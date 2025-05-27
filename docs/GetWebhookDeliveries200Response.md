# GetWebhookDeliveries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WebhookDelivery**](WebhookDelivery.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMetaSimple**](PaginationMetaSimple.md) |  | [optional] 

## Methods

### NewGetWebhookDeliveries200Response

`func NewGetWebhookDeliveries200Response() *GetWebhookDeliveries200Response`

NewGetWebhookDeliveries200Response instantiates a new GetWebhookDeliveries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhookDeliveries200ResponseWithDefaults

`func NewGetWebhookDeliveries200ResponseWithDefaults() *GetWebhookDeliveries200Response`

NewGetWebhookDeliveries200ResponseWithDefaults instantiates a new GetWebhookDeliveries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetWebhookDeliveries200Response) GetData() []WebhookDelivery`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetWebhookDeliveries200Response) GetDataOk() (*[]WebhookDelivery, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetWebhookDeliveries200Response) SetData(v []WebhookDelivery)`

SetData sets Data field to given value.

### HasData

`func (o *GetWebhookDeliveries200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetWebhookDeliveries200Response) GetMeta() PaginationMetaSimple`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetWebhookDeliveries200Response) GetMetaOk() (*PaginationMetaSimple, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetWebhookDeliveries200Response) SetMeta(v PaginationMetaSimple)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetWebhookDeliveries200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


