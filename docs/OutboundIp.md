# OutboundIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallParkingEnabled** | Pointer to **NullableBool** | Forces all SIP calls originated on this connection to be \&quot;parked\&quot; instead of \&quot;bridged\&quot; to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next. | [optional] [default to false]
**AniOverride** | Pointer to **string** | Set a phone number as the ani_override value to override caller id number on outbound calls. | [optional] [default to ""]
**AniOverrideType** | Pointer to **string** | Specifies when we apply your ani_override setting. Only applies when ani_override is not blank. | [optional] [default to "always"]
**ChannelLimit** | Pointer to **int32** | When set, this will limit the total number of outbound calls to phone numbers associated with this connection. | [optional] 
**InstantRingbackEnabled** | Pointer to **bool** | When set, ringback will not wait for indication before sending ringback tone to calling party. | [optional] [default to true]
**GenerateRingbackTone** | Pointer to **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**Localization** | Pointer to **string** | A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to &#x60;US&#x60; then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default. | [optional] 
**T38ReinviteSource** | Pointer to **string** | This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite.&lt;br/&gt;&lt;br/&gt; By default, Telnyx will send the re-invite. If set to &#x60;customer&#x60;, the caller is expected to send the t.38 reinvite. | [optional] [default to "customer"]
**TechPrefix** | Pointer to **string** | Numerical chars only, exactly 4 characters. | [optional] [default to ""]
**IpAuthenticationMethod** | Pointer to **string** |  | [optional] [default to "tech-prefixp-charge-info"]
**IpAuthenticationToken** | Pointer to **string** |  | [optional] 
**OutboundVoiceProfileId** | Pointer to **string** | Identifies the associated outbound voice profile. | [optional] 

## Methods

### NewOutboundIp

`func NewOutboundIp() *OutboundIp`

NewOutboundIp instantiates a new OutboundIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundIpWithDefaults

`func NewOutboundIpWithDefaults() *OutboundIp`

NewOutboundIpWithDefaults instantiates a new OutboundIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallParkingEnabled

`func (o *OutboundIp) GetCallParkingEnabled() bool`

GetCallParkingEnabled returns the CallParkingEnabled field if non-nil, zero value otherwise.

### GetCallParkingEnabledOk

`func (o *OutboundIp) GetCallParkingEnabledOk() (*bool, bool)`

GetCallParkingEnabledOk returns a tuple with the CallParkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallParkingEnabled

`func (o *OutboundIp) SetCallParkingEnabled(v bool)`

SetCallParkingEnabled sets CallParkingEnabled field to given value.

### HasCallParkingEnabled

`func (o *OutboundIp) HasCallParkingEnabled() bool`

HasCallParkingEnabled returns a boolean if a field has been set.

### SetCallParkingEnabledNil

`func (o *OutboundIp) SetCallParkingEnabledNil(b bool)`

 SetCallParkingEnabledNil sets the value for CallParkingEnabled to be an explicit nil

### UnsetCallParkingEnabled
`func (o *OutboundIp) UnsetCallParkingEnabled()`

UnsetCallParkingEnabled ensures that no value is present for CallParkingEnabled, not even an explicit nil
### GetAniOverride

`func (o *OutboundIp) GetAniOverride() string`

GetAniOverride returns the AniOverride field if non-nil, zero value otherwise.

### GetAniOverrideOk

`func (o *OutboundIp) GetAniOverrideOk() (*string, bool)`

GetAniOverrideOk returns a tuple with the AniOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniOverride

`func (o *OutboundIp) SetAniOverride(v string)`

SetAniOverride sets AniOverride field to given value.

### HasAniOverride

`func (o *OutboundIp) HasAniOverride() bool`

HasAniOverride returns a boolean if a field has been set.

### GetAniOverrideType

`func (o *OutboundIp) GetAniOverrideType() string`

GetAniOverrideType returns the AniOverrideType field if non-nil, zero value otherwise.

### GetAniOverrideTypeOk

`func (o *OutboundIp) GetAniOverrideTypeOk() (*string, bool)`

GetAniOverrideTypeOk returns a tuple with the AniOverrideType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniOverrideType

`func (o *OutboundIp) SetAniOverrideType(v string)`

SetAniOverrideType sets AniOverrideType field to given value.

### HasAniOverrideType

`func (o *OutboundIp) HasAniOverrideType() bool`

HasAniOverrideType returns a boolean if a field has been set.

### GetChannelLimit

`func (o *OutboundIp) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *OutboundIp) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *OutboundIp) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *OutboundIp) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetInstantRingbackEnabled

`func (o *OutboundIp) GetInstantRingbackEnabled() bool`

GetInstantRingbackEnabled returns the InstantRingbackEnabled field if non-nil, zero value otherwise.

### GetInstantRingbackEnabledOk

`func (o *OutboundIp) GetInstantRingbackEnabledOk() (*bool, bool)`

GetInstantRingbackEnabledOk returns a tuple with the InstantRingbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRingbackEnabled

`func (o *OutboundIp) SetInstantRingbackEnabled(v bool)`

SetInstantRingbackEnabled sets InstantRingbackEnabled field to given value.

### HasInstantRingbackEnabled

`func (o *OutboundIp) HasInstantRingbackEnabled() bool`

HasInstantRingbackEnabled returns a boolean if a field has been set.

### GetGenerateRingbackTone

`func (o *OutboundIp) GetGenerateRingbackTone() bool`

GetGenerateRingbackTone returns the GenerateRingbackTone field if non-nil, zero value otherwise.

### GetGenerateRingbackToneOk

`func (o *OutboundIp) GetGenerateRingbackToneOk() (*bool, bool)`

GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRingbackTone

`func (o *OutboundIp) SetGenerateRingbackTone(v bool)`

SetGenerateRingbackTone sets GenerateRingbackTone field to given value.

### HasGenerateRingbackTone

`func (o *OutboundIp) HasGenerateRingbackTone() bool`

HasGenerateRingbackTone returns a boolean if a field has been set.

### GetLocalization

`func (o *OutboundIp) GetLocalization() string`

GetLocalization returns the Localization field if non-nil, zero value otherwise.

### GetLocalizationOk

`func (o *OutboundIp) GetLocalizationOk() (*string, bool)`

GetLocalizationOk returns a tuple with the Localization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalization

`func (o *OutboundIp) SetLocalization(v string)`

SetLocalization sets Localization field to given value.

### HasLocalization

`func (o *OutboundIp) HasLocalization() bool`

HasLocalization returns a boolean if a field has been set.

### GetT38ReinviteSource

`func (o *OutboundIp) GetT38ReinviteSource() string`

GetT38ReinviteSource returns the T38ReinviteSource field if non-nil, zero value otherwise.

### GetT38ReinviteSourceOk

`func (o *OutboundIp) GetT38ReinviteSourceOk() (*string, bool)`

GetT38ReinviteSourceOk returns a tuple with the T38ReinviteSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38ReinviteSource

`func (o *OutboundIp) SetT38ReinviteSource(v string)`

SetT38ReinviteSource sets T38ReinviteSource field to given value.

### HasT38ReinviteSource

`func (o *OutboundIp) HasT38ReinviteSource() bool`

HasT38ReinviteSource returns a boolean if a field has been set.

### GetTechPrefix

`func (o *OutboundIp) GetTechPrefix() string`

GetTechPrefix returns the TechPrefix field if non-nil, zero value otherwise.

### GetTechPrefixOk

`func (o *OutboundIp) GetTechPrefixOk() (*string, bool)`

GetTechPrefixOk returns a tuple with the TechPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechPrefix

`func (o *OutboundIp) SetTechPrefix(v string)`

SetTechPrefix sets TechPrefix field to given value.

### HasTechPrefix

`func (o *OutboundIp) HasTechPrefix() bool`

HasTechPrefix returns a boolean if a field has been set.

### GetIpAuthenticationMethod

`func (o *OutboundIp) GetIpAuthenticationMethod() string`

GetIpAuthenticationMethod returns the IpAuthenticationMethod field if non-nil, zero value otherwise.

### GetIpAuthenticationMethodOk

`func (o *OutboundIp) GetIpAuthenticationMethodOk() (*string, bool)`

GetIpAuthenticationMethodOk returns a tuple with the IpAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAuthenticationMethod

`func (o *OutboundIp) SetIpAuthenticationMethod(v string)`

SetIpAuthenticationMethod sets IpAuthenticationMethod field to given value.

### HasIpAuthenticationMethod

`func (o *OutboundIp) HasIpAuthenticationMethod() bool`

HasIpAuthenticationMethod returns a boolean if a field has been set.

### GetIpAuthenticationToken

`func (o *OutboundIp) GetIpAuthenticationToken() string`

GetIpAuthenticationToken returns the IpAuthenticationToken field if non-nil, zero value otherwise.

### GetIpAuthenticationTokenOk

`func (o *OutboundIp) GetIpAuthenticationTokenOk() (*string, bool)`

GetIpAuthenticationTokenOk returns a tuple with the IpAuthenticationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAuthenticationToken

`func (o *OutboundIp) SetIpAuthenticationToken(v string)`

SetIpAuthenticationToken sets IpAuthenticationToken field to given value.

### HasIpAuthenticationToken

`func (o *OutboundIp) HasIpAuthenticationToken() bool`

HasIpAuthenticationToken returns a boolean if a field has been set.

### GetOutboundVoiceProfileId

`func (o *OutboundIp) GetOutboundVoiceProfileId() string`

GetOutboundVoiceProfileId returns the OutboundVoiceProfileId field if non-nil, zero value otherwise.

### GetOutboundVoiceProfileIdOk

`func (o *OutboundIp) GetOutboundVoiceProfileIdOk() (*string, bool)`

GetOutboundVoiceProfileIdOk returns a tuple with the OutboundVoiceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundVoiceProfileId

`func (o *OutboundIp) SetOutboundVoiceProfileId(v string)`

SetOutboundVoiceProfileId sets OutboundVoiceProfileId field to given value.

### HasOutboundVoiceProfileId

`func (o *OutboundIp) HasOutboundVoiceProfileId() bool`

HasOutboundVoiceProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


