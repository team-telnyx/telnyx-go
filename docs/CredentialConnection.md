# CredentialConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**UserName** | Pointer to **string** | The user name to be used as part of the credentials. Must be 4-32 characters long and alphanumeric values only (no spaces or special characters). At least one of the first 5 characters must be a letter. | [optional] 
**Password** | Pointer to **string** | The password to be used as part of the credentials. Must be 8 to 128 characters long. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO-8601 formatted date indicating when the resource was updated. | [optional] 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | Pointer to **string** |  | [optional] 
**SipUriCallingPreference** | Pointer to **string** | This feature enables inbound SIP URI calls to your Credential Auth Connection. If enabled for all (unrestricted) then anyone who calls the SIP URI &lt;your-username&gt;@telnyx.com will be connected to your Connection. You can also choose to allow only calls that are originated on any Connections under your account (internal). | [optional] 
**DefaultOnHoldComfortNoiseEnabled** | Pointer to **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to true]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**EncodeContactHeaderEnabled** | Pointer to **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | Pointer to [**NullableEncryptedMedia**](EncryptedMedia.md) |  | [optional] 
**OnnetT38PassthroughEnabled** | Pointer to **bool** | Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg&#39;s settings. | [optional] [default to false]
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to **string** | Determines which webhook format will be used, Telnyx API v1 or v2. | [optional] [default to "1"]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**RtcpSettings** | Pointer to [**ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] 
**Inbound** | Pointer to [**CredentialInbound**](CredentialInbound.md) |  | [optional] 
**Outbound** | Pointer to [**CredentialOutbound**](CredentialOutbound.md) |  | [optional] 

## Methods

### NewCredentialConnection

`func NewCredentialConnection() *CredentialConnection`

NewCredentialConnection instantiates a new CredentialConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialConnectionWithDefaults

`func NewCredentialConnectionWithDefaults() *CredentialConnection`

NewCredentialConnectionWithDefaults instantiates a new CredentialConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *CredentialConnection) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *CredentialConnection) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *CredentialConnection) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *CredentialConnection) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActive

`func (o *CredentialConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CredentialConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CredentialConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CredentialConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUserName

`func (o *CredentialConnection) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CredentialConnection) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CredentialConnection) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *CredentialConnection) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *CredentialConnection) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialConnection) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialConnection) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialConnection) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CredentialConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *CredentialConnection) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *CredentialConnection) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *CredentialConnection) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *CredentialConnection) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *CredentialConnection) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *CredentialConnection) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *CredentialConnection) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *CredentialConnection) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetSipUriCallingPreference

`func (o *CredentialConnection) GetSipUriCallingPreference() string`

GetSipUriCallingPreference returns the SipUriCallingPreference field if non-nil, zero value otherwise.

### GetSipUriCallingPreferenceOk

`func (o *CredentialConnection) GetSipUriCallingPreferenceOk() (*string, bool)`

GetSipUriCallingPreferenceOk returns a tuple with the SipUriCallingPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUriCallingPreference

`func (o *CredentialConnection) SetSipUriCallingPreference(v string)`

SetSipUriCallingPreference sets SipUriCallingPreference field to given value.

### HasSipUriCallingPreference

`func (o *CredentialConnection) HasSipUriCallingPreference() bool`

HasSipUriCallingPreference returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *CredentialConnection) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *CredentialConnection) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *CredentialConnection) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *CredentialConnection) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *CredentialConnection) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *CredentialConnection) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *CredentialConnection) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *CredentialConnection) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *CredentialConnection) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *CredentialConnection) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *CredentialConnection) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *CredentialConnection) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *CredentialConnection) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *CredentialConnection) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *CredentialConnection) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *CredentialConnection) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *CredentialConnection) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *CredentialConnection) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetOnnetT38PassthroughEnabled

`func (o *CredentialConnection) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *CredentialConnection) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *CredentialConnection) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *CredentialConnection) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *CredentialConnection) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *CredentialConnection) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *CredentialConnection) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *CredentialConnection) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *CredentialConnection) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *CredentialConnection) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *CredentialConnection) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *CredentialConnection) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *CredentialConnection) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *CredentialConnection) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *CredentialConnection) GetWebhookApiVersion() string`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *CredentialConnection) GetWebhookApiVersionOk() (*string, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *CredentialConnection) SetWebhookApiVersion(v string)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *CredentialConnection) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *CredentialConnection) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *CredentialConnection) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *CredentialConnection) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *CredentialConnection) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *CredentialConnection) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *CredentialConnection) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetTags

`func (o *CredentialConnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CredentialConnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CredentialConnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CredentialConnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRtcpSettings

`func (o *CredentialConnection) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *CredentialConnection) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *CredentialConnection) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *CredentialConnection) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetInbound

`func (o *CredentialConnection) GetInbound() CredentialInbound`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *CredentialConnection) GetInboundOk() (*CredentialInbound, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *CredentialConnection) SetInbound(v CredentialInbound)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *CredentialConnection) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *CredentialConnection) GetOutbound() CredentialOutbound`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *CredentialConnection) GetOutboundOk() (*CredentialOutbound, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *CredentialConnection) SetOutbound(v CredentialOutbound)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *CredentialConnection) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


