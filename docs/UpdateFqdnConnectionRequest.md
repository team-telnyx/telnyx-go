# UpdateFqdnConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | Pointer to **string** | A user-assigned name to help manage the connection. | [optional] 
**TransportProtocol** | Pointer to [**FqdnConnectionTransportProtocol**](FqdnConnectionTransportProtocol.md) |  | [optional] [default to UDP]
**DefaultOnHoldComfortNoiseEnabled** | Pointer to **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to true]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**EncodeContactHeaderEnabled** | Pointer to **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | Pointer to [**NullableEncryptedMedia**](EncryptedMedia.md) |  | [optional] 
**OnnetT38PassthroughEnabled** | Pointer to **bool** | Enable on-net T38 if you prefer that the sender and receiver negotiate T38 directly when both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call according to each leg&#39;s settings. | [optional] [default to false]
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**IosPushCredentialId** | Pointer to **NullableString** | The uuid of the push credential for Ios | [optional] 
**AndroidPushCredentialId** | Pointer to **NullableString** | The uuid of the push credential for Android | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to [**WebhookApiVersion**](WebhookApiVersion.md) |  | [optional] [default to _1]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**RtcpSettings** | Pointer to [**ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] 
**Inbound** | Pointer to [**InboundFqdn**](InboundFqdn.md) |  | [optional] 
**Outbound** | Pointer to [**OutboundFqdn**](OutboundFqdn.md) |  | [optional] 

## Methods

### NewUpdateFqdnConnectionRequest

`func NewUpdateFqdnConnectionRequest() *UpdateFqdnConnectionRequest`

NewUpdateFqdnConnectionRequest instantiates a new UpdateFqdnConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFqdnConnectionRequestWithDefaults

`func NewUpdateFqdnConnectionRequestWithDefaults() *UpdateFqdnConnectionRequest`

NewUpdateFqdnConnectionRequestWithDefaults instantiates a new UpdateFqdnConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateFqdnConnectionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateFqdnConnectionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateFqdnConnectionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateFqdnConnectionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *UpdateFqdnConnectionRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *UpdateFqdnConnectionRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *UpdateFqdnConnectionRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *UpdateFqdnConnectionRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *UpdateFqdnConnectionRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *UpdateFqdnConnectionRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *UpdateFqdnConnectionRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *UpdateFqdnConnectionRequest) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *UpdateFqdnConnectionRequest) GetTransportProtocol() FqdnConnectionTransportProtocol`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *UpdateFqdnConnectionRequest) GetTransportProtocolOk() (*FqdnConnectionTransportProtocol, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *UpdateFqdnConnectionRequest) SetTransportProtocol(v FqdnConnectionTransportProtocol)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *UpdateFqdnConnectionRequest) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateFqdnConnectionRequest) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *UpdateFqdnConnectionRequest) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateFqdnConnectionRequest) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateFqdnConnectionRequest) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *UpdateFqdnConnectionRequest) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *UpdateFqdnConnectionRequest) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *UpdateFqdnConnectionRequest) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *UpdateFqdnConnectionRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *UpdateFqdnConnectionRequest) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *UpdateFqdnConnectionRequest) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *UpdateFqdnConnectionRequest) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *UpdateFqdnConnectionRequest) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *UpdateFqdnConnectionRequest) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *UpdateFqdnConnectionRequest) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *UpdateFqdnConnectionRequest) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *UpdateFqdnConnectionRequest) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *UpdateFqdnConnectionRequest) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *UpdateFqdnConnectionRequest) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetOnnetT38PassthroughEnabled

`func (o *UpdateFqdnConnectionRequest) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *UpdateFqdnConnectionRequest) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *UpdateFqdnConnectionRequest) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *UpdateFqdnConnectionRequest) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetTags

`func (o *UpdateFqdnConnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateFqdnConnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateFqdnConnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateFqdnConnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIosPushCredentialId

`func (o *UpdateFqdnConnectionRequest) GetIosPushCredentialId() string`

GetIosPushCredentialId returns the IosPushCredentialId field if non-nil, zero value otherwise.

### GetIosPushCredentialIdOk

`func (o *UpdateFqdnConnectionRequest) GetIosPushCredentialIdOk() (*string, bool)`

GetIosPushCredentialIdOk returns a tuple with the IosPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPushCredentialId

`func (o *UpdateFqdnConnectionRequest) SetIosPushCredentialId(v string)`

SetIosPushCredentialId sets IosPushCredentialId field to given value.

### HasIosPushCredentialId

`func (o *UpdateFqdnConnectionRequest) HasIosPushCredentialId() bool`

HasIosPushCredentialId returns a boolean if a field has been set.

### SetIosPushCredentialIdNil

`func (o *UpdateFqdnConnectionRequest) SetIosPushCredentialIdNil(b bool)`

 SetIosPushCredentialIdNil sets the value for IosPushCredentialId to be an explicit nil

### UnsetIosPushCredentialId
`func (o *UpdateFqdnConnectionRequest) UnsetIosPushCredentialId()`

UnsetIosPushCredentialId ensures that no value is present for IosPushCredentialId, not even an explicit nil
### GetAndroidPushCredentialId

`func (o *UpdateFqdnConnectionRequest) GetAndroidPushCredentialId() string`

GetAndroidPushCredentialId returns the AndroidPushCredentialId field if non-nil, zero value otherwise.

### GetAndroidPushCredentialIdOk

`func (o *UpdateFqdnConnectionRequest) GetAndroidPushCredentialIdOk() (*string, bool)`

GetAndroidPushCredentialIdOk returns a tuple with the AndroidPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPushCredentialId

`func (o *UpdateFqdnConnectionRequest) SetAndroidPushCredentialId(v string)`

SetAndroidPushCredentialId sets AndroidPushCredentialId field to given value.

### HasAndroidPushCredentialId

`func (o *UpdateFqdnConnectionRequest) HasAndroidPushCredentialId() bool`

HasAndroidPushCredentialId returns a boolean if a field has been set.

### SetAndroidPushCredentialIdNil

`func (o *UpdateFqdnConnectionRequest) SetAndroidPushCredentialIdNil(b bool)`

 SetAndroidPushCredentialIdNil sets the value for AndroidPushCredentialId to be an explicit nil

### UnsetAndroidPushCredentialId
`func (o *UpdateFqdnConnectionRequest) UnsetAndroidPushCredentialId()`

UnsetAndroidPushCredentialId ensures that no value is present for AndroidPushCredentialId, not even an explicit nil
### GetWebhookEventUrl

`func (o *UpdateFqdnConnectionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *UpdateFqdnConnectionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *UpdateFqdnConnectionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *UpdateFqdnConnectionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *UpdateFqdnConnectionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *UpdateFqdnConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *UpdateFqdnConnectionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *UpdateFqdnConnectionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *UpdateFqdnConnectionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *UpdateFqdnConnectionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *UpdateFqdnConnectionRequest) GetWebhookApiVersion() WebhookApiVersion`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *UpdateFqdnConnectionRequest) GetWebhookApiVersionOk() (*WebhookApiVersion, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *UpdateFqdnConnectionRequest) SetWebhookApiVersion(v WebhookApiVersion)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *UpdateFqdnConnectionRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *UpdateFqdnConnectionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *UpdateFqdnConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *UpdateFqdnConnectionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *UpdateFqdnConnectionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *UpdateFqdnConnectionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *UpdateFqdnConnectionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetRtcpSettings

`func (o *UpdateFqdnConnectionRequest) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *UpdateFqdnConnectionRequest) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *UpdateFqdnConnectionRequest) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *UpdateFqdnConnectionRequest) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetInbound

`func (o *UpdateFqdnConnectionRequest) GetInbound() InboundFqdn`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *UpdateFqdnConnectionRequest) GetInboundOk() (*InboundFqdn, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *UpdateFqdnConnectionRequest) SetInbound(v InboundFqdn)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *UpdateFqdnConnectionRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *UpdateFqdnConnectionRequest) GetOutbound() OutboundFqdn`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *UpdateFqdnConnectionRequest) GetOutboundOk() (*OutboundFqdn, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *UpdateFqdnConnectionRequest) SetOutbound(v OutboundFqdn)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *UpdateFqdnConnectionRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


