# IpConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | Pointer to **string** |  | [optional] 
**TransportProtocol** | Pointer to **string** | One of UDP, TLS, or TCP. Applies only to connections with IP authentication or FQDN authentication. | [optional] [default to "UDP"]
**DefaultOnHoldComfortNoiseEnabled** | Pointer to **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to true]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**EncodeContactHeaderEnabled** | Pointer to **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | Pointer to [**NullableEncryptedMedia**](EncryptedMedia.md) |  | [optional] 
**OnnetT38PassthroughEnabled** | Pointer to **bool** | Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg&#39;s settings. | [optional] [default to false]
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to "1"]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**RtcpSettings** | Pointer to [**ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**Inbound** | Pointer to [**InboundIp**](InboundIp.md) |  | [optional] 
**Outbound** | Pointer to [**OutboundIp**](OutboundIp.md) |  | [optional] 

## Methods

### NewIpConnection

`func NewIpConnection() *IpConnection`

NewIpConnection instantiates a new IpConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpConnectionWithDefaults

`func NewIpConnectionWithDefaults() *IpConnection`

NewIpConnectionWithDefaults instantiates a new IpConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *IpConnection) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *IpConnection) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *IpConnection) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *IpConnection) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActive

`func (o *IpConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IpConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IpConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *IpConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *IpConnection) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *IpConnection) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *IpConnection) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *IpConnection) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *IpConnection) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *IpConnection) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *IpConnection) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *IpConnection) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *IpConnection) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *IpConnection) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *IpConnection) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *IpConnection) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *IpConnection) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *IpConnection) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *IpConnection) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *IpConnection) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *IpConnection) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *IpConnection) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *IpConnection) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *IpConnection) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *IpConnection) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *IpConnection) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *IpConnection) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *IpConnection) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *IpConnection) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *IpConnection) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *IpConnection) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *IpConnection) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *IpConnection) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *IpConnection) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetOnnetT38PassthroughEnabled

`func (o *IpConnection) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *IpConnection) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *IpConnection) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *IpConnection) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *IpConnection) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *IpConnection) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *IpConnection) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *IpConnection) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *IpConnection) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *IpConnection) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *IpConnection) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *IpConnection) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *IpConnection) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *IpConnection) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *IpConnection) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *IpConnection) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *IpConnection) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *IpConnection) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *IpConnection) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *IpConnection) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *IpConnection) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *IpConnection) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *IpConnection) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *IpConnection) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetRtcpSettings

`func (o *IpConnection) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *IpConnection) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *IpConnection) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *IpConnection) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *IpConnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpConnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpConnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpConnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetInbound

`func (o *IpConnection) GetInbound() InboundIp`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *IpConnection) GetInboundOk() (*InboundIp, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *IpConnection) SetInbound(v InboundIp)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *IpConnection) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *IpConnection) GetOutbound() OutboundIp`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *IpConnection) GetOutboundOk() (*OutboundIp, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *IpConnection) SetOutbound(v OutboundIp)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *IpConnection) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


