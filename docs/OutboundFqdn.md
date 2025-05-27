# OutboundFqdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AniOverride** | Pointer to **string** | Set a phone number as the ani_override value to override caller id number on outbound calls. | [optional] 
**AniOverrideType** | Pointer to **string** | Specifies when we should apply your ani_override setting. Only applies when ani_override is not blank. | [optional] [default to "always"]
**CallParkingEnabled** | Pointer to **NullableBool** | Forces all SIP calls originated on this connection to be \\\&quot;parked\\\&quot; instead of \\\&quot;bridged\\\&quot; to the destination specified on the URI. Parked calls will return ringback to the caller and will await for a Call Control command to define which action will be taken next. | [optional] [default to false]
**ChannelLimit** | Pointer to **int32** | When set, this will limit the total number of inbound calls to phone numbers associated with this connection. | [optional] 
**GenerateRingbackTone** | Pointer to **bool** | Generate ringback tone through 183 session progress message with early media. | [optional] [default to false]
**InstantRingbackEnabled** | Pointer to **bool** | When set, ringback will not wait for indication before sending ringback tone to calling party. | [optional] [default to false]
**IpAuthenticationMethod** | Pointer to **string** |  | [optional] [default to "ip-authentication"]
**IpAuthenticationToken** | Pointer to **string** |  | [optional] 
**Localization** | Pointer to **string** | A 2-character country code specifying the country whose national dialing rules should be used. For example, if set to &#x60;US&#x60; then any US number can be dialed without preprending +1 to the number. When left blank, Telnyx will try US and GB dialing rules, in that order, by default.\&quot;, | [optional] [default to "US"]
**OutboundVoiceProfileId** | Pointer to **string** | Identifies the associated outbound voice profile. | [optional] 
**T38ReinviteSource** | Pointer to **string** | This setting only affects connections with Fax-type Outbound Voice Profiles. The setting dictates whether or not Telnyx sends a t.38 reinvite. By default, Telnyx will send the re-invite. If set to &#x60;customer&#x60;, the caller is expected to send the t.38 reinvite. | [optional] [default to "customer"]
**TechPrefix** | Pointer to **string** | Numerical chars only, exactly 4 characters. | [optional] 
**EncryptedMedia** | Pointer to [**NullableEncryptedMedia**](EncryptedMedia.md) |  | [optional] 
**Timeout1xxSecs** | Pointer to **int32** | Time(sec) before aborting if connection is not made. | [optional] [default to 3]
**Timeout2xxSecs** | Pointer to **int32** | Time(sec) before aborting if call is unanswered (min: 1, max: 600). | [optional] [default to 90]

## Methods

### NewOutboundFqdn

`func NewOutboundFqdn() *OutboundFqdn`

NewOutboundFqdn instantiates a new OutboundFqdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundFqdnWithDefaults

`func NewOutboundFqdnWithDefaults() *OutboundFqdn`

NewOutboundFqdnWithDefaults instantiates a new OutboundFqdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAniOverride

`func (o *OutboundFqdn) GetAniOverride() string`

GetAniOverride returns the AniOverride field if non-nil, zero value otherwise.

### GetAniOverrideOk

`func (o *OutboundFqdn) GetAniOverrideOk() (*string, bool)`

GetAniOverrideOk returns a tuple with the AniOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniOverride

`func (o *OutboundFqdn) SetAniOverride(v string)`

SetAniOverride sets AniOverride field to given value.

### HasAniOverride

`func (o *OutboundFqdn) HasAniOverride() bool`

HasAniOverride returns a boolean if a field has been set.

### GetAniOverrideType

`func (o *OutboundFqdn) GetAniOverrideType() string`

GetAniOverrideType returns the AniOverrideType field if non-nil, zero value otherwise.

### GetAniOverrideTypeOk

`func (o *OutboundFqdn) GetAniOverrideTypeOk() (*string, bool)`

GetAniOverrideTypeOk returns a tuple with the AniOverrideType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAniOverrideType

`func (o *OutboundFqdn) SetAniOverrideType(v string)`

SetAniOverrideType sets AniOverrideType field to given value.

### HasAniOverrideType

`func (o *OutboundFqdn) HasAniOverrideType() bool`

HasAniOverrideType returns a boolean if a field has been set.

### GetCallParkingEnabled

`func (o *OutboundFqdn) GetCallParkingEnabled() bool`

GetCallParkingEnabled returns the CallParkingEnabled field if non-nil, zero value otherwise.

### GetCallParkingEnabledOk

`func (o *OutboundFqdn) GetCallParkingEnabledOk() (*bool, bool)`

GetCallParkingEnabledOk returns a tuple with the CallParkingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallParkingEnabled

`func (o *OutboundFqdn) SetCallParkingEnabled(v bool)`

SetCallParkingEnabled sets CallParkingEnabled field to given value.

### HasCallParkingEnabled

`func (o *OutboundFqdn) HasCallParkingEnabled() bool`

HasCallParkingEnabled returns a boolean if a field has been set.

### SetCallParkingEnabledNil

`func (o *OutboundFqdn) SetCallParkingEnabledNil(b bool)`

 SetCallParkingEnabledNil sets the value for CallParkingEnabled to be an explicit nil

### UnsetCallParkingEnabled
`func (o *OutboundFqdn) UnsetCallParkingEnabled()`

UnsetCallParkingEnabled ensures that no value is present for CallParkingEnabled, not even an explicit nil
### GetChannelLimit

`func (o *OutboundFqdn) GetChannelLimit() int32`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *OutboundFqdn) GetChannelLimitOk() (*int32, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *OutboundFqdn) SetChannelLimit(v int32)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *OutboundFqdn) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetGenerateRingbackTone

`func (o *OutboundFqdn) GetGenerateRingbackTone() bool`

GetGenerateRingbackTone returns the GenerateRingbackTone field if non-nil, zero value otherwise.

### GetGenerateRingbackToneOk

`func (o *OutboundFqdn) GetGenerateRingbackToneOk() (*bool, bool)`

GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRingbackTone

`func (o *OutboundFqdn) SetGenerateRingbackTone(v bool)`

SetGenerateRingbackTone sets GenerateRingbackTone field to given value.

### HasGenerateRingbackTone

`func (o *OutboundFqdn) HasGenerateRingbackTone() bool`

HasGenerateRingbackTone returns a boolean if a field has been set.

### GetInstantRingbackEnabled

`func (o *OutboundFqdn) GetInstantRingbackEnabled() bool`

GetInstantRingbackEnabled returns the InstantRingbackEnabled field if non-nil, zero value otherwise.

### GetInstantRingbackEnabledOk

`func (o *OutboundFqdn) GetInstantRingbackEnabledOk() (*bool, bool)`

GetInstantRingbackEnabledOk returns a tuple with the InstantRingbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantRingbackEnabled

`func (o *OutboundFqdn) SetInstantRingbackEnabled(v bool)`

SetInstantRingbackEnabled sets InstantRingbackEnabled field to given value.

### HasInstantRingbackEnabled

`func (o *OutboundFqdn) HasInstantRingbackEnabled() bool`

HasInstantRingbackEnabled returns a boolean if a field has been set.

### GetIpAuthenticationMethod

`func (o *OutboundFqdn) GetIpAuthenticationMethod() string`

GetIpAuthenticationMethod returns the IpAuthenticationMethod field if non-nil, zero value otherwise.

### GetIpAuthenticationMethodOk

`func (o *OutboundFqdn) GetIpAuthenticationMethodOk() (*string, bool)`

GetIpAuthenticationMethodOk returns a tuple with the IpAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAuthenticationMethod

`func (o *OutboundFqdn) SetIpAuthenticationMethod(v string)`

SetIpAuthenticationMethod sets IpAuthenticationMethod field to given value.

### HasIpAuthenticationMethod

`func (o *OutboundFqdn) HasIpAuthenticationMethod() bool`

HasIpAuthenticationMethod returns a boolean if a field has been set.

### GetIpAuthenticationToken

`func (o *OutboundFqdn) GetIpAuthenticationToken() string`

GetIpAuthenticationToken returns the IpAuthenticationToken field if non-nil, zero value otherwise.

### GetIpAuthenticationTokenOk

`func (o *OutboundFqdn) GetIpAuthenticationTokenOk() (*string, bool)`

GetIpAuthenticationTokenOk returns a tuple with the IpAuthenticationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAuthenticationToken

`func (o *OutboundFqdn) SetIpAuthenticationToken(v string)`

SetIpAuthenticationToken sets IpAuthenticationToken field to given value.

### HasIpAuthenticationToken

`func (o *OutboundFqdn) HasIpAuthenticationToken() bool`

HasIpAuthenticationToken returns a boolean if a field has been set.

### GetLocalization

`func (o *OutboundFqdn) GetLocalization() string`

GetLocalization returns the Localization field if non-nil, zero value otherwise.

### GetLocalizationOk

`func (o *OutboundFqdn) GetLocalizationOk() (*string, bool)`

GetLocalizationOk returns a tuple with the Localization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalization

`func (o *OutboundFqdn) SetLocalization(v string)`

SetLocalization sets Localization field to given value.

### HasLocalization

`func (o *OutboundFqdn) HasLocalization() bool`

HasLocalization returns a boolean if a field has been set.

### GetOutboundVoiceProfileId

`func (o *OutboundFqdn) GetOutboundVoiceProfileId() string`

GetOutboundVoiceProfileId returns the OutboundVoiceProfileId field if non-nil, zero value otherwise.

### GetOutboundVoiceProfileIdOk

`func (o *OutboundFqdn) GetOutboundVoiceProfileIdOk() (*string, bool)`

GetOutboundVoiceProfileIdOk returns a tuple with the OutboundVoiceProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundVoiceProfileId

`func (o *OutboundFqdn) SetOutboundVoiceProfileId(v string)`

SetOutboundVoiceProfileId sets OutboundVoiceProfileId field to given value.

### HasOutboundVoiceProfileId

`func (o *OutboundFqdn) HasOutboundVoiceProfileId() bool`

HasOutboundVoiceProfileId returns a boolean if a field has been set.

### GetT38ReinviteSource

`func (o *OutboundFqdn) GetT38ReinviteSource() string`

GetT38ReinviteSource returns the T38ReinviteSource field if non-nil, zero value otherwise.

### GetT38ReinviteSourceOk

`func (o *OutboundFqdn) GetT38ReinviteSourceOk() (*string, bool)`

GetT38ReinviteSourceOk returns a tuple with the T38ReinviteSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38ReinviteSource

`func (o *OutboundFqdn) SetT38ReinviteSource(v string)`

SetT38ReinviteSource sets T38ReinviteSource field to given value.

### HasT38ReinviteSource

`func (o *OutboundFqdn) HasT38ReinviteSource() bool`

HasT38ReinviteSource returns a boolean if a field has been set.

### GetTechPrefix

`func (o *OutboundFqdn) GetTechPrefix() string`

GetTechPrefix returns the TechPrefix field if non-nil, zero value otherwise.

### GetTechPrefixOk

`func (o *OutboundFqdn) GetTechPrefixOk() (*string, bool)`

GetTechPrefixOk returns a tuple with the TechPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechPrefix

`func (o *OutboundFqdn) SetTechPrefix(v string)`

SetTechPrefix sets TechPrefix field to given value.

### HasTechPrefix

`func (o *OutboundFqdn) HasTechPrefix() bool`

HasTechPrefix returns a boolean if a field has been set.

### GetEncryptedMedia

`func (o *OutboundFqdn) GetEncryptedMedia() EncryptedMedia`

GetEncryptedMedia returns the EncryptedMedia field if non-nil, zero value otherwise.

### GetEncryptedMediaOk

`func (o *OutboundFqdn) GetEncryptedMediaOk() (*EncryptedMedia, bool)`

GetEncryptedMediaOk returns a tuple with the EncryptedMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedMedia

`func (o *OutboundFqdn) SetEncryptedMedia(v EncryptedMedia)`

SetEncryptedMedia sets EncryptedMedia field to given value.

### HasEncryptedMedia

`func (o *OutboundFqdn) HasEncryptedMedia() bool`

HasEncryptedMedia returns a boolean if a field has been set.

### SetEncryptedMediaNil

`func (o *OutboundFqdn) SetEncryptedMediaNil(b bool)`

 SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil

### UnsetEncryptedMedia
`func (o *OutboundFqdn) UnsetEncryptedMedia()`

UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
### GetTimeout1xxSecs

`func (o *OutboundFqdn) GetTimeout1xxSecs() int32`

GetTimeout1xxSecs returns the Timeout1xxSecs field if non-nil, zero value otherwise.

### GetTimeout1xxSecsOk

`func (o *OutboundFqdn) GetTimeout1xxSecsOk() (*int32, bool)`

GetTimeout1xxSecsOk returns a tuple with the Timeout1xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout1xxSecs

`func (o *OutboundFqdn) SetTimeout1xxSecs(v int32)`

SetTimeout1xxSecs sets Timeout1xxSecs field to given value.

### HasTimeout1xxSecs

`func (o *OutboundFqdn) HasTimeout1xxSecs() bool`

HasTimeout1xxSecs returns a boolean if a field has been set.

### GetTimeout2xxSecs

`func (o *OutboundFqdn) GetTimeout2xxSecs() int32`

GetTimeout2xxSecs returns the Timeout2xxSecs field if non-nil, zero value otherwise.

### GetTimeout2xxSecsOk

`func (o *OutboundFqdn) GetTimeout2xxSecsOk() (*int32, bool)`

GetTimeout2xxSecsOk returns a tuple with the Timeout2xxSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout2xxSecs

`func (o *OutboundFqdn) SetTimeout2xxSecs(v int32)`

SetTimeout2xxSecs sets Timeout2xxSecs field to given value.

### HasTimeout2xxSecs

`func (o *OutboundFqdn) HasTimeout2xxSecs() bool`

HasTimeout2xxSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


