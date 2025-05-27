# InboundFqdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AniNumberFormat** | Pointer to **string** | This setting allows you to set the format with which the caller&#39;s number (ANI) is sent for inbound phone calls. | [optional] [default to "E.164-national"]
**DnisNumberFormat** | Pointer to **string** |  | [optional] [default to "e164"]
**Codecs** | Pointer to **[]string** | Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP. | [optional] [default to ["G722","G711U","G711A","G729","OPUS","H.264"]]
**DefaultRoutingMethod** | Pointer to **NullableString** | Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or null, other values are not allowed. | [optional] [default to "sequential"]
**DefaultPrimaryFqdnId** | Pointer to **NullableString** | The default primary FQDN to use for the number. Only settable if the connection is of FQDN type. Value must be the ID of an FQDN set on the connection. | [optional] 
**DefaultSecondaryFqdnId** | Pointer to **NullableString** | The default secondary FQDN to use for the number. Only settable if the connection is of FQDN type. Value must be the ID of an FQDN set on the connection. | [optional] 
**DefaultTertiaryFqdnId** | Pointer to **NullableString** | The default tertiary FQDN to use for the number. Only settable if the connection is of FQDN type. Value must be the ID of an FQDN set on the connection. | [optional] 
**ChannelLimit** | Pointer to **NullableInt32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] 
**GenerateRingbackTone** | Pointer to **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**IsupHeadersEnabled** | Pointer to **bool** | When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.) | [optional] [default to false]
**PrackEnabled** | Pointer to **bool** | Enable PRACK messages as defined in RFC3262. | [optional] [default to false]
**SipCompactHeadersEnabled** | Pointer to **bool** | Defaults to true. | [optional] [default to true]
**SipRegion** | Pointer to **string** | Selects which &#x60;sip_region&#x60; to receive inbound calls from. If null, the default region (US) will be used. | [optional] [default to "US"]
**SipSubdomain** | Pointer to **NullableString** | Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \&quot;example.sip.telnyx.com\&quot; can be called from any SIP endpoint by using the SIP URI \&quot;sip:@example.sip.telnyx.com\&quot; where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls. | [optional] 
**SipSubdomainReceiveSettings** | Pointer to **string** | This option can be enabled to receive calls from: \&quot;Anyone\&quot; (any SIP endpoint in the public Internet) or \&quot;Only my connections\&quot; (any connection assigned to the same Telnyx user). | [optional] [default to "from_anyone"]
**Timeout1xxSecs** | Pointer to **int32** | Time(sec) before aborting if connection is not made. | [optional] [default to 3]
**Timeout2xxSecs** | Pointer to **int32** | Time(sec) before aborting if call is unanswered (min: 1, max: 600). | [optional] [default to 90]
**ShakenStirEnabled** | Pointer to **bool** | When enabled the SIP Connection will receive the Identity header with Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP transport. | [optional] [default to false]

## Methods

### NewInboundFqdn

`func NewInboundFqdn() *InboundFqdn`

NewInboundFqdn instantiates a new InboundFqdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundFqdnWithDefaults

`func NewInboundFqdnWithDefaults() *InboundFqdn`

NewInboundFqdnWithDefaults instantiates a new InboundFqdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAniNumberFormat

`func (o *InboundFqdn) GetAniNumberFormat() string`

GetAniNumberFormat returns the AniNumberFormat field if non-nil, zero value otherwise.

### GetAniNumberFormatOk

`func (o *InboundFqdn) GetAniNumberFormatOk() (*string, bool)`

GetAniNumberFormatOk returns a tuple with the AniNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniNumberFormat

`func (o *InboundFqdn) SetAniNumberFormat(v string)`

SetAniNumberFormat sets AniNumberFormat field to given value.

### HasAniNumberFormat

`func (o *InboundFqdn) HasAniNumberFormat() bool`

HasAniNumberFormat returns a boolean if a field has been set.

### GetDnisNumberFormat

`func (o *InboundFqdn) GetDnisNumberFormat() string`

GetDnisNumberFormat returns the DnisNumberFormat field if non-nil, zero value otherwise.

### GetDnisNumberFormatOk

`func (o *InboundFqdn) GetDnisNumberFormatOk() (*string, bool)`

GetDnisNumberFormatOk returns a tuple with the DnisNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnisNumberFormat

`func (o *InboundFqdn) SetDnisNumberFormat(v string)`

SetDnisNumberFormat sets DnisNumberFormat field to given value.

### HasDnisNumberFormat

`func (o *InboundFqdn) HasDnisNumberFormat() bool`

HasDnisNumberFormat returns a boolean if a field has been set.

### GetCodecs

`func (o *InboundFqdn) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *InboundFqdn) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *InboundFqdn) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *InboundFqdn) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### GetDefaultRoutingMethod

`func (o *InboundFqdn) GetDefaultRoutingMethod() string`

GetDefaultRoutingMethod returns the DefaultRoutingMethod field if non-nil, zero value otherwise.

### GetDefaultRoutingMethodOk

`func (o *InboundFqdn) GetDefaultRoutingMethodOk() (*string, bool)`

GetDefaultRoutingMethodOk returns a tuple with the DefaultRoutingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoutingMethod

`func (o *InboundFqdn) SetDefaultRoutingMethod(v string)`

SetDefaultRoutingMethod sets DefaultRoutingMethod field to given value.

### HasDefaultRoutingMethod

`func (o *InboundFqdn) HasDefaultRoutingMethod() bool`

HasDefaultRoutingMethod returns a boolean if a field has been set.

### SetDefaultRoutingMethodNil

`func (o *InboundFqdn) SetDefaultRoutingMethodNil(b bool)`

 SetDefaultRoutingMethodNil sets the value for DefaultRoutingMethod to be an explicit nil

### UnsetDefaultRoutingMethod
`func (o *InboundFqdn) UnsetDefaultRoutingMethod()`

UnsetDefaultRoutingMethod ensures that no value is present for DefaultRoutingMethod, not even an explicit nil
### GetDefaultPrimaryFqdnId

`func (o *InboundFqdn) GetDefaultPrimaryFqdnId() string`

GetDefaultPrimaryFqdnId returns the DefaultPrimaryFqdnId field if non-nil, zero value otherwise.

### GetDefaultPrimaryFqdnIdOk

`func (o *InboundFqdn) GetDefaultPrimaryFqdnIdOk() (*string, bool)`

GetDefaultPrimaryFqdnIdOk returns a tuple with the DefaultPrimaryFqdnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrimaryFqdnId

`func (o *InboundFqdn) SetDefaultPrimaryFqdnId(v string)`

SetDefaultPrimaryFqdnId sets DefaultPrimaryFqdnId field to given value.

### HasDefaultPrimaryFqdnId

`func (o *InboundFqdn) HasDefaultPrimaryFqdnId() bool`

HasDefaultPrimaryFqdnId returns a boolean if a field has been set.

### SetDefaultPrimaryFqdnIdNil

`func (o *InboundFqdn) SetDefaultPrimaryFqdnIdNil(b bool)`

 SetDefaultPrimaryFqdnIdNil sets the value for DefaultPrimaryFqdnId to be an explicit nil

### UnsetDefaultPrimaryFqdnId
`func (o *InboundFqdn) UnsetDefaultPrimaryFqdnId()`

UnsetDefaultPrimaryFqdnId ensures that no value is present for DefaultPrimaryFqdnId, not even an explicit nil
### GetDefaultSecondaryFqdnId

`func (o *InboundFqdn) GetDefaultSecondaryFqdnId() string`

GetDefaultSecondaryFqdnId returns the DefaultSecondaryFqdnId field if non-nil, zero value otherwise.

### GetDefaultSecondaryFqdnIdOk

`func (o *InboundFqdn) GetDefaultSecondaryFqdnIdOk() (*string, bool)`

GetDefaultSecondaryFqdnIdOk returns a tuple with the DefaultSecondaryFqdnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecondaryFqdnId

`func (o *InboundFqdn) SetDefaultSecondaryFqdnId(v string)`

SetDefaultSecondaryFqdnId sets DefaultSecondaryFqdnId field to given value.

### HasDefaultSecondaryFqdnId

`func (o *InboundFqdn) HasDefaultSecondaryFqdnId() bool`

HasDefaultSecondaryFqdnId returns a boolean if a field has been set.

### SetDefaultSecondaryFqdnIdNil

`func (o *InboundFqdn) SetDefaultSecondaryFqdnIdNil(b bool)`

 SetDefaultSecondaryFqdnIdNil sets the value for DefaultSecondaryFqdnId to be an explicit nil

### UnsetDefaultSecondaryFqdnId
`func (o *InboundFqdn) UnsetDefaultSecondaryFqdnId()`

UnsetDefaultSecondaryFqdnId ensures that no value is present for DefaultSecondaryFqdnId, not even an explicit nil
### GetDefaultTertiaryFqdnId

`func (o *InboundFqdn) GetDefaultTertiaryFqdnId() string`

GetDefaultTertiaryFqdnId returns the DefaultTertiaryFqdnId field if non-nil, zero value otherwise.

### GetDefaultTertiaryFqdnIdOk

`func (o *InboundFqdn) GetDefaultTertiaryFqdnIdOk() (*string, bool)`

GetDefaultTertiaryFqdnIdOk returns a tuple with the DefaultTertiaryFqdnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTertiaryFqdnId

`func (o *InboundFqdn) SetDefaultTertiaryFqdnId(v string)`

SetDefaultTertiaryFqdnId sets DefaultTertiaryFqdnId field to given value.

### HasDefaultTertiaryFqdnId

`func (o *InboundFqdn) HasDefaultTertiaryFqdnId() bool`

HasDefaultTertiaryFqdnId returns a boolean if a field has been set.

### SetDefaultTertiaryFqdnIdNil

`func (o *InboundFqdn) SetDefaultTertiaryFqdnIdNil(b bool)`

 SetDefaultTertiaryFqdnIdNil sets the value for DefaultTertiaryFqdnId to be an explicit nil

### UnsetDefaultTertiaryFqdnId
`func (o *InboundFqdn) UnsetDefaultTertiaryFqdnId()`

UnsetDefaultTertiaryFqdnId ensures that no value is present for DefaultTertiaryFqdnId, not even an explicit nil
### GetChannelLimit

`func (o *InboundFqdn) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *InboundFqdn) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *InboundFqdn) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *InboundFqdn) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### SetChannelLimitNil

`func (o *InboundFqdn) SetChannelLimitNil(b bool)`

 SetChannelLimitNil sets the value for ChannelLimit to be an explicit nil

### UnsetChannelLimit
`func (o *InboundFqdn) UnsetChannelLimit()`

UnsetChannelLimit ensures that no value is present for ChannelLimit, not even an explicit nil
### GetGenerateRingbackTone

`func (o *InboundFqdn) GetGenerateRingbackTone() bool`

GetGenerateRingbackTone returns the GenerateRingbackTone field if non-nil, zero value otherwise.

### GetGenerateRingbackToneOk

`func (o *InboundFqdn) GetGenerateRingbackToneOk() (*bool, bool)`

GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRingbackTone

`func (o *InboundFqdn) SetGenerateRingbackTone(v bool)`

SetGenerateRingbackTone sets GenerateRingbackTone field to given value.

### HasGenerateRingbackTone

`func (o *InboundFqdn) HasGenerateRingbackTone() bool`

HasGenerateRingbackTone returns a boolean if a field has been set.

### GetIsupHeadersEnabled

`func (o *InboundFqdn) GetIsupHeadersEnabled() bool`

GetIsupHeadersEnabled returns the IsupHeadersEnabled field if non-nil, zero value otherwise.

### GetIsupHeadersEnabledOk

`func (o *InboundFqdn) GetIsupHeadersEnabledOk() (*bool, bool)`

GetIsupHeadersEnabledOk returns a tuple with the IsupHeadersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsupHeadersEnabled

`func (o *InboundFqdn) SetIsupHeadersEnabled(v bool)`

SetIsupHeadersEnabled sets IsupHeadersEnabled field to given value.

### HasIsupHeadersEnabled

`func (o *InboundFqdn) HasIsupHeadersEnabled() bool`

HasIsupHeadersEnabled returns a boolean if a field has been set.

### GetPrackEnabled

`func (o *InboundFqdn) GetPrackEnabled() bool`

GetPrackEnabled returns the PrackEnabled field if non-nil, zero value otherwise.

### GetPrackEnabledOk

`func (o *InboundFqdn) GetPrackEnabledOk() (*bool, bool)`

GetPrackEnabledOk returns a tuple with the PrackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrackEnabled

`func (o *InboundFqdn) SetPrackEnabled(v bool)`

SetPrackEnabled sets PrackEnabled field to given value.

### HasPrackEnabled

`func (o *InboundFqdn) HasPrackEnabled() bool`

HasPrackEnabled returns a boolean if a field has been set.

### GetSipCompactHeadersEnabled

`func (o *InboundFqdn) GetSipCompactHeadersEnabled() bool`

GetSipCompactHeadersEnabled returns the SipCompactHeadersEnabled field if non-nil, zero value otherwise.

### GetSipCompactHeadersEnabledOk

`func (o *InboundFqdn) GetSipCompactHeadersEnabledOk() (*bool, bool)`

GetSipCompactHeadersEnabledOk returns a tuple with the SipCompactHeadersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipCompactHeadersEnabled

`func (o *InboundFqdn) SetSipCompactHeadersEnabled(v bool)`

SetSipCompactHeadersEnabled sets SipCompactHeadersEnabled field to given value.

### HasSipCompactHeadersEnabled

`func (o *InboundFqdn) HasSipCompactHeadersEnabled() bool`

HasSipCompactHeadersEnabled returns a boolean if a field has been set.

### GetSipRegion

`func (o *InboundFqdn) GetSipRegion() string`

GetSipRegion returns the SipRegion field if non-nil, zero value otherwise.

### GetSipRegionOk

`func (o *InboundFqdn) GetSipRegionOk() (*string, bool)`

GetSipRegionOk returns a tuple with the SipRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipRegion

`func (o *InboundFqdn) SetSipRegion(v string)`

SetSipRegion sets SipRegion field to given value.

### HasSipRegion

`func (o *InboundFqdn) HasSipRegion() bool`

HasSipRegion returns a boolean if a field has been set.

### GetSipSubdomain

`func (o *InboundFqdn) GetSipSubdomain() string`

GetSipSubdomain returns the SipSubdomain field if non-nil, zero value otherwise.

### GetSipSubdomainOk

`func (o *InboundFqdn) GetSipSubdomainOk() (*string, bool)`

GetSipSubdomainOk returns a tuple with the SipSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomain

`func (o *InboundFqdn) SetSipSubdomain(v string)`

SetSipSubdomain sets SipSubdomain field to given value.

### HasSipSubdomain

`func (o *InboundFqdn) HasSipSubdomain() bool`

HasSipSubdomain returns a boolean if a field has been set.

### SetSipSubdomainNil

`func (o *InboundFqdn) SetSipSubdomainNil(b bool)`

 SetSipSubdomainNil sets the value for SipSubdomain to be an explicit nil

### UnsetSipSubdomain
`func (o *InboundFqdn) UnsetSipSubdomain()`

UnsetSipSubdomain ensures that no value is present for SipSubdomain, not even an explicit nil
### GetSipSubdomainReceiveSettings

`func (o *InboundFqdn) GetSipSubdomainReceiveSettings() string`

GetSipSubdomainReceiveSettings returns the SipSubdomainReceiveSettings field if non-nil, zero value otherwise.

### GetSipSubdomainReceiveSettingsOk

`func (o *InboundFqdn) GetSipSubdomainReceiveSettingsOk() (*string, bool)`

GetSipSubdomainReceiveSettingsOk returns a tuple with the SipSubdomainReceiveSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSubdomainReceiveSettings

`func (o *InboundFqdn) SetSipSubdomainReceiveSettings(v string)`

SetSipSubdomainReceiveSettings sets SipSubdomainReceiveSettings field to given value.

### HasSipSubdomainReceiveSettings

`func (o *InboundFqdn) HasSipSubdomainReceiveSettings() bool`

HasSipSubdomainReceiveSettings returns a boolean if a field has been set.

### GetTimeout1xxSecs

`func (o *InboundFqdn) GetTimeout1xxSecs() int32`

GetTimeout1xxSecs returns the Timeout1xxSecs field if non-nil, zero value otherwise.

### GetTimeout1xxSecsOk

`func (o *InboundFqdn) GetTimeout1xxSecsOk() (*int32, bool)`

GetTimeout1xxSecsOk returns a tuple with the Timeout1xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout1xxSecs

`func (o *InboundFqdn) SetTimeout1xxSecs(v int32)`

SetTimeout1xxSecs sets Timeout1xxSecs field to given value.

### HasTimeout1xxSecs

`func (o *InboundFqdn) HasTimeout1xxSecs() bool`

HasTimeout1xxSecs returns a boolean if a field has been set.

### GetTimeout2xxSecs

`func (o *InboundFqdn) GetTimeout2xxSecs() int32`

GetTimeout2xxSecs returns the Timeout2xxSecs field if non-nil, zero value otherwise.

### GetTimeout2xxSecsOk

`func (o *InboundFqdn) GetTimeout2xxSecsOk() (*int32, bool)`

GetTimeout2xxSecsOk returns a tuple with the Timeout2xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout2xxSecs

`func (o *InboundFqdn) SetTimeout2xxSecs(v int32)`

SetTimeout2xxSecs sets Timeout2xxSecs field to given value.

### HasTimeout2xxSecs

`func (o *InboundFqdn) HasTimeout2xxSecs() bool`

HasTimeout2xxSecs returns a boolean if a field has been set.

### GetShakenStirEnabled

`func (o *InboundFqdn) GetShakenStirEnabled() bool`

GetShakenStirEnabled returns the ShakenStirEnabled field if non-nil, zero value otherwise.

### GetShakenStirEnabledOk

`func (o *InboundFqdn) GetShakenStirEnabledOk() (*bool, bool)`

GetShakenStirEnabledOk returns a tuple with the ShakenStirEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShakenStirEnabled

`func (o *InboundFqdn) SetShakenStirEnabled(v bool)`

SetShakenStirEnabled sets ShakenStirEnabled field to given value.

### HasShakenStirEnabled

`func (o *InboundFqdn) HasShakenStirEnabled() bool`

HasShakenStirEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


