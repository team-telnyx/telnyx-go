# CreateExternalConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the connection can be used. | [optional] [default to true]
**ExternalSipConnection** | [**ExternalSipConnectionZoomOnly**](ExternalSipConnectionZoomOnly.md) |  | [default to ZOOM]
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**Inbound** | Pointer to [**CreateExternalConnectionRequestInbound**](CreateExternalConnectionRequestInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CreateExternalConnectionRequestOutbound**](CreateExternalConnectionRequestOutbound.md) |  | [optional] 

## Methods

### NewCreateExternalConnectionRequest

`func NewCreateExternalConnectionRequest(externalSipConnection ExternalSipConnectionZoomOnly, ) *CreateExternalConnectionRequest`

NewCreateExternalConnectionRequest instantiates a new CreateExternalConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExternalConnectionRequestWithDefaults

`func NewCreateExternalConnectionRequestWithDefaults() *CreateExternalConnectionRequest`

NewCreateExternalConnectionRequestWithDefaults instantiates a new CreateExternalConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateExternalConnectionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateExternalConnectionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateExternalConnectionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateExternalConnectionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExternalSipConnection

`func (o *CreateExternalConnectionRequest) GetExternalSipConnection() ExternalSipConnectionZoomOnly`

GetExternalSipConnection returns the ExternalSipConnection field if non-nil, zero value otherwise.

### GetExternalSipConnectionOk

`func (o *CreateExternalConnectionRequest) GetExternalSipConnectionOk() (*ExternalSipConnectionZoomOnly, bool)`

GetExternalSipConnectionOk returns a tuple with the ExternalSipConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSipConnection

`func (o *CreateExternalConnectionRequest) SetExternalSipConnection(v ExternalSipConnectionZoomOnly)`

SetExternalSipConnection sets ExternalSipConnection field to given value.


### GetTags

`func (o *CreateExternalConnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateExternalConnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateExternalConnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateExternalConnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *CreateExternalConnectionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CreateExternalConnectionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CreateExternalConnectionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *CreateExternalConnectionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CreateExternalConnectionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CreateExternalConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CreateExternalConnectionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CreateExternalConnectionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CreateExternalConnectionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CreateExternalConnectionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookTimeoutSecs

`func (o *CreateExternalConnectionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CreateExternalConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CreateExternalConnectionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CreateExternalConnectionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CreateExternalConnectionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CreateExternalConnectionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetInbound

`func (o *CreateExternalConnectionRequest) GetInbound() CreateExternalConnectionRequestInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CreateExternalConnectionRequest) GetInboundOk() (*CreateExternalConnectionRequestInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CreateExternalConnectionRequest) SetInbound(v CreateExternalConnectionRequestInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *CreateExternalConnectionRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *CreateExternalConnectionRequest) GetOutbound() CreateExternalConnectionRequestOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CreateExternalConnectionRequest) GetOutboundOk() (*CreateExternalConnectionRequestOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CreateExternalConnectionRequest) SetOutbound(v CreateExternalConnectionRequestOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *CreateExternalConnectionRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


