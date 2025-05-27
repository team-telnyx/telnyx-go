# UpdateCredentialConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**UserName** | Pointer to **string** | The user name to be used as part of the credentials. Must be 4-32 characters long and alphanumeric values only (no spaces or special characters). | [optional] 
**Password** | Pointer to **string** | The password to be used as part of the credentials. Must be 8 to 128 characters long. | [optional] 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | Pointer to **string** | A user-assigned name to help manage the connection. | [optional] 
**SipUriCallingPreference** | Pointer to **string** | This feature enables inbound SIP URI calls to your Credential Auth Connection. If enabled for all (unrestricted) then anyone who calls the SIP URI &lt;your-username&gt;@telnyx.com will be connected to your Connection. You can also choose to allow only calls that are originated on any Connections under your account (internal). | [optional] 
**DefaultOnHoldComfortNoiseEnabled** | Pointer to **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to false]
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
**Inbound** | Pointer to [**CredentialInbound**](CredentialInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CredentialOutbound**](CredentialOutbound.md) |  | [optional] 

## Methods

### NewUpdateCredentialConnectionRequest

`func NewUpdateCredentialConnectionRequest() *UpdateCredentialConnectionRequest`

NewUpdateCredentialConnectionRequest instantiates a new UpdateCredentialConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentialConnectionRequestWithDefaults

`func NewUpdateCredentialConnectionRequestWithDefaults() *UpdateCredentialConnectionRequest`

NewUpdateCredentialConnectionRequestWithDefaults instantiates a new UpdateCredentialConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateCredentialConnectionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateCredentialConnectionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateCredentialConnectionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateCredentialConnectionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserName

`func (o *UpdateCredentialConnectionRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UpdateCredentialConnectionRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UpdateCredentialConnectionRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UpdateCredentialConnectionRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateCredentialConnectionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCredentialConnectionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCredentialConnectionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateCredentialConnectionRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *UpdateCredentialConnectionRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *UpdateCredentialConnectionRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *UpdateCredentialConnectionRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *UpdateCredentialConnectionRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *UpdateCredentialConnectionRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *UpdateCredentialConnectionRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *UpdateCredentialConnectionRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *UpdateCredentialConnectionRequest) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetSipUriCallingPreference

`func (o *UpdateCredentialConnectionRequest) GetSipUriCallingPreference() string`

GetSipUriCallingPreference returns the SipUriCallingPreference field if non-nil, zero value otherwise.

### GetSipUriCallingPreferenceOk

`func (o *UpdateCredentialConnectionRequest) GetSipUriCallingPreferenceOk() (*string, bool)`

GetSipUriCallingPreferenceOk returns a tuple with the SipUriCallingPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUriCallingPreference

`func (o *UpdateCredentialConnectionRequest) SetSipUriCallingPreference(v string)`

SetSipUriCallingPreference sets SipUriCallingPreference field to given value.

### HasSipUriCallingPreference

`func (o *UpdateCredentialConnectionRequest) HasSipUriCallingPreference() bool`

HasSipUriCallingPreference returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateCredentialConnectionRequest) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *UpdateCredentialConnectionRequest) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateCredentialConnectionRequest) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *UpdateCredentialConnectionRequest) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *UpdateCredentialConnectionRequest) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *UpdateCredentialConnectionRequest) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *UpdateCredentialConnectionRequest) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *UpdateCredentialConnectionRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *UpdateCredentialConnectionRequest) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *UpdateCredentialConnectionRequest) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *UpdateCredentialConnectionRequest) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *UpdateCredentialConnectionRequest) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *UpdateCredentialConnectionRequest) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *UpdateCredentialConnectionRequest) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *UpdateCredentialConnectionRequest) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *UpdateCredentialConnectionRequest) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *UpdateCredentialConnectionRequest) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *UpdateCredentialConnectionRequest) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetOnnetT38PassthroughEnabled

`func (o *UpdateCredentialConnectionRequest) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *UpdateCredentialConnectionRequest) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *UpdateCredentialConnectionRequest) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *UpdateCredentialConnectionRequest) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetIosPushCredentialId

`func (o *UpdateCredentialConnectionRequest) GetIosPushCredentialId() string`

GetIosPushCredentialId returns the IosPushCredentialId field if non-nil, zero value otherwise.

### GetIosPushCredentialIdOk

`func (o *UpdateCredentialConnectionRequest) GetIosPushCredentialIdOk() (*string, bool)`

GetIosPushCredentialIdOk returns a tuple with the IosPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPushCredentialId

`func (o *UpdateCredentialConnectionRequest) SetIosPushCredentialId(v string)`

SetIosPushCredentialId sets IosPushCredentialId field to given value.

### HasIosPushCredentialId

`func (o *UpdateCredentialConnectionRequest) HasIosPushCredentialId() bool`

HasIosPushCredentialId returns a boolean if a field has been set.

### SetIosPushCredentialIdNil

`func (o *UpdateCredentialConnectionRequest) SetIosPushCredentialIdNil(b bool)`

 SetIosPushCredentialIdNil sets the value for IosPushCredentialId to be an explicit nil

### UnsetIosPushCredentialId
`func (o *UpdateCredentialConnectionRequest) UnsetIosPushCredentialId()`

UnsetIosPushCredentialId ensures that no value is present for IosPushCredentialId, not even an explicit nil
### GetAndroidPushCredentialId

`func (o *UpdateCredentialConnectionRequest) GetAndroidPushCredentialId() string`

GetAndroidPushCredentialId returns the AndroidPushCredentialId field if non-nil, zero value otherwise.

### GetAndroidPushCredentialIdOk

`func (o *UpdateCredentialConnectionRequest) GetAndroidPushCredentialIdOk() (*string, bool)`

GetAndroidPushCredentialIdOk returns a tuple with the AndroidPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPushCredentialId

`func (o *UpdateCredentialConnectionRequest) SetAndroidPushCredentialId(v string)`

SetAndroidPushCredentialId sets AndroidPushCredentialId field to given value.

### HasAndroidPushCredentialId

`func (o *UpdateCredentialConnectionRequest) HasAndroidPushCredentialId() bool`

HasAndroidPushCredentialId returns a boolean if a field has been set.

### SetAndroidPushCredentialIdNil

`func (o *UpdateCredentialConnectionRequest) SetAndroidPushCredentialIdNil(b bool)`

 SetAndroidPushCredentialIdNil sets the value for AndroidPushCredentialId to be an explicit nil

### UnsetAndroidPushCredentialId
`func (o *UpdateCredentialConnectionRequest) UnsetAndroidPushCredentialId()`

UnsetAndroidPushCredentialId ensures that no value is present for AndroidPushCredentialId, not even an explicit nil
### GetWebhookEventUrl

`func (o *UpdateCredentialConnectionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *UpdateCredentialConnectionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *UpdateCredentialConnectionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *UpdateCredentialConnectionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *UpdateCredentialConnectionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *UpdateCredentialConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *UpdateCredentialConnectionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *UpdateCredentialConnectionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *UpdateCredentialConnectionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *UpdateCredentialConnectionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *UpdateCredentialConnectionRequest) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *UpdateCredentialConnectionRequest) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *UpdateCredentialConnectionRequest) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *UpdateCredentialConnectionRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *UpdateCredentialConnectionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *UpdateCredentialConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *UpdateCredentialConnectionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *UpdateCredentialConnectionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *UpdateCredentialConnectionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *UpdateCredentialConnectionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetTags

`func (o *UpdateCredentialConnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCredentialConnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateCredentialConnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateCredentialConnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRtcpSettings

`func (o *UpdateCredentialConnectionRequest) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *UpdateCredentialConnectionRequest) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *UpdateCredentialConnectionRequest) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *UpdateCredentialConnectionRequest) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetInbound

`func (o *UpdateCredentialConnectionRequest) GetInbound() CredentialInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *UpdateCredentialConnectionRequest) GetInboundOk() (*CredentialInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *UpdateCredentialConnectionRequest) SetInbound(v CredentialInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *UpdateCredentialConnectionRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *UpdateCredentialConnectionRequest) GetOutbound() CredentialOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *UpdateCredentialConnectionRequest) GetOutboundOk() (*CredentialOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *UpdateCredentialConnectionRequest) SetOutbound(v CredentialOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *UpdateCredentialConnectionRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


