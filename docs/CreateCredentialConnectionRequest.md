# CreateCredentialConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**UserName** | **string** | The user name to be used as part of the credentials. Must be 4-32 characters long and alphanumeric values only (no spaces or special characters). | 
**Password** | **string** | The password to be used as part of the credentials. Must be 8 to 128 characters long. | 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | **string** | A user-assigned name to help manage the connection. | 
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
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1, v2 or texml. Note - texml can only be set when the outbound object parameter call_parking_enabled is included and set to true. | [optional] [default to "1"]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**RtcpSettings** | Pointer to [**ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] 
**Inbound** | Pointer to [**CredentialInbound**](CredentialInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CredentialOutbound**](CredentialOutbound.md) |  | [optional] 

## Methods

### NewCreateCredentialConnectionRequest

`func NewCreateCredentialConnectionRequest(userName string, password string, connectionName string, ) *CreateCredentialConnectionRequest`

NewCreateCredentialConnectionRequest instantiates a new CreateCredentialConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCredentialConnectionRequestWithDefaults

`func NewCreateCredentialConnectionRequestWithDefaults() *CreateCredentialConnectionRequest`

NewCreateCredentialConnectionRequestWithDefaults instantiates a new CreateCredentialConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CreateCredentialConnectionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateCredentialConnectionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateCredentialConnectionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateCredentialConnectionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserName

`func (o *CreateCredentialConnectionRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateCredentialConnectionRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateCredentialConnectionRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *CreateCredentialConnectionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCredentialConnectionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCredentialConnectionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAnchorsiteOverride

`func (o *CreateCredentialConnectionRequest) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *CreateCredentialConnectionRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *CreateCredentialConnectionRequest) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *CreateCredentialConnectionRequest) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *CreateCredentialConnectionRequest) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *CreateCredentialConnectionRequest) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *CreateCredentialConnectionRequest) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetSipUriCallingPreference

`func (o *CreateCredentialConnectionRequest) GetSipUriCallingPreference() string`

GetSipUriCallingPreference returns the SipUriCallingPreference field if non-nil, zero value otherwise.

### GetSipUriCallingPreferenceOk

`func (o *CreateCredentialConnectionRequest) GetSipUriCallingPreferenceOk() (*string, bool)`

GetSipUriCallingPreferenceOk returns a tuple with the SipUriCallingPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUriCallingPreference

`func (o *CreateCredentialConnectionRequest) SetSipUriCallingPreference(v string)`

SetSipUriCallingPreference sets SipUriCallingPreference field to given value.

### HasSipUriCallingPreference

`func (o *CreateCredentialConnectionRequest) HasSipUriCallingPreference() bool`

HasSipUriCallingPreference returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *CreateCredentialConnectionRequest) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *CreateCredentialConnectionRequest) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *CreateCredentialConnectionRequest) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *CreateCredentialConnectionRequest) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *CreateCredentialConnectionRequest) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *CreateCredentialConnectionRequest) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *CreateCredentialConnectionRequest) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *CreateCredentialConnectionRequest) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *CreateCredentialConnectionRequest) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *CreateCredentialConnectionRequest) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *CreateCredentialConnectionRequest) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *CreateCredentialConnectionRequest) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *CreateCredentialConnectionRequest) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *CreateCredentialConnectionRequest) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *CreateCredentialConnectionRequest) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *CreateCredentialConnectionRequest) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *CreateCredentialConnectionRequest) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *CreateCredentialConnectionRequest) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetOnnetT38PassthroughEnabled

`func (o *CreateCredentialConnectionRequest) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *CreateCredentialConnectionRequest) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *CreateCredentialConnectionRequest) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *CreateCredentialConnectionRequest) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetIosPushCredentialId

`func (o *CreateCredentialConnectionRequest) GetIosPushCredentialId() string`

GetIosPushCredentialId returns the IosPushCredentialId field if non-nil, zero value otherwise.

### GetIosPushCredentialIdOk

`func (o *CreateCredentialConnectionRequest) GetIosPushCredentialIdOk() (*string, bool)`

GetIosPushCredentialIdOk returns a tuple with the IosPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPushCredentialId

`func (o *CreateCredentialConnectionRequest) SetIosPushCredentialId(v string)`

SetIosPushCredentialId sets IosPushCredentialId field to given value.

### HasIosPushCredentialId

`func (o *CreateCredentialConnectionRequest) HasIosPushCredentialId() bool`

HasIosPushCredentialId returns a boolean if a field has been set.

### SetIosPushCredentialIdNil

`func (o *CreateCredentialConnectionRequest) SetIosPushCredentialIdNil(b bool)`

 SetIosPushCredentialIdNil sets the value for IosPushCredentialId to be an explicit nil

### UnsetIosPushCredentialId
`func (o *CreateCredentialConnectionRequest) UnsetIosPushCredentialId()`

UnsetIosPushCredentialId ensures that no value is present for IosPushCredentialId, not even an explicit nil
### GetAndroidPushCredentialId

`func (o *CreateCredentialConnectionRequest) GetAndroidPushCredentialId() string`

GetAndroidPushCredentialId returns the AndroidPushCredentialId field if non-nil, zero value otherwise.

### GetAndroidPushCredentialIdOk

`func (o *CreateCredentialConnectionRequest) GetAndroidPushCredentialIdOk() (*string, bool)`

GetAndroidPushCredentialIdOk returns a tuple with the AndroidPushCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPushCredentialId

`func (o *CreateCredentialConnectionRequest) SetAndroidPushCredentialId(v string)`

SetAndroidPushCredentialId sets AndroidPushCredentialId field to given value.

### HasAndroidPushCredentialId

`func (o *CreateCredentialConnectionRequest) HasAndroidPushCredentialId() bool`

HasAndroidPushCredentialId returns a boolean if a field has been set.

### SetAndroidPushCredentialIdNil

`func (o *CreateCredentialConnectionRequest) SetAndroidPushCredentialIdNil(b bool)`

 SetAndroidPushCredentialIdNil sets the value for AndroidPushCredentialId to be an explicit nil

### UnsetAndroidPushCredentialId
`func (o *CreateCredentialConnectionRequest) UnsetAndroidPushCredentialId()`

UnsetAndroidPushCredentialId ensures that no value is present for AndroidPushCredentialId, not even an explicit nil
### GetWebhookEventUrl

`func (o *CreateCredentialConnectionRequest) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CreateCredentialConnectionRequest) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CreateCredentialConnectionRequest) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *CreateCredentialConnectionRequest) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CreateCredentialConnectionRequest) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CreateCredentialConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CreateCredentialConnectionRequest) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CreateCredentialConnectionRequest) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CreateCredentialConnectionRequest) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CreateCredentialConnectionRequest) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *CreateCredentialConnectionRequest) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *CreateCredentialConnectionRequest) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *CreateCredentialConnectionRequest) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *CreateCredentialConnectionRequest) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *CreateCredentialConnectionRequest) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CreateCredentialConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CreateCredentialConnectionRequest) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CreateCredentialConnectionRequest) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CreateCredentialConnectionRequest) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CreateCredentialConnectionRequest) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetTags

`func (o *CreateCredentialConnectionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCredentialConnectionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCredentialConnectionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCredentialConnectionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRtcpSettings

`func (o *CreateCredentialConnectionRequest) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *CreateCredentialConnectionRequest) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *CreateCredentialConnectionRequest) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *CreateCredentialConnectionRequest) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetInbound

`func (o *CreateCredentialConnectionRequest) GetInbound() CredentialInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CreateCredentialConnectionRequest) GetInboundOk() (*CredentialInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CreateCredentialConnectionRequest) SetInbound(v CredentialInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *CreateCredentialConnectionRequest) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *CreateCredentialConnectionRequest) GetOutbound() CredentialOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CreateCredentialConnectionRequest) GetOutboundOk() (*CredentialOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CreateCredentialConnectionRequest) SetOutbound(v CredentialOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *CreateCredentialConnectionRequest) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


