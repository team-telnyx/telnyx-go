# InboundIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AniNumberFormat** | Pointer to **string** | This setting allows you to set the format with which the caller&#39;s number (ANI) is sent for inbound phone calls. | [optional] [default to "E.164-national"]
**DnisNumberFormat** | Pointer to **string** |  | [optional] [default to "e164"]
**Codecs** | Pointer to **[]string** | Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP. | [optional] [default to ["G722","G711U","G711A","G729","OPUS","H.264"]]
**DefaultPrimaryIpId** | Pointer to **string** | The default primary IP to use for the number. Only settable if the connection is               of IP authentication type. Value must be the ID of an authorized IP set on the connection. | [optional] 
**DefaultSecondaryIpId** | Pointer to **string** | The default secondary IP to use for the number. Only settable if the connection is               of IP authentication type. Value must be the ID of an authorized IP set on the connection. | [optional] 
**DefaultTertiaryIpId** | Pointer to **string** | The default tertiary IP to use for the number. Only settable if the connection is               of IP authentication type. Value must be the ID of an authorized IP set on the connection. | [optional] 
**DefaultRoutingMethod** | Pointer to **string** | Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or left blank, other values are not allowed. | [optional] 
**ChannelLimit** | Pointer to **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] 
**GenerateRingbackTone** | Pointer to **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**IsupHeadersEnabled** | Pointer to **bool** | When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.) | [optional] [default to false]
**PrackEnabled** | Pointer to **bool** | Enable PRACK messages as defined in RFC3262. | [optional] [default to false]
**SipCompactHeadersEnabled** | Pointer to **bool** | Defaults to true. | [optional] [default to true]
**SipRegion** | Pointer to **string** | Selects which &#x60;sip_region&#x60; to receive inbound calls from. If null, the default region (US) will be used. | [optional] [default to "US"]
**SipSubdomain** | Pointer to **string** | Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \&quot;example.sip.telnyx.com\&quot; can be called from any SIP endpoint by using the SIP URI \&quot;sip:@example.sip.telnyx.com\&quot; where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls. | [optional] 
**SipSubdomainReceiveSettings** | Pointer to **string** | This option can be enabled to receive calls from: \&quot;Anyone\&quot; (any SIP endpoint in the public Internet) or \&quot;Only my connections\&quot; (any connection assigned to the same Telnyx user). | [optional] 
**Timeout1xxSecs** | Pointer to **int32** | Time(sec) before aborting if connection is not made. | [optional] [default to 3]
**Timeout2xxSecs** | Pointer to **int32** | Time(sec) before aborting if call is unanswered (min: 1, max: 600). | [optional] [default to 90]
**ShakenStirEnabled** | Pointer to **bool** | When enabled the SIP Connection will receive the Identity header with Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP transport. | [optional] [default to false]

## Methods

### NewInboundIp

`func NewInboundIp() *InboundIp`

NewInboundIp instantiates a new InboundIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundIpWithDefaults

`func NewInboundIpWithDefaults() *InboundIp`

NewInboundIpWithDefaults instantiates a new InboundIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAniNumberFormat

`func (o *InboundIp) GetAniNumberFormat() string`

GetAniNumberFormat returns the AniNumberFormat field if non-nil, zero value otherwise.

### GetAniNumberFormatOk

`func (o *InboundIp) GetAniNumberFormatOk() (*string, bool)`

GetAniNumberFormatOk returns a tuple with the AniNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniNumberFormat

`func (o *InboundIp) SetAniNumberFormat(v string)`

SetAniNumberFormat sets AniNumberFormat field to given value.

### HasAniNumberFormat

`func (o *InboundIp) HasAniNumberFormat() bool`

HasAniNumberFormat returns a boolean if a field has been set.

### GetDnisNumberFormat

`func (o *InboundIp) GetDnisNumberFormat() string`

GetDnisNumberFormat returns the DnisNumberFormat field if non-nil, zero value otherwise.

### GetDnisNumberFormatOk

`func (o *InboundIp) GetDnisNumberFormatOk() (*string, bool)`

GetDnisNumberFormatOk returns a tuple with the DnisNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnisNumberFormat

`func (o *InboundIp) SetDnisNumberFormat(v string)`

SetDnisNumberFormat sets DnisNumberFormat field to given value.

### HasDnisNumberFormat

`func (o *InboundIp) HasDnisNumberFormat() bool`

HasDnisNumberFormat returns a boolean if a field has been set.

### GetCodecs

`func (o *InboundIp) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *InboundIp) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *InboundIp) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *InboundIp) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### GetDefaultPrimaryIpId

`func (o *InboundIp) GetDefaultPrimaryIpId() string`

GetDefaultPrimaryIpId returns the DefaultPrimaryIpId field if non-nil, zero value otherwise.

### GetDefaultPrimaryIpIdOk

`func (o *InboundIp) GetDefaultPrimaryIpIdOk() (*string, bool)`

GetDefaultPrimaryIpIdOk returns a tuple with the DefaultPrimaryIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrimaryIpId

`func (o *InboundIp) SetDefaultPrimaryIpId(v string)`

SetDefaultPrimaryIpId sets DefaultPrimaryIpId field to given value.

### HasDefaultPrimaryIpId

`func (o *InboundIp) HasDefaultPrimaryIpId() bool`

HasDefaultPrimaryIpId returns a boolean if a field has been set.

### GetDefaultSecondaryIpId

`func (o *InboundIp) GetDefaultSecondaryIpId() string`

GetDefaultSecondaryIpId returns the DefaultSecondaryIpId field if non-nil, zero value otherwise.

### GetDefaultSecondaryIpIdOk

`func (o *InboundIp) GetDefaultSecondaryIpIdOk() (*string, bool)`

GetDefaultSecondaryIpIdOk returns a tuple with the DefaultSecondaryIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecondaryIpId

`func (o *InboundIp) SetDefaultSecondaryIpId(v string)`

SetDefaultSecondaryIpId sets DefaultSecondaryIpId field to given value.

### HasDefaultSecondaryIpId

`func (o *InboundIp) HasDefaultSecondaryIpId() bool`

HasDefaultSecondaryIpId returns a boolean if a field has been set.

### GetDefaultTertiaryIpId

`func (o *InboundIp) GetDefaultTertiaryIpId() string`

GetDefaultTertiaryIpId returns the DefaultTertiaryIpId field if non-nil, zero value otherwise.

### GetDefaultTertiaryIpIdOk

`func (o *InboundIp) GetDefaultTertiaryIpIdOk() (*string, bool)`

GetDefaultTertiaryIpIdOk returns a tuple with the DefaultTertiaryIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTertiaryIpId

`func (o *InboundIp) SetDefaultTertiaryIpId(v string)`

SetDefaultTertiaryIpId sets DefaultTertiaryIpId field to given value.

### HasDefaultTertiaryIpId

`func (o *InboundIp) HasDefaultTertiaryIpId() bool`

HasDefaultTertiaryIpId returns a boolean if a field has been set.

### GetDefaultRoutingMethod

`func (o *InboundIp) GetDefaultRoutingMethod() string`

GetDefaultRoutingMethod returns the DefaultRoutingMethod field if non-nil, zero value otherwise.

### GetDefaultRoutingMethodOk

`func (o *InboundIp) GetDefaultRoutingMethodOk() (*string, bool)`

GetDefaultRoutingMethodOk returns a tuple with the DefaultRoutingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoutingMethod

`func (o *InboundIp) SetDefaultRoutingMethod(v string)`

SetDefaultRoutingMethod sets DefaultRoutingMethod field to given value.

### HasDefaultRoutingMethod

`func (o *InboundIp) HasDefaultRoutingMethod() bool`

HasDefaultRoutingMethod returns a boolean if a field has been set.

### GetChannelLimit

`func (o *InboundIp) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *InboundIp) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *InboundIp) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *InboundIp) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetGenerateRingbackTone

`func (o *InboundIp) GetGenerateRingbackTone() bool`

GetGenerateRingbackTone returns the GenerateRingbackTone field if non-nil, zero value otherwise.

### GetGenerateRingbackToneOk

`func (o *InboundIp) GetGenerateRingbackToneOk() (*bool, bool)`

GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRingbackTone

`func (o *InboundIp) SetGenerateRingbackTone(v bool)`

SetGenerateRingbackTone sets GenerateRingbackTone field to given value.

### HasGenerateRingbackTone

`func (o *InboundIp) HasGenerateRingbackTone() bool`

HasGenerateRingbackTone returns a boolean if a field has been set.

### GetIsupHeadersEnabled

`func (o *InboundIp) GetIsupHeadersEnabled() bool`

GetIsupHeadersEnabled returns the IsupHeadersEnabled field if non-nil, zero value otherwise.

### GetIsupHeadersEnabledOk

`func (o *InboundIp) GetIsupHeadersEnabledOk() (*bool, bool)`

GetIsupHeadersEnabledOk returns a tuple with the IsupHeadersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsupHeadersEnabled

`func (o *InboundIp) SetIsupHeadersEnabled(v bool)`

SetIsupHeadersEnabled sets IsupHeadersEnabled field to given value.

### HasIsupHeadersEnabled

`func (o *InboundIp) HasIsupHeadersEnabled() bool`

HasIsupHeadersEnabled returns a boolean if a field has been set.

### GetPrackEnabled

`func (o *InboundIp) GetPrackEnabled() bool`

GetPrackEnabled returns the PrackEnabled field if non-nil, zero value otherwise.

### GetPrackEnabledOk

`func (o *InboundIp) GetPrackEnabledOk() (*bool, bool)`

GetPrackEnabledOk returns a tuple with the PrackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrackEnabled

`func (o *InboundIp) SetPrackEnabled(v bool)`

SetPrackEnabled sets PrackEnabled field to given value.

### HasPrackEnabled

`func (o *InboundIp) HasPrackEnabled() bool`

HasPrackEnabled returns a boolean if a field has been set.

### GetSipCompactHeadersEnabled

`func (o *InboundIp) GetSipCompactHeadersEnabled() bool`

GetSipCompactHeadersEnabled returns the SipCompactHeadersEnabled field if non-nil, zero value otherwise.

### GetSipCompactHeadersEnabledOk

`func (o *InboundIp) GetSipCompactHeadersEnabledOk() (*bool, bool)`

GetSipCompactHeadersEnabledOk returns a tuple with the SipCompactHeadersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipCompactHeadersEnabled

`func (o *InboundIp) SetSipCompactHeadersEnabled(v bool)`

SetSipCompactHeadersEnabled sets SipCompactHeadersEnabled field to given value.

### HasSipCompactHeadersEnabled

`func (o *InboundIp) HasSipCompactHeadersEnabled() bool`

HasSipCompactHeadersEnabled returns a boolean if a field has been set.

### GetSipRegion

`func (o *InboundIp) GetSipRegion() string`

GetSipRegion returns the SipRegion field if non-nil, zero value otherwise.

### GetSipRegionOk

`func (o *InboundIp) GetSipRegionOk() (*string, bool)`

GetSipRegionOk returns a tuple with the SipRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipRegion

`func (o *InboundIp) SetSipRegion(v string)`

SetSipRegion sets SipRegion field to given value.

### HasSipRegion

`func (o *InboundIp) HasSipRegion() bool`

HasSipRegion returns a boolean if a field has been set.

### GetSipSubdomain

`func (o *InboundIp) GetSipSubdomain() string`

GetSipSubdomain returns the SipSubdomain field if non-nil, zero value otherwise.

### GetSipSubdomainOk

`func (o *InboundIp) GetSipSubdomainOk() (*string, bool)`

GetSipSubdomainOk returns a tuple with the SipSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomain

`func (o *InboundIp) SetSipSubdomain(v string)`

SetSipSubdomain sets SipSubdomain field to given value.

### HasSipSubdomain

`func (o *InboundIp) HasSipSubdomain() bool`

HasSipSubdomain returns a boolean if a field has been set.

### GetSipSubdomainReceiveSettings

`func (o *InboundIp) GetSipSubdomainReceiveSettings() string`

GetSipSubdomainReceiveSettings returns the SipSubdomainReceiveSettings field if non-nil, zero value otherwise.

### GetSipSubdomainReceiveSettingsOk

`func (o *InboundIp) GetSipSubdomainReceiveSettingsOk() (*string, bool)`

GetSipSubdomainReceiveSettingsOk returns a tuple with the SipSubdomainReceiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomainReceiveSettings

`func (o *InboundIp) SetSipSubdomainReceiveSettings(v string)`

SetSipSubdomainReceiveSettings sets SipSubdomainReceiveSettings field to given value.

### HasSipSubdomainReceiveSettings

`func (o *InboundIp) HasSipSubdomainReceiveSettings() bool`

HasSipSubdomainReceiveSettings returns a boolean if a field has been set.

### GetTimeout1xxSecs

`func (o *InboundIp) GetTimeout1xxSecs() int32`

GetTimeout1xxSecs returns the Timeout1xxSecs field if non-nil, zero value otherwise.

### GetTimeout1xxSecsOk

`func (o *InboundIp) GetTimeout1xxSecsOk() (*int32, bool)`

GetTimeout1xxSecsOk returns a tuple with the Timeout1xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout1xxSecs

`func (o *InboundIp) SetTimeout1xxSecs(v int32)`

SetTimeout1xxSecs sets Timeout1xxSecs field to given value.

### HasTimeout1xxSecs

`func (o *InboundIp) HasTimeout1xxSecs() bool`

HasTimeout1xxSecs returns a boolean if a field has been set.

### GetTimeout2xxSecs

`func (o *InboundIp) GetTimeout2xxSecs() int32`

GetTimeout2xxSecs returns the Timeout2xxSecs field if non-nil, zero value otherwise.

### GetTimeout2xxSecsOk

`func (o *InboundIp) GetTimeout2xxSecsOk() (*int32, bool)`

GetTimeout2xxSecsOk returns a tuple with the Timeout2xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout2xxSecs

`func (o *InboundIp) SetTimeout2xxSecs(v int32)`

SetTimeout2xxSecs sets Timeout2xxSecs field to given value.

### HasTimeout2xxSecs

`func (o *InboundIp) HasTimeout2xxSecs() bool`

HasTimeout2xxSecs returns a boolean if a field has been set.

### GetShakenStirEnabled

`func (o *InboundIp) GetShakenStirEnabled() bool`

GetShakenStirEnabled returns the ShakenStirEnabled field if non-nil, zero value otherwise.

### GetShakenStirEnabledOk

`func (o *InboundIp) GetShakenStirEnabledOk() (*bool, bool)`

GetShakenStirEnabledOk returns a tuple with the ShakenStirEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShakenStirEnabled

`func (o *InboundIp) SetShakenStirEnabled(v bool)`

SetShakenStirEnabled sets ShakenStirEnabled field to given value.

### HasShakenStirEnabled

`func (o *InboundIp) HasShakenStirEnabled() bool`

HasShakenStirEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


