# UpdateFaxApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | **string** | A user-assigned name to help manage the application. | 
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**WebhookEventUrl** | **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**FaxEmailRecipient** | Pointer to **NullableString** | Specifies an email address where faxes sent to this application will be forwarded to (as pdf or tiff attachments) | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the Fax Application. | [optional] 
**Inbound** | Pointer to [**CreateFaxApplicationRequestInbound**](CreateFaxApplicationRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateExternalConnectionRequestOutbound**](CreateExternalConnectionRequestOutbound.md) |  | [optional] 

## Methods

### NewUpdateFaxApplicationRequest

`func NewUpdateFaxApplicationRequest(applicationName string, webhookEventUrl string, ) *UpdateFaxApplicationRequest`

NewUpdateFaxApplicationRequest instantiates a new UpdateFaxApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFaxApplicationRequestWithDefaults

`func NewUpdateFaxApplicationRequestWithDefaults() *UpdateFaxApplicationRequest`

NewUpdateFaxApplicationRequestWithDefaults instantiates a new UpdateFaxApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *UpdateFaxApplicationRequest) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *UpdateFaxApplicationRequest) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *UpdateFaxApplicationRequest) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### GetActive

`func (o *UpdateFaxApplicationRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateFaxApplicationRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateFaxApplicationRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateFaxApplicationRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *UpdateFaxApplicationRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *UpdateFaxApplicationRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *UpdateFaxApplicationRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *UpdateFaxApplicationRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *UpdateFaxApplicationRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *UpdateFaxApplicationRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *UpdateFaxApplicationRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.


### GetWebhookEventFailoverUrl

`func (o *UpdateFaxApplicationRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *UpdateFaxApplicationRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *UpdateFaxApplicationRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *UpdateFaxApplicationRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *UpdateFaxApplicationRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *UpdateFaxApplicationRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *UpdateFaxApplicationRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *UpdateFaxApplicationRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *UpdateFaxApplicationRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *UpdateFaxApplicationRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *UpdateFaxApplicationRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *UpdateFaxApplicationRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetFaxEmailRecipient

`func (o *UpdateFaxApplicationRequest) GetFaxEmailRecipient() string`

GetFaxEmailRecipient returns the FaxEmailRecipient field if non-nil, zero value otherwise.

### GetFaxEmailRecipientOk

`func (o *UpdateFaxApplicationRequest) GetFaxEmailRecipientOk() (*string, bool)`

GetFaxEmailRecipientOk returns a tuple with the FaxEmailRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxEmailRecipient

`func (o *UpdateFaxApplicationRequest) SetFaxEmailRecipient(v string)`

SetFaxEmailRecipient sets FaxEmailRecipient field to given value.

### HasFaxEmailRecipient

`func (o *UpdateFaxApplicationRequest) HasFaxEmailRecipient() bool`

HasFaxEmailRecipient returns a boolean if a field has been set.

### SetFaxEmailRecipientNil

`func (o *UpdateFaxApplicationRequest) SetFaxEmailRecipientNil(b bool)`

 SetFaxEmailRecipientNil sets the value for FaxEmailRecipient to be an explicit nil

### UnsetFaxEmailRecipient
`func (o *UpdateFaxApplicationRequest) UnsetFaxEmailRecipient()`

UnsetFaxEmailRecipient ensures that no value is present for FaxEmailRecipient, not even an explicit nil
### GetTags

`func (o *UpdateFaxApplicationRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateFaxApplicationRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateFaxApplicationRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateFaxApplicationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetInbound

`func (o *UpdateFaxApplicationRequest) GetInbound() CreateFaxApplicationRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *UpdateFaxApplicationRequest) GetInboundOk() (*CreateFaxApplicationRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *UpdateFaxApplicationRequest) SetInbound(v CreateFaxApplicationRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *UpdateFaxApplicationRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *UpdateFaxApplicationRequest) GetOutbound() CreateExternalConnectionRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *UpdateFaxApplicationRequest) GetOutboundOk() (*CreateExternalConnectionRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *UpdateFaxApplicationRequest) SetOutbound(v CreateExternalConnectionRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *UpdateFaxApplicationRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


