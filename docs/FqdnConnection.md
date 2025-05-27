# FqdnConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Active** | Pointer to **bool** | Defaults to true | [optional] 
**AnchorsiteOverride** | Pointer to [**AnchorsiteOverride**](AnchorsiteOverride.md) |  | [optional] [default to LATENCY]
**ConnectionName** | **string** | A user-assigned name to help manage the connection. | 
**TransportProtocol** | Pointer to [**FqdnConnectionTransportProtocol**](FqdnConnectionTransportProtocol.md) |  | [optional] [default to UDP]
**DefaultOnHoldComfortNoiseEnabled** | Pointer to **bool** | When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout. | [optional] [default to true]
**DtmfType** | Pointer to [**DtmfType**](DtmfType.md) |  | [optional] [default to RFC_2833]
**EncodeContactHeaderEnabled** | Pointer to **bool** | Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios. | [optional] [default to false]
**EncryptedMedia** | Pointer to [**NullableEncryptedMedia**](EncryptedMedia.md) |  | [optional] 
**MicrosoftTeamsSbc** | Pointer to **bool** | The connection is enabled for Microsoft Teams Direct Routing. | [optional] [default to false]
**OnnetT38PassthroughEnabled** | Pointer to **bool** | Enable on-net T38 if you prefer that the sender and receiver negotiate T38 directly when both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call according to each leg&#39;s settings. | [optional] [default to false]
**UserName** | Pointer to **string** | The username for the FQDN connection. | [optional] 
**Password** | Pointer to **string** | The password for the FQDN connection. | [optional] 
**RtpPassCodecsOnStreamChange** | Pointer to **bool** | Defines if codecs should be passed on stream change. | [optional] 
**AdjustDtmfTimestamp** | Pointer to **bool** | Indicates whether DTMF timestamp adjustment is enabled. | [optional] 
**IgnoreDtmfDuration** | Pointer to **bool** | Indicates whether DTMF duration should be ignored. | [optional] 
**IgnoreMarkBit** | Pointer to **bool** | Indicates whether the mark bit should be ignored. | [optional] 
**CallCostEnabled** | Pointer to **bool** | Indicates whether call cost calculation is enabled. | [optional] 
**NoiseSuppression** | Pointer to **bool** | Indicates whether noise suppression is enabled. | [optional] 
**SendNormalizedTimestamps** | Pointer to **bool** | Indicates whether normalized timestamps should be sent. | [optional] 
**ThirdPartyControlEnabled** | Pointer to **bool** | Indicates whether third-party control is enabled. | [optional] 
**TxtName** | Pointer to **string** | The name for the TXT record associated with the FQDN connection. | [optional] 
**TxtValue** | Pointer to **string** | The value for the TXT record associated with the FQDN connection. | [optional] 
**TxtTtl** | Pointer to **int32** | The time to live for the TXT record associated with the FQDN connection. | [optional] 
**Tags** | Pointer to **[]string** | Tags associated with the connection. | [optional] 
**WebhookEventUrl** | Pointer to **string** | The URL where webhooks related to this connection will be sent. Must include a scheme, such as &#39;https&#39;. | [optional] 
**WebhookEventFailoverUrl** | Pointer to **NullableString** | The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as &#39;https&#39;. | [optional] [default to ""]
**WebhookApiVersion** | Pointer to [**WebhookApiVersion**](WebhookApiVersion.md) |  | [optional] [default to _1]
**WebhookTimeoutSecs** | Pointer to **NullableInt32** | Specifies how many seconds to wait before timing out a webhook. | [optional] 
**RtcpSettings** | Pointer to [**ConnectionRtcpSettings**](ConnectionRtcpSettings.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 
**Inbound** | Pointer to [**InboundFqdn**](InboundFqdn.md) |  | [optional] 
**Outbound** | Pointer to [**OutboundFqdn**](OutboundFqdn.md) |  | [optional] 

## Methods

### NewFqdnConnection

`func NewFqdnConnection(connectionName string, ) *FqdnConnection`

NewFqdnConnection instantiates a new FqdnConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFqdnConnectionWithDefaults

`func NewFqdnConnectionWithDefaults() *FqdnConnection`

NewFqdnConnectionWithDefaults instantiates a new FqdnConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FqdnConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FqdnConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FqdnConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FqdnConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *FqdnConnection) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *FqdnConnection) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *FqdnConnection) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *FqdnConnection) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetActive

`func (o *FqdnConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FqdnConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FqdnConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FqdnConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAnchorsiteOverride

`func (o *FqdnConnection) GetAnchorsiteOverride() AnchorsiteOverride`

GetAnchorsiteOverride returns the AnchorsiteOverride field if non-nil, zero value otherwise.

### GetAnchorsiteOverrideOk

`func (o *FqdnConnection) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool)`

GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorsiteOverride

`func (o *FqdnConnection) SetAnchorsiteOverride(v AnchorsiteOverride)`

SetAnchorsiteOverride sets AnchorsiteOverride field to given value.

### HasAnchorsiteOverride

`func (o *FqdnConnection) HasAnchorsiteOverride() bool`

HasAnchorsiteOverride returns a boolean if a field has been set.

### GetConnectionName

`func (o *FqdnConnection) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *FqdnConnection) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *FqdnConnection) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetTransportProtocol

`func (o *FqdnConnection) GetTransportProtocol() FqdnConnectionTransportProtocol`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *FqdnConnection) GetTransportProtocolOk() (*FqdnConnectionTransportProtocol, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *FqdnConnection) SetTransportProtocol(v FqdnConnectionTransportProtocol)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *FqdnConnection) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetDefaultOnHoldComfortNoiseEnabled

`func (o *FqdnConnection) GetDefaultOnHoldComfortNoiseEnabled() bool`

GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field if non-nil, zero value otherwise.

### GetDefaultOnHoldComfortNoiseEnabledOk

`func (o *FqdnConnection) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool)`

GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOnHoldComfortNoiseEnabled

`func (o *FqdnConnection) SetDefaultOnHoldComfortNoiseEnabled(v bool)`

SetDefaultOnHoldComfortNoiseEnabled sets DefaultOnHoldComfortNoiseEnabled field to given value.

### HasDefaultOnHoldComfortNoiseEnabled

`func (o *FqdnConnection) HasDefaultOnHoldComfortNoiseEnabled() bool`

HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.

### GetDtmfType

`func (o *FqdnConnection) GetDtmfType() DtmfType`

GetDtmfType returns the DtmfType field if non-nil, zero value otherwise.

### GetDtmfTypeOk

`func (o *FqdnConnection) GetDtmfTypeOk() (*DtmfType, bool)`

GetDtmfTypeOk returns a tuple with the DtmfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfType

`func (o *FqdnConnection) SetDtmfType(v DtmfType)`

SetDtmfType sets DtmfType field to given value.

### HasDtmfType

`func (o *FqdnConnection) HasDtmfType() bool`

HasDtmfType returns a boolean if a field has been set.

### GetEncodeContactHeaderEnabled

`func (o *FqdnConnection) GetEncodeContactHeaderEnabled() bool`

GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field if non-nil, zero value otherwise.

### GetEncodeContactHeaderEnabledOk

`func (o *FqdnConnection) GetEncodeContactHeaderEnabledOk() (*bool, bool)`

GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodeContactHeaderEnabled

`func (o *FqdnConnection) SetEncodeContactHeaderEnabled(v bool)`

SetEncodeContactHeaderEnabled sets EncodeContactHeaderEnabled field to given value.

### HasEncodeContactHeaderEnabled

`func (o *FqdnConnection) HasEncodeContactHeaderEnabled() bool`

HasEncodeContactHeaderEnabled returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *FqdnConnection) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *FqdnConnection) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *FqdnConnection) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *FqdnConnection) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *FqdnConnection) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *FqdnConnection) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetMicrosoftTeamsSbc

`func (o *FqdnConnection) GetMicrosoftTeamsSbc() bool`

GetMicrosoftTeamsSbc returns the MicrosoftTeamsSbc field if non-nil, zero value otherwise.

### GetMicrosoftTeamsSbcOk

`func (o *FqdnConnection) GetMicrosoftTeamsSbcOk() (*bool, bool)`

GetMicrosoftTeamsSbcOk returns a tuple with the MicrosoftTeamsSbc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTeamsSbc

`func (o *FqdnConnection) SetMicrosoftTeamsSbc(v bool)`

SetMicrosoftTeamsSbc sets MicrosoftTeamsSbc field to given value.

### HasMicrosoftTeamsSbc

`func (o *FqdnConnection) HasMicrosoftTeamsSbc() bool`

HasMicrosoftTeamsSbc returns a boolean if a field has been set.

### GetOnnetT38PassthroughEnabled

`func (o *FqdnConnection) GetOnnetT38PassthroughEnabled() bool`

GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field if non-nil, zero value otherwise.

### GetOnnetT38PassthroughEnabledOk

`func (o *FqdnConnection) GetOnnetT38PassthroughEnabledOk() (*bool, bool)`

GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnnetT38PassthroughEnabled

`func (o *FqdnConnection) SetOnnetT38PassthroughEnabled(v bool)`

SetOnnetT38PassthroughEnabled sets OnnetT38PassthroughEnabled field to given value.

### HasOnnetT38PassthroughEnabled

`func (o *FqdnConnection) HasOnnetT38PassthroughEnabled() bool`

HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.

### GetUserName

`func (o *FqdnConnection) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FqdnConnection) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FqdnConnection) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *FqdnConnection) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *FqdnConnection) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FqdnConnection) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FqdnConnection) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FqdnConnection) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRtpPassCodecsOnStreamChange

`func (o *FqdnConnection) GetRtpPassCodecsOnStreamChange() bool`

GetRtpPassCodecsOnStreamChange returns the RtpPassCodecsOnStreamChange field if non-nil, zero value otherwise.

### GetRtpPassCodecsOnStreamChangeOk

`func (o *FqdnConnection) GetRtpPassCodecsOnStreamChangeOk() (*bool, bool)`

GetRtpPassCodecsOnStreamChangeOk returns a tuple with the RtpPassCodecsOnStreamChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtpPassCodecsOnStreamChange

`func (o *FqdnConnection) SetRtpPassCodecsOnStreamChange(v bool)`

SetRtpPassCodecsOnStreamChange sets RtpPassCodecsOnStreamChange field to given value.

### HasRtpPassCodecsOnStreamChange

`func (o *FqdnConnection) HasRtpPassCodecsOnStreamChange() bool`

HasRtpPassCodecsOnStreamChange returns a boolean if a field has been set.

### GetAdjustDtmfTimestamp

`func (o *FqdnConnection) GetAdjustDtmfTimestamp() bool`

GetAdjustDtmfTimestamp returns the AdjustDtmfTimestamp field if non-nil, zero value otherwise.

### GetAdjustDtmfTimestampOk

`func (o *FqdnConnection) GetAdjustDtmfTimestampOk() (*bool, bool)`

GetAdjustDtmfTimestampOk returns a tuple with the AdjustDtmfTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustDtmfTimestamp

`func (o *FqdnConnection) SetAdjustDtmfTimestamp(v bool)`

SetAdjustDtmfTimestamp sets AdjustDtmfTimestamp field to given value.

### HasAdjustDtmfTimestamp

`func (o *FqdnConnection) HasAdjustDtmfTimestamp() bool`

HasAdjustDtmfTimestamp returns a boolean if a field has been set.

### GetIgnoreDtmfDuration

`func (o *FqdnConnection) GetIgnoreDtmfDuration() bool`

GetIgnoreDtmfDuration returns the IgnoreDtmfDuration field if non-nil, zero value otherwise.

### GetIgnoreDtmfDurationOk

`func (o *FqdnConnection) GetIgnoreDtmfDurationOk() (*bool, bool)`

GetIgnoreDtmfDurationOk returns a tuple with the IgnoreDtmfDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDtmfDuration

`func (o *FqdnConnection) SetIgnoreDtmfDuration(v bool)`

SetIgnoreDtmfDuration sets IgnoreDtmfDuration field to given value.

### HasIgnoreDtmfDuration

`func (o *FqdnConnection) HasIgnoreDtmfDuration() bool`

HasIgnoreDtmfDuration returns a boolean if a field has been set.

### GetIgnoreMarkBit

`func (o *FqdnConnection) GetIgnoreMarkBit() bool`

GetIgnoreMarkBit returns the IgnoreMarkBit field if non-nil, zero value otherwise.

### GetIgnoreMarkBitOk

`func (o *FqdnConnection) GetIgnoreMarkBitOk() (*bool, bool)`

GetIgnoreMarkBitOk returns a tuple with the IgnoreMarkBit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreMarkBit

`func (o *FqdnConnection) SetIgnoreMarkBit(v bool)`

SetIgnoreMarkBit sets IgnoreMarkBit field to given value.

### HasIgnoreMarkBit

`func (o *FqdnConnection) HasIgnoreMarkBit() bool`

HasIgnoreMarkBit returns a boolean if a field has been set.

### GetCallCostEnabled

`func (o *FqdnConnection) GetCallCostEnabled() bool`

GetCallCostEnabled returns the CallCostEnabled field if non-nil, zero value otherwise.

### GetCallCostEnabledOk

`func (o *FqdnConnection) GetCallCostEnabledOk() (*bool, bool)`

GetCallCostEnabledOk returns a tuple with the CallCostEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallCostEnabled

`func (o *FqdnConnection) SetCallCostEnabled(v bool)`

SetCallCostEnabled sets CallCostEnabled field to given value.

### HasCallCostEnabled

`func (o *FqdnConnection) HasCallCostEnabled() bool`

HasCallCostEnabled returns a boolean if a field has been set.

### GetNoiseSuppression

`func (o *FqdnConnection) GetNoiseSuppression() bool`

GetNoiseSuppression returns the NoiseSuppression field if non-nil, zero value otherwise.

### GetNoiseSuppressionOk

`func (o *FqdnConnection) GetNoiseSuppressionOk() (*bool, bool)`

GetNoiseSuppressionOk returns a tuple with the NoiseSuppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoiseSuppression

`func (o *FqdnConnection) SetNoiseSuppression(v bool)`

SetNoiseSuppression sets NoiseSuppression field to given value.

### HasNoiseSuppression

`func (o *FqdnConnection) HasNoiseSuppression() bool`

HasNoiseSuppression returns a boolean if a field has been set.

### GetSendNormalizedTimestamps

`func (o *FqdnConnection) GetSendNormalizedTimestamps() bool`

GetSendNormalizedTimestamps returns the SendNormalizedTimestamps field if non-nil, zero value otherwise.

### GetSendNormalizedTimestampsOk

`func (o *FqdnConnection) GetSendNormalizedTimestampsOk() (*bool, bool)`

GetSendNormalizedTimestampsOk returns a tuple with the SendNormalizedTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNormalizedTimestamps

`func (o *FqdnConnection) SetSendNormalizedTimestamps(v bool)`

SetSendNormalizedTimestamps sets SendNormalizedTimestamps field to given value.

### HasSendNormalizedTimestamps

`func (o *FqdnConnection) HasSendNormalizedTimestamps() bool`

HasSendNormalizedTimestamps returns a boolean if a field has been set.

### GetThirdPartyControlEnabled

`func (o *FqdnConnection) GetThirdPartyControlEnabled() bool`

GetThirdPartyControlEnabled returns the ThirdPartyControlEnabled field if non-nil, zero value otherwise.

### GetThirdPartyControlEnabledOk

`func (o *FqdnConnection) GetThirdPartyControlEnabledOk() (*bool, bool)`

GetThirdPartyControlEnabledOk returns a tuple with the ThirdPartyControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyControlEnabled

`func (o *FqdnConnection) SetThirdPartyControlEnabled(v bool)`

SetThirdPartyControlEnabled sets ThirdPartyControlEnabled field to given value.

### HasThirdPartyControlEnabled

`func (o *FqdnConnection) HasThirdPartyControlEnabled() bool`

HasThirdPartyControlEnabled returns a boolean if a field has been set.

### GetTxtName

`func (o *FqdnConnection) GetTxtName() string`

GetTxtName returns the TxtName field if non-nil, zero value otherwise.

### GetTxtNameOk

`func (o *FqdnConnection) GetTxtNameOk() (*string, bool)`

GetTxtNameOk returns a tuple with the TxtName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtName

`func (o *FqdnConnection) SetTxtName(v string)`

SetTxtName sets TxtName field to given value.

### HasTxtName

`func (o *FqdnConnection) HasTxtName() bool`

HasTxtName returns a boolean if a field has been set.

### GetTxtValue

`func (o *FqdnConnection) GetTxtValue() string`

GetTxtValue returns the TxtValue field if non-nil, zero value otherwise.

### GetTxtValueOk

`func (o *FqdnConnection) GetTxtValueOk() (*string, bool)`

GetTxtValueOk returns a tuple with the TxtValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtValue

`func (o *FqdnConnection) SetTxtValue(v string)`

SetTxtValue sets TxtValue field to given value.

### HasTxtValue

`func (o *FqdnConnection) HasTxtValue() bool`

HasTxtValue returns a boolean if a field has been set.

### GetTxtTtl

`func (o *FqdnConnection) GetTxtTtl() int32`

GetTxtTtl returns the TxtTtl field if non-nil, zero value otherwise.

### GetTxtTtlOk

`func (o *FqdnConnection) GetTxtTtlOk() (*int32, bool)`

GetTxtTtlOk returns a tuple with the TxtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtTtl

`func (o *FqdnConnection) SetTxtTtl(v int32)`

SetTxtTtl sets TxtTtl field to given value.

### HasTxtTtl

`func (o *FqdnConnection) HasTxtTtl() bool`

HasTxtTtl returns a boolean if a field has been set.

### GetTags

`func (o *FqdnConnection) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FqdnConnection) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FqdnConnection) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FqdnConnection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWebhookEventUrl

`func (o *FqdnConnection) GetWebhookEventUrl() string`

GetWebhookEventUrl returns the WebhookEventUrl field if non-nil, zero value otherwise.

### GetWebhookEventUrlOk

`func (o *FqdnConnection) GetWebhookEventUrlOk() (*string, bool)`

GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventUrl

`func (o *FqdnConnection) SetWebhookEventUrl(v string)`

SetWebhookEventUrl sets WebhookEventUrl field to given value.

### HasWebhookEventUrl

`func (o *FqdnConnection) HasWebhookEventUrl() bool`

HasWebhookEventUrl returns a boolean if a field has been set.

### GetWebhookEventFailoverUrl

`func (o *FqdnConnection) GetWebhookEventFailoverUrl() string`

GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field if non-nil, zero value otherwise.

### GetWebhookEventFailoverUrlOk

`func (o *FqdnConnection) GetWebhookEventFailoverUrlOk() (*string, bool)`

GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEventFailoverUrl

`func (o *FqdnConnection) SetWebhookEventFailoverUrl(v string)`

SetWebhookEventFailoverUrl sets WebhookEventFailoverUrl field to given value.

### HasWebhookEventFailoverUrl

`func (o *FqdnConnection) HasWebhookEventFailoverUrl() bool`

HasWebhookEventFailoverUrl returns a boolean if a field has been set.

### SetWebhookEventFailoverUrlNil

`func (o *FqdnConnection) SetWebhookEventFailoverUrlNil(b bool)`

 SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil

### UnsetWebhookEventFailoverUrl
`func (o *FqdnConnection) UnsetWebhookEventFailoverUrl()`

UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
### GetWebhookApiVersion

`func (o *FqdnConnection) GetWebhookApiVersion() WebhookApiVersion`

GetWebhookApiVersion returns the WebhookApiVersion field if non-nil, zero value otherwise.

### GetWebhookApiVersionOk

`func (o *FqdnConnection) GetWebhookApiVersionOk() (*WebhookApiVersion, bool)`

GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookApiVersion

`func (o *FqdnConnection) SetWebhookApiVersion(v WebhookApiVersion)`

SetWebhookApiVersion sets WebhookApiVersion field to given value.

### HasWebhookApiVersion

`func (o *FqdnConnection) HasWebhookApiVersion() bool`

HasWebhookApiVersion returns a boolean if a field has been set.

### GetWebhookTimeoutSecs

`func (o *FqdnConnection) GetWebhookTimeoutSecs() int32`

GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field if non-nil, zero value otherwise.

### GetWebhookTimeoutSecsOk

`func (o *FqdnConnection) GetWebhookTimeoutSecsOk() (*int32, bool)`

GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookTimeoutSecs

`func (o *FqdnConnection) SetWebhookTimeoutSecs(v int32)`

SetWebhookTimeoutSecs sets WebhookTimeoutSecs field to given value.

### HasWebhookTimeoutSecs

`func (o *FqdnConnection) HasWebhookTimeoutSecs() bool`

HasWebhookTimeoutSecs returns a boolean if a field has been set.

### SetWebhookTimeoutSecsNil

`func (o *FqdnConnection) SetWebhookTimeoutSecsNil(b bool)`

 SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil

### UnsetWebhookTimeoutSecs
`func (o *FqdnConnection) UnsetWebhookTimeoutSecs()`

UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
### GetRtcpSettings

`func (o *FqdnConnection) GetRtcpSettings() ConnectionRtcpSettings`

GetRtcpSettings returns the RtcpSettings field if non-nil, zero value otherwise.

### GetRtcpSettingsOk

`func (o *FqdnConnection) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool)`

GetRtcpSettingsOk returns a tuple with the RtcpSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtcpSettings

`func (o *FqdnConnection) SetRtcpSettings(v ConnectionRtcpSettings)`

SetRtcpSettings sets RtcpSettings field to given value.

### HasRtcpSettings

`func (o *FqdnConnection) HasRtcpSettings() bool`

HasRtcpSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FqdnConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FqdnConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FqdnConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FqdnConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FqdnConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FqdnConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FqdnConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FqdnConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetInbound

`func (o *FqdnConnection) GetInbound() InboundFqdn`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *FqdnConnection) GetInboundOk() (*InboundFqdn, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *FqdnConnection) SetInbound(v InboundFqdn)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *FqdnConnection) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetOutbound

`func (o *FqdnConnection) GetOutbound() OutboundFqdn`

GetOutbound returns the Outbound field if non-nil, zero value otherwise.

### GetOutboundOk

`func (o *FqdnConnection) GetOutboundOk() (*OutboundFqdn, bool)`

GetOutboundOk returns a tuple with the Outbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbound

`func (o *FqdnConnection) SetOutbound(v OutboundFqdn)`

SetOutbound sets Outbound field to given value.

### HasOutbound

`func (o *FqdnConnection) HasOutbound() bool`

HasOutbound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


