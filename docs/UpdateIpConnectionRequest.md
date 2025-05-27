# UpdateIpConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | Pointer to **string** |  | [optional] 
**TransportProtocol** | Pointer to **string** | One of UDP, TLS, or TCP. Applies only to connections with IP authentication or FQDN authentication. | [optional] [default to "UDP"]
**DefaultOnHoldComfortNoiseEnabled** | Pointer to **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to true]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**EncodeContactHeaderEnabled** | Pointer to **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | Pointer to [**NullableEncryptedMedia**](EncryptedMedia.md) |  | [optional] 
**OnnetT38PassthroughEnabled** | Pointer to **bool** | Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg&#39;s settings. | [optional] [default to false]
**IosPushCredentialId** | Pointer to **NullableString** | The uuid of the push credential for Ios | [optional] 
**AndroidPushCredentialId** | Pointer to **NullableString** | The uuid of the push credential for Android | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to "1"]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**RtcpSettings** | Pointer to [**ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] 
**Inbound** | Pointer to [**InboundIp**](InboundIp.md) |  | [optional] 
**Outbound** | Pointer to [**OutboundIp**](OutboundIp.md) |  | [optional] 

## Methods

### NewUpdateIpConnectionRequest

`func NewUpdateIpConnectionRequest() *UpdateIpConnectionRequest`

NewUpdateIpConnectionRequest instantiates a new UpdateIpConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpConnectionRequestWithDefaults

`func NewUpdateIpConnectionRequestWithDefaults() *UpdateIpConnectionRequest`

NewUpdateIpConnectionRequestWithDefaults instantiates a new UpdateIpConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateIpConnectionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateIpConnectionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateIpConnectionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateIpConnectionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *UpdateIpConnectionRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *UpdateIpConnectionRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *UpdateIpConnectionRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *UpdateIpConnectionRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *UpdateIpConnectionRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *UpdateIpConnectionRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *UpdateIpConnectionRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *UpdateIpConnectionRequest) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *UpdateIpConnectionRequest) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *UpdateIpConnectionRequest) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *UpdateIpConnectionRequest) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *UpdateIpConnectionRequest) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateIpConnectionRequest) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *UpdateIpConnectionRequest) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateIpConnectionRequest) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateIpConnectionRequest) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *UpdateIpConnectionRequest) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *UpdateIpConnectionRequest) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *UpdateIpConnectionRequest) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *UpdateIpConnectionRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *UpdateIpConnectionRequest) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *UpdateIpConnectionRequest) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *UpdateIpConnectionRequest) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *UpdateIpConnectionRequest) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *UpdateIpConnectionRequest) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *UpdateIpConnectionRequest) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *UpdateIpConnectionRequest) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *UpdateIpConnectionRequest) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *UpdateIpConnectionRequest) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *UpdateIpConnectionRequest) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetOnnetT38PassthroughEnabled

`func (o *UpdateIpConnectionRequest) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *UpdateIpConnectionRequest) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *UpdateIpConnectionRequest) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *UpdateIpConnectionRequest) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetIosPushCredentialId

`func (o *UpdateIpConnectionRequest) GetIosPushCredentialId() string`

GetIosPushCredentialId returns the IosPushCredentialId field if non-nil, zero value otherwise.

### GetIosPushCredentialIdOk

`func (o *UpdateIpConnectionRequest) GetIosPushCredentialIdOk() (*string, bool)`

GetIosPushCredentialIdOk returns a tuple with the IosPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPushCredentialId

`func (o *UpdateIpConnectionRequest) SetIosPushCredentialId(v string)`

SetIosPushCredentialId sets IosPushCredentialId field to given value.

### HasIosPushCredentialId

`func (o *UpdateIpConnectionRequest) HasIosPushCredentialId() bool`

HasIosPushCredentialId returns a boolean if a field has been set.

### SetIosPushCredentialIdNil

`func (o *UpdateIpConnectionRequest) SetIosPushCredentialIdNil(b bool)`

 SetIosPushCredentialIdNil sets the value for IosPushCredentialId to be an explicit nil

### UnsetIosPushCredentialId
`func (o *UpdateIpConnectionRequest) UnsetIosPushCredentialId()`

UnsetIosPushCredentialId ensures that no value is present for IosPushCredentialId, not even an explicit nil
### GetAndroidPushCredentialId

`func (o *UpdateIpConnectionRequest) GetAndroidPushCredentialId() string`

GetAndroidPushCredentialId returns the AndroidPushCredentialId field if non-nil, zero value otherwise.

### GetAndroidPushCredentialIdOk

`func (o *UpdateIpConnectionRequest) GetAndroidPushCredentialIdOk() (*string, bool)`

GetAndroidPushCredentialIdOk returns a tuple with the AndroidPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPushCredentialId

`func (o *UpdateIpConnectionRequest) SetAndroidPushCredentialId(v string)`

SetAndroidPushCredentialId sets AndroidPushCredentialId field to given value.

### HasAndroidPushCredentialId

`func (o *UpdateIpConnectionRequest) HasAndroidPushCredentialId() bool`

HasAndroidPushCredentialId returns a boolean if a field has been set.

### SetAndroidPushCredentialIdNil

`func (o *UpdateIpConnectionRequest) SetAndroidPushCredentialIdNil(b bool)`

 SetAndroidPushCredentialIdNil sets the value for AndroidPushCredentialId to be an explicit nil

### UnsetAndroidPushCredentialId
`func (o *UpdateIpConnectionRequest) UnsetAndroidPushCredentialId()`

UnsetAndroidPushCredentialId ensures that no value is present for AndroidPushCredentialId, not even an explicit nil
### GetWebhookEventUrl

`func (o *UpdateIpConnectionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *UpdateIpConnectionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *UpdateIpConnectionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *UpdateIpConnectionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *UpdateIpConnectionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *UpdateIpConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *UpdateIpConnectionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *UpdateIpConnectionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *UpdateIpConnectionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *UpdateIpConnectionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *UpdateIpConnectionRequest) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *UpdateIpConnectionRequest) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *UpdateIpConnectionRequest) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *UpdateIpConnectionRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *UpdateIpConnectionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *UpdateIpConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *UpdateIpConnectionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *UpdateIpConnectionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *UpdateIpConnectionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *UpdateIpConnectionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetTags

`func (o *UpdateIpConnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateIpConnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateIpConnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateIpConnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRtcpSettings

`func (o *UpdateIpConnectionRequest) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *UpdateIpConnectionRequest) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *UpdateIpConnectionRequest) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *UpdateIpConnectionRequest) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetInbound

`func (o *UpdateIpConnectionRequest) GetInbound() InboundIp`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *UpdateIpConnectionRequest) GetInboundOk() (*InboundIp, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *UpdateIpConnectionRequest) SetInbound(v InboundIp)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *UpdateIpConnectionRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *UpdateIpConnectionRequest) GetOutbound() OutboundIp`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *UpdateIpConnectionRequest) GetOutboundOk() (*OutboundIp, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *UpdateIpConnectionRequest) SetOutbound(v OutboundIp)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *UpdateIpConnectionRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


