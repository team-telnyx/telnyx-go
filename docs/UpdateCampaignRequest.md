# UpdateCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResellerId** | Pointer to **string** | Alphanumeric identifier of the reseller that you want to associate with this campaign. | [optional] 
**Sample1** | Pointer to **string** | Message sample. Some campaign tiers require 1 or more message samples. | [optional] 
**Sample2** | Pointer to **string** | Message sample. Some campaign tiers require 2 or more message samples. | [optional] 
**Sample3** | Pointer to **string** | Message sample. Some campaign tiers require 3 or more message samples. | [optional] 
**Sample4** | Pointer to **string** | Message sample. Some campaign tiers require 4 or more message samples. | [optional] 
**Sample5** | Pointer to **string** | Message sample. Some campaign tiers require 5 or more message samples. | [optional] 
**MessageFlow** | Pointer to **string** | Message flow description. | [optional] 
**HelpMessage** | Pointer to **string** | Help message of the campaign. | [optional] 
**AutoRenewal** | Pointer to **bool** | Help message of the campaign. | [optional] [default to true]
**WebhookURL** | Pointer to **string** | Webhook to which campaign status updates are sent. | [optional] 
**WebhookFailoverURL** | Pointer to **string** | Webhook failover to which campaign status updates are sent. | [optional] 

## Methods

### NewUpdateCampaignRequest

`func NewUpdateCampaignRequest() *UpdateCampaignRequest`

NewUpdateCampaignRequest instantiates a new UpdateCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignRequestWithDefaults

`func NewUpdateCampaignRequestWithDefaults() *UpdateCampaignRequest`

NewUpdateCampaignRequestWithDefaults instantiates a new UpdateCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResellerId

`func (o *UpdateCampaignRequest) GetResellerId() string`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *UpdateCampaignRequest) GetResellerIdOk() (*string, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *UpdateCampaignRequest) SetResellerId(v string)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *UpdateCampaignRequest) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### GetSample1

`func (o *UpdateCampaignRequest) GetSample1() string`

GetSample1 returns the Sample1 field if non-nil, zero value otherwise.

### GetSample1Ok

`func (o *UpdateCampaignRequest) GetSample1Ok() (*string, bool)`

GetSample1Ok returns a tuple with the Sample1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample1

`func (o *UpdateCampaignRequest) SetSample1(v string)`

SetSample1 sets Sample1 field to given value.

### HasSample1

`func (o *UpdateCampaignRequest) HasSample1() bool`

HasSample1 returns a boolean if a field has been set.

### GetSample2

`func (o *UpdateCampaignRequest) GetSample2() string`

GetSample2 returns the Sample2 field if non-nil, zero value otherwise.

### GetSample2Ok

`func (o *UpdateCampaignRequest) GetSample2Ok() (*string, bool)`

GetSample2Ok returns a tuple with the Sample2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample2

`func (o *UpdateCampaignRequest) SetSample2(v string)`

SetSample2 sets Sample2 field to given value.

### HasSample2

`func (o *UpdateCampaignRequest) HasSample2() bool`

HasSample2 returns a boolean if a field has been set.

### GetSample3

`func (o *UpdateCampaignRequest) GetSample3() string`

GetSample3 returns the Sample3 field if non-nil, zero value otherwise.

### GetSample3Ok

`func (o *UpdateCampaignRequest) GetSample3Ok() (*string, bool)`

GetSample3Ok returns a tuple with the Sample3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample3

`func (o *UpdateCampaignRequest) SetSample3(v string)`

SetSample3 sets Sample3 field to given value.

### HasSample3

`func (o *UpdateCampaignRequest) HasSample3() bool`

HasSample3 returns a boolean if a field has been set.

### GetSample4

`func (o *UpdateCampaignRequest) GetSample4() string`

GetSample4 returns the Sample4 field if non-nil, zero value otherwise.

### GetSample4Ok

`func (o *UpdateCampaignRequest) GetSample4Ok() (*string, bool)`

GetSample4Ok returns a tuple with the Sample4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample4

`func (o *UpdateCampaignRequest) SetSample4(v string)`

SetSample4 sets Sample4 field to given value.

### HasSample4

`func (o *UpdateCampaignRequest) HasSample4() bool`

HasSample4 returns a boolean if a field has been set.

### GetSample5

`func (o *UpdateCampaignRequest) GetSample5() string`

GetSample5 returns the Sample5 field if non-nil, zero value otherwise.

### GetSample5Ok

`func (o *UpdateCampaignRequest) GetSample5Ok() (*string, bool)`

GetSample5Ok returns a tuple with the Sample5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample5

`func (o *UpdateCampaignRequest) SetSample5(v string)`

SetSample5 sets Sample5 field to given value.

### HasSample5

`func (o *UpdateCampaignRequest) HasSample5() bool`

HasSample5 returns a boolean if a field has been set.

### GetMessageFlow

`func (o *UpdateCampaignRequest) GetMessageFlow() string`

GetMessageFlow returns the MessageFlow field if non-nil, zero value otherwise.

### GetMessageFlowOk

`func (o *UpdateCampaignRequest) GetMessageFlowOk() (*string, bool)`

GetMessageFlowOk returns a tuple with the MessageFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFlow

`func (o *UpdateCampaignRequest) SetMessageFlow(v string)`

SetMessageFlow sets MessageFlow field to given value.

### HasMessageFlow

`func (o *UpdateCampaignRequest) HasMessageFlow() bool`

HasMessageFlow returns a boolean if a field has been set.

### GetHelpMessage

`func (o *UpdateCampaignRequest) GetHelpMessage() string`

GetHelpMessage returns the HelpMessage field if non-nil, zero value otherwise.

### GetHelpMessageOk

`func (o *UpdateCampaignRequest) GetHelpMessageOk() (*string, bool)`

GetHelpMessageOk returns a tuple with the HelpMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpMessage

`func (o *UpdateCampaignRequest) SetHelpMessage(v string)`

SetHelpMessage sets HelpMessage field to given value.

### HasHelpMessage

`func (o *UpdateCampaignRequest) HasHelpMessage() bool`

HasHelpMessage returns a boolean if a field has been set.

### GetAutoRenewal

`func (o *UpdateCampaignRequest) GetAutoRenewal() bool`

GetAutoRenewal returns the AutoRenewal field if non-nil, zero value otherwise.

### GetAutoRenewalOk

`func (o *UpdateCampaignRequest) GetAutoRenewalOk() (*bool, bool)`

GetAutoRenewalOk returns a tuple with the AutoRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewal

`func (o *UpdateCampaignRequest) SetAutoRenewal(v bool)`

SetAutoRenewal sets AutoRenewal field to given value.

### HasAutoRenewal

`func (o *UpdateCampaignRequest) HasAutoRenewal() bool`

HasAutoRenewal returns a boolean if a field has been set.

### GetWebhookURL

`func (o *UpdateCampaignRequest) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *UpdateCampaignRequest) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *UpdateCampaignRequest) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *UpdateCampaignRequest) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetWebhookFailoverURL

`func (o *UpdateCampaignRequest) GetWebhookFailoverURL() string`

GetWebhookFailoverURL returns the WebhookFailoverURL field if non-nil, zero value otherwise.

### GetWebhookFailoverURLOk

`func (o *UpdateCampaignRequest) GetWebhookFailoverURLOk() (*string, bool)`

GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverURL

`func (o *UpdateCampaignRequest) SetWebhookFailoverURL(v string)`

SetWebhookFailoverURL sets WebhookFailoverURL field to given value.

### HasWebhookFailoverURL

`func (o *UpdateCampaignRequest) HasWebhookFailoverURL() bool`

HasWebhookFailoverURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


