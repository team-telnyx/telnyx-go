# UpdateExternalConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**Inbound** | Pointer to [**CreateExternalConnectionRequestInbound**](CreateExternalConnectionRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateExternalConnectionRequestOutbound**](CreateExternalConnectionRequestOutbound.md) |  | [optional] 

## Methods

### NewUpdateExternalConnectionRequest

`func NewUpdateExternalConnectionRequest() *UpdateExternalConnectionRequest`

NewUpdateExternalConnectionRequest instantiates a new UpdateExternalConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExternalConnectionRequestWithDefaults

`func NewUpdateExternalConnectionRequestWithDefaults() *UpdateExternalConnectionRequest`

NewUpdateExternalConnectionRequestWithDefaults instantiates a new UpdateExternalConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateExternalConnectionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateExternalConnectionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateExternalConnectionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateExternalConnectionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *UpdateExternalConnectionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *UpdateExternalConnectionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *UpdateExternalConnectionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *UpdateExternalConnectionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *UpdateExternalConnectionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *UpdateExternalConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *UpdateExternalConnectionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *UpdateExternalConnectionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *UpdateExternalConnectionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *UpdateExternalConnectionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetTags

`func (o *UpdateExternalConnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateExternalConnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateExternalConnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateExternalConnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *UpdateExternalConnectionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *UpdateExternalConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *UpdateExternalConnectionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *UpdateExternalConnectionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *UpdateExternalConnectionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *UpdateExternalConnectionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetInbound

`func (o *UpdateExternalConnectionRequest) GetInbound() CreateExternalConnectionRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *UpdateExternalConnectionRequest) GetInboundOk() (*CreateExternalConnectionRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *UpdateExternalConnectionRequest) SetInbound(v CreateExternalConnectionRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *UpdateExternalConnectionRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *UpdateExternalConnectionRequest) GetOutbound() CreateExternalConnectionRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *UpdateExternalConnectionRequest) GetOutboundOk() (*CreateExternalConnectionRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *UpdateExternalConnectionRequest) SetOutbound(v CreateExternalConnectionRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *UpdateExternalConnectionRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


