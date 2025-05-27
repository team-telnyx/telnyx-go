# UpdatePartnerCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookURL** | Pointer to **string** | Webhook to which campaign status updates are sent. | [optional] 
**WebhookFailoverURL** | Pointer to **string** | Webhook failover to which campaign status updates are sent. | [optional] 

## Methods

### NewUpdatePartnerCampaignRequest

`func NewUpdatePartnerCampaignRequest() *UpdatePartnerCampaignRequest`

NewUpdatePartnerCampaignRequest instantiates a new UpdatePartnerCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePartnerCampaignRequestWithDefaults

`func NewUpdatePartnerCampaignRequestWithDefaults() *UpdatePartnerCampaignRequest`

NewUpdatePartnerCampaignRequestWithDefaults instantiates a new UpdatePartnerCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookURL

`func (o *UpdatePartnerCampaignRequest) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *UpdatePartnerCampaignRequest) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *UpdatePartnerCampaignRequest) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *UpdatePartnerCampaignRequest) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetWebhookFailoverURL

`func (o *UpdatePartnerCampaignRequest) GetWebhookFailoverURL() string`

GetWebhookFailoverURL returns the WebhookFailoverURL field if non-nil, zero value otherwise.

### GetWebhookFailoverURLOk

`func (o *UpdatePartnerCampaignRequest) GetWebhookFailoverURLOk() (*string, bool)`

GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverURL

`func (o *UpdatePartnerCampaignRequest) SetWebhookFailoverURL(v string)`

SetWebhookFailoverURL sets WebhookFailoverURL field to given value.

### HasWebhookFailoverURL

`func (o *UpdatePartnerCampaignRequest) HasWebhookFailoverURL() bool`

HasWebhookFailoverURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


