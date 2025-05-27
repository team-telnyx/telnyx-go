# CredentialInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AniNumberFormat** | Pointer to **string** | This setting allows you to set the format with which the caller&#39;s number (ANI) is sent for inbound phone calls. | [optional] [default to "E.164-national"]
**DnisNumberFormat** | Pointer to **string** |  | [optional] [default to "e164"]
**Codecs** | Pointer to **[]string** | Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP. | [optional] [default to ["G722","G711U","G711A","G729","OPUS","H.264"]]
**ChannelLimit** | Pointer to **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] 
**GenerateRingbackTone** | Pointer to **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**IsupHeadersEnabled** | Pointer to **bool** | When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.) | [optional] [default to false]
**PrackEnabled** | Pointer to **bool** | Enable PRACK messages as defined in RFC3262. | [optional] [default to false]
**SipCompactHeadersEnabled** | Pointer to **bool** | Defaults to true. | [optional] [default to true]
**Timeout1xxSecs** | Pointer to **int32** | Time(sec) before aborting if connection is not made. | [optional] [default to 3]
**Timeout2xxSecs** | Pointer to **int32** | Time(sec) before aborting if call is unanswered (min: 1, max: 600). | [optional] [default to 90]
**ShakenStirEnabled** | Pointer to **bool** | When enabled the SIP Connection will receive the Identity header with Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP transport. | [optional] [default to false]

## Methods

### NewCredentialInbound

`func NewCredentialInbound() *CredentialInbound`

NewCredentialInbound instantiates a new CredentialInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialInboundWithDefaults

`func NewCredentialInboundWithDefaults() *CredentialInbound`

NewCredentialInboundWithDefaults instantiates a new CredentialInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAniNumberFormat

`func (o *CredentialInbound) GetAniNumberFormat() string`

GetAniNumberFormat returns the AniNumberFormat field if non-nil, zero value otherwise.

### GetAniNumberFormatOk

`func (o *CredentialInbound) GetAniNumberFormatOk() (*string, bool)`

GetAniNumberFormatOk returns a tuple with the AniNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniNumberFormat

`func (o *CredentialInbound) SetAniNumberFormat(v string)`

SetAniNumberFormat sets AniNumberFormat field to given value.

### HasAniNumberFormat

`func (o *CredentialInbound) HasAniNumberFormat() bool`

HasAniNumberFormat returns a boolean if a field has been set.

### GetDnisNumberFormat

`func (o *CredentialInbound) GetDnisNumberFormat() string`

GetDnisNumberFormat returns the DnisNumberFormat field if non-nil, zero value otherwise.

### GetDnisNumberFormatOk

`func (o *CredentialInbound) GetDnisNumberFormatOk() (*string, bool)`

GetDnisNumberFormatOk returns a tuple with the DnisNumberFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnisNumberFormat

`func (o *CredentialInbound) SetDnisNumberFormat(v string)`

SetDnisNumberFormat sets DnisNumberFormat field to given value.

### HasDnisNumberFormat

`func (o *CredentialInbound) HasDnisNumberFormat() bool`

HasDnisNumberFormat returns a boolean if a field has been set.

### GetCodecs

`func (o *CredentialInbound) GetCodecs() []string`

GetCodecs returns the Codecs field if non-nil, zero value otherwise.

### GetCodecsOk

`func (o *CredentialInbound) GetCodecsOk() (*[]string, bool)`

GetCodecsOk returns a tuple with the Codecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecs

`func (o *CredentialInbound) SetCodecs(v []string)`

SetCodecs sets Codecs field to given value.

### HasCodecs

`func (o *CredentialInbound) HasCodecs() bool`

HasCodecs returns a boolean if a field has been set.

### GetChannelLimit

`func (o *CredentialInbound) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *CredentialInbound) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *CredentialInbound) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *CredentialInbound) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetGenerateRingbackTone

`func (o *CredentialInbound) GetGenerateRingbackTone() bool`

GetGenerateRingbackTone returns the GenerateRingbackTone field if non-nil, zero value otherwise.

### GetGenerateRingbackToneOk

`func (o *CredentialInbound) GetGenerateRingbackToneOk() (*bool, bool)`

GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRingbackTone

`func (o *CredentialInbound) SetGenerateRingbackTone(v bool)`

SetGenerateRingbackTone sets GenerateRingbackTone field to given value.

### HasGenerateRingbackTone

`func (o *CredentialInbound) HasGenerateRingbackTone() bool`

HasGenerateRingbackTone returns a boolean if a field has been set.

### GetIsupHeadersEnabled

`func (o *CredentialInbound) GetIsupHeadersEnabled() bool`

GetIsupHeadersEnabled returns the IsupHeadersEnabled field if non-nil, zero value otherwise.

### GetIsupHeadersEnabledOk

`func (o *CredentialInbound) GetIsupHeadersEnabledOk() (*bool, bool)`

GetIsupHeadersEnabledOk returns a tuple with the IsupHeadersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsupHeadersEnabled

`func (o *CredentialInbound) SetIsupHeadersEnabled(v bool)`

SetIsupHeadersEnabled sets IsupHeadersEnabled field to given value.

### HasIsupHeadersEnabled

`func (o *CredentialInbound) HasIsupHeadersEnabled() bool`

HasIsupHeadersEnabled returns a boolean if a field has been set.

### GetPrackEnabled

`func (o *CredentialInbound) GetPrackEnabled() bool`

GetPrackEnabled returns the PrackEnabled field if non-nil, zero value otherwise.

### GetPrackEnabledOk

`func (o *CredentialInbound) GetPrackEnabledOk() (*bool, bool)`

GetPrackEnabledOk returns a tuple with the PrackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrackEnabled

`func (o *CredentialInbound) SetPrackEnabled(v bool)`

SetPrackEnabled sets PrackEnabled field to given value.

### HasPrackEnabled

`func (o *CredentialInbound) HasPrackEnabled() bool`

HasPrackEnabled returns a boolean if a field has been set.

### GetSipCompactHeadersEnabled

`func (o *CredentialInbound) GetSipCompactHeadersEnabled() bool`

GetSipCompactHeadersEnabled returns the SipCompactHeadersEnabled field if non-nil, zero value otherwise.

### GetSipCompactHeadersEnabledOk

`func (o *CredentialInbound) GetSipCompactHeadersEnabledOk() (*bool, bool)`

GetSipCompactHeadersEnabledOk returns a tuple with the SipCompactHeadersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipCompactHeadersEnabled

`func (o *CredentialInbound) SetSipCompactHeadersEnabled(v bool)`

SetSipCompactHeadersEnabled sets SipCompactHeadersEnabled field to given value.

### HasSipCompactHeadersEnabled

`func (o *CredentialInbound) HasSipCompactHeadersEnabled() bool`

HasSipCompactHeadersEnabled returns a boolean if a field has been set.

### GetTimeout1xxSecs

`func (o *CredentialInbound) GetTimeout1xxSecs() int32`

GetTimeout1xxSecs returns the Timeout1xxSecs field if non-nil, zero value otherwise.

### GetTimeout1xxSecsOk

`func (o *CredentialInbound) GetTimeout1xxSecsOk() (*int32, bool)`

GetTimeout1xxSecsOk returns a tuple with the Timeout1xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout1xxSecs

`func (o *CredentialInbound) SetTimeout1xxSecs(v int32)`

SetTimeout1xxSecs sets Timeout1xxSecs field to given value.

### HasTimeout1xxSecs

`func (o *CredentialInbound) HasTimeout1xxSecs() bool`

HasTimeout1xxSecs returns a boolean if a field has been set.

### GetTimeout2xxSecs

`func (o *CredentialInbound) GetTimeout2xxSecs() int32`

GetTimeout2xxSecs returns the Timeout2xxSecs field if non-nil, zero value otherwise.

### GetTimeout2xxSecsOk

`func (o *CredentialInbound) GetTimeout2xxSecsOk() (*int32, bool)`

GetTimeout2xxSecsOk returns a tuple with the Timeout2xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout2xxSecs

`func (o *CredentialInbound) SetTimeout2xxSecs(v int32)`

SetTimeout2xxSecs sets Timeout2xxSecs field to given value.

### HasTimeout2xxSecs

`func (o *CredentialInbound) HasTimeout2xxSecs() bool`

HasTimeout2xxSecs returns a boolean if a field has been set.

### GetShakenStirEnabled

`func (o *CredentialInbound) GetShakenStirEnabled() bool`

GetShakenStirEnabled returns the ShakenStirEnabled field if non-nil, zero value otherwise.

### GetShakenStirEnabledOk

`func (o *CredentialInbound) GetShakenStirEnabledOk() (*bool, bool)`

GetShakenStirEnabledOk returns a tuple with the ShakenStirEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShakenStirEnabled

`func (o *CredentialInbound) SetShakenStirEnabled(v bool)`

SetShakenStirEnabled sets ShakenStirEnabled field to given value.

### HasShakenStirEnabled

`func (o *CredentialInbound) HasShakenStirEnabled() bool`

HasShakenStirEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


